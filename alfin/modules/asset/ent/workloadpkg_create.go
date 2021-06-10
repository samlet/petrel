// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/asset/ent/asset"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"
)

// WorkloadPkgCreate is the builder for creating a WorkloadPkg entity.
type WorkloadPkgCreate struct {
	config
	mutation *WorkloadPkgMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wpc *WorkloadPkgCreate) SetName(s string) *WorkloadPkgCreate {
	wpc.mutation.SetName(s)
	return wpc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wpc *WorkloadPkgCreate) SetNillableName(s *string) *WorkloadPkgCreate {
	if s != nil {
		wpc.SetName(*s)
	}
	return wpc
}

// AddAssetIDs adds the "assets" edge to the Asset entity by IDs.
func (wpc *WorkloadPkgCreate) AddAssetIDs(ids ...int) *WorkloadPkgCreate {
	wpc.mutation.AddAssetIDs(ids...)
	return wpc
}

// AddAssets adds the "assets" edges to the Asset entity.
func (wpc *WorkloadPkgCreate) AddAssets(a ...*Asset) *WorkloadPkgCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wpc.AddAssetIDs(ids...)
}

// Mutation returns the WorkloadPkgMutation object of the builder.
func (wpc *WorkloadPkgCreate) Mutation() *WorkloadPkgMutation {
	return wpc.mutation
}

// Save creates the WorkloadPkg in the database.
func (wpc *WorkloadPkgCreate) Save(ctx context.Context) (*WorkloadPkg, error) {
	var (
		err  error
		node *WorkloadPkg
	)
	wpc.defaults()
	if len(wpc.hooks) == 0 {
		if err = wpc.check(); err != nil {
			return nil, err
		}
		node, err = wpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadPkgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wpc.check(); err != nil {
				return nil, err
			}
			wpc.mutation = mutation
			node, err = wpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wpc.hooks) - 1; i >= 0; i-- {
			mut = wpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wpc *WorkloadPkgCreate) SaveX(ctx context.Context) *WorkloadPkg {
	v, err := wpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wpc *WorkloadPkgCreate) defaults() {
	if _, ok := wpc.mutation.Name(); !ok {
		v := workloadpkg.DefaultName
		wpc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wpc *WorkloadPkgCreate) check() error {
	if _, ok := wpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (wpc *WorkloadPkgCreate) sqlSave(ctx context.Context) (*WorkloadPkg, error) {
	_node, _spec := wpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wpc *WorkloadPkgCreate) createSpec() (*WorkloadPkg, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkloadPkg{config: wpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workloadpkg.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadpkg.FieldID,
			},
		}
	)
	if value, ok := wpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workloadpkg.FieldName,
		})
		_node.Name = value
	}
	if nodes := wpc.mutation.AssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workloadpkg.AssetsTable,
			Columns: []string{workloadpkg.AssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: asset.FieldID,
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

// WorkloadPkgCreateBulk is the builder for creating many WorkloadPkg entities in bulk.
type WorkloadPkgCreateBulk struct {
	config
	builders []*WorkloadPkgCreate
}

// Save creates the WorkloadPkg entities in the database.
func (wpcb *WorkloadPkgCreateBulk) Save(ctx context.Context) ([]*WorkloadPkg, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wpcb.builders))
	nodes := make([]*WorkloadPkg, len(wpcb.builders))
	mutators := make([]Mutator, len(wpcb.builders))
	for i := range wpcb.builders {
		func(i int, root context.Context) {
			builder := wpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkloadPkgMutation)
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
					_, err = mutators[i+1].Mutate(root, wpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wpcb *WorkloadPkgCreateBulk) SaveX(ctx context.Context) []*WorkloadPkg {
	v, err := wpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}