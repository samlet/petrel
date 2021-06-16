package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateVarianceReason(ctx context.Context) error {
	log.Println("VarianceReason creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.VarianceReason

	c, err = client.VarianceReason.Create().SetStringRef("var_lost__variancereason").
		SetDescription("Lost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_lost__variancereason: %v", err)
		return err
	}
	cache.Put("var_lost__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_stolen__variancereason").
		SetDescription("Stolen").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_stolen__variancereason: %v", err)
		return err
	}
	cache.Put("var_stolen__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_found__variancereason").
		SetDescription("Found").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_found__variancereason: %v", err)
		return err
	}
	cache.Put("var_found__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_damaged__variancereason").
		SetDescription("Damaged").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_damaged__variancereason: %v", err)
		return err
	}
	cache.Put("var_damaged__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_integr__variancereason").
		SetDescription("Integration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_integr__variancereason: %v", err)
		return err
	}
	cache.Put("var_integr__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_sample__variancereason").
		SetDescription("Sample (Giveaway)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_sample__variancereason: %v", err)
		return err
	}
	cache.Put("var_sample__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_misship_ordered__variancereason").
		SetDescription("Mis-shipped Item Ordered (+)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_misship_ordered__variancereason: %v", err)
		return err
	}
	cache.Put("var_misship_ordered__variancereason", c)

	c, err = client.VarianceReason.Create().SetStringRef("var_misship_shipped__variancereason").
		SetDescription("Mis-shipped Item Shipped (-)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_misship_shipped__variancereason: %v", err)
		return err
	}
	cache.Put("var_misship_shipped__variancereason", c)

	return nil
}
