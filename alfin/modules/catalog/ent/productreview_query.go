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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/product"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productreview"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstore"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
)

// ProductReviewQuery is the builder for querying ProductReview entities.
type ProductReviewQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductReview
	// eager-loading edges.
	withProductStore *ProductStoreQuery
	withProduct      *ProductQuery
	withStatusItem   *StatusItemQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductReviewQuery builder.
func (prq *ProductReviewQuery) Where(ps ...predicate.ProductReview) *ProductReviewQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit adds a limit step to the query.
func (prq *ProductReviewQuery) Limit(limit int) *ProductReviewQuery {
	prq.limit = &limit
	return prq
}

// Offset adds an offset step to the query.
func (prq *ProductReviewQuery) Offset(offset int) *ProductReviewQuery {
	prq.offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *ProductReviewQuery) Unique(unique bool) *ProductReviewQuery {
	prq.unique = &unique
	return prq
}

// Order adds an order step to the query.
func (prq *ProductReviewQuery) Order(o ...OrderFunc) *ProductReviewQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// QueryProductStore chains the current query on the "product_store" edge.
func (prq *ProductReviewQuery) QueryProductStore() *ProductStoreQuery {
	query := &ProductStoreQuery{config: prq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productreview.Table, productreview.FieldID, selector),
			sqlgraph.To(productstore.Table, productstore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productreview.ProductStoreTable, productreview.ProductStoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (prq *ProductReviewQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: prq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productreview.Table, productreview.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productreview.ProductTable, productreview.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatusItem chains the current query on the "status_item" edge.
func (prq *ProductReviewQuery) QueryStatusItem() *StatusItemQuery {
	query := &StatusItemQuery{config: prq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := prq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productreview.Table, productreview.FieldID, selector),
			sqlgraph.To(statusitem.Table, statusitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productreview.StatusItemTable, productreview.StatusItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(prq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductReview entity from the query.
// Returns a *NotFoundError when no ProductReview was found.
func (prq *ProductReviewQuery) First(ctx context.Context) (*ProductReview, error) {
	nodes, err := prq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productreview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *ProductReviewQuery) FirstX(ctx context.Context) *ProductReview {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductReview ID from the query.
// Returns a *NotFoundError when no ProductReview ID was found.
func (prq *ProductReviewQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productreview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *ProductReviewQuery) FirstIDX(ctx context.Context) int {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductReview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductReview entity is not found.
// Returns a *NotFoundError when no ProductReview entities are found.
func (prq *ProductReviewQuery) Only(ctx context.Context) (*ProductReview, error) {
	nodes, err := prq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productreview.Label}
	default:
		return nil, &NotSingularError{productreview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *ProductReviewQuery) OnlyX(ctx context.Context) *ProductReview {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductReview ID in the query.
// Returns a *NotSingularError when exactly one ProductReview ID is not found.
// Returns a *NotFoundError when no entities are found.
func (prq *ProductReviewQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = prq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = &NotSingularError{productreview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *ProductReviewQuery) OnlyIDX(ctx context.Context) int {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductReviews.
func (prq *ProductReviewQuery) All(ctx context.Context) ([]*ProductReview, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return prq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (prq *ProductReviewQuery) AllX(ctx context.Context) []*ProductReview {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductReview IDs.
func (prq *ProductReviewQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := prq.Select(productreview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *ProductReviewQuery) IDsX(ctx context.Context) []int {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *ProductReviewQuery) Count(ctx context.Context) (int, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return prq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (prq *ProductReviewQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *ProductReviewQuery) Exist(ctx context.Context) (bool, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return prq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *ProductReviewQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductReviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *ProductReviewQuery) Clone() *ProductReviewQuery {
	if prq == nil {
		return nil
	}
	return &ProductReviewQuery{
		config:           prq.config,
		limit:            prq.limit,
		offset:           prq.offset,
		order:            append([]OrderFunc{}, prq.order...),
		predicates:       append([]predicate.ProductReview{}, prq.predicates...),
		withProductStore: prq.withProductStore.Clone(),
		withProduct:      prq.withProduct.Clone(),
		withStatusItem:   prq.withStatusItem.Clone(),
		// clone intermediate query.
		sql:  prq.sql.Clone(),
		path: prq.path,
	}
}

// WithProductStore tells the query-builder to eager-load the nodes that are connected to
// the "product_store" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProductReviewQuery) WithProductStore(opts ...func(*ProductStoreQuery)) *ProductReviewQuery {
	query := &ProductStoreQuery{config: prq.config}
	for _, opt := range opts {
		opt(query)
	}
	prq.withProductStore = query
	return prq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProductReviewQuery) WithProduct(opts ...func(*ProductQuery)) *ProductReviewQuery {
	query := &ProductQuery{config: prq.config}
	for _, opt := range opts {
		opt(query)
	}
	prq.withProduct = query
	return prq
}

// WithStatusItem tells the query-builder to eager-load the nodes that are connected to
// the "status_item" edge. The optional arguments are used to configure the query builder of the edge.
func (prq *ProductReviewQuery) WithStatusItem(opts ...func(*StatusItemQuery)) *ProductReviewQuery {
	query := &StatusItemQuery{config: prq.config}
	for _, opt := range opts {
		opt(query)
	}
	prq.withStatusItem = query
	return prq
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
//	client.ProductReview.Query().
//		GroupBy(productreview.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (prq *ProductReviewQuery) GroupBy(field string, fields ...string) *ProductReviewGroupBy {
	group := &ProductReviewGroupBy{config: prq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return prq.sqlQuery(ctx), nil
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
//	client.ProductReview.Query().
//		Select(productreview.FieldCreateTime).
//		Scan(ctx, &v)
//
func (prq *ProductReviewQuery) Select(field string, fields ...string) *ProductReviewSelect {
	prq.fields = append([]string{field}, fields...)
	return &ProductReviewSelect{ProductReviewQuery: prq}
}

func (prq *ProductReviewQuery) prepareQuery(ctx context.Context) error {
	for _, f := range prq.fields {
		if !productreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	return nil
}

func (prq *ProductReviewQuery) sqlAll(ctx context.Context) ([]*ProductReview, error) {
	var (
		nodes       = []*ProductReview{}
		withFKs     = prq.withFKs
		_spec       = prq.querySpec()
		loadedTypes = [3]bool{
			prq.withProductStore != nil,
			prq.withProduct != nil,
			prq.withStatusItem != nil,
		}
	)
	if prq.withProductStore != nil || prq.withProduct != nil || prq.withStatusItem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, productreview.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductReview{config: prq.config}
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
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := prq.withProductStore; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductReview)
		for i := range nodes {
			if nodes[i].product_store_product_reviews == nil {
				continue
			}
			fk := *nodes[i].product_store_product_reviews
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(productstore.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_store_product_reviews" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ProductStore = n
			}
		}
	}

	if query := prq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductReview)
		for i := range nodes {
			if nodes[i].product_product_reviews == nil {
				continue
			}
			fk := *nodes[i].product_product_reviews
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_product_reviews" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	if query := prq.withStatusItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductReview)
		for i := range nodes {
			if nodes[i].status_item_product_reviews == nil {
				continue
			}
			fk := *nodes[i].status_item_product_reviews
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
				return nil, fmt.Errorf(`unexpected foreign-key "status_item_product_reviews" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.StatusItem = n
			}
		}
	}

	return nodes, nil
}

func (prq *ProductReviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *ProductReviewQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := prq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (prq *ProductReviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productreview.Table,
			Columns: productreview.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productreview.FieldID,
			},
		},
		From:   prq.sql,
		Unique: true,
	}
	if unique := prq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := prq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productreview.FieldID)
		for i := range fields {
			if fields[i] != productreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *ProductReviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(productreview.Table)
	columns := prq.fields
	if len(columns) == 0 {
		columns = productreview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductReviewGroupBy is the group-by builder for ProductReview entities.
type ProductReviewGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *ProductReviewGroupBy) Aggregate(fns ...AggregateFunc) *ProductReviewGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the group-by query and scans the result into the given value.
func (prgb *ProductReviewGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := prgb.path(ctx)
	if err != nil {
		return err
	}
	prgb.sql = query
	return prgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := prgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: ProductReviewGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) StringsX(ctx context.Context) []string {
	v, err := prgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = prgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) StringX(ctx context.Context) string {
	v, err := prgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: ProductReviewGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) IntsX(ctx context.Context) []int {
	v, err := prgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = prgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) IntX(ctx context.Context) int {
	v, err := prgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: ProductReviewGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := prgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = prgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) Float64X(ctx context.Context) float64 {
	v, err := prgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(prgb.fields) > 1 {
		return nil, errors.New("ent: ProductReviewGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := prgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := prgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (prgb *ProductReviewGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = prgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (prgb *ProductReviewGroupBy) BoolX(ctx context.Context) bool {
	v, err := prgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (prgb *ProductReviewGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range prgb.fields {
		if !productreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := prgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (prgb *ProductReviewGroupBy) sqlQuery() *sql.Selector {
	selector := prgb.sql.Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(prgb.fields)+len(prgb.fns))
		for _, f := range prgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(prgb.fields...)...)
}

// ProductReviewSelect is the builder for selecting fields of ProductReview entities.
type ProductReviewSelect struct {
	*ProductReviewQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (prs *ProductReviewSelect) Scan(ctx context.Context, v interface{}) error {
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	prs.sql = prs.ProductReviewQuery.sqlQuery(ctx)
	return prs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (prs *ProductReviewSelect) ScanX(ctx context.Context, v interface{}) {
	if err := prs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Strings(ctx context.Context) ([]string, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: ProductReviewSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (prs *ProductReviewSelect) StringsX(ctx context.Context) []string {
	v, err := prs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = prs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (prs *ProductReviewSelect) StringX(ctx context.Context) string {
	v, err := prs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Ints(ctx context.Context) ([]int, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: ProductReviewSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (prs *ProductReviewSelect) IntsX(ctx context.Context) []int {
	v, err := prs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = prs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (prs *ProductReviewSelect) IntX(ctx context.Context) int {
	v, err := prs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: ProductReviewSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (prs *ProductReviewSelect) Float64sX(ctx context.Context) []float64 {
	v, err := prs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = prs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (prs *ProductReviewSelect) Float64X(ctx context.Context) float64 {
	v, err := prs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(prs.fields) > 1 {
		return nil, errors.New("ent: ProductReviewSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := prs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (prs *ProductReviewSelect) BoolsX(ctx context.Context) []bool {
	v, err := prs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (prs *ProductReviewSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = prs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productreview.Label}
	default:
		err = fmt.Errorf("ent: ProductReviewSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (prs *ProductReviewSelect) BoolX(ctx context.Context) bool {
	v, err := prs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (prs *ProductReviewSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := prs.sql.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
