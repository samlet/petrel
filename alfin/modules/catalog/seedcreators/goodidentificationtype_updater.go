package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateGoodIdentificationType(ctx context.Context) error {
	log.Println("GoodIdentificationType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.GoodIdentificationType
	failures := 0

	c = cache.Get("isbn__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("ISBN").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update isbn__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("isbn__goodidentificationtype", c)
	}

	c = cache.Get("manufacturer_id_no__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manufacturer (Model) Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update manufacturer_id_no__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("manufacturer_id_no__goodidentificationtype", c)
	}

	c = cache.Get("model_year__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Model Year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update model_year__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("model_year__goodidentificationtype", c)
	}

	c = cache.Get("other_id__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_id__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_id__goodidentificationtype", c)
	}

	c = cache.Get("sku__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("SKU").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sku__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sku__goodidentificationtype", c)
	}

	c = cache.Get("upca__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("UPCA").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update upca__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("upca__goodidentificationtype", c)
	}

	c = cache.Get("upce__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("UPCE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update upce__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("upce__goodidentificationtype", c)
	}

	c = cache.Get("ean__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("EAN").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ean__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ean__goodidentificationtype", c)
	}

	c = cache.Get("loc__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Library of Congress").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update loc__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("loc__goodidentificationtype", c)
	}

	c = cache.Get("hs_code__goodidentificationtype").(*ent.GoodIdentificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Harmonized System Codes (HS Code)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hs_code__goodidentificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hs_code__goodidentificationtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
