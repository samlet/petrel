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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productassoctype"
)

// ProductAssocTypeQuery is the builder for querying ProductAssocType entities.
type ProductAssocTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProductAssocType
	// eager-loading edges.
	withParent                 *ProductAssocTypeQuery
	withChildren               *ProductAssocTypeQuery
	withProductAssocs          *ProductAssocQuery
	withChildProductAssocTypes *ProductAssocTypeQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductAssocTypeQuery builder.
func (patq *ProductAssocTypeQuery) Where(ps ...predicate.ProductAssocType) *ProductAssocTypeQuery {
	patq.predicates = append(patq.predicates, ps...)
	return patq
}

// Limit adds a limit step to the query.
func (patq *ProductAssocTypeQuery) Limit(limit int) *ProductAssocTypeQuery {
	patq.limit = &limit
	return patq
}

// Offset adds an offset step to the query.
func (patq *ProductAssocTypeQuery) Offset(offset int) *ProductAssocTypeQuery {
	patq.offset = &offset
	return patq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (patq *ProductAssocTypeQuery) Unique(unique bool) *ProductAssocTypeQuery {
	patq.unique = &unique
	return patq
}

// Order adds an order step to the query.
func (patq *ProductAssocTypeQuery) Order(o ...OrderFunc) *ProductAssocTypeQuery {
	patq.order = append(patq.order, o...)
	return patq
}

// QueryParent chains the current query on the "parent" edge.
func (patq *ProductAssocTypeQuery) QueryParent() *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := patq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoctype.Table, productassoctype.FieldID, selector),
			sqlgraph.To(productassoctype.Table, productassoctype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, productassoctype.ParentTable, productassoctype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(patq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (patq *ProductAssocTypeQuery) QueryChildren() *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := patq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoctype.Table, productassoctype.FieldID, selector),
			sqlgraph.To(productassoctype.Table, productassoctype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productassoctype.ChildrenTable, productassoctype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(patq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductAssocs chains the current query on the "product_assocs" edge.
func (patq *ProductAssocTypeQuery) QueryProductAssocs() *ProductAssocQuery {
	query := &ProductAssocQuery{config: patq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := patq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoctype.Table, productassoctype.FieldID, selector),
			sqlgraph.To(productassoc.Table, productassoc.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productassoctype.ProductAssocsTable, productassoctype.ProductAssocsColumn),
		)
		fromU = sqlgraph.SetNeighbors(patq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildProductAssocTypes chains the current query on the "child_product_assoc_types" edge.
func (patq *ProductAssocTypeQuery) QueryChildProductAssocTypes() *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := patq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productassoctype.Table, productassoctype.FieldID, selector),
			sqlgraph.To(productassoctype.Table, productassoctype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, productassoctype.ChildProductAssocTypesTable, productassoctype.ChildProductAssocTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(patq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductAssocType entity from the query.
// Returns a *NotFoundError when no ProductAssocType was found.
func (patq *ProductAssocTypeQuery) First(ctx context.Context) (*ProductAssocType, error) {
	nodes, err := patq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productassoctype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) FirstX(ctx context.Context) *ProductAssocType {
	node, err := patq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductAssocType ID from the query.
// Returns a *NotFoundError when no ProductAssocType ID was found.
func (patq *ProductAssocTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = patq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productassoctype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := patq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductAssocType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProductAssocType entity is not found.
// Returns a *NotFoundError when no ProductAssocType entities are found.
func (patq *ProductAssocTypeQuery) Only(ctx context.Context) (*ProductAssocType, error) {
	nodes, err := patq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productassoctype.Label}
	default:
		return nil, &NotSingularError{productassoctype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) OnlyX(ctx context.Context) *ProductAssocType {
	node, err := patq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductAssocType ID in the query.
// Returns a *NotSingularError when exactly one ProductAssocType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (patq *ProductAssocTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = patq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = &NotSingularError{productassoctype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := patq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductAssocTypes.
func (patq *ProductAssocTypeQuery) All(ctx context.Context) ([]*ProductAssocType, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return patq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) AllX(ctx context.Context) []*ProductAssocType {
	nodes, err := patq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductAssocType IDs.
func (patq *ProductAssocTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := patq.Select(productassoctype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := patq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (patq *ProductAssocTypeQuery) Count(ctx context.Context) (int, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return patq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) CountX(ctx context.Context) int {
	count, err := patq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (patq *ProductAssocTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := patq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return patq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (patq *ProductAssocTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := patq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductAssocTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (patq *ProductAssocTypeQuery) Clone() *ProductAssocTypeQuery {
	if patq == nil {
		return nil
	}
	return &ProductAssocTypeQuery{
		config:                     patq.config,
		limit:                      patq.limit,
		offset:                     patq.offset,
		order:                      append([]OrderFunc{}, patq.order...),
		predicates:                 append([]predicate.ProductAssocType{}, patq.predicates...),
		withParent:                 patq.withParent.Clone(),
		withChildren:               patq.withChildren.Clone(),
		withProductAssocs:          patq.withProductAssocs.Clone(),
		withChildProductAssocTypes: patq.withChildProductAssocTypes.Clone(),
		// clone intermediate query.
		sql:  patq.sql.Clone(),
		path: patq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (patq *ProductAssocTypeQuery) WithParent(opts ...func(*ProductAssocTypeQuery)) *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	for _, opt := range opts {
		opt(query)
	}
	patq.withParent = query
	return patq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (patq *ProductAssocTypeQuery) WithChildren(opts ...func(*ProductAssocTypeQuery)) *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	for _, opt := range opts {
		opt(query)
	}
	patq.withChildren = query
	return patq
}

// WithProductAssocs tells the query-builder to eager-load the nodes that are connected to
// the "product_assocs" edge. The optional arguments are used to configure the query builder of the edge.
func (patq *ProductAssocTypeQuery) WithProductAssocs(opts ...func(*ProductAssocQuery)) *ProductAssocTypeQuery {
	query := &ProductAssocQuery{config: patq.config}
	for _, opt := range opts {
		opt(query)
	}
	patq.withProductAssocs = query
	return patq
}

// WithChildProductAssocTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_product_assoc_types" edge. The optional arguments are used to configure the query builder of the edge.
func (patq *ProductAssocTypeQuery) WithChildProductAssocTypes(opts ...func(*ProductAssocTypeQuery)) *ProductAssocTypeQuery {
	query := &ProductAssocTypeQuery{config: patq.config}
	for _, opt := range opts {
		opt(query)
	}
	patq.withChildProductAssocTypes = query
	return patq
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
//	client.ProductAssocType.Query().
//		GroupBy(productassoctype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (patq *ProductAssocTypeQuery) GroupBy(field string, fields ...string) *ProductAssocTypeGroupBy {
	group := &ProductAssocTypeGroupBy{config: patq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := patq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return patq.sqlQuery(ctx), nil
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
//	client.ProductAssocType.Query().
//		Select(productassoctype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (patq *ProductAssocTypeQuery) Select(field string, fields ...string) *ProductAssocTypeSelect {
	patq.fields = append([]string{field}, fields...)
	return &ProductAssocTypeSelect{ProductAssocTypeQuery: patq}
}

func (patq *ProductAssocTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range patq.fields {
		if !productassoctype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if patq.path != nil {
		prev, err := patq.path(ctx)
		if err != nil {
			return err
		}
		patq.sql = prev
	}
	return nil
}

func (patq *ProductAssocTypeQuery) sqlAll(ctx context.Context) ([]*ProductAssocType, error) {
	var (
		nodes       = []*ProductAssocType{}
		withFKs     = patq.withFKs
		_spec       = patq.querySpec()
		loadedTypes = [4]bool{
			patq.withParent != nil,
			patq.withChildren != nil,
			patq.withProductAssocs != nil,
			patq.withChildProductAssocTypes != nil,
		}
	)
	if patq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, productassoctype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProductAssocType{config: patq.config}
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
	if err := sqlgraph.QueryNodes(ctx, patq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := patq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProductAssocType)
		for i := range nodes {
			if nodes[i].product_assoc_type_children == nil {
				continue
			}
			fk := *nodes[i].product_assoc_type_children
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
				return nil, fmt.Errorf(`unexpected foreign-key "product_assoc_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := patq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ProductAssocType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*ProductAssocType{}
		}
		query.withFKs = true
		query.Where(predicate.ProductAssocType(func(s *sql.Selector) {
			s.Where(sql.InValues(productassoctype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.product_assoc_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "product_assoc_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_assoc_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := patq.withProductAssocs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ProductAssocType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ProductAssocs = []*ProductAssoc{}
		}
		query.withFKs = true
		query.Where(predicate.ProductAssoc(func(s *sql.Selector) {
			s.Where(sql.InValues(productassoctype.ProductAssocsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.product_assoc_type_product_assocs
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "product_assoc_type_product_assocs" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_assoc_type_product_assocs" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ProductAssocs = append(node.Edges.ProductAssocs, n)
		}
	}

	if query := patq.withChildProductAssocTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ProductAssocType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildProductAssocTypes = []*ProductAssocType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ProductAssocType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   productassoctype.ChildProductAssocTypesTable,
				Columns: productassoctype.ChildProductAssocTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(productassoctype.ChildProductAssocTypesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, patq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_product_assoc_types": %w`, err)
		}
		query.Where(productassoctype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_product_assoc_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildProductAssocTypes = append(nodes[i].Edges.ChildProductAssocTypes, n)
			}
		}
	}

	return nodes, nil
}

func (patq *ProductAssocTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := patq.querySpec()
	return sqlgraph.CountNodes(ctx, patq.driver, _spec)
}

func (patq *ProductAssocTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := patq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (patq *ProductAssocTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productassoctype.Table,
			Columns: productassoctype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productassoctype.FieldID,
			},
		},
		From:   patq.sql,
		Unique: true,
	}
	if unique := patq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := patq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productassoctype.FieldID)
		for i := range fields {
			if fields[i] != productassoctype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := patq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := patq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := patq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := patq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (patq *ProductAssocTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(patq.driver.Dialect())
	t1 := builder.Table(productassoctype.Table)
	columns := patq.fields
	if len(columns) == 0 {
		columns = productassoctype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if patq.sql != nil {
		selector = patq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range patq.predicates {
		p(selector)
	}
	for _, p := range patq.order {
		p(selector)
	}
	if offset := patq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := patq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductAssocTypeGroupBy is the group-by builder for ProductAssocType entities.
type ProductAssocTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (patgb *ProductAssocTypeGroupBy) Aggregate(fns ...AggregateFunc) *ProductAssocTypeGroupBy {
	patgb.fns = append(patgb.fns, fns...)
	return patgb
}

// Scan applies the group-by query and scans the result into the given value.
func (patgb *ProductAssocTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := patgb.path(ctx)
	if err != nil {
		return err
	}
	patgb.sql = query
	return patgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := patgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := patgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = patgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) StringX(ctx context.Context) string {
	v, err := patgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := patgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = patgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) IntX(ctx context.Context) int {
	v, err := patgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := patgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = patgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := patgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(patgb.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := patgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := patgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (patgb *ProductAssocTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = patgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (patgb *ProductAssocTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := patgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (patgb *ProductAssocTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range patgb.fields {
		if !productassoctype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := patgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := patgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (patgb *ProductAssocTypeGroupBy) sqlQuery() *sql.Selector {
	selector := patgb.sql.Select()
	aggregation := make([]string, 0, len(patgb.fns))
	for _, fn := range patgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(patgb.fields)+len(patgb.fns))
		for _, f := range patgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(patgb.fields...)...)
}

// ProductAssocTypeSelect is the builder for selecting fields of ProductAssocType entities.
type ProductAssocTypeSelect struct {
	*ProductAssocTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pats *ProductAssocTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pats.prepareQuery(ctx); err != nil {
		return err
	}
	pats.sql = pats.ProductAssocTypeQuery.sqlQuery(ctx)
	return pats.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pats.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) StringsX(ctx context.Context) []string {
	v, err := pats.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pats.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) StringX(ctx context.Context) string {
	v, err := pats.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) IntsX(ctx context.Context) []int {
	v, err := pats.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pats.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) IntX(ctx context.Context) int {
	v, err := pats.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pats.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pats.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := pats.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pats.fields) > 1 {
		return nil, errors.New("ent: ProductAssocTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pats.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pats.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pats *ProductAssocTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pats.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{productassoctype.Label}
	default:
		err = fmt.Errorf("ent: ProductAssocTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pats *ProductAssocTypeSelect) BoolX(ctx context.Context) bool {
	v, err := pats.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pats *ProductAssocTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pats.sql.Query()
	if err := pats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
