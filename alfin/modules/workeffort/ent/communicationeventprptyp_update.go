// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventprptyp"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// CommunicationEventPrpTypUpdate is the builder for updating CommunicationEventPrpTyp entities.
type CommunicationEventPrpTypUpdate struct {
	config
	hooks    []Hook
	mutation *CommunicationEventPrpTypMutation
}

// Where adds a new predicate for the CommunicationEventPrpTypUpdate builder.
func (ceptu *CommunicationEventPrpTypUpdate) Where(ps ...predicate.CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.predicates = append(ceptu.mutation.predicates, ps...)
	return ceptu
}

// SetStringRef sets the "string_ref" field.
func (ceptu *CommunicationEventPrpTypUpdate) SetStringRef(s string) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.SetStringRef(s)
	return ceptu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ceptu *CommunicationEventPrpTypUpdate) SetNillableStringRef(s *string) *CommunicationEventPrpTypUpdate {
	if s != nil {
		ceptu.SetStringRef(*s)
	}
	return ceptu
}

// ClearStringRef clears the value of the "string_ref" field.
func (ceptu *CommunicationEventPrpTypUpdate) ClearStringRef() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearStringRef()
	return ceptu
}

// SetHasTable sets the "has_table" field.
func (ceptu *CommunicationEventPrpTypUpdate) SetHasTable(ct communicationeventprptyp.HasTable) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.SetHasTable(ct)
	return ceptu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (ceptu *CommunicationEventPrpTypUpdate) SetNillableHasTable(ct *communicationeventprptyp.HasTable) *CommunicationEventPrpTypUpdate {
	if ct != nil {
		ceptu.SetHasTable(*ct)
	}
	return ceptu
}

// ClearHasTable clears the value of the "has_table" field.
func (ceptu *CommunicationEventPrpTypUpdate) ClearHasTable() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearHasTable()
	return ceptu
}

// SetDescription sets the "description" field.
func (ceptu *CommunicationEventPrpTypUpdate) SetDescription(s string) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.SetDescription(s)
	return ceptu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ceptu *CommunicationEventPrpTypUpdate) SetNillableDescription(s *string) *CommunicationEventPrpTypUpdate {
	if s != nil {
		ceptu.SetDescription(*s)
	}
	return ceptu
}

// ClearDescription clears the value of the "description" field.
func (ceptu *CommunicationEventPrpTypUpdate) ClearDescription() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearDescription()
	return ceptu
}

// SetParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID.
func (ceptu *CommunicationEventPrpTypUpdate) SetParentID(id int) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.SetParentID(id)
	return ceptu
}

// SetNillableParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID if the given value is not nil.
func (ceptu *CommunicationEventPrpTypUpdate) SetNillableParentID(id *int) *CommunicationEventPrpTypUpdate {
	if id != nil {
		ceptu = ceptu.SetParentID(*id)
	}
	return ceptu
}

// SetParent sets the "parent" edge to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) SetParent(c *CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	return ceptu.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptu *CommunicationEventPrpTypUpdate) AddChildIDs(ids ...int) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.AddChildIDs(ids...)
	return ceptu
}

// AddChildren adds the "children" edges to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) AddChildren(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptu.AddChildIDs(ids...)
}

// AddChildCommunicationEventPrpTypIDs adds the "child_communication_event_prp_typs" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptu *CommunicationEventPrpTypUpdate) AddChildCommunicationEventPrpTypIDs(ids ...int) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.AddChildCommunicationEventPrpTypIDs(ids...)
	return ceptu
}

// AddChildCommunicationEventPrpTyps adds the "child_communication_event_prp_typs" edges to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) AddChildCommunicationEventPrpTyps(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptu.AddChildCommunicationEventPrpTypIDs(ids...)
}

// Mutation returns the CommunicationEventPrpTypMutation object of the builder.
func (ceptu *CommunicationEventPrpTypUpdate) Mutation() *CommunicationEventPrpTypMutation {
	return ceptu.mutation
}

