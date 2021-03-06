// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderstatus"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/predicate"
)

// OrderStatusUpdate is the builder for updating OrderStatus entities.
type OrderStatusUpdate struct {
	config
	hooks    []Hook
	mutation *OrderStatusMutation
}

// Where adds a new predicate for the OrderStatusUpdate builder.
func (osu *OrderStatusUpdate) Where(ps ...predicate.OrderStatus) *OrderStatusUpdate {
	osu.mutation.predicates = append(osu.mutation.predicates, ps...)
	return osu
}

// SetStatusID sets the "status_id" field.
func (osu *OrderStatusUpdate) SetStatusID(i int) *OrderStatusUpdate {
	osu.mutation.ResetStatusID()
	osu.mutation.SetStatusID(i)
	return osu
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableStatusID(i *int) *OrderStatusUpdate {
	if i != nil {
		osu.SetStatusID(*i)
	}
	return osu
}

// AddStatusID adds i to the "status_id" field.
func (osu *OrderStatusUpdate) AddStatusID(i int) *OrderStatusUpdate {
	osu.mutation.AddStatusID(i)
	return osu
}

// ClearStatusID clears the value of the "status_id" field.
func (osu *OrderStatusUpdate) ClearStatusID() *OrderStatusUpdate {
	osu.mutation.ClearStatusID()
	return osu
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (osu *OrderStatusUpdate) SetOrderItemSeqID(i int) *OrderStatusUpdate {
	osu.mutation.ResetOrderItemSeqID()
	osu.mutation.SetOrderItemSeqID(i)
	return osu
}

// SetNillableOrderItemSeqID sets the "order_item_seq_id" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableOrderItemSeqID(i *int) *OrderStatusUpdate {
	if i != nil {
		osu.SetOrderItemSeqID(*i)
	}
	return osu
}

// AddOrderItemSeqID adds i to the "order_item_seq_id" field.
func (osu *OrderStatusUpdate) AddOrderItemSeqID(i int) *OrderStatusUpdate {
	osu.mutation.AddOrderItemSeqID(i)
	return osu
}

// ClearOrderItemSeqID clears the value of the "order_item_seq_id" field.
func (osu *OrderStatusUpdate) ClearOrderItemSeqID() *OrderStatusUpdate {
	osu.mutation.ClearOrderItemSeqID()
	return osu
}

// SetOrderPaymentPreferenceID sets the "order_payment_preference_id" field.
func (osu *OrderStatusUpdate) SetOrderPaymentPreferenceID(i int) *OrderStatusUpdate {
	osu.mutation.ResetOrderPaymentPreferenceID()
	osu.mutation.SetOrderPaymentPreferenceID(i)
	return osu
}

// SetNillableOrderPaymentPreferenceID sets the "order_payment_preference_id" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableOrderPaymentPreferenceID(i *int) *OrderStatusUpdate {
	if i != nil {
		osu.SetOrderPaymentPreferenceID(*i)
	}
	return osu
}

// AddOrderPaymentPreferenceID adds i to the "order_payment_preference_id" field.
func (osu *OrderStatusUpdate) AddOrderPaymentPreferenceID(i int) *OrderStatusUpdate {
	osu.mutation.AddOrderPaymentPreferenceID(i)
	return osu
}

// ClearOrderPaymentPreferenceID clears the value of the "order_payment_preference_id" field.
func (osu *OrderStatusUpdate) ClearOrderPaymentPreferenceID() *OrderStatusUpdate {
	osu.mutation.ClearOrderPaymentPreferenceID()
	return osu
}

// SetStatusDatetime sets the "status_datetime" field.
func (osu *OrderStatusUpdate) SetStatusDatetime(t time.Time) *OrderStatusUpdate {
	osu.mutation.SetStatusDatetime(t)
	return osu
}

// SetNillableStatusDatetime sets the "status_datetime" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableStatusDatetime(t *time.Time) *OrderStatusUpdate {
	if t != nil {
		osu.SetStatusDatetime(*t)
	}
	return osu
}

// ClearStatusDatetime clears the value of the "status_datetime" field.
func (osu *OrderStatusUpdate) ClearStatusDatetime() *OrderStatusUpdate {
	osu.mutation.ClearStatusDatetime()
	return osu
}

// SetStatusUserLogin sets the "status_user_login" field.
func (osu *OrderStatusUpdate) SetStatusUserLogin(s string) *OrderStatusUpdate {
	osu.mutation.SetStatusUserLogin(s)
	return osu
}

// SetNillableStatusUserLogin sets the "status_user_login" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableStatusUserLogin(s *string) *OrderStatusUpdate {
	if s != nil {
		osu.SetStatusUserLogin(*s)
	}
	return osu
}

