package seedtypes
import "github.com/samlet/petrel/services"


type TemporalExpression struct {
    TempExprId *string `json:"tempExprId,omitempty" xml:"tempExprId,attr"`
    TempExprTypeId *string `json:"tempExprTypeId,omitempty" xml:"tempExprTypeId,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    Date1 *services.DateTime `json:"date1,omitempty" xml:"date1,attr"`
    Date2 *services.DateTime `json:"date2,omitempty" xml:"date2,attr"`
    Integer1 *int64 `json:"integer1,omitempty" xml:"integer1,attr"`
    Integer2 *int64 `json:"integer2,omitempty" xml:"integer2,attr"`
    String1 *string `json:"string1,omitempty" xml:"string1,attr"`
    String2 *string `json:"string2,omitempty" xml:"string2,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

