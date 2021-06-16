package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentGatewayConfigType(ctx context.Context) error {
	log.Println("ShipmentGatewayConfigType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayConfigType
	failures := 0

	c = cache.Get("ship_gateway_dhl__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Shipment Gateway DHL").
		AddShipmentGatewayConfigs(cache.Get("dhl_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_gateway_dhl__shipmentgatewayconfigtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_gateway_dhl__shipmentgatewayconfigtype", c)
	}

	c = cache.Get("ship_gateway_fedex__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Shipment Gateway Fedex").
		AddShipmentGatewayConfigs(cache.Get("fedex_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_gateway_fedex__shipmentgatewayconfigtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_gateway_fedex__shipmentgatewayconfigtype", c)
	}

	c = cache.Get("ship_gateway_ups__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Shipment Gateway UPS").
		AddShipmentGatewayConfigs(cache.Get("ups_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_gateway_ups__shipmentgatewayconfigtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_gateway_ups__shipmentgatewayconfigtype", c)
	}

	c = cache.Get("ship_gateway_usps__shipmentgatewayconfigtype").(*ent.ShipmentGatewayConfigType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Shipment Gateway USPS").
		AddShipmentGatewayConfigs(cache.Get("usps_config__shipmentgatewayconfig").(*ent.ShipmentGatewayConfig)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_gateway_usps__shipmentgatewayconfigtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_gateway_usps__shipmentgatewayconfigtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
