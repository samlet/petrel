package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductMaintType(ctx context.Context) error {
	log.Println("ProductMaintType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductMaintType

	c, err = client.ProductMaintType.Create().SetStringRef("vehicle_maint__productmainttype").
		SetDescription("Vehicle Maintenance").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vehicle_maint__productmainttype: %v", err)
		return err
	}
	cache.Put("vehicle_maint__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("oil_change__productmainttype").
		SetDescription("Oil Change").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create oil_change__productmainttype: %v", err)
		return err
	}
	cache.Put("oil_change__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("serp_belt__productmainttype").
		SetDescription("Serpentine Belt Replacement").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create serp_belt__productmainttype: %v", err)
		return err
	}
	cache.Put("serp_belt__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("refuel__productmainttype").
		SetDescription("Re-Fuel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create refuel__productmainttype: %v", err)
		return err
	}
	cache.Put("refuel__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("replace_battery__productmainttype").
		SetDescription("Replace Battery").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create replace_battery__productmainttype: %v", err)
		return err
	}
	cache.Put("replace_battery__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("tune_up__productmainttype").
		SetDescription("Tune Up").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tune_up__productmainttype: %v", err)
		return err
	}
	cache.Put("tune_up__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("check_battery__productmainttype").
		SetDescription("Check Battery").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create check_battery__productmainttype: %v", err)
		return err
	}
	cache.Put("check_battery__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("chassis_lube__productmainttype").
		SetDescription("Chassis Lubrication").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create chassis_lube__productmainttype: %v", err)
		return err
	}
	cache.Put("chassis_lube__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("rotate_tires__productmainttype").
		SetDescription("Rotate Tires").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create rotate_tires__productmainttype: %v", err)
		return err
	}
	cache.Put("rotate_tires__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("hvac_maint__productmainttype").
		SetDescription("HVAC Maintenance").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hvac_maint__productmainttype: %v", err)
		return err
	}
	cache.Put("hvac_maint__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("hvac_replace_filter__productmainttype").
		SetDescription("Replace Air Filter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hvac_replace_filter__productmainttype: %v", err)
		return err
	}
	cache.Put("hvac_replace_filter__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("hvac_check_refr__productmainttype").
		SetDescription("Check/Recharge Refrigerant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hvac_check_refr__productmainttype: %v", err)
		return err
	}
	cache.Put("hvac_check_refr__productmainttype", c)

	c, err = client.ProductMaintType.Create().SetStringRef("wash__productmainttype").
		SetDescription("Wash").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create wash__productmainttype: %v", err)
		return err
	}
	cache.Put("wash__productmainttype", c)

	return nil
}
