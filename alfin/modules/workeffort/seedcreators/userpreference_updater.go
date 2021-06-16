package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateUserPreference(ctx context.Context) error {
	log.Println("UserPreference updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserPreference
	failures := 0

	c = cache.Get("_na___organization_party__userpreference").(*ent.UserPreference)
	c, err = c.Update().
		SetUserPrefTypeID("ORGANIZATION_PARTY").
		SetUserPrefGroupTypeID("GLOBAL_PREFERENCES").
		SetUserPrefValue("DEFAULT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update _na___organization_party__userpreference: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("_na___organization_party__userpreference", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
