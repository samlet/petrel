// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/orderrole"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/predicate"
)

// OrderRoleUpdate is the builder for updating OrderRole entities.
type OrderRoleUpdate struct {
	config
	hooks    []Hook
	mutation *OrderRoleMutation
}

// Where adds a new predicate for the OrderRoleUpdate builder.
func (oru *OrderRoleUpdate) Where(ps ...predicate.OrderRole) *OrderRoleUpdate {
	oru.mutation.predicates = append(oru.mutation.predicates, ps...)
	return oru
}

// SetPartyID sets the "party_id" field.
func (oru *OrderRoleUpdate) SetPartyID(i int) *OrderRoleUpdate {
	oru.mutation.ResetPartyID()
	oru.mutation.SetPartyID(i)
	return oru
}

// AddPartyID adds i to the "party_id" field.
func (oru *OrderRoleUpdate) AddPartyID(i int) *OrderRoleUpdate {
	oru.mutation.AddPartyID(i)
	return oru
}

// SetRoleTypeID sets the "role_type_id" field.
func (oru *OrderRoleUpdate) SetRoleTypeID(i int) *OrderRoleUpdate {
	oru.mutation.ResetRoleTypeID()
	oru.mutation.SetRoleTypeID(i)
	return oru
}

// AddRoleTypeID adds i to the "role_type_id" field.
func (oru *OrderRoleUpdate) AddRoleTypeID(i int) *OrderRoleUpdate {
	oru.mutation.AddRoleTypeID(i)
	return oru
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (oru *OrderRoleUpdate) SetOrderHeaderID(id int) *OrderRoleUpdate {
	oru.mutation.SetOrderHeaderID(id)
	return oru
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (oru *OrderRoleUpdate) SetNillableOrderHeaderID(id *int) *OrderRoleUpdate {
	if id != nil {
		oru = oru.SetOrderHeaderID(*id)
	}
	return oru
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (oru *OrderRoleUpdate) SetOrderHeader(o *OrderHeader) *OrderRoleUpdate {
	return oru.SetOrderHeaderID(o.ID)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (oru *OrderRoleUpdate) AddOrderItemIDs(ids ...int) *OrderRoleUpdate {
	oru.mutation.AddOrderItemIDs(ids...)
	return oru
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (oru *OrderRoleUpdate) AddOrderItems(o ...*OrderItem) *OrderRoleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oru.AddOrderItemIDs(ids...)
}

// Mutation returns the OrderRoleMutation object of the builder.
func (oru *OrderRoleUpdate) Mutation() *OrderRoleMutation {
	return oru.mutation
}

// ClearOrderHeader clears the "order_header" edge to the OrderHeader entity.
func (oru *OrderRoleUpdate) ClearOrderHeader() *OrderRoleUpdate {
	oru.mutation.ClearOrderHeader()
	return oru
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (oru *OrderRoleUpdate) ClearOrderItems() *OrderRoleUpdate {
	oru.mutation.ClearOrderItems()
	return oru
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (oru *OrderRoleUpdate) RemoveOrderItemIDs(ids ...int) *OrderRoleUpdate {
	oru.mutation.RemoveOrderItemIDs(ids...)
	return oru
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (oru *OrderRoleUpdate) RemoveOrderItems(o ...*OrderItem) *OrderRoleUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oru.RemoveOrderItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oru *OrderRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	oru.defaults()
	if len(oru.hooks) == 0 {
		affected, err = oru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oru.mutation = mutation
			affected, err = oru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oru.hooks) - 1; i >= 0; i-- {
			mut = oru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oru *OrderRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := oru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oru *OrderRoleUpdate) Exec(ctx context.Context) error {
	_, err := oru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oru *OrderRoleUpdate) ExecX(ctx context.Context) {
	if err := oru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oru *OrderRoleUpdate) defaults() {
	if _, ok := oru.mutation.UpdateTime(); !ok {
		v := orderrole.UpdateDefaultUpdateTime()
		oru.mutation.SetUpdateTime(v)
	}
}

func (oru *OrderRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderrole.Table,
			Columns: orderrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderrole.FieldID,
			},
		},
	}
	if ps := oru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderrole.FieldUpdateTime,
		})
	}
	if value, ok := oru.mutation.PartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldPartyID,
		})
	}
	if value, ok := oru.mutation.AddedPartyID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldPartyID,
		})
	}
	if value, ok := oru.mutation.RoleTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldRoleTypeID,
		})
	}
	if value, ok := oru.mutation.AddedRoleTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldRoleTypeID,
		})
	}
	if oru.mutation.OrderHeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderrole.OrderHeaderTable,
			Columns: []string{orderrole.OrderHeaderColumn},
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
	if nodes := oru.mutation.OrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderrole.OrderHeaderTable,
			Columns: []string{orderrole.OrderHeaderColumn},
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
	if oru.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
	if nodes := oru.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !oru.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oru.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, oru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderRoleUpdateOne is the builder for updating a single OrderRole entity.
type OrderRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderRoleMutation
}

// SetPartyID sets the "party_id" field.
func (oruo *OrderRoleUpdateOne) SetPartyID(i int) *OrderRoleUpdateOne {
	oruo.mutation.ResetPartyID()
	oruo.mutation.SetPartyID(i)
	return oruo
}

