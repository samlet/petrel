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
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappl"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappltype"
)

// WorkloadFeatureApplTypeQuery is the builder for querying WorkloadFeatureApplType entities.
type WorkloadFeatureApplTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkloadFeatureApplType
	// eager-loading edges.
	withParent               *WorkloadFeatureApplTypeQuery
	withChildren             *WorkloadFeatureApplTypeQuery
	withWorkloadFeatureAppls *WorkloadFeatureApplQuery
	withFKs                  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkloadFeatureApplTypeQuery builder.
func (wfatq *WorkloadFeatureApplTypeQuery) Where(ps ...predicate.WorkloadFeatureApplType) *WorkloadFeatureApplTypeQuery {
	wfatq.predicates = append(wfatq.predicates, ps...)
	return wfatq
}

// Limit adds a limit step to the query.
func (wfatq *WorkloadFeatureApplTypeQuery) Limit(limit int) *WorkloadFeatureApplTypeQuery {
	wfatq.limit = &limit
	return wfatq
}

// Offset adds an offset step to the query.
func (wfatq *WorkloadFeatureApplTypeQuery) Offset(offset int) *WorkloadFeatureApplTypeQuery {
	wfatq.offset = &offset
	return wfatq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wfatq *WorkloadFeatureApplTypeQuery) Unique(unique bool) *WorkloadFeatureApplTypeQuery {
	wfatq.unique = &unique
	return wfatq
}

// Order adds an order step to the query.
func (wfatq *WorkloadFeatureApplTypeQuery) Order(o ...OrderFunc) *WorkloadFeatureApplTypeQuery {
	wfatq.order = append(wfatq.order, o...)
	return wfatq
}

