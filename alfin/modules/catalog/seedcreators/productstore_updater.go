package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func UpdateProductStore(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductStore

	c = cache.Get("test_store__productstore").(*ent.ProductStore)
	c, err = c.Update().
		AddProductReviews(cache.Get("test_review__productreview").(*ent.ProductReview)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_store__productstore: %v", err)
		return err
	}
	cache.Put("test_store__productstore", c)

	return nil
}
