package meta

import (
	"bytes"
	"encoding/json"
	"github.com/samlet/petrel/services"
)

type P map[string]interface{}

type MetaResponse struct {
	StatusCode        int               `json:"statusCode"`
	StatusDescription string            `json:"statusDescription"`
	Data              EntityMetaWrapper `json:"data"`
}

type EntityMetaWrapper struct {
	Entity EntityMeta `json:"entity"`
}

type EntityMeta struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	PackageName string           `json:"packageName"`
	Fields      []FieldMeta      `json:"fields"`
	Relations   []EntityRelation `json:"relations"`
}

type FieldMeta struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Pk      bool   `json:"pk"`
	NotNull bool   `json:"notNull"`
}

type EntityRelation struct {
	Name          string `json:"name"`
	RelEntityName string `json:"rel_entity_name"`
	Type          string `json:"type"`
}

type ServiceMetaResponse struct {
	StatusCode        int                `json:"statusCode"`
	StatusDescription string             `json:"statusDescription"`
	Data              ServiceMetaWrapper `json:"data"`
}

type ServiceMetaWrapper struct {
	Service ServiceMeta `json:"service"`
}

type ServiceMeta struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Action      string          `json:"action"`
	PackageName string          `json:"packageName,omitempty"`
	Parameters  []ParameterMeta `json:"parameters"`
}

type ParameterMeta struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Mode     string `json:"mode"`
	Repr     string `json:"repr"`
	Optional bool   `json:"optional"`
}

func WrapResult(err error, body string, resp interface{}) error {
	r := bytes.NewReader([]byte(body))
	err = json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return err
	}
	return nil
}

func GetEntityMeta(entityName string) (*MetaResponse, error) {
	var resp MetaResponse
	body, err := services.Post("getEntityMeta", P{"entityName": entityName})
	if err != nil {
		return nil, err
	}
	err = WrapResult(err, body, &resp)
	return &resp, err
}

func GetServiceMeta(serviceName string) (*ServiceMetaResponse, error) {
	var resp ServiceMetaResponse
	body, err := services.Post("getServiceMeta", P{"serviceName": serviceName})
	if err != nil {
		return nil, err
	}
	err = WrapResult(err, body, &resp)
	return &resp, err
}
