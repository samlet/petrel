package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductFeatureApplType(ctx context.Context) error {
	log.Println("ProductFeatureApplType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureApplType

	c, err = client.ProductFeatureApplType.Create().SetStringRef("optional_feature__productfeatureappltype").
		SetHasTable("No").
		SetDescription("Optional").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create optional_feature__productfeatureappltype: %v", err)
		return err
	}
	cache.Put("optional_feature__productfeatureappltype", c)

	c, err = client.ProductFeatureApplType.Create().SetStringRef("required_feature__productfeatureappltype").
		SetHasTable("No").
		SetDescription("Required").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create required_feature__productfeatureappltype: %v", err)
		return err
	}
	cache.Put("required_feature__productfeatureappltype", c)

	c, err = client.ProductFeatureApplType.Create().SetStringRef("selectable_feature__productfeatureappltype").
		SetHasTable("No").
		SetDescription("Selectable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create selectable_feature__productfeatureappltype: %v", err)
		return err
	}
	cache.Put("selectable_feature__productfeatureappltype", c)

	c, err = client.ProductFeatureApplType.Create().SetStringRef("standard_feature__productfeatureappltype").
		SetHasTable("No").
		SetDescription("Standard").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create standard_feature__productfeatureappltype: %v", err)
		return err
	}
	cache.Put("standard_feature__productfeatureappltype", c)

	c, err = client.ProductFeatureApplType.Create().SetStringRef("distinguishing_feat__productfeatureappltype").
		SetHasTable("No").
		SetDescription("Distinguishing").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create distinguishing_feat__productfeatureappltype: %v", err)
		return err
	}
	cache.Put("distinguishing_feat__productfeatureappltype", c)

	return nil
}
