package alfin

import (
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
	return &SeedGen{mani,doc,pkg,entName}
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
	return sh.RunV("go", "fmt", targetFile)
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

	return sh.RunV("go", "fmt", targetFile)
}

