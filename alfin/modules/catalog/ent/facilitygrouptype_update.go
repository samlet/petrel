// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygrouptype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// FacilityGroupTypeUpdate is the builder for updating FacilityGroupType entities.
type FacilityGroupTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FacilityGroupTypeMutation
}

// Where adds a new predicate for the FacilityGroupTypeUpdate builder.
func (fgtu *FacilityGroupTypeUpdate) Where(ps ...predicate.FacilityGroupType) *FacilityGroupTypeUpdate {
	fgtu.mutation.predicates = append(fgtu.mutation.predicates, ps...)
	return fgtu
}

// SetStringRef sets the "string_ref" field.
func (fgtu *FacilityGroupTypeUpdate) SetStringRef(s string) *FacilityGroupTypeUpdate {
	fgtu.mutation.SetStringRef(s)
	return fgtu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (fgtu *FacilityGroupTypeUpdate) SetNillableStringRef(s *string) *FacilityGroupTypeUpdate {
	if s != nil {
		fgtu.SetStringRef(*s)
	}
	return fgtu
}

// ClearStringRef clears the value of the "string_ref" field.
func (fgtu *FacilityGroupTypeUpdate) ClearStringRef() *FacilityGroupTypeUpdate {
	fgtu.mutation.ClearStringRef()
	return fgtu
}

// SetDescription sets the "description" field.
func (fgtu *FacilityGroupTypeUpdate) SetDescription(s string) *FacilityGroupTypeUpdate {
	fgtu.mutation.SetDescription(s)
	return fgtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fgtu *FacilityGroupTypeUpdate) SetNillableDescription(s *string) *FacilityGroupTypeUpdate {
	if s != nil {
		fgtu.SetDescription(*s)
	}
	return fgtu
}

// ClearDescription clears the value of the "description" field.
func (fgtu *FacilityGroupTypeUpdate) ClearDescription() *FacilityGroupTypeUpdate {
	fgtu.mutation.ClearDescription()
	return fgtu
}

// AddFacilityGroupIDs adds the "facility_groups" edge to the FacilityGroup entity by IDs.
func (fgtu *FacilityGroupTypeUpdate) AddFacilityGroupIDs(ids ...int) *FacilityGroupTypeUpdate {
	fgtu.mutation.AddFacilityGroupIDs(ids...)
	return fgtu
}

// AddFacilityGroups adds the "facility_groups" edges to the FacilityGroup entity.
func (fgtu *FacilityGroupTypeUpdate) AddFacilityGroups(f ...*FacilityGroup) *FacilityGroupTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fgtu.AddFacilityGroupIDs(ids...)
}

// Mutation returns the FacilityGroupTypeMutation object of the builder.
func (fgtu *FacilityGroupTypeUpdate) Mutation() *FacilityGroupTypeMutation {
	return fgtu.mutation
}

// ClearFacilityGroups clears all "facility_groups" edges to the FacilityGroup entity.
func (fgtu *FacilityGroupTypeUpdate) ClearFacilityGroups() *FacilityGroupTypeUpdate {
	fgtu.mutation.ClearFacilityGroups()
	return fgtu
}

// RemoveFacilityGroupIDs removes the "facility_groups" edge to FacilityGroup entities by IDs.
func (fgtu *FacilityGroupTypeUpdate) RemoveFacilityGroupIDs(ids ...int) *FacilityGroupTypeUpdate {
	fgtu.mutation.RemoveFacilityGroupIDs(ids...)
	return fgtu
}

// RemoveFacilityGroups removes "facility_groups" edges to FacilityGroup entities.
func (fgtu *FacilityGroupTypeUpdate) RemoveFacilityGroups(f ...*FacilityGroup) *FacilityGroupTypeUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fgtu.RemoveFacilityGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fgtu *FacilityGroupTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fgtu.defaults()
	if len(fgtu.hooks) == 0 {
		affected, err = fgtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacilityGroupTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fgtu.mutation = mutation
			affected, err = fgtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fgtu.hooks) - 1; i >= 0; i-- {
			mut = fgtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fgtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fgtu *FacilityGroupTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := fgtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fgtu *FacilityGroupTypeUpdate) Exec(ctx context.Context) error {
	_, err := fgtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgtu *FacilityGroupTypeUpdate) ExecX(ctx context.Context) {
	if err := fgtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fgtu *FacilityGroupTypeUpdate) defaults() {
	if _, ok := fgtu.mutation.UpdateTime(); !ok {
		v := facilitygrouptype.UpdateDefaultUpdateTime()
		fgtu.mutation.SetUpdateTime(v)
	}
}

func (fgtu *FacilityGroupTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilitygrouptype.Table,
			Columns: facilitygrouptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilitygrouptype.FieldID,
			},
		},
	}
	if ps := fgtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fgtu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilitygrouptype.FieldUpdateTime,
		})
	}
	if value, ok := fgtu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitygrouptype.FieldStringRef,
		})
	}
	if fgtu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitygrouptype.FieldStringRef,
		})
	}
	if value, ok := fgtu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitygrouptype.FieldDescription,
		})
	}
	if fgtu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitygrouptype.FieldDescription,
		})
	}
	if fgtu.mutation.FacilityGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fgtu.mutation.RemovedFacilityGroupsIDs(); len(nodes) > 0 && !fgtu.mutation.FacilityGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fgtu.mutation.FacilityGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fgtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facilitygrouptype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FacilityGroupTypeUpdateOne is the builder for updating a single FacilityGroupType entity.
type FacilityGroupTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FacilityGroupTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (fgtuo *FacilityGroupTypeUpdateOne) SetStringRef(s string) *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.SetStringRef(s)
	return fgtuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (fgtuo *FacilityGroupTypeUpdateOne) SetNillableStringRef(s *string) *FacilityGroupTypeUpdateOne {
	if s != nil {
		fgtuo.SetStringRef(*s)
	}
	return fgtuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (fgtuo *FacilityGroupTypeUpdateOne) ClearStringRef() *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.ClearStringRef()
	return fgtuo
}

// SetDescription sets the "description" field.
func (fgtuo *FacilityGroupTypeUpdateOne) SetDescription(s string) *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.SetDescription(s)
	return fgtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fgtuo *FacilityGroupTypeUpdateOne) SetNillableDescription(s *string) *FacilityGroupTypeUpdateOne {
	if s != nil {
		fgtuo.SetDescription(*s)
	}
	return fgtuo
}

// ClearDescription clears the value of the "description" field.
func (fgtuo *FacilityGroupTypeUpdateOne) ClearDescription() *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.ClearDescription()
	return fgtuo
}

// AddFacilityGroupIDs adds the "facility_groups" edge to the FacilityGroup entity by IDs.
func (fgtuo *FacilityGroupTypeUpdateOne) AddFacilityGroupIDs(ids ...int) *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.AddFacilityGroupIDs(ids...)
	return fgtuo
}

// AddFacilityGroups adds the "facility_groups" edges to the FacilityGroup entity.
func (fgtuo *FacilityGroupTypeUpdateOne) AddFacilityGroups(f ...*FacilityGroup) *FacilityGroupTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fgtuo.AddFacilityGroupIDs(ids...)
}

// Mutation returns the FacilityGroupTypeMutation object of the builder.
func (fgtuo *FacilityGroupTypeUpdateOne) Mutation() *FacilityGroupTypeMutation {
	return fgtuo.mutation
}

// ClearFacilityGroups clears all "facility_groups" edges to the FacilityGroup entity.
func (fgtuo *FacilityGroupTypeUpdateOne) ClearFacilityGroups() *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.ClearFacilityGroups()
	return fgtuo
}

// RemoveFacilityGroupIDs removes the "facility_groups" edge to FacilityGroup entities by IDs.
func (fgtuo *FacilityGroupTypeUpdateOne) RemoveFacilityGroupIDs(ids ...int) *FacilityGroupTypeUpdateOne {
	fgtuo.mutation.RemoveFacilityGroupIDs(ids...)
	return fgtuo
}

// RemoveFacilityGroups removes "facility_groups" edges to FacilityGroup entities.
func (fgtuo *FacilityGroupTypeUpdateOne) RemoveFacilityGroups(f ...*FacilityGroup) *FacilityGroupTypeUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fgtuo.RemoveFacilityGroupIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fgtuo *FacilityGroupTypeUpdateOne) Select(field string, fields ...string) *FacilityGroupTypeUpdateOne {
	fgtuo.fields = append([]string{field}, fields...)
	return fgtuo
}

// Save executes the query and returns the updated FacilityGroupType entity.
func (fgtuo *FacilityGroupTypeUpdateOne) Save(ctx context.Context) (*FacilityGroupType, error) {
	var (
		err  error
		node *FacilityGroupType
	)
	fgtuo.defaults()
	if len(fgtuo.hooks) == 0 {
		node, err = fgtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FacilityGroupTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fgtuo.mutation = mutation
			node, err = fgtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fgtuo.hooks) - 1; i >= 0; i-- {
			mut = fgtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fgtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fgtuo *FacilityGroupTypeUpdateOne) SaveX(ctx context.Context) *FacilityGroupType {
	node, err := fgtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fgtuo *FacilityGroupTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := fgtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fgtuo *FacilityGroupTypeUpdateOne) ExecX(ctx context.Context) {
	if err := fgtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fgtuo *FacilityGroupTypeUpdateOne) defaults() {
	if _, ok := fgtuo.mutation.UpdateTime(); !ok {
		v := facilitygrouptype.UpdateDefaultUpdateTime()
		fgtuo.mutation.SetUpdateTime(v)
	}
}

func (fgtuo *FacilityGroupTypeUpdateOne) sqlSave(ctx context.Context) (_node *FacilityGroupType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilitygrouptype.Table,
			Columns: facilitygrouptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilitygrouptype.FieldID,
			},
		},
	}
	id, ok := fgtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FacilityGroupType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := fgtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, facilitygrouptype.FieldID)
		for _, f := range fields {
			if !facilitygrouptype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != facilitygrouptype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fgtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fgtuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: facilitygrouptype.FieldUpdateTime,
		})
	}
	if value, ok := fgtuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitygrouptype.FieldStringRef,
		})
	}
	if fgtuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitygrouptype.FieldStringRef,
		})
	}
	if value, ok := fgtuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: facilitygrouptype.FieldDescription,
		})
	}
	if fgtuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: facilitygrouptype.FieldDescription,
		})
	}
	if fgtuo.mutation.FacilityGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fgtuo.mutation.RemovedFacilityGroupsIDs(); len(nodes) > 0 && !fgtuo.mutation.FacilityGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fgtuo.mutation.FacilityGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   facilitygrouptype.FacilityGroupsTable,
			Columns: []string{facilitygrouptype.FacilityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: facilitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FacilityGroupType{config: fgtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fgtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{facilitygrouptype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
