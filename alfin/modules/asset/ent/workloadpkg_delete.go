// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/asset/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"
)

// WorkloadPkgDelete is the builder for deleting a WorkloadPkg entity.
type WorkloadPkgDelete struct {
	config
	hooks    []Hook
	mutation *WorkloadPkgMutation
}

// Where adds a new predicate to the WorkloadPkgDelete builder.
func (wpd *WorkloadPkgDelete) Where(ps ...predicate.WorkloadPkg) *WorkloadPkgDelete {
	wpd.mutation.predicates = append(wpd.mutation.predicates, ps...)
	return wpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wpd *WorkloadPkgDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wpd.hooks) == 0 {
		affected, err = wpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadPkgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wpd.mutation = mutation
			affected, err = wpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wpd.hooks) - 1; i >= 0; i-- {
			mut = wpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpd *WorkloadPkgDelete) ExecX(ctx context.Context) int {
	n, err := wpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wpd *WorkloadPkgDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workloadpkg.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadpkg.FieldID,
			},
		},
	}
	if ps := wpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wpd.driver, _spec)
}

// WorkloadPkgDeleteOne is the builder for deleting a single WorkloadPkg entity.
type WorkloadPkgDeleteOne struct {
	wpd *WorkloadPkgDelete
}

// Exec executes the deletion query.
func (wpdo *WorkloadPkgDeleteOne) Exec(ctx context.Context) error {
	n, err := wpdo.wpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workloadpkg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wpdo *WorkloadPkgDeleteOne) ExecX(ctx context.Context) {
	wpdo.wpd.ExecX(ctx)
}