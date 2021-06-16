package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateContentAssocType(ctx context.Context) error {
	log.Println("ContentAssocType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContentAssocType
	failures := 0

	c = cache.Get("image_thumbnail__contentassoctype").(*ent.ContentAssocType)
	c, err = c.Update().
		SetDescription("Image Thumbnail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_thumbnail__contentassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_thumbnail__contentassoctype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
