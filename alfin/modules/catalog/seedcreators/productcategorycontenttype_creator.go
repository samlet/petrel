package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductCategoryContentType(ctx context.Context) error {
	log.Println("ProductCategoryContentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductCategoryContentType

	c, err = client.ProductCategoryContentType.Create().SetStringRef("category_name__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Category Name").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create category_name__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("category_name__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("description__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create description__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("description__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("long_description__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create long_description__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("long_description__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("alternative_url__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Alternative URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create alternative_url__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("alternative_url__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("category_image__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Category Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create category_image__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("category_image__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("category_image_url__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Category Image URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create category_image_url__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("category_image_url__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("category_image_alt__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Category Image Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create category_image_alt__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("category_image_alt__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("link1_alt_text__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Link 1 Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create link1_alt_text__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("link1_alt_text__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("link2_alt_text__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Link 2 Alt Text").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create link2_alt_text__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("link2_alt_text__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("footer__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Footer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create footer__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("footer__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("page_title__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Page Title").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create page_title__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("page_title__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("meta_keyword__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Meta Keyword").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create meta_keyword__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("meta_keyword__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("meta_description__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Meta Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create meta_description__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("meta_description__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("related_url__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Related URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create related_url__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("related_url__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("video__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Video").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create video__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("video__productcategorycontenttype", c)

	c, err = client.ProductCategoryContentType.Create().SetStringRef("video_url__productcategorycontenttype").
		SetHasTable("No").
		SetDescription("Video URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create video_url__productcategorycontenttype: %v", err)
		return err
	}
	cache.Put("video_url__productcategorycontenttype", c)

	return nil
}
