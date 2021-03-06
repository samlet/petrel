// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/temporalexpression"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/temporalexpressionassoc"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
)

// TemporalExpressionCreate is the builder for creating a TemporalExpression entity.
type TemporalExpressionCreate struct {
	config
	mutation *TemporalExpressionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tec *TemporalExpressionCreate) SetCreateTime(t time.Time) *TemporalExpressionCreate {
	tec.mutation.SetCreateTime(t)
	return tec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableCreateTime(t *time.Time) *TemporalExpressionCreate {
	if t != nil {
		tec.SetCreateTime(*t)
	}
	return tec
}

// SetUpdateTime sets the "update_time" field.
func (tec *TemporalExpressionCreate) SetUpdateTime(t time.Time) *TemporalExpressionCreate {
	tec.mutation.SetUpdateTime(t)
	return tec
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableUpdateTime(t *time.Time) *TemporalExpressionCreate {
	if t != nil {
		tec.SetUpdateTime(*t)
	}
	return tec
}

// SetStringRef sets the "string_ref" field.
func (tec *TemporalExpressionCreate) SetStringRef(s string) *TemporalExpressionCreate {
	tec.mutation.SetStringRef(s)
	return tec
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableStringRef(s *string) *TemporalExpressionCreate {
	if s != nil {
		tec.SetStringRef(*s)
	}
	return tec
}

// SetTempExprTypeID sets the "temp_expr_type_id" field.
func (tec *TemporalExpressionCreate) SetTempExprTypeID(i int) *TemporalExpressionCreate {
	tec.mutation.SetTempExprTypeID(i)
	return tec
}

// SetNillableTempExprTypeID sets the "temp_expr_type_id" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableTempExprTypeID(i *int) *TemporalExpressionCreate {
	if i != nil {
		tec.SetTempExprTypeID(*i)
	}
	return tec
}

// SetDescription sets the "description" field.
func (tec *TemporalExpressionCreate) SetDescription(s string) *TemporalExpressionCreate {
	tec.mutation.SetDescription(s)
	return tec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableDescription(s *string) *TemporalExpressionCreate {
	if s != nil {
		tec.SetDescription(*s)
	}
	return tec
}

// SetDate1 sets the "date_1" field.
func (tec *TemporalExpressionCreate) SetDate1(t time.Time) *TemporalExpressionCreate {
	tec.mutation.SetDate1(t)
	return tec
}

// SetNillableDate1 sets the "date_1" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableDate1(t *time.Time) *TemporalExpressionCreate {
	if t != nil {
		tec.SetDate1(*t)
	}
	return tec
}

// SetDate2 sets the "date_2" field.
func (tec *TemporalExpressionCreate) SetDate2(t time.Time) *TemporalExpressionCreate {
	tec.mutation.SetDate2(t)
	return tec
}

// SetNillableDate2 sets the "date_2" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableDate2(t *time.Time) *TemporalExpressionCreate {
	if t != nil {
		tec.SetDate2(*t)
	}
	return tec
}

// SetInteger1 sets the "integer_1" field.
func (tec *TemporalExpressionCreate) SetInteger1(i int) *TemporalExpressionCreate {
	tec.mutation.SetInteger1(i)
	return tec
}

// SetNillableInteger1 sets the "integer_1" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableInteger1(i *int) *TemporalExpressionCreate {
	if i != nil {
		tec.SetInteger1(*i)
	}
	return tec
}

// SetInteger2 sets the "integer_2" field.
func (tec *TemporalExpressionCreate) SetInteger2(i int) *TemporalExpressionCreate {
	tec.mutation.SetInteger2(i)
	return tec
}

// SetNillableInteger2 sets the "integer_2" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableInteger2(i *int) *TemporalExpressionCreate {
	if i != nil {
		tec.SetInteger2(*i)
	}
	return tec
}

// SetString1 sets the "string_1" field.
func (tec *TemporalExpressionCreate) SetString1(i int) *TemporalExpressionCreate {
	tec.mutation.SetString1(i)
	return tec
}

// SetNillableString1 sets the "string_1" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableString1(i *int) *TemporalExpressionCreate {
	if i != nil {
		tec.SetString1(*i)
	}
	return tec
}

// SetString2 sets the "string_2" field.
func (tec *TemporalExpressionCreate) SetString2(i int) *TemporalExpressionCreate {
	tec.mutation.SetString2(i)
	return tec
}

// SetNillableString2 sets the "string_2" field if the given value is not nil.
func (tec *TemporalExpressionCreate) SetNillableString2(i *int) *TemporalExpressionCreate {
	if i != nil {
		tec.SetString2(*i)
	}
	return tec
}

// AddFromTemporalExpressionAssocIDs adds the "from_temporal_expression_assocs" edge to the TemporalExpressionAssoc entity by IDs.
func (tec *TemporalExpressionCreate) AddFromTemporalExpressionAssocIDs(ids ...int) *TemporalExpressionCreate {
	tec.mutation.AddFromTemporalExpressionAssocIDs(ids...)
	return tec
}

// AddFromTemporalExpressionAssocs adds the "from_temporal_expression_assocs" edges to the TemporalExpressionAssoc entity.
func (tec *TemporalExpressionCreate) AddFromTemporalExpressionAssocs(t ...*TemporalExpressionAssoc) *TemporalExpressionCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tec.AddFromTemporalExpressionAssocIDs(ids...)
}

// AddToTemporalExpressionAssocIDs adds the "to_temporal_expression_assocs" edge to the TemporalExpressionAssoc entity by IDs.
func (tec *TemporalExpressionCreate) AddToTemporalExpressionAssocIDs(ids ...int) *TemporalExpressionCreate {
	tec.mutation.AddToTemporalExpressionAssocIDs(ids...)
	return tec
}

