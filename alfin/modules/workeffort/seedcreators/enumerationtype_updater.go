package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateEnumerationType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.EnumerationType

	c = cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Employment Status").
		AddEnumerations(cache.Get("emps_fulltime__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_parttime__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_self__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_house__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_retired__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_student__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("emps_unemp__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create employ_stts__enumerationtype: %v", err)
		return err
	}
	cache.Put("employ_stts__enumerationtype", c)

	c = cache.Get("marital_status__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetDescription("Marital Status").
		AddEnumerations(cache.Get("single__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("married__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marital_status__enumerationtype: %v", err)
		return err
	}
	cache.Put("marital_status__enumerationtype", c)

	c = cache.Get("pty_resid_stts__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Residence Status").
		AddEnumerations(cache.Get("press_own__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("press_pvt_tenant__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("press_pub_tenant__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("press_parents__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pty_resid_stts__enumerationtype: %v", err)
		return err
	}
	cache.Put("pty_resid_stts__enumerationtype", c)

	c = cache.Get("party_email__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Party Email Notification").
		AddEnumerations(cache.Get("party_regis_confirm__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("upd_prsnl_inf_cnfrm__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_email_verify__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_partyinv_email__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("cont_noti_email__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("prds_cust_activated__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("unsub_cont_list_noti__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("sub_cont_list_noti__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("unsub_cont_list_veri__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("cont_email_template__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_email__enumerationtype: %v", err)
		return err
	}
	cache.Put("party_email__enumerationtype", c)

	c = cache.Get("user_pref_groups__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("User preference groups").
		AddChildren(cache.Get("global_preferences__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create user_pref_groups__enumerationtype: %v", err)
		return err
	}
	cache.Put("user_pref_groups__enumerationtype", c)

	c = cache.Get("global_preferences__enumerationtype").(*ent.EnumerationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Global preferences").
		SetParent(cache.Get("user_pref_groups__enumerationtype").(*ent.EnumerationType)).
		AddEnumerations(cache.Get("organization_party__enumeration").(*ent.Enumeration)).
		AddEnumerations(cache.Get("visual_theme__enumeration").(*ent.Enumeration)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create global_preferences__enumerationtype: %v", err)
		return err
	}
	cache.Put("global_preferences__enumerationtype", c)

	return nil
}
