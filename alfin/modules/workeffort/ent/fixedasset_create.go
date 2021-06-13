// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/fixedasset"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortfixedassetassign"
)

// FixedAssetCreate is the builder for creating a FixedAsset entity.
type FixedAssetCreate struct {
	config
	mutation *FixedAssetMutation
	hooks    []Hook
}

// SetFixedAssetTypeID sets the "fixed_asset_type_id" field.
func (fac *FixedAssetCreate) SetFixedAssetTypeID(i int) *FixedAssetCreate {
	fac.mutation.SetFixedAssetTypeID(i)
	return fac
}

// SetNillableFixedAssetTypeID sets the "fixed_asset_type_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableFixedAssetTypeID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetFixedAssetTypeID(*i)
	}
	return fac
}

// SetInstanceOfProductID sets the "instance_of_product_id" field.
func (fac *FixedAssetCreate) SetInstanceOfProductID(i int) *FixedAssetCreate {
	fac.mutation.SetInstanceOfProductID(i)
	return fac
}

// SetNillableInstanceOfProductID sets the "instance_of_product_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableInstanceOfProductID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetInstanceOfProductID(*i)
	}
	return fac
}

// SetClassEnumID sets the "class_enum_id" field.
func (fac *FixedAssetCreate) SetClassEnumID(i int) *FixedAssetCreate {
	fac.mutation.SetClassEnumID(i)
	return fac
}

// SetNillableClassEnumID sets the "class_enum_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableClassEnumID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetClassEnumID(*i)
	}
	return fac
}

// SetRoleTypeID sets the "role_type_id" field.
func (fac *FixedAssetCreate) SetRoleTypeID(i int) *FixedAssetCreate {
	fac.mutation.SetRoleTypeID(i)
	return fac
}

// SetNillableRoleTypeID sets the "role_type_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableRoleTypeID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetRoleTypeID(*i)
	}
	return fac
}

// SetFixedAssetName sets the "fixed_asset_name" field.
func (fac *FixedAssetCreate) SetFixedAssetName(s string) *FixedAssetCreate {
	fac.mutation.SetFixedAssetName(s)
	return fac
}

// SetNillableFixedAssetName sets the "fixed_asset_name" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableFixedAssetName(s *string) *FixedAssetCreate {
	if s != nil {
		fac.SetFixedAssetName(*s)
	}
	return fac
}

// SetAcquireOrderID sets the "acquire_order_id" field.
func (fac *FixedAssetCreate) SetAcquireOrderID(i int) *FixedAssetCreate {
	fac.mutation.SetAcquireOrderID(i)
	return fac
}

// SetNillableAcquireOrderID sets the "acquire_order_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableAcquireOrderID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetAcquireOrderID(*i)
	}
	return fac
}

// SetAcquireOrderItemSeqID sets the "acquire_order_item_seq_id" field.
func (fac *FixedAssetCreate) SetAcquireOrderItemSeqID(i int) *FixedAssetCreate {
	fac.mutation.SetAcquireOrderItemSeqID(i)
	return fac
}

// SetNillableAcquireOrderItemSeqID sets the "acquire_order_item_seq_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableAcquireOrderItemSeqID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetAcquireOrderItemSeqID(*i)
	}
	return fac
}

// SetDateAcquired sets the "date_acquired" field.
func (fac *FixedAssetCreate) SetDateAcquired(t time.Time) *FixedAssetCreate {
	fac.mutation.SetDateAcquired(t)
	return fac
}

// SetNillableDateAcquired sets the "date_acquired" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableDateAcquired(t *time.Time) *FixedAssetCreate {
	if t != nil {
		fac.SetDateAcquired(*t)
	}
	return fac
}

// SetDateLastServiced sets the "date_last_serviced" field.
func (fac *FixedAssetCreate) SetDateLastServiced(t time.Time) *FixedAssetCreate {
	fac.mutation.SetDateLastServiced(t)
	return fac
}

// SetNillableDateLastServiced sets the "date_last_serviced" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableDateLastServiced(t *time.Time) *FixedAssetCreate {
	if t != nil {
		fac.SetDateLastServiced(*t)
	}
	return fac
}

// SetDateNextService sets the "date_next_service" field.
func (fac *FixedAssetCreate) SetDateNextService(t time.Time) *FixedAssetCreate {
	fac.mutation.SetDateNextService(t)
	return fac
}

