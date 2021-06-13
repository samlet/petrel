// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechpurposetype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtypepurpose"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ContactMechPurposeTypeUpdate is the builder for updating ContactMechPurposeType entities.
type ContactMechPurposeTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ContactMechPurposeTypeMutation
}

// Where adds a new predicate for the ContactMechPurposeTypeUpdate builder.
func (cmptu *ContactMechPurposeTypeUpdate) Where(ps ...predicate.ContactMechPurposeType) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.predicates = append(cmptu.mutation.predicates, ps...)
	return cmptu
}

// SetStringRef sets the "string_ref" field.
func (cmptu *ContactMechPurposeTypeUpdate) SetStringRef(s string) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.SetStringRef(s)
	return cmptu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (cmptu *ContactMechPurposeTypeUpdate) SetNillableStringRef(s *string) *ContactMechPurposeTypeUpdate {
	if s != nil {
		cmptu.SetStringRef(*s)
	}
	return cmptu
}

// ClearStringRef clears the value of the "string_ref" field.
func (cmptu *ContactMechPurposeTypeUpdate) ClearStringRef() *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ClearStringRef()
	return cmptu
}

// SetParentTypeID sets the "parent_type_id" field.
func (cmptu *ContactMechPurposeTypeUpdate) SetParentTypeID(i int) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ResetParentTypeID()
	cmptu.mutation.SetParentTypeID(i)
	return cmptu
}

// SetNillableParentTypeID sets the "parent_type_id" field if the given value is not nil.
func (cmptu *ContactMechPurposeTypeUpdate) SetNillableParentTypeID(i *int) *ContactMechPurposeTypeUpdate {
	if i != nil {
		cmptu.SetParentTypeID(*i)
	}
	return cmptu
}

// AddParentTypeID adds i to the "parent_type_id" field.
func (cmptu *ContactMechPurposeTypeUpdate) AddParentTypeID(i int) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.AddParentTypeID(i)
	return cmptu
}

// ClearParentTypeID clears the value of the "parent_type_id" field.
func (cmptu *ContactMechPurposeTypeUpdate) ClearParentTypeID() *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ClearParentTypeID()
	return cmptu
}

// SetHasTable sets the "has_table" field.
func (cmptu *ContactMechPurposeTypeUpdate) SetHasTable(ct contactmechpurposetype.HasTable) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.SetHasTable(ct)
	return cmptu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (cmptu *ContactMechPurposeTypeUpdate) SetNillableHasTable(ct *contactmechpurposetype.HasTable) *ContactMechPurposeTypeUpdate {
	if ct != nil {
		cmptu.SetHasTable(*ct)
	}
	return cmptu
}

// ClearHasTable clears the value of the "has_table" field.
func (cmptu *ContactMechPurposeTypeUpdate) ClearHasTable() *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ClearHasTable()
	return cmptu
}

// SetDescription sets the "description" field.
func (cmptu *ContactMechPurposeTypeUpdate) SetDescription(s string) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.SetDescription(s)
	return cmptu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cmptu *ContactMechPurposeTypeUpdate) SetNillableDescription(s *string) *ContactMechPurposeTypeUpdate {
	if s != nil {
		cmptu.SetDescription(*s)
	}
	return cmptu
}

// ClearDescription clears the value of the "description" field.
func (cmptu *ContactMechPurposeTypeUpdate) ClearDescription() *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ClearDescription()
	return cmptu
}

// AddContactMechTypePurposeIDs adds the "contact_mech_type_purposes" edge to the ContactMechTypePurpose entity by IDs.
func (cmptu *ContactMechPurposeTypeUpdate) AddContactMechTypePurposeIDs(ids ...int) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.AddContactMechTypePurposeIDs(ids...)
	return cmptu
}

// AddContactMechTypePurposes adds the "contact_mech_type_purposes" edges to the ContactMechTypePurpose entity.
func (cmptu *ContactMechPurposeTypeUpdate) AddContactMechTypePurposes(c ...*ContactMechTypePurpose) *ContactMechPurposeTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmptu.AddContactMechTypePurposeIDs(ids...)
}

// Mutation returns the ContactMechPurposeTypeMutation object of the builder.
func (cmptu *ContactMechPurposeTypeUpdate) Mutation() *ContactMechPurposeTypeMutation {
	return cmptu.mutation
}

// ClearContactMechTypePurposes clears all "contact_mech_type_purposes" edges to the ContactMechTypePurpose entity.
func (cmptu *ContactMechPurposeTypeUpdate) ClearContactMechTypePurposes() *ContactMechPurposeTypeUpdate {
	cmptu.mutation.ClearContactMechTypePurposes()
	return cmptu
}

// RemoveContactMechTypePurposeIDs removes the "contact_mech_type_purposes" edge to ContactMechTypePurpose entities by IDs.
func (cmptu *ContactMechPurposeTypeUpdate) RemoveContactMechTypePurposeIDs(ids ...int) *ContactMechPurposeTypeUpdate {
	cmptu.mutation.RemoveContactMechTypePurposeIDs(ids...)
	return cmptu
}

