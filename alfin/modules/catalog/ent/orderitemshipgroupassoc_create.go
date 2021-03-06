// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderadjustment"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroupassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgrpinvres"
)

// OrderItemShipGroupAssocCreate is the builder for creating a OrderItemShipGroupAssoc entity.
type OrderItemShipGroupAssocCreate struct {
	config
	mutation *OrderItemShipGroupAssocMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetCreateTime(t time.Time) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetCreateTime(t)
	return oisgac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableCreateTime(t *time.Time) *OrderItemShipGroupAssocCreate {
	if t != nil {
		oisgac.SetCreateTime(*t)
	}
	return oisgac
}

// SetUpdateTime sets the "update_time" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetUpdateTime(t time.Time) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetUpdateTime(t)
	return oisgac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableUpdateTime(t *time.Time) *OrderItemShipGroupAssocCreate {
	if t != nil {
		oisgac.SetUpdateTime(*t)
	}
	return oisgac
}

// SetStringRef sets the "string_ref" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetStringRef(s string) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetStringRef(s)
	return oisgac
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableStringRef(s *string) *OrderItemShipGroupAssocCreate {
	if s != nil {
		oisgac.SetStringRef(*s)
	}
	return oisgac
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderItemSeqID(i int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetOrderItemSeqID(i)
	return oisgac
}

// SetShipGroupSeqID sets the "ship_group_seq_id" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetShipGroupSeqID(i int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetShipGroupSeqID(i)
	return oisgac
}

// SetQuantity sets the "quantity" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetQuantity(f float64) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetQuantity(f)
	return oisgac
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableQuantity(f *float64) *OrderItemShipGroupAssocCreate {
	if f != nil {
		oisgac.SetQuantity(*f)
	}
	return oisgac
}

// SetCancelQuantity sets the "cancel_quantity" field.
func (oisgac *OrderItemShipGroupAssocCreate) SetCancelQuantity(f float64) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetCancelQuantity(f)
	return oisgac
}

