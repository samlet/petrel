package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"

	"fmt"
)

func UpdateShipmentContactMechType(ctx context.Context) error {
	log.Println("ShipmentContactMechType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentContactMechType
	failures := 0

	c = cache.Get("ship_to_address__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-To Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_to_address__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_to_address__shipmentcontactmechtype", c)
	}

	c = cache.Get("ship_from_address__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-From Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_from_address__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_from_address__shipmentcontactmechtype", c)
	}

	c = cache.Get("ship_to_telecom__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-To Telecom Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_to_telecom__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_to_telecom__shipmentcontactmechtype", c)
	}

	c = cache.Get("ship_from_telecom__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-From Telecom Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_from_telecom__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_from_telecom__shipmentcontactmechtype", c)
	}

	c = cache.Get("ship_to_email__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-To E-Mail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_to_email__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_to_email__shipmentcontactmechtype", c)
	}

	c = cache.Get("ship_from_email__shipmentcontactmechtype").(*ent.ShipmentContactMechType)
	c, err = c.Update().
		SetDescription("Ship-From E-Mail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ship_from_email__shipmentcontactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ship_from_email__shipmentcontactmechtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
