package alfin

import (
	"fmt"
	"strings"
	"testing"
)

func TestMetaManipulate(t *testing.T) {
	f := fmt.Sprintf
	entsMap := make(map[string]*ModelEntity)
	ents := []string{"Example", "ExampleItem"}
	for _, e := range ents {
		ex, err := LoadModelEntity(f("./assets/%s.json", strings.ToLower(e)))
		if err != nil {
			panic(err)
		}
		entsMap[ex.Name] = ex
	}

	// get keys
	keys := make([]string, 0, len(entsMap))
	for k := range entsMap {
		keys = append(keys, k)
	}
	fmt.Printf("%s\n", keys)

	for _, e := range entsMap {
		for _, rel := range e.Relations {
			relEntName := rel.RelEntityName
			if rel.HashBackref() {
				fmt.Printf(".. check %s.%s(%s)\n", e.Name, rel.Name, relEntName)
				relEnt, ok := entsMap[relEntName]
				if ok {
					refName, err := relEnt.GetBackRefName(e.Name, rel)
					if err != nil {
						panic(err)
					}
					rel.Backref=refName
					fmt.Printf("*** %s.%s backref: %s.%s\n",
						e.Name, rel.Name,
						relEntName, refName)
				} else {
					fmt.Printf("\tEntity %s is not exists in meta-collection\n", relEntName)
				}
			}
		}
	}
}
