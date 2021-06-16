package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductFeatureType(ctx context.Context) error {
	log.Println("ProductFeatureType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureType

	c, err = client.ProductFeatureType.Create().SetStringRef("accessory__productfeaturetype").
		SetHasTable("No").
		SetDescription("Accessory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create accessory__productfeaturetype: %v", err)
		return err
	}
	cache.Put("accessory__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("amount__productfeaturetype").
		SetHasTable("No").
		SetDescription("Amount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create amount__productfeaturetype: %v", err)
		return err
	}
	cache.Put("amount__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("net_weight__productfeaturetype").
		SetHasTable("No").
		SetDescription("Net Weight").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create net_weight__productfeaturetype: %v", err)
		return err
	}
	cache.Put("net_weight__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("artist__productfeaturetype").
		SetHasTable("No").
		SetDescription("Artist").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create artist__productfeaturetype: %v", err)
		return err
	}
	cache.Put("artist__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("billing_feature__productfeaturetype").
		SetHasTable("No").
		SetDescription("Billing Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create billing_feature__productfeaturetype: %v", err)
		return err
	}
	cache.Put("billing_feature__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("brand__productfeaturetype").
		SetHasTable("No").
		SetDescription("Brand").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create brand__productfeaturetype: %v", err)
		return err
	}
	cache.Put("brand__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("care__productfeaturetype").
		SetHasTable("No").
		SetDescription("Care").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create care__productfeaturetype: %v", err)
		return err
	}
	cache.Put("care__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("color__productfeaturetype").
		SetHasTable("No").
		SetDescription("Color").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create color__productfeaturetype: %v", err)
		return err
	}
	cache.Put("color__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("dimension__productfeaturetype").
		SetHasTable("No").
		SetDescription("Dimension").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create dimension__productfeaturetype: %v", err)
		return err
	}
	cache.Put("dimension__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("equip_class__productfeaturetype").
		SetHasTable("No").
		SetDescription("Equipment Class").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create equip_class__productfeaturetype: %v", err)
		return err
	}
	cache.Put("equip_class__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("fabric__productfeaturetype").
		SetHasTable("No").
		SetDescription("Fabric").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fabric__productfeaturetype: %v", err)
		return err
	}
	cache.Put("fabric__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("genre__productfeaturetype").
		SetHasTable("No").
		SetDescription("Genre").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create genre__productfeaturetype: %v", err)
		return err
	}
	cache.Put("genre__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("gift_wrap__productfeaturetype").
		SetHasTable("No").
		SetDescription("Gift Wrap").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create gift_wrap__productfeaturetype: %v", err)
		return err
	}
	cache.Put("gift_wrap__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("hardware_feature__productfeaturetype").
		SetHasTable("No").
		SetDescription("Hardware Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hardware_feature__productfeaturetype: %v", err)
		return err
	}
	cache.Put("hardware_feature__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("hazmat__productfeaturetype").
		SetHasTable("No").
		SetDescription("Hazmat").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create hazmat__productfeaturetype: %v", err)
		return err
	}
	cache.Put("hazmat__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("license__productfeaturetype").
		SetHasTable("No").
		SetDescription("License").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create license__productfeaturetype: %v", err)
		return err
	}
	cache.Put("license__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("origin__productfeaturetype").
		SetHasTable("No").
		SetDescription("Origin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create origin__productfeaturetype: %v", err)
		return err
	}
	cache.Put("origin__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("other_feature__productfeaturetype").
		SetHasTable("No").
		SetDescription("Other Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_feature__productfeaturetype: %v", err)
		return err
	}
	cache.Put("other_feature__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("product_quality__productfeaturetype").
		SetHasTable("No").
		SetDescription("Product Quality").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_quality__productfeaturetype: %v", err)
		return err
	}
	cache.Put("product_quality__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("size__productfeaturetype").
		SetHasTable("No").
		SetDescription("Size").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create size__productfeaturetype: %v", err)
		return err
	}
	cache.Put("size__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("software_feature__productfeaturetype").
		SetHasTable("No").
		SetDescription("Software Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create software_feature__productfeaturetype: %v", err)
		return err
	}
	cache.Put("software_feature__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("style__productfeaturetype").
		SetHasTable("No").
		SetDescription("Style").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create style__productfeaturetype: %v", err)
		return err
	}
	cache.Put("style__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("symptom__productfeaturetype").
		SetHasTable("No").
		SetDescription("Symptom").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create symptom__productfeaturetype: %v", err)
		return err
	}
	cache.Put("symptom__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("topic__productfeaturetype").
		SetHasTable("No").
		SetDescription("Topic").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create topic__productfeaturetype: %v", err)
		return err
	}
	cache.Put("topic__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("type__productfeaturetype").
		SetHasTable("No").
		SetDescription("Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create type__productfeaturetype: %v", err)
		return err
	}
	cache.Put("type__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("warranty__productfeaturetype").
		SetHasTable("No").
		SetDescription("Warranty").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create warranty__productfeaturetype: %v", err)
		return err
	}
	cache.Put("warranty__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("model_year__productfeaturetype").
		SetHasTable("No").
		SetDescription("Model Year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create model_year__productfeaturetype: %v", err)
		return err
	}
	cache.Put("model_year__productfeaturetype", c)

	c, err = client.ProductFeatureType.Create().SetStringRef("year_made__productfeaturetype").
		SetHasTable("No").
		SetDescription("Year Made").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create year_made__productfeaturetype: %v", err)
		return err
	}
	cache.Put("year_made__productfeaturetype", c)

	return nil
}
