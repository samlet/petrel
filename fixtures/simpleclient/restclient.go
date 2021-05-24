package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func InvokeService(srv string, reqstr string) (string, error) {
	url := fmt.Sprintf("https://localhost:8443/rest/services/%s", srv)

	payload := strings.NewReader(reqstr)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJ1c2VyTG9naW5JZCI6ImFkbWluIiwiaXNzIjoiQXBhY2hlT0ZCaXoiLCJleHAiOjE2MzQ1OTc3NTcsImlhdCI6MTYxNjU5Nzc1N30.Luuf2bK7ZJ8KE_CtsA3iPZ189i-Qbm2qK5r5VfeQcJqIyTKy4DHf2fBAp37W8OtU6SIplwCdnbTMtHuCZ5h8cA")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	return string(body), nil
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
