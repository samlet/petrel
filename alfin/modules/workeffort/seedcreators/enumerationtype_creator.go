package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateEnumerationType(ctx context.Context) error {
	log.Println("EnumerationType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.EnumerationType

	c, err = client.EnumerationType.Create().SetStringRef("employ_stts__enumerationtype").
		SetHasTable("No").
		SetDescription("Employment Status").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create employ_stts__enumerationtype: %v", err)
		return err
	}
	cache.Put("employ_stts__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("marital_status__enumerationtype").
		SetDescription("Marital Status").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marital_status__enumerationtype: %v", err)
		return err
	}
	cache.Put("marital_status__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("pty_resid_stts__enumerationtype").
		SetHasTable("No").
		SetDescription("Residence Status").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pty_resid_stts__enumerationtype: %v", err)
		return err
	}
	cache.Put("pty_resid_stts__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("party_email__enumerationtype").
		SetHasTable("No").
		SetDescription("Party Email Notification").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create party_email__enumerationtype: %v", err)
		return err
	}
	cache.Put("party_email__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("user_pref_groups__enumerationtype").
		SetHasTable("No").
		SetDescription("User preference groups").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create user_pref_groups__enumerationtype: %v", err)
		return err
	}
	cache.Put("user_pref_groups__enumerationtype", c)

	c, err = client.EnumerationType.Create().SetStringRef("global_preferences__enumerationtype").
		SetHasTable("No").
		SetDescription("Global preferences").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create global_preferences__enumerationtype: %v", err)
		return err
	}
	cache.Put("global_preferences__enumerationtype", c)

	return nil
}
