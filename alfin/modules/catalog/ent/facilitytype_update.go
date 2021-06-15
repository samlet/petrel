// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitytype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// FacilityTypeUpdate is the builder for updating FacilityType entities.
type FacilityTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FacilityTypeMutation
}

// Where adds a new predicate for the FacilityTypeUpdate builder.
func (ftu *FacilityTypeUpdate) Where(ps ...predicate.FacilityType) *FacilityTypeUpdate {
	ftu.mutation.predicates = append(ftu.mutation.predicates, ps...)
	return ftu
}

// SetStringRef sets the "string_ref" field.
func (ftu *FacilityTypeUpdate) SetStringRef(s string) *FacilityTypeUpdate {
	ftu.mutation.SetStringRef(s)
	return ftu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ftu *FacilityTypeUpdate) SetNillableStringRef(s *string) *FacilityTypeUpdate {
	if s != nil {
		ftu.SetStringRef(*s)
	}
	return ftu
}

// ClearStringRef clears the value of the "string_ref" field.
func (ftu *FacilityTypeUpdate) ClearStringRef() *FacilityTypeUpdate {
	ftu.mutation.ClearStringRef()
	return ftu
}

// SetHasTable sets the "has_table" field.
func (ftu *FacilityTypeUpdate) SetHasTable(ft facilitytype.HasTable) *FacilityTypeUpdate {
	ftu.mutation.SetHasTable(ft)
	return ftu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (ftu *FacilityTypeUpdate) SetNillableHasTable(ft *facilitytype.HasTable) *FacilityTypeUpdate {
	if ft != nil {
		ftu.SetHasTable(*ft)
	}
	return ftu
}

// ClearHasTable clears the value of the "has_table" field.
func (ftu *FacilityTypeUpdate) ClearHasTable() *FacilityTypeUpdate {
	ftu.mutation.ClearHasTable()
	return ftu
}

// SetDescription sets the "description" field.
func (ftu *FacilityTypeUpdate) SetDescription(s string) *FacilityTypeUpdate {
	ftu.mutation.SetDescription(s)
	return ftu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ftu *FacilityTypeUpdate) SetNillableDescription(s *string) *FacilityTypeUpdate {
	if s != nil {
		ftu.SetDescription(*s)
	}
	return ftu
}

// ClearDescription clears the value of the "description" field.
func (ftu *FacilityTypeUpdate) ClearDescription() *FacilityTypeUpdate {
	ftu.mutation.ClearDescription()
	return ftu
}

// SetParentID sets the "parent" edge to the FacilityType entity by ID.
func (ftu *FacilityTypeUpdate) SetParentID(id int) *FacilityTypeUpdate {
	ftu.mutation.SetParentID(id)
	return ftu
}

// SetNillableParentID sets the "parent" edge to the FacilityType entity by ID if the given value is not nil.
func (ftu *FacilityTypeUpdate) SetNillableParentID(id *int) *FacilityTypeUpdate {
	if id != nil {
		ftu = ftu.SetParentID(*id)
	}
	return ftu
}

// SetParent sets the "parent" edge to the FacilityType entity.
func (ftu *FacilityTypeUpdate) SetParent(f *FacilityType) *FacilityTypeUpdate {
	return ftu.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the FacilityType entity by IDs.
func (ftu *FacilityTypeUpdate) AddChildIDs(ids ...int) *FacilityTypeUpdate {
	ftu.mutation.AddChildIDs(ids...)
	return ftu
}

// AddChildren adds the "children" edges to the FacilityType entity.
func (ftu *FacilityTypeUpdate) AddChildren(f ...*FacilityType) *FacilityTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddChildIDs(ids...)
}

// AddChildFacilityTypeIDs adds the "child_facility_types" edge to the FacilityType entity by IDs.
func (ftu *FacilityTypeUpdate) AddChildFacilityTypeIDs(ids ...int) *FacilityTypeUpdate {
	ftu.mutation.AddChildFacilityTypeIDs(ids...)
	return ftu
}

// AddChildFacilityTypes adds the "child_facility_types" edges to the FacilityType entity.
func (ftu *FacilityTypeUpdate) AddChildFacilityTypes(f ...*FacilityType) *FacilityTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddChildFacilityTypeIDs(ids...)
}

