// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"

type InventoryItem struct {
	InventoryItemId         string             `json:"inventoryItemId,omitempty"`
	InventoryItemTypeId     string             `json:"inventoryItemTypeId,omitempty"`
	ProductId               string             `json:"productId,omitempty"`
	PartyId                 string             `json:"partyId,omitempty"`
	OwnerPartyId            string             `json:"ownerPartyId,omitempty"`
	StatusId                string             `json:"statusId,omitempty"`
	DatetimeReceived        *services.DateTime `json:"datetimeReceived,omitempty"`
	DatetimeManufactured    *services.DateTime `json:"datetimeManufactured,omitempty"`
	ExpireDate              *services.DateTime `json:"expireDate,omitempty"`
	FacilityId              string             `json:"facilityId,omitempty"`
	ContainerId             string             `json:"containerId,omitempty"`
	LotId                   string             `json:"lotId,omitempty"`
	UomId                   string             `json:"uomId,omitempty"`
	BinNumber               string             `json:"binNumber,omitempty"`
	LocationSeqId           string             `json:"locationSeqId,omitempty"`
	Comments                string             `json:"comments,omitempty"`
	QuantityOnHandTotal     *services.Decimal  `json:"quantityOnHandTotal,omitempty"`
	AvailableToPromiseTotal *services.Decimal  `json:"availableToPromiseTotal,omitempty"`
	AccountingQuantityTotal *services.Decimal  `json:"accountingQuantityTotal,omitempty"`
	OldQuantityOnHand       *services.Decimal  `json:"oldQuantityOnHand,omitempty"`
	OldAvailableToPromise   *services.Decimal  `json:"oldAvailableToPromise,omitempty"`
	SerialNumber            string             `json:"serialNumber,omitempty"`
	SoftIdentifier          string             `json:"softIdentifier,omitempty"`
	ActivationNumber        string             `json:"activationNumber,omitempty"`
	ActivationValidThru     *services.DateTime `json:"activationValidThru,omitempty"`
	UnitCost                *services.Decimal  `json:"unitCost,omitempty"`
	CurrencyUomId           string             `json:"currencyUomId,omitempty"`
	FixedAssetId            string             `json:"fixedAssetId,omitempty"`
	LastUpdatedStamp        *services.DateTime `json:"lastUpdatedStamp,omitempty"`
	LastUpdatedTxStamp      *services.DateTime `json:"lastUpdatedTxStamp,omitempty"`
	CreatedStamp            *services.DateTime `json:"createdStamp,omitempty"`
	CreatedTxStamp          *services.DateTime `json:"createdTxStamp,omitempty"`
}
