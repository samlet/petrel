package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdatePartyContentType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyContentType

	c = cache.Get("internal__partycontenttype").(*ent.PartyContentType)
	c, err = c.Update().
		SetDescription("Internal Content").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create internal__partycontenttype: %v", err)
		return err
	}
	cache.Put("internal__partycontenttype", c)

	c = cache.Get("userdef__partycontenttype").(*ent.PartyContentType)
	c, err = c.Update().
		SetDescription("User Defined Content").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create userdef__partycontenttype: %v", err)
		return err
	}
	cache.Put("userdef__partycontenttype", c)

	c = cache.Get("lgoimgurl__partycontenttype").(*ent.PartyContentType)
	c, err = c.Update().
		SetDescription("Logo Image URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lgoimgurl__partycontenttype: %v", err)
		return err
	}
	cache.Put("lgoimgurl__partycontenttype", c)

	c = cache.Get("vndshpinf__partycontenttype").(*ent.PartyContentType)
	c, err = c.Update().
		SetDescription("Vendor Shipping Info").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vndshpinf__partycontenttype: %v", err)
		return err
	}
	cache.Put("vndshpinf__partycontenttype", c)

	return nil
}
