package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateEnumeration(ctx context.Context) error {
	log.Println("Enumeration creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Enumeration

	c, err = client.Enumeration.Create().SetStringRef("prip_product_id__enumeration").
		SetEnumCode("PRODUCT_ID").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_product_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_product_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_prod_cat_id__enumeration").
		SetEnumCode("PRODUCT_CATEGORY_ID").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Product Category").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_prod_cat_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_prod_cat_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_prod_clg_id__enumeration").
		SetEnumCode("PROD_CATALOG_ID").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Product Catalog").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_prod_clg_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_prod_clg_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_prod_feat_id__enumeration").
		SetEnumCode("PRODUCT_FEATURE_ID").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Product Feature").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_prod_feat_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_prod_feat_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_prod_sgrp_id__enumeration").
		SetEnumCode("PROD_STORE_GRP_ID").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Product Store Group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_prod_sgrp_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_prod_sgrp_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_website_id__enumeration").
		SetEnumCode("WEBSITE_ID").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Website").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_website_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_website_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_quantity__enumeration").
		SetEnumCode("QUANTITY").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Quantity").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_quantity__enumeration: %v", err)
		return err
	}
	cache.Put("prip_quantity__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_party_id__enumeration").
		SetEnumCode("PARTY_ID").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_party_id__enumeration: %v", err)
		return err
	}
	cache.Put("prip_party_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_party_grp_mem__enumeration").
		SetEnumCode("PARTY_GROUP_MEMBER").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Party Group Member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_party_grp_mem__enumeration: %v", err)
		return err
	}
	cache.Put("prip_party_grp_mem__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_party_class__enumeration").
		SetEnumCode("PARTY_CLASS").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Party Classification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_party_class__enumeration: %v", err)
		return err
	}
	cache.Put("prip_party_class__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_role_type__enumeration").
		SetEnumCode("ROLE_TYPE").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Role Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_role_type__enumeration: %v", err)
		return err
	}
	cache.Put("prip_role_type__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_list_price__enumeration").
		SetEnumCode("LIST_PRICE").
		SetSequenceID(common.ParseId("12")).
		SetDescription("List Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_list_price__enumeration: %v", err)
		return err
	}
	cache.Put("prip_list_price__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prip_currency_uomid__enumeration").
		SetEnumCode("CURRENCY_UOMID").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Currency UomId").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prip_currency_uomid__enumeration: %v", err)
		return err
	}
	cache.Put("prip_currency_uomid__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_eq__enumeration").
		SetEnumCode("EQ").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Is").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_eq__enumeration: %v", err)
		return err
	}
	cache.Put("prc_eq__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_neq__enumeration").
		SetEnumCode("NEQ").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Is Not").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_neq__enumeration: %v", err)
		return err
	}
	cache.Put("prc_neq__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_lt__enumeration").
		SetEnumCode("LT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Is Less Than").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_lt__enumeration: %v", err)
		return err
	}
	cache.Put("prc_lt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_lte__enumeration").
		SetEnumCode("LTE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Is Less Than or Equal To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_lte__enumeration: %v", err)
		return err
	}
	cache.Put("prc_lte__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_gt__enumeration").
		SetEnumCode("GT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Is Greater Than").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_gt__enumeration: %v", err)
		return err
	}
	cache.Put("prc_gt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prc_gte__enumeration").
		SetEnumCode("GTE").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Is Greater Than or Equal To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prc_gte__enumeration: %v", err)
		return err
	}
	cache.Put("prc_gte__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_fifo_rec__enumeration").
		SetEnumCode("FIFO_REC").
		SetSequenceID(common.ParseId("01")).
		SetDescription("FIFO Received").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_fifo_rec__enumeration: %v", err)
		return err
	}
	cache.Put("invro_fifo_rec__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_lifo_rec__enumeration").
		SetEnumCode("LIFO_REC").
		SetSequenceID(common.ParseId("02")).
		SetDescription("LIFO Received").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_lifo_rec__enumeration: %v", err)
		return err
	}
	cache.Put("invro_lifo_rec__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_fifo_exp__enumeration").
		SetEnumCode("FIFO_EXP").
		SetSequenceID(common.ParseId("03")).
		SetDescription("FIFO Expire").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_fifo_exp__enumeration: %v", err)
		return err
	}
	cache.Put("invro_fifo_exp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_lifo_exp__enumeration").
		SetEnumCode("LIFO_EXP").
		SetSequenceID(common.ParseId("04")).
		SetDescription("LIFO Expire").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_lifo_exp__enumeration: %v", err)
		return err
	}
	cache.Put("invro_lifo_exp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_gunit_cost__enumeration").
		SetEnumCode("GUNIT_COST").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Greater Unit Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_gunit_cost__enumeration: %v", err)
		return err
	}
	cache.Put("invro_gunit_cost__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("invro_lunit_cost__enumeration").
		SetEnumCode("LUNIT_COST").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Less Unit Cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create invro_lunit_cost__enumeration: %v", err)
		return err
	}
	cache.Put("invro_lunit_cost__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_lost__enumeration").
		SetEnumCode("VAR_LOST").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Lost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_lost__enumeration: %v", err)
		return err
	}
	cache.Put("var_lost__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_stolen__enumeration").
		SetEnumCode("VAR_STOLEN").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Stolen").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_stolen__enumeration: %v", err)
		return err
	}
	cache.Put("var_stolen__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_found__enumeration").
		SetEnumCode("VAR_FOUND").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Found").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_found__enumeration: %v", err)
		return err
	}
	cache.Put("var_found__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_damaged__enumeration").
		SetEnumCode("VAR_DAMAGED").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Damaged").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_damaged__enumeration: %v", err)
		return err
	}
	cache.Put("var_damaged__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_sample__enumeration").
		SetEnumCode("VAR_SAMPLE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Sample (Giveaway)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_sample__enumeration: %v", err)
		return err
	}
	cache.Put("var_sample__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_integr__enumeration").
		SetEnumCode("VAR_INTEGR").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Integration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_integr__enumeration: %v", err)
		return err
	}
	cache.Put("var_integr__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_misship_ordered__enumeration").
		SetEnumCode("VAR_MISSHIP_ORDERED").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Mis-shipped Item Ordered (+)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_misship_ordered__enumeration: %v", err)
		return err
	}
	cache.Put("var_misship_ordered__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("var_misship_shipped__enumeration").
		SetEnumCode("VAR_MISSHIP_SHIPPED").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Mis-shipped Item Shipped (-)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create var_misship_shipped__enumeration: %v", err)
		return err
	}
	cache.Put("var_misship_shipped__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_auth__enumeration").
		SetEnumCode("PAY_AUTH").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Payment Authorization Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_auth__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_auth__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_auth_verify__enumeration").
		SetEnumCode("PAY_AUTH_VERIFY").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Auth Verification Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_auth_verify__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_auth_verify__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_reauth__enumeration").
		SetEnumCode("PAY_REAUTH").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Payment Re-Authorization Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_reauth__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_reauth__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_release__enumeration").
		SetEnumCode("PAY_RELEASE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Payment Release Authorization Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_release__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_release__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_capture__enumeration").
		SetEnumCode("PAY_CAPTURE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Payment Capture Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_capture__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_capture__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_refund__enumeration").
		SetEnumCode("PAY_REFUND").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Payment Refund Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_refund__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_refund__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_credit__enumeration").
		SetEnumCode("PAY_CREDIT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Credit Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_credit__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_credit__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pay_external__enumeration").
		SetEnumCode("PAY_EXTERNAL").
		SetSequenceID(common.ParseId("07")).
		SetDescription("External Payment (No Service)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pay_external__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pay_external__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_cust_register__enumeration").
		SetEnumCode("CUST_REGISTER").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Registration").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_cust_register__enumeration: %v", err)
		return err
	}
	cache.Put("prds_cust_register__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_confirm__enumeration").
		SetEnumCode("ODR_CONFIRM").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Confirmation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_confirm__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_confirm__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_complete__enumeration").
		SetEnumCode("ODR_COMPLETE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_complete__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_complete__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_backorder__enumeration").
		SetEnumCode("ODR_BACKORDER").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Back-Order").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_backorder__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_backorder__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_change__enumeration").
		SetEnumCode("ODR_CHANGE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Order Change").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_change__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_change__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_payretry__enumeration").
		SetEnumCode("ODR_PAYRETRY").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Payment Retry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_payretry__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_payretry__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_rtn_accept__enumeration").
		SetEnumCode("RTN_ACCEPT").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Return Accepted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_rtn_accept__enumeration: %v", err)
		return err
	}
	cache.Put("prds_rtn_accept__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_rtn_complete__enumeration").
		SetEnumCode("RTN_COMPLETE").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Return Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_rtn_complete__enumeration: %v", err)
		return err
	}
	cache.Put("prds_rtn_complete__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_rtn_cancel__enumeration").
		SetEnumCode("RTN_CANCEL").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Return Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_rtn_cancel__enumeration: %v", err)
		return err
	}
	cache.Put("prds_rtn_cancel__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_pwd_retrieve__enumeration").
		SetEnumCode("PWD_RETRIEVE").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Retrieve Password").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_pwd_retrieve__enumeration: %v", err)
		return err
	}
	cache.Put("prds_pwd_retrieve__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_tell_friend__enumeration").
		SetEnumCode("TELL_FRIEND").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Tell-A-Friend").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_tell_friend__enumeration: %v", err)
		return err
	}
	cache.Put("prds_tell_friend__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_gc_purchase__enumeration").
		SetEnumCode("GC_PURCHASE").
		SetSequenceID(common.ParseId("12")).
		SetDescription("Gift-Card Purchase").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_gc_purchase__enumeration: %v", err)
		return err
	}
	cache.Put("prds_gc_purchase__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_gc_reload__enumeration").
		SetEnumCode("GC_RELOAD").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Gift-Card Reload").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_gc_reload__enumeration: %v", err)
		return err
	}
	cache.Put("prds_gc_reload__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_quo_created__enumeration").
		SetEnumCode("QUO_CREATED").
		SetSequenceID(common.ParseId("14")).
		SetDescription("Quote Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_quo_created__enumeration: %v", err)
		return err
	}
	cache.Put("prds_quo_created__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_quo_confirm__enumeration").
		SetEnumCode("QUO_CONFIRM").
		SetSequenceID(common.ParseId("15")).
		SetDescription("Quote Confirmation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_quo_confirm__enumeration: %v", err)
		return err
	}
	cache.Put("prds_quo_confirm__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_odr_ship_complt__enumeration").
		SetEnumCode("SHP_COMPLETE").
		SetSequenceID(common.ParseId("16")).
		SetDescription("Shipment Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_odr_ship_complt__enumeration: %v", err)
		return err
	}
	cache.Put("prds_odr_ship_complt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("flt_pickloc__enumeration").
		SetEnumCode("PICKLOC").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Pick/Primary").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flt_pickloc__enumeration: %v", err)
		return err
	}
	cache.Put("flt_pickloc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("flt_bulk__enumeration").
		SetEnumCode("BULK").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Bulk").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flt_bulk__enumeration: %v", err)
		return err
	}
	cache.Put("flt_bulk__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("kott_prodcat__enumeration").
		SetEnumCode("PRODCAT").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Product Category").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create kott_prodcat__enumeration: %v", err)
		return err
	}
	cache.Put("kott_prodcat__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("kott_product__enumeration").
		SetEnumCode("PRODUCT").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create kott_product__enumeration: %v", err)
		return err
	}
	cache.Put("kott_product__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("kott_ofburl__enumeration").
		SetEnumCode("OFBURL").
		SetSequenceID(common.ParseId("03")).
		SetDescription("OFBiz URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create kott_ofburl__enumeration: %v", err)
		return err
	}
	cache.Put("kott_ofburl__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("kott_aurl__enumeration").
		SetEnumCode("AURL").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Absolute URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create kott_aurl__enumeration: %v", err)
		return err
	}
	cache.Put("kott_aurl__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pclt_search_param__enumeration").
		SetEnumCode("SEARCH_PARAM").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Search Parameters").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pclt_search_param__enumeration: %v", err)
		return err
	}
	cache.Put("pclt_search_param__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pclt_abs_url__enumeration").
		SetEnumCode("ABS_URL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Absolute URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pclt_abs_url__enumeration: %v", err)
		return err
	}
	cache.Put("pclt_abs_url__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pclt_cat_id__enumeration").
		SetEnumCode("CAT_ID").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Category ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pclt_cat_id__enumeration: %v", err)
		return err
	}
	cache.Put("pclt_cat_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pg_purch_include__enumeration").
		SetEnumCode("PURCHASE_INCLUDE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Purchase Include Geo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pg_purch_include__enumeration: %v", err)
		return err
	}
	cache.Put("pg_purch_include__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pg_purch_exclude__enumeration").
		SetEnumCode("PURCHASE_EXCLUDE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Purchase Exclude Geo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pg_purch_exclude__enumeration: %v", err)
		return err
	}
	cache.Put("pg_purch_exclude__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pg_ship_include__enumeration").
		SetEnumCode("SHIPMENT_INCLUDE").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Shipment Include Geo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pg_ship_include__enumeration: %v", err)
		return err
	}
	cache.Put("pg_ship_include__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pg_ship_exclude__enumeration").
		SetEnumCode("SHIPMENT_EXCLUDE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Shipment Exclude Geo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pg_ship_exclude__enumeration: %v", err)
		return err
	}
	cache.Put("pg_ship_exclude__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_order_total__enumeration").
		SetEnumCode("ORDER_TOTAL").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Cart Sub-total").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_order_total__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_order_total__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_total__enumeration").
		SetEnumCode("PRODUCT_TOTAL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Total Amount of Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_total__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_total__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_amount__enumeration").
		SetEnumCode("PRODUCT_AMOUNT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("X Amount of Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_amount__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_amount__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_quant__enumeration").
		SetEnumCode("PRODUCT_QUANT").
		SetSequenceID(common.ParseId("04")).
		SetDescription("X Quantity of Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_quant__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_quant__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_new_acct__enumeration").
		SetEnumCode("NEW_ACCT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Account Days Since Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_new_acct__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_new_acct__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_id__enumeration").
		SetEnumCode("PARTY_ID").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_id__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_grp_mem__enumeration").
		SetEnumCode("PARTY_GROUP_MEMBER").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Party Group Member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_grp_mem__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_grp_mem__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_class__enumeration").
		SetEnumCode("PARTY_CLASS").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Party Classification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_class__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_class__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_role_type__enumeration").
		SetEnumCode("ROLE_TYPE").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Role Type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_role_type__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_role_type__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_hist__enumeration").
		SetEnumCode("ORST_HIST").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Order sub-total X in last Y Months").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_hist__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_hist__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_recurrence__enumeration").
		SetEnumCode("PROMO_RECURRENCE").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Promotion Recurrence").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_recurrence__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_recurrence__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_year__enumeration").
		SetEnumCode("ORST_YEAR").
		SetSequenceID(common.ParseId("12")).
		SetDescription("Order sub-total X since beginning of current year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_year__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_year__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_last_year__enumeration").
		SetEnumCode("ORST_LAST_YEAR").
		SetSequenceID(common.ParseId("13")).
		SetDescription("Order sub-total X last year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_last_year__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_last_year__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_lpmup_amt__enumeration").
		SetEnumCode("LPMUP_AMT").
		SetSequenceID(common.ParseId("14")).
		SetDescription("List Price minus Unit Price (Amount)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_lpmup_amt__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_lpmup_amt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_lpmup_per__enumeration").
		SetEnumCode("LPMUP_PER").
		SetSequenceID(common.ParseId("15")).
		SetDescription("List Price minus Unit Price (Percent)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_lpmup_per__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_lpmup_per__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_order_shiptotal__enumeration").
		SetEnumCode("ORDER_SHIP_TOTAL").
		SetSequenceID(common.ParseId("16")).
		SetDescription("Shipping Total").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_order_shiptotal__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_order_shiptotal__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_service__enumeration").
		SetEnumCode("SERVICE").
		SetSequenceID(common.ParseId("17")).
		SetDescription("Call Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_service__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_service__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_geo_id__enumeration").
		SetEnumCode("GEO_ID").
		SetSequenceID(common.ParseId("18")).
		SetDescription("Shipping Destination").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_geo_id__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_geo_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_eq__enumeration").
		SetEnumCode("EQ").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Is").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_eq__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_eq__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_neq__enumeration").
		SetEnumCode("NEQ").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Is Not").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_neq__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_neq__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_lt__enumeration").
		SetEnumCode("LT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Is Less Than").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_lt__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_lt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_lte__enumeration").
		SetEnumCode("LTE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Is Less Than or Equal To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_lte__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_lte__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_gt__enumeration").
		SetEnumCode("GT").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Is Greater Than").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_gt__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_gt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppc_gte__enumeration").
		SetEnumCode("GTE").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Is Greater Than or Equal To").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_gte__enumeration: %v", err)
		return err
	}
	cache.Put("ppc_gte__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pppa_include__enumeration").
		SetEnumCode("INCLUDE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Include").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pppa_include__enumeration: %v", err)
		return err
	}
	cache.Put("pppa_include__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pppa_exclude__enumeration").
		SetEnumCode("EXCLUDE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Exclude").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pppa_exclude__enumeration: %v", err)
		return err
	}
	cache.Put("pppa_exclude__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("pppa_always__enumeration").
		SetEnumCode("ALWAYS").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Always Include").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pppa_always__enumeration: %v", err)
		return err
	}
	cache.Put("pppa_always__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_gwp__enumeration").
		SetEnumCode("GWP").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Gift With Purchase").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_gwp__enumeration: %v", err)
		return err
	}
	cache.Put("promo_gwp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_disc__enumeration").
		SetEnumCode("PROD_DISC").
		SetSequenceID(common.ParseId("03")).
		SetDescription("X Product for Y% Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_disc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_disc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_amdisc__enumeration").
		SetEnumCode("PROD_AMDISC").
		SetSequenceID(common.ParseId("04")).
		SetDescription("X Product for Y Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_amdisc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_amdisc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_price__enumeration").
		SetEnumCode("PROD_PRICE").
		SetSequenceID(common.ParseId("05")).
		SetDescription("X Product for Y Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_price__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_price__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_order_percent__enumeration").
		SetEnumCode("ORDER_PERCENT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Order Percent Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_order_percent__enumeration: %v", err)
		return err
	}
	cache.Put("promo_order_percent__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_order_amount__enumeration").
		SetEnumCode("ORDER_AMOUNT").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Order Amount Flat").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_order_amount__enumeration: %v", err)
		return err
	}
	cache.Put("promo_order_amount__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_spprc__enumeration").
		SetEnumCode("PROD_SPPRC").
		SetSequenceID(common.ParseId("08")).
		SetDescription("Product for [Special Promo] Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_spprc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_spprc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_ship_charge__enumeration").
		SetEnumCode("SHIP_CHARGE").
		SetSequenceID(common.ParseId("09")).
		SetDescription("Shipping X% discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_ship_charge__enumeration: %v", err)
		return err
	}
	cache.Put("promo_ship_charge__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_service__enumeration").
		SetEnumCode("SERVICE").
		SetSequenceID(common.ParseId("10")).
		SetDescription("Call Service").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_service__enumeration: %v", err)
		return err
	}
	cache.Put("promo_service__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_tax_percent__enumeration").
		SetEnumCode("TAX_PERCENT").
		SetSequenceID(common.ParseId("11")).
		SetDescription("Tax % Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_tax_percent__enumeration: %v", err)
		return err
	}
	cache.Put("promo_tax_percent__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prdr_min__enumeration").
		SetEnumCode("MIN").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Min Rating").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prdr_min__enumeration: %v", err)
		return err
	}
	cache.Put("prdr_min__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prdr_max__enumeration").
		SetEnumCode("MAX").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Max Rating").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prdr_max__enumeration: %v", err)
		return err
	}
	cache.Put("prdr_max__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prdr_flat__enumeration").
		SetEnumCode("FLAT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Rating Override").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prdr_flat__enumeration: %v", err)
		return err
	}
	cache.Put("prdr_flat__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("vv_featuretree__enumeration").
		SetEnumCode("FEATURETREE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Feature tree Generation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vv_featuretree__enumeration: %v", err)
		return err
	}
	cache.Put("vv_featuretree__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("vv_varianttree__enumeration").
		SetEnumCode("VARIANTTREE").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Variant Tree generation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vv_varianttree__enumeration: %v", err)
		return err
	}
	cache.Put("vv_varianttree__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_none__enumeration").
		SetEnumCode("NONE").
		SetSequenceID(common.ParseId("01")).
		SetDescription("No Requirement Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_none__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_none__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_auto__enumeration").
		SetEnumCode("AUTO").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Automatic For Every Sales Order").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_auto__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_auto__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_stock__enumeration").
		SetEnumCode("STOCK_QOH").
		SetSequenceID(common.ParseId("03")).
		SetDescription("When QOH Reaches Minimum Stock for Product-Facility").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_stock__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_stock__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_stock_atp__enumeration").
		SetEnumCode("STOCK_ATP").
		SetSequenceID(common.ParseId("04")).
		SetDescription("When ATP Reaches Minimum Stock for Product-Facility").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_stock_atp__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_stock_atp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_atp__enumeration").
		SetEnumCode("ATP").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Requirement for order when ATP Reaches Minimum Stock for Product-Facility").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_atp__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_atp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_ds__enumeration").
		SetEnumCode("DROPS").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Drop-ship only").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_ds__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_ds__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prodrqm_dsatp__enumeration").
		SetEnumCode("DROPS_ATP").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Auto drop-ship on low quantity").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prodrqm_dsatp__enumeration: %v", err)
		return err
	}
	cache.Put("prodrqm_dsatp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("retake_photo__enumeration").
		SetEnumCode("RETAKE_PHOTO").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Re-take Photo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create retake_photo__enumeration: %v", err)
		return err
	}
	cache.Put("retake_photo__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("remove_logo__enumeration").
		SetEnumCode("REMOVE_LOGO").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Remove Logo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create remove_logo__enumeration: %v", err)
		return err
	}
	cache.Put("remove_logo__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("other__enumeration").
		SetEnumCode("OTHER").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Other").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other__enumeration: %v", err)
		return err
	}
	cache.Put("other__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("direct_store_delivry__enumeration").
		SetEnumCode("DIRECT_STORE_DELIVRY").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Direct Store Delivery").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create direct_store_delivry__enumeration: %v", err)
		return err
	}
	cache.Put("direct_store_delivry__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_amount__enumeration").
		SetEnumCode("PPC_PRODUCT_AMOUNT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_amount__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_amount__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_total__enumeration").
		SetEnumCode("PPC_PRODUCT_TOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_total__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_total__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_product_quant__enumeration").
		SetEnumCode("PPC_PRODUCT_QUANT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_product_quant__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_product_quant__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_new_acct__enumeration").
		SetEnumCode("PPC_NEW_ACCT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_new_acct__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_new_acct__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_id__enumeration").
		SetEnumCode("PPC_PARTY_ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_id__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_grp_mem__enumeration").
		SetEnumCode("PPC_PARTY_GRP_MEM").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_grp_mem__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_grp_mem__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_party_class__enumeration").
		SetEnumCode("PPC_PARTY_CLASS").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_party_class__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_party_class__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_role_type__enumeration").
		SetEnumCode("PPC_ROLE_TYPE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_role_type__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_role_type__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_geo_id__enumeration").
		SetEnumCode("PPC_GEO_ID").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_geo_id__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_geo_id__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_order_total__enumeration").
		SetEnumCode("PPC_ORDER_TOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_order_total__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_order_total__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_hist__enumeration").
		SetEnumCode("PPC_ORST_HIST").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_hist__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_hist__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_year__enumeration").
		SetEnumCode("PPC_ORST_YEAR").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_year__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_year__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_orst_last_year__enumeration").
		SetEnumCode("PPC_ORST_LAST_YEAR").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_orst_last_year__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_orst_last_year__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_recurrence__enumeration").
		SetEnumCode("PPC_RECURRENCE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_recurrence__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_recurrence__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_order_shiptotal__enumeration").
		SetEnumCode("PPC_ORDER_SHIPTOTAL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_order_shiptotal__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_order_shiptotal__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_lpmup_amt__enumeration").
		SetEnumCode("PPC_LPMUP_AMT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_lpmup_amt__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_lpmup_amt__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("ppip_lpmup_per__enumeration").
		SetEnumCode("PPC_LPMUP_PER").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppip_lpmup_per__enumeration: %v", err)
		return err
	}
	cache.Put("ppip_lpmup_per__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_gwp__enumeration").
		SetEnumCode("PPA_GWP").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_gwp__enumeration: %v", err)
		return err
	}
	cache.Put("promo_gwp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_free_shipping__enumeration").
		SetEnumCode("PPA_FREE_SHIPPING").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_free_shipping__enumeration: %v", err)
		return err
	}
	cache.Put("promo_free_shipping__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_disc__enumeration").
		SetEnumCode("PPA_PROD_DISC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_disc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_disc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_amdisc__enumeration").
		SetEnumCode("PPA_PROD_AMDISC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_amdisc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_amdisc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_price__enumeration").
		SetEnumCode("PPA_PROD_PRICE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_price__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_price__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_order_percent__enumeration").
		SetEnumCode("PPA_ORDER_PERCENT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_order_percent__enumeration: %v", err)
		return err
	}
	cache.Put("promo_order_percent__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_order_amount__enumeration").
		SetEnumCode("PPA_ORDER_AMOUNT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_order_amount__enumeration: %v", err)
		return err
	}
	cache.Put("promo_order_amount__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_prod_spprc__enumeration").
		SetEnumCode("PPA_PROD_SPPRC").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_prod_spprc__enumeration: %v", err)
		return err
	}
	cache.Put("promo_prod_spprc__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_tax_percent__enumeration").
		SetEnumCode("PPA_TAX_PERCENT").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_tax_percent__enumeration: %v", err)
		return err
	}
	cache.Put("promo_tax_percent__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("promo_ship_charge__enumeration").
		SetEnumCode("PPA_SHIP_CHARGE").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create promo_ship_charge__enumeration: %v", err)
		return err
	}
	cache.Put("promo_ship_charge__enumeration", c)

	return nil
}
