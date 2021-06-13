// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/fixedasset"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortfixedassetassign"
)

// WorkEffortFixedAssetAssignCreate is the builder for creating a WorkEffortFixedAssetAssign entity.
type WorkEffortFixedAssetAssignCreate struct {
	config
	mutation *WorkEffortFixedAssetAssignMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetCreateTime(t time.Time) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetCreateTime(t)
	return wefaac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableCreateTime(t *time.Time) *WorkEffortFixedAssetAssignCreate {
	if t != nil {
		wefaac.SetCreateTime(*t)
	}
	return wefaac
}

// SetUpdateTime sets the "update_time" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetUpdateTime(t time.Time) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetUpdateTime(t)
	return wefaac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableUpdateTime(t *time.Time) *WorkEffortFixedAssetAssignCreate {
	if t != nil {
		wefaac.SetUpdateTime(*t)
	}
	return wefaac
}

// SetStringRef sets the "string_ref" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetStringRef(s string) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetStringRef(s)
	return wefaac
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableStringRef(s *string) *WorkEffortFixedAssetAssignCreate {
	if s != nil {
		wefaac.SetStringRef(*s)
	}
	return wefaac
}

// SetFromDate sets the "from_date" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetFromDate(t time.Time) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetFromDate(t)
	return wefaac
}

// SetNillableFromDate sets the "from_date" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableFromDate(t *time.Time) *WorkEffortFixedAssetAssignCreate {
	if t != nil {
		wefaac.SetFromDate(*t)
	}
	return wefaac
}

// SetThruDate sets the "thru_date" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetThruDate(t time.Time) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetThruDate(t)
	return wefaac
}

// SetNillableThruDate sets the "thru_date" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableThruDate(t *time.Time) *WorkEffortFixedAssetAssignCreate {
	if t != nil {
		wefaac.SetThruDate(*t)
	}
	return wefaac
}

// SetAllocatedCost sets the "allocated_cost" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetAllocatedCost(f float64) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetAllocatedCost(f)
	return wefaac
}

// SetNillableAllocatedCost sets the "allocated_cost" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableAllocatedCost(f *float64) *WorkEffortFixedAssetAssignCreate {
	if f != nil {
		wefaac.SetAllocatedCost(*f)
	}
	return wefaac
}

// SetComments sets the "comments" field.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetComments(s string) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetComments(s)
	return wefaac
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableComments(s *string) *WorkEffortFixedAssetAssignCreate {
	if s != nil {
		wefaac.SetComments(*s)
	}
	return wefaac
}

// SetWorkEffortID sets the "work_effort" edge to the WorkEffort entity by ID.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetWorkEffortID(id int) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetWorkEffortID(id)
	return wefaac
}

// SetNillableWorkEffortID sets the "work_effort" edge to the WorkEffort entity by ID if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableWorkEffortID(id *int) *WorkEffortFixedAssetAssignCreate {
	if id != nil {
		wefaac = wefaac.SetWorkEffortID(*id)
	}
	return wefaac
}

// SetWorkEffort sets the "work_effort" edge to the WorkEffort entity.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetWorkEffort(w *WorkEffort) *WorkEffortFixedAssetAssignCreate {
	return wefaac.SetWorkEffortID(w.ID)
}

// SetFixedAssetID sets the "fixed_asset" edge to the FixedAsset entity by ID.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetFixedAssetID(id int) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetFixedAssetID(id)
	return wefaac
}

// SetNillableFixedAssetID sets the "fixed_asset" edge to the FixedAsset entity by ID if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableFixedAssetID(id *int) *WorkEffortFixedAssetAssignCreate {
	if id != nil {
		wefaac = wefaac.SetFixedAssetID(*id)
	}
	return wefaac
}

// SetFixedAsset sets the "fixed_asset" edge to the FixedAsset entity.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetFixedAsset(f *FixedAsset) *WorkEffortFixedAssetAssignCreate {
	return wefaac.SetFixedAssetID(f.ID)
}

// SetStatusItemID sets the "status_item" edge to the StatusItem entity by ID.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetStatusItemID(id int) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetStatusItemID(id)
	return wefaac
}

// SetNillableStatusItemID sets the "status_item" edge to the StatusItem entity by ID if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableStatusItemID(id *int) *WorkEffortFixedAssetAssignCreate {
	if id != nil {
		wefaac = wefaac.SetStatusItemID(*id)
	}
	return wefaac
}

// SetStatusItem sets the "status_item" edge to the StatusItem entity.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetStatusItem(s *StatusItem) *WorkEffortFixedAssetAssignCreate {
	return wefaac.SetStatusItemID(s.ID)
}

// SetAvailabilityStatusItemID sets the "availability_status_item" edge to the StatusItem entity by ID.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetAvailabilityStatusItemID(id int) *WorkEffortFixedAssetAssignCreate {
	wefaac.mutation.SetAvailabilityStatusItemID(id)
	return wefaac
}

// SetNillableAvailabilityStatusItemID sets the "availability_status_item" edge to the StatusItem entity by ID if the given value is not nil.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetNillableAvailabilityStatusItemID(id *int) *WorkEffortFixedAssetAssignCreate {
	if id != nil {
		wefaac = wefaac.SetAvailabilityStatusItemID(*id)
	}
	return wefaac
}

// SetAvailabilityStatusItem sets the "availability_status_item" edge to the StatusItem entity.
func (wefaac *WorkEffortFixedAssetAssignCreate) SetAvailabilityStatusItem(s *StatusItem) *WorkEffortFixedAssetAssignCreate {
	return wefaac.SetAvailabilityStatusItemID(s.ID)
}

