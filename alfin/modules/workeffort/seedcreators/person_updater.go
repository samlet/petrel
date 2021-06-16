package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdatePerson(ctx context.Context) error {
	log.Println("Person updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Person
	failures := 0

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
		log.Printf("fail to update demoemployee__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee__person", c)
	}

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
		log.Printf("fail to update demoemployee1__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee1__person", c)
	}

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
		log.Printf("fail to update demoemployee2__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee2__person", c)
	}

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
		log.Printf("fail to update demoemployee3__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee3__person", c)
	}

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
		log.Printf("fail to update democustomer1__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer1__person", c)
	}

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
		log.Printf("fail to update democustomer2__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer2__person", c)
	}

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
		log.Printf("fail to update democustomer3__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer3__person", c)
	}

	c = cache.Get("workeffortuser__person").(*ent.Person)
	c, err = c.Update().
		SetFirstName("WorkEffort").
		SetLastName("User").
		SetParty(cache.Get("workeffortuser__party").(*ent.Party)).
		AddUserLogins(cache.Get("workeffortuser__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffortuser__person: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("workeffortuser__person", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
