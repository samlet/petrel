package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderItemShipGroup(ctx context.Context) error {
	log.Println("OrderItemShipGroup updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItemShipGroup
	failures := 0

	c = cache.Get("demo81015__00001__orderitemshipgroup").(*ent.OrderItemShipGroup)
	c, err = c.Update().
		SetShipGroupSeqID(common.ParseId("00001")).
		SetShipmentMethodTypeID(common.ParseId("NEXT_DAY")).
		SetCarrierPartyID(common.ParseId("UPS")).
		SetCarrierRoleTypeID(common.ParseId("CARRIER")).
		SetContactMechID(common.ParseId("9010")).
		SetMaySplit("No").
		SetIsGift("No").
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderAdjustments(cache.Get("81018__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81019__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81020__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81021__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81022__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderItemShipGroupAssocs(cache.Get("demo81015__00001__00001__orderitemshipgroupassoc").(*ent.OrderItemShipGroupAssoc)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9001__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9025__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__00001__orderitemshipgroup: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__00001__orderitemshipgroup", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
