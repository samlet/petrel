// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partystatus"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
)

// PartyStatusUpdate is the builder for updating PartyStatus entities.
type PartyStatusUpdate struct {
	config
	hooks    []Hook
	mutation *PartyStatusMutation
}

// Where adds a new predicate for the PartyStatusUpdate builder.
func (psu *PartyStatusUpdate) Where(ps ...predicate.PartyStatus) *PartyStatusUpdate {
	psu.mutation.predicates = append(psu.mutation.predicates, ps...)
	return psu
}

// SetStringRef sets the "string_ref" field.
func (psu *PartyStatusUpdate) SetStringRef(s string) *PartyStatusUpdate {
	psu.mutation.SetStringRef(s)
	return psu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (psu *PartyStatusUpdate) SetNillableStringRef(s *string) *PartyStatusUpdate {
	if s != nil {
		psu.SetStringRef(*s)
	}
	return psu
}

// ClearStringRef clears the value of the "string_ref" field.
func (psu *PartyStatusUpdate) ClearStringRef() *PartyStatusUpdate {
	psu.mutation.ClearStringRef()
	return psu
}

// SetStatusDate sets the "status_date" field.
func (psu *PartyStatusUpdate) SetStatusDate(t time.Time) *PartyStatusUpdate {
	psu.mutation.SetStatusDate(t)
	return psu
}

// SetNillableStatusDate sets the "status_date" field if the given value is not nil.
func (psu *PartyStatusUpdate) SetNillableStatusDate(t *time.Time) *PartyStatusUpdate {
	if t != nil {
		psu.SetStatusDate(*t)
	}
	return psu
}

// SetStatusItemID sets the "status_item" edge to the StatusItem entity by ID.
func (psu *PartyStatusUpdate) SetStatusItemID(id int) *PartyStatusUpdate {
	psu.mutation.SetStatusItemID(id)
	return psu
}

// SetNillableStatusItemID sets the "status_item" edge to the StatusItem entity by ID if the given value is not nil.
func (psu *PartyStatusUpdate) SetNillableStatusItemID(id *int) *PartyStatusUpdate {
	if id != nil {
		psu = psu.SetStatusItemID(*id)
	}
	return psu
}

// SetStatusItem sets the "status_item" edge to the StatusItem entity.
func (psu *PartyStatusUpdate) SetStatusItem(s *StatusItem) *PartyStatusUpdate {
	return psu.SetStatusItemID(s.ID)
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (psu *PartyStatusUpdate) SetPartyID(id int) *PartyStatusUpdate {
	psu.mutation.SetPartyID(id)
	return psu
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (psu *PartyStatusUpdate) SetNillablePartyID(id *int) *PartyStatusUpdate {
	if id != nil {
		psu = psu.SetPartyID(*id)
	}
	return psu
}

// SetParty sets the "party" edge to the Party entity.
func (psu *PartyStatusUpdate) SetParty(p *Party) *PartyStatusUpdate {
	return psu.SetPartyID(p.ID)
}

// SetChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID.
func (psu *PartyStatusUpdate) SetChangeByUserLoginID(id int) *PartyStatusUpdate {
	psu.mutation.SetChangeByUserLoginID(id)
	return psu
}

// SetNillableChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID if the given value is not nil.
func (psu *PartyStatusUpdate) SetNillableChangeByUserLoginID(id *int) *PartyStatusUpdate {
	if id != nil {
		psu = psu.SetChangeByUserLoginID(*id)
	}
	return psu
}

// SetChangeByUserLogin sets the "change_by_user_login" edge to the UserLogin entity.
func (psu *PartyStatusUpdate) SetChangeByUserLogin(u *UserLogin) *PartyStatusUpdate {
	return psu.SetChangeByUserLoginID(u.ID)
}

// Mutation returns the PartyStatusMutation object of the builder.
func (psu *PartyStatusUpdate) Mutation() *PartyStatusMutation {
	return psu.mutation
}

// ClearStatusItem clears the "status_item" edge to the StatusItem entity.
func (psu *PartyStatusUpdate) ClearStatusItem() *PartyStatusUpdate {
	psu.mutation.ClearStatusItem()
	return psu
}

// ClearParty clears the "party" edge to the Party entity.
func (psu *PartyStatusUpdate) ClearParty() *PartyStatusUpdate {
	psu.mutation.ClearParty()
	return psu
}

// ClearChangeByUserLogin clears the "change_by_user_login" edge to the UserLogin entity.
func (psu *PartyStatusUpdate) ClearChangeByUserLogin() *PartyStatusUpdate {
	psu.mutation.ClearChangeByUserLogin()
	return psu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *PartyStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	psu.defaults()
	if len(psu.hooks) == 0 {
		affected, err = psu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psu.mutation = mutation
			affected, err = psu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psu.hooks) - 1; i >= 0; i-- {
			mut = psu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psu *PartyStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *PartyStatusUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *PartyStatusUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *PartyStatusUpdate) defaults() {
	if _, ok := psu.mutation.UpdateTime(); !ok {
		v := partystatus.UpdateDefaultUpdateTime()
		psu.mutation.SetUpdateTime(v)
	}
}

func (psu *PartyStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partystatus.Table,
			Columns: partystatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partystatus.FieldID,
			},
		},
	}
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldUpdateTime,
		})
	}
	if value, ok := psu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partystatus.FieldStringRef,
		})
	}
	if psu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partystatus.FieldStringRef,
		})
	}
	if value, ok := psu.mutation.StatusDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldStatusDate,
		})
	}
	if psu.mutation.StatusItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.StatusItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psu.mutation.PartyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.PartyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psu.mutation.ChangeByUserLoginCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.ChangeByUserLoginIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partystatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PartyStatusUpdateOne is the builder for updating a single PartyStatus entity.
type PartyStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PartyStatusMutation
}

// SetStringRef sets the "string_ref" field.
func (psuo *PartyStatusUpdateOne) SetStringRef(s string) *PartyStatusUpdateOne {
	psuo.mutation.SetStringRef(s)
	return psuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (psuo *PartyStatusUpdateOne) SetNillableStringRef(s *string) *PartyStatusUpdateOne {
	if s != nil {
		psuo.SetStringRef(*s)
	}
	return psuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (psuo *PartyStatusUpdateOne) ClearStringRef() *PartyStatusUpdateOne {
	psuo.mutation.ClearStringRef()
	return psuo
}

// SetStatusDate sets the "status_date" field.
func (psuo *PartyStatusUpdateOne) SetStatusDate(t time.Time) *PartyStatusUpdateOne {
	psuo.mutation.SetStatusDate(t)
	return psuo
}

// SetNillableStatusDate sets the "status_date" field if the given value is not nil.
func (psuo *PartyStatusUpdateOne) SetNillableStatusDate(t *time.Time) *PartyStatusUpdateOne {
	if t != nil {
		psuo.SetStatusDate(*t)
	}
	return psuo
}

// SetStatusItemID sets the "status_item" edge to the StatusItem entity by ID.
func (psuo *PartyStatusUpdateOne) SetStatusItemID(id int) *PartyStatusUpdateOne {
	psuo.mutation.SetStatusItemID(id)
	return psuo
}

// SetNillableStatusItemID sets the "status_item" edge to the StatusItem entity by ID if the given value is not nil.
func (psuo *PartyStatusUpdateOne) SetNillableStatusItemID(id *int) *PartyStatusUpdateOne {
	if id != nil {
		psuo = psuo.SetStatusItemID(*id)
	}
	return psuo
}

// SetStatusItem sets the "status_item" edge to the StatusItem entity.
func (psuo *PartyStatusUpdateOne) SetStatusItem(s *StatusItem) *PartyStatusUpdateOne {
	return psuo.SetStatusItemID(s.ID)
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (psuo *PartyStatusUpdateOne) SetPartyID(id int) *PartyStatusUpdateOne {
	psuo.mutation.SetPartyID(id)
	return psuo
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (psuo *PartyStatusUpdateOne) SetNillablePartyID(id *int) *PartyStatusUpdateOne {
	if id != nil {
		psuo = psuo.SetPartyID(*id)
	}
	return psuo
}

// SetParty sets the "party" edge to the Party entity.
func (psuo *PartyStatusUpdateOne) SetParty(p *Party) *PartyStatusUpdateOne {
	return psuo.SetPartyID(p.ID)
}

// SetChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID.
func (psuo *PartyStatusUpdateOne) SetChangeByUserLoginID(id int) *PartyStatusUpdateOne {
	psuo.mutation.SetChangeByUserLoginID(id)
	return psuo
}

// SetNillableChangeByUserLoginID sets the "change_by_user_login" edge to the UserLogin entity by ID if the given value is not nil.
func (psuo *PartyStatusUpdateOne) SetNillableChangeByUserLoginID(id *int) *PartyStatusUpdateOne {
	if id != nil {
		psuo = psuo.SetChangeByUserLoginID(*id)
	}
	return psuo
}

// SetChangeByUserLogin sets the "change_by_user_login" edge to the UserLogin entity.
func (psuo *PartyStatusUpdateOne) SetChangeByUserLogin(u *UserLogin) *PartyStatusUpdateOne {
	return psuo.SetChangeByUserLoginID(u.ID)
}

// Mutation returns the PartyStatusMutation object of the builder.
func (psuo *PartyStatusUpdateOne) Mutation() *PartyStatusMutation {
	return psuo.mutation
}

// ClearStatusItem clears the "status_item" edge to the StatusItem entity.
func (psuo *PartyStatusUpdateOne) ClearStatusItem() *PartyStatusUpdateOne {
	psuo.mutation.ClearStatusItem()
	return psuo
}

// ClearParty clears the "party" edge to the Party entity.
func (psuo *PartyStatusUpdateOne) ClearParty() *PartyStatusUpdateOne {
	psuo.mutation.ClearParty()
	return psuo
}

// ClearChangeByUserLogin clears the "change_by_user_login" edge to the UserLogin entity.
func (psuo *PartyStatusUpdateOne) ClearChangeByUserLogin() *PartyStatusUpdateOne {
	psuo.mutation.ClearChangeByUserLogin()
	return psuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *PartyStatusUpdateOne) Select(field string, fields ...string) *PartyStatusUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated PartyStatus entity.
func (psuo *PartyStatusUpdateOne) Save(ctx context.Context) (*PartyStatus, error) {
	var (
		err  error
		node *PartyStatus
	)
	psuo.defaults()
	if len(psuo.hooks) == 0 {
		node, err = psuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psuo.mutation = mutation
			node, err = psuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psuo.hooks) - 1; i >= 0; i-- {
			mut = psuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *PartyStatusUpdateOne) SaveX(ctx context.Context) *PartyStatus {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *PartyStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *PartyStatusUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *PartyStatusUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdateTime(); !ok {
		v := partystatus.UpdateDefaultUpdateTime()
		psuo.mutation.SetUpdateTime(v)
	}
}

func (psuo *PartyStatusUpdateOne) sqlSave(ctx context.Context) (_node *PartyStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partystatus.Table,
			Columns: partystatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partystatus.FieldID,
			},
		},
	}
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PartyStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partystatus.FieldID)
		for _, f := range fields {
			if !partystatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != partystatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldUpdateTime,
		})
	}
	if value, ok := psuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partystatus.FieldStringRef,
		})
	}
	if psuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: partystatus.FieldStringRef,
		})
	}
	if value, ok := psuo.mutation.StatusDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partystatus.FieldStatusDate,
		})
	}
	if psuo.mutation.StatusItemCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.StatusItemIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psuo.mutation.PartyCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.PartyIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psuo.mutation.ChangeByUserLoginCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.ChangeByUserLoginIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PartyStatus{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{partystatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
