package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateCostComponentType(ctx context.Context) error {
	log.Println("CostComponentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CostComponentType
	failures := 0

	c = cache.Get("mat_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Materials Cost").
		AddChildren(cache.Get("est_std_mat_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_mat_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update mat_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("mat_cost__costcomponenttype", c)
	}

	c = cache.Get("route_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Route (fixed asset usage) Cost").
		AddChildren(cache.Get("est_std_route_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_route_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update route_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("route_cost__costcomponenttype", c)
	}

	c = cache.Get("labor_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Labor Cost").
		AddChildren(cache.Get("est_std_labor_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_labor_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update labor_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("labor_cost__costcomponenttype", c)
	}

	c = cache.Get("gen_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("General Cost").
		AddChildren(cache.Get("est_std_gen_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_gen_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update gen_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("gen_cost__costcomponenttype", c)
	}

	c = cache.Get("ind_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Indirect Cost").
		AddChildren(cache.Get("est_std_ind_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_ind_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ind_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ind_cost__costcomponenttype", c)
	}

	c = cache.Get("other_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Cost").
		AddChildren(cache.Get("est_std_other_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_other_cost__costcomponenttype").(*ent.CostComponentType)).
		AddChildren(cache.Get("actual_other_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_mat_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard Materials Cost").
		SetParent(cache.Get("mat_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_mat_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_mat_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_route_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard Route (fixed asset usage) Cost").
		SetParent(cache.Get("route_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_route_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_route_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_labor_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard Labor Cost").
		SetParent(cache.Get("labor_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_labor_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_labor_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_gen_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard General Cost").
		SetParent(cache.Get("gen_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_gen_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_gen_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_ind_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard Indirect Cost").
		SetParent(cache.Get("ind_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_ind_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_ind_cost__costcomponenttype", c)
	}

	c = cache.Get("est_std_other_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Estimated Standard Other Cost").
		SetParent(cache.Get("other_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update est_std_other_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("est_std_other_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_mat_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Materials Cost").
		SetParent(cache.Get("mat_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_mat_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_mat_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_route_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Route (fixed asset usage) Cost").
		SetParent(cache.Get("route_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_route_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_route_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_labor_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Labor Cost").
		SetParent(cache.Get("labor_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_labor_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_labor_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_other_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Other Cost").
		SetParent(cache.Get("other_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_other_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_other_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_gen_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual General Cost").
		SetParent(cache.Get("gen_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_gen_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_gen_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_ind_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Indirect Cost").
		SetParent(cache.Get("ind_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_ind_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_ind_cost__costcomponenttype", c)
	}

	c = cache.Get("actual_other_cost__costcomponenttype").(*ent.CostComponentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Other Cost").
		SetParent(cache.Get("other_cost__costcomponenttype").(*ent.CostComponentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update actual_other_cost__costcomponenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("actual_other_cost__costcomponenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
