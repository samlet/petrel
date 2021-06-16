package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateSupplierPrefOrder(ctx context.Context) error {
	log.Println("SupplierPrefOrder creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SupplierPrefOrder

	c, err = client.SupplierPrefOrder.Create().SetStringRef("10_main_suppl__supplierpreforder").
		SetDescription("Main Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 10_main_suppl__supplierpreforder: %v", err)
		return err
	}
	cache.Put("10_main_suppl__supplierpreforder", c)

	c, err = client.SupplierPrefOrder.Create().SetStringRef("90_alt_suppl__supplierpreforder").
		SetDescription("Alternative Supplier").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 90_alt_suppl__supplierpreforder: %v", err)
		return err
	}
	cache.Put("90_alt_suppl__supplierpreforder", c)

	return nil
}
