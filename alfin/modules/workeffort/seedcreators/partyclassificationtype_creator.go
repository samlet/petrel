package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyClassificationType(ctx context.Context) error {
	log.Println("PartyClassificationType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyClassificationType

	c, err = client.PartyClassificationType.Create().SetStringRef("eeoc_classification__partyclassificationtype").
		SetHasTable("No").
		SetDescription("EEOC (Equal Opportunity)").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create eeoc_classification__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("eeoc_classification__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("minority_classificat__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Minority").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create minority_classificat__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("minority_classificat__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("income_classificatio__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Income").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create income_classificatio__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("income_classificatio__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("industry_classificat__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Industry").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create industry_classificat__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("industry_classificat__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("organization_classif__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Organization").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create organization_classif__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("organization_classif__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("person_classificatio__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Person").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create person_classificatio__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("person_classificatio__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("size_classification__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Size").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create size_classification__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("size_classification__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("trade_classification__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Trade").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create trade_classification__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("trade_classification__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("trade_whole_classifi__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Wholesale").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create trade_whole_classifi__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("trade_whole_classifi__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("trade_retail_classif__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Retail").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create trade_retail_classif__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("trade_retail_classif__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("annual_revenue__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Annual Revenue").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create annual_revenue__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("annual_revenue__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("number_of_employees__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Number of Employees").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create number_of_employees__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("number_of_employees__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("value_rating__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Value Rating").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create value_rating__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("value_rating__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("sic_code__partyclassificationtype").
		SetHasTable("No").
		SetDescription("SIC Code").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create sic_code__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("sic_code__partyclassificationtype", c)

	c, err = client.PartyClassificationType.Create().SetStringRef("ownership__partyclassificationtype").
		SetHasTable("No").
		SetDescription("Ownership").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ownership__partyclassificationtype: %v", err)
		return err
	}
	cache.Put("ownership__partyclassificationtype", c)

	return nil
}
