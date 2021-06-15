package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func UpdateCommunicationEventType(ctx context.Context) error {
	log.Println("CommunicationEventType updater", common.Version)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CommunicationEventType

	c = cache.Get("email_communication__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Email").
		SetContacMechTypeContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update email_communication__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("email_communication__communicationeventtype", c)
	}

	c = cache.Get("face_to_face_communi__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Face-To-Face").
		Save(ctx)
	if err != nil {
		log.Printf("fail to update face_to_face_communi__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("face_to_face_communi__communicationeventtype", c)
	}

	c = cache.Get("fax_communication__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Fax").
		SetContacMechTypeContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update fax_communication__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("fax_communication__communicationeventtype", c)
	}

	c = cache.Get("letter_correspondenc__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Letter").
		SetContacMechTypeContactMechType(cache.Get("postal_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update letter_correspondenc__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("letter_correspondenc__communicationeventtype", c)
	}

	c = cache.Get("phone_communication__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Phone").
		SetContacMechTypeContactMechType(cache.Get("telecom_number__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update phone_communication__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("phone_communication__communicationeventtype", c)
	}

	c = cache.Get("web_site_communicati__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Web Site").
		SetContacMechTypeContactMechType(cache.Get("web_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update web_site_communicati__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("web_site_communicati__communicationeventtype", c)
	}

	c = cache.Get("comment_note__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Comment/Note").
		SetContacMechTypeContactMechType(cache.Get("internal_partyid__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update comment_note__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("comment_note__communicationeventtype", c)
	}

	c = cache.Get("auto_email_comm__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("Auto Email").
		SetContacMechTypeContactMechType(cache.Get("email_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update auto_email_comm__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("auto_email_comm__communicationeventtype", c)
	}

	c = cache.Get("file_transfer_comm__communicationeventtype").(*ent.CommunicationEventType)
	c, err = c.Update().
		SetHasTable("No").
		SetDescription("File transfer").
		SetContacMechTypeContactMechType(cache.Get("ftp_address__contactmechtype").(*ent.ContactMechType)).
		Save(ctx)
	if err != nil {
		log.Printf("fail to update file_transfer_comm__communicationeventtype: %v", err)
		// return err
		// skip update failure
	} else {
		cache.Put("file_transfer_comm__communicationeventtype", c)
	}

	return nil
}
