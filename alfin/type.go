package alfin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type (
	ModelEntity struct {
		Name          string           `json:"name"`
		Fields        []*ModelField    `json:"fields"`
		Relations     []*ModelRelation `json:"relations"`
		PksSize       int              `json:"pksSize"`
		Pks           []string         `json:"pks"`
		IsView        bool             `json:"isView"`
		EntitiesInPkg *[]string        `json:"-"`
	}

	ModelField struct {
		Name                string `json:"name"`
		Type                string `json:"type"`
		Col                 string `json:"col"`
		Pk                  bool   `json:"pk"`
		NotNull             bool   `json:"notNull"`
		AutoCreatedInternal bool   `json:"autoCreatedInternal"`
		FieldNumber         int    `json:"-"`
	}

	ModelRelation struct {
		Name          string        `json:"name"`
		Type          string        `json:"type"`
		RelEntityName string        `json:"relEntityName"`
		FkName        string        `json:"fkName"`
		Keymaps       []ModelKeymap `json:"keymaps"`
		Backref       string        `json:"-"`
		SelfRelation  bool          `json:"-"`
		BackrefType   string        `json:"-"`
		FieldNumber   int           `json:"-"`
	}

	ModelKeymap struct {
		FieldName    string `json:"fieldName"`
		RelFieldName string `json:"relFieldName"`
	}
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (t ModelEntity) IsUniqueIdField(fldName string) bool {
	return t.PksSize == 1 && contains(t.Pks, fldName)
}

func (t ModelEntity) UniquePk() string {
	if t.PksSize == 1 {
		return t.Pks[0]
	} else {
		return ""
	}
}

func (t ModelEntity) HasCombineIndex() bool {
	return t.PksSize > 1
}

func (t ModelEntity) PksString() string {
	return strings.Join(t.Pks, ", ")
}

func (t ModelEntity) Indexes() string {
	f := fmt.Sprintf
	quotePks := MapStrToStr(t.Pks, func(s string) string {
		return f("\"%s\"", strcase.ToSnake(s))
	})
	quoteString := strings.Join(quotePks, ", ")
	return f(`return []ent.Index{
        index.Fields(%s).
            Unique(),
    }`, quoteString)
}

func (t ModelEntity) NormalFields() []*ModelField {
	result := []*ModelField{}
	for _, f := range t.Fields {
		if !f.AutoCreatedInternal && !t.isRelationField(f.Name) && f.Name!=t.Pks[0]{
			result = append(result, f)
		}
	}
	return result
}

func (t ModelEntity) isRelationField(fldName string) bool {
	for _, r := range t.Edges() {
		if r.FieldName() == fldName {
			return true
		}
	}
	return false
}

func (t ModelEntity) PluralName() string {
	return PluralizeTypeName(t.Name)
}
//
//func (t ModelEntity) AdderName() string{
//	return "Add"+strcase.ToCamel(t.PluralName())
//}

func (t ModelEntity) GetField(fld string) *ModelField {
	for _, f := range t.Fields {
		if f.Name==fld{
			return f
		}
	}
	return nil
}

func (t ModelField) VarName() string {
	return strcase.ToSnake(t.Name)
}

func (t ModelField) IsDateTime() bool{
	if t.Type=="date-time" || t.Type=="date" || t.Type=="time"{
		return true
	}
	return false
}

func (t ModelField) SetterName() string{
	name:= "Set"+strcase.ToCamel(t.Name)
	if strings.HasSuffix(name, "Id"){
		name=strings.TrimRight(name, "Id")+"ID"
	}else if strings.HasSuffix(name, "Url"){
		name=strings.TrimRight(name, "Url")+"URL"
	}else if strings.Contains(name, "Url"){
		name=strings.ReplaceAll(name, "Url","URL")
	}
	return name
}

func (t ModelField) EntFieldType() string {
	f := fmt.Sprintf
	var resultDef string
	switch t.Type {
	case "id":
		//resultDef = f(`field.String("%s").MaxLen(20)`, t.VarName())
		resultDef = f(`field.Int("%s")`, t.VarName())
	case "id-long":
		resultDef = f(`field.String("%s").MaxLen(32)`, t.VarName())
	case "id-vlong":
		resultDef = f(`field.String("%s")`, t.VarName())
	case "blob", "object":
		resultDef = f(`field.JSON("%s", []byte{})`, t.VarName())
	case "byte-array":
		resultDef = f(`field.JSON("%s", []byte{})`, t.VarName())
	case "date-time", "date":
		resultDef = f(`field.Time("%s").
            Default(time.Now)`, t.VarName())
	case "time":
		resultDef = f(`field.Int("%s").
            Positive()`, t.VarName())
	case "currency-amount", "currency-precise", "fixed-point":
		resultDef = f(`field.Float("%s")`, t.VarName())
	case "floating-point":
		resultDef = f(`field.Float("%s")`, t.VarName())
	case "integer":
		resultDef = f(`field.Int("%s")`, t.VarName())
	case "numeric":
		resultDef = f(`field.Int("%s")`, t.VarName())
	case "indicator":
		resultDef = f(`field.Enum("%s").
            Values("Yes", "No", "Unknown")`, t.VarName())
	case "very-short":
		resultDef = f(`field.String("%s").MaxLen(10)`, t.VarName())
	case "url":
		//resultDef = f(`field.JSON("%s", &url.URL{})`, t.VarName())
		resultDef = f(`field.String("%s")`, t.VarName())
	default:
		resultDef = f(`field.String("%s")`, t.VarName())
	}
	if !t.NotNull {
		resultDef = resultDef + ".Optional()"
	}
	return resultDef
}

func (t ModelField) QuoteValue(v string) string {
	switch t.Type {
	case "id":
		return fmt.Sprintf("common.ParseId(\"%s\")", v)
	case "currency-amount", "currency-precise", "fixed-point", "floating-point",
		"integer", "numeric", "time":
		return v
	case "date-time", "date":
		return fmt.Sprintf("common.ParseDateTime(\"%s\")", v)
	case "indicator":
		return "\"" +indicatorValue(v)+ "\""
	default:
		return "\"" + v + "\""
	}
}

func indicatorValue(val string) string {
	switch val {
	case "Y": return "Yes"
	case "N": return "No"
	default:
		return "Unknown"
	}
}

func (t ModelField) EthFieldType() string {
	return EthType(t.Type)
}

func (t ModelField) GoFieldType() string {
	return FieldType(t.Type)
}

func (t ModelField) GoRefFieldType() string {
	typ := FieldType(t.Type)
	if typ[0] == '*' {
		return typ
	}
	return "*" + typ
}

func (t ModelRelation) FieldName() string {
	return t.Keymaps[0].FieldName
}
func (t ModelRelation) Keys() string {
	var keys []string
	for _, k := range t.Keymaps {
		keys = append(keys, k.FieldName)
	}
	return strings.Join(keys, ", ")
}

func (t ModelRelation) PluralName() string {
	return PluralizeTypeName(t.Name)
}

func (t ModelRelation) SnakecaseName() string {
	return strcase.ToSnake(t.Name)
}

func (t ModelRelation) AdderName() string{
	//if t.SelfRelation{
	if strings.HasPrefix(t.Name, "Child"){
		return "AddChildren"
	}
	//}
	return "Add"+strcase.ToCamel(t.PluralName())
}

func (t ModelRelation) SetterName() string{
	if t.SelfRelation{
		if strings.HasPrefix(t.Name, "Parent") ||
			strings.HasPrefix(t.Name, "PrimaryParent"){
			return "SetParent"
		}
	}
	return "Set"+strcase.ToCamel(t.Name)
}

func (t ModelRelation) HashBackref() bool {
	return t.Type == "one" || t.Type == "one-nofk"
}

func (t ModelEntity) Edges() []*ModelRelation {
	var result []*ModelRelation
	for _, r := range t.Relations {
		//if len(r.Keymaps) == 1 {
		if contains(*t.EntitiesInPkg, r.RelEntityName) {
			result = append(result, r)
		}
		//}
	}
	return result
}

func (t ModelEntity) GetBackRef(entName string, rel *ModelRelation) *ModelRelation {
	//log.Printf(".. get backref %s.%s in entity %s", entName, rel.Name, t.Name)
	for _, r := range t.Relations {
		//if len(r.Keymaps) == 1 {
		//log.Printf("%s == %s, %s == %s, %s == %s", r.Keymaps, rel.Keymaps,
		//	rel.RelEntityName, t.Name, r.RelEntityName, entName)
		//log.Println("rel.RelEntityName==t.Name", rel.RelEntityName==t.Name)
		if rel.RelEntityName == t.Name &&
			r.RelEntityName == entName &&
			rel.Keymaps[0].FieldName == r.Keymaps[0].RelFieldName &&
			rel.Keymaps[0].RelFieldName == r.Keymaps[0].FieldName {
			return r
		}
		//}else{
		//log.Println("\tRelation has multi-keymaps", len(rel.Keymaps))
		//}
	}

	return nil
}

func (t ModelEntity) GetBackRefNameAndType(entName string, rel *ModelRelation) (string, string, error) {
	ref := t.GetBackRef(entName, rel)
	if ref == nil {
		return "", "", errors.New(fmt.Sprintf(
			"Cannot find backref for relation %s in entity %s for %s",
			rel.Name, t.Name, entName))
	} else {
		if ref.Type == "one" || ref.Type == "one-nofk" {
			return ref.SnakecaseName(), ref.Type, nil
		} else {
			return ref.PluralName(), ref.Type, nil
		}
	}
}

func LoadModelEntity(filename string) (*ModelEntity, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}
	var modelEntity ModelEntity
	err = json.Unmarshal(content, &modelEntity)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return &modelEntity, err
}

func GenModelEntity(templateFile string, inputFile string, writer io.Writer) error {
	m, err := LoadModelEntity(inputFile)
	if err != nil {
		return err
	}
	return GenModelEntityWithMeta(templateFile, m, writer)
}

// func GenModelEntityWithMeta(templateFile string, m *ModelEntity, writer io.Writer) error {
func GenModelEntityWithMeta(templateFile string, m interface{}, writer io.Writer) error {
	tf := template.FuncMap{
		"title":     strings.Title,
		"lower":     strings.ToLower,
		"snakecase": strcase.ToSnake,
		"isUniquePk": func(ent ModelEntity, fld string) bool {
			return ent.IsUniqueIdField(fld)
		},
	}

	tpl := template.New("tpl").Funcs(tf)
	cnt, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	tt, err := tpl.Parse(string(cnt))
	if err != nil {
		return err
	}

	if err = tt.Execute(writer, m); err != nil {
		return err
	}

	return nil
}
