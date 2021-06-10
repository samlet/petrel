// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workload/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeature"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappl"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappltype"
)

// WorkloadFeatureApplQuery is the builder for querying WorkloadFeatureAppl entities.
type WorkloadFeatureApplQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkloadFeatureAppl
	// eager-loading edges.
	withWorkload                *WorkloadQuery
	withWorkloadFeature         *WorkloadFeatureQuery
	withWorkloadFeatureApplType *WorkloadFeatureApplTypeQuery
	withFKs                     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkloadFeatureApplQuery builder.
func (wfaq *WorkloadFeatureApplQuery) Where(ps ...predicate.WorkloadFeatureAppl) *WorkloadFeatureApplQuery {
	wfaq.predicates = append(wfaq.predicates, ps...)
	return wfaq
}

// Limit adds a limit step to the query.
func (wfaq *WorkloadFeatureApplQuery) Limit(limit int) *WorkloadFeatureApplQuery {
	wfaq.limit = &limit
	return wfaq
}

// Offset adds an offset step to the query.
func (wfaq *WorkloadFeatureApplQuery) Offset(offset int) *WorkloadFeatureApplQuery {
	wfaq.offset = &offset
	return wfaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wfaq *WorkloadFeatureApplQuery) Unique(unique bool) *WorkloadFeatureApplQuery {
	wfaq.unique = &unique
	return wfaq
}

// Order adds an order step to the query.
func (wfaq *WorkloadFeatureApplQuery) Order(o ...OrderFunc) *WorkloadFeatureApplQuery {
	wfaq.order = append(wfaq.order, o...)
	return wfaq
}

