package seedtypes
import "github.com/samlet/petrel/services"


type PartyRole struct {
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    RoleTypeId *string `json:"roleTypeId,omitempty" xml:"roleTypeId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

