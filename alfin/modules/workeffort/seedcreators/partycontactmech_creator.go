package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyContactMech(ctx context.Context) error {
	log.Println("PartyContactMech creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyContactMech

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee1__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee1__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee1__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee1__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee1__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee1__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee2__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee2__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee2__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee2__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee2__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee2__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee3__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee3__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee3__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee3__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("demoemployee3__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("demoemployee3__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer1__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer1__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer1__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer1__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer1__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer1__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer2__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer2__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer2__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer2__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer2__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer2__9023__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer3__9020__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__9020__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer3__9020__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer3__9001__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__9001__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer3__9001__946720908__partycontactmech", c)

	c, err = client.PartyContactMech.Create().SetStringRef("democustomer3__9023__946720908__partycontactmech").
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__9023__946720908__partycontactmech: %v", err)
		return err
	}
	cache.Put("democustomer3__9023__946720908__partycontactmech", c)

	return nil
}
