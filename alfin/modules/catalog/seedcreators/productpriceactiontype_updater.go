package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductPriceActionType(ctx context.Context) error {
	log.Println("ProductPriceActionType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPriceActionType
	failures := 0

	c = cache.Get("price_pol__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Percent Of List Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_pol__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_pol__productpriceactiontype", c)
	}

	c = cache.Get("price_pod__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Percent Of Default Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_pod__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_pod__productpriceactiontype", c)
	}

	c = cache.Get("price_poac__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Percent Of Average Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_poac__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_poac__productpriceactiontype", c)
	}

	c = cache.Get("price_pom__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Percent Of Margin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_pom__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_pom__productpriceactiontype", c)
	}

	c = cache.Get("price_powhs__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Percent Of Wholesale").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_powhs__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_powhs__productpriceactiontype", c)
	}

	c = cache.Get("price_fol__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Flat Amount Modify").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_fol__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_fol__productpriceactiontype", c)
	}

	c = cache.Get("price_flat__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Flat Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_flat__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_flat__productpriceactiontype", c)
	}

	c = cache.Get("price_pflat__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Promo Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_pflat__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_pflat__productpriceactiontype", c)
	}

	c = cache.Get("price_wflat__productpriceactiontype").(*ent.ProductPriceActionType)
	c, err = c.Update().
		SetDescription("Wholesale Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_wflat__productpriceactiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_wflat__productpriceactiontype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
