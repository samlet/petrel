package alfin

import (
	rice "github.com/GeertJohan/go.rice"
	"log"
	"path/filepath"
	"strings"
	"testing"
)

func TestPyGen(t *testing.T) {
	ent := "ExampleItem"
	// $ python service_meta.py entity_abi Example
	out, err := PyGen("service_meta.py", "entity_abi", ent)
	if err != nil {
		panic(err)
	}
	println("PyGen OUT => ", out)
}

func TestMainEntConf(t *testing.T) {
	configFile := "conf/maint_common.yml"
	maintConf := ReadMaintConf(configFile)
	Pretty(maintConf)
}

func TestTemplatePath(t *testing.T) {
	path:=filepath.Join("templates", "service_impl.tmpl")
	dir, file:=filepath.Split(path)
	box:=strings.TrimRight(dir, "/")
	println(box, file)
	// find a rice.Box
	templateString, err := rice.MustFindBox(box).String(file)
	if err != nil {
		log.Fatal(err)
	}
	println(templateString)
}