// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statustype"
)

// StatusTypeCreate is the builder for creating a StatusType entity.
type StatusTypeCreate struct {
	config
	mutation *StatusTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (stc *StatusTypeCreate) SetCreateTime(t time.Time) *StatusTypeCreate {
	stc.mutation.SetCreateTime(t)
	return stc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableCreateTime(t *time.Time) *StatusTypeCreate {
	if t != nil {
		stc.SetCreateTime(*t)
	}
	return stc
}

// SetUpdateTime sets the "update_time" field.
func (stc *StatusTypeCreate) SetUpdateTime(t time.Time) *StatusTypeCreate {
	stc.mutation.SetUpdateTime(t)
	return stc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableUpdateTime(t *time.Time) *StatusTypeCreate {
	if t != nil {
		stc.SetUpdateTime(*t)
	}
	return stc
}

// SetStringRef sets the "string_ref" field.
func (stc *StatusTypeCreate) SetStringRef(s string) *StatusTypeCreate {
	stc.mutation.SetStringRef(s)
	return stc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableStringRef(s *string) *StatusTypeCreate {
	if s != nil {
		stc.SetStringRef(*s)
	}
	return stc
}

// SetHasTable sets the "has_table" field.
func (stc *StatusTypeCreate) SetHasTable(st statustype.HasTable) *StatusTypeCreate {
	stc.mutation.SetHasTable(st)
	return stc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableHasTable(st *statustype.HasTable) *StatusTypeCreate {
	if st != nil {
		stc.SetHasTable(*st)
	}
	return stc
}

// SetDescription sets the "description" field.
func (stc *StatusTypeCreate) SetDescription(s string) *StatusTypeCreate {
	stc.mutation.SetDescription(s)
	return stc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableDescription(s *string) *StatusTypeCreate {
	if s != nil {
		stc.SetDescription(*s)
	}
	return stc
}

// SetParentID sets the "parent" edge to the StatusType entity by ID.
func (stc *StatusTypeCreate) SetParentID(id int) *StatusTypeCreate {
	stc.mutation.SetParentID(id)
	return stc
}

// SetNillableParentID sets the "parent" edge to the StatusType entity by ID if the given value is not nil.
func (stc *StatusTypeCreate) SetNillableParentID(id *int) *StatusTypeCreate {
	if id != nil {
		stc = stc.SetParentID(*id)
	}
	return stc
}

// SetParent sets the "parent" edge to the StatusType entity.
func (stc *StatusTypeCreate) SetParent(s *StatusType) *StatusTypeCreate {
	return stc.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the StatusType entity by IDs.
func (stc *StatusTypeCreate) AddChildIDs(ids ...int) *StatusTypeCreate {
	stc.mutation.AddChildIDs(ids...)
	return stc
}

// AddChildren adds the "children" edges to the StatusType entity.
func (stc *StatusTypeCreate) AddChildren(s ...*StatusType) *StatusTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stc.AddChildIDs(ids...)
}

// AddStatusItemIDs adds the "status_items" edge to the StatusItem entity by IDs.
func (stc *StatusTypeCreate) AddStatusItemIDs(ids ...int) *StatusTypeCreate {
	stc.mutation.AddStatusItemIDs(ids...)
	return stc
}

// AddStatusItems adds the "status_items" edges to the StatusItem entity.
func (stc *StatusTypeCreate) AddStatusItems(s ...*StatusItem) *StatusTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stc.AddStatusItemIDs(ids...)
}

// AddChildStatusTypeIDs adds the "child_status_types" edge to the StatusType entity by IDs.
func (stc *StatusTypeCreate) AddChildStatusTypeIDs(ids ...int) *StatusTypeCreate {
	stc.mutation.AddChildStatusTypeIDs(ids...)
	return stc
}

// AddChildStatusTypes adds the "child_status_types" edges to the StatusType entity.
func (stc *StatusTypeCreate) AddChildStatusTypes(s ...*StatusType) *StatusTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stc.AddChildStatusTypeIDs(ids...)
}

// Mutation returns the StatusTypeMutation object of the builder.
func (stc *StatusTypeCreate) Mutation() *StatusTypeMutation {
	return stc.mutation
}

// Save creates the StatusType in the database.
func (stc *StatusTypeCreate) Save(ctx context.Context) (*StatusType, error) {
	var (
		err  error
		node *StatusType
	)
	stc.defaults()
	if len(stc.hooks) == 0 {
		if err = stc.check(); err != nil {
			return nil, err
		}
		node, err = stc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StatusTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stc.check(); err != nil {
				return nil, err
			}
			stc.mutation = mutation
			if node, err = stc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(stc.hooks) - 1; i >= 0; i-- {
			mut = stc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (stc *StatusTypeCreate) SaveX(ctx context.Context) *StatusType {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (stc *StatusTypeCreate) defaults() {
	if _, ok := stc.mutation.CreateTime(); !ok {
		v := statustype.DefaultCreateTime()
		stc.mutation.SetCreateTime(v)
	}
	if _, ok := stc.mutation.UpdateTime(); !ok {
		v := statustype.DefaultUpdateTime()
		stc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *StatusTypeCreate) check() error {
	if _, ok := stc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := stc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := stc.mutation.HasTable(); ok {
		if err := statustype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (stc *StatusTypeCreate) sqlSave(ctx context.Context) (*StatusType, error) {
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (stc *StatusTypeCreate) createSpec() (*StatusType, *sqlgraph.CreateSpec) {
	var (
		_node = &StatusType{config: stc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: statustype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statustype.FieldID,
			},
		}
	)
	if value, ok := stc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: statustype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := stc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: statustype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := stc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statustype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := stc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: statustype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := stc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: statustype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := stc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statustype.ParentTable,
			Columns: []string{statustype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statustype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.status_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustype.ChildrenTable,
			Columns: []string{statustype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statustype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.StatusItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   statustype.StatusItemsTable,
			Columns: []string{statustype.StatusItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.ChildStatusTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   statustype.ChildStatusTypesTable,
			Columns: statustype.ChildStatusTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statustype.FieldID,
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

// StatusTypeCreateBulk is the builder for creating many StatusType entities in bulk.
type StatusTypeCreateBulk struct {
	config
	builders []*StatusTypeCreate
}

// Save creates the StatusType entities in the database.
func (stcb *StatusTypeCreateBulk) Save(ctx context.Context) ([]*StatusType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*StatusType, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StatusTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *StatusTypeCreateBulk) SaveX(ctx context.Context) []*StatusType {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
