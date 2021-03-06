// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpriceactiontype"
)

// ProductPriceActionTypeDelete is the builder for deleting a ProductPriceActionType entity.
type ProductPriceActionTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProductPriceActionTypeMutation
}

// Where adds a new predicate to the ProductPriceActionTypeDelete builder.
func (ppatd *ProductPriceActionTypeDelete) Where(ps ...predicate.ProductPriceActionType) *ProductPriceActionTypeDelete {
	ppatd.mutation.predicates = append(ppatd.mutation.predicates, ps...)
	return ppatd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ppatd *ProductPriceActionTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ppatd.hooks) == 0 {
		affected, err = ppatd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductPriceActionTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ppatd.mutation = mutation
			affected, err = ppatd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ppatd.hooks) - 1; i >= 0; i-- {
			mut = ppatd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ppatd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppatd *ProductPriceActionTypeDelete) ExecX(ctx context.Context) int {
	n, err := ppatd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ppatd *ProductPriceActionTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productpriceactiontype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productpriceactiontype.FieldID,
			},
		},
	}
	if ps := ppatd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ppatd.driver, _spec)
}

// ProductPriceActionTypeDeleteOne is the builder for deleting a single ProductPriceActionType entity.
type ProductPriceActionTypeDeleteOne struct {
	ppatd *ProductPriceActionTypeDelete
}

// Exec executes the deletion query.
func (ppatdo *ProductPriceActionTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ppatdo.ppatd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productpriceactiontype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ppatdo *ProductPriceActionTypeDeleteOne) ExecX(ctx context.Context) {
	ppatdo.ppatd.ExecX(ctx)
}
