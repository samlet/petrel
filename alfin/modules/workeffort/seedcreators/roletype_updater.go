package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateRoleType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RoleType

	c = cache.Get("main_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Role").
		AddChildren(cache.Get("organization_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create main_role__roletype: %v", err)
		return err
	}
	cache.Put("main_role__roletype", c)

	c = cache.Get("account_lead__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Account Lead").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account_lead__roletype: %v", err)
		return err
	}
	cache.Put("account_lead__roletype", c)

	c = cache.Get("admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Administrator").
		AddChildren(cache.Get("ltd_admin__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__roletype: %v", err)
		return err
	}
	cache.Put("admin__roletype", c)

	c = cache.Get("agent__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create agent__roletype: %v", err)
		return err
	}
	cache.Put("agent__roletype", c)

	c = cache.Get("automated_agent_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Automated Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create automated_agent_role__roletype: %v", err)
		return err
	}
	cache.Put("automated_agent_role__roletype", c)

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
		log.Printf("fail to create calendar_role__roletype: %v", err)
		return err
	}
	cache.Put("calendar_role__roletype", c)

	c = cache.Get("client__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client__roletype: %v", err)
		return err
	}
	cache.Put("client__roletype", c)

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
		log.Printf("fail to create commevent_role__roletype: %v", err)
		return err
	}
	cache.Put("commevent_role__roletype", c)

	c = cache.Get("consumer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Consumer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create consumer__roletype: %v", err)
		return err
	}
	cache.Put("consumer__roletype", c)

	c = cache.Get("contractor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Contractor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contractor__roletype: %v", err)
		return err
	}
	cache.Put("contractor__roletype", c)

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
		log.Printf("fail to create customer__roletype: %v", err)
		return err
	}
	cache.Put("customer__roletype", c)

	c = cache.Get("distribution_channel__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Distribution Channel").
		AddChildren(cache.Get("distributor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distribution_channel__roletype: %v", err)
		return err
	}
	cache.Put("distribution_channel__roletype", c)

	c = cache.Get("isp__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("ISP").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create isp__roletype: %v", err)
		return err
	}
	cache.Put("isp__roletype", c)

	c = cache.Get("hosting_server__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Hosting Server").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hosting_server__roletype: %v", err)
		return err
	}
	cache.Put("hosting_server__roletype", c)

	c = cache.Get("manufacturer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manufacturer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manufacturer__roletype: %v", err)
		return err
	}
	cache.Put("manufacturer__roletype", c)

	c = cache.Get("_na___roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___roletype: %v", err)
		return err
	}
	cache.Put("_na___roletype", c)

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
		log.Printf("fail to create organization_role__roletype: %v", err)
		return err
	}
	cache.Put("organization_role__roletype", c)

	c = cache.Get("owner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Owner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create owner__roletype: %v", err)
		return err
	}
	cache.Put("owner__roletype", c)

	c = cache.Get("prospect__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Prospect").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prospect__roletype: %v", err)
		return err
	}
	cache.Put("prospect__roletype", c)

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
		log.Printf("fail to create person_role__roletype: %v", err)
		return err
	}
	cache.Put("person_role__roletype", c)

	c = cache.Get("referrer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Referrer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create referrer__roletype: %v", err)
		return err
	}
	cache.Put("referrer__roletype", c)

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
		log.Printf("fail to create request_role__roletype: %v", err)
		return err
	}
	cache.Put("request_role__roletype", c)

	c = cache.Get("req_manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Manager").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_manager__roletype: %v", err)
		return err
	}
	cache.Put("req_manager__roletype", c)

	c = cache.Get("req_requester__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Requesting Party").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_requester__roletype: %v", err)
		return err
	}
	cache.Put("req_requester__roletype", c)

	c = cache.Get("req_taker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Taker").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_taker__roletype: %v", err)
		return err
	}
	cache.Put("req_taker__roletype", c)

	c = cache.Get("req_respond__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Request Respondent").
		SetParent(cache.Get("request_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_respond__roletype: %v", err)
		return err
	}
	cache.Put("req_respond__roletype", c)

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
		log.Printf("fail to create sfa_role__roletype: %v", err)
		return err
	}
	cache.Put("sfa_role__roletype", c)

	c = cache.Get("shareholder__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shareholder").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shareholder__roletype: %v", err)
		return err
	}
	cache.Put("shareholder__roletype", c)

	c = cache.Get("subscriber__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Subscriber").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create subscriber__roletype: %v", err)
		return err
	}
	cache.Put("subscriber__roletype", c)

	c = cache.Get("vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Vendor").
		AddChildren(cache.Get("bill_from_vendor__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("ship_from_vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vendor__roletype: %v", err)
		return err
	}
	cache.Put("vendor__roletype", c)

	c = cache.Get("visitor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create visitor__roletype: %v", err)
		return err
	}
	cache.Put("visitor__roletype", c)

	c = cache.Get("web_master__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Web Master").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_master__roletype: %v", err)
		return err
	}
	cache.Put("web_master__roletype", c)

	c = cache.Get("workflow_role__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Workflow").
		AddChildren(cache.Get("wf_owner__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workflow_role__roletype: %v", err)
		return err
	}
	cache.Put("workflow_role__roletype", c)

	c = cache.Get("accountant__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Accountant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create accountant__roletype: %v", err)
		return err
	}
	cache.Put("accountant__roletype", c)

	c = cache.Get("account__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Account").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account__roletype: %v", err)
		return err
	}
	cache.Put("account__roletype", c)

	c = cache.Get("addressee__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Addressee").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create addressee__roletype: %v", err)
		return err
	}
	cache.Put("addressee__roletype", c)

	c = cache.Get("association__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Association").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create association__roletype: %v", err)
		return err
	}
	cache.Put("association__roletype", c)

	c = cache.Get("bill_from_vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bill-From Vendor").
		SetParent(cache.Get("vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bill_from_vendor__roletype: %v", err)
		return err
	}
	cache.Put("bill_from_vendor__roletype", c)

	c = cache.Get("bill_to_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bill-To Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bill_to_customer__roletype: %v", err)
		return err
	}
	cache.Put("bill_to_customer__roletype", c)

	c = cache.Get("bcc__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Blind Copy").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bcc__roletype: %v", err)
		return err
	}
	cache.Put("bcc__roletype", c)

	c = cache.Get("bulk_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Bulk Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bulk_customer__roletype: %v", err)
		return err
	}
	cache.Put("bulk_customer__roletype", c)

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
		log.Printf("fail to create cal_attendee__roletype: %v", err)
		return err
	}
	cache.Put("cal_attendee__roletype", c)

	c = cache.Get("cal_delegate__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Delegate").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_delegate__roletype: %v", err)
		return err
	}
	cache.Put("cal_delegate__roletype", c)

	c = cache.Get("cal_host__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Host").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_host__roletype: %v", err)
		return err
	}
	cache.Put("cal_host__roletype", c)

	c = cache.Get("cal_organizer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Calendar Organizer").
		SetParent(cache.Get("calendar_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_organizer__roletype: %v", err)
		return err
	}
	cache.Put("cal_organizer__roletype", c)

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
		log.Printf("fail to create cal_owner__roletype: %v", err)
		return err
	}
	cache.Put("cal_owner__roletype", c)

	c = cache.Get("carrier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Carrier").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create carrier__roletype: %v", err)
		return err
	}
	cache.Put("carrier__roletype", c)

	c = cache.Get("competitor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Competitor").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create competitor__roletype: %v", err)
		return err
	}
	cache.Put("competitor__roletype", c)

	c = cache.Get("contact__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Contact").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contact__roletype: %v", err)
		return err
	}
	cache.Put("contact__roletype", c)

	c = cache.Get("cc__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Carbon Copy").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cc__roletype: %v", err)
		return err
	}
	cache.Put("cc__roletype", c)

	c = cache.Get("originator__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Originator").
		SetParent(cache.Get("commevent_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create originator__roletype: %v", err)
		return err
	}
	cache.Put("originator__roletype", c)

	c = cache.Get("distributor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Distributor").
		SetParent(cache.Get("distribution_channel__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distributor__roletype: %v", err)
		return err
	}
	cache.Put("distributor__roletype", c)

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
		log.Printf("fail to create employee__roletype: %v", err)
		return err
	}
	cache.Put("employee__roletype", c)

	c = cache.Get("end_user_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("End-User Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create end_user_customer__roletype: %v", err)
		return err
	}
	cache.Put("end_user_customer__roletype", c)

	c = cache.Get("household__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Household").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create household__roletype: %v", err)
		return err
	}
	cache.Put("household__roletype", c)

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
		log.Printf("fail to create internal_organizatio__roletype: %v", err)
		return err
	}
	cache.Put("internal_organizatio__roletype", c)

	c = cache.Get("lead__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Lead").
		SetParent(cache.Get("sfa_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead__roletype: %v", err)
		return err
	}
	cache.Put("lead__roletype", c)

	c = cache.Get("ltd_admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Limited Administrator").
		SetParent(cache.Get("admin__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ltd_admin__roletype: %v", err)
		return err
	}
	cache.Put("ltd_admin__roletype", c)

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
		log.Printf("fail to create organization_unit__roletype: %v", err)
		return err
	}
	cache.Put("organization_unit__roletype", c)

	c = cache.Get("partner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Partner").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partner__roletype: %v", err)
		return err
	}
	cache.Put("partner__roletype", c)

	c = cache.Get("placing_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Placing Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create placing_customer__roletype: %v", err)
		return err
	}
	cache.Put("placing_customer__roletype", c)

	c = cache.Get("regulatory_agency__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Regulatory Agency").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create regulatory_agency__roletype: %v", err)
		return err
	}
	cache.Put("regulatory_agency__roletype", c)

	c = cache.Get("sales_rep__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Representative").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		AddChildren(cache.Get("affiliate__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_rep__roletype: %v", err)
		return err
	}
	cache.Put("sales_rep__roletype", c)

	c = cache.Get("ship_from_vendor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ship-From Vendor").
		SetParent(cache.Get("vendor__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_from_vendor__roletype: %v", err)
		return err
	}
	cache.Put("ship_from_vendor__roletype", c)

	c = cache.Get("ship_to_customer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ship-To Customer").
		SetParent(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_to_customer__roletype: %v", err)
		return err
	}
	cache.Put("ship_to_customer__roletype", c)

	c = cache.Get("sponsor__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sponsor").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sponsor__roletype: %v", err)
		return err
	}
	cache.Put("sponsor__roletype", c)

	c = cache.Get("spouse__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Spouse").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create spouse__roletype: %v", err)
		return err
	}
	cache.Put("spouse__roletype", c)

	c = cache.Get("supplier_agent__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Supplier Agent").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier_agent__roletype: %v", err)
		return err
	}
	cache.Put("supplier_agent__roletype", c)

	c = cache.Get("supplier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Supplier").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier__roletype: %v", err)
		return err
	}
	cache.Put("supplier__roletype", c)

	c = cache.Get("tax_authority__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Tax Authority").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tax_authority__roletype: %v", err)
		return err
	}
	cache.Put("tax_authority__roletype", c)

	c = cache.Get("union__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Union").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create union__roletype: %v", err)
		return err
	}
	cache.Put("union__roletype", c)

	c = cache.Get("wf_owner__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Workflow Owner").
		SetParent(cache.Get("workflow_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create wf_owner__roletype: %v", err)
		return err
	}
	cache.Put("wf_owner__roletype", c)

	c = cache.Get("gov_agency__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Government Agency").
		SetParent(cache.Get("organization_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create gov_agency__roletype: %v", err)
		return err
	}
	cache.Put("gov_agency__roletype", c)

	c = cache.Get("affiliate__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Affiliate").
		SetParent(cache.Get("sales_rep__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create affiliate__roletype: %v", err)
		return err
	}
	cache.Put("affiliate__roletype", c)

	c = cache.Get("buyer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Buyer").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create buyer__roletype: %v", err)
		return err
	}
	cache.Put("buyer__roletype", c)

	c = cache.Get("cashier__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Cashier").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cashier__roletype: %v", err)
		return err
	}
	cache.Put("cashier__roletype", c)

	c = cache.Get("department__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Department").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create department__roletype: %v", err)
		return err
	}
	cache.Put("department__roletype", c)

	c = cache.Get("division__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Division").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create division__roletype: %v", err)
		return err
	}
	cache.Put("division__roletype", c)

	c = cache.Get("family_member__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Family Member").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create family_member__roletype: %v", err)
		return err
	}
	cache.Put("family_member__roletype", c)

	c = cache.Get("manager__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manager").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manager__roletype: %v", err)
		return err
	}
	cache.Put("manager__roletype", c)

	c = cache.Get("order_clerk__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Order Clerk").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create order_clerk__roletype: %v", err)
		return err
	}
	cache.Put("order_clerk__roletype", c)

	c = cache.Get("other_internal_organ__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Internal").
		SetParent(cache.Get("internal_organizatio__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_internal_organ__roletype: %v", err)
		return err
	}
	cache.Put("other_internal_organ__roletype", c)

	c = cache.Get("other_organization_u__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Organization Unit").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_organization_u__roletype: %v", err)
		return err
	}
	cache.Put("other_organization_u__roletype", c)

	c = cache.Get("parent_organization__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Parent Organization").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent_organization__roletype: %v", err)
		return err
	}
	cache.Put("parent_organization__roletype", c)

	c = cache.Get("packer__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Packer").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create packer__roletype: %v", err)
		return err
	}
	cache.Put("packer__roletype", c)

	c = cache.Get("picker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Picker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picker__roletype: %v", err)
		return err
	}
	cache.Put("picker__roletype", c)

	c = cache.Get("receiver__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Receiver").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create receiver__roletype: %v", err)
		return err
	}
	cache.Put("receiver__roletype", c)

	c = cache.Get("shipment_clerk__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipment Clerk").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_clerk__roletype: %v", err)
		return err
	}
	cache.Put("shipment_clerk__roletype", c)

	c = cache.Get("stocker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Stocker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create stocker__roletype: %v", err)
		return err
	}
	cache.Put("stocker__roletype", c)

	c = cache.Get("subsidiary__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Subsidiary").
		SetParent(cache.Get("organization_unit__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create subsidiary__roletype: %v", err)
		return err
	}
	cache.Put("subsidiary__roletype", c)

	c = cache.Get("worker__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Worker").
		SetParent(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create worker__roletype: %v", err)
		return err
	}
	cache.Put("worker__roletype", c)

	c = cache.Get("email_admin__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Email Administrator").
		SetParent(cache.Get("person_role__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_admin__roletype: %v", err)
		return err
	}
	cache.Put("email_admin__roletype", c)

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
		log.Printf("fail to create project_team__roletype: %v", err)
		return err
	}
	cache.Put("project_team__roletype", c)

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
		log.Printf("fail to create client_manager__roletype: %v", err)
		return err
	}
	cache.Put("client_manager__roletype", c)

	c = cache.Get("client_analyst__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Client Analyst").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client_analyst__roletype: %v", err)
		return err
	}
	cache.Put("client_analyst__roletype", c)

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
		log.Printf("fail to create client_billing__roletype: %v", err)
		return err
	}
	cache.Put("client_billing__roletype", c)

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
		log.Printf("fail to create provider_manager__roletype: %v", err)
		return err
	}
	cache.Put("provider_manager__roletype", c)

	c = cache.Get("provider_accounting__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Accounting").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_accounting__roletype: %v", err)
		return err
	}
	cache.Put("provider_accounting__roletype", c)

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
		log.Printf("fail to create provider_analyst__roletype: %v", err)
		return err
	}
	cache.Put("provider_analyst__roletype", c)

	c = cache.Get("provider_validator__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Validator").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_validator__roletype: %v", err)
		return err
	}
	cache.Put("provider_validator__roletype", c)

	c = cache.Get("provider_functional__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Func. Impl").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_functional__roletype: %v", err)
		return err
	}
	cache.Put("provider_functional__roletype", c)

	c = cache.Get("provider_tester__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Provider Tester").
		SetParent(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_tester__roletype: %v", err)
		return err
	}
	cache.Put("provider_tester__roletype", c)

	return nil
}
