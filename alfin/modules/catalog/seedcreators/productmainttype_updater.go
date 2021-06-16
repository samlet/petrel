package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductMaintType(ctx context.Context) error {
	log.Println("ProductMaintType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductMaintType
	failures := 0

	c = cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Vehicle Maintenance").
		AddChildren(cache.Get("oil_change__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("serp_belt__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("refuel__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("replace_battery__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("tune_up__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("check_battery__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("chassis_lube__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("rotate_tires__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update vehicle_maint__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("vehicle_maint__productmainttype", c)
	}

	c = cache.Get("oil_change__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Oil Change").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update oil_change__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("oil_change__productmainttype", c)
	}

	c = cache.Get("serp_belt__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Serpentine Belt Replacement").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update serp_belt__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("serp_belt__productmainttype", c)
	}

	c = cache.Get("refuel__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Re-Fuel").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update refuel__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("refuel__productmainttype", c)
	}

	c = cache.Get("replace_battery__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Replace Battery").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update replace_battery__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("replace_battery__productmainttype", c)
	}

	c = cache.Get("tune_up__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Tune Up").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tune_up__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tune_up__productmainttype", c)
	}

	c = cache.Get("check_battery__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Check Battery").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update check_battery__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("check_battery__productmainttype", c)
	}

	c = cache.Get("chassis_lube__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Chassis Lubrication").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update chassis_lube__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("chassis_lube__productmainttype", c)
	}

	c = cache.Get("rotate_tires__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Rotate Tires").
		SetParent(cache.Get("vehicle_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update rotate_tires__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("rotate_tires__productmainttype", c)
	}

	c = cache.Get("hvac_maint__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("HVAC Maintenance").
		AddChildren(cache.Get("hvac_replace_filter__productmainttype").(*ent.ProductMaintType)).
		AddChildren(cache.Get("hvac_check_refr__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hvac_maint__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hvac_maint__productmainttype", c)
	}

	c = cache.Get("hvac_replace_filter__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Replace Air Filter").
		SetParent(cache.Get("hvac_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hvac_replace_filter__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hvac_replace_filter__productmainttype", c)
	}

	c = cache.Get("hvac_check_refr__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Check/Recharge Refrigerant").
		SetParent(cache.Get("hvac_maint__productmainttype").(*ent.ProductMaintType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hvac_check_refr__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hvac_check_refr__productmainttype", c)
	}

	c = cache.Get("wash__productmainttype").(*ent.ProductMaintType)
	c, err = c.Update().
		SetDescription("Wash").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update wash__productmainttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("wash__productmainttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
