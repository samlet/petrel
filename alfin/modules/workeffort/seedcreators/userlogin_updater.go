package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateUserLogin(ctx context.Context) error {
	log.Println("UserLogin updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserLogin
	failures := 0

	c = cache.Get("demoemployee__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee__userlogin", c)
	}

	c = cache.Get("demoemployee1__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee1__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee1__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee1__userlogin", c)
	}

	c = cache.Get("demoemployee2__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee2__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee2__userlogin", c)
	}

	c = cache.Get("demoemployee3__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update demoemployee3__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("demoemployee3__userlogin", c)
	}

	c = cache.Get("democustomer1__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer1__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer1__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer1__userlogin", c)
	}

	c = cache.Get("democustomer2__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer2__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer2__userlogin", c)
	}

	c = cache.Get("democustomer3__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update democustomer3__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("democustomer3__userlogin", c)
	}

	c = cache.Get("workeffortuser__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetCurrentPassword("{SHA}47b56994cbc2b6d10aa1be30f70165adb305a41a").
		SetParty(cache.Get("workeffortuser__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update workeffortuser__userlogin: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("workeffortuser__userlogin", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
