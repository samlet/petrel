package seedtypes
import "github.com/samlet/petrel/services"


type PartyStatus struct {
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    StatusDate *services.DateTime `json:"statusDate,omitempty" xml:"statusDate,attr"`
    ChangeByUserLoginId *string `json:"changeByUserLoginId,omitempty" xml:"changeByUserLoginId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

