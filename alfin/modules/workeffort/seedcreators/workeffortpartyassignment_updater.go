package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateWorkEffortPartyAssignment(ctx context.Context) error {
	log.Println("WorkEffortPartyAssignment updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortPartyAssignment
	failures := 0

	c = cache.Get("9000__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__company__internal_organizatio__1238963272__workeffortpartyassignment", c)
	}

	c = cache.Get("9000__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__admin__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("democustomer1__client_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__democustomer1__client_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("democustomer3__client_billing__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__democustomer3__client_billing__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee2__provider_analyst__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9002__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9002__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9002__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9002__admin__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__admin__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__company__internal_organizatio__1238963272__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("democustomer2__client_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__democustomer2__client_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("democustomer3__client_billing__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__democustomer3__client_billing__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9200__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("9200__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		SetWorkEffort(cache.Get("9200__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9200__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9200__company__internal_organizatio__1238963272__workeffortpartyassignment", c)
	}

	c = cache.Get("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.893")).
		SetWorkEffort(cache.Get("9200__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("democustcompany__client_billing__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9200__democustcompany__client_billing__1238963272__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment", c)
	}

	c = cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffort(cache.Get("9200__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		SetAssignmentStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)
	}

	c = cache.Get("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("calendar_pub_demo__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__cal_owner__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("project_pub_demo__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee3__cal_attendee__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__cal_attendee__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("oneoffmeeting__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("oneoffmeeting__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("oneoffmeeting__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("admin__cal_attendee__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment", c)
	}

	c = cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffort(cache.Get("privatedemoemployee1__workeffort").(*ent.WorkEffort)).
		SetPartyRole(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
