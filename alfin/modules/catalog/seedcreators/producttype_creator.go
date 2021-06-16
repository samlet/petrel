package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductType(ctx context.Context) error {
	log.Println("ProductType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductType

	c, err = client.ProductType.Create().SetStringRef("test_type__producttype").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_type__producttype: %v", err)
		return err
	}
	cache.Put("test_type__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("asset_usage__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Fixed Asset Usage").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create asset_usage__producttype: %v", err)
		return err
	}
	cache.Put("asset_usage__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("service__producttype").
		SetIsPhysical("No").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create service__producttype: %v", err)
		return err
	}
	cache.Put("service__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("good__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Good").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create good__producttype: %v", err)
		return err
	}
	cache.Put("good__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("raw_material__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Raw Material").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create raw_material__producttype: %v", err)
		return err
	}
	cache.Put("raw_material__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("subassembly__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Subassembly").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create subassembly__producttype: %v", err)
		return err
	}
	cache.Put("subassembly__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("finished_good__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Finished Good").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create finished_good__producttype: %v", err)
		return err
	}
	cache.Put("finished_good__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("digital_good__producttype").
		SetIsPhysical("No").
		SetIsDigital("Yes").
		SetHasTable("No").
		SetDescription("Digital Good").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create digital_good__producttype: %v", err)
		return err
	}
	cache.Put("digital_good__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("findig_good__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("Yes").
		SetHasTable("No").
		SetDescription("Finished/Digital Good").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create findig_good__producttype: %v", err)
		return err
	}
	cache.Put("findig_good__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("aggregated__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Good").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create aggregated__producttype: %v", err)
		return err
	}
	cache.Put("aggregated__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("marketing_pkg__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marketing_pkg__producttype: %v", err)
		return err
	}
	cache.Put("marketing_pkg__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("marketing_pkg_auto__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package: Auto Manufactured").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marketing_pkg_auto__producttype: %v", err)
		return err
	}
	cache.Put("marketing_pkg_auto__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("marketing_pkg_pick__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package: Pick Assembly").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marketing_pkg_pick__producttype: %v", err)
		return err
	}
	cache.Put("marketing_pkg_pick__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("wip__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Work In Process").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create wip__producttype: %v", err)
		return err
	}
	cache.Put("wip__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("aggregated_conf__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Good Configuration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create aggregated_conf__producttype: %v", err)
		return err
	}
	cache.Put("aggregated_conf__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("asset_usage_out_in__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Fixed Asset Usage For Rental of an asset which is shipped from and returned to inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create asset_usage_out_in__producttype: %v", err)
		return err
	}
	cache.Put("asset_usage_out_in__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("service_product__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Service a product using inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create service_product__producttype: %v", err)
		return err
	}
	cache.Put("service_product__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("aggregated_service__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Service using inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create aggregated_service__producttype: %v", err)
		return err
	}
	cache.Put("aggregated_service__producttype", c)

	c, err = client.ProductType.Create().SetStringRef("aggregatedserv_conf__producttype").
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Service Configuration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create aggregatedserv_conf__producttype: %v", err)
		return err
	}
	cache.Put("aggregatedserv_conf__producttype", c)

	return nil
}
