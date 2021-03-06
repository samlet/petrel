// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/prodconfitemcontenttype"
)

// ProdConfItemContentTypeCreate is the builder for creating a ProdConfItemContentType entity.
type ProdConfItemContentTypeCreate struct {
	config
	mutation *ProdConfItemContentTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pcictc *ProdConfItemContentTypeCreate) SetCreateTime(t time.Time) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetCreateTime(t)
	return pcictc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableCreateTime(t *time.Time) *ProdConfItemContentTypeCreate {
	if t != nil {
		pcictc.SetCreateTime(*t)
	}
	return pcictc
}

// SetUpdateTime sets the "update_time" field.
func (pcictc *ProdConfItemContentTypeCreate) SetUpdateTime(t time.Time) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetUpdateTime(t)
	return pcictc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableUpdateTime(t *time.Time) *ProdConfItemContentTypeCreate {
	if t != nil {
		pcictc.SetUpdateTime(*t)
	}
	return pcictc
}

// SetStringRef sets the "string_ref" field.
func (pcictc *ProdConfItemContentTypeCreate) SetStringRef(s string) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetStringRef(s)
	return pcictc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableStringRef(s *string) *ProdConfItemContentTypeCreate {
	if s != nil {
		pcictc.SetStringRef(*s)
	}
	return pcictc
}

// SetHasTable sets the "has_table" field.
func (pcictc *ProdConfItemContentTypeCreate) SetHasTable(pt prodconfitemcontenttype.HasTable) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetHasTable(pt)
	return pcictc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableHasTable(pt *prodconfitemcontenttype.HasTable) *ProdConfItemContentTypeCreate {
	if pt != nil {
		pcictc.SetHasTable(*pt)
	}
	return pcictc
}

// SetDescription sets the "description" field.
func (pcictc *ProdConfItemContentTypeCreate) SetDescription(s string) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetDescription(s)
	return pcictc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableDescription(s *string) *ProdConfItemContentTypeCreate {
	if s != nil {
		pcictc.SetDescription(*s)
	}
	return pcictc
}

// SetParentID sets the "parent" edge to the ProdConfItemContentType entity by ID.
func (pcictc *ProdConfItemContentTypeCreate) SetParentID(id int) *ProdConfItemContentTypeCreate {
	pcictc.mutation.SetParentID(id)
	return pcictc
}

// SetNillableParentID sets the "parent" edge to the ProdConfItemContentType entity by ID if the given value is not nil.
func (pcictc *ProdConfItemContentTypeCreate) SetNillableParentID(id *int) *ProdConfItemContentTypeCreate {
	if id != nil {
		pcictc = pcictc.SetParentID(*id)
	}
	return pcictc
}

// SetParent sets the "parent" edge to the ProdConfItemContentType entity.
func (pcictc *ProdConfItemContentTypeCreate) SetParent(p *ProdConfItemContentType) *ProdConfItemContentTypeCreate {
	return pcictc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProdConfItemContentType entity by IDs.
func (pcictc *ProdConfItemContentTypeCreate) AddChildIDs(ids ...int) *ProdConfItemContentTypeCreate {
	pcictc.mutation.AddChildIDs(ids...)
	return pcictc
}

// AddChildren adds the "children" edges to the ProdConfItemContentType entity.
func (pcictc *ProdConfItemContentTypeCreate) AddChildren(p ...*ProdConfItemContentType) *ProdConfItemContentTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcictc.AddChildIDs(ids...)
}

// AddChildProdConfItemContentTypeIDs adds the "child_prod_conf_item_content_types" edge to the ProdConfItemContentType entity by IDs.
func (pcictc *ProdConfItemContentTypeCreate) AddChildProdConfItemContentTypeIDs(ids ...int) *ProdConfItemContentTypeCreate {
	pcictc.mutation.AddChildProdConfItemContentTypeIDs(ids...)
	return pcictc
}

