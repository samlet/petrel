// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayups"
)

// ShipmentGatewayUpsDelete is the builder for deleting a ShipmentGatewayUps entity.
type ShipmentGatewayUpsDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentGatewayUpsMutation
}

// Where adds a new predicate to the ShipmentGatewayUpsDelete builder.
func (sgud *ShipmentGatewayUpsDelete) Where(ps ...predicate.ShipmentGatewayUps) *ShipmentGatewayUpsDelete {
	sgud.mutation.predicates = append(sgud.mutation.predicates, ps...)
	return sgud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sgud *ShipmentGatewayUpsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sgud.hooks) == 0 {
		affected, err = sgud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentGatewayUpsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgud.mutation = mutation
			affected, err = sgud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgud.hooks) - 1; i >= 0; i-- {
			mut = sgud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgud *ShipmentGatewayUpsDelete) ExecX(ctx context.Context) int {
	n, err := sgud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sgud *ShipmentGatewayUpsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: shipmentgatewayups.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewayups.FieldID,
			},
		},
	}
	if ps := sgud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sgud.driver, _spec)
}

// ShipmentGatewayUpsDeleteOne is the builder for deleting a single ShipmentGatewayUps entity.
type ShipmentGatewayUpsDeleteOne struct {
	sgud *ShipmentGatewayUpsDelete
}

// Exec executes the deletion query.
func (sgudo *ShipmentGatewayUpsDeleteOne) Exec(ctx context.Context) error {
	n, err := sgudo.sgud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmentgatewayups.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sgudo *ShipmentGatewayUpsDeleteOne) ExecX(ctx context.Context) {
	sgudo.sgud.ExecX(ctx)
}
