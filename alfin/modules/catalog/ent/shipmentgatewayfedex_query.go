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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfig"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayfedex"
)

// ShipmentGatewayFedexQuery is the builder for querying ShipmentGatewayFedex entities.
type ShipmentGatewayFedexQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShipmentGatewayFedex
	// eager-loading edges.
	withShipmentGatewayConfig *ShipmentGatewayConfigQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentGatewayFedexQuery builder.
func (sgfq *ShipmentGatewayFedexQuery) Where(ps ...predicate.ShipmentGatewayFedex) *ShipmentGatewayFedexQuery {
	sgfq.predicates = append(sgfq.predicates, ps...)
	return sgfq
}

// Limit adds a limit step to the query.
func (sgfq *ShipmentGatewayFedexQuery) Limit(limit int) *ShipmentGatewayFedexQuery {
	sgfq.limit = &limit
	return sgfq
}

// Offset adds an offset step to the query.
func (sgfq *ShipmentGatewayFedexQuery) Offset(offset int) *ShipmentGatewayFedexQuery {
	sgfq.offset = &offset
	return sgfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sgfq *ShipmentGatewayFedexQuery) Unique(unique bool) *ShipmentGatewayFedexQuery {
	sgfq.unique = &unique
	return sgfq
}

// Order adds an order step to the query.
func (sgfq *ShipmentGatewayFedexQuery) Order(o ...OrderFunc) *ShipmentGatewayFedexQuery {
	sgfq.order = append(sgfq.order, o...)
	return sgfq
}

