package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductFeature(ctx context.Context) error {
	log.Println("ProductFeature updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeature
	failures := 0

	c = cache.Get("test_feature__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Feature for testing purpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_feature__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_feature__productfeature", c)
	}

	c = cache.Get("image_avatar__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("100 X 75 (avatar)").
		SetDefaultSequenceNum(1).
		SetAbbrev(common.ParseId("100x75")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_avatar__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_avatar__productfeature", c)
	}

	c = cache.Get("image_thumbnail__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("150 X 112 (thumbnail)").
		SetDefaultSequenceNum(2).
		SetAbbrev(common.ParseId("150x112")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_thumbnail__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_thumbnail__productfeature", c)
	}

	c = cache.Get("image_website__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("320 X 240 (for websites and email)").
		SetDefaultSequenceNum(3).
		SetAbbrev(common.ParseId("320x240")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_website__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_website__productfeature", c)
	}

	c = cache.Get("image_board__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("640 X 480 (for message boards)").
		SetDefaultSequenceNum(4).
		SetAbbrev(common.ParseId("640x480")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_board__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_board__productfeature", c)
	}

	c = cache.Get("image_monitor15__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("800 X 600 (15-inch monitor)").
		SetDefaultSequenceNum(5).
		SetAbbrev(common.ParseId("800x600")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_monitor15__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_monitor15__productfeature", c)
	}

	c = cache.Get("image_monitor17__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("1024 X 768 (17-inch monitor)").
		SetDefaultSequenceNum(6).
		SetAbbrev(common.ParseId("1024x768")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_monitor17__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_monitor17__productfeature", c)
	}

	c = cache.Get("image_monitor19__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("1280 X 1024 (19-inch monitor)").
		SetDefaultSequenceNum(7).
		SetAbbrev(common.ParseId("1280x1024")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_monitor19__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_monitor19__productfeature", c)
	}

	c = cache.Get("image_monitor21__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("1600 X 1200 (21-inch monitor)").
		SetDefaultSequenceNum(8).
		SetAbbrev(common.ParseId("1600x1200")).
		SetProductFeatureCategory(cache.Get("image__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_monitor21__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_monitor21__productfeature", c)
	}

	c = cache.Get("text_small__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Small").
		SetDefaultSequenceNum(1).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_small__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_small__productfeature", c)
	}

	c = cache.Get("text_middle__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Middle").
		SetDefaultSequenceNum(2).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_middle__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_middle__productfeature", c)
	}

	c = cache.Get("text_large__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Large").
		SetDefaultSequenceNum(3).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_large__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_large__productfeature", c)
	}

	c = cache.Get("text_verylarge__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Very Large").
		SetDefaultSequenceNum(4).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_verylarge__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_verylarge__productfeature", c)
	}

	c = cache.Get("text_white__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("White").
		SetDefaultSequenceNum(1).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_white__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_white__productfeature", c)
	}

	c = cache.Get("text_gray__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Gray").
		SetDefaultSequenceNum(2).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_gray__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_gray__productfeature", c)
	}

	c = cache.Get("text_black__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Black").
		SetDefaultSequenceNum(3).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_black__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_black__productfeature", c)
	}

	c = cache.Get("text_red__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Red").
		SetDefaultSequenceNum(4).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_red__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_red__productfeature", c)
	}

	c = cache.Get("text_green__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Green").
		SetDefaultSequenceNum(5).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_green__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_green__productfeature", c)
	}

	c = cache.Get("text_blue__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Blue").
		SetDefaultSequenceNum(6).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_blue__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_blue__productfeature", c)
	}

	c = cache.Get("text_yellow__productfeature").(*ent.ProductFeature)
	c, err = c.Update().
		SetDescription("Yellow").
		SetDefaultSequenceNum(7).
		SetProductFeatureCategory(cache.Get("text__productfeaturecategory").(*ent.ProductFeatureCategory)).
		SetProductFeatureType(cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update text_yellow__productfeature: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("text_yellow__productfeature", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
