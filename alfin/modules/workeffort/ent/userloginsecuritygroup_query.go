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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userloginsecuritygroup"
)

// UserLoginSecurityGroupQuery is the builder for querying UserLoginSecurityGroup entities.
type UserLoginSecurityGroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserLoginSecurityGroup
	// eager-loading edges.
	withUserLogin                *UserLoginQuery
	withSecurityGroup            *SecurityGroupQuery
	withSecurityGroupPermissions *SecurityGroupPermissionQuery
	withFKs                      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserLoginSecurityGroupQuery builder.
func (ulsgq *UserLoginSecurityGroupQuery) Where(ps ...predicate.UserLoginSecurityGroup) *UserLoginSecurityGroupQuery {
	ulsgq.predicates = append(ulsgq.predicates, ps...)
	return ulsgq
}

// Limit adds a limit step to the query.
func (ulsgq *UserLoginSecurityGroupQuery) Limit(limit int) *UserLoginSecurityGroupQuery {
	ulsgq.limit = &limit
	return ulsgq
}

// Offset adds an offset step to the query.
func (ulsgq *UserLoginSecurityGroupQuery) Offset(offset int) *UserLoginSecurityGroupQuery {
	ulsgq.offset = &offset
	return ulsgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ulsgq *UserLoginSecurityGroupQuery) Unique(unique bool) *UserLoginSecurityGroupQuery {
	ulsgq.unique = &unique
	return ulsgq
}

// Order adds an order step to the query.
func (ulsgq *UserLoginSecurityGroupQuery) Order(o ...OrderFunc) *UserLoginSecurityGroupQuery {
	ulsgq.order = append(ulsgq.order, o...)
	return ulsgq
}

