// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygrouppermission"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitypermission"
)

// SecurityGroupPermissionUpdate is the builder for updating SecurityGroupPermission entities.
type SecurityGroupPermissionUpdate struct {
	config
	hooks    []Hook
	mutation *SecurityGroupPermissionMutation
}

// Where adds a new predicate for the SecurityGroupPermissionUpdate builder.
func (sgpu *SecurityGroupPermissionUpdate) Where(ps ...predicate.SecurityGroupPermission) *SecurityGroupPermissionUpdate {
	sgpu.mutation.predicates = append(sgpu.mutation.predicates, ps...)
	return sgpu
}

// SetStringRef sets the "string_ref" field.
func (sgpu *SecurityGroupPermissionUpdate) SetStringRef(s string) *SecurityGroupPermissionUpdate {
	sgpu.mutation.SetStringRef(s)
	return sgpu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (sgpu *SecurityGroupPermissionUpdate) SetNillableStringRef(s *string) *SecurityGroupPermissionUpdate {
	if s != nil {
		sgpu.SetStringRef(*s)
	}
	return sgpu
}

// ClearStringRef clears the value of the "string_ref" field.
func (sgpu *SecurityGroupPermissionUpdate) ClearStringRef() *SecurityGroupPermissionUpdate {
	sgpu.mutation.ClearStringRef()
	return sgpu
}

// SetFromDate sets the "from_date" field.
func (sgpu *SecurityGroupPermissionUpdate) SetFromDate(t time.Time) *SecurityGroupPermissionUpdate {
	sgpu.mutation.SetFromDate(t)
	return sgpu
}

// SetNillableFromDate sets the "from_date" field if the given value is not nil.
func (sgpu *SecurityGroupPermissionUpdate) SetNillableFromDate(t *time.Time) *SecurityGroupPermissionUpdate {
	if t != nil {
		sgpu.SetFromDate(*t)
	}
	return sgpu
}

// SetThruDate sets the "thru_date" field.
func (sgpu *SecurityGroupPermissionUpdate) SetThruDate(t time.Time) *SecurityGroupPermissionUpdate {
	sgpu.mutation.SetThruDate(t)
	return sgpu
}

// SetNillableThruDate sets the "thru_date" field if the given value is not nil.
func (sgpu *SecurityGroupPermissionUpdate) SetNillableThruDate(t *time.Time) *SecurityGroupPermissionUpdate {
	if t != nil {
		sgpu.SetThruDate(*t)
	}
	return sgpu
}

// ClearThruDate clears the value of the "thru_date" field.
func (sgpu *SecurityGroupPermissionUpdate) ClearThruDate() *SecurityGroupPermissionUpdate {
	sgpu.mutation.ClearThruDate()
	return sgpu
}

// SetSecurityGroupID sets the "security_group" edge to the SecurityGroup entity by ID.
func (sgpu *SecurityGroupPermissionUpdate) SetSecurityGroupID(id int) *SecurityGroupPermissionUpdate {
	sgpu.mutation.SetSecurityGroupID(id)
	return sgpu
}

// SetNillableSecurityGroupID sets the "security_group" edge to the SecurityGroup entity by ID if the given value is not nil.
func (sgpu *SecurityGroupPermissionUpdate) SetNillableSecurityGroupID(id *int) *SecurityGroupPermissionUpdate {
	if id != nil {
		sgpu = sgpu.SetSecurityGroupID(*id)
	}
	return sgpu
}

// SetSecurityGroup sets the "security_group" edge to the SecurityGroup entity.
func (sgpu *SecurityGroupPermissionUpdate) SetSecurityGroup(s *SecurityGroup) *SecurityGroupPermissionUpdate {
	return sgpu.SetSecurityGroupID(s.ID)
}

// SetSecurityPermissionID sets the "security_permission" edge to the SecurityPermission entity by ID.
func (sgpu *SecurityGroupPermissionUpdate) SetSecurityPermissionID(id int) *SecurityGroupPermissionUpdate {
	sgpu.mutation.SetSecurityPermissionID(id)
	return sgpu
}

// SetNillableSecurityPermissionID sets the "security_permission" edge to the SecurityPermission entity by ID if the given value is not nil.
func (sgpu *SecurityGroupPermissionUpdate) SetNillableSecurityPermissionID(id *int) *SecurityGroupPermissionUpdate {
	if id != nil {
		sgpu = sgpu.SetSecurityPermissionID(*id)
	}
	return sgpu
}

// SetSecurityPermission sets the "security_permission" edge to the SecurityPermission entity.
func (sgpu *SecurityGroupPermissionUpdate) SetSecurityPermission(s *SecurityPermission) *SecurityGroupPermissionUpdate {
	return sgpu.SetSecurityPermissionID(s.ID)
}

// Mutation returns the SecurityGroupPermissionMutation object of the builder.
func (sgpu *SecurityGroupPermissionUpdate) Mutation() *SecurityGroupPermissionMutation {
	return sgpu.mutation
}

// ClearSecurityGroup clears the "security_group" edge to the SecurityGroup entity.
func (sgpu *SecurityGroupPermissionUpdate) ClearSecurityGroup() *SecurityGroupPermissionUpdate {
	sgpu.mutation.ClearSecurityGroup()
	return sgpu
}

// ClearSecurityPermission clears the "security_permission" edge to the SecurityPermission entity.
func (sgpu *SecurityGroupPermissionUpdate) ClearSecurityPermission() *SecurityGroupPermissionUpdate {
	sgpu.mutation.ClearSecurityPermission()
	return sgpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sgpu *SecurityGroupPermissionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	sgpu.defaults()
	if len(sgpu.hooks) == 0 {
		affected, err = sgpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecurityGroupPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgpu.mutation = mutation
			affected, err = sgpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sgpu.hooks) - 1; i >= 0; i-- {
			mut = sgpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sgpu *SecurityGroupPermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := sgpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sgpu *SecurityGroupPermissionUpdate) Exec(ctx context.Context) error {
	_, err := sgpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgpu *SecurityGroupPermissionUpdate) ExecX(ctx context.Context) {
	if err := sgpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sgpu *SecurityGroupPermissionUpdate) defaults() {
	if _, ok := sgpu.mutation.UpdateTime(); !ok {
		v := securitygrouppermission.UpdateDefaultUpdateTime()
		sgpu.mutation.SetUpdateTime(v)
	}
}

func (sgpu *SecurityGroupPermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securitygrouppermission.Table,
			Columns: securitygrouppermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securitygrouppermission.FieldID,
			},
		},
	}
	if ps := sgpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sgpu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldUpdateTime,
		})
	}
	if value, ok := sgpu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securitygrouppermission.FieldStringRef,
		})
	}
	if sgpu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: securitygrouppermission.FieldStringRef,
		})
	}
	if value, ok := sgpu.mutation.FromDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldFromDate,
		})
	}
	if value, ok := sgpu.mutation.ThruDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldThruDate,
		})
	}
	if sgpu.mutation.ThruDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: securitygrouppermission.FieldThruDate,
		})
	}
	if sgpu.mutation.SecurityGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityGroupTable,
			Columns: []string{securitygrouppermission.SecurityGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitygroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sgpu.mutation.SecurityGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityGroupTable,
			Columns: []string{securitygrouppermission.SecurityGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sgpu.mutation.SecurityPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityPermissionTable,
			Columns: []string{securitygrouppermission.SecurityPermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitypermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sgpu.mutation.SecurityPermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityPermissionTable,
			Columns: []string{securitygrouppermission.SecurityPermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sgpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{securitygrouppermission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SecurityGroupPermissionUpdateOne is the builder for updating a single SecurityGroupPermission entity.
type SecurityGroupPermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SecurityGroupPermissionMutation
}

// SetStringRef sets the "string_ref" field.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetStringRef(s string) *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.SetStringRef(s)
	return sgpuo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetNillableStringRef(s *string) *SecurityGroupPermissionUpdateOne {
	if s != nil {
		sgpuo.SetStringRef(*s)
	}
	return sgpuo
}

// ClearStringRef clears the value of the "string_ref" field.
func (sgpuo *SecurityGroupPermissionUpdateOne) ClearStringRef() *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.ClearStringRef()
	return sgpuo
}

// SetFromDate sets the "from_date" field.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetFromDate(t time.Time) *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.SetFromDate(t)
	return sgpuo
}

// SetNillableFromDate sets the "from_date" field if the given value is not nil.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetNillableFromDate(t *time.Time) *SecurityGroupPermissionUpdateOne {
	if t != nil {
		sgpuo.SetFromDate(*t)
	}
	return sgpuo
}

// SetThruDate sets the "thru_date" field.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetThruDate(t time.Time) *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.SetThruDate(t)
	return sgpuo
}

