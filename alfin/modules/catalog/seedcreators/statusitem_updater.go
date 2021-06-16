package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateStatusItem(ctx context.Context) error {
	log.Println("StatusItem updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusItem
	failures := 0

	c = cache.Get("inv_on_order__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ON_ORDER").
		SetSequenceID(common.ParseId("01")).
		SetDescription("On Order").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_on_order__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_on_order__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_on_order__statusitem", c)
	}

	c = cache.Get("inv_available__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("AVAILABLE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Available").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_available__inv_promised__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_available__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_available__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_available__inv_being_transfered__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_on_order__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_on_hold__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_being_transfered__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_promised__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_returned__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_available__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_available__statusitem", c)
	}

	c = cache.Get("inv_promised__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PROMISED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Promised").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_promised__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_promised__inv_delivered__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_promised__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_promised__inv_being_trans_prm__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_available__inv_promised__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_being_trans_prm__inv_promised__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_promised__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_promised__statusitem", c)
	}

	c = cache.Get("inv_delivered__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DELIVERED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Delivered").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_delivered__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_delivered__inv_activated__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_promised__inv_delivered__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_delivered__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_delivered__statusitem", c)
	}

	c = cache.Get("inv_activated__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ACTIVATED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Activated").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_activated__inv_deactivated__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_activated__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_delivered__inv_activated__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_activated__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_activated__statusitem", c)
	}

	c = cache.Get("inv_deactivated__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DEACTIVATED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Deactivated").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_deactivated__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_deactivated__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_activated__inv_deactivated__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_deactivated__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_deactivated__statusitem", c)
	}

	c = cache.Get("inv_on_hold__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ON_HOLD").
		SetSequenceID(common.ParseId("07")).
		SetDescription("On Hold").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_on_hold__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_on_hold__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_available__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_returned__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_deactivated__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_on_hold__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_on_hold__statusitem", c)
	}

	c = cache.Get("inv_being_transfered__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("BEING_TRANSFERED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Being Transfered").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_being_transfered__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_available__inv_being_transfered__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_being_transfered__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_being_transfered__statusitem", c)
	}

	c = cache.Get("inv_being_trans_prm__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("BEING_TRANS_PRM").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Being Transfered (Promised)").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_being_trans_prm__inv_promised__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_promised__inv_being_trans_prm__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_being_trans_prm__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_being_trans_prm__statusitem", c)
	}

	c = cache.Get("inv_returned__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("RETURNED").
		SetSequenceID(common.ParseId("20")).
		SetDescription("Returned").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_returned__inv_available__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_returned__inv_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_returned__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_delivered__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_activated__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_deactivated__inv_returned__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_returned__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_returned__statusitem", c)
	}

	c = cache.Get("inv_defective__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DEFECTIVE").
		SetSequenceID(common.ParseId("21")).
		SetDescription("Defective").
		SetStatusType(cache.Get("inv_serialized_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("inv_available__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_on_hold__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_promised__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("inv_returned__inv_defective__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_defective__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_defective__statusitem", c)
	}

	c = cache.Get("inv_ns_on_hold__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ON_HOLD_NS").
		SetSequenceID(common.ParseId("01")).
		SetDescription("On Hold (Non-Serialized)").
		SetStatusType(cache.Get("inv_non_ser_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("inv_ns_returned__inv_ns_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_ns_on_hold__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_ns_on_hold__statusitem", c)
	}

	c = cache.Get("inv_ns_defective__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DEFECTIVE_NS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Defective (Non-Serialized)").
		SetStatusType(cache.Get("inv_non_ser_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("inv_ns_returned__inv_ns_defective__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_ns_defective__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_ns_defective__statusitem", c)
	}

	c = cache.Get("inv_ns_returned__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("RETURNED_NS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Returned (Non-Serialized)").
		SetStatusType(cache.Get("inv_non_ser_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("inv_ns_returned__inv_ns_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("inv_ns_returned__inv_ns_defective__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_ns_returned__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_ns_returned__statusitem", c)
	}

	c = cache.Get("ixf_requested__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("REQUESTED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Requested").
		SetStatusType(cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("ixf_requested__ixf_scheduled__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("ixf_requested__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("ixf_requested__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("ixf_requested__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_requested__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_requested__statusitem", c)
	}

	c = cache.Get("ixf_scheduled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SCHEDULED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Scheduled").
		SetStatusType(cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("ixf_scheduled__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("ixf_scheduled__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("ixf_scheduled__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_requested__ixf_scheduled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_scheduled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_scheduled__statusitem", c)
	}

	c = cache.Get("ixf_en_route__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("EN_ROUTE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("En-Route").
		SetStatusType(cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("ixf_en_route__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_requested__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_scheduled__ixf_en_route__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_en_route__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_en_route__statusitem", c)
	}

	c = cache.Get("ixf_complete__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("COMPLETE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Complete").
		SetStatusType(cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("ixf_requested__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_scheduled__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_en_route__ixf_complete__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_complete__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_complete__statusitem", c)
	}

	c = cache.Get("ixf_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("inventory_xfer_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("ixf_requested__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("ixf_scheduled__ixf_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ixf_cancelled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ixf_cancelled__statusitem", c)
	}

	c = cache.Get("prr_pending__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		SetStatusType(cache.Get("product_review_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("prr_pending__prr_approved__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("prr_pending__prr_deleted__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_pending__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_pending__statusitem", c)
	}

	c = cache.Get("prr_approved__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("APPROVED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Approved").
		SetStatusType(cache.Get("product_review_stts__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("prr_approved__prr_deleted__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("prr_pending__prr_approved__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_approved__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_approved__statusitem", c)
	}

	c = cache.Get("prr_deleted__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DELETED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Deleted").
		SetStatusType(cache.Get("product_review_stts__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("prr_pending__prr_deleted__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("prr_approved__prr_deleted__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prr_deleted__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prr_deleted__statusitem", c)
	}

	c = cache.Get("im_pending__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		SetStatusType(cache.Get("image_management_st__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update im_pending__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("im_pending__statusitem", c)
	}

	c = cache.Get("im_approved__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("APPROVED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Approved").
		SetStatusType(cache.Get("image_management_st__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update im_approved__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("im_approved__statusitem", c)
	}

	c = cache.Get("im_rejected__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("REJECTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Rejected").
		SetStatusType(cache.Get("image_management_st__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update im_rejected__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("im_rejected__statusitem", c)
	}

	c = cache.Get("go_created__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		SetStatusType(cache.Get("group_order_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update go_created__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("go_created__statusitem", c)
	}

	c = cache.Get("go_success__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SUCCESS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Success").
		SetStatusType(cache.Get("group_order_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update go_success__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("go_success__statusitem", c)
	}

	c = cache.Get("go_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("group_order_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update go_cancelled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("go_cancelled__statusitem", c)
	}

	c = cache.Get("shipment_input__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("INPUT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Input").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shipment_input__shipment_scheduled__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_input__shipment_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_input__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_input__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_input__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_input__statusitem", c)
	}

	c = cache.Get("shipment_scheduled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SCHEDULED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Scheduled").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shipment_scheduled__shipment_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_scheduled__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_scheduled__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_input__shipment_scheduled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_scheduled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_scheduled__statusitem", c)
	}

	c = cache.Get("shipment_picked__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PICKED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Picked").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shipment_picked__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shipment_picked__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_input__shipment_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_scheduled__shipment_picked__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_picked__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_picked__statusitem", c)
	}

	c = cache.Get("shipment_packed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PACKED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Packed").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shipment_packed__shipment_shipped__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_input__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_scheduled__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_picked__shipment_packed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_packed__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_packed__statusitem", c)
	}

	c = cache.Get("shipment_shipped__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SHIPPED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Shipped").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shipment_shipped__shipment_delivered__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_packed__shipment_shipped__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_shipped__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_shipped__statusitem", c)
	}

	c = cache.Get("shipment_delivered__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DELIVERED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Delivered").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("shipment_shipped__shipment_delivered__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_delivered__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_delivered__statusitem", c)
	}

	c = cache.Get("shipment_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("shipment_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("shipment_input__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_scheduled__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shipment_picked__shipment_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_cancelled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_cancelled__statusitem", c)
	}

	c = cache.Get("purch_ship_created__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		SetStatusType(cache.Get("purch_ship_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("purch_ship_created__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("purch_ship_created__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_created__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_created__statusitem", c)
	}

	c = cache.Get("purch_ship_shipped__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SHIPPED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Shipped").
		SetStatusType(cache.Get("purch_ship_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("purch_ship_shipped__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("purch_ship_created__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("purch_ship_received__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_shipped__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_shipped__statusitem", c)
	}

	c = cache.Get("purch_ship_received__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("RECEIVED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Received").
		SetStatusType(cache.Get("purch_ship_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("purch_ship_received__purch_ship_shipped__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("purch_ship_created__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("purch_ship_shipped__purch_ship_received__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purch_ship_received__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purch_ship_received__statusitem", c)
	}

	c = cache.Get("shrscs_not_started__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("NOT_STARTED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Not Started").
		SetStatusType(cache.Get("shprtsg_cs_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shrscs_not_started__shrscs_confirmed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_not_started__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_not_started__statusitem", c)
	}

	c = cache.Get("shrscs_confirmed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CONFIRMED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Confirmed").
		SetStatusType(cache.Get("shprtsg_cs_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shrscs_confirmed__shrscs_accepted__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("shrscs_confirmed__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shrscs_not_started__shrscs_confirmed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_confirmed__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_confirmed__statusitem", c)
	}

	c = cache.Get("shrscs_accepted__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ACCEPTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Accepted").
		SetStatusType(cache.Get("shprtsg_cs_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("shrscs_accepted__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shrscs_confirmed__shrscs_accepted__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_accepted__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_accepted__statusitem", c)
	}

	c = cache.Get("shrscs_voided__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("VOIDED").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Voided").
		SetStatusType(cache.Get("shprtsg_cs_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("shrscs_confirmed__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("shrscs_accepted__shrscs_voided__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shrscs_voided__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shrscs_voided__statusitem", c)
	}

	c = cache.Get("picklist_input__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("INPUT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Input").
		SetStatusType(cache.Get("picklist_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("picklist_input__picklist_assigned__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_input__picklist_printed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_input__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_input__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_input__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_input__statusitem", c)
	}

	c = cache.Get("picklist_assigned__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ASSIGNED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Assigned").
		SetStatusType(cache.Get("picklist_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("picklist_assigned__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_assigned__picklist_printed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_assigned__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_input__picklist_assigned__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_assigned__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_assigned__statusitem", c)
	}

	c = cache.Get("picklist_printed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PRINTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Printed").
		SetStatusType(cache.Get("picklist_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("picklist_printed__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("picklist_printed__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_input__picklist_printed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_assigned__picklist_printed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_printed__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_printed__statusitem", c)
	}

	c = cache.Get("picklist_picked__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PICKED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Picked").
		SetStatusType(cache.Get("picklist_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("picklist_input__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_assigned__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_printed__picklist_picked__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_picked__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_picked__statusitem", c)
	}

	c = cache.Get("picklist_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("picklist_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("picklist_input__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_assigned__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("picklist_printed__picklist_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picklist_cancelled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picklist_cancelled__statusitem", c)
	}

	c = cache.Get("pickitem_pending__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		SetStatusType(cache.Get("pickitem_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("pickitem_pending__pickitem_completed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("pickitem_pending__pickitem_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_pending__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_pending__statusitem", c)
	}

	c = cache.Get("pickitem_completed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("50")).
		SetDescription("Completed").
		SetStatusType(cache.Get("pickitem_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("pickitem_pending__pickitem_completed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_completed__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_completed__statusitem", c)
	}

	c = cache.Get("pickitem_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("pickitem_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("pickitem_pending__pickitem_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pickitem_cancelled__statusitem: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pickitem_cancelled__statusitem", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
