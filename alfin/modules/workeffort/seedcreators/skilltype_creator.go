package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateSkillType(ctx context.Context) error {
	log.Println("SkillType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.SkillType

	c, err = client.SkillType.Create().SetStringRef("9000__skilltype").
		SetDescription("Java/Groovy/BSH").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9000__skilltype: %v", err)
		return err
	}
	cache.Put("9000__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("9001__skilltype").
		SetDescription("Mini Language").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9001__skilltype: %v", err)
		return err
	}
	cache.Put("9001__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("9002__skilltype").
		SetDescription("HTML/FTL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9002__skilltype: %v", err)
		return err
	}
	cache.Put("9002__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("9003__skilltype").
		SetDescription("Javascript/Dojo").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9003__skilltype: %v", err)
		return err
	}
	cache.Put("9003__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("9004__skilltype").
		SetDescription("Screens/Forms").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9004__skilltype: %v", err)
		return err
	}
	cache.Put("9004__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("9005__skilltype").
		SetDescription("OFBiz Installation").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create 9005__skilltype: %v", err)
		return err
	}
	cache.Put("9005__skilltype", c)

	c, err = client.SkillType.Create().SetStringRef("_na___skilltype").
		SetDescription("Not Applicable").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create _na___skilltype: %v", err)
		return err
	}
	cache.Put("_na___skilltype", c)

	return nil
}