// AddPartyID adds i to the "party_id" field.
func (oruo *OrderRoleUpdateOne) AddPartyID(i int) *OrderRoleUpdateOne {
	oruo.mutation.AddPartyID(i)
	return oruo
}

// SetRoleTypeID sets the "role_type_id" field.
func (oruo *OrderRoleUpdateOne) SetRoleTypeID(i int) *OrderRoleUpdateOne {
	oruo.mutation.ResetRoleTypeID()
	oruo.mutation.SetRoleTypeID(i)
	return oruo
}

// AddRoleTypeID adds i to the "role_type_id" field.
func (oruo *OrderRoleUpdateOne) AddRoleTypeID(i int) *OrderRoleUpdateOne {
	oruo.mutation.AddRoleTypeID(i)
	return oruo
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (oruo *OrderRoleUpdateOne) SetOrderHeaderID(id int) *OrderRoleUpdateOne {
	oruo.mutation.SetOrderHeaderID(id)
	return oruo
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (oruo *OrderRoleUpdateOne) SetNillableOrderHeaderID(id *int) *OrderRoleUpdateOne {
	if id != nil {
		oruo = oruo.SetOrderHeaderID(*id)
	}
	return oruo
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (oruo *OrderRoleUpdateOne) SetOrderHeader(o *OrderHeader) *OrderRoleUpdateOne {
	return oruo.SetOrderHeaderID(o.ID)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (oruo *OrderRoleUpdateOne) AddOrderItemIDs(ids ...int) *OrderRoleUpdateOne {
	oruo.mutation.AddOrderItemIDs(ids...)
	return oruo
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (oruo *OrderRoleUpdateOne) AddOrderItems(o ...*OrderItem) *OrderRoleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oruo.AddOrderItemIDs(ids...)
}

// Mutation returns the OrderRoleMutation object of the builder.
func (oruo *OrderRoleUpdateOne) Mutation() *OrderRoleMutation {
	return oruo.mutation
}

// ClearOrderHeader clears the "order_header" edge to the OrderHeader entity.
func (oruo *OrderRoleUpdateOne) ClearOrderHeader() *OrderRoleUpdateOne {
	oruo.mutation.ClearOrderHeader()
	return oruo
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (oruo *OrderRoleUpdateOne) ClearOrderItems() *OrderRoleUpdateOne {
	oruo.mutation.ClearOrderItems()
	return oruo
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (oruo *OrderRoleUpdateOne) RemoveOrderItemIDs(ids ...int) *OrderRoleUpdateOne {
	oruo.mutation.RemoveOrderItemIDs(ids...)
	return oruo
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (oruo *OrderRoleUpdateOne) RemoveOrderItems(o ...*OrderItem) *OrderRoleUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oruo.RemoveOrderItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oruo *OrderRoleUpdateOne) Select(field string, fields ...string) *OrderRoleUpdateOne {
	oruo.fields = append([]string{field}, fields...)
	return oruo
}

// Save executes the query and returns the updated OrderRole entity.
func (oruo *OrderRoleUpdateOne) Save(ctx context.Context) (*OrderRole, error) {
	var (
		err  error
		node *OrderRole
	)
	oruo.defaults()
	if len(oruo.hooks) == 0 {
		node, err = oruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oruo.mutation = mutation
			node, err = oruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oruo.hooks) - 1; i >= 0; i-- {
			mut = oruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oruo *OrderRoleUpdateOne) SaveX(ctx context.Context) *OrderRole {
	node, err := oruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oruo *OrderRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := oruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruo *OrderRoleUpdateOne) ExecX(ctx context.Context) {
	if err := oruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oruo *OrderRoleUpdateOne) defaults() {
	if _, ok := oruo.mutation.UpdateTime(); !ok {
		v := orderrole.UpdateDefaultUpdateTime()
		oruo.mutation.SetUpdateTime(v)
	}
}

func (oruo *OrderRoleUpdateOne) sqlSave(ctx context.Context) (_node *OrderRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderrole.Table,
			Columns: orderrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderrole.FieldID,
			},
		},
	}
	id, ok := oruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing OrderRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := oruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderrole.FieldID)
		for _, f := range fields {
			if !orderrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderrole.FieldUpdateTime,
		})
	}
	if value, ok := oruo.mutation.PartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldPartyID,
		})
	}
	if value, ok := oruo.mutation.AddedPartyID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldPartyID,
		})
	}
	if value, ok := oruo.mutation.RoleTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldRoleTypeID,
		})
	}
	if value, ok := oruo.mutation.AddedRoleTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderrole.FieldRoleTypeID,
		})
	}
	if oruo.mutation.OrderHeaderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderrole.OrderHeaderTable,
			Columns: []string{orderrole.OrderHeaderColumn},
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
	if nodes := oruo.mutation.OrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderrole.OrderHeaderTable,
			Columns: []string{orderrole.OrderHeaderColumn},
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
	if oruo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
	if nodes := oruo.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !oruo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oruo.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderrole.OrderItemsTable,
			Columns: []string{orderrole.OrderItemsColumn},
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
	_node = &OrderRole{config: oruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
