package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderItemShipGrpInvRes(ctx context.Context) error {
	log.Println("OrderItemShipGrpInvRes updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGrpInvRes
	failures := 0

	c = cache.Get("demo81015__00001__00001__9001__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)
	c, err = c.Update().
		SetShipGroupSeqID(common.ParseId("00001")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetInventoryItemID(common.ParseId("9001")).
		SetReserveOrderEnumID(common.ParseId("INVRO_FIFO_REC")).
		SetQuantity(5.000000).
		SetReservedDatetime(common.ParseDateTime("2011-08-12 23:17:12.839")).
		SetCreatedDatetime(common.ParseDateTime("2011-08-12 23:17:12.839")).
		SetPromisedDatetime(common.ParseDateTime("2011-08-27 23:17:11.507")).
		SetPriority("Unknown").
		SetOrderItem(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__00001__00001__9001__orderitemshipgrpinvres: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__00001__00001__9001__orderitemshipgrpinvres", c)
	}

	c = cache.Get("demo81015__00001__00001__9025__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)
	c, err = c.Update().
		SetShipGroupSeqID(common.ParseId("00001")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetInventoryItemID(common.ParseId("9025")).
		SetReserveOrderEnumID(common.ParseId("INVRO_FIFO_REC")).
		SetQuantity(1.000000).
		SetReservedDatetime(common.ParseDateTime("2011-08-12 23:17:12.967")).
		SetCreatedDatetime(common.ParseDateTime("2011-08-12 23:17:12.967")).
		SetPromisedDatetime(common.ParseDateTime("2011-08-27 23:17:11.507")).
		SetPriority("Unknown").
		SetOrderItem(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__00001__00001__9025__orderitemshipgrpinvres: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__00001__00001__9025__orderitemshipgrpinvres", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
