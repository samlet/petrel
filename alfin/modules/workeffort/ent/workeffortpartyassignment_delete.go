// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortpartyassignment"
)

// WorkEffortPartyAssignmentDelete is the builder for deleting a WorkEffortPartyAssignment entity.
type WorkEffortPartyAssignmentDelete struct {
	config
	hooks    []Hook
	mutation *WorkEffortPartyAssignmentMutation
}

// Where adds a new predicate to the WorkEffortPartyAssignmentDelete builder.
func (wepad *WorkEffortPartyAssignmentDelete) Where(ps ...predicate.WorkEffortPartyAssignment) *WorkEffortPartyAssignmentDelete {
	wepad.mutation.predicates = append(wepad.mutation.predicates, ps...)
	return wepad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wepad *WorkEffortPartyAssignmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wepad.hooks) == 0 {
		affected, err = wepad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkEffortPartyAssignmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wepad.mutation = mutation
			affected, err = wepad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wepad.hooks) - 1; i >= 0; i-- {
			mut = wepad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wepad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wepad *WorkEffortPartyAssignmentDelete) ExecX(ctx context.Context) int {
	n, err := wepad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wepad *WorkEffortPartyAssignmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: workeffortpartyassignment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workeffortpartyassignment.FieldID,
			},
		},
	}
	if ps := wepad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wepad.driver, _spec)
}

// WorkEffortPartyAssignmentDeleteOne is the builder for deleting a single WorkEffortPartyAssignment entity.
type WorkEffortPartyAssignmentDeleteOne struct {
	wepad *WorkEffortPartyAssignmentDelete
}

// Exec executes the deletion query.
func (wepado *WorkEffortPartyAssignmentDeleteOne) Exec(ctx context.Context) error {
	n, err := wepado.wepad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workeffortpartyassignment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wepado *WorkEffortPartyAssignmentDeleteOne) ExecX(ctx context.Context) {
	wepado.wepad.ExecX(ctx)
}
