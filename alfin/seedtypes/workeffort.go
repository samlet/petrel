package seedtypes
import "github.com/samlet/petrel/services"


type WorkEffort struct {
    WorkEffortId *string `json:"workEffortId,omitempty" xml:"workEffortId,attr"`
    WorkEffortTypeId *string `json:"workEffortTypeId,omitempty" xml:"workEffortTypeId,attr"`
    CurrentStatusId *string `json:"currentStatusId,omitempty" xml:"currentStatusId,attr"`
    LastStatusUpdate *services.DateTime `json:"lastStatusUpdate,omitempty" xml:"lastStatusUpdate,attr"`
    WorkEffortPurposeTypeId *string `json:"workEffortPurposeTypeId,omitempty" xml:"workEffortPurposeTypeId,attr"`
    WorkEffortParentId *string `json:"workEffortParentId,omitempty" xml:"workEffortParentId,attr"`
    ScopeEnumId *string `json:"scopeEnumId,omitempty" xml:"scopeEnumId,attr"`
    Priority *int64 `json:"priority,omitempty" xml:"priority,attr"`
    PercentComplete *int64 `json:"percentComplete,omitempty" xml:"percentComplete,attr"`
    WorkEffortName *string `json:"workEffortName,omitempty" xml:"workEffortName,attr"`
    ShowAsEnumId *string `json:"showAsEnumId,omitempty" xml:"showAsEnumId,attr"`
    SendNotificationEmail *byte `json:"sendNotificationEmail,omitempty" xml:"sendNotificationEmail,attr"`
    Description *string `json:"description,omitempty" xml:"description,attr"`
    LocationDesc *string `json:"locationDesc,omitempty" xml:"locationDesc,attr"`
    EstimatedStartDate *services.DateTime `json:"estimatedStartDate,omitempty" xml:"estimatedStartDate,attr"`
    EstimatedCompletionDate *services.DateTime `json:"estimatedCompletionDate,omitempty" xml:"estimatedCompletionDate,attr"`
    ActualStartDate *services.DateTime `json:"actualStartDate,omitempty" xml:"actualStartDate,attr"`
    ActualCompletionDate *services.DateTime `json:"actualCompletionDate,omitempty" xml:"actualCompletionDate,attr"`
    EstimatedMilliSeconds *float64 `json:"estimatedMilliSeconds,omitempty" xml:"estimatedMilliSeconds,attr"`
    EstimatedSetupMillis *float64 `json:"estimatedSetupMillis,omitempty" xml:"estimatedSetupMillis,attr"`
    EstimateCalcMethod *string `json:"estimateCalcMethod,omitempty" xml:"estimateCalcMethod,attr"`
    ActualMilliSeconds *float64 `json:"actualMilliSeconds,omitempty" xml:"actualMilliSeconds,attr"`
    ActualSetupMillis *float64 `json:"actualSetupMillis,omitempty" xml:"actualSetupMillis,attr"`
    TotalMilliSecondsAllowed *float64 `json:"totalMilliSecondsAllowed,omitempty" xml:"totalMilliSecondsAllowed,attr"`
    TotalMoneyAllowed *services.Decimal `json:"totalMoneyAllowed,omitempty" xml:"totalMoneyAllowed,attr"`
    MoneyUomId *string `json:"moneyUomId,omitempty" xml:"moneyUomId,attr"`
    SpecialTerms *string `json:"specialTerms,omitempty" xml:"specialTerms,attr"`
    TimeTransparency *int64 `json:"timeTransparency,omitempty" xml:"timeTransparency,attr"`
    UniversalId *string `json:"universalId,omitempty" xml:"universalId,attr"`
    SourceReferenceId *string `json:"sourceReferenceId,omitempty" xml:"sourceReferenceId,attr"`
    FixedAssetId *string `json:"fixedAssetId,omitempty" xml:"fixedAssetId,attr"`
    FacilityId *string `json:"facilityId,omitempty" xml:"facilityId,attr"`
    InfoUrl *string `json:"infoUrl,omitempty" xml:"infoUrl,attr"`
    RecurrenceInfoId *string `json:"recurrenceInfoId,omitempty" xml:"recurrenceInfoId,attr"`
    TempExprId *string `json:"tempExprId,omitempty" xml:"tempExprId,attr"`
    RuntimeDataId *string `json:"runtimeDataId,omitempty" xml:"runtimeDataId,attr"`
    NoteId *string `json:"noteId,omitempty" xml:"noteId,attr"`
    ServiceLoaderName *string `json:"serviceLoaderName,omitempty" xml:"serviceLoaderName,attr"`
    QuantityToProduce *services.Decimal `json:"quantityToProduce,omitempty" xml:"quantityToProduce,attr"`
    QuantityProduced *services.Decimal `json:"quantityProduced,omitempty" xml:"quantityProduced,attr"`
    QuantityRejected *services.Decimal `json:"quantityRejected,omitempty" xml:"quantityRejected,attr"`
    ReservPersons *services.Decimal `json:"reservPersons,omitempty" xml:"reservPersons,attr"`
    Reserv2ndPPPerc *services.Decimal `json:"reserv2ndPPPerc,omitempty" xml:"reserv2ndPPPerc,attr"`
    ReservNthPPPerc *services.Decimal `json:"reservNthPPPerc,omitempty" xml:"reservNthPPPerc,attr"`
    AccommodationMapId *string `json:"accommodationMapId,omitempty" xml:"accommodationMapId,attr"`
    AccommodationSpotId *string `json:"accommodationSpotId,omitempty" xml:"accommodationSpotId,attr"`
    RevisionNumber *int64 `json:"revisionNumber,omitempty" xml:"revisionNumber,attr"`
    CreatedDate *services.DateTime `json:"createdDate,omitempty" xml:"createdDate,attr"`
    CreatedByUserLogin *string `json:"createdByUserLogin,omitempty" xml:"createdByUserLogin,attr"`
    LastModifiedDate *services.DateTime `json:"lastModifiedDate,omitempty" xml:"lastModifiedDate,attr"`
    LastModifiedByUserLogin *string `json:"lastModifiedByUserLogin,omitempty" xml:"lastModifiedByUserLogin,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
    SequenceNum *int64 `json:"sequenceNum,omitempty" xml:"sequenceNum,attr"`
}

