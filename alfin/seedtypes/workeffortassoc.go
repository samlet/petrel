package seedtypes
import "github.com/samlet/petrel/services"


type WorkEffortAssoc struct {
    WorkEffortIdFrom *string `json:"workEffortIdFrom,omitempty" xml:"workEffortIdFrom,attr"`
    WorkEffortIdTo *string `json:"workEffortIdTo,omitempty" xml:"workEffortIdTo,attr"`
    WorkEffortAssocTypeId *string `json:"workEffortAssocTypeId,omitempty" xml:"workEffortAssocTypeId,attr"`
    SequenceNum *int64 `json:"sequenceNum,omitempty" xml:"sequenceNum,attr"`
    FromDate *services.DateTime `json:"fromDate,omitempty" xml:"fromDate,attr"`
    ThruDate *services.DateTime `json:"thruDate,omitempty" xml:"thruDate,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