// AddToTemporalExpressionAssocs adds the "to_temporal_expression_assocs" edges to the TemporalExpressionAssoc entity.
func (tec *TemporalExpressionCreate) AddToTemporalExpressionAssocs(t ...*TemporalExpressionAssoc) *TemporalExpressionCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tec.AddToTemporalExpressionAssocIDs(ids...)
}

// AddWorkEffortIDs adds the "work_efforts" edge to the WorkEffort entity by IDs.
func (tec *TemporalExpressionCreate) AddWorkEffortIDs(ids ...int) *TemporalExpressionCreate {
	tec.mutation.AddWorkEffortIDs(ids...)
	return tec
}

// AddWorkEfforts adds the "work_efforts" edges to the WorkEffort entity.
func (tec *TemporalExpressionCreate) AddWorkEfforts(w ...*WorkEffort) *TemporalExpressionCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return tec.AddWorkEffortIDs(ids...)
}

// Mutation returns the TemporalExpressionMutation object of the builder.
func (tec *TemporalExpressionCreate) Mutation() *TemporalExpressionMutation {
	return tec.mutation
}

// Save creates the TemporalExpression in the database.
func (tec *TemporalExpressionCreate) Save(ctx context.Context) (*TemporalExpression, error) {
	var (
		err  error
		node *TemporalExpression
	)
	tec.defaults()
	if len(tec.hooks) == 0 {
		if err = tec.check(); err != nil {
			return nil, err
		}
		node, err = tec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemporalExpressionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tec.check(); err != nil {
				return nil, err
			}
			tec.mutation = mutation
			if node, err = tec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tec.hooks) - 1; i >= 0; i-- {
			mut = tec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tec *TemporalExpressionCreate) SaveX(ctx context.Context) *TemporalExpression {
	v, err := tec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tec *TemporalExpressionCreate) defaults() {
	if _, ok := tec.mutation.CreateTime(); !ok {
		v := temporalexpression.DefaultCreateTime()
		tec.mutation.SetCreateTime(v)
	}
	if _, ok := tec.mutation.UpdateTime(); !ok {
		v := temporalexpression.DefaultUpdateTime()
		tec.mutation.SetUpdateTime(v)
	}
	if _, ok := tec.mutation.Date1(); !ok {
		v := temporalexpression.DefaultDate1()
		tec.mutation.SetDate1(v)
	}
	if _, ok := tec.mutation.Date2(); !ok {
		v := temporalexpression.DefaultDate2()
		tec.mutation.SetDate2(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tec *TemporalExpressionCreate) check() error {
	if _, ok := tec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := tec.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (tec *TemporalExpressionCreate) sqlSave(ctx context.Context) (*TemporalExpression, error) {
	_node, _spec := tec.createSpec()
	if err := sqlgraph.CreateNode(ctx, tec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tec *TemporalExpressionCreate) createSpec() (*TemporalExpression, *sqlgraph.CreateSpec) {
	var (
		_node = &TemporalExpression{config: tec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: temporalexpression.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: temporalexpression.FieldID,
			},
		}
	)
	if value, ok := tec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: temporalexpression.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tec.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: temporalexpression.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tec.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: temporalexpression.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := tec.mutation.TempExprTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: temporalexpression.FieldTempExprTypeID,
		})
		_node.TempExprTypeID = value
	}
	if value, ok := tec.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: temporalexpression.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := tec.mutation.Date1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: temporalexpression.FieldDate1,
		})
		_node.Date1 = value
	}
	if value, ok := tec.mutation.Date2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: temporalexpression.FieldDate2,
		})
		_node.Date2 = value
	}
	if value, ok := tec.mutation.Integer1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: temporalexpression.FieldInteger1,
		})
		_node.Integer1 = value
	}
	if value, ok := tec.mutation.Integer2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: temporalexpression.FieldInteger2,
		})
		_node.Integer2 = value
	}
	if value, ok := tec.mutation.String1(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: temporalexpression.FieldString1,
		})
		_node.String1 = value
	}
	if value, ok := tec.mutation.String2(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: temporalexpression.FieldString2,
		})
		_node.String2 = value
	}
	if nodes := tec.mutation.FromTemporalExpressionAssocsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   temporalexpression.FromTemporalExpressionAssocsTable,
			Columns: []string{temporalexpression.FromTemporalExpressionAssocsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: temporalexpressionassoc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tec.mutation.ToTemporalExpressionAssocsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   temporalexpression.ToTemporalExpressionAssocsTable,
			Columns: []string{temporalexpression.ToTemporalExpressionAssocsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: temporalexpressionassoc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tec.mutation.WorkEffortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   temporalexpression.WorkEffortsTable,
			Columns: []string{temporalexpression.WorkEffortsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TemporalExpressionCreateBulk is the builder for creating many TemporalExpression entities in bulk.
type TemporalExpressionCreateBulk struct {
	config
	builders []*TemporalExpressionCreate
}

// Save creates the TemporalExpression entities in the database.
func (tecb *TemporalExpressionCreateBulk) Save(ctx context.Context) ([]*TemporalExpression, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tecb.builders))
	nodes := make([]*TemporalExpression, len(tecb.builders))
	mutators := make([]Mutator, len(tecb.builders))
	for i := range tecb.builders {
		func(i int, root context.Context) {
			builder := tecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemporalExpressionMutation)
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
					_, err = mutators[i+1].Mutate(root, tecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tecb *TemporalExpressionCreateBulk) SaveX(ctx context.Context) []*TemporalExpression {
	v, err := tecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
