package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateStatusItem(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusItem

	c = cache.Get("party_enabled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ENABLED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Enabled").
		SetStatusType(cache.Get("party_status__statustype").(*ent.StatusType)).
		AddParties(cache.Get("demoemployee__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee1__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee2__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee3__party").(*ent.Party)).
		AddParties(cache.Get("democustomer1__party").(*ent.Party)).
		AddParties(cache.Get("democustomer2__party").(*ent.Party)).
		AddParties(cache.Get("democustomer3__party").(*ent.Party)).
		AddParties(cache.Get("workeffortuser__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee1__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee2__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee3__party").(*ent.Party)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee1__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee2__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee3__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer1__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer2__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer3__978350400__partystatus").(*ent.PartyStatus)).
		AddPartyStatuses(cache.Get("party_enabled__workeffortuser__978350400__partystatus").(*ent.PartyStatus)).
		AddMainStatusValidChanges(cache.Get("party_enabled__party_disabled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("party_disabled__party_enabled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__statusitem: %v", err)
		return err
	}
	cache.Put("party_enabled__statusitem", c)

	c = cache.Get("party_disabled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DISABLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Disabled").
		SetStatusType(cache.Get("party_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("party_disabled__party_enabled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("party_enabled__party_disabled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_disabled__statusitem: %v", err)
		return err
	}
	cache.Put("party_disabled__statusitem", c)

	c = cache.Get("com_pending__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pending").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_pending__com_entered__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_pending__com_in_progress__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_pending__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_pending__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_entered__com_pending__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__statusitem: %v", err)
		return err
	}
	cache.Put("com_pending__statusitem", c)

	c = cache.Get("com_entered__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ENTERED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Entered").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_entered__com_pending__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_entered__com_in_progress__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_entered__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_entered__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_pending__com_entered__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_unknown_party__com_entered__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__statusitem: %v", err)
		return err
	}
	cache.Put("com_entered__statusitem", c)

	c = cache.Get("com_in_progress__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("IN_PROGRESS").
		SetSequenceID(common.ParseId("05")).
		SetDescription("In-Progress").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_in_progress__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_in_progress__com_bounced__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_in_progress__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_entered__com_in_progress__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_pending__com_in_progress__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_in_progress__statusitem: %v", err)
		return err
	}
	cache.Put("com_in_progress__statusitem", c)

	c = cache.Get("com_unknown_party__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("UNKNOWN_PARTY").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Unknown Party").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_unknown_party__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_unknown_party__com_entered__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_unknown_party__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_unknown_party__statusitem: %v", err)
		return err
	}
	cache.Put("com_unknown_party__statusitem", c)

	c = cache.Get("com_complete__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("COMPLETE").
		SetSequenceID(common.ParseId("20")).
		SetDescription("Closed").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_complete__com_resolved__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_complete__com_referred__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_complete__com_bounced__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_complete__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_entered__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_in_progress__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_unknown_party__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_pending__com_complete__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__statusitem: %v", err)
		return err
	}
	cache.Put("com_complete__statusitem", c)

	c = cache.Get("com_resolved__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("RESOLVED").
		SetSequenceID(common.ParseId("21")).
		SetDescription("Resolved").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_resolved__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_complete__com_resolved__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_resolved__statusitem: %v", err)
		return err
	}
	cache.Put("com_resolved__statusitem", c)

	c = cache.Get("com_referred__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("REFERRED").
		SetSequenceID(common.ParseId("22")).
		SetDescription("Referred").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_referred__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_complete__com_referred__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_referred__statusitem: %v", err)
		return err
	}
	cache.Put("com_referred__statusitem", c)

	c = cache.Get("com_bounced__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("BOUNCED").
		SetSequenceID(common.ParseId("50")).
		SetDescription("Bounced").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("com_in_progress__com_bounced__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_complete__com_bounced__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_bounced__statusitem: %v", err)
		return err
	}
	cache.Put("com_bounced__statusitem", c)

	c = cache.Get("com_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("99")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("com_entered__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_pending__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_in_progress__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_complete__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_resolved__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_referred__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_unknown_party__com_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("com_cancelled__statusitem", c)

	c = cache.Get("com_role_created__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ENTERED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		SetStatusType(cache.Get("com_event_rol_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_role_created__com_role_read__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("com_role_created__com_role_completed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_created__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_created__statusitem", c)

	c = cache.Get("com_role_read__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Read").
		SetStatusType(cache.Get("com_event_rol_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("com_role_read__com_role_completed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_role_created__com_role_read__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_read__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_read__statusitem", c)

	c = cache.Get("com_role_completed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("READ").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Closed").
		SetStatusType(cache.Get("com_event_rol_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("com_role_created__com_role_completed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("com_role_read__com_role_completed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_completed__statusitem: %v", err)
		return err
	}
	cache.Put("com_role_completed__statusitem", c)

	c = cache.Get("partyrel_created__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		SetStatusType(cache.Get("party_rel_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("partyrel_created__partyrel_expired__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyrel_created__statusitem: %v", err)
		return err
	}
	cache.Put("partyrel_created__statusitem", c)

	c = cache.Get("partyrel_expired__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("EXPIRED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Expired").
		SetStatusType(cache.Get("party_rel_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("partyrel_created__partyrel_expired__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyrel_expired__statusitem: %v", err)
		return err
	}
	cache.Put("partyrel_expired__statusitem", c)

	c = cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("SENT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Invitation Sent").
		SetStatusType(cache.Get("party_inv_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("partyinv_sent__partyinv_pending__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("partyinv_sent__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("partyinv_sent__partyinv_declined__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("partyinv_sent__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_sent__statusitem", c)

	c = cache.Get("partyinv_pending__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("PENDING").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Invitation Pending").
		SetStatusType(cache.Get("party_inv_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("partyinv_pending__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("partyinv_pending__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("partyinv_sent__partyinv_pending__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_pending__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_pending__statusitem", c)

	c = cache.Get("partyinv_accepted__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ACCEPTED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Invitation Accepted").
		SetStatusType(cache.Get("party_inv_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("partyinv_sent__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("partyinv_pending__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_accepted__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_accepted__statusitem", c)

	c = cache.Get("partyinv_declined__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("DECLINED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Invitation Declined").
		SetStatusType(cache.Get("party_inv_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("partyinv_sent__partyinv_declined__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_declined__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_declined__statusitem", c)

	c = cache.Get("partyinv_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Invitation Cancelled").
		SetStatusType(cache.Get("party_inv_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("partyinv_sent__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("partyinv_pending__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("partyinv_cancelled__statusitem", c)

	c = cache.Get("pas_assigned__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ASSIGNED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Assigned").
		SetStatusType(cache.Get("project_assgn_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("pas_assigned__pas_completed__statusvalidchange").(*ent.StatusValidChange)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9000__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9002__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9100__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddAssignmentWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pas_assigned__statusitem: %v", err)
		return err
	}
	cache.Put("pas_assigned__statusitem", c)

	c = cache.Get("pas_completed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Completed").
		SetStatusType(cache.Get("project_assgn_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("pas_assigned__pas_completed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pas_completed__statusitem: %v", err)
		return err
	}
	cache.Put("pas_completed__statusitem", c)

	c = cache.Get("pts_created__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Created").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("pts_created__pts_completed__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("pts_created__pts_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		AddMainStatusValidChanges(cache.Get("pts_created__pts_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("pts_on_hold__pts_created__statusvalidchange").(*ent.StatusValidChange)).
		AddCurrentWorkEfforts(cache.Get("9002__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9003__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9005__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9006__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9102__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9103__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9105__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9106__workeffort").(*ent.WorkEffort)).
		AddCurrentWorkEfforts(cache.Get("9202__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created__statusitem", c)

	c = cache.Get("pts_created_ua__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Unassigned").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_ua__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_ua__statusitem", c)

	c = cache.Get("pts_created_as__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Assigned").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_as__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_as__statusitem", c)

	c = cache.Get("pts_created_ip__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CREATED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("In Progress").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created_ip__statusitem: %v", err)
		return err
	}
	cache.Put("pts_created_ip__statusitem", c)

	c = cache.Get("pts_completed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("COMPLETED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Completed").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("pts_created__pts_completed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_completed__statusitem: %v", err)
		return err
	}
	cache.Put("pts_completed__statusitem", c)

	c = cache.Get("pts_on_hold__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ON_HOLD").
		SetSequenceID(common.ParseId("07")).
		SetDescription("On Hold").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("pts_on_hold__pts_created__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("pts_created__pts_on_hold__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_on_hold__statusitem: %v", err)
		return err
	}
	cache.Put("pts_on_hold__statusitem", c)

	c = cache.Get("pts_cancelled__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CANCELLED").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Cancelled").
		SetStatusType(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("pts_created__pts_cancelled__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_cancelled__statusitem: %v", err)
		return err
	}
	cache.Put("pts_cancelled__statusitem", c)

	c = cache.Get("prj_active__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("ACTIVE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Active").
		SetStatusType(cache.Get("project_status__statustype").(*ent.StatusType)).
		AddMainStatusValidChanges(cache.Get("prj_active__prj_closed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prj_active__statusitem: %v", err)
		return err
	}
	cache.Put("prj_active__statusitem", c)

	c = cache.Get("prj_closed__statusitem").(*ent.StatusItem)
	c, err = c.Update().
		SetStatusCode("CLOSED").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Closed").
		SetStatusType(cache.Get("project_status__statustype").(*ent.StatusType)).
		AddToStatusValidChanges(cache.Get("prj_active__prj_closed__statusvalidchange").(*ent.StatusValidChange)).
		AddToStatusValidChanges(cache.Get("_na___prj_closed__statusvalidchange").(*ent.StatusValidChange)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prj_closed__statusitem: %v", err)
		return err
	}
	cache.Put("prj_closed__statusitem", c)

	return nil
}
