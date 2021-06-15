// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/supplierpreforder"
)

// SupplierPrefOrderUpdate is the builder for updating SupplierPrefOrder entities.
type SupplierPrefOrderUpdate struct {
	config
	hooks    []Hook
	mutation *SupplierPrefOrderMutation
}

// Where adds a new predicate for the SupplierPrefOrderUpdate builder.
func (spou *SupplierPrefOrderUpdate) Where(ps ...predicate.SupplierPrefOrder) *SupplierPrefOrderUpdate {
	spou.mutation.predicates = append(spou.mutation.predicates, ps...)
	return spou
}

// SetStringRef sets the "string_ref" field.
func (spou *SupplierPrefOrderUpdate) SetStringRef(s string) *SupplierPrefOrderUpdate {
	spou.mutation.SetStringRef(s)
	return spou
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (spou *SupplierPrefOrderUpdate) SetNillableStringRef(s *string) *SupplierPrefOrderUpdate {
	if s != nil {
		spou.SetStringRef(*s)
	}
	return spou
}

// ClearStringRef clears the value of the "string_ref" field.
func (spou *SupplierPrefOrderUpdate) ClearStringRef() *SupplierPrefOrderUpdate {
	spou.mutation.ClearStringRef()
	return spou
}

// SetDescription sets the "description" field.
func (spou *SupplierPrefOrderUpdate) SetDescription(s string) *SupplierPrefOrderUpdate {
	spou.mutation.SetDescription(s)
	return spou
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (spou *SupplierPrefOrderUpdate) SetNillableDescription(s *string) *SupplierPrefOrderUpdate {
	if s != nil {
		spou.SetDescription(*s)
	}
	return spou
}

// ClearDescription clears the value of the "description" field.
func (spou *SupplierPrefOrderUpdate) ClearDescription() *SupplierPrefOrderUpdate {
	spou.mutation.ClearDescription()
	return spou
}

// Mutation returns the SupplierPrefOrderMutation object of the builder.
func (spou *SupplierPrefOrderUpdate) Mutation() *SupplierPrefOrderMutation {
	return spou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spou *SupplierPrefOrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	spou.defaults()
	if len(spou.hooks) == 0 {
		affected, err = spou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SupplierPrefOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spou.mutation = mutation
			affected, err = spou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spou.hooks) - 1; i >= 0; i-- {
			mut = spou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spou *SupplierPrefOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := spou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spou *SupplierPrefOrderUpdate) Exec(ctx context.Context) error {
	_, err := spou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spou *SupplierPrefOrderUpdate) ExecX(ctx context.Context) {
	if err := spou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spou *SupplierPrefOrderUpdate) defaults() {
	if _, ok := spou.mutation.UpdateTime(); !ok {
		v := supplierpreforder.UpdateDefaultUpdateTime()
		spou.mutation.SetUpdateTime(v)
	}
}

func (spou *SupplierPrefOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   supplierpreforder.Table,
			Columns: supplierpreforder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: supplierpreforder.FieldID,
			},
		},
	}
	if ps := spou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spou.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: supplierpreforder.FieldUpdateTime,
		})
	}
	if value, ok := spou.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: supplierpreforder.FieldStringRef,
		})
	}
	if spou.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: supplierpreforder.FieldStringRef,
		})
	}
	if value, ok := spou.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: supplierpreforder.FieldDescription,
		})
	}
	if spou.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: supplierpreforder.FieldDescription,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{supplierpreforder.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SupplierPrefOrderUpdateOne is the builder for updating a single SupplierPrefOrder entity.
type SupplierPrefOrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SupplierPrefOrderMutation
}

// SetStringRef sets the "string_ref" field.
func (spouo *SupplierPrefOrderUpdateOne) SetStringRef(s string) *SupplierPrefOrderUpdateOne {
	spouo.mutation.SetStringRef(s)
	return spouo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (spouo *SupplierPrefOrderUpdateOne) SetNillableStringRef(s *string) *SupplierPrefOrderUpdateOne {
	if s != nil {
		spouo.SetStringRef(*s)
	}
	return spouo
}

// ClearStringRef clears the value of the "string_ref" field.
func (spouo *SupplierPrefOrderUpdateOne) ClearStringRef() *SupplierPrefOrderUpdateOne {
	spouo.mutation.ClearStringRef()
	return spouo
}

// SetDescription sets the "description" field.
func (spouo *SupplierPrefOrderUpdateOne) SetDescription(s string) *SupplierPrefOrderUpdateOne {
	spouo.mutation.SetDescription(s)
	return spouo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (spouo *SupplierPrefOrderUpdateOne) SetNillableDescription(s *string) *SupplierPrefOrderUpdateOne {
	if s != nil {
		spouo.SetDescription(*s)
	}
	return spouo
}

// ClearDescription clears the value of the "description" field.
func (spouo *SupplierPrefOrderUpdateOne) ClearDescription() *SupplierPrefOrderUpdateOne {
	spouo.mutation.ClearDescription()
	return spouo
}

// Mutation returns the SupplierPrefOrderMutation object of the builder.
func (spouo *SupplierPrefOrderUpdateOne) Mutation() *SupplierPrefOrderMutation {
	return spouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spouo *SupplierPrefOrderUpdateOne) Select(field string, fields ...string) *SupplierPrefOrderUpdateOne {
	spouo.fields = append([]string{field}, fields...)
	return spouo
}

// Save executes the query and returns the updated SupplierPrefOrder entity.
func (spouo *SupplierPrefOrderUpdateOne) Save(ctx context.Context) (*SupplierPrefOrder, error) {
	var (
		err  error
		node *SupplierPrefOrder
	)
	spouo.defaults()
	if len(spouo.hooks) == 0 {
		node, err = spouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SupplierPrefOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spouo.mutation = mutation
			node, err = spouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spouo.hooks) - 1; i >= 0; i-- {
			mut = spouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (spouo *SupplierPrefOrderUpdateOne) SaveX(ctx context.Context) *SupplierPrefOrder {
	node, err := spouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spouo *SupplierPrefOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := spouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spouo *SupplierPrefOrderUpdateOne) ExecX(ctx context.Context) {
	if err := spouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spouo *SupplierPrefOrderUpdateOne) defaults() {
	if _, ok := spouo.mutation.UpdateTime(); !ok {
		v := supplierpreforder.UpdateDefaultUpdateTime()
		spouo.mutation.SetUpdateTime(v)
	}
}

func (spouo *SupplierPrefOrderUpdateOne) sqlSave(ctx context.Context) (_node *SupplierPrefOrder, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   supplierpreforder.Table,
			Columns: supplierpreforder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: supplierpreforder.FieldID,
			},
		},
	}
	id, ok := spouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SupplierPrefOrder.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := spouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, supplierpreforder.FieldID)
		for _, f := range fields {
			if !supplierpreforder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != supplierpreforder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spouo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: supplierpreforder.FieldUpdateTime,
		})
	}
	if value, ok := spouo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: supplierpreforder.FieldStringRef,
		})
	}
	if spouo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: supplierpreforder.FieldStringRef,
		})
	}
	if value, ok := spouo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: supplierpreforder.FieldDescription,
		})
	}
	if spouo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: supplierpreforder.FieldDescription,
		})
	}
	_node = &SupplierPrefOrder{config: spouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{supplierpreforder.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}