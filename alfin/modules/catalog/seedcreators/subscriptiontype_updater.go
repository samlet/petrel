package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateSubscriptionType(ctx context.Context) error {
	log.Println("SubscriptionType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SubscriptionType
	failures := 0

	c = cache.Get("product_subscr__subscriptiontype").(*ent.SubscriptionType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_subscr__subscriptiontype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_subscr__subscriptiontype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