// Mutation returns the FacilityTypeMutation object of the builder.
func (ftu *FacilityTypeUpdate) Mutation() *FacilityTypeMutation {
	return ftu.mutation
}

// ClearParent clears the "parent" edge to the FacilityType entity.
func (ftu *FacilityTypeUpdate) ClearParent() *FacilityTypeUpdate {
	ftu.mutation.ClearParent()
	return ftu
}

// ClearChildren clears all "children" edges to the FacilityType entity.
func (ftu *FacilityTypeUpdate) ClearChildren() *FacilityTypeUpdate {
	ftu.mutation.ClearChildren()
	return ftu
}

// RemoveChildIDs removes the "children" edge to FacilityType entities by IDs.
func (ftu *FacilityTypeUpdate) RemoveChildIDs(ids ...int) *FacilityTypeUpdate {
	ftu.mutation.RemoveChildIDs(ids...)
	return ftu
}

// RemoveChildren removes "children" edges to FacilityType entities.
func (ftu *FacilityTypeUpdate) RemoveChildren(f ...*FacilityType) *FacilityTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveChildIDs(ids...)
}

// ClearChildFacilityTypes clears all "child_facility_types" edges to the FacilityType entity.
func (ftu *FacilityTypeUpdate) ClearChildFacilityTypes() *FacilityTypeUpdate {
	ftu.mutation.ClearChildFacilityTypes()
	return ftu
}

// RemoveChildFacilityTypeIDs removes the "child_facility_types" edge to FacilityType entities by IDs.
func (ftu *FacilityTypeUpdate) RemoveChildFacilityTypeIDs(ids ...int) *FacilityTypeUpdate {
	ftu.mutation.RemoveChildFacilityTypeIDs(ids...)
	return ftu
}

