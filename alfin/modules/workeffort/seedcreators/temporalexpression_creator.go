package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateTemporalExpression(ctx context.Context) error {
	log.Println("TemporalExpression creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TemporalExpression

	c, err = client.TemporalExpression.Create().SetStringRef("staff_mtg__temporalexpression").
		SetTempExprTypeID(common.ParseId("INTERSECTION")).
		SetDescription("First Monday of Each Month").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__temporalexpression: %v", err)
		return err
	}
	cache.Put("staff_mtg__temporalexpression", c)

	return nil
}
