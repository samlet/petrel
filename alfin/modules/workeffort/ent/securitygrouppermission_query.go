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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygrouppermission"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitypermission"
)

// SecurityGroupPermissionQuery is the builder for querying SecurityGroupPermission entities.
type SecurityGroupPermissionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SecurityGroupPermission
	// eager-loading edges.
	withSecurityGroup      *SecurityGroupQuery
	withSecurityPermission *SecurityPermissionQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SecurityGroupPermissionQuery builder.
func (sgpq *SecurityGroupPermissionQuery) Where(ps ...predicate.SecurityGroupPermission) *SecurityGroupPermissionQuery {
	sgpq.predicates = append(sgpq.predicates, ps...)
	return sgpq
}

// Limit adds a limit step to the query.
func (sgpq *SecurityGroupPermissionQuery) Limit(limit int) *SecurityGroupPermissionQuery {
	sgpq.limit = &limit
	return sgpq
}

// Offset adds an offset step to the query.
func (sgpq *SecurityGroupPermissionQuery) Offset(offset int) *SecurityGroupPermissionQuery {
	sgpq.offset = &offset
	return sgpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sgpq *SecurityGroupPermissionQuery) Unique(unique bool) *SecurityGroupPermissionQuery {
	sgpq.unique = &unique
	return sgpq
}

// Order adds an order step to the query.
func (sgpq *SecurityGroupPermissionQuery) Order(o ...OrderFunc) *SecurityGroupPermissionQuery {
	sgpq.order = append(sgpq.order, o...)
	return sgpq
}

