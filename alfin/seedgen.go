package alfin

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/magefile/mage/sh"
	"log"
	"path/filepath"
	"strings"
)

type SeedGen struct{
	mani *MetaManipulate
	doc *etree.Document
	pkg string
	entName string
	dontFmt bool
}

func seedDoc(pkg string) *etree.Document {
	xmlSeedFile:=filepath.Join("assets", pkg, "seeds.xml")
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xmlSeedFile); err != nil {
		log.Fatal(err)
	}
	return doc
}

func NewSeedGen(pkg string, entName string) *SeedGen {
	doc:=seedDoc(pkg)
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)

	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		log.Fatal(err)
	}
	return &SeedGen{mani,doc,pkg,entName, false}
}

func (t SeedGen) ModDir() string{
	return filepath.Join("modules", t.pkg)
}

func (t SeedGen) CreatorsDir() string{
	return filepath.Join(t.ModDir(), "seedcreators")
}

func (t SeedGen) Generate(onlyCreator bool) error {
	err:=EnsureDir(t.CreatorsDir())
	if err != nil {
		log.Fatalf("create dir %s fail: %v", t.CreatorsDir(), err)
	}
	err = t.seedCreator()
	if err != nil {
		return err
	}
	if !onlyCreator {
		err = t.seedUpdater()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t SeedGen) seedCreator() error {
	model := t.mani.MustEntity(t.entName)
	seedProc := NewSeedProcessor(t.mani, false)
	//seedProc.TraceOn()
	//println("trace", seedProc.Trace)
	//elements:=seedProc.GetElementsByArgs(doc, entName, "partyId", "DemoCustomer3")
	elements := t.doc.FindElements("//" + t.entName)

	seedProc.WriteFileHeader(t.pkg)
	seedProc.WriteFunctionHeader(model)
	seedProc.ProcessElements(t.doc, elements, model)
	seedProc.WriteFunctionFooter()

	//println("-------------------")
	outputCode := seedProc.Buffer.String()
	//println(outputCode)
	targetFile := filepath.Join(t.CreatorsDir(), strings.ToLower(t.entName)+"_creator.go")
	err := WriteContent(targetFile, outputCode)
	if err != nil {
		log.Fatalf("write code fail: %v", err)
	}

	if !t.dontFmt {
		sh.RunV("go", "fmt", targetFile)
	}
	return nil
}

func (t SeedGen) seedUpdater() error {
	model := t.mani.MustEntity(t.entName)
	seedProc := NewSeedProcessor(t.mani, true)
	seedProc.Phrase=UpdatePhrase

	elements := t.doc.FindElements("//" + t.entName)
	seedProc.WriteFileHeader(t.pkg)
	seedProc.WriteFunctionHeader(model)
	seedProc.ProcessElements(t.doc, elements, model)
	seedProc.WriteFunctionFooter()

	outputCode := seedProc.Buffer.String()
	targetFile := filepath.Join(t.CreatorsDir(), strings.ToLower(t.entName)+"_updater.go")
	err := WriteContent(targetFile, outputCode)
	if err != nil {
		log.Fatalf("write code fail: %v", err)
	}

	if !t.dontFmt {
		sh.RunV("go", "fmt", targetFile)
	}

	return nil
}


var (
	Header=`package seedcreators

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/%s/ent"
	"log"
)

func LoadSeeds(execUpdaters bool) {
	client, err := ent.Open("sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	ctx = ent.NewContext(ctx, client)
	ctx = cachecomp.NewContext(ctx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

`

	Footer=`}`
)

func GeneratePackage(pkg string) error{
	doc:=seedDoc(pkg)
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)

	mani, err := NewManipulateWithPackage(pkg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _,entName := range mani.EntityNames() {
		seedgen := SeedGen{mani, doc, pkg, entName, true}
		err=seedgen.Generate(false)
		if err != nil {
			log.Fatalf("gen %s fail: %v", entName, err)
			return err
		}
	}

	buf:=strings.Builder{}
	f:=fmt.Sprintf
	buf.WriteString(f(Header, pkg))

	for _,entName := range mani.EntityNames() {
		buf.WriteString(f(`if err := Create%s(ctx); err != nil {
		log.Fatal(err)
	}
	`, entName))
	}

	buf.WriteString("if execUpdaters {")
	for _,entName := range mani.EntityNames() {
		buf.WriteString(f(`if err := Update%s(ctx); err != nil {
		log.Fatal(err)
	}
	`, entName))
	}
	buf.WriteString("}")

	buf.WriteString(Footer)
	creatorDir:=filepath.Join("modules", pkg, "seedcreators")
	targetFile:=filepath.Join(creatorDir, "seedloader.go")
	err=WriteContent(targetFile, buf.String())
	if err != nil {
		log.Fatalf("write seed-loader fail: %v", err)
		return err
	}

	// exec: go fmt modules/workeffort/seedcreators/*.go
	//sh.RunV("go", "fmt", "\""+creatorDir+"/*.go"+"\"")
	return nil
}