package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateStatusValidChange(ctx context.Context) error {
	log.Println("StatusValidChange updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.StatusValidChange
	failures := 0

	c = cache.Get("party_enabled__party_disabled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Disable").
		SetMainStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("party_disabled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_enabled__party_disabled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("party_enabled__party_disabled__statusvalidchange", c)
	}

	c = cache.Get("party_disabled__party_enabled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Re-Enable").
		SetMainStatusItem(cache.Get("party_disabled__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("party_enabled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_disabled__party_enabled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("party_disabled__party_enabled__statusvalidchange", c)
	}

	c = cache.Get("com_entered__com_pending__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Set Pending, only visible to originator").
		SetMainStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_entered__com_pending__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_entered__com_pending__statusvalidchange", c)
	}

	c = cache.Get("com_entered__com_in_progress__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Set In Progress, waiting to be send").
		SetMainStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_entered__com_in_progress__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_entered__com_in_progress__statusvalidchange", c)
	}

	c = cache.Get("com_entered__com_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Complete").
		SetMainStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_entered__com_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_entered__com_complete__statusvalidchange", c)
	}

	c = cache.Get("com_pending__com_entered__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Entered,visible to all participants").
		SetMainStatusItem(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_pending__com_entered__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_pending__com_entered__statusvalidchange", c)
	}

	c = cache.Get("com_pending__com_in_progress__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Set In Progress, waiting to be send").
		SetMainStatusItem(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_pending__com_in_progress__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_pending__com_in_progress__statusvalidchange", c)
	}

	c = cache.Get("com_in_progress__com_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Complete").
		SetMainStatusItem(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_in_progress__com_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_in_progress__com_complete__statusvalidchange", c)
	}

	c = cache.Get("com_in_progress__com_bounced__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Bounced").
		SetMainStatusItem(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_bounced__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_in_progress__com_bounced__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_in_progress__com_bounced__statusvalidchange", c)
	}

	c = cache.Get("com_complete__com_resolved__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Resolve").
		SetMainStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_resolved__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_complete__com_resolved__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_complete__com_resolved__statusvalidchange", c)
	}

	c = cache.Get("com_complete__com_referred__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Refer").
		SetMainStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_referred__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_complete__com_referred__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_complete__com_referred__statusvalidchange", c)
	}

	c = cache.Get("com_complete__com_bounced__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Bounced").
		SetMainStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_bounced__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_complete__com_bounced__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_complete__com_bounced__statusvalidchange", c)
	}

	c = cache.Get("com_unknown_party__com_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Complete").
		SetMainStatusItem(cache.Get("com_unknown_party__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_unknown_party__com_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_unknown_party__com_complete__statusvalidchange", c)
	}

	c = cache.Get("com_unknown_party__com_entered__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Corrected").
		SetMainStatusItem(cache.Get("com_unknown_party__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_unknown_party__com_entered__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_unknown_party__com_entered__statusvalidchange", c)
	}

	c = cache.Get("com_pending__com_complete__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Complete").
		SetMainStatusItem(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_pending__com_complete__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_pending__com_complete__statusvalidchange", c)
	}

	c = cache.Get("com_entered__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_entered__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_entered__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_entered__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_pending__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_pending__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_pending__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_in_progress__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_in_progress__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_in_progress__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_in_progress__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_complete__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_complete__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_complete__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_complete__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_resolved__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_resolved__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_resolved__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_resolved__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_referred__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_referred__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_referred__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_referred__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_unknown_party__com_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancel").
		SetMainStatusItem(cache.Get("com_unknown_party__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_unknown_party__com_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_unknown_party__com_cancelled__statusvalidchange", c)
	}

	c = cache.Get("com_role_created__com_role_read__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Read").
		SetMainStatusItem(cache.Get("com_role_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_role_read__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_role_created__com_role_read__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_role_created__com_role_read__statusvalidchange", c)
	}

	c = cache.Get("com_role_created__com_role_completed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Completed").
		SetMainStatusItem(cache.Get("com_role_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_role_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_role_created__com_role_completed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_role_created__com_role_completed__statusvalidchange", c)
	}

	c = cache.Get("com_role_read__com_role_completed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Completed").
		SetMainStatusItem(cache.Get("com_role_read__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("com_role_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update com_role_read__com_role_completed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("com_role_read__com_role_completed__statusvalidchange", c)
	}

	c = cache.Get("partyrel_created__partyrel_expired__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Expired").
		SetMainStatusItem(cache.Get("partyrel_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyrel_expired__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyrel_created__partyrel_expired__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyrel_created__partyrel_expired__statusvalidchange", c)
	}

	c = cache.Get("partyinv_sent__partyinv_pending__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Pending").
		SetMainStatusItem(cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_pending__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_sent__partyinv_pending__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_sent__partyinv_pending__statusvalidchange", c)
	}

	c = cache.Get("partyinv_sent__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Accepted").
		SetMainStatusItem(cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_accepted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_sent__partyinv_accepted__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_sent__partyinv_accepted__statusvalidchange", c)
	}

	c = cache.Get("partyinv_sent__partyinv_declined__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Declined").
		SetMainStatusItem(cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_declined__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_sent__partyinv_declined__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_sent__partyinv_declined__statusvalidchange", c)
	}

	c = cache.Get("partyinv_sent__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancelled").
		SetMainStatusItem(cache.Get("partyinv_sent__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_sent__partyinv_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_sent__partyinv_cancelled__statusvalidchange", c)
	}

	c = cache.Get("partyinv_pending__partyinv_accepted__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Accepted").
		SetMainStatusItem(cache.Get("partyinv_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_accepted__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_pending__partyinv_accepted__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_pending__partyinv_accepted__statusvalidchange", c)
	}

	c = cache.Get("partyinv_pending__partyinv_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Cancelled").
		SetMainStatusItem(cache.Get("partyinv_pending__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("partyinv_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update partyinv_pending__partyinv_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("partyinv_pending__partyinv_cancelled__statusvalidchange", c)
	}

	c = cache.Get("pas_assigned__pas_completed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Assignment Complete").
		SetMainStatusItem(cache.Get("pas_assigned__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pas_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pas_assigned__pas_completed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pas_assigned__pas_completed__statusvalidchange", c)
	}

	c = cache.Get("pts_created__pts_completed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Task Complete").
		SetMainStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pts_completed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pts_created__pts_completed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pts_created__pts_completed__statusvalidchange", c)
	}

	c = cache.Get("pts_created__pts_on_hold__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Task On Hold").
		SetMainStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pts_on_hold__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pts_created__pts_on_hold__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pts_created__pts_on_hold__statusvalidchange", c)
	}

	c = cache.Get("pts_created__pts_cancelled__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Task On Hold").
		SetMainStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pts_cancelled__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pts_created__pts_cancelled__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pts_created__pts_cancelled__statusvalidchange", c)
	}

	c = cache.Get("pts_on_hold__pts_created__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Activated").
		SetMainStatusItem(cache.Get("pts_on_hold__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("pts_created__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update pts_on_hold__pts_created__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("pts_on_hold__pts_created__statusvalidchange", c)
	}

	c = cache.Get("prj_active__prj_closed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Close project").
		SetMainStatusItem(cache.Get("prj_active__statusitem").(*ent.StatusItem)).
		SetToStatusItem(cache.Get("prj_closed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prj_active__prj_closed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prj_active__prj_closed__statusvalidchange", c)
	}

	c = cache.Get("_na___prj_closed__statusvalidchange").(*ent.StatusValidChange)
	c, err = c.Update().
		SetTransitionName("Close project").
		SetToStatusItem(cache.Get("prj_closed__statusitem").(*ent.StatusItem)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update _na___prj_closed__statusvalidchange: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("_na___prj_closed__statusvalidchange", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