// SetNillableThruDate sets the "thru_date" field if the given value is not nil.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetNillableThruDate(t *time.Time) *SecurityGroupPermissionUpdateOne {
	if t != nil {
		sgpuo.SetThruDate(*t)
	}
	return sgpuo
}

// ClearThruDate clears the value of the "thru_date" field.
func (sgpuo *SecurityGroupPermissionUpdateOne) ClearThruDate() *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.ClearThruDate()
	return sgpuo
}

// SetSecurityGroupID sets the "security_group" edge to the SecurityGroup entity by ID.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetSecurityGroupID(id int) *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.SetSecurityGroupID(id)
	return sgpuo
}

// SetNillableSecurityGroupID sets the "security_group" edge to the SecurityGroup entity by ID if the given value is not nil.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetNillableSecurityGroupID(id *int) *SecurityGroupPermissionUpdateOne {
	if id != nil {
		sgpuo = sgpuo.SetSecurityGroupID(*id)
	}
	return sgpuo
}

// SetSecurityGroup sets the "security_group" edge to the SecurityGroup entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetSecurityGroup(s *SecurityGroup) *SecurityGroupPermissionUpdateOne {
	return sgpuo.SetSecurityGroupID(s.ID)
}

// SetSecurityPermissionID sets the "security_permission" edge to the SecurityPermission entity by ID.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetSecurityPermissionID(id int) *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.SetSecurityPermissionID(id)
	return sgpuo
}

