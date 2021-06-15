package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateTemporalExpression(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TemporalExpression

	c = cache.Get("staff_mtg__temporalexpression").(*ent.TemporalExpression)
	c, err = c.Update().
		SetTempExprTypeID(common.ParseId("INTERSECTION")).
		SetDescription("First Monday of Each Month").
		AddFromTemporalExpressionAssocs(cache.Get("staff_mtg__1st_monday_in_month__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)).
		AddFromTemporalExpressionAssocs(cache.Get("staff_mtg__minute_00__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)).
		AddFromTemporalExpressionAssocs(cache.Get("staff_mtg__hour_10__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)).
		AddWorkEfforts(cache.Get("staff_mtg__workeffort").(*ent.WorkEffort)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create staff_mtg__temporalexpression: %v", err)
		return err
	}
	cache.Put("staff_mtg__temporalexpression", c)

	return nil
}
