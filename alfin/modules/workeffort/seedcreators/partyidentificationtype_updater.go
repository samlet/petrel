package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyIdentificationType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyIdentificationType

	c = cache.Get("party_import__partyidentificationtype").(*ent.PartyIdentificationType)
	c, err = c.Update().
		SetDescription("Original ID in the system where this record was imported from").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_import__partyidentificationtype: %v", err)
		return err
	}
	cache.Put("party_import__partyidentificationtype", c)

	return nil
}
