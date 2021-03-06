// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventprptyp"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// CommunicationEventPrpTypDelete is the builder for deleting a CommunicationEventPrpTyp entity.
type CommunicationEventPrpTypDelete struct {
	config
	hooks    []Hook
	mutation *CommunicationEventPrpTypMutation
}

// Where adds a new predicate to the CommunicationEventPrpTypDelete builder.
func (ceptd *CommunicationEventPrpTypDelete) Where(ps ...predicate.CommunicationEventPrpTyp) *CommunicationEventPrpTypDelete {
	ceptd.mutation.predicates = append(ceptd.mutation.predicates, ps...)
	return ceptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ceptd *CommunicationEventPrpTypDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ceptd.hooks) == 0 {
		affected, err = ceptd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunicationEventPrpTypMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceptd.mutation = mutation
			affected, err = ceptd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceptd.hooks) - 1; i >= 0; i-- {
			mut = ceptd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceptd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceptd *CommunicationEventPrpTypDelete) ExecX(ctx context.Context) int {
	n, err := ceptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ceptd *CommunicationEventPrpTypDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: communicationeventprptyp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationeventprptyp.FieldID,
			},
		},
	}
	if ps := ceptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ceptd.driver, _spec)
}

// CommunicationEventPrpTypDeleteOne is the builder for deleting a single CommunicationEventPrpTyp entity.
type CommunicationEventPrpTypDeleteOne struct {
	ceptd *CommunicationEventPrpTypDelete
}

// Exec executes the deletion query.
func (ceptdo *CommunicationEventPrpTypDeleteOne) Exec(ctx context.Context) error {
	n, err := ceptdo.ceptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{communicationeventprptyp.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ceptdo *CommunicationEventPrpTypDeleteOne) ExecX(ctx context.Context) {
	ceptdo.ceptd.ExecX(ctx)
}
