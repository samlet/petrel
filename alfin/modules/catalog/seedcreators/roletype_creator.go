package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateRoleType(ctx context.Context) error {
	log.Println("RoleType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RoleType

	c, err = client.RoleType.Create().SetStringRef("imageapprover__roletype").
		SetHasTable("No").
		SetDescription("Image Approver").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create imageapprover__roletype: %v", err)
		return err
	}
	cache.Put("imageapprover__roletype", c)

	return nil
}
