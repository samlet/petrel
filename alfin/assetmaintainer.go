package alfin

import (
	rice "github.com/GeertJohan/go.rice"
	"log"
)

type AssetBox struct {
	Assets *rice.Box
}

func NewAssetBox() *AssetBox {
	assertBox, err := rice.FindBox("assets")
	if err != nil {
		log.Fatal(err)
	}
	return &AssetBox{assertBox}
}

func (b AssetBox) String(s string) (string, error) {
	return b.Assets.String(s)
}
