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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderstatus"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
)

// OrderStatusQuery is the builder for querying OrderStatus entities.
type OrderStatusQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderStatus
	// eager-loading edges.
	withStatusItem             *StatusItemQuery
	withOrderHeader            *OrderHeaderQuery
	withOrderItem              *OrderItemQuery
	withOrderPaymentPreference *OrderPaymentPreferenceQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStatusQuery builder.
func (osq *OrderStatusQuery) Where(ps ...predicate.OrderStatus) *OrderStatusQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit adds a limit step to the query.
func (osq *OrderStatusQuery) Limit(limit int) *OrderStatusQuery {
	osq.limit = &limit
	return osq
}

// Offset adds an offset step to the query.
func (osq *OrderStatusQuery) Offset(offset int) *OrderStatusQuery {
	osq.offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrderStatusQuery) Unique(unique bool) *OrderStatusQuery {
	osq.unique = &unique
	return osq
}

// Order adds an order step to the query.
func (osq *OrderStatusQuery) Order(o ...OrderFunc) *OrderStatusQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryStatusItem chains the current query on the "status_item" edge.
func (osq *OrderStatusQuery) QueryStatusItem() *StatusItemQuery {
	query := &StatusItemQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderstatus.Table, orderstatus.FieldID, selector),
			sqlgraph.To(statusitem.Table, statusitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderstatus.StatusItemTable, orderstatus.StatusItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderHeader chains the current query on the "order_header" edge.
func (osq *OrderStatusQuery) QueryOrderHeader() *OrderHeaderQuery {
	query := &OrderHeaderQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderstatus.Table, orderstatus.FieldID, selector),
			sqlgraph.To(orderheader.Table, orderheader.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderstatus.OrderHeaderTable, orderstatus.OrderHeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItem chains the current query on the "order_item" edge.
func (osq *OrderStatusQuery) QueryOrderItem() *OrderItemQuery {
	query := &OrderItemQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderstatus.Table, orderstatus.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderstatus.OrderItemTable, orderstatus.OrderItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderPaymentPreference chains the current query on the "order_payment_preference" edge.
func (osq *OrderStatusQuery) QueryOrderPaymentPreference() *OrderPaymentPreferenceQuery {
	query := &OrderPaymentPreferenceQuery{config: osq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderstatus.Table, orderstatus.FieldID, selector),
			sqlgraph.To(orderpaymentpreference.Table, orderpaymentpreference.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderstatus.OrderPaymentPreferenceTable, orderstatus.OrderPaymentPreferenceColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderStatus entity from the query.
// Returns a *NotFoundError when no OrderStatus was found.
func (osq *OrderStatusQuery) First(ctx context.Context) (*OrderStatus, error) {
	nodes, err := osq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrderStatusQuery) FirstX(ctx context.Context) *OrderStatus {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderStatus ID from the query.
// Returns a *NotFoundError when no OrderStatus ID was found.
func (osq *OrderStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrderStatusQuery) FirstIDX(ctx context.Context) int {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderStatus entity is not found.
// Returns a *NotFoundError when no OrderStatus entities are found.
func (osq *OrderStatusQuery) Only(ctx context.Context) (*OrderStatus, error) {
	nodes, err := osq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstatus.Label}
	default:
		return nil, &NotSingularError{orderstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrderStatusQuery) OnlyX(ctx context.Context) *OrderStatus {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderStatus ID in the query.
// Returns a *NotSingularError when exactly one OrderStatus ID is not found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrderStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = &NotSingularError{orderstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrderStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStatusSlice.
func (osq *OrderStatusQuery) All(ctx context.Context) ([]*OrderStatus, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return osq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrderStatusQuery) AllX(ctx context.Context) []*OrderStatus {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderStatus IDs.
func (osq *OrderStatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := osq.Select(orderstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrderStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrderStatusQuery) Count(ctx context.Context) (int, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return osq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrderStatusQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrderStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return osq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrderStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrderStatusQuery) Clone() *OrderStatusQuery {
	if osq == nil {
		return nil
	}
	return &OrderStatusQuery{
		config:                     osq.config,
		limit:                      osq.limit,
		offset:                     osq.offset,
		order:                      append([]OrderFunc{}, osq.order...),
		predicates:                 append([]predicate.OrderStatus{}, osq.predicates...),
		withStatusItem:             osq.withStatusItem.Clone(),
		withOrderHeader:            osq.withOrderHeader.Clone(),
		withOrderItem:              osq.withOrderItem.Clone(),
		withOrderPaymentPreference: osq.withOrderPaymentPreference.Clone(),
		// clone intermediate query.
		sql:  osq.sql.Clone(),
		path: osq.path,
	}
}

// WithStatusItem tells the query-builder to eager-load the nodes that are connected to
// the "status_item" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithStatusItem(opts ...func(*StatusItemQuery)) *OrderStatusQuery {
	query := &StatusItemQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withStatusItem = query
	return osq
}

// WithOrderHeader tells the query-builder to eager-load the nodes that are connected to
// the "order_header" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithOrderHeader(opts ...func(*OrderHeaderQuery)) *OrderStatusQuery {
	query := &OrderHeaderQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrderHeader = query
	return osq
}

// WithOrderItem tells the query-builder to eager-load the nodes that are connected to
// the "order_item" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithOrderItem(opts ...func(*OrderItemQuery)) *OrderStatusQuery {
	query := &OrderItemQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrderItem = query
	return osq
}

// WithOrderPaymentPreference tells the query-builder to eager-load the nodes that are connected to
// the "order_payment_preference" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithOrderPaymentPreference(opts ...func(*OrderPaymentPreferenceQuery)) *OrderStatusQuery {
	query := &OrderPaymentPreferenceQuery{config: osq.config}
	for _, opt := range opts {
		opt(query)
	}
	osq.withOrderPaymentPreference = query
	return osq
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
//	client.OrderStatus.Query().
//		GroupBy(orderstatus.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (osq *OrderStatusQuery) GroupBy(field string, fields ...string) *OrderStatusGroupBy {
	group := &OrderStatusGroupBy{config: osq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osq.sqlQuery(ctx), nil
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
//	client.OrderStatus.Query().
//		Select(orderstatus.FieldCreateTime).
//		Scan(ctx, &v)
//
func (osq *OrderStatusQuery) Select(field string, fields ...string) *OrderStatusSelect {
	osq.fields = append([]string{field}, fields...)
	return &OrderStatusSelect{OrderStatusQuery: osq}
}

func (osq *OrderStatusQuery) prepareQuery(ctx context.Context) error {
	for _, f := range osq.fields {
		if !orderstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OrderStatusQuery) sqlAll(ctx context.Context) ([]*OrderStatus, error) {
	var (
		nodes       = []*OrderStatus{}
		withFKs     = osq.withFKs
		_spec       = osq.querySpec()
		loadedTypes = [4]bool{
			osq.withStatusItem != nil,
			osq.withOrderHeader != nil,
			osq.withOrderItem != nil,
			osq.withOrderPaymentPreference != nil,
		}
	)
	if osq.withStatusItem != nil || osq.withOrderHeader != nil || osq.withOrderItem != nil || osq.withOrderPaymentPreference != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatus.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderStatus{config: osq.config}
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
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := osq.withStatusItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderStatus)
		for i := range nodes {
			if nodes[i].status_item_order_statuses == nil {
				continue
			}
			fk := *nodes[i].status_item_order_statuses
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
				return nil, fmt.Errorf(`unexpected foreign-key "status_item_order_statuses" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.StatusItem = n
			}
		}
	}

	if query := osq.withOrderHeader; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderStatus)
		for i := range nodes {
			if nodes[i].order_header_order_statuses == nil {
				continue
			}
			fk := *nodes[i].order_header_order_statuses
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderheader.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_header_order_statuses" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderHeader = n
			}
		}
	}

	if query := osq.withOrderItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderStatus)
		for i := range nodes {
			if nodes[i].order_item_order_statuses == nil {
				continue
			}
			fk := *nodes[i].order_item_order_statuses
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderitem.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_item_order_statuses" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderItem = n
			}
		}
	}

	if query := osq.withOrderPaymentPreference; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderStatus)
		for i := range nodes {
			if nodes[i].order_payment_preference_order_statuses == nil {
				continue
			}
			fk := *nodes[i].order_payment_preference_order_statuses
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderpaymentpreference.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_payment_preference_order_statuses" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderPaymentPreference = n
			}
		}
	}

	return nodes, nil
}

func (osq *OrderStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrderStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := osq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (osq *OrderStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatus.Table,
			Columns: orderstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderstatus.FieldID,
			},
		},
		From:   osq.sql,
		Unique: true,
	}
	if unique := osq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := osq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatus.FieldID)
		for i := range fields {
			if fields[i] != orderstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrderStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(orderstatus.Table)
	columns := osq.fields
	if len(columns) == 0 {
		columns = orderstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderStatusGroupBy is the group-by builder for OrderStatus entities.
type OrderStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrderStatusGroupBy) Aggregate(fns ...AggregateFunc) *OrderStatusGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the group-by query and scans the result into the given value.
func (osgb *OrderStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := osgb.path(ctx)
	if err != nil {
		return err
	}
	osgb.sql = query
	return osgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := osgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OrderStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := osgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = osgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) StringX(ctx context.Context) string {
	v, err := osgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OrderStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := osgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = osgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) IntX(ctx context.Context) int {
	v, err := osgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OrderStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := osgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = osgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := osgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(osgb.fields) > 1 {
		return nil, errors.New("ent: OrderStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := osgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := osgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (osgb *OrderStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = osgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (osgb *OrderStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := osgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (osgb *OrderStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range osgb.fields {
		if !orderstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := osgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (osgb *OrderStatusGroupBy) sqlQuery() *sql.Selector {
	selector := osgb.sql.Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(osgb.fields)+len(osgb.fns))
		for _, f := range osgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(osgb.fields...)...)
}

// OrderStatusSelect is the builder for selecting fields of OrderStatus entities.
type OrderStatusSelect struct {
	*OrderStatusQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrderStatusSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	oss.sql = oss.OrderStatusQuery.sqlQuery(ctx)
	return oss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oss *OrderStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OrderStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oss *OrderStatusSelect) StringsX(ctx context.Context) []string {
	v, err := oss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oss *OrderStatusSelect) StringX(ctx context.Context) string {
	v, err := oss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OrderStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oss *OrderStatusSelect) IntsX(ctx context.Context) []int {
	v, err := oss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oss *OrderStatusSelect) IntX(ctx context.Context) int {
	v, err := oss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OrderStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oss *OrderStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oss *OrderStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := oss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oss.fields) > 1 {
		return nil, errors.New("ent: OrderStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oss *OrderStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := oss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oss *OrderStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = fmt.Errorf("ent: OrderStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oss *OrderStatusSelect) BoolX(ctx context.Context) bool {
	v, err := oss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oss *OrderStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oss.sql.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
