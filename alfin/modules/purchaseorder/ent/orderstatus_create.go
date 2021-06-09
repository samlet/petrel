// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderstatus"
)

// OrderStatusCreate is the builder for creating a OrderStatus entity.
type OrderStatusCreate struct {
	config
	mutation *OrderStatusMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (osc *OrderStatusCreate) SetCreateTime(t time.Time) *OrderStatusCreate {
	osc.mutation.SetCreateTime(t)
	return osc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableCreateTime(t *time.Time) *OrderStatusCreate {
	if t != nil {
		osc.SetCreateTime(*t)
	}
	return osc
}

// SetUpdateTime sets the "update_time" field.
func (osc *OrderStatusCreate) SetUpdateTime(t time.Time) *OrderStatusCreate {
	osc.mutation.SetUpdateTime(t)
	return osc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableUpdateTime(t *time.Time) *OrderStatusCreate {
	if t != nil {
		osc.SetUpdateTime(*t)
	}
	return osc
}

// SetStatusID sets the "status_id" field.
func (osc *OrderStatusCreate) SetStatusID(i int) *OrderStatusCreate {
	osc.mutation.SetStatusID(i)
	return osc
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableStatusID(i *int) *OrderStatusCreate {
	if i != nil {
		osc.SetStatusID(*i)
	}
	return osc
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (osc *OrderStatusCreate) SetOrderItemSeqID(i int) *OrderStatusCreate {
	osc.mutation.SetOrderItemSeqID(i)
	return osc
}

// SetNillableOrderItemSeqID sets the "order_item_seq_id" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableOrderItemSeqID(i *int) *OrderStatusCreate {
	if i != nil {
		osc.SetOrderItemSeqID(*i)
	}
	return osc
}

// SetOrderPaymentPreferenceID sets the "order_payment_preference_id" field.
func (osc *OrderStatusCreate) SetOrderPaymentPreferenceID(i int) *OrderStatusCreate {
	osc.mutation.SetOrderPaymentPreferenceID(i)
	return osc
}

// SetNillableOrderPaymentPreferenceID sets the "order_payment_preference_id" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableOrderPaymentPreferenceID(i *int) *OrderStatusCreate {
	if i != nil {
		osc.SetOrderPaymentPreferenceID(*i)
	}
	return osc
}

// SetStatusDatetime sets the "status_datetime" field.
func (osc *OrderStatusCreate) SetStatusDatetime(t time.Time) *OrderStatusCreate {
	osc.mutation.SetStatusDatetime(t)
	return osc
}

// SetNillableStatusDatetime sets the "status_datetime" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableStatusDatetime(t *time.Time) *OrderStatusCreate {
	if t != nil {
		osc.SetStatusDatetime(*t)
	}
	return osc
}

// SetStatusUserLogin sets the "status_user_login" field.
func (osc *OrderStatusCreate) SetStatusUserLogin(s string) *OrderStatusCreate {
	osc.mutation.SetStatusUserLogin(s)
	return osc
}

// SetNillableStatusUserLogin sets the "status_user_login" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableStatusUserLogin(s *string) *OrderStatusCreate {
	if s != nil {
		osc.SetStatusUserLogin(*s)
	}
	return osc
}

// SetChangeReason sets the "change_reason" field.
func (osc *OrderStatusCreate) SetChangeReason(s string) *OrderStatusCreate {
	osc.mutation.SetChangeReason(s)
	return osc
}

// SetNillableChangeReason sets the "change_reason" field if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableChangeReason(s *string) *OrderStatusCreate {
	if s != nil {
		osc.SetChangeReason(*s)
	}
	return osc
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (osc *OrderStatusCreate) SetOrderHeaderID(id int) *OrderStatusCreate {
	osc.mutation.SetOrderHeaderID(id)
	return osc
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableOrderHeaderID(id *int) *OrderStatusCreate {
	if id != nil {
		osc = osc.SetOrderHeaderID(*id)
	}
	return osc
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (osc *OrderStatusCreate) SetOrderHeader(o *OrderHeader) *OrderStatusCreate {
	return osc.SetOrderHeaderID(o.ID)
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (osc *OrderStatusCreate) SetOrderItemID(id int) *OrderStatusCreate {
	osc.mutation.SetOrderItemID(id)
	return osc
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (osc *OrderStatusCreate) SetNillableOrderItemID(id *int) *OrderStatusCreate {
	if id != nil {
		osc = osc.SetOrderItemID(*id)
	}
	return osc
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (osc *OrderStatusCreate) SetOrderItem(o *OrderItem) *OrderStatusCreate {
	return osc.SetOrderItemID(o.ID)
}

// Mutation returns the OrderStatusMutation object of the builder.
func (osc *OrderStatusCreate) Mutation() *OrderStatusMutation {
	return osc.mutation
}

// Save creates the OrderStatus in the database.
func (osc *OrderStatusCreate) Save(ctx context.Context) (*OrderStatus, error) {
	var (
		err  error
		node *OrderStatus
	)
	osc.defaults()
	if len(osc.hooks) == 0 {
		if err = osc.check(); err != nil {
			return nil, err
		}
		node, err = osc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = osc.check(); err != nil {
				return nil, err
			}
			osc.mutation = mutation
			node, err = osc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osc.hooks) - 1; i >= 0; i-- {
			mut = osc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OrderStatusCreate) SaveX(ctx context.Context) *OrderStatus {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (osc *OrderStatusCreate) defaults() {
	if _, ok := osc.mutation.CreateTime(); !ok {
		v := orderstatus.DefaultCreateTime()
		osc.mutation.SetCreateTime(v)
	}
	if _, ok := osc.mutation.UpdateTime(); !ok {
		v := orderstatus.DefaultUpdateTime()
		osc.mutation.SetUpdateTime(v)
	}
	if _, ok := osc.mutation.StatusDatetime(); !ok {
		v := orderstatus.DefaultStatusDatetime()
		osc.mutation.SetStatusDatetime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OrderStatusCreate) check() error {
	if _, ok := osc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := osc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (osc *OrderStatusCreate) sqlSave(ctx context.Context) (*OrderStatus, error) {
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (osc *OrderStatusCreate) createSpec() (*OrderStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderStatus{config: osc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderstatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderstatus.FieldID,
			},
		}
	)
	if value, ok := osc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := osc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := osc.mutation.StatusID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldStatusID,
		})
		_node.StatusID = value
	}
	if value, ok := osc.mutation.OrderItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderItemSeqID,
		})
		_node.OrderItemSeqID = value
	}
	if value, ok := osc.mutation.OrderPaymentPreferenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
		_node.OrderPaymentPreferenceID = value
	}
	if value, ok := osc.mutation.StatusDatetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldStatusDatetime,
		})
		_node.StatusDatetime = value
	}
	if value, ok := osc.mutation.StatusUserLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldStatusUserLogin,
		})
		_node.StatusUserLogin = value
	}
	if value, ok := osc.mutation.ChangeReason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldChangeReason,
		})
		_node.ChangeReason = value
	}
	if nodes := osc.mutation.OrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderstatus.OrderHeaderTable,
			Columns: []string{orderstatus.OrderHeaderColumn},
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
		_node.order_header_order_statuses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := osc.mutation.OrderItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderstatus.OrderItemTable,
			Columns: []string{orderstatus.OrderItemColumn},
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
		_node.order_item_order_statuses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderStatusCreateBulk is the builder for creating many OrderStatus entities in bulk.
type OrderStatusCreateBulk struct {
	config
	builders []*OrderStatusCreate
}

// Save creates the OrderStatus entities in the database.
func (oscb *OrderStatusCreateBulk) Save(ctx context.Context) ([]*OrderStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OrderStatus, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OrderStatusCreateBulk) SaveX(ctx context.Context) []*OrderStatus {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
