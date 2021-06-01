// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package example

import (
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/services"
	"net/http"
)

type ExampleClient struct {
	B   *services.AlfinBackend
	Key string
}

func newClient(key string) ExampleClient {
	return ExampleClient{services.NewAlfinBackend(), key}
}

// CreateExample Create a Example
func (c *ExampleClient) CreateExample(params *alfin.CreateExampleParams) (*alfin.CreateExampleResult, error) {
	reversal := &alfin.CreateExampleResult{}
	err := c.B.Call(http.MethodPost, "createExample", c.Key, params, reversal)
	return reversal, err
}

// DeleteExample Delete a Example
func (c *ExampleClient) DeleteExample(params *alfin.DeleteExampleParams) (*alfin.DeleteExampleResult, error) {
	reversal := &alfin.DeleteExampleResult{}
	err := c.B.Call(http.MethodPost, "deleteExample", c.Key, params, reversal)
	return reversal, err
}

// UpdateExample Update a Example
func (c *ExampleClient) UpdateExample(params *alfin.UpdateExampleParams) (*alfin.UpdateExampleResult, error) {
	reversal := &alfin.UpdateExampleResult{}
	err := c.B.Call(http.MethodPost, "updateExample", c.Key, params, reversal)
	return reversal, err
}

// CreateExampleStatus Create a ExampleStatus
func (c *ExampleClient) CreateExampleStatus(params *alfin.CreateExampleStatusParams) (*alfin.CreateExampleStatusResult, error) {
	reversal := &alfin.CreateExampleStatusResult{}
	err := c.B.Call(http.MethodPost, "createExampleStatus", c.Key, params, reversal)
	return reversal, err
}

// CreateExampleFeature Create a ExampleFeature
func (c *ExampleClient) CreateExampleFeature(params *alfin.CreateExampleFeatureParams) (*alfin.CreateExampleFeatureResult, error) {
	reversal := &alfin.CreateExampleFeatureResult{}
	err := c.B.Call(http.MethodPost, "createExampleFeature", c.Key, params, reversal)
	return reversal, err
}
