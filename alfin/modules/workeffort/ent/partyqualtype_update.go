// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyqualtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// PartyQualTypeUpdate is the builder for updating PartyQualType entities.
type PartyQualTypeUpdate struct {
	config
	hooks    []Hook
	mutation *PartyQualTypeMutation
}

// Where adds a new predicate for the PartyQualTypeUpdate builder.
func (pqtu *PartyQualTypeUpdate) Where(ps ...predicate.PartyQualType) *PartyQualTypeUpdate {
	pqtu.mutation.predicates = append(pqtu.mutation.predicates, ps...)
	return pqtu
}

// SetStringRef sets the "string_ref" field.
func (pqtu *PartyQualTypeUpdate) SetStringRef(s string) *PartyQualTypeUpdate {
	pqtu.mutation.SetStringRef(s)
	return pqtu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pqtu *PartyQualTypeUpdate) SetNillableStringRef(s *string) *PartyQualTypeUpdate {
	if s != nil {
		pqtu.SetStringRef(*s)
	}
	return pqtu
}

// ClearStringRef clears the value of the "string_ref" field.
func (pqtu *PartyQualTypeUpdate) ClearStringRef() *PartyQualTypeUpdate {
	pqtu.mutation.ClearStringRef()
	return pqtu
}

// SetHasTable sets the "has_table" field.
func (pqtu *PartyQualTypeUpdate) SetHasTable(pt partyqualtype.HasTable) *PartyQualTypeUpdate {
	pqtu.mutation.SetHasTable(pt)
	return pqtu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pqtu *PartyQualTypeUpdate) SetNillableHasTable(pt *partyqualtype.HasTable) *PartyQualTypeUpdate {
	if pt != nil {
		pqtu.SetHasTable(*pt)
	}
	return pqtu
}

// ClearHasTable clears the value of the "has_table" field.
func (pqtu *PartyQualTypeUpdate) ClearHasTable() *PartyQualTypeUpdate {
	pqtu.mutation.ClearHasTable()
	return pqtu
}

// SetDescription sets the "description" field.
func (pqtu *PartyQualTypeUpdate) SetDescription(s string) *PartyQualTypeUpdate {
	pqtu.mutation.SetDescription(s)
	return pqtu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pqtu *PartyQualTypeUpdate) SetNillableDescription(s *string) *PartyQualTypeUpdate {
	if s != nil {
		pqtu.SetDescription(*s)
	}
	return pqtu
}

// ClearDescription clears the value of the "description" field.
func (pqtu *PartyQualTypeUpdate) ClearDescription() *PartyQualTypeUpdate {
	pqtu.mutation.ClearDescription()
	return pqtu
}

// SetParentID sets the "parent" edge to the PartyQualType entity by ID.
func (pqtu *PartyQualTypeUpdate) SetParentID(id int) *PartyQualTypeUpdate {
	pqtu.mutation.SetParentID(id)
	return pqtu
}

// SetNillableParentID sets the "parent" edge to the PartyQualType entity by ID if the given value is not nil.
func (pqtu *PartyQualTypeUpdate) SetNillableParentID(id *int) *PartyQualTypeUpdate {
	if id != nil {
		pqtu = pqtu.SetParentID(*id)
	}
	return pqtu
}

// SetParent sets the "parent" edge to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) SetParent(p *PartyQualType) *PartyQualTypeUpdate {
	return pqtu.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the PartyQualType entity by IDs.
func (pqtu *PartyQualTypeUpdate) AddChildIDs(ids ...int) *PartyQualTypeUpdate {
	pqtu.mutation.AddChildIDs(ids...)
	return pqtu
}

// AddChildren adds the "children" edges to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) AddChildren(p ...*PartyQualType) *PartyQualTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtu.AddChildIDs(ids...)
}

// AddChildPartyQualTypeIDs adds the "child_party_qual_types" edge to the PartyQualType entity by IDs.
func (pqtu *PartyQualTypeUpdate) AddChildPartyQualTypeIDs(ids ...int) *PartyQualTypeUpdate {
	pqtu.mutation.AddChildPartyQualTypeIDs(ids...)
	return pqtu
}

// AddChildPartyQualTypes adds the "child_party_qual_types" edges to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) AddChildPartyQualTypes(p ...*PartyQualType) *PartyQualTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtu.AddChildPartyQualTypeIDs(ids...)
}

