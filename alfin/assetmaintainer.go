package alfin

import (
	rice "github.com/GeertJohan/go.rice"
	"log"
)

type AssetBox struct {
	Assets    *rice.Box
	Templates *rice.Box
	Configs   *rice.Box
}

func NewAssetBox() *AssetBox {
	assertBox, err := rice.FindBox("assets")
	if err != nil {
		log.Fatal(err)
	}
	return &AssetBox{assertBox,
		rice.MustFindBox("templates"),
		rice.MustFindBox("conf"),
	}
}

func (b AssetBox) String(s string) (string, error) {
	return b.Assets.String(s)
}

func (b AssetBox) Template(s string) (string, error) {
	return b.Templates.String(s)
}

func (b AssetBox) Conf(s string) (string, error) {
	return b.Configs.String(s)
}
