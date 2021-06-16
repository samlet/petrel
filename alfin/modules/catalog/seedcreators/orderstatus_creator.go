package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderStatus(ctx context.Context) error {
	log.Println("OrderStatus creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderStatus

	c, err = client.OrderStatus.Create().SetStringRef("81015__orderstatus").
		SetStatusDatetime(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetStatusUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81015__orderstatus: %v", err)
		return err
	}
	cache.Put("81015__orderstatus", c)

	c, err = client.OrderStatus.Create().SetStringRef("81016__orderstatus").
		SetOrderItemSeqID(common.ParseId("00001")).
		SetStatusDatetime(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetStatusUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81016__orderstatus: %v", err)
		return err
	}
	cache.Put("81016__orderstatus", c)

	return nil
}
