package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateFacilityType(ctx context.Context) error {
	log.Println("FacilityType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityType
	failures := 0

	c = cache.Get("building__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Building").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update building__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("building__facilitytype", c)
	}

	c = cache.Get("floor__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Floor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update floor__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("floor__facilitytype", c)
	}

	c = cache.Get("office__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Office").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update office__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("office__facilitytype", c)
	}

	c = cache.Get("call_center__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Call Center").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update call_center__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("call_center__facilitytype", c)
	}

	c = cache.Get("plant__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Plant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update plant__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("plant__facilitytype", c)
	}

	c = cache.Get("room__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Room").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update room__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("room__facilitytype", c)
	}

	c = cache.Get("retail_store__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Retail Store").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update retail_store__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("retail_store__facilitytype", c)
	}

	c = cache.Get("warehouse__facilitytype").(*ent.FacilityType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Warehouse").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update warehouse__facilitytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("warehouse__facilitytype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
