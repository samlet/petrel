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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusvalidchange"
)

// StatusValidChangeQuery is the builder for querying StatusValidChange entities.
type StatusValidChangeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StatusValidChange
	// eager-loading edges.
	withMainStatusItem *StatusItemQuery
	withToStatusItem   *StatusItemQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StatusValidChangeQuery builder.
func (svcq *StatusValidChangeQuery) Where(ps ...predicate.StatusValidChange) *StatusValidChangeQuery {
	svcq.predicates = append(svcq.predicates, ps...)
	return svcq
}

// Limit adds a limit step to the query.
func (svcq *StatusValidChangeQuery) Limit(limit int) *StatusValidChangeQuery {
	svcq.limit = &limit
	return svcq
}

// Offset adds an offset step to the query.
func (svcq *StatusValidChangeQuery) Offset(offset int) *StatusValidChangeQuery {
	svcq.offset = &offset
	return svcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (svcq *StatusValidChangeQuery) Unique(unique bool) *StatusValidChangeQuery {
	svcq.unique = &unique
	return svcq
}

// Order adds an order step to the query.
func (svcq *StatusValidChangeQuery) Order(o ...OrderFunc) *StatusValidChangeQuery {
	svcq.order = append(svcq.order, o...)
	return svcq
}

// QueryMainStatusItem chains the current query on the "main_status_item" edge.
func (svcq *StatusValidChangeQuery) QueryMainStatusItem() *StatusItemQuery {
	query := &StatusItemQuery{config: svcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := svcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := svcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statusvalidchange.Table, statusvalidchange.FieldID, selector),
			sqlgraph.To(statusitem.Table, statusitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, statusvalidchange.MainStatusItemTable, statusvalidchange.MainStatusItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(svcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryToStatusItem chains the current query on the "to_status_item" edge.
func (svcq *StatusValidChangeQuery) QueryToStatusItem() *StatusItemQuery {
	query := &StatusItemQuery{config: svcq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := svcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := svcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statusvalidchange.Table, statusvalidchange.FieldID, selector),
			sqlgraph.To(statusitem.Table, statusitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, statusvalidchange.ToStatusItemTable, statusvalidchange.ToStatusItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(svcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StatusValidChange entity from the query.
// Returns a *NotFoundError when no StatusValidChange was found.
func (svcq *StatusValidChangeQuery) First(ctx context.Context) (*StatusValidChange, error) {
	nodes, err := svcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{statusvalidchange.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) FirstX(ctx context.Context) *StatusValidChange {
	node, err := svcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StatusValidChange ID from the query.
// Returns a *NotFoundError when no StatusValidChange ID was found.
func (svcq *StatusValidChangeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statusvalidchange.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) FirstIDX(ctx context.Context) int {
	id, err := svcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StatusValidChange entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one StatusValidChange entity is not found.
// Returns a *NotFoundError when no StatusValidChange entities are found.
func (svcq *StatusValidChangeQuery) Only(ctx context.Context) (*StatusValidChange, error) {
	nodes, err := svcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{statusvalidchange.Label}
	default:
		return nil, &NotSingularError{statusvalidchange.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) OnlyX(ctx context.Context) *StatusValidChange {
	node, err := svcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StatusValidChange ID in the query.
// Returns a *NotSingularError when exactly one StatusValidChange ID is not found.
// Returns a *NotFoundError when no entities are found.
func (svcq *StatusValidChangeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = &NotSingularError{statusvalidchange.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) OnlyIDX(ctx context.Context) int {
	id, err := svcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StatusValidChanges.
func (svcq *StatusValidChangeQuery) All(ctx context.Context) ([]*StatusValidChange, error) {
	if err := svcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return svcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) AllX(ctx context.Context) []*StatusValidChange {
	nodes, err := svcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StatusValidChange IDs.
func (svcq *StatusValidChangeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := svcq.Select(statusvalidchange.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) IDsX(ctx context.Context) []int {
	ids, err := svcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (svcq *StatusValidChangeQuery) Count(ctx context.Context) (int, error) {
	if err := svcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return svcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) CountX(ctx context.Context) int {
	count, err := svcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (svcq *StatusValidChangeQuery) Exist(ctx context.Context) (bool, error) {
	if err := svcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return svcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (svcq *StatusValidChangeQuery) ExistX(ctx context.Context) bool {
	exist, err := svcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StatusValidChangeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (svcq *StatusValidChangeQuery) Clone() *StatusValidChangeQuery {
	if svcq == nil {
		return nil
	}
	return &StatusValidChangeQuery{
		config:             svcq.config,
		limit:              svcq.limit,
		offset:             svcq.offset,
		order:              append([]OrderFunc{}, svcq.order...),
		predicates:         append([]predicate.StatusValidChange{}, svcq.predicates...),
		withMainStatusItem: svcq.withMainStatusItem.Clone(),
		withToStatusItem:   svcq.withToStatusItem.Clone(),
		// clone intermediate query.
		sql:  svcq.sql.Clone(),
		path: svcq.path,
	}
}

// WithMainStatusItem tells the query-builder to eager-load the nodes that are connected to
// the "main_status_item" edge. The optional arguments are used to configure the query builder of the edge.
func (svcq *StatusValidChangeQuery) WithMainStatusItem(opts ...func(*StatusItemQuery)) *StatusValidChangeQuery {
	query := &StatusItemQuery{config: svcq.config}
	for _, opt := range opts {
		opt(query)
	}
	svcq.withMainStatusItem = query
	return svcq
}

// WithToStatusItem tells the query-builder to eager-load the nodes that are connected to
// the "to_status_item" edge. The optional arguments are used to configure the query builder of the edge.
func (svcq *StatusValidChangeQuery) WithToStatusItem(opts ...func(*StatusItemQuery)) *StatusValidChangeQuery {
	query := &StatusItemQuery{config: svcq.config}
	for _, opt := range opts {
		opt(query)
	}
	svcq.withToStatusItem = query
	return svcq
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
//	client.StatusValidChange.Query().
//		GroupBy(statusvalidchange.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (svcq *StatusValidChangeQuery) GroupBy(field string, fields ...string) *StatusValidChangeGroupBy {
	group := &StatusValidChangeGroupBy{config: svcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := svcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return svcq.sqlQuery(ctx), nil
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
//	client.StatusValidChange.Query().
//		Select(statusvalidchange.FieldCreateTime).
//		Scan(ctx, &v)
//
func (svcq *StatusValidChangeQuery) Select(field string, fields ...string) *StatusValidChangeSelect {
	svcq.fields = append([]string{field}, fields...)
	return &StatusValidChangeSelect{StatusValidChangeQuery: svcq}
}

func (svcq *StatusValidChangeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range svcq.fields {
		if !statusvalidchange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if svcq.path != nil {
		prev, err := svcq.path(ctx)
		if err != nil {
			return err
		}
		svcq.sql = prev
	}
	return nil
}

func (svcq *StatusValidChangeQuery) sqlAll(ctx context.Context) ([]*StatusValidChange, error) {
	var (
		nodes       = []*StatusValidChange{}
		withFKs     = svcq.withFKs
		_spec       = svcq.querySpec()
		loadedTypes = [2]bool{
			svcq.withMainStatusItem != nil,
			svcq.withToStatusItem != nil,
		}
	)
	if svcq.withMainStatusItem != nil || svcq.withToStatusItem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, statusvalidchange.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StatusValidChange{config: svcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, svcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := svcq.withMainStatusItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*StatusValidChange)
		for i := range nodes {
			if nodes[i].status_item_main_status_valid_changes == nil {
				continue
			}
			fk := *nodes[i].status_item_main_status_valid_changes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(statusitem.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "status_item_main_status_valid_changes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.MainStatusItem = n
			}
		}
	}

	if query := svcq.withToStatusItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*StatusValidChange)
		for i := range nodes {
			if nodes[i].status_item_to_status_valid_changes == nil {
				continue
			}
			fk := *nodes[i].status_item_to_status_valid_changes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(statusitem.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "status_item_to_status_valid_changes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ToStatusItem = n
			}
		}
	}

	return nodes, nil
}

func (svcq *StatusValidChangeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := svcq.querySpec()
	return sqlgraph.CountNodes(ctx, svcq.driver, _spec)
}

func (svcq *StatusValidChangeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := svcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (svcq *StatusValidChangeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   statusvalidchange.Table,
			Columns: statusvalidchange.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: statusvalidchange.FieldID,
			},
		},
		From:   svcq.sql,
		Unique: true,
	}
	if unique := svcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := svcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statusvalidchange.FieldID)
		for i := range fields {
			if fields[i] != statusvalidchange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := svcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := svcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := svcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := svcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (svcq *StatusValidChangeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(svcq.driver.Dialect())
	t1 := builder.Table(statusvalidchange.Table)
	columns := svcq.fields
	if len(columns) == 0 {
		columns = statusvalidchange.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if svcq.sql != nil {
		selector = svcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range svcq.predicates {
		p(selector)
	}
	for _, p := range svcq.order {
		p(selector)
	}
	if offset := svcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := svcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StatusValidChangeGroupBy is the group-by builder for StatusValidChange entities.
type StatusValidChangeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (svcgb *StatusValidChangeGroupBy) Aggregate(fns ...AggregateFunc) *StatusValidChangeGroupBy {
	svcgb.fns = append(svcgb.fns, fns...)
	return svcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (svcgb *StatusValidChangeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := svcgb.path(ctx)
	if err != nil {
		return err
	}
	svcgb.sql = query
	return svcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := svcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(svcgb.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := svcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) StringsX(ctx context.Context) []string {
	v, err := svcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = svcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) StringX(ctx context.Context) string {
	v, err := svcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(svcgb.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := svcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) IntsX(ctx context.Context) []int {
	v, err := svcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = svcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) IntX(ctx context.Context) int {
	v, err := svcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(svcgb.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := svcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := svcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = svcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := svcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(svcgb.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := svcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := svcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (svcgb *StatusValidChangeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = svcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (svcgb *StatusValidChangeGroupBy) BoolX(ctx context.Context) bool {
	v, err := svcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (svcgb *StatusValidChangeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range svcgb.fields {
		if !statusvalidchange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := svcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := svcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (svcgb *StatusValidChangeGroupBy) sqlQuery() *sql.Selector {
	selector := svcgb.sql.Select()
	aggregation := make([]string, 0, len(svcgb.fns))
	for _, fn := range svcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(svcgb.fields)+len(svcgb.fns))
		for _, f := range svcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(svcgb.fields...)...)
}

// StatusValidChangeSelect is the builder for selecting fields of StatusValidChange entities.
type StatusValidChangeSelect struct {
	*StatusValidChangeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (svcs *StatusValidChangeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := svcs.prepareQuery(ctx); err != nil {
		return err
	}
	svcs.sql = svcs.StatusValidChangeQuery.sqlQuery(ctx)
	return svcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := svcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(svcs.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := svcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) StringsX(ctx context.Context) []string {
	v, err := svcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = svcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) StringX(ctx context.Context) string {
	v, err := svcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(svcs.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := svcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) IntsX(ctx context.Context) []int {
	v, err := svcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = svcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) IntX(ctx context.Context) int {
	v, err := svcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(svcs.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := svcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := svcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = svcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) Float64X(ctx context.Context) float64 {
	v, err := svcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(svcs.fields) > 1 {
		return nil, errors.New("ent: StatusValidChangeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := svcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) BoolsX(ctx context.Context) []bool {
	v, err := svcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (svcs *StatusValidChangeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = svcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{statusvalidchange.Label}
	default:
		err = fmt.Errorf("ent: StatusValidChangeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (svcs *StatusValidChangeSelect) BoolX(ctx context.Context) bool {
	v, err := svcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (svcs *StatusValidChangeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := svcs.sql.Query()
	if err := svcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
