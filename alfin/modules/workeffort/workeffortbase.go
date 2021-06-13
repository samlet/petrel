package workeffort

import (
	"context"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
)

func IdRef(ctx context.Context, client *ent.Client, stringId string) *ent.WorkEffortType {
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

