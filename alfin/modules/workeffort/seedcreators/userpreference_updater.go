package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateUserPreference(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserPreference

	c = cache.Get("_na___organization_party__userpreference").(*ent.UserPreference)
	c, err = c.Update().
		SetUserPrefTypeID("ORGANIZATION_PARTY").
		SetUserPrefGroupTypeID("GLOBAL_PREFERENCES").
		SetUserPrefValue("DEFAULT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___organization_party__userpreference: %v", err)
		return err
	}
	cache.Put("_na___organization_party__userpreference", c)

	return nil
}
