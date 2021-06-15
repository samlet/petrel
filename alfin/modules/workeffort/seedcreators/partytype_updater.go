package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyType(ctx context.Context) error {
	log.Println("PartyType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyType

	c = cache.Get("automated_agent__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Automated Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update automated_agent__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("automated_agent__partytype", c)
	}

	c = cache.Get("person__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Person").
		AddParties(cache.Get("demoemployee__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee1__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee2__party").(*ent.Party)).
		AddParties(cache.Get("demoemployee3__party").(*ent.Party)).
		AddParties(cache.Get("democustomer1__party").(*ent.Party)).
		AddParties(cache.Get("democustomer2__party").(*ent.Party)).
		AddParties(cache.Get("democustomer3__party").(*ent.Party)).
		AddParties(cache.Get("workeffortuser__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update person__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("person__partytype", c)
	}

	c = cache.Get("party_group__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Party Group").
		AddChildren(cache.Get("informal_group__partytype").(*ent.PartyType)).
		AddChildren(cache.Get("legal_organization__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_group__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_group__partytype", c)
	}

	c = cache.Get("informal_group__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Informal Group").
		SetParent(cache.Get("party_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("informal_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("legal_organization__partytype").(*ent.PartyType)).
		AddChildren(cache.Get("family__partytype").(*ent.PartyType)).
		AddChildren(cache.Get("team__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update informal_group__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("informal_group__partytype", c)
	}

	c = cache.Get("family__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Family").
		SetParent(cache.Get("informal_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("family__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("team__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update family__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("family__partytype", c)
	}

	c = cache.Get("team__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Team").
		SetParent(cache.Get("informal_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("family__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("team__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update team__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("team__partytype", c)
	}

	c = cache.Get("legal_organization__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Legal Organization").
		SetParent(cache.Get("party_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("informal_group__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("legal_organization__partytype").(*ent.PartyType)).
		AddChildren(cache.Get("corporation__partytype").(*ent.PartyType)).
		AddChildren(cache.Get("government_agency__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update legal_organization__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("legal_organization__partytype", c)
	}

	c = cache.Get("corporation__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Corporation").
		SetParent(cache.Get("legal_organization__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("corporation__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("government_agency__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update corporation__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("corporation__partytype", c)
	}

	c = cache.Get("government_agency__partytype").(*ent.PartyType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Government Agency").
		SetParent(cache.Get("legal_organization__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("corporation__partytype").(*ent.PartyType)).
		AddSiblingPartyTypes(cache.Get("government_agency__partytype").(*ent.PartyType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update government_agency__partytype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("government_agency__partytype", c)
	}

	return nil
}
