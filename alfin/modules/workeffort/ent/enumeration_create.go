// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumerationtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/fixedasset"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortpartyassignment"
)

// EnumerationCreate is the builder for creating a Enumeration entity.
type EnumerationCreate struct {
	config
	mutation *EnumerationMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ec *EnumerationCreate) SetCreateTime(t time.Time) *EnumerationCreate {
	ec.mutation.SetCreateTime(t)
	return ec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableCreateTime(t *time.Time) *EnumerationCreate {
	if t != nil {
		ec.SetCreateTime(*t)
	}
	return ec
}

// SetUpdateTime sets the "update_time" field.
func (ec *EnumerationCreate) SetUpdateTime(t time.Time) *EnumerationCreate {
	ec.mutation.SetUpdateTime(t)
	return ec
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableUpdateTime(t *time.Time) *EnumerationCreate {
	if t != nil {
		ec.SetUpdateTime(*t)
	}
	return ec
}

// SetStringRef sets the "string_ref" field.
func (ec *EnumerationCreate) SetStringRef(s string) *EnumerationCreate {
	ec.mutation.SetStringRef(s)
	return ec
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableStringRef(s *string) *EnumerationCreate {
	if s != nil {
		ec.SetStringRef(*s)
	}
	return ec
}

// SetEnumCode sets the "enum_code" field.
func (ec *EnumerationCreate) SetEnumCode(s string) *EnumerationCreate {
	ec.mutation.SetEnumCode(s)
	return ec
}

// SetNillableEnumCode sets the "enum_code" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableEnumCode(s *string) *EnumerationCreate {
	if s != nil {
		ec.SetEnumCode(*s)
	}
	return ec
}

// SetSequenceID sets the "sequence_id" field.
func (ec *EnumerationCreate) SetSequenceID(i int) *EnumerationCreate {
	ec.mutation.SetSequenceID(i)
	return ec
}

// SetNillableSequenceID sets the "sequence_id" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableSequenceID(i *int) *EnumerationCreate {
	if i != nil {
		ec.SetSequenceID(*i)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *EnumerationCreate) SetDescription(s string) *EnumerationCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EnumerationCreate) SetNillableDescription(s *string) *EnumerationCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetEnumerationTypeID sets the "enumeration_type" edge to the EnumerationType entity by ID.
func (ec *EnumerationCreate) SetEnumerationTypeID(id int) *EnumerationCreate {
	ec.mutation.SetEnumerationTypeID(id)
	return ec
}

// SetNillableEnumerationTypeID sets the "enumeration_type" edge to the EnumerationType entity by ID if the given value is not nil.
func (ec *EnumerationCreate) SetNillableEnumerationTypeID(id *int) *EnumerationCreate {
	if id != nil {
		ec = ec.SetEnumerationTypeID(*id)
	}
	return ec
}

// SetEnumerationType sets the "enumeration_type" edge to the EnumerationType entity.
func (ec *EnumerationCreate) SetEnumerationType(e *EnumerationType) *EnumerationCreate {
	return ec.SetEnumerationTypeID(e.ID)
}

// AddClassFixedAssetIDs adds the "class_fixed_assets" edge to the FixedAsset entity by IDs.
func (ec *EnumerationCreate) AddClassFixedAssetIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddClassFixedAssetIDs(ids...)
	return ec
}

// AddClassFixedAssets adds the "class_fixed_assets" edges to the FixedAsset entity.
func (ec *EnumerationCreate) AddClassFixedAssets(f ...*FixedAsset) *EnumerationCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ec.AddClassFixedAssetIDs(ids...)
}

// AddEmploymentStatusPersonIDs adds the "employment_status_people" edge to the Person entity by IDs.
func (ec *EnumerationCreate) AddEmploymentStatusPersonIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddEmploymentStatusPersonIDs(ids...)
	return ec
}

// AddEmploymentStatusPeople adds the "employment_status_people" edges to the Person entity.
func (ec *EnumerationCreate) AddEmploymentStatusPeople(p ...*Person) *EnumerationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddEmploymentStatusPersonIDs(ids...)
}

// AddResidenceStatusPersonIDs adds the "residence_status_people" edge to the Person entity by IDs.
func (ec *EnumerationCreate) AddResidenceStatusPersonIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddResidenceStatusPersonIDs(ids...)
	return ec
}

// AddResidenceStatusPeople adds the "residence_status_people" edges to the Person entity.
func (ec *EnumerationCreate) AddResidenceStatusPeople(p ...*Person) *EnumerationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddResidenceStatusPersonIDs(ids...)
}

// AddMaritalStatusPersonIDs adds the "marital_status_people" edge to the Person entity by IDs.
func (ec *EnumerationCreate) AddMaritalStatusPersonIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddMaritalStatusPersonIDs(ids...)
	return ec
}

// AddMaritalStatusPeople adds the "marital_status_people" edges to the Person entity.
func (ec *EnumerationCreate) AddMaritalStatusPeople(p ...*Person) *EnumerationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ec.AddMaritalStatusPersonIDs(ids...)
}

