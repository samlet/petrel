package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateTemporalExpressionAssoc(ctx context.Context) error {
	log.Println("TemporalExpressionAssoc creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TemporalExpressionAssoc

	c, err = client.TemporalExpressionAssoc.Create().SetStringRef("staff_mtg__1st_monday_in_month__temporalexpressionassoc").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__1st_monday_in_month__temporalexpressionassoc: %v", err)
		return err
	}
	cache.Put("staff_mtg__1st_monday_in_month__temporalexpressionassoc", c)

	c, err = client.TemporalExpressionAssoc.Create().SetStringRef("staff_mtg__minute_00__temporalexpressionassoc").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__minute_00__temporalexpressionassoc: %v", err)
		return err
	}
	cache.Put("staff_mtg__minute_00__temporalexpressionassoc", c)

	c, err = client.TemporalExpressionAssoc.Create().SetStringRef("staff_mtg__hour_10__temporalexpressionassoc").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__hour_10__temporalexpressionassoc: %v", err)
		return err
	}
	cache.Put("staff_mtg__hour_10__temporalexpressionassoc", c)

	return nil
}
