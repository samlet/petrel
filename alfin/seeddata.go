package alfin

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/iancoleman/strcase"
	"github.com/samlet/petrel/alfin/common"
	"log"
	"strconv"
	"strings"
)

// ClientKey is the key for lookup
type ClientKey int

const (
	// CadenceClientKey for retrieving cadence client from context
	CadenceClientKey ClientKey = iota
	WorkloadStoreKey
	SeedsKey
)

type SeedGenPhrase int

const (
	CreatePhrase SeedGenPhrase = iota
	UpdatePhrase
)

type SeedElements struct {
	elements map[string]interface{}
}

func NewSeedElements() *SeedElements {
	return &SeedElements{elements: make(map[string]interface{})}
}

func (t *SeedElements) Put(key string, value interface{}) interface{} {
	t.elements[key] = value
	return value
}

func (t *SeedElements) MustGet(key string) interface{} {
	v, ok := t.elements[key]
	if !ok {
		panic(fmt.Errorf("Cannot get value from key %s", key))
	}
	return v
}

func (t *SeedElements) Size() int {
	return len(t.elements)
}

type SeedProcessor struct {
	MetaMan       *MetaManipulate
	Buffer        *strings.Builder
	WriteRelation bool
	Trace         bool
	Phrase        SeedGenPhrase
}

func NewSeedProcessor(MetaMan *MetaManipulate, writeRelation bool) *SeedProcessor {
	return &SeedProcessor{
		MetaMan:       MetaMan,
		Buffer:        new(strings.Builder),
		WriteRelation: writeRelation,
		Phrase:        CreatePhrase,
	}
}

func (t SeedProcessor) Write(format string, a ...interface{}) {
	t.Buffer.WriteString(fmt.Sprintf(format, a...))
}
func (t SeedProcessor) WriteLine(format string, a ...interface{}) {
	t.Buffer.WriteString(fmt.Sprintf(format, a...) + "\n")
}

func (t SeedProcessor) WriteFileHeader(pkg string) {
	t.WriteLine(`package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/%s/ent"
	"github.com/samlet/petrel/alfin/common"
	"log"
)`, pkg)
}

func (t SeedProcessor) WriteFunctionHeader(ent *ModelEntity) {
	switch t.Phrase {
	case CreatePhrase:
		t.WriteLine(`func Create%s(ctx context.Context) error {
	log.Println("creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.%s
`, ent.Name, ent.Name)

	case UpdatePhrase:
		t.WriteLine(`func Update%s(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.%s
`, ent.Name, ent.Name)
	}
}

func (t SeedProcessor) WriteFunctionFooter() {
	t.WriteLine(`return nil
}`)
}

func (t SeedProcessor) Printf(format string, a ...interface{}) {
	if t.Trace {
		fmt.Printf(format, a...)
	}
}

func (t SeedProcessor) Println(a ...interface{}) {
	if t.Trace {
		fmt.Println(a...)
	}
}
func (t SeedProcessor) Print(a ...interface{}) {
	if t.Trace {
		fmt.Print(a...)
	}
}

func (t *SeedProcessor) TraceOn() {
	t.Trace = true
}

func (t SeedProcessor) ProcessElements(doc *etree.Document, elements []*etree.Element, ent *ModelEntity) {
	for _, element := range elements {
		pk := ent.Pks[0]
		t.Println("Tag:", element.Tag, element.SelectAttrValue(pk, ""))
		pkString := GetStringRef(ent, element)
		//t.WriteLine(`var err error`)
		switch t.Phrase {
		case CreatePhrase:
			t.WriteLine("c, err = client.%s.Create().SetStringRef(\"%s\").",
				element.Tag,
				pkString)
		case UpdatePhrase:
			t.WriteLine(`c=cache.Get("%s").(*ent.%s)
			c, err=c.Update().`,
				pkString,
				element.Tag,
			)
		default:
			log.Fatal("Cannot process phrase", t.Phrase)
		}

		for _, fld := range ent.NormalFields() {
			att := element.SelectAttr(fld.Name)
			if att != nil {
				t.Printf("\t[f] %s = %s\n", fld.Name, att.Value)
				code := fmt.Sprintf("\t  %s(%s).\n", fld.SetterName(),
					fld.QuoteValue(att.Value))
				t.Buffer.WriteString(code)
				t.Print(code)
			} else {
				t.Println("\tabsent", fld.Name)
			}
		}
		for _, edge := range ent.Edges() {
			att := element.SelectAttr(edge.FieldName())
			if att != nil {
				t.Printf("\t[r] %s(%s) = %s\n", edge.Name, edge.Keys(), att.Value)
				if edge.Type == "many" {
					code := fmt.Sprintf("\t  Add%s(%s).\n", strcase.ToCamel(edge.Name),
						strings.ToLower(att.Value)+"_"+strings.ToLower(edge.RelEntityName))

					t.Print(code)
					t.relatedQuery(edge, element, doc)
				} else if edge.Type == "one" {
					if edge.BackrefType == "many" {
						t.Println("\t  * backref-type = many")
					}
					code := fmt.Sprintf("\t  Set%s(%s).\n", strcase.ToCamel(edge.Name),
						strings.ToLower(att.Value)+"_"+strings.ToLower(edge.RelEntityName))

					t.Print(code)
					t.relatedQuery(edge, element, doc)

				} else {
					t.Println("\t  skip one-nofk")
				}
			} else {
				t.Println("\tabsent", edge.Name)
			}
		}

		t.WriteLine("\t  Save(ctx)")
		t.WriteLine(`if err != nil {
	log.Printf("fail to create %s: %%v", err)
	return err
}
cache.Put("%s", c)
`, pkString, pkString)
	}
}

