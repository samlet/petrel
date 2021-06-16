package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderHeader(ctx context.Context) error {
	log.Println("OrderHeader creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderHeader

	c, err = client.OrderHeader.Create().SetStringRef("demo81015__orderheader").
		SetOrderTypeID(common.ParseId("SALES_ORDER")).
		SetPriority("Unknown").
		SetEntryDate(common.ParseDateTime("2011-08-12 23:17:11.507")).
		SetVisitID(common.ParseId("10043")).
		SetCreatedBy("admin").
		SetCurrencyUom(common.ParseId("USD")).
		SetRemainingSubTotal(79.56).
		SetGrandTotal(105.14).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__orderheader: %v", err)
		return err
	}
	cache.Put("demo81015__orderheader", c)

	return nil
}
