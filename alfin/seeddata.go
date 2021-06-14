package alfin

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/iancoleman/strcase"
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
	MetaMan *MetaManipulate
	Buffer        *strings.Builder
	WriteRelation bool
}

func NewSeedProcessor(MetaMan *MetaManipulate, writeRelation bool) *SeedProcessor {
	return &SeedProcessor{
		MetaMan: MetaMan,
		Buffer: new(strings.Builder),
		WriteRelation: writeRelation}
}

func (t SeedProcessor) Write(format string, a ...interface{}) {
	t.Buffer.WriteString(fmt.Sprintf(format, a...))
}
func (t SeedProcessor) WriteLine(format string, a ...interface{}) {
	t.Buffer.WriteString(fmt.Sprintf(format, a...) + "\n")
}

func (t SeedProcessor) ProcessElements(doc *etree.Document, elements []*etree.Element, ent *ModelEntity) {
	for _, element := range elements {
		pk := ent.Pks[0]
		fmt.Println("Tag:", element.Tag, element.SelectAttrValue(pk, ""))
		pkString := GetStringRef(ent, element)
		t.WriteLine("%s.Create().SetStringRef(\"%s\").", element.Tag,
			pkString)
		for _, fld := range ent.NormalFields() {
			att := element.SelectAttr(fld.Name)
			if att != nil {
				fmt.Printf("\t[f] %s = %s\n", fld.Name, att.Value)
				code := fmt.Sprintf("\t  Set%s(%s).\n", strcase.ToCamel(fld.Name),
					fld.QuoteValue(att.Value))
				t.Buffer.WriteString(code)
				print(code)
			} else {
				fmt.Println("\tabsent", fld.Name)
			}
		}
		for _, edge := range ent.Edges() {
			att := element.SelectAttr(edge.FieldName())
			if att != nil {
				fmt.Printf("\t[r] %s(%s) = %s\n", edge.Name, edge.Keys(), att.Value)
				if edge.Type == "many" {
					code := fmt.Sprintf("\t  Add%s(%s).\n", strcase.ToCamel(edge.Name),
						strings.ToLower(att.Value)+"_"+strings.ToLower(edge.RelEntityName))

					print(code)
					t.relatedQuery(edge, element, doc)
				} else if edge.Type == "one" {
					if edge.BackrefType == "many" {
						fmt.Println("\t  * backref-type = many")
					}
					code := fmt.Sprintf("\t  Set%s(%s).\n", strcase.ToCamel(edge.Name),
						strings.ToLower(att.Value)+"_"+strings.ToLower(edge.RelEntityName))

					print(code)
					t.relatedQuery(edge, element, doc)

				} else {
					fmt.Println("\t  skip one-nofk")
				}
			} else {
				fmt.Println("\tabsent", edge.Name)
			}
		}

		t.WriteLine("\t  Save(ctx)")
	}
}

func GetStringRef(ent *ModelEntity, element *etree.Element) string {
	var pkValues []string
	for _, pk := range ent.Pks {
		pkValues = append(pkValues, element.SelectAttrValue(pk, ""))
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

var (
	SerialNumbers = []string{"❶", "❷", "❸", "❹", "❺", "❻", "❼", "❽", "❾", "❿"}
)

func (t SeedProcessor) queryElements(doc *etree.Document, edge *ModelRelation, keys []string, values []string) {
	entName:=edge.RelEntityName
	// "//OrderRole[@partyId='DemoSupplier'][@orderId='DEMO10091']"
	var sb strings.Builder
	for i := 0; i < len(keys); i++ {
		sb.WriteString(fmt.Sprintf("[@%s='%s']", keys[i], values[i]))
	}
	entModel:=t.MetaMan.MustEntity(entName)
	for n, element := range doc.FindElements("//" + entName + sb.String()) {
		stringRef:=GetStringRef(entModel, element)
		fmt.Printf("\t\t%s %s %s (%s)\n", SerialNumbers[n],
			element.Tag,
			element.SelectAttrValue(keys[0], ""),
			stringRef,
		)
		for _, att := range element.Attr {
			fmt.Println("\t\t\t-", att.Key, att.Value)
		}

		// codegen
		if t.WriteRelation {
			var prefix string
			if edge.Type=="many"{
				prefix="Add"
			}else{
				prefix="Set"
			}
			code := fmt.Sprintf("\t  %s%s(%s).\n",
				prefix,
				strcase.ToCamel(edge.Name),
				stringRef)
			t.Buffer.WriteString(code)
		}
	}
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