// RemoveChildFacilityTypes removes "child_facility_types" edges to FacilityType entities.
func (ftu *FacilityTypeUpdate) RemoveChildFacilityTypes(f ...*FacilityType) *FacilityTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveChildFacilityTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FacilityTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ftu.defaults()
	if len(ftu.hooks) == 0 {
		if err = ftu.check(); err != nil {
			return 0, err
		}
		affected, err = ftu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacilityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftu.check(); err != nil {
				return 0, err
			}
			ftu.mutation = mutation
			affected, err = ftu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftu.hooks) - 1; i >= 0; i-- {
			mut = ftu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FacilityTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FacilityTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FacilityTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FacilityTypeUpdate) defaults() {
	if _, ok := ftu.mutation.UpdateTime(); !ok {
		v := facilitytype.UpdateDefaultUpdateTime()
		ftu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FacilityTypeUpdate) check() error {
	if v, ok := ftu.mutation.HasTable(); ok {
		if err := facilitytype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (ftu *FacilityTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilitytype.Table,
			Columns: facilitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilitytype.FieldID,
			},
		},
	}
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilitytype.FieldUpdateTime,
		})
	}
	if value, ok := ftu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitytype.FieldStringRef,
		})
	}
	if ftu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitytype.FieldStringRef,
		})
	}
	if value, ok := ftu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: facilitytype.FieldHasTable,
		})
	}
	if ftu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: facilitytype.FieldHasTable,
		})
	}
	if value, ok := ftu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitytype.FieldDescription,
		})
	}
	if ftu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitytype.FieldDescription,
		})
	}
	if ftu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   facilitytype.ParentTable,
			Columns: []string{facilitytype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   facilitytype.ParentTable,
			Columns: []string{facilitytype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ftu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftu.mutation.ChildFacilityTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.RemovedChildFacilityTypesIDs(); len(nodes) > 0 && !ftu.mutation.ChildFacilityTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftu.mutation.ChildFacilityTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facilitytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FacilityTypeUpdateOne is the builder for updating a single FacilityType entity.
type FacilityTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FacilityTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (ftuo *FacilityTypeUpdateOne) SetStringRef(s string) *FacilityTypeUpdateOne {
	ftuo.mutation.SetStringRef(s)
	return ftuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ftuo *FacilityTypeUpdateOne) SetNillableStringRef(s *string) *FacilityTypeUpdateOne {
	if s != nil {
		ftuo.SetStringRef(*s)
	}
	return ftuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (ftuo *FacilityTypeUpdateOne) ClearStringRef() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearStringRef()
	return ftuo
}

// SetHasTable sets the "has_table" field.
func (ftuo *FacilityTypeUpdateOne) SetHasTable(ft facilitytype.HasTable) *FacilityTypeUpdateOne {
	ftuo.mutation.SetHasTable(ft)
	return ftuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (ftuo *FacilityTypeUpdateOne) SetNillableHasTable(ft *facilitytype.HasTable) *FacilityTypeUpdateOne {
	if ft != nil {
		ftuo.SetHasTable(*ft)
	}
	return ftuo
}

// ClearHasTable clears the value of the "has_table" field.
func (ftuo *FacilityTypeUpdateOne) ClearHasTable() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearHasTable()
	return ftuo
}

// SetDescription sets the "description" field.
func (ftuo *FacilityTypeUpdateOne) SetDescription(s string) *FacilityTypeUpdateOne {
	ftuo.mutation.SetDescription(s)
	return ftuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ftuo *FacilityTypeUpdateOne) SetNillableDescription(s *string) *FacilityTypeUpdateOne {
	if s != nil {
		ftuo.SetDescription(*s)
	}
	return ftuo
}

// ClearDescription clears the value of the "description" field.
func (ftuo *FacilityTypeUpdateOne) ClearDescription() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearDescription()
	return ftuo
}

// SetParentID sets the "parent" edge to the FacilityType entity by ID.
func (ftuo *FacilityTypeUpdateOne) SetParentID(id int) *FacilityTypeUpdateOne {
	ftuo.mutation.SetParentID(id)
	return ftuo
}

// SetNillableParentID sets the "parent" edge to the FacilityType entity by ID if the given value is not nil.
func (ftuo *FacilityTypeUpdateOne) SetNillableParentID(id *int) *FacilityTypeUpdateOne {
	if id != nil {
		ftuo = ftuo.SetParentID(*id)
	}
	return ftuo
}

// SetParent sets the "parent" edge to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) SetParent(f *FacilityType) *FacilityTypeUpdateOne {
	return ftuo.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the FacilityType entity by IDs.
func (ftuo *FacilityTypeUpdateOne) AddChildIDs(ids ...int) *FacilityTypeUpdateOne {
	ftuo.mutation.AddChildIDs(ids...)
	return ftuo
}

// AddChildren adds the "children" edges to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) AddChildren(f ...*FacilityType) *FacilityTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddChildIDs(ids...)
}

// AddChildFacilityTypeIDs adds the "child_facility_types" edge to the FacilityType entity by IDs.
func (ftuo *FacilityTypeUpdateOne) AddChildFacilityTypeIDs(ids ...int) *FacilityTypeUpdateOne {
	ftuo.mutation.AddChildFacilityTypeIDs(ids...)
	return ftuo
}

// AddChildFacilityTypes adds the "child_facility_types" edges to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) AddChildFacilityTypes(f ...*FacilityType) *FacilityTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddChildFacilityTypeIDs(ids...)
}

// Mutation returns the FacilityTypeMutation object of the builder.
func (ftuo *FacilityTypeUpdateOne) Mutation() *FacilityTypeMutation {
	return ftuo.mutation
}

// ClearParent clears the "parent" edge to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) ClearParent() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearParent()
	return ftuo
}

