// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleitem

import (
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/services"
	"net/http"
)

type ExampleItemClient struct {
	B   *services.AlfinBackend
	Key string
}

func newClient(key string) ExampleItemClient {
	return ExampleItemClient{services.NewAlfinBackend(), key}
}

// CreateExampleItem Create a ExampleItem
func (c *ExampleItemClient) CreateExampleItem(params *alfin.CreateExampleItemParams) (*alfin.CreateExampleItemResult, error) {
	reversal := &alfin.CreateExampleItemResult{}
	err := c.B.Call(http.MethodPost, "createExampleItem", c.Key, params, reversal)
	return reversal, err
}

// DeleteExampleItem Delete a ExampleItem
func (c *ExampleItemClient) DeleteExampleItem(params *alfin.DeleteExampleItemParams) (*alfin.DeleteExampleItemResult, error) {
	reversal := &alfin.DeleteExampleItemResult{}
	err := c.B.Call(http.MethodPost, "deleteExampleItem", c.Key, params, reversal)
	return reversal, err
}

// UpdateExampleItem Update a ExampleItem
func (c *ExampleItemClient) UpdateExampleItem(params *alfin.UpdateExampleItemParams) (*alfin.UpdateExampleItemResult, error) {
	reversal := &alfin.UpdateExampleItemResult{}
	err := c.B.Call(http.MethodPost, "updateExampleItem", c.Key, params, reversal)
	return reversal, err
}
