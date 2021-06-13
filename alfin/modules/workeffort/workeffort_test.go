package workeffort

import (
	"context"
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestWorkEffort(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	workeff, err := client.WorkEffort.Create().
		SetDescription("first task").Save(ctx)
	if err != nil {
		return err
	}
	alfin.Pretty(workeff)

	return nil
}
