package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
)

func GetInterfaces(root string, files []string) {
	for _, filename := range files {
		if filename == "-" {
			continue
		}
		filename = filepath.Join(root, filename)
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			log.Fatal("Error parsing file: " + filename)
		}

		// traverse all tokens
		ast.Inspect(node, func(n ast.Node) bool {
			switch t := n.(type) {
			// find variable declarations
			case *ast.TypeSpec:
				// which are public
				if t.Name.IsExported() {
					switch t.Type.(type) {
					// and are interfaces
					case *ast.InterfaceType:
						fmt.Println(t.Name.Name)
					}
				}
			}
			return true
		})
	}
}
