package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderPaymentPreference(ctx context.Context) error {
	log.Println("OrderPaymentPreference creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderPaymentPreference

	c, err = client.OrderPaymentPreference.Create().SetStringRef("10020__orderpaymentpreference").
		SetPaymentMethodTypeID(common.ParseId("EXT_OFFLINE")).
		SetPresentFlag("No").
		SetSwipedFlag("No").
		SetOverflowFlag("No").
		SetMaxAmount(105.14).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.789")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 10020__orderpaymentpreference: %v", err)
		return err
	}
	cache.Put("10020__orderpaymentpreference", c)

	return nil
}
