// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontactmech"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/roletype"
)

// PartyContactMechCreate is the builder for creating a PartyContactMech entity.
type PartyContactMechCreate struct {
	config
	mutation *PartyContactMechMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pcmc *PartyContactMechCreate) SetCreateTime(t time.Time) *PartyContactMechCreate {
	pcmc.mutation.SetCreateTime(t)
	return pcmc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableCreateTime(t *time.Time) *PartyContactMechCreate {
	if t != nil {
		pcmc.SetCreateTime(*t)
	}
	return pcmc
}

// SetUpdateTime sets the "update_time" field.
func (pcmc *PartyContactMechCreate) SetUpdateTime(t time.Time) *PartyContactMechCreate {
	pcmc.mutation.SetUpdateTime(t)
	return pcmc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableUpdateTime(t *time.Time) *PartyContactMechCreate {
	if t != nil {
		pcmc.SetUpdateTime(*t)
	}
	return pcmc
}

// SetStringRef sets the "string_ref" field.
func (pcmc *PartyContactMechCreate) SetStringRef(s string) *PartyContactMechCreate {
	pcmc.mutation.SetStringRef(s)
	return pcmc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableStringRef(s *string) *PartyContactMechCreate {
	if s != nil {
		pcmc.SetStringRef(*s)
	}
	return pcmc
}

// SetContactMechID sets the "contact_mech_id" field.
func (pcmc *PartyContactMechCreate) SetContactMechID(i int) *PartyContactMechCreate {
	pcmc.mutation.SetContactMechID(i)
	return pcmc
}

// SetFromDate sets the "from_date" field.
func (pcmc *PartyContactMechCreate) SetFromDate(t time.Time) *PartyContactMechCreate {
	pcmc.mutation.SetFromDate(t)
	return pcmc
}

// SetNillableFromDate sets the "from_date" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableFromDate(t *time.Time) *PartyContactMechCreate {
	if t != nil {
		pcmc.SetFromDate(*t)
	}
	return pcmc
}

// SetThruDate sets the "thru_date" field.
func (pcmc *PartyContactMechCreate) SetThruDate(t time.Time) *PartyContactMechCreate {
	pcmc.mutation.SetThruDate(t)
	return pcmc
}

// SetNillableThruDate sets the "thru_date" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableThruDate(t *time.Time) *PartyContactMechCreate {
	if t != nil {
		pcmc.SetThruDate(*t)
	}
	return pcmc
}

// SetAllowSolicitation sets the "allow_solicitation" field.
func (pcmc *PartyContactMechCreate) SetAllowSolicitation(ps partycontactmech.AllowSolicitation) *PartyContactMechCreate {
	pcmc.mutation.SetAllowSolicitation(ps)
	return pcmc
}

// SetNillableAllowSolicitation sets the "allow_solicitation" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableAllowSolicitation(ps *partycontactmech.AllowSolicitation) *PartyContactMechCreate {
	if ps != nil {
		pcmc.SetAllowSolicitation(*ps)
	}
	return pcmc
}

// SetExtension sets the "extension" field.
func (pcmc *PartyContactMechCreate) SetExtension(s string) *PartyContactMechCreate {
	pcmc.mutation.SetExtension(s)
	return pcmc
}

// SetNillableExtension sets the "extension" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableExtension(s *string) *PartyContactMechCreate {
	if s != nil {
		pcmc.SetExtension(*s)
	}
	return pcmc
}

