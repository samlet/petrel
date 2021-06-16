package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateProductAssocType(ctx context.Context) error {
	log.Println("ProductAssocType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductAssocType
	failures := 0

	c = cache.Get("also_bought__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Also Bought").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update also_bought__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("also_bought__productassoctype", c)
	}

	c = cache.Get("product_upgrade__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Upgrade or Up-Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_upgrade__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_upgrade__productassoctype", c)
	}

	c = cache.Get("product_complement__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Complementary or Cross-Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_complement__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_complement__productassoctype", c)
	}

	c = cache.Get("product_incompatable__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Incompatible").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_incompatable__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_incompatable__productassoctype", c)
	}

	c = cache.Get("product_obsolescence__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("New Version, Replacement").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_obsolescence__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_obsolescence__productassoctype", c)
	}

	c = cache.Get("product_component__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Actual Product Component").
		AddChildren(cache.Get("manuf_component__productassoctype").(*ent.ProductAssocType)).
		AddChildren(cache.Get("engineer_component__productassoctype").(*ent.ProductAssocType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_component__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_component__productassoctype", c)
	}

	c = cache.Get("product_substitute__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Equivalent or Substitute").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_substitute__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_substitute__productassoctype", c)
	}

	c = cache.Get("product_variant__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Variant").
		AddProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddProductAssocs(cache.Get("demoproduct__demoproduct-2__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		AddProductAssocs(cache.Get("demoproduct__demoproduct-3__product_variant__1147521600__productassoc").(*ent.ProductAssoc)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_variant__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_variant__productassoctype", c)
	}

	c = cache.Get("unique_item__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Unique Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update unique_item__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("unique_item__productassoctype", c)
	}

	c = cache.Get("product_accessory__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Accessory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_accessory__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_accessory__productassoctype", c)
	}

	c = cache.Get("product_refurb__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Refurbished Equivalent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_refurb__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_refurb__productassoctype", c)
	}

	c = cache.Get("product_repair_srv__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Repair Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_repair_srv__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_repair_srv__productassoctype", c)
	}

	c = cache.Get("product_autoro__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Auto Reorder (needs recurrenceInfoId)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_autoro__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_autoro__productassoctype", c)
	}

	c = cache.Get("product_revision__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Revision").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_revision__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_revision__productassoctype", c)
	}

	c = cache.Get("manuf_component__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Manufacturing Bill of Materials").
		SetParent(cache.Get("product_component__productassoctype").(*ent.ProductAssocType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update manuf_component__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("manuf_component__productassoctype", c)
	}

	c = cache.Get("engineer_component__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Engineering Bill of Materials").
		SetParent(cache.Get("product_component__productassoctype").(*ent.ProductAssocType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update engineer_component__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("engineer_component__productassoctype", c)
	}

	c = cache.Get("product_manufactured__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Manufactured As").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_manufactured__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_manufactured__productassoctype", c)
	}

	c = cache.Get("product_conf__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Configurable product instance").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_conf__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_conf__productassoctype", c)
	}

	c = cache.Get("alternative_package__productassoctype").(*ent.ProductAssocType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Alternative Packaging").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update alternative_package__productassoctype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("alternative_package__productassoctype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
