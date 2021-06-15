package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateCommunicationEventPrpTyp(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CommunicationEventPrpTyp

	c = cache.Get("activity_request__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Activity Request").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create activity_request__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("activity_request__communicationeventprptyp", c)

	c = cache.Get("conference__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Conference").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create conference__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("conference__communicationeventprptyp", c)

	c = cache.Get("customer_service_cal__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Customer Service Call").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create customer_service_cal__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("customer_service_cal__communicationeventprptyp", c)

	c = cache.Get("inquiry__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Inquiry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create inquiry__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("inquiry__communicationeventprptyp", c)

	c = cache.Get("meeting__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Meeting").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create meeting__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("meeting__communicationeventprptyp", c)

	c = cache.Get("sales_follow_up__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Sales Follow Up").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sales_follow_up__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("sales_follow_up__communicationeventprptyp", c)

	c = cache.Get("seminar__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Seminar").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create seminar__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("seminar__communicationeventprptyp", c)

	c = cache.Get("support_call__communicationeventprptyp").(*ent.CommunicationEventPrpTyp)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Support Call").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create support_call__communicationeventprptyp: %v", err)
		return err
	}
	cache.Put("support_call__communicationeventprptyp", c)

	return nil
}
