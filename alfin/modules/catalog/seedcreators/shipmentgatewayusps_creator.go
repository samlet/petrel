package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayUsps(ctx context.Context) error {
	log.Println("ShipmentGatewayUsps creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayUsps

	c, err = client.ShipmentGatewayUsps.Create().SetStringRef("usps_config__shipmentgatewayusps").
		SetConnectURL("http://production.shippingapis.com/ShippingAPITest.dll").
		SetConnectURLLabels("http://production.shippingapis.com/ShippingAPITest.dll").
		SetConnectTimeout(60).
		SetAccessUserID("000000000000").
		SetAccessPassword("999999999999").
		SetMaxEstimateWeight(70).
		SetTest("N").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create usps_config__shipmentgatewayusps: %v", err)
		return err
	}
	cache.Put("usps_config__shipmentgatewayusps", c)

	return nil
}
