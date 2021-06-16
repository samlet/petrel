package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateFacilityGroup(ctx context.Context) error {
	log.Println("FacilityGroup creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.FacilityGroup

	c, err = client.FacilityGroup.Create().SetStringRef("_na___facilitygroup").
		SetFacilityGroupName("Not Applicable").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___facilitygroup: %v", err)
		return err
	}
	cache.Put("_na___facilitygroup", c)

	return nil
}