// ClearStatusUserLogin clears the value of the "status_user_login" field.
func (osu *OrderStatusUpdate) ClearStatusUserLogin() *OrderStatusUpdate {
	osu.mutation.ClearStatusUserLogin()
	return osu
}

// SetChangeReason sets the "change_reason" field.
func (osu *OrderStatusUpdate) SetChangeReason(s string) *OrderStatusUpdate {
	osu.mutation.SetChangeReason(s)
	return osu
}

// SetNillableChangeReason sets the "change_reason" field if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableChangeReason(s *string) *OrderStatusUpdate {
	if s != nil {
		osu.SetChangeReason(*s)
	}
	return osu
}

// ClearChangeReason clears the value of the "change_reason" field.
func (osu *OrderStatusUpdate) ClearChangeReason() *OrderStatusUpdate {
	osu.mutation.ClearChangeReason()
	return osu
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (osu *OrderStatusUpdate) SetOrderHeaderID(id int) *OrderStatusUpdate {
	osu.mutation.SetOrderHeaderID(id)
	return osu
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableOrderHeaderID(id *int) *OrderStatusUpdate {
	if id != nil {
		osu = osu.SetOrderHeaderID(*id)
	}
	return osu
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (osu *OrderStatusUpdate) SetOrderHeader(o *OrderHeader) *OrderStatusUpdate {
	return osu.SetOrderHeaderID(o.ID)
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (osu *OrderStatusUpdate) SetOrderItemID(id int) *OrderStatusUpdate {
	osu.mutation.SetOrderItemID(id)
	return osu
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (osu *OrderStatusUpdate) SetNillableOrderItemID(id *int) *OrderStatusUpdate {
	if id != nil {
		osu = osu.SetOrderItemID(*id)
	}
	return osu
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (osu *OrderStatusUpdate) SetOrderItem(o *OrderItem) *OrderStatusUpdate {
	return osu.SetOrderItemID(o.ID)
}

// Mutation returns the OrderStatusMutation object of the builder.
func (osu *OrderStatusUpdate) Mutation() *OrderStatusMutation {
	return osu.mutation
}

// ClearOrderHeader clears the "order_header" edge to the OrderHeader entity.
func (osu *OrderStatusUpdate) ClearOrderHeader() *OrderStatusUpdate {
	osu.mutation.ClearOrderHeader()
	return osu
}

// ClearOrderItem clears the "order_item" edge to the OrderItem entity.
func (osu *OrderStatusUpdate) ClearOrderItem() *OrderStatusUpdate {
	osu.mutation.ClearOrderItem()
	return osu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrderStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	osu.defaults()
	if len(osu.hooks) == 0 {
		affected, err = osu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osu.mutation = mutation
			affected, err = osu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osu.hooks) - 1; i >= 0; i-- {
			mut = osu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrderStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrderStatusUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrderStatusUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OrderStatusUpdate) defaults() {
	if _, ok := osu.mutation.UpdateTime(); !ok {
		v := orderstatus.UpdateDefaultUpdateTime()
		osu.mutation.SetUpdateTime(v)
	}
}

func (osu *OrderStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatus.Table,
			Columns: orderstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderstatus.FieldID,
			},
		},
	}
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldUpdateTime,
		})
	}
	if value, ok := osu.mutation.StatusID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldStatusID,
		})
	}
	if value, ok := osu.mutation.AddedStatusID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldStatusID,
		})
	}
	if osu.mutation.StatusIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldStatusID,
		})
	}
	if value, ok := osu.mutation.OrderItemSeqID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if value, ok := osu.mutation.AddedOrderItemSeqID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if osu.mutation.OrderItemSeqIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if value, ok := osu.mutation.OrderPaymentPreferenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if value, ok := osu.mutation.AddedOrderPaymentPreferenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if osu.mutation.OrderPaymentPreferenceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if value, ok := osu.mutation.StatusDatetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldStatusDatetime,
		})
	}
	if osu.mutation.StatusDatetimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: orderstatus.FieldStatusDatetime,
		})
	}
	if value, ok := osu.mutation.StatusUserLogin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldStatusUserLogin,
		})
	}
	if osu.mutation.StatusUserLoginCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatus.FieldStatusUserLogin,
		})
	}
	if value, ok := osu.mutation.ChangeReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldChangeReason,
		})
	}
	if osu.mutation.ChangeReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatus.FieldChangeReason,
		})
	}
	if osu.mutation.OrderHeaderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrderHeaderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osu.mutation.OrderItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrderItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderStatusUpdateOne is the builder for updating a single OrderStatus entity.
type OrderStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderStatusMutation
}

