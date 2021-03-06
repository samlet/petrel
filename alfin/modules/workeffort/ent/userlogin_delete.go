// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
)

// UserLoginDelete is the builder for deleting a UserLogin entity.
type UserLoginDelete struct {
	config
	hooks    []Hook
	mutation *UserLoginMutation
}

// Where adds a new predicate to the UserLoginDelete builder.
func (uld *UserLoginDelete) Where(ps ...predicate.UserLogin) *UserLoginDelete {
	uld.mutation.predicates = append(uld.mutation.predicates, ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UserLoginDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uld.hooks) == 0 {
		affected, err = uld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uld.mutation = mutation
			affected, err = uld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uld.hooks) - 1; i >= 0; i-- {
			mut = uld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UserLoginDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UserLoginDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userlogin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlogin.FieldID,
			},
		},
	}
	if ps := uld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
}

// UserLoginDeleteOne is the builder for deleting a single UserLogin entity.
type UserLoginDeleteOne struct {
	uld *UserLoginDelete
}

// Exec executes the deletion query.
func (uldo *UserLoginDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userlogin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UserLoginDeleteOne) ExecX(ctx context.Context) {
	uldo.uld.ExecX(ctx)
}
