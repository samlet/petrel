package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateCustomMethodType(ctx context.Context) error {
	log.Println("CustomMethodType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CustomMethodType

	c, err = client.CustomMethodType.Create().SetStringRef("cost_formula__custommethodtype").
		SetDescription("Formula for calculating costs for tasks and products").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cost_formula__custommethodtype: %v", err)
		return err
	}
	cache.Put("cost_formula__custommethodtype", c)

	c, err = client.CustomMethodType.Create().SetStringRef("price_formula__custommethodtype").
		SetDescription("Service with formula for calculating the unit price of a product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create price_formula__custommethodtype: %v", err)
		return err
	}
	cache.Put("price_formula__custommethodtype", c)

	c, err = client.CustomMethodType.Create().SetStringRef("ship_est__custommethodtype").
		SetHasTable("No").
		SetDescription("Shipment Gateway rate estimate methods").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est__custommethodtype: %v", err)
		return err
	}
	cache.Put("ship_est__custommethodtype", c)

	c, err = client.CustomMethodType.Create().SetStringRef("product_promo__custommethodtype").
		SetDescription("").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_promo__custommethodtype: %v", err)
		return err
	}
	cache.Put("product_promo__custommethodtype", c)

	c, err = client.CustomMethodType.Create().SetStringRef("product_promo_cond__custommethodtype").
		SetDescription("").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_promo_cond__custommethodtype: %v", err)
		return err
	}
	cache.Put("product_promo_cond__custommethodtype", c)

	c, err = client.CustomMethodType.Create().SetStringRef("product_promo_action__custommethodtype").
		SetDescription("").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_promo_action__custommethodtype: %v", err)
		return err
	}
	cache.Put("product_promo_action__custommethodtype", c)

	return nil
}
