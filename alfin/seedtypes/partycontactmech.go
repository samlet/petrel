package seedtypes
import "github.com/samlet/petrel/services"


type PartyContactMech struct {
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    ContactMechId *string `json:"contactMechId,omitempty" xml:"contactMechId,attr"`
    FromDate *services.DateTime `json:"fromDate,omitempty" xml:"fromDate,attr"`
    ThruDate *services.DateTime `json:"thruDate,omitempty" xml:"thruDate,attr"`
    RoleTypeId *string `json:"roleTypeId,omitempty" xml:"roleTypeId,attr"`
    AllowSolicitation *byte `json:"allowSolicitation,omitempty" xml:"allowSolicitation,attr"`
    Extension *string `json:"extension,omitempty" xml:"extension,attr"`
    Verified *byte `json:"verified,omitempty" xml:"verified,attr"`
    Comments *string `json:"comments,omitempty" xml:"comments,attr"`
    YearsWithContactMech *int64 `json:"yearsWithContactMech,omitempty" xml:"yearsWithContactMech,attr"`
    MonthsWithContactMech *int64 `json:"monthsWithContactMech,omitempty" xml:"monthsWithContactMech,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

