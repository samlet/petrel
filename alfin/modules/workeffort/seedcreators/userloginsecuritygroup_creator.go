package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateUserLoginSecurityGroup(ctx context.Context) error {
	log.Println("UserLoginSecurityGroup creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.UserLoginSecurityGroup

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("demoemployee__projectadmin__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__projectadmin__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("demoemployee__projectadmin__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("demoemployee1__projectadmin__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__projectadmin__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("demoemployee1__projectadmin__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("demoemployee2__projectuser__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__projectuser__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("demoemployee2__projectuser__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("demoemployee3__projectuser__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__projectuser__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("demoemployee3__projectuser__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("democustomer1__projectuser__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__projectuser__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("democustomer1__projectuser__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("democustomer2__projectuser__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__projectuser__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("democustomer2__projectuser__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("democustomer3__projectuser__946684800__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2000-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__projectuser__946684800__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("democustomer3__projectuser__946684800__userloginsecuritygroup", c)

	c, err = client.UserLoginSecurityGroup.Create().SetStringRef("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup").
		SetFromDate(common.ParseDateTime("2011-01-01 00:00:00.0")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__workeffort_user__1293840000__userloginsecuritygroup: %v", err)
		return err
	}
	cache.Put("workeffortuser__workeffort_user__1293840000__userloginsecuritygroup", c)

	return nil
}
