package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductStoreGroup(ctx context.Context) error {
	log.Println("ProductStoreGroup updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductStoreGroup
	failures := 0

	c = cache.Get("test_group__productstoregroup").(*ent.ProductStoreGroup)
	c, err = c.Update().
		SetProductStoreGroupName("test group").
		AddProductPrices(cache.Get("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_group__productstoregroup: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_group__productstoregroup", c)
	}

	c = cache.Get("_na___productstoregroup").(*ent.ProductStoreGroup)
	c, err = c.Update().
		SetProductStoreGroupName("Not Applicable").
		SetDescription("Not Applicable").
		AddProductPrices(cache.Get("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update _na___productstoregroup: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("_na___productstoregroup", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
