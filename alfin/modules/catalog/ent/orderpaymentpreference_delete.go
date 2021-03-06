// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// OrderPaymentPreferenceDelete is the builder for deleting a OrderPaymentPreference entity.
type OrderPaymentPreferenceDelete struct {
	config
	hooks    []Hook
	mutation *OrderPaymentPreferenceMutation
}

// Where adds a new predicate to the OrderPaymentPreferenceDelete builder.
func (oppd *OrderPaymentPreferenceDelete) Where(ps ...predicate.OrderPaymentPreference) *OrderPaymentPreferenceDelete {
	oppd.mutation.predicates = append(oppd.mutation.predicates, ps...)
	return oppd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oppd *OrderPaymentPreferenceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oppd.hooks) == 0 {
		affected, err = oppd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPaymentPreferenceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oppd.mutation = mutation
			affected, err = oppd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oppd.hooks) - 1; i >= 0; i-- {
			mut = oppd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oppd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oppd *OrderPaymentPreferenceDelete) ExecX(ctx context.Context) int {
	n, err := oppd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oppd *OrderPaymentPreferenceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderpaymentpreference.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderpaymentpreference.FieldID,
			},
		},
	}
	if ps := oppd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, oppd.driver, _spec)
}

// OrderPaymentPreferenceDeleteOne is the builder for deleting a single OrderPaymentPreference entity.
type OrderPaymentPreferenceDeleteOne struct {
	oppd *OrderPaymentPreferenceDelete
}

// Exec executes the deletion query.
func (oppdo *OrderPaymentPreferenceDeleteOne) Exec(ctx context.Context) error {
	n, err := oppdo.oppd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderpaymentpreference.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oppdo *OrderPaymentPreferenceDeleteOne) ExecX(ctx context.Context) {
	oppdo.oppd.ExecX(ctx)
}
