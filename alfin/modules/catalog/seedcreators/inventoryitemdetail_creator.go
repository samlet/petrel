package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateInventoryItemDetail(ctx context.Context) error {
	log.Println("InventoryItemDetail creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.InventoryItemDetail

	c, err = client.InventoryItemDetail.Create().SetStringRef("9001__81015__inventoryitemdetail").
		SetInventoryItemDetailSeqID(common.ParseId("81015")).
		SetEffectiveDate(common.ParseDateTime("2011-08-12 23:17:12.433")).
		SetQuantityOnHandDiff(0.000000).
		SetAvailableToPromiseDiff(-5.000000).
		SetAccountingQuantityDiff(0.000000).
		SetOrderItemSeqID(common.ParseId("00001")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9001__81015__inventoryitemdetail: %v", err)
		return err
	}
	cache.Put("9001__81015__inventoryitemdetail", c)

	c, err = client.InventoryItemDetail.Create().SetStringRef("9025__81016__inventoryitemdetail").
		SetInventoryItemDetailSeqID(common.ParseId("81016")).
		SetEffectiveDate(common.ParseDateTime("2011-08-12 23:17:12.85")).
		SetQuantityOnHandDiff(0.000000).
		SetAvailableToPromiseDiff(-1.000000).
		SetAccountingQuantityDiff(0.000000).
		SetOrderItemSeqID(common.ParseId("00001")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9025__81016__inventoryitemdetail: %v", err)
		return err
	}
	cache.Put("9025__81016__inventoryitemdetail", c)

	return nil
}
