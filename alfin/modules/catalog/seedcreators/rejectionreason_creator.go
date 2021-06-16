package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/catalog/ent"
	"log"
)

func CreateRejectionReason(ctx context.Context) error {
	log.Println("RejectionReason creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.RejectionReason

	c, err = client.RejectionReason.Create().SetStringRef("srj_damaged__rejectionreason").
		SetDescription("Damaged").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create srj_damaged__rejectionreason: %v", err)
		return err
	}
	cache.Put("srj_damaged__rejectionreason", c)

	c, err = client.RejectionReason.Create().SetStringRef("srj_not_ordered__rejectionreason").
		SetDescription("Not Ordered").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create srj_not_ordered__rejectionreason: %v", err)
		return err
	}
	cache.Put("srj_not_ordered__rejectionreason", c)

	c, err = client.RejectionReason.Create().SetStringRef("srj_over_shipped__rejectionreason").
		SetDescription("Over Shipped").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create srj_over_shipped__rejectionreason: %v", err)
		return err
	}
	cache.Put("srj_over_shipped__rejectionreason", c)

	return nil
}
