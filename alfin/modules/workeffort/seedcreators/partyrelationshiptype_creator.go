package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyRelationshipType(ctx context.Context) error {
	log.Println("PartyRelationshipType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyRelationshipType

	c, err = client.PartyRelationshipType.Create().SetStringRef("agent__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create agent__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("agent__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("child__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Child").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create child__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("child__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("contact_rel__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Contact Relation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contact_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("contact_rel__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("lead_rel__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Lead Relation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("lead_rel__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("lead_owner__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Lead Owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead_owner__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("lead_owner__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("customer_rel__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Customer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create customer_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("customer_rel__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("distribution_channel__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Distributor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distribution_channel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("distribution_channel__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("employment__partyrelationshiptype").
		SetHasTable("Yes").
		SetPartyRelationshipName("Employee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create employment__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("employment__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("friend__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Friend").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create friend__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("friend__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("group_rollup__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Group Member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create group_rollup__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("group_rollup__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("host_server_visitor__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Host Server Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create host_server_visitor__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("host_server_visitor__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("visitor_isp__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("ISP Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create visitor_isp__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("visitor_isp__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("manager__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manager__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("manager__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("parent__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Parent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("parent__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("partnership__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Partner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partnership__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("partnership__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("sales_affiliate__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Sales Affiliate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_affiliate__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("sales_affiliate__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("spouse__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Spouse").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create spouse__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("spouse__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("supplier_rel__partyrelationshiptype").
		SetHasTable("Yes").
		SetPartyRelationshipName("Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("supplier_rel__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("web_master_assignmen__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Web Master").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_master_assignmen__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("web_master_assignmen__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("account__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Account owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("account__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("assistant__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Assistant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create assistant__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("assistant__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("owner__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create owner__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("owner__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("parent_account__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Parent Account").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent_account__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("parent_account__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("reports_to__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Reports To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create reports_to__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("reports_to__partyrelationshiptype", c)

	c, err = client.PartyRelationshipType.Create().SetStringRef("lead_own_grp_member__partyrelationshiptype").
		SetHasTable("No").
		SetPartyRelationshipName("Lead Owners/Managers").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead_own_grp_member__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("lead_own_grp_member__partyrelationshiptype", c)

	return nil
}
