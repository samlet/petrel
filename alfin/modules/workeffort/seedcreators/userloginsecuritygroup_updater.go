package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateUserLoginSecurityGroup(ctx context.Context) error {
	log.Println("UserLoginSecurityGroup updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserLoginSecurityGroup

	c = cache.Get("demoemployee__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("demoemployee__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__projectadmin__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__projectadmin__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("demoemployee1__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("demoemployee1__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__projectadmin__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__projectadmin__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("demoemployee2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("demoemployee2__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__projectuser__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__projectuser__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("demoemployee3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("demoemployee3__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__projectuser__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__projectuser__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("democustomer1__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("democustomer1__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__projectuser__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__projectuser__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("democustomer2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("democustomer2__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__projectuser__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__projectuser__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("democustomer3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("democustomer3__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__projectuser__946684800__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__projectuser__946684800__userloginsecuritygroup", c)
	}

	c = cache.Get("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2011-01-01 00:00:00.0")).
		SetUserLogin(cache.Get("workeffortuser__userlogin").(*ent.UserLogin)).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		AddSecurityGroupPermissions(cache.Get("workeffort_user__ofbtools_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffortuser__workeffort_user__1293840000__userloginsecuritygroup: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup", c)
	}

	return nil
}
