package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyStatus(ctx context.Context) error {
	log.Println("PartyStatus creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyStatus

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__demoemployee__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__demoemployee__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__demoemployee__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__demoemployee1__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__demoemployee1__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__demoemployee1__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__demoemployee2__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__demoemployee2__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__demoemployee2__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__demoemployee3__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__demoemployee3__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__demoemployee3__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__democustomer1__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__democustomer1__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__democustomer1__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__democustomer2__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__democustomer2__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__democustomer2__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__democustomer3__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__democustomer3__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__democustomer3__978350400__partystatus", c)

	c, err = client.PartyStatus.Create().SetStringRef("party_enabled__workeffortuser__978350400__partystatus").
		SetStatusDate(common.ParseDateTime("2001-01-01 12:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__workeffortuser__978350400__partystatus: %v", err)
		return err
	}
	cache.Put("party_enabled__workeffortuser__978350400__partystatus", c)

	return nil
}
