package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateShipmentContactMechType(ctx context.Context) error {
	log.Println("ShipmentContactMechType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ShipmentContactMechType

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_to_address__shipmentcontactmechtype").
		SetDescription("Ship-To Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_to_address__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_to_address__shipmentcontactmechtype", c)

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_from_address__shipmentcontactmechtype").
		SetDescription("Ship-From Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_from_address__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_from_address__shipmentcontactmechtype", c)

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_to_telecom__shipmentcontactmechtype").
		SetDescription("Ship-To Telecom Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_to_telecom__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_to_telecom__shipmentcontactmechtype", c)

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_from_telecom__shipmentcontactmechtype").
		SetDescription("Ship-From Telecom Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_from_telecom__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_from_telecom__shipmentcontactmechtype", c)

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_to_email__shipmentcontactmechtype").
		SetDescription("Ship-To E-Mail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_to_email__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_to_email__shipmentcontactmechtype", c)

	c, err = client.ShipmentContactMechType.Create().SetStringRef("ship_from_email__shipmentcontactmechtype").
		SetDescription("Ship-From E-Mail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_from_email__shipmentcontactmechtype: %v", err)
		return err
	}
	cache.Put("ship_from_email__shipmentcontactmechtype", c)

	return nil
}
