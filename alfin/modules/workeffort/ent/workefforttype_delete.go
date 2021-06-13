// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
)

// WorkEffortTypeDelete is the builder for deleting a WorkEffortType entity.
type WorkEffortTypeDelete struct {
	config
	hooks    []Hook
	mutation *WorkEffortTypeMutation
}

// Where adds a new predicate to the WorkEffortTypeDelete builder.
func (wetd *WorkEffortTypeDelete) Where(ps ...predicate.WorkEffortType) *WorkEffortTypeDelete {
	wetd.mutation.predicates = append(wetd.mutation.predicates, ps...)
	return wetd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wetd *WorkEffortTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wetd.hooks) == 0 {
		affected, err = wetd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkEffortTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wetd.mutation = mutation
			affected, err = wetd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wetd.hooks) - 1; i >= 0; i-- {
			mut = wetd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wetd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wetd *WorkEffortTypeDelete) ExecX(ctx context.Context) int {
	n, err := wetd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wetd *WorkEffortTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workefforttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workefforttype.FieldID,
			},
		},
	}
	if ps := wetd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wetd.driver, _spec)
}

// WorkEffortTypeDeleteOne is the builder for deleting a single WorkEffortType entity.
type WorkEffortTypeDeleteOne struct {
	wetd *WorkEffortTypeDelete
}

// Exec executes the deletion query.
func (wetdo *WorkEffortTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := wetdo.wetd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workefforttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wetdo *WorkEffortTypeDeleteOne) ExecX(ctx context.Context) {
	wetdo.wetd.ExecX(ctx)
}
