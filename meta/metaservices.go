package meta

import (
	"bytes"
	"encoding/json"
	"github.com/samlet/petrel/services"
	"io/ioutil"
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

func WrapResult(err error, body string, resp *MetaResponse) error {

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

func GenEntityMetaFromTemplate(path string, entityName string) (string, error) {
	meta, err := GetEntityMeta(entityName)
	if err != nil {
		return "", err
	}
	tplData, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	result := GenEntity(&meta.Data.Entity, string(tplData))
	return result, nil
}
