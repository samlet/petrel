package alfin

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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

func TestPackageMeta(t *testing.T) {
	sample:=`{
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
	err:=json.Unmarshal([]byte(sample), &pkgMeta)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", pkgMeta)
}

func TestLoadPackage(t *testing.T) {
	pkg:="workload"
	assetDir:=filepath.Join(AssetRoot, pkg)
	metaFile:=filepath.Join(assetDir, "meta.json")
	var pkgMeta PackageMeta
	err:=ReadJsonFile(metaFile, &pkgMeta)
	if err != nil {
		panic(err)
	}
	println(pkgMeta.Package)
	files:=[]string{}
	for _,entFile:=range pkgMeta.Entities{
		files=append(files, filepath.Join(assetDir, entFile))
	}

	mani, err:=NewMetaManipulate(files)
	if err != nil {
		panic(err)
	}

	e:=mani.MustEntity("Workload")
	println(e.Name)

	e=mani.MustEntity("WorkloadType")
	println(e.Name)
	for _,r:=range e.Relations{
		println(e.Name, r.Name, r.FieldName(), r.SelfRelation)
	}
}

func TestGenSchemas(t *testing.T) {
	pkg:="workload"
	err:=GenSchemas(pkg, os.Stdout)
	if err != nil {
		panic(err)
	}
}

func TestWriteSchemas(t *testing.T) {
	pkg:="workload"
	path:=filepath.Join("modules", pkg, "ent", "schema", "workload.go")
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	err=GenSchemas(pkg, f)
	if err != nil {
		panic(err)
	}

	f.Sync()

	println("write to ", path)
}
