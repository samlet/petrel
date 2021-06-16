package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductConfigItem(ctx context.Context) error {
	log.Println("ProductConfigItem creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductConfigItem

	c, err = client.ProductConfigItem.Create().SetStringRef("testconfigitemid__productconfigitem").
		SetConfigItemTypeID(common.ParseId("SINGLE")).
		SetConfigItemName("Test Config Item Name").
		SetDescription("Test Desc").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create testconfigitemid__productconfigitem: %v", err)
		return err
	}
	cache.Put("testconfigitemid__productconfigitem", c)

	return nil
}
