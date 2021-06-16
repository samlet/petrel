package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateCustomMethod(ctx context.Context) error {
	log.Println("CustomMethod creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CustomMethod

	c, err = client.CustomMethod.Create().SetStringRef("prod_perc_formula__custommethod").
		SetCustomMethodName("productCostPercentageFormula").
		SetDescription("Formula that creates a cost component equal to a percentage of total product cost").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prod_perc_formula__custommethod: %v", err)
		return err
	}
	cache.Put("prod_perc_formula__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ship_est_dhl__custommethod").
		SetCustomMethodName("dhlRateEstimate").
		SetDescription("DHL rate estimate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est_dhl__custommethod: %v", err)
		return err
	}
	cache.Put("ship_est_dhl__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ship_est_fedex__custommethod").
		SetCustomMethodName("fedexRateEstimate").
		SetDescription("FedEx rate estimate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est_fedex__custommethod: %v", err)
		return err
	}
	cache.Put("ship_est_fedex__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ship_est_ups__custommethod").
		SetCustomMethodName("upsRateEstimate").
		SetDescription("UPS rate estimate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est_ups__custommethod: %v", err)
		return err
	}
	cache.Put("ship_est_ups__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ship_est_usps__custommethod").
		SetCustomMethodName("uspsRateInquire").
		SetDescription("USPS rate estimate").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est_usps__custommethod: %v", err)
		return err
	}
	cache.Put("ship_est_usps__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ship_est_usps_int__custommethod").
		SetCustomMethodName("uspsInternationalRateInquire").
		SetDescription("USPS rate estimate international").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_est_usps_int__custommethod: %v", err)
		return err
	}
	cache.Put("ship_est_usps_int__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_product_amount__custommethod").
		SetCustomMethodName("productPromoCondProductAmount").
		SetDescription("Product amount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_product_amount__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_product_amount__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_product_total__custommethod").
		SetCustomMethodName("productPromoCondProductTotal").
		SetDescription("Product Total").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_product_total__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_product_total__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_product_quant__custommethod").
		SetCustomMethodName("productPromoCondProductQuant").
		SetDescription("X Quantity of Product").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_product_quant__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_product_quant__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_new_acct__custommethod").
		SetCustomMethodName("productPromoCondNewACCT").
		SetDescription("Account Days Since Created").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_new_acct__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_new_acct__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_party_id__custommethod").
		SetCustomMethodName("productPromoCondPartyID").
		SetDescription("Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_party_id__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_party_id__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_party_grp_mem__custommethod").
		SetCustomMethodName("productPromoCondPartyGM").
		SetDescription("Party group member").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_party_grp_mem__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_party_grp_mem__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_party_class__custommethod").
		SetCustomMethodName("productPromoCondPartyClass").
		SetDescription("Party classification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_party_class__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_party_class__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_role_type__custommethod").
		SetCustomMethodName("productPromoCondRoleType").
		SetDescription("Role type").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_role_type__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_role_type__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_geo_id__custommethod").
		SetCustomMethodName("productPromoCondGeoID").
		SetDescription("Shipping destination").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_geo_id__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_geo_id__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_order_total__custommethod").
		SetCustomMethodName("productPromoCondOrderTotal").
		SetDescription("Cart sub-total").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_order_total__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_order_total__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_orst_hist__custommethod").
		SetCustomMethodName("productPromoCondOrderHist").
		SetDescription("Order sub-total X in last Y Months").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_orst_hist__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_orst_hist__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_orst_year__custommethod").
		SetCustomMethodName("productPromoCondOrderYear").
		SetDescription("Order sub-total X since beginning of current year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_orst_year__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_orst_year__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_orst_last_year__custommethod").
		SetCustomMethodName("productPromoCondOrderLastYear").
		SetDescription("Order sub-total X last year").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_orst_last_year__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_orst_last_year__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_recurrence__custommethod").
		SetCustomMethodName("productPromoCondPromoRecurrence").
		SetDescription("Promotion Recurrence").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_recurrence__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_recurrence__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_order_shiptotal__custommethod").
		SetCustomMethodName("productPromoCondOrderShipTotal").
		SetDescription("Shipping total").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_order_shiptotal__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_order_shiptotal__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_lpmup_amt__custommethod").
		SetCustomMethodName("productPromoCondListPriceMinAmount").
		SetDescription("List Price minus Unit Price (Amount)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_lpmup_amt__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_lpmup_amt__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppc_lpmup_per__custommethod").
		SetCustomMethodName("productPromoCondListPriceMinPercent").
		SetDescription("List Price minus Unit Price (Percent)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppc_lpmup_per__custommethod: %v", err)
		return err
	}
	cache.Put("ppc_lpmup_per__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_gwp__custommethod").
		SetCustomMethodName("productPromoActGiftGWP").
		SetDescription("gift with purchase").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_gwp__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_gwp__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_free_shipping__custommethod").
		SetCustomMethodName("productPromoActFreeShip").
		SetDescription("free shipping").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_free_shipping__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_free_shipping__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_prod_disc__custommethod").
		SetCustomMethodName("productPromoActProdDISC").
		SetDescription("X Product for Y% Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_prod_disc__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_prod_disc__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_prod_amdisc__custommethod").
		SetCustomMethodName("productPromoActProdAMDISC").
		SetDescription("X Product for Y Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_prod_amdisc__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_prod_amdisc__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_prod_price__custommethod").
		SetCustomMethodName("productPromoActProdPrice").
		SetDescription("X Product for Y Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_prod_price__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_prod_price__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_order_percent__custommethod").
		SetCustomMethodName("productPromoActOrderPercent").
		SetDescription("Order Percent Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_order_percent__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_order_percent__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_order_amount__custommethod").
		SetCustomMethodName("productPromoActOrderAmount").
		SetDescription("Order Amount Flat").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_order_amount__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_order_amount__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_prod_spprc__custommethod").
		SetCustomMethodName("productPromoActProdSpecialPrice").
		SetDescription("Product for [Special Promo] Price").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_prod_spprc__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_prod_spprc__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_tax_percent__custommethod").
		SetCustomMethodName("productPromoActTaxPercent").
		SetDescription("Tax % Discount").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_tax_percent__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_tax_percent__custommethod", c)

	c, err = client.CustomMethod.Create().SetStringRef("ppa_ship_charge__custommethod").
		SetCustomMethodName("productPromoActShipCharge").
		SetDescription("shipping charge").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ppa_ship_charge__custommethod: %v", err)
		return err
	}
	cache.Put("ppa_ship_charge__custommethod", c)

	return nil
}
