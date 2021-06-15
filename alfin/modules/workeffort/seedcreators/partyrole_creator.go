package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreatePartyRole(ctx context.Context) error {
	log.Println("creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.PartyRole

	c, err = client.PartyRole.Create().SetStringRef("demoemployee__employee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__employee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee__employee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee__provider_analyst__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__provider_analyst__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee__provider_analyst__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee1__employee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__employee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee1__employee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee1__provider_manager__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__provider_manager__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee1__provider_manager__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee1__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee1__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee2__provider_analyst__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__provider_analyst__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee2__provider_analyst__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee2__employee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__employee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee2__employee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee2__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee2__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee3__provider_analyst__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__provider_analyst__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee3__provider_analyst__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee3__employee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__employee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee3__employee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee3__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee3__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer1__customer__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__customer__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer1__customer__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer1__client_manager__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__client_manager__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer1__client_manager__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer1__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer1__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer1__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer2__customer__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__customer__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer2__customer__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer2__client_manager__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__client_manager__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer2__client_manager__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer2__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer2__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer2__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer3__customer__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__customer__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer3__customer__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer3__client_billing__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__client_billing__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer3__client_billing__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustomer3__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustomer3__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("democustomer3__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("admin__project_team__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__project_team__partyrole: %v", err)
		return err
	}
	cache.Put("admin__project_team__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("admin__provider_manager__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__provider_manager__partyrole: %v", err)
		return err
	}
	cache.Put("admin__provider_manager__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("democustcompany__client_billing__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create democustcompany__client_billing__partyrole: %v", err)
		return err
	}
	cache.Put("democustcompany__client_billing__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("admin__cal_owner__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__cal_owner__partyrole: %v", err)
		return err
	}
	cache.Put("admin__cal_owner__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("admin__cal_attendee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create admin__cal_attendee__partyrole: %v", err)
		return err
	}
	cache.Put("admin__cal_attendee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee1__cal_owner__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee1__cal_owner__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee1__cal_owner__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee2__cal_attendee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee2__cal_attendee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee2__cal_attendee__partyrole", c)

	c, err = client.PartyRole.Create().SetStringRef("demoemployee3__cal_attendee__partyrole").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create demoemployee3__cal_attendee__partyrole: %v", err)
		return err
	}
	cache.Put("demoemployee3__cal_attendee__partyrole", c)

	return nil
}
