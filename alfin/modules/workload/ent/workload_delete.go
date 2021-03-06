// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
)

// WorkloadDelete is the builder for deleting a Workload entity.
type WorkloadDelete struct {
	config
	hooks    []Hook
	mutation *WorkloadMutation
}

// Where adds a new predicate to the WorkloadDelete builder.
func (wd *WorkloadDelete) Where(ps ...predicate.Workload) *WorkloadDelete {
	wd.mutation.predicates = append(wd.mutation.predicates, ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkloadDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wd.hooks) == 0 {
		affected, err = wd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wd.mutation = mutation
			affected, err = wd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wd.hooks) - 1; i >= 0; i-- {
			mut = wd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkloadDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkloadDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workload.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workload.FieldID,
			},
		},
	}
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
}

// WorkloadDeleteOne is the builder for deleting a single Workload entity.
type WorkloadDeleteOne struct {
	wd *WorkloadDelete
}

// Exec executes the deletion query.
func (wdo *WorkloadDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workload.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkloadDeleteOne) ExecX(ctx context.Context) {
	wdo.wd.ExecX(ctx)
}
