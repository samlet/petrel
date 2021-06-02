// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"


type CreateFacilityParams struct {
    services.AlfinParams
    FacilityTypeId string    `json:"facilityTypeId,omitempty"`
    ParentFacilityId string    `json:"parentFacilityId,omitempty"`
    OwnerPartyId string    `json:"ownerPartyId,omitempty"`
    DefaultInventoryItemTypeId string    `json:"defaultInventoryItemTypeId,omitempty"`
    FacilityName string    `json:"facilityName,omitempty"`
    PrimaryFacilityGroupId string    `json:"primaryFacilityGroupId,omitempty"`
    OldSquareFootage int64    `json:"oldSquareFootage,omitempty"`
    FacilitySize *services.Decimal    `json:"facilitySize,omitempty"`
    FacilitySizeUomId string    `json:"facilitySizeUomId,omitempty"`
    ProductStoreId string    `json:"productStoreId,omitempty"`
    DefaultDaysToShip int64    `json:"defaultDaysToShip,omitempty"`
    OpenedDate *services.Timestamp    `json:"openedDate,omitempty"`
    ClosedDate *services.Timestamp    `json:"closedDate,omitempty"`
    Description string    `json:"description,omitempty"`
    DefaultDimensionUomId string    `json:"defaultDimensionUomId,omitempty"`
    DefaultWeightUomId string    `json:"defaultWeightUomId,omitempty"`
    GeoPointId string    `json:"geoPointId,omitempty"`
    FacilityLevel int64    `json:"facilityLevel,omitempty"`
}

type CreateFacilityResult struct {
    services.APIResource
    FacilityId string   `json:"facilityId,omitempty"`
}

type DeleteFacilityParams struct {
    services.AlfinParams
    FacilityId string    `json:"facilityId,omitempty"`
}

type DeleteFacilityResult struct {
    services.APIResource
}

type UpdateFacilityParams struct {
    services.AlfinParams
    FacilityId string    `json:"facilityId,omitempty"`
    FacilityTypeId string    `json:"facilityTypeId,omitempty"`
    ParentFacilityId string    `json:"parentFacilityId,omitempty"`
    OwnerPartyId string    `json:"ownerPartyId,omitempty"`
    DefaultInventoryItemTypeId string    `json:"defaultInventoryItemTypeId,omitempty"`
    FacilityName string    `json:"facilityName,omitempty"`
    PrimaryFacilityGroupId string    `json:"primaryFacilityGroupId,omitempty"`
    OldSquareFootage int64    `json:"oldSquareFootage,omitempty"`
    FacilitySize *services.Decimal    `json:"facilitySize,omitempty"`
    FacilitySizeUomId string    `json:"facilitySizeUomId,omitempty"`
    ProductStoreId string    `json:"productStoreId,omitempty"`
    DefaultDaysToShip int64    `json:"defaultDaysToShip,omitempty"`
    OpenedDate *services.Timestamp    `json:"openedDate,omitempty"`
    ClosedDate *services.Timestamp    `json:"closedDate,omitempty"`
    Description string    `json:"description,omitempty"`
    DefaultDimensionUomId string    `json:"defaultDimensionUomId,omitempty"`
    DefaultWeightUomId string    `json:"defaultWeightUomId,omitempty"`
    GeoPointId string    `json:"geoPointId,omitempty"`
    FacilityLevel int64    `json:"facilityLevel,omitempty"`
}

type UpdateFacilityResult struct {
    services.APIResource
}

type GetInventoryAvailableByFacilityParams struct {
    services.AlfinParams
    ProductId string    `json:"productId,omitempty"`
    FacilityId string    `json:"facilityId,omitempty"`
    StatusId string    `json:"statusId,omitempty"`
    LotId string    `json:"lotId,omitempty"`
    UseCache bool    `json:"useCache,omitempty"`
}

type GetInventoryAvailableByFacilityResult struct {
    services.APIResource
    QuantityOnHandTotal *services.Decimal   `json:"quantityOnHandTotal,omitempty"`
    AvailableToPromiseTotal *services.Decimal   `json:"availableToPromiseTotal,omitempty"`
    AccountingQuantityTotal *services.Decimal   `json:"accountingQuantityTotal,omitempty"`
}


// Interface
type FacilityOps interface {
    // CreateFacility Create a Facility
    CreateFacility(params *CreateFacilityParams) (*CreateFacilityResult, error)
    // DeleteFacility Delete a Facility
    DeleteFacility(params *DeleteFacilityParams) (*DeleteFacilityResult, error)
    // UpdateFacility Update a Facility
    UpdateFacility(params *UpdateFacilityParams) (*UpdateFacilityResult, error)
    // GetInventoryAvailableByFacility Get Inventory Availability for a Product constrained by a facilityId
    GetInventoryAvailableByFacility(params *GetInventoryAvailableByFacilityParams) (*GetInventoryAvailableByFacilityResult, error)
}

