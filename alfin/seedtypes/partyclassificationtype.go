package seedtypes
import "github.com/samlet/petrel/services"


type PartyClassificationType struct {
    PartyClassificationTypeId *string `json:"partyClassificationTypeId,omitempty" xml:"partyClassificationTypeId,attr"`
    ParentTypeId *string `json:"parentTypeId,omitempty" xml:"parentTypeId,attr"`
    HasTable *byte `json:"hasTable,omitempty" xml:"hasTable,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

