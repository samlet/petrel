package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentType(ctx context.Context) error {
	log.Println("ShipmentType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentType
	failures := 0

	c = cache.Get("incoming_shipment__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Incoming").
		AddChildren(cache.Get("sales_return__shipmenttype").(*ent.ShipmentType)).
		AddChildren(cache.Get("purchase_shipment__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update incoming_shipment__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("incoming_shipment__shipmenttype", c)
	}

	c = cache.Get("outgoing_shipment__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Outgoing").
		AddChildren(cache.Get("sales_shipment__shipmenttype").(*ent.ShipmentType)).
		AddChildren(cache.Get("purchase_return__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update outgoing_shipment__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("outgoing_shipment__shipmenttype", c)
	}

	c = cache.Get("sales_return__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Return").
		SetParent(cache.Get("incoming_shipment__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sales_return__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sales_return__shipmenttype", c)
	}

	c = cache.Get("sales_shipment__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Shipment").
		SetParent(cache.Get("outgoing_shipment__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sales_shipment__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sales_shipment__shipmenttype", c)
	}

	c = cache.Get("purchase_shipment__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Purchase Shipment").
		SetParent(cache.Get("incoming_shipment__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purchase_shipment__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purchase_shipment__shipmenttype", c)
	}

	c = cache.Get("purchase_return__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Purchase Return").
		SetParent(cache.Get("outgoing_shipment__shipmenttype").(*ent.ShipmentType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update purchase_return__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("purchase_return__shipmenttype", c)
	}

	c = cache.Get("drop_shipment__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Drop Shipment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update drop_shipment__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("drop_shipment__shipmenttype", c)
	}

	c = cache.Get("transfer__shipmenttype").(*ent.ShipmentType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Transfer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update transfer__shipmenttype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("transfer__shipmenttype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
