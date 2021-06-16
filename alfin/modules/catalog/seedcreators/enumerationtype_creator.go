package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateEnumerationType(ctx context.Context) error {
	log.Println("EnumerationType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.EnumerationType

	c, err = client.EnumerationType.Create().SetStringRef("prod_price__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Price Parent Enum Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_price__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_price__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_price_in_param__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Price Input Parameter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_price_in_param__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_price_in_param__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_price_cond__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Price Condition").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_price_cond__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_price_cond__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("inv_res_order__enumerationtype").
		SetHasTable("No").
		SetDescription("Inventory Reservation Order").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inv_res_order__enumerationtype: %v", err)
		return err
	}
	cache.Put("inv_res_order__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("iid_reason__enumerationtype").
		SetHasTable("No").
		SetDescription("Inventory Item Detail Reason").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create iid_reason__enumerationtype: %v", err)
		return err
	}
	cache.Put("iid_reason__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prds_paysvc__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Store Payment Service Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_paysvc__enumerationtype: %v", err)
		return err
	}
	cache.Put("prds_paysvc__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prds_email__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Store Email Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_email__enumerationtype: %v", err)
		return err
	}
	cache.Put("prds_email__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("facloc_type__enumerationtype").
		SetHasTable("No").
		SetDescription("Facility Location Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create facloc_type__enumerationtype: %v", err)
		return err
	}
	cache.Put("facloc_type__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("kwovrd_trgt_type__enumerationtype").
		SetHasTable("No").
		SetDescription("Keyword Override Target Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create kwovrd_trgt_type__enumerationtype: %v", err)
		return err
	}
	cache.Put("kwovrd_trgt_type__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("pcat_link_type__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Category Link Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pcat_link_type__enumerationtype: %v", err)
		return err
	}
	cache.Put("pcat_link_type__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_geo__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Geo Data").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_geo__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_geo__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_promo__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Promotion Parent Enum Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_promo__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_promo__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_promo_in_param__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Promotion Input Parameter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_promo_in_param__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_promo_in_param__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_promo_cond__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Promotion Condition").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_promo_cond__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_promo_cond__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_promo_pcappl__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Promotion Product/Category Application Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_promo_pcappl__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_promo_pcappl__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_promo_action__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Promotion Action").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_promo_action__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_promo_action__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_rating_type__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Rating Field Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_rating_type__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_rating_type__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_vvmethod__enumerationtype").
		SetHasTable("No").
		SetDescription("Virtual Variant Method").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_vvmethod__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_vvmethod__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("prod_req_method__enumerationtype").
		SetHasTable("No").
		SetDescription("Product Requirement Method").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_req_method__enumerationtype: %v", err)
		return err
	}
	cache.Put("prod_req_method__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("image_reject_reason__enumerationtype").
		SetHasTable("No").
		SetDescription("Image Reject Reason").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create image_reject_reason__enumerationtype: %v", err)
		return err
	}
	cache.Put("image_reject_reason__enumerationtype", c)

	return nil
}
