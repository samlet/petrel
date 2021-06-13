package alfin

import (
	"errors"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"strings"
)

const (
	AssetRoot="assets"
)
type MetaManipulate struct{
	EntsMap map[string]*ModelEntity
}

func NewMetaManipulateWithMap(entsMap map[string]*ModelEntity) (*MetaManipulate, error) {
	// get keys
	keys := make([]string, 0, len(entsMap))
	for k := range entsMap {
		keys = append(keys, k)
	}

	log.Printf("%s\n", keys)

	for _, e := range entsMap {
		e.EntitiesInPkg=&keys
		// manage entity fields
		var lastNumber int
		for i,fld := range e.NormalFields() {
			fld.FieldNumber=i+2  // start with 2, the id field is number 1
			lastNumber=fld.FieldNumber+1
		}

		// manage entity relations
		for _, rel := range e.Relations {
			rel.FieldNumber=lastNumber
			lastNumber=lastNumber+1

			relEntName := rel.RelEntityName
			if rel.HashBackref() {
				if relEntName==e.Name{
					log.Printf("entity %s has a self-relation\n", e.Name)
					rel.SelfRelation=true
				}else {
					log.Printf(".. check %s.%s(%s)\n", e.Name, rel.Name, relEntName)
					relEnt, ok := entsMap[relEntName]
					if ok {
						refName, refType, err := relEnt.GetBackRefNameAndType(e.Name, rel)
						if err != nil {
							panic(err)
						}
						rel.Backref = refName
						rel.BackrefType=refType
						log.Printf("*** %s.%s backref: %s.%s(%s)\n",
							e.Name, rel.Name,
							relEntName, refName, refType)
					} else {
						log.Printf("\tEntity %s is not exists in meta-collection\n", relEntName)
					}
				}
			}
		}


	}

	return &MetaManipulate{EntsMap: entsMap}, nil
}

// NewMetaManipulate create MetaManipulate
// entity names as `ents := []string{"Example", "ExampleItem"}` or files
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

	return NewMetaManipulateWithMap(entsMap)
}

func (t *MetaManipulate) MustEntity(entName string) *ModelEntity {
	ent, ok:=t.EntsMap[entName]
	if !ok {
		panic(errors.New("Cannot find entity "+entName))
	}
	return ent
}


type PackageMeta struct{
	Name string `json:"name"`
	Package string `json:"package"`
	Entities map[string]string `json:"entities"`
}

const (
	SchemaHeader=`package schema
import (
    "entgo.io/ent"
    // "entgo.io/ent/schema/index"
	"entgo.io/ent/schema"
    // "entgo.io/ent/schema/mixin"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
	"entgo.io/contrib/entproto"
    "time"
)
`
)

func GenSchemas(pkg string, writer io.Writer) error{
	assetDir:=filepath.Join(AssetRoot, pkg)
	metaFile:=filepath.Join(assetDir, "meta.json")
	var pkgMeta PackageMeta
	err:=ReadJsonFile(metaFile, &pkgMeta)
	if err != nil {
		return(err)
	}
	println(pkgMeta.Package)
	files:=[]string{}
	ents:=[]string{}
	for entName,entFile:=range pkgMeta.Entities{
		ents=append(ents, entName)
		files=append(files, filepath.Join(assetDir, entFile))
	}

	return GenSchemaWithFiles(files, ents, writer)
}

func GenSchemaWithFiles(files []string, ents []string, writer io.Writer) error {
	mani, err := NewMetaManipulate(files)
	if err != nil {
		return (err)
	}

	_, err = writer.Write([]byte(SchemaHeader))
	if err != nil {
		return fmt.Errorf("write header: %w", err)
	}
	tmpls := []string{"ent_schema.tmpl", "relation_desc.tmpl"}
	for _, ent := range ents {
		e := mani.MustEntity(ent)
		if !e.IsView {
			for _, tmpl := range tmpls {
				err = GenModelEntityWithMeta(filepath.Join("templates", tmpl),
					e, writer)
				if err != nil {
					return (err)
				}
			}
		}
	}

	return nil
}

