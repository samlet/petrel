// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcontenttype"
)

// ProductContentTypeDelete is the builder for deleting a ProductContentType entity.
type ProductContentTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProductContentTypeMutation
}

// Where adds a new predicate to the ProductContentTypeDelete builder.
func (pctd *ProductContentTypeDelete) Where(ps ...predicate.ProductContentType) *ProductContentTypeDelete {
	pctd.mutation.predicates = append(pctd.mutation.predicates, ps...)
	return pctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pctd *ProductContentTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pctd.hooks) == 0 {
		affected, err = pctd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductContentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pctd.mutation = mutation
			affected, err = pctd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pctd.hooks) - 1; i >= 0; i-- {
			mut = pctd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pctd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pctd *ProductContentTypeDelete) ExecX(ctx context.Context) int {
	n, err := pctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pctd *ProductContentTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productcontenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcontenttype.FieldID,
			},
		},
	}
	if ps := pctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pctd.driver, _spec)
}

// ProductContentTypeDeleteOne is the builder for deleting a single ProductContentType entity.
type ProductContentTypeDeleteOne struct {
	pctd *ProductContentTypeDelete
}

// Exec executes the deletion query.
func (pctdo *ProductContentTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := pctdo.pctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productcontenttype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pctdo *ProductContentTypeDeleteOne) ExecX(ctx context.Context) {
	pctdo.pctd.ExecX(ctx)
}
