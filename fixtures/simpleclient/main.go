package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	result, _:=InvokeService("performFindList", `{
		"entityName":"WorkEffort",
		"viewIndex": 0,
		"viewSize": 20,
		"inputFields":{
			"workEffortTypeId":"ROUTING"
		}
	}`)

	fmt.Println(result)
}
