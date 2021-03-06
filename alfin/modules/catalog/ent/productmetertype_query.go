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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productmetertype"
)

// ProductMeterTypeQuery is the builder for querying ProductMeterType entities.
type ProductMeterTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductMeterType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductMeterTypeQuery builder.
func (pmtq *ProductMeterTypeQuery) Where(ps ...predicate.ProductMeterType) *ProductMeterTypeQuery {
	pmtq.predicates = append(pmtq.predicates, ps...)
	return pmtq
}

// Limit adds a limit step to the query.
func (pmtq *ProductMeterTypeQuery) Limit(limit int) *ProductMeterTypeQuery {
	pmtq.limit = &limit
	return pmtq
}

// Offset adds an offset step to the query.
func (pmtq *ProductMeterTypeQuery) Offset(offset int) *ProductMeterTypeQuery {
	pmtq.offset = &offset
	return pmtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pmtq *ProductMeterTypeQuery) Unique(unique bool) *ProductMeterTypeQuery {
	pmtq.unique = &unique
	return pmtq
}

// Order adds an order step to the query.
func (pmtq *ProductMeterTypeQuery) Order(o ...OrderFunc) *ProductMeterTypeQuery {
	pmtq.order = append(pmtq.order, o...)
	return pmtq
}

// First returns the first ProductMeterType entity from the query.
// Returns a *NotFoundError when no ProductMeterType was found.
func (pmtq *ProductMeterTypeQuery) First(ctx context.Context) (*ProductMeterType, error) {
	nodes, err := pmtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productmetertype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) FirstX(ctx context.Context) *ProductMeterType {
	node, err := pmtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductMeterType ID from the query.
// Returns a *NotFoundError when no ProductMeterType ID was found.
func (pmtq *ProductMeterTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productmetertype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := pmtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductMeterType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductMeterType entity is not found.
// Returns a *NotFoundError when no ProductMeterType entities are found.
func (pmtq *ProductMeterTypeQuery) Only(ctx context.Context) (*ProductMeterType, error) {
	nodes, err := pmtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productmetertype.Label}
	default:
		return nil, &NotSingularError{productmetertype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) OnlyX(ctx context.Context) *ProductMeterType {
	node, err := pmtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductMeterType ID in the query.
// Returns a *NotSingularError when exactly one ProductMeterType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pmtq *ProductMeterTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pmtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = &NotSingularError{productmetertype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pmtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductMeterTypes.
func (pmtq *ProductMeterTypeQuery) All(ctx context.Context) ([]*ProductMeterType, error) {
	if err := pmtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pmtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) AllX(ctx context.Context) []*ProductMeterType {
	nodes, err := pmtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductMeterType IDs.
func (pmtq *ProductMeterTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pmtq.Select(productmetertype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := pmtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pmtq *ProductMeterTypeQuery) Count(ctx context.Context) (int, error) {
	if err := pmtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pmtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) CountX(ctx context.Context) int {
	count, err := pmtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pmtq *ProductMeterTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pmtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pmtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pmtq *ProductMeterTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := pmtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductMeterTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pmtq *ProductMeterTypeQuery) Clone() *ProductMeterTypeQuery {
	if pmtq == nil {
		return nil
	}
	return &ProductMeterTypeQuery{
		config:     pmtq.config,
		limit:      pmtq.limit,
		offset:     pmtq.offset,
		order:      append([]OrderFunc{}, pmtq.order...),
		predicates: append([]predicate.ProductMeterType{}, pmtq.predicates...),
		// clone intermediate query.
		sql:  pmtq.sql.Clone(),
		path: pmtq.path,
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
//	client.ProductMeterType.Query().
//		GroupBy(productmetertype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pmtq *ProductMeterTypeQuery) GroupBy(field string, fields ...string) *ProductMeterTypeGroupBy {
	group := &ProductMeterTypeGroupBy{config: pmtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pmtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pmtq.sqlQuery(ctx), nil
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
//	client.ProductMeterType.Query().
//		Select(productmetertype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pmtq *ProductMeterTypeQuery) Select(field string, fields ...string) *ProductMeterTypeSelect {
	pmtq.fields = append([]string{field}, fields...)
	return &ProductMeterTypeSelect{ProductMeterTypeQuery: pmtq}
}

func (pmtq *ProductMeterTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pmtq.fields {
		if !productmetertype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pmtq.path != nil {
		prev, err := pmtq.path(ctx)
		if err != nil {
			return err
		}
		pmtq.sql = prev
	}
	return nil
}

func (pmtq *ProductMeterTypeQuery) sqlAll(ctx context.Context) ([]*ProductMeterType, error) {
	var (
		nodes = []*ProductMeterType{}
		_spec = pmtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductMeterType{config: pmtq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pmtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pmtq *ProductMeterTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pmtq.querySpec()
	return sqlgraph.CountNodes(ctx, pmtq.driver, _spec)
}

func (pmtq *ProductMeterTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pmtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pmtq *ProductMeterTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productmetertype.Table,
			Columns: productmetertype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productmetertype.FieldID,
			},
		},
		From:   pmtq.sql,
		Unique: true,
	}
	if unique := pmtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pmtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productmetertype.FieldID)
		for i := range fields {
			if fields[i] != productmetertype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pmtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pmtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pmtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pmtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pmtq *ProductMeterTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pmtq.driver.Dialect())
	t1 := builder.Table(productmetertype.Table)
	columns := pmtq.fields
	if len(columns) == 0 {
		columns = productmetertype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pmtq.sql != nil {
		selector = pmtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pmtq.predicates {
		p(selector)
	}
	for _, p := range pmtq.order {
		p(selector)
	}
	if offset := pmtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pmtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductMeterTypeGroupBy is the group-by builder for ProductMeterType entities.
type ProductMeterTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pmtgb *ProductMeterTypeGroupBy) Aggregate(fns ...AggregateFunc) *ProductMeterTypeGroupBy {
	pmtgb.fns = append(pmtgb.fns, fns...)
	return pmtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pmtgb *ProductMeterTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pmtgb.path(ctx)
	if err != nil {
		return err
	}
	pmtgb.sql = query
	return pmtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pmtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pmtgb.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pmtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pmtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pmtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) StringX(ctx context.Context) string {
	v, err := pmtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pmtgb.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pmtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pmtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pmtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) IntX(ctx context.Context) int {
	v, err := pmtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pmtgb.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pmtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pmtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pmtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pmtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pmtgb.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pmtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pmtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pmtgb *ProductMeterTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pmtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pmtgb *ProductMeterTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pmtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pmtgb *ProductMeterTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pmtgb.fields {
		if !productmetertype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pmtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pmtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pmtgb *ProductMeterTypeGroupBy) sqlQuery() *sql.Selector {
	selector := pmtgb.sql.Select()
	aggregation := make([]string, 0, len(pmtgb.fns))
	for _, fn := range pmtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pmtgb.fields)+len(pmtgb.fns))
		for _, f := range pmtgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pmtgb.fields...)...)
}

// ProductMeterTypeSelect is the builder for selecting fields of ProductMeterType entities.
type ProductMeterTypeSelect struct {
	*ProductMeterTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pmts *ProductMeterTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pmts.prepareQuery(ctx); err != nil {
		return err
	}
	pmts.sql = pmts.ProductMeterTypeQuery.sqlQuery(ctx)
	return pmts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pmts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pmts.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pmts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) StringsX(ctx context.Context) []string {
	v, err := pmts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pmts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) StringX(ctx context.Context) string {
	v, err := pmts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pmts.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pmts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) IntsX(ctx context.Context) []int {
	v, err := pmts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pmts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) IntX(ctx context.Context) int {
	v, err := pmts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pmts.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pmts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pmts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pmts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := pmts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pmts.fields) > 1 {
		return nil, errors.New("ent: ProductMeterTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pmts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pmts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pmts *ProductMeterTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pmts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productmetertype.Label}
	default:
		err = fmt.Errorf("ent: ProductMeterTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pmts *ProductMeterTypeSelect) BoolX(ctx context.Context) bool {
	v, err := pmts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pmts *ProductMeterTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pmts.sql.Query()
	if err := pmts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
