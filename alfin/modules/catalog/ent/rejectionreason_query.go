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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/rejectionreason"
)

// RejectionReasonQuery is the builder for querying RejectionReason entities.
type RejectionReasonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RejectionReason
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RejectionReasonQuery builder.
func (rrq *RejectionReasonQuery) Where(ps ...predicate.RejectionReason) *RejectionReasonQuery {
	rrq.predicates = append(rrq.predicates, ps...)
	return rrq
}

// Limit adds a limit step to the query.
func (rrq *RejectionReasonQuery) Limit(limit int) *RejectionReasonQuery {
	rrq.limit = &limit
	return rrq
}

// Offset adds an offset step to the query.
func (rrq *RejectionReasonQuery) Offset(offset int) *RejectionReasonQuery {
	rrq.offset = &offset
	return rrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rrq *RejectionReasonQuery) Unique(unique bool) *RejectionReasonQuery {
	rrq.unique = &unique
	return rrq
}

// Order adds an order step to the query.
func (rrq *RejectionReasonQuery) Order(o ...OrderFunc) *RejectionReasonQuery {
	rrq.order = append(rrq.order, o...)
	return rrq
}

// First returns the first RejectionReason entity from the query.
// Returns a *NotFoundError when no RejectionReason was found.
func (rrq *RejectionReasonQuery) First(ctx context.Context) (*RejectionReason, error) {
	nodes, err := rrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rejectionreason.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rrq *RejectionReasonQuery) FirstX(ctx context.Context) *RejectionReason {
	node, err := rrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RejectionReason ID from the query.
// Returns a *NotFoundError when no RejectionReason ID was found.
func (rrq *RejectionReasonQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rejectionreason.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rrq *RejectionReasonQuery) FirstIDX(ctx context.Context) int {
	id, err := rrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RejectionReason entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one RejectionReason entity is not found.
// Returns a *NotFoundError when no RejectionReason entities are found.
func (rrq *RejectionReasonQuery) Only(ctx context.Context) (*RejectionReason, error) {
	nodes, err := rrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rejectionreason.Label}
	default:
		return nil, &NotSingularError{rejectionreason.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rrq *RejectionReasonQuery) OnlyX(ctx context.Context) *RejectionReason {
	node, err := rrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RejectionReason ID in the query.
// Returns a *NotSingularError when exactly one RejectionReason ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rrq *RejectionReasonQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = &NotSingularError{rejectionreason.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rrq *RejectionReasonQuery) OnlyIDX(ctx context.Context) int {
	id, err := rrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RejectionReasons.
func (rrq *RejectionReasonQuery) All(ctx context.Context) ([]*RejectionReason, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rrq *RejectionReasonQuery) AllX(ctx context.Context) []*RejectionReason {
	nodes, err := rrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RejectionReason IDs.
func (rrq *RejectionReasonQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rrq.Select(rejectionreason.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rrq *RejectionReasonQuery) IDsX(ctx context.Context) []int {
	ids, err := rrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rrq *RejectionReasonQuery) Count(ctx context.Context) (int, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rrq *RejectionReasonQuery) CountX(ctx context.Context) int {
	count, err := rrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rrq *RejectionReasonQuery) Exist(ctx context.Context) (bool, error) {
	if err := rrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rrq *RejectionReasonQuery) ExistX(ctx context.Context) bool {
	exist, err := rrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RejectionReasonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rrq *RejectionReasonQuery) Clone() *RejectionReasonQuery {
	if rrq == nil {
		return nil
	}
	return &RejectionReasonQuery{
		config:     rrq.config,
		limit:      rrq.limit,
		offset:     rrq.offset,
		order:      append([]OrderFunc{}, rrq.order...),
		predicates: append([]predicate.RejectionReason{}, rrq.predicates...),
		// clone intermediate query.
		sql:  rrq.sql.Clone(),
		path: rrq.path,
	}
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
//	client.RejectionReason.Query().
//		GroupBy(rejectionreason.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rrq *RejectionReasonQuery) GroupBy(field string, fields ...string) *RejectionReasonGroupBy {
	group := &RejectionReasonGroupBy{config: rrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rrq.sqlQuery(ctx), nil
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
//	client.RejectionReason.Query().
//		Select(rejectionreason.FieldCreateTime).
//		Scan(ctx, &v)
//
func (rrq *RejectionReasonQuery) Select(field string, fields ...string) *RejectionReasonSelect {
	rrq.fields = append([]string{field}, fields...)
	return &RejectionReasonSelect{RejectionReasonQuery: rrq}
}

func (rrq *RejectionReasonQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rrq.fields {
		if !rejectionreason.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rrq.path != nil {
		prev, err := rrq.path(ctx)
		if err != nil {
			return err
		}
		rrq.sql = prev
	}
	return nil
}

func (rrq *RejectionReasonQuery) sqlAll(ctx context.Context) ([]*RejectionReason, error) {
	var (
		nodes = []*RejectionReason{}
		_spec = rrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &RejectionReason{config: rrq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rrq *RejectionReasonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rrq.querySpec()
	return sqlgraph.CountNodes(ctx, rrq.driver, _spec)
}

func (rrq *RejectionReasonQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rrq *RejectionReasonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rejectionreason.Table,
			Columns: rejectionreason.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rejectionreason.FieldID,
			},
		},
		From:   rrq.sql,
		Unique: true,
	}
	if unique := rrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rejectionreason.FieldID)
		for i := range fields {
			if fields[i] != rejectionreason.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rrq *RejectionReasonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rrq.driver.Dialect())
	t1 := builder.Table(rejectionreason.Table)
	columns := rrq.fields
	if len(columns) == 0 {
		columns = rejectionreason.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rrq.sql != nil {
		selector = rrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range rrq.predicates {
		p(selector)
	}
	for _, p := range rrq.order {
		p(selector)
	}
	if offset := rrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RejectionReasonGroupBy is the group-by builder for RejectionReason entities.
type RejectionReasonGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rrgb *RejectionReasonGroupBy) Aggregate(fns ...AggregateFunc) *RejectionReasonGroupBy {
	rrgb.fns = append(rrgb.fns, fns...)
	return rrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rrgb *RejectionReasonGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rrgb.path(ctx)
	if err != nil {
		return err
	}
	rrgb.sql = query
	return rrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) StringsX(ctx context.Context) []string {
	v, err := rrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) StringX(ctx context.Context) string {
	v, err := rrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) IntsX(ctx context.Context) []int {
	v, err := rrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) IntX(ctx context.Context) int {
	v, err := rrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rrgb.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rrgb *RejectionReasonGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rrgb *RejectionReasonGroupBy) BoolX(ctx context.Context) bool {
	v, err := rrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rrgb *RejectionReasonGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rrgb.fields {
		if !rejectionreason.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rrgb *RejectionReasonGroupBy) sqlQuery() *sql.Selector {
	selector := rrgb.sql.Select()
	aggregation := make([]string, 0, len(rrgb.fns))
	for _, fn := range rrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rrgb.fields)+len(rrgb.fns))
		for _, f := range rrgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rrgb.fields...)...)
}

// RejectionReasonSelect is the builder for selecting fields of RejectionReason entities.
type RejectionReasonSelect struct {
	*RejectionReasonQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rrs *RejectionReasonSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rrs.prepareQuery(ctx); err != nil {
		return err
	}
	rrs.sql = rrs.RejectionReasonQuery.sqlQuery(ctx)
	return rrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rrs *RejectionReasonSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rrs *RejectionReasonSelect) StringsX(ctx context.Context) []string {
	v, err := rrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rrs *RejectionReasonSelect) StringX(ctx context.Context) string {
	v, err := rrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rrs *RejectionReasonSelect) IntsX(ctx context.Context) []int {
	v, err := rrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rrs *RejectionReasonSelect) IntX(ctx context.Context) int {
	v, err := rrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rrs *RejectionReasonSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rrs *RejectionReasonSelect) Float64X(ctx context.Context) float64 {
	v, err := rrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rrs.fields) > 1 {
		return nil, errors.New("ent: RejectionReasonSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rrs *RejectionReasonSelect) BoolsX(ctx context.Context) []bool {
	v, err := rrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rrs *RejectionReasonSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{rejectionreason.Label}
	default:
		err = fmt.Errorf("ent: RejectionReasonSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rrs *RejectionReasonSelect) BoolX(ctx context.Context) bool {
	v, err := rrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rrs *RejectionReasonSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rrs.sql.Query()
	if err := rrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
