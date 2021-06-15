// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/inventoryitemdetail"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgrpinvres"
)

// InventoryItemDetailCreate is the builder for creating a InventoryItemDetail entity.
type InventoryItemDetailCreate struct {
	config
	mutation *InventoryItemDetailMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (iidc *InventoryItemDetailCreate) SetCreateTime(t time.Time) *InventoryItemDetailCreate {
	iidc.mutation.SetCreateTime(t)
	return iidc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableCreateTime(t *time.Time) *InventoryItemDetailCreate {
	if t != nil {
		iidc.SetCreateTime(*t)
	}
	return iidc
}

// SetUpdateTime sets the "update_time" field.
func (iidc *InventoryItemDetailCreate) SetUpdateTime(t time.Time) *InventoryItemDetailCreate {
	iidc.mutation.SetUpdateTime(t)
	return iidc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableUpdateTime(t *time.Time) *InventoryItemDetailCreate {
	if t != nil {
		iidc.SetUpdateTime(*t)
	}
	return iidc
}

// SetStringRef sets the "string_ref" field.
func (iidc *InventoryItemDetailCreate) SetStringRef(s string) *InventoryItemDetailCreate {
	iidc.mutation.SetStringRef(s)
	return iidc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableStringRef(s *string) *InventoryItemDetailCreate {
	if s != nil {
		iidc.SetStringRef(*s)
	}
	return iidc
}

// SetInventoryItemID sets the "inventory_item_id" field.
func (iidc *InventoryItemDetailCreate) SetInventoryItemID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetInventoryItemID(i)
	return iidc
}

// SetInventoryItemDetailSeqID sets the "inventory_item_detail_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetInventoryItemDetailSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetInventoryItemDetailSeqID(i)
	return iidc
}

// SetEffectiveDate sets the "effective_date" field.
func (iidc *InventoryItemDetailCreate) SetEffectiveDate(t time.Time) *InventoryItemDetailCreate {
	iidc.mutation.SetEffectiveDate(t)
	return iidc
}

// SetNillableEffectiveDate sets the "effective_date" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableEffectiveDate(t *time.Time) *InventoryItemDetailCreate {
	if t != nil {
		iidc.SetEffectiveDate(*t)
	}
	return iidc
}

// SetQuantityOnHandDiff sets the "quantity_on_hand_diff" field.
func (iidc *InventoryItemDetailCreate) SetQuantityOnHandDiff(f float64) *InventoryItemDetailCreate {
	iidc.mutation.SetQuantityOnHandDiff(f)
	return iidc
}

// SetNillableQuantityOnHandDiff sets the "quantity_on_hand_diff" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableQuantityOnHandDiff(f *float64) *InventoryItemDetailCreate {
	if f != nil {
		iidc.SetQuantityOnHandDiff(*f)
	}
	return iidc
}

// SetAvailableToPromiseDiff sets the "available_to_promise_diff" field.
func (iidc *InventoryItemDetailCreate) SetAvailableToPromiseDiff(f float64) *InventoryItemDetailCreate {
	iidc.mutation.SetAvailableToPromiseDiff(f)
	return iidc
}

// SetNillableAvailableToPromiseDiff sets the "available_to_promise_diff" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableAvailableToPromiseDiff(f *float64) *InventoryItemDetailCreate {
	if f != nil {
		iidc.SetAvailableToPromiseDiff(*f)
	}
	return iidc
}

// SetAccountingQuantityDiff sets the "accounting_quantity_diff" field.
func (iidc *InventoryItemDetailCreate) SetAccountingQuantityDiff(f float64) *InventoryItemDetailCreate {
	iidc.mutation.SetAccountingQuantityDiff(f)
	return iidc
}

// SetNillableAccountingQuantityDiff sets the "accounting_quantity_diff" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableAccountingQuantityDiff(f *float64) *InventoryItemDetailCreate {
	if f != nil {
		iidc.SetAccountingQuantityDiff(*f)
	}
	return iidc
}

// SetUnitCost sets the "unit_cost" field.
func (iidc *InventoryItemDetailCreate) SetUnitCost(f float64) *InventoryItemDetailCreate {
	iidc.mutation.SetUnitCost(f)
	return iidc
}

// SetNillableUnitCost sets the "unit_cost" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableUnitCost(f *float64) *InventoryItemDetailCreate {
	if f != nil {
		iidc.SetUnitCost(*f)
	}
	return iidc
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetOrderItemSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetOrderItemSeqID(i)
	return iidc
}

