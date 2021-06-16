package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductFeatureApplType(ctx context.Context) error {
	log.Println("ProductFeatureApplType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureApplType
	failures := 0

	c = cache.Get("optional_feature__productfeatureappltype").(*ent.ProductFeatureApplType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Optional").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update optional_feature__productfeatureappltype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("optional_feature__productfeatureappltype", c)
	}

	c = cache.Get("required_feature__productfeatureappltype").(*ent.ProductFeatureApplType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Required").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update required_feature__productfeatureappltype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("required_feature__productfeatureappltype", c)
	}

	c = cache.Get("selectable_feature__productfeatureappltype").(*ent.ProductFeatureApplType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Selectable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update selectable_feature__productfeatureappltype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("selectable_feature__productfeatureappltype", c)
	}

	c = cache.Get("standard_feature__productfeatureappltype").(*ent.ProductFeatureApplType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Standard").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update standard_feature__productfeatureappltype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("standard_feature__productfeatureappltype", c)
	}

	c = cache.Get("distinguishing_feat__productfeatureappltype").(*ent.ProductFeatureApplType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Distinguishing").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update distinguishing_feat__productfeatureappltype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("distinguishing_feat__productfeatureappltype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
