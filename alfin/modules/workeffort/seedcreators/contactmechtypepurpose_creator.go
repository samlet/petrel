package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateContactMechTypePurpose(ctx context.Context) error {
	log.Println("ContactMechTypePurpose creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechTypePurpose

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__billing_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__billing_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__billing_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__marketing_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__marketing_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__marketing_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__payment_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__payment_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__payment_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__order_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__order_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__order_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__other_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__other_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__other_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("email_address__primary_email__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__primary_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__primary_email__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__shipping_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__shipping_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__shipping_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__ship_orig_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__ship_orig_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__ship_orig_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__billing_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__billing_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__billing_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__payment_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__payment_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__payment_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__general_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__general_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__general_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__pur_ret_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__pur_ret_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__pur_ret_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__primary_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__primary_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__primary_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("postal_address__previous_location__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__previous_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__previous_location__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__fax_number__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_number__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_number__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__fax_number_sec__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_number_sec__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_number_sec__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_did__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_did__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_did__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_home__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_home__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_home__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_mobile__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_mobile__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_mobile__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_work__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_work__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_work__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_work_sec__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_work_sec__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_work_sec__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_shipping__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_shipping__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_shipping__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_ship_orig__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_ship_orig__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_ship_orig__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_billing__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_billing__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_billing__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_payment__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_payment__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_payment__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__primary_phone__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__primary_phone__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__primary_phone__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_quick__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_quick__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_quick__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__fax_shipping__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_shipping__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_shipping__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__fax_billing__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_billing__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_billing__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("telecom_number__phone_assistant__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_assistant__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_assistant__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("web_address__primary_web_url__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__primary_web_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__primary_web_url__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("web_address__twitter_url__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__twitter_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__twitter_url__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("web_address__facebook_url__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__facebook_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__facebook_url__contactmechtypepurpose", c)

	c, err = client.ContactMechTypePurpose.Create().SetStringRef("web_address__linkedin_url__contactmechtypepurpose").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__linkedin_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__linkedin_url__contactmechtypepurpose", c)

	return nil
}