// SetNillableDateNextService sets the "date_next_service" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableDateNextService(t *time.Time) *FixedAssetCreate {
	if t != nil {
		fac.SetDateNextService(*t)
	}
	return fac
}

// SetExpectedEndOfLife sets the "expected_end_of_life" field.
func (fac *FixedAssetCreate) SetExpectedEndOfLife(t time.Time) *FixedAssetCreate {
	fac.mutation.SetExpectedEndOfLife(t)
	return fac
}

// SetNillableExpectedEndOfLife sets the "expected_end_of_life" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableExpectedEndOfLife(t *time.Time) *FixedAssetCreate {
	if t != nil {
		fac.SetExpectedEndOfLife(*t)
	}
	return fac
}

// SetActualEndOfLife sets the "actual_end_of_life" field.
func (fac *FixedAssetCreate) SetActualEndOfLife(t time.Time) *FixedAssetCreate {
	fac.mutation.SetActualEndOfLife(t)
	return fac
}

// SetNillableActualEndOfLife sets the "actual_end_of_life" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableActualEndOfLife(t *time.Time) *FixedAssetCreate {
	if t != nil {
		fac.SetActualEndOfLife(*t)
	}
	return fac
}

// SetProductionCapacity sets the "production_capacity" field.
func (fac *FixedAssetCreate) SetProductionCapacity(f float64) *FixedAssetCreate {
	fac.mutation.SetProductionCapacity(f)
	return fac
}

// SetNillableProductionCapacity sets the "production_capacity" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableProductionCapacity(f *float64) *FixedAssetCreate {
	if f != nil {
		fac.SetProductionCapacity(*f)
	}
	return fac
}

// SetUomID sets the "uom_id" field.
func (fac *FixedAssetCreate) SetUomID(i int) *FixedAssetCreate {
	fac.mutation.SetUomID(i)
	return fac
}

// SetNillableUomID sets the "uom_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableUomID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetUomID(*i)
	}
	return fac
}

// SetCalendarID sets the "calendar_id" field.
func (fac *FixedAssetCreate) SetCalendarID(i int) *FixedAssetCreate {
	fac.mutation.SetCalendarID(i)
	return fac
}

// SetNillableCalendarID sets the "calendar_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableCalendarID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetCalendarID(*i)
	}
	return fac
}

// SetSerialNumber sets the "serial_number" field.
func (fac *FixedAssetCreate) SetSerialNumber(s string) *FixedAssetCreate {
	fac.mutation.SetSerialNumber(s)
	return fac
}

// SetNillableSerialNumber sets the "serial_number" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableSerialNumber(s *string) *FixedAssetCreate {
	if s != nil {
		fac.SetSerialNumber(*s)
	}
	return fac
}

// SetLocatedAtFacilityID sets the "located_at_facility_id" field.
func (fac *FixedAssetCreate) SetLocatedAtFacilityID(i int) *FixedAssetCreate {
	fac.mutation.SetLocatedAtFacilityID(i)
	return fac
}

// SetNillableLocatedAtFacilityID sets the "located_at_facility_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableLocatedAtFacilityID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetLocatedAtFacilityID(*i)
	}
	return fac
}

// SetLocatedAtLocationSeqID sets the "located_at_location_seq_id" field.
func (fac *FixedAssetCreate) SetLocatedAtLocationSeqID(i int) *FixedAssetCreate {
	fac.mutation.SetLocatedAtLocationSeqID(i)
	return fac
}

// SetNillableLocatedAtLocationSeqID sets the "located_at_location_seq_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableLocatedAtLocationSeqID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetLocatedAtLocationSeqID(*i)
	}
	return fac
}

// SetSalvageValue sets the "salvage_value" field.
func (fac *FixedAssetCreate) SetSalvageValue(f float64) *FixedAssetCreate {
	fac.mutation.SetSalvageValue(f)
	return fac
}

// SetNillableSalvageValue sets the "salvage_value" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableSalvageValue(f *float64) *FixedAssetCreate {
	if f != nil {
		fac.SetSalvageValue(*f)
	}
	return fac
}

// SetDepreciation sets the "depreciation" field.
func (fac *FixedAssetCreate) SetDepreciation(f float64) *FixedAssetCreate {
	fac.mutation.SetDepreciation(f)
	return fac
}

