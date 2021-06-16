package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateInventoryItemType(ctx context.Context) error {
	log.Println("InventoryItemType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.InventoryItemType

	c, err = client.InventoryItemType.Create().SetStringRef("serialized_inv_item__inventoryitemtype").
		SetHasTable("No").
		SetDescription("Serialized").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create serialized_inv_item__inventoryitemtype: %v", err)
		return err
	}
	cache.Put("serialized_inv_item__inventoryitemtype", c)

	c, err = client.InventoryItemType.Create().SetStringRef("non_serial_inv_item__inventoryitemtype").
		SetHasTable("No").
		SetDescription("Non-Serialized").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create non_serial_inv_item__inventoryitemtype: %v", err)
		return err
	}
	cache.Put("non_serial_inv_item__inventoryitemtype", c)

	return nil
}
