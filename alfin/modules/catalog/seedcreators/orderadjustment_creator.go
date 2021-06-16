package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateOrderAdjustment(ctx context.Context) error {
	log.Println("OrderAdjustment creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.OrderAdjustment

	c, err = client.OrderAdjustment.Create().SetStringRef("81015__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("PROMOTION_ADJUSTMENT")).
		SetOrderItemSeqID(common.ParseId("_NA_")).
		SetShipGroupSeqID(common.ParseId("_NA_")).
		SetDescription("10% off entire purchase").
		SetAmount(-8.840).
		SetProductPromoID(common.ParseId("9011")).
		SetProductPromoRuleID(common.ParseId("01")).
		SetProductPromoActionSeqID(common.ParseId("01")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81015__orderadjustment: %v", err)
		return err
	}
	cache.Put("81015__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81016__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("PROMOTION_ADJUSTMENT")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("_NA_")).
		SetDescription("Buy 4 items for $50 from Purple Gizmo [GZ-5005], Rainbow Gizmo [GZ-1004], Round Gizmo [GZ-2644] or Square Gizmo [GZ-2002] limit 2 per customer").
		SetAmount(-103.600).
		SetProductPromoID(common.ParseId("9015")).
		SetProductPromoRuleID(common.ParseId("01")).
		SetProductPromoActionSeqID(common.ParseId("01")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81016__orderadjustment: %v", err)
		return err
	}
	cache.Put("81016__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81017__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("PROMOTION_ADJUSTMENT")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("_NA_")).
		SetDescription("Get $500 off any item in the Small Gizmos [101] category, limit 1 per order, 2 per customer, 3 for entire promotion. Discount not to exceed the price of the item.").
		SetAmount(-38.400).
		SetProductPromoID(common.ParseId("9016")).
		SetProductPromoRuleID(common.ParseId("01")).
		SetProductPromoActionSeqID(common.ParseId("01")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81017__orderadjustment: %v", err)
		return err
	}
	cache.Put("81017__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81018__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("SHIPPING_CHARGES")).
		SetOrderItemSeqID(common.ParseId("_NA_")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetAmount(24.700).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81018__orderadjustment: %v", err)
		return err
	}
	cache.Put("81018__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81019__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("SALES_TAX")).
		SetOrderItemSeqID(common.ParseId("_NA_")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetComments("Utah County, Utah Sales Tax").
		SetAmount(0.000).
		SetTaxAuthorityRateSeqID(common.ParseId("9005")).
		SetSourcePercentage(0.100000).
		SetCustomerReferenceID("12-3456789").
		SetPrimaryGeoID(common.ParseId("UT-UTAH")).
		SetExemptAmount(0.00).
		SetTaxAuthGeoID(common.ParseId("UT-UTAH")).
		SetTaxAuthPartyID(common.ParseId("UT_UTAH_TAXMAN")).
		SetOverrideGlAccountID(common.ParseId("224153")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81019__orderadjustment: %v", err)
		return err
	}
	cache.Put("81019__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81020__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("SALES_TAX")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetComments("Utah County, Utah Sales Tax").
		SetAmount(0.000).
		SetTaxAuthorityRateSeqID(common.ParseId("9005")).
		SetSourcePercentage(0.100000).
		SetCustomerReferenceID("12-3456789").
		SetPrimaryGeoID(common.ParseId("UT-UTAH")).
		SetExemptAmount(0.08).
		SetTaxAuthGeoID(common.ParseId("UT-UTAH")).
		SetTaxAuthPartyID(common.ParseId("UT_UTAH_TAXMAN")).
		SetOverrideGlAccountID(common.ParseId("224153")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81020__orderadjustment: %v", err)
		return err
	}
	cache.Put("81020__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81021__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("SALES_TAX")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetComments("Utah State Sales Tax").
		SetAmount(0.000).
		SetTaxAuthorityRateSeqID(common.ParseId("9004")).
		SetSourcePercentage(4.750000).
		SetCustomerReferenceID("12-3456789").
		SetPrimaryGeoID(common.ParseId("UT")).
		SetExemptAmount(4.19).
		SetTaxAuthGeoID(common.ParseId("UT")).
		SetTaxAuthPartyID(common.ParseId("UT_TAXMAN")).
		SetOverrideGlAccountID(common.ParseId("224153")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81021__orderadjustment: %v", err)
		return err
	}
	cache.Put("81021__orderadjustment", c)

	c, err = client.OrderAdjustment.Create().SetStringRef("81022__orderadjustment").
		SetOrderAdjustmentTypeID(common.ParseId("SALES_TAX")).
		SetOrderItemSeqID(common.ParseId("00001")).
		SetShipGroupSeqID(common.ParseId("00001")).
		SetComments("1% OFB _NA_ Tax").
		SetAmount(0.884).
		SetTaxAuthorityRateSeqID(common.ParseId("9000")).
		SetSourcePercentage(1.000000).
		SetPrimaryGeoID(common.ParseId("_NA_")).
		SetTaxAuthGeoID(common.ParseId("_NA_")).
		SetTaxAuthPartyID(common.ParseId("_NA_")).
		SetOverrideGlAccountID(common.ParseId("224000")).
		SetCreatedDate(common.ParseDateTime("2011-08-12 23:17:11.708")).
		SetCreatedByUserLogin("admin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 81022__orderadjustment: %v", err)
		return err
	}
	cache.Put("81022__orderadjustment", c)

	return nil
}