// ClearChildren clears all "children" edges to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) ClearChildren() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearChildren()
	return ftuo
}

// RemoveChildIDs removes the "children" edge to FacilityType entities by IDs.
func (ftuo *FacilityTypeUpdateOne) RemoveChildIDs(ids ...int) *FacilityTypeUpdateOne {
	ftuo.mutation.RemoveChildIDs(ids...)
	return ftuo
}

// RemoveChildren removes "children" edges to FacilityType entities.
func (ftuo *FacilityTypeUpdateOne) RemoveChildren(f ...*FacilityType) *FacilityTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveChildIDs(ids...)
}

// ClearChildFacilityTypes clears all "child_facility_types" edges to the FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) ClearChildFacilityTypes() *FacilityTypeUpdateOne {
	ftuo.mutation.ClearChildFacilityTypes()
	return ftuo
}

// RemoveChildFacilityTypeIDs removes the "child_facility_types" edge to FacilityType entities by IDs.
func (ftuo *FacilityTypeUpdateOne) RemoveChildFacilityTypeIDs(ids ...int) *FacilityTypeUpdateOne {
	ftuo.mutation.RemoveChildFacilityTypeIDs(ids...)
	return ftuo
}

// RemoveChildFacilityTypes removes "child_facility_types" edges to FacilityType entities.
func (ftuo *FacilityTypeUpdateOne) RemoveChildFacilityTypes(f ...*FacilityType) *FacilityTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveChildFacilityTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FacilityTypeUpdateOne) Select(field string, fields ...string) *FacilityTypeUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FacilityType entity.
func (ftuo *FacilityTypeUpdateOne) Save(ctx context.Context) (*FacilityType, error) {
	var (
		err  error
		node *FacilityType
	)
	ftuo.defaults()
	if len(ftuo.hooks) == 0 {
		if err = ftuo.check(); err != nil {
			return nil, err
		}
		node, err = ftuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacilityTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftuo.check(); err != nil {
				return nil, err
			}
			ftuo.mutation = mutation
			node, err = ftuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftuo.hooks) - 1; i >= 0; i-- {
			mut = ftuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FacilityTypeUpdateOne) SaveX(ctx context.Context) *FacilityType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FacilityTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FacilityTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FacilityTypeUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdateTime(); !ok {
		v := facilitytype.UpdateDefaultUpdateTime()
		ftuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FacilityTypeUpdateOne) check() error {
	if v, ok := ftuo.mutation.HasTable(); ok {
		if err := facilitytype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (ftuo *FacilityTypeUpdateOne) sqlSave(ctx context.Context) (_node *FacilityType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilitytype.Table,
			Columns: facilitytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilitytype.FieldID,
			},
		},
	}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FacilityType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, facilitytype.FieldID)
		for _, f := range fields {
			if !facilitytype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != facilitytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilitytype.FieldUpdateTime,
		})
	}
	if value, ok := ftuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitytype.FieldStringRef,
		})
	}
	if ftuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitytype.FieldStringRef,
		})
	}
	if value, ok := ftuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: facilitytype.FieldHasTable,
		})
	}
	if ftuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: facilitytype.FieldHasTable,
		})
	}
	if value, ok := ftuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitytype.FieldDescription,
		})
	}
	if ftuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitytype.FieldDescription,
		})
	}
	if ftuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   facilitytype.ParentTable,
			Columns: []string{facilitytype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   facilitytype.ParentTable,
			Columns: []string{facilitytype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ftuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitytype.ChildrenTable,
			Columns: []string{facilitytype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ftuo.mutation.ChildFacilityTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.RemovedChildFacilityTypesIDs(); len(nodes) > 0 && !ftuo.mutation.ChildFacilityTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ftuo.mutation.ChildFacilityTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   facilitytype.ChildFacilityTypesTable,
			Columns: facilitytype.ChildFacilityTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FacilityType{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facilitytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}