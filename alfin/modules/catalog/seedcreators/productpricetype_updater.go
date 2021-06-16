package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductPriceType(ctx context.Context) error {
	log.Println("ProductPriceType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPriceType
	failures := 0

	c = cache.Get("list_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("List Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update list_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("list_price__productpricetype", c)
	}

	c = cache.Get("default_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Default Price").
		AddProductPrices(cache.Get("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		AddProductPrices(cache.Get("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update default_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("default_price__productpricetype", c)
	}

	c = cache.Get("average_cost__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Average Cost").
		AddProductPrices(cache.Get("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").(*ent.ProductPrice)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update average_cost__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("average_cost__productpricetype", c)
	}

	c = cache.Get("minimum_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Minimum Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update minimum_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("minimum_price__productpricetype", c)
	}

	c = cache.Get("maximum_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Maximum Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update maximum_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("maximum_price__productpricetype", c)
	}

	c = cache.Get("promo_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Promotional Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_price__productpricetype", c)
	}

	c = cache.Get("competitive_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Competitive Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update competitive_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("competitive_price__productpricetype", c)
	}

	c = cache.Get("wholesale_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Wholesale Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update wholesale_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("wholesale_price__productpricetype", c)
	}

	c = cache.Get("special_promo_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Special Promo Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update special_promo_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("special_promo_price__productpricetype", c)
	}

	c = cache.Get("box_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Box Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update box_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("box_price__productpricetype", c)
	}

	c = cache.Get("minimum_order_price__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Minimum Order Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update minimum_order_price__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("minimum_order_price__productpricetype", c)
	}

	c = cache.Get("shipping_allowance__productpricetype").(*ent.ProductPriceType)
	c, err = c.Update().
		SetDescription("Shipping Allowance Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipping_allowance__productpricetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipping_allowance__productpricetype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