// RemoveContactMechTypePurposes removes "contact_mech_type_purposes" edges to ContactMechTypePurpose entities.
func (cmptu *ContactMechPurposeTypeUpdate) RemoveContactMechTypePurposes(c ...*ContactMechTypePurpose) *ContactMechPurposeTypeUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmptu.RemoveContactMechTypePurposeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmptu *ContactMechPurposeTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cmptu.defaults()
	if len(cmptu.hooks) == 0 {
		if err = cmptu.check(); err != nil {
			return 0, err
		}
		affected, err = cmptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMechPurposeTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmptu.check(); err != nil {
				return 0, err
			}
			cmptu.mutation = mutation
			affected, err = cmptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmptu.hooks) - 1; i >= 0; i-- {
			mut = cmptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmptu *ContactMechPurposeTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := cmptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmptu *ContactMechPurposeTypeUpdate) Exec(ctx context.Context) error {
	_, err := cmptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmptu *ContactMechPurposeTypeUpdate) ExecX(ctx context.Context) {
	if err := cmptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmptu *ContactMechPurposeTypeUpdate) defaults() {
	if _, ok := cmptu.mutation.UpdateTime(); !ok {
		v := contactmechpurposetype.UpdateDefaultUpdateTime()
		cmptu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmptu *ContactMechPurposeTypeUpdate) check() error {
	if v, ok := cmptu.mutation.HasTable(); ok {
		if err := contactmechpurposetype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (cmptu *ContactMechPurposeTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contactmechpurposetype.Table,
			Columns: contactmechpurposetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contactmechpurposetype.FieldID,
			},
		},
	}
	if ps := cmptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmptu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contactmechpurposetype.FieldUpdateTime,
		})
	}
	if value, ok := cmptu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechpurposetype.FieldStringRef,
		})
	}
	if cmptu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contactmechpurposetype.FieldStringRef,
		})
	}
	if value, ok := cmptu.mutation.ParentTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if value, ok := cmptu.mutation.AddedParentTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if cmptu.mutation.ParentTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if value, ok := cmptu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: contactmechpurposetype.FieldHasTable,
		})
	}
	if cmptu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: contactmechpurposetype.FieldHasTable,
		})
	}
	if value, ok := cmptu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechpurposetype.FieldDescription,
		})
	}
	if cmptu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contactmechpurposetype.FieldDescription,
		})
	}
	if cmptu.mutation.ContactMechTypePurposesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmptu.mutation.RemovedContactMechTypePurposesIDs(); len(nodes) > 0 && !cmptu.mutation.ContactMechTypePurposesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmptu.mutation.ContactMechTypePurposesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contactmechpurposetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ContactMechPurposeTypeUpdateOne is the builder for updating a single ContactMechPurposeType entity.
type ContactMechPurposeTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContactMechPurposeTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetStringRef(s string) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.SetStringRef(s)
	return cmptuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetNillableStringRef(s *string) *ContactMechPurposeTypeUpdateOne {
	if s != nil {
		cmptuo.SetStringRef(*s)
	}
	return cmptuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ClearStringRef() *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ClearStringRef()
	return cmptuo
}

// SetParentTypeID sets the "parent_type_id" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetParentTypeID(i int) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ResetParentTypeID()
	cmptuo.mutation.SetParentTypeID(i)
	return cmptuo
}

// SetNillableParentTypeID sets the "parent_type_id" field if the given value is not nil.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetNillableParentTypeID(i *int) *ContactMechPurposeTypeUpdateOne {
	if i != nil {
		cmptuo.SetParentTypeID(*i)
	}
	return cmptuo
}

// AddParentTypeID adds i to the "parent_type_id" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) AddParentTypeID(i int) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.AddParentTypeID(i)
	return cmptuo
}

// ClearParentTypeID clears the value of the "parent_type_id" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ClearParentTypeID() *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ClearParentTypeID()
	return cmptuo
}

// SetHasTable sets the "has_table" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetHasTable(ct contactmechpurposetype.HasTable) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.SetHasTable(ct)
	return cmptuo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetNillableHasTable(ct *contactmechpurposetype.HasTable) *ContactMechPurposeTypeUpdateOne {
	if ct != nil {
		cmptuo.SetHasTable(*ct)
	}
	return cmptuo
}

// ClearHasTable clears the value of the "has_table" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ClearHasTable() *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ClearHasTable()
	return cmptuo
}

// SetDescription sets the "description" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetDescription(s string) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.SetDescription(s)
	return cmptuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SetNillableDescription(s *string) *ContactMechPurposeTypeUpdateOne {
	if s != nil {
		cmptuo.SetDescription(*s)
	}
	return cmptuo
}

// ClearDescription clears the value of the "description" field.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ClearDescription() *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ClearDescription()
	return cmptuo
}

// AddContactMechTypePurposeIDs adds the "contact_mech_type_purposes" edge to the ContactMechTypePurpose entity by IDs.
func (cmptuo *ContactMechPurposeTypeUpdateOne) AddContactMechTypePurposeIDs(ids ...int) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.AddContactMechTypePurposeIDs(ids...)
	return cmptuo
}

