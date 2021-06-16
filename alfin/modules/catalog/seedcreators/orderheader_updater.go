package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderHeader(ctx context.Context) error {
	log.Println("OrderHeader updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderHeader
	failures := 0

	c = cache.Get("demo81015__orderheader").(*ent.OrderHeader)
	c, err = c.Update().
		SetOrderTypeID(common.ParseId("SALES_ORDER")).
		SetPriority("Unknown").
		SetEntryDate(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetVisitID(common.ParseId("10043")).
		SetCreatedBy("admin").
		SetCurrencyUom(common.ParseId("USD")).
		SetRemainingSubTotal(79.56).
		SetGrandTotal(105.14).
		AddOrderAdjustments(cache.Get("81015__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81016__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81017__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81018__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81019__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81020__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81021__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderAdjustments(cache.Get("81022__orderadjustment").(*ent.OrderAdjustment)).
		AddOrderContactMeches(cache.Get("demo81015__order_email__9021__ordercontactmech").(*ent.OrderContactMech)).
		AddOrderContactMeches(cache.Get("demo81015__shipping_location__9010__ordercontactmech").(*ent.OrderContactMech)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		AddOrderItemShipGroups(cache.Get("demo81015__00001__orderitemshipgroup").(*ent.OrderItemShipGroup)).
		AddOrderItemShipGroupAssocs(cache.Get("demo81015__00001__00001__orderitemshipgroupassoc").(*ent.OrderItemShipGroupAssoc)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9001__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderItemShipGrpInvRes(cache.Get("demo81015__00001__00001__9025__orderitemshipgrpinvres").(*ent.OrderItemShipGrpInvRes)).
		AddOrderPaymentPreferences(cache.Get("10020__orderpaymentpreference").(*ent.OrderPaymentPreference)).
		AddOrderRoles(cache.Get("demo81015__company__bill_from_vendor__orderrole").(*ent.OrderRole)).
		AddOrderRoles(cache.Get("demo81015__democustomer__bill_to_customer__orderrole").(*ent.OrderRole)).
		AddOrderRoles(cache.Get("demo81015__democustomer__end_user_customer__orderrole").(*ent.OrderRole)).
		AddOrderRoles(cache.Get("demo81015__democustomer__placing_customer__orderrole").(*ent.OrderRole)).
		AddOrderRoles(cache.Get("demo81015__democustomer__ship_to_customer__orderrole").(*ent.OrderRole)).
		AddOrderStatuses(cache.Get("81015__orderstatus").(*ent.OrderStatus)).
		AddOrderStatuses(cache.Get("81016__orderstatus").(*ent.OrderStatus)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__orderheader: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__orderheader", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
