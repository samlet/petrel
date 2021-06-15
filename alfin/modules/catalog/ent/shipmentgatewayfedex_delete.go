// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayfedex"
)

// ShipmentGatewayFedexDelete is the builder for deleting a ShipmentGatewayFedex entity.
type ShipmentGatewayFedexDelete struct {
	config
	hooks    []Hook
	mutation *ShipmentGatewayFedexMutation
}

// Where adds a new predicate to the ShipmentGatewayFedexDelete builder.
func (sgfd *ShipmentGatewayFedexDelete) Where(ps ...predicate.ShipmentGatewayFedex) *ShipmentGatewayFedexDelete {
	sgfd.mutation.predicates = append(sgfd.mutation.predicates, ps...)
	return sgfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sgfd *ShipmentGatewayFedexDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sgfd.hooks) == 0 {
		affected, err = sgfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentGatewayFedexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgfd.mutation = mutation
			affected, err = sgfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgfd.hooks) - 1; i >= 0; i-- {
			mut = sgfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgfd *ShipmentGatewayFedexDelete) ExecX(ctx context.Context) int {
	n, err := sgfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sgfd *ShipmentGatewayFedexDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: shipmentgatewayfedex.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewayfedex.FieldID,
			},
		},
	}
	if ps := sgfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sgfd.driver, _spec)
}

// ShipmentGatewayFedexDeleteOne is the builder for deleting a single ShipmentGatewayFedex entity.
type ShipmentGatewayFedexDeleteOne struct {
	sgfd *ShipmentGatewayFedexDelete
}

// Exec executes the deletion query.
func (sgfdo *ShipmentGatewayFedexDeleteOne) Exec(ctx context.Context) error {
	n, err := sgfdo.sgfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shipmentgatewayfedex.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sgfdo *ShipmentGatewayFedexDeleteOne) ExecX(ctx context.Context) {
	sgfdo.sgfd.ExecX(ctx)
}