// AddScopeWorkEffortIDs adds the "scope_work_efforts" edge to the WorkEffort entity by IDs.
func (ec *EnumerationCreate) AddScopeWorkEffortIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddScopeWorkEffortIDs(ids...)
	return ec
}

// AddScopeWorkEfforts adds the "scope_work_efforts" edges to the WorkEffort entity.
func (ec *EnumerationCreate) AddScopeWorkEfforts(w ...*WorkEffort) *EnumerationCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddScopeWorkEffortIDs(ids...)
}

// AddExpectationWorkEffortPartyAssignmentIDs adds the "expectation_work_effort_party_assignments" edge to the WorkEffortPartyAssignment entity by IDs.
func (ec *EnumerationCreate) AddExpectationWorkEffortPartyAssignmentIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddExpectationWorkEffortPartyAssignmentIDs(ids...)
	return ec
}

// AddExpectationWorkEffortPartyAssignments adds the "expectation_work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (ec *EnumerationCreate) AddExpectationWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *EnumerationCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddExpectationWorkEffortPartyAssignmentIDs(ids...)
}

// AddDelegateReasonWorkEffortPartyAssignmentIDs adds the "delegate_reason_work_effort_party_assignments" edge to the WorkEffortPartyAssignment entity by IDs.
func (ec *EnumerationCreate) AddDelegateReasonWorkEffortPartyAssignmentIDs(ids ...int) *EnumerationCreate {
	ec.mutation.AddDelegateReasonWorkEffortPartyAssignmentIDs(ids...)
	return ec
}

// AddDelegateReasonWorkEffortPartyAssignments adds the "delegate_reason_work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (ec *EnumerationCreate) AddDelegateReasonWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *EnumerationCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddDelegateReasonWorkEffortPartyAssignmentIDs(ids...)
}

// Mutation returns the EnumerationMutation object of the builder.
func (ec *EnumerationCreate) Mutation() *EnumerationMutation {
	return ec.mutation
}

// Save creates the Enumeration in the database.
func (ec *EnumerationCreate) Save(ctx context.Context) (*Enumeration, error) {
	var (
		err  error
		node *Enumeration
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EnumerationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EnumerationCreate) SaveX(ctx context.Context) *Enumeration {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ec *EnumerationCreate) defaults() {
	if _, ok := ec.mutation.CreateTime(); !ok {
		v := enumeration.DefaultCreateTime()
		ec.mutation.SetCreateTime(v)
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		v := enumeration.DefaultUpdateTime()
		ec.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EnumerationCreate) check() error {
	if _, ok := ec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	return nil
}

func (ec *EnumerationCreate) sqlSave(ctx context.Context) (*Enumeration, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ec *EnumerationCreate) createSpec() (*Enumeration, *sqlgraph.CreateSpec) {
	var (
		_node = &Enumeration{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: enumeration.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enumeration.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enumeration.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ec.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enumeration.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ec.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enumeration.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := ec.mutation.EnumCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enumeration.FieldEnumCode,
		})
		_node.EnumCode = value
	}
	if value, ok := ec.mutation.SequenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: enumeration.FieldSequenceID,
		})
		_node.SequenceID = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: enumeration.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := ec.mutation.EnumerationTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   enumeration.EnumerationTypeTable,
			Columns: []string{enumeration.EnumerationTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumerationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_type_enumerations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ClassFixedAssetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.ClassFixedAssetsTable,
			Columns: []string{enumeration.ClassFixedAssetsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EmploymentStatusPeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.EmploymentStatusPeopleTable,
			Columns: []string{enumeration.EmploymentStatusPeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ResidenceStatusPeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.ResidenceStatusPeopleTable,
			Columns: []string{enumeration.ResidenceStatusPeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.MaritalStatusPeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.MaritalStatusPeopleTable,
			Columns: []string{enumeration.MaritalStatusPeopleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ScopeWorkEffortsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.ScopeWorkEffortsTable,
			Columns: []string{enumeration.ScopeWorkEffortsColumn},
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
	if nodes := ec.mutation.ExpectationWorkEffortPartyAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.ExpectationWorkEffortPartyAssignmentsTable,
			Columns: []string{enumeration.ExpectationWorkEffortPartyAssignmentsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.DelegateReasonWorkEffortPartyAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   enumeration.DelegateReasonWorkEffortPartyAssignmentsTable,
			Columns: []string{enumeration.DelegateReasonWorkEffortPartyAssignmentsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EnumerationCreateBulk is the builder for creating many Enumeration entities in bulk.
type EnumerationCreateBulk struct {
	config
	builders []*EnumerationCreate
}

// Save creates the Enumeration entities in the database.
func (ecb *EnumerationCreateBulk) Save(ctx context.Context) ([]*Enumeration, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Enumeration, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnumerationMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EnumerationCreateBulk) SaveX(ctx context.Context) []*Enumeration {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
