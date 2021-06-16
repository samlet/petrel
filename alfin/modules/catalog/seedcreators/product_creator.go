package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProduct(ctx context.Context) error {
	log.Println("Product creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Product

	c, err = client.Product.Create().SetStringRef("demoproduct__product").
		SetInternalName("Test Product").
		SetProductName("Demo Product").
		SetIsVirtual("Yes").
		SetIsVariant("No").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct__product: %v", err)
		return err
	}
	cache.Put("demoproduct__product", c)

	c, err = client.Product.Create().SetStringRef("demoproduct-1__product").
		SetProductName("Demo Product 1").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-1__product: %v", err)
		return err
	}
	cache.Put("demoproduct-1__product", c)

	c, err = client.Product.Create().SetStringRef("demoproduct-2__product").
		SetSalesDiscontinuationDate(common.ParseDateTime("2007-03-23 23:05:32.915")).
		SetProductName("Demo Product 2").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-2__product: %v", err)
		return err
	}
	cache.Put("demoproduct-2__product", c)

	c, err = client.Product.Create().SetStringRef("demoproduct-3__product").
		SetProductName("Demo Product 3").
		SetIsVirtual("No").
		SetIsVariant("Yes").
		SetCreatedDate(common.ParseDateTime("2006-03-23 23:05:32.915")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-3__product: %v", err)
		return err
	}
	cache.Put("demoproduct-3__product", c)

	c, err = client.Product.Create().SetStringRef("test_product_a__product").
		SetProductName("Test_name_A").
		SetDescription("This is original description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_product_a__product: %v", err)
		return err
	}
	cache.Put("test_product_a__product", c)

	c, err = client.Product.Create().SetStringRef("test_product_b__product").
		SetProductName("Test_name_C").
		SetDescription("This is product description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_product_b__product: %v", err)
		return err
	}
	cache.Put("test_product_b__product", c)

	c, err = client.Product.Create().SetStringRef("test_product_c__product").
		SetDescription("description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_product_c__product: %v", err)
		return err
	}
	cache.Put("test_product_c__product", c)

	return nil
}