// SetNillableDepreciation sets the "depreciation" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableDepreciation(f *float64) *FixedAssetCreate {
	if f != nil {
		fac.SetDepreciation(*f)
	}
	return fac
}

// SetPurchaseCost sets the "purchase_cost" field.
func (fac *FixedAssetCreate) SetPurchaseCost(f float64) *FixedAssetCreate {
	fac.mutation.SetPurchaseCost(f)
	return fac
}

// SetNillablePurchaseCost sets the "purchase_cost" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillablePurchaseCost(f *float64) *FixedAssetCreate {
	if f != nil {
		fac.SetPurchaseCost(*f)
	}
	return fac
}

// SetPurchaseCostUomID sets the "purchase_cost_uom_id" field.
func (fac *FixedAssetCreate) SetPurchaseCostUomID(i int) *FixedAssetCreate {
	fac.mutation.SetPurchaseCostUomID(i)
	return fac
}

// SetNillablePurchaseCostUomID sets the "purchase_cost_uom_id" field if the given value is not nil.
func (fac *FixedAssetCreate) SetNillablePurchaseCostUomID(i *int) *FixedAssetCreate {
	if i != nil {
		fac.SetPurchaseCostUomID(*i)
	}
	return fac
}

// SetParentID sets the "parent" edge to the FixedAsset entity by ID.
func (fac *FixedAssetCreate) SetParentID(id int) *FixedAssetCreate {
	fac.mutation.SetParentID(id)
	return fac
}

// SetNillableParentID sets the "parent" edge to the FixedAsset entity by ID if the given value is not nil.
func (fac *FixedAssetCreate) SetNillableParentID(id *int) *FixedAssetCreate {
	if id != nil {
		fac = fac.SetParentID(*id)
	}
	return fac
}

// SetParent sets the "parent" edge to the FixedAsset entity.
func (fac *FixedAssetCreate) SetParent(f *FixedAsset) *FixedAssetCreate {
	return fac.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the FixedAsset entity by IDs.
func (fac *FixedAssetCreate) AddChildIDs(ids ...int) *FixedAssetCreate {
	fac.mutation.AddChildIDs(ids...)
	return fac
}

// AddChildren adds the "children" edges to the FixedAsset entity.
func (fac *FixedAssetCreate) AddChildren(f ...*FixedAsset) *FixedAssetCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fac.AddChildIDs(ids...)
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (fac *FixedAssetCreate) SetPartyID(id int) *FixedAssetCreate {
	fac.mutation.SetPartyID(id)
	return fac
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (fac *FixedAssetCreate) SetNillablePartyID(id *int) *FixedAssetCreate {
	if id != nil {
		fac = fac.SetPartyID(*id)
	}
	return fac
}

// SetParty sets the "party" edge to the Party entity.
func (fac *FixedAssetCreate) SetParty(p *Party) *FixedAssetCreate {
	return fac.SetPartyID(p.ID)
}

// SetPartyRoleID sets the "party_role" edge to the PartyRole entity by ID.
func (fac *FixedAssetCreate) SetPartyRoleID(id int) *FixedAssetCreate {
	fac.mutation.SetPartyRoleID(id)
	return fac
}

// SetNillablePartyRoleID sets the "party_role" edge to the PartyRole entity by ID if the given value is not nil.
func (fac *FixedAssetCreate) SetNillablePartyRoleID(id *int) *FixedAssetCreate {
	if id != nil {
		fac = fac.SetPartyRoleID(*id)
	}
	return fac
}

// SetPartyRole sets the "party_role" edge to the PartyRole entity.
func (fac *FixedAssetCreate) SetPartyRole(p *PartyRole) *FixedAssetCreate {
	return fac.SetPartyRoleID(p.ID)
}

// AddChildFixedAssetIDs adds the "child_fixed_assets" edge to the FixedAsset entity by IDs.
func (fac *FixedAssetCreate) AddChildFixedAssetIDs(ids ...int) *FixedAssetCreate {
	fac.mutation.AddChildFixedAssetIDs(ids...)
	return fac
}

// AddChildFixedAssets adds the "child_fixed_assets" edges to the FixedAsset entity.
func (fac *FixedAssetCreate) AddChildFixedAssets(f ...*FixedAsset) *FixedAssetCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fac.AddChildFixedAssetIDs(ids...)
}

// AddWorkEffortIDs adds the "work_efforts" edge to the WorkEffort entity by IDs.
func (fac *FixedAssetCreate) AddWorkEffortIDs(ids ...int) *FixedAssetCreate {
	fac.mutation.AddWorkEffortIDs(ids...)
	return fac
}

// AddWorkEfforts adds the "work_efforts" edges to the WorkEffort entity.
func (fac *FixedAssetCreate) AddWorkEfforts(w ...*WorkEffort) *FixedAssetCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return fac.AddWorkEffortIDs(ids...)
}

