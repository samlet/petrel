package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderContactMech(ctx context.Context) error {
	log.Println("OrderContactMech updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderContactMech
	failures := 0

	c = cache.Get("demo81015__order_email__9021__ordercontactmech").(*ent.OrderContactMech)
	c, err = c.Update().
		SetContactMechPurposeTypeID(common.ParseId("ORDER_EMAIL")).
		SetContactMechID(common.ParseId("9021")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__order_email__9021__ordercontactmech: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__order_email__9021__ordercontactmech", c)
	}

	c = cache.Get("demo81015__shipping_location__9010__ordercontactmech").(*ent.OrderContactMech)
	c, err = c.Update().
		SetContactMechPurposeTypeID(common.ParseId("SHIPPING_LOCATION")).
		SetContactMechID(common.ParseId("9010")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__shipping_location__9010__ordercontactmech: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__shipping_location__9010__ordercontactmech", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
