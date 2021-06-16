package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateProductAssocType(ctx context.Context) error {
	log.Println("ProductAssocType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ProductAssocType

	c, err = client.ProductAssocType.Create().SetStringRef("also_bought__productassoctype").
		SetHasTable("No").
		SetDescription("Also Bought").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create also_bought__productassoctype: %v", err)
		return err
	}
	cache.Put("also_bought__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_upgrade__productassoctype").
		SetHasTable("No").
		SetDescription("Upgrade or Up-Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_upgrade__productassoctype: %v", err)
		return err
	}
	cache.Put("product_upgrade__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_complement__productassoctype").
		SetHasTable("No").
		SetDescription("Complementary or Cross-Sell").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_complement__productassoctype: %v", err)
		return err
	}
	cache.Put("product_complement__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_incompatable__productassoctype").
		SetHasTable("No").
		SetDescription("Incompatible").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_incompatable__productassoctype: %v", err)
		return err
	}
	cache.Put("product_incompatable__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_obsolescence__productassoctype").
		SetHasTable("No").
		SetDescription("New Version, Replacement").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_obsolescence__productassoctype: %v", err)
		return err
	}
	cache.Put("product_obsolescence__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_component__productassoctype").
		SetHasTable("No").
		SetDescription("Actual Product Component").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_component__productassoctype: %v", err)
		return err
	}
	cache.Put("product_component__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_substitute__productassoctype").
		SetHasTable("No").
		SetDescription("Equivalent or Substitute").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_substitute__productassoctype: %v", err)
		return err
	}
	cache.Put("product_substitute__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_variant__productassoctype").
		SetHasTable("No").
		SetDescription("Product Variant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_variant__productassoctype: %v", err)
		return err
	}
	cache.Put("product_variant__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("unique_item__productassoctype").
		SetHasTable("No").
		SetDescription("Unique Item").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create unique_item__productassoctype: %v", err)
		return err
	}
	cache.Put("unique_item__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_accessory__productassoctype").
		SetHasTable("No").
		SetDescription("Accessory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_accessory__productassoctype: %v", err)
		return err
	}
	cache.Put("product_accessory__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_refurb__productassoctype").
		SetHasTable("No").
		SetDescription("Refurbished Equivalent").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_refurb__productassoctype: %v", err)
		return err
	}
	cache.Put("product_refurb__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_repair_srv__productassoctype").
		SetHasTable("No").
		SetDescription("Repair Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_repair_srv__productassoctype: %v", err)
		return err
	}
	cache.Put("product_repair_srv__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_autoro__productassoctype").
		SetHasTable("No").
		SetDescription("Auto Reorder (needs recurrenceInfoId)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_autoro__productassoctype: %v", err)
		return err
	}
	cache.Put("product_autoro__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_revision__productassoctype").
		SetHasTable("No").
		SetDescription("Revision").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_revision__productassoctype: %v", err)
		return err
	}
	cache.Put("product_revision__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("manuf_component__productassoctype").
		SetHasTable("No").
		SetDescription("Manufacturing Bill of Materials").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create manuf_component__productassoctype: %v", err)
		return err
	}
	cache.Put("manuf_component__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("engineer_component__productassoctype").
		SetHasTable("No").
		SetDescription("Engineering Bill of Materials").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create engineer_component__productassoctype: %v", err)
		return err
	}
	cache.Put("engineer_component__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_manufactured__productassoctype").
		SetHasTable("No").
		SetDescription("Product Manufactured As").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_manufactured__productassoctype: %v", err)
		return err
	}
	cache.Put("product_manufactured__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("product_conf__productassoctype").
		SetHasTable("No").
		SetDescription("Configurable product instance").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create product_conf__productassoctype: %v", err)
		return err
	}
	cache.Put("product_conf__productassoctype", c)

	c, err = client.ProductAssocType.Create().SetStringRef("alternative_package__productassoctype").
		SetHasTable("No").
		SetDescription("Alternative Packaging").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create alternative_package__productassoctype: %v", err)
		return err
	}
	cache.Put("alternative_package__productassoctype", c)

	return nil
}
