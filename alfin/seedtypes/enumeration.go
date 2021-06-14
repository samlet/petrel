package seedtypes
import "github.com/samlet/petrel/services"


type Enumeration struct {
    EnumId *string `json:"enumId,omitempty" xml:"enumId,attr"`
    EnumTypeId *string `json:"enumTypeId,omitempty" xml:"enumTypeId,attr"`
    EnumCode *string `json:"enumCode,omitempty" xml:"enumCode,attr"`
    SequenceId *string `json:"sequenceId,omitempty" xml:"sequenceId,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

