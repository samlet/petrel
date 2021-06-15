package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffortSkillStandard(ctx context.Context) error {
	log.Println("WorkEffortSkillStandard creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortSkillStandard

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9002__9000__workeffortskillstandard").
		SetEstimatedDuration(16.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9002__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9003__9000__workeffortskillstandard").
		SetEstimatedDuration(24.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9003__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9003__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9005__9000__workeffortskillstandard").
		SetEstimatedDuration(32.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9005__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9005__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9006__9000__workeffortskillstandard").
		SetEstimatedDuration(40.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9006__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9006__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9102__9000__workeffortskillstandard").
		SetEstimatedDuration(16.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9102__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9102__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9103__9000__workeffortskillstandard").
		SetEstimatedDuration(24.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9103__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9103__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9105__9000__workeffortskillstandard").
		SetEstimatedDuration(32.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9105__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9105__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9106__9000__workeffortskillstandard").
		SetEstimatedDuration(40.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9106__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9106__9000__workeffortskillstandard", c)

	c, err = client.WorkEffortSkillStandard.Create().SetStringRef("9202__9000__workeffortskillstandard").
		SetEstimatedDuration(16.0).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9202__9000__workeffortskillstandard: %v", err)
		return err
	}
	cache.Put("9202__9000__workeffortskillstandard", c)

	return nil
}
