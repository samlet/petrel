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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/prodcatalogcategorytype"
)

// ProdCatalogCategoryTypeQuery is the builder for querying ProdCatalogCategoryType entities.
type ProdCatalogCategoryTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProdCatalogCategoryType
	// eager-loading edges.
	withParent                        *ProdCatalogCategoryTypeQuery
	withChildren                      *ProdCatalogCategoryTypeQuery
	withChildProdCatalogCategoryTypes *ProdCatalogCategoryTypeQuery
	withFKs                           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProdCatalogCategoryTypeQuery builder.
func (pcctq *ProdCatalogCategoryTypeQuery) Where(ps ...predicate.ProdCatalogCategoryType) *ProdCatalogCategoryTypeQuery {
	pcctq.predicates = append(pcctq.predicates, ps...)
	return pcctq
}

// Limit adds a limit step to the query.
func (pcctq *ProdCatalogCategoryTypeQuery) Limit(limit int) *ProdCatalogCategoryTypeQuery {
	pcctq.limit = &limit
	return pcctq
}

// Offset adds an offset step to the query.
func (pcctq *ProdCatalogCategoryTypeQuery) Offset(offset int) *ProdCatalogCategoryTypeQuery {
	pcctq.offset = &offset
	return pcctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcctq *ProdCatalogCategoryTypeQuery) Unique(unique bool) *ProdCatalogCategoryTypeQuery {
	pcctq.unique = &unique
	return pcctq
}

// Order adds an order step to the query.
func (pcctq *ProdCatalogCategoryTypeQuery) Order(o ...OrderFunc) *ProdCatalogCategoryTypeQuery {
	pcctq.order = append(pcctq.order, o...)
	return pcctq
}

