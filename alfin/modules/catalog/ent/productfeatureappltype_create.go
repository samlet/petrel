// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeatureappltype"
)

// ProductFeatureApplTypeCreate is the builder for creating a ProductFeatureApplType entity.
type ProductFeatureApplTypeCreate struct {
	config
	mutation *ProductFeatureApplTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pfatc *ProductFeatureApplTypeCreate) SetCreateTime(t time.Time) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetCreateTime(t)
	return pfatc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableCreateTime(t *time.Time) *ProductFeatureApplTypeCreate {
	if t != nil {
		pfatc.SetCreateTime(*t)
	}
	return pfatc
}

// SetUpdateTime sets the "update_time" field.
func (pfatc *ProductFeatureApplTypeCreate) SetUpdateTime(t time.Time) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetUpdateTime(t)
	return pfatc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductFeatureApplTypeCreate {
	if t != nil {
		pfatc.SetUpdateTime(*t)
	}
	return pfatc
}

// SetStringRef sets the "string_ref" field.
func (pfatc *ProductFeatureApplTypeCreate) SetStringRef(s string) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetStringRef(s)
	return pfatc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableStringRef(s *string) *ProductFeatureApplTypeCreate {
	if s != nil {
		pfatc.SetStringRef(*s)
	}
	return pfatc
}

// SetHasTable sets the "has_table" field.
func (pfatc *ProductFeatureApplTypeCreate) SetHasTable(pt productfeatureappltype.HasTable) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetHasTable(pt)
	return pfatc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableHasTable(pt *productfeatureappltype.HasTable) *ProductFeatureApplTypeCreate {
	if pt != nil {
		pfatc.SetHasTable(*pt)
	}
	return pfatc
}

// SetDescription sets the "description" field.
func (pfatc *ProductFeatureApplTypeCreate) SetDescription(s string) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetDescription(s)
	return pfatc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableDescription(s *string) *ProductFeatureApplTypeCreate {
	if s != nil {
		pfatc.SetDescription(*s)
	}
	return pfatc
}

// SetParentID sets the "parent" edge to the ProductFeatureApplType entity by ID.
func (pfatc *ProductFeatureApplTypeCreate) SetParentID(id int) *ProductFeatureApplTypeCreate {
	pfatc.mutation.SetParentID(id)
	return pfatc
}

// SetNillableParentID sets the "parent" edge to the ProductFeatureApplType entity by ID if the given value is not nil.
func (pfatc *ProductFeatureApplTypeCreate) SetNillableParentID(id *int) *ProductFeatureApplTypeCreate {
	if id != nil {
		pfatc = pfatc.SetParentID(*id)
	}
	return pfatc
}

// SetParent sets the "parent" edge to the ProductFeatureApplType entity.
func (pfatc *ProductFeatureApplTypeCreate) SetParent(p *ProductFeatureApplType) *ProductFeatureApplTypeCreate {
	return pfatc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductFeatureApplType entity by IDs.
func (pfatc *ProductFeatureApplTypeCreate) AddChildIDs(ids ...int) *ProductFeatureApplTypeCreate {
	pfatc.mutation.AddChildIDs(ids...)
	return pfatc
}

// AddChildren adds the "children" edges to the ProductFeatureApplType entity.
func (pfatc *ProductFeatureApplTypeCreate) AddChildren(p ...*ProductFeatureApplType) *ProductFeatureApplTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pfatc.AddChildIDs(ids...)
}

// AddChildProductFeatureApplTypeIDs adds the "child_product_feature_appl_types" edge to the ProductFeatureApplType entity by IDs.
func (pfatc *ProductFeatureApplTypeCreate) AddChildProductFeatureApplTypeIDs(ids ...int) *ProductFeatureApplTypeCreate {
	pfatc.mutation.AddChildProductFeatureApplTypeIDs(ids...)
	return pfatc
}

