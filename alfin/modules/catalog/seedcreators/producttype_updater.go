package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductType(ctx context.Context) error {
	log.Println("ProductType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductType
	failures := 0

	c = cache.Get("test_type__producttype").(*ent.ProductType)
	c, err = c.Update().
		AddProducts(cache.Get("test_product_a__product").(*ent.Product)).
		AddProducts(cache.Get("test_product_b__product").(*ent.Product)).
		AddProducts(cache.Get("test_product_c__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_type__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_type__producttype", c)
	}

	c = cache.Get("asset_usage__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Fixed Asset Usage").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update asset_usage__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("asset_usage__producttype", c)
	}

	c = cache.Get("service__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("No").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update service__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("service__producttype", c)
	}

	c = cache.Get("good__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Good").
		AddChildren(cache.Get("raw_material__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("subassembly__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("finished_good__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("digital_good__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("findig_good__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("aggregated__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("marketing_pkg__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("wip__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update good__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("good__producttype", c)
	}

	c = cache.Get("raw_material__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Raw Material").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update raw_material__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("raw_material__producttype", c)
	}

	c = cache.Get("subassembly__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Subassembly").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update subassembly__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("subassembly__producttype", c)
	}

	c = cache.Get("finished_good__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Finished Good").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		AddProducts(cache.Get("demoproduct__product").(*ent.Product)).
		AddProducts(cache.Get("demoproduct-1__product").(*ent.Product)).
		AddProducts(cache.Get("demoproduct-2__product").(*ent.Product)).
		AddProducts(cache.Get("demoproduct-3__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update finished_good__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("finished_good__producttype", c)
	}

	c = cache.Get("digital_good__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("No").
		SetIsDigital("Yes").
		SetHasTable("No").
		SetDescription("Digital Good").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update digital_good__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("digital_good__producttype", c)
	}

	c = cache.Get("findig_good__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("Yes").
		SetHasTable("No").
		SetDescription("Finished/Digital Good").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update findig_good__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("findig_good__producttype", c)
	}

	c = cache.Get("aggregated__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Good").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("aggregated_conf__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("aggregatedserv_conf__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update aggregated__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("aggregated__producttype", c)
	}

	c = cache.Get("marketing_pkg__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("marketing_pkg_auto__producttype").(*ent.ProductType)).
		AddChildren(cache.Get("marketing_pkg_pick__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update marketing_pkg__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("marketing_pkg__producttype", c)
	}

	c = cache.Get("marketing_pkg_auto__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package: Auto Manufactured").
		SetParent(cache.Get("marketing_pkg__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update marketing_pkg_auto__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("marketing_pkg_auto__producttype", c)
	}

	c = cache.Get("marketing_pkg_pick__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Marketing Package: Pick Assembly").
		SetParent(cache.Get("marketing_pkg__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update marketing_pkg_pick__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("marketing_pkg_pick__producttype", c)
	}

	c = cache.Get("wip__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Work In Process").
		SetParent(cache.Get("good__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update wip__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("wip__producttype", c)
	}

	c = cache.Get("aggregated_conf__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Good Configuration").
		SetParent(cache.Get("aggregated__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update aggregated_conf__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("aggregated_conf__producttype", c)
	}

	c = cache.Get("asset_usage_out_in__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Fixed Asset Usage For Rental of an asset which is shipped from and returned to inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update asset_usage_out_in__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("asset_usage_out_in__producttype", c)
	}

	c = cache.Get("service_product__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Service a product using inventory").
		AddChildren(cache.Get("aggregated_service__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update service_product__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("service_product__producttype", c)
	}

	c = cache.Get("aggregated_service__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Service using inventory").
		SetParent(cache.Get("service_product__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update aggregated_service__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("aggregated_service__producttype", c)
	}

	c = cache.Get("aggregatedserv_conf__producttype").(*ent.ProductType)
	c, err = c.Update().
		SetIsPhysical("Yes").
		SetIsDigital("No").
		SetHasTable("No").
		SetDescription("Configurable Service Configuration").
		SetParent(cache.Get("aggregated__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update aggregatedserv_conf__producttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("aggregatedserv_conf__producttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
