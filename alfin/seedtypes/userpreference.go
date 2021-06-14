package seedtypes
import "github.com/samlet/petrel/services"


type UserPreference struct {
    UserLoginId *string `json:"userLoginId,omitempty" xml:"userLoginId,attr"`
    UserPrefTypeId *string `json:"userPrefTypeId,omitempty" xml:"userPrefTypeId,attr"`
    UserPrefGroupTypeId *string `json:"userPrefGroupTypeId,omitempty" xml:"userPrefGroupTypeId,attr"`
    UserPrefValue *string `json:"userPrefValue,omitempty" xml:"userPrefValue,attr"`
    UserPrefDataType *string `json:"userPrefDataType,omitempty" xml:"userPrefDataType,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

