package seedtypes
import "github.com/samlet/petrel/services"


type Party struct {
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    PartyTypeId *string `json:"partyTypeId,omitempty" xml:"partyTypeId,attr"`
    ExternalId *string `json:"externalId,omitempty" xml:"externalId,attr"`
    PreferredCurrencyUomId *string `json:"preferredCurrencyUomId,omitempty" xml:"preferredCurrencyUomId,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    CreatedDate *services.DateTime `json:"createdDate,omitempty" xml:"createdDate,attr"`
    CreatedByUserLogin *string `json:"createdByUserLogin,omitempty" xml:"createdByUserLogin,attr"`
    LastModifiedDate *services.DateTime `json:"lastModifiedDate,omitempty" xml:"lastModifiedDate,attr"`
    LastModifiedByUserLogin *string `json:"lastModifiedByUserLogin,omitempty" xml:"lastModifiedByUserLogin,attr"`
    DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,attr"`
    IsUnread *byte `json:"isUnread,omitempty" xml:"isUnread,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

