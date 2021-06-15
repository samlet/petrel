package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateParty(ctx context.Context) error {
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Party

	c, err = client.Party.Create().SetStringRef("demoemployee__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__party: %v", err)
		return err
	}
	cache.Put("demoemployee__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee1__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__party: %v", err)
		return err
	}
	cache.Put("demoemployee1__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee2__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__party: %v", err)
		return err
	}
	cache.Put("demoemployee2__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee3__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__party: %v", err)
		return err
	}
	cache.Put("demoemployee3__party", c)

	c, err = client.Party.Create().SetStringRef("democustomer1__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__party: %v", err)
		return err
	}
	cache.Put("democustomer1__party", c)

	c, err = client.Party.Create().SetStringRef("democustomer2__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__party: %v", err)
		return err
	}
	cache.Put("democustomer2__party", c)

	c, err = client.Party.Create().SetStringRef("democustomer3__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__party: %v", err)
		return err
	}
	cache.Put("democustomer3__party", c)

	c, err = client.Party.Create().SetStringRef("workeffortuser__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__party: %v", err)
		return err
	}
	cache.Put("workeffortuser__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee1__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__party: %v", err)
		return err
	}
	cache.Put("demoemployee1__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee2__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__party: %v", err)
		return err
	}
	cache.Put("demoemployee2__party", c)

	c, err = client.Party.Create().SetStringRef("demoemployee3__party").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__party: %v", err)
		return err
	}
	cache.Put("demoemployee3__party", c)

	return nil
}
