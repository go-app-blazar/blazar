package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	directory, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get CWD.", "err", err)
		os.Exit(1)
	}
	fmt.Printf("CWD: %s\n", directory)

	t, err := template.ParseFiles("interface-template.gotmpl")
	if err != nil {
		slog.Error("Failed to parse template", "err", err)
		os.Exit(1)
	}

	files, err := filepath.Glob("*.go")
	if err != nil {
		slog.Error("Failed to glob files", "err", err)
		os.Exit(1)
	}

	fset := token.NewFileSet()
	var astFiles []*ast.File
	for _, file := range files {
		slog.Info("Go file.", "file", file)

		f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			slog.Error("Failed to parse file.", "file", file, "err", err)
			os.Exit(1)
		}

		astFiles = append(astFiles, f)
	}

	structFileNames := map[string]string{}
	structMethods := map[string][]string{}

	for _, f := range astFiles {
		ast.Inspect(f, func(n ast.Node) bool {
			// Look for type specifications (for example, type Name struct { ... }).
			typeSpec, isTypeSpec := n.(*ast.TypeSpec)
			if !isTypeSpec {
				return true
			}

			// Check if the type specification is a struct.
			structType, isStructType := typeSpec.Type.(*ast.StructType)
			if !isStructType {
				return true
			}

			slog.Info("Found struct.", "struct", typeSpec.Name.Name)

			isCompo := false

			// Iterate over the fields of the struct.
			for _, field := range structType.Fields.List {
				//slog.Info("  Field.", "field", fieldNames, "type", field.Type)
				if fieldExprToString(field.Type) == "app.Compo" {
					isCompo = true
				}
			}

			if !isCompo {
				return true
			}

			typeName := typeSpec.Name.Name
			if typeSpec.TypeParams != nil && len(typeSpec.TypeParams.List) > 0 {
				typeParts := []string{}
				for _, typeParam := range typeSpec.TypeParams.List {
					typeParts = append(typeParts, typeParam.Names[0].Name)
				}
				typeName += "[" + strings.Join(typeParts, ", ") + "]"
			}

			structMethods[typeName] = []string{}
			structFileNames[typeName] = strings.TrimSuffix(fset.Position(f.Package).Filename, ".go")

			_ = structType

			return true
		})
	}

	for _, f := range astFiles {
		ast.Inspect(f, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			slog.Info("Found function.", "function", fn.Name.Name)

			if len(fn.Name.Name) == 0 {
				return true
			}
			if fn.Name.Name[0] < 'A' || fn.Name.Name[0] > 'Z' {
				return true
			}

			if fn.Recv == nil {
				return true
			}
			if len(fn.Recv.List) == 0 {
				return true
			}

			receiverTypeName := getReceiverTypeName(fn.Recv.List[0].Type)
			if _, ok := structMethods[receiverTypeName]; !ok {
				return true
			}

			slog.Info("Receiver type name.", "receiverTypeName", receiverTypeName)
			slog.Info("Function type.", "functionType", types.ExprString(fn.Type))
			functionDeclaration := types.ExprString(fn.Type)
			functionDeclaration = strings.Replace(functionDeclaration, "func", fn.Name.Name, 1)
			if !strings.HasSuffix(functionDeclaration, "*"+receiverTypeName) {
				return true
			}
			functionDeclaration = strings.Replace(functionDeclaration, "*"+receiverTypeName, "%%TYPE%%", 1)

			structMethods[receiverTypeName] = append(structMethods[receiverTypeName], functionDeclaration)

			return true
		})
	}

	for structName := range structMethods {
		slices.Sort(structMethods[structName])
	}

	for structName, methods := range structMethods {
		slog.Info("Struct.", "struct", structName, "methods", len(methods))
		fileName := structFileNames[structName]
		slog.Info("File name.", "fileName", fileName)
		for _, method := range methods {
			slog.Info("   Method.", "method", method)
		}

		if !strings.HasPrefix(structName, "blazar") {
			continue
		}

		imports := []string{}
		// TODO: FIGURE OUT THE PACKAGES NEEDED BASED ON THE FUNCTION SIGNATURES AND INCLUDE THEM.
		for _, method := range methods {
			if strings.Contains(method, "app.") {
				imports = append(imports, "github.com/maxence-charriere/go-app/v11/pkg/app")
			}
			if strings.Contains(method, "htmlevent.") {
				imports = append(imports, "github.com/go-app-blazar/blazar/htmlevent")
			}
		}
		slices.Sort(imports)
		imports = slices.Compact(imports)

		interfaceName := strings.Replace(structName, "blazar", "Blazar", 1)

		// TODO: FIGURE OUT THE PROPER "T any" SYNTAX FOR THE INTERFACE GENERIC.
		fullInterfaceName := strings.Replace(interfaceName, "[T]", "[T any]", 1)

		var methodStrings []string
		for _, method := range methods {
			method = strings.ReplaceAll(method, "%%TYPE%%", interfaceName)
			methodStrings = append(methodStrings, method)
		}
		slices.Sort(methodStrings)

		packageName := "blazar"
		data := map[string]any{
			"Package": packageName,
			"Imports": imports,
			"Type":    fullInterfaceName,
			"Methods": methodStrings,
		}

		newFilename := "z-interface." + strings.TrimSuffix(fileName, ".go") + "_generated.go"
		slog.Info("New filename.", "newFilename", newFilename)

		newFile, err := os.Create(newFilename)
		if err != nil {
			slog.Error("Failed to create new file.", "err", err)
			os.Exit(1)
		}
		defer newFile.Close()

		err = t.Execute(newFile, data)
		if err != nil {
			slog.Error("Failed to execute template.", "err", err)
			os.Exit(1)
		}

	}

	_ = t

	/*
		packageName := fset.
		slog.Info("Package.", "package", packageName)
		ast.Inspect(f, func(n ast.Node) bool {
			// Look for type specifications (for example, type Name struct { ... }).
			typeSpec, isTypeSpec := n.(*ast.TypeSpec)
			if !isTypeSpec {
				return true
			}

			// Check if the type specification is a struct.
			structType, isStructType := typeSpec.Type.(*ast.StructType)
			if !isStructType {
				return true
			}

			slog.Info("Found struct.", "struct", typeSpec.Name.Name)

			isMixHTML := false

			// Iterate over the fields of the struct.
			for _, field := range structType.Fields.List {
				//slog.Info("  Field.", "field", fieldNames, "type", field.Type)
				if fieldExprToString(field.Type) == "MixHTML" {
					isMixHTML = true
				}
			}

			if !isMixHTML {
				return true
			}

			typeName := typeSpec.Name.Name
			if typeSpec.TypeParams != nil && len(typeSpec.TypeParams.List) > 0 {
				typeParts := []string{}
				for _, typeParam := range typeSpec.TypeParams.List {
					typeParts = append(typeParts, typeParam.Names[0].Name)
				}
				typeName += "[" + strings.Join(typeParts, ", ") + "]"
			}

			slog.Info("MixHTML found.", "struct", typeName)

			newFilename := "z-mix-html." + strings.TrimSuffix(file, ".go") + "_generated.go"
			slog.Info("New filename.", "newFilename", newFilename)

			newFile, err := os.Create(newFilename)
			if err != nil {
				slog.Error("Failed to create new file.", "err", err)
				os.Exit(1)
			}
			defer newFile.Close()

			data := map[string]any{
				"Package": packageName,
				"Type":    typeName,
			}

			err = t.Execute(newFile, data)
			if err != nil {
				slog.Error("Failed to execute template.", "err", err)
				os.Exit(1)
			}

			return true
		})
	*/
}

// Simple helper to convert field type expressions to string representation
func fieldExprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fieldExprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + fieldExprToString(t.Elt)
	default:
		return fmt.Sprintf("%T", expr)
	}
}

// getReceiverTypeName extracts the receiver type name, preserving generic type parameters.
// Examples:
//
//	(s MyStruct)       -> "MyStruct"
//	(s *MyStruct)      -> "MyStruct"
//	(s *Box[T])        -> "Box[T]"
//	(s Pair[K, V])     -> "Pair[K, V]"
func getReceiverTypeName(expr ast.Expr) string {
	// Unwrap pointer receiver if present
	if star, ok := expr.(*ast.StarExpr); ok {
		expr = star.X
	}

	// types.ExprString cleanly formats *ast.Ident, *ast.IndexExpr, and *ast.IndexListExpr
	return types.ExprString(expr)
}
