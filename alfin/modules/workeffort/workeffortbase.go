package workeffort

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
	"log"
)

func WorkEffortTypeRef(ctx context.Context, client *ent.Client, stringId string) *ent.WorkEffortType {
	rec, err := client.WorkEffortType.
		Query().
		Where(workefforttype.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		panic(err)
	}
	//log.Println("returned: ", rec)
	return rec
}


func StatusItemRef(ctx context.Context, client *ent.Client, stringId string) *ent.StatusItem {
	rec, err := client.StatusItem.
		Query().
		Where(statusitem.StringRefEQ(stringId)).
		// `Only` fails if no record found,
		// or more than 1 record returned.
		Only(ctx)
	if err != nil {
		log.Fatalf("Cannot find status item %s: %v", stringId, err)
	}
	return rec
}

