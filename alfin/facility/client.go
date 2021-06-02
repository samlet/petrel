// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package facility

import (
    "github.com/samlet/petrel/alfin"
    "github.com/samlet/petrel/services"
    "net/http"
)

type FacilityClient struct {
    B   *services.AlfinBackend
    Key string
}

func newClient(key string) FacilityClient {
    return FacilityClient{services.NewAlfinBackend(), key}
}
// CreateFacility Create a Facility
func (c *FacilityClient) CreateFacility(params *alfin.CreateFacilityParams) (*alfin.CreateFacilityResult, error) {
    reversal := &alfin.CreateFacilityResult{}
    err := c.B.Call(http.MethodPost, "createFacility", c.Key, params, reversal)
    return reversal, err
}
// DeleteFacility Delete a Facility
func (c *FacilityClient) DeleteFacility(params *alfin.DeleteFacilityParams) (*alfin.DeleteFacilityResult, error) {
    reversal := &alfin.DeleteFacilityResult{}
    err := c.B.Call(http.MethodPost, "deleteFacility", c.Key, params, reversal)
    return reversal, err
}
// UpdateFacility Update a Facility
func (c *FacilityClient) UpdateFacility(params *alfin.UpdateFacilityParams) (*alfin.UpdateFacilityResult, error) {
    reversal := &alfin.UpdateFacilityResult{}
    err := c.B.Call(http.MethodPost, "updateFacility", c.Key, params, reversal)
    return reversal, err
}
// GetInventoryAvailableByFacility Get Inventory Availability for a Product constrained by a facilityId
func (c *FacilityClient) GetInventoryAvailableByFacility(params *alfin.GetInventoryAvailableByFacilityParams) (*alfin.GetInventoryAvailableByFacilityResult, error) {
    reversal := &alfin.GetInventoryAvailableByFacilityResult{}
    err := c.B.Call(http.MethodPost, "getInventoryAvailableByFacility", c.Key, params, reversal)
    return reversal, err
}
