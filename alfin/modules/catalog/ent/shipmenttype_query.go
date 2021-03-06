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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmenttype"
)

// ShipmentTypeQuery is the builder for querying ShipmentType entities.
type ShipmentTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ShipmentType
	// eager-loading edges.
	withParent             *ShipmentTypeQuery
	withChildren           *ShipmentTypeQuery
	withChildShipmentTypes *ShipmentTypeQuery
	withFKs                bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentTypeQuery builder.
func (stq *ShipmentTypeQuery) Where(ps ...predicate.ShipmentType) *ShipmentTypeQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *ShipmentTypeQuery) Limit(limit int) *ShipmentTypeQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *ShipmentTypeQuery) Offset(offset int) *ShipmentTypeQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *ShipmentTypeQuery) Unique(unique bool) *ShipmentTypeQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *ShipmentTypeQuery) Order(o ...OrderFunc) *ShipmentTypeQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// QueryParent chains the current query on the "parent" edge.
func (stq *ShipmentTypeQuery) QueryParent() *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmenttype.Table, shipmenttype.FieldID, selector),
			sqlgraph.To(shipmenttype.Table, shipmenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shipmenttype.ParentTable, shipmenttype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (stq *ShipmentTypeQuery) QueryChildren() *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmenttype.Table, shipmenttype.FieldID, selector),
			sqlgraph.To(shipmenttype.Table, shipmenttype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, shipmenttype.ChildrenTable, shipmenttype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildShipmentTypes chains the current query on the "child_shipment_types" edge.
func (stq *ShipmentTypeQuery) QueryChildShipmentTypes() *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmenttype.Table, shipmenttype.FieldID, selector),
			sqlgraph.To(shipmenttype.Table, shipmenttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, shipmenttype.ChildShipmentTypesTable, shipmenttype.ChildShipmentTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentType entity from the query.
// Returns a *NotFoundError when no ShipmentType was found.
func (stq *ShipmentTypeQuery) First(ctx context.Context) (*ShipmentType, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmenttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *ShipmentTypeQuery) FirstX(ctx context.Context) *ShipmentType {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentType ID from the query.
// Returns a *NotFoundError when no ShipmentType ID was found.
func (stq *ShipmentTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmenttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *ShipmentTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ShipmentType entity is not found.
// Returns a *NotFoundError when no ShipmentType entities are found.
func (stq *ShipmentTypeQuery) Only(ctx context.Context) (*ShipmentType, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmenttype.Label}
	default:
		return nil, &NotSingularError{shipmenttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *ShipmentTypeQuery) OnlyX(ctx context.Context) *ShipmentType {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentType ID in the query.
// Returns a *NotSingularError when exactly one ShipmentType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (stq *ShipmentTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = &NotSingularError{shipmenttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *ShipmentTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentTypes.
func (stq *ShipmentTypeQuery) All(ctx context.Context) ([]*ShipmentType, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *ShipmentTypeQuery) AllX(ctx context.Context) []*ShipmentType {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentType IDs.
func (stq *ShipmentTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := stq.Select(shipmenttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *ShipmentTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *ShipmentTypeQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *ShipmentTypeQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *ShipmentTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *ShipmentTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *ShipmentTypeQuery) Clone() *ShipmentTypeQuery {
	if stq == nil {
		return nil
	}
	return &ShipmentTypeQuery{
		config:                 stq.config,
		limit:                  stq.limit,
		offset:                 stq.offset,
		order:                  append([]OrderFunc{}, stq.order...),
		predicates:             append([]predicate.ShipmentType{}, stq.predicates...),
		withParent:             stq.withParent.Clone(),
		withChildren:           stq.withChildren.Clone(),
		withChildShipmentTypes: stq.withChildShipmentTypes.Clone(),
		// clone intermediate query.
		sql:  stq.sql.Clone(),
		path: stq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShipmentTypeQuery) WithParent(opts ...func(*ShipmentTypeQuery)) *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withParent = query
	return stq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShipmentTypeQuery) WithChildren(opts ...func(*ShipmentTypeQuery)) *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withChildren = query
	return stq
}

// WithChildShipmentTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_shipment_types" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShipmentTypeQuery) WithChildShipmentTypes(opts ...func(*ShipmentTypeQuery)) *ShipmentTypeQuery {
	query := &ShipmentTypeQuery{config: stq.config}
	for _, opt := range opts {
		opt(query)
	}
	stq.withChildShipmentTypes = query
	return stq
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
//	client.ShipmentType.Query().
//		GroupBy(shipmenttype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (stq *ShipmentTypeQuery) GroupBy(field string, fields ...string) *ShipmentTypeGroupBy {
	group := &ShipmentTypeGroupBy{config: stq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
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
//	client.ShipmentType.Query().
//		Select(shipmenttype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (stq *ShipmentTypeQuery) Select(field string, fields ...string) *ShipmentTypeSelect {
	stq.fields = append([]string{field}, fields...)
	return &ShipmentTypeSelect{ShipmentTypeQuery: stq}
}

func (stq *ShipmentTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !shipmenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *ShipmentTypeQuery) sqlAll(ctx context.Context) ([]*ShipmentType, error) {
	var (
		nodes       = []*ShipmentType{}
		withFKs     = stq.withFKs
		_spec       = stq.querySpec()
		loadedTypes = [3]bool{
			stq.withParent != nil,
			stq.withChildren != nil,
			stq.withChildShipmentTypes != nil,
		}
	)
	if stq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shipmenttype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ShipmentType{config: stq.config}
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
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := stq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ShipmentType)
		for i := range nodes {
			if nodes[i].shipment_type_children == nil {
				continue
			}
			fk := *nodes[i].shipment_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(shipmenttype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := stq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*ShipmentType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*ShipmentType{}
		}
		query.withFKs = true
		query.Where(predicate.ShipmentType(func(s *sql.Selector) {
			s.Where(sql.InValues(shipmenttype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.shipment_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "shipment_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shipment_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := stq.withChildShipmentTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*ShipmentType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildShipmentTypes = []*ShipmentType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*ShipmentType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   shipmenttype.ChildShipmentTypesTable,
				Columns: shipmenttype.ChildShipmentTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(shipmenttype.ChildShipmentTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, stq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_shipment_types": %w`, err)
		}
		query.Where(shipmenttype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_shipment_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildShipmentTypes = append(nodes[i].Edges.ChildShipmentTypes, n)
			}
		}
	}

	return nodes, nil
}

func (stq *ShipmentTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *ShipmentTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := stq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (stq *ShipmentTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shipmenttype.Table,
			Columns: shipmenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shipmenttype.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmenttype.FieldID)
		for i := range fields {
			if fields[i] != shipmenttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *ShipmentTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(shipmenttype.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = shipmenttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShipmentTypeGroupBy is the group-by builder for ShipmentType entities.
type ShipmentTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *ShipmentTypeGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentTypeGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *ShipmentTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := stgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := stgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = stgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) StringX(ctx context.Context) string {
	v, err := stgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := stgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = stgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) IntX(ctx context.Context) int {
	v, err := stgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := stgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = stgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := stgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(stgb.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := stgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := stgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (stgb *ShipmentTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = stgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (stgb *ShipmentTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := stgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (stgb *ShipmentTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range stgb.fields {
		if !shipmenttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *ShipmentTypeGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// ShipmentTypeSelect is the builder for selecting fields of ShipmentType entities.
type ShipmentTypeSelect struct {
	*ShipmentTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *ShipmentTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.ShipmentTypeQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sts *ShipmentTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sts *ShipmentTypeSelect) StringsX(ctx context.Context) []string {
	v, err := sts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sts *ShipmentTypeSelect) StringX(ctx context.Context) string {
	v, err := sts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sts *ShipmentTypeSelect) IntsX(ctx context.Context) []int {
	v, err := sts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sts *ShipmentTypeSelect) IntX(ctx context.Context) int {
	v, err := sts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sts *ShipmentTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sts *ShipmentTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := sts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sts.fields) > 1 {
		return nil, errors.New("ent: ShipmentTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sts *ShipmentTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := sts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sts *ShipmentTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{shipmenttype.Label}
	default:
		err = fmt.Errorf("ent: ShipmentTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sts *ShipmentTypeSelect) BoolX(ctx context.Context) bool {
	v, err := sts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sts *ShipmentTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
