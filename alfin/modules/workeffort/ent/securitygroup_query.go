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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygrouppermission"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userloginsecuritygroup"
)

// SecurityGroupQuery is the builder for querying SecurityGroup entities.
type SecurityGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SecurityGroup
	// eager-loading edges.
	withSecurityGroupPermissions *SecurityGroupPermissionQuery
	withUserLoginSecurityGroups  *UserLoginSecurityGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SecurityGroupQuery builder.
func (sgq *SecurityGroupQuery) Where(ps ...predicate.SecurityGroup) *SecurityGroupQuery {
	sgq.predicates = append(sgq.predicates, ps...)
	return sgq
}

// Limit adds a limit step to the query.
func (sgq *SecurityGroupQuery) Limit(limit int) *SecurityGroupQuery {
	sgq.limit = &limit
	return sgq
}

// Offset adds an offset step to the query.
func (sgq *SecurityGroupQuery) Offset(offset int) *SecurityGroupQuery {
	sgq.offset = &offset
	return sgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sgq *SecurityGroupQuery) Unique(unique bool) *SecurityGroupQuery {
	sgq.unique = &unique
	return sgq
}

// Order adds an order step to the query.
func (sgq *SecurityGroupQuery) Order(o ...OrderFunc) *SecurityGroupQuery {
	sgq.order = append(sgq.order, o...)
	return sgq
}

