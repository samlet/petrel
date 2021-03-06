// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productprice"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpricepurpose"
)

// ProductPricePurposeCreate is the builder for creating a ProductPricePurpose entity.
type ProductPricePurposeCreate struct {
	config
	mutation *ProductPricePurposeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pppc *ProductPricePurposeCreate) SetCreateTime(t time.Time) *ProductPricePurposeCreate {
	pppc.mutation.SetCreateTime(t)
	return pppc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pppc *ProductPricePurposeCreate) SetNillableCreateTime(t *time.Time) *ProductPricePurposeCreate {
	if t != nil {
		pppc.SetCreateTime(*t)
	}
	return pppc
}

// SetUpdateTime sets the "update_time" field.
func (pppc *ProductPricePurposeCreate) SetUpdateTime(t time.Time) *ProductPricePurposeCreate {
	pppc.mutation.SetUpdateTime(t)
	return pppc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pppc *ProductPricePurposeCreate) SetNillableUpdateTime(t *time.Time) *ProductPricePurposeCreate {
	if t != nil {
		pppc.SetUpdateTime(*t)
	}
	return pppc
}

// SetStringRef sets the "string_ref" field.
func (pppc *ProductPricePurposeCreate) SetStringRef(s string) *ProductPricePurposeCreate {
	pppc.mutation.SetStringRef(s)
	return pppc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pppc *ProductPricePurposeCreate) SetNillableStringRef(s *string) *ProductPricePurposeCreate {
	if s != nil {
		pppc.SetStringRef(*s)
	}
	return pppc
}

// SetDescription sets the "description" field.
func (pppc *ProductPricePurposeCreate) SetDescription(s string) *ProductPricePurposeCreate {
	pppc.mutation.SetDescription(s)
	return pppc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pppc *ProductPricePurposeCreate) SetNillableDescription(s *string) *ProductPricePurposeCreate {
	if s != nil {
		pppc.SetDescription(*s)
	}
	return pppc
}

// AddOrderPaymentPreferenceIDs adds the "order_payment_preferences" edge to the OrderPaymentPreference entity by IDs.
func (pppc *ProductPricePurposeCreate) AddOrderPaymentPreferenceIDs(ids ...int) *ProductPricePurposeCreate {
	pppc.mutation.AddOrderPaymentPreferenceIDs(ids...)
	return pppc
}

// AddOrderPaymentPreferences adds the "order_payment_preferences" edges to the OrderPaymentPreference entity.
func (pppc *ProductPricePurposeCreate) AddOrderPaymentPreferences(o ...*OrderPaymentPreference) *ProductPricePurposeCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pppc.AddOrderPaymentPreferenceIDs(ids...)
}

// AddProductPriceIDs adds the "product_prices" edge to the ProductPrice entity by IDs.
func (pppc *ProductPricePurposeCreate) AddProductPriceIDs(ids ...int) *ProductPricePurposeCreate {
	pppc.mutation.AddProductPriceIDs(ids...)
	return pppc
}

// AddProductPrices adds the "product_prices" edges to the ProductPrice entity.
func (pppc *ProductPricePurposeCreate) AddProductPrices(p ...*ProductPrice) *ProductPricePurposeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pppc.AddProductPriceIDs(ids...)
}

// Mutation returns the ProductPricePurposeMutation object of the builder.
func (pppc *ProductPricePurposeCreate) Mutation() *ProductPricePurposeMutation {
	return pppc.mutation
}

// Save creates the ProductPricePurpose in the database.
func (pppc *ProductPricePurposeCreate) Save(ctx context.Context) (*ProductPricePurpose, error) {
	var (
		err  error
		node *ProductPricePurpose
	)
	pppc.defaults()
	if len(pppc.hooks) == 0 {
		if err = pppc.check(); err != nil {
			return nil, err
		}
		node, err = pppc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductPricePurposeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pppc.check(); err != nil {
				return nil, err
			}
			pppc.mutation = mutation
			if node, err = pppc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pppc.hooks) - 1; i >= 0; i-- {
			mut = pppc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pppc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pppc *ProductPricePurposeCreate) SaveX(ctx context.Context) *ProductPricePurpose {
	v, err := pppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pppc *ProductPricePurposeCreate) defaults() {
	if _, ok := pppc.mutation.CreateTime(); !ok {
		v := productpricepurpose.DefaultCreateTime()
		pppc.mutation.SetCreateTime(v)
	}
	if _, ok := pppc.mutation.UpdateTime(); !ok {
		v := productpricepurpose.DefaultUpdateTime()
		pppc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pppc *ProductPricePurposeCreate) check() error {
	if _, ok := pppc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pppc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (pppc *ProductPricePurposeCreate) sqlSave(ctx context.Context) (*ProductPricePurpose, error) {
	_node, _spec := pppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pppc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pppc *ProductPricePurposeCreate) createSpec() (*ProductPricePurpose, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductPricePurpose{config: pppc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productpricepurpose.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productpricepurpose.FieldID,
			},
		}
	)
	if value, ok := pppc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productpricepurpose.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pppc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productpricepurpose.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pppc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productpricepurpose.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pppc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productpricepurpose.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pppc.mutation.OrderPaymentPreferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productpricepurpose.OrderPaymentPreferencesTable,
			Columns: []string{productpricepurpose.OrderPaymentPreferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderpaymentpreference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pppc.mutation.ProductPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productpricepurpose.ProductPricesTable,
			Columns: []string{productpricepurpose.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductPricePurposeCreateBulk is the builder for creating many ProductPricePurpose entities in bulk.
type ProductPricePurposeCreateBulk struct {
	config
	builders []*ProductPricePurposeCreate
}

// Save creates the ProductPricePurpose entities in the database.
func (pppcb *ProductPricePurposeCreateBulk) Save(ctx context.Context) ([]*ProductPricePurpose, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pppcb.builders))
	nodes := make([]*ProductPricePurpose, len(pppcb.builders))
	mutators := make([]Mutator, len(pppcb.builders))
	for i := range pppcb.builders {
		func(i int, root context.Context) {
			builder := pppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductPricePurposeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pppcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pppcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pppcb *ProductPricePurposeCreateBulk) SaveX(ctx context.Context) []*ProductPricePurpose {
	v, err := pppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
