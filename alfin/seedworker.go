package alfin

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"github.com/samlet/petrel/alfin/seedcreators"
	"log"
)

func LoadSeedData() {
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

	// Creators
	if err := seedcreators.CreatePerson(ctx); err != nil {
		log.Fatal(err)
	}
}