func GetStringRef(ent *ModelEntity, element *etree.Element) string {
	var pkValues []string
	for _, pk := range ent.Pks {
		val := element.SelectAttrValue(pk, "")
		fld := ent.GetField(pk)
		if fld.IsDateTime() {
			intval, err := common.ToSecs(val)
			if err != nil {
				log.Fatalf(" fail: %v", err)
			}
			val = strconv.FormatInt(intval, 10)
		}
		pkValues = append(pkValues, val)
	}
	pkString := strings.ToLower(strings.Join(pkValues, "__") + "__" + element.Tag)
	return pkString
}

func (t SeedProcessor) relatedQuery(edge *ModelRelation, element *etree.Element, doc *etree.Document) {
	var keys []string
	var values []string
	for _, km := range edge.Keymaps {
		keys = append(keys, km.RelFieldName)
		values = append(values, element.SelectAttrValue(km.FieldName, ""))
	}
	t.queryElements(doc, edge, keys, values)
}

//var (
//	SerialNumbers = []string{"❶", "❷", "❸", "❹", "❺", "❻", "❼", "❽", "❾", "❿"}
//)

func (t SeedProcessor) queryElements(doc *etree.Document, edge *ModelRelation, keys []string, values []string) {
	entName := edge.RelEntityName
	// "//OrderRole[@partyId='DemoSupplier'][@orderId='DEMO10091']"
	var sb strings.Builder
	for i := 0; i < len(keys); i++ {
		sb.WriteString(fmt.Sprintf("[@%s='%s']", keys[i], values[i]))
	}
	entModel := t.MetaMan.MustEntity(entName)
	for n, element := range doc.FindElements("//" + entName + sb.String()) {
		stringRef := GetStringRef(entModel, element)
		t.Printf("\t\t%d. %s %s (%s)\n", n,
			element.Tag,
			element.SelectAttrValue(keys[0], ""),
			stringRef,
		)
		for _, att := range element.Attr {
			t.Println("\t\t\t-", att.Key, att.Value)
		}

		// codegen
		if t.WriteRelation {
			var procName string
			if edge.Type == "many" {
				//procName = entModel.AdderName()
				procName = edge.AdderName()
			} else {
				procName = edge.SetterName()
			}
			code := fmt.Sprintf("\t  %s(%s).\n",
				procName,
				quoteRef(entName, stringRef))
			t.Buffer.WriteString(code)
		}
	}
}

func quoteRef(entName string, stringRef string) interface{} {
	return fmt.Sprintf("cache.Get(\"%s\").(*ent.%s)", stringRef, entName)
}

func (t SeedProcessor) GetElements(doc *etree.Document, entName string, keys []string, values []string) []*etree.Element {
	// "//OrderRole[@partyId='DemoSupplier'][@orderId='DEMO10091']"
	var sb strings.Builder
	for i := 0; i < len(keys); i++ {
		sb.WriteString(fmt.Sprintf("[@%s='%s']", keys[i], values[i]))
	}
	return doc.FindElements("//" + entName + sb.String())
}

func (t SeedProcessor) GetElementsByArgs(doc *etree.Document, entName string, args ...string) []*etree.Element {
	// "//OrderRole[@partyId='DemoSupplier'][@orderId='DEMO10091']"
	var sb strings.Builder
	for k, v := range Mapify(args) {
		sb.WriteString(fmt.Sprintf("[@%s='%s']", k, v))
	}
	return doc.FindElements("//" + entName + sb.String())
}
