// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadstatus"
)

// WorkloadStatusUpdate is the builder for updating WorkloadStatus entities.
type WorkloadStatusUpdate struct {
	config
	hooks    []Hook
	mutation *WorkloadStatusMutation
}

// Where adds a new predicate for the WorkloadStatusUpdate builder.
func (wsu *WorkloadStatusUpdate) Where(ps ...predicate.WorkloadStatus) *WorkloadStatusUpdate {
	wsu.mutation.predicates = append(wsu.mutation.predicates, ps...)
	return wsu
}

// SetStatusDate sets the "status_date" field.
func (wsu *WorkloadStatusUpdate) SetStatusDate(t time.Time) *WorkloadStatusUpdate {
	wsu.mutation.SetStatusDate(t)
	return wsu
}

// SetNillableStatusDate sets the "status_date" field if the given value is not nil.
func (wsu *WorkloadStatusUpdate) SetNillableStatusDate(t *time.Time) *WorkloadStatusUpdate {
	if t != nil {
		wsu.SetStatusDate(*t)
	}
	return wsu
}

// SetStatusEndDate sets the "status_end_date" field.
func (wsu *WorkloadStatusUpdate) SetStatusEndDate(t time.Time) *WorkloadStatusUpdate {
	wsu.mutation.SetStatusEndDate(t)
	return wsu
}

// SetNillableStatusEndDate sets the "status_end_date" field if the given value is not nil.
func (wsu *WorkloadStatusUpdate) SetNillableStatusEndDate(t *time.Time) *WorkloadStatusUpdate {
	if t != nil {
		wsu.SetStatusEndDate(*t)
	}
	return wsu
}

// ClearStatusEndDate clears the value of the "status_end_date" field.
func (wsu *WorkloadStatusUpdate) ClearStatusEndDate() *WorkloadStatusUpdate {
	wsu.mutation.ClearStatusEndDate()
	return wsu
}

// SetChangeByUserLoginID sets the "change_by_user_login_id" field.
func (wsu *WorkloadStatusUpdate) SetChangeByUserLoginID(s string) *WorkloadStatusUpdate {
	wsu.mutation.SetChangeByUserLoginID(s)
	return wsu
}

// SetNillableChangeByUserLoginID sets the "change_by_user_login_id" field if the given value is not nil.
func (wsu *WorkloadStatusUpdate) SetNillableChangeByUserLoginID(s *string) *WorkloadStatusUpdate {
	if s != nil {
		wsu.SetChangeByUserLoginID(*s)
	}
	return wsu
}

// ClearChangeByUserLoginID clears the value of the "change_by_user_login_id" field.
func (wsu *WorkloadStatusUpdate) ClearChangeByUserLoginID() *WorkloadStatusUpdate {
	wsu.mutation.ClearChangeByUserLoginID()
	return wsu
}

