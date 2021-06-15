// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/rejectionreason"
)

// RejectionReasonCreate is the builder for creating a RejectionReason entity.
type RejectionReasonCreate struct {
	config
	mutation *RejectionReasonMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rrc *RejectionReasonCreate) SetCreateTime(t time.Time) *RejectionReasonCreate {
	rrc.mutation.SetCreateTime(t)
	return rrc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rrc *RejectionReasonCreate) SetNillableCreateTime(t *time.Time) *RejectionReasonCreate {
	if t != nil {
		rrc.SetCreateTime(*t)
	}
	return rrc
}

// SetUpdateTime sets the "update_time" field.
func (rrc *RejectionReasonCreate) SetUpdateTime(t time.Time) *RejectionReasonCreate {
	rrc.mutation.SetUpdateTime(t)
	return rrc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rrc *RejectionReasonCreate) SetNillableUpdateTime(t *time.Time) *RejectionReasonCreate {
	if t != nil {
		rrc.SetUpdateTime(*t)
	}
	return rrc
}

// SetStringRef sets the "string_ref" field.
func (rrc *RejectionReasonCreate) SetStringRef(s string) *RejectionReasonCreate {
	rrc.mutation.SetStringRef(s)
	return rrc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (rrc *RejectionReasonCreate) SetNillableStringRef(s *string) *RejectionReasonCreate {
	if s != nil {
		rrc.SetStringRef(*s)
	}
	return rrc
}

// SetDescription sets the "description" field.
func (rrc *RejectionReasonCreate) SetDescription(s string) *RejectionReasonCreate {
	rrc.mutation.SetDescription(s)
	return rrc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rrc *RejectionReasonCreate) SetNillableDescription(s *string) *RejectionReasonCreate {
	if s != nil {
		rrc.SetDescription(*s)
	}
	return rrc
}

// Mutation returns the RejectionReasonMutation object of the builder.
func (rrc *RejectionReasonCreate) Mutation() *RejectionReasonMutation {
	return rrc.mutation
}

// Save creates the RejectionReason in the database.
func (rrc *RejectionReasonCreate) Save(ctx context.Context) (*RejectionReason, error) {
	var (
		err  error
		node *RejectionReason
	)
	rrc.defaults()
	if len(rrc.hooks) == 0 {
		if err = rrc.check(); err != nil {
			return nil, err
		}
		node, err = rrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RejectionReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rrc.check(); err != nil {
				return nil, err
			}
			rrc.mutation = mutation
			if node, err = rrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rrc.hooks) - 1; i >= 0; i-- {
			mut = rrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rrc *RejectionReasonCreate) SaveX(ctx context.Context) *RejectionReason {
	v, err := rrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (rrc *RejectionReasonCreate) defaults() {
	if _, ok := rrc.mutation.CreateTime(); !ok {
		v := rejectionreason.DefaultCreateTime()
		rrc.mutation.SetCreateTime(v)
	}
	if _, ok := rrc.mutation.UpdateTime(); !ok {
		v := rejectionreason.DefaultUpdateTime()
		rrc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rrc *RejectionReasonCreate) check() error {
	if _, ok := rrc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := rrc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (rrc *RejectionReasonCreate) sqlSave(ctx context.Context) (*RejectionReason, error) {
	_node, _spec := rrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rrc *RejectionReasonCreate) createSpec() (*RejectionReason, *sqlgraph.CreateSpec) {
	var (
		_node = &RejectionReason{config: rrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rejectionreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rejectionreason.FieldID,
			},
		}
	)
	if value, ok := rrc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rejectionreason.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := rrc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rejectionreason.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := rrc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rejectionreason.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := rrc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rejectionreason.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// RejectionReasonCreateBulk is the builder for creating many RejectionReason entities in bulk.
type RejectionReasonCreateBulk struct {
	config
	builders []*RejectionReasonCreate
}

// Save creates the RejectionReason entities in the database.
func (rrcb *RejectionReasonCreateBulk) Save(ctx context.Context) ([]*RejectionReason, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rrcb.builders))
	nodes := make([]*RejectionReason, len(rrcb.builders))
	mutators := make([]Mutator, len(rrcb.builders))
	for i := range rrcb.builders {
		func(i int, root context.Context) {
			builder := rrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RejectionReasonMutation)
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
					_, err = mutators[i+1].Mutate(root, rrcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rrcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rrcb *RejectionReasonCreateBulk) SaveX(ctx context.Context) []*RejectionReason {
	v, err := rrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}