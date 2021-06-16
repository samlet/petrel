package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateInventoryItemType(ctx context.Context) error {
	log.Println("InventoryItemType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.InventoryItemType
	failures := 0

	c = cache.Get("serialized_inv_item__inventoryitemtype").(*ent.InventoryItemType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Serialized").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update serialized_inv_item__inventoryitemtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("serialized_inv_item__inventoryitemtype", c)
	}

	c = cache.Get("non_serial_inv_item__inventoryitemtype").(*ent.InventoryItemType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Non-Serialized").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update non_serial_inv_item__inventoryitemtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("non_serial_inv_item__inventoryitemtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
