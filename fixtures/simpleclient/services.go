package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

type P map[string]interface{}

type Response struct {
	StatusCode        int                    `json:"statusCode"`
	StatusDescription string                 `json:"statusDescription"`
	Data              map[string]interface{} `json:"data"`
}

func WrapResult(err error, body string, resp Response) (*Response, error) {
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader([]byte(body))
	err = json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func FindOrder(orderId string) (*Response, error) {
	var resp Response
	body, err := InvokeService("performFindItem",
		fmt.Sprintf(`{
				"entityName":"OrderHeader",
				"inputFields":{
					"orderId":"%s"
				}
			}`, orderId))
	return WrapResult(err, body, resp)
}

func Render(format string, p P) string {
	b := &bytes.Buffer{}
	template.Must(template.New("").Parse(format)).Execute(b, p)
	return b.String()
}

func CreateGood(productName string) (*Response, error) {
	var resp Response
	body, err := InvokeService("createProduct",
		Render(`{
					"internalName": "{{.productName}}",
					"description": "{{.productName}}",
					"productName": "{{.productName}}",
					"productTypeId": "FINISHED_GOOD",
					"billOfMaterialLevel": 0,
					"isVirtual": "N",
					"isVariant": "N"
			}`, P{"productName": productName}))
	return WrapResult(err, body, resp)
}
