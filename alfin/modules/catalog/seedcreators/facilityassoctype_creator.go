package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateFacilityAssocType(ctx context.Context) error {
	log.Println("FacilityAssocType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityAssocType

	c, err = client.FacilityAssocType.Create().SetStringRef("backup_warehouse__facilityassoctype").
		SetDescription("Facility that serves another facility in terms of inventory").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create backup_warehouse__facilityassoctype: %v", err)
		return err
	}
	cache.Put("backup_warehouse__facilityassoctype", c)

	return nil
}
