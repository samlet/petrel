package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateStatusItem(ctx context.Context) error {
	log.Println("creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusItem

	c, err = client.StatusItem.Create().SetStringRef("party_enabled__statusitem").
		SetStatusCode("ENABLED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Enabled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__statusitem: %v", err)
		return err
	}
	cache.Put("party_enabled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("party_disabled__statusitem").
		SetStatusCode("DISABLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Disabled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_disabled__statusitem: %v", err)
		return err
	}
	cache.Put("party_disabled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_pending__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__statusitem: %v", err)
		return err
	}
	cache.Put("com_pending__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_entered__statusitem").
		SetStatusCode("ENTERED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Entered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__statusitem: %v", err)
		return err
	}
	cache.Put("com_entered__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_in_progress__statusitem").
		SetStatusCode("IN_PROGRESS").
		SetSequenceID(common.ParseId("05")).
		SetDescription("In-Progress").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_in_progress__statusitem: %v", err)
		return err
	}
	cache.Put("com_in_progress__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_unknown_party__statusitem").
		SetStatusCode("UNKNOWN_PARTY").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Unknown Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_unknown_party__statusitem: %v", err)
		return err
	}
	cache.Put("com_unknown_party__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_complete__statusitem").
		SetStatusCode("COMPLETE").
		SetSequenceID(common.ParseId("20")).
		SetDescription("Closed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__statusitem: %v", err)
		return err
	}
	cache.Put("com_complete__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_resolved__statusitem").
		SetStatusCode("RESOLVED").
		SetSequenceID(common.ParseId("21")).
		SetDescription("Resolved").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_resolved__statusitem: %v", err)
		return err
	}
	cache.Put("com_resolved__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_referred__statusitem").
		SetStatusCode("REFERRED").
		SetSequenceID(common.ParseId("22")).
		SetDescription("Referred").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_referred__statusitem: %v", err)
		return err
	}
	cache.Put("com_referred__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_bounced__statusitem").
		SetStatusCode("BOUNCED").
		SetSequenceID(common.ParseId("50")).
		SetDescription("Bounced").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_bounced__statusitem: %v", err)
		return err
	}
	cache.Put("com_bounced__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("com_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_role_created__statusitem").
		SetStatusCode("ENTERED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_created__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_created__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_role_read__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Read").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_read__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_read__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("com_role_completed__statusitem").
		SetStatusCode("READ").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Closed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_completed__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_completed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyrel_created__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyrel_created__statusitem: %v", err)
		return err
	}
	cache.Put("partyrel_created__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyrel_expired__statusitem").
		SetStatusCode("EXPIRED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Expired").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyrel_expired__statusitem: %v", err)
		return err
	}
	cache.Put("partyrel_expired__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyinv_sent__statusitem").
		SetStatusCode("SENT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Invitation Sent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_sent__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyinv_pending__statusitem").
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Invitation Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_pending__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_pending__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyinv_accepted__statusitem").
		SetStatusCode("ACCEPTED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Invitation Accepted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_accepted__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_accepted__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyinv_declined__statusitem").
		SetStatusCode("DECLINED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Invitation Declined").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_declined__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_declined__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("partyinv_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Invitation Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pas_assigned__statusitem").
		SetStatusCode("ASSIGNED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Assigned").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pas_assigned__statusitem: %v", err)
		return err
	}
	cache.Put("pas_assigned__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pas_completed__statusitem").
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pas_completed__statusitem: %v", err)
		return err
	}
	cache.Put("pas_completed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_created__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_created_ua__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Unassigned").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_ua__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_ua__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_created_as__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Assigned").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_as__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_as__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_created_ip__statusitem").
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("In Progress").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_ip__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_ip__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_completed__statusitem").
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_completed__statusitem: %v", err)
		return err
	}
	cache.Put("pts_completed__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_on_hold__statusitem").
		SetStatusCode("ON_HOLD").
		SetSequenceID(common.ParseId("07")).
		SetDescription("On Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_on_hold__statusitem: %v", err)
		return err
	}
	cache.Put("pts_on_hold__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("pts_cancelled__statusitem").
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("pts_cancelled__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("prj_active__statusitem").
		SetStatusCode("ACTIVE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Active").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prj_active__statusitem: %v", err)
		return err
	}
	cache.Put("prj_active__statusitem", c)

	c, err = client.StatusItem.Create().SetStringRef("prj_closed__statusitem").
		SetStatusCode("CLOSED").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Closed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prj_closed__statusitem: %v", err)
		return err
	}
	cache.Put("prj_closed__statusitem", c)

	return nil
}
