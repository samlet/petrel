package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductStore(ctx context.Context) error {
	log.Println("ProductStore updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductStore
	failures := 0

	c = cache.Get("test_store__productstore").(*ent.ProductStore)
	c, err = c.Update().
		AddProductReviews(cache.Get("test_review__productreview").(*ent.ProductReview)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_store__productstore: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_store__productstore", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