// ClearParent clears the "parent" edge to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) ClearParent() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearParent()
	return ceptu
}

// ClearChildren clears all "children" edges to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) ClearChildren() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearChildren()
	return ceptu
}

// RemoveChildIDs removes the "children" edge to CommunicationEventPrpTyp entities by IDs.
func (ceptu *CommunicationEventPrpTypUpdate) RemoveChildIDs(ids ...int) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.RemoveChildIDs(ids...)
	return ceptu
}

// RemoveChildren removes "children" edges to CommunicationEventPrpTyp entities.
func (ceptu *CommunicationEventPrpTypUpdate) RemoveChildren(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptu.RemoveChildIDs(ids...)
}

// ClearChildCommunicationEventPrpTyps clears all "child_communication_event_prp_typs" edges to the CommunicationEventPrpTyp entity.
func (ceptu *CommunicationEventPrpTypUpdate) ClearChildCommunicationEventPrpTyps() *CommunicationEventPrpTypUpdate {
	ceptu.mutation.ClearChildCommunicationEventPrpTyps()
	return ceptu
}

// RemoveChildCommunicationEventPrpTypIDs removes the "child_communication_event_prp_typs" edge to CommunicationEventPrpTyp entities by IDs.
func (ceptu *CommunicationEventPrpTypUpdate) RemoveChildCommunicationEventPrpTypIDs(ids ...int) *CommunicationEventPrpTypUpdate {
	ceptu.mutation.RemoveChildCommunicationEventPrpTypIDs(ids...)
	return ceptu
}

// RemoveChildCommunicationEventPrpTyps removes "child_communication_event_prp_typs" edges to CommunicationEventPrpTyp entities.
func (ceptu *CommunicationEventPrpTypUpdate) RemoveChildCommunicationEventPrpTyps(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptu.RemoveChildCommunicationEventPrpTypIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceptu *CommunicationEventPrpTypUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ceptu.defaults()
	if len(ceptu.hooks) == 0 {
		if err = ceptu.check(); err != nil {
			return 0, err
		}
		affected, err = ceptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunicationEventPrpTypMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ceptu.check(); err != nil {
				return 0, err
			}
			ceptu.mutation = mutation
			affected, err = ceptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceptu.hooks) - 1; i >= 0; i-- {
			mut = ceptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceptu *CommunicationEventPrpTypUpdate) SaveX(ctx context.Context) int {
	affected, err := ceptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceptu *CommunicationEventPrpTypUpdate) Exec(ctx context.Context) error {
	_, err := ceptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceptu *CommunicationEventPrpTypUpdate) ExecX(ctx context.Context) {
	if err := ceptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceptu *CommunicationEventPrpTypUpdate) defaults() {
	if _, ok := ceptu.mutation.UpdateTime(); !ok {
		v := communicationeventprptyp.UpdateDefaultUpdateTime()
		ceptu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ceptu *CommunicationEventPrpTypUpdate) check() error {
	if v, ok := ceptu.mutation.HasTable(); ok {
		if err := communicationeventprptyp.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (ceptu *CommunicationEventPrpTypUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   communicationeventprptyp.Table,
			Columns: communicationeventprptyp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationeventprptyp.FieldID,
			},
		},
	}
	if ps := ceptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceptu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: communicationeventprptyp.FieldUpdateTime,
		})
	}
	if value, ok := ceptu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldStringRef,
		})
	}
	if ceptu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: communicationeventprptyp.FieldStringRef,
		})
	}
	if value, ok := ceptu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: communicationeventprptyp.FieldHasTable,
		})
	}
	if ceptu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: communicationeventprptyp.FieldHasTable,
		})
	}
	if value, ok := ceptu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldDescription,
		})
	}
	if ceptu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: communicationeventprptyp.FieldDescription,
		})
	}
	if ceptu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ceptu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ceptu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptu.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ceptu.mutation.ChildCommunicationEventPrpTypsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptu.mutation.RemovedChildCommunicationEventPrpTypsIDs(); len(nodes) > 0 && !ceptu.mutation.ChildCommunicationEventPrpTypsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptu.mutation.ChildCommunicationEventPrpTypsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ceptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{communicationeventprptyp.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CommunicationEventPrpTypUpdateOne is the builder for updating a single CommunicationEventPrpTyp entity.
type CommunicationEventPrpTypUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommunicationEventPrpTypMutation
}

