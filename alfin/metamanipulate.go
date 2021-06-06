package alfin

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type MetaManipulate struct{
	EntsMap map[string]*ModelEntity
}

// NewMetaManipulate create MetaManipulate
// `ents := []string{"Example", "ExampleItem"}`
func NewMetaManipulate(ents []string) (*MetaManipulate, error) {
	f := fmt.Sprintf
	entsMap := make(map[string]*ModelEntity)
	for _, e := range ents {
		ex, err := LoadModelEntity(f("./assets/%s.json", strings.ToLower(e)))
		if err != nil {
			return nil, err
		}
		entsMap[ex.Name] = ex
	}

	// get keys
	keys := make([]string, 0, len(entsMap))
	for k := range entsMap {
		keys = append(keys, k)
	}

	log.Printf("%s\n", keys)

	for _, e := range entsMap {
		for _, rel := range e.Relations {
			relEntName := rel.RelEntityName
			if rel.HashBackref() {
				log.Printf(".. check %s.%s(%s)\n", e.Name, rel.Name, relEntName)
				relEnt, ok := entsMap[relEntName]
				if ok {
					refName, err := relEnt.GetBackRefName(e.Name, rel)
					if err != nil {
						panic(err)
					}
					rel.Backref=refName
					log.Printf("*** %s.%s backref: %s.%s\n",
						e.Name, rel.Name,
						relEntName, refName)
				} else {
					log.Printf("\tEntity %s is not exists in meta-collection\n", relEntName)
				}
			}
		}
	}

	return &MetaManipulate{EntsMap: entsMap}, nil
}

func (t *MetaManipulate) MustEntity(entName string) *ModelEntity {
	ent, ok:=t.EntsMap[entName]
	if !ok {
		panic(errors.New("Cannot find entity "+entName))
	}
	return ent
}
