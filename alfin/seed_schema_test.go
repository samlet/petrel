package alfin

import (
	"encoding/xml"
	"github.com/beevik/etree"
	"github.com/samlet/petrel/alfin/seedtypes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSeedSchema(t *testing.T) {
	tmpl := "templates/seed_schema.tmpl"

	//ents := []string{"Example", "ExampleItem", "ExampleType"}
	ents := []string{"Example"}
	mani, err := NewMetaManipulate(ents)
	if err != nil {
		panic(err)
	}

	for _, ent := range ents {
		e := mani.MustEntity(ent)
		//e:=mani.MustEntity("Example")
		err = GenModelEntityWithMeta(tmpl, e, os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}

func TestSeedSchemas(t *testing.T) {
	tmpl := "templates/seed_schema.tmpl"
	header := `package seedtypes
import "github.com/samlet/petrel/services"
`

	mani, err := NewManipulateWithPackage("workeffort")
	if err != nil {
		panic(err)
	}

	for _, ent := range mani.Entities().Entities {
		filePath := strings.ToLower(ent.Name) + ".go"
		f, err := os.Create(filepath.Join("seedtypes", filePath))
		if err != nil {
			log.Fatalf("create file fail: %v", err)
		}
		_, err = f.WriteString(header)
		if err != nil {
			log.Fatalf("write header fail: %v", err)
		}
		err = GenModelEntityWithMeta(tmpl, ent, f)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}

type SeedTypes struct {
	UserLogin []seedtypes.UserLogin
}

func TestSeedUnmarshal(t *testing.T) {
	seeds := `<entity-engine-xml>
    <Party partyId="DemoCustomer3" partyTypeId="PERSON" statusId="PARTY_ENABLED"/>
    <Person partyId="DemoCustomer3" firstName="Billing" lastName="Customer 3"/>
    <UserLogin userLoginId="DemoCustomer3"  partyId="DemoCustomer3"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CUSTOMER"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CLIENT_BILLING"/>
</entity-engine-xml>`
	var v SeedTypes
	err := xml.Unmarshal([]byte(seeds), &v)
	if err != nil {
		log.Fatalf(" fail: %v", err)
	}
	for _, r := range v.UserLogin {
		pretty(r)
	}

}

func TestSetters(tt *testing.T) {
	xmlSeed := `<entity-engine-xml>
    <Party partyId="DemoCustomer3" partyTypeId="PERSON" statusId="PARTY_ENABLED"/>
    <Person partyId="DemoCustomer3" firstName="Billing" lastName="Customer 3"/>
    <UserLogin userLoginId="DemoCustomer3"  partyId="DemoCustomer3"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CUSTOMER"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CLIENT_BILLING"/>
</entity-engine-xml>`

	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlSeed); err != nil {
		panic(err)
	}
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)

	// ...
	entName := "Party"
	//entName:="Person"
	mani, err := NewManipulateWithPackage("workeffort")
	if err != nil {
		panic(err)
	}
	ent := mani.MustEntity(entName)
	elements:=doc.FindElements("//" + entName)

	seedProc:=NewSeedProcessor(mani,true)
	seedProc.ProcessElements(doc, elements, ent)
}

func TestSettersByArgs(tt *testing.T) {
	xmlSeed := `<entity-engine-xml>
    <Party partyId="DemoCustomer3" partyTypeId="PERSON" statusId="PARTY_ENABLED"/>
    <Person partyId="DemoCustomer3" firstName="Billing" lastName="Customer 3"/>
    <UserLogin userLoginId="DemoCustomer3"  partyId="DemoCustomer3"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CUSTOMER"/>
    <PartyRole partyId="DemoCustomer3" roleTypeId="CLIENT_BILLING"/>
</entity-engine-xml>`

	doc := etree.NewDocument()
	if err := doc.ReadFromString(xmlSeed); err != nil {
		panic(err)
	}
	root := doc.SelectElement("entity-engine-xml")
	log.Println("ROOT element:", root.Tag)

	// ...
	entName := "Party"
	//entName:="Person"
	mani, err := NewManipulateWithPackage("workeffort")
	if err != nil {
		panic(err)
	}
	ent := mani.MustEntity(entName)
	seedProc:=NewSeedProcessor(mani,true)
	//seedProc:=NewSeedProcessor(false)
	elements:=seedProc.GetElementsByArgs(doc, entName, "partyId", "DemoCustomer3")

	seedProc.ProcessElements(doc, elements, ent)
	println("-------------------")
	println(seedProc.Buffer.String())
}
