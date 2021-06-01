package alfin

import (
	"encoding/json"
	"github.com/samlet/petrel/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func pretty(v interface{}) {
	r, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(r))
}

func TestBackendErrCall(t *testing.T) {
	backend := services.NewAlfinBackend()
	//params, err:= services.NewParams(CreatePersonParams{FirstName: "samlet"})
	//if err != nil {
	//	panic(err)
	//}
	params := CreateExampleParams{ExampleName: "samlet"}
	reversal := &CreateExampleResult{}
	err := backend.Call(http.MethodPost, "createExample", "",
		&params, reversal)
	if err != nil {
		panic(err)
	}
	println(reversal.LastResponse.StatusCode, reversal.LastResponse.Status)
	//println(string(reversal.LastResponse.RawJSON))
	pretty(reversal.LastTypedResponse)
}

func TestBackendOkCall(t *testing.T) {
	backend := services.NewAlfinBackend()
	params := CreateExampleParams{
		ExampleTypeId: "CONTRIVED",
		StatusId:      "EXST_IN_DESIGN",
		ExampleName:   "Example 100",
	}
	reversal := &CreateExampleResult{}
	err := backend.Call(http.MethodPost, "createExample", "", &params, reversal)
	if err != nil {
		panic(err)
	}
	println(reversal.LastResponse.StatusCode, reversal.LastResponse.Status)
	//println(string(reversal.LastResponse.RawJSON))
	pretty(reversal.LastTypedResponse)
	println("created:", reversal.ExampleId)
}

type ExampleClient struct {
	B   *services.AlfinBackend
	Key string
}

func newClient() ExampleClient {
	return ExampleClient{services.NewAlfinBackend(), ""}
}

func (c *ExampleClient) CreateExample(params *CreateExampleParams) (*CreateExampleResult, error) {
	reversal := &CreateExampleResult{}
	err := c.B.Call(http.MethodPost, "createExample", "", params, reversal)
	return reversal, err
}

func TestExampleClient(t *testing.T) {
	client := newClient()
	rs, err := client.CreateExample(&CreateExampleParams{
		ExampleTypeId: "CONTRIVED",
		StatusId:      "EXST_IN_DESIGN",
		ExampleName:   "Example 100",
	})
	assert.Nil(t, err)
	assert.NotNil(t, rs)
	println(rs.ExampleId)
}
