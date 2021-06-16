package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateStatusValidChange(ctx context.Context) error {
	log.Println("StatusValidChange creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusValidChange

	c, err = client.StatusValidChange.Create().SetStringRef("inv_on_order__inv_available__statusvalidchange").
		SetTransitionName("Order Arrived").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_on_order__inv_available__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_on_order__inv_available__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_available__inv_promised__statusvalidchange").
		SetTransitionName("Promise").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_available__inv_promised__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_available__inv_promised__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_available__inv_on_hold__statusvalidchange").
		SetTransitionName("Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_available__inv_on_hold__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_available__inv_on_hold__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_available__inv_defective__statusvalidchange").
		SetTransitionName("Mark As Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_available__inv_defective__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_available__inv_defective__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_available__inv_being_transfered__statusvalidchange").
		SetTransitionName("Being Transfered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_available__inv_being_transfered__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_available__inv_being_transfered__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_on_hold__inv_available__statusvalidchange").
		SetTransitionName("Release Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_on_hold__inv_available__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_on_hold__inv_available__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_on_hold__inv_defective__statusvalidchange").
		SetTransitionName("Mark Held Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_on_hold__inv_defective__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_on_hold__inv_defective__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_being_transfered__inv_available__statusvalidchange").
		SetTransitionName("Transfer Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_being_transfered__inv_available__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_being_transfered__inv_available__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_promised__inv_available__statusvalidchange").
		SetTransitionName("Cancel Promise").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_promised__inv_available__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_promised__inv_available__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_promised__inv_delivered__statusvalidchange").
		SetTransitionName("Deliver").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_promised__inv_delivered__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_promised__inv_delivered__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_promised__inv_defective__statusvalidchange").
		SetTransitionName("Mark As Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_promised__inv_defective__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_promised__inv_defective__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_promised__inv_being_trans_prm__statusvalidchange").
		SetTransitionName("Being Transfered (Promised)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_promised__inv_being_trans_prm__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_promised__inv_being_trans_prm__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_being_trans_prm__inv_promised__statusvalidchange").
		SetTransitionName("Transfer Complete (Promised)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_being_trans_prm__inv_promised__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_being_trans_prm__inv_promised__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_delivered__inv_returned__statusvalidchange").
		SetTransitionName("Return Status Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_delivered__inv_returned__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_delivered__inv_returned__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_returned__inv_available__statusvalidchange").
		SetTransitionName("Make Return Available").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_returned__inv_available__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_returned__inv_available__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_returned__inv_on_hold__statusvalidchange").
		SetTransitionName("Make Return Held").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_returned__inv_on_hold__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_returned__inv_on_hold__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_returned__inv_defective__statusvalidchange").
		SetTransitionName("Mark Return Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_returned__inv_defective__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_returned__inv_defective__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_delivered__inv_activated__statusvalidchange").
		SetTransitionName("Activate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_delivered__inv_activated__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_delivered__inv_activated__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_activated__inv_deactivated__statusvalidchange").
		SetTransitionName("Deactivate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_activated__inv_deactivated__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_activated__inv_deactivated__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_activated__inv_returned__statusvalidchange").
		SetTransitionName("Return Status Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_activated__inv_returned__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_activated__inv_returned__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_deactivated__inv_on_hold__statusvalidchange").
		SetTransitionName("Hold Inactive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_deactivated__inv_on_hold__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_deactivated__inv_on_hold__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_deactivated__inv_returned__statusvalidchange").
		SetTransitionName("Return Inactive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_deactivated__inv_returned__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_deactivated__inv_returned__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_ns_returned__inv_ns_on_hold__statusvalidchange").
		SetTransitionName("Make Return Held").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_ns_returned__inv_ns_on_hold__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_ns_returned__inv_ns_on_hold__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("inv_ns_returned__inv_ns_defective__statusvalidchange").
		SetTransitionName("Mark Return Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_ns_returned__inv_ns_defective__statusvalidchange: %v", err)
		return err
	}
	cache.Put("inv_ns_returned__inv_ns_defective__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_requested__ixf_scheduled__statusvalidchange").
		SetTransitionName("Inventory Request Scheduled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_requested__ixf_scheduled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_requested__ixf_scheduled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_requested__ixf_en_route__statusvalidchange").
		SetTransitionName("Inventory Request In Route").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_requested__ixf_en_route__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_requested__ixf_en_route__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_requested__ixf_complete__statusvalidchange").
		SetTransitionName("Inventory Request Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_requested__ixf_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_requested__ixf_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_requested__ixf_cancelled__statusvalidchange").
		SetTransitionName("Inventory Request Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_requested__ixf_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_requested__ixf_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_scheduled__ixf_en_route__statusvalidchange").
		SetTransitionName("Scheduled Inventory In Route").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_scheduled__ixf_en_route__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_scheduled__ixf_en_route__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_scheduled__ixf_complete__statusvalidchange").
		SetTransitionName("Scheduled Inventory Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_scheduled__ixf_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_scheduled__ixf_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_scheduled__ixf_cancelled__statusvalidchange").
		SetTransitionName("Scheduled Inventory Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_scheduled__ixf_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_scheduled__ixf_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("ixf_en_route__ixf_complete__statusvalidchange").
		SetTransitionName("In Route Transfer Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_en_route__ixf_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("ixf_en_route__ixf_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("prr_pending__prr_approved__statusvalidchange").
		SetTransitionName("Review Approved").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_pending__prr_approved__statusvalidchange: %v", err)
		return err
	}
	cache.Put("prr_pending__prr_approved__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("prr_pending__prr_deleted__statusvalidchange").
		SetTransitionName("Review Deleted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_pending__prr_deleted__statusvalidchange: %v", err)
		return err
	}
	cache.Put("prr_pending__prr_deleted__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("prr_approved__prr_deleted__statusvalidchange").
		SetTransitionName("Review Deleted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_approved__prr_deleted__statusvalidchange: %v", err)
		return err
	}
	cache.Put("prr_approved__prr_deleted__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_input__shipment_scheduled__statusvalidchange").
		SetTransitionName("Schedule").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_input__shipment_scheduled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_input__shipment_scheduled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_input__shipment_picked__statusvalidchange").
		SetTransitionName("Pick").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_input__shipment_picked__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_input__shipment_picked__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_input__shipment_packed__statusvalidchange").
		SetTransitionName("Pack").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_input__shipment_packed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_input__shipment_packed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_scheduled__shipment_picked__statusvalidchange").
		SetTransitionName("Pick").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_scheduled__shipment_picked__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_scheduled__shipment_picked__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_scheduled__shipment_packed__statusvalidchange").
		SetTransitionName("Pack").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_scheduled__shipment_packed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_scheduled__shipment_packed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_picked__shipment_packed__statusvalidchange").
		SetTransitionName("Pack").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_picked__shipment_packed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_picked__shipment_packed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_packed__shipment_shipped__statusvalidchange").
		SetTransitionName("Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_packed__shipment_shipped__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_packed__shipment_shipped__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_shipped__shipment_delivered__statusvalidchange").
		SetTransitionName("Deliver").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_shipped__shipment_delivered__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_shipped__shipment_delivered__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_input__shipment_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_input__shipment_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_input__shipment_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_scheduled__shipment_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_scheduled__shipment_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_scheduled__shipment_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shipment_picked__shipment_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_picked__shipment_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shipment_picked__shipment_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("purch_ship_created__purch_ship_received__statusvalidchange").
		SetTransitionName("Receive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_created__purch_ship_received__statusvalidchange: %v", err)
		return err
	}
	cache.Put("purch_ship_created__purch_ship_received__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("purch_ship_created__purch_ship_shipped__statusvalidchange").
		SetTransitionName("Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_created__purch_ship_shipped__statusvalidchange: %v", err)
		return err
	}
	cache.Put("purch_ship_created__purch_ship_shipped__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("purch_ship_shipped__purch_ship_received__statusvalidchange").
		SetTransitionName("Receive").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_shipped__purch_ship_received__statusvalidchange: %v", err)
		return err
	}
	cache.Put("purch_ship_shipped__purch_ship_received__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("purch_ship_received__purch_ship_shipped__statusvalidchange").
		SetTransitionName("Ship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_received__purch_ship_shipped__statusvalidchange: %v", err)
		return err
	}
	cache.Put("purch_ship_received__purch_ship_shipped__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shrscs_not_started__shrscs_confirmed__statusvalidchange").
		SetTransitionName("Confirm").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_not_started__shrscs_confirmed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shrscs_not_started__shrscs_confirmed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shrscs_confirmed__shrscs_accepted__statusvalidchange").
		SetTransitionName("Accept").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_confirmed__shrscs_accepted__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shrscs_confirmed__shrscs_accepted__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shrscs_confirmed__shrscs_voided__statusvalidchange").
		SetTransitionName("Void").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_confirmed__shrscs_voided__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shrscs_confirmed__shrscs_voided__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("shrscs_accepted__shrscs_voided__statusvalidchange").
		SetTransitionName("Void").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_accepted__shrscs_voided__statusvalidchange: %v", err)
		return err
	}
	cache.Put("shrscs_accepted__shrscs_voided__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_input__picklist_assigned__statusvalidchange").
		SetTransitionName("Assign").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_input__picklist_assigned__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_input__picklist_assigned__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_input__picklist_printed__statusvalidchange").
		SetTransitionName("Print").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_input__picklist_printed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_input__picklist_printed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_input__picklist_picked__statusvalidchange").
		SetTransitionName("Pick").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_input__picklist_picked__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_input__picklist_picked__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_input__picklist_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_input__picklist_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_input__picklist_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_assigned__picklist_picked__statusvalidchange").
		SetTransitionName("Pick").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_assigned__picklist_picked__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_assigned__picklist_picked__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_assigned__picklist_printed__statusvalidchange").
		SetTransitionName("Print").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_assigned__picklist_printed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_assigned__picklist_printed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_assigned__picklist_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_assigned__picklist_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_assigned__picklist_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_printed__picklist_picked__statusvalidchange").
		SetTransitionName("Pick").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_printed__picklist_picked__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_printed__picklist_picked__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("picklist_printed__picklist_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_printed__picklist_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("picklist_printed__picklist_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pickitem_pending__pickitem_completed__statusvalidchange").
		SetTransitionName("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_pending__pickitem_completed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pickitem_pending__pickitem_completed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pickitem_pending__pickitem_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_pending__pickitem_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pickitem_pending__pickitem_cancelled__statusvalidchange", c)

	return nil
}
