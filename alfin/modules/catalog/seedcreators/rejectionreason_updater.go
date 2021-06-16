package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateRejectionReason(ctx context.Context) error {
	log.Println("RejectionReason updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RejectionReason
	failures := 0

	c = cache.Get("srj_damaged__rejectionreason").(*ent.RejectionReason)
	c, err = c.Update().
		SetDescription("Damaged").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update srj_damaged__rejectionreason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("srj_damaged__rejectionreason", c)
	}

	c = cache.Get("srj_not_ordered__rejectionreason").(*ent.RejectionReason)
	c, err = c.Update().
		SetDescription("Not Ordered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update srj_not_ordered__rejectionreason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("srj_not_ordered__rejectionreason", c)
	}

	c = cache.Get("srj_over_shipped__rejectionreason").(*ent.RejectionReason)
	c, err = c.Update().
		SetDescription("Over Shipped").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update srj_over_shipped__rejectionreason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("srj_over_shipped__rejectionreason", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
