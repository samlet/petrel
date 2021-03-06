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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/facilitygrouptype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// FacilityGroupTypeQuery is the builder for querying FacilityGroupType entities.
type FacilityGroupTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FacilityGroupType
	// eager-loading edges.
	withFacilityGroups *FacilityGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FacilityGroupTypeQuery builder.
func (fgtq *FacilityGroupTypeQuery) Where(ps ...predicate.FacilityGroupType) *FacilityGroupTypeQuery {
	fgtq.predicates = append(fgtq.predicates, ps...)
	return fgtq
}

// Limit adds a limit step to the query.
func (fgtq *FacilityGroupTypeQuery) Limit(limit int) *FacilityGroupTypeQuery {
	fgtq.limit = &limit
	return fgtq
}

// Offset adds an offset step to the query.
func (fgtq *FacilityGroupTypeQuery) Offset(offset int) *FacilityGroupTypeQuery {
	fgtq.offset = &offset
	return fgtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fgtq *FacilityGroupTypeQuery) Unique(unique bool) *FacilityGroupTypeQuery {
	fgtq.unique = &unique
	return fgtq
}

// Order adds an order step to the query.
func (fgtq *FacilityGroupTypeQuery) Order(o ...OrderFunc) *FacilityGroupTypeQuery {
	fgtq.order = append(fgtq.order, o...)
	return fgtq
}

