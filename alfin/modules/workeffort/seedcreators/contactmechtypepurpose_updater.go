package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateContactMechTypePurpose(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechTypePurpose

	c = cache.Get("email_address__billing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("billing_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__billing_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__billing_email__contactmechtypepurpose", c)

	c = cache.Get("email_address__marketing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("marketing_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__marketing_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__marketing_email__contactmechtypepurpose", c)

	c = cache.Get("email_address__payment_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("payment_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__payment_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__payment_email__contactmechtypepurpose", c)

	c = cache.Get("email_address__order_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("order_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__order_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__order_email__contactmechtypepurpose", c)

	c = cache.Get("email_address__other_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("other_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__other_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__other_email__contactmechtypepurpose", c)

	c = cache.Get("email_address__primary_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("primary_email__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_address__primary_email__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("email_address__primary_email__contactmechtypepurpose", c)

	c = cache.Get("postal_address__shipping_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("shipping_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__shipping_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__shipping_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__ship_orig_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("ship_orig_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__ship_orig_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__ship_orig_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__billing_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("billing_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__billing_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__billing_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__payment_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("payment_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__payment_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__payment_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__general_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("general_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__general_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__general_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__pur_ret_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("pur_ret_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__pur_ret_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__pur_ret_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__primary_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("primary_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__primary_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__primary_location__contactmechtypepurpose", c)

	c = cache.Get("postal_address__previous_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("previous_location__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create postal_address__previous_location__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("postal_address__previous_location__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__fax_number__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("fax_number__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_number__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_number__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__fax_number_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("fax_number_sec__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_number_sec__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_number_sec__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_did__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_did__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_did__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_did__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_home__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_home__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_home__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_home__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_mobile__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_mobile__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_mobile__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_mobile__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_work__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_work__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_work__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_work__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_work_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_work_sec__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_work_sec__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_work_sec__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_shipping__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_shipping__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_shipping__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_ship_orig__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_ship_orig__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_ship_orig__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_ship_orig__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_billing__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_billing__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_billing__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_payment__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_payment__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_payment__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_payment__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__primary_phone__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("primary_phone__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__primary_phone__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__primary_phone__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_quick__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_quick__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_quick__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_quick__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__fax_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("fax_shipping__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_shipping__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_shipping__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__fax_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("fax_billing__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__fax_billing__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__fax_billing__contactmechtypepurpose", c)

	c = cache.Get("telecom_number__phone_assistant__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("phone_assistant__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create telecom_number__phone_assistant__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("telecom_number__phone_assistant__contactmechtypepurpose", c)

	c = cache.Get("web_address__primary_web_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("primary_web_url__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__primary_web_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__primary_web_url__contactmechtypepurpose", c)

	c = cache.Get("web_address__twitter_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("twitter_url__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__twitter_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__twitter_url__contactmechtypepurpose", c)

	c = cache.Get("web_address__facebook_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("facebook_url__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__facebook_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__facebook_url__contactmechtypepurpose", c)

	c = cache.Get("web_address__linkedin_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)
	c, err = c.Update().
		SetContactMechType(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		SetContactMechPurposeType(cache.Get("linkedin_url__contactmechpurposetype").(*ent.ContactMechPurposeType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_address__linkedin_url__contactmechtypepurpose: %v", err)
		return err
	}
	cache.Put("web_address__linkedin_url__contactmechtypepurpose", c)

	return nil
}
