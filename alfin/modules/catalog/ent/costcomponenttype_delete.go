// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/costcomponenttype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// CostComponentTypeDelete is the builder for deleting a CostComponentType entity.
type CostComponentTypeDelete struct {
	config
	hooks    []Hook
	mutation *CostComponentTypeMutation
}

// Where adds a new predicate to the CostComponentTypeDelete builder.
func (cctd *CostComponentTypeDelete) Where(ps ...predicate.CostComponentType) *CostComponentTypeDelete {
	cctd.mutation.predicates = append(cctd.mutation.predicates, ps...)
	return cctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cctd *CostComponentTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cctd.hooks) == 0 {
		affected, err = cctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CostComponentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cctd.mutation = mutation
			affected, err = cctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cctd.hooks) - 1; i >= 0; i-- {
			mut = cctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cctd *CostComponentTypeDelete) ExecX(ctx context.Context) int {
	n, err := cctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cctd *CostComponentTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: costcomponenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: costcomponenttype.FieldID,
			},
		},
	}
	if ps := cctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cctd.driver, _spec)
}

// CostComponentTypeDeleteOne is the builder for deleting a single CostComponentType entity.
type CostComponentTypeDeleteOne struct {
	cctd *CostComponentTypeDelete
}

// Exec executes the deletion query.
func (cctdo *CostComponentTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := cctdo.cctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{costcomponenttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cctdo *CostComponentTypeDeleteOne) ExecX(ctx context.Context) {
	cctdo.cctd.ExecX(ctx)
}
