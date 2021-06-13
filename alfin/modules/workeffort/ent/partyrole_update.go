// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/fixedasset"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontactmech"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/roletype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortpartyassignment"
)

// PartyRoleUpdate is the builder for updating PartyRole entities.
type PartyRoleUpdate struct {
	config
	hooks    []Hook
	mutation *PartyRoleMutation
}

// Where adds a new predicate for the PartyRoleUpdate builder.
func (pru *PartyRoleUpdate) Where(ps ...predicate.PartyRole) *PartyRoleUpdate {
	pru.mutation.predicates = append(pru.mutation.predicates, ps...)
	return pru
}

// SetStringRef sets the "string_ref" field.
func (pru *PartyRoleUpdate) SetStringRef(s string) *PartyRoleUpdate {
	pru.mutation.SetStringRef(s)
	return pru
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pru *PartyRoleUpdate) SetNillableStringRef(s *string) *PartyRoleUpdate {
	if s != nil {
		pru.SetStringRef(*s)
	}
	return pru
}

// ClearStringRef clears the value of the "string_ref" field.
func (pru *PartyRoleUpdate) ClearStringRef() *PartyRoleUpdate {
	pru.mutation.ClearStringRef()
	return pru
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (pru *PartyRoleUpdate) SetPartyID(id int) *PartyRoleUpdate {
	pru.mutation.SetPartyID(id)
	return pru
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (pru *PartyRoleUpdate) SetNillablePartyID(id *int) *PartyRoleUpdate {
	if id != nil {
		pru = pru.SetPartyID(*id)
	}
	return pru
}

// SetParty sets the "party" edge to the Party entity.
func (pru *PartyRoleUpdate) SetParty(p *Party) *PartyRoleUpdate {
	return pru.SetPartyID(p.ID)
}

// SetRoleTypeID sets the "role_type" edge to the RoleType entity by ID.
func (pru *PartyRoleUpdate) SetRoleTypeID(id int) *PartyRoleUpdate {
	pru.mutation.SetRoleTypeID(id)
	return pru
}

// SetNillableRoleTypeID sets the "role_type" edge to the RoleType entity by ID if the given value is not nil.
func (pru *PartyRoleUpdate) SetNillableRoleTypeID(id *int) *PartyRoleUpdate {
	if id != nil {
		pru = pru.SetRoleTypeID(*id)
	}
	return pru
}

// SetRoleType sets the "role_type" edge to the RoleType entity.
func (pru *PartyRoleUpdate) SetRoleType(r *RoleType) *PartyRoleUpdate {
	return pru.SetRoleTypeID(r.ID)
}

// AddFixedAssetIDs adds the "fixed_assets" edge to the FixedAsset entity by IDs.
func (pru *PartyRoleUpdate) AddFixedAssetIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.AddFixedAssetIDs(ids...)
	return pru
}

// AddFixedAssets adds the "fixed_assets" edges to the FixedAsset entity.
func (pru *PartyRoleUpdate) AddFixedAssets(f ...*FixedAsset) *PartyRoleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pru.AddFixedAssetIDs(ids...)
}

// AddPartyContactMechIDs adds the "party_contact_meches" edge to the PartyContactMech entity by IDs.
func (pru *PartyRoleUpdate) AddPartyContactMechIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.AddPartyContactMechIDs(ids...)
	return pru
}

// AddPartyContactMeches adds the "party_contact_meches" edges to the PartyContactMech entity.
func (pru *PartyRoleUpdate) AddPartyContactMeches(p ...*PartyContactMech) *PartyRoleUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pru.AddPartyContactMechIDs(ids...)
}

// AddWorkEffortPartyAssignmentIDs adds the "work_effort_party_assignments" edge to the WorkEffortPartyAssignment entity by IDs.
func (pru *PartyRoleUpdate) AddWorkEffortPartyAssignmentIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.AddWorkEffortPartyAssignmentIDs(ids...)
	return pru
}

// AddWorkEffortPartyAssignments adds the "work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (pru *PartyRoleUpdate) AddWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *PartyRoleUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pru.AddWorkEffortPartyAssignmentIDs(ids...)
}

// Mutation returns the PartyRoleMutation object of the builder.
func (pru *PartyRoleUpdate) Mutation() *PartyRoleMutation {
	return pru.mutation
}

// ClearParty clears the "party" edge to the Party entity.
func (pru *PartyRoleUpdate) ClearParty() *PartyRoleUpdate {
	pru.mutation.ClearParty()
	return pru
}

// ClearRoleType clears the "role_type" edge to the RoleType entity.
func (pru *PartyRoleUpdate) ClearRoleType() *PartyRoleUpdate {
	pru.mutation.ClearRoleType()
	return pru
}

