// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"

type CreateExampleItemParams struct {
	services.AlfinParams
	ExampleId   string  `json:"exampleId,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	AmountUomId string  `json:"amountUomId,omitempty"`
}

type CreateExampleItemResult struct {
	services.APIResource
	ExampleItemSeqId string `json:"exampleItemSeqId,omitempty"`
}

type DeleteExampleItemParams struct {
	services.AlfinParams
	ExampleId        string `json:"exampleId,omitempty"`
	ExampleItemSeqId string `json:"exampleItemSeqId,omitempty"`
}

type DeleteExampleItemResult struct {
	services.APIResource
}

type UpdateExampleItemParams struct {
	services.AlfinParams
	ExampleId        string  `json:"exampleId,omitempty"`
	ExampleItemSeqId string  `json:"exampleItemSeqId,omitempty"`
	Description      string  `json:"description,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
	AmountUomId      string  `json:"amountUomId,omitempty"`
}

type UpdateExampleItemResult struct {
	services.APIResource
}

// Interface
type ExampleItemOps interface {
	// CreateExampleItem Create a ExampleItem
	CreateExampleItem(params *CreateExampleItemParams) (*CreateExampleItemResult, error)
	// DeleteExampleItem Delete a ExampleItem
	DeleteExampleItem(params *DeleteExampleItemParams) (*DeleteExampleItemResult, error)
	// UpdateExampleItem Update a ExampleItem
	UpdateExampleItem(params *UpdateExampleItemParams) (*UpdateExampleItemResult, error)
}
