package seedcreators

import (
	"context"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func LoadSeeds(execUpdaters bool) {
	client, err := ent.Open("sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %!v(MISSING)", err)
	}
	defer client.Close()

	ctx := context.Background()
	ctx = ent.NewContext(ctx, client)
	ctx = cachecomp.NewContext(ctx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %!v(MISSING)", err)
	}

	if err := CreateQuantityBreakType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumerationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateInventoryItemType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductPricePurpose(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductPriceActionType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayDhl(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderRole(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductStore(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFacilityGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSubscriptionType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProduct(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayFedex(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProdConfItemContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductConfigItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductFeature(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductCategory(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayConfigType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderItemShipGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayUps(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateEnumeration(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductPrice(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductMeterType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderItemShipGrpInvRes(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProdCatalogCategoryType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderStatus(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderPaymentPreference(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderHeader(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCustomMethodType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductStoreGroup(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderAdjustment(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductFeatureIactnType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductReview(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductPriceType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContentAssocType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFacilityAssocType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductAssocType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentContactMechType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductCategoryType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductFeatureType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateInventoryItemDetail(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateGoodIdentificationType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateVarianceReason(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayConfig(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateRejectionReason(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateRoleType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCustomMethod(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateShipmentGatewayUsps(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductFeatureCategory(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderItem(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductMaintType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateCostComponentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateStatusValidChange(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFacilityGroupType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderContactMech(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateFacilityType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateOrderItemShipGroupAssoc(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateSupplierPrefOrder(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductFeatureApplType(ctx); err != nil {
		log.Fatal(err)
	}
	if err := CreateProductCategoryContentType(ctx); err != nil {
		log.Fatal(err)
	}
	if execUpdaters {
		prompt := color.New(color.FgGreen).PrintfFunc()
		warn := color.New(color.FgRed).PrintfFunc()
		if err := UpdateShipmentType(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentType fail: %v\n", err)
		} else {
			prompt("Update ShipmentType ok.\n")
		}
		if err := UpdateCustomMethod(ctx); err != nil {
			log.Println(err)
			warn("Update CustomMethod fail: %v\n", err)
		} else {
			prompt("Update CustomMethod ok.\n")
		}
		if err := UpdateShipmentGatewayUsps(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayUsps fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayUsps ok.\n")
		}
		if err := UpdateRejectionReason(ctx); err != nil {
			log.Println(err)
			warn("Update RejectionReason fail: %v\n", err)
		} else {
			prompt("Update RejectionReason ok.\n")
		}
		if err := UpdateStatusType(ctx); err != nil {
			log.Println(err)
			warn("Update StatusType fail: %v\n", err)
		} else {
			prompt("Update StatusType ok.\n")
		}
		if err := UpdateRoleType(ctx); err != nil {
			log.Println(err)
			warn("Update RoleType fail: %v\n", err)
		} else {
			prompt("Update RoleType ok.\n")
		}
		if err := UpdateCostComponentType(ctx); err != nil {
			log.Println(err)
			warn("Update CostComponentType fail: %v\n", err)
		} else {
			prompt("Update CostComponentType ok.\n")
		}
		if err := UpdateStatusValidChange(ctx); err != nil {
			log.Println(err)
			warn("Update StatusValidChange fail: %v\n", err)
		} else {
			prompt("Update StatusValidChange ok.\n")
		}
		if err := UpdateFacilityGroupType(ctx); err != nil {
			log.Println(err)
			warn("Update FacilityGroupType fail: %v\n", err)
		} else {
			prompt("Update FacilityGroupType ok.\n")
		}
		if err := UpdateProductFeatureCategory(ctx); err != nil {
			log.Println(err)
			warn("Update ProductFeatureCategory fail: %v\n", err)
		} else {
			prompt("Update ProductFeatureCategory ok.\n")
		}
		if err := UpdateOrderItem(ctx); err != nil {
			log.Println(err)
			warn("Update OrderItem fail: %v\n", err)
		} else {
			prompt("Update OrderItem ok.\n")
		}
		if err := UpdateProductMaintType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductMaintType fail: %v\n", err)
		} else {
			prompt("Update ProductMaintType ok.\n")
		}
		if err := UpdateContentType(ctx); err != nil {
			log.Println(err)
			warn("Update ContentType fail: %v\n", err)
		} else {
			prompt("Update ContentType ok.\n")
		}
		if err := UpdateSupplierPrefOrder(ctx); err != nil {
			log.Println(err)
			warn("Update SupplierPrefOrder fail: %v\n", err)
		} else {
			prompt("Update SupplierPrefOrder ok.\n")
		}
		if err := UpdateOrderContactMech(ctx); err != nil {
			log.Println(err)
			warn("Update OrderContactMech fail: %v\n", err)
		} else {
			prompt("Update OrderContactMech ok.\n")
		}
		if err := UpdateFacilityType(ctx); err != nil {
			log.Println(err)
			warn("Update FacilityType fail: %v\n", err)
		} else {
			prompt("Update FacilityType ok.\n")
		}
		if err := UpdateOrderItemShipGroupAssoc(ctx); err != nil {
			log.Println(err)
			warn("Update OrderItemShipGroupAssoc fail: %v\n", err)
		} else {
			prompt("Update OrderItemShipGroupAssoc ok.\n")
		}
		if err := UpdateProductFeatureApplType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductFeatureApplType fail: %v\n", err)
		} else {
			prompt("Update ProductFeatureApplType ok.\n")
		}
		if err := UpdateProductCategoryContentType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductCategoryContentType fail: %v\n", err)
		} else {
			prompt("Update ProductCategoryContentType ok.\n")
		}
		if err := UpdateQuantityBreakType(ctx); err != nil {
			log.Println(err)
			warn("Update QuantityBreakType fail: %v\n", err)
		} else {
			prompt("Update QuantityBreakType ok.\n")
		}
		if err := UpdateEnumerationType(ctx); err != nil {
			log.Println(err)
			warn("Update EnumerationType fail: %v\n", err)
		} else {
			prompt("Update EnumerationType ok.\n")
		}
		if err := UpdateInventoryItemType(ctx); err != nil {
			log.Println(err)
			warn("Update InventoryItemType fail: %v\n", err)
		} else {
			prompt("Update InventoryItemType ok.\n")
		}
		if err := UpdateOrderRole(ctx); err != nil {
			log.Println(err)
			warn("Update OrderRole fail: %v\n", err)
		} else {
			prompt("Update OrderRole ok.\n")
		}
		if err := UpdateProductPricePurpose(ctx); err != nil {
			log.Println(err)
			warn("Update ProductPricePurpose fail: %v\n", err)
		} else {
			prompt("Update ProductPricePurpose ok.\n")
		}
		if err := UpdateProductPriceActionType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductPriceActionType fail: %v\n", err)
		} else {
			prompt("Update ProductPriceActionType ok.\n")
		}
		if err := UpdateShipmentGatewayDhl(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayDhl fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayDhl ok.\n")
		}
		if err := UpdateProductStore(ctx); err != nil {
			log.Println(err)
			warn("Update ProductStore fail: %v\n", err)
		} else {
			prompt("Update ProductStore ok.\n")
		}
		if err := UpdateSubscriptionType(ctx); err != nil {
			log.Println(err)
			warn("Update SubscriptionType fail: %v\n", err)
		} else {
			prompt("Update SubscriptionType ok.\n")
		}
		if err := UpdateProduct(ctx); err != nil {
			log.Println(err)
			warn("Update Product fail: %v\n", err)
		} else {
			prompt("Update Product ok.\n")
		}
		if err := UpdateShipmentGatewayFedex(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayFedex fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayFedex ok.\n")
		}
		if err := UpdateProductAssoc(ctx); err != nil {
			log.Println(err)
			warn("Update ProductAssoc fail: %v\n", err)
		} else {
			prompt("Update ProductAssoc ok.\n")
		}
		if err := UpdateProductType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductType fail: %v\n", err)
		} else {
			prompt("Update ProductType ok.\n")
		}
		if err := UpdateFacilityGroup(ctx); err != nil {
			log.Println(err)
			warn("Update FacilityGroup fail: %v\n", err)
		} else {
			prompt("Update FacilityGroup ok.\n")
		}
		if err := UpdateShipmentGatewayConfigType(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayConfigType fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayConfigType ok.\n")
		}
		if err := UpdateOrderItemShipGroup(ctx); err != nil {
			log.Println(err)
			warn("Update OrderItemShipGroup fail: %v\n", err)
		} else {
			prompt("Update OrderItemShipGroup ok.\n")
		}
		if err := UpdateShipmentGatewayUps(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayUps fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayUps ok.\n")
		}
		if err := UpdateProdConfItemContentType(ctx); err != nil {
			log.Println(err)
			warn("Update ProdConfItemContentType fail: %v\n", err)
		} else {
			prompt("Update ProdConfItemContentType ok.\n")
		}
		if err := UpdateProductConfigItem(ctx); err != nil {
			log.Println(err)
			warn("Update ProductConfigItem fail: %v\n", err)
		} else {
			prompt("Update ProductConfigItem ok.\n")
		}
		if err := UpdateProductFeature(ctx); err != nil {
			log.Println(err)
			warn("Update ProductFeature fail: %v\n", err)
		} else {
			prompt("Update ProductFeature ok.\n")
		}
		if err := UpdateProductCategory(ctx); err != nil {
			log.Println(err)
			warn("Update ProductCategory fail: %v\n", err)
		} else {
			prompt("Update ProductCategory ok.\n")
		}
		if err := UpdateOrderItemShipGrpInvRes(ctx); err != nil {
			log.Println(err)
			warn("Update OrderItemShipGrpInvRes fail: %v\n", err)
		} else {
			prompt("Update OrderItemShipGrpInvRes ok.\n")
		}
		if err := UpdateProdCatalogCategoryType(ctx); err != nil {
			log.Println(err)
			warn("Update ProdCatalogCategoryType fail: %v\n", err)
		} else {
			prompt("Update ProdCatalogCategoryType ok.\n")
		}
		if err := UpdateOrderStatus(ctx); err != nil {
			log.Println(err)
			warn("Update OrderStatus fail: %v\n", err)
		} else {
			prompt("Update OrderStatus ok.\n")
		}
		if err := UpdateEnumeration(ctx); err != nil {
			log.Println(err)
			warn("Update Enumeration fail: %v\n", err)
		} else {
			prompt("Update Enumeration ok.\n")
		}
		if err := UpdateProductPrice(ctx); err != nil {
			log.Println(err)
			warn("Update ProductPrice fail: %v\n", err)
		} else {
			prompt("Update ProductPrice ok.\n")
		}
		if err := UpdateProductMeterType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductMeterType fail: %v\n", err)
		} else {
			prompt("Update ProductMeterType ok.\n")
		}
		if err := UpdateStatusItem(ctx); err != nil {
			log.Println(err)
			warn("Update StatusItem fail: %v\n", err)
		} else {
			prompt("Update StatusItem ok.\n")
		}
		if err := UpdateOrderPaymentPreference(ctx); err != nil {
			log.Println(err)
			warn("Update OrderPaymentPreference fail: %v\n", err)
		} else {
			prompt("Update OrderPaymentPreference ok.\n")
		}
		if err := UpdateOrderHeader(ctx); err != nil {
			log.Println(err)
			warn("Update OrderHeader fail: %v\n", err)
		} else {
			prompt("Update OrderHeader ok.\n")
		}
		if err := UpdateCustomMethodType(ctx); err != nil {
			log.Println(err)
			warn("Update CustomMethodType fail: %v\n", err)
		} else {
			prompt("Update CustomMethodType ok.\n")
		}
		if err := UpdateProductReview(ctx); err != nil {
			log.Println(err)
			warn("Update ProductReview fail: %v\n", err)
		} else {
			prompt("Update ProductReview ok.\n")
		}
		if err := UpdateProductPriceType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductPriceType fail: %v\n", err)
		} else {
			prompt("Update ProductPriceType ok.\n")
		}
		if err := UpdateContentAssocType(ctx); err != nil {
			log.Println(err)
			warn("Update ContentAssocType fail: %v\n", err)
		} else {
			prompt("Update ContentAssocType ok.\n")
		}
		if err := UpdateProductStoreGroup(ctx); err != nil {
			log.Println(err)
			warn("Update ProductStoreGroup fail: %v\n", err)
		} else {
			prompt("Update ProductStoreGroup ok.\n")
		}
		if err := UpdateOrderAdjustment(ctx); err != nil {
			log.Println(err)
			warn("Update OrderAdjustment fail: %v\n", err)
		} else {
			prompt("Update OrderAdjustment ok.\n")
		}
		if err := UpdateProductFeatureIactnType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductFeatureIactnType fail: %v\n", err)
		} else {
			prompt("Update ProductFeatureIactnType ok.\n")
		}
		if err := UpdateProductCategoryType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductCategoryType fail: %v\n", err)
		} else {
			prompt("Update ProductCategoryType ok.\n")
		}
		if err := UpdateProductContentType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductContentType fail: %v\n", err)
		} else {
			prompt("Update ProductContentType ok.\n")
		}
		if err := UpdateProductFeatureType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductFeatureType fail: %v\n", err)
		} else {
			prompt("Update ProductFeatureType ok.\n")
		}
		if err := UpdateFacilityAssocType(ctx); err != nil {
			log.Println(err)
			warn("Update FacilityAssocType fail: %v\n", err)
		} else {
			prompt("Update FacilityAssocType ok.\n")
		}
		if err := UpdateProductAssocType(ctx); err != nil {
			log.Println(err)
			warn("Update ProductAssocType fail: %v\n", err)
		} else {
			prompt("Update ProductAssocType ok.\n")
		}
		if err := UpdateShipmentContactMechType(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentContactMechType fail: %v\n", err)
		} else {
			prompt("Update ShipmentContactMechType ok.\n")
		}
		if err := UpdateShipmentGatewayConfig(ctx); err != nil {
			log.Println(err)
			warn("Update ShipmentGatewayConfig fail: %v\n", err)
		} else {
			prompt("Update ShipmentGatewayConfig ok.\n")
		}
		if err := UpdateInventoryItemDetail(ctx); err != nil {
			log.Println(err)
			warn("Update InventoryItemDetail fail: %v\n", err)
		} else {
			prompt("Update InventoryItemDetail ok.\n")
		}
		if err := UpdateGoodIdentificationType(ctx); err != nil {
			log.Println(err)
			warn("Update GoodIdentificationType fail: %v\n", err)
		} else {
			prompt("Update GoodIdentificationType ok.\n")
		}
		if err := UpdateVarianceReason(ctx); err != nil {
			log.Println(err)
			warn("Update VarianceReason fail: %v\n", err)
		} else {
			prompt("Update VarianceReason ok.\n")
		}
	}
}
