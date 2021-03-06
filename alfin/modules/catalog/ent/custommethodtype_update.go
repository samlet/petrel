// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethod"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethodtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// CustomMethodTypeUpdate is the builder for updating CustomMethodType entities.
type CustomMethodTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CustomMethodTypeMutation
}

// Where adds a new predicate for the CustomMethodTypeUpdate builder.
func (cmtu *CustomMethodTypeUpdate) Where(ps ...predicate.CustomMethodType) *CustomMethodTypeUpdate {
	cmtu.mutation.predicates = append(cmtu.mutation.predicates, ps...)
	return cmtu
}

// SetStringRef sets the "string_ref" field.
func (cmtu *CustomMethodTypeUpdate) SetStringRef(s string) *CustomMethodTypeUpdate {
	cmtu.mutation.SetStringRef(s)
	return cmtu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (cmtu *CustomMethodTypeUpdate) SetNillableStringRef(s *string) *CustomMethodTypeUpdate {
	if s != nil {
		cmtu.SetStringRef(*s)
	}
	return cmtu
}

// ClearStringRef clears the value of the "string_ref" field.
func (cmtu *CustomMethodTypeUpdate) ClearStringRef() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearStringRef()
	return cmtu
}

// SetHasTable sets the "has_table" field.
func (cmtu *CustomMethodTypeUpdate) SetHasTable(ct custommethodtype.HasTable) *CustomMethodTypeUpdate {
	cmtu.mutation.SetHasTable(ct)
	return cmtu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (cmtu *CustomMethodTypeUpdate) SetNillableHasTable(ct *custommethodtype.HasTable) *CustomMethodTypeUpdate {
	if ct != nil {
		cmtu.SetHasTable(*ct)
	}
	return cmtu
}

// ClearHasTable clears the value of the "has_table" field.
func (cmtu *CustomMethodTypeUpdate) ClearHasTable() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearHasTable()
	return cmtu
}

// SetDescription sets the "description" field.
func (cmtu *CustomMethodTypeUpdate) SetDescription(s string) *CustomMethodTypeUpdate {
	cmtu.mutation.SetDescription(s)
	return cmtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cmtu *CustomMethodTypeUpdate) SetNillableDescription(s *string) *CustomMethodTypeUpdate {
	if s != nil {
		cmtu.SetDescription(*s)
	}
	return cmtu
}

// ClearDescription clears the value of the "description" field.
func (cmtu *CustomMethodTypeUpdate) ClearDescription() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearDescription()
	return cmtu
}

// SetParentID sets the "parent" edge to the CustomMethodType entity by ID.
func (cmtu *CustomMethodTypeUpdate) SetParentID(id int) *CustomMethodTypeUpdate {
	cmtu.mutation.SetParentID(id)
	return cmtu
}

// SetNillableParentID sets the "parent" edge to the CustomMethodType entity by ID if the given value is not nil.
func (cmtu *CustomMethodTypeUpdate) SetNillableParentID(id *int) *CustomMethodTypeUpdate {
	if id != nil {
		cmtu = cmtu.SetParentID(*id)
	}
	return cmtu
}

// SetParent sets the "parent" edge to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) SetParent(c *CustomMethodType) *CustomMethodTypeUpdate {
	return cmtu.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the CustomMethodType entity by IDs.
func (cmtu *CustomMethodTypeUpdate) AddChildIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.AddChildIDs(ids...)
	return cmtu
}

// AddChildren adds the "children" edges to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) AddChildren(c ...*CustomMethodType) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.AddChildIDs(ids...)
}

// AddCustomMethodIDs adds the "custom_methods" edge to the CustomMethod entity by IDs.
func (cmtu *CustomMethodTypeUpdate) AddCustomMethodIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.AddCustomMethodIDs(ids...)
	return cmtu
}

// AddCustomMethods adds the "custom_methods" edges to the CustomMethod entity.
func (cmtu *CustomMethodTypeUpdate) AddCustomMethods(c ...*CustomMethod) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.AddCustomMethodIDs(ids...)
}

// AddChildCustomMethodTypeIDs adds the "child_custom_method_types" edge to the CustomMethodType entity by IDs.
func (cmtu *CustomMethodTypeUpdate) AddChildCustomMethodTypeIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.AddChildCustomMethodTypeIDs(ids...)
	return cmtu
}

// AddChildCustomMethodTypes adds the "child_custom_method_types" edges to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) AddChildCustomMethodTypes(c ...*CustomMethodType) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.AddChildCustomMethodTypeIDs(ids...)
}

// Mutation returns the CustomMethodTypeMutation object of the builder.
func (cmtu *CustomMethodTypeUpdate) Mutation() *CustomMethodTypeMutation {
	return cmtu.mutation
}

