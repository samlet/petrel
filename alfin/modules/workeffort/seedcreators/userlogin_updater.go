package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateUserLogin(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserLogin

	c = cache.Get("demoemployee__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee__userlogin", c)

	c = cache.Get("demoemployee1__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee1__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee1__projectadmin__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee1__userlogin", c)

	c = cache.Get("demoemployee2__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee2__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee2__userlogin", c)

	c = cache.Get("demoemployee3__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		SetParty(cache.Get("demoemployee3__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("demoemployee3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee3__userlogin", c)

	c = cache.Get("democustomer1__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer1__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer1__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer1__userlogin", c)

	c = cache.Get("democustomer2__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer2__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer2__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer2__userlogin", c)

	c = cache.Get("democustomer3__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetParty(cache.Get("democustomer3__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("democustomer3__projectuser__946684800__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer3__userlogin", c)

	c = cache.Get("workeffortuser__userlogin").(*ent.UserLogin)
	c, err = c.Update().
		SetCurrentPassword("{SHA}47b56994cbc2b6d10aa1be30f70165adb305a41a").
		SetParty(cache.Get("workeffortuser__party").(*ent.Party)).
		AddUserLoginSecurityGroups(cache.Get("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup").(*ent.UserLoginSecurityGroup)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__userlogin: %v", err)
		return err
	}
	cache.Put("workeffortuser__userlogin", c)

	return nil
}
