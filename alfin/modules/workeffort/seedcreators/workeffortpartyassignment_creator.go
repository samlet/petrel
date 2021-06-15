package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffortPartyAssignment(ctx context.Context) error {
	log.Println("WorkEffortPartyAssignment creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortPartyAssignment

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__company__internal_organizatio__1238963272__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__company__internal_organizatio__1238963272__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__admin__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__admin__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__democustomer1__client_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__democustomer3__client_billing__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9002__admin__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9002__admin__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__admin__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__admin__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__admin__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__company__internal_organizatio__1238963272__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__company__internal_organizatio__1238963272__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__democustomer2__client_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__democustomer3__client_billing__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9200__company__internal_organizatio__1238963272__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.818")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9200__company__internal_organizatio__1238963272__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9200__company__internal_organizatio__1238963272__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2009-04-05 20:27:52.893")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9200__democustcompany__client_billing__1238963272__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment", c)

	c, err = client.WorkEffortPartyAssignment.Create().SetStringRef("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment: %v", err)
		return err
	}
	cache.Put("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment", c)

	return nil
}
