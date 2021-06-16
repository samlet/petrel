package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateVarianceReason(ctx context.Context) error {
	log.Println("VarianceReason updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.VarianceReason
	failures := 0

	c = cache.Get("var_lost__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Lost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_lost__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_lost__variancereason", c)
	}

	c = cache.Get("var_stolen__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Stolen").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_stolen__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_stolen__variancereason", c)
	}

	c = cache.Get("var_found__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Found").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_found__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_found__variancereason", c)
	}

	c = cache.Get("var_damaged__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Damaged").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_damaged__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_damaged__variancereason", c)
	}

	c = cache.Get("var_integr__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Integration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_integr__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_integr__variancereason", c)
	}

	c = cache.Get("var_sample__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Sample (Giveaway)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_sample__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_sample__variancereason", c)
	}

	c = cache.Get("var_misship_ordered__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Mis-shipped Item Ordered (+)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_misship_ordered__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_misship_ordered__variancereason", c)
	}

	c = cache.Get("var_misship_shipped__variancereason").(*ent.VarianceReason)
	c, err = c.Update().
		SetDescription("Mis-shipped Item Shipped (-)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_misship_shipped__variancereason: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_misship_shipped__variancereason", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
