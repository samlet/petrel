package seedcreators

import (
	"context"
	cachecomp "github.com/samlet/petrel/alfin/cache"
	"github.com/samlet/petrel/alfin/common"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent"
	"log"
)

func CreateCommunicationEventType(ctx context.Context) error {
	log.Println("CommunicationEventType creator", common.Version)
	client := ent.FromContext(ctx)
	cache := cachecomp.FromContext(ctx)

	var err error
	var c *ent.CommunicationEventType

	c, err = client.CommunicationEventType.Create().SetStringRef("email_communication__communicationeventtype").
		SetHasTable("No").
		SetDescription("Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create email_communication__communicationeventtype: %v", err)
		return err
	}
	cache.Put("email_communication__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("face_to_face_communi__communicationeventtype").
		SetHasTable("No").
		SetDescription("Face-To-Face").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create face_to_face_communi__communicationeventtype: %v", err)
		return err
	}
	cache.Put("face_to_face_communi__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("fax_communication__communicationeventtype").
		SetHasTable("No").
		SetDescription("Fax").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create fax_communication__communicationeventtype: %v", err)
		return err
	}
	cache.Put("fax_communication__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("letter_correspondenc__communicationeventtype").
		SetHasTable("No").
		SetDescription("Letter").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create letter_correspondenc__communicationeventtype: %v", err)
		return err
	}
	cache.Put("letter_correspondenc__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("phone_communication__communicationeventtype").
		SetHasTable("No").
		SetDescription("Phone").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create phone_communication__communicationeventtype: %v", err)
		return err
	}
	cache.Put("phone_communication__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("web_site_communicati__communicationeventtype").
		SetHasTable("No").
		SetDescription("Web Site").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create web_site_communicati__communicationeventtype: %v", err)
		return err
	}
	cache.Put("web_site_communicati__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("comment_note__communicationeventtype").
		SetHasTable("No").
		SetDescription("Comment/Note").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create comment_note__communicationeventtype: %v", err)
		return err
	}
	cache.Put("comment_note__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("auto_email_comm__communicationeventtype").
		SetHasTable("No").
		SetDescription("Auto Email").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create auto_email_comm__communicationeventtype: %v", err)
		return err
	}
	cache.Put("auto_email_comm__communicationeventtype", c)

	c, err = client.CommunicationEventType.Create().SetStringRef("file_transfer_comm__communicationeventtype").
		SetHasTable("No").
		SetDescription("File transfer").
		Save(ctx)
	if err != nil {
		log.Printf("fail to create file_transfer_comm__communicationeventtype: %v", err)
		return err
	}
	cache.Put("file_transfer_comm__communicationeventtype", c)

	return nil
}
