package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductPriceActionType(ctx context.Context) error {
	log.Println("ProductPriceActionType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPriceActionType

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_pol__productpriceactiontype").
		SetDescription("Percent Of List Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_pol__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_pol__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_pod__productpriceactiontype").
		SetDescription("Percent Of Default Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_pod__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_pod__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_poac__productpriceactiontype").
		SetDescription("Percent Of Average Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_poac__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_poac__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_pom__productpriceactiontype").
		SetDescription("Percent Of Margin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_pom__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_pom__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_powhs__productpriceactiontype").
		SetDescription("Percent Of Wholesale").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_powhs__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_powhs__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_fol__productpriceactiontype").
		SetDescription("Flat Amount Modify").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_fol__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_fol__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_flat__productpriceactiontype").
		SetDescription("Flat Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_flat__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_flat__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_pflat__productpriceactiontype").
		SetDescription("Promo Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_pflat__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_pflat__productpriceactiontype", c)

	c, err = client.ProductPriceActionType.Create().SetStringRef("price_wflat__productpriceactiontype").
		SetDescription("Wholesale Amount Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_wflat__productpriceactiontype: %v", err)
		return err
	}
	cache.Put("price_wflat__productpriceactiontype", c)

	return nil
}