// SetVerified sets the "verified" field.
func (pcmc *PartyContactMechCreate) SetVerified(pa partycontactmech.Verified) *PartyContactMechCreate {
	pcmc.mutation.SetVerified(pa)
	return pcmc
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableVerified(pa *partycontactmech.Verified) *PartyContactMechCreate {
	if pa != nil {
		pcmc.SetVerified(*pa)
	}
	return pcmc
}

// SetComments sets the "comments" field.
func (pcmc *PartyContactMechCreate) SetComments(s string) *PartyContactMechCreate {
	pcmc.mutation.SetComments(s)
	return pcmc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableComments(s *string) *PartyContactMechCreate {
	if s != nil {
		pcmc.SetComments(*s)
	}
	return pcmc
}

// SetYearsWithContactMech sets the "years_with_contact_mech" field.
func (pcmc *PartyContactMechCreate) SetYearsWithContactMech(i int) *PartyContactMechCreate {
	pcmc.mutation.SetYearsWithContactMech(i)
	return pcmc
}

// SetNillableYearsWithContactMech sets the "years_with_contact_mech" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableYearsWithContactMech(i *int) *PartyContactMechCreate {
	if i != nil {
		pcmc.SetYearsWithContactMech(*i)
	}
	return pcmc
}

// SetMonthsWithContactMech sets the "months_with_contact_mech" field.
func (pcmc *PartyContactMechCreate) SetMonthsWithContactMech(i int) *PartyContactMechCreate {
	pcmc.mutation.SetMonthsWithContactMech(i)
	return pcmc
}

// SetNillableMonthsWithContactMech sets the "months_with_contact_mech" field if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableMonthsWithContactMech(i *int) *PartyContactMechCreate {
	if i != nil {
		pcmc.SetMonthsWithContactMech(*i)
	}
	return pcmc
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (pcmc *PartyContactMechCreate) SetPartyID(id int) *PartyContactMechCreate {
	pcmc.mutation.SetPartyID(id)
	return pcmc
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillablePartyID(id *int) *PartyContactMechCreate {
	if id != nil {
		pcmc = pcmc.SetPartyID(*id)
	}
	return pcmc
}

// SetParty sets the "party" edge to the Party entity.
func (pcmc *PartyContactMechCreate) SetParty(p *Party) *PartyContactMechCreate {
	return pcmc.SetPartyID(p.ID)
}

// SetPersonID sets the "person" edge to the Person entity by ID.
func (pcmc *PartyContactMechCreate) SetPersonID(id int) *PartyContactMechCreate {
	pcmc.mutation.SetPersonID(id)
	return pcmc
}

// SetNillablePersonID sets the "person" edge to the Person entity by ID if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillablePersonID(id *int) *PartyContactMechCreate {
	if id != nil {
		pcmc = pcmc.SetPersonID(*id)
	}
	return pcmc
}

// SetPerson sets the "person" edge to the Person entity.
func (pcmc *PartyContactMechCreate) SetPerson(p *Person) *PartyContactMechCreate {
	return pcmc.SetPersonID(p.ID)
}

// SetPartyRoleID sets the "party_role" edge to the PartyRole entity by ID.
func (pcmc *PartyContactMechCreate) SetPartyRoleID(id int) *PartyContactMechCreate {
	pcmc.mutation.SetPartyRoleID(id)
	return pcmc
}

// SetNillablePartyRoleID sets the "party_role" edge to the PartyRole entity by ID if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillablePartyRoleID(id *int) *PartyContactMechCreate {
	if id != nil {
		pcmc = pcmc.SetPartyRoleID(*id)
	}
	return pcmc
}

// SetPartyRole sets the "party_role" edge to the PartyRole entity.
func (pcmc *PartyContactMechCreate) SetPartyRole(p *PartyRole) *PartyContactMechCreate {
	return pcmc.SetPartyRoleID(p.ID)
}

// SetRoleTypeID sets the "role_type" edge to the RoleType entity by ID.
func (pcmc *PartyContactMechCreate) SetRoleTypeID(id int) *PartyContactMechCreate {
	pcmc.mutation.SetRoleTypeID(id)
	return pcmc
}

// SetNillableRoleTypeID sets the "role_type" edge to the RoleType entity by ID if the given value is not nil.
func (pcmc *PartyContactMechCreate) SetNillableRoleTypeID(id *int) *PartyContactMechCreate {
	if id != nil {
		pcmc = pcmc.SetRoleTypeID(*id)
	}
	return pcmc
}

// SetRoleType sets the "role_type" edge to the RoleType entity.
func (pcmc *PartyContactMechCreate) SetRoleType(r *RoleType) *PartyContactMechCreate {
	return pcmc.SetRoleTypeID(r.ID)
}

// Mutation returns the PartyContactMechMutation object of the builder.
func (pcmc *PartyContactMechCreate) Mutation() *PartyContactMechMutation {
	return pcmc.mutation
}

