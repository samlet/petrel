package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateSecurityGroupPermission(ctx context.Context) error {
	log.Println("SecurityGroupPermission updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityGroupPermission

	c = cache.Get("super__partymgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update super__partymgr_admin__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("super__partymgr_admin__989755200__securitygrouppermission", c)
	}

	c = cache.Get("super__security_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update super__security_admin__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("super__security_admin__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__projectmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__projectmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("viewadmin__projectmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update viewadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("viewadmin__projectmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("bizadmin__projectmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bizadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("bizadmin__projectmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("fulladmin__workeffortmgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fulladmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("fulladmin__workeffortmgr_admin__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_delete__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_delete__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_delete__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_update__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_update__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_update__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission", c)
	}

	c = cache.Get("flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("viewadmin__workeffortmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update viewadmin__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("viewadmin__workeffortmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("bizadmin__workeffortmgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update bizadmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("bizadmin__workeffortmgr_admin__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffort_user__ofbtools_view__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffort_user__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffort_user__ofbtools_view__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffort_user__ofbtools_view__989755200__securitygrouppermission", c)
	}

	c = cache.Get("workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)
	c, err = c.Update().
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		SetSecurityGroup(cache.Get("workeffortadmin__securitygroup").(*ent.SecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission", c)
	}

	return nil
}
