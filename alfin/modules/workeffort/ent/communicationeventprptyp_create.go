// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventprptyp"
)

// CommunicationEventPrpTypCreate is the builder for creating a CommunicationEventPrpTyp entity.
type CommunicationEventPrpTypCreate struct {
	config
	mutation *CommunicationEventPrpTypMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ceptc *CommunicationEventPrpTypCreate) SetCreateTime(t time.Time) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetCreateTime(t)
	return ceptc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableCreateTime(t *time.Time) *CommunicationEventPrpTypCreate {
	if t != nil {
		ceptc.SetCreateTime(*t)
	}
	return ceptc
}

// SetUpdateTime sets the "update_time" field.
func (ceptc *CommunicationEventPrpTypCreate) SetUpdateTime(t time.Time) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetUpdateTime(t)
	return ceptc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableUpdateTime(t *time.Time) *CommunicationEventPrpTypCreate {
	if t != nil {
		ceptc.SetUpdateTime(*t)
	}
	return ceptc
}

// SetStringRef sets the "string_ref" field.
func (ceptc *CommunicationEventPrpTypCreate) SetStringRef(s string) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetStringRef(s)
	return ceptc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableStringRef(s *string) *CommunicationEventPrpTypCreate {
	if s != nil {
		ceptc.SetStringRef(*s)
	}
	return ceptc
}

// SetHasTable sets the "has_table" field.
func (ceptc *CommunicationEventPrpTypCreate) SetHasTable(ct communicationeventprptyp.HasTable) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetHasTable(ct)
	return ceptc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableHasTable(ct *communicationeventprptyp.HasTable) *CommunicationEventPrpTypCreate {
	if ct != nil {
		ceptc.SetHasTable(*ct)
	}
	return ceptc
}

// SetDescription sets the "description" field.
func (ceptc *CommunicationEventPrpTypCreate) SetDescription(s string) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetDescription(s)
	return ceptc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableDescription(s *string) *CommunicationEventPrpTypCreate {
	if s != nil {
		ceptc.SetDescription(*s)
	}
	return ceptc
}

// SetParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID.
func (ceptc *CommunicationEventPrpTypCreate) SetParentID(id int) *CommunicationEventPrpTypCreate {
	ceptc.mutation.SetParentID(id)
	return ceptc
}

// SetNillableParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID if the given value is not nil.
func (ceptc *CommunicationEventPrpTypCreate) SetNillableParentID(id *int) *CommunicationEventPrpTypCreate {
	if id != nil {
		ceptc = ceptc.SetParentID(*id)
	}
	return ceptc
}

// SetParent sets the "parent" edge to the CommunicationEventPrpTyp entity.
func (ceptc *CommunicationEventPrpTypCreate) SetParent(c *CommunicationEventPrpTyp) *CommunicationEventPrpTypCreate {
	return ceptc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptc *CommunicationEventPrpTypCreate) AddChildIDs(ids ...int) *CommunicationEventPrpTypCreate {
	ceptc.mutation.AddChildIDs(ids...)
	return ceptc
}

// AddChildren adds the "children" edges to the CommunicationEventPrpTyp entity.
func (ceptc *CommunicationEventPrpTypCreate) AddChildren(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptc.AddChildIDs(ids...)
}

// AddChildCommunicationEventPrpTypIDs adds the "child_communication_event_prp_typs" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptc *CommunicationEventPrpTypCreate) AddChildCommunicationEventPrpTypIDs(ids ...int) *CommunicationEventPrpTypCreate {
	ceptc.mutation.AddChildCommunicationEventPrpTypIDs(ids...)
	return ceptc
}

// AddChildCommunicationEventPrpTyps adds the "child_communication_event_prp_typs" edges to the CommunicationEventPrpTyp entity.
func (ceptc *CommunicationEventPrpTypCreate) AddChildCommunicationEventPrpTyps(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptc.AddChildCommunicationEventPrpTypIDs(ids...)
}

