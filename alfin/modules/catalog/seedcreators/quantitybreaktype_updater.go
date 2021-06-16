package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateQuantityBreakType(ctx context.Context) error {
	log.Println("QuantityBreakType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.QuantityBreakType
	failures := 0

	c = cache.Get("ship_weight__quantitybreaktype").(*ent.QuantityBreakType)
	c, err = c.Update().
		SetDescription("Shipping Weight Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_weight__quantitybreaktype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_weight__quantitybreaktype", c)
	}

	c = cache.Get("ship_quantity__quantitybreaktype").(*ent.QuantityBreakType)
	c, err = c.Update().
		SetDescription("Shipping Quantity Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_quantity__quantitybreaktype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_quantity__quantitybreaktype", c)
	}

	c = cache.Get("ship_price__quantitybreaktype").(*ent.QuantityBreakType)
	c, err = c.Update().
		SetDescription("Shipping Price Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_price__quantitybreaktype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_price__quantitybreaktype", c)
	}

	c = cache.Get("quantity__quantitybreaktype").(*ent.QuantityBreakType)
	c, err = c.Update().
		SetDescription("Price Component Quantity Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update quantity__quantitybreaktype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("quantity__quantitybreaktype", c)
	}

	c = cache.Get("order_value__quantitybreaktype").(*ent.QuantityBreakType)
	c, err = c.Update().
		SetDescription("Price Component Order Value Break").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update order_value__quantitybreaktype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("order_value__quantitybreaktype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
