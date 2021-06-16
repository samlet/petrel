package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateSubscriptionType(ctx context.Context) error {
	log.Println("SubscriptionType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SubscriptionType

	c, err = client.SubscriptionType.Create().SetStringRef("product_subscr__subscriptiontype").
		SetHasTable("No").
		SetDescription("Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_subscr__subscriptiontype: %v", err)
		return err
	}
	cache.Put("product_subscr__subscriptiontype", c)

	return nil
}
