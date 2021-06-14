package alfin

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xlab/treeprint"
	"os"
	"path/filepath"
	"testing"
)

func TestMetaManipulate(t *testing.T) {
	ents := []string{"Example", "ExampleItem"}
	mani, err := NewMetaManipulate(ents)
	if err != nil {
		panic(err)
	}

	e := mani.MustEntity("ExampleItem")
	//jsonr, err:=json.MarshalIndent(e, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s", jsonr)
	for _, r := range e.Relations {
		println(r.Backref)
	}
}

func TestPackageMeta(t *testing.T) {
	sample := `{
  "name": "workload",
  "package": "com.bluecc.workload",
  "entities": {
    "WorkloadFeature": "workloadfeature.json",
    "WorkloadFeatureAppl": "workloadfeatureappl.json",
    "WorkloadFeatureApplType": "workloadfeatureappltype.json",
    "Workload": "workload.json",
    "WorkloadItem": "workloaditem.json",
    "WorkloadStatus": "workloadstatus.json",
    "WorkloadType": "workloadtype.json"
  }
}`
	var pkgMeta PackageMeta
	err := json.Unmarshal([]byte(sample), &pkgMeta)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", pkgMeta)
}

func TestLoadPackage(t *testing.T) {
	pkg := "workload"
	assetDir := filepath.Join(AssetRoot, pkg)
	metaFile := filepath.Join(assetDir, "meta.json")
	var pkgMeta PackageMeta
	err := ReadJsonFile(metaFile, &pkgMeta)
	if err != nil {
		panic(err)
	}
	println(pkgMeta.Package)
	files := []string{}
	for _, entFile := range pkgMeta.Entities {
		files = append(files, filepath.Join(assetDir, entFile))
	}

	mani, err := NewMetaManipulate(files)
	if err != nil {
		panic(err)
	}

	e := mani.MustEntity("Workload")
	println(e.Name)

	e = mani.MustEntity("WorkloadType")
	println(e.Name)
	for _, r := range e.Relations {
		println(e.Name, r.Name, r.FieldName(), r.SelfRelation)
	}
}

func TestGenSchemas(t *testing.T) {
	pkg := "workload"
	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		panic(err)
	}
	err = mani.GenSchemaWithFiles(os.Stdout)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
}

func TestGenSchemas_purchaseorder(t *testing.T) {
	pkg := "purchaseorder"
	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		panic(err)
	}
	err = mani.GenSchemaWithFiles(os.Stdout)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
}

func TestWriteSchemas(t *testing.T) {
	pkg := "workload"
	path := filepath.Join("modules", pkg, "ent", "schema", "modent.go")
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		panic(err)
	}
	err = mani.GenSchemaWithFiles(f)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}

	f.Sync()

	println("write to ", path)
}

func TestWriteHelpers(t *testing.T) {
	pkg := "workeffort"
	path := filepath.Join("modules", pkg, "helper", pkg+"helper.go")
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		panic(err)
	}
	mani.GenHelpers(f)

	f.Sync()

	println("write to ", path)
}

func TestCreateTree(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	pkg := "workeffort"
	entName := "WorkEffort"
	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		panic(err)
	}
	//for _, ent := range mani.Entities().Entities {
	//	DisplayEntInfo(ent)
	//}
	tree := treeprint.New()
	ent := mani.MustEntity(entName)
	DisplayEntInfo(ent, tree, true)
	fmt.Println(tree.String())
}
