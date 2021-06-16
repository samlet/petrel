package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductFeatureIactnType(ctx context.Context) error {
	log.Println("ProductFeatureIactnType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureIactnType
	failures := 0

	c = cache.Get("feature_iactn_depend__productfeatureiactntype").(*ent.ProductFeatureIactnType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Dependencies").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update feature_iactn_depend__productfeatureiactntype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("feature_iactn_depend__productfeatureiactntype", c)
	}

	c = cache.Get("feature_iactn_incomp__productfeatureiactntype").(*ent.ProductFeatureIactnType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Incompatibility").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update feature_iactn_incomp__productfeatureiactntype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("feature_iactn_incomp__productfeatureiactntype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