// QueryParent chains the current query on the "parent" edge.
func (pcctq *ProdCatalogCategoryTypeQuery) QueryParent() *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID, selector),
			sqlgraph.To(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prodcatalogcategorytype.ParentTable, prodcatalogcategorytype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (pcctq *ProdCatalogCategoryTypeQuery) QueryChildren() *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID, selector),
			sqlgraph.To(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, prodcatalogcategorytype.ChildrenTable, prodcatalogcategorytype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildProdCatalogCategoryTypes chains the current query on the "child_prod_catalog_category_types" edge.
func (pcctq *ProdCatalogCategoryTypeQuery) QueryChildProdCatalogCategoryTypes() *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID, selector),
			sqlgraph.To(prodcatalogcategorytype.Table, prodcatalogcategorytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, prodcatalogcategorytype.ChildProdCatalogCategoryTypesTable, prodcatalogcategorytype.ChildProdCatalogCategoryTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pcctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProdCatalogCategoryType entity from the query.
// Returns a *NotFoundError when no ProdCatalogCategoryType was found.
func (pcctq *ProdCatalogCategoryTypeQuery) First(ctx context.Context) (*ProdCatalogCategoryType, error) {
	nodes, err := pcctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prodcatalogcategorytype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) FirstX(ctx context.Context) *ProdCatalogCategoryType {
	node, err := pcctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProdCatalogCategoryType ID from the query.
// Returns a *NotFoundError when no ProdCatalogCategoryType ID was found.
func (pcctq *ProdCatalogCategoryTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prodcatalogcategorytype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := pcctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProdCatalogCategoryType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProdCatalogCategoryType entity is not found.
// Returns a *NotFoundError when no ProdCatalogCategoryType entities are found.
func (pcctq *ProdCatalogCategoryTypeQuery) Only(ctx context.Context) (*ProdCatalogCategoryType, error) {
	nodes, err := pcctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prodcatalogcategorytype.Label}
	default:
		return nil, &NotSingularError{prodcatalogcategorytype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) OnlyX(ctx context.Context) *ProdCatalogCategoryType {
	node, err := pcctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProdCatalogCategoryType ID in the query.
// Returns a *NotSingularError when exactly one ProdCatalogCategoryType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pcctq *ProdCatalogCategoryTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pcctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = &NotSingularError{prodcatalogcategorytype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pcctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProdCatalogCategoryTypes.
func (pcctq *ProdCatalogCategoryTypeQuery) All(ctx context.Context) ([]*ProdCatalogCategoryType, error) {
	if err := pcctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) AllX(ctx context.Context) []*ProdCatalogCategoryType {
	nodes, err := pcctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProdCatalogCategoryType IDs.
func (pcctq *ProdCatalogCategoryTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pcctq.Select(prodcatalogcategorytype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := pcctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcctq *ProdCatalogCategoryTypeQuery) Count(ctx context.Context) (int, error) {
	if err := pcctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) CountX(ctx context.Context) int {
	count, err := pcctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcctq *ProdCatalogCategoryTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcctq *ProdCatalogCategoryTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := pcctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProdCatalogCategoryTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcctq *ProdCatalogCategoryTypeQuery) Clone() *ProdCatalogCategoryTypeQuery {
	if pcctq == nil {
		return nil
	}
	return &ProdCatalogCategoryTypeQuery{
		config:                            pcctq.config,
		limit:                             pcctq.limit,
		offset:                            pcctq.offset,
		order:                             append([]OrderFunc{}, pcctq.order...),
		predicates:                        append([]predicate.ProdCatalogCategoryType{}, pcctq.predicates...),
		withParent:                        pcctq.withParent.Clone(),
		withChildren:                      pcctq.withChildren.Clone(),
		withChildProdCatalogCategoryTypes: pcctq.withChildProdCatalogCategoryTypes.Clone(),
		// clone intermediate query.
		sql:  pcctq.sql.Clone(),
		path: pcctq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (pcctq *ProdCatalogCategoryTypeQuery) WithParent(opts ...func(*ProdCatalogCategoryTypeQuery)) *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcctq.withParent = query
	return pcctq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (pcctq *ProdCatalogCategoryTypeQuery) WithChildren(opts ...func(*ProdCatalogCategoryTypeQuery)) *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcctq.withChildren = query
	return pcctq
}

// WithChildProdCatalogCategoryTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_prod_catalog_category_types" edge. The optional arguments are used to configure the query builder of the edge.
func (pcctq *ProdCatalogCategoryTypeQuery) WithChildProdCatalogCategoryTypes(opts ...func(*ProdCatalogCategoryTypeQuery)) *ProdCatalogCategoryTypeQuery {
	query := &ProdCatalogCategoryTypeQuery{config: pcctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pcctq.withChildProdCatalogCategoryTypes = query
	return pcctq
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
//	client.ProdCatalogCategoryType.Query().
//		GroupBy(prodcatalogcategorytype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pcctq *ProdCatalogCategoryTypeQuery) GroupBy(field string, fields ...string) *ProdCatalogCategoryTypeGroupBy {
	group := &ProdCatalogCategoryTypeGroupBy{config: pcctq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcctq.sqlQuery(ctx), nil
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
//	client.ProdCatalogCategoryType.Query().
//		Select(prodcatalogcategorytype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pcctq *ProdCatalogCategoryTypeQuery) Select(field string, fields ...string) *ProdCatalogCategoryTypeSelect {
	pcctq.fields = append([]string{field}, fields...)
	return &ProdCatalogCategoryTypeSelect{ProdCatalogCategoryTypeQuery: pcctq}
}

func (pcctq *ProdCatalogCategoryTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcctq.fields {
		if !prodcatalogcategorytype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcctq.path != nil {
		prev, err := pcctq.path(ctx)
		if err != nil {
			return err
		}
		pcctq.sql = prev
	}
	return nil
}

func (pcctq *ProdCatalogCategoryTypeQuery) sqlAll(ctx context.Context) ([]*ProdCatalogCategoryType, error) {
	var (
		nodes       = []*ProdCatalogCategoryType{}
		withFKs     = pcctq.withFKs
		_spec       = pcctq.querySpec()
		loadedTypes = [3]bool{
			pcctq.withParent != nil,
			pcctq.withChildren != nil,
			pcctq.withChildProdCatalogCategoryTypes != nil,
		}
	)
	if pcctq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, prodcatalogcategorytype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProdCatalogCategoryType{config: pcctq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pcctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pcctq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ProdCatalogCategoryType)
		for i := range nodes {
			if nodes[i].prod_catalog_category_type_children == nil {
				continue
			}
			fk := *nodes[i].prod_catalog_category_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(prodcatalogcategorytype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "prod_catalog_category_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := pcctq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ProdCatalogCategoryType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*ProdCatalogCategoryType{}
		}
		query.withFKs = true
		query.Where(predicate.ProdCatalogCategoryType(func(s *sql.Selector) {
			s.Where(sql.InValues(prodcatalogcategorytype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.prod_catalog_category_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "prod_catalog_category_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "prod_catalog_category_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := pcctq.withChildProdCatalogCategoryTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ProdCatalogCategoryType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildProdCatalogCategoryTypes = []*ProdCatalogCategoryType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ProdCatalogCategoryType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   prodcatalogcategorytype.ChildProdCatalogCategoryTypesTable,
				Columns: prodcatalogcategorytype.ChildProdCatalogCategoryTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(prodcatalogcategorytype.ChildProdCatalogCategoryTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, pcctq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_prod_catalog_category_types": %w`, err)
		}
		query.Where(prodcatalogcategorytype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_prod_catalog_category_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildProdCatalogCategoryTypes = append(nodes[i].Edges.ChildProdCatalogCategoryTypes, n)
			}
		}
	}

	return nodes, nil
}

func (pcctq *ProdCatalogCategoryTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcctq.querySpec()
	return sqlgraph.CountNodes(ctx, pcctq.driver, _spec)
}

func (pcctq *ProdCatalogCategoryTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcctq *ProdCatalogCategoryTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prodcatalogcategorytype.Table,
			Columns: prodcatalogcategorytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prodcatalogcategorytype.FieldID,
			},
		},
		From:   pcctq.sql,
		Unique: true,
	}
	if unique := pcctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prodcatalogcategorytype.FieldID)
		for i := range fields {
			if fields[i] != prodcatalogcategorytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcctq *ProdCatalogCategoryTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcctq.driver.Dialect())
	t1 := builder.Table(prodcatalogcategorytype.Table)
	columns := pcctq.fields
	if len(columns) == 0 {
		columns = prodcatalogcategorytype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcctq.sql != nil {
		selector = pcctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pcctq.predicates {
		p(selector)
	}
	for _, p := range pcctq.order {
		p(selector)
	}
	if offset := pcctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProdCatalogCategoryTypeGroupBy is the group-by builder for ProdCatalogCategoryType entities.
type ProdCatalogCategoryTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Aggregate(fns ...AggregateFunc) *ProdCatalogCategoryTypeGroupBy {
	pcctgb.fns = append(pcctgb.fns, fns...)
	return pcctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcctgb.path(ctx)
	if err != nil {
		return err
	}
	pcctgb.sql = query
	return pcctgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pcctgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pcctgb.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pcctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pcctgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcctgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) StringX(ctx context.Context) string {
	v, err := pcctgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pcctgb.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pcctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pcctgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcctgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) IntX(ctx context.Context) int {
	v, err := pcctgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcctgb.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pcctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pcctgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcctgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pcctgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pcctgb.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pcctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pcctgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcctgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcctgb *ProdCatalogCategoryTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pcctgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcctgb *ProdCatalogCategoryTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcctgb.fields {
		if !prodcatalogcategorytype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcctgb *ProdCatalogCategoryTypeGroupBy) sqlQuery() *sql.Selector {
	selector := pcctgb.sql.Select()
	aggregation := make([]string, 0, len(pcctgb.fns))
	for _, fn := range pcctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcctgb.fields)+len(pcctgb.fns))
		for _, f := range pcctgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcctgb.fields...)...)
}

// ProdCatalogCategoryTypeSelect is the builder for selecting fields of ProdCatalogCategoryType entities.
type ProdCatalogCategoryTypeSelect struct {
	*ProdCatalogCategoryTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pccts *ProdCatalogCategoryTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pccts.prepareQuery(ctx); err != nil {
		return err
	}
	pccts.sql = pccts.ProdCatalogCategoryTypeQuery.sqlQuery(ctx)
	return pccts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pccts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pccts.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pccts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) StringsX(ctx context.Context) []string {
	v, err := pccts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pccts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) StringX(ctx context.Context) string {
	v, err := pccts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pccts.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pccts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) IntsX(ctx context.Context) []int {
	v, err := pccts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pccts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) IntX(ctx context.Context) int {
	v, err := pccts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pccts.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pccts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pccts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pccts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := pccts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pccts.fields) > 1 {
		return nil, errors.New("ent: ProdCatalogCategoryTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pccts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pccts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pccts *ProdCatalogCategoryTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pccts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prodcatalogcategorytype.Label}
	default:
		err = fmt.Errorf("ent: ProdCatalogCategoryTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pccts *ProdCatalogCategoryTypeSelect) BoolX(ctx context.Context) bool {
	v, err := pccts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pccts *ProdCatalogCategoryTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pccts.sql.Query()
	if err := pccts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
