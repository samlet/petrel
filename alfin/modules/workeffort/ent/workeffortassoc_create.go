// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortassoc"
)

// WorkEffortAssocCreate is the builder for creating a WorkEffortAssoc entity.
type WorkEffortAssocCreate struct {
	config
	mutation *WorkEffortAssocMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (weac *WorkEffortAssocCreate) SetCreateTime(t time.Time) *WorkEffortAssocCreate {
	weac.mutation.SetCreateTime(t)
	return weac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableCreateTime(t *time.Time) *WorkEffortAssocCreate {
	if t != nil {
		weac.SetCreateTime(*t)
	}
	return weac
}

// SetUpdateTime sets the "update_time" field.
func (weac *WorkEffortAssocCreate) SetUpdateTime(t time.Time) *WorkEffortAssocCreate {
	weac.mutation.SetUpdateTime(t)
	return weac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableUpdateTime(t *time.Time) *WorkEffortAssocCreate {
	if t != nil {
		weac.SetUpdateTime(*t)
	}
	return weac
}

// SetStringRef sets the "string_ref" field.
func (weac *WorkEffortAssocCreate) SetStringRef(s string) *WorkEffortAssocCreate {
	weac.mutation.SetStringRef(s)
	return weac
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableStringRef(s *string) *WorkEffortAssocCreate {
	if s != nil {
		weac.SetStringRef(*s)
	}
	return weac
}

// SetWorkEffortAssocTypeID sets the "work_effort_assoc_type_id" field.
func (weac *WorkEffortAssocCreate) SetWorkEffortAssocTypeID(i int) *WorkEffortAssocCreate {
	weac.mutation.SetWorkEffortAssocTypeID(i)
	return weac
}

// SetSequenceNum sets the "sequence_num" field.
func (weac *WorkEffortAssocCreate) SetSequenceNum(i int) *WorkEffortAssocCreate {
	weac.mutation.SetSequenceNum(i)
	return weac
}

// SetNillableSequenceNum sets the "sequence_num" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableSequenceNum(i *int) *WorkEffortAssocCreate {
	if i != nil {
		weac.SetSequenceNum(*i)
	}
	return weac
}

// SetFromDate sets the "from_date" field.
func (weac *WorkEffortAssocCreate) SetFromDate(t time.Time) *WorkEffortAssocCreate {
	weac.mutation.SetFromDate(t)
	return weac
}

// SetNillableFromDate sets the "from_date" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableFromDate(t *time.Time) *WorkEffortAssocCreate {
	if t != nil {
		weac.SetFromDate(*t)
	}
	return weac
}

// SetThruDate sets the "thru_date" field.
func (weac *WorkEffortAssocCreate) SetThruDate(t time.Time) *WorkEffortAssocCreate {
	weac.mutation.SetThruDate(t)
	return weac
}

// SetNillableThruDate sets the "thru_date" field if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableThruDate(t *time.Time) *WorkEffortAssocCreate {
	if t != nil {
		weac.SetThruDate(*t)
	}
	return weac
}

// SetFromWorkEffortID sets the "from_work_effort" edge to the WorkEffort entity by ID.
func (weac *WorkEffortAssocCreate) SetFromWorkEffortID(id int) *WorkEffortAssocCreate {
	weac.mutation.SetFromWorkEffortID(id)
	return weac
}

// SetNillableFromWorkEffortID sets the "from_work_effort" edge to the WorkEffort entity by ID if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableFromWorkEffortID(id *int) *WorkEffortAssocCreate {
	if id != nil {
		weac = weac.SetFromWorkEffortID(*id)
	}
	return weac
}

// SetFromWorkEffort sets the "from_work_effort" edge to the WorkEffort entity.
func (weac *WorkEffortAssocCreate) SetFromWorkEffort(w *WorkEffort) *WorkEffortAssocCreate {
	return weac.SetFromWorkEffortID(w.ID)
}

// SetToWorkEffortID sets the "to_work_effort" edge to the WorkEffort entity by ID.
func (weac *WorkEffortAssocCreate) SetToWorkEffortID(id int) *WorkEffortAssocCreate {
	weac.mutation.SetToWorkEffortID(id)
	return weac
}

// SetNillableToWorkEffortID sets the "to_work_effort" edge to the WorkEffort entity by ID if the given value is not nil.
func (weac *WorkEffortAssocCreate) SetNillableToWorkEffortID(id *int) *WorkEffortAssocCreate {
	if id != nil {
		weac = weac.SetToWorkEffortID(*id)
	}
	return weac
}

// SetToWorkEffort sets the "to_work_effort" edge to the WorkEffort entity.
func (weac *WorkEffortAssocCreate) SetToWorkEffort(w *WorkEffort) *WorkEffortAssocCreate {
	return weac.SetToWorkEffortID(w.ID)
}