// Mutation returns the CommunicationEventPrpTypMutation object of the builder.
func (ceptc *CommunicationEventPrpTypCreate) Mutation() *CommunicationEventPrpTypMutation {
	return ceptc.mutation
}

// Save creates the CommunicationEventPrpTyp in the database.
func (ceptc *CommunicationEventPrpTypCreate) Save(ctx context.Context) (*CommunicationEventPrpTyp, error) {
	var (
		err  error
		node *CommunicationEventPrpTyp
	)
	ceptc.defaults()
	if len(ceptc.hooks) == 0 {
		if err = ceptc.check(); err != nil {
			return nil, err
		}
		node, err = ceptc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunicationEventPrpTypMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ceptc.check(); err != nil {
				return nil, err
			}
			ceptc.mutation = mutation
			if node, err = ceptc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ceptc.hooks) - 1; i >= 0; i-- {
			mut = ceptc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceptc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ceptc *CommunicationEventPrpTypCreate) SaveX(ctx context.Context) *CommunicationEventPrpTyp {
	v, err := ceptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ceptc *CommunicationEventPrpTypCreate) defaults() {
	if _, ok := ceptc.mutation.CreateTime(); !ok {
		v := communicationeventprptyp.DefaultCreateTime()
		ceptc.mutation.SetCreateTime(v)
	}
	if _, ok := ceptc.mutation.UpdateTime(); !ok {
		v := communicationeventprptyp.DefaultUpdateTime()
		ceptc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ceptc *CommunicationEventPrpTypCreate) check() error {
	if _, ok := ceptc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ceptc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := ceptc.mutation.HasTable(); ok {
		if err := communicationeventprptyp.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (ceptc *CommunicationEventPrpTypCreate) sqlSave(ctx context.Context) (*CommunicationEventPrpTyp, error) {
	_node, _spec := ceptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ceptc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ceptc *CommunicationEventPrpTypCreate) createSpec() (*CommunicationEventPrpTyp, *sqlgraph.CreateSpec) {
	var (
		_node = &CommunicationEventPrpTyp{config: ceptc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: communicationeventprptyp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationeventprptyp.FieldID,
			},
		}
	)
	if value, ok := ceptc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: communicationeventprptyp.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ceptc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: communicationeventprptyp.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ceptc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := ceptc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: communicationeventprptyp.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := ceptc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := ceptc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   communicationeventprptyp.ParentTable,
			Columns: []string{communicationeventprptyp.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: communicationeventprptyp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.communication_event_prp_typ_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ceptc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   communicationeventprptyp.ChildrenTable,
			Columns: []string{communicationeventprptyp.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: communicationeventprptyp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ceptc.mutation.ChildCommunicationEventPrpTypsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   communicationeventprptyp.ChildCommunicationEventPrpTypsTable,
			Columns: communicationeventprptyp.ChildCommunicationEventPrpTypsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: communicationeventprptyp.FieldID,
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

// CommunicationEventPrpTypCreateBulk is the builder for creating many CommunicationEventPrpTyp entities in bulk.
type CommunicationEventPrpTypCreateBulk struct {
	config
	builders []*CommunicationEventPrpTypCreate
}

// Save creates the CommunicationEventPrpTyp entities in the database.
func (ceptcb *CommunicationEventPrpTypCreateBulk) Save(ctx context.Context) ([]*CommunicationEventPrpTyp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ceptcb.builders))
	nodes := make([]*CommunicationEventPrpTyp, len(ceptcb.builders))
	mutators := make([]Mutator, len(ceptcb.builders))
	for i := range ceptcb.builders {
		func(i int, root context.Context) {
			builder := ceptcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommunicationEventPrpTypMutation)
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
					_, err = mutators[i+1].Mutate(root, ceptcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ceptcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ceptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ceptcb *CommunicationEventPrpTypCreateBulk) SaveX(ctx context.Context) []*CommunicationEventPrpTyp {
	v, err := ceptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
