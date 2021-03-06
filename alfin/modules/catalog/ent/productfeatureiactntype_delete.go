// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeatureiactntype"
)

// ProductFeatureIactnTypeDelete is the builder for deleting a ProductFeatureIactnType entity.
type ProductFeatureIactnTypeDelete struct {
	config
	hooks    []Hook
	mutation *ProductFeatureIactnTypeMutation
}

// Where adds a new predicate to the ProductFeatureIactnTypeDelete builder.
func (pfitd *ProductFeatureIactnTypeDelete) Where(ps ...predicate.ProductFeatureIactnType) *ProductFeatureIactnTypeDelete {
	pfitd.mutation.predicates = append(pfitd.mutation.predicates, ps...)
	return pfitd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pfitd *ProductFeatureIactnTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pfitd.hooks) == 0 {
		affected, err = pfitd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductFeatureIactnTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pfitd.mutation = mutation
			affected, err = pfitd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pfitd.hooks) - 1; i >= 0; i-- {
			mut = pfitd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfitd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pfitd *ProductFeatureIactnTypeDelete) ExecX(ctx context.Context) int {
	n, err := pfitd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pfitd *ProductFeatureIactnTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: productfeatureiactntype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productfeatureiactntype.FieldID,
			},
		},
	}
	if ps := pfitd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pfitd.driver, _spec)
}

// ProductFeatureIactnTypeDeleteOne is the builder for deleting a single ProductFeatureIactnType entity.
type ProductFeatureIactnTypeDeleteOne struct {
	pfitd *ProductFeatureIactnTypeDelete
}

// Exec executes the deletion query.
func (pfitdo *ProductFeatureIactnTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := pfitdo.pfitd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productfeatureiactntype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pfitdo *ProductFeatureIactnTypeDeleteOne) ExecX(ctx context.Context) {
	pfitdo.pfitd.ExecX(ctx)
}
