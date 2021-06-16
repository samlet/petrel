package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderItem(ctx context.Context) error {
	log.Println("OrderItem creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderItem

	c, err = client.OrderItem.Create().SetStringRef("demo81015__00001__orderitem").
		SetOrderItemSeqID(common.ParseId("00001")).
		SetOrderItemTypeID(common.ParseId("PRODUCT_ORDER_ITEM")).
		SetProdCatalogID(common.ParseId("DemoCatalog")).
		SetIsPromo("No").
		SetQuantity(6.000000).
		SetSelectedAmount(0.000000).
		SetUnitPrice(38.400).
		SetUnitListPrice(48.000).
		SetIsModifiedPrice("No").
		SetItemDescription("Round Gizmo").
		SetChangeByUserLoginID("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demo81015__00001__orderitem: %v", err)
		return err
	}
	cache.Put("demo81015__00001__orderitem", c)

	return nil
}
