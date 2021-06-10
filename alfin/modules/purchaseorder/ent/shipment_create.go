// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/itemissuance"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderitemshipgroup"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipment"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipmentitem"
)

// ShipmentCreate is the builder for creating a Shipment entity.
type ShipmentCreate struct {
	config
	mutation *ShipmentMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *ShipmentCreate) SetCreateTime(t time.Time) *ShipmentCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableCreateTime(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ShipmentCreate) SetUpdateTime(t time.Time) *ShipmentCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableUpdateTime(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetShipmentTypeID sets the "shipment_type_id" field.
func (sc *ShipmentCreate) SetShipmentTypeID(i int) *ShipmentCreate {
	sc.mutation.SetShipmentTypeID(i)
	return sc
}

// SetNillableShipmentTypeID sets the "shipment_type_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableShipmentTypeID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetShipmentTypeID(*i)
	}
	return sc
}

// SetStatusID sets the "status_id" field.
func (sc *ShipmentCreate) SetStatusID(i int) *ShipmentCreate {
	sc.mutation.SetStatusID(i)
	return sc
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableStatusID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetStatusID(*i)
	}
	return sc
}

// SetPrimaryReturnID sets the "primary_return_id" field.
func (sc *ShipmentCreate) SetPrimaryReturnID(i int) *ShipmentCreate {
	sc.mutation.SetPrimaryReturnID(i)
	return sc
}

// SetNillablePrimaryReturnID sets the "primary_return_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePrimaryReturnID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetPrimaryReturnID(*i)
	}
	return sc
}

// SetPrimaryShipGroupSeqID sets the "primary_ship_group_seq_id" field.
func (sc *ShipmentCreate) SetPrimaryShipGroupSeqID(i int) *ShipmentCreate {
	sc.mutation.SetPrimaryShipGroupSeqID(i)
	return sc
}

// SetNillablePrimaryShipGroupSeqID sets the "primary_ship_group_seq_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePrimaryShipGroupSeqID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetPrimaryShipGroupSeqID(*i)
	}
	return sc
}

// SetPicklistBinID sets the "picklist_bin_id" field.
func (sc *ShipmentCreate) SetPicklistBinID(i int) *ShipmentCreate {
	sc.mutation.SetPicklistBinID(i)
	return sc
}

// SetNillablePicklistBinID sets the "picklist_bin_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePicklistBinID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetPicklistBinID(*i)
	}
	return sc
}

// SetEstimatedReadyDate sets the "estimated_ready_date" field.
func (sc *ShipmentCreate) SetEstimatedReadyDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetEstimatedReadyDate(t)
	return sc
}

// SetNillableEstimatedReadyDate sets the "estimated_ready_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedReadyDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetEstimatedReadyDate(*t)
	}
	return sc
}

// SetEstimatedShipDate sets the "estimated_ship_date" field.
func (sc *ShipmentCreate) SetEstimatedShipDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetEstimatedShipDate(t)
	return sc
}

// SetNillableEstimatedShipDate sets the "estimated_ship_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedShipDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetEstimatedShipDate(*t)
	}
	return sc
}

// SetEstimatedShipWorkEffID sets the "estimated_ship_work_eff_id" field.
func (sc *ShipmentCreate) SetEstimatedShipWorkEffID(i int) *ShipmentCreate {
	sc.mutation.SetEstimatedShipWorkEffID(i)
	return sc
}

// SetNillableEstimatedShipWorkEffID sets the "estimated_ship_work_eff_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedShipWorkEffID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetEstimatedShipWorkEffID(*i)
	}
	return sc
}

// SetEstimatedArrivalDate sets the "estimated_arrival_date" field.
func (sc *ShipmentCreate) SetEstimatedArrivalDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetEstimatedArrivalDate(t)
	return sc
}

// SetNillableEstimatedArrivalDate sets the "estimated_arrival_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedArrivalDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetEstimatedArrivalDate(*t)
	}
	return sc
}

// SetEstimatedArrivalWorkEffID sets the "estimated_arrival_work_eff_id" field.
func (sc *ShipmentCreate) SetEstimatedArrivalWorkEffID(i int) *ShipmentCreate {
	sc.mutation.SetEstimatedArrivalWorkEffID(i)
	return sc
}

// SetNillableEstimatedArrivalWorkEffID sets the "estimated_arrival_work_eff_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedArrivalWorkEffID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetEstimatedArrivalWorkEffID(*i)
	}
	return sc
}

