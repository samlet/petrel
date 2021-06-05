// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"os"
	"strings"

	"github.com/magefile/mage/sh"
)

var Default = Build

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

/**
$ mage check
$ mage build
*/
