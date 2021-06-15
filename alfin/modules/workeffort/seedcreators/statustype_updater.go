package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateStatusType(ctx context.Context) error {
	log.Println("StatusType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusType

	c = cache.Get("party_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Party").
		AddStatusItems(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("party_disabled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_status__statustype", c)
	}

	c = cache.Get("case_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Case").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update case_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("case_status__statustype", c)
	}

	c = cache.Get("com_event_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Communication Event").
		AddStatusItems(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_unknown_party__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_resolved__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_referred__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_bounced__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		AddChildren(cache.Get("com_event_rol_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_event_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("com_event_status__statustype", c)
	}

	c = cache.Get("com_event_rol_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Communication Event Role").
		SetParent(cache.Get("com_event_status__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("com_role_created__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_role_read__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("com_role_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_event_rol_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("com_event_rol_status__statustype", c)
	}

	c = cache.Get("party_rel_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Party Relationship").
		AddStatusItems(cache.Get("partyrel_created__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("partyrel_expired__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_rel_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_rel_status__statustype", c)
	}

	c = cache.Get("party_inv_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Party Invitation").
		AddStatusItems(cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("partyinv_pending__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("partyinv_accepted__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("partyinv_declined__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("partyinv_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_inv_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_inv_status__statustype", c)
	}

	c = cache.Get("project__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Project").
		AddChildren(cache.Get("project_status__statustype").(*ent.StatusType)).
		AddChildren(cache.Get("project_task_status__statustype").(*ent.StatusType)).
		AddChildren(cache.Get("project_assgn_status__statustype").(*ent.StatusType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project__statustype", c)
	}

	c = cache.Get("project_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Project status").
		SetParent(cache.Get("project__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("prj_active__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("prj_closed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_status__statustype", c)
	}

	c = cache.Get("project_task_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Project Task").
		SetParent(cache.Get("project__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_created_ua__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_created_as__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_created_ip__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_completed__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_on_hold__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pts_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_task_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_task_status__statustype", c)
	}

	c = cache.Get("project_assgn_status__statustype").(*ent.StatusType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Project Assignment").
		SetParent(cache.Get("project__statustype").(*ent.StatusType)).
		AddStatusItems(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		AddStatusItems(cache.Get("pas_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_assgn_status__statustype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_assgn_status__statustype", c)
	}

	return nil
}
