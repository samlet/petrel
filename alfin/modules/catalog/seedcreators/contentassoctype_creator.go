package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateContentAssocType(ctx context.Context) error {
	log.Println("ContentAssocType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContentAssocType

	c, err = client.ContentAssocType.Create().SetStringRef("image_thumbnail__contentassoctype").
		SetDescription("Image Thumbnail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_thumbnail__contentassoctype: %v", err)
		return err
	}
	cache.Put("image_thumbnail__contentassoctype", c)

	return nil
}
