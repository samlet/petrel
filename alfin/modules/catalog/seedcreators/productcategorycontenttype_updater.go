package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductCategoryContentType(ctx context.Context) error {
	log.Println("ProductCategoryContentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategoryContentType
	failures := 0

	c = cache.Get("category_name__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Category Name").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update category_name__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("category_name__productcategorycontenttype", c)
	}

	c = cache.Get("description__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update description__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("description__productcategorycontenttype", c)
	}

	c = cache.Get("long_description__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update long_description__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("long_description__productcategorycontenttype", c)
	}

	c = cache.Get("alternative_url__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Alternative URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update alternative_url__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("alternative_url__productcategorycontenttype", c)
	}

	c = cache.Get("category_image__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Category Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update category_image__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("category_image__productcategorycontenttype", c)
	}

	c = cache.Get("category_image_url__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Category Image URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update category_image_url__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("category_image_url__productcategorycontenttype", c)
	}

	c = cache.Get("category_image_alt__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Category Image Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update category_image_alt__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("category_image_alt__productcategorycontenttype", c)
	}

	c = cache.Get("link1_alt_text__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Link 1 Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update link1_alt_text__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("link1_alt_text__productcategorycontenttype", c)
	}

	c = cache.Get("link2_alt_text__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Link 2 Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update link2_alt_text__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("link2_alt_text__productcategorycontenttype", c)
	}

	c = cache.Get("footer__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Footer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update footer__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("footer__productcategorycontenttype", c)
	}

	c = cache.Get("page_title__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Page Title").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update page_title__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("page_title__productcategorycontenttype", c)
	}

	c = cache.Get("meta_keyword__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meta Keyword").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update meta_keyword__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("meta_keyword__productcategorycontenttype", c)
	}

	c = cache.Get("meta_description__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meta Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update meta_description__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("meta_description__productcategorycontenttype", c)
	}

	c = cache.Get("related_url__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Related URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update related_url__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("related_url__productcategorycontenttype", c)
	}

	c = cache.Get("video__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Video").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update video__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("video__productcategorycontenttype", c)
	}

	c = cache.Get("video_url__productcategorycontenttype").(*ent.ProductCategoryContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Video URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update video_url__productcategorycontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("video_url__productcategorycontenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
