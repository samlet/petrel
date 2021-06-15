package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateUserLogin(ctx context.Context) error {
	log.Println("UserLogin creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserLogin

	c, err = client.UserLogin.Create().SetStringRef("demoemployee__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("demoemployee1__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee1__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("demoemployee2__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee2__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("demoemployee3__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__userlogin: %v", err)
		return err
	}
	cache.Put("demoemployee3__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("democustomer1__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer1__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("democustomer2__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer2__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("democustomer3__userlogin").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__userlogin: %v", err)
		return err
	}
	cache.Put("democustomer3__userlogin", c)

	c, err = client.UserLogin.Create().SetStringRef("workeffortuser__userlogin").
		SetCurrentPassword("{SHA}47b56994cbc2b6d10aa1be30f70165adb305a41a").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__userlogin: %v", err)
		return err
	}
	cache.Put("workeffortuser__userlogin", c)

	return nil
}
