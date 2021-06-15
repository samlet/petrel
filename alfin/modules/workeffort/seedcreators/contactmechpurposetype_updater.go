package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateContactMechPurposeType(ctx context.Context) error {
	log.Println("updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechPurposeType

	c = cache.Get("shipping_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipping Destination Address").
		AddContactMechTypePurposes(cache.Get("postal_address__shipping_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipping_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("shipping_location__contactmechpurposetype", c)

	c = cache.Get("ship_orig_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipping Origin Address").
		AddContactMechTypePurposes(cache.Get("postal_address__ship_orig_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_orig_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("ship_orig_location__contactmechpurposetype", c)

	c = cache.Get("billing_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Billing (AP) Address").
		AddContactMechTypePurposes(cache.Get("postal_address__billing_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create billing_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("billing_location__contactmechpurposetype", c)

	c = cache.Get("payment_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (AR) Address").
		AddContactMechTypePurposes(cache.Get("postal_address__payment_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create payment_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("payment_location__contactmechpurposetype", c)

	c = cache.Get("general_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("General Correspondence Address").
		AddContactMechTypePurposes(cache.Get("postal_address__general_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create general_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("general_location__contactmechpurposetype", c)

	c = cache.Get("pur_ret_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Purchase Return Address").
		AddContactMechTypePurposes(cache.Get("postal_address__pur_ret_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pur_ret_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("pur_ret_location__contactmechpurposetype", c)

	c = cache.Get("primary_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Primary Address").
		AddContactMechTypePurposes(cache.Get("postal_address__primary_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_location__contactmechpurposetype", c)

	c = cache.Get("previous_location__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Previous Address").
		AddContactMechTypePurposes(cache.Get("postal_address__previous_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create previous_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("previous_location__contactmechpurposetype", c)

	c = cache.Get("phone_shipping__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipping Destination Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_shipping__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_shipping__contactmechpurposetype", c)

	c = cache.Get("phone_ship_orig__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipping Origin Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_ship_orig__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_ship_orig__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_ship_orig__contactmechpurposetype", c)

	c = cache.Get("phone_billing__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Billing (AP) Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_billing__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_billing__contactmechpurposetype", c)

	c = cache.Get("phone_did__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Direct Inward Dialing Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_did__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_did__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_did__contactmechpurposetype", c)

	c = cache.Get("phone_payment__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (AR) Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_payment__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_payment__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_payment__contactmechpurposetype", c)

	c = cache.Get("phone_home__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Home Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_home__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_home__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_home__contactmechpurposetype", c)

	c = cache.Get("phone_work__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Work Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_work__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_work__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_work__contactmechpurposetype", c)

	c = cache.Get("phone_work_sec__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Secondary Work Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_work_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_work_sec__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_work_sec__contactmechpurposetype", c)

	c = cache.Get("fax_number__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Fax Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_number__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_number__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_number__contactmechpurposetype", c)

	c = cache.Get("fax_number_sec__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Secondary Fax Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_number_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_number_sec__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_number_sec__contactmechpurposetype", c)

	c = cache.Get("fax_shipping__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Shipping Destination Fax Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_shipping__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_shipping__contactmechpurposetype", c)

	c = cache.Get("fax_billing__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Billing Destination Fax Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_billing__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_billing__contactmechpurposetype", c)

	c = cache.Get("phone_mobile__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Main Mobile Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_mobile__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_mobile__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_mobile__contactmechpurposetype", c)

	c = cache.Get("phone_assistant__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Assistant's Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_assistant__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_assistant__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_assistant__contactmechpurposetype", c)

	c = cache.Get("primary_phone__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Primary Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__primary_phone__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_phone__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_phone__contactmechpurposetype", c)

	c = cache.Get("phone_quick__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Quick Calls Phone Number").
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_quick__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_quick__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_quick__contactmechpurposetype", c)

	c = cache.Get("marketing_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Primary Marketing Email Address").
		AddContactMechTypePurposes(cache.Get("email_address__marketing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marketing_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("marketing_email__contactmechpurposetype", c)

	c = cache.Get("primary_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Primary Email Address").
		AddContactMechTypePurposes(cache.Get("email_address__primary_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_email__contactmechpurposetype", c)

	c = cache.Get("billing_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Billing (AP) Email").
		AddContactMechTypePurposes(cache.Get("email_address__billing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create billing_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("billing_email__contactmechpurposetype", c)

	c = cache.Get("payment_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Payment (AR) Email").
		AddContactMechTypePurposes(cache.Get("email_address__payment_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create payment_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("payment_email__contactmechpurposetype", c)

	c = cache.Get("other_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Other Email Address").
		AddContactMechTypePurposes(cache.Get("email_address__other_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("other_email__contactmechpurposetype", c)

	c = cache.Get("support_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Support Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create support_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("support_email__contactmechpurposetype", c)

	c = cache.Get("order_email__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Order Notification Email Address").
		AddContactMechTypePurposes(cache.Get("email_address__order_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create order_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("order_email__contactmechpurposetype", c)

	c = cache.Get("primary_web_url__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Primary Website URL").
		AddContactMechTypePurposes(cache.Get("web_address__primary_web_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_web_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_web_url__contactmechpurposetype", c)

	c = cache.Get("twitter_url__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Twitter Website URL").
		AddContactMechTypePurposes(cache.Get("web_address__twitter_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create twitter_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("twitter_url__contactmechpurposetype", c)

	c = cache.Get("facebook_url__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Facebook Website URL").
		AddContactMechTypePurposes(cache.Get("web_address__facebook_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create facebook_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("facebook_url__contactmechpurposetype", c)

	c = cache.Get("linkedin_url__contactmechpurposetype").(*ent.ContactMechPurposeType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("LinkedIn Website URL").
		AddContactMechTypePurposes(cache.Get("web_address__linkedin_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to create linkedin_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("linkedin_url__contactmechpurposetype", c)

	return nil
}
