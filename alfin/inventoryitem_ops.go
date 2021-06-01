// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"

type CreateInventoryItemParams struct {
	services.AlfinParams
	InventoryItemTypeId     string              `json:"inventoryItemTypeId,omitempty"`
	ProductId               string              `json:"productId,omitempty"`
	PartyId                 string              `json:"partyId,omitempty"`
	OwnerPartyId            string              `json:"ownerPartyId,omitempty"`
	StatusId                string              `json:"statusId,omitempty"`
	DatetimeReceived        *services.Timestamp `json:"datetimeReceived,omitempty"`
	DatetimeManufactured    *services.Timestamp `json:"datetimeManufactured,omitempty"`
	ExpireDate              *services.Timestamp `json:"expireDate,omitempty"`
	FacilityId              string              `json:"facilityId,omitempty"`
	ContainerId             string              `json:"containerId,omitempty"`
	LotId                   string              `json:"lotId,omitempty"`
	UomId                   string              `json:"uomId,omitempty"`
	BinNumber               string              `json:"binNumber,omitempty"`
	LocationSeqId           string              `json:"locationSeqId,omitempty"`
	Comments                string              `json:"comments,omitempty"`
	AccountingQuantityTotal *services.Decimal   `json:"accountingQuantityTotal,omitempty"`
	OldQuantityOnHand       *services.Decimal   `json:"oldQuantityOnHand,omitempty"`
	OldAvailableToPromise   *services.Decimal   `json:"oldAvailableToPromise,omitempty"`
	SerialNumber            string              `json:"serialNumber,omitempty"`
	SoftIdentifier          string              `json:"softIdentifier,omitempty"`
	ActivationNumber        string              `json:"activationNumber,omitempty"`
	ActivationValidThru     *services.Timestamp `json:"activationValidThru,omitempty"`
	UnitCost                *services.Decimal   `json:"unitCost,omitempty"`
	CurrencyUomId           string              `json:"currencyUomId,omitempty"`
	FixedAssetId            string              `json:"fixedAssetId,omitempty"`
	IsReturned              string              `json:"isReturned,omitempty"`
}

type CreateInventoryItemResult struct {
	services.APIResource
	InventoryItemId string `json:"inventoryItemId,omitempty"`
}

type GetInventoryItemOwnerParams struct {
	services.AlfinParams
	InventoryItemId string `json:"inventoryItemId,omitempty"`
}

type GetInventoryItemOwnerResult struct {
	services.APIResource
	OwnerPartyId string `json:"ownerPartyId,omitempty"`
}

type InventoryItemCheckSetDefaultValuesParams struct {
	services.AlfinParams
	InventoryItemId string                 `json:"inventoryItemId,omitempty"`
	InventoryItem   map[string]interface{} `json:"inventoryItem,omitempty"`
}

type InventoryItemCheckSetDefaultValuesResult struct {
	services.APIResource
}

type UpdateInventoryItemParams struct {
	services.AlfinParams
	InventoryItemId         string              `json:"inventoryItemId,omitempty"`
	InventoryItemTypeId     string              `json:"inventoryItemTypeId,omitempty"`
	ProductId               string              `json:"productId,omitempty"`
	PartyId                 string              `json:"partyId,omitempty"`
	OwnerPartyId            string              `json:"ownerPartyId,omitempty"`
	StatusId                string              `json:"statusId,omitempty"`
	DatetimeReceived        *services.Timestamp `json:"datetimeReceived,omitempty"`
	DatetimeManufactured    *services.Timestamp `json:"datetimeManufactured,omitempty"`
	ExpireDate              *services.Timestamp `json:"expireDate,omitempty"`
	FacilityId              string              `json:"facilityId,omitempty"`
	ContainerId             string              `json:"containerId,omitempty"`
	LotId                   string              `json:"lotId,omitempty"`
	UomId                   string              `json:"uomId,omitempty"`
	BinNumber               string              `json:"binNumber,omitempty"`
	LocationSeqId           string              `json:"locationSeqId,omitempty"`
	Comments                string              `json:"comments,omitempty"`
	AccountingQuantityTotal *services.Decimal   `json:"accountingQuantityTotal,omitempty"`
	OldQuantityOnHand       *services.Decimal   `json:"oldQuantityOnHand,omitempty"`
	OldAvailableToPromise   *services.Decimal   `json:"oldAvailableToPromise,omitempty"`
	SerialNumber            string              `json:"serialNumber,omitempty"`
	SoftIdentifier          string              `json:"softIdentifier,omitempty"`
	ActivationNumber        string              `json:"activationNumber,omitempty"`
	ActivationValidThru     *services.Timestamp `json:"activationValidThru,omitempty"`
	UnitCost                *services.Decimal   `json:"unitCost,omitempty"`
	CurrencyUomId           string              `json:"currencyUomId,omitempty"`
	FixedAssetId            string              `json:"fixedAssetId,omitempty"`
}

type UpdateInventoryItemResult struct {
	services.APIResource
	OldOwnerPartyId string `json:"oldOwnerPartyId,omitempty"`
	OldProductId    string `json:"oldProductId,omitempty"`
	OldStatusId     string `json:"oldStatusId,omitempty"`
}

type UpdateSerializedInventoryTotalsParams struct {
	services.AlfinParams
	InventoryItemId string `json:"inventoryItemId,omitempty"`
}

type UpdateSerializedInventoryTotalsResult struct {
	services.APIResource
}

// Interface
type InventoryItemOps interface {
	// CreateInventoryItem Create an InventoryItem
	CreateInventoryItem(params *CreateInventoryItemParams) (*CreateInventoryItemResult, error)
	// GetInventoryItemOwner get an ownerPartyId from inventoryItemId
	GetInventoryItemOwner(params *GetInventoryItemOwnerParams) (*GetInventoryItemOwnerResult, error)
	// InventoryItemCheckSetDefaultValues Check and, if empty, fills with default values ownerPartyId, currencyUomId, unitCost
	InventoryItemCheckSetDefaultValues(params *InventoryItemCheckSetDefaultValuesParams) (*InventoryItemCheckSetDefaultValuesResult, error)
	// UpdateInventoryItem Update an InventoryItem
	UpdateInventoryItem(params *UpdateInventoryItemParams) (*UpdateInventoryItemResult, error)
	// UpdateSerializedInventoryTotals Sets the ATP/QOH totals for serialized inventory items
	UpdateSerializedInventoryTotals(params *UpdateSerializedInventoryTotalsParams) (*UpdateSerializedInventoryTotalsResult, error)
}
