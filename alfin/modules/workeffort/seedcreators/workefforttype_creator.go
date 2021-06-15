package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffortType(ctx context.Context) error {
	log.Println("WorkEffortType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortType

	c, err = client.WorkEffortType.Create().SetStringRef("template__workefforttype").
		SetDescription("Template").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create template__workefforttype: %v", err)
		return err
	}
	cache.Put("template__workefforttype", c)

	c, err = client.WorkEffortType.Create().SetStringRef("project_template__workefforttype").
		SetDescription("Project Template").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_template__workefforttype: %v", err)
		return err
	}
	cache.Put("project_template__workefforttype", c)

	c, err = client.WorkEffortType.Create().SetStringRef("phase_template__workefforttype").
		SetDescription("Project Phase Template").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phase_template__workefforttype: %v", err)
		return err
	}
	cache.Put("phase_template__workefforttype", c)

	c, err = client.WorkEffortType.Create().SetStringRef("task_template__workefforttype").
		SetDescription("Project Task Template").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create task_template__workefforttype: %v", err)
		return err
	}
	cache.Put("task_template__workefforttype", c)

	return nil
}
