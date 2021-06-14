package seedtypes
import "github.com/samlet/petrel/services"


type WorkEffortSkillStandard struct {
    WorkEffortId *string `json:"workEffortId,omitempty" xml:"workEffortId,attr"`
    SkillTypeId *string `json:"skillTypeId,omitempty" xml:"skillTypeId,attr"`
    EstimatedNumPeople *float64 `json:"estimatedNumPeople,omitempty" xml:"estimatedNumPeople,attr"`
    EstimatedDuration *float64 `json:"estimatedDuration,omitempty" xml:"estimatedDuration,attr"`
    EstimatedCost *services.Decimal `json:"estimatedCost,omitempty" xml:"estimatedCost,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

