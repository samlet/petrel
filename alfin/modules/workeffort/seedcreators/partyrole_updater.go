package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyRole(ctx context.Context) error {
	log.Println("PartyRole updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyRole

	c = cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("demoemployee__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__employee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__employee__partyrole", c)
	}

	c = cache.Get("demoemployee__provider_analyst__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetRoleType(cache.Get("provider_analyst__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__provider_analyst__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__provider_analyst__partyrole", c)
	}

	c = cache.Get("demoemployee__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__project_team__partyrole", c)
	}

	c = cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__employee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__employee__partyrole", c)
	}

	c = cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetRoleType(cache.Get("provider_manager__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__provider_manager__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__provider_manager__partyrole", c)
	}

	c = cache.Get("demoemployee1__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__project_team__partyrole", c)
	}

	c = cache.Get("demoemployee2__provider_analyst__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetRoleType(cache.Get("provider_analyst__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__provider_analyst__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__provider_analyst__partyrole", c)
	}

	c = cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__employee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__employee__partyrole", c)
	}

	c = cache.Get("demoemployee2__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__project_team__partyrole", c)
	}

	c = cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetRoleType(cache.Get("provider_analyst__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__provider_analyst__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__provider_analyst__partyrole", c)
	}

	c = cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetRoleType(cache.Get("employee__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__employee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__employee__partyrole", c)
	}

	c = cache.Get("demoemployee3__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__project_team__partyrole", c)
	}

	c = cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("democustomer1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__customer__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__customer__partyrole", c)
	}

	c = cache.Get("democustomer1__client_manager__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetRoleType(cache.Get("client_manager__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__client_manager__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__client_manager__partyrole", c)
	}

	c = cache.Get("democustomer1__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__project_team__partyrole", c)
	}

	c = cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("democustomer2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__customer__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__customer__partyrole", c)
	}

	c = cache.Get("democustomer2__client_manager__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetRoleType(cache.Get("client_manager__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__client_manager__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__client_manager__partyrole", c)
	}

	c = cache.Get("democustomer2__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__project_team__partyrole", c)
	}

	c = cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetRoleType(cache.Get("customer__roletype").(*ent.RoleType)).
		AddPartyContactMeches(cache.Get("democustomer3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__customer__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__customer__partyrole", c)
	}

	c = cache.Get("democustomer3__client_billing__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetRoleType(cache.Get("client_billing__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__client_billing__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__client_billing__partyrole", c)
	}

	c = cache.Get("democustomer3__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__project_team__partyrole", c)
	}

	c = cache.Get("admin__project_team__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetRoleType(cache.Get("project_team__roletype").(*ent.RoleType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update admin__project_team__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("admin__project_team__partyrole", c)
	}

	c = cache.Get("admin__provider_manager__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetRoleType(cache.Get("provider_manager__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9000__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9002__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__admin__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update admin__provider_manager__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("admin__provider_manager__partyrole", c)
	}

	c = cache.Get("democustcompany__client_billing__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetRoleType(cache.Get("client_billing__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("9200__democustcompany__client_billing__1238963272__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustcompany__client_billing__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustcompany__client_billing__partyrole", c)
	}

	c = cache.Get("admin__cal_owner__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetRoleType(cache.Get("cal_owner__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("calendar_pub_demo__admin__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update admin__cal_owner__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("admin__cal_owner__partyrole", c)
	}

	c = cache.Get("admin__cal_attendee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetRoleType(cache.Get("cal_attendee__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__admin__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update admin__cal_attendee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("admin__cal_attendee__partyrole", c)
	}

	c = cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetRoleType(cache.Get("cal_owner__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__cal_owner__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__cal_owner__partyrole", c)
	}

	c = cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetRoleType(cache.Get("cal_attendee__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__cal_attendee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__cal_attendee__partyrole", c)
	}

	c = cache.Get("demoemployee3__cal_attendee__partyrole").(*ent.PartyRole)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetRoleType(cache.Get("cal_attendee__roletype").(*ent.RoleType)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__cal_attendee__partyrole: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__cal_attendee__partyrole", c)
	}

	return nil
}
