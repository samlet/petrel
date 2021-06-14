package seedtypes
import "github.com/samlet/petrel/services"


type TemporalExpressionAssoc struct {
    FromTempExprId *string `json:"fromTempExprId,omitempty" xml:"fromTempExprId,attr"`
    ToTempExprId *string `json:"toTempExprId,omitempty" xml:"toTempExprId,attr"`
    ExprAssocType *string `json:"exprAssocType,omitempty" xml:"exprAssocType,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

