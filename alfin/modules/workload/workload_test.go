package workload

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workload/ent"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestWorkload(t *testing.T) {
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
	item,err:=client.WorkloadItem.Create().
		SetWorkloadItemSeqID(1).
		SetDescription("item1").Save(ctx)
	if err != nil {
		panic(err)
	}
	w,err:=client.Workload.Create().
		SetWorkloadName("tester").
		SetStatusID(0).
		AddWorkloadItems(item).
		Save(ctx)
	if err != nil {
		panic(err)
	}
	log.Printf("%s", w)
	return nil
}
