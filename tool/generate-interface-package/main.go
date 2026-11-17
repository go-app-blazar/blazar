package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/go-app-blazar/blazar/mix"
)

type Args struct {
	SourceDirectory string `arg:"--source-directory,required"`
	TargetDirectory string `arg:"--target-directory,required"`
	PrivatePrefix   string `arg:"--private-prefix,required"`
	PublicPrefix    string `arg:"--public-prefix,required"`
}

func main() {
	{
		currentDirectory, err := os.Getwd()
		if err != nil {
			slog.Error("Failed to get CWD.", "err", err)
			os.Exit(1)
		}
		fmt.Printf("CWD: %s\n", currentDirectory)
	}

	var args Args
	arg.MustParse(&args)

	sourceDirectory := args.SourceDirectory
	sourceDirectory, err := filepath.Abs(sourceDirectory)
	if err != nil {
		slog.Error("Failed to get absolute path of source directory.", "directory", sourceDirectory, "err", err)
		os.Exit(1)
	}
	fmt.Printf("Source directory: %s\n", sourceDirectory)

	destinationDirectory := args.TargetDirectory
	destinationDirectory, err = filepath.Abs(destinationDirectory)
	if err != nil {
		slog.Error("Failed to get absolute path of destination directory.", "directory", destinationDirectory, "err", err)
		os.Exit(1)
	}
	fmt.Printf("Destination directory: %s\n", destinationDirectory)

	privatePrefix := args.PrivatePrefix
	publicPrefix := args.PublicPrefix

	err = os.Chdir(sourceDirectory)
	if err != nil {
		slog.Error("Failed to change to directory.", "directory", sourceDirectory, "err", err)
		os.Exit(1)
	}
	fmt.Printf("Changed to: %s\n", sourceDirectory)

	t, err := template.ParseFS(mix.FS, "interface.gotmpl")
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
	filenameToASTMap := map[string]*ast.File{}
	for _, filename := range files {
		slog.Info("Go file.", "file", filename)

		if strings.HasSuffix(filename, "_test.go") {
			continue
		}

		contents, err := os.ReadFile(filename)
		if err != nil {
			slog.Error("Failed to read file.", "filename", filename, "err", err)
			os.Exit(1)
		}

		// Prepend two blank lines to the source before parsing.
		// This gives us real line numbers (lines 1, 2, and 3) in the token.File before the "package" declaration (now on line 4).
		modifiedFileContents := "\n\n\n" + string(contents)

		f, err := parser.ParseFile(fset, filename, modifiedFileContents, parser.ParseComments)
		if err != nil {
			slog.Error("Failed to parse file.", "filename", filename, "err", err)
			os.Exit(1)
		}

		filenameToASTMap[filename] = f
	}

	var packageName string
	structFileNames := map[string]string{}
	structImports := map[string][]string{}
	structInterfaces := map[string][]string{}
	structMethods := map[string][]string{}

	for _, f := range filenameToASTMap {
		if packageName == "" {
			packageName = f.Name.Name
		}
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
			structInterfaces[typeName] = []string{"app.UI"}
			structFileNames[typeName] = strings.TrimSuffix(fset.Position(f.Package).Filename, ".go")

			// Iterate over the fields of the struct.
			for _, field := range structType.Fields.List {
				//slog.Info("  Field.", "field", fieldNames, "type", field.Type)
				if fieldExprToString(field.Type) == "mix.MixComposer" {
					structInterfaces[typeName] = append(structInterfaces[typeName], "app.Composer")
				}
				if fieldExprToString(field.Type) == "mix.MixRouterViewComponent" {
					structInterfaces[typeName] = append(structInterfaces[typeName], "router.RouterViewInterface")
					structImports[typeName] = append(structImports[typeName], "github.com/go-app-blazar/router")
				}
			}

			_ = structType

			return true
		})
	}

	for _, f := range filenameToASTMap {
		headerComment := &ast.CommentGroup{
			List: []*ast.Comment{
				{
					Slash: f.FileStart,
					Text:  "// This file was automatically generated as part of the Blazar build process via \"go generate\".\n// DO NOT EDIT THIS FILE.\n",
				},
			},
		}
		f.Comments = append([]*ast.CommentGroup{headerComment}, f.Comments...)

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
				slog.Info("No receiver type name.", "function", fn.Name.Name)

				if fn.Type.Results != nil && len(fn.Type.Results.List) == 1 {
					returnType := types.ExprString(fn.Type.Results.List[0].Type)
					slog.Info("Return type.", "returnType", returnType)

					if !strings.HasPrefix(returnType, "*"+privatePrefix) {
						return true
					}
					structName := strings.TrimPrefix(returnType, "*")
					if _, ok := structMethods[structName]; !ok {
						return true
					}

					if returnType == "*"+privatePrefix+fn.Name.Name {
						fn.Type.Results.List[0].Type = &ast.Ident{Name: publicPrefix + fn.Name.Name}
					}
				}

				return true
			}
			if len(fn.Recv.List) == 0 {
				return true
			}

			receiverTypeName := getReceiverTypeName(fn.Recv.List[0].Type)

			// If this isn't one of the structs that we care about, then skip it.
			if _, ok := structMethods[receiverTypeName]; !ok {
				return true
			}

			desiredInterfaceName := strings.Replace(receiverTypeName, privatePrefix, publicPrefix, 1)

			slog.Info("Receiver type name.", "receiverTypeName", receiverTypeName)
			slog.Info("Function type.", "functionType", types.ExprString(fn.Type))
			functionDeclaration := types.ExprString(fn.Type)
			functionDeclaration = strings.Replace(functionDeclaration, "func", fn.Name.Name, 1)
			if !strings.HasSuffix(functionDeclaration, "*"+receiverTypeName) {
				//return true
			}
			functionDeclaration = strings.Replace(functionDeclaration, "*"+receiverTypeName, "%%TYPE%%", 1)

			if fn.Type.Results != nil && len(fn.Type.Results.List) == 1 {
				returnType := types.ExprString(fn.Type.Results.List[0].Type)
				slog.Info("Return type.", "returnType", returnType)
				if returnType == "*"+receiverTypeName {
					fn.Type.Results.List[0].Type = &ast.Ident{Name: desiredInterfaceName}
				}
			}

			comment := strings.TrimSpace(fn.Doc.Text())
			if len(comment) > 0 {
				newComment := ""
				lines := strings.Split(comment, "\n")
				for _, line := range lines {
					line = strings.TrimSpace(line)

					newComment = newComment + "// " + line + "\n"
				}

				functionDeclaration = strings.TrimSpace(newComment) + "\n" + functionDeclaration
			}

			structMethods[receiverTypeName] = append(structMethods[receiverTypeName], functionDeclaration)

			return true
		})
	}

	for structName := range structMethods {
		slices.Sort(structMethods[structName])
	}

	err = os.Chdir(destinationDirectory)
	if err != nil {
		slog.Error("Failed to change to destination directory.", "directory", destinationDirectory, "err", err)
		os.Exit(1)
	}
	fmt.Printf("Changed to: %s\n", destinationDirectory)

	for filename, f := range filenameToASTMap {
		if !strings.HasSuffix(filename, "_generated.go") {
			filename = strings.TrimSuffix(filename, ".go") + "_generated.go"
		}

		var buffer bytes.Buffer
		err = format.Node(&buffer, fset, f)
		if err != nil {
			slog.Error("Failed to format file.", "err", err)
			os.Exit(1)
		}

		err = os.WriteFile(filename, buffer.Bytes(), 0644)
		if err != nil {
			slog.Error("Failed to write file.", "err", err)
			os.Exit(1)
		}
	}

	for structName, methods := range structMethods {
		slog.Info("Struct.", "struct", structName, "methods", len(methods))
		fileName := structFileNames[structName]
		slog.Info("File name.", "fileName", fileName)
		for _, method := range methods {
			slog.Info("   Method.", "method", method)
		}

		if !strings.HasPrefix(structName, privatePrefix) {
			continue
		}

		imports := []string{
			"github.com/maxence-charriere/go-app/v11/pkg/app", // JSValue returns "app.Value".
		}
		for _, include := range structImports[structName] {
			imports = append(imports, include)
		}
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

		interfaceName := strings.Replace(structName, privatePrefix, publicPrefix, 1)

		// TODO: FIGURE OUT THE PROPER "T any" SYNTAX FOR THE INTERFACE GENERIC.
		fullInterfaceName := strings.Replace(interfaceName, "[T]", "[T any]", 1)

		var methodStrings []string
		for _, method := range methods {
			method = strings.ReplaceAll(method, "%%TYPE%%", interfaceName)
			methodStrings = append(methodStrings, method)
		}
		slices.Sort(methodStrings)

		data := map[string]any{
			"Package":    packageName,
			"Imports":    imports,
			"Type":       fullInterfaceName,
			"Interfaces": structInterfaces[structName],
			"Methods":    methodStrings,
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
