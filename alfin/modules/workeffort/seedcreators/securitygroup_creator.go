package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateSecurityGroup(ctx context.Context) error {
	log.Println("SecurityGroup creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityGroup

	c, err = client.SecurityGroup.Create().SetStringRef("workeffort_user__securitygroup").
		SetGroupName("WorkEffort User").
		SetDescription("WorkEffort user group; all limited workeffort permissions.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__securitygroup: %v", err)
		return err
	}
	cache.Put("workeffort_user__securitygroup", c)

	c, err = client.SecurityGroup.Create().SetStringRef("workeffortadmin__securitygroup").
		SetGroupName("WorkEffort Admin").
		SetDescription("WorkEffort Admin group, has all workeffort permissions.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortadmin__securitygroup: %v", err)
		return err
	}
	cache.Put("workeffortadmin__securitygroup", c)

	return nil
}
