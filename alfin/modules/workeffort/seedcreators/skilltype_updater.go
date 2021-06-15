package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateSkillType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SkillType

	c = cache.Get("9000__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("Java/Groovy/BSH").
		AddWorkEffortSkillStandards(cache.Get("9002__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9003__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9005__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9006__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9102__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9103__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9105__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9106__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		AddWorkEffortSkillStandards(cache.Get("9202__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__skilltype: %v", err)
		return err
	}
	cache.Put("9000__skilltype", c)

	c = cache.Get("9001__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("Mini Language").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9001__skilltype: %v", err)
		return err
	}
	cache.Put("9001__skilltype", c)

	c = cache.Get("9002__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("HTML/FTL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__skilltype: %v", err)
		return err
	}
	cache.Put("9002__skilltype", c)

	c = cache.Get("9003__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("Javascript/Dojo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9003__skilltype: %v", err)
		return err
	}
	cache.Put("9003__skilltype", c)

	c = cache.Get("9004__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("Screens/Forms").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9004__skilltype: %v", err)
		return err
	}
	cache.Put("9004__skilltype", c)

	c = cache.Get("9005__skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("OFBiz Installation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9005__skilltype: %v", err)
		return err
	}
	cache.Put("9005__skilltype", c)

	c = cache.Get("_na___skilltype").(*ent.SkillType)
	c, err = c.Update().
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___skilltype: %v", err)
		return err
	}
	cache.Put("_na___skilltype", c)

	return nil
}
