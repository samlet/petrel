package seedtypes
import "github.com/samlet/petrel/services"


type StatusItem struct {
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    StatusTypeId *string `json:"statusTypeId,omitempty" xml:"statusTypeId,attr"`
    StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,attr"`
    SequenceId *string `json:"sequenceId,omitempty" xml:"sequenceId,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

