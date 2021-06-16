package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateFacilityGroup(ctx context.Context) error {
	log.Println("FacilityGroup updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityGroup
	failures := 0

	c = cache.Get("_na___facilitygroup").(*ent.FacilityGroup)
	c, err = c.Update().
		SetFacilityGroupName("Not Applicable").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update _na___facilitygroup: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("_na___facilitygroup", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
