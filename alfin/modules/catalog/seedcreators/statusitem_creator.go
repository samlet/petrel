package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateStatusItem(ctx context.Context) error {
	log.Println("StatusItem creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusItem

	c, err = client.StatusItem.Create().SetStringRef("inv_on_order__statusitem").
		SetStatusCode("ON_ORDER").
		SetSequenceID(common.ParseId("01")).
		SetDescription("On Order").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_on_order__statusitem: %v", err)
		return err
	}
	cache.Put("inv_on_order__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_available__statusitem").
		SetStatusCode("AVAILABLE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Available").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_available__statusitem: %v", err)
		return err
	}
	cache.Put("inv_available__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_promised__statusitem").
		SetStatusCode("PROMISED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Promised").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_promised__statusitem: %v", err)
		return err
	}
	cache.Put("inv_promised__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_delivered__statusitem").
		SetStatusCode("DELIVERED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Delivered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_delivered__statusitem: %v", err)
		return err
	}
	cache.Put("inv_delivered__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_activated__statusitem").
		SetStatusCode("ACTIVATED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Activated").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_activated__statusitem: %v", err)
		return err
	}
	cache.Put("inv_activated__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_deactivated__statusitem").
		SetStatusCode("DEACTIVATED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Deactivated").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_deactivated__statusitem: %v", err)
		return err
	}
	cache.Put("inv_deactivated__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_on_hold__statusitem").
		SetStatusCode("ON_HOLD").
		SetSequenceID(common.ParseId("07")).
		SetDescription("On Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_on_hold__statusitem: %v", err)
		return err
	}
	cache.Put("inv_on_hold__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_being_transfered__statusitem").
		SetStatusCode("BEING_TRANSFERED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Being Transfered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_being_transfered__statusitem: %v", err)
		return err
	}
	cache.Put("inv_being_transfered__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_being_trans_prm__statusitem").
		SetStatusCode("BEING_TRANS_PRM").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Being Transfered (Promised)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_being_trans_prm__statusitem: %v", err)
		return err
	}
	cache.Put("inv_being_trans_prm__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_returned__statusitem").
		SetStatusCode("RETURNED").
		SetSequenceID(common.ParseId("20")).
		SetDescription("Returned").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_returned__statusitem: %v", err)
		return err
	}
	cache.Put("inv_returned__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_defective__statusitem").
		SetStatusCode("DEFECTIVE").
		SetSequenceID(common.ParseId("21")).
		SetDescription("Defective").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_defective__statusitem: %v", err)
		return err
	}
	cache.Put("inv_defective__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_ns_on_hold__statusitem").
		SetStatusCode("ON_HOLD_NS").
		SetSequenceID(common.ParseId("01")).
		SetDescription("On Hold (Non-Serialized)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_ns_on_hold__statusitem: %v", err)
		return err
	}
	cache.Put("inv_ns_on_hold__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_ns_defective__statusitem").
		SetStatusCode("DEFECTIVE_NS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Defective (Non-Serialized)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_ns_defective__statusitem: %v", err)
		return err
	}
	cache.Put("inv_ns_defective__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("inv_ns_returned__statusitem").
		SetStatusCode("RETURNED_NS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Returned (Non-Serialized)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_ns_returned__statusitem: %v", err)
		return err
	}
	cache.Put("inv_ns_returned__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("ixf_requested__statusitem").
		SetStatusCode("REQUESTED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Requested").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_requested__statusitem: %v", err)
		return err
	}
	cache.Put("ixf_requested__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("ixf_scheduled__statusitem").
		SetStatusCode("SCHEDULED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Scheduled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_scheduled__statusitem: %v", err)
		return err
	}
	cache.Put("ixf_scheduled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("ixf_en_route__statusitem").
		SetStatusCode("EN_ROUTE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("En-Route").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_en_route__statusitem: %v", err)
		return err
	}
	cache.Put("ixf_en_route__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("ixf_complete__statusitem").
		SetStatusCode("COMPLETE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_complete__statusitem: %v", err)
		return err
	}
	cache.Put("ixf_complete__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("ixf_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ixf_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("ixf_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("prr_pending__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_pending__statusitem: %v", err)
		return err
	}
	cache.Put("prr_pending__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("prr_approved__statusitem").
		SetStatusCode("APPROVED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Approved").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_approved__statusitem: %v", err)
		return err
	}
	cache.Put("prr_approved__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("prr_deleted__statusitem").
		SetStatusCode("DELETED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Deleted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prr_deleted__statusitem: %v", err)
		return err
	}
	cache.Put("prr_deleted__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("im_pending__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create im_pending__statusitem: %v", err)
		return err
	}
	cache.Put("im_pending__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("im_approved__statusitem").
		SetStatusCode("APPROVED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Approved").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create im_approved__statusitem: %v", err)
		return err
	}
	cache.Put("im_approved__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("im_rejected__statusitem").
		SetStatusCode("REJECTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Rejected").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create im_rejected__statusitem: %v", err)
		return err
	}
	cache.Put("im_rejected__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("go_created__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create go_created__statusitem: %v", err)
		return err
	}
	cache.Put("go_created__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("go_success__statusitem").
		SetStatusCode("SUCCESS").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Success").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create go_success__statusitem: %v", err)
		return err
	}
	cache.Put("go_success__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("go_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create go_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("go_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_input__statusitem").
		SetStatusCode("INPUT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Input").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_input__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_input__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_scheduled__statusitem").
		SetStatusCode("SCHEDULED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Scheduled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_scheduled__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_scheduled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_picked__statusitem").
		SetStatusCode("PICKED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Picked").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_picked__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_picked__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_packed__statusitem").
		SetStatusCode("PACKED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Packed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_packed__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_packed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_shipped__statusitem").
		SetStatusCode("SHIPPED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Shipped").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_shipped__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_shipped__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_delivered__statusitem").
		SetStatusCode("DELIVERED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Delivered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_delivered__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_delivered__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shipment_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("shipment_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("purch_ship_created__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_created__statusitem: %v", err)
		return err
	}
	cache.Put("purch_ship_created__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("purch_ship_shipped__statusitem").
		SetStatusCode("SHIPPED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Shipped").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_shipped__statusitem: %v", err)
		return err
	}
	cache.Put("purch_ship_shipped__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("purch_ship_received__statusitem").
		SetStatusCode("RECEIVED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Received").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purch_ship_received__statusitem: %v", err)
		return err
	}
	cache.Put("purch_ship_received__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shrscs_not_started__statusitem").
		SetStatusCode("NOT_STARTED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Not Started").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_not_started__statusitem: %v", err)
		return err
	}
	cache.Put("shrscs_not_started__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shrscs_confirmed__statusitem").
		SetStatusCode("CONFIRMED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Confirmed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_confirmed__statusitem: %v", err)
		return err
	}
	cache.Put("shrscs_confirmed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shrscs_accepted__statusitem").
		SetStatusCode("ACCEPTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Accepted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_accepted__statusitem: %v", err)
		return err
	}
	cache.Put("shrscs_accepted__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("shrscs_voided__statusitem").
		SetStatusCode("VOIDED").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Voided").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shrscs_voided__statusitem: %v", err)
		return err
	}
	cache.Put("shrscs_voided__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("picklist_input__statusitem").
		SetStatusCode("INPUT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Input").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_input__statusitem: %v", err)
		return err
	}
	cache.Put("picklist_input__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("picklist_assigned__statusitem").
		SetStatusCode("ASSIGNED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Assigned").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_assigned__statusitem: %v", err)
		return err
	}
	cache.Put("picklist_assigned__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("picklist_printed__statusitem").
		SetStatusCode("PRINTED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Printed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_printed__statusitem: %v", err)
		return err
	}
	cache.Put("picklist_printed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("picklist_picked__statusitem").
		SetStatusCode("PICKED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Picked").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_picked__statusitem: %v", err)
		return err
	}
	cache.Put("picklist_picked__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("picklist_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picklist_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("picklist_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pickitem_pending__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_pending__statusitem: %v", err)
		return err
	}
	cache.Put("pickitem_pending__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pickitem_completed__statusitem").
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("50")).
		SetDescription("Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_completed__statusitem: %v", err)
		return err
	}
	cache.Put("pickitem_completed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pickitem_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pickitem_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("pickitem_cancelled__statusitem", c)

	return nil
}