// ClearFixedAssets clears all "fixed_assets" edges to the FixedAsset entity.
func (pru *PartyRoleUpdate) ClearFixedAssets() *PartyRoleUpdate {
	pru.mutation.ClearFixedAssets()
	return pru
}

// RemoveFixedAssetIDs removes the "fixed_assets" edge to FixedAsset entities by IDs.
func (pru *PartyRoleUpdate) RemoveFixedAssetIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.RemoveFixedAssetIDs(ids...)
	return pru
}

// RemoveFixedAssets removes "fixed_assets" edges to FixedAsset entities.
func (pru *PartyRoleUpdate) RemoveFixedAssets(f ...*FixedAsset) *PartyRoleUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pru.RemoveFixedAssetIDs(ids...)
}

// ClearPartyContactMeches clears all "party_contact_meches" edges to the PartyContactMech entity.
func (pru *PartyRoleUpdate) ClearPartyContactMeches() *PartyRoleUpdate {
	pru.mutation.ClearPartyContactMeches()
	return pru
}

// RemovePartyContactMechIDs removes the "party_contact_meches" edge to PartyContactMech entities by IDs.
func (pru *PartyRoleUpdate) RemovePartyContactMechIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.RemovePartyContactMechIDs(ids...)
	return pru
}

// RemovePartyContactMeches removes "party_contact_meches" edges to PartyContactMech entities.
func (pru *PartyRoleUpdate) RemovePartyContactMeches(p ...*PartyContactMech) *PartyRoleUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pru.RemovePartyContactMechIDs(ids...)
}

// ClearWorkEffortPartyAssignments clears all "work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (pru *PartyRoleUpdate) ClearWorkEffortPartyAssignments() *PartyRoleUpdate {
	pru.mutation.ClearWorkEffortPartyAssignments()
	return pru
}

// RemoveWorkEffortPartyAssignmentIDs removes the "work_effort_party_assignments" edge to WorkEffortPartyAssignment entities by IDs.
func (pru *PartyRoleUpdate) RemoveWorkEffortPartyAssignmentIDs(ids ...int) *PartyRoleUpdate {
	pru.mutation.RemoveWorkEffortPartyAssignmentIDs(ids...)
	return pru
}

// RemoveWorkEffortPartyAssignments removes "work_effort_party_assignments" edges to WorkEffortPartyAssignment entities.
func (pru *PartyRoleUpdate) RemoveWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *PartyRoleUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pru.RemoveWorkEffortPartyAssignmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *PartyRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pru.defaults()
	if len(pru.hooks) == 0 {
		affected, err = pru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pru.mutation = mutation
			affected, err = pru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pru.hooks) - 1; i >= 0; i-- {
			mut = pru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pru *PartyRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *PartyRoleUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *PartyRoleUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pru *PartyRoleUpdate) defaults() {
	if _, ok := pru.mutation.UpdateTime(); !ok {
		v := partyrole.UpdateDefaultUpdateTime()
		pru.mutation.SetUpdateTime(v)
	}
}

