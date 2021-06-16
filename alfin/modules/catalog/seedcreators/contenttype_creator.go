package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateContentType(ctx context.Context) error {
	log.Println("ContentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContentType

	c, err = client.ContentType.Create().SetStringRef("image_frame__contenttype").
		SetDescription("Frame Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_frame__contenttype: %v", err)
		return err
	}
	cache.Put("image_frame__contenttype", c)

	return nil
}