// ClearParent clears the "parent" edge to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) ClearParent() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearParent()
	return cmtu
}

// ClearChildren clears all "children" edges to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) ClearChildren() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearChildren()
	return cmtu
}

// RemoveChildIDs removes the "children" edge to CustomMethodType entities by IDs.
func (cmtu *CustomMethodTypeUpdate) RemoveChildIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.RemoveChildIDs(ids...)
	return cmtu
}

// RemoveChildren removes "children" edges to CustomMethodType entities.
func (cmtu *CustomMethodTypeUpdate) RemoveChildren(c ...*CustomMethodType) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.RemoveChildIDs(ids...)
}

// ClearCustomMethods clears all "custom_methods" edges to the CustomMethod entity.
func (cmtu *CustomMethodTypeUpdate) ClearCustomMethods() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearCustomMethods()
	return cmtu
}

// RemoveCustomMethodIDs removes the "custom_methods" edge to CustomMethod entities by IDs.
func (cmtu *CustomMethodTypeUpdate) RemoveCustomMethodIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.RemoveCustomMethodIDs(ids...)
	return cmtu
}

// RemoveCustomMethods removes "custom_methods" edges to CustomMethod entities.
func (cmtu *CustomMethodTypeUpdate) RemoveCustomMethods(c ...*CustomMethod) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.RemoveCustomMethodIDs(ids...)
}

// ClearChildCustomMethodTypes clears all "child_custom_method_types" edges to the CustomMethodType entity.
func (cmtu *CustomMethodTypeUpdate) ClearChildCustomMethodTypes() *CustomMethodTypeUpdate {
	cmtu.mutation.ClearChildCustomMethodTypes()
	return cmtu
}

// RemoveChildCustomMethodTypeIDs removes the "child_custom_method_types" edge to CustomMethodType entities by IDs.
func (cmtu *CustomMethodTypeUpdate) RemoveChildCustomMethodTypeIDs(ids ...int) *CustomMethodTypeUpdate {
	cmtu.mutation.RemoveChildCustomMethodTypeIDs(ids...)
	return cmtu
}

// RemoveChildCustomMethodTypes removes "child_custom_method_types" edges to CustomMethodType entities.
func (cmtu *CustomMethodTypeUpdate) RemoveChildCustomMethodTypes(c ...*CustomMethodType) *CustomMethodTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtu.RemoveChildCustomMethodTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmtu *CustomMethodTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cmtu.defaults()
	if len(cmtu.hooks) == 0 {
		if err = cmtu.check(); err != nil {
			return 0, err
		}
		affected, err = cmtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomMethodTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmtu.check(); err != nil {
				return 0, err
			}
			cmtu.mutation = mutation
			affected, err = cmtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmtu.hooks) - 1; i >= 0; i-- {
			mut = cmtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmtu *CustomMethodTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := cmtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmtu *CustomMethodTypeUpdate) Exec(ctx context.Context) error {
	_, err := cmtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmtu *CustomMethodTypeUpdate) ExecX(ctx context.Context) {
	if err := cmtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmtu *CustomMethodTypeUpdate) defaults() {
	if _, ok := cmtu.mutation.UpdateTime(); !ok {
		v := custommethodtype.UpdateDefaultUpdateTime()
		cmtu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmtu *CustomMethodTypeUpdate) check() error {
	if v, ok := cmtu.mutation.HasTable(); ok {
		if err := custommethodtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (cmtu *CustomMethodTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   custommethodtype.Table,
			Columns: custommethodtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: custommethodtype.FieldID,
			},
		},
	}
	if ps := cmtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmtu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: custommethodtype.FieldUpdateTime,
		})
	}
	if value, ok := cmtu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: custommethodtype.FieldStringRef,
		})
	}
	if cmtu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: custommethodtype.FieldStringRef,
		})
	}
	if value, ok := cmtu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: custommethodtype.FieldHasTable,
		})
	}
	if cmtu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: custommethodtype.FieldHasTable,
		})
	}
	if value, ok := cmtu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: custommethodtype.FieldDescription,
		})
	}
	if cmtu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: custommethodtype.FieldDescription,
		})
	}
	if cmtu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   custommethodtype.ParentTable,
			Columns: []string{custommethodtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   custommethodtype.ParentTable,
			Columns: []string{custommethodtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cmtu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtu.mutation.CustomMethodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.RemovedCustomMethodsIDs(); len(nodes) > 0 && !cmtu.mutation.CustomMethodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.CustomMethodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtu.mutation.ChildCustomMethodTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.RemovedChildCustomMethodTypesIDs(); len(nodes) > 0 && !cmtu.mutation.ChildCustomMethodTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtu.mutation.ChildCustomMethodTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{custommethodtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomMethodTypeUpdateOne is the builder for updating a single CustomMethodType entity.
type CustomMethodTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomMethodTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (cmtuo *CustomMethodTypeUpdateOne) SetStringRef(s string) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.SetStringRef(s)
	return cmtuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (cmtuo *CustomMethodTypeUpdateOne) SetNillableStringRef(s *string) *CustomMethodTypeUpdateOne {
	if s != nil {
		cmtuo.SetStringRef(*s)
	}
	return cmtuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (cmtuo *CustomMethodTypeUpdateOne) ClearStringRef() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearStringRef()
	return cmtuo
}

// SetHasTable sets the "has_table" field.
func (cmtuo *CustomMethodTypeUpdateOne) SetHasTable(ct custommethodtype.HasTable) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.SetHasTable(ct)
	return cmtuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (cmtuo *CustomMethodTypeUpdateOne) SetNillableHasTable(ct *custommethodtype.HasTable) *CustomMethodTypeUpdateOne {
	if ct != nil {
		cmtuo.SetHasTable(*ct)
	}
	return cmtuo
}

// ClearHasTable clears the value of the "has_table" field.
func (cmtuo *CustomMethodTypeUpdateOne) ClearHasTable() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearHasTable()
	return cmtuo
}

// SetDescription sets the "description" field.
func (cmtuo *CustomMethodTypeUpdateOne) SetDescription(s string) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.SetDescription(s)
	return cmtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cmtuo *CustomMethodTypeUpdateOne) SetNillableDescription(s *string) *CustomMethodTypeUpdateOne {
	if s != nil {
		cmtuo.SetDescription(*s)
	}
	return cmtuo
}

// ClearDescription clears the value of the "description" field.
func (cmtuo *CustomMethodTypeUpdateOne) ClearDescription() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearDescription()
	return cmtuo
}

