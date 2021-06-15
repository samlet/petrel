// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/inventoryitemtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/product"
)

// InventoryItemTypeCreate is the builder for creating a InventoryItemType entity.
type InventoryItemTypeCreate struct {
	config
	mutation *InventoryItemTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (iitc *InventoryItemTypeCreate) SetCreateTime(t time.Time) *InventoryItemTypeCreate {
	iitc.mutation.SetCreateTime(t)
	return iitc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableCreateTime(t *time.Time) *InventoryItemTypeCreate {
	if t != nil {
		iitc.SetCreateTime(*t)
	}
	return iitc
}

// SetUpdateTime sets the "update_time" field.
func (iitc *InventoryItemTypeCreate) SetUpdateTime(t time.Time) *InventoryItemTypeCreate {
	iitc.mutation.SetUpdateTime(t)
	return iitc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableUpdateTime(t *time.Time) *InventoryItemTypeCreate {
	if t != nil {
		iitc.SetUpdateTime(*t)
	}
	return iitc
}

// SetStringRef sets the "string_ref" field.
func (iitc *InventoryItemTypeCreate) SetStringRef(s string) *InventoryItemTypeCreate {
	iitc.mutation.SetStringRef(s)
	return iitc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableStringRef(s *string) *InventoryItemTypeCreate {
	if s != nil {
		iitc.SetStringRef(*s)
	}
	return iitc
}

// SetHasTable sets the "has_table" field.
func (iitc *InventoryItemTypeCreate) SetHasTable(it inventoryitemtype.HasTable) *InventoryItemTypeCreate {
	iitc.mutation.SetHasTable(it)
	return iitc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableHasTable(it *inventoryitemtype.HasTable) *InventoryItemTypeCreate {
	if it != nil {
		iitc.SetHasTable(*it)
	}
	return iitc
}

// SetDescription sets the "description" field.
func (iitc *InventoryItemTypeCreate) SetDescription(s string) *InventoryItemTypeCreate {
	iitc.mutation.SetDescription(s)
	return iitc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableDescription(s *string) *InventoryItemTypeCreate {
	if s != nil {
		iitc.SetDescription(*s)
	}
	return iitc
}

// SetParentID sets the "parent" edge to the InventoryItemType entity by ID.
func (iitc *InventoryItemTypeCreate) SetParentID(id int) *InventoryItemTypeCreate {
	iitc.mutation.SetParentID(id)
	return iitc
}

// SetNillableParentID sets the "parent" edge to the InventoryItemType entity by ID if the given value is not nil.
func (iitc *InventoryItemTypeCreate) SetNillableParentID(id *int) *InventoryItemTypeCreate {
	if id != nil {
		iitc = iitc.SetParentID(*id)
	}
	return iitc
}

// SetParent sets the "parent" edge to the InventoryItemType entity.
func (iitc *InventoryItemTypeCreate) SetParent(i *InventoryItemType) *InventoryItemTypeCreate {
	return iitc.SetParentID(i.ID)
}

// AddChildIDs adds the "children" edge to the InventoryItemType entity by IDs.
func (iitc *InventoryItemTypeCreate) AddChildIDs(ids ...int) *InventoryItemTypeCreate {
	iitc.mutation.AddChildIDs(ids...)
	return iitc
}

// AddChildren adds the "children" edges to the InventoryItemType entity.
func (iitc *InventoryItemTypeCreate) AddChildren(i ...*InventoryItemType) *InventoryItemTypeCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iitc.AddChildIDs(ids...)
}

// AddChildInventoryItemTypeIDs adds the "child_inventory_item_types" edge to the InventoryItemType entity by IDs.
func (iitc *InventoryItemTypeCreate) AddChildInventoryItemTypeIDs(ids ...int) *InventoryItemTypeCreate {
	iitc.mutation.AddChildInventoryItemTypeIDs(ids...)
	return iitc
}

// AddChildInventoryItemTypes adds the "child_inventory_item_types" edges to the InventoryItemType entity.
func (iitc *InventoryItemTypeCreate) AddChildInventoryItemTypes(i ...*InventoryItemType) *InventoryItemTypeCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iitc.AddChildInventoryItemTypeIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (iitc *InventoryItemTypeCreate) AddProductIDs(ids ...int) *InventoryItemTypeCreate {
	iitc.mutation.AddProductIDs(ids...)
	return iitc
}

// AddProducts adds the "products" edges to the Product entity.
func (iitc *InventoryItemTypeCreate) AddProducts(p ...*Product) *InventoryItemTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iitc.AddProductIDs(ids...)
}

// Mutation returns the InventoryItemTypeMutation object of the builder.
func (iitc *InventoryItemTypeCreate) Mutation() *InventoryItemTypeMutation {
	return iitc.mutation
}

// Save creates the InventoryItemType in the database.
func (iitc *InventoryItemTypeCreate) Save(ctx context.Context) (*InventoryItemType, error) {
	var (
		err  error
		node *InventoryItemType
	)
	iitc.defaults()
	if len(iitc.hooks) == 0 {
		if err = iitc.check(); err != nil {
			return nil, err
		}
		node, err = iitc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InventoryItemTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iitc.check(); err != nil {
				return nil, err
			}
			iitc.mutation = mutation
			if node, err = iitc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(iitc.hooks) - 1; i >= 0; i-- {
			mut = iitc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iitc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (iitc *InventoryItemTypeCreate) SaveX(ctx context.Context) *InventoryItemType {
	v, err := iitc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (iitc *InventoryItemTypeCreate) defaults() {
	if _, ok := iitc.mutation.CreateTime(); !ok {
		v := inventoryitemtype.DefaultCreateTime()
		iitc.mutation.SetCreateTime(v)
	}
	if _, ok := iitc.mutation.UpdateTime(); !ok {
		v := inventoryitemtype.DefaultUpdateTime()
		iitc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iitc *InventoryItemTypeCreate) check() error {
	if _, ok := iitc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := iitc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := iitc.mutation.HasTable(); ok {
		if err := inventoryitemtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (iitc *InventoryItemTypeCreate) sqlSave(ctx context.Context) (*InventoryItemType, error) {
	_node, _spec := iitc.createSpec()
	if err := sqlgraph.CreateNode(ctx, iitc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (iitc *InventoryItemTypeCreate) createSpec() (*InventoryItemType, *sqlgraph.CreateSpec) {
	var (
		_node = &InventoryItemType{config: iitc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: inventoryitemtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inventoryitemtype.FieldID,
			},
		}
	)
	if value, ok := iitc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inventoryitemtype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := iitc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inventoryitemtype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := iitc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryitemtype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := iitc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: inventoryitemtype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := iitc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inventoryitemtype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := iitc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inventoryitemtype.ParentTable,
			Columns: []string{inventoryitemtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inventoryitemtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inventory_item_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iitc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventoryitemtype.ChildrenTable,
			Columns: []string{inventoryitemtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inventoryitemtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iitc.mutation.ChildInventoryItemTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   inventoryitemtype.ChildInventoryItemTypesTable,
			Columns: inventoryitemtype.ChildInventoryItemTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inventoryitemtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := iitc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventoryitemtype.ProductsTable,
			Columns: []string{inventoryitemtype.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
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

// InventoryItemTypeCreateBulk is the builder for creating many InventoryItemType entities in bulk.
type InventoryItemTypeCreateBulk struct {
	config
	builders []*InventoryItemTypeCreate
}

// Save creates the InventoryItemType entities in the database.
func (iitcb *InventoryItemTypeCreateBulk) Save(ctx context.Context) ([]*InventoryItemType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(iitcb.builders))
	nodes := make([]*InventoryItemType, len(iitcb.builders))
	mutators := make([]Mutator, len(iitcb.builders))
	for i := range iitcb.builders {
		func(i int, root context.Context) {
			builder := iitcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InventoryItemTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, iitcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iitcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iitcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iitcb *InventoryItemTypeCreateBulk) SaveX(ctx context.Context) []*InventoryItemType {
	v, err := iitcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
