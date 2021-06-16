package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderContactMech(ctx context.Context) error {
	log.Println("OrderContactMech creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderContactMech

	c, err = client.OrderContactMech.Create().SetStringRef("demo81015__order_email__9021__ordercontactmech").
		SetContactMechPurposeTypeID(common.ParseId("ORDER_EMAIL")).
		SetContactMechID(common.ParseId("9021")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__order_email__9021__ordercontactmech: %v", err)
		return err
	}
	cache.Put("demo81015__order_email__9021__ordercontactmech", c)

	c, err = client.OrderContactMech.Create().SetStringRef("demo81015__shipping_location__9010__ordercontactmech").
		SetContactMechPurposeTypeID(common.ParseId("SHIPPING_LOCATION")).
		SetContactMechID(common.ParseId("9010")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__shipping_location__9010__ordercontactmech: %v", err)
		return err
	}
	cache.Put("demo81015__shipping_location__9010__ordercontactmech", c)

	return nil
}
