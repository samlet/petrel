package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductFeatureCategory(ctx context.Context) error {
	log.Println("ProductFeatureCategory updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureCategory
	failures := 0

	c = cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)
	c, err = c.Update().
		SetDescription("Image").
		AddProductFeatures(cache.Get("image_avatar__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_thumbnail__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_website__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_board__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor15__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor17__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor19__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor21__productfeature").(*ent.ProductFeature)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image__productfeaturecategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image__productfeaturecategory", c)
	}

	c = cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)
	c, err = c.Update().
		SetDescription("Text").
		AddProductFeatures(cache.Get("text_small__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_middle__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_large__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_verylarge__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_white__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_gray__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_black__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_red__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_green__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_blue__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_yellow__productfeature").(*ent.ProductFeature)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text__productfeaturecategory: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text__productfeaturecategory", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
