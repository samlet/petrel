// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/subscriptiontype"
)

// SubscriptionTypeUpdate is the builder for updating SubscriptionType entities.
type SubscriptionTypeUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriptionTypeMutation
}

// Where adds a new predicate for the SubscriptionTypeUpdate builder.
func (stu *SubscriptionTypeUpdate) Where(ps ...predicate.SubscriptionType) *SubscriptionTypeUpdate {
	stu.mutation.predicates = append(stu.mutation.predicates, ps...)
	return stu
}

// SetStringRef sets the "string_ref" field.
func (stu *SubscriptionTypeUpdate) SetStringRef(s string) *SubscriptionTypeUpdate {
	stu.mutation.SetStringRef(s)
	return stu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (stu *SubscriptionTypeUpdate) SetNillableStringRef(s *string) *SubscriptionTypeUpdate {
	if s != nil {
		stu.SetStringRef(*s)
	}
	return stu
}

// ClearStringRef clears the value of the "string_ref" field.
func (stu *SubscriptionTypeUpdate) ClearStringRef() *SubscriptionTypeUpdate {
	stu.mutation.ClearStringRef()
	return stu
}

// SetHasTable sets the "has_table" field.
func (stu *SubscriptionTypeUpdate) SetHasTable(st subscriptiontype.HasTable) *SubscriptionTypeUpdate {
	stu.mutation.SetHasTable(st)
	return stu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (stu *SubscriptionTypeUpdate) SetNillableHasTable(st *subscriptiontype.HasTable) *SubscriptionTypeUpdate {
	if st != nil {
		stu.SetHasTable(*st)
	}
	return stu
}

// ClearHasTable clears the value of the "has_table" field.
func (stu *SubscriptionTypeUpdate) ClearHasTable() *SubscriptionTypeUpdate {
	stu.mutation.ClearHasTable()
	return stu
}

// SetDescription sets the "description" field.
func (stu *SubscriptionTypeUpdate) SetDescription(s string) *SubscriptionTypeUpdate {
	stu.mutation.SetDescription(s)
	return stu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stu *SubscriptionTypeUpdate) SetNillableDescription(s *string) *SubscriptionTypeUpdate {
	if s != nil {
		stu.SetDescription(*s)
	}
	return stu
}

// ClearDescription clears the value of the "description" field.
func (stu *SubscriptionTypeUpdate) ClearDescription() *SubscriptionTypeUpdate {
	stu.mutation.ClearDescription()
	return stu
}

// SetParentID sets the "parent" edge to the SubscriptionType entity by ID.
func (stu *SubscriptionTypeUpdate) SetParentID(id int) *SubscriptionTypeUpdate {
	stu.mutation.SetParentID(id)
	return stu
}

// SetNillableParentID sets the "parent" edge to the SubscriptionType entity by ID if the given value is not nil.
func (stu *SubscriptionTypeUpdate) SetNillableParentID(id *int) *SubscriptionTypeUpdate {
	if id != nil {
		stu = stu.SetParentID(*id)
	}
	return stu
}

// SetParent sets the "parent" edge to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) SetParent(s *SubscriptionType) *SubscriptionTypeUpdate {
	return stu.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the SubscriptionType entity by IDs.
func (stu *SubscriptionTypeUpdate) AddChildIDs(ids ...int) *SubscriptionTypeUpdate {
	stu.mutation.AddChildIDs(ids...)
	return stu
}

// AddChildren adds the "children" edges to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) AddChildren(s ...*SubscriptionType) *SubscriptionTypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddChildIDs(ids...)
}

// AddChildSubscriptionTypeIDs adds the "child_subscription_types" edge to the SubscriptionType entity by IDs.
func (stu *SubscriptionTypeUpdate) AddChildSubscriptionTypeIDs(ids ...int) *SubscriptionTypeUpdate {
	stu.mutation.AddChildSubscriptionTypeIDs(ids...)
	return stu
}

// AddChildSubscriptionTypes adds the "child_subscription_types" edges to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) AddChildSubscriptionTypes(s ...*SubscriptionType) *SubscriptionTypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.AddChildSubscriptionTypeIDs(ids...)
}

// Mutation returns the SubscriptionTypeMutation object of the builder.
func (stu *SubscriptionTypeUpdate) Mutation() *SubscriptionTypeMutation {
	return stu.mutation
}

// ClearParent clears the "parent" edge to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) ClearParent() *SubscriptionTypeUpdate {
	stu.mutation.ClearParent()
	return stu
}

// ClearChildren clears all "children" edges to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) ClearChildren() *SubscriptionTypeUpdate {
	stu.mutation.ClearChildren()
	return stu
}

// RemoveChildIDs removes the "children" edge to SubscriptionType entities by IDs.
func (stu *SubscriptionTypeUpdate) RemoveChildIDs(ids ...int) *SubscriptionTypeUpdate {
	stu.mutation.RemoveChildIDs(ids...)
	return stu
}

// RemoveChildren removes "children" edges to SubscriptionType entities.
func (stu *SubscriptionTypeUpdate) RemoveChildren(s ...*SubscriptionType) *SubscriptionTypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveChildIDs(ids...)
}

