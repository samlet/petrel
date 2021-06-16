package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderItemShipGroupAssoc(ctx context.Context) error {
	log.Println("OrderItemShipGroupAssoc updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGroupAssoc
	failures := 0

	c = cache.Get("demo81015__00001__00001__orderitemshipgroupassoc").(*ent.OrderItemShipGroupAssoc)
	c, err = c.Update().
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetQuantity(6.000000).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		SetOrderItem(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		SetOrderItemShipGroup(cache.Get("demo81015__00001__orderitemshipgroup").(*ent.OrderItemShipGroup)).
		AddOrderAdjustments(cache.Get("81020__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81021__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81022__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9001__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9025__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__00001__00001__orderitemshipgroupassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__00001__00001__orderitemshipgroupassoc", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
