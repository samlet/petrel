package helper

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"github.com/samlet/petrel/alfin/modules/catalog/ent/contentassoctype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/contenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/costcomponenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethod"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethodtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/enumerationtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilityassoctype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygrouptype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitytype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/goodidentificationtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/inventoryitemdetail"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/inventoryitemtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderadjustment"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/ordercontactmech"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroupassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgrpinvres"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderrole"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderstatus"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/prodcatalogcategorytype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/prodconfitemcontenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/product"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoctype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategory"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategorycontenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategorytype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productconfigitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcontenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeature"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeatureappltype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeaturecategory"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeatureiactntype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeaturetype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productmainttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productmetertype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productprice"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpriceactiontype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpricepurpose"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpricetype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productreview"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstore"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstoregroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/producttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/quantitybreaktype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/rejectionreason"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/roletype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentcontactmechtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfig"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfigtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewaydhl"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayfedex"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayups"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayusps"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statustype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusvalidchange"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/subscriptiontype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/supplierpreforder"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/variancereason"
)

func ProductCategoryTypeRef(ctx context.Context, stringId string) *ent.ProductCategoryType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductCategoryType.
		Query().
		Where(productcategorytype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductCategoryType %s: %v", stringId, err)
	}
	return rec
}

func ShipmentContactMechTypeRef(ctx context.Context, stringId string) *ent.ShipmentContactMechType {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentContactMechType.
		Query().
		Where(shipmentcontactmechtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentContactMechType %s: %v", stringId, err)
	}
	return rec
}

func ProductMaintTypeRef(ctx context.Context, stringId string) *ent.ProductMaintType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductMaintType.
		Query().
		Where(productmainttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductMaintType %s: %v", stringId, err)
	}
	return rec
}

func ProductRef(ctx context.Context, stringId string) *ent.Product {
	client := ent.FromContext(ctx)
	rec, err := client.Product.
		Query().
		Where(product.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find Product %s: %v", stringId, err)
	}
	return rec
}

func EnumerationRef(ctx context.Context, stringId string) *ent.Enumeration {
	client := ent.FromContext(ctx)
	rec, err := client.Enumeration.
		Query().
		Where(enumeration.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find Enumeration %s: %v", stringId, err)
	}
	return rec
}

func FacilityGroupRef(ctx context.Context, stringId string) *ent.FacilityGroup {
	client := ent.FromContext(ctx)
	rec, err := client.FacilityGroup.
		Query().
		Where(facilitygroup.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find FacilityGroup %s: %v", stringId, err)
	}
	return rec
}