// SetParentID sets the "parent" edge to the CustomMethodType entity by ID.
func (cmtuo *CustomMethodTypeUpdateOne) SetParentID(id int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.SetParentID(id)
	return cmtuo
}

// SetNillableParentID sets the "parent" edge to the CustomMethodType entity by ID if the given value is not nil.
func (cmtuo *CustomMethodTypeUpdateOne) SetNillableParentID(id *int) *CustomMethodTypeUpdateOne {
	if id != nil {
		cmtuo = cmtuo.SetParentID(*id)
	}
	return cmtuo
}

// SetParent sets the "parent" edge to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) SetParent(c *CustomMethodType) *CustomMethodTypeUpdateOne {
	return cmtuo.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the CustomMethodType entity by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) AddChildIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.AddChildIDs(ids...)
	return cmtuo
}

// AddChildren adds the "children" edges to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) AddChildren(c ...*CustomMethodType) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.AddChildIDs(ids...)
}

// AddCustomMethodIDs adds the "custom_methods" edge to the CustomMethod entity by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) AddCustomMethodIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.AddCustomMethodIDs(ids...)
	return cmtuo
}

// AddCustomMethods adds the "custom_methods" edges to the CustomMethod entity.
func (cmtuo *CustomMethodTypeUpdateOne) AddCustomMethods(c ...*CustomMethod) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.AddCustomMethodIDs(ids...)
}

// AddChildCustomMethodTypeIDs adds the "child_custom_method_types" edge to the CustomMethodType entity by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) AddChildCustomMethodTypeIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.AddChildCustomMethodTypeIDs(ids...)
	return cmtuo
}

// AddChildCustomMethodTypes adds the "child_custom_method_types" edges to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) AddChildCustomMethodTypes(c ...*CustomMethodType) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.AddChildCustomMethodTypeIDs(ids...)
}

// Mutation returns the CustomMethodTypeMutation object of the builder.
func (cmtuo *CustomMethodTypeUpdateOne) Mutation() *CustomMethodTypeMutation {
	return cmtuo.mutation
}

// ClearParent clears the "parent" edge to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) ClearParent() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearParent()
	return cmtuo
}

// ClearChildren clears all "children" edges to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) ClearChildren() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearChildren()
	return cmtuo
}

// RemoveChildIDs removes the "children" edge to CustomMethodType entities by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveChildIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.RemoveChildIDs(ids...)
	return cmtuo
}

// RemoveChildren removes "children" edges to CustomMethodType entities.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveChildren(c ...*CustomMethodType) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.RemoveChildIDs(ids...)
}

// ClearCustomMethods clears all "custom_methods" edges to the CustomMethod entity.
func (cmtuo *CustomMethodTypeUpdateOne) ClearCustomMethods() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearCustomMethods()
	return cmtuo
}

