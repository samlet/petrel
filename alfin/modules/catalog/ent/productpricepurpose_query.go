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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productprice"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpricepurpose"
)

// ProductPricePurposeQuery is the builder for querying ProductPricePurpose entities.
type ProductPricePurposeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductPricePurpose
	// eager-loading edges.
	withOrderPaymentPreferences *OrderPaymentPreferenceQuery
	withProductPrices           *ProductPriceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductPricePurposeQuery builder.
func (pppq *ProductPricePurposeQuery) Where(ps ...predicate.ProductPricePurpose) *ProductPricePurposeQuery {
	pppq.predicates = append(pppq.predicates, ps...)
	return pppq
}

// Limit adds a limit step to the query.
func (pppq *ProductPricePurposeQuery) Limit(limit int) *ProductPricePurposeQuery {
	pppq.limit = &limit
	return pppq
}

// Offset adds an offset step to the query.
func (pppq *ProductPricePurposeQuery) Offset(offset int) *ProductPricePurposeQuery {
	pppq.offset = &offset
	return pppq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pppq *ProductPricePurposeQuery) Unique(unique bool) *ProductPricePurposeQuery {
	pppq.unique = &unique
	return pppq
}

// Order adds an order step to the query.
func (pppq *ProductPricePurposeQuery) Order(o ...OrderFunc) *ProductPricePurposeQuery {
	pppq.order = append(pppq.order, o...)
	return pppq
}

