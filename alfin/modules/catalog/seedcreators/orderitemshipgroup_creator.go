package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderItemShipGroup(ctx context.Context) error {
	log.Println("OrderItemShipGroup creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGroup

	c, err = client.OrderItemShipGroup.Create().SetStringRef("demo81015__00001__orderitemshipgroup").
		SetShipGroupSeqID(common.ParseId("00001")).
		SetShipmentMethodTypeID(common.ParseId("NEXT_DAY")).
		SetCarrierPartyID(common.ParseId("UPS")).
		SetCarrierRoleTypeID(common.ParseId("CARRIER")).
		SetContactMechID(common.ParseId("9010")).
		SetMaySplit("No").
		SetIsGift("No").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__00001__orderitemshipgroup: %v", err)
		return err
	}
	cache.Put("demo81015__00001__orderitemshipgroup", c)

	return nil
}
