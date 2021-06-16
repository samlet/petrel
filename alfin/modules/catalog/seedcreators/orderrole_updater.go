package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateOrderRole(ctx context.Context) error {
	log.Println("OrderRole updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderRole
	failures := 0

	c = cache.Get("demo81015__company__bill_from_vendor__orderrole").(*ent.OrderRole)
	c, err = c.Update().
		SetPartyID(common.ParseId("Company")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__company__bill_from_vendor__orderrole: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__company__bill_from_vendor__orderrole", c)
	}

	c = cache.Get("demo81015__democustomer__bill_to_customer__orderrole").(*ent.OrderRole)
	c, err = c.Update().
		SetPartyID(common.ParseId("DemoCustomer")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__democustomer__bill_to_customer__orderrole: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__democustomer__bill_to_customer__orderrole", c)
	}

	c = cache.Get("demo81015__democustomer__end_user_customer__orderrole").(*ent.OrderRole)
	c, err = c.Update().
		SetPartyID(common.ParseId("DemoCustomer")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__democustomer__end_user_customer__orderrole: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__democustomer__end_user_customer__orderrole", c)
	}

	c = cache.Get("demo81015__democustomer__placing_customer__orderrole").(*ent.OrderRole)
	c, err = c.Update().
		SetPartyID(common.ParseId("DemoCustomer")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__democustomer__placing_customer__orderrole: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__democustomer__placing_customer__orderrole", c)
	}

	c = cache.Get("demo81015__democustomer__ship_to_customer__orderrole").(*ent.OrderRole)
	c, err = c.Update().
		SetPartyID(common.ParseId("DemoCustomer")).
		SetOrderHeader(cache.Get("demo81015__orderheader").(*ent.OrderHeader)).
		AddOrderItems(cache.Get("demo81015__00001__orderitem").(*ent.OrderItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demo81015__democustomer__ship_to_customer__orderrole: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demo81015__democustomer__ship_to_customer__orderrole", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