// AddChildProdConfItemContentTypes adds the "child_prod_conf_item_content_types" edges to the ProdConfItemContentType entity.
func (pcictc *ProdConfItemContentTypeCreate) AddChildProdConfItemContentTypes(p ...*ProdConfItemContentType) *ProdConfItemContentTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcictc.AddChildProdConfItemContentTypeIDs(ids...)
}

// Mutation returns the ProdConfItemContentTypeMutation object of the builder.
func (pcictc *ProdConfItemContentTypeCreate) Mutation() *ProdConfItemContentTypeMutation {
	return pcictc.mutation
}

// Save creates the ProdConfItemContentType in the database.
func (pcictc *ProdConfItemContentTypeCreate) Save(ctx context.Context) (*ProdConfItemContentType, error) {
	var (
		err  error
		node *ProdConfItemContentType
	)
	pcictc.defaults()
	if len(pcictc.hooks) == 0 {
		if err = pcictc.check(); err != nil {
			return nil, err
		}
		node, err = pcictc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProdConfItemContentTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcictc.check(); err != nil {
				return nil, err
			}
			pcictc.mutation = mutation
			if node, err = pcictc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcictc.hooks) - 1; i >= 0; i-- {
			mut = pcictc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcictc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcictc *ProdConfItemContentTypeCreate) SaveX(ctx context.Context) *ProdConfItemContentType {
	v, err := pcictc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pcictc *ProdConfItemContentTypeCreate) defaults() {
	if _, ok := pcictc.mutation.CreateTime(); !ok {
		v := prodconfitemcontenttype.DefaultCreateTime()
		pcictc.mutation.SetCreateTime(v)
	}
	if _, ok := pcictc.mutation.UpdateTime(); !ok {
		v := prodconfitemcontenttype.DefaultUpdateTime()
		pcictc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcictc *ProdConfItemContentTypeCreate) check() error {
	if _, ok := pcictc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pcictc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pcictc.mutation.HasTable(); ok {
		if err := prodconfitemcontenttype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pcictc *ProdConfItemContentTypeCreate) sqlSave(ctx context.Context) (*ProdConfItemContentType, error) {
	_node, _spec := pcictc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcictc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pcictc *ProdConfItemContentTypeCreate) createSpec() (*ProdConfItemContentType, *sqlgraph.CreateSpec) {
	var (
		_node = &ProdConfItemContentType{config: pcictc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prodconfitemcontenttype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prodconfitemcontenttype.FieldID,
			},
		}
	)
	if value, ok := pcictc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prodconfitemcontenttype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pcictc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: prodconfitemcontenttype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pcictc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prodconfitemcontenttype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pcictc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: prodconfitemcontenttype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pcictc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prodconfitemcontenttype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pcictc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prodconfitemcontenttype.ParentTable,
			Columns: []string{prodconfitemcontenttype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prodconfitemcontenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.prod_conf_item_content_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcictc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prodconfitemcontenttype.ChildrenTable,
			Columns: []string{prodconfitemcontenttype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prodconfitemcontenttype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcictc.mutation.ChildProdConfItemContentTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   prodconfitemcontenttype.ChildProdConfItemContentTypesTable,
			Columns: prodconfitemcontenttype.ChildProdConfItemContentTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prodconfitemcontenttype.FieldID,
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

// ProdConfItemContentTypeCreateBulk is the builder for creating many ProdConfItemContentType entities in bulk.
type ProdConfItemContentTypeCreateBulk struct {
	config
	builders []*ProdConfItemContentTypeCreate
}

// Save creates the ProdConfItemContentType entities in the database.
func (pcictcb *ProdConfItemContentTypeCreateBulk) Save(ctx context.Context) ([]*ProdConfItemContentType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcictcb.builders))
	nodes := make([]*ProdConfItemContentType, len(pcictcb.builders))
	mutators := make([]Mutator, len(pcictcb.builders))
	for i := range pcictcb.builders {
		func(i int, root context.Context) {
			builder := pcictcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProdConfItemContentTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pcictcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcictcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcictcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcictcb *ProdConfItemContentTypeCreateBulk) SaveX(ctx context.Context) []*ProdConfItemContentType {
	v, err := pcictcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