func (pru *PartyRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partyrole.Table,
			Columns: partyrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyrole.FieldID,
			},
		},
	}
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyrole.FieldUpdateTime,
		})
	}
	if value, ok := pru.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyrole.FieldStringRef,
		})
	}
	if pru.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyrole.FieldStringRef,
		})
	}
	if pru.mutation.PartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.PartyTable,
			Columns: []string{partyrole.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.PartyTable,
			Columns: []string{partyrole.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.RoleTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.RoleTypeTable,
			Columns: []string{partyrole.RoleTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roletype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RoleTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.RoleTypeTable,
			Columns: []string{partyrole.RoleTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roletype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.FixedAssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedFixedAssetsIDs(); len(nodes) > 0 && !pru.mutation.FixedAssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.FixedAssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.PartyContactMechesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedPartyContactMechesIDs(); len(nodes) > 0 && !pru.mutation.PartyContactMechesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.PartyContactMechesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pru.mutation.WorkEffortPartyAssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.RemovedWorkEffortPartyAssignmentsIDs(); len(nodes) > 0 && !pru.mutation.WorkEffortPartyAssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.WorkEffortPartyAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partyrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PartyRoleUpdateOne is the builder for updating a single PartyRole entity.
type PartyRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PartyRoleMutation
}

// SetStringRef sets the "string_ref" field.
func (pruo *PartyRoleUpdateOne) SetStringRef(s string) *PartyRoleUpdateOne {
	pruo.mutation.SetStringRef(s)
	return pruo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pruo *PartyRoleUpdateOne) SetNillableStringRef(s *string) *PartyRoleUpdateOne {
	if s != nil {
		pruo.SetStringRef(*s)
	}
	return pruo
}

// ClearStringRef clears the value of the "string_ref" field.
func (pruo *PartyRoleUpdateOne) ClearStringRef() *PartyRoleUpdateOne {
	pruo.mutation.ClearStringRef()
	return pruo
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (pruo *PartyRoleUpdateOne) SetPartyID(id int) *PartyRoleUpdateOne {
	pruo.mutation.SetPartyID(id)
	return pruo
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (pruo *PartyRoleUpdateOne) SetNillablePartyID(id *int) *PartyRoleUpdateOne {
	if id != nil {
		pruo = pruo.SetPartyID(*id)
	}
	return pruo
}

// SetParty sets the "party" edge to the Party entity.
func (pruo *PartyRoleUpdateOne) SetParty(p *Party) *PartyRoleUpdateOne {
	return pruo.SetPartyID(p.ID)
}

// SetRoleTypeID sets the "role_type" edge to the RoleType entity by ID.
func (pruo *PartyRoleUpdateOne) SetRoleTypeID(id int) *PartyRoleUpdateOne {
	pruo.mutation.SetRoleTypeID(id)
	return pruo
}

// SetNillableRoleTypeID sets the "role_type" edge to the RoleType entity by ID if the given value is not nil.
func (pruo *PartyRoleUpdateOne) SetNillableRoleTypeID(id *int) *PartyRoleUpdateOne {
	if id != nil {
		pruo = pruo.SetRoleTypeID(*id)
	}
	return pruo
}

// SetRoleType sets the "role_type" edge to the RoleType entity.
func (pruo *PartyRoleUpdateOne) SetRoleType(r *RoleType) *PartyRoleUpdateOne {
	return pruo.SetRoleTypeID(r.ID)
}

// AddFixedAssetIDs adds the "fixed_assets" edge to the FixedAsset entity by IDs.
func (pruo *PartyRoleUpdateOne) AddFixedAssetIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.AddFixedAssetIDs(ids...)
	return pruo
}

// AddFixedAssets adds the "fixed_assets" edges to the FixedAsset entity.
func (pruo *PartyRoleUpdateOne) AddFixedAssets(f ...*FixedAsset) *PartyRoleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pruo.AddFixedAssetIDs(ids...)
}

// AddPartyContactMechIDs adds the "party_contact_meches" edge to the PartyContactMech entity by IDs.
func (pruo *PartyRoleUpdateOne) AddPartyContactMechIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.AddPartyContactMechIDs(ids...)
	return pruo
}

// AddPartyContactMeches adds the "party_contact_meches" edges to the PartyContactMech entity.
func (pruo *PartyRoleUpdateOne) AddPartyContactMeches(p ...*PartyContactMech) *PartyRoleUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pruo.AddPartyContactMechIDs(ids...)
}

// AddWorkEffortPartyAssignmentIDs adds the "work_effort_party_assignments" edge to the WorkEffortPartyAssignment entity by IDs.
func (pruo *PartyRoleUpdateOne) AddWorkEffortPartyAssignmentIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.AddWorkEffortPartyAssignmentIDs(ids...)
	return pruo
}

// AddWorkEffortPartyAssignments adds the "work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (pruo *PartyRoleUpdateOne) AddWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *PartyRoleUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pruo.AddWorkEffortPartyAssignmentIDs(ids...)
}

// Mutation returns the PartyRoleMutation object of the builder.
func (pruo *PartyRoleUpdateOne) Mutation() *PartyRoleMutation {
	return pruo.mutation
}

// ClearParty clears the "party" edge to the Party entity.
func (pruo *PartyRoleUpdateOne) ClearParty() *PartyRoleUpdateOne {
	pruo.mutation.ClearParty()
	return pruo
}

// ClearRoleType clears the "role_type" edge to the RoleType entity.
func (pruo *PartyRoleUpdateOne) ClearRoleType() *PartyRoleUpdateOne {
	pruo.mutation.ClearRoleType()
	return pruo
}

// ClearFixedAssets clears all "fixed_assets" edges to the FixedAsset entity.
func (pruo *PartyRoleUpdateOne) ClearFixedAssets() *PartyRoleUpdateOne {
	pruo.mutation.ClearFixedAssets()
	return pruo
}

// RemoveFixedAssetIDs removes the "fixed_assets" edge to FixedAsset entities by IDs.
func (pruo *PartyRoleUpdateOne) RemoveFixedAssetIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.RemoveFixedAssetIDs(ids...)
	return pruo
}

