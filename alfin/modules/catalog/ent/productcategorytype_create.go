// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategory"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategorytype"
)

// ProductCategoryTypeCreate is the builder for creating a ProductCategoryType entity.
type ProductCategoryTypeCreate struct {
	config
	mutation *ProductCategoryTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pctc *ProductCategoryTypeCreate) SetCreateTime(t time.Time) *ProductCategoryTypeCreate {
	pctc.mutation.SetCreateTime(t)
	return pctc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableCreateTime(t *time.Time) *ProductCategoryTypeCreate {
	if t != nil {
		pctc.SetCreateTime(*t)
	}
	return pctc
}

// SetUpdateTime sets the "update_time" field.
func (pctc *ProductCategoryTypeCreate) SetUpdateTime(t time.Time) *ProductCategoryTypeCreate {
	pctc.mutation.SetUpdateTime(t)
	return pctc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductCategoryTypeCreate {
	if t != nil {
		pctc.SetUpdateTime(*t)
	}
	return pctc
}

// SetStringRef sets the "string_ref" field.
func (pctc *ProductCategoryTypeCreate) SetStringRef(s string) *ProductCategoryTypeCreate {
	pctc.mutation.SetStringRef(s)
	return pctc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableStringRef(s *string) *ProductCategoryTypeCreate {
	if s != nil {
		pctc.SetStringRef(*s)
	}
	return pctc
}

// SetHasTable sets the "has_table" field.
func (pctc *ProductCategoryTypeCreate) SetHasTable(pt productcategorytype.HasTable) *ProductCategoryTypeCreate {
	pctc.mutation.SetHasTable(pt)
	return pctc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableHasTable(pt *productcategorytype.HasTable) *ProductCategoryTypeCreate {
	if pt != nil {
		pctc.SetHasTable(*pt)
	}
	return pctc
}

// SetDescription sets the "description" field.
func (pctc *ProductCategoryTypeCreate) SetDescription(s string) *ProductCategoryTypeCreate {
	pctc.mutation.SetDescription(s)
	return pctc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableDescription(s *string) *ProductCategoryTypeCreate {
	if s != nil {
		pctc.SetDescription(*s)
	}
	return pctc
}

// SetParentID sets the "parent" edge to the ProductCategoryType entity by ID.
func (pctc *ProductCategoryTypeCreate) SetParentID(id int) *ProductCategoryTypeCreate {
	pctc.mutation.SetParentID(id)
	return pctc
}

// SetNillableParentID sets the "parent" edge to the ProductCategoryType entity by ID if the given value is not nil.
func (pctc *ProductCategoryTypeCreate) SetNillableParentID(id *int) *ProductCategoryTypeCreate {
	if id != nil {
		pctc = pctc.SetParentID(*id)
	}
	return pctc
}

// SetParent sets the "parent" edge to the ProductCategoryType entity.
func (pctc *ProductCategoryTypeCreate) SetParent(p *ProductCategoryType) *ProductCategoryTypeCreate {
	return pctc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductCategoryType entity by IDs.
func (pctc *ProductCategoryTypeCreate) AddChildIDs(ids ...int) *ProductCategoryTypeCreate {
	pctc.mutation.AddChildIDs(ids...)
	return pctc
}

// AddChildren adds the "children" edges to the ProductCategoryType entity.
func (pctc *ProductCategoryTypeCreate) AddChildren(p ...*ProductCategoryType) *ProductCategoryTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pctc.AddChildIDs(ids...)
}

// AddProductCategoryIDs adds the "product_categories" edge to the ProductCategory entity by IDs.
func (pctc *ProductCategoryTypeCreate) AddProductCategoryIDs(ids ...int) *ProductCategoryTypeCreate {
	pctc.mutation.AddProductCategoryIDs(ids...)
	return pctc
}

// AddProductCategories adds the "product_categories" edges to the ProductCategory entity.
func (pctc *ProductCategoryTypeCreate) AddProductCategories(p ...*ProductCategory) *ProductCategoryTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pctc.AddProductCategoryIDs(ids...)
}

// AddChildProductCategoryTypeIDs adds the "child_product_category_types" edge to the ProductCategoryType entity by IDs.
func (pctc *ProductCategoryTypeCreate) AddChildProductCategoryTypeIDs(ids ...int) *ProductCategoryTypeCreate {
	pctc.mutation.AddChildProductCategoryTypeIDs(ids...)
	return pctc
}

// AddChildProductCategoryTypes adds the "child_product_category_types" edges to the ProductCategoryType entity.
func (pctc *ProductCategoryTypeCreate) AddChildProductCategoryTypes(p ...*ProductCategoryType) *ProductCategoryTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pctc.AddChildProductCategoryTypeIDs(ids...)
}

// Mutation returns the ProductCategoryTypeMutation object of the builder.
func (pctc *ProductCategoryTypeCreate) Mutation() *ProductCategoryTypeMutation {
	return pctc.mutation
}

// Save creates the ProductCategoryType in the database.
func (pctc *ProductCategoryTypeCreate) Save(ctx context.Context) (*ProductCategoryType, error) {
	var (
		err  error
		node *ProductCategoryType
	)
	pctc.defaults()
	if len(pctc.hooks) == 0 {
		if err = pctc.check(); err != nil {
			return nil, err
		}
		node, err = pctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pctc.check(); err != nil {
				return nil, err
			}
			pctc.mutation = mutation
			if node, err = pctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pctc.hooks) - 1; i >= 0; i-- {
			mut = pctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pctc *ProductCategoryTypeCreate) SaveX(ctx context.Context) *ProductCategoryType {
	v, err := pctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pctc *ProductCategoryTypeCreate) defaults() {
	if _, ok := pctc.mutation.CreateTime(); !ok {
		v := productcategorytype.DefaultCreateTime()
		pctc.mutation.SetCreateTime(v)
	}
	if _, ok := pctc.mutation.UpdateTime(); !ok {
		v := productcategorytype.DefaultUpdateTime()
		pctc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pctc *ProductCategoryTypeCreate) check() error {
	if _, ok := pctc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pctc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pctc.mutation.HasTable(); ok {
		if err := productcategorytype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pctc *ProductCategoryTypeCreate) sqlSave(ctx context.Context) (*ProductCategoryType, error) {
	_node, _spec := pctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pctc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pctc *ProductCategoryTypeCreate) createSpec() (*ProductCategoryType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductCategoryType{config: pctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productcategorytype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcategorytype.FieldID,
			},
		}
	)
	if value, ok := pctc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productcategorytype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pctc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productcategorytype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pctc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productcategorytype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pctc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: productcategorytype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pctc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productcategorytype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pctc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategorytype.ParentTable,
			Columns: []string{productcategorytype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_category_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pctc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorytype.ChildrenTable,
			Columns: []string{productcategorytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pctc.mutation.ProductCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorytype.ProductCategoriesTable,
			Columns: []string{productcategorytype.ProductCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pctc.mutation.ChildProductCategoryTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productcategorytype.ChildProductCategoryTypesTable,
			Columns: productcategorytype.ChildProductCategoryTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorytype.FieldID,
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

// ProductCategoryTypeCreateBulk is the builder for creating many ProductCategoryType entities in bulk.
type ProductCategoryTypeCreateBulk struct {
	config
	builders []*ProductCategoryTypeCreate
}

// Save creates the ProductCategoryType entities in the database.
func (pctcb *ProductCategoryTypeCreateBulk) Save(ctx context.Context) ([]*ProductCategoryType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pctcb.builders))
	nodes := make([]*ProductCategoryType, len(pctcb.builders))
	mutators := make([]Mutator, len(pctcb.builders))
	for i := range pctcb.builders {
		func(i int, root context.Context) {
			builder := pctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductCategoryTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pctcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pctcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pctcb *ProductCategoryTypeCreateBulk) SaveX(ctx context.Context) []*ProductCategoryType {
	v, err := pctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
