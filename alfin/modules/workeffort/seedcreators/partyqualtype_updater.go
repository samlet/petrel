package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyQualType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyQualType

	c = cache.Get("experience__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Work experience").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create experience__partyqualtype: %v", err)
		return err
	}
	cache.Put("experience__partyqualtype", c)

	c = cache.Get("certification__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Certification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create certification__partyqualtype: %v", err)
		return err
	}
	cache.Put("certification__partyqualtype", c)

	c = cache.Get("degree__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
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
