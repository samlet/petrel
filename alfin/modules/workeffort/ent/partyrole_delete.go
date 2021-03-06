// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// PartyRoleDelete is the builder for deleting a PartyRole entity.
type PartyRoleDelete struct {
	config
	hooks    []Hook
	mutation *PartyRoleMutation
}

// Where adds a new predicate to the PartyRoleDelete builder.
func (prd *PartyRoleDelete) Where(ps ...predicate.PartyRole) *PartyRoleDelete {
	prd.mutation.predicates = append(prd.mutation.predicates, ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *PartyRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(prd.hooks) == 0 {
		affected, err = prd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prd.mutation = mutation
			affected, err = prd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(prd.hooks) - 1; i >= 0; i-- {
			mut = prd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *PartyRoleDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *PartyRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: partyrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyrole.FieldID,
			},
		},
	}
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
}

// PartyRoleDeleteOne is the builder for deleting a single PartyRole entity.
type PartyRoleDeleteOne struct {
	prd *PartyRoleDelete
}

// Exec executes the deletion query.
func (prdo *PartyRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{partyrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *PartyRoleDeleteOne) ExecX(ctx context.Context) {
	prdo.prd.ExecX(ctx)
}
