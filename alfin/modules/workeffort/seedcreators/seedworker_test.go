package seedcreators

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
	"os"
	"testing"
)

func TestLoadSeedData(t *testing.T) {
	LoadSeedData()
}

func TestSchemas(t *testing.T) {
	client, err := ent.Open("sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.WriteTo(ctx, os.Stdout)
	if err != nil {
		log.Fatalf("write schemas fail: %v", err)
	}
}