// SetNillableOrderItemSeqID sets the "order_item_seq_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableOrderItemSeqID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetOrderItemSeqID(*i)
	}
	return iidc
}

// SetShipGroupSeqID sets the "ship_group_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetShipGroupSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetShipGroupSeqID(i)
	return iidc
}

// SetNillableShipGroupSeqID sets the "ship_group_seq_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableShipGroupSeqID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetShipGroupSeqID(*i)
	}
	return iidc
}

// SetShipmentID sets the "shipment_id" field.
func (iidc *InventoryItemDetailCreate) SetShipmentID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetShipmentID(i)
	return iidc
}

// SetNillableShipmentID sets the "shipment_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableShipmentID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetShipmentID(*i)
	}
	return iidc
}

// SetShipmentItemSeqID sets the "shipment_item_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetShipmentItemSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetShipmentItemSeqID(i)
	return iidc
}

// SetNillableShipmentItemSeqID sets the "shipment_item_seq_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableShipmentItemSeqID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetShipmentItemSeqID(*i)
	}
	return iidc
}

// SetReturnID sets the "return_id" field.
func (iidc *InventoryItemDetailCreate) SetReturnID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetReturnID(i)
	return iidc
}

// SetNillableReturnID sets the "return_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableReturnID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetReturnID(*i)
	}
	return iidc
}

// SetReturnItemSeqID sets the "return_item_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetReturnItemSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetReturnItemSeqID(i)
	return iidc
}

// SetNillableReturnItemSeqID sets the "return_item_seq_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableReturnItemSeqID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetReturnItemSeqID(*i)
	}
	return iidc
}

// SetWorkEffortID sets the "work_effort_id" field.
func (iidc *InventoryItemDetailCreate) SetWorkEffortID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetWorkEffortID(i)
	return iidc
}

// SetNillableWorkEffortID sets the "work_effort_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableWorkEffortID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetWorkEffortID(*i)
	}
	return iidc
}

// SetFixedAssetID sets the "fixed_asset_id" field.
func (iidc *InventoryItemDetailCreate) SetFixedAssetID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetFixedAssetID(i)
	return iidc
}

// SetNillableFixedAssetID sets the "fixed_asset_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableFixedAssetID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetFixedAssetID(*i)
	}
	return iidc
}

// SetMaintHistSeqID sets the "maint_hist_seq_id" field.
func (iidc *InventoryItemDetailCreate) SetMaintHistSeqID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetMaintHistSeqID(i)
	return iidc
}

// SetNillableMaintHistSeqID sets the "maint_hist_seq_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableMaintHistSeqID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetMaintHistSeqID(*i)
	}
	return iidc
}

// SetItemIssuanceID sets the "item_issuance_id" field.
func (iidc *InventoryItemDetailCreate) SetItemIssuanceID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetItemIssuanceID(i)
	return iidc
}

// SetNillableItemIssuanceID sets the "item_issuance_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableItemIssuanceID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetItemIssuanceID(*i)
	}
	return iidc
}

// SetReceiptID sets the "receipt_id" field.
func (iidc *InventoryItemDetailCreate) SetReceiptID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetReceiptID(i)
	return iidc
}

// SetNillableReceiptID sets the "receipt_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableReceiptID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetReceiptID(*i)
	}
	return iidc
}

// SetPhysicalInventoryID sets the "physical_inventory_id" field.
func (iidc *InventoryItemDetailCreate) SetPhysicalInventoryID(i int) *InventoryItemDetailCreate {
	iidc.mutation.SetPhysicalInventoryID(i)
	return iidc
}

// SetNillablePhysicalInventoryID sets the "physical_inventory_id" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillablePhysicalInventoryID(i *int) *InventoryItemDetailCreate {
	if i != nil {
		iidc.SetPhysicalInventoryID(*i)
	}
	return iidc
}

// SetDescription sets the "description" field.
func (iidc *InventoryItemDetailCreate) SetDescription(s string) *InventoryItemDetailCreate {
	iidc.mutation.SetDescription(s)
	return iidc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableDescription(s *string) *InventoryItemDetailCreate {
	if s != nil {
		iidc.SetDescription(*s)
	}
	return iidc
}

// SetOrderItemShipGrpInvResID sets the "order_item_ship_grp_inv_res" edge to the OrderItemShipGrpInvRes entity by ID.
func (iidc *InventoryItemDetailCreate) SetOrderItemShipGrpInvResID(id int) *InventoryItemDetailCreate {
	iidc.mutation.SetOrderItemShipGrpInvResID(id)
	return iidc
}

// SetNillableOrderItemShipGrpInvResID sets the "order_item_ship_grp_inv_res" edge to the OrderItemShipGrpInvRes entity by ID if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableOrderItemShipGrpInvResID(id *int) *InventoryItemDetailCreate {
	if id != nil {
		iidc = iidc.SetOrderItemShipGrpInvResID(*id)
	}
	return iidc
}