// ClearChildSubscriptionTypes clears all "child_subscription_types" edges to the SubscriptionType entity.
func (stu *SubscriptionTypeUpdate) ClearChildSubscriptionTypes() *SubscriptionTypeUpdate {
	stu.mutation.ClearChildSubscriptionTypes()
	return stu
}

// RemoveChildSubscriptionTypeIDs removes the "child_subscription_types" edge to SubscriptionType entities by IDs.
func (stu *SubscriptionTypeUpdate) RemoveChildSubscriptionTypeIDs(ids ...int) *SubscriptionTypeUpdate {
	stu.mutation.RemoveChildSubscriptionTypeIDs(ids...)
	return stu
}

// RemoveChildSubscriptionTypes removes "child_subscription_types" edges to SubscriptionType entities.
func (stu *SubscriptionTypeUpdate) RemoveChildSubscriptionTypes(s ...*SubscriptionType) *SubscriptionTypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stu.RemoveChildSubscriptionTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SubscriptionTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	stu.defaults()
	if len(stu.hooks) == 0 {
		if err = stu.check(); err != nil {
			return 0, err
		}
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stu.check(); err != nil {
				return 0, err
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SubscriptionTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SubscriptionTypeUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SubscriptionTypeUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stu *SubscriptionTypeUpdate) defaults() {
	if _, ok := stu.mutation.UpdateTime(); !ok {
		v := subscriptiontype.UpdateDefaultUpdateTime()
		stu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *SubscriptionTypeUpdate) check() error {
	if v, ok := stu.mutation.HasTable(); ok {
		if err := subscriptiontype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (stu *SubscriptionTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriptiontype.Table,
			Columns: subscriptiontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriptiontype.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptiontype.FieldUpdateTime,
		})
	}
	if value, ok := stu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldStringRef,
		})
	}
	if stu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriptiontype.FieldStringRef,
		})
	}
	if value, ok := stu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: subscriptiontype.FieldHasTable,
		})
	}
	if stu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: subscriptiontype.FieldHasTable,
		})
	}
	if value, ok := stu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldDescription,
		})
	}
	if stu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriptiontype.FieldDescription,
		})
	}
	if stu.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !stu.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stu.mutation.ChildSubscriptionTypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedChildSubscriptionTypesIDs(); len(nodes) > 0 && !stu.mutation.ChildSubscriptionTypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.ChildSubscriptionTypesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriptiontype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SubscriptionTypeUpdateOne is the builder for updating a single SubscriptionType entity.
type SubscriptionTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriptionTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (stuo *SubscriptionTypeUpdateOne) SetStringRef(s string) *SubscriptionTypeUpdateOne {
	stuo.mutation.SetStringRef(s)
	return stuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (stuo *SubscriptionTypeUpdateOne) SetNillableStringRef(s *string) *SubscriptionTypeUpdateOne {
	if s != nil {
		stuo.SetStringRef(*s)
	}
	return stuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (stuo *SubscriptionTypeUpdateOne) ClearStringRef() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearStringRef()
	return stuo
}

// SetHasTable sets the "has_table" field.
func (stuo *SubscriptionTypeUpdateOne) SetHasTable(st subscriptiontype.HasTable) *SubscriptionTypeUpdateOne {
	stuo.mutation.SetHasTable(st)
	return stuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (stuo *SubscriptionTypeUpdateOne) SetNillableHasTable(st *subscriptiontype.HasTable) *SubscriptionTypeUpdateOne {
	if st != nil {
		stuo.SetHasTable(*st)
	}
	return stuo
}

// ClearHasTable clears the value of the "has_table" field.
func (stuo *SubscriptionTypeUpdateOne) ClearHasTable() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearHasTable()
	return stuo
}

// SetDescription sets the "description" field.
func (stuo *SubscriptionTypeUpdateOne) SetDescription(s string) *SubscriptionTypeUpdateOne {
	stuo.mutation.SetDescription(s)
	return stuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (stuo *SubscriptionTypeUpdateOne) SetNillableDescription(s *string) *SubscriptionTypeUpdateOne {
	if s != nil {
		stuo.SetDescription(*s)
	}
	return stuo
}

// ClearDescription clears the value of the "description" field.
func (stuo *SubscriptionTypeUpdateOne) ClearDescription() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearDescription()
	return stuo
}

// SetParentID sets the "parent" edge to the SubscriptionType entity by ID.
func (stuo *SubscriptionTypeUpdateOne) SetParentID(id int) *SubscriptionTypeUpdateOne {
	stuo.mutation.SetParentID(id)
	return stuo
}

// SetNillableParentID sets the "parent" edge to the SubscriptionType entity by ID if the given value is not nil.
func (stuo *SubscriptionTypeUpdateOne) SetNillableParentID(id *int) *SubscriptionTypeUpdateOne {
	if id != nil {
		stuo = stuo.SetParentID(*id)
	}
	return stuo
}

// SetParent sets the "parent" edge to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) SetParent(s *SubscriptionType) *SubscriptionTypeUpdateOne {
	return stuo.SetParentID(s.ID)
}

