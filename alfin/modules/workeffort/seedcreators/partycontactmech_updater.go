package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyContactMech(ctx context.Context) error {
	log.Println("PartyContactMech updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyContactMech

	c = cache.Get("demoemployee__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetPartyRole(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer1__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer1__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer1__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer2__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer2__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer2__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__9023__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer3__9020__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9020")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__9020__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__9020__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer3__9001__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9001")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__9001__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__9001__946720908__partycontactmech", c)
	}

	c = cache.Get("democustomer3__9023__946720908__partycontactmech").(*ent.PartyContactMech)
	c, err = c.Update().
		SetContactMechID(common.ParseId("9023")).
		SetFromDate(common.ParseDateTime("2000-01-01 10:01:48.933")).
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetPartyRole(cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__9023__946720908__partycontactmech: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__9023__946720908__partycontactmech", c)
	}

	return nil
}
