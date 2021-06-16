package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderRole(ctx context.Context) error {
	log.Println("OrderRole creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderRole

	c, err = client.OrderRole.Create().SetStringRef("demo81015__company__bill_from_vendor__orderrole").
		SetPartyID(common.ParseId("Company")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__company__bill_from_vendor__orderrole: %v", err)
		return err
	}
	cache.Put("demo81015__company__bill_from_vendor__orderrole", c)

	c, err = client.OrderRole.Create().SetStringRef("demo81015__democustomer__bill_to_customer__orderrole").
		SetPartyID(common.ParseId("DemoCustomer")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__democustomer__bill_to_customer__orderrole: %v", err)
		return err
	}
	cache.Put("demo81015__democustomer__bill_to_customer__orderrole", c)

	c, err = client.OrderRole.Create().SetStringRef("demo81015__democustomer__end_user_customer__orderrole").
		SetPartyID(common.ParseId("DemoCustomer")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__democustomer__end_user_customer__orderrole: %v", err)
		return err
	}
	cache.Put("demo81015__democustomer__end_user_customer__orderrole", c)

	c, err = client.OrderRole.Create().SetStringRef("demo81015__democustomer__placing_customer__orderrole").
		SetPartyID(common.ParseId("DemoCustomer")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__democustomer__placing_customer__orderrole: %v", err)
		return err
	}
	cache.Put("demo81015__democustomer__placing_customer__orderrole", c)

	c, err = client.OrderRole.Create().SetStringRef("demo81015__democustomer__ship_to_customer__orderrole").
		SetPartyID(common.ParseId("DemoCustomer")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__democustomer__ship_to_customer__orderrole: %v", err)
		return err
	}
	cache.Put("demo81015__democustomer__ship_to_customer__orderrole", c)

	return nil
}
