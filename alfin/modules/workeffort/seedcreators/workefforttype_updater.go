package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateWorkEffortType(ctx context.Context) error {
	log.Println("WorkEffortType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortType

	c = cache.Get("template__workefforttype").(*ent.WorkEffortType)
	c, err = c.Update().
		SetDescription("Template").
		AddChildren(cache.Get("project_template__workefforttype").(*ent.WorkEffortType)).
		AddChildren(cache.Get("phase_template__workefforttype").(*ent.WorkEffortType)).
		AddChildren(cache.Get("task_template__workefforttype").(*ent.WorkEffortType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update template__workefforttype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("template__workefforttype", c)
	}

	c = cache.Get("project_template__workefforttype").(*ent.WorkEffortType)
	c, err = c.Update().
		SetDescription("Project Template").
		SetParent(cache.Get("template__workefforttype").(*ent.WorkEffortType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_template__workefforttype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_template__workefforttype", c)
	}

	c = cache.Get("phase_template__workefforttype").(*ent.WorkEffortType)
	c, err = c.Update().
		SetDescription("Project Phase Template").
		SetParent(cache.Get("template__workefforttype").(*ent.WorkEffortType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update phase_template__workefforttype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("phase_template__workefforttype", c)
	}

	c = cache.Get("task_template__workefforttype").(*ent.WorkEffortType)
	c, err = c.Update().
		SetDescription("Project Task Template").
		SetParent(cache.Get("template__workefforttype").(*ent.WorkEffortType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update task_template__workefforttype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("task_template__workefforttype", c)
	}

	return nil
}