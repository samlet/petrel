package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayUps(ctx context.Context) error {
	log.Println("ShipmentGatewayUps creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayUps

	c, err = client.ShipmentGatewayUps.Create().SetStringRef("ups_config__shipmentgatewayups").
		SetConnectURL("https://wwwcie.ups.com/ups.app/xml").
		SetConnectTimeout(60).
		SetShipperNumber("12345E").
		SetBillShipperAccountNumber("12345E").
		SetAccessLicenseNumber("TEST262223144CAT").
		SetAccessUserID("REG111111").
		SetAccessPassword("REG111111").
		SetSaveCertInfo("true").
		SetSaveCertPath("${sys:getProperty('ofbiz.home')}/runtime/output/upscert").
		SetShipperPickupType("06").
		SetCustomerClassification("03").
		SetMaxEstimateWeight(90).
		SetMinEstimateWeight(0.1).
		SetCodAllowCod("true").
		SetCodSurchargeAmount(9).
		SetCodSurchargeCurrencyUomID("USD").
		SetCodSurchargeApplyToPackage("first").
		SetCodFundsCode("0").
		SetDefaultReturnLabelMemo("UPS Shipment Return Memo").
		SetDefaultReturnLabelSubject("UPS Shipment Return Label").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ups_config__shipmentgatewayups: %v", err)
		return err
	}
	cache.Put("ups_config__shipmentgatewayups", c)

	return nil
}