// QuerySecurityGroupPermissions chains the current query on the "security_group_permissions" edge.
func (sgq *SecurityGroupQuery) QuerySecurityGroupPermissions() *SecurityGroupPermissionQuery {
	query := &SecurityGroupPermissionQuery{config: sgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(securitygroup.Table, securitygroup.FieldID, selector),
			sqlgraph.To(securitygrouppermission.Table, securitygrouppermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, securitygroup.SecurityGroupPermissionsTable, securitygroup.SecurityGroupPermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserLoginSecurityGroups chains the current query on the "user_login_security_groups" edge.
func (sgq *SecurityGroupQuery) QueryUserLoginSecurityGroups() *UserLoginSecurityGroupQuery {
	query := &UserLoginSecurityGroupQuery{config: sgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(securitygroup.Table, securitygroup.FieldID, selector),
			sqlgraph.To(userloginsecuritygroup.Table, userloginsecuritygroup.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, securitygroup.UserLoginSecurityGroupsTable, securitygroup.UserLoginSecurityGroupsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SecurityGroup entity from the query.
// Returns a *NotFoundError when no SecurityGroup was found.
func (sgq *SecurityGroupQuery) First(ctx context.Context) (*SecurityGroup, error) {
	nodes, err := sgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{securitygroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sgq *SecurityGroupQuery) FirstX(ctx context.Context) *SecurityGroup {
	node, err := sgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SecurityGroup ID from the query.
// Returns a *NotFoundError when no SecurityGroup ID was found.
func (sgq *SecurityGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{securitygroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sgq *SecurityGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := sgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SecurityGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SecurityGroup entity is not found.
// Returns a *NotFoundError when no SecurityGroup entities are found.
func (sgq *SecurityGroupQuery) Only(ctx context.Context) (*SecurityGroup, error) {
	nodes, err := sgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{securitygroup.Label}
	default:
		return nil, &NotSingularError{securitygroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sgq *SecurityGroupQuery) OnlyX(ctx context.Context) *SecurityGroup {
	node, err := sgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SecurityGroup ID in the query.
// Returns a *NotSingularError when exactly one SecurityGroup ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sgq *SecurityGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = &NotSingularError{securitygroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sgq *SecurityGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := sgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SecurityGroups.
func (sgq *SecurityGroupQuery) All(ctx context.Context) ([]*SecurityGroup, error) {
	if err := sgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sgq *SecurityGroupQuery) AllX(ctx context.Context) []*SecurityGroup {
	nodes, err := sgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SecurityGroup IDs.
func (sgq *SecurityGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sgq.Select(securitygroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sgq *SecurityGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := sgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sgq *SecurityGroupQuery) Count(ctx context.Context) (int, error) {
	if err := sgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sgq *SecurityGroupQuery) CountX(ctx context.Context) int {
	count, err := sgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sgq *SecurityGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := sgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sgq *SecurityGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := sgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SecurityGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sgq *SecurityGroupQuery) Clone() *SecurityGroupQuery {
	if sgq == nil {
		return nil
	}
	return &SecurityGroupQuery{
		config:                       sgq.config,
		limit:                        sgq.limit,
		offset:                       sgq.offset,
		order:                        append([]OrderFunc{}, sgq.order...),
		predicates:                   append([]predicate.SecurityGroup{}, sgq.predicates...),
		withSecurityGroupPermissions: sgq.withSecurityGroupPermissions.Clone(),
		withUserLoginSecurityGroups:  sgq.withUserLoginSecurityGroups.Clone(),
		// clone intermediate query.
		sql:  sgq.sql.Clone(),
		path: sgq.path,
	}
}

// WithSecurityGroupPermissions tells the query-builder to eager-load the nodes that are connected to
// the "security_group_permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (sgq *SecurityGroupQuery) WithSecurityGroupPermissions(opts ...func(*SecurityGroupPermissionQuery)) *SecurityGroupQuery {
	query := &SecurityGroupPermissionQuery{config: sgq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgq.withSecurityGroupPermissions = query
	return sgq
}

// WithUserLoginSecurityGroups tells the query-builder to eager-load the nodes that are connected to
// the "user_login_security_groups" edge. The optional arguments are used to configure the query builder of the edge.
func (sgq *SecurityGroupQuery) WithUserLoginSecurityGroups(opts ...func(*UserLoginSecurityGroupQuery)) *SecurityGroupQuery {
	query := &UserLoginSecurityGroupQuery{config: sgq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgq.withUserLoginSecurityGroups = query
	return sgq
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
//	client.SecurityGroup.Query().
//		GroupBy(securitygroup.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sgq *SecurityGroupQuery) GroupBy(field string, fields ...string) *SecurityGroupGroupBy {
	group := &SecurityGroupGroupBy{config: sgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sgq.sqlQuery(ctx), nil
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
//	client.SecurityGroup.Query().
//		Select(securitygroup.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sgq *SecurityGroupQuery) Select(field string, fields ...string) *SecurityGroupSelect {
	sgq.fields = append([]string{field}, fields...)
	return &SecurityGroupSelect{SecurityGroupQuery: sgq}
}

func (sgq *SecurityGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sgq.fields {
		if !securitygroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sgq.path != nil {
		prev, err := sgq.path(ctx)
		if err != nil {
			return err
		}
		sgq.sql = prev
	}
	return nil
}

func (sgq *SecurityGroupQuery) sqlAll(ctx context.Context) ([]*SecurityGroup, error) {
	var (
		nodes       = []*SecurityGroup{}
		_spec       = sgq.querySpec()
		loadedTypes = [2]bool{
			sgq.withSecurityGroupPermissions != nil,
			sgq.withUserLoginSecurityGroups != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SecurityGroup{config: sgq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sgq.withSecurityGroupPermissions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*SecurityGroup)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.SecurityGroupPermissions = []*SecurityGroupPermission{}
		}
		query.withFKs = true
		query.Where(predicate.SecurityGroupPermission(func(s *sql.Selector) {
			s.Where(sql.InValues(securitygroup.SecurityGroupPermissionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.security_group_security_group_permissions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "security_group_security_group_permissions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "security_group_security_group_permissions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.SecurityGroupPermissions = append(node.Edges.SecurityGroupPermissions, n)
		}
	}

	if query := sgq.withUserLoginSecurityGroups; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*SecurityGroup)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.UserLoginSecurityGroups = []*UserLoginSecurityGroup{}
		}
		query.withFKs = true
		query.Where(predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
			s.Where(sql.InValues(securitygroup.UserLoginSecurityGroupsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.security_group_user_login_security_groups
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "security_group_user_login_security_groups" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "security_group_user_login_security_groups" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.UserLoginSecurityGroups = append(node.Edges.UserLoginSecurityGroups, n)
		}
	}

	return nodes, nil
}

func (sgq *SecurityGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sgq.querySpec()
	return sqlgraph.CountNodes(ctx, sgq.driver, _spec)
}

func (sgq *SecurityGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sgq *SecurityGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securitygroup.Table,
			Columns: securitygroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securitygroup.FieldID,
			},
		},
		From:   sgq.sql,
		Unique: true,
	}
	if unique := sgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, securitygroup.FieldID)
		for i := range fields {
			if fields[i] != securitygroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sgq *SecurityGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sgq.driver.Dialect())
	t1 := builder.Table(securitygroup.Table)
	columns := sgq.fields
	if len(columns) == 0 {
		columns = securitygroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sgq.sql != nil {
		selector = sgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sgq.predicates {
		p(selector)
	}
	for _, p := range sgq.order {
		p(selector)
	}
	if offset := sgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SecurityGroupGroupBy is the group-by builder for SecurityGroup entities.
type SecurityGroupGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sggb *SecurityGroupGroupBy) Aggregate(fns ...AggregateFunc) *SecurityGroupGroupBy {
	sggb.fns = append(sggb.fns, fns...)
	return sggb
}

// Scan applies the group-by query and scans the result into the given value.
func (sggb *SecurityGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sggb.path(ctx)
	if err != nil {
		return err
	}
	sggb.sql = query
	return sggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sggb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := sggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) StringX(ctx context.Context) string {
	v, err := sggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sggb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := sggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) IntX(ctx context.Context) int {
	v, err := sggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sggb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sggb.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sggb *SecurityGroupGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sggb *SecurityGroupGroupBy) BoolX(ctx context.Context) bool {
	v, err := sggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sggb *SecurityGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sggb.fields {
		if !securitygroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sggb *SecurityGroupGroupBy) sqlQuery() *sql.Selector {
	selector := sggb.sql.Select()
	aggregation := make([]string, 0, len(sggb.fns))
	for _, fn := range sggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sggb.fields)+len(sggb.fns))
		for _, f := range sggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sggb.fields...)...)
}

// SecurityGroupSelect is the builder for selecting fields of SecurityGroup entities.
type SecurityGroupSelect struct {
	*SecurityGroupQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sgs *SecurityGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sgs.prepareQuery(ctx); err != nil {
		return err
	}
	sgs.sql = sgs.SecurityGroupQuery.sqlQuery(ctx)
	return sgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgs *SecurityGroupSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sgs.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgs *SecurityGroupSelect) StringsX(ctx context.Context) []string {
	v, err := sgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgs *SecurityGroupSelect) StringX(ctx context.Context) string {
	v, err := sgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sgs.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgs *SecurityGroupSelect) IntsX(ctx context.Context) []int {
	v, err := sgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgs *SecurityGroupSelect) IntX(ctx context.Context) int {
	v, err := sgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgs.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgs *SecurityGroupSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgs *SecurityGroupSelect) Float64X(ctx context.Context) float64 {
	v, err := sgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sgs.fields) > 1 {
		return nil, errors.New("ent: SecurityGroupSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgs *SecurityGroupSelect) BoolsX(ctx context.Context) []bool {
	v, err := sgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sgs *SecurityGroupSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{securitygroup.Label}
	default:
		err = fmt.Errorf("ent: SecurityGroupSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgs *SecurityGroupSelect) BoolX(ctx context.Context) bool {
	v, err := sgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgs *SecurityGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgs.sql.Query()
	if err := sgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}