package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductAssoc(ctx context.Context) error {
	log.Println("ProductAssoc updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductAssoc
	failures := 0

	c = cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		SetProductAssocType(cache.Get("product_variant__productassoctype").(*ent.ProductAssocType)).
		SetMainProduct(cache.Get("demoproduct__product").(*ent.Product)).
		SetAssocProduct(cache.Get("demoproduct-2__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct__demoproduct-2__product_variant__1147521600__productassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct__demoproduct-2__product_variant__1147521600__productassoc", c)
	}

	c = cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		SetProductAssocType(cache.Get("product_variant__productassoctype").(*ent.ProductAssocType)).
		SetMainProduct(cache.Get("demoproduct__product").(*ent.Product)).
		SetAssocProduct(cache.Get("demoproduct-2__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct__demoproduct-2__product_variant__1147521600__productassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct__demoproduct-2__product_variant__1147521600__productassoc", c)
	}

	c = cache.Get("demoproduct__demoproduct-3__product_variant__1147521600__productassoc").(*ent.ProductAssoc)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		SetProductAssocType(cache.Get("product_variant__productassoctype").(*ent.ProductAssocType)).
		SetMainProduct(cache.Get("demoproduct__product").(*ent.Product)).
		SetAssocProduct(cache.Get("demoproduct-3__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoproduct__demoproduct-3__product_variant__1147521600__productassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoproduct__demoproduct-3__product_variant__1147521600__productassoc", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
