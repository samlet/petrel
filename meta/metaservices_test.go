package meta

import (
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEntityMeta(t *testing.T) {
	meta, err := GetEntityMeta("Product")
	assert.True(t, err == nil)
	println(meta.StatusCode)
	assert.Equal(t, "Product", meta.Data.Entity.Name)
	for i, rel := range meta.Data.Entity.Relations {
		_, _ = pretty.Printf("%d. %s\n", i, rel)
	}
}

func TestGetServiceMeta(t *testing.T) {
	meta, err := GetServiceMeta("createPerson")
	assert.True(t, err == nil)
	println(meta.StatusCode)
	assert.Equal(t, "createPerson", meta.Data.Service.Name)
	for i, rel := range meta.Data.Service.Parameters {
		_, _ = pretty.Printf("%d. %s\n", i, rel)
	}
}
