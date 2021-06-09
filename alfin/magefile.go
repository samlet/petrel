// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/samlet/petrel/alfin"
	"os"
	"strings"
)

//var Default = Build
var Default = Check

var Aliases = map[string]interface{} {
	"gen": GenRelations,
}

var GOPATH = os.Getenv("GOPATH")

func gopath(path string) string {
	return GOPATH + path
}

func getPackages() ([]string, error) {
	out, err := sh.Output("go", "list", "./...")
	if err != nil {
		return nil, err
	}

	return strings.Split(out, "\n"), nil
}

func getPackageFiles(name string) ([]string, error) {
	out, err := sh.Output("go", "list", "-f", "{{range .GoFiles}}{{$.Dir}}/{{.}} {{end}}", name)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.Trim(out, " "), " "), nil
}

func Check() {
	//mg.SerialDeps(Gofmt, Govet)
	println(".. check")
}

// Build binary
func Build() {
	//mg.SerialDeps(Vendor, Check)
	mg.SerialDeps(Check)
	sh.RunV("go", "build")
}

// Generate entity relations, to run:
// $ mage genrelations example
func GenRelations(what string){
	tmpl:="templates/relation_desc.tmpl"
	entMaps:=map[string][]string{
		"example": []string{"Example", "ExampleItem"},
	}
	ents,ok := entMaps[what]
	if ok {
		mani, err := alfin.NewMetaManipulate(ents)
		if err != nil {
			panic(err)
		}

		for _, ent := range ents {
			e := mani.MustEntity(ent)
			//e:=mani.MustEntity("Example")
			err = alfin.GenModelEntityWithMeta(tmpl, e, os.Stdout)
			if err != nil {
				panic(err)
			}
		}
	}else{
		fmt.Printf("Cannot find entity config for %s", what)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteSchemas(pkg string) {
	//pkg:="workload"
	alfin.WriteSchemas(pkg)
}

func CreateMod(modName string){
	alfin.CreateMod(modName)
}

/**
$ mage check
$ mage build
$ mage -l
$ mage gen example
	or $ mage genrelations example
$ mage WriteSchemas workload
*/