// SetStringRef sets the "string_ref" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetStringRef(s string) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.SetStringRef(s)
	return ceptuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetNillableStringRef(s *string) *CommunicationEventPrpTypUpdateOne {
	if s != nil {
		ceptuo.SetStringRef(*s)
	}
	return ceptuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearStringRef() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearStringRef()
	return ceptuo
}

// SetHasTable sets the "has_table" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetHasTable(ct communicationeventprptyp.HasTable) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.SetHasTable(ct)
	return ceptuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetNillableHasTable(ct *communicationeventprptyp.HasTable) *CommunicationEventPrpTypUpdateOne {
	if ct != nil {
		ceptuo.SetHasTable(*ct)
	}
	return ceptuo
}

// ClearHasTable clears the value of the "has_table" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearHasTable() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearHasTable()
	return ceptuo
}

// SetDescription sets the "description" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetDescription(s string) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.SetDescription(s)
	return ceptuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetNillableDescription(s *string) *CommunicationEventPrpTypUpdateOne {
	if s != nil {
		ceptuo.SetDescription(*s)
	}
	return ceptuo
}

// ClearDescription clears the value of the "description" field.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearDescription() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearDescription()
	return ceptuo
}

// SetParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetParentID(id int) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.SetParentID(id)
	return ceptuo
}

// SetNillableParentID sets the "parent" edge to the CommunicationEventPrpTyp entity by ID if the given value is not nil.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetNillableParentID(id *int) *CommunicationEventPrpTypUpdateOne {
	if id != nil {
		ceptuo = ceptuo.SetParentID(*id)
	}
	return ceptuo
}

// SetParent sets the "parent" edge to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SetParent(c *CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdateOne {
	return ceptuo.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) AddChildIDs(ids ...int) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.AddChildIDs(ids...)
	return ceptuo
}

// AddChildren adds the "children" edges to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) AddChildren(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptuo.AddChildIDs(ids...)
}

// AddChildCommunicationEventPrpTypIDs adds the "child_communication_event_prp_typs" edge to the CommunicationEventPrpTyp entity by IDs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) AddChildCommunicationEventPrpTypIDs(ids ...int) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.AddChildCommunicationEventPrpTypIDs(ids...)
	return ceptuo
}

// AddChildCommunicationEventPrpTyps adds the "child_communication_event_prp_typs" edges to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) AddChildCommunicationEventPrpTyps(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptuo.AddChildCommunicationEventPrpTypIDs(ids...)
}

// Mutation returns the CommunicationEventPrpTypMutation object of the builder.
func (ceptuo *CommunicationEventPrpTypUpdateOne) Mutation() *CommunicationEventPrpTypMutation {
	return ceptuo.mutation
}

// ClearParent clears the "parent" edge to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearParent() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearParent()
	return ceptuo
}

// ClearChildren clears all "children" edges to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearChildren() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearChildren()
	return ceptuo
}

// RemoveChildIDs removes the "children" edge to CommunicationEventPrpTyp entities by IDs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) RemoveChildIDs(ids ...int) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.RemoveChildIDs(ids...)
	return ceptuo
}

// RemoveChildren removes "children" edges to CommunicationEventPrpTyp entities.
func (ceptuo *CommunicationEventPrpTypUpdateOne) RemoveChildren(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptuo.RemoveChildIDs(ids...)
}

// ClearChildCommunicationEventPrpTyps clears all "child_communication_event_prp_typs" edges to the CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ClearChildCommunicationEventPrpTyps() *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.ClearChildCommunicationEventPrpTyps()
	return ceptuo
}

