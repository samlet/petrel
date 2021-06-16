package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProdConfItemContentType(ctx context.Context) error {
	log.Println("ProdConfItemContentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProdConfItemContentType

	c, err = client.ProdConfItemContentType.Create().SetStringRef("image_url__prodconfitemcontenttype").
		SetHasTable("No").
		SetDescription("Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_url__prodconfitemcontenttype: %v", err)
		return err
	}
	cache.Put("image_url__prodconfitemcontenttype", c)

	c, err = client.ProdConfItemContentType.Create().SetStringRef("description__prodconfitemcontenttype").
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create description__prodconfitemcontenttype: %v", err)
		return err
	}
	cache.Put("description__prodconfitemcontenttype", c)

	c, err = client.ProdConfItemContentType.Create().SetStringRef("long_description__prodconfitemcontenttype").
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create long_description__prodconfitemcontenttype: %v", err)
		return err
	}
	cache.Put("long_description__prodconfitemcontenttype", c)

	c, err = client.ProdConfItemContentType.Create().SetStringRef("instructions__prodconfitemcontenttype").
		SetHasTable("No").
		SetDescription("Instructions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create instructions__prodconfitemcontenttype: %v", err)
		return err
	}
	cache.Put("instructions__prodconfitemcontenttype", c)

	return nil
}
