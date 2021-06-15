package alfin

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"github.com/stretchr/testify/assert"
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
	ctx = cachecomp.NewContext(ctx)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := createPersons(ctx); err != nil {
		log.Fatal(err)
	}

	// Do query
	cache := cachecomp.FromContext(ctx)
	person:=cache.Get("democustomer3__person").(*ent.Person)
	println(person.FirstName, person.LastName, person.BirthDate.String())

	person, err=person.Update().SetFirstName("Samlet").Save(ctx)
	if err != nil {
		log.Fatalf("update person fail: %v", err)
	}
	Pretty(person)

	err=updatePersons(ctx)
	if err != nil {
		log.Fatalf("update fail: %v", err)
	}
	person=cache.Get("democustomer3__person").(*ent.Person)
	assert.Equal(t, "Jim", person.FirstName)
	Pretty(person)
}

func createPersons(ctx context.Context) error {
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	c, err := client.Person.Create().SetStringRef("democustomer3__person").
		SetFirstName("Billing").
		SetLastName("Customer 3").
		SetBirthDate(ParseDateTime("2000-01-01 10:01:48.933")).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__person: %v", err)
		return err
	}
	cache.Put("democustomer3__person", c)

	return nil
}

func updatePersons(ctx context.Context) error {
	//client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.Person

	c=cache.Get("democustomer3__person").(*ent.Person)
	c, err=c.Update().
		SetFirstName("Jim").
		Save(ctx)
	if err != nil {
		log.Fatalf("update person fail: %v", err)
	}
	cache.Put("democustomer3__person", c)

	return nil
}

