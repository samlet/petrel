package alfin

import (
	"testing"
)

func TestMetaManipulate(t *testing.T) {
	ents := []string{"Example", "ExampleItem"}
	mani, err:=NewMetaManipulate(ents)
	if err != nil {
		panic(err)
	}

	e:=mani.MustEntity("ExampleItem")
	//jsonr, err:=json.MarshalIndent(e, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s", jsonr)
	for _,r := range e.Relations {
		println(r.Backref)
	}
}
