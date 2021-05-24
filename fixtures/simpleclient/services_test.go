package main

import (
	"fmt"
	"github.com/samlet/petrel/fixtures/common"
	"testing"
)

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestFindOrder(t *testing.T) {
	resp, err := FindOrder("WSCO10010")
	if err != nil {
		errHandler(err)
		return
	}
	println(resp.StatusCode)
	//fmt.Printf("%s\n", resp.Data["item"])
	common.DumpMap("", resp.Data)
}

func TestCreateGood(t *testing.T) {
	resp, err := CreateGood("Demo Product")
	if err != nil {
		errHandler(err)
		return
	}
	println(resp.StatusCode)
	common.DumpMap("", resp.Data)

	id := resp.Data["productId"].(string)
	println(id)
}
