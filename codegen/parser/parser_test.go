package main

import (
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

	fmt.Printf("%s\n", typesMap)
}
