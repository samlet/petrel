package seedtypes
import "github.com/samlet/petrel/services"


type SecurityGroupPermission struct {
    GroupId *string `json:"groupId,omitempty" xml:"groupId,attr"`
    PermissionId *string `json:"permissionId,omitempty" xml:"permissionId,attr"`
    FromDate *services.DateTime `json:"fromDate,omitempty" xml:"fromDate,attr"`
    ThruDate *services.DateTime `json:"thruDate,omitempty" xml:"thruDate,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