// Mutation returns the PartyQualTypeMutation object of the builder.
func (pqtu *PartyQualTypeUpdate) Mutation() *PartyQualTypeMutation {
	return pqtu.mutation
}

// ClearParent clears the "parent" edge to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) ClearParent() *PartyQualTypeUpdate {
	pqtu.mutation.ClearParent()
	return pqtu
}

// ClearChildren clears all "children" edges to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) ClearChildren() *PartyQualTypeUpdate {
	pqtu.mutation.ClearChildren()
	return pqtu
}

// RemoveChildIDs removes the "children" edge to PartyQualType entities by IDs.
func (pqtu *PartyQualTypeUpdate) RemoveChildIDs(ids ...int) *PartyQualTypeUpdate {
	pqtu.mutation.RemoveChildIDs(ids...)
	return pqtu
}

// RemoveChildren removes "children" edges to PartyQualType entities.
func (pqtu *PartyQualTypeUpdate) RemoveChildren(p ...*PartyQualType) *PartyQualTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtu.RemoveChildIDs(ids...)
}

// ClearChildPartyQualTypes clears all "child_party_qual_types" edges to the PartyQualType entity.
func (pqtu *PartyQualTypeUpdate) ClearChildPartyQualTypes() *PartyQualTypeUpdate {
	pqtu.mutation.ClearChildPartyQualTypes()
	return pqtu
}

// RemoveChildPartyQualTypeIDs removes the "child_party_qual_types" edge to PartyQualType entities by IDs.
func (pqtu *PartyQualTypeUpdate) RemoveChildPartyQualTypeIDs(ids ...int) *PartyQualTypeUpdate {
	pqtu.mutation.RemoveChildPartyQualTypeIDs(ids...)
	return pqtu
}

// RemoveChildPartyQualTypes removes "child_party_qual_types" edges to PartyQualType entities.
func (pqtu *PartyQualTypeUpdate) RemoveChildPartyQualTypes(p ...*PartyQualType) *PartyQualTypeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtu.RemoveChildPartyQualTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pqtu *PartyQualTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pqtu.defaults()
	if len(pqtu.hooks) == 0 {
		if err = pqtu.check(); err != nil {
			return 0, err
		}
		affected, err = pqtu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyQualTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pqtu.check(); err != nil {
				return 0, err
			}
			pqtu.mutation = mutation
			affected, err = pqtu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pqtu.hooks) - 1; i >= 0; i-- {
			mut = pqtu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pqtu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pqtu *PartyQualTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := pqtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pqtu *PartyQualTypeUpdate) Exec(ctx context.Context) error {
	_, err := pqtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqtu *PartyQualTypeUpdate) ExecX(ctx context.Context) {
	if err := pqtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pqtu *PartyQualTypeUpdate) defaults() {
	if _, ok := pqtu.mutation.UpdateTime(); !ok {
		v := partyqualtype.UpdateDefaultUpdateTime()
		pqtu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pqtu *PartyQualTypeUpdate) check() error {
	if v, ok := pqtu.mutation.HasTable(); ok {
		if err := partyqualtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pqtu *PartyQualTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partyqualtype.Table,
			Columns: partyqualtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyqualtype.FieldID,
			},
		},
	}
	if ps := pqtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pqtu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyqualtype.FieldUpdateTime,
		})
	}
	if value, ok := pqtu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldStringRef,
		})
	}
	if pqtu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyqualtype.FieldStringRef,
		})
	}
	if value, ok := pqtu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: partyqualtype.FieldHasTable,
		})
	}
	if pqtu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: partyqualtype.FieldHasTable,
		})
	}
	if value, ok := pqtu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldDescription,
		})
	}
	if pqtu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyqualtype.FieldDescription,
		})
	}
	if pqtu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyqualtype.ParentTable,
			Columns: []string{partyqualtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyqualtype.ParentTable,
			Columns: []string{partyqualtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pqtu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pqtu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pqtu.mutation.ChildPartyQualTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtu.mutation.RemovedChildPartyQualTypesIDs(); len(nodes) > 0 && !pqtu.mutation.ChildPartyQualTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtu.mutation.ChildPartyQualTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pqtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partyqualtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PartyQualTypeUpdateOne is the builder for updating a single PartyQualType entity.
type PartyQualTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PartyQualTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (pqtuo *PartyQualTypeUpdateOne) SetStringRef(s string) *PartyQualTypeUpdateOne {
	pqtuo.mutation.SetStringRef(s)
	return pqtuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pqtuo *PartyQualTypeUpdateOne) SetNillableStringRef(s *string) *PartyQualTypeUpdateOne {
	if s != nil {
		pqtuo.SetStringRef(*s)
	}
	return pqtuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (pqtuo *PartyQualTypeUpdateOne) ClearStringRef() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearStringRef()
	return pqtuo
}

// SetHasTable sets the "has_table" field.
func (pqtuo *PartyQualTypeUpdateOne) SetHasTable(pt partyqualtype.HasTable) *PartyQualTypeUpdateOne {
	pqtuo.mutation.SetHasTable(pt)
	return pqtuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pqtuo *PartyQualTypeUpdateOne) SetNillableHasTable(pt *partyqualtype.HasTable) *PartyQualTypeUpdateOne {
	if pt != nil {
		pqtuo.SetHasTable(*pt)
	}
	return pqtuo
}

// ClearHasTable clears the value of the "has_table" field.
func (pqtuo *PartyQualTypeUpdateOne) ClearHasTable() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearHasTable()
	return pqtuo
}

// SetDescription sets the "description" field.
func (pqtuo *PartyQualTypeUpdateOne) SetDescription(s string) *PartyQualTypeUpdateOne {
	pqtuo.mutation.SetDescription(s)
	return pqtuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pqtuo *PartyQualTypeUpdateOne) SetNillableDescription(s *string) *PartyQualTypeUpdateOne {
	if s != nil {
		pqtuo.SetDescription(*s)
	}
	return pqtuo
}

// ClearDescription clears the value of the "description" field.
func (pqtuo *PartyQualTypeUpdateOne) ClearDescription() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearDescription()
	return pqtuo
}

