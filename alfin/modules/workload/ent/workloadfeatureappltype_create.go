// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappl"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappltype"
)

// WorkloadFeatureApplTypeCreate is the builder for creating a WorkloadFeatureApplType entity.
type WorkloadFeatureApplTypeCreate struct {
	config
	mutation *WorkloadFeatureApplTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (wfatc *WorkloadFeatureApplTypeCreate) SetCreateTime(t time.Time) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.SetCreateTime(t)
	return wfatc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wfatc *WorkloadFeatureApplTypeCreate) SetNillableCreateTime(t *time.Time) *WorkloadFeatureApplTypeCreate {
	if t != nil {
		wfatc.SetCreateTime(*t)
	}
	return wfatc
}

// SetUpdateTime sets the "update_time" field.
func (wfatc *WorkloadFeatureApplTypeCreate) SetUpdateTime(t time.Time) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.SetUpdateTime(t)
	return wfatc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (wfatc *WorkloadFeatureApplTypeCreate) SetNillableUpdateTime(t *time.Time) *WorkloadFeatureApplTypeCreate {
	if t != nil {
		wfatc.SetUpdateTime(*t)
	}
	return wfatc
}

// SetDescription sets the "description" field.
func (wfatc *WorkloadFeatureApplTypeCreate) SetDescription(s string) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.SetDescription(s)
	return wfatc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wfatc *WorkloadFeatureApplTypeCreate) SetNillableDescription(s *string) *WorkloadFeatureApplTypeCreate {
	if s != nil {
		wfatc.SetDescription(*s)
	}
	return wfatc
}

// SetParentID sets the "parent" edge to the WorkloadFeatureApplType entity by ID.
func (wfatc *WorkloadFeatureApplTypeCreate) SetParentID(id int) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.SetParentID(id)
	return wfatc
}

// SetNillableParentID sets the "parent" edge to the WorkloadFeatureApplType entity by ID if the given value is not nil.
func (wfatc *WorkloadFeatureApplTypeCreate) SetNillableParentID(id *int) *WorkloadFeatureApplTypeCreate {
	if id != nil {
		wfatc = wfatc.SetParentID(*id)
	}
	return wfatc
}

// SetParent sets the "parent" edge to the WorkloadFeatureApplType entity.
func (wfatc *WorkloadFeatureApplTypeCreate) SetParent(w *WorkloadFeatureApplType) *WorkloadFeatureApplTypeCreate {
	return wfatc.SetParentID(w.ID)
}

// AddChildIDs adds the "children" edge to the WorkloadFeatureApplType entity by IDs.
func (wfatc *WorkloadFeatureApplTypeCreate) AddChildIDs(ids ...int) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.AddChildIDs(ids...)
	return wfatc
}

// AddChildren adds the "children" edges to the WorkloadFeatureApplType entity.
func (wfatc *WorkloadFeatureApplTypeCreate) AddChildren(w ...*WorkloadFeatureApplType) *WorkloadFeatureApplTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wfatc.AddChildIDs(ids...)
}

// AddWorkloadFeatureApplIDs adds the "workload_feature_appls" edge to the WorkloadFeatureAppl entity by IDs.
func (wfatc *WorkloadFeatureApplTypeCreate) AddWorkloadFeatureApplIDs(ids ...int) *WorkloadFeatureApplTypeCreate {
	wfatc.mutation.AddWorkloadFeatureApplIDs(ids...)
	return wfatc
}

// AddWorkloadFeatureAppls adds the "workload_feature_appls" edges to the WorkloadFeatureAppl entity.
func (wfatc *WorkloadFeatureApplTypeCreate) AddWorkloadFeatureAppls(w ...*WorkloadFeatureAppl) *WorkloadFeatureApplTypeCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wfatc.AddWorkloadFeatureApplIDs(ids...)
}

// Mutation returns the WorkloadFeatureApplTypeMutation object of the builder.
func (wfatc *WorkloadFeatureApplTypeCreate) Mutation() *WorkloadFeatureApplTypeMutation {
	return wfatc.mutation
}

// Save creates the WorkloadFeatureApplType in the database.
func (wfatc *WorkloadFeatureApplTypeCreate) Save(ctx context.Context) (*WorkloadFeatureApplType, error) {
	var (
		err  error
		node *WorkloadFeatureApplType
	)
	wfatc.defaults()
	if len(wfatc.hooks) == 0 {
		if err = wfatc.check(); err != nil {
			return nil, err
		}
		node, err = wfatc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadFeatureApplTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wfatc.check(); err != nil {
				return nil, err
			}
			wfatc.mutation = mutation
			node, err = wfatc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wfatc.hooks) - 1; i >= 0; i-- {
			mut = wfatc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfatc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wfatc *WorkloadFeatureApplTypeCreate) SaveX(ctx context.Context) *WorkloadFeatureApplType {
	v, err := wfatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wfatc *WorkloadFeatureApplTypeCreate) defaults() {
	if _, ok := wfatc.mutation.CreateTime(); !ok {
		v := workloadfeatureappltype.DefaultCreateTime()
		wfatc.mutation.SetCreateTime(v)
	}
	if _, ok := wfatc.mutation.UpdateTime(); !ok {
		v := workloadfeatureappltype.DefaultUpdateTime()
		wfatc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wfatc *WorkloadFeatureApplTypeCreate) check() error {
	if _, ok := wfatc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := wfatc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (wfatc *WorkloadFeatureApplTypeCreate) sqlSave(ctx context.Context) (*WorkloadFeatureApplType, error) {
	_node, _spec := wfatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wfatc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wfatc *WorkloadFeatureApplTypeCreate) createSpec() (*WorkloadFeatureApplType, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkloadFeatureApplType{config: wfatc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workloadfeatureappltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadfeatureappltype.FieldID,
			},
		}
	)
	if value, ok := wfatc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadfeatureappltype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := wfatc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadfeatureappltype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := wfatc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workloadfeatureappltype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := wfatc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadfeatureappltype.ParentTable,
			Columns: []string{workloadfeatureappltype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workloadfeatureappltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workload_feature_appl_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wfatc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workloadfeatureappltype.ChildrenTable,
			Columns: []string{workloadfeatureappltype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workloadfeatureappltype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wfatc.mutation.WorkloadFeatureApplsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workloadfeatureappltype.WorkloadFeatureApplsTable,
			Columns: []string{workloadfeatureappltype.WorkloadFeatureApplsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workloadfeatureappl.FieldID,
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

// WorkloadFeatureApplTypeCreateBulk is the builder for creating many WorkloadFeatureApplType entities in bulk.
type WorkloadFeatureApplTypeCreateBulk struct {
	config
	builders []*WorkloadFeatureApplTypeCreate
}

// Save creates the WorkloadFeatureApplType entities in the database.
func (wfatcb *WorkloadFeatureApplTypeCreateBulk) Save(ctx context.Context) ([]*WorkloadFeatureApplType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wfatcb.builders))
	nodes := make([]*WorkloadFeatureApplType, len(wfatcb.builders))
	mutators := make([]Mutator, len(wfatcb.builders))
	for i := range wfatcb.builders {
		func(i int, root context.Context) {
			builder := wfatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkloadFeatureApplTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, wfatcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wfatcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wfatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wfatcb *WorkloadFeatureApplTypeCreateBulk) SaveX(ctx context.Context) []*WorkloadFeatureApplType {
	v, err := wfatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
