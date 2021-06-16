package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateStatusType(ctx context.Context) error {
	log.Println("StatusType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusType

	c, err = client.StatusType.Create().SetStringRef("inventory_item_stts__statustype").
		SetHasTable("No").
		SetDescription("Inventory Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inventory_item_stts__statustype: %v", err)
		return err
	}
	cache.Put("inventory_item_stts__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("inv_serialized_stts__statustype").
		SetHasTable("No").
		SetDescription("Serialized Inventory Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_serialized_stts__statustype: %v", err)
		return err
	}
	cache.Put("inv_serialized_stts__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("inv_non_ser_stts__statustype").
		SetHasTable("No").
		SetDescription("Non-Serialized Inventory Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_non_ser_stts__statustype: %v", err)
		return err
	}
	cache.Put("inv_non_ser_stts__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("inventory_xfer_stts__statustype").
		SetHasTable("No").
		SetDescription("Inventory Transfer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inventory_xfer_stts__statustype: %v", err)
		return err
	}
	cache.Put("inventory_xfer_stts__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("product_review_stts__statustype").
		SetHasTable("No").
		SetDescription("Product Review").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_review_stts__statustype: %v", err)
		return err
	}
	cache.Put("product_review_stts__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("image_management_st__statustype").
		SetHasTable("No").
		SetDescription("Image Management").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_management_st__statustype: %v", err)
		return err
	}
	cache.Put("image_management_st__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("group_order_status__statustype").
		SetHasTable("No").
		SetDescription("Group Order Status").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create group_order_status__statustype: %v", err)
		return err
	}
	cache.Put("group_order_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("shipment_status__statustype").
		SetHasTable("No").
		SetDescription("Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_status__statustype: %v", err)
		return err
	}
	cache.Put("shipment_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("purch_ship_status__statustype").
		SetHasTable("No").
		SetDescription("Purchase Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_status__statustype: %v", err)
		return err
	}
	cache.Put("purch_ship_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("shprtsg_cs_status__statustype").
		SetHasTable("No").
		SetDescription("ShipmentRouteSegment:CarrierService").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shprtsg_cs_status__statustype: %v", err)
		return err
	}
	cache.Put("shprtsg_cs_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("picklist_status__statustype").
		SetHasTable("No").
		SetDescription("Picklist").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_status__statustype: %v", err)
		return err
	}
	cache.Put("picklist_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("pickitem_status__statustype").
		SetHasTable("No").
		SetDescription("Picklist Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_status__statustype: %v", err)
		return err
	}
	cache.Put("pickitem_status__statustype", c)

	return nil
}
