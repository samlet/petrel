package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentGatewayConfigType(ctx context.Context) error {
	log.Println("ShipmentGatewayConfigType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentGatewayConfigType

	c, err = client.ShipmentGatewayConfigType.Create().SetStringRef("ship_gateway_dhl__shipmentgatewayconfigtype").
		SetHasTable("Yes").
		SetDescription("Shipment Gateway DHL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_gateway_dhl__shipmentgatewayconfigtype: %v", err)
		return err
	}
	cache.Put("ship_gateway_dhl__shipmentgatewayconfigtype", c)

	c, err = client.ShipmentGatewayConfigType.Create().SetStringRef("ship_gateway_fedex__shipmentgatewayconfigtype").
		SetHasTable("Yes").
		SetDescription("Shipment Gateway Fedex").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_gateway_fedex__shipmentgatewayconfigtype: %v", err)
		return err
	}
	cache.Put("ship_gateway_fedex__shipmentgatewayconfigtype", c)

	c, err = client.ShipmentGatewayConfigType.Create().SetStringRef("ship_gateway_ups__shipmentgatewayconfigtype").
		SetHasTable("Yes").
		SetDescription("Shipment Gateway UPS").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_gateway_ups__shipmentgatewayconfigtype: %v", err)
		return err
	}
	cache.Put("ship_gateway_ups__shipmentgatewayconfigtype", c)

	c, err = client.ShipmentGatewayConfigType.Create().SetStringRef("ship_gateway_usps__shipmentgatewayconfigtype").
		SetHasTable("Yes").
		SetDescription("Shipment Gateway USPS").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_gateway_usps__shipmentgatewayconfigtype: %v", err)
		return err
	}
	cache.Put("ship_gateway_usps__shipmentgatewayconfigtype", c)

	return nil
}
