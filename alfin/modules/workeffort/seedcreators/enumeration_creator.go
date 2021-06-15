package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateEnumeration(ctx context.Context) error {
	log.Println("Enumeration creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Enumeration

	c, err = client.Enumeration.Create().SetStringRef("emps_fulltime__enumeration").
		SetEnumCode("FULLTIME").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Full-time Employed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_fulltime__enumeration: %v", err)
		return err
	}
	cache.Put("emps_fulltime__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_parttime__enumeration").
		SetEnumCode("PARTTIME").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Part-time Employed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_parttime__enumeration: %v", err)
		return err
	}
	cache.Put("emps_parttime__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_self__enumeration").
		SetEnumCode("SELF").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Self Employed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_self__enumeration: %v", err)
		return err
	}
	cache.Put("emps_self__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_house__enumeration").
		SetEnumCode("HOUSE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("House-Person").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_house__enumeration: %v", err)
		return err
	}
	cache.Put("emps_house__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_retired__enumeration").
		SetEnumCode("RETIRED").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Retired").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_retired__enumeration: %v", err)
		return err
	}
	cache.Put("emps_retired__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_student__enumeration").
		SetEnumCode("STUDENT").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Student").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_student__enumeration: %v", err)
		return err
	}
	cache.Put("emps_student__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("emps_unemp__enumeration").
		SetEnumCode("UNEMP").
		SetSequenceID(common.ParseId("07")).
		SetDescription("Unemployed").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create emps_unemp__enumeration: %v", err)
		return err
	}
	cache.Put("emps_unemp__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("single__enumeration").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Single").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create single__enumeration: %v", err)
		return err
	}
	cache.Put("single__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("married__enumeration").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Married").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create married__enumeration: %v", err)
		return err
	}
	cache.Put("married__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("press_own__enumeration").
		SetEnumCode("OWN").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Own Home").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create press_own__enumeration: %v", err)
		return err
	}
	cache.Put("press_own__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("press_pvt_tenant__enumeration").
		SetEnumCode("PVT_TENANT").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Private Tenant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create press_pvt_tenant__enumeration: %v", err)
		return err
	}
	cache.Put("press_pvt_tenant__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("press_pub_tenant__enumeration").
		SetEnumCode("PUB_TENANT").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Public Tenant").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create press_pub_tenant__enumeration: %v", err)
		return err
	}
	cache.Put("press_pub_tenant__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("press_parents__enumeration").
		SetEnumCode("PARENTS").
		SetSequenceID(common.ParseId("04")).
		SetDescription("With Parents").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create press_parents__enumeration: %v", err)
		return err
	}
	cache.Put("press_parents__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("party_regis_confirm__enumeration").
		SetEnumCode("REGIS_CONFIRM").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Party Registration Confirmation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_regis_confirm__enumeration: %v", err)
		return err
	}
	cache.Put("party_regis_confirm__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("upd_prsnl_inf_cnfrm__enumeration").
		SetEnumCode("UPDAT_CONFIRM").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Update Personal Info Confirmation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create upd_prsnl_inf_cnfrm__enumeration: %v", err)
		return err
	}
	cache.Put("upd_prsnl_inf_cnfrm__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_email_verify__enumeration").
		SetEnumCode("EMAIL_VERIFY").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Party Email Address Verification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_email_verify__enumeration: %v", err)
		return err
	}
	cache.Put("prds_email_verify__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_partyinv_email__enumeration").
		SetEnumCode("INVITE_EMAIL").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Party Invitation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_partyinv_email__enumeration: %v", err)
		return err
	}
	cache.Put("prds_partyinv_email__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("cont_noti_email__enumeration").
		SetEnumCode("CONT_EMAIL").
		SetSequenceID(common.ParseId("05")).
		SetDescription("Contact-Us Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cont_noti_email__enumeration: %v", err)
		return err
	}
	cache.Put("cont_noti_email__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("prds_cust_activated__enumeration").
		SetEnumCode("ACCOUNT_ACTIVATED").
		SetSequenceID(common.ParseId("06")).
		SetDescription("Account Activated Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create prds_cust_activated__enumeration: %v", err)
		return err
	}
	cache.Put("prds_cust_activated__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("unsub_cont_list_noti__enumeration").
		SetEnumCode("UNSUB_CONT_EMAIL").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Unsubscribe Contact List Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create unsub_cont_list_noti__enumeration: %v", err)
		return err
	}
	cache.Put("unsub_cont_list_noti__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("sub_cont_list_noti__enumeration").
		SetEnumCode("SUB_CONT_EMAIL").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Subscribe Contact List Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sub_cont_list_noti__enumeration: %v", err)
		return err
	}
	cache.Put("sub_cont_list_noti__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("unsub_cont_list_veri__enumeration").
		SetEnumCode("UNSUB_CONT_VERIFY_EMAIL").
		SetSequenceID(common.ParseId("03")).
		SetDescription("Unsubscribe Contact List Verify").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create unsub_cont_list_veri__enumeration: %v", err)
		return err
	}
	cache.Put("unsub_cont_list_veri__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("cont_email_template__enumeration").
		SetEnumCode("CONT_EMAIL_TEMPLATE").
		SetSequenceID(common.ParseId("04")).
		SetDescription("Contact List E-mail Template").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create cont_email_template__enumeration: %v", err)
		return err
	}
	cache.Put("cont_email_template__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("organization_party__enumeration").
		SetSequenceID(common.ParseId("01")).
		SetDescription("Organization party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create organization_party__enumeration: %v", err)
		return err
	}
	cache.Put("organization_party__enumeration", c)

	c, err = client.Enumeration.Create().SetStringRef("visual_theme__enumeration").
		SetSequenceID(common.ParseId("02")).
		SetDescription("Visual Theme").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create visual_theme__enumeration: %v", err)
		return err
	}
	cache.Put("visual_theme__enumeration", c)

	return nil
}
