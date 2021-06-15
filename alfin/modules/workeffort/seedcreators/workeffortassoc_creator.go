package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffortAssoc(ctx context.Context) error {
	log.Println("WorkEffortAssoc creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffortAssoc

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9002__9003__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__9003__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9002__9003__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9003__9005__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9003__9005__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9003__9005__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9005__9006__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9005__9006__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9005__9006__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9102__9103__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9102__9103__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9102__9103__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9103__9105__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9103__9105__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9103__9105__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("9105__9106__work_eff_dependency__946684800__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetSequenceNum(0).
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9105__9106__work_eff_dependency__946684800__workeffortassoc: %v", err)
		return err
	}
	cache.Put("9105__9106__work_eff_dependency__946684800__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		return err
	}
	cache.Put("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		return err
	}
	cache.Put("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		return err
	}
	cache.Put("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc", c)

	c, err = client.WorkEffortAssoc.Create().SetStringRef("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc").
		SetWorkEffortAssocTypeID(common.ParseId("WORK_EFF_DEPENDENCY")).
		SetFromDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc: %v", err)
		return err
	}
	cache.Put("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc", c)

	return nil
}
