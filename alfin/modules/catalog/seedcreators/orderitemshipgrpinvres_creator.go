package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderItemShipGrpInvRes(ctx context.Context) error {
	log.Println("OrderItemShipGrpInvRes creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGrpInvRes

	c, err = client.OrderItemShipGrpInvRes.Create().SetStringRef("demo81015__00001__00001__9001__orderitemshipgrpinvres").
		SetShipGroupSeqID(common.ParseId("00001")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetInventoryItemID(common.ParseId("9001")).
		SetReserveOrderEnumID(common.ParseId("INVRO_FIFO_REC")).
		SetQuantity(5.000000).
		SetReservedDatetime(common.ParseDateTime("2011-08-12 23:17:12.839")).
		SetCreatedDatetime(common.ParseDateTime("2011-08-12 23:17:12.839")).
		SetPromisedDatetime(common.ParseDateTime("2011-08-27 23:17:11.507")).
		SetPriority("Unknown").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__00001__00001__9001__orderitemshipgrpinvres: %v", err)
		return err
	}
	cache.Put("demo81015__00001__00001__9001__orderitemshipgrpinvres", c)

	c, err = client.OrderItemShipGrpInvRes.Create().SetStringRef("demo81015__00001__00001__9025__orderitemshipgrpinvres").
		SetShipGroupSeqID(common.ParseId("00001")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetInventoryItemID(common.ParseId("9025")).
		SetReserveOrderEnumID(common.ParseId("INVRO_FIFO_REC")).
		SetQuantity(1.000000).
		SetReservedDatetime(common.ParseDateTime("2011-08-12 23:17:12.967")).
		SetCreatedDatetime(common.ParseDateTime("2011-08-12 23:17:12.967")).
		SetPromisedDatetime(common.ParseDateTime("2011-08-27 23:17:11.507")).
		SetPriority("Unknown").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__00001__00001__9025__orderitemshipgrpinvres: %v", err)
		return err
	}
	cache.Put("demo81015__00001__00001__9025__orderitemshipgrpinvres", c)

	return nil
}
