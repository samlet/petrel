package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyRelationshipType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyRelationshipType

	c = cache.Get("agent__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create agent__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("agent__partyrelationshiptype", c)

	c = cache.Get("child__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Child").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create child__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("child__partyrelationshiptype", c)

	c = cache.Get("contact_rel__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Contact Relation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create contact_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("contact_rel__partyrelationshiptype", c)

	c = cache.Get("lead_rel__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Lead Relation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("lead_rel__partyrelationshiptype", c)

	c = cache.Get("lead_owner__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Lead Owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lead_owner__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("lead_owner__partyrelationshiptype", c)

	c = cache.Get("customer_rel__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Customer").
		SetValidFromRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create customer_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("customer_rel__partyrelationshiptype", c)

	c = cache.Get("distribution_channel__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Distributor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distribution_channel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("distribution_channel__partyrelationshiptype", c)

	c = cache.Get("employment__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetPartyRelationshipName("Employee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create employment__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("employment__partyrelationshiptype", c)

	c = cache.Get("friend__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Friend").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create friend__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("friend__partyrelationshiptype", c)

	c = cache.Get("group_rollup__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Group Member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create group_rollup__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("group_rollup__partyrelationshiptype", c)

	c = cache.Get("host_server_visitor__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Host Server Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create host_server_visitor__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("host_server_visitor__partyrelationshiptype", c)

	c = cache.Get("visitor_isp__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("ISP Visitor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create visitor_isp__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("visitor_isp__partyrelationshiptype", c)

	c = cache.Get("manager__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manager__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("manager__partyrelationshiptype", c)

	c = cache.Get("parent__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Parent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("parent__partyrelationshiptype", c)

	c = cache.Get("partnership__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Partner").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partnership__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("partnership__partyrelationshiptype", c)

	c = cache.Get("sales_affiliate__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Sales Affiliate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_affiliate__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("sales_affiliate__partyrelationshiptype", c)

	c = cache.Get("spouse__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Spouse").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create spouse__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("spouse__partyrelationshiptype", c)

	c = cache.Get("supplier_rel__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetPartyRelationshipName("Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create supplier_rel__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("supplier_rel__partyrelationshiptype", c)

	c = cache.Get("web_master_assignmen__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Web Master").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_master_assignmen__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("web_master_assignmen__partyrelationshiptype", c)

	c = cache.Get("account__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Account owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create account__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("account__partyrelationshiptype", c)

	c = cache.Get("assistant__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Assistant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create assistant__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("assistant__partyrelationshiptype", c)

	c = cache.Get("owner__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Owned by").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create owner__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("owner__partyrelationshiptype", c)

	c = cache.Get("parent_account__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Parent Account").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create parent_account__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("parent_account__partyrelationshiptype", c)

	c = cache.Get("reports_to__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
		SetHasTable("No").
		SetPartyRelationshipName("Reports To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create reports_to__partyrelationshiptype: %v", err)
		return err
	}
	cache.Put("reports_to__partyrelationshiptype", c)

	c = cache.Get("lead_own_grp_member__partyrelationshiptype").(*ent.PartyRelationshipType)
	c, err = c.Update().
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
