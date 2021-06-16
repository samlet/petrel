package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProdConfItemContentType(ctx context.Context) error {
	log.Println("ProdConfItemContentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProdConfItemContentType
	failures := 0

	c = cache.Get("image_url__prodconfitemcontenttype").(*ent.ProdConfItemContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_url__prodconfitemcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_url__prodconfitemcontenttype", c)
	}

	c = cache.Get("description__prodconfitemcontenttype").(*ent.ProdConfItemContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update description__prodconfitemcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("description__prodconfitemcontenttype", c)
	}

	c = cache.Get("long_description__prodconfitemcontenttype").(*ent.ProdConfItemContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Description - Long").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update long_description__prodconfitemcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("long_description__prodconfitemcontenttype", c)
	}

	c = cache.Get("instructions__prodconfitemcontenttype").(*ent.ProdConfItemContentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Instructions").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update instructions__prodconfitemcontenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("instructions__prodconfitemcontenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
