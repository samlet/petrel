package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductPriceType(ctx context.Context) error {
	log.Println("ProductPriceType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPriceType

	c, err = client.ProductPriceType.Create().SetStringRef("list_price__productpricetype").
		SetDescription("List Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create list_price__productpricetype: %v", err)
		return err
	}
	cache.Put("list_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("default_price__productpricetype").
		SetDescription("Default Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create default_price__productpricetype: %v", err)
		return err
	}
	cache.Put("default_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("average_cost__productpricetype").
		SetDescription("Average Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create average_cost__productpricetype: %v", err)
		return err
	}
	cache.Put("average_cost__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("minimum_price__productpricetype").
		SetDescription("Minimum Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create minimum_price__productpricetype: %v", err)
		return err
	}
	cache.Put("minimum_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("maximum_price__productpricetype").
		SetDescription("Maximum Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create maximum_price__productpricetype: %v", err)
		return err
	}
	cache.Put("maximum_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("promo_price__productpricetype").
		SetDescription("Promotional Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_price__productpricetype: %v", err)
		return err
	}
	cache.Put("promo_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("competitive_price__productpricetype").
		SetDescription("Competitive Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create competitive_price__productpricetype: %v", err)
		return err
	}
	cache.Put("competitive_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("wholesale_price__productpricetype").
		SetDescription("Wholesale Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create wholesale_price__productpricetype: %v", err)
		return err
	}
	cache.Put("wholesale_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("special_promo_price__productpricetype").
		SetDescription("Special Promo Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create special_promo_price__productpricetype: %v", err)
		return err
	}
	cache.Put("special_promo_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("box_price__productpricetype").
		SetDescription("Box Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create box_price__productpricetype: %v", err)
		return err
	}
	cache.Put("box_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("minimum_order_price__productpricetype").
		SetDescription("Minimum Order Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create minimum_order_price__productpricetype: %v", err)
		return err
	}
	cache.Put("minimum_order_price__productpricetype", c)

	c, err = client.ProductPriceType.Create().SetStringRef("shipping_allowance__productpricetype").
		SetDescription("Shipping Allowance Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipping_allowance__productpricetype: %v", err)
		return err
	}
	cache.Put("shipping_allowance__productpricetype", c)

	return nil
}
