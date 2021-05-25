package meta

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const ENTITY_DEF = `
type {{.Name}} struct {
{{- range .Fields}}
	{{title .Name}} {{.Type}}
{{- end}}
}
`

func TestGenEntity(t *testing.T) {
	meta, err := GetEntityMeta("Product")
	assert.True(t, err == nil)
	result := GenEntity(&meta.Data.Entity, ENTITY_DEF)
	println(result)
}

func TestGenEntityFromFile(t *testing.T) {
	meta, err := GetEntityMeta("Product")
	assert.True(t, err == nil)
	tplData, err := ioutil.ReadFile("../incls/entity_def.tmpl")
	if err != nil {
		panic(err)
	}
	result := GenEntity(&meta.Data.Entity, string(tplData))
	println(result)
}
