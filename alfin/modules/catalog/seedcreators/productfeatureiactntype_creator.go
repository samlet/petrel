package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductFeatureIactnType(ctx context.Context) error {
	log.Println("ProductFeatureIactnType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureIactnType

	c, err = client.ProductFeatureIactnType.Create().SetStringRef("feature_iactn_depend__productfeatureiactntype").
		SetHasTable("No").
		SetDescription("Dependencies").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create feature_iactn_depend__productfeatureiactntype: %v", err)
		return err
	}
	cache.Put("feature_iactn_depend__productfeatureiactntype", c)

	c, err = client.ProductFeatureIactnType.Create().SetStringRef("feature_iactn_incomp__productfeatureiactntype").
		SetHasTable("No").
		SetDescription("Incompatibility").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create feature_iactn_incomp__productfeatureiactntype: %v", err)
		return err
	}
	cache.Put("feature_iactn_incomp__productfeatureiactntype", c)

	return nil
}
