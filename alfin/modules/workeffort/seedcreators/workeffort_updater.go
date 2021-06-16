package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateWorkEffort(ctx context.Context) error {
	log.Println("WorkEffort updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffort
	failures := 0

	c = cache.Get("9000__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.901")).
		SetWorkEffortName("Demo Project1 Cust1").
		SetRevisionNumber(1).
		AddChildren(cache.Get("9001__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9004__workeffort").(*ent.WorkEffort)).
		AddWorkEffortPartyAssignments(cache.Get("9000__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9000__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9000__workeffort", c)
	}

	c = cache.Get("9001__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9002__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9003__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9001__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9001__workeffort", c)
	}

	c = cache.Get("9002__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9001__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9002__9003__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortPartyAssignments(cache.Get("9002__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortSkillStandards(cache.Get("9002__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9002__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9002__workeffort", c)
	}

	c = cache.Get("9003__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:29.453")).
		SetWorkEffortName("task2").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9001__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9003__9005__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddToWorkEffortAssocs(cache.Get("9002__9003__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9003__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9003__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9003__workeffort", c)
	}

	c = cache.Get("9004__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:35.939")).
		SetWorkEffortName("phase2").
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9000__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9005__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9006__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9004__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9004__workeffort", c)
	}

	c = cache.Get("9005__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:50.84")).
		SetWorkEffortName("task3").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-17 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9004__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9005__9006__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddToWorkEffortAssocs(cache.Get("9003__9005__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9005__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9005__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9005__workeffort", c)
	}

	c = cache.Get("9006__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:58.857")).
		SetWorkEffortName("task4").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-18 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-20 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9004__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddToWorkEffortAssocs(cache.Get("9005__9006__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9006__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9006__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9006__workeffort", c)
	}

	c = cache.Get("9100__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.911")).
		SetWorkEffortName("Demo Project2 Cust 2").
		SetRevisionNumber(1).
		AddChildren(cache.Get("9101__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9104__workeffort").(*ent.WorkEffort)).
		AddWorkEffortPartyAssignments(cache.Get("9100__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9100__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9100__workeffort", c)
	}

	c = cache.Get("9101__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9102__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9103__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9101__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9101__workeffort", c)
	}

	c = cache.Get("9102__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9101__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9102__9103__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9102__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9102__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9102__workeffort", c)
	}

	c = cache.Get("9103__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:29.453")).
		SetWorkEffortName("task2").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9101__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9103__9105__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddToWorkEffortAssocs(cache.Get("9102__9103__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9103__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9103__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9103__workeffort", c)
	}

	c = cache.Get("9104__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:35.939")).
		SetWorkEffortName("phase2").
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9100__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9105__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9106__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9104__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9104__workeffort", c)
	}

	c = cache.Get("9105__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:50.84")).
		SetWorkEffortName("task3").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-17 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		SetParent(cache.Get("9104__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddFromWorkEffortAssocs(cache.Get("9105__9106__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddToWorkEffortAssocs(cache.Get("9103__9105__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9105__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9105__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9105__workeffort", c)
	}

	c = cache.Get("9106__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:58.857")).
		SetWorkEffortName("task4").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-18 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-20 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		SetParent(cache.Get("9104__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddToWorkEffortAssocs(cache.Get("9105__9106__work_eff_dependency__946684800__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortSkillStandards(cache.Get("9106__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9106__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9106__workeffort", c)
	}

	c = cache.Get("9200__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.911")).
		SetWorkEffortName("Demo Project3 DemoCustomerCompany").
		SetRevisionNumber(1).
		AddChildren(cache.Get("9201__workeffort").(*ent.WorkEffort)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__company__internal_organizatio__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9200__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9200__workeffort", c)
	}

	c = cache.Get("9201__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		SetParent(cache.Get("9200__workeffort").(*ent.WorkEffort)).
		AddChildren(cache.Get("9202__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9201__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9201__workeffort", c)
	}

	c = cache.Get("9202__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		SetParent(cache.Get("9201__workeffort").(*ent.WorkEffort)).
		SetCurrentStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		AddWorkEffortSkillStandards(cache.Get("9202__9000__workeffortskillstandard").(*ent.WorkEffortSkillStandard)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update 9202__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("9202__workeffort", c)
	}

	c = cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("Staff Meeting").
		SetDescription("Staff Meeting").
		SetEstimatedStartDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetEstimatedMilliSeconds(3600000).
		SetTemporalExpression(cache.Get("staff_mtg__temporalexpression").(*ent.TemporalExpression)).
		AddToWorkEffortAssocs(cache.Get("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortFixedAssetAssigns(cache.Get("staff_mtg__demo_projector__1199145600__workeffortfixedassetassign").(*ent.WorkEffortFixedAssetAssign)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__workeffort", c)
	}

	c = cache.Get("calendar_pub_demo__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetWorkEffortName("iCalendar Publish Demonstration").
		SetDescription("Demo Project 1 Customer 1").
		AddFromWorkEffortAssocs(cache.Get("calendar_pub_demo__staff_mtg__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortPartyAssignments(cache.Get("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update calendar_pub_demo__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("calendar_pub_demo__workeffort", c)
	}

	c = cache.Get("project_pub_demo__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetWorkEffortName("Project iCalendar Publish Demonstration").
		SetDescription("Demo Project 1 Customer 1").
		AddFromWorkEffortAssocs(cache.Get("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddFromWorkEffortAssocs(cache.Get("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddFromWorkEffortAssocs(cache.Get("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortPartyAssignments(cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update project_pub_demo__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("project_pub_demo__workeffort", c)
	}

	c = cache.Get("oneoffmeeting__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("One off Meeting june 12").
		SetDescription("Staff Meeting").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-12 09:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-12 10:00:00.0")).
		AddToWorkEffortAssocs(cache.Get("project_pub_demo__oneoffmeeting__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update oneoffmeeting__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("oneoffmeeting__workeffort", c)
	}

	c = cache.Get("privatedemoemployee1__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("My Wifes birthday june 29").
		SetDescription("The birthday i should never forget").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-29 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-29 23:00:00.0")).
		AddToWorkEffortAssocs(cache.Get("project_pub_demo__privatedemoemployee1__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		AddWorkEffortPartyAssignments(cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update privatedemoemployee1__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("privatedemoemployee1__workeffort", c)
	}

	c = cache.Get("publicevent__workeffort").(*ent.WorkEffort)
	c, err = c.Update().
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("The general company party june 17").
		SetDescription("General Party").
		SetLocationDesc("Tom's Banquet Hall").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-17 19:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-17 23:00:00.0")).
		AddToWorkEffortAssocs(cache.Get("project_pub_demo__publicevent__work_eff_dependency__1199145600__workeffortassoc").(*ent.WorkEffortAssoc)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update publicevent__workeffort: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("publicevent__workeffort", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