// QueryUserLogin chains the current query on the "user_login" edge.
func (ulsgq *UserLoginSecurityGroupQuery) QueryUserLogin() *UserLoginQuery {
	query := &UserLoginQuery{config: ulsgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ulsgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ulsgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userloginsecuritygroup.Table, userloginsecuritygroup.FieldID, selector),
			sqlgraph.To(userlogin.Table, userlogin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userloginsecuritygroup.UserLoginTable, userloginsecuritygroup.UserLoginColumn),
		)
		fromU = sqlgraph.SetNeighbors(ulsgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySecurityGroup chains the current query on the "security_group" edge.
func (ulsgq *UserLoginSecurityGroupQuery) QuerySecurityGroup() *SecurityGroupQuery {
	query := &SecurityGroupQuery{config: ulsgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ulsgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ulsgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userloginsecuritygroup.Table, userloginsecuritygroup.FieldID, selector),
			sqlgraph.To(securitygroup.Table, securitygroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userloginsecuritygroup.SecurityGroupTable, userloginsecuritygroup.SecurityGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(ulsgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySecurityGroupPermissions chains the current query on the "security_group_permissions" edge.
func (ulsgq *UserLoginSecurityGroupQuery) QuerySecurityGroupPermissions() *SecurityGroupPermissionQuery {
	query := &SecurityGroupPermissionQuery{config: ulsgq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ulsgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ulsgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userloginsecuritygroup.Table, userloginsecuritygroup.FieldID, selector),
			sqlgraph.To(securitygrouppermission.Table, securitygrouppermission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userloginsecuritygroup.SecurityGroupPermissionsTable, userloginsecuritygroup.SecurityGroupPermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ulsgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserLoginSecurityGroup entity from the query.
// Returns a *NotFoundError when no UserLoginSecurityGroup was found.
func (ulsgq *UserLoginSecurityGroupQuery) First(ctx context.Context) (*UserLoginSecurityGroup, error) {
	nodes, err := ulsgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userloginsecuritygroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) FirstX(ctx context.Context) *UserLoginSecurityGroup {
	node, err := ulsgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserLoginSecurityGroup ID from the query.
// Returns a *NotFoundError when no UserLoginSecurityGroup ID was found.
func (ulsgq *UserLoginSecurityGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulsgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userloginsecuritygroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := ulsgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserLoginSecurityGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one UserLoginSecurityGroup entity is not found.
// Returns a *NotFoundError when no UserLoginSecurityGroup entities are found.
func (ulsgq *UserLoginSecurityGroupQuery) Only(ctx context.Context) (*UserLoginSecurityGroup, error) {
	nodes, err := ulsgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userloginsecuritygroup.Label}
	default:
		return nil, &NotSingularError{userloginsecuritygroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) OnlyX(ctx context.Context) *UserLoginSecurityGroup {
	node, err := ulsgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserLoginSecurityGroup ID in the query.
// Returns a *NotSingularError when exactly one UserLoginSecurityGroup ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ulsgq *UserLoginSecurityGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulsgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = &NotSingularError{userloginsecuritygroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := ulsgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserLoginSecurityGroups.
func (ulsgq *UserLoginSecurityGroupQuery) All(ctx context.Context) ([]*UserLoginSecurityGroup, error) {
	if err := ulsgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ulsgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) AllX(ctx context.Context) []*UserLoginSecurityGroup {
	nodes, err := ulsgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserLoginSecurityGroup IDs.
func (ulsgq *UserLoginSecurityGroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ulsgq.Select(userloginsecuritygroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := ulsgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ulsgq *UserLoginSecurityGroupQuery) Count(ctx context.Context) (int, error) {
	if err := ulsgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ulsgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) CountX(ctx context.Context) int {
	count, err := ulsgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ulsgq *UserLoginSecurityGroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := ulsgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ulsgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ulsgq *UserLoginSecurityGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := ulsgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserLoginSecurityGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ulsgq *UserLoginSecurityGroupQuery) Clone() *UserLoginSecurityGroupQuery {
	if ulsgq == nil {
		return nil
	}
	return &UserLoginSecurityGroupQuery{
		config:                       ulsgq.config,
		limit:                        ulsgq.limit,
		offset:                       ulsgq.offset,
		order:                        append([]OrderFunc{}, ulsgq.order...),
		predicates:                   append([]predicate.UserLoginSecurityGroup{}, ulsgq.predicates...),
		withUserLogin:                ulsgq.withUserLogin.Clone(),
		withSecurityGroup:            ulsgq.withSecurityGroup.Clone(),
		withSecurityGroupPermissions: ulsgq.withSecurityGroupPermissions.Clone(),
		// clone intermediate query.
		sql:  ulsgq.sql.Clone(),
		path: ulsgq.path,
	}
}

// WithUserLogin tells the query-builder to eager-load the nodes that are connected to
// the "user_login" edge. The optional arguments are used to configure the query builder of the edge.
func (ulsgq *UserLoginSecurityGroupQuery) WithUserLogin(opts ...func(*UserLoginQuery)) *UserLoginSecurityGroupQuery {
	query := &UserLoginQuery{config: ulsgq.config}
	for _, opt := range opts {
		opt(query)
	}
	ulsgq.withUserLogin = query
	return ulsgq
}

// WithSecurityGroup tells the query-builder to eager-load the nodes that are connected to
// the "security_group" edge. The optional arguments are used to configure the query builder of the edge.
func (ulsgq *UserLoginSecurityGroupQuery) WithSecurityGroup(opts ...func(*SecurityGroupQuery)) *UserLoginSecurityGroupQuery {
	query := &SecurityGroupQuery{config: ulsgq.config}
	for _, opt := range opts {
		opt(query)
	}
	ulsgq.withSecurityGroup = query
	return ulsgq
}

// WithSecurityGroupPermissions tells the query-builder to eager-load the nodes that are connected to
// the "security_group_permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (ulsgq *UserLoginSecurityGroupQuery) WithSecurityGroupPermissions(opts ...func(*SecurityGroupPermissionQuery)) *UserLoginSecurityGroupQuery {
	query := &SecurityGroupPermissionQuery{config: ulsgq.config}
	for _, opt := range opts {
		opt(query)
	}
	ulsgq.withSecurityGroupPermissions = query
	return ulsgq
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
//	client.UserLoginSecurityGroup.Query().
//		GroupBy(userloginsecuritygroup.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ulsgq *UserLoginSecurityGroupQuery) GroupBy(field string, fields ...string) *UserLoginSecurityGroupGroupBy {
	group := &UserLoginSecurityGroupGroupBy{config: ulsgq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ulsgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ulsgq.sqlQuery(ctx), nil
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
//	client.UserLoginSecurityGroup.Query().
//		Select(userloginsecuritygroup.FieldCreateTime).
//		Scan(ctx, &v)
//
func (ulsgq *UserLoginSecurityGroupQuery) Select(field string, fields ...string) *UserLoginSecurityGroupSelect {
	ulsgq.fields = append([]string{field}, fields...)
	return &UserLoginSecurityGroupSelect{UserLoginSecurityGroupQuery: ulsgq}
}

func (ulsgq *UserLoginSecurityGroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ulsgq.fields {
		if !userloginsecuritygroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ulsgq.path != nil {
		prev, err := ulsgq.path(ctx)
		if err != nil {
			return err
		}
		ulsgq.sql = prev
	}
	return nil
}

func (ulsgq *UserLoginSecurityGroupQuery) sqlAll(ctx context.Context) ([]*UserLoginSecurityGroup, error) {
	var (
		nodes       = []*UserLoginSecurityGroup{}
		withFKs     = ulsgq.withFKs
		_spec       = ulsgq.querySpec()
		loadedTypes = [3]bool{
			ulsgq.withUserLogin != nil,
			ulsgq.withSecurityGroup != nil,
			ulsgq.withSecurityGroupPermissions != nil,
		}
	)
	if ulsgq.withUserLogin != nil || ulsgq.withSecurityGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userloginsecuritygroup.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserLoginSecurityGroup{config: ulsgq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ulsgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ulsgq.withUserLogin; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserLoginSecurityGroup)
		for i := range nodes {
			if nodes[i].user_login_user_login_security_groups == nil {
				continue
			}
			fk := *nodes[i].user_login_user_login_security_groups
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(userlogin.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_login_user_login_security_groups" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.UserLogin = n
			}
		}
	}

	if query := ulsgq.withSecurityGroup; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserLoginSecurityGroup)
		for i := range nodes {
			if nodes[i].security_group_user_login_security_groups == nil {
				continue
			}
			fk := *nodes[i].security_group_user_login_security_groups
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
				return nil, fmt.Errorf(`unexpected foreign-key "security_group_user_login_security_groups" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SecurityGroup = n
			}
		}
	}

	if query := ulsgq.withSecurityGroupPermissions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*UserLoginSecurityGroup)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.SecurityGroupPermissions = []*SecurityGroupPermission{}
		}
		query.withFKs = true
		query.Where(predicate.SecurityGroupPermission(func(s *sql.Selector) {
			s.Where(sql.InValues(userloginsecuritygroup.SecurityGroupPermissionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_login_security_group_security_group_permissions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_login_security_group_security_group_permissions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_login_security_group_security_group_permissions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.SecurityGroupPermissions = append(node.Edges.SecurityGroupPermissions, n)
		}
	}

	return nodes, nil
}

func (ulsgq *UserLoginSecurityGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ulsgq.querySpec()
	return sqlgraph.CountNodes(ctx, ulsgq.driver, _spec)
}

func (ulsgq *UserLoginSecurityGroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ulsgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ulsgq *UserLoginSecurityGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userloginsecuritygroup.Table,
			Columns: userloginsecuritygroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userloginsecuritygroup.FieldID,
			},
		},
		From:   ulsgq.sql,
		Unique: true,
	}
	if unique := ulsgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ulsgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userloginsecuritygroup.FieldID)
		for i := range fields {
			if fields[i] != userloginsecuritygroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ulsgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ulsgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ulsgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ulsgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ulsgq *UserLoginSecurityGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ulsgq.driver.Dialect())
	t1 := builder.Table(userloginsecuritygroup.Table)
	columns := ulsgq.fields
	if len(columns) == 0 {
		columns = userloginsecuritygroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ulsgq.sql != nil {
		selector = ulsgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ulsgq.predicates {
		p(selector)
	}
	for _, p := range ulsgq.order {
		p(selector)
	}
	if offset := ulsgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ulsgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserLoginSecurityGroupGroupBy is the group-by builder for UserLoginSecurityGroup entities.
type UserLoginSecurityGroupGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ulsggb *UserLoginSecurityGroupGroupBy) Aggregate(fns ...AggregateFunc) *UserLoginSecurityGroupGroupBy {
	ulsggb.fns = append(ulsggb.fns, fns...)
	return ulsggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ulsggb *UserLoginSecurityGroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ulsggb.path(ctx)
	if err != nil {
		return err
	}
	ulsggb.sql = query
	return ulsggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ulsggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ulsggb.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ulsggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := ulsggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ulsggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) StringX(ctx context.Context) string {
	v, err := ulsggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ulsggb.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ulsggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := ulsggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ulsggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) IntX(ctx context.Context) int {
	v, err := ulsggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ulsggb.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ulsggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ulsggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ulsggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ulsggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ulsggb.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ulsggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ulsggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ulsggb *UserLoginSecurityGroupGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ulsggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ulsggb *UserLoginSecurityGroupGroupBy) BoolX(ctx context.Context) bool {
	v, err := ulsggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ulsggb *UserLoginSecurityGroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ulsggb.fields {
		if !userloginsecuritygroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ulsggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ulsggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ulsggb *UserLoginSecurityGroupGroupBy) sqlQuery() *sql.Selector {
	selector := ulsggb.sql.Select()
	aggregation := make([]string, 0, len(ulsggb.fns))
	for _, fn := range ulsggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ulsggb.fields)+len(ulsggb.fns))
		for _, f := range ulsggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ulsggb.fields...)...)
}

// UserLoginSecurityGroupSelect is the builder for selecting fields of UserLoginSecurityGroup entities.
type UserLoginSecurityGroupSelect struct {
	*UserLoginSecurityGroupQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ulsgs *UserLoginSecurityGroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ulsgs.prepareQuery(ctx); err != nil {
		return err
	}
	ulsgs.sql = ulsgs.UserLoginSecurityGroupQuery.sqlQuery(ctx)
	return ulsgs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ulsgs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ulsgs.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ulsgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) StringsX(ctx context.Context) []string {
	v, err := ulsgs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ulsgs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) StringX(ctx context.Context) string {
	v, err := ulsgs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ulsgs.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ulsgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) IntsX(ctx context.Context) []int {
	v, err := ulsgs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ulsgs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) IntX(ctx context.Context) int {
	v, err := ulsgs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ulsgs.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ulsgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ulsgs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ulsgs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) Float64X(ctx context.Context) float64 {
	v, err := ulsgs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ulsgs.fields) > 1 {
		return nil, errors.New("ent: UserLoginSecurityGroupSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ulsgs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) BoolsX(ctx context.Context) []bool {
	v, err := ulsgs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ulsgs *UserLoginSecurityGroupSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ulsgs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userloginsecuritygroup.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSecurityGroupSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ulsgs *UserLoginSecurityGroupSelect) BoolX(ctx context.Context) bool {
	v, err := ulsgs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ulsgs *UserLoginSecurityGroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ulsgs.sql.Query()
	if err := ulsgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}