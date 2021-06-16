package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductStore(ctx context.Context) error {
	log.Println("ProductStore creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductStore

	c, err = client.ProductStore.Create().SetStringRef("test_store__productstore").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_store__productstore: %v", err)
		return err
	}
	cache.Put("test_store__productstore", c)

	return nil
}
