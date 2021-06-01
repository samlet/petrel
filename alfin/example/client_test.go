package example

import (
	"github.com/samlet/petrel/alfin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleClient_CreateExample(t *testing.T) {
	client := newClient("")
	rs, err := client.CreateExample(&alfin.CreateExampleParams{
		ExampleTypeId: "CONTRIVED",
		StatusId:      "EXST_IN_DESIGN",
		ExampleName:   "Example 100",
	})
	assert.Nil(t, err)
	assert.NotNil(t, rs)
	println(rs.ExampleId)
}
