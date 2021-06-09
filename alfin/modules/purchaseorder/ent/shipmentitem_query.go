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
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/itemissuance"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipment"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/shipmentitem"
)

// ShipmentItemQuery is the builder for querying ShipmentItem entities.
type ShipmentItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShipmentItem
	// eager-loading edges.
	withShipment      *ShipmentQuery
	withItemIssuances *ItemIssuanceQuery
	withFKs           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentItemQuery builder.
func (siq *ShipmentItemQuery) Where(ps ...predicate.ShipmentItem) *ShipmentItemQuery {
	siq.predicates = append(siq.predicates, ps...)
	return siq
}

// Limit adds a limit step to the query.
func (siq *ShipmentItemQuery) Limit(limit int) *ShipmentItemQuery {
	siq.limit = &limit
	return siq
}

// Offset adds an offset step to the query.
func (siq *ShipmentItemQuery) Offset(offset int) *ShipmentItemQuery {
	siq.offset = &offset
	return siq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (siq *ShipmentItemQuery) Unique(unique bool) *ShipmentItemQuery {
	siq.unique = &unique
	return siq
}

// Order adds an order step to the query.
func (siq *ShipmentItemQuery) Order(o ...OrderFunc) *ShipmentItemQuery {
	siq.order = append(siq.order, o...)
	return siq
}

// QueryShipment chains the current query on the "shipment" edge.
func (siq *ShipmentItemQuery) QueryShipment() *ShipmentQuery {
	query := &ShipmentQuery{config: siq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := siq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentitem.Table, shipmentitem.FieldID, selector),
			sqlgraph.To(shipment.Table, shipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmentitem.ShipmentTable, shipmentitem.ShipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(siq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryItemIssuances chains the current query on the "item_issuances" edge.
func (siq *ShipmentItemQuery) QueryItemIssuances() *ItemIssuanceQuery {
	query := &ItemIssuanceQuery{config: siq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := siq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentitem.Table, shipmentitem.FieldID, selector),
			sqlgraph.To(itemissuance.Table, itemissuance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, shipmentitem.ItemIssuancesTable, shipmentitem.ItemIssuancesColumn),
		)
		fromU = sqlgraph.SetNeighbors(siq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentItem entity from the query.
// Returns a *NotFoundError when no ShipmentItem was found.
func (siq *ShipmentItemQuery) First(ctx context.Context) (*ShipmentItem, error) {
	nodes, err := siq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (siq *ShipmentItemQuery) FirstX(ctx context.Context) *ShipmentItem {
	node, err := siq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentItem ID from the query.
// Returns a *NotFoundError when no ShipmentItem ID was found.
func (siq *ShipmentItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = siq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (siq *ShipmentItemQuery) FirstIDX(ctx context.Context) int {
	id, err := siq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShipmentItem entity is not found.
// Returns a *NotFoundError when no ShipmentItem entities are found.
func (siq *ShipmentItemQuery) Only(ctx context.Context) (*ShipmentItem, error) {
	nodes, err := siq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentitem.Label}
	default:
		return nil, &NotSingularError{shipmentitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (siq *ShipmentItemQuery) OnlyX(ctx context.Context) *ShipmentItem {
	node, err := siq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentItem ID in the query.
// Returns a *NotSingularError when exactly one ShipmentItem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (siq *ShipmentItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = siq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = &NotSingularError{shipmentitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (siq *ShipmentItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := siq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentItems.
func (siq *ShipmentItemQuery) All(ctx context.Context) ([]*ShipmentItem, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return siq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (siq *ShipmentItemQuery) AllX(ctx context.Context) []*ShipmentItem {
	nodes, err := siq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentItem IDs.
func (siq *ShipmentItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := siq.Select(shipmentitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (siq *ShipmentItemQuery) IDsX(ctx context.Context) []int {
	ids, err := siq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (siq *ShipmentItemQuery) Count(ctx context.Context) (int, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return siq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (siq *ShipmentItemQuery) CountX(ctx context.Context) int {
	count, err := siq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (siq *ShipmentItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return siq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (siq *ShipmentItemQuery) ExistX(ctx context.Context) bool {
	exist, err := siq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (siq *ShipmentItemQuery) Clone() *ShipmentItemQuery {
	if siq == nil {
		return nil
	}
	return &ShipmentItemQuery{
		config:            siq.config,
		limit:             siq.limit,
		offset:            siq.offset,
		order:             append([]OrderFunc{}, siq.order...),
		predicates:        append([]predicate.ShipmentItem{}, siq.predicates...),
		withShipment:      siq.withShipment.Clone(),
		withItemIssuances: siq.withItemIssuances.Clone(),
		// clone intermediate query.
		sql:  siq.sql.Clone(),
		path: siq.path,
	}
}

// WithShipment tells the query-builder to eager-load the nodes that are connected to
// the "shipment" edge. The optional arguments are used to configure the query builder of the edge.
func (siq *ShipmentItemQuery) WithShipment(opts ...func(*ShipmentQuery)) *ShipmentItemQuery {
	query := &ShipmentQuery{config: siq.config}
	for _, opt := range opts {
		opt(query)
	}
	siq.withShipment = query
	return siq
}

// WithItemIssuances tells the query-builder to eager-load the nodes that are connected to
// the "item_issuances" edge. The optional arguments are used to configure the query builder of the edge.
func (siq *ShipmentItemQuery) WithItemIssuances(opts ...func(*ItemIssuanceQuery)) *ShipmentItemQuery {
	query := &ItemIssuanceQuery{config: siq.config}
	for _, opt := range opts {
		opt(query)
	}
	siq.withItemIssuances = query
	return siq
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
//	client.ShipmentItem.Query().
//		GroupBy(shipmentitem.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (siq *ShipmentItemQuery) GroupBy(field string, fields ...string) *ShipmentItemGroupBy {
	group := &ShipmentItemGroupBy{config: siq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return siq.sqlQuery(ctx), nil
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
//	client.ShipmentItem.Query().
//		Select(shipmentitem.FieldCreateTime).
//		Scan(ctx, &v)
//
func (siq *ShipmentItemQuery) Select(field string, fields ...string) *ShipmentItemSelect {
	siq.fields = append([]string{field}, fields...)
	return &ShipmentItemSelect{ShipmentItemQuery: siq}
}

func (siq *ShipmentItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range siq.fields {
		if !shipmentitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if siq.path != nil {
		prev, err := siq.path(ctx)
		if err != nil {
			return err
		}
		siq.sql = prev
	}
	return nil
}

func (siq *ShipmentItemQuery) sqlAll(ctx context.Context) ([]*ShipmentItem, error) {
	var (
		nodes       = []*ShipmentItem{}
		withFKs     = siq.withFKs
		_spec       = siq.querySpec()
		loadedTypes = [2]bool{
			siq.withShipment != nil,
			siq.withItemIssuances != nil,
		}
	)
	if siq.withShipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShipmentItem{config: siq.config}
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
	if err := sqlgraph.QueryNodes(ctx, siq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := siq.withShipment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ShipmentItem)
		for i := range nodes {
			if nodes[i].shipment_shipment_items == nil {
				continue
			}
			fk := *nodes[i].shipment_shipment_items
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(shipment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_shipment_items" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Shipment = n
			}
		}
	}

	if query := siq.withItemIssuances; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ShipmentItem)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ItemIssuances = []*ItemIssuance{}
		}
		query.withFKs = true
		query.Where(predicate.ItemIssuance(func(s *sql.Selector) {
			s.Where(sql.InValues(shipmentitem.ItemIssuancesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.shipment_item_item_issuances
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "shipment_item_item_issuances" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_item_item_issuances" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ItemIssuances = append(node.Edges.ItemIssuances, n)
		}
	}

	return nodes, nil
}

func (siq *ShipmentItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := siq.querySpec()
	return sqlgraph.CountNodes(ctx, siq.driver, _spec)
}

func (siq *ShipmentItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := siq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (siq *ShipmentItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shipmentitem.Table,
			Columns: shipmentitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentitem.FieldID,
			},
		},
		From:   siq.sql,
		Unique: true,
	}
	if unique := siq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := siq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentitem.FieldID)
		for i := range fields {
			if fields[i] != shipmentitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := siq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := siq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := siq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := siq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (siq *ShipmentItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(siq.driver.Dialect())
	t1 := builder.Table(shipmentitem.Table)
	selector := builder.Select(t1.Columns(shipmentitem.Columns...)...).From(t1)
	if siq.sql != nil {
		selector = siq.sql
		selector.Select(selector.Columns(shipmentitem.Columns...)...)
	}
	for _, p := range siq.predicates {
		p(selector)
	}
	for _, p := range siq.order {
		p(selector)
	}
	if offset := siq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := siq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentItemGroupBy is the group-by builder for ShipmentItem entities.
type ShipmentItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sigb *ShipmentItemGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentItemGroupBy {
	sigb.fns = append(sigb.fns, fns...)
	return sigb
}

// Scan applies the group-by query and scans the result into the given value.
func (sigb *ShipmentItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sigb.path(ctx)
	if err != nil {
		return err
	}
	sigb.sql = query
	return sigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := sigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) StringX(ctx context.Context) string {
	v, err := sigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := sigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) IntX(ctx context.Context) int {
	v, err := sigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sigb *ShipmentItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sigb *ShipmentItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := sigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sigb *ShipmentItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sigb.fields {
		if !shipmentitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sigb *ShipmentItemGroupBy) sqlQuery() *sql.Selector {
	selector := sigb.sql
	columns := make([]string, 0, len(sigb.fields)+len(sigb.fns))
	columns = append(columns, sigb.fields...)
	for _, fn := range sigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sigb.fields...)
}

// ShipmentItemSelect is the builder for selecting fields of ShipmentItem entities.
type ShipmentItemSelect struct {
	*ShipmentItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sis *ShipmentItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sis.prepareQuery(ctx); err != nil {
		return err
	}
	sis.sql = sis.ShipmentItemQuery.sqlQuery(ctx)
	return sis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sis *ShipmentItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sis *ShipmentItemSelect) StringsX(ctx context.Context) []string {
	v, err := sis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sis *ShipmentItemSelect) StringX(ctx context.Context) string {
	v, err := sis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sis *ShipmentItemSelect) IntsX(ctx context.Context) []int {
	v, err := sis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sis *ShipmentItemSelect) IntX(ctx context.Context) int {
	v, err := sis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sis *ShipmentItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sis *ShipmentItemSelect) Float64X(ctx context.Context) float64 {
	v, err := sis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ShipmentItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sis *ShipmentItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := sis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sis *ShipmentItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentitem.Label}
	default:
		err = fmt.Errorf("ent: ShipmentItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sis *ShipmentItemSelect) BoolX(ctx context.Context) bool {
	v, err := sis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sis *ShipmentItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sis.sqlQuery().Query()
	if err := sis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sis *ShipmentItemSelect) sqlQuery() sql.Querier {
	selector := sis.sql
	selector.Select(selector.Columns(sis.fields...)...)
	return selector
}
