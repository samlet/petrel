package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentGatewayUsps(ctx context.Context) error {
	log.Println("ShipmentGatewayUsps updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayUsps
	failures := 0

	c = cache.Get("usps_config__shipmentgatewayusps").(*ent.ShipmentGatewayUsps)
	c, err = c.Update().
		SetConnectURL("http://production.shippingapis.com/ShippingAPITest.dll").
		SetConnectURLLabels("http://production.shippingapis.com/ShippingAPITest.dll").
		SetConnectTimeout(60).
		SetAccessUserID("000000000000").
		SetAccessPassword("999999999999").
		SetMaxEstimateWeight(70).
		SetTest("N").
		SetShipmentGatewayConfig(cache.Get("usps_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update usps_config__shipmentgatewayusps: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("usps_config__shipmentgatewayusps", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
