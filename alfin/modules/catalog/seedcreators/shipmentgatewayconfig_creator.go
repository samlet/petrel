package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayConfig(ctx context.Context) error {
	log.Println("ShipmentGatewayConfig creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayConfig

	c, err = client.ShipmentGatewayConfig.Create().SetStringRef("dhl_config__shipmentgatewayconfig").
		SetDescription("DHL Config").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create dhl_config__shipmentgatewayconfig: %v", err)
		return err
	}
	cache.Put("dhl_config__shipmentgatewayconfig", c)

	c, err = client.ShipmentGatewayConfig.Create().SetStringRef("fedex_config__shipmentgatewayconfig").
		SetDescription("Fedex Config").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fedex_config__shipmentgatewayconfig: %v", err)
		return err
	}
	cache.Put("fedex_config__shipmentgatewayconfig", c)

	c, err = client.ShipmentGatewayConfig.Create().SetStringRef("ups_config__shipmentgatewayconfig").
		SetDescription("UPS Config").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ups_config__shipmentgatewayconfig: %v", err)
		return err
	}
	cache.Put("ups_config__shipmentgatewayconfig", c)

	c, err = client.ShipmentGatewayConfig.Create().SetStringRef("usps_config__shipmentgatewayconfig").
		SetDescription("USPS Config").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create usps_config__shipmentgatewayconfig: %v", err)
		return err
	}
	cache.Put("usps_config__shipmentgatewayconfig", c)

	return nil
}
