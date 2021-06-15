package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateSecurityGroupPermission(ctx context.Context) error {
	log.Println("SecurityGroupPermission creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityGroupPermission

	c, err = client.SecurityGroupPermission.Create().SetStringRef("super__partymgr_admin__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create super__partymgr_admin__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("super__partymgr_admin__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("super__security_admin__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create super__security_admin__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("super__security_admin__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__projectmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__projectmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("viewadmin__projectmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create viewadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("viewadmin__projectmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("viewadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("bizadmin__projectmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bizadmin__projectmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("bizadmin__projectmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("bizadmin__projectmgr_role_timesheet_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("fulladmin__workeffortmgr_admin__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fulladmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("fulladmin__workeffortmgr_admin__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_delete__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_delete__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_delete__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_update__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_update__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_update__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_role_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_role_update__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("flexadmin__workeffortmgr_role_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("viewadmin__workeffortmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create viewadmin__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("viewadmin__workeffortmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("bizadmin__workeffortmgr_admin__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create bizadmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("bizadmin__workeffortmgr_admin__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__workeffortmgr_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__workeffortmgr_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__workeffortmgr_role_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__workeffortmgr_role_create__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__workeffortmgr_role_update__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__workeffortmgr_role_delete__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffort_user__ofbtools_view__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffort_user__ofbtools_view__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffort_user__ofbtools_view__989755200__securitygrouppermission", c)

	c, err = client.SecurityGroupPermission.Create().SetStringRef("workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission").
		SetFromDate(common.ParseDateTime("2001-05-13 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission: %v", err)
		return err
	}
	cache.Put("workeffortadmin__workeffortmgr_admin__989755200__securitygrouppermission", c)

	return nil
}
