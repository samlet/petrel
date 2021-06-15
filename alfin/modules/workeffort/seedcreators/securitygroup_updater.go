package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateSecurityGroup(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityGroup

	c = cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)
	c, err = c.Update().
		SetGroupName("WorkEffort User").
		SetDescription("WorkEffort user group; all limited workeffort permissions.").
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__ofbtools_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddUserLoginSecurityGroups(cache.Get("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__securitygroup: %v", err)
		return err
	}
	cache.Put("workeffort_user__securitygroup", c)

	c = cache.Get("workeffortadmin__securitygroup").(*ent.SecurityGroup)
	c, err = c.Update().
		SetGroupName("WorkEffort Admin").
		SetDescription("WorkEffort Admin group, has all workeffort permissions.").
		AddSecurityGroupPermissions(cache.Get("workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortadmin__securitygroup: %v", err)
		return err
	}
	cache.Put("workeffortadmin__securitygroup", c)

	return nil
}