// Mutation returns the WorkEffortAssocMutation object of the builder.
func (weac *WorkEffortAssocCreate) Mutation() *WorkEffortAssocMutation {
	return weac.mutation
}

// Save creates the WorkEffortAssoc in the database.
func (weac *WorkEffortAssocCreate) Save(ctx context.Context) (*WorkEffortAssoc, error) {
	var (
		err  error
		node *WorkEffortAssoc
	)
	weac.defaults()
	if len(weac.hooks) == 0 {
		if err = weac.check(); err != nil {
			return nil, err
		}
		node, err = weac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkEffortAssocMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = weac.check(); err != nil {
				return nil, err
			}
			weac.mutation = mutation
			if node, err = weac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(weac.hooks) - 1; i >= 0; i-- {
			mut = weac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, weac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (weac *WorkEffortAssocCreate) SaveX(ctx context.Context) *WorkEffortAssoc {
	v, err := weac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (weac *WorkEffortAssocCreate) defaults() {
	if _, ok := weac.mutation.CreateTime(); !ok {
		v := workeffortassoc.DefaultCreateTime()
		weac.mutation.SetCreateTime(v)
	}
	if _, ok := weac.mutation.UpdateTime(); !ok {
		v := workeffortassoc.DefaultUpdateTime()
		weac.mutation.SetUpdateTime(v)
	}
	if _, ok := weac.mutation.FromDate(); !ok {
		v := workeffortassoc.DefaultFromDate()
		weac.mutation.SetFromDate(v)
	}
	if _, ok := weac.mutation.ThruDate(); !ok {
		v := workeffortassoc.DefaultThruDate()
		weac.mutation.SetThruDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (weac *WorkEffortAssocCreate) check() error {
	if _, ok := weac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := weac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := weac.mutation.WorkEffortAssocTypeID(); !ok {
		return &ValidationError{Name: "work_effort_assoc_type_id", err: errors.New("ent: missing required field \"work_effort_assoc_type_id\"")}
	}
	if _, ok := weac.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New("ent: missing required field \"from_date\"")}
	}
	return nil
}

func (weac *WorkEffortAssocCreate) sqlSave(ctx context.Context) (*WorkEffortAssoc, error) {
	_node, _spec := weac.createSpec()
	if err := sqlgraph.CreateNode(ctx, weac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (weac *WorkEffortAssocCreate) createSpec() (*WorkEffortAssoc, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkEffortAssoc{config: weac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workeffortassoc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workeffortassoc.FieldID,
			},
		}
	)
	if value, ok := weac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortassoc.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := weac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortassoc.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := weac.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workeffortassoc.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := weac.mutation.WorkEffortAssocTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workeffortassoc.FieldWorkEffortAssocTypeID,
		})
		_node.WorkEffortAssocTypeID = value
	}
	if value, ok := weac.mutation.SequenceNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workeffortassoc.FieldSequenceNum,
		})
		_node.SequenceNum = value
	}
	if value, ok := weac.mutation.FromDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortassoc.FieldFromDate,
		})
		_node.FromDate = value
	}
	if value, ok := weac.mutation.ThruDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workeffortassoc.FieldThruDate,
		})
		_node.ThruDate = value
	}
	if nodes := weac.mutation.FromWorkEffortIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortassoc.FromWorkEffortTable,
			Columns: []string{workeffortassoc.FromWorkEffortColumn},
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
		_node.work_effort_from_work_effort_assocs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := weac.mutation.ToWorkEffortIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workeffortassoc.ToWorkEffortTable,
			Columns: []string{workeffortassoc.ToWorkEffortColumn},
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
		_node.work_effort_to_work_effort_assocs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkEffortAssocCreateBulk is the builder for creating many WorkEffortAssoc entities in bulk.
type WorkEffortAssocCreateBulk struct {
	config
	builders []*WorkEffortAssocCreate
}

// Save creates the WorkEffortAssoc entities in the database.
func (weacb *WorkEffortAssocCreateBulk) Save(ctx context.Context) ([]*WorkEffortAssoc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(weacb.builders))
	nodes := make([]*WorkEffortAssoc, len(weacb.builders))
	mutators := make([]Mutator, len(weacb.builders))
	for i := range weacb.builders {
		func(i int, root context.Context) {
			builder := weacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkEffortAssocMutation)
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
					_, err = mutators[i+1].Mutate(root, weacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, weacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, weacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (weacb *WorkEffortAssocCreateBulk) SaveX(ctx context.Context) []*WorkEffortAssoc {
	v, err := weacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}