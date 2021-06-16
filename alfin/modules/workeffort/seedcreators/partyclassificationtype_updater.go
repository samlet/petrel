package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdatePartyClassificationType(ctx context.Context) error {
	log.Println("PartyClassificationType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyClassificationType
	failures := 0

	c = cache.Get("eeoc_classification__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("EEOC (Equal Opportunity)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update eeoc_classification__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("eeoc_classification__partyclassificationtype", c)
	}

	c = cache.Get("minority_classificat__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Minority").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update minority_classificat__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("minority_classificat__partyclassificationtype", c)
	}

	c = cache.Get("income_classificatio__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Income").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update income_classificatio__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("income_classificatio__partyclassificationtype", c)
	}

	c = cache.Get("industry_classificat__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Industry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update industry_classificat__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("industry_classificat__partyclassificationtype", c)
	}

	c = cache.Get("organization_classif__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update organization_classif__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("organization_classif__partyclassificationtype", c)
	}

	c = cache.Get("person_classificatio__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Person").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update person_classificatio__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("person_classificatio__partyclassificationtype", c)
	}

	c = cache.Get("size_classification__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Size").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update size_classification__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("size_classification__partyclassificationtype", c)
	}

	c = cache.Get("trade_classification__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Trade").
		AddChildren(cache.Get("trade_whole_classifi__partyclassificationtype").(*ent.PartyClassificationType)).
		AddChildren(cache.Get("trade_retail_classif__partyclassificationtype").(*ent.PartyClassificationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update trade_classification__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("trade_classification__partyclassificationtype", c)
	}

	c = cache.Get("trade_whole_classifi__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Wholesale").
		SetParent(cache.Get("trade_classification__partyclassificationtype").(*ent.PartyClassificationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update trade_whole_classifi__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("trade_whole_classifi__partyclassificationtype", c)
	}

	c = cache.Get("trade_retail_classif__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Retail").
		SetParent(cache.Get("trade_classification__partyclassificationtype").(*ent.PartyClassificationType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update trade_retail_classif__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("trade_retail_classif__partyclassificationtype", c)
	}

	c = cache.Get("annual_revenue__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Annual Revenue").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update annual_revenue__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("annual_revenue__partyclassificationtype", c)
	}

	c = cache.Get("number_of_employees__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Number of Employees").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update number_of_employees__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("number_of_employees__partyclassificationtype", c)
	}

	c = cache.Get("value_rating__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Value Rating").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update value_rating__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("value_rating__partyclassificationtype", c)
	}

	c = cache.Get("sic_code__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("SIC Code").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update sic_code__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("sic_code__partyclassificationtype", c)
	}

	c = cache.Get("ownership__partyclassificationtype").(*ent.PartyClassificationType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Ownership").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ownership__partyclassificationtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ownership__partyclassificationtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
