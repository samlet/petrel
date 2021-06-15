package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateWorkEffortAssoc(ctx context.Context) error {
	log.Println("WorkEffortAssoc updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortAssoc

	c = cache.Get("9002__9003__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9002__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9003__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9002__9003__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9002__9003__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("9003__9005__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9003__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9005__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9003__9005__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9003__9005__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("9005__9006__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9005__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9006__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9005__9006__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9005__9006__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("9102__9103__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9102__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9103__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9102__9103__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9102__9103__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("9103__9105__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9103__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9105__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9103__9105__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9103__9105__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("9105__9106__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("9105__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("9106__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9105__9106__work_eff_dependency__946684800__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("9105__9106__work_eff_dependency__946684800__workeffortassoc", c)
	}

	c = cache.Get("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("calendar_pub_demo__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc", c)
	}

	c = cache.Get("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("project_pub_demo__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("oneoffmeeting__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc", c)
	}

	c = cache.Get("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("project_pub_demo__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("privatedemoemployee1__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc", c)
	}

	c = cache.Get("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)
	c, err = c.Update().
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetFromWorkEffort(cache.Get("project_pub_demo__workeffort").(*ent.WorkEffort)).
		SetToWorkEffort(cache.Get("publicevent__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc", c)
	}

	return nil
}
