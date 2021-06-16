package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductCategoryType(ctx context.Context) error {
	log.Println("ProductCategoryType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategoryType

	c, err = client.ProductCategoryType.Create().SetStringRef("test_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Test Category").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("test_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("catalog_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Catalog").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create catalog_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("catalog_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("industry_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Industry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create industry_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("industry_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("internal_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Internal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create internal_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("internal_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("materials_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Materials").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create materials_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("materials_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("quickadd_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Quick Add").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create quickadd_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("quickadd_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("search_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Search").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create search_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("search_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("usage_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Usage").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create usage_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("usage_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("mixmatch_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Mix and Match").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create mixmatch_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("mixmatch_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("cross_sell_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Cross Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cross_sell_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("cross_sell_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("tax_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Tax").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tax_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("tax_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("google_base_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Google Base").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create google_base_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("google_base_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("gift_card_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Gift Cards").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create gift_card_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("gift_card_category__productcategorytype", c)

	c, err = client.ProductCategoryType.Create().SetStringRef("best_sell_category__productcategorytype").
		SetHasTable("No").
		SetDescription("Best Selling").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create best_sell_category__productcategorytype: %v", err)
		return err
	}
	cache.Put("best_sell_category__productcategorytype", c)

	return nil
}
