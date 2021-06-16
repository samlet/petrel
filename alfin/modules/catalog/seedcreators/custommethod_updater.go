package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateCustomMethod(ctx context.Context) error {
	log.Println("CustomMethod updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CustomMethod
	failures := 0

	c = cache.Get("prod_perc_formula__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productCostPercentageFormula").
		SetDescription("Formula that creates a cost component equal to a percentage of total product cost").
		SetCustomMethodType(cache.Get("cost_formula__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prod_perc_formula__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prod_perc_formula__custommethod", c)
	}

	c = cache.Get("ship_est_dhl__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("dhlRateEstimate").
		SetDescription("DHL rate estimate").
		SetCustomMethodType(cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est_dhl__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est_dhl__custommethod", c)
	}

	c = cache.Get("ship_est_fedex__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("fedexRateEstimate").
		SetDescription("FedEx rate estimate").
		SetCustomMethodType(cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est_fedex__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est_fedex__custommethod", c)
	}

	c = cache.Get("ship_est_ups__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("upsRateEstimate").
		SetDescription("UPS rate estimate").
		SetCustomMethodType(cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est_ups__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est_ups__custommethod", c)
	}

	c = cache.Get("ship_est_usps__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("uspsRateInquire").
		SetDescription("USPS rate estimate").
		SetCustomMethodType(cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est_usps__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est_usps__custommethod", c)
	}

	c = cache.Get("ship_est_usps_int__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("uspsInternationalRateInquire").
		SetDescription("USPS rate estimate international").
		SetCustomMethodType(cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est_usps_int__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est_usps_int__custommethod", c)
	}

	c = cache.Get("ppc_product_amount__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondProductAmount").
		SetDescription("Product amount").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_product_amount__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_product_amount__custommethod", c)
	}

	c = cache.Get("ppc_product_total__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondProductTotal").
		SetDescription("Product Total").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_product_total__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_product_total__custommethod", c)
	}

	c = cache.Get("ppc_product_quant__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondProductQuant").
		SetDescription("X Quantity of Product").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_product_quant__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_product_quant__custommethod", c)
	}

	c = cache.Get("ppc_new_acct__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondNewACCT").
		SetDescription("Account Days Since Created").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_new_acct__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_new_acct__custommethod", c)
	}

	c = cache.Get("ppc_party_id__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondPartyID").
		SetDescription("Party").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_party_id__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_party_id__custommethod", c)
	}

	c = cache.Get("ppc_party_grp_mem__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondPartyGM").
		SetDescription("Party group member").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_party_grp_mem__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_party_grp_mem__custommethod", c)
	}

	c = cache.Get("ppc_party_class__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondPartyClass").
		SetDescription("Party classification").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_party_class__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_party_class__custommethod", c)
	}

	c = cache.Get("ppc_role_type__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondRoleType").
		SetDescription("Role type").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_role_type__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_role_type__custommethod", c)
	}

	c = cache.Get("ppc_geo_id__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondGeoID").
		SetDescription("Shipping destination").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_geo_id__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_geo_id__custommethod", c)
	}

	c = cache.Get("ppc_order_total__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondOrderTotal").
		SetDescription("Cart sub-total").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_order_total__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_order_total__custommethod", c)
	}

	c = cache.Get("ppc_orst_hist__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondOrderHist").
		SetDescription("Order sub-total X in last Y Months").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_orst_hist__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_orst_hist__custommethod", c)
	}

	c = cache.Get("ppc_orst_year__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondOrderYear").
		SetDescription("Order sub-total X since beginning of current year").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_orst_year__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_orst_year__custommethod", c)
	}

	c = cache.Get("ppc_orst_last_year__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondOrderLastYear").
		SetDescription("Order sub-total X last year").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_orst_last_year__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_orst_last_year__custommethod", c)
	}

	c = cache.Get("ppc_recurrence__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondPromoRecurrence").
		SetDescription("Promotion Recurrence").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_recurrence__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_recurrence__custommethod", c)
	}

	c = cache.Get("ppc_order_shiptotal__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondOrderShipTotal").
		SetDescription("Shipping total").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_order_shiptotal__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_order_shiptotal__custommethod", c)
	}

	c = cache.Get("ppc_lpmup_amt__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondListPriceMinAmount").
		SetDescription("List Price minus Unit Price (Amount)").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_lpmup_amt__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_lpmup_amt__custommethod", c)
	}

	c = cache.Get("ppc_lpmup_per__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoCondListPriceMinPercent").
		SetDescription("List Price minus Unit Price (Percent)").
		SetCustomMethodType(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppc_lpmup_per__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppc_lpmup_per__custommethod", c)
	}

	c = cache.Get("ppa_gwp__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActGiftGWP").
		SetDescription("gift with purchase").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_gwp__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_gwp__custommethod", c)
	}

	c = cache.Get("ppa_free_shipping__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActFreeShip").
		SetDescription("free shipping").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_free_shipping__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_free_shipping__custommethod", c)
	}

	c = cache.Get("ppa_prod_disc__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActProdDISC").
		SetDescription("X Product for Y% Discount").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_prod_disc__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_prod_disc__custommethod", c)
	}

	c = cache.Get("ppa_prod_amdisc__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActProdAMDISC").
		SetDescription("X Product for Y Discount").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_prod_amdisc__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_prod_amdisc__custommethod", c)
	}

	c = cache.Get("ppa_prod_price__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActProdPrice").
		SetDescription("X Product for Y Price").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_prod_price__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_prod_price__custommethod", c)
	}

	c = cache.Get("ppa_order_percent__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActOrderPercent").
		SetDescription("Order Percent Discount").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_order_percent__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_order_percent__custommethod", c)
	}

	c = cache.Get("ppa_order_amount__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActOrderAmount").
		SetDescription("Order Amount Flat").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_order_amount__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_order_amount__custommethod", c)
	}

	c = cache.Get("ppa_prod_spprc__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActProdSpecialPrice").
		SetDescription("Product for [Special Promo] Price").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_prod_spprc__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_prod_spprc__custommethod", c)
	}

	c = cache.Get("ppa_tax_percent__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActTaxPercent").
		SetDescription("Tax % Discount").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_tax_percent__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_tax_percent__custommethod", c)
	}

	c = cache.Get("ppa_ship_charge__custommethod").(*ent.CustomMethod)
	c, err = c.Update().
		SetCustomMethodName("productPromoActShipCharge").
		SetDescription("shipping charge").
		SetCustomMethodType(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ppa_ship_charge__custommethod: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ppa_ship_charge__custommethod", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
