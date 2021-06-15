// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/subscriptiontype"
)

// SubscriptionTypeCreate is the builder for creating a SubscriptionType entity.
type SubscriptionTypeCreate struct {
	config
	mutation *SubscriptionTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (stc *SubscriptionTypeCreate) SetCreateTime(t time.Time) *SubscriptionTypeCreate {
	stc.mutation.SetCreateTime(t)
	return stc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableCreateTime(t *time.Time) *SubscriptionTypeCreate {
	if t != nil {
		stc.SetCreateTime(*t)
	}
	return stc
}

// SetUpdateTime sets the "update_time" field.
func (stc *SubscriptionTypeCreate) SetUpdateTime(t time.Time) *SubscriptionTypeCreate {
	stc.mutation.SetUpdateTime(t)
	return stc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableUpdateTime(t *time.Time) *SubscriptionTypeCreate {
	if t != nil {
		stc.SetUpdateTime(*t)
	}
	return stc
}

// SetStringRef sets the "string_ref" field.
func (stc *SubscriptionTypeCreate) SetStringRef(s string) *SubscriptionTypeCreate {
	stc.mutation.SetStringRef(s)
	return stc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableStringRef(s *string) *SubscriptionTypeCreate {
	if s != nil {
		stc.SetStringRef(*s)
	}
	return stc
}

// SetHasTable sets the "has_table" field.
func (stc *SubscriptionTypeCreate) SetHasTable(st subscriptiontype.HasTable) *SubscriptionTypeCreate {
	stc.mutation.SetHasTable(st)
	return stc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableHasTable(st *subscriptiontype.HasTable) *SubscriptionTypeCreate {
	if st != nil {
		stc.SetHasTable(*st)
	}
	return stc
}

// SetDescription sets the "description" field.
func (stc *SubscriptionTypeCreate) SetDescription(s string) *SubscriptionTypeCreate {
	stc.mutation.SetDescription(s)
	return stc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableDescription(s *string) *SubscriptionTypeCreate {
	if s != nil {
		stc.SetDescription(*s)
	}
	return stc
}

// SetParentID sets the "parent" edge to the SubscriptionType entity by ID.
func (stc *SubscriptionTypeCreate) SetParentID(id int) *SubscriptionTypeCreate {
	stc.mutation.SetParentID(id)
	return stc
}

// SetNillableParentID sets the "parent" edge to the SubscriptionType entity by ID if the given value is not nil.
func (stc *SubscriptionTypeCreate) SetNillableParentID(id *int) *SubscriptionTypeCreate {
	if id != nil {
		stc = stc.SetParentID(*id)
	}
	return stc
}

// SetParent sets the "parent" edge to the SubscriptionType entity.
func (stc *SubscriptionTypeCreate) SetParent(s *SubscriptionType) *SubscriptionTypeCreate {
	return stc.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the SubscriptionType entity by IDs.
func (stc *SubscriptionTypeCreate) AddChildIDs(ids ...int) *SubscriptionTypeCreate {
	stc.mutation.AddChildIDs(ids...)
	return stc
}

// AddChildren adds the "children" edges to the SubscriptionType entity.
func (stc *SubscriptionTypeCreate) AddChildren(s ...*SubscriptionType) *SubscriptionTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stc.AddChildIDs(ids...)
}

// AddChildSubscriptionTypeIDs adds the "child_subscription_types" edge to the SubscriptionType entity by IDs.
func (stc *SubscriptionTypeCreate) AddChildSubscriptionTypeIDs(ids ...int) *SubscriptionTypeCreate {
	stc.mutation.AddChildSubscriptionTypeIDs(ids...)
	return stc
}

// AddChildSubscriptionTypes adds the "child_subscription_types" edges to the SubscriptionType entity.
func (stc *SubscriptionTypeCreate) AddChildSubscriptionTypes(s ...*SubscriptionType) *SubscriptionTypeCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stc.AddChildSubscriptionTypeIDs(ids...)
}

// Mutation returns the SubscriptionTypeMutation object of the builder.
func (stc *SubscriptionTypeCreate) Mutation() *SubscriptionTypeMutation {
	return stc.mutation
}

// Save creates the SubscriptionType in the database.
func (stc *SubscriptionTypeCreate) Save(ctx context.Context) (*SubscriptionType, error) {
	var (
		err  error
		node *SubscriptionType
	)
	stc.defaults()
	if len(stc.hooks) == 0 {
		if err = stc.check(); err != nil {
			return nil, err
		}
		node, err = stc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionTypeMutation)
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
func (stc *SubscriptionTypeCreate) SaveX(ctx context.Context) *SubscriptionType {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (stc *SubscriptionTypeCreate) defaults() {
	if _, ok := stc.mutation.CreateTime(); !ok {
		v := subscriptiontype.DefaultCreateTime()
		stc.mutation.SetCreateTime(v)
	}
	if _, ok := stc.mutation.UpdateTime(); !ok {
		v := subscriptiontype.DefaultUpdateTime()
		stc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *SubscriptionTypeCreate) check() error {
	if _, ok := stc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := stc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := stc.mutation.HasTable(); ok {
		if err := subscriptiontype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (stc *SubscriptionTypeCreate) sqlSave(ctx context.Context) (*SubscriptionType, error) {
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

func (stc *SubscriptionTypeCreate) createSpec() (*SubscriptionType, *sqlgraph.CreateSpec) {
	var (
		_node = &SubscriptionType{config: stc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subscriptiontype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriptiontype.FieldID,
			},
		}
	)
	if value, ok := stc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptiontype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := stc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptiontype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := stc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := stc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: subscriptiontype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := stc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := stc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscriptiontype.ParentTable,
			Columns: []string{subscriptiontype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subscriptiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.subscription_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subscriptiontype.ChildrenTable,
			Columns: []string{subscriptiontype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subscriptiontype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.ChildSubscriptionTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subscriptiontype.ChildSubscriptionTypesTable,
			Columns: subscriptiontype.ChildSubscriptionTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subscriptiontype.FieldID,
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

// SubscriptionTypeCreateBulk is the builder for creating many SubscriptionType entities in bulk.
type SubscriptionTypeCreateBulk struct {
	config
	builders []*SubscriptionTypeCreate
}

// Save creates the SubscriptionType entities in the database.
func (stcb *SubscriptionTypeCreateBulk) Save(ctx context.Context) ([]*SubscriptionType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*SubscriptionType, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriptionTypeMutation)
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
func (stcb *SubscriptionTypeCreateBulk) SaveX(ctx context.Context) []*SubscriptionType {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
