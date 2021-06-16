package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductPricePurpose(ctx context.Context) error {
	log.Println("ProductPricePurpose creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPricePurpose

	c, err = client.ProductPricePurpose.Create().SetStringRef("deposit__productpricepurpose").
		SetDescription("Deposit price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create deposit__productpricepurpose: %v", err)
		return err
	}
	cache.Put("deposit__productpricepurpose", c)

	c, err = client.ProductPricePurpose.Create().SetStringRef("purchase__productpricepurpose").
		SetDescription("Purchase/Initial").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purchase__productpricepurpose: %v", err)
		return err
	}
	cache.Put("purchase__productpricepurpose", c)

	c, err = client.ProductPricePurpose.Create().SetStringRef("recurring_charge__productpricepurpose").
		SetDescription("Recurring Charge").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create recurring_charge__productpricepurpose: %v", err)
		return err
	}
	cache.Put("recurring_charge__productpricepurpose", c)

	c, err = client.ProductPricePurpose.Create().SetStringRef("usage_charge__productpricepurpose").
		SetDescription("Usage Charge").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create usage_charge__productpricepurpose: %v", err)
		return err
	}
	cache.Put("usage_charge__productpricepurpose", c)

	c, err = client.ProductPricePurpose.Create().SetStringRef("component_price__productpricepurpose").
		SetDescription("Component Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create component_price__productpricepurpose: %v", err)
		return err
	}
	cache.Put("component_price__productpricepurpose", c)

	return nil
}
