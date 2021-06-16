package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductFeature(ctx context.Context) error {
	log.Println("ProductFeature creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeature

	c, err = client.ProductFeature.Create().SetStringRef("test_feature__productfeature").
		SetDescription("Feature for testing purpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_feature__productfeature: %v", err)
		return err
	}
	cache.Put("test_feature__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_avatar__productfeature").
		SetDescription("100 X 75 (avatar)").
		SetDefaultSequenceNum(1).
		SetAbbrev(common.ParseId("100x75")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_avatar__productfeature: %v", err)
		return err
	}
	cache.Put("image_avatar__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_thumbnail__productfeature").
		SetDescription("150 X 112 (thumbnail)").
		SetDefaultSequenceNum(2).
		SetAbbrev(common.ParseId("150x112")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_thumbnail__productfeature: %v", err)
		return err
	}
	cache.Put("image_thumbnail__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_website__productfeature").
		SetDescription("320 X 240 (for websites and email)").
		SetDefaultSequenceNum(3).
		SetAbbrev(common.ParseId("320x240")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_website__productfeature: %v", err)
		return err
	}
	cache.Put("image_website__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_board__productfeature").
		SetDescription("640 X 480 (for message boards)").
		SetDefaultSequenceNum(4).
		SetAbbrev(common.ParseId("640x480")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_board__productfeature: %v", err)
		return err
	}
	cache.Put("image_board__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_monitor15__productfeature").
		SetDescription("800 X 600 (15-inch monitor)").
		SetDefaultSequenceNum(5).
		SetAbbrev(common.ParseId("800x600")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_monitor15__productfeature: %v", err)
		return err
	}
	cache.Put("image_monitor15__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_monitor17__productfeature").
		SetDescription("1024 X 768 (17-inch monitor)").
		SetDefaultSequenceNum(6).
		SetAbbrev(common.ParseId("1024x768")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_monitor17__productfeature: %v", err)
		return err
	}
	cache.Put("image_monitor17__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_monitor19__productfeature").
		SetDescription("1280 X 1024 (19-inch monitor)").
		SetDefaultSequenceNum(7).
		SetAbbrev(common.ParseId("1280x1024")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_monitor19__productfeature: %v", err)
		return err
	}
	cache.Put("image_monitor19__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("image_monitor21__productfeature").
		SetDescription("1600 X 1200 (21-inch monitor)").
		SetDefaultSequenceNum(8).
		SetAbbrev(common.ParseId("1600x1200")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_monitor21__productfeature: %v", err)
		return err
	}
	cache.Put("image_monitor21__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_small__productfeature").
		SetDescription("Small").
		SetDefaultSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_small__productfeature: %v", err)
		return err
	}
	cache.Put("text_small__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_middle__productfeature").
		SetDescription("Middle").
		SetDefaultSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_middle__productfeature: %v", err)
		return err
	}
	cache.Put("text_middle__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_large__productfeature").
		SetDescription("Large").
		SetDefaultSequenceNum(3).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_large__productfeature: %v", err)
		return err
	}
	cache.Put("text_large__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_verylarge__productfeature").
		SetDescription("Very Large").
		SetDefaultSequenceNum(4).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_verylarge__productfeature: %v", err)
		return err
	}
	cache.Put("text_verylarge__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_white__productfeature").
		SetDescription("White").
		SetDefaultSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_white__productfeature: %v", err)
		return err
	}
	cache.Put("text_white__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_gray__productfeature").
		SetDescription("Gray").
		SetDefaultSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_gray__productfeature: %v", err)
		return err
	}
	cache.Put("text_gray__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_black__productfeature").
		SetDescription("Black").
		SetDefaultSequenceNum(3).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_black__productfeature: %v", err)
		return err
	}
	cache.Put("text_black__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_red__productfeature").
		SetDescription("Red").
		SetDefaultSequenceNum(4).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_red__productfeature: %v", err)
		return err
	}
	cache.Put("text_red__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_green__productfeature").
		SetDescription("Green").
		SetDefaultSequenceNum(5).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_green__productfeature: %v", err)
		return err
	}
	cache.Put("text_green__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_blue__productfeature").
		SetDescription("Blue").
		SetDefaultSequenceNum(6).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_blue__productfeature: %v", err)
		return err
	}
	cache.Put("text_blue__productfeature", c)

	c, err = client.ProductFeature.Create().SetStringRef("text_yellow__productfeature").
		SetDescription("Yellow").
		SetDefaultSequenceNum(7).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create text_yellow__productfeature: %v", err)
		return err
	}
	cache.Put("text_yellow__productfeature", c)

	return nil
}
