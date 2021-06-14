package seedtypes
import "github.com/samlet/petrel/services"


type WorkEffortPartyAssignment struct {
    WorkEffortId *string `json:"workEffortId,omitempty" xml:"workEffortId,attr"`
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    RoleTypeId *string `json:"roleTypeId,omitempty" xml:"roleTypeId,attr"`
    FromDate *services.DateTime `json:"fromDate,omitempty" xml:"fromDate,attr"`
    ThruDate *services.DateTime `json:"thruDate,omitempty" xml:"thruDate,attr"`
    AssignedByUserLoginId *string `json:"assignedByUserLoginId,omitempty" xml:"assignedByUserLoginId,attr"`
    StatusId *string `json:"statusId,omitempty" xml:"statusId,attr"`
    StatusDateTime *services.DateTime `json:"statusDateTime,omitempty" xml:"statusDateTime,attr"`
    ExpectationEnumId *string `json:"expectationEnumId,omitempty" xml:"expectationEnumId,attr"`
    DelegateReasonEnumId *string `json:"delegateReasonEnumId,omitempty" xml:"delegateReasonEnumId,attr"`
    FacilityId *string `json:"facilityId,omitempty" xml:"facilityId,attr"`
    Comments *string `json:"comments,omitempty" xml:"comments,attr"`
    MustRsvp *byte `json:"mustRsvp,omitempty" xml:"mustRsvp,attr"`
    AvailabilityStatusId *string `json:"availabilityStatusId,omitempty" xml:"availabilityStatusId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