// QueryWorkload chains the current query on the "workload" edge.
func (wfaq *WorkloadFeatureApplQuery) QueryWorkload() *WorkloadQuery {
	query := &WorkloadQuery{config: wfaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappl.Table, workloadfeatureappl.FieldID, selector),
			sqlgraph.To(workload.Table, workload.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workloadfeatureappl.WorkloadTable, workloadfeatureappl.WorkloadColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkloadFeature chains the current query on the "workload_feature" edge.
func (wfaq *WorkloadFeatureApplQuery) QueryWorkloadFeature() *WorkloadFeatureQuery {
	query := &WorkloadFeatureQuery{config: wfaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappl.Table, workloadfeatureappl.FieldID, selector),
			sqlgraph.To(workloadfeature.Table, workloadfeature.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workloadfeatureappl.WorkloadFeatureTable, workloadfeatureappl.WorkloadFeatureColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkloadFeatureApplType chains the current query on the "workload_feature_appl_type" edge.
func (wfaq *WorkloadFeatureApplQuery) QueryWorkloadFeatureApplType() *WorkloadFeatureApplTypeQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wfaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wfaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workloadfeatureappl.Table, workloadfeatureappl.FieldID, selector),
			sqlgraph.To(workloadfeatureappltype.Table, workloadfeatureappltype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workloadfeatureappl.WorkloadFeatureApplTypeTable, workloadfeatureappl.WorkloadFeatureApplTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(wfaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkloadFeatureAppl entity from the query.
// Returns a *NotFoundError when no WorkloadFeatureAppl was found.
func (wfaq *WorkloadFeatureApplQuery) First(ctx context.Context) (*WorkloadFeatureAppl, error) {
	nodes, err := wfaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workloadfeatureappl.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) FirstX(ctx context.Context) *WorkloadFeatureAppl {
	node, err := wfaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkloadFeatureAppl ID from the query.
// Returns a *NotFoundError when no WorkloadFeatureAppl ID was found.
func (wfaq *WorkloadFeatureApplQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workloadfeatureappl.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) FirstIDX(ctx context.Context) int {
	id, err := wfaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkloadFeatureAppl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WorkloadFeatureAppl entity is not found.
// Returns a *NotFoundError when no WorkloadFeatureAppl entities are found.
func (wfaq *WorkloadFeatureApplQuery) Only(ctx context.Context) (*WorkloadFeatureAppl, error) {
	nodes, err := wfaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workloadfeatureappl.Label}
	default:
		return nil, &NotSingularError{workloadfeatureappl.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) OnlyX(ctx context.Context) *WorkloadFeatureAppl {
	node, err := wfaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkloadFeatureAppl ID in the query.
// Returns a *NotSingularError when exactly one WorkloadFeatureAppl ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wfaq *WorkloadFeatureApplQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = &NotSingularError{workloadfeatureappl.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) OnlyIDX(ctx context.Context) int {
	id, err := wfaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkloadFeatureAppls.
func (wfaq *WorkloadFeatureApplQuery) All(ctx context.Context) ([]*WorkloadFeatureAppl, error) {
	if err := wfaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wfaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) AllX(ctx context.Context) []*WorkloadFeatureAppl {
	nodes, err := wfaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkloadFeatureAppl IDs.
func (wfaq *WorkloadFeatureApplQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wfaq.Select(workloadfeatureappl.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) IDsX(ctx context.Context) []int {
	ids, err := wfaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wfaq *WorkloadFeatureApplQuery) Count(ctx context.Context) (int, error) {
	if err := wfaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wfaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) CountX(ctx context.Context) int {
	count, err := wfaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wfaq *WorkloadFeatureApplQuery) Exist(ctx context.Context) (bool, error) {
	if err := wfaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wfaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wfaq *WorkloadFeatureApplQuery) ExistX(ctx context.Context) bool {
	exist, err := wfaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkloadFeatureApplQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wfaq *WorkloadFeatureApplQuery) Clone() *WorkloadFeatureApplQuery {
	if wfaq == nil {
		return nil
	}
	return &WorkloadFeatureApplQuery{
		config:                      wfaq.config,
		limit:                       wfaq.limit,
		offset:                      wfaq.offset,
		order:                       append([]OrderFunc{}, wfaq.order...),
		predicates:                  append([]predicate.WorkloadFeatureAppl{}, wfaq.predicates...),
		withWorkload:                wfaq.withWorkload.Clone(),
		withWorkloadFeature:         wfaq.withWorkloadFeature.Clone(),
		withWorkloadFeatureApplType: wfaq.withWorkloadFeatureApplType.Clone(),
		// clone intermediate query.
		sql:  wfaq.sql.Clone(),
		path: wfaq.path,
	}
}

// WithWorkload tells the query-builder to eager-load the nodes that are connected to
// the "workload" edge. The optional arguments are used to configure the query builder of the edge.
func (wfaq *WorkloadFeatureApplQuery) WithWorkload(opts ...func(*WorkloadQuery)) *WorkloadFeatureApplQuery {
	query := &WorkloadQuery{config: wfaq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfaq.withWorkload = query
	return wfaq
}

// WithWorkloadFeature tells the query-builder to eager-load the nodes that are connected to
// the "workload_feature" edge. The optional arguments are used to configure the query builder of the edge.
func (wfaq *WorkloadFeatureApplQuery) WithWorkloadFeature(opts ...func(*WorkloadFeatureQuery)) *WorkloadFeatureApplQuery {
	query := &WorkloadFeatureQuery{config: wfaq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfaq.withWorkloadFeature = query
	return wfaq
}

// WithWorkloadFeatureApplType tells the query-builder to eager-load the nodes that are connected to
// the "workload_feature_appl_type" edge. The optional arguments are used to configure the query builder of the edge.
func (wfaq *WorkloadFeatureApplQuery) WithWorkloadFeatureApplType(opts ...func(*WorkloadFeatureApplTypeQuery)) *WorkloadFeatureApplQuery {
	query := &WorkloadFeatureApplTypeQuery{config: wfaq.config}
	for _, opt := range opts {
		opt(query)
	}
	wfaq.withWorkloadFeatureApplType = query
	return wfaq
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
//	client.WorkloadFeatureAppl.Query().
//		GroupBy(workloadfeatureappl.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wfaq *WorkloadFeatureApplQuery) GroupBy(field string, fields ...string) *WorkloadFeatureApplGroupBy {
	group := &WorkloadFeatureApplGroupBy{config: wfaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wfaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wfaq.sqlQuery(ctx), nil
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
//	client.WorkloadFeatureAppl.Query().
//		Select(workloadfeatureappl.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wfaq *WorkloadFeatureApplQuery) Select(field string, fields ...string) *WorkloadFeatureApplSelect {
	wfaq.fields = append([]string{field}, fields...)
	return &WorkloadFeatureApplSelect{WorkloadFeatureApplQuery: wfaq}
}

func (wfaq *WorkloadFeatureApplQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wfaq.fields {
		if !workloadfeatureappl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wfaq.path != nil {
		prev, err := wfaq.path(ctx)
		if err != nil {
			return err
		}
		wfaq.sql = prev
	}
	return nil
}

func (wfaq *WorkloadFeatureApplQuery) sqlAll(ctx context.Context) ([]*WorkloadFeatureAppl, error) {
	var (
		nodes       = []*WorkloadFeatureAppl{}
		withFKs     = wfaq.withFKs
		_spec       = wfaq.querySpec()
		loadedTypes = [3]bool{
			wfaq.withWorkload != nil,
			wfaq.withWorkloadFeature != nil,
			wfaq.withWorkloadFeatureApplType != nil,
		}
	)
	if wfaq.withWorkload != nil || wfaq.withWorkloadFeature != nil || wfaq.withWorkloadFeatureApplType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workloadfeatureappl.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WorkloadFeatureAppl{config: wfaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wfaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wfaq.withWorkload; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkloadFeatureAppl)
		for i := range nodes {
			if nodes[i].workload_workload_feature_appls == nil {
				continue
			}
			fk := *nodes[i].workload_workload_feature_appls
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workload.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_workload_feature_appls" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Workload = n
			}
		}
	}

	if query := wfaq.withWorkloadFeature; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkloadFeatureAppl)
		for i := range nodes {
			if nodes[i].workload_feature_workload_feature_appls == nil {
				continue
			}
			fk := *nodes[i].workload_feature_workload_feature_appls
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workloadfeature.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "workload_feature_workload_feature_appls" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.WorkloadFeature = n
			}
		}
	}

	if query := wfaq.withWorkloadFeatureApplType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkloadFeatureAppl)
		for i := range nodes {
			if nodes[i].workload_feature_appl_type_workload_feature_appls == nil {
				continue
			}
			fk := *nodes[i].workload_feature_appl_type_workload_feature_appls
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
				return nil, fmt.Errorf(`unexpected foreign-key "workload_feature_appl_type_workload_feature_appls" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.WorkloadFeatureApplType = n
			}
		}
	}

	return nodes, nil
}

func (wfaq *WorkloadFeatureApplQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wfaq.querySpec()
	return sqlgraph.CountNodes(ctx, wfaq.driver, _spec)
}

func (wfaq *WorkloadFeatureApplQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wfaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wfaq *WorkloadFeatureApplQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workloadfeatureappl.Table,
			Columns: workloadfeatureappl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workloadfeatureappl.FieldID,
			},
		},
		From:   wfaq.sql,
		Unique: true,
	}
	if unique := wfaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wfaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workloadfeatureappl.FieldID)
		for i := range fields {
			if fields[i] != workloadfeatureappl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wfaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wfaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wfaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wfaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wfaq *WorkloadFeatureApplQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wfaq.driver.Dialect())
	t1 := builder.Table(workloadfeatureappl.Table)
	selector := builder.Select(t1.Columns(workloadfeatureappl.Columns...)...).From(t1)
	if wfaq.sql != nil {
		selector = wfaq.sql
		selector.Select(selector.Columns(workloadfeatureappl.Columns...)...)
	}
	for _, p := range wfaq.predicates {
		p(selector)
	}
	for _, p := range wfaq.order {
		p(selector)
	}
	if offset := wfaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wfaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkloadFeatureApplGroupBy is the group-by builder for WorkloadFeatureAppl entities.
type WorkloadFeatureApplGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wfagb *WorkloadFeatureApplGroupBy) Aggregate(fns ...AggregateFunc) *WorkloadFeatureApplGroupBy {
	wfagb.fns = append(wfagb.fns, fns...)
	return wfagb
}

// Scan applies the group-by query and scans the result into the given value.
func (wfagb *WorkloadFeatureApplGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wfagb.path(ctx)
	if err != nil {
		return err
	}
	wfagb.sql = query
	return wfagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wfagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wfagb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wfagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) StringsX(ctx context.Context) []string {
	v, err := wfagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) StringX(ctx context.Context) string {
	v, err := wfagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wfagb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wfagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) IntsX(ctx context.Context) []int {
	v, err := wfagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) IntX(ctx context.Context) int {
	v, err := wfagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfagb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wfagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wfagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wfagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wfagb.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wfagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wfagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfagb *WorkloadFeatureApplGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfagb *WorkloadFeatureApplGroupBy) BoolX(ctx context.Context) bool {
	v, err := wfagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfagb *WorkloadFeatureApplGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wfagb.fields {
		if !workloadfeatureappl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wfagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfagb *WorkloadFeatureApplGroupBy) sqlQuery() *sql.Selector {
	selector := wfagb.sql
	columns := make([]string, 0, len(wfagb.fields)+len(wfagb.fns))
	columns = append(columns, wfagb.fields...)
	for _, fn := range wfagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wfagb.fields...)
}

// WorkloadFeatureApplSelect is the builder for selecting fields of WorkloadFeatureAppl entities.
type WorkloadFeatureApplSelect struct {
	*WorkloadFeatureApplQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wfas *WorkloadFeatureApplSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wfas.prepareQuery(ctx); err != nil {
		return err
	}
	wfas.sql = wfas.WorkloadFeatureApplQuery.sqlQuery(ctx)
	return wfas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wfas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wfas.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wfas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) StringsX(ctx context.Context) []string {
	v, err := wfas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) StringX(ctx context.Context) string {
	v, err := wfas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wfas.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wfas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) IntsX(ctx context.Context) []int {
	v, err := wfas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) IntX(ctx context.Context) int {
	v, err := wfas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfas.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wfas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wfas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) Float64X(ctx context.Context) float64 {
	v, err := wfas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wfas.fields) > 1 {
		return nil, errors.New("ent: WorkloadFeatureApplSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wfas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) BoolsX(ctx context.Context) []bool {
	v, err := wfas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wfas *WorkloadFeatureApplSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workloadfeatureappl.Label}
	default:
		err = fmt.Errorf("ent: WorkloadFeatureApplSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfas *WorkloadFeatureApplSelect) BoolX(ctx context.Context) bool {
	v, err := wfas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfas *WorkloadFeatureApplSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wfas.sqlQuery().Query()
	if err := wfas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfas *WorkloadFeatureApplSelect) sqlQuery() sql.Querier {
	selector := wfas.sql
	selector.Select(selector.Columns(wfas.fields...)...)
	return selector
}