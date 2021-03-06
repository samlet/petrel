// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeatureiactntype"
)

// ProductFeatureIactnTypeCreate is the builder for creating a ProductFeatureIactnType entity.
type ProductFeatureIactnTypeCreate struct {
	config
	mutation *ProductFeatureIactnTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pfitc *ProductFeatureIactnTypeCreate) SetCreateTime(t time.Time) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetCreateTime(t)
	return pfitc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableCreateTime(t *time.Time) *ProductFeatureIactnTypeCreate {
	if t != nil {
		pfitc.SetCreateTime(*t)
	}
	return pfitc
}

// SetUpdateTime sets the "update_time" field.
func (pfitc *ProductFeatureIactnTypeCreate) SetUpdateTime(t time.Time) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetUpdateTime(t)
	return pfitc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductFeatureIactnTypeCreate {
	if t != nil {
		pfitc.SetUpdateTime(*t)
	}
	return pfitc
}

// SetStringRef sets the "string_ref" field.
func (pfitc *ProductFeatureIactnTypeCreate) SetStringRef(s string) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetStringRef(s)
	return pfitc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableStringRef(s *string) *ProductFeatureIactnTypeCreate {
	if s != nil {
		pfitc.SetStringRef(*s)
	}
	return pfitc
}

// SetHasTable sets the "has_table" field.
func (pfitc *ProductFeatureIactnTypeCreate) SetHasTable(pt productfeatureiactntype.HasTable) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetHasTable(pt)
	return pfitc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableHasTable(pt *productfeatureiactntype.HasTable) *ProductFeatureIactnTypeCreate {
	if pt != nil {
		pfitc.SetHasTable(*pt)
	}
	return pfitc
}

// SetDescription sets the "description" field.
func (pfitc *ProductFeatureIactnTypeCreate) SetDescription(s string) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetDescription(s)
	return pfitc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableDescription(s *string) *ProductFeatureIactnTypeCreate {
	if s != nil {
		pfitc.SetDescription(*s)
	}
	return pfitc
}

// SetParentID sets the "parent" edge to the ProductFeatureIactnType entity by ID.
func (pfitc *ProductFeatureIactnTypeCreate) SetParentID(id int) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.SetParentID(id)
	return pfitc
}

// SetNillableParentID sets the "parent" edge to the ProductFeatureIactnType entity by ID if the given value is not nil.
func (pfitc *ProductFeatureIactnTypeCreate) SetNillableParentID(id *int) *ProductFeatureIactnTypeCreate {
	if id != nil {
		pfitc = pfitc.SetParentID(*id)
	}
	return pfitc
}

// SetParent sets the "parent" edge to the ProductFeatureIactnType entity.
func (pfitc *ProductFeatureIactnTypeCreate) SetParent(p *ProductFeatureIactnType) *ProductFeatureIactnTypeCreate {
	return pfitc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductFeatureIactnType entity by IDs.
func (pfitc *ProductFeatureIactnTypeCreate) AddChildIDs(ids ...int) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.AddChildIDs(ids...)
	return pfitc
}

// AddChildren adds the "children" edges to the ProductFeatureIactnType entity.
func (pfitc *ProductFeatureIactnTypeCreate) AddChildren(p ...*ProductFeatureIactnType) *ProductFeatureIactnTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pfitc.AddChildIDs(ids...)
}

// AddChildProductFeatureIactnTypeIDs adds the "child_product_feature_iactn_types" edge to the ProductFeatureIactnType entity by IDs.
func (pfitc *ProductFeatureIactnTypeCreate) AddChildProductFeatureIactnTypeIDs(ids ...int) *ProductFeatureIactnTypeCreate {
	pfitc.mutation.AddChildProductFeatureIactnTypeIDs(ids...)
	return pfitc
}