// SetStatusID sets the "status_id" field.
func (osuo *OrderStatusUpdateOne) SetStatusID(i int) *OrderStatusUpdateOne {
	osuo.mutation.ResetStatusID()
	osuo.mutation.SetStatusID(i)
	return osuo
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableStatusID(i *int) *OrderStatusUpdateOne {
	if i != nil {
		osuo.SetStatusID(*i)
	}
	return osuo
}

// AddStatusID adds i to the "status_id" field.
func (osuo *OrderStatusUpdateOne) AddStatusID(i int) *OrderStatusUpdateOne {
	osuo.mutation.AddStatusID(i)
	return osuo
}

// ClearStatusID clears the value of the "status_id" field.
func (osuo *OrderStatusUpdateOne) ClearStatusID() *OrderStatusUpdateOne {
	osuo.mutation.ClearStatusID()
	return osuo
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (osuo *OrderStatusUpdateOne) SetOrderItemSeqID(i int) *OrderStatusUpdateOne {
	osuo.mutation.ResetOrderItemSeqID()
	osuo.mutation.SetOrderItemSeqID(i)
	return osuo
}

// SetNillableOrderItemSeqID sets the "order_item_seq_id" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableOrderItemSeqID(i *int) *OrderStatusUpdateOne {
	if i != nil {
		osuo.SetOrderItemSeqID(*i)
	}
	return osuo
}

// AddOrderItemSeqID adds i to the "order_item_seq_id" field.
func (osuo *OrderStatusUpdateOne) AddOrderItemSeqID(i int) *OrderStatusUpdateOne {
	osuo.mutation.AddOrderItemSeqID(i)
	return osuo
}

// ClearOrderItemSeqID clears the value of the "order_item_seq_id" field.
func (osuo *OrderStatusUpdateOne) ClearOrderItemSeqID() *OrderStatusUpdateOne {
	osuo.mutation.ClearOrderItemSeqID()
	return osuo
}

// SetOrderPaymentPreferenceID sets the "order_payment_preference_id" field.
func (osuo *OrderStatusUpdateOne) SetOrderPaymentPreferenceID(i int) *OrderStatusUpdateOne {
	osuo.mutation.ResetOrderPaymentPreferenceID()
	osuo.mutation.SetOrderPaymentPreferenceID(i)
	return osuo
}

// SetNillableOrderPaymentPreferenceID sets the "order_payment_preference_id" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableOrderPaymentPreferenceID(i *int) *OrderStatusUpdateOne {
	if i != nil {
		osuo.SetOrderPaymentPreferenceID(*i)
	}
	return osuo
}

// AddOrderPaymentPreferenceID adds i to the "order_payment_preference_id" field.
func (osuo *OrderStatusUpdateOne) AddOrderPaymentPreferenceID(i int) *OrderStatusUpdateOne {
	osuo.mutation.AddOrderPaymentPreferenceID(i)
	return osuo
}

// ClearOrderPaymentPreferenceID clears the value of the "order_payment_preference_id" field.
func (osuo *OrderStatusUpdateOne) ClearOrderPaymentPreferenceID() *OrderStatusUpdateOne {
	osuo.mutation.ClearOrderPaymentPreferenceID()
	return osuo
}

// SetStatusDatetime sets the "status_datetime" field.
func (osuo *OrderStatusUpdateOne) SetStatusDatetime(t time.Time) *OrderStatusUpdateOne {
	osuo.mutation.SetStatusDatetime(t)
	return osuo
}

// SetNillableStatusDatetime sets the "status_datetime" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableStatusDatetime(t *time.Time) *OrderStatusUpdateOne {
	if t != nil {
		osuo.SetStatusDatetime(*t)
	}
	return osuo
}

// ClearStatusDatetime clears the value of the "status_datetime" field.
func (osuo *OrderStatusUpdateOne) ClearStatusDatetime() *OrderStatusUpdateOne {
	osuo.mutation.ClearStatusDatetime()
	return osuo
}

// SetStatusUserLogin sets the "status_user_login" field.
func (osuo *OrderStatusUpdateOne) SetStatusUserLogin(s string) *OrderStatusUpdateOne {
	osuo.mutation.SetStatusUserLogin(s)
	return osuo
}

// SetNillableStatusUserLogin sets the "status_user_login" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableStatusUserLogin(s *string) *OrderStatusUpdateOne {
	if s != nil {
		osuo.SetStatusUserLogin(*s)
	}
	return osuo
}

// ClearStatusUserLogin clears the value of the "status_user_login" field.
func (osuo *OrderStatusUpdateOne) ClearStatusUserLogin() *OrderStatusUpdateOne {
	osuo.mutation.ClearStatusUserLogin()
	return osuo
}

// SetChangeReason sets the "change_reason" field.
func (osuo *OrderStatusUpdateOne) SetChangeReason(s string) *OrderStatusUpdateOne {
	osuo.mutation.SetChangeReason(s)
	return osuo
}

// SetNillableChangeReason sets the "change_reason" field if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableChangeReason(s *string) *OrderStatusUpdateOne {
	if s != nil {
		osuo.SetChangeReason(*s)
	}
	return osuo
}

