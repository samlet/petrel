package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductCategoryType(ctx context.Context) error {
	log.Println("ProductCategoryType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategoryType
	failures := 0

	c = cache.Get("test_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Test Category").
		AddProductCategories(cache.Get("tpcp__productcategory").(*ent.ProductCategory)).
		AddProductCategories(cache.Get("tpc__productcategory").(*ent.ProductCategory)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_category__productcategorytype", c)
	}

	c = cache.Get("catalog_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Catalog").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update catalog_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("catalog_category__productcategorytype", c)
	}

	c = cache.Get("industry_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Industry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update industry_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("industry_category__productcategorytype", c)
	}

	c = cache.Get("internal_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Internal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update internal_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("internal_category__productcategorytype", c)
	}

	c = cache.Get("materials_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Materials").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update materials_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("materials_category__productcategorytype", c)
	}

	c = cache.Get("quickadd_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Quick Add").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update quickadd_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("quickadd_category__productcategorytype", c)
	}

	c = cache.Get("search_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Search").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update search_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("search_category__productcategorytype", c)
	}

	c = cache.Get("usage_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Usage").
		AddProductCategories(cache.Get("test_category_a__productcategory").(*ent.ProductCategory)).
		AddProductCategories(cache.Get("test_category_b__productcategory").(*ent.ProductCategory)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update usage_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("usage_category__productcategorytype", c)
	}

	c = cache.Get("mixmatch_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Mix and Match").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update mixmatch_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("mixmatch_category__productcategorytype", c)
	}

	c = cache.Get("cross_sell_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Cross Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cross_sell_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cross_sell_category__productcategorytype", c)
	}

	c = cache.Get("tax_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Tax").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tax_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tax_category__productcategorytype", c)
	}

	c = cache.Get("google_base_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Google Base").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update google_base_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("google_base_category__productcategorytype", c)
	}

	c = cache.Get("gift_card_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Gift Cards").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update gift_card_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("gift_card_category__productcategorytype", c)
	}

	c = cache.Get("best_sell_category__productcategorytype").(*ent.ProductCategoryType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Best Selling").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update best_sell_category__productcategorytype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("best_sell_category__productcategorytype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
