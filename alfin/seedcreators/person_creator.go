package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePersons(ctx context.Context){
	client := ent.FromContext(ctx)
	cache:= cachecomp.FromContext(ctx)

	c, err := client.Person.Create().SetStringRef("democustomer3__person").
		SetFirstName("Billing").
		SetLastName("Customer 3").
		Save(ctx)
	if err != nil {
		log.Fatalf("fail to create democustomer3__person: %v", err)
	}
	cache.Put("democustomer3__person", c)

}