// AddContactMechTypePurposes adds the "contact_mech_type_purposes" edges to the ContactMechTypePurpose entity.
func (cmptuo *ContactMechPurposeTypeUpdateOne) AddContactMechTypePurposes(c ...*ContactMechTypePurpose) *ContactMechPurposeTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmptuo.AddContactMechTypePurposeIDs(ids...)
}

// Mutation returns the ContactMechPurposeTypeMutation object of the builder.
func (cmptuo *ContactMechPurposeTypeUpdateOne) Mutation() *ContactMechPurposeTypeMutation {
	return cmptuo.mutation
}

// ClearContactMechTypePurposes clears all "contact_mech_type_purposes" edges to the ContactMechTypePurpose entity.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ClearContactMechTypePurposes() *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.ClearContactMechTypePurposes()
	return cmptuo
}

// RemoveContactMechTypePurposeIDs removes the "contact_mech_type_purposes" edge to ContactMechTypePurpose entities by IDs.
func (cmptuo *ContactMechPurposeTypeUpdateOne) RemoveContactMechTypePurposeIDs(ids ...int) *ContactMechPurposeTypeUpdateOne {
	cmptuo.mutation.RemoveContactMechTypePurposeIDs(ids...)
	return cmptuo
}

// RemoveContactMechTypePurposes removes "contact_mech_type_purposes" edges to ContactMechTypePurpose entities.
func (cmptuo *ContactMechPurposeTypeUpdateOne) RemoveContactMechTypePurposes(c ...*ContactMechTypePurpose) *ContactMechPurposeTypeUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmptuo.RemoveContactMechTypePurposeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmptuo *ContactMechPurposeTypeUpdateOne) Select(field string, fields ...string) *ContactMechPurposeTypeUpdateOne {
	cmptuo.fields = append([]string{field}, fields...)
	return cmptuo
}

// Save executes the query and returns the updated ContactMechPurposeType entity.
func (cmptuo *ContactMechPurposeTypeUpdateOne) Save(ctx context.Context) (*ContactMechPurposeType, error) {
	var (
		err  error
		node *ContactMechPurposeType
	)
	cmptuo.defaults()
	if len(cmptuo.hooks) == 0 {
		if err = cmptuo.check(); err != nil {
			return nil, err
		}
		node, err = cmptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMechPurposeTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmptuo.check(); err != nil {
				return nil, err
			}
			cmptuo.mutation = mutation
			node, err = cmptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cmptuo.hooks) - 1; i >= 0; i-- {
			mut = cmptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmptuo *ContactMechPurposeTypeUpdateOne) SaveX(ctx context.Context) *ContactMechPurposeType {
	node, err := cmptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmptuo *ContactMechPurposeTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := cmptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmptuo *ContactMechPurposeTypeUpdateOne) ExecX(ctx context.Context) {
	if err := cmptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cmptuo *ContactMechPurposeTypeUpdateOne) defaults() {
	if _, ok := cmptuo.mutation.UpdateTime(); !ok {
		v := contactmechpurposetype.UpdateDefaultUpdateTime()
		cmptuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmptuo *ContactMechPurposeTypeUpdateOne) check() error {
	if v, ok := cmptuo.mutation.HasTable(); ok {
		if err := contactmechpurposetype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (cmptuo *ContactMechPurposeTypeUpdateOne) sqlSave(ctx context.Context) (_node *ContactMechPurposeType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contactmechpurposetype.Table,
			Columns: contactmechpurposetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contactmechpurposetype.FieldID,
			},
		},
	}
	id, ok := cmptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ContactMechPurposeType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cmptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contactmechpurposetype.FieldID)
		for _, f := range fields {
			if !contactmechpurposetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contactmechpurposetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmptuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contactmechpurposetype.FieldUpdateTime,
		})
	}
	if value, ok := cmptuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechpurposetype.FieldStringRef,
		})
	}
	if cmptuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contactmechpurposetype.FieldStringRef,
		})
	}
	if value, ok := cmptuo.mutation.ParentTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if value, ok := cmptuo.mutation.AddedParentTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if cmptuo.mutation.ParentTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: contactmechpurposetype.FieldParentTypeID,
		})
	}
	if value, ok := cmptuo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: contactmechpurposetype.FieldHasTable,
		})
	}
	if cmptuo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: contactmechpurposetype.FieldHasTable,
		})
	}
	if value, ok := cmptuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechpurposetype.FieldDescription,
		})
	}
	if cmptuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contactmechpurposetype.FieldDescription,
		})
	}
	if cmptuo.mutation.ContactMechTypePurposesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmptuo.mutation.RemovedContactMechTypePurposesIDs(); len(nodes) > 0 && !cmptuo.mutation.ContactMechTypePurposesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmptuo.mutation.ContactMechTypePurposesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechpurposetype.ContactMechTypePurposesTable,
			Columns: []string{contactmechpurposetype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContactMechPurposeType{config: cmptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contactmechpurposetype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
