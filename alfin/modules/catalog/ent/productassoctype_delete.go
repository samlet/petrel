// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoctype"
)

// ProductAssocTypeDelete is the builder for deleting a ProductAssocType entity.
type ProductAssocTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProductAssocTypeMutation
}

// Where adds a new predicate to the ProductAssocTypeDelete builder.
func (patd *ProductAssocTypeDelete) Where(ps ...predicate.ProductAssocType) *ProductAssocTypeDelete {
	patd.mutation.predicates = append(patd.mutation.predicates, ps...)
	return patd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (patd *ProductAssocTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(patd.hooks) == 0 {
		affected, err = patd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductAssocTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			patd.mutation = mutation
			affected, err = patd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(patd.hooks) - 1; i >= 0; i-- {
			mut = patd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, patd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (patd *ProductAssocTypeDelete) ExecX(ctx context.Context) int {
	n, err := patd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (patd *ProductAssocTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productassoctype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productassoctype.FieldID,
			},
		},
	}
	if ps := patd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, patd.driver, _spec)
}

// ProductAssocTypeDeleteOne is the builder for deleting a single ProductAssocType entity.
type ProductAssocTypeDeleteOne struct {
	patd *ProductAssocTypeDelete
}

// Exec executes the deletion query.
func (patdo *ProductAssocTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := patdo.patd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productassoctype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (patdo *ProductAssocTypeDeleteOne) ExecX(ctx context.Context) {
	patdo.patd.ExecX(ctx)
}