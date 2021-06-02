// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"


type CreateExampleParams struct {
    services.AlfinParams
    ExampleTypeId string    `json:"exampleTypeId,omitempty"`
    StatusId string    `json:"statusId,omitempty"`
    ExampleName string    `json:"exampleName,omitempty"`
    Description string    `json:"description,omitempty"`
    LongDescription string    `json:"longDescription,omitempty"`
    Comments string    `json:"comments,omitempty"`
    ExampleSize int64    `json:"exampleSize,omitempty"`
    ExampleDate *services.Timestamp    `json:"exampleDate,omitempty"`
    AnotherDate *services.Timestamp    `json:"anotherDate,omitempty"`
    AnotherText string    `json:"anotherText,omitempty"`
}

type CreateExampleResult struct {
    services.APIResource
    ExampleId string   `json:"exampleId,omitempty"`
}

type DeleteExampleParams struct {
    services.AlfinParams
    ExampleId string    `json:"exampleId,omitempty"`
}

type DeleteExampleResult struct {
    services.APIResource
}

type UpdateExampleParams struct {
    services.AlfinParams
    ExampleId string    `json:"exampleId,omitempty"`
    ExampleTypeId string    `json:"exampleTypeId,omitempty"`
    StatusId string    `json:"statusId,omitempty"`
    ExampleName string    `json:"exampleName,omitempty"`
    Description string    `json:"description,omitempty"`
    LongDescription string    `json:"longDescription,omitempty"`
    Comments string    `json:"comments,omitempty"`
    ExampleSize int64    `json:"exampleSize,omitempty"`
    ExampleDate *services.Timestamp    `json:"exampleDate,omitempty"`
    AnotherDate *services.Timestamp    `json:"anotherDate,omitempty"`
    AnotherText string    `json:"anotherText,omitempty"`
}

type UpdateExampleResult struct {
    services.APIResource
    OldStatusId string   `json:"oldStatusId,omitempty"`
}

type CreateExampleStatusParams struct {
    services.AlfinParams
    ExampleId string    `json:"exampleId,omitempty"`
    StatusId string    `json:"statusId,omitempty"`
}

type CreateExampleStatusResult struct {
    services.APIResource
}

type CreateExampleFeatureParams struct {
    services.AlfinParams
    FeatureSourceEnumId string    `json:"featureSourceEnumId,omitempty"`
    Description string    `json:"description,omitempty"`
}

type CreateExampleFeatureResult struct {
    services.APIResource
    ExampleFeatureId string   `json:"exampleFeatureId,omitempty"`
}


// Interface
type ExampleOps interface {
    // CreateExample Create a Example
    CreateExample(params *CreateExampleParams) (*CreateExampleResult, error)
    // DeleteExample Delete a Example
    DeleteExample(params *DeleteExampleParams) (*DeleteExampleResult, error)
    // UpdateExample Update a Example
    UpdateExample(params *UpdateExampleParams) (*UpdateExampleResult, error)
    // CreateExampleStatus Create a ExampleStatus
    CreateExampleStatus(params *CreateExampleStatusParams) (*CreateExampleStatusResult, error)
    // CreateExampleFeature Create a ExampleFeature
    CreateExampleFeature(params *CreateExampleFeatureParams) (*CreateExampleFeatureResult, error)
}

