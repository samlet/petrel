// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/todo/ent/attachment"
	"github.com/samlet/petrel/alfin/modules/todo/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/todo/ent/user"
)

// AttachmentUpdate is the builder for updating Attachment entities.
type AttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where adds a new predicate for the AttachmentUpdate builder.
func (au *AttachmentUpdate) Where(ps ...predicate.Attachment) *AttachmentUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AttachmentUpdate) SetUserID(id int) *AttachmentUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (au *AttachmentUpdate) SetNillableUserID(id *int) *AttachmentUpdate {
	if id != nil {
		au = au.SetUserID(*id)
	}
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AttachmentUpdate) SetUser(u *User) *AttachmentUpdate {
	return au.SetUserID(u.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (au *AttachmentUpdate) Mutation() *AttachmentMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AttachmentUpdate) ClearUser() *AttachmentUpdate {
	au.mutation.ClearUser()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttachmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttachmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttachmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attachment.UserTable,
			Columns: []string{attachment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attachment.UserTable,
			Columns: []string{attachment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AttachmentUpdateOne is the builder for updating a single Attachment entity.
type AttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AttachmentUpdateOne) SetUserID(id int) *AttachmentUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableUserID(id *int) *AttachmentUpdateOne {
	if id != nil {
		auo = auo.SetUserID(*id)
	}
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AttachmentUpdateOne) SetUser(u *User) *AttachmentUpdateOne {
	return auo.SetUserID(u.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (auo *AttachmentUpdateOne) Mutation() *AttachmentMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AttachmentUpdateOne) ClearUser() *AttachmentUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttachmentUpdateOne) Select(field string, fields ...string) *AttachmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attachment entity.
func (auo *AttachmentUpdateOne) Save(ctx context.Context) (*Attachment, error) {
	var (
		err  error
		node *Attachment
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttachmentUpdateOne) SaveX(ctx context.Context) *Attachment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AttachmentUpdateOne) sqlSave(ctx context.Context) (_node *Attachment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Attachment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for _, f := range fields {
			if !attachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attachment.UserTable,
			Columns: []string{attachment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attachment.UserTable,
			Columns: []string{attachment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attachment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
