// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// OrderHeaderDelete is the builder for deleting a OrderHeader entity.
type OrderHeaderDelete struct {
	config
	hooks    []Hook
	mutation *OrderHeaderMutation
}

// Where adds a new predicate to the OrderHeaderDelete builder.
func (ohd *OrderHeaderDelete) Where(ps ...predicate.OrderHeader) *OrderHeaderDelete {
	ohd.mutation.predicates = append(ohd.mutation.predicates, ps...)
	return ohd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ohd *OrderHeaderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ohd.hooks) == 0 {
		affected, err = ohd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderHeaderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ohd.mutation = mutation
			affected, err = ohd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ohd.hooks) - 1; i >= 0; i-- {
			mut = ohd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ohd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohd *OrderHeaderDelete) ExecX(ctx context.Context) int {
	n, err := ohd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ohd *OrderHeaderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderheader.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderheader.FieldID,
			},
		},
	}
	if ps := ohd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ohd.driver, _spec)
}

// OrderHeaderDeleteOne is the builder for deleting a single OrderHeader entity.
type OrderHeaderDeleteOne struct {
	ohd *OrderHeaderDelete
}

// Exec executes the deletion query.
func (ohdo *OrderHeaderDeleteOne) Exec(ctx context.Context) error {
	n, err := ohdo.ohd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderheader.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ohdo *OrderHeaderDeleteOne) ExecX(ctx context.Context) {
	ohdo.ohd.ExecX(ctx)
}
