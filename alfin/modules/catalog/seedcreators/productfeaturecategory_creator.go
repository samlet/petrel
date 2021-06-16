package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductFeatureCategory(ctx context.Context) error {
	log.Println("ProductFeatureCategory creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureCategory

	c, err = client.ProductFeatureCategory.Create().SetStringRef("image__productfeaturecategory").
		SetDescription("Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image__productfeaturecategory: %v", err)
		return err
	}
	cache.Put("image__productfeaturecategory", c)

	c, err = client.ProductFeatureCategory.Create().SetStringRef("text__productfeaturecategory").
		SetDescription("Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text__productfeaturecategory: %v", err)
		return err
	}
	cache.Put("text__productfeaturecategory", c)

	return nil
}
