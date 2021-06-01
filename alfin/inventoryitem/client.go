// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inventoryitem

import (
	"github.com/samlet/petrel/alfin"
	"github.com/samlet/petrel/services"
	"net/http"
)

type InventoryItemClient struct {
	B   *services.AlfinBackend
	Key string
}

func newClient(key string) InventoryItemClient {
	return InventoryItemClient{services.NewAlfinBackend(), key}
}

// CreateInventoryItem Create an InventoryItem
func (c *InventoryItemClient) CreateInventoryItem(params *alfin.CreateInventoryItemParams) (*alfin.CreateInventoryItemResult, error) {
	reversal := &alfin.CreateInventoryItemResult{}
	err := c.B.Call(http.MethodPost, "createInventoryItem", c.Key, params, reversal)
	return reversal, err
}

// GetInventoryItemOwner get an ownerPartyId from inventoryItemId
func (c *InventoryItemClient) GetInventoryItemOwner(params *alfin.GetInventoryItemOwnerParams) (*alfin.GetInventoryItemOwnerResult, error) {
	reversal := &alfin.GetInventoryItemOwnerResult{}
	err := c.B.Call(http.MethodPost, "getInventoryItemOwner", c.Key, params, reversal)
	return reversal, err
}

// InventoryItemCheckSetDefaultValues Check and, if empty, fills with default values ownerPartyId, currencyUomId, unitCost
func (c *InventoryItemClient) InventoryItemCheckSetDefaultValues(params *alfin.InventoryItemCheckSetDefaultValuesParams) (*alfin.InventoryItemCheckSetDefaultValuesResult, error) {
	reversal := &alfin.InventoryItemCheckSetDefaultValuesResult{}
	err := c.B.Call(http.MethodPost, "inventoryItemCheckSetDefaultValues", c.Key, params, reversal)
	return reversal, err
}

// UpdateInventoryItem Update an InventoryItem
func (c *InventoryItemClient) UpdateInventoryItem(params *alfin.UpdateInventoryItemParams) (*alfin.UpdateInventoryItemResult, error) {
	reversal := &alfin.UpdateInventoryItemResult{}
	err := c.B.Call(http.MethodPost, "updateInventoryItem", c.Key, params, reversal)
	return reversal, err
}

// UpdateSerializedInventoryTotals Sets the ATP/QOH totals for serialized inventory items
func (c *InventoryItemClient) UpdateSerializedInventoryTotals(params *alfin.UpdateSerializedInventoryTotalsParams) (*alfin.UpdateSerializedInventoryTotalsResult, error) {
	reversal := &alfin.UpdateSerializedInventoryTotalsResult{}
	err := c.B.Call(http.MethodPost, "updateSerializedInventoryTotals", c.Key, params, reversal)
	return reversal, err
}
