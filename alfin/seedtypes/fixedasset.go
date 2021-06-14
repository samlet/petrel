package seedtypes
import "github.com/samlet/petrel/services"


type FixedAsset struct {
    FixedAssetId *string `json:"fixedAssetId,omitempty" xml:"fixedAssetId,attr"`
    FixedAssetTypeId *string `json:"fixedAssetTypeId,omitempty" xml:"fixedAssetTypeId,attr"`
    ParentFixedAssetId *string `json:"parentFixedAssetId,omitempty" xml:"parentFixedAssetId,attr"`
    InstanceOfProductId *string `json:"instanceOfProductId,omitempty" xml:"instanceOfProductId,attr"`
    ClassEnumId *string `json:"classEnumId,omitempty" xml:"classEnumId,attr"`
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    RoleTypeId *string `json:"roleTypeId,omitempty" xml:"roleTypeId,attr"`
    FixedAssetName *string `json:"fixedAssetName,omitempty" xml:"fixedAssetName,attr"`
    AcquireOrderId *string `json:"acquireOrderId,omitempty" xml:"acquireOrderId,attr"`
    AcquireOrderItemSeqId *string `json:"acquireOrderItemSeqId,omitempty" xml:"acquireOrderItemSeqId,attr"`
    DateAcquired *services.DateTime `json:"dateAcquired,omitempty" xml:"dateAcquired,attr"`
    DateLastServiced *services.DateTime `json:"dateLastServiced,omitempty" xml:"dateLastServiced,attr"`
    DateNextService *services.DateTime `json:"dateNextService,omitempty" xml:"dateNextService,attr"`
    ExpectedEndOfLife *services.DateTime `json:"expectedEndOfLife,omitempty" xml:"expectedEndOfLife,attr"`
    ActualEndOfLife *services.DateTime `json:"actualEndOfLife,omitempty" xml:"actualEndOfLife,attr"`
    ProductionCapacity *services.Decimal `json:"productionCapacity,omitempty" xml:"productionCapacity,attr"`
    UomId *string `json:"uomId,omitempty" xml:"uomId,attr"`
    CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,attr"`
    SerialNumber *string `json:"serialNumber,omitempty" xml:"serialNumber,attr"`
    LocatedAtFacilityId *string `json:"locatedAtFacilityId,omitempty" xml:"locatedAtFacilityId,attr"`
    LocatedAtLocationSeqId *string `json:"locatedAtLocationSeqId,omitempty" xml:"locatedAtLocationSeqId,attr"`
    SalvageValue *services.Decimal `json:"salvageValue,omitempty" xml:"salvageValue,attr"`
    Depreciation *services.Decimal `json:"depreciation,omitempty" xml:"depreciation,attr"`
    PurchaseCost *services.Decimal `json:"purchaseCost,omitempty" xml:"purchaseCost,attr"`
    PurchaseCostUomId *string `json:"purchaseCostUomId,omitempty" xml:"purchaseCostUomId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

