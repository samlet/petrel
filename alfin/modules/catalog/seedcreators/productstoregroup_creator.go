package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductStoreGroup(ctx context.Context) error {
	log.Println("ProductStoreGroup creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductStoreGroup

	c, err = client.ProductStoreGroup.Create().SetStringRef("test_group__productstoregroup").
		SetProductStoreGroupName("test group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_group__productstoregroup: %v", err)
		return err
	}
	cache.Put("test_group__productstoregroup", c)

	c, err = client.ProductStoreGroup.Create().SetStringRef("_na___productstoregroup").
		SetProductStoreGroupName("Not Applicable").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___productstoregroup: %v", err)
		return err
	}
	cache.Put("_na___productstoregroup", c)

	return nil
}
