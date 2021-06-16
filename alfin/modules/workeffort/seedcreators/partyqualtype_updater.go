package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdatePartyQualType(ctx context.Context) error {
	log.Println("PartyQualType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyQualType
	failures := 0

	c = cache.Get("experience__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Work experience").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update experience__partyqualtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("experience__partyqualtype", c)
	}

	c = cache.Get("certification__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Certification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update certification__partyqualtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("certification__partyqualtype", c)
	}

	c = cache.Get("degree__partyqualtype").(*ent.PartyQualType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Degree").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update degree__partyqualtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("degree__partyqualtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
