package seedtypes
import "github.com/samlet/petrel/services"


type ContactMechTypePurpose struct {
    ContactMechTypeId *string `json:"contactMechTypeId,omitempty" xml:"contactMechTypeId,attr"`
    ContactMechPurposeTypeId *string `json:"contactMechPurposeTypeId,omitempty" xml:"contactMechPurposeTypeId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

