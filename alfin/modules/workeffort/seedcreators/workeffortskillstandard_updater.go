package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateWorkEffortSkillStandard(ctx context.Context) error {
	log.Println("WorkEffortSkillStandard updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortSkillStandard

	c = cache.Get("9002__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(16.0).
		SetWorkEffort(cache.Get("9002__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9002__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9002__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9003__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(24.0).
		SetWorkEffort(cache.Get("9003__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9003__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9003__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9005__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(32.0).
		SetWorkEffort(cache.Get("9005__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9005__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9005__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9006__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(40.0).
		SetWorkEffort(cache.Get("9006__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9006__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9006__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9102__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(16.0).
		SetWorkEffort(cache.Get("9102__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9102__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9102__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9103__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(24.0).
		SetWorkEffort(cache.Get("9103__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9103__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9103__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9105__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(32.0).
		SetWorkEffort(cache.Get("9105__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9105__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9105__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9106__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(40.0).
		SetWorkEffort(cache.Get("9106__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9106__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9106__9000__workeffortskillstandard", c)
	}

	c = cache.Get("9202__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)
	c, err = c.Update().
		SetEstimatedDuration(16.0).
		SetWorkEffort(cache.Get("9202__workeffort").(*ent.WorkEffort)).
		SetSkillType(cache.Get("9000__skilltype").(*ent.SkillType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9202__9000__workeffortskillstandard: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9202__9000__workeffortskillstandard", c)
	}

	return nil
}
