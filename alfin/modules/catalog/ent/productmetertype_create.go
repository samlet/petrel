// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productmetertype"
)

// ProductMeterTypeCreate is the builder for creating a ProductMeterType entity.
type ProductMeterTypeCreate struct {
	config
	mutation *ProductMeterTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pmtc *ProductMeterTypeCreate) SetCreateTime(t time.Time) *ProductMeterTypeCreate {
	pmtc.mutation.SetCreateTime(t)
	return pmtc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pmtc *ProductMeterTypeCreate) SetNillableCreateTime(t *time.Time) *ProductMeterTypeCreate {
	if t != nil {
		pmtc.SetCreateTime(*t)
	}
	return pmtc
}

// SetUpdateTime sets the "update_time" field.
func (pmtc *ProductMeterTypeCreate) SetUpdateTime(t time.Time) *ProductMeterTypeCreate {
	pmtc.mutation.SetUpdateTime(t)
	return pmtc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pmtc *ProductMeterTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductMeterTypeCreate {
	if t != nil {
		pmtc.SetUpdateTime(*t)
	}
	return pmtc
}

// SetStringRef sets the "string_ref" field.
func (pmtc *ProductMeterTypeCreate) SetStringRef(s string) *ProductMeterTypeCreate {
	pmtc.mutation.SetStringRef(s)
	return pmtc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pmtc *ProductMeterTypeCreate) SetNillableStringRef(s *string) *ProductMeterTypeCreate {
	if s != nil {
		pmtc.SetStringRef(*s)
	}
	return pmtc
}

// SetDescription sets the "description" field.
func (pmtc *ProductMeterTypeCreate) SetDescription(s string) *ProductMeterTypeCreate {
	pmtc.mutation.SetDescription(s)
	return pmtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pmtc *ProductMeterTypeCreate) SetNillableDescription(s *string) *ProductMeterTypeCreate {
	if s != nil {
		pmtc.SetDescription(*s)
	}
	return pmtc
}

// SetDefaultUomID sets the "default_uom_id" field.
func (pmtc *ProductMeterTypeCreate) SetDefaultUomID(i int) *ProductMeterTypeCreate {
	pmtc.mutation.SetDefaultUomID(i)
	return pmtc
}

// SetNillableDefaultUomID sets the "default_uom_id" field if the given value is not nil.
func (pmtc *ProductMeterTypeCreate) SetNillableDefaultUomID(i *int) *ProductMeterTypeCreate {
	if i != nil {
		pmtc.SetDefaultUomID(*i)
	}
	return pmtc
}

// Mutation returns the ProductMeterTypeMutation object of the builder.
func (pmtc *ProductMeterTypeCreate) Mutation() *ProductMeterTypeMutation {
	return pmtc.mutation
}

// Save creates the ProductMeterType in the database.
func (pmtc *ProductMeterTypeCreate) Save(ctx context.Context) (*ProductMeterType, error) {
	var (
		err  error
		node *ProductMeterType
	)
	pmtc.defaults()
	if len(pmtc.hooks) == 0 {
		if err = pmtc.check(); err != nil {
			return nil, err
		}
		node, err = pmtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMeterTypeMutation)
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
func (pmtc *ProductMeterTypeCreate) SaveX(ctx context.Context) *ProductMeterType {
	v, err := pmtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pmtc *ProductMeterTypeCreate) defaults() {
	if _, ok := pmtc.mutation.CreateTime(); !ok {
		v := productmetertype.DefaultCreateTime()
		pmtc.mutation.SetCreateTime(v)
	}
	if _, ok := pmtc.mutation.UpdateTime(); !ok {
		v := productmetertype.DefaultUpdateTime()
		pmtc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pmtc *ProductMeterTypeCreate) check() error {
	if _, ok := pmtc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pmtc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (pmtc *ProductMeterTypeCreate) sqlSave(ctx context.Context) (*ProductMeterType, error) {
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

func (pmtc *ProductMeterTypeCreate) createSpec() (*ProductMeterType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductMeterType{config: pmtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productmetertype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productmetertype.FieldID,
			},
		}
	)
	if value, ok := pmtc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productmetertype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pmtc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productmetertype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pmtc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productmetertype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pmtc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productmetertype.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := pmtc.mutation.DefaultUomID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productmetertype.FieldDefaultUomID,
		})
		_node.DefaultUomID = value
	}
	return _node, _spec
}

// ProductMeterTypeCreateBulk is the builder for creating many ProductMeterType entities in bulk.
type ProductMeterTypeCreateBulk struct {
	config
	builders []*ProductMeterTypeCreate
}

// Save creates the ProductMeterType entities in the database.
func (pmtcb *ProductMeterTypeCreateBulk) Save(ctx context.Context) ([]*ProductMeterType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pmtcb.builders))
	nodes := make([]*ProductMeterType, len(pmtcb.builders))
	mutators := make([]Mutator, len(pmtcb.builders))
	for i := range pmtcb.builders {
		func(i int, root context.Context) {
			builder := pmtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMeterTypeMutation)
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
func (pmtcb *ProductMeterTypeCreateBulk) SaveX(ctx context.Context) []*ProductMeterType {
	v, err := pmtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}