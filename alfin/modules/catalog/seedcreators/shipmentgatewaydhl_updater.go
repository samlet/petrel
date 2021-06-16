package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentGatewayDhl(ctx context.Context) error {
	log.Println("ShipmentGatewayDhl updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayDhl
	failures := 0

	c = cache.Get("dhl_config__shipmentgatewaydhl").(*ent.ShipmentGatewayDhl)
	c, err = c.Update().
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
		SetShipmentGatewayConfig(cache.Get("dhl_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update dhl_config__shipmentgatewaydhl: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("dhl_config__shipmentgatewaydhl", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
