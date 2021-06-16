package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductReview(ctx context.Context) error {
	log.Println("ProductReview updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductReview
	failures := 0

	c = cache.Get("test_review__productreview").(*ent.ProductReview)
	c, err = c.Update().
		SetProductRating(1.000000).
		SetProductReview("Original review").
		SetProductStore(cache.Get("test_store__productstore").(*ent.ProductStore)).
		SetProduct(cache.Get("test_product_c__product").(*ent.Product)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update test_review__productreview: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("test_review__productreview", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
