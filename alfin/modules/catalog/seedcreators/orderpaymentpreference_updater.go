package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderPaymentPreference(ctx context.Context) error {
	log.Println("OrderPaymentPreference updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderPaymentPreference
	failures := 0

	c = cache.Get("10020__orderpaymentpreference").(*ent.OrderPaymentPreference)
	c, err = c.Update().
		SetPaymentMethodTypeID(common.ParseId("EXT_OFFLINE")).
		SetPresentFlag("No").
		SetSwipedFlag("No").
		SetOverflowFlag("No").
		SetMaxAmount(105.14).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.789")).
		SetCreatedByUserLogin("admin").
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 10020__orderpaymentpreference: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("10020__orderpaymentpreference", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
