package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateRoleType(ctx context.Context) error {
	log.Println("RoleType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RoleType
	failures := 0

	c = cache.Get("imageapprover__roletype").(*ent.RoleType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Approver").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update imageapprover__roletype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("imageapprover__roletype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
