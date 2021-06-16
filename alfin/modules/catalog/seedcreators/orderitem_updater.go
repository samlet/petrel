package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderItem(ctx context.Context) error {
	log.Println("OrderItem updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItem
	failures := 0

	c = cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)
	c, err = c.Update().
		SetOrderItemSeqID(common.ParseId("00001")).
		SetOrderItemTypeID(common.ParseId("PRODUCT_ORDER_ITEM")).
		SetProdCatalogID(common.ParseId("DemoCatalog")).
		SetIsPromo("No").
		SetQuantity(6.000000).
		SetSelectedAmount(0.000000).
		SetUnitPrice(38.400).
		SetUnitListPrice(48.000).
		SetIsModifiedPrice("No").
		SetItemDescription("Round Gizmo").
		SetChangeByUserLoginID("admin").
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderAdjustments(cache.Get("81016__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81017__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81020__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81021__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81022__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderItemShipGroupAssocs(cache.Get("demo81015__00001__00001__orderitemshipgroupassoc").(*ent.OrderItemShipGroupAssoc)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9001__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9025__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderStatuses(cache.Get("81016__orderstatus").(*ent.OrderStatus)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__00001__orderitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__00001__orderitem", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
