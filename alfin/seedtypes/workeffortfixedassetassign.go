package seedtypes
import "github.com/samlet/petrel/services"


type WorkEffortFixedAssetAssign struct {
    WorkEffortId *string `json:"workEffortId,omitempty" xml:"workEffortId,attr"`
    FixedAssetId *string `json:"fixedAssetId,omitempty" xml:"fixedAssetId,attr"`
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    FromDate *services.DateTime `json:"fromDate,omitempty" xml:"fromDate,attr"`
    ThruDate *services.DateTime `json:"thruDate,omitempty" xml:"thruDate,attr"`
    AvailabilityStatusId *string `json:"availabilityStatusId,omitempty" xml:"availabilityStatusId,attr"`
    AllocatedCost *services.Decimal `json:"allocatedCost,omitempty" xml:"allocatedCost,attr"`
    Comments *string `json:"comments,omitempty" xml:"comments,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

