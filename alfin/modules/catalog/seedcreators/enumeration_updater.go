package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateEnumeration(ctx context.Context) error {
	log.Println("Enumeration updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Enumeration
	failures := 0

	c = cache.Get("prip_product_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_ID").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Product").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_product_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_product_id__enumeration", c)
	}

	c = cache.Get("prip_prod_cat_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_CATEGORY_ID").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Product Category").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_prod_cat_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_prod_cat_id__enumeration", c)
	}

	c = cache.Get("prip_prod_clg_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_CATALOG_ID").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Product Catalog").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_prod_clg_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_prod_clg_id__enumeration", c)
	}

	c = cache.Get("prip_prod_feat_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_FEATURE_ID").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Product Feature").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_prod_feat_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_prod_feat_id__enumeration", c)
	}

	c = cache.Get("prip_prod_sgrp_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_STORE_GRP_ID").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Product Store Group").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_prod_sgrp_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_prod_sgrp_id__enumeration", c)
	}

	c = cache.Get("prip_website_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("WEBSITE_ID").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Website").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_website_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_website_id__enumeration", c)
	}

	c = cache.Get("prip_quantity__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("QUANTITY").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Quantity").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_quantity__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_quantity__enumeration", c)
	}

	c = cache.Get("prip_party_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_ID").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Party").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_party_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_party_id__enumeration", c)
	}

	c = cache.Get("prip_party_grp_mem__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_GROUP_MEMBER").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Party Group Member").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_party_grp_mem__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_party_grp_mem__enumeration", c)
	}

	c = cache.Get("prip_party_class__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_CLASS").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Party Classification").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_party_class__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_party_class__enumeration", c)
	}

	c = cache.Get("prip_role_type__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ROLE_TYPE").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Role Type").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_role_type__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_role_type__enumeration", c)
	}

	c = cache.Get("prip_list_price__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LIST_PRICE").
		SetSequenceID(common.ParseId("12")).
		SetDescription("List Price").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_list_price__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_list_price__enumeration", c)
	}

	c = cache.Get("prip_currency_uomid__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("CURRENCY_UOMID").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Currency UomId").
		SetEnumerationType(cache.Get("prod_price_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prip_currency_uomid__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prip_currency_uomid__enumeration", c)
	}

	c = cache.Get("prc_eq__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("EQ").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Is").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_eq__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_eq__enumeration", c)
	}

	c = cache.Get("prc_neq__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("NEQ").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Is Not").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_neq__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_neq__enumeration", c)
	}

	c = cache.Get("prc_lt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Is Less Than").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_lt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_lt__enumeration", c)
	}

	c = cache.Get("prc_lte__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LTE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Is Less Than or Equal To").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_lte__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_lte__enumeration", c)
	}

	c = cache.Get("prc_gt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Is Greater Than").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_gt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_gt__enumeration", c)
	}

	c = cache.Get("prc_gte__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GTE").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Is Greater Than or Equal To").
		SetEnumerationType(cache.Get("prod_price_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prc_gte__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prc_gte__enumeration", c)
	}

	c = cache.Get("invro_fifo_rec__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("FIFO_REC").
		SetSequenceID(common.ParseId("01")).
		SetDescription("FIFO Received").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_fifo_rec__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_fifo_rec__enumeration", c)
	}

	c = cache.Get("invro_lifo_rec__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LIFO_REC").
		SetSequenceID(common.ParseId("02")).
		SetDescription("LIFO Received").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_lifo_rec__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_lifo_rec__enumeration", c)
	}

	c = cache.Get("invro_fifo_exp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("FIFO_EXP").
		SetSequenceID(common.ParseId("03")).
		SetDescription("FIFO Expire").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_fifo_exp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_fifo_exp__enumeration", c)
	}

	c = cache.Get("invro_lifo_exp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LIFO_EXP").
		SetSequenceID(common.ParseId("04")).
		SetDescription("LIFO Expire").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_lifo_exp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_lifo_exp__enumeration", c)
	}

	c = cache.Get("invro_gunit_cost__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GUNIT_COST").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Greater Unit Cost").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_gunit_cost__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_gunit_cost__enumeration", c)
	}

	c = cache.Get("invro_lunit_cost__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LUNIT_COST").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Less Unit Cost").
		SetEnumerationType(cache.Get("inv_res_order__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update invro_lunit_cost__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("invro_lunit_cost__enumeration", c)
	}

	c = cache.Get("var_lost__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_LOST").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Lost").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_lost__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_lost__enumeration", c)
	}

	c = cache.Get("var_stolen__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_STOLEN").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Stolen").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_stolen__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_stolen__enumeration", c)
	}

	c = cache.Get("var_found__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_FOUND").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Found").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_found__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_found__enumeration", c)
	}

	c = cache.Get("var_damaged__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_DAMAGED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Damaged").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_damaged__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_damaged__enumeration", c)
	}

	c = cache.Get("var_sample__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_SAMPLE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Sample (Giveaway)").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_sample__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_sample__enumeration", c)
	}

	c = cache.Get("var_integr__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_INTEGR").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Integration").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_integr__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_integr__enumeration", c)
	}

	c = cache.Get("var_misship_ordered__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_MISSHIP_ORDERED").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Mis-shipped Item Ordered (+)").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_misship_ordered__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_misship_ordered__enumeration", c)
	}

	c = cache.Get("var_misship_shipped__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VAR_MISSHIP_SHIPPED").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Mis-shipped Item Shipped (-)").
		SetEnumerationType(cache.Get("iid_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update var_misship_shipped__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("var_misship_shipped__enumeration", c)
	}

	c = cache.Get("prds_pay_auth__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_AUTH").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Payment Authorization Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_auth__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_auth__enumeration", c)
	}

	c = cache.Get("prds_pay_auth_verify__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_AUTH_VERIFY").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Auth Verification Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_auth_verify__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_auth_verify__enumeration", c)
	}

	c = cache.Get("prds_pay_reauth__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_REAUTH").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Payment Re-Authorization Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_reauth__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_reauth__enumeration", c)
	}

	c = cache.Get("prds_pay_release__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_RELEASE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Payment Release Authorization Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_release__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_release__enumeration", c)
	}

	c = cache.Get("prds_pay_capture__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_CAPTURE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Payment Capture Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_capture__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_capture__enumeration", c)
	}

	c = cache.Get("prds_pay_refund__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_REFUND").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Payment Refund Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_refund__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_refund__enumeration", c)
	}

	c = cache.Get("prds_pay_credit__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_CREDIT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Credit Service").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_credit__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_credit__enumeration", c)
	}

	c = cache.Get("prds_pay_external__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PAY_EXTERNAL").
		SetSequenceID(common.ParseId("07")).
		SetDescription("External Payment (No Service)").
		SetEnumerationType(cache.Get("prds_paysvc__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pay_external__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pay_external__enumeration", c)
	}

	c = cache.Get("prds_cust_register__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("CUST_REGISTER").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Registration").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_cust_register__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_cust_register__enumeration", c)
	}

	c = cache.Get("prds_odr_confirm__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ODR_CONFIRM").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Confirmation").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_confirm__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_confirm__enumeration", c)
	}

	c = cache.Get("prds_odr_complete__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ODR_COMPLETE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Complete").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_complete__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_complete__enumeration", c)
	}

	c = cache.Get("prds_odr_backorder__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ODR_BACKORDER").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Back-Order").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_backorder__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_backorder__enumeration", c)
	}

	c = cache.Get("prds_odr_change__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ODR_CHANGE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Order Change").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_change__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_change__enumeration", c)
	}

	c = cache.Get("prds_odr_payretry__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ODR_PAYRETRY").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Retry").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_payretry__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_payretry__enumeration", c)
	}

	c = cache.Get("prds_rtn_accept__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("RTN_ACCEPT").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Return Accepted").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_rtn_accept__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_rtn_accept__enumeration", c)
	}

	c = cache.Get("prds_rtn_complete__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("RTN_COMPLETE").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Return Completed").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_rtn_complete__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_rtn_complete__enumeration", c)
	}

	c = cache.Get("prds_rtn_cancel__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("RTN_CANCEL").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Return Cancelled").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_rtn_cancel__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_rtn_cancel__enumeration", c)
	}

	c = cache.Get("prds_pwd_retrieve__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PWD_RETRIEVE").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Retrieve Password").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_pwd_retrieve__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_pwd_retrieve__enumeration", c)
	}

	c = cache.Get("prds_tell_friend__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("TELL_FRIEND").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Tell-A-Friend").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_tell_friend__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_tell_friend__enumeration", c)
	}

	c = cache.Get("prds_gc_purchase__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GC_PURCHASE").
		SetSequenceID(common.ParseId("12")).
		SetDescription("Gift-Card Purchase").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_gc_purchase__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_gc_purchase__enumeration", c)
	}

	c = cache.Get("prds_gc_reload__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GC_RELOAD").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Gift-Card Reload").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_gc_reload__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_gc_reload__enumeration", c)
	}

	c = cache.Get("prds_quo_created__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("QUO_CREATED").
		SetSequenceID(common.ParseId("14")).
		SetDescription("Quote Created").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_quo_created__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_quo_created__enumeration", c)
	}

	c = cache.Get("prds_quo_confirm__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("QUO_CONFIRM").
		SetSequenceID(common.ParseId("15")).
		SetDescription("Quote Confirmation").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_quo_confirm__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_quo_confirm__enumeration", c)
	}

	c = cache.Get("prds_odr_ship_complt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SHP_COMPLETE").
		SetSequenceID(common.ParseId("16")).
		SetDescription("Shipment Complete").
		SetEnumerationType(cache.Get("prds_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_odr_ship_complt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_odr_ship_complt__enumeration", c)
	}

	c = cache.Get("flt_pickloc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PICKLOC").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pick/Primary").
		SetEnumerationType(cache.Get("facloc_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flt_pickloc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("flt_pickloc__enumeration", c)
	}

	c = cache.Get("flt_bulk__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("BULK").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Bulk").
		SetEnumerationType(cache.Get("facloc_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flt_bulk__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("flt_bulk__enumeration", c)
	}

	c = cache.Get("kott_prodcat__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODCAT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Product Category").
		SetEnumerationType(cache.Get("kwovrd_trgt_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update kott_prodcat__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("kott_prodcat__enumeration", c)
	}

	c = cache.Get("kott_product__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Product").
		SetEnumerationType(cache.Get("kwovrd_trgt_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update kott_product__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("kott_product__enumeration", c)
	}

	c = cache.Get("kott_ofburl__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("OFBURL").
		SetSequenceID(common.ParseId("03")).
		SetDescription("OFBiz URL").
		SetEnumerationType(cache.Get("kwovrd_trgt_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update kott_ofburl__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("kott_ofburl__enumeration", c)
	}

	c = cache.Get("kott_aurl__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("AURL").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Absolute URL").
		SetEnumerationType(cache.Get("kwovrd_trgt_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update kott_aurl__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("kott_aurl__enumeration", c)
	}

	c = cache.Get("pclt_search_param__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SEARCH_PARAM").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Search Parameters").
		SetEnumerationType(cache.Get("pcat_link_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pclt_search_param__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pclt_search_param__enumeration", c)
	}

	c = cache.Get("pclt_abs_url__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ABS_URL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Absolute URL").
		SetEnumerationType(cache.Get("pcat_link_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pclt_abs_url__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pclt_abs_url__enumeration", c)
	}

	c = cache.Get("pclt_cat_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("CAT_ID").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Category ID").
		SetEnumerationType(cache.Get("pcat_link_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pclt_cat_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pclt_cat_id__enumeration", c)
	}

	c = cache.Get("pg_purch_include__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PURCHASE_INCLUDE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Purchase Include Geo").
		SetEnumerationType(cache.Get("prod_geo__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pg_purch_include__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pg_purch_include__enumeration", c)
	}

	c = cache.Get("pg_purch_exclude__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PURCHASE_EXCLUDE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Purchase Exclude Geo").
		SetEnumerationType(cache.Get("prod_geo__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pg_purch_exclude__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pg_purch_exclude__enumeration", c)
	}

	c = cache.Get("pg_ship_include__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SHIPMENT_INCLUDE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Shipment Include Geo").
		SetEnumerationType(cache.Get("prod_geo__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pg_ship_include__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pg_ship_include__enumeration", c)
	}

	c = cache.Get("pg_ship_exclude__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SHIPMENT_EXCLUDE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Shipment Exclude Geo").
		SetEnumerationType(cache.Get("prod_geo__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pg_ship_exclude__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pg_ship_exclude__enumeration", c)
	}

	c = cache.Get("ppip_order_total__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORDER_TOTAL").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Cart Sub-total").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_order_total__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_order_total__enumeration", c)
	}

	c = cache.Get("ppip_product_total__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_TOTAL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Total Amount of Product").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_total__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_total__enumeration", c)
	}

	c = cache.Get("ppip_product_amount__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_AMOUNT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("X Amount of Product").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_amount__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_amount__enumeration", c)
	}

	c = cache.Get("ppip_product_quant__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PRODUCT_QUANT").
		SetSequenceID(common.ParseId("04")).
		SetDescription("X Quantity of Product").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_quant__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_quant__enumeration", c)
	}

	c = cache.Get("ppip_new_acct__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("NEW_ACCT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Account Days Since Created").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_new_acct__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_new_acct__enumeration", c)
	}

	c = cache.Get("ppip_party_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_ID").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Party").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_id__enumeration", c)
	}

	c = cache.Get("ppip_party_grp_mem__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_GROUP_MEMBER").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Party Group Member").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_grp_mem__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_grp_mem__enumeration", c)
	}

	c = cache.Get("ppip_party_class__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTY_CLASS").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Party Classification").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_class__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_class__enumeration", c)
	}

	c = cache.Get("ppip_role_type__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ROLE_TYPE").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Role Type").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_role_type__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_role_type__enumeration", c)
	}

	c = cache.Get("ppip_orst_hist__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORST_HIST").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Order sub-total X in last Y Months").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_hist__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_hist__enumeration", c)
	}

	c = cache.Get("ppip_recurrence__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROMO_RECURRENCE").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Promotion Recurrence").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_recurrence__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_recurrence__enumeration", c)
	}

	c = cache.Get("ppip_orst_year__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORST_YEAR").
		SetSequenceID(common.ParseId("12")).
		SetDescription("Order sub-total X since beginning of current year").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_year__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_year__enumeration", c)
	}

	c = cache.Get("ppip_orst_last_year__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORST_LAST_YEAR").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Order sub-total X last year").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_last_year__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_last_year__enumeration", c)
	}

	c = cache.Get("ppip_lpmup_amt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LPMUP_AMT").
		SetSequenceID(common.ParseId("14")).
		SetDescription("List Price minus Unit Price (Amount)").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_lpmup_amt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_lpmup_amt__enumeration", c)
	}

	c = cache.Get("ppip_lpmup_per__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LPMUP_PER").
		SetSequenceID(common.ParseId("15")).
		SetDescription("List Price minus Unit Price (Percent)").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_lpmup_per__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_lpmup_per__enumeration", c)
	}

	c = cache.Get("ppip_order_shiptotal__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORDER_SHIP_TOTAL").
		SetSequenceID(common.ParseId("16")).
		SetDescription("Shipping Total").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_order_shiptotal__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_order_shiptotal__enumeration", c)
	}

	c = cache.Get("ppip_service__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SERVICE").
		SetSequenceID(common.ParseId("17")).
		SetDescription("Call Service").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_service__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_service__enumeration", c)
	}

	c = cache.Get("ppip_geo_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GEO_ID").
		SetSequenceID(common.ParseId("18")).
		SetDescription("Shipping Destination").
		SetEnumerationType(cache.Get("prod_promo_in_param__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_geo_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_geo_id__enumeration", c)
	}

	c = cache.Get("ppc_eq__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("EQ").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Is").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_eq__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_eq__enumeration", c)
	}

	c = cache.Get("ppc_neq__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("NEQ").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Is Not").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_neq__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_neq__enumeration", c)
	}

	c = cache.Get("ppc_lt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Is Less Than").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_lt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_lt__enumeration", c)
	}

	c = cache.Get("ppc_lte__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("LTE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Is Less Than or Equal To").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_lte__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_lte__enumeration", c)
	}

	c = cache.Get("ppc_gt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Is Greater Than").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_gt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_gt__enumeration", c)
	}

	c = cache.Get("ppc_gte__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GTE").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Is Greater Than or Equal To").
		SetEnumerationType(cache.Get("prod_promo_cond__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_gte__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_gte__enumeration", c)
	}

	c = cache.Get("pppa_include__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("INCLUDE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Include").
		SetEnumerationType(cache.Get("prod_promo_pcappl__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pppa_include__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pppa_include__enumeration", c)
	}

	c = cache.Get("pppa_exclude__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("EXCLUDE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Exclude").
		SetEnumerationType(cache.Get("prod_promo_pcappl__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pppa_exclude__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pppa_exclude__enumeration", c)
	}

	c = cache.Get("pppa_always__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ALWAYS").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Always Include").
		SetEnumerationType(cache.Get("prod_promo_pcappl__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pppa_always__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pppa_always__enumeration", c)
	}

	c = cache.Get("promo_gwp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("GWP").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Gift With Purchase").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_gwp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_gwp__enumeration", c)
	}

	c = cache.Get("promo_prod_disc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_DISC").
		SetSequenceID(common.ParseId("03")).
		SetDescription("X Product for Y% Discount").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_disc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_disc__enumeration", c)
	}

	c = cache.Get("promo_prod_amdisc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_AMDISC").
		SetSequenceID(common.ParseId("04")).
		SetDescription("X Product for Y Discount").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_amdisc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_amdisc__enumeration", c)
	}

	c = cache.Get("promo_prod_price__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_PRICE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("X Product for Y Price").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_price__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_price__enumeration", c)
	}

	c = cache.Get("promo_order_percent__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORDER_PERCENT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Order Percent Discount").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_order_percent__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_order_percent__enumeration", c)
	}

	c = cache.Get("promo_order_amount__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ORDER_AMOUNT").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Order Amount Flat").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_order_amount__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_order_amount__enumeration", c)
	}

	c = cache.Get("promo_prod_spprc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PROD_SPPRC").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Product for [Special Promo] Price").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_spprc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_spprc__enumeration", c)
	}

	c = cache.Get("promo_ship_charge__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SHIP_CHARGE").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Shipping X% discount").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_ship_charge__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_ship_charge__enumeration", c)
	}

	c = cache.Get("promo_service__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SERVICE").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Call Service").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_service__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_service__enumeration", c)
	}

	c = cache.Get("promo_tax_percent__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("TAX_PERCENT").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Tax % Discount").
		SetEnumerationType(cache.Get("prod_promo_action__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_tax_percent__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_tax_percent__enumeration", c)
	}

	c = cache.Get("prdr_min__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("MIN").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Min Rating").
		SetEnumerationType(cache.Get("prod_rating_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prdr_min__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prdr_min__enumeration", c)
	}

	c = cache.Get("prdr_max__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("MAX").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Max Rating").
		SetEnumerationType(cache.Get("prod_rating_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prdr_max__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prdr_max__enumeration", c)
	}

	c = cache.Get("prdr_flat__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("FLAT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Rating Override").
		SetEnumerationType(cache.Get("prod_rating_type__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prdr_flat__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prdr_flat__enumeration", c)
	}

	c = cache.Get("vv_featuretree__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("FEATURETREE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Feature tree Generation").
		SetEnumerationType(cache.Get("prod_vvmethod__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update vv_featuretree__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("vv_featuretree__enumeration", c)
	}

	c = cache.Get("vv_varianttree__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("VARIANTTREE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Variant Tree generation").
		SetEnumerationType(cache.Get("prod_vvmethod__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update vv_varianttree__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("vv_varianttree__enumeration", c)
	}

	c = cache.Get("prodrqm_none__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("NONE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("No Requirement Created").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_none__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_none__enumeration", c)
	}

	c = cache.Get("prodrqm_auto__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("AUTO").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Automatic For Every Sales Order").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_auto__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_auto__enumeration", c)
	}

	c = cache.Get("prodrqm_stock__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("STOCK_QOH").
		SetSequenceID(common.ParseId("03")).
		SetDescription("When QOH Reaches Minimum Stock for Product-Facility").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_stock__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_stock__enumeration", c)
	}

	c = cache.Get("prodrqm_stock_atp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("STOCK_ATP").
		SetSequenceID(common.ParseId("04")).
		SetDescription("When ATP Reaches Minimum Stock for Product-Facility").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_stock_atp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_stock_atp__enumeration", c)
	}

	c = cache.Get("prodrqm_atp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ATP").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Requirement for order when ATP Reaches Minimum Stock for Product-Facility").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_atp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_atp__enumeration", c)
	}

	c = cache.Get("prodrqm_ds__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("DROPS").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Drop-ship only").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_ds__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_ds__enumeration", c)
	}

	c = cache.Get("prodrqm_dsatp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("DROPS_ATP").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Auto drop-ship on low quantity").
		SetEnumerationType(cache.Get("prod_req_method__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prodrqm_dsatp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prodrqm_dsatp__enumeration", c)
	}

	c = cache.Get("retake_photo__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("RETAKE_PHOTO").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Re-take Photo").
		SetEnumerationType(cache.Get("image_reject_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update retake_photo__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("retake_photo__enumeration", c)
	}

	c = cache.Get("remove_logo__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("REMOVE_LOGO").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Remove Logo").
		SetEnumerationType(cache.Get("image_reject_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update remove_logo__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("remove_logo__enumeration", c)
	}

	c = cache.Get("other__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("OTHER").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Other").
		SetEnumerationType(cache.Get("image_reject_reason__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update other__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("other__enumeration", c)
	}

	c = cache.Get("direct_store_delivry__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("DIRECT_STORE_DELIVRY").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Direct Store Delivery").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update direct_store_delivry__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("direct_store_delivry__enumeration", c)
	}

	c = cache.Get("ppip_product_amount__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PRODUCT_AMOUNT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_amount__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_amount__enumeration", c)
	}

	c = cache.Get("ppip_product_total__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PRODUCT_TOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_total__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_total__enumeration", c)
	}

	c = cache.Get("ppip_product_quant__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PRODUCT_QUANT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_product_quant__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_product_quant__enumeration", c)
	}

	c = cache.Get("ppip_new_acct__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_NEW_ACCT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_new_acct__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_new_acct__enumeration", c)
	}

	c = cache.Get("ppip_party_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PARTY_ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_id__enumeration", c)
	}

	c = cache.Get("ppip_party_grp_mem__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PARTY_GRP_MEM").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_grp_mem__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_grp_mem__enumeration", c)
	}

	c = cache.Get("ppip_party_class__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_PARTY_CLASS").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_party_class__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_party_class__enumeration", c)
	}

	c = cache.Get("ppip_role_type__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ROLE_TYPE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_role_type__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_role_type__enumeration", c)
	}

	c = cache.Get("ppip_geo_id__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_GEO_ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_geo_id__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_geo_id__enumeration", c)
	}

	c = cache.Get("ppip_order_total__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ORDER_TOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_order_total__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_order_total__enumeration", c)
	}

	c = cache.Get("ppip_orst_hist__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ORST_HIST").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_hist__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_hist__enumeration", c)
	}

	c = cache.Get("ppip_orst_year__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ORST_YEAR").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_year__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_year__enumeration", c)
	}

	c = cache.Get("ppip_orst_last_year__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ORST_LAST_YEAR").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_orst_last_year__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_orst_last_year__enumeration", c)
	}

	c = cache.Get("ppip_recurrence__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_RECURRENCE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_recurrence__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_recurrence__enumeration", c)
	}

	c = cache.Get("ppip_order_shiptotal__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_ORDER_SHIPTOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_order_shiptotal__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_order_shiptotal__enumeration", c)
	}

	c = cache.Get("ppip_lpmup_amt__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_LPMUP_AMT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_lpmup_amt__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_lpmup_amt__enumeration", c)
	}

	c = cache.Get("ppip_lpmup_per__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPC_LPMUP_PER").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppip_lpmup_per__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppip_lpmup_per__enumeration", c)
	}

	c = cache.Get("promo_gwp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_GWP").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_gwp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_gwp__enumeration", c)
	}

	c = cache.Get("promo_free_shipping__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_FREE_SHIPPING").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_free_shipping__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_free_shipping__enumeration", c)
	}

	c = cache.Get("promo_prod_disc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_PROD_DISC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_disc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_disc__enumeration", c)
	}

	c = cache.Get("promo_prod_amdisc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_PROD_AMDISC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_amdisc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_amdisc__enumeration", c)
	}

	c = cache.Get("promo_prod_price__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_PROD_PRICE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_price__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_price__enumeration", c)
	}

	c = cache.Get("promo_order_percent__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_ORDER_PERCENT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_order_percent__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_order_percent__enumeration", c)
	}

	c = cache.Get("promo_order_amount__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_ORDER_AMOUNT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_order_amount__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_order_amount__enumeration", c)
	}

	c = cache.Get("promo_prod_spprc__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_PROD_SPPRC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_prod_spprc__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_prod_spprc__enumeration", c)
	}

	c = cache.Get("promo_tax_percent__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_TAX_PERCENT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_tax_percent__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_tax_percent__enumeration", c)
	}

	c = cache.Get("promo_ship_charge__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PPA_SHIP_CHARGE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update promo_ship_charge__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("promo_ship_charge__enumeration", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
