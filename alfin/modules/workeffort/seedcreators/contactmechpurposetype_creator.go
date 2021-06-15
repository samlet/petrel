package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateContactMechPurposeType(ctx context.Context) error {
	log.Println("ContactMechPurposeType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechPurposeType

	c, err = client.ContactMechPurposeType.Create().SetStringRef("shipping_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Shipping Destination Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create shipping_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("shipping_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("ship_orig_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Shipping Origin Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create ship_orig_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("ship_orig_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("billing_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Billing (AP) Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create billing_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("billing_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("payment_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Payment (AR) Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create payment_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("payment_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("general_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("General Correspondence Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create general_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("general_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("pur_ret_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Purchase Return Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create pur_ret_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("pur_ret_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("primary_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Primary Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("previous_location__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Previous Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create previous_location__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("previous_location__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_shipping__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Shipping Destination Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_shipping__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_shipping__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_ship_orig__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Shipping Origin Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_ship_orig__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_ship_orig__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_billing__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Billing (AP) Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_billing__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_billing__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_did__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Direct Inward Dialing Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_did__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_did__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_payment__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Payment (AR) Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_payment__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_payment__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_home__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Main Home Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_home__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_home__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_work__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Main Work Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_work__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_work__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_work_sec__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Secondary Work Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_work_sec__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_work_sec__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("fax_number__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Main Fax Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_number__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_number__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("fax_number_sec__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Secondary Fax Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_number_sec__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_number_sec__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("fax_shipping__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Shipping Destination Fax Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_shipping__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_shipping__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("fax_billing__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Billing Destination Fax Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_billing__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("fax_billing__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_mobile__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Main Mobile Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_mobile__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_mobile__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_assistant__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Assistant's Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_assistant__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_assistant__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("primary_phone__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Primary Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_phone__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_phone__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("phone_quick__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Quick Calls Phone Number").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_quick__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("phone_quick__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("marketing_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Primary Marketing Email Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create marketing_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("marketing_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("primary_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Primary Email Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("billing_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Billing (AP) Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create billing_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("billing_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("payment_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Payment (AR) Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create payment_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("payment_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("other_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Other Email Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create other_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("other_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("support_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Support Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create support_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("support_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("order_email__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Order Notification Email Address").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create order_email__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("order_email__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("primary_web_url__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Primary Website URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create primary_web_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("primary_web_url__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("twitter_url__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Twitter Website URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create twitter_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("twitter_url__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("facebook_url__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("Facebook Website URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create facebook_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("facebook_url__contactmechpurposetype", c)

	c, err = client.ContactMechPurposeType.Create().SetStringRef("linkedin_url__contactmechpurposetype").
		SetHasTable("No").
		SetDescription("LinkedIn Website URL").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create linkedin_url__contactmechpurposetype: %v", err)
		return err
	}
	cache.Put("linkedin_url__contactmechpurposetype", c)

	return nil
}
