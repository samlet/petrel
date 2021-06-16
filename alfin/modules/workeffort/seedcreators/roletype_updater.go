package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateRoleType(ctx context.Context) error {
	log.Println("RoleType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RoleType
	failures := 0

	c = cache.Get("main_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Role").
		AddChildren(cache.Get("organization_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update main_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("main_role__roletype", c)
	}

	c = cache.Get("account_lead__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Account Lead").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update account_lead__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("account_lead__roletype", c)
	}

	c = cache.Get("admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Administrator").
		AddChildren(cache.Get("ltd_admin__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update admin__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("admin__roletype", c)
	}

	c = cache.Get("agent__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update agent__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("agent__roletype", c)
	}

	c = cache.Get("automated_agent_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Automated Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update automated_agent_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("automated_agent_role__roletype", c)
	}

	c = cache.Get("calendar_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar").
		AddChildren(cache.Get("cal_attendee__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cal_delegate__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cal_host__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cal_organizer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cal_owner__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update calendar_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("calendar_role__roletype", c)
	}

	c = cache.Get("client__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update client__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("client__roletype", c)
	}

	c = cache.Get("commevent_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Communication Participant").
		AddChildren(cache.Get("addressee__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("bcc__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cc__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("originator__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update commevent_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("commevent_role__roletype", c)
	}

	c = cache.Get("consumer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Consumer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update consumer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("consumer__roletype", c)
	}

	c = cache.Get("contractor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Contractor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update contractor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("contractor__roletype", c)
	}

	c = cache.Get("customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Customer").
		AddPartyContactMeches(cache.Get("democustomer1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddValidFromPartyRelationshipTypes(cache.Get("customer_rel__partyrelationshiptype").(*ent.PartyRelationshipType)).
		AddPartyRoles(cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)).
		AddChildren(cache.Get("bill_to_customer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("bulk_customer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("end_user_customer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("placing_customer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("ship_to_customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("customer__roletype", c)
	}

	c = cache.Get("distribution_channel__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Distribution Channel").
		AddChildren(cache.Get("distributor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update distribution_channel__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("distribution_channel__roletype", c)
	}

	c = cache.Get("isp__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("ISP").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update isp__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("isp__roletype", c)
	}

	c = cache.Get("hosting_server__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Hosting Server").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hosting_server__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hosting_server__roletype", c)
	}

	c = cache.Get("manufacturer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manufacturer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update manufacturer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("manufacturer__roletype", c)
	}

	c = cache.Get("_na___roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update _na___roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("_na___roletype", c)
	}

	c = cache.Get("organization_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Organization").
		SetParent(cache.Get("main_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("association__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("carrier__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("competitor__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("household__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("internal_organizatio__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("partner__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("regulatory_agency__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("supplier__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("tax_authority__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("union__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("gov_agency__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update organization_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("organization_role__roletype", c)
	}

	c = cache.Get("owner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Owner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update owner__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("owner__roletype", c)
	}

	c = cache.Get("prospect__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Prospect").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prospect__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prospect__roletype", c)
	}

	c = cache.Get("person_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Person").
		AddChildren(cache.Get("employee__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("sales_rep__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("sponsor__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("spouse__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("supplier_agent__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("family_member__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("email_admin__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update person_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("person_role__roletype", c)
	}

	c = cache.Get("referrer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Referrer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update referrer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("referrer__roletype", c)
	}

	c = cache.Get("request_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Role").
		AddChildren(cache.Get("req_manager__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("req_requester__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("req_taker__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("req_respond__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update request_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("request_role__roletype", c)
	}

	c = cache.Get("req_manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Manager").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update req_manager__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("req_manager__roletype", c)
	}

	c = cache.Get("req_requester__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Requesting Party").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update req_requester__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("req_requester__roletype", c)
	}

	c = cache.Get("req_taker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Taker").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update req_taker__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("req_taker__roletype", c)
	}

	c = cache.Get("req_respond__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Respondent").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update req_respond__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("req_respond__roletype", c)
	}

	c = cache.Get("sfa_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Force Autm.").
		SetParent(cache.Get("main_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("account__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("contact__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("lead__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sfa_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sfa_role__roletype", c)
	}

	c = cache.Get("shareholder__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shareholder").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shareholder__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shareholder__roletype", c)
	}

	c = cache.Get("subscriber__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Subscriber").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update subscriber__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("subscriber__roletype", c)
	}

	c = cache.Get("vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Vendor").
		AddChildren(cache.Get("bill_from_vendor__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("ship_from_vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update vendor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("vendor__roletype", c)
	}

	c = cache.Get("visitor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update visitor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("visitor__roletype", c)
	}

	c = cache.Get("web_master__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Web Master").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update web_master__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("web_master__roletype", c)
	}

	c = cache.Get("workflow_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Workflow").
		AddChildren(cache.Get("wf_owner__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workflow_role__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("workflow_role__roletype", c)
	}

	c = cache.Get("accountant__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Accountant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update accountant__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("accountant__roletype", c)
	}

	c = cache.Get("account__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Account").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update account__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("account__roletype", c)
	}

	c = cache.Get("addressee__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Addressee").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update addressee__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("addressee__roletype", c)
	}

	c = cache.Get("association__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Association").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update association__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("association__roletype", c)
	}

	c = cache.Get("bill_from_vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bill-From Vendor").
		SetParent(cache.Get("vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bill_from_vendor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("bill_from_vendor__roletype", c)
	}

	c = cache.Get("bill_to_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bill-To Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bill_to_customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("bill_to_customer__roletype", c)
	}

	c = cache.Get("bcc__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Blind Copy").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bcc__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("bcc__roletype", c)
	}

	c = cache.Get("bulk_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bulk Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bulk_customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("bulk_customer__roletype", c)
	}

	c = cache.Get("cal_attendee__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Attendee").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("admin__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__cal_attendee__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cal_attendee__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cal_attendee__roletype", c)
	}

	c = cache.Get("cal_delegate__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Delegate").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cal_delegate__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cal_delegate__roletype", c)
	}

	c = cache.Get("cal_host__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Host").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cal_host__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cal_host__roletype", c)
	}

	c = cache.Get("cal_organizer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Organizer").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cal_organizer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cal_organizer__roletype", c)
	}

	c = cache.Get("cal_owner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Owner").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("admin__cal_owner__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cal_owner__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cal_owner__roletype", c)
	}

	c = cache.Get("carrier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Carrier").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update carrier__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("carrier__roletype", c)
	}

	c = cache.Get("competitor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Competitor").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update competitor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("competitor__roletype", c)
	}

	c = cache.Get("contact__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Contact").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update contact__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("contact__roletype", c)
	}

	c = cache.Get("cc__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Carbon Copy").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cc__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cc__roletype", c)
	}

	c = cache.Get("originator__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Originator").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update originator__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("originator__roletype", c)
	}

	c = cache.Get("distributor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Distributor").
		SetParent(cache.Get("distribution_channel__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update distributor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("distributor__roletype", c)
	}

	c = cache.Get("employee__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Employee").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("demoemployee__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		AddChildren(cache.Get("buyer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("cashier__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("manager__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("order_clerk__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("packer__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("picker__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("receiver__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("shipment_clerk__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("stocker__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("worker__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update employee__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("employee__roletype", c)
	}

	c = cache.Get("end_user_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("End-User Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update end_user_customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("end_user_customer__roletype", c)
	}

	c = cache.Get("household__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Household").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update household__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("household__roletype", c)
	}

	c = cache.Get("internal_organizatio__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Internal Organization").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("other_internal_organ__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update internal_organizatio__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("internal_organizatio__roletype", c)
	}

	c = cache.Get("lead__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Lead").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update lead__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("lead__roletype", c)
	}

	c = cache.Get("ltd_admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Limited Administrator").
		SetParent(cache.Get("admin__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ltd_admin__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ltd_admin__roletype", c)
	}

	c = cache.Get("organization_unit__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Organization Unit").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("department__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("division__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("other_organization_u__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("parent_organization__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("subsidiary__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update organization_unit__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("organization_unit__roletype", c)
	}

	c = cache.Get("partner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Partner").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partner__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partner__roletype", c)
	}

	c = cache.Get("placing_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Placing Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update placing_customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("placing_customer__roletype", c)
	}

	c = cache.Get("regulatory_agency__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Regulatory Agency").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update regulatory_agency__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("regulatory_agency__roletype", c)
	}

	c = cache.Get("sales_rep__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Representative").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("affiliate__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sales_rep__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sales_rep__roletype", c)
	}

	c = cache.Get("ship_from_vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ship-From Vendor").
		SetParent(cache.Get("vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_from_vendor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_from_vendor__roletype", c)
	}

	c = cache.Get("ship_to_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ship-To Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_to_customer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_to_customer__roletype", c)
	}

	c = cache.Get("sponsor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sponsor").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sponsor__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sponsor__roletype", c)
	}

	c = cache.Get("spouse__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Spouse").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update spouse__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("spouse__roletype", c)
	}

	c = cache.Get("supplier_agent__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Supplier Agent").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update supplier_agent__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("supplier_agent__roletype", c)
	}

	c = cache.Get("supplier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Supplier").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update supplier__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("supplier__roletype", c)
	}

	c = cache.Get("tax_authority__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Tax Authority").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update tax_authority__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("tax_authority__roletype", c)
	}

	c = cache.Get("union__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Union").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update union__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("union__roletype", c)
	}

	c = cache.Get("wf_owner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Workflow Owner").
		SetParent(cache.Get("workflow_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update wf_owner__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("wf_owner__roletype", c)
	}

	c = cache.Get("gov_agency__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Government Agency").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update gov_agency__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("gov_agency__roletype", c)
	}

	c = cache.Get("affiliate__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Affiliate").
		SetParent(cache.Get("sales_rep__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update affiliate__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("affiliate__roletype", c)
	}

	c = cache.Get("buyer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Buyer").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update buyer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("buyer__roletype", c)
	}

	c = cache.Get("cashier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Cashier").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cashier__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cashier__roletype", c)
	}

	c = cache.Get("department__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Department").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update department__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("department__roletype", c)
	}

	c = cache.Get("division__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Division").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update division__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("division__roletype", c)
	}

	c = cache.Get("family_member__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Family Member").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update family_member__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("family_member__roletype", c)
	}

	c = cache.Get("manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manager").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update manager__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("manager__roletype", c)
	}

	c = cache.Get("order_clerk__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Order Clerk").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update order_clerk__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("order_clerk__roletype", c)
	}

	c = cache.Get("other_internal_organ__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Internal").
		SetParent(cache.Get("internal_organizatio__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_internal_organ__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_internal_organ__roletype", c)
	}

	c = cache.Get("other_organization_u__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Organization Unit").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_organization_u__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_organization_u__roletype", c)
	}

	c = cache.Get("parent_organization__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Parent Organization").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update parent_organization__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("parent_organization__roletype", c)
	}

	c = cache.Get("packer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Packer").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update packer__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("packer__roletype", c)
	}

	c = cache.Get("picker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Picker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update picker__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("picker__roletype", c)
	}

	c = cache.Get("receiver__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Receiver").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update receiver__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("receiver__roletype", c)
	}

	c = cache.Get("shipment_clerk__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipment Clerk").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update shipment_clerk__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("shipment_clerk__roletype", c)
	}

	c = cache.Get("stocker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Stocker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update stocker__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("stocker__roletype", c)
	}

	c = cache.Get("subsidiary__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Subsidiary").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update subsidiary__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("subsidiary__roletype", c)
	}

	c = cache.Get("worker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Worker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update worker__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("worker__roletype", c)
	}

	c = cache.Get("email_admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Email Administrator").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update email_admin__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("email_admin__roletype", c)
	}

	c = cache.Get("project_team__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Project Team").
		AddPartyRoles(cache.Get("demoemployee__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer1__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer2__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer3__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("admin__project_team__partyrole").(*ent.PartyRole)).
		AddChildren(cache.Get("client_manager__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("client_analyst__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("client_billing__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_manager__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_accounting__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_analyst__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_validator__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_functional__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("provider_tester__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_team__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("project_team__roletype", c)
	}

	c = cache.Get("client_manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client Manager").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("democustomer1__client_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer2__client_manager__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update client_manager__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("client_manager__roletype", c)
	}

	c = cache.Get("client_analyst__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client Analyst").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update client_analyst__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("client_analyst__roletype", c)
	}

	c = cache.Get("client_billing__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client Billing").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("democustomer3__client_billing__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustcompany__client_billing__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update client_billing__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("client_billing__roletype", c)
	}

	c = cache.Get("provider_manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Manager").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("admin__provider_manager__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("9000__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9002__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_manager__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_manager__roletype", c)
	}

	c = cache.Get("provider_accounting__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Accounting").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_accounting__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_accounting__roletype", c)
	}

	c = cache.Get("provider_analyst__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Analyst").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		AddPartyRoles(cache.Get("demoemployee__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_analyst__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_analyst__roletype", c)
	}

	c = cache.Get("provider_validator__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Validator").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_validator__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_validator__roletype", c)
	}

	c = cache.Get("provider_functional__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Func. Impl").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_functional__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_functional__roletype", c)
	}

	c = cache.Get("provider_tester__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Tester").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update provider_tester__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("provider_tester__roletype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
