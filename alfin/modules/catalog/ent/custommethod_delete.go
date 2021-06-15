// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethod"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// CustomMethodDelete is the builder for deleting a CustomMethod entity.
type CustomMethodDelete struct {
	config
	hooks    []Hook
	mutation *CustomMethodMutation
}

// Where adds a new predicate to the CustomMethodDelete builder.
func (cmd *CustomMethodDelete) Where(ps ...predicate.CustomMethod) *CustomMethodDelete {
	cmd.mutation.predicates = append(cmd.mutation.predicates, ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *CustomMethodDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmd.hooks) == 0 {
		affected, err = cmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomMethodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmd.mutation = mutation
			affected, err = cmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmd.hooks) - 1; i >= 0; i-- {
			mut = cmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *CustomMethodDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *CustomMethodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: custommethod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: custommethod.FieldID,
			},
		},
	}
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
}

// CustomMethodDeleteOne is the builder for deleting a single CustomMethod entity.
type CustomMethodDeleteOne struct {
	cmd *CustomMethodDelete
}

// Exec executes the deletion query.
func (cmdo *CustomMethodDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{custommethod.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *CustomMethodDeleteOne) ExecX(ctx context.Context) {
	cmdo.cmd.ExecX(ctx)
}