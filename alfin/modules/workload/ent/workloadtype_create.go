// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadtype"
)

// WorkloadTypeCreate is the builder for creating a WorkloadType entity.
type WorkloadTypeCreate struct {
	config
	mutation *WorkloadTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (wtc *WorkloadTypeCreate) SetCreateTime(t time.Time) *WorkloadTypeCreate {
	wtc.mutation.SetCreateTime(t)
	return wtc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wtc *WorkloadTypeCreate) SetNillableCreateTime(t *time.Time) *WorkloadTypeCreate {
	if t != nil {
		wtc.SetCreateTime(*t)
	}
	return wtc
}

// SetUpdateTime sets the "update_time" field.
func (wtc *WorkloadTypeCreate) SetUpdateTime(t time.Time) *WorkloadTypeCreate {
	wtc.mutation.SetUpdateTime(t)
	return wtc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (wtc *WorkloadTypeCreate) SetNillableUpdateTime(t *time.Time) *WorkloadTypeCreate {
	if t != nil {
		wtc.SetUpdateTime(*t)
	}
	return wtc
}

// SetDescription sets the "description" field.
func (wtc *WorkloadTypeCreate) SetDescription(s string) *WorkloadTypeCreate {
	wtc.mutation.SetDescription(s)
	return wtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wtc *WorkloadTypeCreate) SetNillableDescription(s *string) *WorkloadTypeCreate {
	if s != nil {
		wtc.SetDescription(*s)
	}
	return wtc
}

// SetParentID sets the "parent" edge to the WorkloadType entity by ID.
func (wtc *WorkloadTypeCreate) SetParentID(id int) *WorkloadTypeCreate {
	wtc.mutation.SetParentID(id)
	return wtc
}

// SetNillableParentID sets the "parent" edge to the WorkloadType entity by ID if the given value is not nil.
func (wtc *WorkloadTypeCreate) SetNillableParentID(id *int) *WorkloadTypeCreate {
	if id != nil {
		wtc = wtc.SetParentID(*id)
	}
	return wtc
}

// SetParent sets the "parent" edge to the WorkloadType entity.
func (wtc *WorkloadTypeCreate) SetParent(w *WorkloadType) *WorkloadTypeCreate {
	return wtc.SetParentID(w.ID)
}

// AddChildIDs adds the "children" edge to the WorkloadType entity by IDs.
func (wtc *WorkloadTypeCreate) AddChildIDs(ids ...int) *WorkloadTypeCreate {
	wtc.mutation.AddChildIDs(ids...)
	return wtc
}

// AddChildren adds the "children" edges to the WorkloadType entity.
func (wtc *WorkloadTypeCreate) AddChildren(w ...*WorkloadType) *WorkloadTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtc.AddChildIDs(ids...)
}

// AddWorkloadIDs adds the "workloads" edge to the Workload entity by IDs.
func (wtc *WorkloadTypeCreate) AddWorkloadIDs(ids ...int) *WorkloadTypeCreate {
	wtc.mutation.AddWorkloadIDs(ids...)
	return wtc
}

// AddWorkloads adds the "workloads" edges to the Workload entity.
func (wtc *WorkloadTypeCreate) AddWorkloads(w ...*Workload) *WorkloadTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wtc.AddWorkloadIDs(ids...)
}

// Mutation returns the WorkloadTypeMutation object of the builder.
func (wtc *WorkloadTypeCreate) Mutation() *WorkloadTypeMutation {
	return wtc.mutation
}

// Save creates the WorkloadType in the database.
func (wtc *WorkloadTypeCreate) Save(ctx context.Context) (*WorkloadType, error) {
	var (
		err  error
		node *WorkloadType
	)
	wtc.defaults()
	if len(wtc.hooks) == 0 {
		if err = wtc.check(); err != nil {
			return nil, err
		}
		node, err = wtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wtc.check(); err != nil {
				return nil, err
			}
			wtc.mutation = mutation
			node, err = wtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wtc.hooks) - 1; i >= 0; i-- {
			mut = wtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wtc *WorkloadTypeCreate) SaveX(ctx context.Context) *WorkloadType {
	v, err := wtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wtc *WorkloadTypeCreate) defaults() {
	if _, ok := wtc.mutation.CreateTime(); !ok {
		v := workloadtype.DefaultCreateTime()
		wtc.mutation.SetCreateTime(v)
	}
	if _, ok := wtc.mutation.UpdateTime(); !ok {
		v := workloadtype.DefaultUpdateTime()
		wtc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wtc *WorkloadTypeCreate) check() error {
	if _, ok := wtc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := wtc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (wtc *WorkloadTypeCreate) sqlSave(ctx context.Context) (*WorkloadType, error) {
	_node, _spec := wtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wtc *WorkloadTypeCreate) createSpec() (*WorkloadType, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkloadType{config: wtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workloadtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadtype.FieldID,
			},
		}
	)
	if value, ok := wtc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadtype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := wtc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadtype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := wtc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workloadtype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := wtc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadtype.ParentTable,
			Columns: []string{workloadtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workloadtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workload_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wtc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workloadtype.ChildrenTable,
			Columns: []string{workloadtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workloadtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wtc.mutation.WorkloadsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workloadtype.WorkloadsTable,
			Columns: []string{workloadtype.WorkloadsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workload.FieldID,
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

// WorkloadTypeCreateBulk is the builder for creating many WorkloadType entities in bulk.
type WorkloadTypeCreateBulk struct {
	config
	builders []*WorkloadTypeCreate
}

// Save creates the WorkloadType entities in the database.
func (wtcb *WorkloadTypeCreateBulk) Save(ctx context.Context) ([]*WorkloadType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wtcb.builders))
	nodes := make([]*WorkloadType, len(wtcb.builders))
	mutators := make([]Mutator, len(wtcb.builders))
	for i := range wtcb.builders {
		func(i int, root context.Context) {
			builder := wtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkloadTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, wtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wtcb *WorkloadTypeCreateBulk) SaveX(ctx context.Context) []*WorkloadType {
	v, err := wtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