// QueryFacilityGroups chains the current query on the "facility_groups" edge.
func (fgtq *FacilityGroupTypeQuery) QueryFacilityGroups() *FacilityGroupQuery {
	query := &FacilityGroupQuery{config: fgtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fgtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fgtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(facilitygrouptype.Table, facilitygrouptype.FieldID, selector),
			sqlgraph.To(facilitygroup.Table, facilitygroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, facilitygrouptype.FacilityGroupsTable, facilitygrouptype.FacilityGroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fgtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FacilityGroupType entity from the query.
// Returns a *NotFoundError when no FacilityGroupType was found.
func (fgtq *FacilityGroupTypeQuery) First(ctx context.Context) (*FacilityGroupType, error) {
	nodes, err := fgtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{facilitygrouptype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) FirstX(ctx context.Context) *FacilityGroupType {
	node, err := fgtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FacilityGroupType ID from the query.
// Returns a *NotFoundError when no FacilityGroupType ID was found.
func (fgtq *FacilityGroupTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{facilitygrouptype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := fgtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FacilityGroupType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one FacilityGroupType entity is not found.
// Returns a *NotFoundError when no FacilityGroupType entities are found.
func (fgtq *FacilityGroupTypeQuery) Only(ctx context.Context) (*FacilityGroupType, error) {
	nodes, err := fgtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{facilitygrouptype.Label}
	default:
		return nil, &NotSingularError{facilitygrouptype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) OnlyX(ctx context.Context) *FacilityGroupType {
	node, err := fgtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FacilityGroupType ID in the query.
// Returns a *NotSingularError when exactly one FacilityGroupType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fgtq *FacilityGroupTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fgtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = &NotSingularError{facilitygrouptype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := fgtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FacilityGroupTypes.
func (fgtq *FacilityGroupTypeQuery) All(ctx context.Context) ([]*FacilityGroupType, error) {
	if err := fgtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fgtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) AllX(ctx context.Context) []*FacilityGroupType {
	nodes, err := fgtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FacilityGroupType IDs.
func (fgtq *FacilityGroupTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fgtq.Select(facilitygrouptype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := fgtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fgtq *FacilityGroupTypeQuery) Count(ctx context.Context) (int, error) {
	if err := fgtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fgtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) CountX(ctx context.Context) int {
	count, err := fgtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fgtq *FacilityGroupTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := fgtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fgtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fgtq *FacilityGroupTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := fgtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FacilityGroupTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fgtq *FacilityGroupTypeQuery) Clone() *FacilityGroupTypeQuery {
	if fgtq == nil {
		return nil
	}
	return &FacilityGroupTypeQuery{
		config:             fgtq.config,
		limit:              fgtq.limit,
		offset:             fgtq.offset,
		order:              append([]OrderFunc{}, fgtq.order...),
		predicates:         append([]predicate.FacilityGroupType{}, fgtq.predicates...),
		withFacilityGroups: fgtq.withFacilityGroups.Clone(),
		// clone intermediate query.
		sql:  fgtq.sql.Clone(),
		path: fgtq.path,
	}
}

// WithFacilityGroups tells the query-builder to eager-load the nodes that are connected to
// the "facility_groups" edge. The optional arguments are used to configure the query builder of the edge.
func (fgtq *FacilityGroupTypeQuery) WithFacilityGroups(opts ...func(*FacilityGroupQuery)) *FacilityGroupTypeQuery {
	query := &FacilityGroupQuery{config: fgtq.config}
	for _, opt := range opts {
		opt(query)
	}
	fgtq.withFacilityGroups = query
	return fgtq
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
//	client.FacilityGroupType.Query().
//		GroupBy(facilitygrouptype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fgtq *FacilityGroupTypeQuery) GroupBy(field string, fields ...string) *FacilityGroupTypeGroupBy {
	group := &FacilityGroupTypeGroupBy{config: fgtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fgtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fgtq.sqlQuery(ctx), nil
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
//	client.FacilityGroupType.Query().
//		Select(facilitygrouptype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (fgtq *FacilityGroupTypeQuery) Select(field string, fields ...string) *FacilityGroupTypeSelect {
	fgtq.fields = append([]string{field}, fields...)
	return &FacilityGroupTypeSelect{FacilityGroupTypeQuery: fgtq}
}

func (fgtq *FacilityGroupTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fgtq.fields {
		if !facilitygrouptype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fgtq.path != nil {
		prev, err := fgtq.path(ctx)
		if err != nil {
			return err
		}
		fgtq.sql = prev
	}
	return nil
}

func (fgtq *FacilityGroupTypeQuery) sqlAll(ctx context.Context) ([]*FacilityGroupType, error) {
	var (
		nodes       = []*FacilityGroupType{}
		_spec       = fgtq.querySpec()
		loadedTypes = [1]bool{
			fgtq.withFacilityGroups != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FacilityGroupType{config: fgtq.config}
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
	if err := sqlgraph.QueryNodes(ctx, fgtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fgtq.withFacilityGroups; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*FacilityGroupType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.FacilityGroups = []*FacilityGroup{}
		}
		query.withFKs = true
		query.Where(predicate.FacilityGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(facilitygrouptype.FacilityGroupsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.facility_group_type_facility_groups
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "facility_group_type_facility_groups" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "facility_group_type_facility_groups" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.FacilityGroups = append(node.Edges.FacilityGroups, n)
		}
	}

	return nodes, nil
}

func (fgtq *FacilityGroupTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fgtq.querySpec()
	return sqlgraph.CountNodes(ctx, fgtq.driver, _spec)
}

func (fgtq *FacilityGroupTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fgtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fgtq *FacilityGroupTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   facilitygrouptype.Table,
			Columns: facilitygrouptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: facilitygrouptype.FieldID,
			},
		},
		From:   fgtq.sql,
		Unique: true,
	}
	if unique := fgtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fgtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, facilitygrouptype.FieldID)
		for i := range fields {
			if fields[i] != facilitygrouptype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fgtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fgtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fgtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fgtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fgtq *FacilityGroupTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fgtq.driver.Dialect())
	t1 := builder.Table(facilitygrouptype.Table)
	columns := fgtq.fields
	if len(columns) == 0 {
		columns = facilitygrouptype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fgtq.sql != nil {
		selector = fgtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range fgtq.predicates {
		p(selector)
	}
	for _, p := range fgtq.order {
		p(selector)
	}
	if offset := fgtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fgtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FacilityGroupTypeGroupBy is the group-by builder for FacilityGroupType entities.
type FacilityGroupTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgtgb *FacilityGroupTypeGroupBy) Aggregate(fns ...AggregateFunc) *FacilityGroupTypeGroupBy {
	fgtgb.fns = append(fgtgb.fns, fns...)
	return fgtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgtgb *FacilityGroupTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgtgb.path(ctx)
	if err != nil {
		return err
	}
	fgtgb.sql = query
	return fgtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgtgb.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) StringX(ctx context.Context) string {
	v, err := fgtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgtgb.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) IntX(ctx context.Context) int {
	v, err := fgtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgtgb.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgtgb.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgtgb *FacilityGroupTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgtgb *FacilityGroupTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgtgb *FacilityGroupTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgtgb.fields {
		if !facilitygrouptype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgtgb *FacilityGroupTypeGroupBy) sqlQuery() *sql.Selector {
	selector := fgtgb.sql.Select()
	aggregation := make([]string, 0, len(fgtgb.fns))
	for _, fn := range fgtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgtgb.fields)+len(fgtgb.fns))
		for _, f := range fgtgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgtgb.fields...)...)
}

// FacilityGroupTypeSelect is the builder for selecting fields of FacilityGroupType entities.
type FacilityGroupTypeSelect struct {
	*FacilityGroupTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fgts *FacilityGroupTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fgts.prepareQuery(ctx); err != nil {
		return err
	}
	fgts.sql = fgts.FacilityGroupTypeQuery.sqlQuery(ctx)
	return fgts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fgts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fgts.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fgts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fgts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) StringX(ctx context.Context) string {
	v, err := fgts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fgts.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fgts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fgts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) IntX(ctx context.Context) int {
	v, err := fgts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgts.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fgts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fgts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := fgts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fgts.fields) > 1 {
		return nil, errors.New("ent: FacilityGroupTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fgts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fgts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fgts *FacilityGroupTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{facilitygrouptype.Label}
	default:
		err = fmt.Errorf("ent: FacilityGroupTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgts *FacilityGroupTypeSelect) BoolX(ctx context.Context) bool {
	v, err := fgts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgts *FacilityGroupTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fgts.sql.Query()
	if err := fgts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
