package seedtypes
import "github.com/samlet/petrel/services"


type StatusValidChange struct {
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    StatusIdTo *string `json:"statusIdTo,omitempty" xml:"statusIdTo,attr"`
    ConditionExpression *string `json:"conditionExpression,omitempty" xml:"conditionExpression,attr"`
    TransitionName *string `json:"transitionName,omitempty" xml:"transitionName,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

