package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductPrice(ctx context.Context) error {
	log.Println("ProductPrice updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPrice
	failures := 0

	c = cache.Get("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)
	c, err = c.Update().
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(15.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetProduct(cache.Get("demoproduct-1__product").(*ent.Product)).
		SetProductPriceType(cache.Get("default_price__productpricetype").(*ent.ProductPriceType)).
		SetProductPricePurpose(cache.Get("purchase__productpricepurpose").(*ent.ProductPricePurpose)).
		SetProductStoreGroup(cache.Get("_na___productstoregroup").(*ent.ProductStoreGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-1__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice", c)
	}

	c = cache.Get("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)
	c, err = c.Update().
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(5.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetProduct(cache.Get("demoproduct-2__product").(*ent.Product)).
		SetProductPriceType(cache.Get("default_price__productpricetype").(*ent.ProductPriceType)).
		SetProductPricePurpose(cache.Get("purchase__productpricepurpose").(*ent.ProductPricePurpose)).
		SetProductStoreGroup(cache.Get("_na___productstoregroup").(*ent.ProductStoreGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-2__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice", c)
	}

	c = cache.Get("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)
	c, err = c.Update().
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(10.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetProduct(cache.Get("demoproduct-3__product").(*ent.Product)).
		SetProductPriceType(cache.Get("default_price__productpricetype").(*ent.ProductPriceType)).
		SetProductPricePurpose(cache.Get("purchase__productpricepurpose").(*ent.ProductPricePurpose)).
		SetProductStoreGroup(cache.Get("_na___productstoregroup").(*ent.ProductStoreGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-3__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice", c)
	}

	c = cache.Get("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").(*ent.ProductPrice)
	c, err = c.Update().
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2013-07-04 00:00:00")).
		SetPrice(20).
		SetProduct(cache.Get("test_product_c__product").(*ent.Product)).
		SetProductPriceType(cache.Get("average_cost__productpricetype").(*ent.ProductPriceType)).
		SetProductPricePurpose(cache.Get("component_price__productpricepurpose").(*ent.ProductPricePurpose)).
		SetProductStoreGroup(cache.Get("test_group__productstoregroup").(*ent.ProductStoreGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
