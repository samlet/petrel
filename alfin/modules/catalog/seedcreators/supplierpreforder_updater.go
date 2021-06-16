package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateSupplierPrefOrder(ctx context.Context) error {
	log.Println("SupplierPrefOrder updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SupplierPrefOrder
	failures := 0

	c = cache.Get("10_main_suppl__supplierpreforder").(*ent.SupplierPrefOrder)
	c, err = c.Update().
		SetDescription("Main Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 10_main_suppl__supplierpreforder: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("10_main_suppl__supplierpreforder", c)
	}

	c = cache.Get("90_alt_suppl__supplierpreforder").(*ent.SupplierPrefOrder)
	c, err = c.Update().
		SetDescription("Alternative Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 90_alt_suppl__supplierpreforder: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("90_alt_suppl__supplierpreforder", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
