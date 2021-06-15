package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyQualType(ctx context.Context) error {
	log.Println("PartyQualType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyQualType

	c, err = client.PartyQualType.Create().SetStringRef("experience__partyqualtype").
		SetHasTable("No").
		SetDescription("Work experience").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create experience__partyqualtype: %v", err)
		return err
	}
	cache.Put("experience__partyqualtype", c)

	c, err = client.PartyQualType.Create().SetStringRef("certification__partyqualtype").
		SetHasTable("No").
		SetDescription("Certification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create certification__partyqualtype: %v", err)
		return err
	}
	cache.Put("certification__partyqualtype", c)

	c, err = client.PartyQualType.Create().SetStringRef("degree__partyqualtype").
		SetHasTable("No").
		SetDescription("Degree").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create degree__partyqualtype: %v", err)
		return err
	}
	cache.Put("degree__partyqualtype", c)

	return nil
}
