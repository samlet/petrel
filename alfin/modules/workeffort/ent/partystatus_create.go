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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partystatus"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
)

// PartyStatusCreate is the builder for creating a PartyStatus entity.
type PartyStatusCreate struct {
	config
	mutation *PartyStatusMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (psc *PartyStatusCreate) SetCreateTime(t time.Time) *PartyStatusCreate {
	psc.mutation.SetCreateTime(t)
	return psc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableCreateTime(t *time.Time) *PartyStatusCreate {
	if t != nil {
		psc.SetCreateTime(*t)
	}
	return psc
}

// SetUpdateTime sets the "update_time" field.
func (psc *PartyStatusCreate) SetUpdateTime(t time.Time) *PartyStatusCreate {
	psc.mutation.SetUpdateTime(t)
	return psc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableUpdateTime(t *time.Time) *PartyStatusCreate {
	if t != nil {
		psc.SetUpdateTime(*t)
	}
	return psc
}

// SetStringRef sets the "string_ref" field.
func (psc *PartyStatusCreate) SetStringRef(s string) *PartyStatusCreate {
	psc.mutation.SetStringRef(s)
	return psc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableStringRef(s *string) *PartyStatusCreate {
	if s != nil {
		psc.SetStringRef(*s)
	}
	return psc
}

// SetStatusDate sets the "status_date" field.
func (psc *PartyStatusCreate) SetStatusDate(t time.Time) *PartyStatusCreate {
	psc.mutation.SetStatusDate(t)
	return psc
}

// SetNillableStatusDate sets the "status_date" field if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableStatusDate(t *time.Time) *PartyStatusCreate {
	if t != nil {
		psc.SetStatusDate(*t)
	}
	return psc
}

// SetStatusItemID sets the "status_item" edge to the StatusItem entity by ID.
func (psc *PartyStatusCreate) SetStatusItemID(id int) *PartyStatusCreate {
	psc.mutation.SetStatusItemID(id)
	return psc
}

// SetNillableStatusItemID sets the "status_item" edge to the StatusItem entity by ID if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableStatusItemID(id *int) *PartyStatusCreate {
	if id != nil {
		psc = psc.SetStatusItemID(*id)
	}
	return psc
}

// SetStatusItem sets the "status_item" edge to the StatusItem entity.
func (psc *PartyStatusCreate) SetStatusItem(s *StatusItem) *PartyStatusCreate {
	return psc.SetStatusItemID(s.ID)
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (psc *PartyStatusCreate) SetPartyID(id int) *PartyStatusCreate {
	psc.mutation.SetPartyID(id)
	return psc
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (psc *PartyStatusCreate) SetNillablePartyID(id *int) *PartyStatusCreate {
	if id != nil {
		psc = psc.SetPartyID(*id)
	}
	return psc
}

// SetParty sets the "party" edge to the Party entity.
func (psc *PartyStatusCreate) SetParty(p *Party) *PartyStatusCreate {
	return psc.SetPartyID(p.ID)
}

// SetChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID.
func (psc *PartyStatusCreate) SetChangeByUserLoginID(id int) *PartyStatusCreate {
	psc.mutation.SetChangeByUserLoginID(id)
	return psc
}

// SetNillableChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID if the given value is not nil.
func (psc *PartyStatusCreate) SetNillableChangeByUserLoginID(id *int) *PartyStatusCreate {
	if id != nil {
		psc = psc.SetChangeByUserLoginID(*id)
	}
	return psc
}

// SetChangeByUserLogin sets the "change_by_user_login" edge to the UserLogin entity.
func (psc *PartyStatusCreate) SetChangeByUserLogin(u *UserLogin) *PartyStatusCreate {
	return psc.SetChangeByUserLoginID(u.ID)
}

// Mutation returns the PartyStatusMutation object of the builder.
func (psc *PartyStatusCreate) Mutation() *PartyStatusMutation {
	return psc.mutation
}

// Save creates the PartyStatus in the database.
func (psc *PartyStatusCreate) Save(ctx context.Context) (*PartyStatus, error) {
	var (
		err  error
		node *PartyStatus
	)
	psc.defaults()
	if len(psc.hooks) == 0 {
		if err = psc.check(); err != nil {
			return nil, err
		}
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psc.check(); err != nil {
				return nil, err
			}
			psc.mutation = mutation
			if node, err = psc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PartyStatusCreate) SaveX(ctx context.Context) *PartyStatus {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (psc *PartyStatusCreate) defaults() {
	if _, ok := psc.mutation.CreateTime(); !ok {
		v := partystatus.DefaultCreateTime()
		psc.mutation.SetCreateTime(v)
	}
	if _, ok := psc.mutation.UpdateTime(); !ok {
		v := partystatus.DefaultUpdateTime()
		psc.mutation.SetUpdateTime(v)
	}
	if _, ok := psc.mutation.StatusDate(); !ok {
		v := partystatus.DefaultStatusDate()
		psc.mutation.SetStatusDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PartyStatusCreate) check() error {
	if _, ok := psc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := psc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := psc.mutation.StatusDate(); !ok {
		return &ValidationError{Name: "status_date", err: errors.New("ent: missing required field \"status_date\"")}
	}
	return nil
}

func (psc *PartyStatusCreate) sqlSave(ctx context.Context) (*PartyStatus, error) {
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (psc *PartyStatusCreate) createSpec() (*PartyStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &PartyStatus{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: partystatus.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partystatus.FieldID,
			},
		}
	)
	if value, ok := psc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := psc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := psc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partystatus.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := psc.mutation.StatusDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldStatusDate,
		})
		_node.StatusDate = value
	}
	if nodes := psc.mutation.StatusItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partystatus.StatusItemTable,
			Columns: []string{partystatus.StatusItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statusitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.status_item_party_statuses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partystatus.PartyTable,
			Columns: []string{partystatus.PartyColumn},
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
		_node.party_party_statuses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ChangeByUserLoginIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partystatus.ChangeByUserLoginTable,
			Columns: []string{partystatus.ChangeByUserLoginColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlogin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_login_change_by_party_statuses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PartyStatusCreateBulk is the builder for creating many PartyStatus entities in bulk.
type PartyStatusCreateBulk struct {
	config
	builders []*PartyStatusCreate
}

// Save creates the PartyStatus entities in the database.
func (pscb *PartyStatusCreateBulk) Save(ctx context.Context) ([]*PartyStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PartyStatus, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartyStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PartyStatusCreateBulk) SaveX(ctx context.Context) []*PartyStatus {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