// QueryParent chains the current query on the "parent" edge.
func (wfatq *WorkloadFeatureApplTypeQuery) QueryParent() *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID, selector),
			sqlgraph.To(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workloadfeatureappltype.ParentTable, workloadfeatureappltype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (wfatq *WorkloadFeatureApplTypeQuery) QueryChildren() *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID, selector),
			sqlgraph.To(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workloadfeatureappltype.ChildrenTable, workloadfeatureappltype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkloadFeatureAppls chains the current query on the "workload_feature_appls" edge.
func (wfatq *WorkloadFeatureApplTypeQuery) QueryWorkloadFeatureAppls() *WorkloadFeatureApplQuery {
	query := &WorkloadFeatureApplQuery{config: wfatq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfatq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID, selector),
			sqlgraph.To(workloadfeatureappl.Table, workloadfeatureappl.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workloadfeatureappltype.WorkloadFeatureApplsTable, workloadfeatureappltype.WorkloadFeatureApplsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfatq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkloadFeatureApplType entity from the query.
// Returns a *NotFoundError when no WorkloadFeatureApplType was found.
func (wfatq *WorkloadFeatureApplTypeQuery) First(ctx context.Context) (*WorkloadFeatureApplType, error) {
	nodes, err := wfatq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workloadfeatureappltype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) FirstX(ctx context.Context) *WorkloadFeatureApplType {
	node, err := wfatq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkloadFeatureApplType ID from the query.
// Returns a *NotFoundError when no WorkloadFeatureApplType ID was found.
func (wfatq *WorkloadFeatureApplTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfatq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workloadfeatureappltype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := wfatq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkloadFeatureApplType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WorkloadFeatureApplType entity is not found.
// Returns a *NotFoundError when no WorkloadFeatureApplType entities are found.
func (wfatq *WorkloadFeatureApplTypeQuery) Only(ctx context.Context) (*WorkloadFeatureApplType, error) {
	nodes, err := wfatq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workloadfeatureappltype.Label}
	default:
		return nil, &NotSingularError{workloadfeatureappltype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) OnlyX(ctx context.Context) *WorkloadFeatureApplType {
	node, err := wfatq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkloadFeatureApplType ID in the query.
// Returns a *NotSingularError when exactly one WorkloadFeatureApplType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wfatq *WorkloadFeatureApplTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfatq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = &NotSingularError{workloadfeatureappltype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := wfatq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkloadFeatureApplTypes.
func (wfatq *WorkloadFeatureApplTypeQuery) All(ctx context.Context) ([]*WorkloadFeatureApplType, error) {
	if err := wfatq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wfatq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) AllX(ctx context.Context) []*WorkloadFeatureApplType {
	nodes, err := wfatq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkloadFeatureApplType IDs.
func (wfatq *WorkloadFeatureApplTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wfatq.Select(workloadfeatureappltype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := wfatq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wfatq *WorkloadFeatureApplTypeQuery) Count(ctx context.Context) (int, error) {
	if err := wfatq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wfatq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) CountX(ctx context.Context) int {
	count, err := wfatq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wfatq *WorkloadFeatureApplTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wfatq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wfatq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wfatq *WorkloadFeatureApplTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := wfatq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkloadFeatureApplTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wfatq *WorkloadFeatureApplTypeQuery) Clone() *WorkloadFeatureApplTypeQuery {
	if wfatq == nil {
		return nil
	}
	return &WorkloadFeatureApplTypeQuery{
		config:                   wfatq.config,
		limit:                    wfatq.limit,
		offset:                   wfatq.offset,
		order:                    append([]OrderFunc{}, wfatq.order...),
		predicates:               append([]predicate.WorkloadFeatureApplType{}, wfatq.predicates...),
		withParent:               wfatq.withParent.Clone(),
		withChildren:             wfatq.withChildren.Clone(),
		withWorkloadFeatureAppls: wfatq.withWorkloadFeatureAppls.Clone(),
		// clone intermediate query.
		sql:  wfatq.sql.Clone(),
		path: wfatq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (wfatq *WorkloadFeatureApplTypeQuery) WithParent(opts ...func(*WorkloadFeatureApplTypeQuery)) *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfatq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfatq.withParent = query
	return wfatq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (wfatq *WorkloadFeatureApplTypeQuery) WithChildren(opts ...func(*WorkloadFeatureApplTypeQuery)) *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfatq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfatq.withChildren = query
	return wfatq
}

// WithWorkloadFeatureAppls tells the query-builder to eager-load the nodes that are connected to
// the "workload_feature_appls" edge. The optional arguments are used to configure the query builder of the edge.
func (wfatq *WorkloadFeatureApplTypeQuery) WithWorkloadFeatureAppls(opts ...func(*WorkloadFeatureApplQuery)) *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplQuery{config: wfatq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfatq.withWorkloadFeatureAppls = query
	return wfatq
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
//	client.WorkloadFeatureApplType.Query().
//		GroupBy(workloadfeatureappltype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wfatq *WorkloadFeatureApplTypeQuery) GroupBy(field string, fields ...string) *WorkloadFeatureApplTypeGroupBy {
	group := &WorkloadFeatureApplTypeGroupBy{config: wfatq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wfatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wfatq.sqlQuery(ctx), nil
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
//	client.WorkloadFeatureApplType.Query().
//		Select(workloadfeatureappltype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wfatq *WorkloadFeatureApplTypeQuery) Select(field string, fields ...string) *WorkloadFeatureApplTypeSelect {
	wfatq.fields = append([]string{field}, fields...)
	return &WorkloadFeatureApplTypeSelect{WorkloadFeatureApplTypeQuery: wfatq}
}

func (wfatq *WorkloadFeatureApplTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wfatq.fields {
		if !workloadfeatureappltype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wfatq.path != nil {
		prev, err := wfatq.path(ctx)
		if err != nil {
			return err
		}
		wfatq.sql = prev
	}
	return nil
}

func (wfatq *WorkloadFeatureApplTypeQuery) sqlAll(ctx context.Context) ([]*WorkloadFeatureApplType, error) {
	var (
		nodes       = []*WorkloadFeatureApplType{}
		withFKs     = wfatq.withFKs
		_spec       = wfatq.querySpec()
		loadedTypes = [3]bool{
			wfatq.withParent != nil,
			wfatq.withChildren != nil,
			wfatq.withWorkloadFeatureAppls != nil,
		}
	)
	if wfatq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workloadfeatureappltype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WorkloadFeatureApplType{config: wfatq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wfatq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wfatq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkloadFeatureApplType)
		for i := range nodes {
			if nodes[i].workload_feature_appl_type_children == nil {
				continue
			}
			fk := *nodes[i].workload_feature_appl_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workloadfeatureappltype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_feature_appl_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := wfatq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkloadFeatureApplType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*WorkloadFeatureApplType{}
		}
		query.withFKs = true
		query.Where(predicate.WorkloadFeatureApplType(func(s *sql.Selector) {
			s.Where(sql.InValues(workloadfeatureappltype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.workload_feature_appl_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "workload_feature_appl_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_feature_appl_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := wfatq.withWorkloadFeatureAppls; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkloadFeatureApplType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.WorkloadFeatureAppls = []*WorkloadFeatureAppl{}
		}
		query.withFKs = true
		query.Where(predicate.WorkloadFeatureAppl(func(s *sql.Selector) {
			s.Where(sql.InValues(workloadfeatureappltype.WorkloadFeatureApplsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.workload_feature_appl_type_workload_feature_appls
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "workload_feature_appl_type_workload_feature_appls" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_feature_appl_type_workload_feature_appls" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.WorkloadFeatureAppls = append(node.Edges.WorkloadFeatureAppls, n)
		}
	}

	return nodes, nil
}

func (wfatq *WorkloadFeatureApplTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wfatq.querySpec()
	return sqlgraph.CountNodes(ctx, wfatq.driver, _spec)
}

func (wfatq *WorkloadFeatureApplTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wfatq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wfatq *WorkloadFeatureApplTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workloadfeatureappltype.Table,
			Columns: workloadfeatureappltype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadfeatureappltype.FieldID,
			},
		},
		From:   wfatq.sql,
		Unique: true,
	}
	if unique := wfatq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wfatq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workloadfeatureappltype.FieldID)
		for i := range fields {
			if fields[i] != workloadfeatureappltype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wfatq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wfatq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wfatq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wfatq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wfatq *WorkloadFeatureApplTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wfatq.driver.Dialect())
	t1 := builder.Table(workloadfeatureappltype.Table)
	selector := builder.Select(t1.Columns(workloadfeatureappltype.Columns...)...).From(t1)
	if wfatq.sql != nil {
		selector = wfatq.sql
		selector.Select(selector.Columns(workloadfeatureappltype.Columns...)...)
	}
	for _, p := range wfatq.predicates {
		p(selector)
	}
	for _, p := range wfatq.order {
		p(selector)
	}
	if offset := wfatq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wfatq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkloadFeatureApplTypeGroupBy is the group-by builder for WorkloadFeatureApplType entities.
type WorkloadFeatureApplTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Aggregate(fns ...AggregateFunc) *WorkloadFeatureApplTypeGroupBy {
	wfatgb.fns = append(wfatgb.fns, fns...)
	return wfatgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wfatgb.path(ctx)
	if err != nil {
		return err
	}
	wfatgb.sql = query
	return wfatgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wfatgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wfatgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wfatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wfatgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfatgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) StringX(ctx context.Context) string {
	v, err := wfatgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wfatgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wfatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wfatgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfatgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) IntX(ctx context.Context) int {
	v, err := wfatgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfatgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wfatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wfatgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfatgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wfatgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wfatgb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wfatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wfatgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfatgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfatgb *WorkloadFeatureApplTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wfatgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfatgb *WorkloadFeatureApplTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wfatgb.fields {
		if !workloadfeatureappltype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wfatgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfatgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfatgb *WorkloadFeatureApplTypeGroupBy) sqlQuery() *sql.Selector {
	selector := wfatgb.sql
	columns := make([]string, 0, len(wfatgb.fields)+len(wfatgb.fns))
	columns = append(columns, wfatgb.fields...)
	for _, fn := range wfatgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wfatgb.fields...)
}

// WorkloadFeatureApplTypeSelect is the builder for selecting fields of WorkloadFeatureApplType entities.
type WorkloadFeatureApplTypeSelect struct {
	*WorkloadFeatureApplTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wfats *WorkloadFeatureApplTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wfats.prepareQuery(ctx); err != nil {
		return err
	}
	wfats.sql = wfats.WorkloadFeatureApplTypeQuery.sqlQuery(ctx)
	return wfats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wfats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wfats.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wfats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) StringsX(ctx context.Context) []string {
	v, err := wfats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) StringX(ctx context.Context) string {
	v, err := wfats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wfats.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wfats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) IntsX(ctx context.Context) []int {
	v, err := wfats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) IntX(ctx context.Context) int {
	v, err := wfats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfats.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wfats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wfats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := wfats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wfats.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wfats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := wfats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wfats *WorkloadFeatureApplTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappltype.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfats *WorkloadFeatureApplTypeSelect) BoolX(ctx context.Context) bool {
	v, err := wfats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfats *WorkloadFeatureApplTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wfats.sqlQuery().Query()
	if err := wfats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfats *WorkloadFeatureApplTypeSelect) sqlQuery() sql.Querier {
	selector := wfats.sql
	selector.Select(selector.Columns(wfats.fields...)...)
	return selector
}
