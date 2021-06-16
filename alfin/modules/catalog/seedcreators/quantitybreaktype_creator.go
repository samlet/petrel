package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateQuantityBreakType(ctx context.Context) error {
	log.Println("QuantityBreakType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.QuantityBreakType

	c, err = client.QuantityBreakType.Create().SetStringRef("ship_weight__quantitybreaktype").
		SetDescription("Shipping Weight Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_weight__quantitybreaktype: %v", err)
		return err
	}
	cache.Put("ship_weight__quantitybreaktype", c)

	c, err = client.QuantityBreakType.Create().SetStringRef("ship_quantity__quantitybreaktype").
		SetDescription("Shipping Quantity Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_quantity__quantitybreaktype: %v", err)
		return err
	}
	cache.Put("ship_quantity__quantitybreaktype", c)

	c, err = client.QuantityBreakType.Create().SetStringRef("ship_price__quantitybreaktype").
		SetDescription("Shipping Price Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_price__quantitybreaktype: %v", err)
		return err
	}
	cache.Put("ship_price__quantitybreaktype", c)

	c, err = client.QuantityBreakType.Create().SetStringRef("quantity__quantitybreaktype").
		SetDescription("Price Component Quantity Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create quantity__quantitybreaktype: %v", err)
		return err
	}
	cache.Put("quantity__quantitybreaktype", c)

	c, err = client.QuantityBreakType.Create().SetStringRef("order_value__quantitybreaktype").
		SetDescription("Price Component Order Value Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create order_value__quantitybreaktype: %v", err)
		return err
	}
	cache.Put("order_value__quantitybreaktype", c)

	return nil
}