// SetParentID sets the "parent" edge to the PartyQualType entity by ID.
func (pqtuo *PartyQualTypeUpdateOne) SetParentID(id int) *PartyQualTypeUpdateOne {
	pqtuo.mutation.SetParentID(id)
	return pqtuo
}

// SetNillableParentID sets the "parent" edge to the PartyQualType entity by ID if the given value is not nil.
func (pqtuo *PartyQualTypeUpdateOne) SetNillableParentID(id *int) *PartyQualTypeUpdateOne {
	if id != nil {
		pqtuo = pqtuo.SetParentID(*id)
	}
	return pqtuo
}

// SetParent sets the "parent" edge to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) SetParent(p *PartyQualType) *PartyQualTypeUpdateOne {
	return pqtuo.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the PartyQualType entity by IDs.
func (pqtuo *PartyQualTypeUpdateOne) AddChildIDs(ids ...int) *PartyQualTypeUpdateOne {
	pqtuo.mutation.AddChildIDs(ids...)
	return pqtuo
}

// AddChildren adds the "children" edges to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) AddChildren(p ...*PartyQualType) *PartyQualTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtuo.AddChildIDs(ids...)
}

// AddChildPartyQualTypeIDs adds the "child_party_qual_types" edge to the PartyQualType entity by IDs.
func (pqtuo *PartyQualTypeUpdateOne) AddChildPartyQualTypeIDs(ids ...int) *PartyQualTypeUpdateOne {
	pqtuo.mutation.AddChildPartyQualTypeIDs(ids...)
	return pqtuo
}

// AddChildPartyQualTypes adds the "child_party_qual_types" edges to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) AddChildPartyQualTypes(p ...*PartyQualType) *PartyQualTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtuo.AddChildPartyQualTypeIDs(ids...)
}

// Mutation returns the PartyQualTypeMutation object of the builder.
func (pqtuo *PartyQualTypeUpdateOne) Mutation() *PartyQualTypeMutation {
	return pqtuo.mutation
}

// ClearParent clears the "parent" edge to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) ClearParent() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearParent()
	return pqtuo
}

// ClearChildren clears all "children" edges to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) ClearChildren() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearChildren()
	return pqtuo
}