// RemoveFixedAssets removes "fixed_assets" edges to FixedAsset entities.
func (pruo *PartyRoleUpdateOne) RemoveFixedAssets(f ...*FixedAsset) *PartyRoleUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pruo.RemoveFixedAssetIDs(ids...)
}

// ClearPartyContactMeches clears all "party_contact_meches" edges to the PartyContactMech entity.
func (pruo *PartyRoleUpdateOne) ClearPartyContactMeches() *PartyRoleUpdateOne {
	pruo.mutation.ClearPartyContactMeches()
	return pruo
}

// RemovePartyContactMechIDs removes the "party_contact_meches" edge to PartyContactMech entities by IDs.
func (pruo *PartyRoleUpdateOne) RemovePartyContactMechIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.RemovePartyContactMechIDs(ids...)
	return pruo
}

// RemovePartyContactMeches removes "party_contact_meches" edges to PartyContactMech entities.
func (pruo *PartyRoleUpdateOne) RemovePartyContactMeches(p ...*PartyContactMech) *PartyRoleUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pruo.RemovePartyContactMechIDs(ids...)
}

// ClearWorkEffortPartyAssignments clears all "work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (pruo *PartyRoleUpdateOne) ClearWorkEffortPartyAssignments() *PartyRoleUpdateOne {
	pruo.mutation.ClearWorkEffortPartyAssignments()
	return pruo
}

// RemoveWorkEffortPartyAssignmentIDs removes the "work_effort_party_assignments" edge to WorkEffortPartyAssignment entities by IDs.
func (pruo *PartyRoleUpdateOne) RemoveWorkEffortPartyAssignmentIDs(ids ...int) *PartyRoleUpdateOne {
	pruo.mutation.RemoveWorkEffortPartyAssignmentIDs(ids...)
	return pruo
}

// RemoveWorkEffortPartyAssignments removes "work_effort_party_assignments" edges to WorkEffortPartyAssignment entities.
func (pruo *PartyRoleUpdateOne) RemoveWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *PartyRoleUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pruo.RemoveWorkEffortPartyAssignmentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *PartyRoleUpdateOne) Select(field string, fields ...string) *PartyRoleUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated PartyRole entity.
func (pruo *PartyRoleUpdateOne) Save(ctx context.Context) (*PartyRole, error) {
	var (
		err  error
		node *PartyRole
	)
	pruo.defaults()
	if len(pruo.hooks) == 0 {
		node, err = pruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pruo.mutation = mutation
			node, err = pruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pruo.hooks) - 1; i >= 0; i-- {
			mut = pruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *PartyRoleUpdateOne) SaveX(ctx context.Context) *PartyRole {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *PartyRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *PartyRoleUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pruo *PartyRoleUpdateOne) defaults() {
	if _, ok := pruo.mutation.UpdateTime(); !ok {
		v := partyrole.UpdateDefaultUpdateTime()
		pruo.mutation.SetUpdateTime(v)
	}
}

func (pruo *PartyRoleUpdateOne) sqlSave(ctx context.Context) (_node *PartyRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partyrole.Table,
			Columns: partyrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyrole.FieldID,
			},
		},
	}
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PartyRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partyrole.FieldID)
		for _, f := range fields {
			if !partyrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != partyrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyrole.FieldUpdateTime,
		})
	}
	if value, ok := pruo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyrole.FieldStringRef,
		})
	}
	if pruo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partyrole.FieldStringRef,
		})
	}
	if pruo.mutation.PartyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.PartyTable,
			Columns: []string{partyrole.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.PartyTable,
			Columns: []string{partyrole.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.RoleTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.RoleTypeTable,
			Columns: []string{partyrole.RoleTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roletype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RoleTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyrole.RoleTypeTable,
			Columns: []string{partyrole.RoleTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roletype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.FixedAssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fixedasset.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedFixedAssetsIDs(); len(nodes) > 0 && !pruo.mutation.FixedAssetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.FixedAssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.FixedAssetsTable,
			Columns: []string{partyrole.FixedAssetsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.PartyContactMechesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedPartyContactMechesIDs(); len(nodes) > 0 && !pruo.mutation.PartyContactMechesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.PartyContactMechesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.PartyContactMechesTable,
			Columns: []string{partyrole.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pruo.mutation.WorkEffortPartyAssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.RemovedWorkEffortPartyAssignmentsIDs(); len(nodes) > 0 && !pruo.mutation.WorkEffortPartyAssignmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.WorkEffortPartyAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyrole.WorkEffortPartyAssignmentsTable,
			Columns: []string{partyrole.WorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PartyRole{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partyrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
