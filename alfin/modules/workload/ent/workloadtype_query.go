// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadtype"
)

// WorkloadTypeQuery is the builder for querying WorkloadType entities.
type WorkloadTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkloadType
	// eager-loading edges.
	withParent    *WorkloadTypeQuery
	withChildren  *WorkloadTypeQuery
	withWorkloads *WorkloadQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkloadTypeQuery builder.
func (wtq *WorkloadTypeQuery) Where(ps ...predicate.WorkloadType) *WorkloadTypeQuery {
	wtq.predicates = append(wtq.predicates, ps...)
	return wtq
}

// Limit adds a limit step to the query.
func (wtq *WorkloadTypeQuery) Limit(limit int) *WorkloadTypeQuery {
	wtq.limit = &limit
	return wtq
}

// Offset adds an offset step to the query.
func (wtq *WorkloadTypeQuery) Offset(offset int) *WorkloadTypeQuery {
	wtq.offset = &offset
	return wtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wtq *WorkloadTypeQuery) Unique(unique bool) *WorkloadTypeQuery {
	wtq.unique = &unique
	return wtq
}

// Order adds an order step to the query.
func (wtq *WorkloadTypeQuery) Order(o ...OrderFunc) *WorkloadTypeQuery {
	wtq.order = append(wtq.order, o...)
	return wtq
}

