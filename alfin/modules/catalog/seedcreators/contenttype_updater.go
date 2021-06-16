package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateContentType(ctx context.Context) error {
	log.Println("ContentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContentType
	failures := 0

	c = cache.Get("image_frame__contenttype").(*ent.ContentType)
	c, err = c.Update().
		SetDescription("Frame Image").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_frame__contenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_frame__contenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
