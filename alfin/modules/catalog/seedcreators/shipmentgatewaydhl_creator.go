package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayDhl(ctx context.Context) error {
	log.Println("ShipmentGatewayDhl creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayDhl

	c, err = client.ShipmentGatewayDhl.Create().SetStringRef("dhl_config__shipmentgatewaydhl").
		SetConnectURL("https://eCommerce.airborne.com/ApiLandingTest.asp").
		SetConnectTimeout(60).
		SetHeadVersion("1.1").
		SetHeadAction("Request").
		SetAccessUserID("YOUR DHL ShipIT USER ID").
		SetAccessPassword("YOUR DHL ShipIT ACCESS PASSWORD").
		SetAccessAccountNbr("YOUR DHL ShipIT ACCOUNT NUMBER").
		SetAccessShippingKey("YOUR DHL ShipIT SHIPPING KEY").
		SetLabelImageFormat("PNG").
		SetRateEstimateTemplate("api.schema.DHL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create dhl_config__shipmentgatewaydhl: %v", err)
		return err
	}
	cache.Put("dhl_config__shipmentgatewaydhl", c)

	return nil
}