// ClearChangeReason clears the value of the "change_reason" field.
func (osuo *OrderStatusUpdateOne) ClearChangeReason() *OrderStatusUpdateOne {
	osuo.mutation.ClearChangeReason()
	return osuo
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (osuo *OrderStatusUpdateOne) SetOrderHeaderID(id int) *OrderStatusUpdateOne {
	osuo.mutation.SetOrderHeaderID(id)
	return osuo
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableOrderHeaderID(id *int) *OrderStatusUpdateOne {
	if id != nil {
		osuo = osuo.SetOrderHeaderID(*id)
	}
	return osuo
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (osuo *OrderStatusUpdateOne) SetOrderHeader(o *OrderHeader) *OrderStatusUpdateOne {
	return osuo.SetOrderHeaderID(o.ID)
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (osuo *OrderStatusUpdateOne) SetOrderItemID(id int) *OrderStatusUpdateOne {
	osuo.mutation.SetOrderItemID(id)
	return osuo
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (osuo *OrderStatusUpdateOne) SetNillableOrderItemID(id *int) *OrderStatusUpdateOne {
	if id != nil {
		osuo = osuo.SetOrderItemID(*id)
	}
	return osuo
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (osuo *OrderStatusUpdateOne) SetOrderItem(o *OrderItem) *OrderStatusUpdateOne {
	return osuo.SetOrderItemID(o.ID)
}

// Mutation returns the OrderStatusMutation object of the builder.
func (osuo *OrderStatusUpdateOne) Mutation() *OrderStatusMutation {
	return osuo.mutation
}

// ClearOrderHeader clears the "order_header" edge to the OrderHeader entity.
func (osuo *OrderStatusUpdateOne) ClearOrderHeader() *OrderStatusUpdateOne {
	osuo.mutation.ClearOrderHeader()
	return osuo
}

// ClearOrderItem clears the "order_item" edge to the OrderItem entity.
func (osuo *OrderStatusUpdateOne) ClearOrderItem() *OrderStatusUpdateOne {
	osuo.mutation.ClearOrderItem()
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrderStatusUpdateOne) Select(field string, fields ...string) *OrderStatusUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrderStatus entity.
func (osuo *OrderStatusUpdateOne) Save(ctx context.Context) (*OrderStatus, error) {
	var (
		err  error
		node *OrderStatus
	)
	osuo.defaults()
	if len(osuo.hooks) == 0 {
		node, err = osuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osuo.mutation = mutation
			node, err = osuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osuo.hooks) - 1; i >= 0; i-- {
			mut = osuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrderStatusUpdateOne) SaveX(ctx context.Context) *OrderStatus {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrderStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrderStatusUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OrderStatusUpdateOne) defaults() {
	if _, ok := osuo.mutation.UpdateTime(); !ok {
		v := orderstatus.UpdateDefaultUpdateTime()
		osuo.mutation.SetUpdateTime(v)
	}
}

func (osuo *OrderStatusUpdateOne) sqlSave(ctx context.Context) (_node *OrderStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatus.Table,
			Columns: orderstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderstatus.FieldID,
			},
		},
	}
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OrderStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatus.FieldID)
		for _, f := range fields {
			if !orderstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldUpdateTime,
		})
	}
	if value, ok := osuo.mutation.StatusID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldStatusID,
		})
	}
	if value, ok := osuo.mutation.AddedStatusID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldStatusID,
		})
	}
	if osuo.mutation.StatusIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldStatusID,
		})
	}
	if value, ok := osuo.mutation.OrderItemSeqID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if value, ok := osuo.mutation.AddedOrderItemSeqID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if osuo.mutation.OrderItemSeqIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldOrderItemSeqID,
		})
	}
	if value, ok := osuo.mutation.OrderPaymentPreferenceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if value, ok := osuo.mutation.AddedOrderPaymentPreferenceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if osuo.mutation.OrderPaymentPreferenceIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: orderstatus.FieldOrderPaymentPreferenceID,
		})
	}
	if value, ok := osuo.mutation.StatusDatetime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderstatus.FieldStatusDatetime,
		})
	}
	if osuo.mutation.StatusDatetimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: orderstatus.FieldStatusDatetime,
		})
	}
	if value, ok := osuo.mutation.StatusUserLogin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldStatusUserLogin,
		})
	}
	if osuo.mutation.StatusUserLoginCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatus.FieldStatusUserLogin,
		})
	}
	if value, ok := osuo.mutation.ChangeReason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatus.FieldChangeReason,
		})
	}
	if osuo.mutation.ChangeReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatus.FieldChangeReason,
		})
	}
	if osuo.mutation.OrderHeaderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrderHeaderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if osuo.mutation.OrderItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrderItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderStatus{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
