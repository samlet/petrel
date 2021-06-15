package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyType(ctx context.Context) error {
	log.Println("PartyType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyType

	c, err = client.PartyType.Create().SetStringRef("automated_agent__partytype").
		SetHasTable("No").
		SetDescription("Automated Agent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create automated_agent__partytype: %v", err)
		return err
	}
	cache.Put("automated_agent__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("person__partytype").
		SetHasTable("Yes").
		SetDescription("Person").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create person__partytype: %v", err)
		return err
	}
	cache.Put("person__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("party_group__partytype").
		SetHasTable("Yes").
		SetDescription("Party Group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_group__partytype: %v", err)
		return err
	}
	cache.Put("party_group__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("informal_group__partytype").
		SetHasTable("No").
		SetDescription("Informal Group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create informal_group__partytype: %v", err)
		return err
	}
	cache.Put("informal_group__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("family__partytype").
		SetHasTable("No").
		SetDescription("Family").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create family__partytype: %v", err)
		return err
	}
	cache.Put("family__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("team__partytype").
		SetHasTable("No").
		SetDescription("Team").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create team__partytype: %v", err)
		return err
	}
	cache.Put("team__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("legal_organization__partytype").
		SetHasTable("No").
		SetDescription("Legal Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create legal_organization__partytype: %v", err)
		return err
	}
	cache.Put("legal_organization__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("corporation__partytype").
		SetHasTable("No").
		SetDescription("Corporation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create corporation__partytype: %v", err)
		return err
	}
	cache.Put("corporation__partytype", c)

	c, err = client.PartyType.Create().SetStringRef("government_agency__partytype").
		SetHasTable("No").
		SetDescription("Government Agency").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create government_agency__partytype: %v", err)
		return err
	}
	cache.Put("government_agency__partytype", c)

	return nil
}
