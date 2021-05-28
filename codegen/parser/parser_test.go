package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
	"testing"
)

func TestParseConstants(t *testing.T) {
	gofile := "constants.go"
	fset := token.NewFileSet()
	//解析go文件
	f, err := parser.ParseFile(fset, gofile, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	typeNames := "Status"
	types := strings.Split(typeNames, ",")
	typesMap := make(map[string][]string, len(types))
	for _, v := range types {
		typesMap[strings.TrimSpace(v)] = []string{}
	}

	//遍历每个树节点
	typ := ""
	ast.Inspect(f, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		// 只需要const
		if !ok || decl.Tok != token.CONST {
			return true
		}
		for _, spec := range decl.Specs {
			vspec := spec.(*ast.ValueSpec)
			if vspec.Type == nil && len(vspec.Values) > 0 {
				// 排除 v = 1 这种结构
				typ = ""
				continue
			}
			//如果Type不为空，则确认typ
			if vspec.Type != nil {
				ident, ok := vspec.Type.(*ast.Ident)
				if !ok {
					continue
				}
				typ = ident.Name
			}
			//typ是否是需处理的类型
			consts, ok := typesMap[typ]
			if !ok {
				continue
			}
			//将所有const变量名保存
			for _, n := range vspec.Names {
				consts = append(consts, n.Name)
			}
			typesMap[typ] = consts
		}
		return true
	})

	jsonstr, _ := json.MarshalIndent(typesMap, "", "  ")
	fmt.Printf("%s\n", jsonstr)
}

func TestParseSource(t *testing.T) {
	gofile := "interfaces.go"
	fset := token.NewFileSet()
	//解析go文件
	f, err := parser.ParseFile(fset, gofile, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}

func TestParseInterfaces(t *testing.T) {
	gofile := "interfaces.go"
	fset := token.NewFileSet()
	//解析go文件
	f, err := parser.ParseFile(fset, gofile, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Print the AST.
	//ast.Print(fset, f)

	// traverse all tokens
	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		// find variable declarations
		case *ast.TypeSpec:
			// which are public
			if t.Name.IsExported() {
				switch t.Type.(type) {
				// and are interfaces
				case *ast.InterfaceType:
					fmt.Println(t.Name.Name)
					ispec := t.Type.(*ast.InterfaceType)
					fieldsLen := len(ispec.Methods.List)
					fmt.Printf(".. methods: %d\n", fieldsLen)
					for _, m := range ispec.Methods.List {
						//ast.Print(fset, m)
						ast.Print(fset, m.Names[0].Name)
					}

					HandleInterface(fset, ispec)
				}
			}
		}
		return true
	})

}

func HandleInterface(fset *token.FileSet, iface *ast.InterfaceType) {
	if iface.Methods != nil || iface.Methods.List != nil {
		for _, v := range iface.Methods.List {
			print(v.Names[0].Name, " => ")
			ft := v.Type.(*ast.FuncType)
			for _, v := range ft.Params.List {
				//fmt.Printf("%s: %s, ", v.Names[0].Name, v.Type)
				fmt.Printf("%s: ", v.Names[0].Name)
				//ast.Print(fset, v.Type)
				switch n := v.Type.(type) {
				case *ast.MapType:
					ast.Print(fset, n)
				case *ast.Ident:
					print(n.Name, ", ")
				default:
					// ...
				}
			}
			println()
		}
	}
}