// SetLatestCancelDate sets the "latest_cancel_date" field.
func (sc *ShipmentCreate) SetLatestCancelDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetLatestCancelDate(t)
	return sc
}

// SetNillableLatestCancelDate sets the "latest_cancel_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableLatestCancelDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetLatestCancelDate(*t)
	}
	return sc
}

// SetEstimatedShipCost sets the "estimated_ship_cost" field.
func (sc *ShipmentCreate) SetEstimatedShipCost(f float64) *ShipmentCreate {
	sc.mutation.SetEstimatedShipCost(f)
	return sc
}

// SetNillableEstimatedShipCost sets the "estimated_ship_cost" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableEstimatedShipCost(f *float64) *ShipmentCreate {
	if f != nil {
		sc.SetEstimatedShipCost(*f)
	}
	return sc
}

// SetCurrencyUomID sets the "currency_uom_id" field.
func (sc *ShipmentCreate) SetCurrencyUomID(i int) *ShipmentCreate {
	sc.mutation.SetCurrencyUomID(i)
	return sc
}

// SetNillableCurrencyUomID sets the "currency_uom_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableCurrencyUomID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetCurrencyUomID(*i)
	}
	return sc
}

// SetHandlingInstructions sets the "handling_instructions" field.
func (sc *ShipmentCreate) SetHandlingInstructions(s string) *ShipmentCreate {
	sc.mutation.SetHandlingInstructions(s)
	return sc
}

// SetNillableHandlingInstructions sets the "handling_instructions" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableHandlingInstructions(s *string) *ShipmentCreate {
	if s != nil {
		sc.SetHandlingInstructions(*s)
	}
	return sc
}

// SetOriginFacilityID sets the "origin_facility_id" field.
func (sc *ShipmentCreate) SetOriginFacilityID(i int) *ShipmentCreate {
	sc.mutation.SetOriginFacilityID(i)
	return sc
}

// SetNillableOriginFacilityID sets the "origin_facility_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableOriginFacilityID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetOriginFacilityID(*i)
	}
	return sc
}

// SetDestinationFacilityID sets the "destination_facility_id" field.
func (sc *ShipmentCreate) SetDestinationFacilityID(i int) *ShipmentCreate {
	sc.mutation.SetDestinationFacilityID(i)
	return sc
}

// SetNillableDestinationFacilityID sets the "destination_facility_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableDestinationFacilityID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetDestinationFacilityID(*i)
	}
	return sc
}

// SetOriginContactMechID sets the "origin_contact_mech_id" field.
func (sc *ShipmentCreate) SetOriginContactMechID(i int) *ShipmentCreate {
	sc.mutation.SetOriginContactMechID(i)
	return sc
}

// SetNillableOriginContactMechID sets the "origin_contact_mech_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableOriginContactMechID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetOriginContactMechID(*i)
	}
	return sc
}

// SetOriginTelecomNumberID sets the "origin_telecom_number_id" field.
func (sc *ShipmentCreate) SetOriginTelecomNumberID(i int) *ShipmentCreate {
	sc.mutation.SetOriginTelecomNumberID(i)
	return sc
}

// SetNillableOriginTelecomNumberID sets the "origin_telecom_number_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableOriginTelecomNumberID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetOriginTelecomNumberID(*i)
	}
	return sc
}

// SetDestinationContactMechID sets the "destination_contact_mech_id" field.
func (sc *ShipmentCreate) SetDestinationContactMechID(i int) *ShipmentCreate {
	sc.mutation.SetDestinationContactMechID(i)
	return sc
}

// SetNillableDestinationContactMechID sets the "destination_contact_mech_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableDestinationContactMechID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetDestinationContactMechID(*i)
	}
	return sc
}

// SetDestinationTelecomNumberID sets the "destination_telecom_number_id" field.
func (sc *ShipmentCreate) SetDestinationTelecomNumberID(i int) *ShipmentCreate {
	sc.mutation.SetDestinationTelecomNumberID(i)
	return sc
}

// SetNillableDestinationTelecomNumberID sets the "destination_telecom_number_id" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableDestinationTelecomNumberID(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetDestinationTelecomNumberID(*i)
	}
	return sc
}

// SetPartyIDTo sets the "party_id_to" field.
func (sc *ShipmentCreate) SetPartyIDTo(i int) *ShipmentCreate {
	sc.mutation.SetPartyIDTo(i)
	return sc
}

// SetNillablePartyIDTo sets the "party_id_to" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePartyIDTo(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetPartyIDTo(*i)
	}
	return sc
}

