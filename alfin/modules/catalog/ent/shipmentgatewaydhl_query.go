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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewaydhl"
)

// ShipmentGatewayDhlQuery is the builder for querying ShipmentGatewayDhl entities.
type ShipmentGatewayDhlQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShipmentGatewayDhl
	// eager-loading edges.
	withShipmentGatewayConfig *ShipmentGatewayConfigQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentGatewayDhlQuery builder.
func (sgdq *ShipmentGatewayDhlQuery) Where(ps ...predicate.ShipmentGatewayDhl) *ShipmentGatewayDhlQuery {
	sgdq.predicates = append(sgdq.predicates, ps...)
	return sgdq
}

// Limit adds a limit step to the query.
func (sgdq *ShipmentGatewayDhlQuery) Limit(limit int) *ShipmentGatewayDhlQuery {
	sgdq.limit = &limit
	return sgdq
}

// Offset adds an offset step to the query.
func (sgdq *ShipmentGatewayDhlQuery) Offset(offset int) *ShipmentGatewayDhlQuery {
	sgdq.offset = &offset
	return sgdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sgdq *ShipmentGatewayDhlQuery) Unique(unique bool) *ShipmentGatewayDhlQuery {
	sgdq.unique = &unique
	return sgdq
}

// Order adds an order step to the query.
func (sgdq *ShipmentGatewayDhlQuery) Order(o ...OrderFunc) *ShipmentGatewayDhlQuery {
	sgdq.order = append(sgdq.order, o...)
	return sgdq
}

