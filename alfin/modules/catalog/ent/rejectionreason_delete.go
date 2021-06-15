// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/rejectionreason"
)

// RejectionReasonDelete is the builder for deleting a RejectionReason entity.
type RejectionReasonDelete struct {
	config
	hooks    []Hook
	mutation *RejectionReasonMutation
}

// Where adds a new predicate to the RejectionReasonDelete builder.
func (rrd *RejectionReasonDelete) Where(ps ...predicate.RejectionReason) *RejectionReasonDelete {
	rrd.mutation.predicates = append(rrd.mutation.predicates, ps...)
	return rrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rrd *RejectionReasonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rrd.hooks) == 0 {
		affected, err = rrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RejectionReasonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rrd.mutation = mutation
			affected, err = rrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rrd.hooks) - 1; i >= 0; i-- {
			mut = rrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrd *RejectionReasonDelete) ExecX(ctx context.Context) int {
	n, err := rrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rrd *RejectionReasonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: rejectionreason.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rejectionreason.FieldID,
			},
		},
	}
	if ps := rrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rrd.driver, _spec)
}

// RejectionReasonDeleteOne is the builder for deleting a single RejectionReason entity.
type RejectionReasonDeleteOne struct {
	rrd *RejectionReasonDelete
}

// Exec executes the deletion query.
func (rrdo *RejectionReasonDeleteOne) Exec(ctx context.Context) error {
	n, err := rrdo.rrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rejectionreason.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rrdo *RejectionReasonDeleteOne) ExecX(ctx context.Context) {
	rrdo.rrd.ExecX(ctx)
}
