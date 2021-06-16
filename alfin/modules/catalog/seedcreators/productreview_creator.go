package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductReview(ctx context.Context) error {
	log.Println("ProductReview creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductReview

	c, err = client.ProductReview.Create().SetStringRef("test_review__productreview").
		SetProductRating(1.000000).
		SetProductReview("Original review").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create test_review__productreview: %v", err)
		return err
	}
	cache.Put("test_review__productreview", c)

	return nil
}