// AddWorkEffortFixedAssetAssignIDs adds the "work_effort_fixed_asset_assigns" edge to the WorkEffortFixedAssetAssign entity by IDs.
func (fac *FixedAssetCreate) AddWorkEffortFixedAssetAssignIDs(ids ...int) *FixedAssetCreate {
	fac.mutation.AddWorkEffortFixedAssetAssignIDs(ids...)
	return fac
}

// AddWorkEffortFixedAssetAssigns adds the "work_effort_fixed_asset_assigns" edges to the WorkEffortFixedAssetAssign entity.
func (fac *FixedAssetCreate) AddWorkEffortFixedAssetAssigns(w ...*WorkEffortFixedAssetAssign) *FixedAssetCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return fac.AddWorkEffortFixedAssetAssignIDs(ids...)
}

// Mutation returns the FixedAssetMutation object of the builder.
func (fac *FixedAssetCreate) Mutation() *FixedAssetMutation {
	return fac.mutation
}

// Save creates the FixedAsset in the database.
func (fac *FixedAssetCreate) Save(ctx context.Context) (*FixedAsset, error) {
	var (
		err  error
		node *FixedAsset
	)
	fac.defaults()
	if len(fac.hooks) == 0 {
		if err = fac.check(); err != nil {
			return nil, err
		}
		node, err = fac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FixedAssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fac.check(); err != nil {
				return nil, err
			}
			fac.mutation = mutation
			node, err = fac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fac.hooks) - 1; i >= 0; i-- {
			mut = fac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fac *FixedAssetCreate) SaveX(ctx context.Context) *FixedAsset {
	v, err := fac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fac *FixedAssetCreate) defaults() {
	if _, ok := fac.mutation.DateAcquired(); !ok {
		v := fixedasset.DefaultDateAcquired()
		fac.mutation.SetDateAcquired(v)
	}
	if _, ok := fac.mutation.DateLastServiced(); !ok {
		v := fixedasset.DefaultDateLastServiced()
		fac.mutation.SetDateLastServiced(v)
	}
	if _, ok := fac.mutation.DateNextService(); !ok {
		v := fixedasset.DefaultDateNextService()
		fac.mutation.SetDateNextService(v)
	}
	if _, ok := fac.mutation.ExpectedEndOfLife(); !ok {
		v := fixedasset.DefaultExpectedEndOfLife()
		fac.mutation.SetExpectedEndOfLife(v)
	}
	if _, ok := fac.mutation.ActualEndOfLife(); !ok {
		v := fixedasset.DefaultActualEndOfLife()
		fac.mutation.SetActualEndOfLife(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fac *FixedAssetCreate) check() error {
	return nil
}

func (fac *FixedAssetCreate) sqlSave(ctx context.Context) (*FixedAsset, error) {
	_node, _spec := fac.createSpec()
	if err := sqlgraph.CreateNode(ctx, fac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fac *FixedAssetCreate) createSpec() (*FixedAsset, *sqlgraph.CreateSpec) {
	var (
		_node = &FixedAsset{config: fac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fixedasset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fixedasset.FieldID,
			},
		}
	)
	if value, ok := fac.mutation.FixedAssetTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldFixedAssetTypeID,
		})
		_node.FixedAssetTypeID = value
	}
	if value, ok := fac.mutation.InstanceOfProductID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldInstanceOfProductID,
		})
		_node.InstanceOfProductID = value
	}
	if value, ok := fac.mutation.ClassEnumID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldClassEnumID,
		})
		_node.ClassEnumID = value
	}
	if value, ok := fac.mutation.RoleTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldRoleTypeID,
		})
		_node.RoleTypeID = value
	}
	if value, ok := fac.mutation.FixedAssetName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixedasset.FieldFixedAssetName,
		})
		_node.FixedAssetName = value
	}
	if value, ok := fac.mutation.AcquireOrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldAcquireOrderID,
		})
		_node.AcquireOrderID = value
	}
	if value, ok := fac.mutation.AcquireOrderItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldAcquireOrderItemSeqID,
		})
		_node.AcquireOrderItemSeqID = value
	}
	if value, ok := fac.mutation.DateAcquired(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fixedasset.FieldDateAcquired,
		})
		_node.DateAcquired = value
	}
	if value, ok := fac.mutation.DateLastServiced(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fixedasset.FieldDateLastServiced,
		})
		_node.DateLastServiced = value
	}
	if value, ok := fac.mutation.DateNextService(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fixedasset.FieldDateNextService,
		})
		_node.DateNextService = value
	}
	if value, ok := fac.mutation.ExpectedEndOfLife(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fixedasset.FieldExpectedEndOfLife,
		})
		_node.ExpectedEndOfLife = value
	}
	if value, ok := fac.mutation.ActualEndOfLife(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fixedasset.FieldActualEndOfLife,
		})
		_node.ActualEndOfLife = value
	}
	if value, ok := fac.mutation.ProductionCapacity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fixedasset.FieldProductionCapacity,
		})
		_node.ProductionCapacity = value
	}
	if value, ok := fac.mutation.UomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldUomID,
		})
		_node.UomID = value
	}
	if value, ok := fac.mutation.CalendarID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldCalendarID,
		})
		_node.CalendarID = value
	}
	if value, ok := fac.mutation.SerialNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fixedasset.FieldSerialNumber,
		})
		_node.SerialNumber = value
	}
	if value, ok := fac.mutation.LocatedAtFacilityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldLocatedAtFacilityID,
		})
		_node.LocatedAtFacilityID = value
	}
	if value, ok := fac.mutation.LocatedAtLocationSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldLocatedAtLocationSeqID,
		})
		_node.LocatedAtLocationSeqID = value
	}
	if value, ok := fac.mutation.SalvageValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fixedasset.FieldSalvageValue,
		})
		_node.SalvageValue = value
	}
	if value, ok := fac.mutation.Depreciation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fixedasset.FieldDepreciation,
		})
		_node.Depreciation = value
	}
	if value, ok := fac.mutation.PurchaseCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: fixedasset.FieldPurchaseCost,
		})
		_node.PurchaseCost = value
	}
	if value, ok := fac.mutation.PurchaseCostUomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: fixedasset.FieldPurchaseCostUomID,
		})
		_node.PurchaseCostUomID = value
	}
	if nodes := fac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixedasset.ParentTable,
			Columns: []string{fixedasset.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.fixed_asset_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixedasset.ChildrenTable,
			Columns: []string{fixedasset.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixedasset.PartyTable,
			Columns: []string{fixedasset.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.party_fixed_assets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.PartyRoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fixedasset.PartyRoleTable,
			Columns: []string{fixedasset.PartyRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.party_role_fixed_assets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.ChildFixedAssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   fixedasset.ChildFixedAssetsTable,
			Columns: fixedasset.ChildFixedAssetsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.WorkEffortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixedasset.WorkEffortsTable,
			Columns: []string{fixedasset.WorkEffortsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffort.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fac.mutation.WorkEffortFixedAssetAssignsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fixedasset.WorkEffortFixedAssetAssignsTable,
			Columns: []string{fixedasset.WorkEffortFixedAssetAssignsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortfixedassetassign.FieldID,
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

// FixedAssetCreateBulk is the builder for creating many FixedAsset entities in bulk.
type FixedAssetCreateBulk struct {
	config
	builders []*FixedAssetCreate
}

// Save creates the FixedAsset entities in the database.
func (facb *FixedAssetCreateBulk) Save(ctx context.Context) ([]*FixedAsset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(facb.builders))
	nodes := make([]*FixedAsset, len(facb.builders))
	mutators := make([]Mutator, len(facb.builders))
	for i := range facb.builders {
		func(i int, root context.Context) {
			builder := facb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FixedAssetMutation)
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
					_, err = mutators[i+1].Mutate(root, facb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, facb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, facb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (facb *FixedAssetCreateBulk) SaveX(ctx context.Context) []*FixedAsset {
	v, err := facb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