func OrderItemShipGrpInvResRef(ctx context.Context, stringId string) *ent.OrderItemShipGrpInvRes {
	client := ent.FromContext(ctx)
	rec, err := client.OrderItemShipGrpInvRes.
		Query().
		Where(orderitemshipgrpinvres.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderItemShipGrpInvRes %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayFedexRef(ctx context.Context, stringId string) *ent.ShipmentGatewayFedex {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayFedex.
		Query().
		Where(shipmentgatewayfedex.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayFedex %s: %v", stringId, err)
	}
	return rec
}

func ProductCategoryContentTypeRef(ctx context.Context, stringId string) *ent.ProductCategoryContentType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductCategoryContentType.
		Query().
		Where(productcategorycontenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductCategoryContentType %s: %v", stringId, err)
	}
	return rec
}

func OrderAdjustmentRef(ctx context.Context, stringId string) *ent.OrderAdjustment {
	client := ent.FromContext(ctx)
	rec, err := client.OrderAdjustment.
		Query().
		Where(orderadjustment.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderAdjustment %s: %v", stringId, err)
	}
	return rec
}

func SubscriptionTypeRef(ctx context.Context, stringId string) *ent.SubscriptionType {
	client := ent.FromContext(ctx)
	rec, err := client.SubscriptionType.
		Query().
		Where(subscriptiontype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find SubscriptionType %s: %v", stringId, err)
	}
	return rec
}

func InventoryItemDetailRef(ctx context.Context, stringId string) *ent.InventoryItemDetail {
	client := ent.FromContext(ctx)
	rec, err := client.InventoryItemDetail.
		Query().
		Where(inventoryitemdetail.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find InventoryItemDetail %s: %v", stringId, err)
	}
	return rec
}

func ProductCategoryRef(ctx context.Context, stringId string) *ent.ProductCategory {
	client := ent.FromContext(ctx)
	rec, err := client.ProductCategory.
		Query().
		Where(productcategory.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductCategory %s: %v", stringId, err)
	}
	return rec
}

func CostComponentTypeRef(ctx context.Context, stringId string) *ent.CostComponentType {
	client := ent.FromContext(ctx)
	rec, err := client.CostComponentType.
		Query().
		Where(costcomponenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find CostComponentType %s: %v", stringId, err)
	}
	return rec
}

func ProductConfigItemRef(ctx context.Context, stringId string) *ent.ProductConfigItem {
	client := ent.FromContext(ctx)
	rec, err := client.ProductConfigItem.
		Query().
		Where(productconfigitem.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductConfigItem %s: %v", stringId, err)
	}
	return rec
}

func ProductAssocRef(ctx context.Context, stringId string) *ent.ProductAssoc {
	client := ent.FromContext(ctx)
	rec, err := client.ProductAssoc.
		Query().
		Where(productassoc.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductAssoc %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayDhlRef(ctx context.Context, stringId string) *ent.ShipmentGatewayDhl {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayDhl.
		Query().
		Where(shipmentgatewaydhl.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayDhl %s: %v", stringId, err)
	}
	return rec
}

func ProductFeatureCategoryRef(ctx context.Context, stringId string) *ent.ProductFeatureCategory {
	client := ent.FromContext(ctx)
	rec, err := client.ProductFeatureCategory.
		Query().
		Where(productfeaturecategory.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductFeatureCategory %s: %v", stringId, err)
	}
	return rec
}

func ProductFeatureTypeRef(ctx context.Context, stringId string) *ent.ProductFeatureType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductFeatureType.
		Query().
		Where(productfeaturetype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductFeatureType %s: %v", stringId, err)
	}
	return rec
}

func VarianceReasonRef(ctx context.Context, stringId string) *ent.VarianceReason {
	client := ent.FromContext(ctx)
	rec, err := client.VarianceReason.
		Query().
		Where(variancereason.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find VarianceReason %s: %v", stringId, err)
	}
	return rec
}

func RoleTypeRef(ctx context.Context, stringId string) *ent.RoleType {
	client := ent.FromContext(ctx)
	rec, err := client.RoleType.
		Query().
		Where(roletype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find RoleType %s: %v", stringId, err)
	}
	return rec
}

func ProductReviewRef(ctx context.Context, stringId string) *ent.ProductReview {
	client := ent.FromContext(ctx)
	rec, err := client.ProductReview.
		Query().
		Where(productreview.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductReview %s: %v", stringId, err)
	}
	return rec
}

func GoodIdentificationTypeRef(ctx context.Context, stringId string) *ent.GoodIdentificationType {
	client := ent.FromContext(ctx)
	rec, err := client.GoodIdentificationType.
		Query().
		Where(goodidentificationtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find GoodIdentificationType %s: %v", stringId, err)
	}
	return rec
}

func StatusValidChangeRef(ctx context.Context, stringId string) *ent.StatusValidChange {
	client := ent.FromContext(ctx)
	rec, err := client.StatusValidChange.
		Query().
		Where(statusvalidchange.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find StatusValidChange %s: %v", stringId, err)
	}
	return rec
}

func OrderRoleRef(ctx context.Context, stringId string) *ent.OrderRole {
	client := ent.FromContext(ctx)
	rec, err := client.OrderRole.
		Query().
		Where(orderrole.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderRole %s: %v", stringId, err)
	}
	return rec
}

func ProductTypeRef(ctx context.Context, stringId string) *ent.ProductType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductType.
		Query().
		Where(producttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductType %s: %v", stringId, err)
	}
	return rec
}

func ContentAssocTypeRef(ctx context.Context, stringId string) *ent.ContentAssocType {
	client := ent.FromContext(ctx)
	rec, err := client.ContentAssocType.
		Query().
		Where(contentassoctype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ContentAssocType %s: %v", stringId, err)
	}
	return rec
}

func ProductPriceRef(ctx context.Context, stringId string) *ent.ProductPrice {
	client := ent.FromContext(ctx)
	rec, err := client.ProductPrice.
		Query().
		Where(productprice.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductPrice %s: %v", stringId, err)
	}
	return rec
}

func ProdCatalogCategoryTypeRef(ctx context.Context, stringId string) *ent.ProdCatalogCategoryType {
	client := ent.FromContext(ctx)
	rec, err := client.ProdCatalogCategoryType.
		Query().
		Where(prodcatalogcategorytype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProdCatalogCategoryType %s: %v", stringId, err)
	}
	return rec
}

func CustomMethodRef(ctx context.Context, stringId string) *ent.CustomMethod {
	client := ent.FromContext(ctx)
	rec, err := client.CustomMethod.
		Query().
		Where(custommethod.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find CustomMethod %s: %v", stringId, err)
	}
	return rec
}

func StatusTypeRef(ctx context.Context, stringId string) *ent.StatusType {
	client := ent.FromContext(ctx)
	rec, err := client.StatusType.
		Query().
		Where(statustype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find StatusType %s: %v", stringId, err)
	}
	return rec
}

func ProductMeterTypeRef(ctx context.Context, stringId string) *ent.ProductMeterType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductMeterType.
		Query().
		Where(productmetertype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductMeterType %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayUspsRef(ctx context.Context, stringId string) *ent.ShipmentGatewayUsps {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayUsps.
		Query().
		Where(shipmentgatewayusps.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayUsps %s: %v", stringId, err)
	}
	return rec
}

func ProductFeatureIactnTypeRef(ctx context.Context, stringId string) *ent.ProductFeatureIactnType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductFeatureIactnType.
		Query().
		Where(productfeatureiactntype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductFeatureIactnType %s: %v", stringId, err)
	}
	return rec
}

func OrderContactMechRef(ctx context.Context, stringId string) *ent.OrderContactMech {
	client := ent.FromContext(ctx)
	rec, err := client.OrderContactMech.
		Query().
		Where(ordercontactmech.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderContactMech %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayConfigRef(ctx context.Context, stringId string) *ent.ShipmentGatewayConfig {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayConfig.
		Query().
		Where(shipmentgatewayconfig.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayConfig %s: %v", stringId, err)
	}
	return rec
}

func ProductPriceTypeRef(ctx context.Context, stringId string) *ent.ProductPriceType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductPriceType.
		Query().
		Where(productpricetype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductPriceType %s: %v", stringId, err)
	}
	return rec
}

func ProdConfItemContentTypeRef(ctx context.Context, stringId string) *ent.ProdConfItemContentType {
	client := ent.FromContext(ctx)
	rec, err := client.ProdConfItemContentType.
		Query().
		Where(prodconfitemcontenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProdConfItemContentType %s: %v", stringId, err)
	}
	return rec
}

func CustomMethodTypeRef(ctx context.Context, stringId string) *ent.CustomMethodType {
	client := ent.FromContext(ctx)
	rec, err := client.CustomMethodType.
		Query().
		Where(custommethodtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find CustomMethodType %s: %v", stringId, err)
	}
	return rec
}

func FacilityGroupTypeRef(ctx context.Context, stringId string) *ent.FacilityGroupType {
	client := ent.FromContext(ctx)
	rec, err := client.FacilityGroupType.
		Query().
		Where(facilitygrouptype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find FacilityGroupType %s: %v", stringId, err)
	}
	return rec
}

func OrderItemShipGroupAssocRef(ctx context.Context, stringId string) *ent.OrderItemShipGroupAssoc {
	client := ent.FromContext(ctx)
	rec, err := client.OrderItemShipGroupAssoc.
		Query().
		Where(orderitemshipgroupassoc.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderItemShipGroupAssoc %s: %v", stringId, err)
	}
	return rec
}

func StatusItemRef(ctx context.Context, stringId string) *ent.StatusItem {
	client := ent.FromContext(ctx)
	rec, err := client.StatusItem.
		Query().
		Where(statusitem.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find StatusItem %s: %v", stringId, err)
	}
	return rec
}

func ProductFeatureApplTypeRef(ctx context.Context, stringId string) *ent.ProductFeatureApplType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductFeatureApplType.
		Query().
		Where(productfeatureappltype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductFeatureApplType %s: %v", stringId, err)
	}
	return rec
}

func InventoryItemTypeRef(ctx context.Context, stringId string) *ent.InventoryItemType {
	client := ent.FromContext(ctx)
	rec, err := client.InventoryItemType.
		Query().
		Where(inventoryitemtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find InventoryItemType %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayUpsRef(ctx context.Context, stringId string) *ent.ShipmentGatewayUps {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayUps.
		Query().
		Where(shipmentgatewayups.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayUps %s: %v", stringId, err)
	}
	return rec
}

func ProductStoreRef(ctx context.Context, stringId string) *ent.ProductStore {
	client := ent.FromContext(ctx)
	rec, err := client.ProductStore.
		Query().
		Where(productstore.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductStore %s: %v", stringId, err)
	}
	return rec
}

func ProductPriceActionTypeRef(ctx context.Context, stringId string) *ent.ProductPriceActionType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductPriceActionType.
		Query().
		Where(productpriceactiontype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductPriceActionType %s: %v", stringId, err)
	}
	return rec
}

func ShipmentTypeRef(ctx context.Context, stringId string) *ent.ShipmentType {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentType.
		Query().
		Where(shipmenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentType %s: %v", stringId, err)
	}
	return rec
}

func ProductAssocTypeRef(ctx context.Context, stringId string) *ent.ProductAssocType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductAssocType.
		Query().
		Where(productassoctype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductAssocType %s: %v", stringId, err)
	}
	return rec
}

func OrderItemShipGroupRef(ctx context.Context, stringId string) *ent.OrderItemShipGroup {
	client := ent.FromContext(ctx)
	rec, err := client.OrderItemShipGroup.
		Query().
		Where(orderitemshipgroup.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderItemShipGroup %s: %v", stringId, err)
	}
	return rec
}

func FacilityTypeRef(ctx context.Context, stringId string) *ent.FacilityType {
	client := ent.FromContext(ctx)
	rec, err := client.FacilityType.
		Query().
		Where(facilitytype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find FacilityType %s: %v", stringId, err)
	}
	return rec
}

func ProductPricePurposeRef(ctx context.Context, stringId string) *ent.ProductPricePurpose {
	client := ent.FromContext(ctx)
	rec, err := client.ProductPricePurpose.
		Query().
		Where(productpricepurpose.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductPricePurpose %s: %v", stringId, err)
	}
	return rec
}

func SupplierPrefOrderRef(ctx context.Context, stringId string) *ent.SupplierPrefOrder {
	client := ent.FromContext(ctx)
	rec, err := client.SupplierPrefOrder.
		Query().
		Where(supplierpreforder.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find SupplierPrefOrder %s: %v", stringId, err)
	}
	return rec
}

func ProductStoreGroupRef(ctx context.Context, stringId string) *ent.ProductStoreGroup {
	client := ent.FromContext(ctx)
	rec, err := client.ProductStoreGroup.
		Query().
		Where(productstoregroup.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductStoreGroup %s: %v", stringId, err)
	}
	return rec
}

func ContentTypeRef(ctx context.Context, stringId string) *ent.ContentType {
	client := ent.FromContext(ctx)
	rec, err := client.ContentType.
		Query().
		Where(contenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ContentType %s: %v", stringId, err)
	}
	return rec
}

func ShipmentGatewayConfigTypeRef(ctx context.Context, stringId string) *ent.ShipmentGatewayConfigType {
	client := ent.FromContext(ctx)
	rec, err := client.ShipmentGatewayConfigType.
		Query().
		Where(shipmentgatewayconfigtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ShipmentGatewayConfigType %s: %v", stringId, err)
	}
	return rec
}

func ProductContentTypeRef(ctx context.Context, stringId string) *ent.ProductContentType {
	client := ent.FromContext(ctx)
	rec, err := client.ProductContentType.
		Query().
		Where(productcontenttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductContentType %s: %v", stringId, err)
	}
	return rec
}

func OrderStatusRef(ctx context.Context, stringId string) *ent.OrderStatus {
	client := ent.FromContext(ctx)
	rec, err := client.OrderStatus.
		Query().
		Where(orderstatus.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderStatus %s: %v", stringId, err)
	}
	return rec
}

func RejectionReasonRef(ctx context.Context, stringId string) *ent.RejectionReason {
	client := ent.FromContext(ctx)
	rec, err := client.RejectionReason.
		Query().
		Where(rejectionreason.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find RejectionReason %s: %v", stringId, err)
	}
	return rec
}

func OrderHeaderRef(ctx context.Context, stringId string) *ent.OrderHeader {
	client := ent.FromContext(ctx)
	rec, err := client.OrderHeader.
		Query().
		Where(orderheader.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderHeader %s: %v", stringId, err)
	}
	return rec
}

func OrderItemRef(ctx context.Context, stringId string) *ent.OrderItem {
	client := ent.FromContext(ctx)
	rec, err := client.OrderItem.
		Query().
		Where(orderitem.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderItem %s: %v", stringId, err)
	}
	return rec
}

func QuantityBreakTypeRef(ctx context.Context, stringId string) *ent.QuantityBreakType {
	client := ent.FromContext(ctx)
	rec, err := client.QuantityBreakType.
		Query().
		Where(quantitybreaktype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find QuantityBreakType %s: %v", stringId, err)
	}
	return rec
}

func OrderPaymentPreferenceRef(ctx context.Context, stringId string) *ent.OrderPaymentPreference {
	client := ent.FromContext(ctx)
	rec, err := client.OrderPaymentPreference.
		Query().
		Where(orderpaymentpreference.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find OrderPaymentPreference %s: %v", stringId, err)
	}
	return rec
}

func ProductFeatureRef(ctx context.Context, stringId string) *ent.ProductFeature {
	client := ent.FromContext(ctx)
	rec, err := client.ProductFeature.
		Query().
		Where(productfeature.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find ProductFeature %s: %v", stringId, err)
	}
	return rec
}

func EnumerationTypeRef(ctx context.Context, stringId string) *ent.EnumerationType {
	client := ent.FromContext(ctx)
	rec, err := client.EnumerationType.
		Query().
		Where(enumerationtype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find EnumerationType %s: %v", stringId, err)
	}
	return rec
}

func FacilityAssocTypeRef(ctx context.Context, stringId string) *ent.FacilityAssocType {
	client := ent.FromContext(ctx)
	rec, err := client.FacilityAssocType.
		Query().
		Where(facilityassoctype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find FacilityAssocType %s: %v", stringId, err)
	}
	return rec
}
