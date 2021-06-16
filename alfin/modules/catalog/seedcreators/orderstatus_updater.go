package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderStatus(ctx context.Context) error {
	log.Println("OrderStatus updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderStatus
	failures := 0

	c = cache.Get("81015__orderstatus").(*ent.OrderStatus)
	c, err = c.Update().
		SetStatusDatetime(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetStatusUserLogin("admin").
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 81015__orderstatus: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("81015__orderstatus", c)
	}

	c = cache.Get("81016__orderstatus").(*ent.OrderStatus)
	c, err = c.Update().
		SetOrderItemSeqID(common.ParseId("00001")).
		SetStatusDatetime(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetStatusUserLogin("admin").
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 81016__orderstatus: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("81016__orderstatus", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
