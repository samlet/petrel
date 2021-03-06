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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilityassoctype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// FacilityAssocTypeQuery is the builder for querying FacilityAssocType entities.
type FacilityAssocTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FacilityAssocType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FacilityAssocTypeQuery builder.
func (fatq *FacilityAssocTypeQuery) Where(ps ...predicate.FacilityAssocType) *FacilityAssocTypeQuery {
	fatq.predicates = append(fatq.predicates, ps...)
	return fatq
}

// Limit adds a limit step to the query.
func (fatq *FacilityAssocTypeQuery) Limit(limit int) *FacilityAssocTypeQuery {
	fatq.limit = &limit
	return fatq
}

// Offset adds an offset step to the query.
func (fatq *FacilityAssocTypeQuery) Offset(offset int) *FacilityAssocTypeQuery {
	fatq.offset = &offset
	return fatq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fatq *FacilityAssocTypeQuery) Unique(unique bool) *FacilityAssocTypeQuery {
	fatq.unique = &unique
	return fatq
}

// Order adds an order step to the query.
func (fatq *FacilityAssocTypeQuery) Order(o ...OrderFunc) *FacilityAssocTypeQuery {
	fatq.order = append(fatq.order, o...)
	return fatq
}

// First returns the first FacilityAssocType entity from the query.
// Returns a *NotFoundError when no FacilityAssocType was found.
func (fatq *FacilityAssocTypeQuery) First(ctx context.Context) (*FacilityAssocType, error) {
	nodes, err := fatq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{facilityassoctype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) FirstX(ctx context.Context) *FacilityAssocType {
	node, err := fatq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FacilityAssocType ID from the query.
// Returns a *NotFoundError when no FacilityAssocType ID was found.
func (fatq *FacilityAssocTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fatq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{facilityassoctype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := fatq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FacilityAssocType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one FacilityAssocType entity is not found.
// Returns a *NotFoundError when no FacilityAssocType entities are found.
func (fatq *FacilityAssocTypeQuery) Only(ctx context.Context) (*FacilityAssocType, error) {
	nodes, err := fatq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{facilityassoctype.Label}
	default:
		return nil, &NotSingularError{facilityassoctype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) OnlyX(ctx context.Context) *FacilityAssocType {
	node, err := fatq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FacilityAssocType ID in the query.
// Returns a *NotSingularError when exactly one FacilityAssocType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fatq *FacilityAssocTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fatq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = &NotSingularError{facilityassoctype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := fatq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FacilityAssocTypes.
func (fatq *FacilityAssocTypeQuery) All(ctx context.Context) ([]*FacilityAssocType, error) {
	if err := fatq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fatq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) AllX(ctx context.Context) []*FacilityAssocType {
	nodes, err := fatq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FacilityAssocType IDs.
func (fatq *FacilityAssocTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fatq.Select(facilityassoctype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := fatq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fatq *FacilityAssocTypeQuery) Count(ctx context.Context) (int, error) {
	if err := fatq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fatq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) CountX(ctx context.Context) int {
	count, err := fatq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fatq *FacilityAssocTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := fatq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fatq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fatq *FacilityAssocTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := fatq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FacilityAssocTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fatq *FacilityAssocTypeQuery) Clone() *FacilityAssocTypeQuery {
	if fatq == nil {
		return nil
	}
	return &FacilityAssocTypeQuery{
		config:     fatq.config,
		limit:      fatq.limit,
		offset:     fatq.offset,
		order:      append([]OrderFunc{}, fatq.order...),
		predicates: append([]predicate.FacilityAssocType{}, fatq.predicates...),
		// clone intermediate query.
		sql:  fatq.sql.Clone(),
		path: fatq.path,
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
//	client.FacilityAssocType.Query().
//		GroupBy(facilityassoctype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fatq *FacilityAssocTypeQuery) GroupBy(field string, fields ...string) *FacilityAssocTypeGroupBy {
	group := &FacilityAssocTypeGroupBy{config: fatq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fatq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fatq.sqlQuery(ctx), nil
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
//	client.FacilityAssocType.Query().
//		Select(facilityassoctype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (fatq *FacilityAssocTypeQuery) Select(field string, fields ...string) *FacilityAssocTypeSelect {
	fatq.fields = append([]string{field}, fields...)
	return &FacilityAssocTypeSelect{FacilityAssocTypeQuery: fatq}
}

func (fatq *FacilityAssocTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fatq.fields {
		if !facilityassoctype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fatq.path != nil {
		prev, err := fatq.path(ctx)
		if err != nil {
			return err
		}
		fatq.sql = prev
	}
	return nil
}

func (fatq *FacilityAssocTypeQuery) sqlAll(ctx context.Context) ([]*FacilityAssocType, error) {
	var (
		nodes = []*FacilityAssocType{}
		_spec = fatq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FacilityAssocType{config: fatq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fatq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fatq *FacilityAssocTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fatq.querySpec()
	return sqlgraph.CountNodes(ctx, fatq.driver, _spec)
}

func (fatq *FacilityAssocTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fatq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fatq *FacilityAssocTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilityassoctype.Table,
			Columns: facilityassoctype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilityassoctype.FieldID,
			},
		},
		From:   fatq.sql,
		Unique: true,
	}
	if unique := fatq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fatq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, facilityassoctype.FieldID)
		for i := range fields {
			if fields[i] != facilityassoctype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fatq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fatq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fatq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fatq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fatq *FacilityAssocTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fatq.driver.Dialect())
	t1 := builder.Table(facilityassoctype.Table)
	columns := fatq.fields
	if len(columns) == 0 {
		columns = facilityassoctype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fatq.sql != nil {
		selector = fatq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range fatq.predicates {
		p(selector)
	}
	for _, p := range fatq.order {
		p(selector)
	}
	if offset := fatq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fatq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FacilityAssocTypeGroupBy is the group-by builder for FacilityAssocType entities.
type FacilityAssocTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fatgb *FacilityAssocTypeGroupBy) Aggregate(fns ...AggregateFunc) *FacilityAssocTypeGroupBy {
	fatgb.fns = append(fatgb.fns, fns...)
	return fatgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fatgb *FacilityAssocTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fatgb.path(ctx)
	if err != nil {
		return err
	}
	fatgb.sql = query
	return fatgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fatgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fatgb.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := fatgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fatgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) StringX(ctx context.Context) string {
	v, err := fatgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fatgb.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := fatgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fatgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) IntX(ctx context.Context) int {
	v, err := fatgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fatgb.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fatgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fatgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fatgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fatgb.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fatgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fatgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fatgb *FacilityAssocTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fatgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fatgb *FacilityAssocTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := fatgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fatgb *FacilityAssocTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fatgb.fields {
		if !facilityassoctype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fatgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fatgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fatgb *FacilityAssocTypeGroupBy) sqlQuery() *sql.Selector {
	selector := fatgb.sql.Select()
	aggregation := make([]string, 0, len(fatgb.fns))
	for _, fn := range fatgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fatgb.fields)+len(fatgb.fns))
		for _, f := range fatgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fatgb.fields...)...)
}

// FacilityAssocTypeSelect is the builder for selecting fields of FacilityAssocType entities.
type FacilityAssocTypeSelect struct {
	*FacilityAssocTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fats *FacilityAssocTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fats.prepareQuery(ctx); err != nil {
		return err
	}
	fats.sql = fats.FacilityAssocTypeQuery.sqlQuery(ctx)
	return fats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fats.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) StringX(ctx context.Context) string {
	v, err := fats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fats.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) IntX(ctx context.Context) int {
	v, err := fats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fats.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := fats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fats.fields) > 1 {
		return nil, errors.New("ent: FacilityAssocTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fats *FacilityAssocTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilityassoctype.Label}
	default:
		err = fmt.Errorf("ent: FacilityAssocTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fats *FacilityAssocTypeSelect) BoolX(ctx context.Context) bool {
	v, err := fats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fats *FacilityAssocTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fats.sql.Query()
	if err := fats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
