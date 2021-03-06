package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateTemporalExpressionAssoc(ctx context.Context) error {
	log.Println("TemporalExpressionAssoc updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.TemporalExpressionAssoc
	failures := 0

	c = cache.Get("staff_mtg__1st_monday_in_month__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)
	c, err = c.Update().
		SetFromTemporalExpression(cache.Get("staff_mtg__temporalexpression").(*ent.TemporalExpression)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__1st_monday_in_month__temporalexpressionassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__1st_monday_in_month__temporalexpressionassoc", c)
	}

	c = cache.Get("staff_mtg__minute_00__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)
	c, err = c.Update().
		SetFromTemporalExpression(cache.Get("staff_mtg__temporalexpression").(*ent.TemporalExpression)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__minute_00__temporalexpressionassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__minute_00__temporalexpressionassoc", c)
	}

	c = cache.Get("staff_mtg__hour_10__temporalexpressionassoc").(*ent.TemporalExpressionAssoc)
	c, err = c.Update().
		SetFromTemporalExpression(cache.Get("staff_mtg__temporalexpression").(*ent.TemporalExpression)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update staff_mtg__hour_10__temporalexpressionassoc: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("staff_mtg__hour_10__temporalexpressionassoc", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
