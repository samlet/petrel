package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductCategory(ctx context.Context) error {
	log.Println("ProductCategory creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategory

	c, err = client.ProductCategory.Create().SetStringRef("tpcp__productcategory").
		SetCategoryName("Test Product Category Parent").
		SetLongDescription("Long Test Product Category Parent Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tpcp__productcategory: %v", err)
		return err
	}
	cache.Put("tpcp__productcategory", c)

	c, err = client.ProductCategory.Create().SetStringRef("tpc__productcategory").
		SetCategoryName("Test Product Category").
		SetLongDescription("Long Test Product Category Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tpc__productcategory: %v", err)
		return err
	}
	cache.Put("tpc__productcategory", c)

	c, err = client.ProductCategory.Create().SetStringRef("test_category_a__productcategory").
		SetCategoryName("Original_name").
		SetLongDescription("Original description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_category_a__productcategory: %v", err)
		return err
	}
	cache.Put("test_category_a__productcategory", c)

	c, err = client.ProductCategory.Create().SetStringRef("test_category_b__productcategory").
		SetCategoryName("Category_name").
		SetLongDescription("Category description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_category_b__productcategory: %v", err)
		return err
	}
	cache.Put("test_category_b__productcategory", c)

	return nil
}