// SetNillableCancelQuantity sets the "cancel_quantity" field if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableCancelQuantity(f *float64) *OrderItemShipGroupAssocCreate {
	if f != nil {
		oisgac.SetCancelQuantity(*f)
	}
	return oisgac
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderHeaderID(id int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetOrderHeaderID(id)
	return oisgac
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableOrderHeaderID(id *int) *OrderItemShipGroupAssocCreate {
	if id != nil {
		oisgac = oisgac.SetOrderHeaderID(*id)
	}
	return oisgac
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderHeader(o *OrderHeader) *OrderItemShipGroupAssocCreate {
	return oisgac.SetOrderHeaderID(o.ID)
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderItemID(id int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetOrderItemID(id)
	return oisgac
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableOrderItemID(id *int) *OrderItemShipGroupAssocCreate {
	if id != nil {
		oisgac = oisgac.SetOrderItemID(*id)
	}
	return oisgac
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderItem(o *OrderItem) *OrderItemShipGroupAssocCreate {
	return oisgac.SetOrderItemID(o.ID)
}

// SetOrderItemShipGroupID sets the "order_item_ship_group" edge to the OrderItemShipGroup entity by ID.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderItemShipGroupID(id int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.SetOrderItemShipGroupID(id)
	return oisgac
}

// SetNillableOrderItemShipGroupID sets the "order_item_ship_group" edge to the OrderItemShipGroup entity by ID if the given value is not nil.
func (oisgac *OrderItemShipGroupAssocCreate) SetNillableOrderItemShipGroupID(id *int) *OrderItemShipGroupAssocCreate {
	if id != nil {
		oisgac = oisgac.SetOrderItemShipGroupID(*id)
	}
	return oisgac
}

// SetOrderItemShipGroup sets the "order_item_ship_group" edge to the OrderItemShipGroup entity.
func (oisgac *OrderItemShipGroupAssocCreate) SetOrderItemShipGroup(o *OrderItemShipGroup) *OrderItemShipGroupAssocCreate {
	return oisgac.SetOrderItemShipGroupID(o.ID)
}

// AddOrderAdjustmentIDs adds the "order_adjustments" edge to the OrderAdjustment entity by IDs.
func (oisgac *OrderItemShipGroupAssocCreate) AddOrderAdjustmentIDs(ids ...int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.AddOrderAdjustmentIDs(ids...)
	return oisgac
}

// AddOrderAdjustments adds the "order_adjustments" edges to the OrderAdjustment entity.
func (oisgac *OrderItemShipGroupAssocCreate) AddOrderAdjustments(o ...*OrderAdjustment) *OrderItemShipGroupAssocCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oisgac.AddOrderAdjustmentIDs(ids...)
}

// AddOrderItemShipGrpInvReIDs adds the "order_item_ship_grp_inv_res" edge to the OrderItemShipGrpInvRes entity by IDs.
func (oisgac *OrderItemShipGroupAssocCreate) AddOrderItemShipGrpInvReIDs(ids ...int) *OrderItemShipGroupAssocCreate {
	oisgac.mutation.AddOrderItemShipGrpInvReIDs(ids...)
	return oisgac
}

// AddOrderItemShipGrpInvRes adds the "order_item_ship_grp_inv_res" edges to the OrderItemShipGrpInvRes entity.
func (oisgac *OrderItemShipGroupAssocCreate) AddOrderItemShipGrpInvRes(o ...*OrderItemShipGrpInvRes) *OrderItemShipGroupAssocCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oisgac.AddOrderItemShipGrpInvReIDs(ids...)
}

// Mutation returns the OrderItemShipGroupAssocMutation object of the builder.
func (oisgac *OrderItemShipGroupAssocCreate) Mutation() *OrderItemShipGroupAssocMutation {
	return oisgac.mutation
}

// Save creates the OrderItemShipGroupAssoc in the database.
func (oisgac *OrderItemShipGroupAssocCreate) Save(ctx context.Context) (*OrderItemShipGroupAssoc, error) {
	var (
		err  error
		node *OrderItemShipGroupAssoc
	)
	oisgac.defaults()
	if len(oisgac.hooks) == 0 {
		if err = oisgac.check(); err != nil {
			return nil, err
		}
		node, err = oisgac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderItemShipGroupAssocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oisgac.check(); err != nil {
				return nil, err
			}
			oisgac.mutation = mutation
			if node, err = oisgac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oisgac.hooks) - 1; i >= 0; i-- {
			mut = oisgac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oisgac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oisgac *OrderItemShipGroupAssocCreate) SaveX(ctx context.Context) *OrderItemShipGroupAssoc {
	v, err := oisgac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (oisgac *OrderItemShipGroupAssocCreate) defaults() {
	if _, ok := oisgac.mutation.CreateTime(); !ok {
		v := orderitemshipgroupassoc.DefaultCreateTime()
		oisgac.mutation.SetCreateTime(v)
	}
	if _, ok := oisgac.mutation.UpdateTime(); !ok {
		v := orderitemshipgroupassoc.DefaultUpdateTime()
		oisgac.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oisgac *OrderItemShipGroupAssocCreate) check() error {
	if _, ok := oisgac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := oisgac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := oisgac.mutation.OrderItemSeqID(); !ok {
		return &ValidationError{Name: "order_item_seq_id", err: errors.New("ent: missing required field \"order_item_seq_id\"")}
	}
	if _, ok := oisgac.mutation.ShipGroupSeqID(); !ok {
		return &ValidationError{Name: "ship_group_seq_id", err: errors.New("ent: missing required field \"ship_group_seq_id\"")}
	}
	return nil
}

func (oisgac *OrderItemShipGroupAssocCreate) sqlSave(ctx context.Context) (*OrderItemShipGroupAssoc, error) {
	_node, _spec := oisgac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oisgac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oisgac *OrderItemShipGroupAssocCreate) createSpec() (*OrderItemShipGroupAssoc, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItemShipGroupAssoc{config: oisgac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderitemshipgroupassoc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderitemshipgroupassoc.FieldID,
			},
		}
	)
	if value, ok := oisgac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oisgac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oisgac.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := oisgac.mutation.OrderItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldOrderItemSeqID,
		})
		_node.OrderItemSeqID = value
	}
	if value, ok := oisgac.mutation.ShipGroupSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldShipGroupSeqID,
		})
		_node.ShipGroupSeqID = value
	}
	if value, ok := oisgac.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldQuantity,
		})
		_node.Quantity = value
	}
	if value, ok := oisgac.mutation.CancelQuantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderitemshipgroupassoc.FieldCancelQuantity,
		})
		_node.CancelQuantity = value
	}
	if nodes := oisgac.mutation.OrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitemshipgroupassoc.OrderHeaderTable,
			Columns: []string{orderitemshipgroupassoc.OrderHeaderColumn},
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
		_node.order_header_order_item_ship_group_assocs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oisgac.mutation.OrderItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitemshipgroupassoc.OrderItemTable,
			Columns: []string{orderitemshipgroupassoc.OrderItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_order_item_ship_group_assocs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oisgac.mutation.OrderItemShipGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitemshipgroupassoc.OrderItemShipGroupTable,
			Columns: []string{orderitemshipgroupassoc.OrderItemShipGroupColumn},
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
		_node.order_item_ship_group_order_item_ship_group_assocs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oisgac.mutation.OrderAdjustmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderitemshipgroupassoc.OrderAdjustmentsTable,
			Columns: []string{orderitemshipgroupassoc.OrderAdjustmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderadjustment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oisgac.mutation.OrderItemShipGrpInvResIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderitemshipgroupassoc.OrderItemShipGrpInvResTable,
			Columns: []string{orderitemshipgroupassoc.OrderItemShipGrpInvResColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderItemShipGroupAssocCreateBulk is the builder for creating many OrderItemShipGroupAssoc entities in bulk.
type OrderItemShipGroupAssocCreateBulk struct {
	config
	builders []*OrderItemShipGroupAssocCreate
}

// Save creates the OrderItemShipGroupAssoc entities in the database.
func (oisgacb *OrderItemShipGroupAssocCreateBulk) Save(ctx context.Context) ([]*OrderItemShipGroupAssoc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oisgacb.builders))
	nodes := make([]*OrderItemShipGroupAssoc, len(oisgacb.builders))
	mutators := make([]Mutator, len(oisgacb.builders))
	for i := range oisgacb.builders {
		func(i int, root context.Context) {
			builder := oisgacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemShipGroupAssocMutation)
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
					_, err = mutators[i+1].Mutate(root, oisgacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oisgacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oisgacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oisgacb *OrderItemShipGroupAssocCreateBulk) SaveX(ctx context.Context) []*OrderItemShipGroupAssoc {
	v, err := oisgacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
