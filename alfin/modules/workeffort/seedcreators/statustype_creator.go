package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateStatusType(ctx context.Context) error {
	log.Println("StatusType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusType

	c, err = client.StatusType.Create().SetStringRef("party_status__statustype").
		SetHasTable("No").
		SetDescription("Party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_status__statustype: %v", err)
		return err
	}
	cache.Put("party_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("case_status__statustype").
		SetHasTable("No").
		SetDescription("Case").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create case_status__statustype: %v", err)
		return err
	}
	cache.Put("case_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("com_event_status__statustype").
		SetHasTable("No").
		SetDescription("Communication Event").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_event_status__statustype: %v", err)
		return err
	}
	cache.Put("com_event_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("com_event_rol_status__statustype").
		SetHasTable("No").
		SetDescription("Communication Event Role").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_event_rol_status__statustype: %v", err)
		return err
	}
	cache.Put("com_event_rol_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("party_rel_status__statustype").
		SetHasTable("No").
		SetDescription("Party Relationship").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_rel_status__statustype: %v", err)
		return err
	}
	cache.Put("party_rel_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("party_inv_status__statustype").
		SetHasTable("No").
		SetDescription("Party Invitation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_inv_status__statustype: %v", err)
		return err
	}
	cache.Put("party_inv_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("project__statustype").
		SetHasTable("No").
		SetDescription("Project").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project__statustype: %v", err)
		return err
	}
	cache.Put("project__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("project_status__statustype").
		SetHasTable("No").
		SetDescription("Project status").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_status__statustype: %v", err)
		return err
	}
	cache.Put("project_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("project_task_status__statustype").
		SetHasTable("No").
		SetDescription("Project Task").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_task_status__statustype: %v", err)
		return err
	}
	cache.Put("project_task_status__statustype", c)

	c, err = client.StatusType.Create().SetStringRef("project_assgn_status__statustype").
		SetHasTable("No").
		SetDescription("Project Assignment").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_assgn_status__statustype: %v", err)
		return err
	}
	cache.Put("project_assgn_status__statustype", c)

	return nil
}
