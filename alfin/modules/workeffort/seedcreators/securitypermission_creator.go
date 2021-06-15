package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateSecurityPermission(ctx context.Context) error {
	log.Println("SecurityPermission creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityPermission

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_view__securitypermission").
		SetDescription("View operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_view__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_view__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_create__securitypermission").
		SetDescription("Create operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_update__securitypermission").
		SetDescription("Update operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_delete__securitypermission").
		SetDescription("Delete operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_admin__securitypermission").
		SetDescription("ALL operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_admin__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_admin__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_note__securitypermission").
		SetDescription("Create notes in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_note__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_note__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_sts_update__securitypermission").
		SetDescription("Update party status in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_sts_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_sts_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_grp_update__securitypermission").
		SetDescription("Update PartyGroup or Person detail information.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_grp_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_grp_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_rel_create__securitypermission").
		SetDescription("Create party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_rel_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_rel_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_rel_update__securitypermission").
		SetDescription("Update party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_rel_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_rel_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_rel_delete__securitypermission").
		SetDescription("Delete party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_rel_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_rel_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_role_create__securitypermission").
		SetDescription("Create party roles in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_role_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_role_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_role_delete__securitypermission").
		SetDescription("Delete party roles in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_role_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_role_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_pcm_create__securitypermission").
		SetDescription("Create party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_pcm_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_pcm_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_pcm_update__securitypermission").
		SetDescription("Update party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_pcm_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_pcm_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_pcm_delete__securitypermission").
		SetDescription("Delete party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_pcm_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_pcm_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_src_create__securitypermission").
		SetDescription("Create party to data source relations.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_src_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_src_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_qal_create__securitypermission").
		SetDescription("Create party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_qal_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_qal_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_qal_update__securitypermission").
		SetDescription("Update party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_qal_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_qal_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_qal_delete__securitypermission").
		SetDescription("Delete party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_qal_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_qal_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme_create__securitypermission").
		SetDescription("Create communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme_update__securitypermission").
		SetDescription("Update communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme_delete__securitypermission").
		SetDescription("Delete communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme-email_create__securitypermission").
		SetDescription("Can create Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme-email_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme-email_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme-email_update__securitypermission").
		SetDescription("Can update Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme-email_update__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme-email_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme-email_delete__securitypermission").
		SetDescription("Can delete Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme-email_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme-email_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("partymgr_cme-note_create__securitypermission").
		SetDescription("Can create Internal note communication event for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partymgr_cme-note_create__securitypermission: %v", err)
		return err
	}
	cache.Put("partymgr_cme-note_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("security_view__securitypermission").
		SetDescription("View operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create security_view__securitypermission: %v", err)
		return err
	}
	cache.Put("security_view__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("security_create__securitypermission").
		SetDescription("Create operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create security_create__securitypermission: %v", err)
		return err
	}
	cache.Put("security_create__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("security_update__securitypermission").
		SetDescription("Update operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create security_update__securitypermission: %v", err)
		return err
	}
	cache.Put("security_update__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("security_delete__securitypermission").
		SetDescription("Delete operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create security_delete__securitypermission: %v", err)
		return err
	}
	cache.Put("security_delete__securitypermission", c)

	c, err = client.SecurityPermission.Create().SetStringRef("security_admin__securitypermission").
		SetDescription("ALL operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create security_admin__securitypermission: %v", err)
		return err
	}
	cache.Put("security_admin__securitypermission", c)

	return nil
}
