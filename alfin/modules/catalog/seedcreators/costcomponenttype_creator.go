package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateCostComponentType(ctx context.Context) error {
	log.Println("CostComponentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CostComponentType

	c, err = client.CostComponentType.Create().SetStringRef("mat_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Materials Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create mat_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("mat_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("route_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Route (fixed asset usage) Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create route_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("route_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("labor_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Labor Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create labor_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("labor_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("gen_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("General Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create gen_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("gen_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("ind_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Indirect Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ind_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("ind_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("other_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Other Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("other_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_mat_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard Materials Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_mat_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_mat_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_route_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard Route (fixed asset usage) Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_route_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_route_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_labor_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard Labor Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_labor_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_labor_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_gen_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard General Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_gen_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_gen_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_ind_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard Indirect Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_ind_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_ind_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("est_std_other_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Estimated Standard Other Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create est_std_other_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("est_std_other_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_mat_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Materials Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_mat_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_mat_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_route_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Route (fixed asset usage) Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_route_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_route_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_labor_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Labor Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_labor_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_labor_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_other_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Other Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_other_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_other_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_gen_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual General Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_gen_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_gen_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_ind_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Indirect Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_ind_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_ind_cost__costcomponenttype", c)

	c, err = client.CostComponentType.Create().SetStringRef("actual_other_cost__costcomponenttype").
		SetHasTable("No").
		SetDescription("Actual Other Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create actual_other_cost__costcomponenttype: %v", err)
		return err
	}
	cache.Put("actual_other_cost__costcomponenttype", c)

	return nil
}
