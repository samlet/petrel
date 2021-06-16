package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductPricePurpose(ctx context.Context) error {
	log.Println("ProductPricePurpose updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPricePurpose
	failures := 0

	c = cache.Get("deposit__productpricepurpose").(*ent.ProductPricePurpose)
	c, err = c.Update().
		SetDescription("Deposit price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update deposit__productpricepurpose: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("deposit__productpricepurpose", c)
	}

	c = cache.Get("purchase__productpricepurpose").(*ent.ProductPricePurpose)
	c, err = c.Update().
		SetDescription("Purchase/Initial").
		AddProductPrices(cache.Get("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purchase__productpricepurpose: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purchase__productpricepurpose", c)
	}

	c = cache.Get("recurring_charge__productpricepurpose").(*ent.ProductPricePurpose)
	c, err = c.Update().
		SetDescription("Recurring Charge").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update recurring_charge__productpricepurpose: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("recurring_charge__productpricepurpose", c)
	}

	c = cache.Get("usage_charge__productpricepurpose").(*ent.ProductPricePurpose)
	c, err = c.Update().
		SetDescription("Usage Charge").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update usage_charge__productpricepurpose: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("usage_charge__productpricepurpose", c)
	}

	c = cache.Get("component_price__productpricepurpose").(*ent.ProductPricePurpose)
	c, err = c.Update().
		SetDescription("Component Price").
		AddProductPrices(cache.Get("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update component_price__productpricepurpose: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("component_price__productpricepurpose", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
