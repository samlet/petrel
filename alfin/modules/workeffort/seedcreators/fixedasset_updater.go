package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateFixedAsset(ctx context.Context) error {
	log.Println("FixedAsset updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FixedAsset

	c = cache.Get("demo_projector__fixedasset").(*ent.FixedAsset)
	c, err = c.Update().
		SetFixedAssetTypeID(common.ParseId("EQUIPMENT")).
		SetFixedAssetName("Overhead Projector").
		SetPurchaseCost(2000).
		SetPurchaseCostUomID(common.ParseId("USD")).
		AddWorkEffortFixedAssetAssigns(cache.Get("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign").(*ent.WorkEffortFixedAssetAssign)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo_projector__fixedasset: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demo_projector__fixedasset", c)
	}

	return nil
}
