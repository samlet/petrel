package alfin

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	AssetRoot="assets"
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
		metaFile:=e
		if !strings.HasSuffix(e, ".json"){
			metaFile=f("./assets/%s.json", strings.ToLower(e))
		}
		ex, err := LoadModelEntity(metaFile)
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
				if relEntName==e.Name{
					log.Printf("entity %s has a self-relation\n", e.Name)
					rel.SelfRelation=true
				}else {
					log.Printf(".. check %s.%s(%s)\n", e.Name, rel.Name, relEntName)
					relEnt, ok := entsMap[relEntName]
					if ok {
						refName, err := relEnt.GetBackRefName(e.Name, rel)
						if err != nil {
							panic(err)
						}
						rel.Backref = refName
						log.Printf("*** %s.%s backref: %s.%s\n",
							e.Name, rel.Name,
							relEntName, refName)
					} else {
						log.Printf("\tEntity %s is not exists in meta-collection\n", relEntName)
					}
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

