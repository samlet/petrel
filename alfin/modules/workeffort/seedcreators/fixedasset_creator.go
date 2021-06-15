package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateFixedAsset(ctx context.Context) error {
	log.Println("FixedAsset creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FixedAsset

	c, err = client.FixedAsset.Create().SetStringRef("demo_projector__fixedasset").
		SetFixedAssetTypeID(common.ParseId("EQUIPMENT")).
		SetFixedAssetName("Overhead Projector").
		SetPurchaseCost(2000).
		SetPurchaseCostUomID(common.ParseId("USD")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo_projector__fixedasset: %v", err)
		return err
	}
	cache.Put("demo_projector__fixedasset", c)

	return nil
}