// RemoveChildIDs removes the "children" edge to PartyQualType entities by IDs.
func (pqtuo *PartyQualTypeUpdateOne) RemoveChildIDs(ids ...int) *PartyQualTypeUpdateOne {
	pqtuo.mutation.RemoveChildIDs(ids...)
	return pqtuo
}

// RemoveChildren removes "children" edges to PartyQualType entities.
func (pqtuo *PartyQualTypeUpdateOne) RemoveChildren(p ...*PartyQualType) *PartyQualTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtuo.RemoveChildIDs(ids...)
}

// ClearChildPartyQualTypes clears all "child_party_qual_types" edges to the PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) ClearChildPartyQualTypes() *PartyQualTypeUpdateOne {
	pqtuo.mutation.ClearChildPartyQualTypes()
	return pqtuo
}

// RemoveChildPartyQualTypeIDs removes the "child_party_qual_types" edge to PartyQualType entities by IDs.
func (pqtuo *PartyQualTypeUpdateOne) RemoveChildPartyQualTypeIDs(ids ...int) *PartyQualTypeUpdateOne {
	pqtuo.mutation.RemoveChildPartyQualTypeIDs(ids...)
	return pqtuo
}

// RemoveChildPartyQualTypes removes "child_party_qual_types" edges to PartyQualType entities.
func (pqtuo *PartyQualTypeUpdateOne) RemoveChildPartyQualTypes(p ...*PartyQualType) *PartyQualTypeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtuo.RemoveChildPartyQualTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pqtuo *PartyQualTypeUpdateOne) Select(field string, fields ...string) *PartyQualTypeUpdateOne {
	pqtuo.fields = append([]string{field}, fields...)
	return pqtuo
}

// Save executes the query and returns the updated PartyQualType entity.
func (pqtuo *PartyQualTypeUpdateOne) Save(ctx context.Context) (*PartyQualType, error) {
	var (
		err  error
		node *PartyQualType
	)
	pqtuo.defaults()
	if len(pqtuo.hooks) == 0 {
		if err = pqtuo.check(); err != nil {
			return nil, err
		}
		node, err = pqtuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyQualTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pqtuo.check(); err != nil {
				return nil, err
			}
			pqtuo.mutation = mutation
			node, err = pqtuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pqtuo.hooks) - 1; i >= 0; i-- {
			mut = pqtuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pqtuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pqtuo *PartyQualTypeUpdateOne) SaveX(ctx context.Context) *PartyQualType {
	node, err := pqtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pqtuo *PartyQualTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := pqtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pqtuo *PartyQualTypeUpdateOne) ExecX(ctx context.Context) {
	if err := pqtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pqtuo *PartyQualTypeUpdateOne) defaults() {
	if _, ok := pqtuo.mutation.UpdateTime(); !ok {
		v := partyqualtype.UpdateDefaultUpdateTime()
		pqtuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pqtuo *PartyQualTypeUpdateOne) check() error {
	if v, ok := pqtuo.mutation.HasTable(); ok {
		if err := partyqualtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pqtuo *PartyQualTypeUpdateOne) sqlSave(ctx context.Context) (_node *PartyQualType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partyqualtype.Table,
			Columns: partyqualtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyqualtype.FieldID,
			},
		},
	}
	id, ok := pqtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PartyQualType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pqtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partyqualtype.FieldID)
		for _, f := range fields {
			if !partyqualtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != partyqualtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pqtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pqtuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyqualtype.FieldUpdateTime,
		})
	}
	if value, ok := pqtuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldStringRef,
		})
	}
	if pqtuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyqualtype.FieldStringRef,
		})
	}
	if value, ok := pqtuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: partyqualtype.FieldHasTable,
		})
	}
	if pqtuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: partyqualtype.FieldHasTable,
		})
	}
	if value, ok := pqtuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldDescription,
		})
	}
	if pqtuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyqualtype.FieldDescription,
		})
	}
	if pqtuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyqualtype.ParentTable,
			Columns: []string{partyqualtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyqualtype.ParentTable,
			Columns: []string{partyqualtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pqtuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pqtuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pqtuo.mutation.ChildPartyQualTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtuo.mutation.RemovedChildPartyQualTypesIDs(); len(nodes) > 0 && !pqtuo.mutation.ChildPartyQualTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pqtuo.mutation.ChildPartyQualTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PartyQualType{config: pqtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pqtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partyqualtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
