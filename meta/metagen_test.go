package meta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenEntity(t *testing.T) {
	meta, err := GetEntityMeta("Product")
	assert.True(t, err == nil)
	result := GenEntity(&meta.Data.Entity, ENTITY_DEF)
	println(result)
}
