package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateCustomMethodType(ctx context.Context) error {
	log.Println("CustomMethodType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CustomMethodType
	failures := 0

	c = cache.Get("cost_formula__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetDescription("Formula for calculating costs for tasks and products").
		AddCustomMethods(cache.Get("prod_perc_formula__custommethod").(*ent.CustomMethod)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cost_formula__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cost_formula__custommethodtype", c)
	}

	c = cache.Get("price_formula__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetDescription("Service with formula for calculating the unit price of a product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update price_formula__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("price_formula__custommethodtype", c)
	}

	c = cache.Get("ship_est__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipment Gateway rate estimate methods").
		AddCustomMethods(cache.Get("ship_est_dhl__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ship_est_fedex__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ship_est_ups__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ship_est_usps__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ship_est_usps_int__custommethod").(*ent.CustomMethod)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_est__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_est__custommethodtype", c)
	}

	c = cache.Get("product_promo__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetDescription("").
		AddChildren(cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)).
		AddChildren(cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_promo__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_promo__custommethodtype", c)
	}

	c = cache.Get("product_promo_cond__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetDescription("").
		SetParent(cache.Get("product_promo__custommethodtype").(*ent.CustomMethodType)).
		AddCustomMethods(cache.Get("ppc_product_amount__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_product_total__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_product_quant__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_new_acct__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_party_id__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_party_grp_mem__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_party_class__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_role_type__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_geo_id__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_order_total__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_orst_hist__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_orst_year__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_orst_last_year__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_recurrence__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_order_shiptotal__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_lpmup_amt__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppc_lpmup_per__custommethod").(*ent.CustomMethod)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_promo_cond__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_promo_cond__custommethodtype", c)
	}

	c = cache.Get("product_promo_action__custommethodtype").(*ent.CustomMethodType)
	c, err = c.Update().
		SetDescription("").
		SetParent(cache.Get("product_promo__custommethodtype").(*ent.CustomMethodType)).
		AddCustomMethods(cache.Get("ppa_gwp__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_free_shipping__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_prod_disc__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_prod_amdisc__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_prod_price__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_order_percent__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_order_amount__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_prod_spprc__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_tax_percent__custommethod").(*ent.CustomMethod)).
		AddCustomMethods(cache.Get("ppa_ship_charge__custommethod").(*ent.CustomMethod)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update product_promo_action__custommethodtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("product_promo_action__custommethodtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
