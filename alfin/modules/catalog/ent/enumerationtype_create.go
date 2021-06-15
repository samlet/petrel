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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/enumerationtype"
)

// EnumerationTypeCreate is the builder for creating a EnumerationType entity.
type EnumerationTypeCreate struct {
	config
	mutation *EnumerationTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (etc *EnumerationTypeCreate) SetCreateTime(t time.Time) *EnumerationTypeCreate {
	etc.mutation.SetCreateTime(t)
	return etc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableCreateTime(t *time.Time) *EnumerationTypeCreate {
	if t != nil {
		etc.SetCreateTime(*t)
	}
	return etc
}

// SetUpdateTime sets the "update_time" field.
func (etc *EnumerationTypeCreate) SetUpdateTime(t time.Time) *EnumerationTypeCreate {
	etc.mutation.SetUpdateTime(t)
	return etc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableUpdateTime(t *time.Time) *EnumerationTypeCreate {
	if t != nil {
		etc.SetUpdateTime(*t)
	}
	return etc
}

// SetStringRef sets the "string_ref" field.
func (etc *EnumerationTypeCreate) SetStringRef(s string) *EnumerationTypeCreate {
	etc.mutation.SetStringRef(s)
	return etc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableStringRef(s *string) *EnumerationTypeCreate {
	if s != nil {
		etc.SetStringRef(*s)
	}
	return etc
}

// SetHasTable sets the "has_table" field.
func (etc *EnumerationTypeCreate) SetHasTable(et enumerationtype.HasTable) *EnumerationTypeCreate {
	etc.mutation.SetHasTable(et)
	return etc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableHasTable(et *enumerationtype.HasTable) *EnumerationTypeCreate {
	if et != nil {
		etc.SetHasTable(*et)
	}
	return etc
}

// SetDescription sets the "description" field.
func (etc *EnumerationTypeCreate) SetDescription(s string) *EnumerationTypeCreate {
	etc.mutation.SetDescription(s)
	return etc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableDescription(s *string) *EnumerationTypeCreate {
	if s != nil {
		etc.SetDescription(*s)
	}
	return etc
}

// SetParentID sets the "parent" edge to the EnumerationType entity by ID.
func (etc *EnumerationTypeCreate) SetParentID(id int) *EnumerationTypeCreate {
	etc.mutation.SetParentID(id)
	return etc
}

// SetNillableParentID sets the "parent" edge to the EnumerationType entity by ID if the given value is not nil.
func (etc *EnumerationTypeCreate) SetNillableParentID(id *int) *EnumerationTypeCreate {
	if id != nil {
		etc = etc.SetParentID(*id)
	}
	return etc
}

// SetParent sets the "parent" edge to the EnumerationType entity.
func (etc *EnumerationTypeCreate) SetParent(e *EnumerationType) *EnumerationTypeCreate {
	return etc.SetParentID(e.ID)
}

// AddChildIDs adds the "children" edge to the EnumerationType entity by IDs.
func (etc *EnumerationTypeCreate) AddChildIDs(ids ...int) *EnumerationTypeCreate {
	etc.mutation.AddChildIDs(ids...)
	return etc
}

// AddChildren adds the "children" edges to the EnumerationType entity.
func (etc *EnumerationTypeCreate) AddChildren(e ...*EnumerationType) *EnumerationTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddChildIDs(ids...)
}

// AddEnumerationIDs adds the "enumerations" edge to the Enumeration entity by IDs.
func (etc *EnumerationTypeCreate) AddEnumerationIDs(ids ...int) *EnumerationTypeCreate {
	etc.mutation.AddEnumerationIDs(ids...)
	return etc
}

// AddEnumerations adds the "enumerations" edges to the Enumeration entity.
func (etc *EnumerationTypeCreate) AddEnumerations(e ...*Enumeration) *EnumerationTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddEnumerationIDs(ids...)
}

// AddChildEnumerationTypeIDs adds the "child_enumeration_types" edge to the EnumerationType entity by IDs.
func (etc *EnumerationTypeCreate) AddChildEnumerationTypeIDs(ids ...int) *EnumerationTypeCreate {
	etc.mutation.AddChildEnumerationTypeIDs(ids...)
	return etc
}

// AddChildEnumerationTypes adds the "child_enumeration_types" edges to the EnumerationType entity.
func (etc *EnumerationTypeCreate) AddChildEnumerationTypes(e ...*EnumerationType) *EnumerationTypeCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddChildEnumerationTypeIDs(ids...)
}

// Mutation returns the EnumerationTypeMutation object of the builder.
func (etc *EnumerationTypeCreate) Mutation() *EnumerationTypeMutation {
	return etc.mutation
}

// Save creates the EnumerationType in the database.
func (etc *EnumerationTypeCreate) Save(ctx context.Context) (*EnumerationType, error) {
	var (
		err  error
		node *EnumerationType
	)
	etc.defaults()
	if len(etc.hooks) == 0 {
		if err = etc.check(); err != nil {
			return nil, err
		}
		node, err = etc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EnumerationTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etc.check(); err != nil {
				return nil, err
			}
			etc.mutation = mutation
			if node, err = etc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(etc.hooks) - 1; i >= 0; i-- {
			mut = etc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EnumerationTypeCreate) SaveX(ctx context.Context) *EnumerationType {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (etc *EnumerationTypeCreate) defaults() {
	if _, ok := etc.mutation.CreateTime(); !ok {
		v := enumerationtype.DefaultCreateTime()
		etc.mutation.SetCreateTime(v)
	}
	if _, ok := etc.mutation.UpdateTime(); !ok {
		v := enumerationtype.DefaultUpdateTime()
		etc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EnumerationTypeCreate) check() error {
	if _, ok := etc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := etc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := etc.mutation.HasTable(); ok {
		if err := enumerationtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (etc *EnumerationTypeCreate) sqlSave(ctx context.Context) (*EnumerationType, error) {
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (etc *EnumerationTypeCreate) createSpec() (*EnumerationType, *sqlgraph.CreateSpec) {
	var (
		_node = &EnumerationType{config: etc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: enumerationtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enumerationtype.FieldID,
			},
		}
	)
	if value, ok := etc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enumerationtype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := etc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enumerationtype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := etc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enumerationtype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := etc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: enumerationtype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := etc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enumerationtype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := etc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enumerationtype.ParentTable,
			Columns: []string{enumerationtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumerationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumerationtype.ChildrenTable,
			Columns: []string{enumerationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumerationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.EnumerationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumerationtype.EnumerationsTable,
			Columns: []string{enumerationtype.EnumerationsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := etc.mutation.ChildEnumerationTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   enumerationtype.ChildEnumerationTypesTable,
			Columns: enumerationtype.ChildEnumerationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumerationtype.FieldID,
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

// EnumerationTypeCreateBulk is the builder for creating many EnumerationType entities in bulk.
type EnumerationTypeCreateBulk struct {
	config
	builders []*EnumerationTypeCreate
}

// Save creates the EnumerationType entities in the database.
func (etcb *EnumerationTypeCreateBulk) Save(ctx context.Context) ([]*EnumerationType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EnumerationType, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnumerationTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EnumerationTypeCreateBulk) SaveX(ctx context.Context) []*EnumerationType {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}