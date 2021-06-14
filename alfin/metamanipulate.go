package alfin

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"strings"
)

const (
	AssetRoot = "assets"
)

type MetaManipulate struct {
	Name    string
	EntsMap map[string]*ModelEntity
}

func NewMetaManipulateWithMap(entsMap map[string]*ModelEntity) (*MetaManipulate, error) {
	// get keys
	keys := make([]string, 0, len(entsMap))
	for k := range entsMap {
		keys = append(keys, k)
	}

	log.Debugf("%s\n", keys)

	for _, e := range entsMap {
		e.EntitiesInPkg = &keys
		// manage entity fields
		var lastNumber int
		for i, fld := range e.NormalFields() {
			fld.FieldNumber = i + 2 // start with 2, the id field is number 1
			lastNumber = fld.FieldNumber + 1
		}

		// manage entity relations
		for _, rel := range e.Relations {
			rel.FieldNumber = lastNumber
			lastNumber = lastNumber + 1

			relEntName := rel.RelEntityName
			if rel.HashBackref() {
				if relEntName == e.Name {
					log.Debugf("entity %s has a self-relation\n", e.Name)
					rel.SelfRelation = true
				} else {
					log.Debugf(".. check %s.%s(%s)\n", e.Name, rel.Name, relEntName)
					relEnt, ok := entsMap[relEntName]
					if ok {
						refName, refType, err := relEnt.GetBackRefNameAndType(e.Name, rel)
						if err != nil {
							panic(err)
						}
						rel.Backref = refName
						rel.BackrefType = refType
						log.Debugf("*** %s.%s backref: %s.%s(%s)\n",
							e.Name, rel.Name,
							relEntName, refName, refType)
					} else {
						log.Debugf("\tEntity %s is not exists in meta-collection\n", relEntName)
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
		metaFile := e
		if !strings.HasSuffix(e, ".json") {
			metaFile = f("./assets/%s.json", strings.ToLower(e))
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
	ent, ok := t.EntsMap[entName]
	if !ok {
		log.Fatal(errors.New("Cannot find entity " + entName))
	}
	return ent
}

func (t MetaManipulate) EntityNames() []string {
	keys := make([]string, 0, len(t.EntsMap))
	for k := range t.EntsMap {
		keys = append(keys, k)
	}
	return keys
}

type ModelPackage struct {
	Name     string
	Entities []*ModelEntity
}

func (t MetaManipulate) Entities() ModelPackage {
	vals := make([]*ModelEntity, 0, len(t.EntsMap))
	for _, v := range t.EntsMap {
		vals = append(vals, v)
	}
	return ModelPackage{Name: t.Name, Entities: vals}
}

type PackageMeta struct {
	Name     string            `json:"name"`
	Package  string            `json:"package"`
	Entities map[string]string `json:"entities"`
}

const (
	SchemaHeader = `package schema
import (
    "entgo.io/ent"
    // "entgo.io/ent/schema/index"
	"entgo.io/ent/schema"
    "entgo.io/ent/schema/mixin"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
	"entgo.io/contrib/entproto"
	"github.com/samlet/petrel/alfin/schemamixins"
    "time"
)
`
)

func NewManipulateWithPackage(pkg string) (*MetaManipulate, error) {
	assetDir := filepath.Join(AssetRoot, pkg)
	metaFile := filepath.Join(assetDir, "meta.json")
	var pkgMeta PackageMeta
	err := ReadJsonFile(metaFile, &pkgMeta)
	if err != nil {
		return nil, err
	}
	log.Debugf("process pkg %s", pkgMeta.Package)
	files := []string{}
	//ents:=[]string{}
	for _, entFile := range pkgMeta.Entities {
		//ents=append(ents, entName)
		files = append(files, filepath.Join(assetDir, entFile))
	}

	mani, err := NewMetaManipulate(files)
	if err != nil {
		return nil, err
	}
	mani.Name = pkg
	//err = mani.genSchemaWithFiles(writer)
	//if err != nil {
	//	log.Fatalf(" fail: %v", err)
	//}

	return mani, nil
}

func (mani *MetaManipulate) GenSchemaWithFiles(writer io.Writer) error {
	_, err := writer.Write([]byte(SchemaHeader))
	if err != nil {
		return fmt.Errorf("write header: %w", err)
	}
	tmpls := []string{"ent_schema.tmpl", "relation_desc.tmpl"}
	for _, ent := range mani.EntityNames() {
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

func (mani *MetaManipulate) GenHelpers(writer io.Writer) {
	tmpl := "ent_ref.tmpl"
	models := mani.Entities()

	err := GenModelEntityWithMeta(filepath.Join("templates", tmpl),
		models, writer)
	if err != nil {
		log.Fatal(err)
	}
}

