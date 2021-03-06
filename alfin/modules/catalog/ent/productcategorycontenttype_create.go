// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productcategorycontenttype"
)

// ProductCategoryContentTypeCreate is the builder for creating a ProductCategoryContentType entity.
type ProductCategoryContentTypeCreate struct {
	config
	mutation *ProductCategoryContentTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pcctc *ProductCategoryContentTypeCreate) SetCreateTime(t time.Time) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetCreateTime(t)
	return pcctc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableCreateTime(t *time.Time) *ProductCategoryContentTypeCreate {
	if t != nil {
		pcctc.SetCreateTime(*t)
	}
	return pcctc
}

// SetUpdateTime sets the "update_time" field.
func (pcctc *ProductCategoryContentTypeCreate) SetUpdateTime(t time.Time) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetUpdateTime(t)
	return pcctc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableUpdateTime(t *time.Time) *ProductCategoryContentTypeCreate {
	if t != nil {
		pcctc.SetUpdateTime(*t)
	}
	return pcctc
}

// SetStringRef sets the "string_ref" field.
func (pcctc *ProductCategoryContentTypeCreate) SetStringRef(s string) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetStringRef(s)
	return pcctc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableStringRef(s *string) *ProductCategoryContentTypeCreate {
	if s != nil {
		pcctc.SetStringRef(*s)
	}
	return pcctc
}

// SetHasTable sets the "has_table" field.
func (pcctc *ProductCategoryContentTypeCreate) SetHasTable(pt productcategorycontenttype.HasTable) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetHasTable(pt)
	return pcctc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableHasTable(pt *productcategorycontenttype.HasTable) *ProductCategoryContentTypeCreate {
	if pt != nil {
		pcctc.SetHasTable(*pt)
	}
	return pcctc
}

// SetDescription sets the "description" field.
func (pcctc *ProductCategoryContentTypeCreate) SetDescription(s string) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetDescription(s)
	return pcctc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableDescription(s *string) *ProductCategoryContentTypeCreate {
	if s != nil {
		pcctc.SetDescription(*s)
	}
	return pcctc
}

// SetParentID sets the "parent" edge to the ProductCategoryContentType entity by ID.
func (pcctc *ProductCategoryContentTypeCreate) SetParentID(id int) *ProductCategoryContentTypeCreate {
	pcctc.mutation.SetParentID(id)
	return pcctc
}

// SetNillableParentID sets the "parent" edge to the ProductCategoryContentType entity by ID if the given value is not nil.
func (pcctc *ProductCategoryContentTypeCreate) SetNillableParentID(id *int) *ProductCategoryContentTypeCreate {
	if id != nil {
		pcctc = pcctc.SetParentID(*id)
	}
	return pcctc
}

// SetParent sets the "parent" edge to the ProductCategoryContentType entity.
func (pcctc *ProductCategoryContentTypeCreate) SetParent(p *ProductCategoryContentType) *ProductCategoryContentTypeCreate {
	return pcctc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductCategoryContentType entity by IDs.
func (pcctc *ProductCategoryContentTypeCreate) AddChildIDs(ids ...int) *ProductCategoryContentTypeCreate {
	pcctc.mutation.AddChildIDs(ids...)
	return pcctc
}

// AddChildren adds the "children" edges to the ProductCategoryContentType entity.
func (pcctc *ProductCategoryContentTypeCreate) AddChildren(p ...*ProductCategoryContentType) *ProductCategoryContentTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcctc.AddChildIDs(ids...)
}

// AddChildProductCategoryContentTypeIDs adds the "child_product_category_content_types" edge to the ProductCategoryContentType entity by IDs.
func (pcctc *ProductCategoryContentTypeCreate) AddChildProductCategoryContentTypeIDs(ids ...int) *ProductCategoryContentTypeCreate {
	pcctc.mutation.AddChildProductCategoryContentTypeIDs(ids...)
	return pcctc
}

