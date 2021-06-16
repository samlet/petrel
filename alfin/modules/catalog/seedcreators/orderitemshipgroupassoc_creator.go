package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderItemShipGroupAssoc(ctx context.Context) error {
	log.Println("OrderItemShipGroupAssoc creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGroupAssoc

	c, err = client.OrderItemShipGroupAssoc.Create().SetStringRef("demo81015__00001__00001__orderitemshipgroupassoc").
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetQuantity(6.000000).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__00001__00001__orderitemshipgroupassoc: %v", err)
		return err
	}
	cache.Put("demo81015__00001__00001__orderitemshipgroupassoc", c)

	return nil
}