// Mutation returns the WorkEffortFixedAssetAssignMutation object of the builder.
func (wefaac *WorkEffortFixedAssetAssignCreate) Mutation() *WorkEffortFixedAssetAssignMutation {
	return wefaac.mutation
}

// Save creates the WorkEffortFixedAssetAssign in the database.
func (wefaac *WorkEffortFixedAssetAssignCreate) Save(ctx context.Context) (*WorkEffortFixedAssetAssign, error) {
	var (
		err  error
		node *WorkEffortFixedAssetAssign
	)
	wefaac.defaults()
	if len(wefaac.hooks) == 0 {
		if err = wefaac.check(); err != nil {
			return nil, err
		}
		node, err = wefaac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkEffortFixedAssetAssignMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wefaac.check(); err != nil {
				return nil, err
			}
			wefaac.mutation = mutation
			if node, err = wefaac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wefaac.hooks) - 1; i >= 0; i-- {
			mut = wefaac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wefaac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wefaac *WorkEffortFixedAssetAssignCreate) SaveX(ctx context.Context) *WorkEffortFixedAssetAssign {
	v, err := wefaac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (wefaac *WorkEffortFixedAssetAssignCreate) defaults() {
	if _, ok := wefaac.mutation.CreateTime(); !ok {
		v := workeffortfixedassetassign.DefaultCreateTime()
		wefaac.mutation.SetCreateTime(v)
	}
	if _, ok := wefaac.mutation.UpdateTime(); !ok {
		v := workeffortfixedassetassign.DefaultUpdateTime()
		wefaac.mutation.SetUpdateTime(v)
	}
	if _, ok := wefaac.mutation.FromDate(); !ok {
		v := workeffortfixedassetassign.DefaultFromDate()
		wefaac.mutation.SetFromDate(v)
	}
	if _, ok := wefaac.mutation.ThruDate(); !ok {
		v := workeffortfixedassetassign.DefaultThruDate()
		wefaac.mutation.SetThruDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wefaac *WorkEffortFixedAssetAssignCreate) check() error {
	if _, ok := wefaac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := wefaac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := wefaac.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New("ent: missing required field \"from_date\"")}
	}
	return nil
}

func (wefaac *WorkEffortFixedAssetAssignCreate) sqlSave(ctx context.Context) (*WorkEffortFixedAssetAssign, error) {
	_node, _spec := wefaac.createSpec()
	if err := sqlgraph.CreateNode(ctx, wefaac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wefaac *WorkEffortFixedAssetAssignCreate) createSpec() (*WorkEffortFixedAssetAssign, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkEffortFixedAssetAssign{config: wefaac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workeffortfixedassetassign.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workeffortfixedassetassign.FieldID,
			},
		}
	)
	if value, ok := wefaac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortfixedassetassign.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := wefaac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortfixedassetassign.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := wefaac.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workeffortfixedassetassign.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := wefaac.mutation.FromDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortfixedassetassign.FieldFromDate,
		})
		_node.FromDate = value
	}
	if value, ok := wefaac.mutation.ThruDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortfixedassetassign.FieldThruDate,
		})
		_node.ThruDate = value
	}
	if value, ok := wefaac.mutation.AllocatedCost(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: workeffortfixedassetassign.FieldAllocatedCost,
		})
		_node.AllocatedCost = value
	}
	if value, ok := wefaac.mutation.Comments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workeffortfixedassetassign.FieldComments,
		})
		_node.Comments = value
	}
	if nodes := wefaac.mutation.WorkEffortIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortfixedassetassign.WorkEffortTable,
			Columns: []string{workeffortfixedassetassign.WorkEffortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffort.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.work_effort_work_effort_fixed_asset_assigns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wefaac.mutation.FixedAssetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortfixedassetassign.FixedAssetTable,
			Columns: []string{workeffortfixedassetassign.FixedAssetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.fixed_asset_work_effort_fixed_asset_assigns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wefaac.mutation.StatusItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortfixedassetassign.StatusItemTable,
			Columns: []string{workeffortfixedassetassign.StatusItemColumn},
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
		_node.status_item_work_effort_fixed_asset_assigns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wefaac.mutation.AvailabilityStatusItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortfixedassetassign.AvailabilityStatusItemTable,
			Columns: []string{workeffortfixedassetassign.AvailabilityStatusItemColumn},
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
		_node.status_item_availability_work_effort_fixed_asset_assigns = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkEffortFixedAssetAssignCreateBulk is the builder for creating many WorkEffortFixedAssetAssign entities in bulk.
type WorkEffortFixedAssetAssignCreateBulk struct {
	config
	builders []*WorkEffortFixedAssetAssignCreate
}

// Save creates the WorkEffortFixedAssetAssign entities in the database.
func (wefaacb *WorkEffortFixedAssetAssignCreateBulk) Save(ctx context.Context) ([]*WorkEffortFixedAssetAssign, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wefaacb.builders))
	nodes := make([]*WorkEffortFixedAssetAssign, len(wefaacb.builders))
	mutators := make([]Mutator, len(wefaacb.builders))
	for i := range wefaacb.builders {
		func(i int, root context.Context) {
			builder := wefaacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkEffortFixedAssetAssignMutation)
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
					_, err = mutators[i+1].Mutate(root, wefaacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wefaacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wefaacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wefaacb *WorkEffortFixedAssetAssignCreateBulk) SaveX(ctx context.Context) []*WorkEffortFixedAssetAssign {
	v, err := wefaacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
