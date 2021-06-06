package alfin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestUnmashallModelEntity(te *testing.T) {
	content, err := ioutil.ReadFile("./assets/example.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var modelEntity ModelEntity
	err=json.Unmarshal(content, &modelEntity)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	println(modelEntity.Name)
}

// $ go test -run TestModelEntityTemplate
func TestModelEntityTemplate(te *testing.T) {
	tmpl:="templates/ent_schema.tmpl"
	err:=GenModelEntity(tmpl, "./assets/example.json", os.Stdout)
	if err != nil {
		panic(err)
	}
	println("// ------------------------------")
	err=GenModelEntity(tmpl, "./assets/exampleitem.json", os.Stdout)
	if err != nil {
		panic(err)
	}
	println("ok")
}

