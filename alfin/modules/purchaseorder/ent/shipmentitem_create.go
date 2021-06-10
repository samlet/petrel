// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/itemissuance"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipment"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipmentitem"
)

// ShipmentItemCreate is the builder for creating a ShipmentItem entity.
type ShipmentItemCreate struct {
	config
	mutation *ShipmentItemMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sic *ShipmentItemCreate) SetCreateTime(t time.Time) *ShipmentItemCreate {
	sic.mutation.SetCreateTime(t)
	return sic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableCreateTime(t *time.Time) *ShipmentItemCreate {
	if t != nil {
		sic.SetCreateTime(*t)
	}
	return sic
}

// SetUpdateTime sets the "update_time" field.
func (sic *ShipmentItemCreate) SetUpdateTime(t time.Time) *ShipmentItemCreate {
	sic.mutation.SetUpdateTime(t)
	return sic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableUpdateTime(t *time.Time) *ShipmentItemCreate {
	if t != nil {
		sic.SetUpdateTime(*t)
	}
	return sic
}

// SetShipmentItemSeqID sets the "shipment_item_seq_id" field.
func (sic *ShipmentItemCreate) SetShipmentItemSeqID(i int) *ShipmentItemCreate {
	sic.mutation.SetShipmentItemSeqID(i)
	return sic
}

// SetProductID sets the "product_id" field.
func (sic *ShipmentItemCreate) SetProductID(i int) *ShipmentItemCreate {
	sic.mutation.SetProductID(i)
	return sic
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableProductID(i *int) *ShipmentItemCreate {
	if i != nil {
		sic.SetProductID(*i)
	}
	return sic
}

// SetQuantity sets the "quantity" field.
func (sic *ShipmentItemCreate) SetQuantity(f float64) *ShipmentItemCreate {
	sic.mutation.SetQuantity(f)
	return sic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableQuantity(f *float64) *ShipmentItemCreate {
	if f != nil {
		sic.SetQuantity(*f)
	}
	return sic
}

// SetShipmentContentDescription sets the "shipment_content_description" field.
func (sic *ShipmentItemCreate) SetShipmentContentDescription(s string) *ShipmentItemCreate {
	sic.mutation.SetShipmentContentDescription(s)
	return sic
}

// SetNillableShipmentContentDescription sets the "shipment_content_description" field if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableShipmentContentDescription(s *string) *ShipmentItemCreate {
	if s != nil {
		sic.SetShipmentContentDescription(*s)
	}
	return sic
}

// SetShipmentID sets the "shipment" edge to the Shipment entity by ID.
func (sic *ShipmentItemCreate) SetShipmentID(id int) *ShipmentItemCreate {
	sic.mutation.SetShipmentID(id)
	return sic
}

// SetNillableShipmentID sets the "shipment" edge to the Shipment entity by ID if the given value is not nil.
func (sic *ShipmentItemCreate) SetNillableShipmentID(id *int) *ShipmentItemCreate {
	if id != nil {
		sic = sic.SetShipmentID(*id)
	}
	return sic
}

// SetShipment sets the "shipment" edge to the Shipment entity.
func (sic *ShipmentItemCreate) SetShipment(s *Shipment) *ShipmentItemCreate {
	return sic.SetShipmentID(s.ID)
}

// AddItemIssuanceIDs adds the "item_issuances" edge to the ItemIssuance entity by IDs.
func (sic *ShipmentItemCreate) AddItemIssuanceIDs(ids ...int) *ShipmentItemCreate {
	sic.mutation.AddItemIssuanceIDs(ids...)
	return sic
}

// AddItemIssuances adds the "item_issuances" edges to the ItemIssuance entity.
func (sic *ShipmentItemCreate) AddItemIssuances(i ...*ItemIssuance) *ShipmentItemCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sic.AddItemIssuanceIDs(ids...)
}

// Mutation returns the ShipmentItemMutation object of the builder.
func (sic *ShipmentItemCreate) Mutation() *ShipmentItemMutation {
	return sic.mutation
}

// Save creates the ShipmentItem in the database.
func (sic *ShipmentItemCreate) Save(ctx context.Context) (*ShipmentItem, error) {
	var (
		err  error
		node *ShipmentItem
	)
	sic.defaults()
	if len(sic.hooks) == 0 {
		if err = sic.check(); err != nil {
			return nil, err
		}
		node, err = sic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShipmentItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sic.check(); err != nil {
				return nil, err
			}
			sic.mutation = mutation
			node, err = sic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sic.hooks) - 1; i >= 0; i-- {
			mut = sic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sic *ShipmentItemCreate) SaveX(ctx context.Context) *ShipmentItem {
	v, err := sic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (sic *ShipmentItemCreate) defaults() {
	if _, ok := sic.mutation.CreateTime(); !ok {
		v := shipmentitem.DefaultCreateTime()
		sic.mutation.SetCreateTime(v)
	}
	if _, ok := sic.mutation.UpdateTime(); !ok {
		v := shipmentitem.DefaultUpdateTime()
		sic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sic *ShipmentItemCreate) check() error {
	if _, ok := sic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := sic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := sic.mutation.ShipmentItemSeqID(); !ok {
		return &ValidationError{Name: "shipment_item_seq_id", err: errors.New("ent: missing required field \"shipment_item_seq_id\"")}
	}
	return nil
}

func (sic *ShipmentItemCreate) sqlSave(ctx context.Context) (*ShipmentItem, error) {
	_node, _spec := sic.createSpec()
	if err := sqlgraph.CreateNode(ctx, sic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sic *ShipmentItemCreate) createSpec() (*ShipmentItem, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentItem{config: sic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shipmentitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentitem.FieldID,
			},
		}
	)
	if value, ok := sic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentitem.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shipmentitem.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sic.mutation.ShipmentItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipmentitem.FieldShipmentItemSeqID,
		})
		_node.ShipmentItemSeqID = value
	}
	if value, ok := sic.mutation.ProductID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: shipmentitem.FieldProductID,
		})
		_node.ProductID = value
	}
	if value, ok := sic.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: shipmentitem.FieldQuantity,
		})
		_node.Quantity = value
	}
	if value, ok := sic.mutation.ShipmentContentDescription(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shipmentitem.FieldShipmentContentDescription,
		})
		_node.ShipmentContentDescription = value
	}
	if nodes := sic.mutation.ShipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shipmentitem.ShipmentTable,
			Columns: []string{shipmentitem.ShipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.shipment_shipment_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sic.mutation.ItemIssuancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shipmentitem.ItemIssuancesTable,
			Columns: []string{shipmentitem.ItemIssuancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: itemissuance.FieldID,
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

// ShipmentItemCreateBulk is the builder for creating many ShipmentItem entities in bulk.
type ShipmentItemCreateBulk struct {
	config
	builders []*ShipmentItemCreate
}

// Save creates the ShipmentItem entities in the database.
func (sicb *ShipmentItemCreateBulk) Save(ctx context.Context) ([]*ShipmentItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sicb.builders))
	nodes := make([]*ShipmentItem, len(sicb.builders))
	mutators := make([]Mutator, len(sicb.builders))
	for i := range sicb.builders {
		func(i int, root context.Context) {
			builder := sicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentItemMutation)
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
					_, err = mutators[i+1].Mutate(root, sicb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sicb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
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
		if _, err := mutators[0].Mutate(ctx, sicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sicb *ShipmentItemCreateBulk) SaveX(ctx context.Context) []*ShipmentItem {
	v, err := sicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}