// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstore"
)

// ProductStoreDelete is the builder for deleting a ProductStore entity.
type ProductStoreDelete struct {
	config
	hooks    []Hook
	mutation *ProductStoreMutation
}

// Where adds a new predicate to the ProductStoreDelete builder.
func (psd *ProductStoreDelete) Where(ps ...predicate.ProductStore) *ProductStoreDelete {
	psd.mutation.predicates = append(psd.mutation.predicates, ps...)
	return psd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psd *ProductStoreDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(psd.hooks) == 0 {
		affected, err = psd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductStoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psd.mutation = mutation
			affected, err = psd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psd.hooks) - 1; i >= 0; i-- {
			mut = psd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (psd *ProductStoreDelete) ExecX(ctx context.Context) int {
	n, err := psd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psd *ProductStoreDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productstore.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productstore.FieldID,
			},
		},
	}
	if ps := psd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, psd.driver, _spec)
}

// ProductStoreDeleteOne is the builder for deleting a single ProductStore entity.
type ProductStoreDeleteOne struct {
	psd *ProductStoreDelete
}

// Exec executes the deletion query.
func (psdo *ProductStoreDeleteOne) Exec(ctx context.Context) error {
	n, err := psdo.psd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productstore.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psdo *ProductStoreDeleteOne) ExecX(ctx context.Context) {
	psdo.psd.ExecX(ctx)
}