// Save creates the PartyContactMech in the database.
func (pcmc *PartyContactMechCreate) Save(ctx context.Context) (*PartyContactMech, error) {
	var (
		err  error
		node *PartyContactMech
	)
	pcmc.defaults()
	if len(pcmc.hooks) == 0 {
		if err = pcmc.check(); err != nil {
			return nil, err
		}
		node, err = pcmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyContactMechMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcmc.check(); err != nil {
				return nil, err
			}
			pcmc.mutation = mutation
			if node, err = pcmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcmc.hooks) - 1; i >= 0; i-- {
			mut = pcmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcmc *PartyContactMechCreate) SaveX(ctx context.Context) *PartyContactMech {
	v, err := pcmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pcmc *PartyContactMechCreate) defaults() {
	if _, ok := pcmc.mutation.CreateTime(); !ok {
		v := partycontactmech.DefaultCreateTime()
		pcmc.mutation.SetCreateTime(v)
	}
	if _, ok := pcmc.mutation.UpdateTime(); !ok {
		v := partycontactmech.DefaultUpdateTime()
		pcmc.mutation.SetUpdateTime(v)
	}
	if _, ok := pcmc.mutation.FromDate(); !ok {
		v := partycontactmech.DefaultFromDate()
		pcmc.mutation.SetFromDate(v)
	}
	if _, ok := pcmc.mutation.ThruDate(); !ok {
		v := partycontactmech.DefaultThruDate()
		pcmc.mutation.SetThruDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcmc *PartyContactMechCreate) check() error {
	if _, ok := pcmc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pcmc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := pcmc.mutation.ContactMechID(); !ok {
		return &ValidationError{Name: "contact_mech_id", err: errors.New("ent: missing required field \"contact_mech_id\"")}
	}
	if _, ok := pcmc.mutation.FromDate(); !ok {
		return &ValidationError{Name: "from_date", err: errors.New("ent: missing required field \"from_date\"")}
	}
	if v, ok := pcmc.mutation.AllowSolicitation(); ok {
		if err := partycontactmech.AllowSolicitationValidator(v); err != nil {
			return &ValidationError{Name: "allow_solicitation", err: fmt.Errorf("ent: validator failed for field \"allow_solicitation\": %w", err)}
		}
	}
	if v, ok := pcmc.mutation.Verified(); ok {
		if err := partycontactmech.VerifiedValidator(v); err != nil {
			return &ValidationError{Name: "verified", err: fmt.Errorf("ent: validator failed for field \"verified\": %w", err)}
		}
	}
	return nil
}

func (pcmc *PartyContactMechCreate) sqlSave(ctx context.Context) (*PartyContactMech, error) {
	_node, _spec := pcmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcmc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pcmc *PartyContactMechCreate) createSpec() (*PartyContactMech, *sqlgraph.CreateSpec) {
	var (
		_node = &PartyContactMech{config: pcmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: partycontactmech.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partycontactmech.FieldID,
			},
		}
	)
	if value, ok := pcmc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partycontactmech.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pcmc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partycontactmech.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pcmc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partycontactmech.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pcmc.mutation.ContactMechID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: partycontactmech.FieldContactMechID,
		})
		_node.ContactMechID = value
	}
	if value, ok := pcmc.mutation.FromDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partycontactmech.FieldFromDate,
		})
		_node.FromDate = value
	}
	if value, ok := pcmc.mutation.ThruDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partycontactmech.FieldThruDate,
		})
		_node.ThruDate = value
	}
	if value, ok := pcmc.mutation.AllowSolicitation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: partycontactmech.FieldAllowSolicitation,
		})
		_node.AllowSolicitation = value
	}
	if value, ok := pcmc.mutation.Extension(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partycontactmech.FieldExtension,
		})
		_node.Extension = value
	}
	if value, ok := pcmc.mutation.Verified(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: partycontactmech.FieldVerified,
		})
		_node.Verified = value
	}
	if value, ok := pcmc.mutation.Comments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partycontactmech.FieldComments,
		})
		_node.Comments = value
	}
	if value, ok := pcmc.mutation.YearsWithContactMech(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: partycontactmech.FieldYearsWithContactMech,
		})
		_node.YearsWithContactMech = value
	}
	if value, ok := pcmc.mutation.MonthsWithContactMech(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: partycontactmech.FieldMonthsWithContactMech,
		})
		_node.MonthsWithContactMech = value
	}
	if nodes := pcmc.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partycontactmech.PartyTable,
			Columns: []string{partycontactmech.PartyColumn},
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
		_node.party_party_contact_meches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcmc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partycontactmech.PersonTable,
			Columns: []string{partycontactmech.PersonColumn},
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
		_node.person_party_contact_meches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcmc.mutation.PartyRoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partycontactmech.PartyRoleTable,
			Columns: []string{partycontactmech.PartyRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.party_role_party_contact_meches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcmc.mutation.RoleTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partycontactmech.RoleTypeTable,
			Columns: []string{partycontactmech.RoleTypeColumn},
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
		_node.role_type_party_contact_meches = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PartyContactMechCreateBulk is the builder for creating many PartyContactMech entities in bulk.
type PartyContactMechCreateBulk struct {
	config
	builders []*PartyContactMechCreate
}

// Save creates the PartyContactMech entities in the database.
func (pcmcb *PartyContactMechCreateBulk) Save(ctx context.Context) ([]*PartyContactMech, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcmcb.builders))
	nodes := make([]*PartyContactMech, len(pcmcb.builders))
	mutators := make([]Mutator, len(pcmcb.builders))
	for i := range pcmcb.builders {
		func(i int, root context.Context) {
			builder := pcmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartyContactMechMutation)
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
					_, err = mutators[i+1].Mutate(root, pcmcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcmcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcmcb *PartyContactMechCreateBulk) SaveX(ctx context.Context) []*PartyContactMech {
	v, err := pcmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