// AddChildProductFeatureApplTypes adds the "child_product_feature_appl_types" edges to the ProductFeatureApplType entity.
func (pfatc *ProductFeatureApplTypeCreate) AddChildProductFeatureApplTypes(p ...*ProductFeatureApplType) *ProductFeatureApplTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pfatc.AddChildProductFeatureApplTypeIDs(ids...)
}

// Mutation returns the ProductFeatureApplTypeMutation object of the builder.
func (pfatc *ProductFeatureApplTypeCreate) Mutation() *ProductFeatureApplTypeMutation {
	return pfatc.mutation
}

// Save creates the ProductFeatureApplType in the database.
func (pfatc *ProductFeatureApplTypeCreate) Save(ctx context.Context) (*ProductFeatureApplType, error) {
	var (
		err  error
		node *ProductFeatureApplType
	)
	pfatc.defaults()
	if len(pfatc.hooks) == 0 {
		if err = pfatc.check(); err != nil {
			return nil, err
		}
		node, err = pfatc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductFeatureApplTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pfatc.check(); err != nil {
				return nil, err
			}
			pfatc.mutation = mutation
			if node, err = pfatc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pfatc.hooks) - 1; i >= 0; i-- {
			mut = pfatc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pfatc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pfatc *ProductFeatureApplTypeCreate) SaveX(ctx context.Context) *ProductFeatureApplType {
	v, err := pfatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pfatc *ProductFeatureApplTypeCreate) defaults() {
	if _, ok := pfatc.mutation.CreateTime(); !ok {
		v := productfeatureappltype.DefaultCreateTime()
		pfatc.mutation.SetCreateTime(v)
	}
	if _, ok := pfatc.mutation.UpdateTime(); !ok {
		v := productfeatureappltype.DefaultUpdateTime()
		pfatc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pfatc *ProductFeatureApplTypeCreate) check() error {
	if _, ok := pfatc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pfatc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pfatc.mutation.HasTable(); ok {
		if err := productfeatureappltype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pfatc *ProductFeatureApplTypeCreate) sqlSave(ctx context.Context) (*ProductFeatureApplType, error) {
	_node, _spec := pfatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pfatc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pfatc *ProductFeatureApplTypeCreate) createSpec() (*ProductFeatureApplType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductFeatureApplType{config: pfatc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productfeatureappltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productfeatureappltype.FieldID,
			},
		}
	)
	if value, ok := pfatc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productfeatureappltype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pfatc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productfeatureappltype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pfatc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productfeatureappltype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pfatc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: productfeatureappltype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pfatc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productfeatureappltype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pfatc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productfeatureappltype.ParentTable,
			Columns: []string{productfeatureappltype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureappltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_feature_appl_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfatc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productfeatureappltype.ChildrenTable,
			Columns: []string{productfeatureappltype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureappltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pfatc.mutation.ChildProductFeatureApplTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productfeatureappltype.ChildProductFeatureApplTypesTable,
			Columns: productfeatureappltype.ChildProductFeatureApplTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productfeatureappltype.FieldID,
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

// ProductFeatureApplTypeCreateBulk is the builder for creating many ProductFeatureApplType entities in bulk.
type ProductFeatureApplTypeCreateBulk struct {
	config
	builders []*ProductFeatureApplTypeCreate
}

// Save creates the ProductFeatureApplType entities in the database.
func (pfatcb *ProductFeatureApplTypeCreateBulk) Save(ctx context.Context) ([]*ProductFeatureApplType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pfatcb.builders))
	nodes := make([]*ProductFeatureApplType, len(pfatcb.builders))
	mutators := make([]Mutator, len(pfatcb.builders))
	for i := range pfatcb.builders {
		func(i int, root context.Context) {
			builder := pfatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductFeatureApplTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pfatcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pfatcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pfatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pfatcb *ProductFeatureApplTypeCreateBulk) SaveX(ctx context.Context) []*ProductFeatureApplType {
	v, err := pfatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
