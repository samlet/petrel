package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateWorkEffortFixedAssetAssign(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortFixedAssetAssign

	c = cache.Get("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign").(*ent.WorkEffortFixedAssetAssign)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetFixedAsset(cache.Get("demo_projector__fixedasset").(*ent.FixedAsset)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__demo_projector__1199145600__workeffortfixedassetassign: %v", err)
		return err
	}
	cache.Put("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign", c)

	return nil
}
