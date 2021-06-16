package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateInventoryItemDetail(ctx context.Context) error {
	log.Println("InventoryItemDetail updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.InventoryItemDetail
	failures := 0

	c = cache.Get("9001__81015__inventoryitemdetail").(*ent.InventoryItemDetail)
	c, err = c.Update().
		SetInventoryItemDetailSeqID(common.ParseId("81015")).
		SetEffectiveDate(common.ParseDateTime("2011-08-12 23:17:12.433")).
		SetQuantityOnHandDiff(0.000000).
		SetAvailableToPromiseDiff(-5.000000).
		SetAccountingQuantityDiff(0.000000).
		SetOrderItemSeqID(common.ParseId("00001")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9001__81015__inventoryitemdetail: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9001__81015__inventoryitemdetail", c)
	}

	c = cache.Get("9025__81016__inventoryitemdetail").(*ent.InventoryItemDetail)
	c, err = c.Update().
		SetInventoryItemDetailSeqID(common.ParseId("81016")).
		SetEffectiveDate(common.ParseDateTime("2011-08-12 23:17:12.85")).
		SetQuantityOnHandDiff(0.000000).
		SetAvailableToPromiseDiff(-1.000000).
		SetAccountingQuantityDiff(0.000000).
		SetOrderItemSeqID(common.ParseId("00001")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9025__81016__inventoryitemdetail: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9025__81016__inventoryitemdetail", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
