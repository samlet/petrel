package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateGoodIdentificationType(ctx context.Context) error {
	log.Println("GoodIdentificationType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.GoodIdentificationType

	c, err = client.GoodIdentificationType.Create().SetStringRef("isbn__goodidentificationtype").
		SetHasTable("No").
		SetDescription("ISBN").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create isbn__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("isbn__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("manufacturer_id_no__goodidentificationtype").
		SetHasTable("No").
		SetDescription("Manufacturer (Model) Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manufacturer_id_no__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("manufacturer_id_no__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("model_year__goodidentificationtype").
		SetHasTable("No").
		SetDescription("Model Year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create model_year__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("model_year__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("other_id__goodidentificationtype").
		SetHasTable("No").
		SetDescription("Other").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_id__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("other_id__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("sku__goodidentificationtype").
		SetHasTable("No").
		SetDescription("SKU").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sku__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("sku__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("upca__goodidentificationtype").
		SetHasTable("No").
		SetDescription("UPCA").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create upca__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("upca__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("upce__goodidentificationtype").
		SetHasTable("No").
		SetDescription("UPCE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create upce__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("upce__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("ean__goodidentificationtype").
		SetHasTable("No").
		SetDescription("EAN").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ean__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("ean__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("loc__goodidentificationtype").
		SetHasTable("No").
		SetDescription("Library of Congress").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create loc__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("loc__goodidentificationtype", c)

	c, err = client.GoodIdentificationType.Create().SetStringRef("hs_code__goodidentificationtype").
		SetHasTable("No").
		SetDescription("Harmonized System Codes (HS Code)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hs_code__goodidentificationtype: %v", err)
		return err
	}
	cache.Put("hs_code__goodidentificationtype", c)

	return nil
}
