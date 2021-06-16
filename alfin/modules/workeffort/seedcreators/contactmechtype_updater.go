package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"

	"fmt"
)

func UpdateContactMechType(ctx context.Context) error {
	log.Println("ContactMechType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.ContactMechType
	failures := 0

	c = cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Electronic Address").
		AddChildren(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		AddChildren(cache.Get("ip_address__contactmechtype").(*ent.ContactMechType)).
		AddChildren(cache.Get("domain_name__contactmechtype").(*ent.ContactMechType)).
		AddChildren(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		AddChildren(cache.Get("internal_partyid__contactmechtype").(*ent.ContactMechType)).
		AddChildren(cache.Get("ftp_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update electronic_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("electronic_address__contactmechtype", c)
	}

	c = cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Postal Address").
		AddContacMechTypeCommunicationEventTypes(cache.Get("letter_correspondenc__communicationeventtype").(*ent.CommunicationEventType)).
		AddContactMechTypePurposes(cache.Get("postal_address__shipping_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__ship_orig_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__billing_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__payment_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__general_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__pur_ret_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__primary_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("postal_address__previous_location__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update postal_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("postal_address__contactmechtype", c)
	}

	c = cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Phone Number").
		AddContacMechTypeCommunicationEventTypes(cache.Get("fax_communication__communicationeventtype").(*ent.CommunicationEventType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("phone_communication__communicationeventtype").(*ent.CommunicationEventType)).
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_number__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_number_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_did__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_home__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_mobile__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_work__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_work_sec__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_ship_orig__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_payment__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__primary_phone__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_quick__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_shipping__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__fax_billing__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("telecom_number__phone_assistant__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update telecom_number__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("telecom_number__contactmechtype", c)
	}

	c = cache.Get("email_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Email Address").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("email_communication__communicationeventtype").(*ent.CommunicationEventType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("auto_email_comm__communicationeventtype").(*ent.CommunicationEventType)).
		AddContactMechTypePurposes(cache.Get("email_address__billing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("email_address__marketing_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("email_address__payment_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("email_address__order_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("email_address__other_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("email_address__primary_email__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update email_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("email_address__contactmechtype", c)
	}

	c = cache.Get("ip_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Internet IP Address").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ip_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ip_address__contactmechtype", c)
	}

	c = cache.Get("domain_name__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Internet Domain Name").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update domain_name__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("domain_name__contactmechtype", c)
	}

	c = cache.Get("web_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Web URL/Address").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("web_site_communicati__communicationeventtype").(*ent.CommunicationEventType)).
		AddContactMechTypePurposes(cache.Get("web_address__primary_web_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("web_address__twitter_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("web_address__facebook_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		AddContactMechTypePurposes(cache.Get("web_address__linkedin_url__contactmechtypepurpose").(*ent.ContactMechTypePurpose)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update web_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("web_address__contactmechtype", c)
	}

	c = cache.Get("internal_partyid__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Internal Note via partyId").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("comment_note__communicationeventtype").(*ent.CommunicationEventType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update internal_partyid__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("internal_partyid__contactmechtype", c)
	}

	c = cache.Get("ftp_address__contactmechtype").(*ent.ContactMechType)
	c, err = c.Update().
		SetHasTable("Yes").
		SetDescription("Ftp server connection").
		SetParent(cache.Get("electronic_address__contactmechtype").(*ent.ContactMechType)).
		AddContacMechTypeCommunicationEventTypes(cache.Get("file_transfer_comm__communicationeventtype").(*ent.CommunicationEventType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update ftp_address__contactmechtype: %v", err)
		// return err
		// skip update failure
		failures = failures + 1
	} else {
		cache.Put("ftp_address__contactmechtype", c)
	}

	if failures != 0 {
		return fmt.Errorf("occurs %d failtures", failures)
	}
	return nil
}