// RemoveChildCommunicationEventPrpTypIDs removes the "child_communication_event_prp_typs" edge to CommunicationEventPrpTyp entities by IDs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) RemoveChildCommunicationEventPrpTypIDs(ids ...int) *CommunicationEventPrpTypUpdateOne {
	ceptuo.mutation.RemoveChildCommunicationEventPrpTypIDs(ids...)
	return ceptuo
}

// RemoveChildCommunicationEventPrpTyps removes "child_communication_event_prp_typs" edges to CommunicationEventPrpTyp entities.
func (ceptuo *CommunicationEventPrpTypUpdateOne) RemoveChildCommunicationEventPrpTyps(c ...*CommunicationEventPrpTyp) *CommunicationEventPrpTypUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ceptuo.RemoveChildCommunicationEventPrpTypIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ceptuo *CommunicationEventPrpTypUpdateOne) Select(field string, fields ...string) *CommunicationEventPrpTypUpdateOne {
	ceptuo.fields = append([]string{field}, fields...)
	return ceptuo
}

// Save executes the query and returns the updated CommunicationEventPrpTyp entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) Save(ctx context.Context) (*CommunicationEventPrpTyp, error) {
	var (
		err  error
		node *CommunicationEventPrpTyp
	)
	ceptuo.defaults()
	if len(ceptuo.hooks) == 0 {
		if err = ceptuo.check(); err != nil {
			return nil, err
		}
		node, err = ceptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunicationEventPrpTypMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ceptuo.check(); err != nil {
				return nil, err
			}
			ceptuo.mutation = mutation
			node, err = ceptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ceptuo.hooks) - 1; i >= 0; i-- {
			mut = ceptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) SaveX(ctx context.Context) *CommunicationEventPrpTyp {
	node, err := ceptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceptuo *CommunicationEventPrpTypUpdateOne) Exec(ctx context.Context) error {
	_, err := ceptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceptuo *CommunicationEventPrpTypUpdateOne) ExecX(ctx context.Context) {
	if err := ceptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceptuo *CommunicationEventPrpTypUpdateOne) defaults() {
	if _, ok := ceptuo.mutation.UpdateTime(); !ok {
		v := communicationeventprptyp.UpdateDefaultUpdateTime()
		ceptuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ceptuo *CommunicationEventPrpTypUpdateOne) check() error {
	if v, ok := ceptuo.mutation.HasTable(); ok {
		if err := communicationeventprptyp.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (ceptuo *CommunicationEventPrpTypUpdateOne) sqlSave(ctx context.Context) (_node *CommunicationEventPrpTyp, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   communicationeventprptyp.Table,
			Columns: communicationeventprptyp.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationeventprptyp.FieldID,
			},
		},
	}
	id, ok := ceptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CommunicationEventPrpTyp.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ceptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, communicationeventprptyp.FieldID)
		for _, f := range fields {
			if !communicationeventprptyp.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != communicationeventprptyp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ceptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceptuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: communicationeventprptyp.FieldUpdateTime,
		})
	}
	if value, ok := ceptuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldStringRef,
		})
	}
	if ceptuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: communicationeventprptyp.FieldStringRef,
		})
	}
	if value, ok := ceptuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: communicationeventprptyp.FieldHasTable,
		})
	}
	if ceptuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: communicationeventprptyp.FieldHasTable,
		})
	}
	if value, ok := ceptuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: communicationeventprptyp.FieldDescription,
		})
	}
	if ceptuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: communicationeventprptyp.FieldDescription,
		})
	}
	if ceptuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ceptuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ceptuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptuo.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ceptuo.mutation.ChildCommunicationEventPrpTypsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptuo.mutation.RemovedChildCommunicationEventPrpTypsIDs(); len(nodes) > 0 && !ceptuo.mutation.ChildCommunicationEventPrpTypsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ceptuo.mutation.ChildCommunicationEventPrpTypsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CommunicationEventPrpTyp{config: ceptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{communicationeventprptyp.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}