package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductCategory(ctx context.Context) error {
	log.Println("ProductCategory updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategory
	failures := 0

	c = cache.Get("tpcp__productcategory").(*ent.ProductCategory)
	c, err = c.Update().
		SetCategoryName("Test Product Category Parent").
		SetLongDescription("Long Test Product Category Parent Description").
		SetProductCategoryType(cache.Get("test_category__productcategorytype").(*ent.ProductCategoryType)).
		AddPrimaryChildProductCategories(cache.Get("tpc__productcategory").(*ent.ProductCategory)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tpcp__productcategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tpcp__productcategory", c)
	}

	c = cache.Get("tpc__productcategory").(*ent.ProductCategory)
	c, err = c.Update().
		SetCategoryName("Test Product Category").
		SetLongDescription("Long Test Product Category Description").
		SetProductCategoryType(cache.Get("test_category__productcategorytype").(*ent.ProductCategoryType)).
		SetParent(cache.Get("tpcp__productcategory").(*ent.ProductCategory)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tpc__productcategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tpc__productcategory", c)
	}

	c = cache.Get("test_category_a__productcategory").(*ent.ProductCategory)
	c, err = c.Update().
		SetCategoryName("Original_name").
		SetLongDescription("Original description").
		SetProductCategoryType(cache.Get("usage_category__productcategorytype").(*ent.ProductCategoryType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_category_a__productcategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_category_a__productcategory", c)
	}

	c = cache.Get("test_category_b__productcategory").(*ent.ProductCategory)
	c, err = c.Update().
		SetCategoryName("Category_name").
		SetLongDescription("Category description").
		SetProductCategoryType(cache.Get("usage_category__productcategorytype").(*ent.ProductCategoryType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_category_b__productcategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_category_b__productcategory", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
