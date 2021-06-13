package workeffort

import (
	"context"
	"github.com/iancoleman/strcase"
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"github.com/samlet/petrel/alfin/modules/workeffort/helper"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestProject(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	ctx=ent.NewContext(ctx, client)
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := DoProject(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func DoProject(ctx context.Context, client *ent.Client) error {
	createWorkEffortTypes(ctx, client)
	createParties(ctx, client)

	workeff, err := client.WorkEffort.Create().
		SetWorkEffortType(WorkEffortTypeRef(ctx, client, "PROJECT")).
		SetWorkEffortName("Demo Project1 Cust1").
		Save(ctx)
	if err != nil {
		log.Fatalf("create workeffort fail: %v", err)
	}
	alfin.Pretty(workeff)

	return nil
}

func createParties(ctx context.Context, client *ent.Client) {
	client.PartyType.Create().SetStringRef("PERSON").SetDescription("Person").Save(ctx)
	client.StatusItem.Create().SetStringRef("PARTY_ENABLED").SetDescription("Party enabled").Save(ctx)

	client.Party.Create().SetStringRef("DemoEmployee").
		SetPartyType(helper.PartyTypeRef(ctx, "PERSON")).
		SetStatusItem(StatusItemRef(ctx,client,"PARTY_ENABLED")).
		Save(ctx)
}

func createWorkEffortTypes(ctx context.Context, client *ent.Client) {
	names := []string{"Project", "Task", "layla"}
	bulk := make([]*ent.WorkEffortTypeCreate, len(names))
	for i, name := range names {
		bulk[i] = client.WorkEffortType.Create().
			SetDescription(name).
			SetStringRef(strcase.ToScreamingSnake(name))
	}
	rs, err := client.WorkEffortType.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		log.Fatalf("create type fail: %v", err)
	}
	for _,r := range rs {
		log.Printf("created %d for ref %s", r.ID, r.StringRef)
	}
}