// SetPartyIDFrom sets the "party_id_from" field.
func (sc *ShipmentCreate) SetPartyIDFrom(i int) *ShipmentCreate {
	sc.mutation.SetPartyIDFrom(i)
	return sc
}

// SetNillablePartyIDFrom sets the "party_id_from" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePartyIDFrom(i *int) *ShipmentCreate {
	if i != nil {
		sc.SetPartyIDFrom(*i)
	}
	return sc
}

// SetAdditionalShippingCharge sets the "additional_shipping_charge" field.
func (sc *ShipmentCreate) SetAdditionalShippingCharge(f float64) *ShipmentCreate {
	sc.mutation.SetAdditionalShippingCharge(f)
	return sc
}

// SetNillableAdditionalShippingCharge sets the "additional_shipping_charge" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableAdditionalShippingCharge(f *float64) *ShipmentCreate {
	if f != nil {
		sc.SetAdditionalShippingCharge(*f)
	}
	return sc
}

// SetAddtlShippingChargeDesc sets the "addtl_shipping_charge_desc" field.
func (sc *ShipmentCreate) SetAddtlShippingChargeDesc(s string) *ShipmentCreate {
	sc.mutation.SetAddtlShippingChargeDesc(s)
	return sc
}

// SetNillableAddtlShippingChargeDesc sets the "addtl_shipping_charge_desc" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableAddtlShippingChargeDesc(s *string) *ShipmentCreate {
	if s != nil {
		sc.SetAddtlShippingChargeDesc(*s)
	}
	return sc
}

// SetCreatedDate sets the "created_date" field.
func (sc *ShipmentCreate) SetCreatedDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetCreatedDate(t)
	return sc
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableCreatedDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetCreatedDate(*t)
	}
	return sc
}

// SetCreatedByUserLogin sets the "created_by_user_login" field.
func (sc *ShipmentCreate) SetCreatedByUserLogin(s string) *ShipmentCreate {
	sc.mutation.SetCreatedByUserLogin(s)
	return sc
}

// SetNillableCreatedByUserLogin sets the "created_by_user_login" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableCreatedByUserLogin(s *string) *ShipmentCreate {
	if s != nil {
		sc.SetCreatedByUserLogin(*s)
	}
	return sc
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (sc *ShipmentCreate) SetLastModifiedDate(t time.Time) *ShipmentCreate {
	sc.mutation.SetLastModifiedDate(t)
	return sc
}

// SetNillableLastModifiedDate sets the "last_modified_date" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableLastModifiedDate(t *time.Time) *ShipmentCreate {
	if t != nil {
		sc.SetLastModifiedDate(*t)
	}
	return sc
}

// SetLastModifiedByUserLogin sets the "last_modified_by_user_login" field.
func (sc *ShipmentCreate) SetLastModifiedByUserLogin(s string) *ShipmentCreate {
	sc.mutation.SetLastModifiedByUserLogin(s)
	return sc
}

// SetNillableLastModifiedByUserLogin sets the "last_modified_by_user_login" field if the given value is not nil.
func (sc *ShipmentCreate) SetNillableLastModifiedByUserLogin(s *string) *ShipmentCreate {
	if s != nil {
		sc.SetLastModifiedByUserLogin(*s)
	}
	return sc
}

// SetPrimaryOrderHeaderID sets the "primary_order_header" edge to the OrderHeader entity by ID.
func (sc *ShipmentCreate) SetPrimaryOrderHeaderID(id int) *ShipmentCreate {
	sc.mutation.SetPrimaryOrderHeaderID(id)
	return sc
}

// SetNillablePrimaryOrderHeaderID sets the "primary_order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePrimaryOrderHeaderID(id *int) *ShipmentCreate {
	if id != nil {
		sc = sc.SetPrimaryOrderHeaderID(*id)
	}
	return sc
}

// SetPrimaryOrderHeader sets the "primary_order_header" edge to the OrderHeader entity.
func (sc *ShipmentCreate) SetPrimaryOrderHeader(o *OrderHeader) *ShipmentCreate {
	return sc.SetPrimaryOrderHeaderID(o.ID)
}

// SetPrimaryOrderItemShipGroupID sets the "primary_order_item_ship_group" edge to the OrderItemShipGroup entity by ID.
func (sc *ShipmentCreate) SetPrimaryOrderItemShipGroupID(id int) *ShipmentCreate {
	sc.mutation.SetPrimaryOrderItemShipGroupID(id)
	return sc
}

