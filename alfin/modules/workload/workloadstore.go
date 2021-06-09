package workload

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workload/ent"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type WorkloadStore struct {
	client *ent.Client
}

func NewDevStore(ctx context.Context) (*WorkloadStore, error){
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil,err
	}
	//defer client.Close()
	//ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil,err
	}
	return &WorkloadStore{client: client}, nil
}

func (t WorkloadStore) Client() *ent.Client {
	return t.client
}