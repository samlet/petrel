package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateCommunicationEventPrpTyp(ctx context.Context) error {
	log.Println("CommunicationEventPrpTyp updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CommunicationEventPrpTyp
	failures := 0

	c = cache.Get("activity_request__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Activity Request").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update activity_request__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("activity_request__communicationeventprptyp", c)
	}

	c = cache.Get("conference__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Conference").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update conference__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("conference__communicationeventprptyp", c)
	}

	c = cache.Get("customer_service_cal__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Customer Service Call").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update customer_service_cal__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("customer_service_cal__communicationeventprptyp", c)
	}

	c = cache.Get("inquiry__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inquiry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update inquiry__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("inquiry__communicationeventprptyp", c)
	}

	c = cache.Get("meeting__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meeting").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update meeting__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("meeting__communicationeventprptyp", c)
	}

	c = cache.Get("sales_follow_up__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Follow Up").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sales_follow_up__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sales_follow_up__communicationeventprptyp", c)
	}

	c = cache.Get("seminar__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Seminar").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update seminar__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("seminar__communicationeventprptyp", c)
	}

	c = cache.Get("support_call__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Support Call").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update support_call__communicationeventprptyp: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("support_call__communicationeventprptyp", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
