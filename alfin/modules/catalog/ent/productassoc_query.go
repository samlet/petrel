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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/custommethod"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/product"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoctype"
)

// ProductAssocQuery is the builder for querying ProductAssoc entities.
type ProductAssocQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductAssoc
	// eager-loading edges.
	withProductAssocType *ProductAssocTypeQuery
	withMainProduct      *ProductQuery
	withAssocProduct     *ProductQuery
	withCustomMethod     *CustomMethodQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductAssocQuery builder.
func (paq *ProductAssocQuery) Where(ps ...predicate.ProductAssoc) *ProductAssocQuery {
	paq.predicates = append(paq.predicates, ps...)
	return paq
}

// Limit adds a limit step to the query.
func (paq *ProductAssocQuery) Limit(limit int) *ProductAssocQuery {
	paq.limit = &limit
	return paq
}

// Offset adds an offset step to the query.
func (paq *ProductAssocQuery) Offset(offset int) *ProductAssocQuery {
	paq.offset = &offset
	return paq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (paq *ProductAssocQuery) Unique(unique bool) *ProductAssocQuery {
	paq.unique = &unique
	return paq
}

// Order adds an order step to the query.
func (paq *ProductAssocQuery) Order(o ...OrderFunc) *ProductAssocQuery {
	paq.order = append(paq.order, o...)
	return paq
}

// QueryProductAssocType chains the current query on the "product_assoc_type" edge.
func (paq *ProductAssocQuery) QueryProductAssocType() *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: paq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoc.Table, productassoc.FieldID, selector),
			sqlgraph.To(productassoctype.Table, productassoctype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productassoc.ProductAssocTypeTable, productassoc.ProductAssocTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMainProduct chains the current query on the "main_product" edge.
func (paq *ProductAssocQuery) QueryMainProduct() *ProductQuery {
	query := &ProductQuery{config: paq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoc.Table, productassoc.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productassoc.MainProductTable, productassoc.MainProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssocProduct chains the current query on the "assoc_product" edge.
func (paq *ProductAssocQuery) QueryAssocProduct() *ProductQuery {
	query := &ProductQuery{config: paq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoc.Table, productassoc.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productassoc.AssocProductTable, productassoc.AssocProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomMethod chains the current query on the "custom_method" edge.
func (paq *ProductAssocQuery) QueryCustomMethod() *CustomMethodQuery {
	query := &CustomMethodQuery{config: paq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := paq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoc.Table, productassoc.FieldID, selector),
			sqlgraph.To(custommethod.Table, custommethod.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productassoc.CustomMethodTable, productassoc.CustomMethodColumn),
		)
		fromU = sqlgraph.SetNeighbors(paq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductAssoc entity from the query.
// Returns a *NotFoundError when no ProductAssoc was found.
func (paq *ProductAssocQuery) First(ctx context.Context) (*ProductAssoc, error) {
	nodes, err := paq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productassoc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (paq *ProductAssocQuery) FirstX(ctx context.Context) *ProductAssoc {
	node, err := paq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductAssoc ID from the query.
// Returns a *NotFoundError when no ProductAssoc ID was found.
func (paq *ProductAssocQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productassoc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (paq *ProductAssocQuery) FirstIDX(ctx context.Context) int {
	id, err := paq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductAssoc entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductAssoc entity is not found.
// Returns a *NotFoundError when no ProductAssoc entities are found.
func (paq *ProductAssocQuery) Only(ctx context.Context) (*ProductAssoc, error) {
	nodes, err := paq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productassoc.Label}
	default:
		return nil, &NotSingularError{productassoc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (paq *ProductAssocQuery) OnlyX(ctx context.Context) *ProductAssoc {
	node, err := paq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductAssoc ID in the query.
// Returns a *NotSingularError when exactly one ProductAssoc ID is not found.
// Returns a *NotFoundError when no entities are found.
func (paq *ProductAssocQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = paq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = &NotSingularError{productassoc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (paq *ProductAssocQuery) OnlyIDX(ctx context.Context) int {
	id, err := paq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductAssocs.
func (paq *ProductAssocQuery) All(ctx context.Context) ([]*ProductAssoc, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return paq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (paq *ProductAssocQuery) AllX(ctx context.Context) []*ProductAssoc {
	nodes, err := paq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductAssoc IDs.
func (paq *ProductAssocQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := paq.Select(productassoc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (paq *ProductAssocQuery) IDsX(ctx context.Context) []int {
	ids, err := paq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (paq *ProductAssocQuery) Count(ctx context.Context) (int, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return paq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (paq *ProductAssocQuery) CountX(ctx context.Context) int {
	count, err := paq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (paq *ProductAssocQuery) Exist(ctx context.Context) (bool, error) {
	if err := paq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return paq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (paq *ProductAssocQuery) ExistX(ctx context.Context) bool {
	exist, err := paq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductAssocQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (paq *ProductAssocQuery) Clone() *ProductAssocQuery {
	if paq == nil {
		return nil
	}
	return &ProductAssocQuery{
		config:               paq.config,
		limit:                paq.limit,
		offset:               paq.offset,
		order:                append([]OrderFunc{}, paq.order...),
		predicates:           append([]predicate.ProductAssoc{}, paq.predicates...),
		withProductAssocType: paq.withProductAssocType.Clone(),
		withMainProduct:      paq.withMainProduct.Clone(),
		withAssocProduct:     paq.withAssocProduct.Clone(),
		withCustomMethod:     paq.withCustomMethod.Clone(),
		// clone intermediate query.
		sql:  paq.sql.Clone(),
		path: paq.path,
	}
}

// WithProductAssocType tells the query-builder to eager-load the nodes that are connected to
// the "product_assoc_type" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProductAssocQuery) WithProductAssocType(opts ...func(*ProductAssocTypeQuery)) *ProductAssocQuery {
	query := &ProductAssocTypeQuery{config: paq.config}
	for _, opt := range opts {
		opt(query)
	}
	paq.withProductAssocType = query
	return paq
}

// WithMainProduct tells the query-builder to eager-load the nodes that are connected to
// the "main_product" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProductAssocQuery) WithMainProduct(opts ...func(*ProductQuery)) *ProductAssocQuery {
	query := &ProductQuery{config: paq.config}
	for _, opt := range opts {
		opt(query)
	}
	paq.withMainProduct = query
	return paq
}

// WithAssocProduct tells the query-builder to eager-load the nodes that are connected to
// the "assoc_product" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProductAssocQuery) WithAssocProduct(opts ...func(*ProductQuery)) *ProductAssocQuery {
	query := &ProductQuery{config: paq.config}
	for _, opt := range opts {
		opt(query)
	}
	paq.withAssocProduct = query
	return paq
}

// WithCustomMethod tells the query-builder to eager-load the nodes that are connected to
// the "custom_method" edge. The optional arguments are used to configure the query builder of the edge.
func (paq *ProductAssocQuery) WithCustomMethod(opts ...func(*CustomMethodQuery)) *ProductAssocQuery {
	query := &CustomMethodQuery{config: paq.config}
	for _, opt := range opts {
		opt(query)
	}
	paq.withCustomMethod = query
	return paq
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
//	client.ProductAssoc.Query().
//		GroupBy(productassoc.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (paq *ProductAssocQuery) GroupBy(field string, fields ...string) *ProductAssocGroupBy {
	group := &ProductAssocGroupBy{config: paq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := paq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return paq.sqlQuery(ctx), nil
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
//	client.ProductAssoc.Query().
//		Select(productassoc.FieldCreateTime).
//		Scan(ctx, &v)
//
func (paq *ProductAssocQuery) Select(field string, fields ...string) *ProductAssocSelect {
	paq.fields = append([]string{field}, fields...)
	return &ProductAssocSelect{ProductAssocQuery: paq}
}

func (paq *ProductAssocQuery) prepareQuery(ctx context.Context) error {
	for _, f := range paq.fields {
		if !productassoc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if paq.path != nil {
		prev, err := paq.path(ctx)
		if err != nil {
			return err
		}
		paq.sql = prev
	}
	return nil
}

func (paq *ProductAssocQuery) sqlAll(ctx context.Context) ([]*ProductAssoc, error) {
	var (
		nodes       = []*ProductAssoc{}
		withFKs     = paq.withFKs
		_spec       = paq.querySpec()
		loadedTypes = [4]bool{
			paq.withProductAssocType != nil,
			paq.withMainProduct != nil,
			paq.withAssocProduct != nil,
			paq.withCustomMethod != nil,
		}
	)
	if paq.withProductAssocType != nil || paq.withMainProduct != nil || paq.withAssocProduct != nil || paq.withCustomMethod != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, productassoc.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductAssoc{config: paq.config}
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
	if err := sqlgraph.QueryNodes(ctx, paq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := paq.withProductAssocType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductAssoc)
		for i := range nodes {
			if nodes[i].product_assoc_type_product_assocs == nil {
				continue
			}
			fk := *nodes[i].product_assoc_type_product_assocs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(productassoctype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_assoc_type_product_assocs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ProductAssocType = n
			}
		}
	}

	if query := paq.withMainProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductAssoc)
		for i := range nodes {
			if nodes[i].product_main_product_assocs == nil {
				continue
			}
			fk := *nodes[i].product_main_product_assocs
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
				return nil, fmt.Errorf(`unexpected foreign-key "product_main_product_assocs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.MainProduct = n
			}
		}
	}

	if query := paq.withAssocProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductAssoc)
		for i := range nodes {
			if nodes[i].product_assoc_product_assocs == nil {
				continue
			}
			fk := *nodes[i].product_assoc_product_assocs
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
				return nil, fmt.Errorf(`unexpected foreign-key "product_assoc_product_assocs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.AssocProduct = n
			}
		}
	}

	if query := paq.withCustomMethod; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductAssoc)
		for i := range nodes {
			if nodes[i].custom_method_product_assocs == nil {
				continue
			}
			fk := *nodes[i].custom_method_product_assocs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(custommethod.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "custom_method_product_assocs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CustomMethod = n
			}
		}
	}

	return nodes, nil
}

func (paq *ProductAssocQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := paq.querySpec()
	return sqlgraph.CountNodes(ctx, paq.driver, _spec)
}

func (paq *ProductAssocQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := paq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (paq *ProductAssocQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productassoc.Table,
			Columns: productassoc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productassoc.FieldID,
			},
		},
		From:   paq.sql,
		Unique: true,
	}
	if unique := paq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := paq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productassoc.FieldID)
		for i := range fields {
			if fields[i] != productassoc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := paq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := paq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := paq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := paq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (paq *ProductAssocQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(paq.driver.Dialect())
	t1 := builder.Table(productassoc.Table)
	columns := paq.fields
	if len(columns) == 0 {
		columns = productassoc.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if paq.sql != nil {
		selector = paq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range paq.predicates {
		p(selector)
	}
	for _, p := range paq.order {
		p(selector)
	}
	if offset := paq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := paq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductAssocGroupBy is the group-by builder for ProductAssoc entities.
type ProductAssocGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pagb *ProductAssocGroupBy) Aggregate(fns ...AggregateFunc) *ProductAssocGroupBy {
	pagb.fns = append(pagb.fns, fns...)
	return pagb
}

// Scan applies the group-by query and scans the result into the given value.
func (pagb *ProductAssocGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pagb.path(ctx)
	if err != nil {
		return err
	}
	pagb.sql = query
	return pagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) StringsX(ctx context.Context) []string {
	v, err := pagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) StringX(ctx context.Context) string {
	v, err := pagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) IntsX(ctx context.Context) []int {
	v, err := pagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) IntX(ctx context.Context) int {
	v, err := pagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pagb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pagb *ProductAssocGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pagb *ProductAssocGroupBy) BoolX(ctx context.Context) bool {
	v, err := pagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pagb *ProductAssocGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pagb.fields {
		if !productassoc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pagb *ProductAssocGroupBy) sqlQuery() *sql.Selector {
	selector := pagb.sql.Select()
	aggregation := make([]string, 0, len(pagb.fns))
	for _, fn := range pagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pagb.fields)+len(pagb.fns))
		for _, f := range pagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pagb.fields...)...)
}

// ProductAssocSelect is the builder for selecting fields of ProductAssoc entities.
type ProductAssocSelect struct {
	*ProductAssocQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pas *ProductAssocSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pas.prepareQuery(ctx); err != nil {
		return err
	}
	pas.sql = pas.ProductAssocQuery.sqlQuery(ctx)
	return pas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pas *ProductAssocSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: ProductAssocSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pas *ProductAssocSelect) StringsX(ctx context.Context) []string {
	v, err := pas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pas *ProductAssocSelect) StringX(ctx context.Context) string {
	v, err := pas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: ProductAssocSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pas *ProductAssocSelect) IntsX(ctx context.Context) []int {
	v, err := pas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pas *ProductAssocSelect) IntX(ctx context.Context) int {
	v, err := pas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: ProductAssocSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pas *ProductAssocSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pas *ProductAssocSelect) Float64X(ctx context.Context) float64 {
	v, err := pas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pas.fields) > 1 {
		return nil, errors.New("ent: ProductAssocSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pas *ProductAssocSelect) BoolsX(ctx context.Context) []bool {
	v, err := pas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pas *ProductAssocSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoc.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pas *ProductAssocSelect) BoolX(ctx context.Context) bool {
	v, err := pas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pas *ProductAssocSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pas.sql.Query()
	if err := pas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
