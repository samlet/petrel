package services

import (
	"fmt"
	"github.com/crossdock/crossdock-go/assert"
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
	if resp.StatusCode != 200 {
		t.Fatalf("Create good failed")
	}
	common.DumpMap("", resp.Data)

	if id, found := resp.Data["productId"].(string); found {
		println(id)

		// find it
		resp, err = FindProduct(id)
		if err != nil {
			errHandler(err)
			return
		}
		println(resp.StatusCode)
		common.DumpMap("", resp.Data)
	} else {
		t.Errorf("Cannot get product id from response")
	}
}

func TestFindProduct(t *testing.T) {
	id := "_UNK_"
	resp, err := FindProduct(id)
	if err != nil {
		errHandler(err)
		return
	}
	println(resp.StatusCode)
	common.DumpMap("", resp.Data)
	_, found := resp.Data["item"]
	assert.False(t, found)
}
