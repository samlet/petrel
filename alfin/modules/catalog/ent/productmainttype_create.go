// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productmainttype"
)

// ProductMaintTypeCreate is the builder for creating a ProductMaintType entity.
type ProductMaintTypeCreate struct {
	config
	mutation *ProductMaintTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pmtc *ProductMaintTypeCreate) SetCreateTime(t time.Time) *ProductMaintTypeCreate {
	pmtc.mutation.SetCreateTime(t)
	return pmtc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pmtc *ProductMaintTypeCreate) SetNillableCreateTime(t *time.Time) *ProductMaintTypeCreate {
	if t != nil {
		pmtc.SetCreateTime(*t)
	}
	return pmtc
}

// SetUpdateTime sets the "update_time" field.
func (pmtc *ProductMaintTypeCreate) SetUpdateTime(t time.Time) *ProductMaintTypeCreate {
	pmtc.mutation.SetUpdateTime(t)
	return pmtc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pmtc *ProductMaintTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductMaintTypeCreate {
	if t != nil {
		pmtc.SetUpdateTime(*t)
	}
	return pmtc
}

// SetStringRef sets the "string_ref" field.
func (pmtc *ProductMaintTypeCreate) SetStringRef(s string) *ProductMaintTypeCreate {
	pmtc.mutation.SetStringRef(s)
	return pmtc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pmtc *ProductMaintTypeCreate) SetNillableStringRef(s *string) *ProductMaintTypeCreate {
	if s != nil {
		pmtc.SetStringRef(*s)
	}
	return pmtc
}

// SetDescription sets the "description" field.
func (pmtc *ProductMaintTypeCreate) SetDescription(s string) *ProductMaintTypeCreate {
	pmtc.mutation.SetDescription(s)
	return pmtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pmtc *ProductMaintTypeCreate) SetNillableDescription(s *string) *ProductMaintTypeCreate {
	if s != nil {
		pmtc.SetDescription(*s)
	}
	return pmtc
}

// SetParentID sets the "parent" edge to the ProductMaintType entity by ID.
func (pmtc *ProductMaintTypeCreate) SetParentID(id int) *ProductMaintTypeCreate {
	pmtc.mutation.SetParentID(id)
	return pmtc
}

// SetNillableParentID sets the "parent" edge to the ProductMaintType entity by ID if the given value is not nil.
func (pmtc *ProductMaintTypeCreate) SetNillableParentID(id *int) *ProductMaintTypeCreate {
	if id != nil {
		pmtc = pmtc.SetParentID(*id)
	}
	return pmtc
}

// SetParent sets the "parent" edge to the ProductMaintType entity.
func (pmtc *ProductMaintTypeCreate) SetParent(p *ProductMaintType) *ProductMaintTypeCreate {
	return pmtc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductMaintType entity by IDs.
func (pmtc *ProductMaintTypeCreate) AddChildIDs(ids ...int) *ProductMaintTypeCreate {
	pmtc.mutation.AddChildIDs(ids...)
	return pmtc
}

// AddChildren adds the "children" edges to the ProductMaintType entity.
func (pmtc *ProductMaintTypeCreate) AddChildren(p ...*ProductMaintType) *ProductMaintTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmtc.AddChildIDs(ids...)
}

// AddChildProductMaintTypeIDs adds the "child_product_maint_types" edge to the ProductMaintType entity by IDs.
func (pmtc *ProductMaintTypeCreate) AddChildProductMaintTypeIDs(ids ...int) *ProductMaintTypeCreate {
	pmtc.mutation.AddChildProductMaintTypeIDs(ids...)
	return pmtc
}

// AddChildProductMaintTypes adds the "child_product_maint_types" edges to the ProductMaintType entity.
func (pmtc *ProductMaintTypeCreate) AddChildProductMaintTypes(p ...*ProductMaintType) *ProductMaintTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmtc.AddChildProductMaintTypeIDs(ids...)
}

// Mutation returns the ProductMaintTypeMutation object of the builder.
func (pmtc *ProductMaintTypeCreate) Mutation() *ProductMaintTypeMutation {
	return pmtc.mutation
}

// Save creates the ProductMaintType in the database.
func (pmtc *ProductMaintTypeCreate) Save(ctx context.Context) (*ProductMaintType, error) {
	var (
		err  error
		node *ProductMaintType
	)
	pmtc.defaults()
	if len(pmtc.hooks) == 0 {
		if err = pmtc.check(); err != nil {
			return nil, err
		}
		node, err = pmtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMaintTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pmtc.check(); err != nil {
				return nil, err
			}
			pmtc.mutation = mutation
			if node, err = pmtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pmtc.hooks) - 1; i >= 0; i-- {
			mut = pmtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pmtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pmtc *ProductMaintTypeCreate) SaveX(ctx context.Context) *ProductMaintType {
	v, err := pmtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pmtc *ProductMaintTypeCreate) defaults() {
	if _, ok := pmtc.mutation.CreateTime(); !ok {
		v := productmainttype.DefaultCreateTime()
		pmtc.mutation.SetCreateTime(v)
	}
	if _, ok := pmtc.mutation.UpdateTime(); !ok {
		v := productmainttype.DefaultUpdateTime()
		pmtc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmtc *ProductMaintTypeCreate) check() error {
	if _, ok := pmtc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pmtc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (pmtc *ProductMaintTypeCreate) sqlSave(ctx context.Context) (*ProductMaintType, error) {
	_node, _spec := pmtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pmtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pmtc *ProductMaintTypeCreate) createSpec() (*ProductMaintType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductMaintType{config: pmtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productmainttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productmainttype.FieldID,
			},
		}
	)
	if value, ok := pmtc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productmainttype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pmtc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productmainttype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pmtc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productmainttype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pmtc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productmainttype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pmtc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productmainttype.ParentTable,
			Columns: []string{productmainttype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productmainttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_maint_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pmtc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productmainttype.ChildrenTable,
			Columns: []string{productmainttype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productmainttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pmtc.mutation.ChildProductMaintTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productmainttype.ChildProductMaintTypesTable,
			Columns: productmainttype.ChildProductMaintTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productmainttype.FieldID,
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

// ProductMaintTypeCreateBulk is the builder for creating many ProductMaintType entities in bulk.
type ProductMaintTypeCreateBulk struct {
	config
	builders []*ProductMaintTypeCreate
}

// Save creates the ProductMaintType entities in the database.
func (pmtcb *ProductMaintTypeCreateBulk) Save(ctx context.Context) ([]*ProductMaintType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pmtcb.builders))
	nodes := make([]*ProductMaintType, len(pmtcb.builders))
	mutators := make([]Mutator, len(pmtcb.builders))
	for i := range pmtcb.builders {
		func(i int, root context.Context) {
			builder := pmtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMaintTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pmtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pmtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pmtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pmtcb *ProductMaintTypeCreateBulk) SaveX(ctx context.Context) []*ProductMaintType {
	v, err := pmtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
