package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateStatusValidChange(ctx context.Context) error {
	log.Println("StatusValidChange creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusValidChange

	c, err = client.StatusValidChange.Create().SetStringRef("party_enabled__party_disabled__statusvalidchange").
		SetTransitionName("Disable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_enabled__party_disabled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("party_enabled__party_disabled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("party_disabled__party_enabled__statusvalidchange").
		SetTransitionName("Re-Enable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_disabled__party_enabled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("party_disabled__party_enabled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_entered__com_pending__statusvalidchange").
		SetTransitionName("Set Pending, only visible to originator").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__com_pending__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_entered__com_pending__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_entered__com_in_progress__statusvalidchange").
		SetTransitionName("Set In Progress, waiting to be send").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__com_in_progress__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_entered__com_in_progress__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_entered__com_complete__statusvalidchange").
		SetTransitionName("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__com_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_entered__com_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_pending__com_entered__statusvalidchange").
		SetTransitionName("Entered,visible to all participants").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__com_entered__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_pending__com_entered__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_pending__com_in_progress__statusvalidchange").
		SetTransitionName("Set In Progress, waiting to be send").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__com_in_progress__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_pending__com_in_progress__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_in_progress__com_complete__statusvalidchange").
		SetTransitionName("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_in_progress__com_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_in_progress__com_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_in_progress__com_bounced__statusvalidchange").
		SetTransitionName("Bounced").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_in_progress__com_bounced__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_in_progress__com_bounced__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_complete__com_resolved__statusvalidchange").
		SetTransitionName("Resolve").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__com_resolved__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_complete__com_resolved__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_complete__com_referred__statusvalidchange").
		SetTransitionName("Refer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__com_referred__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_complete__com_referred__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_complete__com_bounced__statusvalidchange").
		SetTransitionName("Bounced").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__com_bounced__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_complete__com_bounced__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_unknown_party__com_complete__statusvalidchange").
		SetTransitionName("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_unknown_party__com_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_unknown_party__com_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_unknown_party__com_entered__statusvalidchange").
		SetTransitionName("Corrected").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_unknown_party__com_entered__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_unknown_party__com_entered__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_pending__com_complete__statusvalidchange").
		SetTransitionName("Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__com_complete__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_pending__com_complete__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_entered__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_entered__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_entered__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_pending__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_pending__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_pending__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_in_progress__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_in_progress__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_in_progress__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_complete__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_complete__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_complete__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_resolved__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_resolved__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_resolved__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_referred__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_referred__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_referred__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_unknown_party__com_cancelled__statusvalidchange").
		SetTransitionName("Cancel").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_unknown_party__com_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_unknown_party__com_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_role_created__com_role_read__statusvalidchange").
		SetTransitionName("Read").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_created__com_role_read__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_role_created__com_role_read__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_role_created__com_role_completed__statusvalidchange").
		SetTransitionName("Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_created__com_role_completed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_role_created__com_role_completed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("com_role_read__com_role_completed__statusvalidchange").
		SetTransitionName("Completed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create com_role_read__com_role_completed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("com_role_read__com_role_completed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyrel_created__partyrel_expired__statusvalidchange").
		SetTransitionName("Expired").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyrel_created__partyrel_expired__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyrel_created__partyrel_expired__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_sent__partyinv_pending__statusvalidchange").
		SetTransitionName("Pending").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__partyinv_pending__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_sent__partyinv_pending__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_sent__partyinv_accepted__statusvalidchange").
		SetTransitionName("Accepted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__partyinv_accepted__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_sent__partyinv_accepted__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_sent__partyinv_declined__statusvalidchange").
		SetTransitionName("Declined").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__partyinv_declined__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_sent__partyinv_declined__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_sent__partyinv_cancelled__statusvalidchange").
		SetTransitionName("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_sent__partyinv_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_sent__partyinv_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_pending__partyinv_accepted__statusvalidchange").
		SetTransitionName("Accepted").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_pending__partyinv_accepted__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_pending__partyinv_accepted__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("partyinv_pending__partyinv_cancelled__statusvalidchange").
		SetTransitionName("Cancelled").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create partyinv_pending__partyinv_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("partyinv_pending__partyinv_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pas_assigned__pas_completed__statusvalidchange").
		SetTransitionName("Assignment Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pas_assigned__pas_completed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pas_assigned__pas_completed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pts_created__pts_completed__statusvalidchange").
		SetTransitionName("Task Complete").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created__pts_completed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pts_created__pts_completed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pts_created__pts_on_hold__statusvalidchange").
		SetTransitionName("Task On Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created__pts_on_hold__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pts_created__pts_on_hold__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pts_created__pts_cancelled__statusvalidchange").
		SetTransitionName("Task On Hold").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_created__pts_cancelled__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pts_created__pts_cancelled__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("pts_on_hold__pts_created__statusvalidchange").
		SetTransitionName("Activated").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pts_on_hold__pts_created__statusvalidchange: %v", err)
		return err
	}
	cache.Put("pts_on_hold__pts_created__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("prj_active__prj_closed__statusvalidchange").
		SetTransitionName("Close project").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prj_active__prj_closed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("prj_active__prj_closed__statusvalidchange", c)

	c, err = client.StatusValidChange.Create().SetStringRef("_na___prj_closed__statusvalidchange").
		SetTransitionName("Close project").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___prj_closed__statusvalidchange: %v", err)
		return err
	}
	cache.Put("_na___prj_closed__statusvalidchange", c)

	return nil
}
