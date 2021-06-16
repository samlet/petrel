package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateFacilityGroupType(ctx context.Context) error {
	log.Println("FacilityGroupType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityGroupType
	failures := 0

	c = cache.Get("mgmt_structure__facilitygrouptype").(*ent.FacilityGroupType)
	c, err = c.Update().
		SetDescription("Management Structure").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update mgmt_structure__facilitygrouptype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("mgmt_structure__facilitygrouptype", c)
	}

	c = cache.Get("pricing_group__facilitygrouptype").(*ent.FacilityGroupType)
	c, err = c.Update().
		SetDescription("Pricing Group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pricing_group__facilitygrouptype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pricing_group__facilitygrouptype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