// SetOrderItemShipGrpInvRes sets the "order_item_ship_grp_inv_res" edge to the OrderItemShipGrpInvRes entity.
func (iidc *InventoryItemDetailCreate) SetOrderItemShipGrpInvRes(o *OrderItemShipGrpInvRes) *InventoryItemDetailCreate {
	return iidc.SetOrderItemShipGrpInvResID(o.ID)
}

// SetReasonEnumerationID sets the "reason_enumeration" edge to the Enumeration entity by ID.
func (iidc *InventoryItemDetailCreate) SetReasonEnumerationID(id int) *InventoryItemDetailCreate {
	iidc.mutation.SetReasonEnumerationID(id)
	return iidc
}

// SetNillableReasonEnumerationID sets the "reason_enumeration" edge to the Enumeration entity by ID if the given value is not nil.
func (iidc *InventoryItemDetailCreate) SetNillableReasonEnumerationID(id *int) *InventoryItemDetailCreate {
	if id != nil {
		iidc = iidc.SetReasonEnumerationID(*id)
	}
	return iidc
}

// SetReasonEnumeration sets the "reason_enumeration" edge to the Enumeration entity.
func (iidc *InventoryItemDetailCreate) SetReasonEnumeration(e *Enumeration) *InventoryItemDetailCreate {
	return iidc.SetReasonEnumerationID(e.ID)
}

// Mutation returns the InventoryItemDetailMutation object of the builder.
func (iidc *InventoryItemDetailCreate) Mutation() *InventoryItemDetailMutation {
	return iidc.mutation
}

