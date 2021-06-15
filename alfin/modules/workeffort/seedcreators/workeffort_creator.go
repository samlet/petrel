package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateWorkEffort(ctx context.Context) error {
	log.Println("WorkEffort creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.WorkEffort

	c, err = client.WorkEffort.Create().SetStringRef("9000__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.901")).
		SetWorkEffortName("Demo Project1 Cust1").
		SetRevisionNumber(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__workeffort: %v", err)
		return err
	}
	cache.Put("9000__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9001__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9001__workeffort: %v", err)
		return err
	}
	cache.Put("9001__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9002__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__workeffort: %v", err)
		return err
	}
	cache.Put("9002__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9003__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:29.453")).
		SetWorkEffortName("task2").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9003__workeffort: %v", err)
		return err
	}
	cache.Put("9003__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9004__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:35.939")).
		SetWorkEffortName("phase2").
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9004__workeffort: %v", err)
		return err
	}
	cache.Put("9004__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9005__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:50.84")).
		SetWorkEffortName("task3").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-17 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9005__workeffort: %v", err)
		return err
	}
	cache.Put("9005__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9006__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:58.857")).
		SetWorkEffortName("task4").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-18 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-20 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9006__workeffort: %v", err)
		return err
	}
	cache.Put("9006__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9100__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.911")).
		SetWorkEffortName("Demo Project2 Cust 2").
		SetRevisionNumber(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9100__workeffort: %v", err)
		return err
	}
	cache.Put("9100__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9101__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9101__workeffort: %v", err)
		return err
	}
	cache.Put("9101__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9102__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9102__workeffort: %v", err)
		return err
	}
	cache.Put("9102__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9103__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:29.453")).
		SetWorkEffortName("task2").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9103__workeffort: %v", err)
		return err
	}
	cache.Put("9103__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9104__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:35.939")).
		SetWorkEffortName("phase2").
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9104__workeffort: %v", err)
		return err
	}
	cache.Put("9104__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9105__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:50.84")).
		SetWorkEffortName("task3").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-16 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-17 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9105__workeffort: %v", err)
		return err
	}
	cache.Put("9105__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9106__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:58.857")).
		SetWorkEffortName("task4").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-18 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-20 00:00:00.0")).
		SetRevisionNumber(1).
		SetSequenceNum(2).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9106__workeffort: %v", err)
		return err
	}
	cache.Put("9106__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9200__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 15:07:52.911")).
		SetWorkEffortName("Demo Project3 DemoCustomerCompany").
		SetRevisionNumber(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9200__workeffort: %v", err)
		return err
	}
	cache.Put("9200__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9201__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:14.226")).
		SetWorkEffortName("phase1").
		SetRevisionNumber(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9201__workeffort: %v", err)
		return err
	}
	cache.Put("9201__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("9202__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2007-12-14 16:45:21.831")).
		SetWorkEffortName("task1").
		SetEstimatedStartDate(common.ParseDateTime("2007-12-03 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2007-12-05 00:00:00.0")).
		SetRevisionNumber(1).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9202__workeffort: %v", err)
		return err
	}
	cache.Put("9202__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("staff_mtg__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("Staff Meeting").
		SetDescription("Staff Meeting").
		SetEstimatedStartDate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetEstimatedMilliSeconds(3600000).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__workeffort: %v", err)
		return err
	}
	cache.Put("staff_mtg__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("calendar_pub_demo__workeffort").
		SetWorkEffortName("iCalendar Publish Demonstration").
		SetDescription("Demo Project 1 Customer 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create calendar_pub_demo__workeffort: %v", err)
		return err
	}
	cache.Put("calendar_pub_demo__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("project_pub_demo__workeffort").
		SetWorkEffortName("Project iCalendar Publish Demonstration").
		SetDescription("Demo Project 1 Customer 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create project_pub_demo__workeffort: %v", err)
		return err
	}
	cache.Put("project_pub_demo__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("oneoffmeeting__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("One off Meeting june 12").
		SetDescription("Staff Meeting").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-12 09:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-12 10:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create oneoffmeeting__workeffort: %v", err)
		return err
	}
	cache.Put("oneoffmeeting__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("privatedemoemployee1__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("My Wifes birthday june 29").
		SetDescription("The birthday i should never forget").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-29 00:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-29 23:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create privatedemoemployee1__workeffort: %v", err)
		return err
	}
	cache.Put("privatedemoemployee1__workeffort", c)

	c, err = client.WorkEffort.Create().SetStringRef("publicevent__workeffort").
		SetLastStatusUpdate(common.ParseDateTime("2008-01-01 00:00:00.0")).
		SetWorkEffortName("The general company party june 17").
		SetDescription("General Party").
		SetLocationDesc("Tom's Banquet Hall").
		SetEstimatedStartDate(common.ParseDateTime("2009-06-17 19:00:00.0")).
		SetEstimatedCompletionDate(common.ParseDateTime("2009-06-17 23:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create publicevent__workeffort: %v", err)
		return err
	}
	cache.Put("publicevent__workeffort", c)

	return nil
}
