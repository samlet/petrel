package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateParty(ctx context.Context) error {
	log.Println("Party updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Party

	c = cache.Get("demoemployee__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee__project_team__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee__party", c)
	}

	c = cache.Get("demoemployee1__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee1__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee1__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__party", c)
	}

	c = cache.Get("demoemployee2__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee2__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee2__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee2__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__party", c)
	}

	c = cache.Get("demoemployee3__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee3__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee3__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__party", c)
	}

	c = cache.Get("democustomer1__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("democustomer1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("democustomer1__customer__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer1__client_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer1__project_team__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer1__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("democustomer1__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer1__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer1__party", c)
	}

	c = cache.Get("democustomer2__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("democustomer2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("democustomer2__customer__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer2__client_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer2__project_team__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer2__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("democustomer2__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer2__client_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer2__party", c)
	}

	c = cache.Get("democustomer3__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("democustomer3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("democustomer3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("democustomer3__customer__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer3__client_billing__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("democustomer3__project_team__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__democustomer3__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("democustomer3__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__democustomer3__client_billing__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("democustomer3__party", c)
	}

	c = cache.Get("workeffortuser__party").(*ent.Party)
	c, err = c.Update().
		SetPartyType(cache.Get("person__partytype").(*ent.PartyType)).
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyStatuses(cache.Get("party_enabled__workeffortuser__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("workeffortuser__userlogin").(*ent.UserLogin)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffortuser__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("workeffortuser__party", c)
	}

	c = cache.Get("demoemployee1__party").(*ent.Party)
	c, err = c.Update().
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee1__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee1__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee1__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__provider_manager__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee1__cal_owner__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee1__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee1__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee1__provider_manager__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("project_pub_demo__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("privatedemoemployee1__demoemployee1__cal_owner__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee1__party", c)
	}

	c = cache.Get("demoemployee2__party").(*ent.Party)
	c, err = c.Update().
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee2__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee2__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee2__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee2__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee2__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee2__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9000__demoemployee2__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("oneoffmeeting__demoemployee2__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee2__party", c)
	}

	c = cache.Get("demoemployee3__party").(*ent.Party)
	c, err = c.Update().
		SetStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		AddPartyContactMeches(cache.Get("demoemployee3__9020__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9001__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyContactMeches(cache.Get("demoemployee3__9023__946720908__partycontactmech").(*ent.PartyContactMech)).
		AddPartyRoles(cache.Get("demoemployee3__provider_analyst__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__employee__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__project_team__partyrole").(*ent.PartyRole)).
		AddPartyRoles(cache.Get("demoemployee3__cal_attendee__partyrole").(*ent.PartyRole)).
		AddPartyStatuses(cache.Get("party_enabled__demoemployee3__978350400__partystatus").(*ent.PartyStatus)).
		AddUserLogins(cache.Get("demoemployee3__userlogin").(*ent.UserLogin)).
		AddWorkEffortPartyAssignments(cache.Get("9100__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("9200__demoemployee3__provider_analyst__1197650721__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		AddWorkEffortPartyAssignments(cache.Get("staff_mtg__demoemployee3__cal_attendee__1199145600__workeffortpartyassignment").(*ent.WorkEffortPartyAssignment)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__party: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("demoemployee3__party", c)
	}

	return nil
}