// Save creates the InventoryItemDetail in the database.
func (iidc *InventoryItemDetailCreate) Save(ctx context.Context) (*InventoryItemDetail, error) {
	var (
		err  error
		node *InventoryItemDetail
	)
	iidc.defaults()
	if len(iidc.hooks) == 0 {
		if err = iidc.check(); err != nil {
			return nil, err
		}
		node, err = iidc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InventoryItemDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iidc.check(); err != nil {
				return nil, err
			}
			iidc.mutation = mutation
			if node, err = iidc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(iidc.hooks) - 1; i >= 0; i-- {
			mut = iidc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iidc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iidc *InventoryItemDetailCreate) SaveX(ctx context.Context) *InventoryItemDetail {
	v, err := iidc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (iidc *InventoryItemDetailCreate) defaults() {
	if _, ok := iidc.mutation.CreateTime(); !ok {
		v := inventoryitemdetail.DefaultCreateTime()
		iidc.mutation.SetCreateTime(v)
	}
	if _, ok := iidc.mutation.UpdateTime(); !ok {
		v := inventoryitemdetail.DefaultUpdateTime()
		iidc.mutation.SetUpdateTime(v)
	}
	if _, ok := iidc.mutation.EffectiveDate(); !ok {
		v := inventoryitemdetail.DefaultEffectiveDate()
		iidc.mutation.SetEffectiveDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iidc *InventoryItemDetailCreate) check() error {
	if _, ok := iidc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := iidc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := iidc.mutation.InventoryItemID(); !ok {
		return &ValidationError{Name: "inventory_item_id", err: errors.New("ent: missing required field \"inventory_item_id\"")}
	}
	if _, ok := iidc.mutation.InventoryItemDetailSeqID(); !ok {
		return &ValidationError{Name: "inventory_item_detail_seq_id", err: errors.New("ent: missing required field \"inventory_item_detail_seq_id\"")}
	}
	return nil
}

func (iidc *InventoryItemDetailCreate) sqlSave(ctx context.Context) (*InventoryItemDetail, error) {
	_node, _spec := iidc.createSpec()
	if err := sqlgraph.CreateNode(ctx, iidc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (iidc *InventoryItemDetailCreate) createSpec() (*InventoryItemDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &InventoryItemDetail{config: iidc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: inventoryitemdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inventoryitemdetail.FieldID,
			},
		}
	)
	if value, ok := iidc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inventoryitemdetail.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := iidc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inventoryitemdetail.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := iidc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryitemdetail.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := iidc.mutation.InventoryItemID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldInventoryItemID,
		})
		_node.InventoryItemID = value
	}
	if value, ok := iidc.mutation.InventoryItemDetailSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldInventoryItemDetailSeqID,
		})
		_node.InventoryItemDetailSeqID = value
	}
	if value, ok := iidc.mutation.EffectiveDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inventoryitemdetail.FieldEffectiveDate,
		})
		_node.EffectiveDate = value
	}
	if value, ok := iidc.mutation.QuantityOnHandDiff(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: inventoryitemdetail.FieldQuantityOnHandDiff,
		})
		_node.QuantityOnHandDiff = value
	}
	if value, ok := iidc.mutation.AvailableToPromiseDiff(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: inventoryitemdetail.FieldAvailableToPromiseDiff,
		})
		_node.AvailableToPromiseDiff = value
	}
	if value, ok := iidc.mutation.AccountingQuantityDiff(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: inventoryitemdetail.FieldAccountingQuantityDiff,
		})
		_node.AccountingQuantityDiff = value
	}
	if value, ok := iidc.mutation.UnitCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: inventoryitemdetail.FieldUnitCost,
		})
		_node.UnitCost = value
	}
	if value, ok := iidc.mutation.OrderItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldOrderItemSeqID,
		})
		_node.OrderItemSeqID = value
	}
	if value, ok := iidc.mutation.ShipGroupSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldShipGroupSeqID,
		})
		_node.ShipGroupSeqID = value
	}
	if value, ok := iidc.mutation.ShipmentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldShipmentID,
		})
		_node.ShipmentID = value
	}
	if value, ok := iidc.mutation.ShipmentItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldShipmentItemSeqID,
		})
		_node.ShipmentItemSeqID = value
	}
	if value, ok := iidc.mutation.ReturnID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldReturnID,
		})
		_node.ReturnID = value
	}
	if value, ok := iidc.mutation.ReturnItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldReturnItemSeqID,
		})
		_node.ReturnItemSeqID = value
	}
	if value, ok := iidc.mutation.WorkEffortID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldWorkEffortID,
		})
		_node.WorkEffortID = value
	}
	if value, ok := iidc.mutation.FixedAssetID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldFixedAssetID,
		})
		_node.FixedAssetID = value
	}
	if value, ok := iidc.mutation.MaintHistSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldMaintHistSeqID,
		})
		_node.MaintHistSeqID = value
	}
	if value, ok := iidc.mutation.ItemIssuanceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldItemIssuanceID,
		})
		_node.ItemIssuanceID = value
	}
	if value, ok := iidc.mutation.ReceiptID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldReceiptID,
		})
		_node.ReceiptID = value
	}
	if value, ok := iidc.mutation.PhysicalInventoryID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: inventoryitemdetail.FieldPhysicalInventoryID,
		})
		_node.PhysicalInventoryID = value
	}
	if value, ok := iidc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryitemdetail.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := iidc.mutation.OrderItemShipGrpInvResIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventoryitemdetail.OrderItemShipGrpInvResTable,
			Columns: []string{inventoryitemdetail.OrderItemShipGrpInvResColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitemshipgrpinvres.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_ship_grp_inv_res_inventory_item_details = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iidc.mutation.ReasonEnumerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventoryitemdetail.ReasonEnumerationTable,
			Columns: []string{inventoryitemdetail.ReasonEnumerationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumeration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_reason_inventory_item_details = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InventoryItemDetailCreateBulk is the builder for creating many InventoryItemDetail entities in bulk.
type InventoryItemDetailCreateBulk struct {
	config
	builders []*InventoryItemDetailCreate
}

// Save creates the InventoryItemDetail entities in the database.
func (iidcb *InventoryItemDetailCreateBulk) Save(ctx context.Context) ([]*InventoryItemDetail, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iidcb.builders))
	nodes := make([]*InventoryItemDetail, len(iidcb.builders))
	mutators := make([]Mutator, len(iidcb.builders))
	for i := range iidcb.builders {
		func(i int, root context.Context) {
			builder := iidcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InventoryItemDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, iidcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iidcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
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
		if _, err := mutators[0].Mutate(ctx, iidcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iidcb *InventoryItemDetailCreateBulk) SaveX(ctx context.Context) []*InventoryItemDetail {
	v, err := iidcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}