// QueryOrderPaymentPreferences chains the current query on the "order_payment_preferences" edge.
func (pppq *ProductPricePurposeQuery) QueryOrderPaymentPreferences() *OrderPaymentPreferenceQuery {
	query := &OrderPaymentPreferenceQuery{config: pppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productpricepurpose.Table, productpricepurpose.FieldID, selector),
			sqlgraph.To(orderpaymentpreference.Table, orderpaymentpreference.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productpricepurpose.OrderPaymentPreferencesTable, productpricepurpose.OrderPaymentPreferencesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductPrices chains the current query on the "product_prices" edge.
func (pppq *ProductPricePurposeQuery) QueryProductPrices() *ProductPriceQuery {
	query := &ProductPriceQuery{config: pppq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pppq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productpricepurpose.Table, productpricepurpose.FieldID, selector),
			sqlgraph.To(productprice.Table, productprice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productpricepurpose.ProductPricesTable, productpricepurpose.ProductPricesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pppq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductPricePurpose entity from the query.
// Returns a *NotFoundError when no ProductPricePurpose was found.
func (pppq *ProductPricePurposeQuery) First(ctx context.Context) (*ProductPricePurpose, error) {
	nodes, err := pppq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productpricepurpose.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) FirstX(ctx context.Context) *ProductPricePurpose {
	node, err := pppq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductPricePurpose ID from the query.
// Returns a *NotFoundError when no ProductPricePurpose ID was found.
func (pppq *ProductPricePurposeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pppq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productpricepurpose.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) FirstIDX(ctx context.Context) int {
	id, err := pppq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductPricePurpose entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductPricePurpose entity is not found.
// Returns a *NotFoundError when no ProductPricePurpose entities are found.
func (pppq *ProductPricePurposeQuery) Only(ctx context.Context) (*ProductPricePurpose, error) {
	nodes, err := pppq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productpricepurpose.Label}
	default:
		return nil, &NotSingularError{productpricepurpose.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) OnlyX(ctx context.Context) *ProductPricePurpose {
	node, err := pppq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductPricePurpose ID in the query.
// Returns a *NotSingularError when exactly one ProductPricePurpose ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pppq *ProductPricePurposeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pppq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = &NotSingularError{productpricepurpose.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pppq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductPricePurposes.
func (pppq *ProductPricePurposeQuery) All(ctx context.Context) ([]*ProductPricePurpose, error) {
	if err := pppq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pppq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) AllX(ctx context.Context) []*ProductPricePurpose {
	nodes, err := pppq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductPricePurpose IDs.
func (pppq *ProductPricePurposeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pppq.Select(productpricepurpose.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) IDsX(ctx context.Context) []int {
	ids, err := pppq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pppq *ProductPricePurposeQuery) Count(ctx context.Context) (int, error) {
	if err := pppq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pppq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) CountX(ctx context.Context) int {
	count, err := pppq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pppq *ProductPricePurposeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pppq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pppq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pppq *ProductPricePurposeQuery) ExistX(ctx context.Context) bool {
	exist, err := pppq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductPricePurposeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pppq *ProductPricePurposeQuery) Clone() *ProductPricePurposeQuery {
	if pppq == nil {
		return nil
	}
	return &ProductPricePurposeQuery{
		config:                      pppq.config,
		limit:                       pppq.limit,
		offset:                      pppq.offset,
		order:                       append([]OrderFunc{}, pppq.order...),
		predicates:                  append([]predicate.ProductPricePurpose{}, pppq.predicates...),
		withOrderPaymentPreferences: pppq.withOrderPaymentPreferences.Clone(),
		withProductPrices:           pppq.withProductPrices.Clone(),
		// clone intermediate query.
		sql:  pppq.sql.Clone(),
		path: pppq.path,
	}
}

// WithOrderPaymentPreferences tells the query-builder to eager-load the nodes that are connected to
// the "order_payment_preferences" edge. The optional arguments are used to configure the query builder of the edge.
func (pppq *ProductPricePurposeQuery) WithOrderPaymentPreferences(opts ...func(*OrderPaymentPreferenceQuery)) *ProductPricePurposeQuery {
	query := &OrderPaymentPreferenceQuery{config: pppq.config}
	for _, opt := range opts {
		opt(query)
	}
	pppq.withOrderPaymentPreferences = query
	return pppq
}

// WithProductPrices tells the query-builder to eager-load the nodes that are connected to
// the "product_prices" edge. The optional arguments are used to configure the query builder of the edge.
func (pppq *ProductPricePurposeQuery) WithProductPrices(opts ...func(*ProductPriceQuery)) *ProductPricePurposeQuery {
	query := &ProductPriceQuery{config: pppq.config}
	for _, opt := range opts {
		opt(query)
	}
	pppq.withProductPrices = query
	return pppq
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
//	client.ProductPricePurpose.Query().
//		GroupBy(productpricepurpose.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pppq *ProductPricePurposeQuery) GroupBy(field string, fields ...string) *ProductPricePurposeGroupBy {
	group := &ProductPricePurposeGroupBy{config: pppq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pppq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pppq.sqlQuery(ctx), nil
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
//	client.ProductPricePurpose.Query().
//		Select(productpricepurpose.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pppq *ProductPricePurposeQuery) Select(field string, fields ...string) *ProductPricePurposeSelect {
	pppq.fields = append([]string{field}, fields...)
	return &ProductPricePurposeSelect{ProductPricePurposeQuery: pppq}
}

func (pppq *ProductPricePurposeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pppq.fields {
		if !productpricepurpose.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pppq.path != nil {
		prev, err := pppq.path(ctx)
		if err != nil {
			return err
		}
		pppq.sql = prev
	}
	return nil
}

func (pppq *ProductPricePurposeQuery) sqlAll(ctx context.Context) ([]*ProductPricePurpose, error) {
	var (
		nodes       = []*ProductPricePurpose{}
		_spec       = pppq.querySpec()
		loadedTypes = [2]bool{
			pppq.withOrderPaymentPreferences != nil,
			pppq.withProductPrices != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductPricePurpose{config: pppq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pppq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pppq.withOrderPaymentPreferences; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ProductPricePurpose)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.OrderPaymentPreferences = []*OrderPaymentPreference{}
		}
		query.withFKs = true
		query.Where(predicate.OrderPaymentPreference(func(s *sql.Selector) {
			s.Where(sql.InValues(productpricepurpose.OrderPaymentPreferencesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.product_price_purpose_order_payment_preferences
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "product_price_purpose_order_payment_preferences" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_price_purpose_order_payment_preferences" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.OrderPaymentPreferences = append(node.Edges.OrderPaymentPreferences, n)
		}
	}

	if query := pppq.withProductPrices; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ProductPricePurpose)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ProductPrices = []*ProductPrice{}
		}
		query.withFKs = true
		query.Where(predicate.ProductPrice(func(s *sql.Selector) {
			s.Where(sql.InValues(productpricepurpose.ProductPricesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.product_price_purpose_product_prices
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "product_price_purpose_product_prices" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_price_purpose_product_prices" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ProductPrices = append(node.Edges.ProductPrices, n)
		}
	}

	return nodes, nil
}

func (pppq *ProductPricePurposeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pppq.querySpec()
	return sqlgraph.CountNodes(ctx, pppq.driver, _spec)
}

func (pppq *ProductPricePurposeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pppq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pppq *ProductPricePurposeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productpricepurpose.Table,
			Columns: productpricepurpose.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productpricepurpose.FieldID,
			},
		},
		From:   pppq.sql,
		Unique: true,
	}
	if unique := pppq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pppq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productpricepurpose.FieldID)
		for i := range fields {
			if fields[i] != productpricepurpose.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pppq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pppq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pppq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pppq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pppq *ProductPricePurposeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pppq.driver.Dialect())
	t1 := builder.Table(productpricepurpose.Table)
	columns := pppq.fields
	if len(columns) == 0 {
		columns = productpricepurpose.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pppq.sql != nil {
		selector = pppq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pppq.predicates {
		p(selector)
	}
	for _, p := range pppq.order {
		p(selector)
	}
	if offset := pppq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pppq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductPricePurposeGroupBy is the group-by builder for ProductPricePurpose entities.
type ProductPricePurposeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pppgb *ProductPricePurposeGroupBy) Aggregate(fns ...AggregateFunc) *ProductPricePurposeGroupBy {
	pppgb.fns = append(pppgb.fns, fns...)
	return pppgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pppgb *ProductPricePurposeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pppgb.path(ctx)
	if err != nil {
		return err
	}
	pppgb.sql = query
	return pppgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pppgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pppgb.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pppgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pppgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) StringX(ctx context.Context) string {
	v, err := pppgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pppgb.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pppgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pppgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) IntX(ctx context.Context) int {
	v, err := pppgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pppgb.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pppgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pppgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pppgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pppgb.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pppgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pppgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pppgb *ProductPricePurposeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pppgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pppgb *ProductPricePurposeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pppgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pppgb *ProductPricePurposeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pppgb.fields {
		if !productpricepurpose.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pppgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pppgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pppgb *ProductPricePurposeGroupBy) sqlQuery() *sql.Selector {
	selector := pppgb.sql.Select()
	aggregation := make([]string, 0, len(pppgb.fns))
	for _, fn := range pppgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pppgb.fields)+len(pppgb.fns))
		for _, f := range pppgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pppgb.fields...)...)
}

// ProductPricePurposeSelect is the builder for selecting fields of ProductPricePurpose entities.
type ProductPricePurposeSelect struct {
	*ProductPricePurposeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ppps *ProductPricePurposeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ppps.prepareQuery(ctx); err != nil {
		return err
	}
	ppps.sql = ppps.ProductPricePurposeQuery.sqlQuery(ctx)
	return ppps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ppps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ppps.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ppps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) StringsX(ctx context.Context) []string {
	v, err := ppps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ppps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) StringX(ctx context.Context) string {
	v, err := ppps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ppps.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ppps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) IntsX(ctx context.Context) []int {
	v, err := ppps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ppps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) IntX(ctx context.Context) int {
	v, err := ppps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ppps.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ppps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ppps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ppps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) Float64X(ctx context.Context) float64 {
	v, err := ppps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ppps.fields) > 1 {
		return nil, errors.New("ent: ProductPricePurposeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ppps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ppps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ppps *ProductPricePurposeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ppps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productpricepurpose.Label}
	default:
		err = fmt.Errorf("ent: ProductPricePurposeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ppps *ProductPricePurposeSelect) BoolX(ctx context.Context) bool {
	v, err := ppps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ppps *ProductPricePurposeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ppps.sql.Query()
	if err := ppps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
