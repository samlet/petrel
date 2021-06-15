package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateRoleType(ctx context.Context) error {
	log.Println("RoleType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RoleType

	c, err = client.RoleType.Create().SetStringRef("main_role__roletype").
		SetHasTable("No").
		SetDescription("Main Role").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create main_role__roletype: %v", err)
		return err
	}
	cache.Put("main_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("account_lead__roletype").
		SetHasTable("No").
		SetDescription("Account Lead").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account_lead__roletype: %v", err)
		return err
	}
	cache.Put("account_lead__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("admin__roletype").
		SetHasTable("No").
		SetDescription("Administrator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__roletype: %v", err)
		return err
	}
	cache.Put("admin__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("agent__roletype").
		SetHasTable("No").
		SetDescription("Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create agent__roletype: %v", err)
		return err
	}
	cache.Put("agent__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("automated_agent_role__roletype").
		SetHasTable("No").
		SetDescription("Automated Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create automated_agent_role__roletype: %v", err)
		return err
	}
	cache.Put("automated_agent_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("calendar_role__roletype").
		SetHasTable("No").
		SetDescription("Calendar").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create calendar_role__roletype: %v", err)
		return err
	}
	cache.Put("calendar_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("client__roletype").
		SetHasTable("No").
		SetDescription("Client").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client__roletype: %v", err)
		return err
	}
	cache.Put("client__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("commevent_role__roletype").
		SetHasTable("No").
		SetDescription("Communication Participant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create commevent_role__roletype: %v", err)
		return err
	}
	cache.Put("commevent_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("consumer__roletype").
		SetHasTable("No").
		SetDescription("Consumer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create consumer__roletype: %v", err)
		return err
	}
	cache.Put("consumer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("contractor__roletype").
		SetHasTable("No").
		SetDescription("Contractor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contractor__roletype: %v", err)
		return err
	}
	cache.Put("contractor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("customer__roletype").
		SetHasTable("No").
		SetDescription("Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create customer__roletype: %v", err)
		return err
	}
	cache.Put("customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("distribution_channel__roletype").
		SetHasTable("No").
		SetDescription("Distribution Channel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distribution_channel__roletype: %v", err)
		return err
	}
	cache.Put("distribution_channel__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("isp__roletype").
		SetHasTable("No").
		SetDescription("ISP").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create isp__roletype: %v", err)
		return err
	}
	cache.Put("isp__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("hosting_server__roletype").
		SetHasTable("No").
		SetDescription("Hosting Server").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hosting_server__roletype: %v", err)
		return err
	}
	cache.Put("hosting_server__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("manufacturer__roletype").
		SetHasTable("No").
		SetDescription("Manufacturer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manufacturer__roletype: %v", err)
		return err
	}
	cache.Put("manufacturer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("_na___roletype").
		SetHasTable("No").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___roletype: %v", err)
		return err
	}
	cache.Put("_na___roletype", c)

	c, err = client.RoleType.Create().SetStringRef("organization_role__roletype").
		SetHasTable("No").
		SetDescription("Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create organization_role__roletype: %v", err)
		return err
	}
	cache.Put("organization_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("owner__roletype").
		SetHasTable("No").
		SetDescription("Owner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create owner__roletype: %v", err)
		return err
	}
	cache.Put("owner__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("prospect__roletype").
		SetHasTable("No").
		SetDescription("Prospect").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prospect__roletype: %v", err)
		return err
	}
	cache.Put("prospect__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("person_role__roletype").
		SetHasTable("No").
		SetDescription("Person").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create person_role__roletype: %v", err)
		return err
	}
	cache.Put("person_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("referrer__roletype").
		SetHasTable("No").
		SetDescription("Referrer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create referrer__roletype: %v", err)
		return err
	}
	cache.Put("referrer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("request_role__roletype").
		SetHasTable("No").
		SetDescription("Request Role").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create request_role__roletype: %v", err)
		return err
	}
	cache.Put("request_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("req_manager__roletype").
		SetHasTable("No").
		SetDescription("Request Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_manager__roletype: %v", err)
		return err
	}
	cache.Put("req_manager__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("req_requester__roletype").
		SetHasTable("No").
		SetDescription("Requesting Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_requester__roletype: %v", err)
		return err
	}
	cache.Put("req_requester__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("req_taker__roletype").
		SetHasTable("No").
		SetDescription("Request Taker").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_taker__roletype: %v", err)
		return err
	}
	cache.Put("req_taker__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("req_respond__roletype").
		SetHasTable("No").
		SetDescription("Request Respondent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create req_respond__roletype: %v", err)
		return err
	}
	cache.Put("req_respond__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("sfa_role__roletype").
		SetHasTable("No").
		SetDescription("Sales Force Autm.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sfa_role__roletype: %v", err)
		return err
	}
	cache.Put("sfa_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("shareholder__roletype").
		SetHasTable("No").
		SetDescription("Shareholder").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shareholder__roletype: %v", err)
		return err
	}
	cache.Put("shareholder__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("subscriber__roletype").
		SetHasTable("No").
		SetDescription("Subscriber").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create subscriber__roletype: %v", err)
		return err
	}
	cache.Put("subscriber__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("vendor__roletype").
		SetHasTable("No").
		SetDescription("Vendor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vendor__roletype: %v", err)
		return err
	}
	cache.Put("vendor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("visitor__roletype").
		SetHasTable("No").
		SetDescription("Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create visitor__roletype: %v", err)
		return err
	}
	cache.Put("visitor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("web_master__roletype").
		SetHasTable("No").
		SetDescription("Web Master").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_master__roletype: %v", err)
		return err
	}
	cache.Put("web_master__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("workflow_role__roletype").
		SetHasTable("No").
		SetDescription("Workflow").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workflow_role__roletype: %v", err)
		return err
	}
	cache.Put("workflow_role__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("accountant__roletype").
		SetHasTable("No").
		SetDescription("Accountant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create accountant__roletype: %v", err)
		return err
	}
	cache.Put("accountant__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("account__roletype").
		SetHasTable("No").
		SetDescription("Account").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account__roletype: %v", err)
		return err
	}
	cache.Put("account__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("addressee__roletype").
		SetHasTable("No").
		SetDescription("Addressee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create addressee__roletype: %v", err)
		return err
	}
	cache.Put("addressee__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("association__roletype").
		SetHasTable("No").
		SetDescription("Association").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create association__roletype: %v", err)
		return err
	}
	cache.Put("association__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("bill_from_vendor__roletype").
		SetHasTable("No").
		SetDescription("Bill-From Vendor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bill_from_vendor__roletype: %v", err)
		return err
	}
	cache.Put("bill_from_vendor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("bill_to_customer__roletype").
		SetHasTable("No").
		SetDescription("Bill-To Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bill_to_customer__roletype: %v", err)
		return err
	}
	cache.Put("bill_to_customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("bcc__roletype").
		SetHasTable("No").
		SetDescription("Blind Copy").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bcc__roletype: %v", err)
		return err
	}
	cache.Put("bcc__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("bulk_customer__roletype").
		SetHasTable("No").
		SetDescription("Bulk Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bulk_customer__roletype: %v", err)
		return err
	}
	cache.Put("bulk_customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cal_attendee__roletype").
		SetHasTable("No").
		SetDescription("Calendar Attendee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_attendee__roletype: %v", err)
		return err
	}
	cache.Put("cal_attendee__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cal_delegate__roletype").
		SetHasTable("No").
		SetDescription("Calendar Delegate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_delegate__roletype: %v", err)
		return err
	}
	cache.Put("cal_delegate__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cal_host__roletype").
		SetHasTable("No").
		SetDescription("Calendar Host").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_host__roletype: %v", err)
		return err
	}
	cache.Put("cal_host__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cal_organizer__roletype").
		SetHasTable("No").
		SetDescription("Calendar Organizer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_organizer__roletype: %v", err)
		return err
	}
	cache.Put("cal_organizer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cal_owner__roletype").
		SetHasTable("No").
		SetDescription("Calendar Owner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cal_owner__roletype: %v", err)
		return err
	}
	cache.Put("cal_owner__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("carrier__roletype").
		SetHasTable("No").
		SetDescription("Carrier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create carrier__roletype: %v", err)
		return err
	}
	cache.Put("carrier__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("competitor__roletype").
		SetHasTable("No").
		SetDescription("Competitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create competitor__roletype: %v", err)
		return err
	}
	cache.Put("competitor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("contact__roletype").
		SetHasTable("No").
		SetDescription("Contact").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contact__roletype: %v", err)
		return err
	}
	cache.Put("contact__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cc__roletype").
		SetHasTable("No").
		SetDescription("Carbon Copy").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cc__roletype: %v", err)
		return err
	}
	cache.Put("cc__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("originator__roletype").
		SetHasTable("No").
		SetDescription("Originator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create originator__roletype: %v", err)
		return err
	}
	cache.Put("originator__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("distributor__roletype").
		SetHasTable("No").
		SetDescription("Distributor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distributor__roletype: %v", err)
		return err
	}
	cache.Put("distributor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("employee__roletype").
		SetHasTable("No").
		SetDescription("Employee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create employee__roletype: %v", err)
		return err
	}
	cache.Put("employee__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("end_user_customer__roletype").
		SetHasTable("No").
		SetDescription("End-User Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create end_user_customer__roletype: %v", err)
		return err
	}
	cache.Put("end_user_customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("household__roletype").
		SetHasTable("No").
		SetDescription("Household").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create household__roletype: %v", err)
		return err
	}
	cache.Put("household__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("internal_organizatio__roletype").
		SetHasTable("Yes").
		SetDescription("Internal Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create internal_organizatio__roletype: %v", err)
		return err
	}
	cache.Put("internal_organizatio__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("lead__roletype").
		SetHasTable("No").
		SetDescription("Lead").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead__roletype: %v", err)
		return err
	}
	cache.Put("lead__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("ltd_admin__roletype").
		SetHasTable("No").
		SetDescription("Limited Administrator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ltd_admin__roletype: %v", err)
		return err
	}
	cache.Put("ltd_admin__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("organization_unit__roletype").
		SetHasTable("No").
		SetDescription("Organization Unit").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create organization_unit__roletype: %v", err)
		return err
	}
	cache.Put("organization_unit__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("partner__roletype").
		SetHasTable("No").
		SetDescription("Partner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partner__roletype: %v", err)
		return err
	}
	cache.Put("partner__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("placing_customer__roletype").
		SetHasTable("No").
		SetDescription("Placing Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create placing_customer__roletype: %v", err)
		return err
	}
	cache.Put("placing_customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("regulatory_agency__roletype").
		SetHasTable("No").
		SetDescription("Regulatory Agency").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create regulatory_agency__roletype: %v", err)
		return err
	}
	cache.Put("regulatory_agency__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("sales_rep__roletype").
		SetHasTable("No").
		SetDescription("Sales Representative").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_rep__roletype: %v", err)
		return err
	}
	cache.Put("sales_rep__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("ship_from_vendor__roletype").
		SetHasTable("No").
		SetDescription("Ship-From Vendor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_from_vendor__roletype: %v", err)
		return err
	}
	cache.Put("ship_from_vendor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("ship_to_customer__roletype").
		SetHasTable("No").
		SetDescription("Ship-To Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_to_customer__roletype: %v", err)
		return err
	}
	cache.Put("ship_to_customer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("sponsor__roletype").
		SetHasTable("No").
		SetDescription("Sponsor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sponsor__roletype: %v", err)
		return err
	}
	cache.Put("sponsor__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("spouse__roletype").
		SetHasTable("No").
		SetDescription("Spouse").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create spouse__roletype: %v", err)
		return err
	}
	cache.Put("spouse__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("supplier_agent__roletype").
		SetHasTable("No").
		SetDescription("Supplier Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier_agent__roletype: %v", err)
		return err
	}
	cache.Put("supplier_agent__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("supplier__roletype").
		SetHasTable("No").
		SetDescription("Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier__roletype: %v", err)
		return err
	}
	cache.Put("supplier__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("tax_authority__roletype").
		SetHasTable("No").
		SetDescription("Tax Authority").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create tax_authority__roletype: %v", err)
		return err
	}
	cache.Put("tax_authority__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("union__roletype").
		SetHasTable("No").
		SetDescription("Union").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create union__roletype: %v", err)
		return err
	}
	cache.Put("union__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("wf_owner__roletype").
		SetHasTable("No").
		SetDescription("Workflow Owner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create wf_owner__roletype: %v", err)
		return err
	}
	cache.Put("wf_owner__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("gov_agency__roletype").
		SetHasTable("No").
		SetDescription("Government Agency").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create gov_agency__roletype: %v", err)
		return err
	}
	cache.Put("gov_agency__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("affiliate__roletype").
		SetHasTable("No").
		SetDescription("Affiliate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create affiliate__roletype: %v", err)
		return err
	}
	cache.Put("affiliate__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("buyer__roletype").
		SetHasTable("No").
		SetDescription("Buyer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create buyer__roletype: %v", err)
		return err
	}
	cache.Put("buyer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("cashier__roletype").
		SetHasTable("No").
		SetDescription("Cashier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cashier__roletype: %v", err)
		return err
	}
	cache.Put("cashier__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("department__roletype").
		SetHasTable("No").
		SetDescription("Department").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create department__roletype: %v", err)
		return err
	}
	cache.Put("department__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("division__roletype").
		SetHasTable("No").
		SetDescription("Division").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create division__roletype: %v", err)
		return err
	}
	cache.Put("division__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("family_member__roletype").
		SetHasTable("No").
		SetDescription("Family Member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create family_member__roletype: %v", err)
		return err
	}
	cache.Put("family_member__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("manager__roletype").
		SetHasTable("No").
		SetDescription("Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manager__roletype: %v", err)
		return err
	}
	cache.Put("manager__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("order_clerk__roletype").
		SetHasTable("No").
		SetDescription("Order Clerk").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create order_clerk__roletype: %v", err)
		return err
	}
	cache.Put("order_clerk__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("other_internal_organ__roletype").
		SetHasTable("No").
		SetDescription("Other Internal").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_internal_organ__roletype: %v", err)
		return err
	}
	cache.Put("other_internal_organ__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("other_organization_u__roletype").
		SetHasTable("No").
		SetDescription("Other Organization Unit").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_organization_u__roletype: %v", err)
		return err
	}
	cache.Put("other_organization_u__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("parent_organization__roletype").
		SetHasTable("No").
		SetDescription("Parent Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent_organization__roletype: %v", err)
		return err
	}
	cache.Put("parent_organization__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("packer__roletype").
		SetHasTable("No").
		SetDescription("Packer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create packer__roletype: %v", err)
		return err
	}
	cache.Put("packer__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("picker__roletype").
		SetHasTable("No").
		SetDescription("Picker").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create picker__roletype: %v", err)
		return err
	}
	cache.Put("picker__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("receiver__roletype").
		SetHasTable("No").
		SetDescription("Receiver").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create receiver__roletype: %v", err)
		return err
	}
	cache.Put("receiver__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("shipment_clerk__roletype").
		SetHasTable("No").
		SetDescription("Shipment Clerk").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipment_clerk__roletype: %v", err)
		return err
	}
	cache.Put("shipment_clerk__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("stocker__roletype").
		SetHasTable("No").
		SetDescription("Stocker").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create stocker__roletype: %v", err)
		return err
	}
	cache.Put("stocker__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("subsidiary__roletype").
		SetHasTable("No").
		SetDescription("Subsidiary").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create subsidiary__roletype: %v", err)
		return err
	}
	cache.Put("subsidiary__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("worker__roletype").
		SetHasTable("No").
		SetDescription("Worker").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create worker__roletype: %v", err)
		return err
	}
	cache.Put("worker__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("email_admin__roletype").
		SetHasTable("No").
		SetDescription("Email Administrator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_admin__roletype: %v", err)
		return err
	}
	cache.Put("email_admin__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("project_team__roletype").
		SetHasTable("No").
		SetDescription("Project Team").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_team__roletype: %v", err)
		return err
	}
	cache.Put("project_team__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("client_manager__roletype").
		SetHasTable("No").
		SetDescription("Client Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client_manager__roletype: %v", err)
		return err
	}
	cache.Put("client_manager__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("client_analyst__roletype").
		SetHasTable("No").
		SetDescription("Client Analyst").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client_analyst__roletype: %v", err)
		return err
	}
	cache.Put("client_analyst__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("client_billing__roletype").
		SetHasTable("No").
		SetDescription("Client Billing").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create client_billing__roletype: %v", err)
		return err
	}
	cache.Put("client_billing__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_manager__roletype").
		SetHasTable("No").
		SetDescription("Provider Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_manager__roletype: %v", err)
		return err
	}
	cache.Put("provider_manager__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_accounting__roletype").
		SetHasTable("No").
		SetDescription("Provider Accounting").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_accounting__roletype: %v", err)
		return err
	}
	cache.Put("provider_accounting__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_analyst__roletype").
		SetHasTable("No").
		SetDescription("Provider Analyst").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_analyst__roletype: %v", err)
		return err
	}
	cache.Put("provider_analyst__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_validator__roletype").
		SetHasTable("No").
		SetDescription("Provider Validator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_validator__roletype: %v", err)
		return err
	}
	cache.Put("provider_validator__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_functional__roletype").
		SetHasTable("No").
		SetDescription("Provider Func. Impl").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_functional__roletype: %v", err)
		return err
	}
	cache.Put("provider_functional__roletype", c)

	c, err = client.RoleType.Create().SetStringRef("provider_tester__roletype").
		SetHasTable("No").
		SetDescription("Provider Tester").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create provider_tester__roletype: %v", err)
		return err
	}
	cache.Put("provider_tester__roletype", c)

	return nil
}
