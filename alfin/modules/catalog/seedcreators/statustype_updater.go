package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateStatusType(ctx context.Context) error {
	log.Println("StatusType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusType
	failures := 0

	c = cache.Get("inventory_item_stts__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inventory Item").
		AddChildren(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddChildren(cache.Get("inv_non_ser_stts__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inventory_item_stts__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inventory_item_stts__statustype", c)
	}

	c = cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Serialized Inventory Item").
		SetParent(cache.Get("inventory_item_stts__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("inv_on_order__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_delivered__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_activated__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_deactivated__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_being_transfered__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_being_trans_prm__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_serialized_stts__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_serialized_stts__statustype", c)
	}

	c = cache.Get("inv_non_ser_stts__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Non-Serialized Inventory Item").
		SetParent(cache.Get("inventory_item_stts__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("inv_ns_on_hold__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_ns_defective__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("inv_ns_returned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_non_ser_stts__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_non_ser_stts__statustype", c)
	}

	c = cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inventory Transfer").
		AddStatusItems(cache.Get("ixf_requested__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("ixf_en_route__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("ixf_complete__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("ixf_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inventory_xfer_stts__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inventory_xfer_stts__statustype", c)
	}

	c = cache.Get("product_review_stts__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Review").
		AddStatusItems(cache.Get("prr_pending__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("prr_approved__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("prr_deleted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_review_stts__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_review_stts__statustype", c)
	}

	c = cache.Get("image_management_st__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Management").
		AddStatusItems(cache.Get("im_pending__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("im_approved__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("im_rejected__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_management_st__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_management_st__statustype", c)
	}

	c = cache.Get("group_order_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Group Order Status").
		AddStatusItems(cache.Get("go_created__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("go_success__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("go_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update group_order_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("group_order_status__statustype", c)
	}

	c = cache.Get("shipment_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipment").
		AddStatusItems(cache.Get("shipment_input__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_picked__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_packed__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_shipped__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_delivered__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shipment_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_status__statustype", c)
	}

	c = cache.Get("purch_ship_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Purchase Shipment").
		AddStatusItems(cache.Get("purch_ship_created__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("purch_ship_shipped__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("purch_ship_received__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_status__statustype", c)
	}

	c = cache.Get("shprtsg_cs_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("ShipmentRouteSegment:CarrierService").
		AddStatusItems(cache.Get("shrscs_not_started__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shrscs_confirmed__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shrscs_accepted__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("shrscs_voided__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shprtsg_cs_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shprtsg_cs_status__statustype", c)
	}

	c = cache.Get("picklist_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Picklist").
		AddStatusItems(cache.Get("picklist_input__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("picklist_printed__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("picklist_picked__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("picklist_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_status__statustype", c)
	}

	c = cache.Get("pickitem_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Picklist Item").
		AddStatusItems(cache.Get("pickitem_pending__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pickitem_completed__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pickitem_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_status__statustype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_status__statustype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
