package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateStatusValidChange(ctx context.Context) error {
	log.Println("StatusValidChange updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusValidChange
	failures := 0

	c = cache.Get("inv_on_order__inv_available__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Order Arrived").
		SetMainStatusItem(cache.Get("inv_on_order__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_on_order__inv_available__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_on_order__inv_available__statusvalidchange", c)
	}

	c = cache.Get("inv_available__inv_promised__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Promise").
		SetMainStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_available__inv_promised__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_available__inv_promised__statusvalidchange", c)
	}

	c = cache.Get("inv_available__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Hold").
		SetMainStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_available__inv_on_hold__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_available__inv_on_hold__statusvalidchange", c)
	}

	c = cache.Get("inv_available__inv_defective__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Mark As Defective").
		SetMainStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_available__inv_defective__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_available__inv_defective__statusvalidchange", c)
	}

	c = cache.Get("inv_available__inv_being_transfered__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Being Transfered").
		SetMainStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_being_transfered__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_available__inv_being_transfered__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_available__inv_being_transfered__statusvalidchange", c)
	}

	c = cache.Get("inv_on_hold__inv_available__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Release Hold").
		SetMainStatusItem(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_on_hold__inv_available__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_on_hold__inv_available__statusvalidchange", c)
	}

	c = cache.Get("inv_on_hold__inv_defective__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Mark Held Defective").
		SetMainStatusItem(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_on_hold__inv_defective__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_on_hold__inv_defective__statusvalidchange", c)
	}

	c = cache.Get("inv_being_transfered__inv_available__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Transfer Complete").
		SetMainStatusItem(cache.Get("inv_being_transfered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_being_transfered__inv_available__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_being_transfered__inv_available__statusvalidchange", c)
	}

	c = cache.Get("inv_promised__inv_available__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel Promise").
		SetMainStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_promised__inv_available__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_promised__inv_available__statusvalidchange", c)
	}

	c = cache.Get("inv_promised__inv_delivered__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Deliver").
		SetMainStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_delivered__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_promised__inv_delivered__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_promised__inv_delivered__statusvalidchange", c)
	}

	c = cache.Get("inv_promised__inv_defective__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Mark As Defective").
		SetMainStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_promised__inv_defective__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_promised__inv_defective__statusvalidchange", c)
	}

	c = cache.Get("inv_promised__inv_being_trans_prm__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Being Transfered (Promised)").
		SetMainStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_being_trans_prm__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_promised__inv_being_trans_prm__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_promised__inv_being_trans_prm__statusvalidchange", c)
	}

	c = cache.Get("inv_being_trans_prm__inv_promised__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Transfer Complete (Promised)").
		SetMainStatusItem(cache.Get("inv_being_trans_prm__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_promised__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_being_trans_prm__inv_promised__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_being_trans_prm__inv_promised__statusvalidchange", c)
	}

	c = cache.Get("inv_delivered__inv_returned__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Return Status Pending").
		SetMainStatusItem(cache.Get("inv_delivered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_delivered__inv_returned__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_delivered__inv_returned__statusvalidchange", c)
	}

	c = cache.Get("inv_returned__inv_available__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Make Return Available").
		SetMainStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_available__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_returned__inv_available__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_returned__inv_available__statusvalidchange", c)
	}

	c = cache.Get("inv_returned__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Make Return Held").
		SetMainStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_returned__inv_on_hold__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_returned__inv_on_hold__statusvalidchange", c)
	}

	c = cache.Get("inv_returned__inv_defective__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Mark Return Defective").
		SetMainStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_returned__inv_defective__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_returned__inv_defective__statusvalidchange", c)
	}

	c = cache.Get("inv_delivered__inv_activated__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Activate").
		SetMainStatusItem(cache.Get("inv_delivered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_activated__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_delivered__inv_activated__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_delivered__inv_activated__statusvalidchange", c)
	}

	c = cache.Get("inv_activated__inv_deactivated__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Deactivate").
		SetMainStatusItem(cache.Get("inv_activated__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_deactivated__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_activated__inv_deactivated__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_activated__inv_deactivated__statusvalidchange", c)
	}

	c = cache.Get("inv_activated__inv_returned__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Return Status Pending").
		SetMainStatusItem(cache.Get("inv_activated__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_activated__inv_returned__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_activated__inv_returned__statusvalidchange", c)
	}

	c = cache.Get("inv_deactivated__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Hold Inactive").
		SetMainStatusItem(cache.Get("inv_deactivated__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_deactivated__inv_on_hold__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_deactivated__inv_on_hold__statusvalidchange", c)
	}

	c = cache.Get("inv_deactivated__inv_returned__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Return Inactive").
		SetMainStatusItem(cache.Get("inv_deactivated__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_returned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_deactivated__inv_returned__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_deactivated__inv_returned__statusvalidchange", c)
	}

	c = cache.Get("inv_ns_returned__inv_ns_on_hold__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Make Return Held").
		SetMainStatusItem(cache.Get("inv_ns_returned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_ns_on_hold__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_ns_returned__inv_ns_on_hold__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_ns_returned__inv_ns_on_hold__statusvalidchange", c)
	}

	c = cache.Get("inv_ns_returned__inv_ns_defective__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Mark Return Defective").
		SetMainStatusItem(cache.Get("inv_ns_returned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("inv_ns_defective__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_ns_returned__inv_ns_defective__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_ns_returned__inv_ns_defective__statusvalidchange", c)
	}

	c = cache.Get("ixf_requested__ixf_scheduled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Inventory Request Scheduled").
		SetMainStatusItem(cache.Get("ixf_requested__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_requested__ixf_scheduled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_requested__ixf_scheduled__statusvalidchange", c)
	}

	c = cache.Get("ixf_requested__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Inventory Request In Route").
		SetMainStatusItem(cache.Get("ixf_requested__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_en_route__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_requested__ixf_en_route__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_requested__ixf_en_route__statusvalidchange", c)
	}

	c = cache.Get("ixf_requested__ixf_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Inventory Request Completed").
		SetMainStatusItem(cache.Get("ixf_requested__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_requested__ixf_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_requested__ixf_complete__statusvalidchange", c)
	}

	c = cache.Get("ixf_requested__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Inventory Request Cancelled").
		SetMainStatusItem(cache.Get("ixf_requested__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_requested__ixf_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_requested__ixf_cancelled__statusvalidchange", c)
	}

	c = cache.Get("ixf_scheduled__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Scheduled Inventory In Route").
		SetMainStatusItem(cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_en_route__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_scheduled__ixf_en_route__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_scheduled__ixf_en_route__statusvalidchange", c)
	}

	c = cache.Get("ixf_scheduled__ixf_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Scheduled Inventory Completed").
		SetMainStatusItem(cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_scheduled__ixf_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_scheduled__ixf_complete__statusvalidchange", c)
	}

	c = cache.Get("ixf_scheduled__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Scheduled Inventory Cancelled").
		SetMainStatusItem(cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_scheduled__ixf_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_scheduled__ixf_cancelled__statusvalidchange", c)
	}

	c = cache.Get("ixf_en_route__ixf_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("In Route Transfer Completed").
		SetMainStatusItem(cache.Get("ixf_en_route__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("ixf_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_en_route__ixf_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_en_route__ixf_complete__statusvalidchange", c)
	}

	c = cache.Get("prr_pending__prr_approved__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Review Approved").
		SetMainStatusItem(cache.Get("prr_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("prr_approved__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_pending__prr_approved__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_pending__prr_approved__statusvalidchange", c)
	}

	c = cache.Get("prr_pending__prr_deleted__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Review Deleted").
		SetMainStatusItem(cache.Get("prr_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("prr_deleted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_pending__prr_deleted__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_pending__prr_deleted__statusvalidchange", c)
	}

	c = cache.Get("prr_approved__prr_deleted__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Review Deleted").
		SetMainStatusItem(cache.Get("prr_approved__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("prr_deleted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_approved__prr_deleted__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_approved__prr_deleted__statusvalidchange", c)
	}

	c = cache.Get("shipment_input__shipment_scheduled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Schedule").
		SetMainStatusItem(cache.Get("shipment_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_input__shipment_scheduled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_input__shipment_scheduled__statusvalidchange", c)
	}

	c = cache.Get("shipment_input__shipment_picked__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pick").
		SetMainStatusItem(cache.Get("shipment_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_picked__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_input__shipment_picked__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_input__shipment_picked__statusvalidchange", c)
	}

	c = cache.Get("shipment_input__shipment_packed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pack").
		SetMainStatusItem(cache.Get("shipment_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_packed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_input__shipment_packed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_input__shipment_packed__statusvalidchange", c)
	}

	c = cache.Get("shipment_scheduled__shipment_picked__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pick").
		SetMainStatusItem(cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_picked__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_scheduled__shipment_picked__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_scheduled__shipment_picked__statusvalidchange", c)
	}

	c = cache.Get("shipment_scheduled__shipment_packed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pack").
		SetMainStatusItem(cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_packed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_scheduled__shipment_packed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_scheduled__shipment_packed__statusvalidchange", c)
	}

	c = cache.Get("shipment_picked__shipment_packed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pack").
		SetMainStatusItem(cache.Get("shipment_picked__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_packed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_picked__shipment_packed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_picked__shipment_packed__statusvalidchange", c)
	}

	c = cache.Get("shipment_packed__shipment_shipped__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Ship").
		SetMainStatusItem(cache.Get("shipment_packed__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_shipped__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_packed__shipment_shipped__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_packed__shipment_shipped__statusvalidchange", c)
	}

	c = cache.Get("shipment_shipped__shipment_delivered__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Deliver").
		SetMainStatusItem(cache.Get("shipment_shipped__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_delivered__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_shipped__shipment_delivered__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_shipped__shipment_delivered__statusvalidchange", c)
	}

	c = cache.Get("shipment_input__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("shipment_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_input__shipment_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_input__shipment_cancelled__statusvalidchange", c)
	}

	c = cache.Get("shipment_scheduled__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_scheduled__shipment_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_scheduled__shipment_cancelled__statusvalidchange", c)
	}

	c = cache.Get("shipment_picked__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("shipment_picked__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shipment_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_picked__shipment_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_picked__shipment_cancelled__statusvalidchange", c)
	}

	c = cache.Get("purch_ship_created__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Receive").
		SetMainStatusItem(cache.Get("purch_ship_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("purch_ship_received__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_created__purch_ship_received__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_created__purch_ship_received__statusvalidchange", c)
	}

	c = cache.Get("purch_ship_created__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Ship").
		SetMainStatusItem(cache.Get("purch_ship_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("purch_ship_shipped__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_created__purch_ship_shipped__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_created__purch_ship_shipped__statusvalidchange", c)
	}

	c = cache.Get("purch_ship_shipped__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Receive").
		SetMainStatusItem(cache.Get("purch_ship_shipped__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("purch_ship_received__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_shipped__purch_ship_received__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_shipped__purch_ship_received__statusvalidchange", c)
	}

	c = cache.Get("purch_ship_received__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Ship").
		SetMainStatusItem(cache.Get("purch_ship_received__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("purch_ship_shipped__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_received__purch_ship_shipped__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_received__purch_ship_shipped__statusvalidchange", c)
	}

	c = cache.Get("shrscs_not_started__shrscs_confirmed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Confirm").
		SetMainStatusItem(cache.Get("shrscs_not_started__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shrscs_confirmed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_not_started__shrscs_confirmed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_not_started__shrscs_confirmed__statusvalidchange", c)
	}

	c = cache.Get("shrscs_confirmed__shrscs_accepted__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Accept").
		SetMainStatusItem(cache.Get("shrscs_confirmed__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shrscs_accepted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_confirmed__shrscs_accepted__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_confirmed__shrscs_accepted__statusvalidchange", c)
	}

	c = cache.Get("shrscs_confirmed__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Void").
		SetMainStatusItem(cache.Get("shrscs_confirmed__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shrscs_voided__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_confirmed__shrscs_voided__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_confirmed__shrscs_voided__statusvalidchange", c)
	}

	c = cache.Get("shrscs_accepted__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Void").
		SetMainStatusItem(cache.Get("shrscs_accepted__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("shrscs_voided__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_accepted__shrscs_voided__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_accepted__shrscs_voided__statusvalidchange", c)
	}

	c = cache.Get("picklist_input__picklist_assigned__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Assign").
		SetMainStatusItem(cache.Get("picklist_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_input__picklist_assigned__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_input__picklist_assigned__statusvalidchange", c)
	}

	c = cache.Get("picklist_input__picklist_printed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Print").
		SetMainStatusItem(cache.Get("picklist_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_printed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_input__picklist_printed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_input__picklist_printed__statusvalidchange", c)
	}

	c = cache.Get("picklist_input__picklist_picked__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pick").
		SetMainStatusItem(cache.Get("picklist_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_picked__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_input__picklist_picked__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_input__picklist_picked__statusvalidchange", c)
	}

	c = cache.Get("picklist_input__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("picklist_input__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_input__picklist_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_input__picklist_cancelled__statusvalidchange", c)
	}

	c = cache.Get("picklist_assigned__picklist_picked__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pick").
		SetMainStatusItem(cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_picked__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_assigned__picklist_picked__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_assigned__picklist_picked__statusvalidchange", c)
	}

	c = cache.Get("picklist_assigned__picklist_printed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Print").
		SetMainStatusItem(cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_printed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_assigned__picklist_printed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_assigned__picklist_printed__statusvalidchange", c)
	}

	c = cache.Get("picklist_assigned__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_assigned__picklist_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_assigned__picklist_cancelled__statusvalidchange", c)
	}

	c = cache.Get("picklist_printed__picklist_picked__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pick").
		SetMainStatusItem(cache.Get("picklist_printed__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_picked__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_printed__picklist_picked__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_printed__picklist_picked__statusvalidchange", c)
	}

	c = cache.Get("picklist_printed__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("picklist_printed__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("picklist_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_printed__picklist_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_printed__picklist_cancelled__statusvalidchange", c)
	}

	c = cache.Get("pickitem_pending__pickitem_completed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Complete").
		SetMainStatusItem(cache.Get("pickitem_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pickitem_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_pending__pickitem_completed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_pending__pickitem_completed__statusvalidchange", c)
	}

	c = cache.Get("pickitem_pending__pickitem_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("pickitem_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pickitem_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_pending__pickitem_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_pending__pickitem_cancelled__statusvalidchange", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
