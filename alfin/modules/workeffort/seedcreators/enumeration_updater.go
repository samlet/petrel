package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateEnumeration(ctx context.Context) error {
	log.Println("Enumeration updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Enumeration
	failures := 0

	c = cache.Get("emps_fulltime__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("FULLTIME").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Full-time Employed").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_fulltime__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_fulltime__enumeration", c)
	}

	c = cache.Get("emps_parttime__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARTTIME").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Part-time Employed").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_parttime__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_parttime__enumeration", c)
	}

	c = cache.Get("emps_self__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SELF").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Self Employed").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_self__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_self__enumeration", c)
	}

	c = cache.Get("emps_house__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("HOUSE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("House-Person").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_house__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_house__enumeration", c)
	}

	c = cache.Get("emps_retired__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("RETIRED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Retired").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_retired__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_retired__enumeration", c)
	}

	c = cache.Get("emps_student__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("STUDENT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Student").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_student__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_student__enumeration", c)
	}

	c = cache.Get("emps_unemp__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("UNEMP").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Unemployed").
		SetEnumerationType(cache.Get("employ_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update emps_unemp__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("emps_unemp__enumeration", c)
	}

	c = cache.Get("single__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetSequenceID(common.ParseId("01")).
		SetDescription("Single").
		SetEnumerationType(cache.Get("marital_status__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update single__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("single__enumeration", c)
	}

	c = cache.Get("married__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetSequenceID(common.ParseId("02")).
		SetDescription("Married").
		SetEnumerationType(cache.Get("marital_status__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update married__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("married__enumeration", c)
	}

	c = cache.Get("press_own__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("OWN").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Own Home").
		SetEnumerationType(cache.Get("pty_resid_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update press_own__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("press_own__enumeration", c)
	}

	c = cache.Get("press_pvt_tenant__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PVT_TENANT").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Private Tenant").
		SetEnumerationType(cache.Get("pty_resid_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update press_pvt_tenant__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("press_pvt_tenant__enumeration", c)
	}

	c = cache.Get("press_pub_tenant__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PUB_TENANT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Public Tenant").
		SetEnumerationType(cache.Get("pty_resid_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update press_pub_tenant__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("press_pub_tenant__enumeration", c)
	}

	c = cache.Get("press_parents__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("PARENTS").
		SetSequenceID(common.ParseId("04")).
		SetDescription("With Parents").
		SetEnumerationType(cache.Get("pty_resid_stts__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update press_parents__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("press_parents__enumeration", c)
	}

	c = cache.Get("party_regis_confirm__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("REGIS_CONFIRM").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Party Registration Confirmation").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update party_regis_confirm__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("party_regis_confirm__enumeration", c)
	}

	c = cache.Get("upd_prsnl_inf_cnfrm__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("UPDAT_CONFIRM").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Update Personal Info Confirmation").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update upd_prsnl_inf_cnfrm__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("upd_prsnl_inf_cnfrm__enumeration", c)
	}

	c = cache.Get("prds_email_verify__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("EMAIL_VERIFY").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Party Email Address Verification").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_email_verify__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_email_verify__enumeration", c)
	}

	c = cache.Get("prds_partyinv_email__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("INVITE_EMAIL").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Party Invitation").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_partyinv_email__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_partyinv_email__enumeration", c)
	}

	c = cache.Get("cont_noti_email__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("CONT_EMAIL").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Contact-Us Notification").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cont_noti_email__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cont_noti_email__enumeration", c)
	}

	c = cache.Get("prds_cust_activated__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("ACCOUNT_ACTIVATED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Account Activated Notification").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update prds_cust_activated__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("prds_cust_activated__enumeration", c)
	}

	c = cache.Get("unsub_cont_list_noti__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("UNSUB_CONT_EMAIL").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Unsubscribe Contact List Notification").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update unsub_cont_list_noti__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("unsub_cont_list_noti__enumeration", c)
	}

	c = cache.Get("sub_cont_list_noti__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("SUB_CONT_EMAIL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Subscribe Contact List Notification").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sub_cont_list_noti__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sub_cont_list_noti__enumeration", c)
	}

	c = cache.Get("unsub_cont_list_veri__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("UNSUB_CONT_VERIFY_EMAIL").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Unsubscribe Contact List Verify").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update unsub_cont_list_veri__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("unsub_cont_list_veri__enumeration", c)
	}

	c = cache.Get("cont_email_template__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetEnumCode("CONT_EMAIL_TEMPLATE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Contact List E-mail Template").
		SetEnumerationType(cache.Get("party_email__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update cont_email_template__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("cont_email_template__enumeration", c)
	}

	c = cache.Get("organization_party__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetSequenceID(common.ParseId("01")).
		SetDescription("Organization party").
		SetEnumerationType(cache.Get("global_preferences__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update organization_party__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("organization_party__enumeration", c)
	}

	c = cache.Get("visual_theme__enumeration").(*ent.Enumeration)
	c, err = c.Update().
		SetSequenceID(common.ParseId("02")).
		SetDescription("Visual Theme").
		SetEnumerationType(cache.Get("global_preferences__enumerationtype").(*ent.EnumerationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update visual_theme__enumeration: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("visual_theme__enumeration", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