// AddChildProductFeatureIactnTypes adds the "child_product_feature_iactn_types" edges to the ProductFeatureIactnType entity.
func (pfitc *ProductFeatureIactnTypeCreate) AddChildProductFeatureIactnTypes(p ...*ProductFeatureIactnType) *ProductFeatureIactnTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pfitc.AddChildProductFeatureIactnTypeIDs(ids...)
}

// Mutation returns the ProductFeatureIactnTypeMutation object of the builder.
func (pfitc *ProductFeatureIactnTypeCreate) Mutation() *ProductFeatureIactnTypeMutation {
	return pfitc.mutation
}

// Save creates the ProductFeatureIactnType in the database.
func (pfitc *ProductFeatureIactnTypeCreate) Save(ctx context.Context) (*ProductFeatureIactnType, error) {
	var (
		err  error
		node *ProductFeatureIactnType
	)
	pfitc.defaults()
	if len(pfitc.hooks) == 0 {
		if err = pfitc.check(); err != nil {
			return nil, err
		}
		node, err = pfitc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductFeatureIactnTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pfitc.check(); err != nil {
				return nil, err
			}
			pfitc.mutation = mutation
			if node, err = pfitc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pfitc.hooks) - 1; i >= 0; i-- {
			mut = pfitc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfitc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pfitc *ProductFeatureIactnTypeCreate) SaveX(ctx context.Context) *ProductFeatureIactnType {
	v, err := pfitc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pfitc *ProductFeatureIactnTypeCreate) defaults() {
	if _, ok := pfitc.mutation.CreateTime(); !ok {
		v := productfeatureiactntype.DefaultCreateTime()
		pfitc.mutation.SetCreateTime(v)
	}
	if _, ok := pfitc.mutation.UpdateTime(); !ok {
		v := productfeatureiactntype.DefaultUpdateTime()
		pfitc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfitc *ProductFeatureIactnTypeCreate) check() error {
	if _, ok := pfitc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pfitc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pfitc.mutation.HasTable(); ok {
		if err := productfeatureiactntype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pfitc *ProductFeatureIactnTypeCreate) sqlSave(ctx context.Context) (*ProductFeatureIactnType, error) {
	_node, _spec := pfitc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfitc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pfitc *ProductFeatureIactnTypeCreate) createSpec() (*ProductFeatureIactnType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductFeatureIactnType{config: pfitc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productfeatureiactntype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productfeatureiactntype.FieldID,
			},
		}
	)
	if value, ok := pfitc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productfeatureiactntype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pfitc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productfeatureiactntype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pfitc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productfeatureiactntype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pfitc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: productfeatureiactntype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pfitc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productfeatureiactntype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pfitc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productfeatureiactntype.ParentTable,
			Columns: []string{productfeatureiactntype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureiactntype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_feature_iactn_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfitc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productfeatureiactntype.ChildrenTable,
			Columns: []string{productfeatureiactntype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureiactntype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfitc.mutation.ChildProductFeatureIactnTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productfeatureiactntype.ChildProductFeatureIactnTypesTable,
			Columns: productfeatureiactntype.ChildProductFeatureIactnTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureiactntype.FieldID,
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

// ProductFeatureIactnTypeCreateBulk is the builder for creating many ProductFeatureIactnType entities in bulk.
type ProductFeatureIactnTypeCreateBulk struct {
	config
	builders []*ProductFeatureIactnTypeCreate
}

// Save creates the ProductFeatureIactnType entities in the database.
func (pfitcb *ProductFeatureIactnTypeCreateBulk) Save(ctx context.Context) ([]*ProductFeatureIactnType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pfitcb.builders))
	nodes := make([]*ProductFeatureIactnType, len(pfitcb.builders))
	mutators := make([]Mutator, len(pfitcb.builders))
	for i := range pfitcb.builders {
		func(i int, root context.Context) {
			builder := pfitcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductFeatureIactnTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pfitcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfitcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pfitcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfitcb *ProductFeatureIactnTypeCreateBulk) SaveX(ctx context.Context) []*ProductFeatureIactnType {
	v, err := pfitcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