// SetNillablePrimaryOrderItemShipGroupID sets the "primary_order_item_ship_group" edge to the OrderItemShipGroup entity by ID if the given value is not nil.
func (sc *ShipmentCreate) SetNillablePrimaryOrderItemShipGroupID(id *int) *ShipmentCreate {
	if id != nil {
		sc = sc.SetPrimaryOrderItemShipGroupID(*id)
	}
	return sc
}

// SetPrimaryOrderItemShipGroup sets the "primary_order_item_ship_group" edge to the OrderItemShipGroup entity.
func (sc *ShipmentCreate) SetPrimaryOrderItemShipGroup(o *OrderItemShipGroup) *ShipmentCreate {
	return sc.SetPrimaryOrderItemShipGroupID(o.ID)
}

// AddItemIssuanceIDs adds the "item_issuances" edge to the ItemIssuance entity by IDs.
func (sc *ShipmentCreate) AddItemIssuanceIDs(ids ...int) *ShipmentCreate {
	sc.mutation.AddItemIssuanceIDs(ids...)
	return sc
}

// AddItemIssuances adds the "item_issuances" edges to the ItemIssuance entity.
func (sc *ShipmentCreate) AddItemIssuances(i ...*ItemIssuance) *ShipmentCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddItemIssuanceIDs(ids...)
}

// AddShipmentItemIDs adds the "shipment_items" edge to the ShipmentItem entity by IDs.
func (sc *ShipmentCreate) AddShipmentItemIDs(ids ...int) *ShipmentCreate {
	sc.mutation.AddShipmentItemIDs(ids...)
	return sc
}

// AddShipmentItems adds the "shipment_items" edges to the ShipmentItem entity.
func (sc *ShipmentCreate) AddShipmentItems(s ...*ShipmentItem) *ShipmentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddShipmentItemIDs(ids...)
}

// Mutation returns the ShipmentMutation object of the builder.
func (sc *ShipmentCreate) Mutation() *ShipmentMutation {
	return sc.mutation
}