// AddChildProductCategoryContentTypes adds the "child_product_category_content_types" edges to the ProductCategoryContentType entity.
func (pcctc *ProductCategoryContentTypeCreate) AddChildProductCategoryContentTypes(p ...*ProductCategoryContentType) *ProductCategoryContentTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcctc.AddChildProductCategoryContentTypeIDs(ids...)
}

// Mutation returns the ProductCategoryContentTypeMutation object of the builder.
func (pcctc *ProductCategoryContentTypeCreate) Mutation() *ProductCategoryContentTypeMutation {
	return pcctc.mutation
}

// Save creates the ProductCategoryContentType in the database.
func (pcctc *ProductCategoryContentTypeCreate) Save(ctx context.Context) (*ProductCategoryContentType, error) {
	var (
		err  error
		node *ProductCategoryContentType
	)
	pcctc.defaults()
	if len(pcctc.hooks) == 0 {
		if err = pcctc.check(); err != nil {
			return nil, err
		}
		node, err = pcctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryContentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcctc.check(); err != nil {
				return nil, err
			}
			pcctc.mutation = mutation
			if node, err = pcctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcctc.hooks) - 1; i >= 0; i-- {
			mut = pcctc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcctc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcctc *ProductCategoryContentTypeCreate) SaveX(ctx context.Context) *ProductCategoryContentType {
	v, err := pcctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pcctc *ProductCategoryContentTypeCreate) defaults() {
	if _, ok := pcctc.mutation.CreateTime(); !ok {
		v := productcategorycontenttype.DefaultCreateTime()
		pcctc.mutation.SetCreateTime(v)
	}
	if _, ok := pcctc.mutation.UpdateTime(); !ok {
		v := productcategorycontenttype.DefaultUpdateTime()
		pcctc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcctc *ProductCategoryContentTypeCreate) check() error {
	if _, ok := pcctc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pcctc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pcctc.mutation.HasTable(); ok {
		if err := productcategorycontenttype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pcctc *ProductCategoryContentTypeCreate) sqlSave(ctx context.Context) (*ProductCategoryContentType, error) {
	_node, _spec := pcctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcctc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pcctc *ProductCategoryContentTypeCreate) createSpec() (*ProductCategoryContentType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductCategoryContentType{config: pcctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: productcategorycontenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productcategorycontenttype.FieldID,
			},
		}
	)
	if value, ok := pcctc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productcategorycontenttype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pcctc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productcategorycontenttype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pcctc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productcategorycontenttype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pcctc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: productcategorycontenttype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pcctc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productcategorycontenttype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pcctc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategorycontenttype.ParentTable,
			Columns: []string{productcategorycontenttype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorycontenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_category_content_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcctc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategorycontenttype.ChildrenTable,
			Columns: []string{productcategorycontenttype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorycontenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcctc.mutation.ChildProductCategoryContentTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   productcategorycontenttype.ChildProductCategoryContentTypesTable,
			Columns: productcategorycontenttype.ChildProductCategoryContentTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productcategorycontenttype.FieldID,
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

// ProductCategoryContentTypeCreateBulk is the builder for creating many ProductCategoryContentType entities in bulk.
type ProductCategoryContentTypeCreateBulk struct {
	config
	builders []*ProductCategoryContentTypeCreate
}

// Save creates the ProductCategoryContentType entities in the database.
func (pcctcb *ProductCategoryContentTypeCreateBulk) Save(ctx context.Context) ([]*ProductCategoryContentType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcctcb.builders))
	nodes := make([]*ProductCategoryContentType, len(pcctcb.builders))
	mutators := make([]Mutator, len(pcctcb.builders))
	for i := range pcctcb.builders {
		func(i int, root context.Context) {
			builder := pcctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductCategoryContentTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pcctcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcctcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcctcb *ProductCategoryContentTypeCreateBulk) SaveX(ctx context.Context) []*ProductCategoryContentType {
	v, err := pcctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
