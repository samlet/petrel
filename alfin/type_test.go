package alfin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func TestRelationDesc(t *testing.T) {
	tmpl:="templates/relation_desc.tmpl"

	ents := []string{"Example", "ExampleItem", "ExampleType"}
	mani, err:=NewMetaManipulate(ents)
	if err != nil {
		panic(err)
	}

	for _,ent := range ents {
		e := mani.MustEntity(ent)
		//e:=mani.MustEntity("Example")
		err = GenModelEntityWithMeta(tmpl, e, os.Stdout)
		if err != nil {
			panic(err)
		}
	}

}

func TestGenEntitySchemas(t *testing.T) {
	tmpls:=[]string{"ent_schema.tmpl", "relation_desc.tmpl"}
	//ents := []string{"Example", "ExampleItem", "ExampleType"}
	ents := []string{"Example"}
	mani, err:=NewMetaManipulate(ents)
	if err != nil {
		panic(err)
	}

	for _,ent := range ents {
		e := mani.MustEntity(ent)
		for _,tmpl := range tmpls {
			err = GenModelEntityWithMeta(filepath.Join("templates", tmpl),
				e, os.Stdout)
			if err != nil {
				panic(err)
			}
		}
	}

}
