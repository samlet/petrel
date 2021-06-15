package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePerson(ctx context.Context) error {
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Person

	c = cache.Get("demoemployee__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("Demo").
		SetLastName("Employee").
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("demoemployee__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("demoemployee__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__person: %v", err)
		return err
	}
	cache.Put("demoemployee__person", c)

	c = cache.Get("demoemployee1__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("Peter").
		SetLastName("Manager").
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("demoemployee1__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__person: %v", err)
		return err
	}
	cache.Put("demoemployee1__person", c)

	c = cache.Get("demoemployee2__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("Jo").
		SetLastName("Analist1").
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("demoemployee2__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__person: %v", err)
		return err
	}
	cache.Put("demoemployee2__person", c)

	c = cache.Get("demoemployee3__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("Tom").
		SetLastName("Analist2").
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("demoemployee3__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__person: %v", err)
		return err
	}
	cache.Put("demoemployee3__person", c)

	c = cache.Get("democustomer1__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("ManagerP1").
		SetLastName("Customer 1").
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("democustomer1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("democustomer1__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__person: %v", err)
		return err
	}
	cache.Put("democustomer1__person", c)

	c = cache.Get("democustomer2__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("ManagerP2").
		SetLastName("Customer 2").
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("democustomer2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("democustomer2__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__person: %v", err)
		return err
	}
	cache.Put("democustomer2__person", c)

	c = cache.Get("democustomer3__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("Billing").
		SetLastName("Customer 3").
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		AddPartyContactMeches(cache.Get("democustomer3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddUserLogins(cache.Get("democustomer3__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__person: %v", err)
		return err
	}
	cache.Put("democustomer3__person", c)

	c = cache.Get("workeffortuser__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("WorkEffort").
		SetLastName("User").
		SetParty(cache.Get("workeffortuser__party").(*ent.Party)).
		AddUserLogins(cache.Get("workeffortuser__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__person: %v", err)
		return err
	}
	cache.Put("workeffortuser__person", c)

	return nil
}