// SetNillableSecurityPermissionID sets the "security_permission" edge to the SecurityPermission entity by ID if the given value is not nil.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetNillableSecurityPermissionID(id *int) *SecurityGroupPermissionUpdateOne {
	if id != nil {
		sgpuo = sgpuo.SetSecurityPermissionID(*id)
	}
	return sgpuo
}

// SetSecurityPermission sets the "security_permission" edge to the SecurityPermission entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) SetSecurityPermission(s *SecurityPermission) *SecurityGroupPermissionUpdateOne {
	return sgpuo.SetSecurityPermissionID(s.ID)
}

// Mutation returns the SecurityGroupPermissionMutation object of the builder.
func (sgpuo *SecurityGroupPermissionUpdateOne) Mutation() *SecurityGroupPermissionMutation {
	return sgpuo.mutation
}

// ClearSecurityGroup clears the "security_group" edge to the SecurityGroup entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) ClearSecurityGroup() *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.ClearSecurityGroup()
	return sgpuo
}

// ClearSecurityPermission clears the "security_permission" edge to the SecurityPermission entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) ClearSecurityPermission() *SecurityGroupPermissionUpdateOne {
	sgpuo.mutation.ClearSecurityPermission()
	return sgpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sgpuo *SecurityGroupPermissionUpdateOne) Select(field string, fields ...string) *SecurityGroupPermissionUpdateOne {
	sgpuo.fields = append([]string{field}, fields...)
	return sgpuo
}

// Save executes the query and returns the updated SecurityGroupPermission entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) Save(ctx context.Context) (*SecurityGroupPermission, error) {
	var (
		err  error
		node *SecurityGroupPermission
	)
	sgpuo.defaults()
	if len(sgpuo.hooks) == 0 {
		node, err = sgpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SecurityGroupPermissionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sgpuo.mutation = mutation
			node, err = sgpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sgpuo.hooks) - 1; i >= 0; i-- {
			mut = sgpuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sgpuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sgpuo *SecurityGroupPermissionUpdateOne) SaveX(ctx context.Context) *SecurityGroupPermission {
	node, err := sgpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sgpuo *SecurityGroupPermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := sgpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sgpuo *SecurityGroupPermissionUpdateOne) ExecX(ctx context.Context) {
	if err := sgpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sgpuo *SecurityGroupPermissionUpdateOne) defaults() {
	if _, ok := sgpuo.mutation.UpdateTime(); !ok {
		v := securitygrouppermission.UpdateDefaultUpdateTime()
		sgpuo.mutation.SetUpdateTime(v)
	}
}

func (sgpuo *SecurityGroupPermissionUpdateOne) sqlSave(ctx context.Context) (_node *SecurityGroupPermission, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securitygrouppermission.Table,
			Columns: securitygrouppermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securitygrouppermission.FieldID,
			},
		},
	}
	id, ok := sgpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SecurityGroupPermission.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := sgpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, securitygrouppermission.FieldID)
		for _, f := range fields {
			if !securitygrouppermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != securitygrouppermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sgpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sgpuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldUpdateTime,
		})
	}
	if value, ok := sgpuo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: securitygrouppermission.FieldStringRef,
		})
	}
	if sgpuo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: securitygrouppermission.FieldStringRef,
		})
	}
	if value, ok := sgpuo.mutation.FromDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldFromDate,
		})
	}
	if value, ok := sgpuo.mutation.ThruDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: securitygrouppermission.FieldThruDate,
		})
	}
	if sgpuo.mutation.ThruDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: securitygrouppermission.FieldThruDate,
		})
	}
	if sgpuo.mutation.SecurityGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityGroupTable,
			Columns: []string{securitygrouppermission.SecurityGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitygroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sgpuo.mutation.SecurityGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityGroupTable,
			Columns: []string{securitygrouppermission.SecurityGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sgpuo.mutation.SecurityPermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityPermissionTable,
			Columns: []string{securitygrouppermission.SecurityPermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitypermission.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sgpuo.mutation.SecurityPermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   securitygrouppermission.SecurityPermissionTable,
			Columns: []string{securitygrouppermission.SecurityPermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: securitypermission.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SecurityGroupPermission{config: sgpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sgpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{securitygrouppermission.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}