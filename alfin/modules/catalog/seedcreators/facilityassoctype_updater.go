package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateFacilityAssocType(ctx context.Context) error {
	log.Println("FacilityAssocType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityAssocType
	failures := 0

	c = cache.Get("backup_warehouse__facilityassoctype").(*ent.FacilityAssocType)
	c, err = c.Update().
		SetDescription("Facility that serves another facility in terms of inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update backup_warehouse__facilityassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("backup_warehouse__facilityassoctype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
