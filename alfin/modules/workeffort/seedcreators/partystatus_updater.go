package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyStatus(ctx context.Context) error {
	log.Println("PartyStatus updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyStatus

	c = cache.Get("party_enabled__demoemployee__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__demoemployee__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__demoemployee__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__demoemployee1__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__demoemployee1__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__demoemployee1__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__demoemployee2__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__demoemployee2__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__demoemployee2__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__demoemployee3__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__demoemployee3__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__demoemployee3__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__democustomer1__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__democustomer1__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__democustomer1__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__democustomer2__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__democustomer2__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__democustomer2__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__democustomer3__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__democustomer3__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__democustomer3__978350400__partystatus", c)
	}

	c = cache.Get("party_enabled__workeffortuser__978350400__partystatus").(*ent.PartyStatus)
	c, err = c.Update().
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetParty(cache.Get("workeffortuser__party").(*ent.Party)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__workeffortuser__978350400__partystatus: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("party_enabled__workeffortuser__978350400__partystatus", c)
	}

	return nil
}