// SetStatusID sets the "status_id" field.
func (wsu *WorkloadStatusUpdate) SetStatusID(i int) *WorkloadStatusUpdate {
	wsu.mutation.ResetStatusID()
	wsu.mutation.SetStatusID(i)
	return wsu
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (wsu *WorkloadStatusUpdate) SetNillableStatusID(i *int) *WorkloadStatusUpdate {
	if i != nil {
		wsu.SetStatusID(*i)
	}
	return wsu
}

// AddStatusID adds i to the "status_id" field.
func (wsu *WorkloadStatusUpdate) AddStatusID(i int) *WorkloadStatusUpdate {
	wsu.mutation.AddStatusID(i)
	return wsu
}

// ClearStatusID clears the value of the "status_id" field.
func (wsu *WorkloadStatusUpdate) ClearStatusID() *WorkloadStatusUpdate {
	wsu.mutation.ClearStatusID()
	return wsu
}

// SetWorkloadID sets the "workload" edge to the Workload entity by ID.
func (wsu *WorkloadStatusUpdate) SetWorkloadID(id int) *WorkloadStatusUpdate {
	wsu.mutation.SetWorkloadID(id)
	return wsu
}

// SetNillableWorkloadID sets the "workload" edge to the Workload entity by ID if the given value is not nil.
func (wsu *WorkloadStatusUpdate) SetNillableWorkloadID(id *int) *WorkloadStatusUpdate {
	if id != nil {
		wsu = wsu.SetWorkloadID(*id)
	}
	return wsu
}

// SetWorkload sets the "workload" edge to the Workload entity.
func (wsu *WorkloadStatusUpdate) SetWorkload(w *Workload) *WorkloadStatusUpdate {
	return wsu.SetWorkloadID(w.ID)
}

// Mutation returns the WorkloadStatusMutation object of the builder.
func (wsu *WorkloadStatusUpdate) Mutation() *WorkloadStatusMutation {
	return wsu.mutation
}

// ClearWorkload clears the "workload" edge to the Workload entity.
func (wsu *WorkloadStatusUpdate) ClearWorkload() *WorkloadStatusUpdate {
	wsu.mutation.ClearWorkload()
	return wsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wsu *WorkloadStatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wsu.defaults()
	if len(wsu.hooks) == 0 {
		affected, err = wsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wsu.mutation = mutation
			affected, err = wsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wsu.hooks) - 1; i >= 0; i-- {
			mut = wsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wsu *WorkloadStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := wsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wsu *WorkloadStatusUpdate) Exec(ctx context.Context) error {
	_, err := wsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsu *WorkloadStatusUpdate) ExecX(ctx context.Context) {
	if err := wsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsu *WorkloadStatusUpdate) defaults() {
	if _, ok := wsu.mutation.UpdateTime(); !ok {
		v := workloadstatus.UpdateDefaultUpdateTime()
		wsu.mutation.SetUpdateTime(v)
	}
}

func (wsu *WorkloadStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workloadstatus.Table,
			Columns: workloadstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadstatus.FieldID,
			},
		},
	}
	if ps := wsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldUpdateTime,
		})
	}
	if value, ok := wsu.mutation.StatusDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldStatusDate,
		})
	}
	if value, ok := wsu.mutation.StatusEndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldStatusEndDate,
		})
	}
	if wsu.mutation.StatusEndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workloadstatus.FieldStatusEndDate,
		})
	}
	if value, ok := wsu.mutation.ChangeByUserLoginID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workloadstatus.FieldChangeByUserLoginID,
		})
	}
	if wsu.mutation.ChangeByUserLoginIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workloadstatus.FieldChangeByUserLoginID,
		})
	}
	if value, ok := wsu.mutation.StatusID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if value, ok := wsu.mutation.AddedStatusID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if wsu.mutation.StatusIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if wsu.mutation.WorkloadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadstatus.WorkloadTable,
			Columns: []string{workloadstatus.WorkloadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workload.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsu.mutation.WorkloadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadstatus.WorkloadTable,
			Columns: []string{workloadstatus.WorkloadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workload.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workloadstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WorkloadStatusUpdateOne is the builder for updating a single WorkloadStatus entity.
type WorkloadStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkloadStatusMutation
}

// SetStatusDate sets the "status_date" field.
func (wsuo *WorkloadStatusUpdateOne) SetStatusDate(t time.Time) *WorkloadStatusUpdateOne {
	wsuo.mutation.SetStatusDate(t)
	return wsuo
}

// SetNillableStatusDate sets the "status_date" field if the given value is not nil.
func (wsuo *WorkloadStatusUpdateOne) SetNillableStatusDate(t *time.Time) *WorkloadStatusUpdateOne {
	if t != nil {
		wsuo.SetStatusDate(*t)
	}
	return wsuo
}

// SetStatusEndDate sets the "status_end_date" field.
func (wsuo *WorkloadStatusUpdateOne) SetStatusEndDate(t time.Time) *WorkloadStatusUpdateOne {
	wsuo.mutation.SetStatusEndDate(t)
	return wsuo
}

// SetNillableStatusEndDate sets the "status_end_date" field if the given value is not nil.
func (wsuo *WorkloadStatusUpdateOne) SetNillableStatusEndDate(t *time.Time) *WorkloadStatusUpdateOne {
	if t != nil {
		wsuo.SetStatusEndDate(*t)
	}
	return wsuo
}

// ClearStatusEndDate clears the value of the "status_end_date" field.
func (wsuo *WorkloadStatusUpdateOne) ClearStatusEndDate() *WorkloadStatusUpdateOne {
	wsuo.mutation.ClearStatusEndDate()
	return wsuo
}

// SetChangeByUserLoginID sets the "change_by_user_login_id" field.
func (wsuo *WorkloadStatusUpdateOne) SetChangeByUserLoginID(s string) *WorkloadStatusUpdateOne {
	wsuo.mutation.SetChangeByUserLoginID(s)
	return wsuo
}

// SetNillableChangeByUserLoginID sets the "change_by_user_login_id" field if the given value is not nil.
func (wsuo *WorkloadStatusUpdateOne) SetNillableChangeByUserLoginID(s *string) *WorkloadStatusUpdateOne {
	if s != nil {
		wsuo.SetChangeByUserLoginID(*s)
	}
	return wsuo
}

// ClearChangeByUserLoginID clears the value of the "change_by_user_login_id" field.
func (wsuo *WorkloadStatusUpdateOne) ClearChangeByUserLoginID() *WorkloadStatusUpdateOne {
	wsuo.mutation.ClearChangeByUserLoginID()
	return wsuo
}

// SetStatusID sets the "status_id" field.
func (wsuo *WorkloadStatusUpdateOne) SetStatusID(i int) *WorkloadStatusUpdateOne {
	wsuo.mutation.ResetStatusID()
	wsuo.mutation.SetStatusID(i)
	return wsuo
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (wsuo *WorkloadStatusUpdateOne) SetNillableStatusID(i *int) *WorkloadStatusUpdateOne {
	if i != nil {
		wsuo.SetStatusID(*i)
	}
	return wsuo
}

// AddStatusID adds i to the "status_id" field.
func (wsuo *WorkloadStatusUpdateOne) AddStatusID(i int) *WorkloadStatusUpdateOne {
	wsuo.mutation.AddStatusID(i)
	return wsuo
}

// ClearStatusID clears the value of the "status_id" field.
func (wsuo *WorkloadStatusUpdateOne) ClearStatusID() *WorkloadStatusUpdateOne {
	wsuo.mutation.ClearStatusID()
	return wsuo
}

// SetWorkloadID sets the "workload" edge to the Workload entity by ID.
func (wsuo *WorkloadStatusUpdateOne) SetWorkloadID(id int) *WorkloadStatusUpdateOne {
	wsuo.mutation.SetWorkloadID(id)
	return wsuo
}

// SetNillableWorkloadID sets the "workload" edge to the Workload entity by ID if the given value is not nil.
func (wsuo *WorkloadStatusUpdateOne) SetNillableWorkloadID(id *int) *WorkloadStatusUpdateOne {
	if id != nil {
		wsuo = wsuo.SetWorkloadID(*id)
	}
	return wsuo
}

// SetWorkload sets the "workload" edge to the Workload entity.
func (wsuo *WorkloadStatusUpdateOne) SetWorkload(w *Workload) *WorkloadStatusUpdateOne {
	return wsuo.SetWorkloadID(w.ID)
}

// Mutation returns the WorkloadStatusMutation object of the builder.
func (wsuo *WorkloadStatusUpdateOne) Mutation() *WorkloadStatusMutation {
	return wsuo.mutation
}

// ClearWorkload clears the "workload" edge to the Workload entity.
func (wsuo *WorkloadStatusUpdateOne) ClearWorkload() *WorkloadStatusUpdateOne {
	wsuo.mutation.ClearWorkload()
	return wsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wsuo *WorkloadStatusUpdateOne) Select(field string, fields ...string) *WorkloadStatusUpdateOne {
	wsuo.fields = append([]string{field}, fields...)
	return wsuo
}

// Save executes the query and returns the updated WorkloadStatus entity.
func (wsuo *WorkloadStatusUpdateOne) Save(ctx context.Context) (*WorkloadStatus, error) {
	var (
		err  error
		node *WorkloadStatus
	)
	wsuo.defaults()
	if len(wsuo.hooks) == 0 {
		node, err = wsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkloadStatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wsuo.mutation = mutation
			node, err = wsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wsuo.hooks) - 1; i >= 0; i-- {
			mut = wsuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wsuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wsuo *WorkloadStatusUpdateOne) SaveX(ctx context.Context) *WorkloadStatus {
	node, err := wsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wsuo *WorkloadStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := wsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wsuo *WorkloadStatusUpdateOne) ExecX(ctx context.Context) {
	if err := wsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wsuo *WorkloadStatusUpdateOne) defaults() {
	if _, ok := wsuo.mutation.UpdateTime(); !ok {
		v := workloadstatus.UpdateDefaultUpdateTime()
		wsuo.mutation.SetUpdateTime(v)
	}
}

func (wsuo *WorkloadStatusUpdateOne) sqlSave(ctx context.Context) (_node *WorkloadStatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workloadstatus.Table,
			Columns: workloadstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadstatus.FieldID,
			},
		},
	}
	id, ok := wsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing WorkloadStatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workloadstatus.FieldID)
		for _, f := range fields {
			if !workloadstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workloadstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wsuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldUpdateTime,
		})
	}
	if value, ok := wsuo.mutation.StatusDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldStatusDate,
		})
	}
	if value, ok := wsuo.mutation.StatusEndDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workloadstatus.FieldStatusEndDate,
		})
	}
	if wsuo.mutation.StatusEndDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: workloadstatus.FieldStatusEndDate,
		})
	}
	if value, ok := wsuo.mutation.ChangeByUserLoginID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workloadstatus.FieldChangeByUserLoginID,
		})
	}
	if wsuo.mutation.ChangeByUserLoginIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: workloadstatus.FieldChangeByUserLoginID,
		})
	}
	if value, ok := wsuo.mutation.StatusID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if value, ok := wsuo.mutation.AddedStatusID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if wsuo.mutation.StatusIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: workloadstatus.FieldStatusID,
		})
	}
	if wsuo.mutation.WorkloadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadstatus.WorkloadTable,
			Columns: []string{workloadstatus.WorkloadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workload.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wsuo.mutation.WorkloadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workloadstatus.WorkloadTable,
			Columns: []string{workloadstatus.WorkloadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workload.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkloadStatus{config: wsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workloadstatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}