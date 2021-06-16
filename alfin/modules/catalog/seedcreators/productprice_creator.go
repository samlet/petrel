package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductPrice(ctx context.Context) error {
	log.Println("ProductPrice creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductPrice

	c, err = client.ProductPrice.Create().SetStringRef("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice").
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(15.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-1__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		return err
	}
	cache.Put("demoproduct-1__default_price__purchase__usd___na___1147521600__productprice", c)

	c, err = client.ProductPrice.Create().SetStringRef("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice").
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(5.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-2__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		return err
	}
	cache.Put("demoproduct-2__default_price__purchase__usd___na___1147521600__productprice", c)

	c, err = client.ProductPrice.Create().SetStringRef("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice").
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetPrice(10.00).
		SetCreatedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		SetLastModifiedDate(common.ParseDateTime("2006-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoproduct-3__default_price__purchase__usd___na___1147521600__productprice: %v", err)
		return err
	}
	cache.Put("demoproduct-3__default_price__purchase__usd___na___1147521600__productprice", c)

	c, err = client.ProductPrice.Create().SetStringRef("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice").
		SetCurrencyUomID(common.ParseId("USD")).
		SetFromDate(common.ParseDateTime("2013-07-04 00:00:00")).
		SetPrice(20).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice: %v", err)
		return err
	}
	cache.Put("test_product_c__average_cost__component_price__usd__test_group__1372896000__productprice", c)

	return nil
}
