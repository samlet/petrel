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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productconfigitem"
)

// ProductConfigItemQuery is the builder for querying ProductConfigItem entities.
type ProductConfigItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductConfigItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductConfigItemQuery builder.
func (pciq *ProductConfigItemQuery) Where(ps ...predicate.ProductConfigItem) *ProductConfigItemQuery {
	pciq.predicates = append(pciq.predicates, ps...)
	return pciq
}

// Limit adds a limit step to the query.
func (pciq *ProductConfigItemQuery) Limit(limit int) *ProductConfigItemQuery {
	pciq.limit = &limit
	return pciq
}

// Offset adds an offset step to the query.
func (pciq *ProductConfigItemQuery) Offset(offset int) *ProductConfigItemQuery {
	pciq.offset = &offset
	return pciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pciq *ProductConfigItemQuery) Unique(unique bool) *ProductConfigItemQuery {
	pciq.unique = &unique
	return pciq
}

// Order adds an order step to the query.
func (pciq *ProductConfigItemQuery) Order(o ...OrderFunc) *ProductConfigItemQuery {
	pciq.order = append(pciq.order, o...)
	return pciq
}

// First returns the first ProductConfigItem entity from the query.
// Returns a *NotFoundError when no ProductConfigItem was found.
func (pciq *ProductConfigItemQuery) First(ctx context.Context) (*ProductConfigItem, error) {
	nodes, err := pciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productconfigitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) FirstX(ctx context.Context) *ProductConfigItem {
	node, err := pciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductConfigItem ID from the query.
// Returns a *NotFoundError when no ProductConfigItem ID was found.
func (pciq *ProductConfigItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productconfigitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) FirstIDX(ctx context.Context) int {
	id, err := pciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductConfigItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductConfigItem entity is not found.
// Returns a *NotFoundError when no ProductConfigItem entities are found.
func (pciq *ProductConfigItemQuery) Only(ctx context.Context) (*ProductConfigItem, error) {
	nodes, err := pciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productconfigitem.Label}
	default:
		return nil, &NotSingularError{productconfigitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) OnlyX(ctx context.Context) *ProductConfigItem {
	node, err := pciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductConfigItem ID in the query.
// Returns a *NotSingularError when exactly one ProductConfigItem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pciq *ProductConfigItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = &NotSingularError{productconfigitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := pciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductConfigItems.
func (pciq *ProductConfigItemQuery) All(ctx context.Context) ([]*ProductConfigItem, error) {
	if err := pciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) AllX(ctx context.Context) []*ProductConfigItem {
	nodes, err := pciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductConfigItem IDs.
func (pciq *ProductConfigItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pciq.Select(productconfigitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) IDsX(ctx context.Context) []int {
	ids, err := pciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pciq *ProductConfigItemQuery) Count(ctx context.Context) (int, error) {
	if err := pciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) CountX(ctx context.Context) int {
	count, err := pciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pciq *ProductConfigItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := pciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pciq *ProductConfigItemQuery) ExistX(ctx context.Context) bool {
	exist, err := pciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductConfigItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pciq *ProductConfigItemQuery) Clone() *ProductConfigItemQuery {
	if pciq == nil {
		return nil
	}
	return &ProductConfigItemQuery{
		config:     pciq.config,
		limit:      pciq.limit,
		offset:     pciq.offset,
		order:      append([]OrderFunc{}, pciq.order...),
		predicates: append([]predicate.ProductConfigItem{}, pciq.predicates...),
		// clone intermediate query.
		sql:  pciq.sql.Clone(),
		path: pciq.path,
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
//	client.ProductConfigItem.Query().
//		GroupBy(productconfigitem.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pciq *ProductConfigItemQuery) GroupBy(field string, fields ...string) *ProductConfigItemGroupBy {
	group := &ProductConfigItemGroupBy{config: pciq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pciq.sqlQuery(ctx), nil
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
//	client.ProductConfigItem.Query().
//		Select(productconfigitem.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pciq *ProductConfigItemQuery) Select(field string, fields ...string) *ProductConfigItemSelect {
	pciq.fields = append([]string{field}, fields...)
	return &ProductConfigItemSelect{ProductConfigItemQuery: pciq}
}

func (pciq *ProductConfigItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pciq.fields {
		if !productconfigitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pciq.path != nil {
		prev, err := pciq.path(ctx)
		if err != nil {
			return err
		}
		pciq.sql = prev
	}
	return nil
}

func (pciq *ProductConfigItemQuery) sqlAll(ctx context.Context) ([]*ProductConfigItem, error) {
	var (
		nodes = []*ProductConfigItem{}
		_spec = pciq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductConfigItem{config: pciq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pciq *ProductConfigItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pciq.querySpec()
	return sqlgraph.CountNodes(ctx, pciq.driver, _spec)
}

func (pciq *ProductConfigItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pciq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pciq *ProductConfigItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productconfigitem.Table,
			Columns: productconfigitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productconfigitem.FieldID,
			},
		},
		From:   pciq.sql,
		Unique: true,
	}
	if unique := pciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productconfigitem.FieldID)
		for i := range fields {
			if fields[i] != productconfigitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pciq *ProductConfigItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pciq.driver.Dialect())
	t1 := builder.Table(productconfigitem.Table)
	columns := pciq.fields
	if len(columns) == 0 {
		columns = productconfigitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pciq.sql != nil {
		selector = pciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pciq.predicates {
		p(selector)
	}
	for _, p := range pciq.order {
		p(selector)
	}
	if offset := pciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductConfigItemGroupBy is the group-by builder for ProductConfigItem entities.
type ProductConfigItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcigb *ProductConfigItemGroupBy) Aggregate(fns ...AggregateFunc) *ProductConfigItemGroupBy {
	pcigb.fns = append(pcigb.fns, fns...)
	return pcigb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcigb *ProductConfigItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcigb.path(ctx)
	if err != nil {
		return err
	}
	pcigb.sql = query
	return pcigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pcigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pcigb.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pcigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := pcigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) StringX(ctx context.Context) string {
	v, err := pcigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pcigb.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pcigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := pcigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) IntX(ctx context.Context) int {
	v, err := pcigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcigb.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pcigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pcigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pcigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pcigb.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pcigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pcigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcigb *ProductConfigItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcigb *ProductConfigItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := pcigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcigb *ProductConfigItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcigb.fields {
		if !productconfigitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcigb *ProductConfigItemGroupBy) sqlQuery() *sql.Selector {
	selector := pcigb.sql.Select()
	aggregation := make([]string, 0, len(pcigb.fns))
	for _, fn := range pcigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcigb.fields)+len(pcigb.fns))
		for _, f := range pcigb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcigb.fields...)...)
}

// ProductConfigItemSelect is the builder for selecting fields of ProductConfigItem entities.
type ProductConfigItemSelect struct {
	*ProductConfigItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcis *ProductConfigItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcis.prepareQuery(ctx); err != nil {
		return err
	}
	pcis.sql = pcis.ProductConfigItemQuery.sqlQuery(ctx)
	return pcis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pcis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pcis.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pcis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) StringsX(ctx context.Context) []string {
	v, err := pcis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) StringX(ctx context.Context) string {
	v, err := pcis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pcis.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pcis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) IntsX(ctx context.Context) []int {
	v, err := pcis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) IntX(ctx context.Context) int {
	v, err := pcis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcis.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pcis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pcis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) Float64X(ctx context.Context) float64 {
	v, err := pcis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pcis.fields) > 1 {
		return nil, errors.New("ent: ProductConfigItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pcis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := pcis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pcis *ProductConfigItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productconfigitem.Label}
	default:
		err = fmt.Errorf("ent: ProductConfigItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcis *ProductConfigItemSelect) BoolX(ctx context.Context) bool {
	v, err := pcis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcis *ProductConfigItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcis.sql.Query()
	if err := pcis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