// QueryShipmentGatewayConfig chains the current query on the "shipment_gateway_config" edge.
func (sgdq *ShipmentGatewayDhlQuery) QueryShipmentGatewayConfig() *ShipmentGatewayConfigQuery {
	query := &ShipmentGatewayConfigQuery{config: sgdq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sgdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sgdq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentgatewaydhl.Table, shipmentgatewaydhl.FieldID, selector),
			sqlgraph.To(shipmentgatewayconfig.Table, shipmentgatewayconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, shipmentgatewaydhl.ShipmentGatewayConfigTable, shipmentgatewaydhl.ShipmentGatewayConfigColumn),
		)
		fromU = sqlgraph.SetNeighbors(sgdq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentGatewayDhl entity from the query.
// Returns a *NotFoundError when no ShipmentGatewayDhl was found.
func (sgdq *ShipmentGatewayDhlQuery) First(ctx context.Context) (*ShipmentGatewayDhl, error) {
	nodes, err := sgdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentgatewaydhl.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) FirstX(ctx context.Context) *ShipmentGatewayDhl {
	node, err := sgdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentGatewayDhl ID from the query.
// Returns a *NotFoundError when no ShipmentGatewayDhl ID was found.
func (sgdq *ShipmentGatewayDhlQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentgatewaydhl.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) FirstIDX(ctx context.Context) int {
	id, err := sgdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentGatewayDhl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShipmentGatewayDhl entity is not found.
// Returns a *NotFoundError when no ShipmentGatewayDhl entities are found.
func (sgdq *ShipmentGatewayDhlQuery) Only(ctx context.Context) (*ShipmentGatewayDhl, error) {
	nodes, err := sgdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentgatewaydhl.Label}
	default:
		return nil, &NotSingularError{shipmentgatewaydhl.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) OnlyX(ctx context.Context) *ShipmentGatewayDhl {
	node, err := sgdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentGatewayDhl ID in the query.
// Returns a *NotSingularError when exactly one ShipmentGatewayDhl ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sgdq *ShipmentGatewayDhlQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sgdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = &NotSingularError{shipmentgatewaydhl.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) OnlyIDX(ctx context.Context) int {
	id, err := sgdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentGatewayDhls.
func (sgdq *ShipmentGatewayDhlQuery) All(ctx context.Context) ([]*ShipmentGatewayDhl, error) {
	if err := sgdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sgdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) AllX(ctx context.Context) []*ShipmentGatewayDhl {
	nodes, err := sgdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentGatewayDhl IDs.
func (sgdq *ShipmentGatewayDhlQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sgdq.Select(shipmentgatewaydhl.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) IDsX(ctx context.Context) []int {
	ids, err := sgdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sgdq *ShipmentGatewayDhlQuery) Count(ctx context.Context) (int, error) {
	if err := sgdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sgdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) CountX(ctx context.Context) int {
	count, err := sgdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sgdq *ShipmentGatewayDhlQuery) Exist(ctx context.Context) (bool, error) {
	if err := sgdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sgdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sgdq *ShipmentGatewayDhlQuery) ExistX(ctx context.Context) bool {
	exist, err := sgdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentGatewayDhlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sgdq *ShipmentGatewayDhlQuery) Clone() *ShipmentGatewayDhlQuery {
	if sgdq == nil {
		return nil
	}
	return &ShipmentGatewayDhlQuery{
		config:                    sgdq.config,
		limit:                     sgdq.limit,
		offset:                    sgdq.offset,
		order:                     append([]OrderFunc{}, sgdq.order...),
		predicates:                append([]predicate.ShipmentGatewayDhl{}, sgdq.predicates...),
		withShipmentGatewayConfig: sgdq.withShipmentGatewayConfig.Clone(),
		// clone intermediate query.
		sql:  sgdq.sql.Clone(),
		path: sgdq.path,
	}
}

// WithShipmentGatewayConfig tells the query-builder to eager-load the nodes that are connected to
// the "shipment_gateway_config" edge. The optional arguments are used to configure the query builder of the edge.
func (sgdq *ShipmentGatewayDhlQuery) WithShipmentGatewayConfig(opts ...func(*ShipmentGatewayConfigQuery)) *ShipmentGatewayDhlQuery {
	query := &ShipmentGatewayConfigQuery{config: sgdq.config}
	for _, opt := range opts {
		opt(query)
	}
	sgdq.withShipmentGatewayConfig = query
	return sgdq
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
//	client.ShipmentGatewayDhl.Query().
//		GroupBy(shipmentgatewaydhl.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sgdq *ShipmentGatewayDhlQuery) GroupBy(field string, fields ...string) *ShipmentGatewayDhlGroupBy {
	group := &ShipmentGatewayDhlGroupBy{config: sgdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sgdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sgdq.sqlQuery(ctx), nil
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
//	client.ShipmentGatewayDhl.Query().
//		Select(shipmentgatewaydhl.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sgdq *ShipmentGatewayDhlQuery) Select(field string, fields ...string) *ShipmentGatewayDhlSelect {
	sgdq.fields = append([]string{field}, fields...)
	return &ShipmentGatewayDhlSelect{ShipmentGatewayDhlQuery: sgdq}
}

func (sgdq *ShipmentGatewayDhlQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sgdq.fields {
		if !shipmentgatewaydhl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sgdq.path != nil {
		prev, err := sgdq.path(ctx)
		if err != nil {
			return err
		}
		sgdq.sql = prev
	}
	return nil
}

func (sgdq *ShipmentGatewayDhlQuery) sqlAll(ctx context.Context) ([]*ShipmentGatewayDhl, error) {
	var (
		nodes       = []*ShipmentGatewayDhl{}
		withFKs     = sgdq.withFKs
		_spec       = sgdq.querySpec()
		loadedTypes = [1]bool{
			sgdq.withShipmentGatewayConfig != nil,
		}
	)
	if sgdq.withShipmentGatewayConfig != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentgatewaydhl.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShipmentGatewayDhl{config: sgdq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sgdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sgdq.withShipmentGatewayConfig; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ShipmentGatewayDhl)
		for i := range nodes {
			if nodes[i].shipment_gateway_config_shipment_gateway_dhl == nil {
				continue
			}
			fk := *nodes[i].shipment_gateway_config_shipment_gateway_dhl
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
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_gateway_config_shipment_gateway_dhl" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ShipmentGatewayConfig = n
			}
		}
	}

	return nodes, nil
}

func (sgdq *ShipmentGatewayDhlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sgdq.querySpec()
	return sqlgraph.CountNodes(ctx, sgdq.driver, _spec)
}

func (sgdq *ShipmentGatewayDhlQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sgdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sgdq *ShipmentGatewayDhlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shipmentgatewaydhl.Table,
			Columns: shipmentgatewaydhl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmentgatewaydhl.FieldID,
			},
		},
		From:   sgdq.sql,
		Unique: true,
	}
	if unique := sgdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sgdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentgatewaydhl.FieldID)
		for i := range fields {
			if fields[i] != shipmentgatewaydhl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sgdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sgdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sgdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sgdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sgdq *ShipmentGatewayDhlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sgdq.driver.Dialect())
	t1 := builder.Table(shipmentgatewaydhl.Table)
	columns := sgdq.fields
	if len(columns) == 0 {
		columns = shipmentgatewaydhl.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sgdq.sql != nil {
		selector = sgdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range sgdq.predicates {
		p(selector)
	}
	for _, p := range sgdq.order {
		p(selector)
	}
	if offset := sgdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sgdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentGatewayDhlGroupBy is the group-by builder for ShipmentGatewayDhl entities.
type ShipmentGatewayDhlGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgdgb *ShipmentGatewayDhlGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentGatewayDhlGroupBy {
	sgdgb.fns = append(sgdgb.fns, fns...)
	return sgdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgdgb *ShipmentGatewayDhlGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgdgb.path(ctx)
	if err != nil {
		return err
	}
	sgdgb.sql = query
	return sgdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgdgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) StringX(ctx context.Context) string {
	v, err := sgdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgdgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) IntX(ctx context.Context) int {
	v, err := sgdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgdgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgdgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sgdgb *ShipmentGatewayDhlGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgdgb *ShipmentGatewayDhlGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgdgb *ShipmentGatewayDhlGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgdgb.fields {
		if !shipmentgatewaydhl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgdgb *ShipmentGatewayDhlGroupBy) sqlQuery() *sql.Selector {
	selector := sgdgb.sql.Select()
	aggregation := make([]string, 0, len(sgdgb.fns))
	for _, fn := range sgdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgdgb.fields)+len(sgdgb.fns))
		for _, f := range sgdgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgdgb.fields...)...)
}

// ShipmentGatewayDhlSelect is the builder for selecting fields of ShipmentGatewayDhl entities.
type ShipmentGatewayDhlSelect struct {
	*ShipmentGatewayDhlQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sgds *ShipmentGatewayDhlSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sgds.prepareQuery(ctx); err != nil {
		return err
	}
	sgds.sql = sgds.ShipmentGatewayDhlQuery.sqlQuery(ctx)
	return sgds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sgds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sgds.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sgds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) StringsX(ctx context.Context) []string {
	v, err := sgds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) StringX(ctx context.Context) string {
	v, err := sgds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sgds.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sgds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) IntsX(ctx context.Context) []int {
	v, err := sgds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) IntX(ctx context.Context) int {
	v, err := sgds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgds.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sgds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sgds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) Float64X(ctx context.Context) float64 {
	v, err := sgds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sgds.fields) > 1 {
		return nil, errors.New("ent: ShipmentGatewayDhlSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sgds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) BoolsX(ctx context.Context) []bool {
	v, err := sgds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sgds *ShipmentGatewayDhlSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmentgatewaydhl.Label}
	default:
		err = fmt.Errorf("ent: ShipmentGatewayDhlSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgds *ShipmentGatewayDhlSelect) BoolX(ctx context.Context) bool {
	v, err := sgds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgds *ShipmentGatewayDhlSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgds.sql.Query()
	if err := sgds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}