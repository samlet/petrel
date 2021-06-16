package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentType(ctx context.Context) error {
	log.Println("ShipmentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentType

	c, err = client.ShipmentType.Create().SetStringRef("incoming_shipment__shipmenttype").
		SetHasTable("No").
		SetDescription("Incoming").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create incoming_shipment__shipmenttype: %v", err)
		return err
	}
	cache.Put("incoming_shipment__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("outgoing_shipment__shipmenttype").
		SetHasTable("No").
		SetDescription("Outgoing").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create outgoing_shipment__shipmenttype: %v", err)
		return err
	}
	cache.Put("outgoing_shipment__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("sales_return__shipmenttype").
		SetHasTable("No").
		SetDescription("Sales Return").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_return__shipmenttype: %v", err)
		return err
	}
	cache.Put("sales_return__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("sales_shipment__shipmenttype").
		SetHasTable("No").
		SetDescription("Sales Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_shipment__shipmenttype: %v", err)
		return err
	}
	cache.Put("sales_shipment__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("purchase_shipment__shipmenttype").
		SetHasTable("No").
		SetDescription("Purchase Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purchase_shipment__shipmenttype: %v", err)
		return err
	}
	cache.Put("purchase_shipment__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("purchase_return__shipmenttype").
		SetHasTable("No").
		SetDescription("Purchase Return").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create purchase_return__shipmenttype: %v", err)
		return err
	}
	cache.Put("purchase_return__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("drop_shipment__shipmenttype").
		SetHasTable("No").
		SetDescription("Drop Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create drop_shipment__shipmenttype: %v", err)
		return err
	}
	cache.Put("drop_shipment__shipmenttype", c)

	c, err = client.ShipmentType.Create().SetStringRef("transfer__shipmenttype").
		SetHasTable("No").
		SetDescription("Transfer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create transfer__shipmenttype: %v", err)
		return err
	}
	cache.Put("transfer__shipmenttype", c)

	return nil
}
