package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateEnumerationType(ctx context.Context) error {
	log.Println("EnumerationType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.EnumerationType
	failures := 0

	c = cache.Get("prod_price__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Price Parent Enum Type").
		AddChildren(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		AddChildren(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_price__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_price__enumerationtype", c)
	}

	c = cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Price Input Parameter").
		SetParent(cache.Get("prod_price__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("prip_product_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_prod_cat_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_prod_clg_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_prod_feat_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_prod_sgrp_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_website_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_quantity__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_party_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_party_grp_mem__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_party_class__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_role_type__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_list_price__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prip_currency_uomid__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_price_in_param__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_price_in_param__enumerationtype", c)
	}

	c = cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Price Condition").
		SetParent(cache.Get("prod_price__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("prc_eq__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prc_neq__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prc_lt__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prc_lte__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prc_gt__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prc_gte__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_price_cond__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_price_cond__enumerationtype", c)
	}

	c = cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inventory Reservation Order").
		AddEnumerations(cache.Get("invro_fifo_rec__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("invro_lifo_rec__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("invro_fifo_exp__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("invro_lifo_exp__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("invro_gunit_cost__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("invro_lunit_cost__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inv_res_order__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inv_res_order__enumerationtype", c)
	}

	c = cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inventory Item Detail Reason").
		AddEnumerations(cache.Get("var_lost__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_stolen__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_found__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_damaged__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_sample__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_integr__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_misship_ordered__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("var_misship_shipped__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update iid_reason__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("iid_reason__enumerationtype", c)
	}

	c = cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Store Payment Service Type").
		AddEnumerations(cache.Get("prds_pay_auth__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_auth_verify__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_reauth__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_release__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_capture__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_refund__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_credit__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pay_external__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_paysvc__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_paysvc__enumerationtype", c)
	}

	c = cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Store Email Notification").
		AddEnumerations(cache.Get("prds_cust_register__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_confirm__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_complete__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_backorder__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_change__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_payretry__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_rtn_accept__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_rtn_complete__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_rtn_cancel__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_pwd_retrieve__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_tell_friend__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_gc_purchase__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_gc_reload__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_quo_created__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_quo_confirm__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_odr_ship_complt__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_email__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_email__enumerationtype", c)
	}

	c = cache.Get("facloc_type__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Facility Location Type").
		AddEnumerations(cache.Get("flt_pickloc__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("flt_bulk__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update facloc_type__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("facloc_type__enumerationtype", c)
	}

	c = cache.Get("kwovrd_trgt_type__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Keyword Override Target Type").
		AddEnumerations(cache.Get("kott_prodcat__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("kott_product__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("kott_ofburl__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("kott_aurl__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update kwovrd_trgt_type__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("kwovrd_trgt_type__enumerationtype", c)
	}

	c = cache.Get("pcat_link_type__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Category Link Type").
		AddEnumerations(cache.Get("pclt_search_param__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pclt_abs_url__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pclt_cat_id__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pcat_link_type__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pcat_link_type__enumerationtype", c)
	}

	c = cache.Get("prod_geo__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Geo Data").
		AddEnumerations(cache.Get("pg_purch_include__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pg_purch_exclude__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pg_ship_include__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pg_ship_exclude__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_geo__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_geo__enumerationtype", c)
	}

	c = cache.Get("prod_promo__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Promotion Parent Enum Type").
		AddChildren(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		AddChildren(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		AddChildren(cache.Get("prod_promo_pcappl__enumerationtype").(*ent.EnumerationType)).
		AddChildren(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_promo__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_promo__enumerationtype", c)
	}

	c = cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Promotion Input Parameter").
		SetParent(cache.Get("prod_promo__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("ppip_order_total__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_product_total__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_product_amount__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_product_quant__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_new_acct__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_party_id__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_party_grp_mem__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_party_class__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_role_type__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_orst_hist__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_recurrence__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_orst_year__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_orst_last_year__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_lpmup_amt__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_lpmup_per__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_order_shiptotal__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_service__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppip_geo_id__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_promo_in_param__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_promo_in_param__enumerationtype", c)
	}

	c = cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Promotion Condition").
		SetParent(cache.Get("prod_promo__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("ppc_eq__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppc_neq__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppc_lt__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppc_lte__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppc_gt__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("ppc_gte__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_promo_cond__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_promo_cond__enumerationtype", c)
	}

	c = cache.Get("prod_promo_pcappl__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Promotion Product/Category Application Type").
		SetParent(cache.Get("prod_promo__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("pppa_include__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pppa_exclude__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("pppa_always__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_promo_pcappl__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_promo_pcappl__enumerationtype", c)
	}

	c = cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Promotion Action").
		SetParent(cache.Get("prod_promo__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("promo_gwp__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_prod_disc__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_prod_amdisc__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_prod_price__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_order_percent__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_order_amount__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_prod_spprc__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_ship_charge__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_service__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("promo_tax_percent__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_promo_action__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_promo_action__enumerationtype", c)
	}

	c = cache.Get("prod_rating_type__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Rating Field Type").
		AddEnumerations(cache.Get("prdr_min__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prdr_max__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prdr_flat__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_rating_type__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_rating_type__enumerationtype", c)
	}

	c = cache.Get("prod_vvmethod__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Virtual Variant Method").
		AddEnumerations(cache.Get("vv_featuretree__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("vv_varianttree__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_vvmethod__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_vvmethod__enumerationtype", c)
	}

	c = cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Product Requirement Method").
		AddEnumerations(cache.Get("prodrqm_none__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_auto__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_stock__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_stock_atp__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_atp__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_ds__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prodrqm_dsatp__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_req_method__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_req_method__enumerationtype", c)
	}

	c = cache.Get("image_reject_reason__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Image Reject Reason").
		AddEnumerations(cache.Get("retake_photo__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("remove_logo__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("other__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update image_reject_reason__enumerationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("image_reject_reason__enumerationtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
