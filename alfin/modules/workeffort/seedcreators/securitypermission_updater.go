package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateSecurityPermission(ctx context.Context) error {
	log.Println("SecurityPermission updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SecurityPermission
	failures := 0

	c = cache.Get("partymgr_view__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("View operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_view__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_view__securitypermission", c)
	}

	c = cache.Get("partymgr_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_create__securitypermission", c)
	}

	c = cache.Get("partymgr_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_update__securitypermission", c)
	}

	c = cache.Get("partymgr_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete operations in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_admin__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("ALL operations in the Party Manager.").
		AddSecurityGroupPermissions(cache.Get("super__partymgr_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_admin__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_admin__securitypermission", c)
	}

	c = cache.Get("partymgr_note__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create notes in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_note__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_note__securitypermission", c)
	}

	c = cache.Get("partymgr_sts_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update party status in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_sts_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_sts_update__securitypermission", c)
	}

	c = cache.Get("partymgr_grp_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update PartyGroup or Person detail information.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_grp_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_grp_update__securitypermission", c)
	}

	c = cache.Get("partymgr_rel_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_rel_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_rel_create__securitypermission", c)
	}

	c = cache.Get("partymgr_rel_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_rel_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_rel_update__securitypermission", c)
	}

	c = cache.Get("partymgr_rel_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete party relationships in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_rel_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_rel_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_role_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create party roles in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_role_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_role_create__securitypermission", c)
	}

	c = cache.Get("partymgr_role_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete party roles in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_role_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_role_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_pcm_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_pcm_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_pcm_create__securitypermission", c)
	}

	c = cache.Get("partymgr_pcm_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_pcm_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_pcm_update__securitypermission", c)
	}

	c = cache.Get("partymgr_pcm_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete party contact mechs in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_pcm_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_pcm_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_src_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create party to data source relations.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_src_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_src_create__securitypermission", c)
	}

	c = cache.Get("partymgr_qal_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_qal_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_qal_create__securitypermission", c)
	}

	c = cache.Get("partymgr_qal_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_qal_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_qal_update__securitypermission", c)
	}

	c = cache.Get("partymgr_qal_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete party quals in the Party Manager.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_qal_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_qal_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_cme_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme_create__securitypermission", c)
	}

	c = cache.Get("partymgr_cme_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme_update__securitypermission", c)
	}

	c = cache.Get("partymgr_cme_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete communication event, any from/to party.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_cme-email_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Can create Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme-email_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme-email_create__securitypermission", c)
	}

	c = cache.Get("partymgr_cme-email_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Can update Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme-email_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme-email_update__securitypermission", c)
	}

	c = cache.Get("partymgr_cme-email_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Can delete Email communication events for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme-email_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme-email_delete__securitypermission", c)
	}

	c = cache.Get("partymgr_cme-note_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Can create Internal note communication event for logged-in user.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partymgr_cme-note_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partymgr_cme-note_create__securitypermission", c)
	}

	c = cache.Get("security_view__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("View operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update security_view__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("security_view__securitypermission", c)
	}

	c = cache.Get("security_create__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Create operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update security_create__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("security_create__securitypermission", c)
	}

	c = cache.Get("security_update__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Update operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update security_update__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("security_update__securitypermission", c)
	}

	c = cache.Get("security_delete__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("Delete operations in the Security Management Screens.").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update security_delete__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("security_delete__securitypermission", c)
	}

	c = cache.Get("security_admin__securitypermission").(*ent.SecurityPermission)
	c, err = c.Update().
		SetDescription("ALL operations in the Security Management Screens.").
		AddSecurityGroupPermissions(cache.Get("super__security_admin__989755200__securitygrouppermission").(*ent.SecurityGroupPermission)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update security_admin__securitypermission: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("security_admin__securitypermission", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