// Save creates the Shipment in the database.
func (sc *ShipmentCreate) Save(ctx context.Context) (*Shipment, error) {
	var (
		err  error
		node *Shipment
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShipmentCreate) SaveX(ctx context.Context) *Shipment {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sc *ShipmentCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := shipment.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := shipment.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.EstimatedReadyDate(); !ok {
		v := shipment.DefaultEstimatedReadyDate()
		sc.mutation.SetEstimatedReadyDate(v)
	}
	if _, ok := sc.mutation.EstimatedShipDate(); !ok {
		v := shipment.DefaultEstimatedShipDate()
		sc.mutation.SetEstimatedShipDate(v)
	}
	if _, ok := sc.mutation.EstimatedArrivalDate(); !ok {
		v := shipment.DefaultEstimatedArrivalDate()
		sc.mutation.SetEstimatedArrivalDate(v)
	}
	if _, ok := sc.mutation.LatestCancelDate(); !ok {
		v := shipment.DefaultLatestCancelDate()
		sc.mutation.SetLatestCancelDate(v)
	}
	if _, ok := sc.mutation.CreatedDate(); !ok {
		v := shipment.DefaultCreatedDate()
		sc.mutation.SetCreatedDate(v)
	}
	if _, ok := sc.mutation.LastModifiedDate(); !ok {
		v := shipment.DefaultLastModifiedDate()
		sc.mutation.SetLastModifiedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShipmentCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (sc *ShipmentCreate) sqlSave(ctx context.Context) (*Shipment, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ShipmentCreate) createSpec() (*Shipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Shipment{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shipment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipment.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.ShipmentTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldShipmentTypeID,
		})
		_node.ShipmentTypeID = value
	}
	if value, ok := sc.mutation.StatusID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldStatusID,
		})
		_node.StatusID = value
	}
	if value, ok := sc.mutation.PrimaryReturnID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldPrimaryReturnID,
		})
		_node.PrimaryReturnID = value
	}
	if value, ok := sc.mutation.PrimaryShipGroupSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldPrimaryShipGroupSeqID,
		})
		_node.PrimaryShipGroupSeqID = value
	}
	if value, ok := sc.mutation.PicklistBinID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldPicklistBinID,
		})
		_node.PicklistBinID = value
	}
	if value, ok := sc.mutation.EstimatedReadyDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldEstimatedReadyDate,
		})
		_node.EstimatedReadyDate = value
	}
	if value, ok := sc.mutation.EstimatedShipDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldEstimatedShipDate,
		})
		_node.EstimatedShipDate = value
	}
	if value, ok := sc.mutation.EstimatedShipWorkEffID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldEstimatedShipWorkEffID,
		})
		_node.EstimatedShipWorkEffID = value
	}
	if value, ok := sc.mutation.EstimatedArrivalDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldEstimatedArrivalDate,
		})
		_node.EstimatedArrivalDate = value
	}
	if value, ok := sc.mutation.EstimatedArrivalWorkEffID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldEstimatedArrivalWorkEffID,
		})
		_node.EstimatedArrivalWorkEffID = value
	}
	if value, ok := sc.mutation.LatestCancelDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldLatestCancelDate,
		})
		_node.LatestCancelDate = value
	}
	if value, ok := sc.mutation.EstimatedShipCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipment.FieldEstimatedShipCost,
		})
		_node.EstimatedShipCost = value
	}
	if value, ok := sc.mutation.CurrencyUomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldCurrencyUomID,
		})
		_node.CurrencyUomID = value
	}
	if value, ok := sc.mutation.HandlingInstructions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipment.FieldHandlingInstructions,
		})
		_node.HandlingInstructions = value
	}
	if value, ok := sc.mutation.OriginFacilityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldOriginFacilityID,
		})
		_node.OriginFacilityID = value
	}
	if value, ok := sc.mutation.DestinationFacilityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldDestinationFacilityID,
		})
		_node.DestinationFacilityID = value
	}
	if value, ok := sc.mutation.OriginContactMechID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldOriginContactMechID,
		})
		_node.OriginContactMechID = value
	}
	if value, ok := sc.mutation.OriginTelecomNumberID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldOriginTelecomNumberID,
		})
		_node.OriginTelecomNumberID = value
	}
	if value, ok := sc.mutation.DestinationContactMechID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldDestinationContactMechID,
		})
		_node.DestinationContactMechID = value
	}
	if value, ok := sc.mutation.DestinationTelecomNumberID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldDestinationTelecomNumberID,
		})
		_node.DestinationTelecomNumberID = value
	}
	if value, ok := sc.mutation.PartyIDTo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldPartyIDTo,
		})
		_node.PartyIDTo = value
	}
	if value, ok := sc.mutation.PartyIDFrom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipment.FieldPartyIDFrom,
		})
		_node.PartyIDFrom = value
	}
	if value, ok := sc.mutation.AdditionalShippingCharge(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipment.FieldAdditionalShippingCharge,
		})
		_node.AdditionalShippingCharge = value
	}
	if value, ok := sc.mutation.AddtlShippingChargeDesc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipment.FieldAddtlShippingChargeDesc,
		})
		_node.AddtlShippingChargeDesc = value
	}
	if value, ok := sc.mutation.CreatedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldCreatedDate,
		})
		_node.CreatedDate = value
	}
	if value, ok := sc.mutation.CreatedByUserLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipment.FieldCreatedByUserLogin,
		})
		_node.CreatedByUserLogin = value
	}
	if value, ok := sc.mutation.LastModifiedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipment.FieldLastModifiedDate,
		})
		_node.LastModifiedDate = value
	}
	if value, ok := sc.mutation.LastModifiedByUserLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipment.FieldLastModifiedByUserLogin,
		})
		_node.LastModifiedByUserLogin = value
	}
	if nodes := sc.mutation.PrimaryOrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shipment.PrimaryOrderHeaderTable,
			Columns: []string{shipment.PrimaryOrderHeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderheader.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_header_primary_shipments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PrimaryOrderItemShipGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shipment.PrimaryOrderItemShipGroupTable,
			Columns: []string{shipment.PrimaryOrderItemShipGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitemshipgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_ship_group_primary_shipments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ItemIssuancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shipment.ItemIssuancesTable,
			Columns: []string{shipment.ItemIssuancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: itemissuance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ShipmentItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shipment.ShipmentItemsTable,
			Columns: []string{shipment.ShipmentItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shipmentitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentCreateBulk is the builder for creating many Shipment entities in bulk.
type ShipmentCreateBulk struct {
	config
	builders []*ShipmentCreate
}

// Save creates the Shipment entities in the database.
func (scb *ShipmentCreateBulk) Save(ctx context.Context) ([]*Shipment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shipment, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShipmentCreateBulk) SaveX(ctx context.Context) []*Shipment {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}