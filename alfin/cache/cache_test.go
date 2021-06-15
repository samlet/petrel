package cache

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestEntSave(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	ctx = ent.NewContext(ctx, client)
	ctx = NewContext(ctx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := createPersons(ctx); err != nil {
		log.Fatal(err)
	}

	// Do query
	cache := FromContext(ctx)
	person:=cache.Get("democustomer3__person").(*ent.Person)
	println(person.FirstName, person.LastName)
}

func createPersons(ctx context.Context) error {
	client := ent.FromContext(ctx)
	cache := FromContext(ctx)

	c, err := client.Person.Create().SetStringRef("democustomer3__person").
		SetFirstName("Billing").
		SetLastName("Customer 3").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__person: %v", err)
		return err
	}
	cache.Put("democustomer3__person", c)

	return nil
}

