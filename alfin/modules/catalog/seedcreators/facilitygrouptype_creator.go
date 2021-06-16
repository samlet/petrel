package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateFacilityGroupType(ctx context.Context) error {
	log.Println("FacilityGroupType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityGroupType

	c, err = client.FacilityGroupType.Create().SetStringRef("mgmt_structure__facilitygrouptype").
		SetDescription("Management Structure").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create mgmt_structure__facilitygrouptype: %v", err)
		return err
	}
	cache.Put("mgmt_structure__facilitygrouptype", c)

	c, err = client.FacilityGroupType.Create().SetStringRef("pricing_group__facilitygrouptype").
		SetDescription("Pricing Group").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pricing_group__facilitygrouptype: %v", err)
		return err
	}
	cache.Put("pricing_group__facilitygrouptype", c)

	return nil
}
