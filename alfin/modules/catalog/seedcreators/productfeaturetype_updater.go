package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductFeatureType(ctx context.Context) error {
	log.Println("ProductFeatureType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductFeatureType
	failures := 0

	c = cache.Get("accessory__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Accessory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update accessory__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("accessory__productfeaturetype", c)
	}

	c = cache.Get("amount__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Amount").
		AddChildren(cache.Get("net_weight__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update amount__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("amount__productfeaturetype", c)
	}

	c = cache.Get("net_weight__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Net Weight").
		SetParent(cache.Get("amount__productfeaturetype").(*ent.ProductFeatureType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update net_weight__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("net_weight__productfeaturetype", c)
	}

	c = cache.Get("artist__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Artist").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update artist__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("artist__productfeaturetype", c)
	}

	c = cache.Get("billing_feature__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Billing Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update billing_feature__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("billing_feature__productfeaturetype", c)
	}

	c = cache.Get("brand__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Brand").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update brand__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("brand__productfeaturetype", c)
	}

	c = cache.Get("care__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Care").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update care__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("care__productfeaturetype", c)
	}

	c = cache.Get("color__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Color").
		AddProductFeatures(cache.Get("text_white__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_gray__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_black__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_red__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_green__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_blue__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_yellow__productfeature").(*ent.ProductFeature)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update color__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("color__productfeaturetype", c)
	}

	c = cache.Get("dimension__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Dimension").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update dimension__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("dimension__productfeaturetype", c)
	}

	c = cache.Get("equip_class__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Equipment Class").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update equip_class__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("equip_class__productfeaturetype", c)
	}

	c = cache.Get("fabric__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Fabric").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fabric__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fabric__productfeaturetype", c)
	}

	c = cache.Get("genre__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Genre").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update genre__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("genre__productfeaturetype", c)
	}

	c = cache.Get("gift_wrap__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Gift Wrap").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update gift_wrap__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("gift_wrap__productfeaturetype", c)
	}

	c = cache.Get("hardware_feature__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Hardware Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hardware_feature__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hardware_feature__productfeaturetype", c)
	}

	c = cache.Get("hazmat__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Hazmat").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update hazmat__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("hazmat__productfeaturetype", c)
	}

	c = cache.Get("license__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("License").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update license__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("license__productfeaturetype", c)
	}

	c = cache.Get("origin__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Origin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update origin__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("origin__productfeaturetype", c)
	}

	c = cache.Get("other_feature__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other_feature__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other_feature__productfeaturetype", c)
	}

	c = cache.Get("product_quality__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Quality").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_quality__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_quality__productfeaturetype", c)
	}

	c = cache.Get("size__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Size").
		AddProductFeatures(cache.Get("image_avatar__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_thumbnail__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_website__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_board__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor15__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor17__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor19__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("image_monitor21__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_small__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_middle__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_large__productfeature").(*ent.ProductFeature)).
		AddProductFeatures(cache.Get("text_verylarge__productfeature").(*ent.ProductFeature)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update size__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("size__productfeaturetype", c)
	}

	c = cache.Get("software_feature__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Software Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update software_feature__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("software_feature__productfeaturetype", c)
	}

	c = cache.Get("style__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Style").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update style__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("style__productfeaturetype", c)
	}

	c = cache.Get("symptom__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Symptom").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update symptom__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("symptom__productfeaturetype", c)
	}

	c = cache.Get("topic__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Topic").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update topic__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("topic__productfeaturetype", c)
	}

	c = cache.Get("type__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update type__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("type__productfeaturetype", c)
	}

	c = cache.Get("warranty__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Warranty").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update warranty__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("warranty__productfeaturetype", c)
	}

	c = cache.Get("model_year__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Model Year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update model_year__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("model_year__productfeaturetype", c)
	}

	c = cache.Get("year_made__productfeaturetype").(*ent.ProductFeatureType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Year Made").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update year_made__productfeaturetype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("year_made__productfeaturetype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