// QueryParent chains the current query on the "parent" edge.
func (wtq *WorkloadTypeQuery) QueryParent() *WorkloadTypeQuery {
	query := &WorkloadTypeQuery{config: wtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadtype.Table, workloadtype.FieldID, selector),
			sqlgraph.To(workloadtype.Table, workloadtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workloadtype.ParentTable, workloadtype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(wtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (wtq *WorkloadTypeQuery) QueryChildren() *WorkloadTypeQuery {
	query := &WorkloadTypeQuery{config: wtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadtype.Table, workloadtype.FieldID, selector),
			sqlgraph.To(workloadtype.Table, workloadtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workloadtype.ChildrenTable, workloadtype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(wtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkloads chains the current query on the "workloads" edge.
func (wtq *WorkloadTypeQuery) QueryWorkloads() *WorkloadQuery {
	query := &WorkloadQuery{config: wtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadtype.Table, workloadtype.FieldID, selector),
			sqlgraph.To(workload.Table, workload.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workloadtype.WorkloadsTable, workloadtype.WorkloadsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkloadType entity from the query.
// Returns a *NotFoundError when no WorkloadType was found.
func (wtq *WorkloadTypeQuery) First(ctx context.Context) (*WorkloadType, error) {
	nodes, err := wtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workloadtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) FirstX(ctx context.Context) *WorkloadType {
	node, err := wtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkloadType ID from the query.
// Returns a *NotFoundError when no WorkloadType ID was found.
func (wtq *WorkloadTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workloadtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := wtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkloadType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WorkloadType entity is not found.
// Returns a *NotFoundError when no WorkloadType entities are found.
func (wtq *WorkloadTypeQuery) Only(ctx context.Context) (*WorkloadType, error) {
	nodes, err := wtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workloadtype.Label}
	default:
		return nil, &NotSingularError{workloadtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) OnlyX(ctx context.Context) *WorkloadType {
	node, err := wtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkloadType ID in the query.
// Returns a *NotSingularError when exactly one WorkloadType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wtq *WorkloadTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = &NotSingularError{workloadtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := wtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkloadTypes.
func (wtq *WorkloadTypeQuery) All(ctx context.Context) ([]*WorkloadType, error) {
	if err := wtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) AllX(ctx context.Context) []*WorkloadType {
	nodes, err := wtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkloadType IDs.
func (wtq *WorkloadTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wtq.Select(workloadtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := wtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wtq *WorkloadTypeQuery) Count(ctx context.Context) (int, error) {
	if err := wtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) CountX(ctx context.Context) int {
	count, err := wtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wtq *WorkloadTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wtq *WorkloadTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := wtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkloadTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wtq *WorkloadTypeQuery) Clone() *WorkloadTypeQuery {
	if wtq == nil {
		return nil
	}
	return &WorkloadTypeQuery{
		config:        wtq.config,
		limit:         wtq.limit,
		offset:        wtq.offset,
		order:         append([]OrderFunc{}, wtq.order...),
		predicates:    append([]predicate.WorkloadType{}, wtq.predicates...),
		withParent:    wtq.withParent.Clone(),
		withChildren:  wtq.withChildren.Clone(),
		withWorkloads: wtq.withWorkloads.Clone(),
		// clone intermediate query.
		sql:  wtq.sql.Clone(),
		path: wtq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (wtq *WorkloadTypeQuery) WithParent(opts ...func(*WorkloadTypeQuery)) *WorkloadTypeQuery {
	query := &WorkloadTypeQuery{config: wtq.config}
	for _, opt := range opts {
		opt(query)
	}
	wtq.withParent = query
	return wtq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (wtq *WorkloadTypeQuery) WithChildren(opts ...func(*WorkloadTypeQuery)) *WorkloadTypeQuery {
	query := &WorkloadTypeQuery{config: wtq.config}
	for _, opt := range opts {
		opt(query)
	}
	wtq.withChildren = query
	return wtq
}

// WithWorkloads tells the query-builder to eager-load the nodes that are connected to
// the "workloads" edge. The optional arguments are used to configure the query builder of the edge.
func (wtq *WorkloadTypeQuery) WithWorkloads(opts ...func(*WorkloadQuery)) *WorkloadTypeQuery {
	query := &WorkloadQuery{config: wtq.config}
	for _, opt := range opts {
		opt(query)
	}
	wtq.withWorkloads = query
	return wtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkloadType.Query().
//		GroupBy(workloadtype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wtq *WorkloadTypeQuery) GroupBy(field string, fields ...string) *WorkloadTypeGroupBy {
	group := &WorkloadTypeGroupBy{config: wtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wtq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.WorkloadType.Query().
//		Select(workloadtype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wtq *WorkloadTypeQuery) Select(field string, fields ...string) *WorkloadTypeSelect {
	wtq.fields = append([]string{field}, fields...)
	return &WorkloadTypeSelect{WorkloadTypeQuery: wtq}
}

func (wtq *WorkloadTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wtq.fields {
		if !workloadtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wtq.path != nil {
		prev, err := wtq.path(ctx)
		if err != nil {
			return err
		}
		wtq.sql = prev
	}
	return nil
}

func (wtq *WorkloadTypeQuery) sqlAll(ctx context.Context) ([]*WorkloadType, error) {
	var (
		nodes       = []*WorkloadType{}
		withFKs     = wtq.withFKs
		_spec       = wtq.querySpec()
		loadedTypes = [3]bool{
			wtq.withParent != nil,
			wtq.withChildren != nil,
			wtq.withWorkloads != nil,
		}
	)
	if wtq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workloadtype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WorkloadType{config: wtq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, wtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wtq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkloadType)
		for i := range nodes {
			if nodes[i].workload_type_children == nil {
				continue
			}
			fk := *nodes[i].workload_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workloadtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := wtq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkloadType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*WorkloadType{}
		}
		query.withFKs = true
		query.Where(predicate.WorkloadType(func(s *sql.Selector) {
			s.Where(sql.InValues(workloadtype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.workload_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "workload_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := wtq.withWorkloads; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkloadType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Workloads = []*Workload{}
		}
		query.withFKs = true
		query.Where(predicate.Workload(func(s *sql.Selector) {
			s.Where(sql.InValues(workloadtype.WorkloadsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.workload_type_workloads
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "workload_type_workloads" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_type_workloads" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Workloads = append(node.Edges.Workloads, n)
		}
	}

	return nodes, nil
}

func (wtq *WorkloadTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wtq.querySpec()
	return sqlgraph.CountNodes(ctx, wtq.driver, _spec)
}

func (wtq *WorkloadTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wtq *WorkloadTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workloadtype.Table,
			Columns: workloadtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadtype.FieldID,
			},
		},
		From:   wtq.sql,
		Unique: true,
	}
	if unique := wtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workloadtype.FieldID)
		for i := range fields {
			if fields[i] != workloadtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wtq *WorkloadTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wtq.driver.Dialect())
	t1 := builder.Table(workloadtype.Table)
	selector := builder.Select(t1.Columns(workloadtype.Columns...)...).From(t1)
	if wtq.sql != nil {
		selector = wtq.sql
		selector.Select(selector.Columns(workloadtype.Columns...)...)
	}
	for _, p := range wtq.predicates {
		p(selector)
	}
	for _, p := range wtq.order {
		p(selector)
	}
	if offset := wtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkloadTypeGroupBy is the group-by builder for WorkloadType entities.
type WorkloadTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wtgb *WorkloadTypeGroupBy) Aggregate(fns ...AggregateFunc) *WorkloadTypeGroupBy {
	wtgb.fns = append(wtgb.fns, fns...)
	return wtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wtgb *WorkloadTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wtgb.path(ctx)
	if err != nil {
		return err
	}
	wtgb.sql = query
	return wtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wtgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) StringX(ctx context.Context) string {
	v, err := wtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wtgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) IntX(ctx context.Context) int {
	v, err := wtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wtgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wtgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wtgb *WorkloadTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wtgb *WorkloadTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wtgb *WorkloadTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wtgb.fields {
		if !workloadtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wtgb *WorkloadTypeGroupBy) sqlQuery() *sql.Selector {
	selector := wtgb.sql
	columns := make([]string, 0, len(wtgb.fields)+len(wtgb.fns))
	columns = append(columns, wtgb.fields...)
	for _, fn := range wtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wtgb.fields...)
}

// WorkloadTypeSelect is the builder for selecting fields of WorkloadType entities.
type WorkloadTypeSelect struct {
	*WorkloadTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wts *WorkloadTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wts.prepareQuery(ctx); err != nil {
		return err
	}
	wts.sql = wts.WorkloadTypeQuery.sqlQuery(ctx)
	return wts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wts *WorkloadTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wts.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wts *WorkloadTypeSelect) StringsX(ctx context.Context) []string {
	v, err := wts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wts *WorkloadTypeSelect) StringX(ctx context.Context) string {
	v, err := wts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wts.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wts *WorkloadTypeSelect) IntsX(ctx context.Context) []int {
	v, err := wts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wts *WorkloadTypeSelect) IntX(ctx context.Context) int {
	v, err := wts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wts.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wts *WorkloadTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wts *WorkloadTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := wts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wts.fields) > 1 {
		return nil, errors.New("ent: WorkloadTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wts *WorkloadTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := wts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wts *WorkloadTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadtype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wts *WorkloadTypeSelect) BoolX(ctx context.Context) bool {
	v, err := wts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wts *WorkloadTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wts.sqlQuery().Query()
	if err := wts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wts *WorkloadTypeSelect) sqlQuery() sql.Querier {
	selector := wts.sql
	selector.Select(selector.Columns(wts.fields...)...)
	return selector
}
