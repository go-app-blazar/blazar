package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	directory, err := os.Getwd()
	if err != nil {
		slog.Error("Failed to get CWD.", "err", err)
		os.Exit(1)
	}
	fmt.Printf("CWD: %s\n", directory)

	t, err := template.ParseFiles("mix-html-template.gotmpl")
	if err != nil {
		slog.Error("Failed to parse template", "err", err)
		os.Exit(1)
	}

	files, err := filepath.Glob("*.go")
	if err != nil {
		slog.Error("Failed to glob files", "err", err)
		os.Exit(1)
	}

	for _, file := range files {
		if strings.HasSuffix(file, "_generated.go") || strings.HasSuffix(file, "-generated.go") {
			continue
		}

		slog.Info("Go file.", "file", file)

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			slog.Error("Failed to parse file.", "file", file, "err", err)
			os.Exit(1)
		}

		packageName := f.Name.Name
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
	}
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