// QuerySecurityGroup chains the current query on the "security_group" edge.
func (sgpq *SecurityGroupPermissionQuery) QuerySecurityGroup() *SecurityGroupQuery {
	query := &SecurityGroupQuery{config: sgpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(securitygrouppermission.Table, securitygrouppermission.FieldID, selector),
			sqlgraph.To(securitygroup.Table, securitygroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, securitygrouppermission.SecurityGroupTable, securitygrouppermission.SecurityGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySecurityPermission chains the current query on the "security_permission" edge.
func (sgpq *SecurityGroupPermissionQuery) QuerySecurityPermission() *SecurityPermissionQuery {
	query := &SecurityPermissionQuery{config: sgpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(securitygrouppermission.Table, securitygrouppermission.FieldID, selector),
			sqlgraph.To(securitypermission.Table, securitypermission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, securitygrouppermission.SecurityPermissionTable, securitygrouppermission.SecurityPermissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SecurityGroupPermission entity from the query.
// Returns a *NotFoundError when no SecurityGroupPermission was found.
func (sgpq *SecurityGroupPermissionQuery) First(ctx context.Context) (*SecurityGroupPermission, error) {
	nodes, err := sgpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{securitygrouppermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) FirstX(ctx context.Context) *SecurityGroupPermission {
	node, err := sgpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SecurityGroupPermission ID from the query.
// Returns a *NotFoundError when no SecurityGroupPermission ID was found.
func (sgpq *SecurityGroupPermissionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{securitygrouppermission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) FirstIDX(ctx context.Context) int {
	id, err := sgpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SecurityGroupPermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SecurityGroupPermission entity is not found.
// Returns a *NotFoundError when no SecurityGroupPermission entities are found.
func (sgpq *SecurityGroupPermissionQuery) Only(ctx context.Context) (*SecurityGroupPermission, error) {
	nodes, err := sgpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{securitygrouppermission.Label}
	default:
		return nil, &NotSingularError{securitygrouppermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) OnlyX(ctx context.Context) *SecurityGroupPermission {
	node, err := sgpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SecurityGroupPermission ID in the query.
// Returns a *NotSingularError when exactly one SecurityGroupPermission ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sgpq *SecurityGroupPermissionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = &NotSingularError{securitygrouppermission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) OnlyIDX(ctx context.Context) int {
	id, err := sgpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SecurityGroupPermissions.
func (sgpq *SecurityGroupPermissionQuery) All(ctx context.Context) ([]*SecurityGroupPermission, error) {
	if err := sgpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sgpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) AllX(ctx context.Context) []*SecurityGroupPermission {
	nodes, err := sgpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SecurityGroupPermission IDs.
func (sgpq *SecurityGroupPermissionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sgpq.Select(securitygrouppermission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) IDsX(ctx context.Context) []int {
	ids, err := sgpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sgpq *SecurityGroupPermissionQuery) Count(ctx context.Context) (int, error) {
	if err := sgpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sgpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) CountX(ctx context.Context) int {
	count, err := sgpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sgpq *SecurityGroupPermissionQuery) Exist(ctx context.Context) (bool, error) {
	if err := sgpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sgpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sgpq *SecurityGroupPermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := sgpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SecurityGroupPermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sgpq *SecurityGroupPermissionQuery) Clone() *SecurityGroupPermissionQuery {
	if sgpq == nil {
		return nil
	}
	return &SecurityGroupPermissionQuery{
		config:                 sgpq.config,
		limit:                  sgpq.limit,
		offset:                 sgpq.offset,
		order:                  append([]OrderFunc{}, sgpq.order...),
		predicates:             append([]predicate.SecurityGroupPermission{}, sgpq.predicates...),
		withSecurityGroup:      sgpq.withSecurityGroup.Clone(),
		withSecurityPermission: sgpq.withSecurityPermission.Clone(),
		// clone intermediate query.
		sql:  sgpq.sql.Clone(),
		path: sgpq.path,
	}
}

// WithSecurityGroup tells the query-builder to eager-load the nodes that are connected to
// the "security_group" edge. The optional arguments are used to configure the query builder of the edge.
func (sgpq *SecurityGroupPermissionQuery) WithSecurityGroup(opts ...func(*SecurityGroupQuery)) *SecurityGroupPermissionQuery {
	query := &SecurityGroupQuery{config: sgpq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgpq.withSecurityGroup = query
	return sgpq
}

// WithSecurityPermission tells the query-builder to eager-load the nodes that are connected to
// the "security_permission" edge. The optional arguments are used to configure the query builder of the edge.
func (sgpq *SecurityGroupPermissionQuery) WithSecurityPermission(opts ...func(*SecurityPermissionQuery)) *SecurityGroupPermissionQuery {
	query := &SecurityPermissionQuery{config: sgpq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgpq.withSecurityPermission = query
	return sgpq
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
//	client.SecurityGroupPermission.Query().
//		GroupBy(securitygrouppermission.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sgpq *SecurityGroupPermissionQuery) GroupBy(field string, fields ...string) *SecurityGroupPermissionGroupBy {
	group := &SecurityGroupPermissionGroupBy{config: sgpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sgpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sgpq.sqlQuery(ctx), nil
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
//	client.SecurityGroupPermission.Query().
//		Select(securitygrouppermission.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sgpq *SecurityGroupPermissionQuery) Select(field string, fields ...string) *SecurityGroupPermissionSelect {
	sgpq.fields = append([]string{field}, fields...)
	return &SecurityGroupPermissionSelect{SecurityGroupPermissionQuery: sgpq}
}

func (sgpq *SecurityGroupPermissionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sgpq.fields {
		if !securitygrouppermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sgpq.path != nil {
		prev, err := sgpq.path(ctx)
		if err != nil {
			return err
		}
		sgpq.sql = prev
	}
	return nil
}

func (sgpq *SecurityGroupPermissionQuery) sqlAll(ctx context.Context) ([]*SecurityGroupPermission, error) {
	var (
		nodes       = []*SecurityGroupPermission{}
		withFKs     = sgpq.withFKs
		_spec       = sgpq.querySpec()
		loadedTypes = [2]bool{
			sgpq.withSecurityGroup != nil,
			sgpq.withSecurityPermission != nil,
		}
	)
	if sgpq.withSecurityGroup != nil || sgpq.withSecurityPermission != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, securitygrouppermission.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SecurityGroupPermission{config: sgpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sgpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sgpq.withSecurityGroup; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*SecurityGroupPermission)
		for i := range nodes {
			if nodes[i].security_group_security_group_permissions == nil {
				continue
			}
			fk := *nodes[i].security_group_security_group_permissions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(securitygroup.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "security_group_security_group_permissions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SecurityGroup = n
			}
		}
	}

	if query := sgpq.withSecurityPermission; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*SecurityGroupPermission)
		for i := range nodes {
			if nodes[i].security_permission_security_group_permissions == nil {
				continue
			}
			fk := *nodes[i].security_permission_security_group_permissions
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(securitypermission.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "security_permission_security_group_permissions" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SecurityPermission = n
			}
		}
	}

	return nodes, nil
}

func (sgpq *SecurityGroupPermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sgpq.querySpec()
	return sqlgraph.CountNodes(ctx, sgpq.driver, _spec)
}

func (sgpq *SecurityGroupPermissionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sgpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sgpq *SecurityGroupPermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securitygrouppermission.Table,
			Columns: securitygrouppermission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securitygrouppermission.FieldID,
			},
		},
		From:   sgpq.sql,
		Unique: true,
	}
	if unique := sgpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sgpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, securitygrouppermission.FieldID)
		for i := range fields {
			if fields[i] != securitygrouppermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sgpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sgpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sgpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sgpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sgpq *SecurityGroupPermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sgpq.driver.Dialect())
	t1 := builder.Table(securitygrouppermission.Table)
	columns := sgpq.fields
	if len(columns) == 0 {
		columns = securitygrouppermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sgpq.sql != nil {
		selector = sgpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sgpq.predicates {
		p(selector)
	}
	for _, p := range sgpq.order {
		p(selector)
	}
	if offset := sgpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sgpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SecurityGroupPermissionGroupBy is the group-by builder for SecurityGroupPermission entities.
type SecurityGroupPermissionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgpgb *SecurityGroupPermissionGroupBy) Aggregate(fns ...AggregateFunc) *SecurityGroupPermissionGroupBy {
	sgpgb.fns = append(sgpgb.fns, fns...)
	return sgpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgpgb *SecurityGroupPermissionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgpgb.path(ctx)
	if err != nil {
		return err
	}
	sgpgb.sql = query
	return sgpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgpgb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) StringX(ctx context.Context) string {
	v, err := sgpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgpgb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) IntX(ctx context.Context) int {
	v, err := sgpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgpgb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgpgb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgpgb *SecurityGroupPermissionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgpgb *SecurityGroupPermissionGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgpgb *SecurityGroupPermissionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgpgb.fields {
		if !securitygrouppermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgpgb *SecurityGroupPermissionGroupBy) sqlQuery() *sql.Selector {
	selector := sgpgb.sql.Select()
	aggregation := make([]string, 0, len(sgpgb.fns))
	for _, fn := range sgpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgpgb.fields)+len(sgpgb.fns))
		for _, f := range sgpgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgpgb.fields...)...)
}

// SecurityGroupPermissionSelect is the builder for selecting fields of SecurityGroupPermission entities.
type SecurityGroupPermissionSelect struct {
	*SecurityGroupPermissionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sgps *SecurityGroupPermissionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sgps.prepareQuery(ctx); err != nil {
		return err
	}
	sgps.sql = sgps.SecurityGroupPermissionQuery.sqlQuery(ctx)
	return sgps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sgps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sgps.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sgps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) StringsX(ctx context.Context) []string {
	v, err := sgps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) StringX(ctx context.Context) string {
	v, err := sgps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sgps.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sgps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) IntsX(ctx context.Context) []int {
	v, err := sgps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) IntX(ctx context.Context) int {
	v, err := sgps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgps.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sgps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sgps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) Float64X(ctx context.Context) float64 {
	v, err := sgps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sgps.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupPermissionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sgps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) BoolsX(ctx context.Context) []bool {
	v, err := sgps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sgps *SecurityGroupPermissionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygrouppermission.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupPermissionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgps *SecurityGroupPermissionSelect) BoolX(ctx context.Context) bool {
	v, err := sgps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgps *SecurityGroupPermissionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgps.sql.Query()
	if err := sgps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}