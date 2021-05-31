package services

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

func WrapResult(err error, body string) (*Response, error) {
	var resp Response
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

func FindRoutings(size int) (string, error) {
	return InvokeService("performFindList",
		fmt.Sprintf(`{
				"entityName":"WorkEffort",
				"viewIndex": 0,
				"viewSize": %d,
				"inputFields":{
					"workEffortTypeId":"ROUTING"
				}
			}`, size))
}

func FindOrder(orderId string) (*Response, error) {
	body, err := InvokeService("performFindItem",
		fmt.Sprintf(`{
				"entityName":"OrderHeader",
				"inputFields":{
					"orderId":"%s"
				}
			}`, orderId))
	return WrapResult(err, body)
}

func Render(format string, p map[string]interface{}) string {
	b := &bytes.Buffer{}
	template.Must(template.New("").Parse(format)).Execute(b, p)
	return b.String()
}

func CreateGood(productName string) (*Response, error) {
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
	return WrapResult(err, body)
}

func FindProduct(productId string) (*Response, error) {
	body, err := InvokeService("performFindItem",
		Render(`{
				"entityName":"Product",
				"inputFields":{
					"productId":"{{.productId}}"
				}
			}`, P{"productId": productId}))
	return WrapResult(err, body)
}