// RemoveCustomMethodIDs removes the "custom_methods" edge to CustomMethod entities by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveCustomMethodIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.RemoveCustomMethodIDs(ids...)
	return cmtuo
}

// RemoveCustomMethods removes "custom_methods" edges to CustomMethod entities.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveCustomMethods(c ...*CustomMethod) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.RemoveCustomMethodIDs(ids...)
}

// ClearChildCustomMethodTypes clears all "child_custom_method_types" edges to the CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) ClearChildCustomMethodTypes() *CustomMethodTypeUpdateOne {
	cmtuo.mutation.ClearChildCustomMethodTypes()
	return cmtuo
}

// RemoveChildCustomMethodTypeIDs removes the "child_custom_method_types" edge to CustomMethodType entities by IDs.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveChildCustomMethodTypeIDs(ids ...int) *CustomMethodTypeUpdateOne {
	cmtuo.mutation.RemoveChildCustomMethodTypeIDs(ids...)
	return cmtuo
}

// RemoveChildCustomMethodTypes removes "child_custom_method_types" edges to CustomMethodType entities.
func (cmtuo *CustomMethodTypeUpdateOne) RemoveChildCustomMethodTypes(c ...*CustomMethodType) *CustomMethodTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtuo.RemoveChildCustomMethodTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmtuo *CustomMethodTypeUpdateOne) Select(field string, fields ...string) *CustomMethodTypeUpdateOne {
	cmtuo.fields = append([]string{field}, fields...)
	return cmtuo
}

// Save executes the query and returns the updated CustomMethodType entity.
func (cmtuo *CustomMethodTypeUpdateOne) Save(ctx context.Context) (*CustomMethodType, error) {
	var (
		err  error
		node *CustomMethodType
	)
	cmtuo.defaults()
	if len(cmtuo.hooks) == 0 {
		if err = cmtuo.check(); err != nil {
			return nil, err
		}
		node, err = cmtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomMethodTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmtuo.check(); err != nil {
				return nil, err
			}
			cmtuo.mutation = mutation
			node, err = cmtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cmtuo.hooks) - 1; i >= 0; i-- {
			mut = cmtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmtuo *CustomMethodTypeUpdateOne) SaveX(ctx context.Context) *CustomMethodType {
	node, err := cmtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmtuo *CustomMethodTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := cmtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmtuo *CustomMethodTypeUpdateOne) ExecX(ctx context.Context) {
	if err := cmtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmtuo *CustomMethodTypeUpdateOne) defaults() {
	if _, ok := cmtuo.mutation.UpdateTime(); !ok {
		v := custommethodtype.UpdateDefaultUpdateTime()
		cmtuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmtuo *CustomMethodTypeUpdateOne) check() error {
	if v, ok := cmtuo.mutation.HasTable(); ok {
		if err := custommethodtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (cmtuo *CustomMethodTypeUpdateOne) sqlSave(ctx context.Context) (_node *CustomMethodType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   custommethodtype.Table,
			Columns: custommethodtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: custommethodtype.FieldID,
			},
		},
	}
	id, ok := cmtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CustomMethodType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cmtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, custommethodtype.FieldID)
		for _, f := range fields {
			if !custommethodtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != custommethodtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmtuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: custommethodtype.FieldUpdateTime,
		})
	}
	if value, ok := cmtuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: custommethodtype.FieldStringRef,
		})
	}
	if cmtuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: custommethodtype.FieldStringRef,
		})
	}
	if value, ok := cmtuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: custommethodtype.FieldHasTable,
		})
	}
	if cmtuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: custommethodtype.FieldHasTable,
		})
	}
	if value, ok := cmtuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: custommethodtype.FieldDescription,
		})
	}
	if cmtuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: custommethodtype.FieldDescription,
		})
	}
	if cmtuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   custommethodtype.ParentTable,
			Columns: []string{custommethodtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   custommethodtype.ParentTable,
			Columns: []string{custommethodtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cmtuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.ChildrenTable,
			Columns: []string{custommethodtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtuo.mutation.CustomMethodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.RemovedCustomMethodsIDs(); len(nodes) > 0 && !cmtuo.mutation.CustomMethodsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.CustomMethodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custommethodtype.CustomMethodsTable,
			Columns: []string{custommethodtype.CustomMethodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmtuo.mutation.ChildCustomMethodTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.RemovedChildCustomMethodTypesIDs(); len(nodes) > 0 && !cmtuo.mutation.ChildCustomMethodTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmtuo.mutation.ChildCustomMethodTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   custommethodtype.ChildCustomMethodTypesTable,
			Columns: custommethodtype.ChildCustomMethodTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: custommethodtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CustomMethodType{config: cmtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{custommethodtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