// AddChildIDs adds the "children" edge to the SubscriptionType entity by IDs.
func (stuo *SubscriptionTypeUpdateOne) AddChildIDs(ids ...int) *SubscriptionTypeUpdateOne {
	stuo.mutation.AddChildIDs(ids...)
	return stuo
}

// AddChildren adds the "children" edges to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) AddChildren(s ...*SubscriptionType) *SubscriptionTypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddChildIDs(ids...)
}

// AddChildSubscriptionTypeIDs adds the "child_subscription_types" edge to the SubscriptionType entity by IDs.
func (stuo *SubscriptionTypeUpdateOne) AddChildSubscriptionTypeIDs(ids ...int) *SubscriptionTypeUpdateOne {
	stuo.mutation.AddChildSubscriptionTypeIDs(ids...)
	return stuo
}

// AddChildSubscriptionTypes adds the "child_subscription_types" edges to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) AddChildSubscriptionTypes(s ...*SubscriptionType) *SubscriptionTypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.AddChildSubscriptionTypeIDs(ids...)
}

// Mutation returns the SubscriptionTypeMutation object of the builder.
func (stuo *SubscriptionTypeUpdateOne) Mutation() *SubscriptionTypeMutation {
	return stuo.mutation
}

// ClearParent clears the "parent" edge to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) ClearParent() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearParent()
	return stuo
}

// ClearChildren clears all "children" edges to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) ClearChildren() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearChildren()
	return stuo
}

// RemoveChildIDs removes the "children" edge to SubscriptionType entities by IDs.
func (stuo *SubscriptionTypeUpdateOne) RemoveChildIDs(ids ...int) *SubscriptionTypeUpdateOne {
	stuo.mutation.RemoveChildIDs(ids...)
	return stuo
}

// RemoveChildren removes "children" edges to SubscriptionType entities.
func (stuo *SubscriptionTypeUpdateOne) RemoveChildren(s ...*SubscriptionType) *SubscriptionTypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveChildIDs(ids...)
}

// ClearChildSubscriptionTypes clears all "child_subscription_types" edges to the SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) ClearChildSubscriptionTypes() *SubscriptionTypeUpdateOne {
	stuo.mutation.ClearChildSubscriptionTypes()
	return stuo
}

// RemoveChildSubscriptionTypeIDs removes the "child_subscription_types" edge to SubscriptionType entities by IDs.
func (stuo *SubscriptionTypeUpdateOne) RemoveChildSubscriptionTypeIDs(ids ...int) *SubscriptionTypeUpdateOne {
	stuo.mutation.RemoveChildSubscriptionTypeIDs(ids...)
	return stuo
}

// RemoveChildSubscriptionTypes removes "child_subscription_types" edges to SubscriptionType entities.
func (stuo *SubscriptionTypeUpdateOne) RemoveChildSubscriptionTypes(s ...*SubscriptionType) *SubscriptionTypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return stuo.RemoveChildSubscriptionTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SubscriptionTypeUpdateOne) Select(field string, fields ...string) *SubscriptionTypeUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SubscriptionType entity.
func (stuo *SubscriptionTypeUpdateOne) Save(ctx context.Context) (*SubscriptionType, error) {
	var (
		err  error
		node *SubscriptionType
	)
	stuo.defaults()
	if len(stuo.hooks) == 0 {
		if err = stuo.check(); err != nil {
			return nil, err
		}
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriptionTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = stuo.check(); err != nil {
				return nil, err
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			mut = stuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SubscriptionTypeUpdateOne) SaveX(ctx context.Context) *SubscriptionType {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SubscriptionTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SubscriptionTypeUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stuo *SubscriptionTypeUpdateOne) defaults() {
	if _, ok := stuo.mutation.UpdateTime(); !ok {
		v := subscriptiontype.UpdateDefaultUpdateTime()
		stuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *SubscriptionTypeUpdateOne) check() error {
	if v, ok := stuo.mutation.HasTable(); ok {
		if err := subscriptiontype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (stuo *SubscriptionTypeUpdateOne) sqlSave(ctx context.Context) (_node *SubscriptionType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriptiontype.Table,
			Columns: subscriptiontype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriptiontype.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SubscriptionType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscriptiontype.FieldID)
		for _, f := range fields {
			if !subscriptiontype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscriptiontype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: subscriptiontype.FieldUpdateTime,
		})
	}
	if value, ok := stuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldStringRef,
		})
	}
	if stuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriptiontype.FieldStringRef,
		})
	}
	if value, ok := stuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: subscriptiontype.FieldHasTable,
		})
	}
	if stuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: subscriptiontype.FieldHasTable,
		})
	}
	if value, ok := stuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriptiontype.FieldDescription,
		})
	}
	if stuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriptiontype.FieldDescription,
		})
	}
	if stuo.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !stuo.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stuo.mutation.ChildSubscriptionTypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedChildSubscriptionTypesIDs(); len(nodes) > 0 && !stuo.mutation.ChildSubscriptionTypesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.ChildSubscriptionTypesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SubscriptionType{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriptiontype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
