package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateFacilityType(ctx context.Context) error {
	log.Println("FacilityType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityType

	c, err = client.FacilityType.Create().SetStringRef("building__facilitytype").
		SetHasTable("No").
		SetDescription("Building").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create building__facilitytype: %v", err)
		return err
	}
	cache.Put("building__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("floor__facilitytype").
		SetHasTable("No").
		SetDescription("Floor").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create floor__facilitytype: %v", err)
		return err
	}
	cache.Put("floor__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("office__facilitytype").
		SetHasTable("No").
		SetDescription("Office").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create office__facilitytype: %v", err)
		return err
	}
	cache.Put("office__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("call_center__facilitytype").
		SetHasTable("No").
		SetDescription("Call Center").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create call_center__facilitytype: %v", err)
		return err
	}
	cache.Put("call_center__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("plant__facilitytype").
		SetHasTable("No").
		SetDescription("Plant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create plant__facilitytype: %v", err)
		return err
	}
	cache.Put("plant__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("room__facilitytype").
		SetHasTable("No").
		SetDescription("Room").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create room__facilitytype: %v", err)
		return err
	}
	cache.Put("room__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("retail_store__facilitytype").
		SetHasTable("No").
		SetDescription("Retail Store").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create retail_store__facilitytype: %v", err)
		return err
	}
	cache.Put("retail_store__facilitytype", c)

	c, err = client.FacilityType.Create().SetStringRef("warehouse__facilitytype").
		SetHasTable("No").
		SetDescription("Warehouse").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create warehouse__facilitytype: %v", err)
		return err
	}
	cache.Put("warehouse__facilitytype", c)

	return nil
}
