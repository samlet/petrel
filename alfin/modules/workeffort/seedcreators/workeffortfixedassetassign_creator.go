package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffortFixedAssetAssign(ctx context.Context) error {
	log.Println("WorkEffortFixedAssetAssign creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortFixedAssetAssign

	c, err = client.WorkEffortFixedAssetAssign.Create().SetStringRef("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__demo_projector__1199145600__workeffortfixedassetassign: %v", err)
		return err
	}
	cache.Put("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign", c)

	return nil
}
