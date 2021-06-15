package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyContentType(ctx context.Context) error {
	log.Println("PartyContentType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyContentType

	c, err = client.PartyContentType.Create().SetStringRef("internal__partycontenttype").
		SetDescription("Internal Content").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create internal__partycontenttype: %v", err)
		return err
	}
	cache.Put("internal__partycontenttype", c)

	c, err = client.PartyContentType.Create().SetStringRef("userdef__partycontenttype").
		SetDescription("User Defined Content").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create userdef__partycontenttype: %v", err)
		return err
	}
	cache.Put("userdef__partycontenttype", c)

	c, err = client.PartyContentType.Create().SetStringRef("lgoimgurl__partycontenttype").
		SetDescription("Logo Image URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create lgoimgurl__partycontenttype: %v", err)
		return err
	}
	cache.Put("lgoimgurl__partycontenttype", c)

	c, err = client.PartyContentType.Create().SetStringRef("vndshpinf__partycontenttype").
		SetDescription("Vendor Shipping Info").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create vndshpinf__partycontenttype: %v", err)
		return err
	}
	cache.Put("vndshpinf__partycontenttype", c)

	return nil
}
