package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateWorkEffortFixedAssetAssign(ctx context.Context) error {
	log.Println("WorkEffortFixedAssetAssign updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortFixedAssetAssign
	failures := 0

	c = cache.Get("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign").(*ent.WorkEffortFixedAssetAssign)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetFixedAsset(cache.Get("demo_projector__fixedasset").(*ent.FixedAsset)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__demo_projector__1199145600__workeffortfixedassetassign: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