// QueryShipmentGatewayConfig chains the current query on the "shipment_gateway_config" edge.
func (sgfq *ShipmentGatewayFedexQuery) QueryShipmentGatewayConfig() *ShipmentGatewayConfigQuery {
	query := &ShipmentGatewayConfigQuery{config: sgfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentgatewayfedex.Table, shipmentgatewayfedex.FieldID, selector),
			sqlgraph.To(shipmentgatewayconfig.Table, shipmentgatewayconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, shipmentgatewayfedex.ShipmentGatewayConfigTable, shipmentgatewayfedex.ShipmentGatewayConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentGatewayFedex entity from the query.
// Returns a *NotFoundError when no ShipmentGatewayFedex was found.
func (sgfq *ShipmentGatewayFedexQuery) First(ctx context.Context) (*ShipmentGatewayFedex, error) {
	nodes, err := sgfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentgatewayfedex.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) FirstX(ctx context.Context) *ShipmentGatewayFedex {
	node, err := sgfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentGatewayFedex ID from the query.
// Returns a *NotFoundError when no ShipmentGatewayFedex ID was found.
func (sgfq *ShipmentGatewayFedexQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentgatewayfedex.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) FirstIDX(ctx context.Context) int {
	id, err := sgfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentGatewayFedex entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShipmentGatewayFedex entity is not found.
// Returns a *NotFoundError when no ShipmentGatewayFedex entities are found.
func (sgfq *ShipmentGatewayFedexQuery) Only(ctx context.Context) (*ShipmentGatewayFedex, error) {
	nodes, err := sgfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentgatewayfedex.Label}
	default:
		return nil, &NotSingularError{shipmentgatewayfedex.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) OnlyX(ctx context.Context) *ShipmentGatewayFedex {
	node, err := sgfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentGatewayFedex ID in the query.
// Returns a *NotSingularError when exactly one ShipmentGatewayFedex ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sgfq *ShipmentGatewayFedexQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = &NotSingularError{shipmentgatewayfedex.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) OnlyIDX(ctx context.Context) int {
	id, err := sgfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentGatewayFedexes.
func (sgfq *ShipmentGatewayFedexQuery) All(ctx context.Context) ([]*ShipmentGatewayFedex, error) {
	if err := sgfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sgfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) AllX(ctx context.Context) []*ShipmentGatewayFedex {
	nodes, err := sgfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentGatewayFedex IDs.
func (sgfq *ShipmentGatewayFedexQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sgfq.Select(shipmentgatewayfedex.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) IDsX(ctx context.Context) []int {
	ids, err := sgfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sgfq *ShipmentGatewayFedexQuery) Count(ctx context.Context) (int, error) {
	if err := sgfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sgfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) CountX(ctx context.Context) int {
	count, err := sgfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sgfq *ShipmentGatewayFedexQuery) Exist(ctx context.Context) (bool, error) {
	if err := sgfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sgfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sgfq *ShipmentGatewayFedexQuery) ExistX(ctx context.Context) bool {
	exist, err := sgfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentGatewayFedexQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sgfq *ShipmentGatewayFedexQuery) Clone() *ShipmentGatewayFedexQuery {
	if sgfq == nil {
		return nil
	}
	return &ShipmentGatewayFedexQuery{
		config:                    sgfq.config,
		limit:                     sgfq.limit,
		offset:                    sgfq.offset,
		order:                     append([]OrderFunc{}, sgfq.order...),
		predicates:                append([]predicate.ShipmentGatewayFedex{}, sgfq.predicates...),
		withShipmentGatewayConfig: sgfq.withShipmentGatewayConfig.Clone(),
		// clone intermediate query.
		sql:  sgfq.sql.Clone(),
		path: sgfq.path,
	}
}

// WithShipmentGatewayConfig tells the query-builder to eager-load the nodes that are connected to
// the "shipment_gateway_config" edge. The optional arguments are used to configure the query builder of the edge.
func (sgfq *ShipmentGatewayFedexQuery) WithShipmentGatewayConfig(opts ...func(*ShipmentGatewayConfigQuery)) *ShipmentGatewayFedexQuery {
	query := &ShipmentGatewayConfigQuery{config: sgfq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgfq.withShipmentGatewayConfig = query
	return sgfq
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
//	client.ShipmentGatewayFedex.Query().
//		GroupBy(shipmentgatewayfedex.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sgfq *ShipmentGatewayFedexQuery) GroupBy(field string, fields ...string) *ShipmentGatewayFedexGroupBy {
	group := &ShipmentGatewayFedexGroupBy{config: sgfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sgfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sgfq.sqlQuery(ctx), nil
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
//	client.ShipmentGatewayFedex.Query().
//		Select(shipmentgatewayfedex.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sgfq *ShipmentGatewayFedexQuery) Select(field string, fields ...string) *ShipmentGatewayFedexSelect {
	sgfq.fields = append([]string{field}, fields...)
	return &ShipmentGatewayFedexSelect{ShipmentGatewayFedexQuery: sgfq}
}

func (sgfq *ShipmentGatewayFedexQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sgfq.fields {
		if !shipmentgatewayfedex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sgfq.path != nil {
		prev, err := sgfq.path(ctx)
		if err != nil {
			return err
		}
		sgfq.sql = prev
	}
	return nil
}

func (sgfq *ShipmentGatewayFedexQuery) sqlAll(ctx context.Context) ([]*ShipmentGatewayFedex, error) {
	var (
		nodes       = []*ShipmentGatewayFedex{}
		withFKs     = sgfq.withFKs
		_spec       = sgfq.querySpec()
		loadedTypes = [1]bool{
			sgfq.withShipmentGatewayConfig != nil,
		}
	)
	if sgfq.withShipmentGatewayConfig != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentgatewayfedex.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShipmentGatewayFedex{config: sgfq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sgfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sgfq.withShipmentGatewayConfig; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ShipmentGatewayFedex)
		for i := range nodes {
			if nodes[i].shipment_gateway_config_shipment_gateway_fedex == nil {
				continue
			}
			fk := *nodes[i].shipment_gateway_config_shipment_gateway_fedex
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(shipmentgatewayconfig.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_gateway_config_shipment_gateway_fedex" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ShipmentGatewayConfig = n
			}
		}
	}

	return nodes, nil
}

func (sgfq *ShipmentGatewayFedexQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sgfq.querySpec()
	return sqlgraph.CountNodes(ctx, sgfq.driver, _spec)
}

func (sgfq *ShipmentGatewayFedexQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sgfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sgfq *ShipmentGatewayFedexQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shipmentgatewayfedex.Table,
			Columns: shipmentgatewayfedex.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewayfedex.FieldID,
			},
		},
		From:   sgfq.sql,
		Unique: true,
	}
	if unique := sgfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sgfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentgatewayfedex.FieldID)
		for i := range fields {
			if fields[i] != shipmentgatewayfedex.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sgfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sgfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sgfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sgfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sgfq *ShipmentGatewayFedexQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sgfq.driver.Dialect())
	t1 := builder.Table(shipmentgatewayfedex.Table)
	columns := sgfq.fields
	if len(columns) == 0 {
		columns = shipmentgatewayfedex.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sgfq.sql != nil {
		selector = sgfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sgfq.predicates {
		p(selector)
	}
	for _, p := range sgfq.order {
		p(selector)
	}
	if offset := sgfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sgfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentGatewayFedexGroupBy is the group-by builder for ShipmentGatewayFedex entities.
type ShipmentGatewayFedexGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgfgb *ShipmentGatewayFedexGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentGatewayFedexGroupBy {
	sgfgb.fns = append(sgfgb.fns, fns...)
	return sgfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgfgb *ShipmentGatewayFedexGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgfgb.path(ctx)
	if err != nil {
		return err
	}
	sgfgb.sql = query
	return sgfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgfgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) StringX(ctx context.Context) string {
	v, err := sgfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgfgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) IntX(ctx context.Context) int {
	v, err := sgfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgfgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgfgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgfgb *ShipmentGatewayFedexGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgfgb *ShipmentGatewayFedexGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgfgb *ShipmentGatewayFedexGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgfgb.fields {
		if !shipmentgatewayfedex.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgfgb *ShipmentGatewayFedexGroupBy) sqlQuery() *sql.Selector {
	selector := sgfgb.sql.Select()
	aggregation := make([]string, 0, len(sgfgb.fns))
	for _, fn := range sgfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgfgb.fields)+len(sgfgb.fns))
		for _, f := range sgfgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgfgb.fields...)...)
}

// ShipmentGatewayFedexSelect is the builder for selecting fields of ShipmentGatewayFedex entities.
type ShipmentGatewayFedexSelect struct {
	*ShipmentGatewayFedexQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sgfs *ShipmentGatewayFedexSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sgfs.prepareQuery(ctx); err != nil {
		return err
	}
	sgfs.sql = sgfs.ShipmentGatewayFedexQuery.sqlQuery(ctx)
	return sgfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sgfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sgfs.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sgfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) StringsX(ctx context.Context) []string {
	v, err := sgfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) StringX(ctx context.Context) string {
	v, err := sgfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sgfs.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sgfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) IntsX(ctx context.Context) []int {
	v, err := sgfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) IntX(ctx context.Context) int {
	v, err := sgfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgfs.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sgfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sgfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) Float64X(ctx context.Context) float64 {
	v, err := sgfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sgfs.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayFedexSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sgfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) BoolsX(ctx context.Context) []bool {
	v, err := sgfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sgfs *ShipmentGatewayFedexSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewayfedex.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayFedexSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgfs *ShipmentGatewayFedexSelect) BoolX(ctx context.Context) bool {
	v, err := sgfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgfs *ShipmentGatewayFedexSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgfs.sql.Query()
	if err := sgfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
