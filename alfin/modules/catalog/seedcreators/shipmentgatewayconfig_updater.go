package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentGatewayConfig(ctx context.Context) error {
	log.Println("ShipmentGatewayConfig updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayConfig
	failures := 0

	c = cache.Get("dhl_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)
	c, err = c.Update().
		SetDescription("DHL Config").
		SetShipmentGatewayConfigType(cache.Get("ship_gateway_dhl__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update dhl_config__shipmentgatewayconfig: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("dhl_config__shipmentgatewayconfig", c)
	}

	c = cache.Get("fedex_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)
	c, err = c.Update().
		SetDescription("Fedex Config").
		SetShipmentGatewayConfigType(cache.Get("ship_gateway_fedex__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fedex_config__shipmentgatewayconfig: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("fedex_config__shipmentgatewayconfig", c)
	}

	c = cache.Get("ups_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)
	c, err = c.Update().
		SetDescription("UPS Config").
		SetShipmentGatewayConfigType(cache.Get("ship_gateway_ups__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ups_config__shipmentgatewayconfig: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ups_config__shipmentgatewayconfig", c)
	}

	c = cache.Get("usps_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)
	c, err = c.Update().
		SetDescription("USPS Config").
		SetShipmentGatewayConfigType(cache.Get("ship_gateway_usps__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update usps_config__shipmentgatewayconfig: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("usps_config__shipmentgatewayconfig", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
