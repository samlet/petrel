// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
)

// SecurityGroupDelete is the builder for deleting a SecurityGroup entity.
type SecurityGroupDelete struct {
	config
	hooks    []Hook
	mutation *SecurityGroupMutation
}

// Where adds a new predicate to the SecurityGroupDelete builder.
func (sgd *SecurityGroupDelete) Where(ps ...predicate.SecurityGroup) *SecurityGroupDelete {
	sgd.mutation.predicates = append(sgd.mutation.predicates, ps...)
	return sgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sgd *SecurityGroupDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sgd.hooks) == 0 {
		affected, err = sgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecurityGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgd.mutation = mutation
			affected, err = sgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgd.hooks) - 1; i >= 0; i-- {
			mut = sgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgd *SecurityGroupDelete) ExecX(ctx context.Context) int {
	n, err := sgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sgd *SecurityGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: securitygroup.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securitygroup.FieldID,
			},
		},
	}
	if ps := sgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sgd.driver, _spec)
}

// SecurityGroupDeleteOne is the builder for deleting a single SecurityGroup entity.
type SecurityGroupDeleteOne struct {
	sgd *SecurityGroupDelete
}

// Exec executes the deletion query.
func (sgdo *SecurityGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := sgdo.sgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{securitygroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sgdo *SecurityGroupDeleteOne) ExecX(ctx context.Context) {
	sgdo.sgd.ExecX(ctx)
}
