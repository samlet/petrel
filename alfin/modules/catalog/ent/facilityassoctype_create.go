// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilityassoctype"
)

// FacilityAssocTypeCreate is the builder for creating a FacilityAssocType entity.
type FacilityAssocTypeCreate struct {
	config
	mutation *FacilityAssocTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (fatc *FacilityAssocTypeCreate) SetCreateTime(t time.Time) *FacilityAssocTypeCreate {
	fatc.mutation.SetCreateTime(t)
	return fatc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fatc *FacilityAssocTypeCreate) SetNillableCreateTime(t *time.Time) *FacilityAssocTypeCreate {
	if t != nil {
		fatc.SetCreateTime(*t)
	}
	return fatc
}

// SetUpdateTime sets the "update_time" field.
func (fatc *FacilityAssocTypeCreate) SetUpdateTime(t time.Time) *FacilityAssocTypeCreate {
	fatc.mutation.SetUpdateTime(t)
	return fatc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (fatc *FacilityAssocTypeCreate) SetNillableUpdateTime(t *time.Time) *FacilityAssocTypeCreate {
	if t != nil {
		fatc.SetUpdateTime(*t)
	}
	return fatc
}

// SetStringRef sets the "string_ref" field.
func (fatc *FacilityAssocTypeCreate) SetStringRef(s string) *FacilityAssocTypeCreate {
	fatc.mutation.SetStringRef(s)
	return fatc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (fatc *FacilityAssocTypeCreate) SetNillableStringRef(s *string) *FacilityAssocTypeCreate {
	if s != nil {
		fatc.SetStringRef(*s)
	}
	return fatc
}

// SetDescription sets the "description" field.
func (fatc *FacilityAssocTypeCreate) SetDescription(s string) *FacilityAssocTypeCreate {
	fatc.mutation.SetDescription(s)
	return fatc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fatc *FacilityAssocTypeCreate) SetNillableDescription(s *string) *FacilityAssocTypeCreate {
	if s != nil {
		fatc.SetDescription(*s)
	}
	return fatc
}

// Mutation returns the FacilityAssocTypeMutation object of the builder.
func (fatc *FacilityAssocTypeCreate) Mutation() *FacilityAssocTypeMutation {
	return fatc.mutation
}

// Save creates the FacilityAssocType in the database.
func (fatc *FacilityAssocTypeCreate) Save(ctx context.Context) (*FacilityAssocType, error) {
	var (
		err  error
		node *FacilityAssocType
	)
	fatc.defaults()
	if len(fatc.hooks) == 0 {
		if err = fatc.check(); err != nil {
			return nil, err
		}
		node, err = fatc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacilityAssocTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fatc.check(); err != nil {
				return nil, err
			}
			fatc.mutation = mutation
			if node, err = fatc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fatc.hooks) - 1; i >= 0; i-- {
			mut = fatc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fatc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fatc *FacilityAssocTypeCreate) SaveX(ctx context.Context) *FacilityAssocType {
	v, err := fatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fatc *FacilityAssocTypeCreate) defaults() {
	if _, ok := fatc.mutation.CreateTime(); !ok {
		v := facilityassoctype.DefaultCreateTime()
		fatc.mutation.SetCreateTime(v)
	}
	if _, ok := fatc.mutation.UpdateTime(); !ok {
		v := facilityassoctype.DefaultUpdateTime()
		fatc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fatc *FacilityAssocTypeCreate) check() error {
	if _, ok := fatc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := fatc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (fatc *FacilityAssocTypeCreate) sqlSave(ctx context.Context) (*FacilityAssocType, error) {
	_node, _spec := fatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fatc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fatc *FacilityAssocTypeCreate) createSpec() (*FacilityAssocType, *sqlgraph.CreateSpec) {
	var (
		_node = &FacilityAssocType{config: fatc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: facilityassoctype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilityassoctype.FieldID,
			},
		}
	)
	if value, ok := fatc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilityassoctype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := fatc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilityassoctype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := fatc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilityassoctype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := fatc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilityassoctype.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// FacilityAssocTypeCreateBulk is the builder for creating many FacilityAssocType entities in bulk.
type FacilityAssocTypeCreateBulk struct {
	config
	builders []*FacilityAssocTypeCreate
}

// Save creates the FacilityAssocType entities in the database.
func (fatcb *FacilityAssocTypeCreateBulk) Save(ctx context.Context) ([]*FacilityAssocType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fatcb.builders))
	nodes := make([]*FacilityAssocType, len(fatcb.builders))
	mutators := make([]Mutator, len(fatcb.builders))
	for i := range fatcb.builders {
		func(i int, root context.Context) {
			builder := fatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FacilityAssocTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, fatcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fatcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fatcb *FacilityAssocTypeCreateBulk) SaveX(ctx context.Context) []*FacilityAssocType {
	v, err := fatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
