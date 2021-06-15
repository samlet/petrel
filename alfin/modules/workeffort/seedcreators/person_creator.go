package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePerson(ctx context.Context) error {
	log.Println("Person creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Person

	c, err = client.Person.Create().SetStringRef("demoemployee__person").
		SetFirstName("Demo").
		SetLastName("Employee").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__person: %v", err)
		return err
	}
	cache.Put("demoemployee__person", c)

	c, err = client.Person.Create().SetStringRef("demoemployee1__person").
		SetFirstName("Peter").
		SetLastName("Manager").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__person: %v", err)
		return err
	}
	cache.Put("demoemployee1__person", c)

	c, err = client.Person.Create().SetStringRef("demoemployee2__person").
		SetFirstName("Jo").
		SetLastName("Analist1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__person: %v", err)
		return err
	}
	cache.Put("demoemployee2__person", c)

	c, err = client.Person.Create().SetStringRef("demoemployee3__person").
		SetFirstName("Tom").
		SetLastName("Analist2").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__person: %v", err)
		return err
	}
	cache.Put("demoemployee3__person", c)

	c, err = client.Person.Create().SetStringRef("democustomer1__person").
		SetFirstName("ManagerP1").
		SetLastName("Customer 1").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__person: %v", err)
		return err
	}
	cache.Put("democustomer1__person", c)

	c, err = client.Person.Create().SetStringRef("democustomer2__person").
		SetFirstName("ManagerP2").
		SetLastName("Customer 2").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__person: %v", err)
		return err
	}
	cache.Put("democustomer2__person", c)

	c, err = client.Person.Create().SetStringRef("democustomer3__person").
		SetFirstName("Billing").
		SetLastName("Customer 3").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__person: %v", err)
		return err
	}
	cache.Put("democustomer3__person", c)

	c, err = client.Person.Create().SetStringRef("workeffortuser__person").
		SetFirstName("WorkEffort").
		SetLastName("User").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create workeffortuser__person: %v", err)
		return err
	}
	cache.Put("workeffortuser__person", c)

	return nil
}
