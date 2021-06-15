package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateContactMechType(ctx context.Context) error {
	log.Println("ContactMechType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechType

	c, err = client.ContactMechType.Create().SetStringRef("electronic_address__contactmechtype").
		SetHasTable("No").
		SetDescription("Electronic Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create electronic_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("electronic_address__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("postal_address__contactmechtype").
		SetHasTable("Yes").
		SetDescription("Postal Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("postal_address__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("telecom_number__contactmechtype").
		SetHasTable("Yes").
		SetDescription("Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__contactmechtype: %v", err)
		return err
	}
	cache.Put("telecom_number__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("email_address__contactmechtype").
		SetHasTable("No").
		SetDescription("Email Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("email_address__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("ip_address__contactmechtype").
		SetHasTable("No").
		SetDescription("Internet IP Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ip_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("ip_address__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("domain_name__contactmechtype").
		SetHasTable("No").
		SetDescription("Internet Domain Name").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create domain_name__contactmechtype: %v", err)
		return err
	}
	cache.Put("domain_name__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("web_address__contactmechtype").
		SetHasTable("No").
		SetDescription("Web URL/Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("web_address__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("internal_partyid__contactmechtype").
		SetHasTable("No").
		SetDescription("Internal Note via partyId").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create internal_partyid__contactmechtype: %v", err)
		return err
	}
	cache.Put("internal_partyid__contactmechtype", c)

	c, err = client.ContactMechType.Create().SetStringRef("ftp_address__contactmechtype").
		SetHasTable("Yes").
		SetDescription("Ftp server connection").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ftp_address__contactmechtype: %v", err)
		return err
	}
	cache.Put("ftp_address__contactmechtype", c)

	return nil
}
