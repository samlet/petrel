package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProduct(ctx context.Context) error {
	log.Println("Product updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Product
	failures := 0

	c = cache.Get("demoproduct__product").(*ent.Product)
	c, err = c.Update().
		SetInternalName("Test Product").
		SetProductName("Demo Product").
		SetIsVirtual("Yes").
		SetIsVariant("No").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		SetProductType(cache.Get("finished_good__producttype").(*ent.ProductType)).
		AddMainProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddMainProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddMainProductAssocs(cache.Get("demoproduct__demoproduct-3__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct__product", c)
	}

	c = cache.Get("demoproduct-1__product").(*ent.Product)
	c, err = c.Update().
		SetProductName("Demo Product 1").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		SetProductType(cache.Get("finished_good__producttype").(*ent.ProductType)).
		AddProductPrices(cache.Get("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-1__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-1__product", c)
	}

	c = cache.Get("demoproduct-2__product").(*ent.Product)
	c, err = c.Update().
		SetSalesDiscontinuationDate(common.ParseDateTime("2007-03-23 23:05:32.915")).
		SetProductName("Demo Product 2").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		SetProductType(cache.Get("finished_good__producttype").(*ent.ProductType)).
		AddAssocProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddAssocProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddProductPrices(cache.Get("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-2__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-2__product", c)
	}

	c = cache.Get("demoproduct-3__product").(*ent.Product)
	c, err = c.Update().
		SetProductName("Demo Product 3").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		SetProductType(cache.Get("finished_good__producttype").(*ent.ProductType)).
		AddAssocProductAssocs(cache.Get("demoproduct__demoproduct-3__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddProductPrices(cache.Get("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct-3__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct-3__product", c)
	}

	c = cache.Get("test_product_a__product").(*ent.Product)
	c, err = c.Update().
		SetProductName("Test_name_A").
		SetDescription("This is original description").
		SetProductType(cache.Get("test_type__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_product_a__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_product_a__product", c)
	}

	c = cache.Get("test_product_b__product").(*ent.Product)
	c, err = c.Update().
		SetProductName("Test_name_C").
		SetDescription("This is product description").
		SetProductType(cache.Get("test_type__producttype").(*ent.ProductType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_product_b__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_product_b__product", c)
	}

	c = cache.Get("test_product_c__product").(*ent.Product)
	c, err = c.Update().
		SetDescription("description").
		SetProductType(cache.Get("test_type__producttype").(*ent.ProductType)).
		AddProductPrices(cache.Get("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").(*ent.ProductPrice)).
		AddProductReviews(cache.Get("test_review__productreview").(*ent.ProductReview)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_product_c__product: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_product_c__product", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
