package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductAssoc(ctx context.Context) error {
	log.Println("ProductAssoc creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductAssoc

	c, err = client.ProductAssoc.Create().SetStringRef("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct__demoproduct-2__product_variant__1147521600__productassoc: %v", err)
		return err
	}
	cache.Put("demoproduct__demoproduct-2__product_variant__1147521600__productassoc", c)

	c, err = client.ProductAssoc.Create().SetStringRef("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct__demoproduct-2__product_variant__1147521600__productassoc: %v", err)
		return err
	}
	cache.Put("demoproduct__demoproduct-2__product_variant__1147521600__productassoc", c)

	c, err = client.ProductAssoc.Create().SetStringRef("demoproduct__demoproduct-3__product_variant__1147521600__productassoc").
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetQuantity(1.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct__demoproduct-3__product_variant__1147521600__productassoc: %v", err)
		return err
	}
	cache.Put("demoproduct__demoproduct-3__product_variant__1147521600__productassoc", c)

	return nil
}
