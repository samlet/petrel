package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductConfigItem(ctx context.Context) error {
	log.Println("ProductConfigItem updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductConfigItem
	failures := 0

	c = cache.Get("testconfigitemid__productconfigitem").(*ent.ProductConfigItem)
	c, err = c.Update().
		SetConfigItemTypeID(common.ParseId("SINGLE")).
		SetConfigItemName("Test Config Item Name").
		SetDescription("Test Desc").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update testconfigitemid__productconfigitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("testconfigitemid__productconfigitem", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
