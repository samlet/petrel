package alfin

import (
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type (
	ModelEntity struct{
		Name string `json:"name"`
		Fields []*ModelField `json:"fields"`
		Relations []*ModelRelation `json:"relations"`
		PksSize int `json:"pksSize"`
		Pks []string `json:"pks"`
	}

	ModelField struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Col string `json:"col"`
		Pk bool `json:"pk"`
		AutoCreatedInternal bool `json:"autoCreatedInternal"`
	}

	ModelRelation struct{
		Name string `json:"name"`
		Type string `json:"type"`
		RelEntityName string `json:"relEntityName"`
		FkName string `json:"fkName"`
		Keymaps []ModelKeymap `json:"keymaps"`
	}

	ModelKeymap struct {
		FieldName string `json:"fieldName"`
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

func (t ModelEntity) IsUniqueIdField(fldName string) bool{
	return t.PksSize==1 && contains(t.Pks, fldName)
}

func (t ModelEntity) UniquePk() string{
	if t.PksSize==1{
		return t.Pks[0]
	}else{
		return ""
	}
}

func (t ModelEntity) HasCombineIndex() bool{
	return t.PksSize>1
}

func (t ModelEntity) Indexes() string{
	f:=fmt.Sprintf
	quotePks := MapStrToStr(t.Pks, func(s string) string {
		return f("\"%s\"", strcase.ToSnake(s))
	})
	quoteString:=strings.Join(quotePks, ", ")
	return f(`return []ent.Index{
        index.Fields(%s).
            Unique(),
    }`, quoteString)
}

func (t ModelEntity) NormalFields() []*ModelField{
	result:=[]*ModelField{}
	for _,f := range t.Fields {
		if !f.AutoCreatedInternal{
			result=append(result, f)
		}
	}
	return result
}

func (t ModelEntity) PluralName() string{
	return PluralizeTypeName(t.Name)
}

func (t ModelField) VarName() string{
	return strcase.ToSnake(t.Name)
}

func (t ModelField) EntFieldType() string{
	f:=fmt.Sprintf
	var resultDef string
	switch t.Type {
	case "id":
		resultDef = f(`field.String("%s").MaxLen(20)`, t.VarName())
	case "id-long":
		resultDef = f(`field.String("%s").MaxLen(32)`, t.VarName())
	case "id-vlong":
		resultDef = f(`field.String("%s")`, t.VarName())
	case "blob", "object":
		resultDef = f(`field.JSON("%s", []byte{}).
            Optional()`, t.VarName())
	case "byte-array":
		resultDef = f(`field.JSON("%s", []byte{}).
            Optional()`, t.VarName())
	case "date-time", "date":
		resultDef = f(`field.Time("%s").
            Default(time.Now)`, t.VarName())
	case "time":
		resultDef = f(`field.Int("%s").
            Positive()`, t.VarName())
	case "currency-amount", "currency-precise", "fixed-point":
		resultDef = f(`field.Float("%s").
            Optional()`, t.VarName())
	case "floating-point":
		resultDef = f(`field.Float("%s").
            Optional()`, t.VarName())
	case "integer":
		resultDef = f(`field.Int("%s").
            Optional()`, t.VarName())
	case "numeric":
		resultDef = f(`field.Int("%s").
            Optional()`, t.VarName())
	case "indicator":
		resultDef = f(`field.Enum("%s").
            Values("Y", "N", "-").
            Optional()`, t.VarName())
	case "very-short":
		resultDef = f(`field.String("%s").MaxLen(10)`, t.VarName())
	case "url":
		resultDef = f(`field.JSON("%s", &url.URL{}).
            Optional()`, t.VarName())
	default:
		resultDef = f(`field.String("%s")`, t.VarName())
	}
	return resultDef
}

func (t ModelRelation) FieldName() string {
	return t.Keymaps[0].FieldName
}
func (t ModelRelation) PluralName() string{
	return PluralizeTypeName(t.Name)
}

func (t ModelEntity) Edges() []*ModelRelation {
	var result []*ModelRelation
	for _, r := range t.Relations {
		if len(r.Keymaps)==1{
			result=append(result, r)
		}
	}
	return result
}

func LoadModelEntity(filename string) (*ModelEntity, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}
	var modelEntity ModelEntity
	err=json.Unmarshal(content, &modelEntity)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return &modelEntity, err
}

func GenModelEntity(templateFile string, inputFile string, writer io.Writer) error{
	tf := template.FuncMap{
		"title":     strings.Title,
		"snakecase": strcase.ToSnake,
		"isUniquePk": func(ent ModelEntity, fld string) bool {
			return ent.IsUniqueIdField(fld)
		},
	}

	tpl := template.New("tpl").Funcs(tf)
	cnt, err:=ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	tt, err := tpl.Parse(string(cnt))
	if err != nil {
		return err
	}

	m, err:=LoadModelEntity(inputFile)
	if err = tt.Execute(writer, m); err != nil {
		return err
	}

	return nil
}