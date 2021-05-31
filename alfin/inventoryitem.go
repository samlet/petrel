// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

type InventoryItem struct {
	InventoryItemId         string   `json:"inventoryItemId,omitempty"`
	InventoryItemTypeId     string   `json:"inventoryItemTypeId,omitempty"`
	ProductId               string   `json:"productId,omitempty"`
	PartyId                 string   `json:"partyId,omitempty"`
	OwnerPartyId            string   `json:"ownerPartyId,omitempty"`
	StatusId                string   `json:"statusId,omitempty"`
	DatetimeReceived        DateTime `json:"datetimeReceived,omitempty"`
	DatetimeManufactured    DateTime `json:"datetimeManufactured,omitempty"`
	ExpireDate              DateTime `json:"expireDate,omitempty"`
	FacilityId              string   `json:"facilityId,omitempty"`
	ContainerId             string   `json:"containerId,omitempty"`
	LotId                   string   `json:"lotId,omitempty"`
	UomId                   string   `json:"uomId,omitempty"`
	BinNumber               string   `json:"binNumber,omitempty"`
	LocationSeqId           string   `json:"locationSeqId,omitempty"`
	Comments                string   `json:"comments,omitempty"`
	QuantityOnHandTotal     float64  `json:"quantityOnHandTotal,omitempty"`
	AvailableToPromiseTotal float64  `json:"availableToPromiseTotal,omitempty"`
	AccountingQuantityTotal float64  `json:"accountingQuantityTotal,omitempty"`
	OldQuantityOnHand       float64  `json:"oldQuantityOnHand,omitempty"`
	OldAvailableToPromise   float64  `json:"oldAvailableToPromise,omitempty"`
	SerialNumber            string   `json:"serialNumber,omitempty"`
	SoftIdentifier          string   `json:"softIdentifier,omitempty"`
	ActivationNumber        string   `json:"activationNumber,omitempty"`
	ActivationValidThru     DateTime `json:"activationValidThru,omitempty"`
	UnitCost                float64  `json:"unitCost,omitempty"`
	CurrencyUomId           string   `json:"currencyUomId,omitempty"`
	FixedAssetId            string   `json:"fixedAssetId,omitempty"`
	LastUpdatedStamp        DateTime `json:"lastUpdatedStamp,omitempty"`
	LastUpdatedTxStamp      DateTime `json:"lastUpdatedTxStamp,omitempty"`
	CreatedStamp            DateTime `json:"createdStamp,omitempty"`
	CreatedTxStamp          DateTime `json:"createdTxStamp,omitempty"`
}
