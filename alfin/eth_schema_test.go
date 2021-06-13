package alfin

import (
	"os"
	"testing"
)

func TestEthSchemaGen(t *testing.T) {
	tmpl:="templates/eth_schema.tmpl"

	//ents := []string{"Example", "ExampleItem", "ExampleType"}
	ents := []string{"Example"}
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


