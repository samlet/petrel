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
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderadjustment"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroupassoc"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// OrderAdjustmentQuery is the builder for querying OrderAdjustment entities.
type OrderAdjustmentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderAdjustment
	// eager-loading edges.
	withOrderHeader             *OrderHeaderQuery
	withOrderItem               *OrderItemQuery
	withOrderItemShipGroup      *OrderItemShipGroupQuery
	withOrderItemShipGroupAssoc *OrderItemShipGroupAssocQuery
	withParent                  *OrderAdjustmentQuery
	withChildren                *OrderAdjustmentQuery
	withFKs                     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderAdjustmentQuery builder.
func (oaq *OrderAdjustmentQuery) Where(ps ...predicate.OrderAdjustment) *OrderAdjustmentQuery {
	oaq.predicates = append(oaq.predicates, ps...)
	return oaq
}

// Limit adds a limit step to the query.
func (oaq *OrderAdjustmentQuery) Limit(limit int) *OrderAdjustmentQuery {
	oaq.limit = &limit
	return oaq
}

// Offset adds an offset step to the query.
func (oaq *OrderAdjustmentQuery) Offset(offset int) *OrderAdjustmentQuery {
	oaq.offset = &offset
	return oaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oaq *OrderAdjustmentQuery) Unique(unique bool) *OrderAdjustmentQuery {
	oaq.unique = &unique
	return oaq
}

// Order adds an order step to the query.
func (oaq *OrderAdjustmentQuery) Order(o ...OrderFunc) *OrderAdjustmentQuery {
	oaq.order = append(oaq.order, o...)
	return oaq
}

// QueryOrderHeader chains the current query on the "order_header" edge.
func (oaq *OrderAdjustmentQuery) QueryOrderHeader() *OrderHeaderQuery {
	query := &OrderHeaderQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderheader.Table, orderheader.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderadjustment.OrderHeaderTable, orderadjustment.OrderHeaderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItem chains the current query on the "order_item" edge.
func (oaq *OrderAdjustmentQuery) QueryOrderItem() *OrderItemQuery {
	query := &OrderItemQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderadjustment.OrderItemTable, orderadjustment.OrderItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItemShipGroup chains the current query on the "order_item_ship_group" edge.
func (oaq *OrderAdjustmentQuery) QueryOrderItemShipGroup() *OrderItemShipGroupQuery {
	query := &OrderItemShipGroupQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderitemshipgroup.Table, orderitemshipgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderadjustment.OrderItemShipGroupTable, orderadjustment.OrderItemShipGroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItemShipGroupAssoc chains the current query on the "order_item_ship_group_assoc" edge.
func (oaq *OrderAdjustmentQuery) QueryOrderItemShipGroupAssoc() *OrderItemShipGroupAssocQuery {
	query := &OrderItemShipGroupAssocQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderitemshipgroupassoc.Table, orderitemshipgroupassoc.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderadjustment.OrderItemShipGroupAssocTable, orderadjustment.OrderItemShipGroupAssocColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (oaq *OrderAdjustmentQuery) QueryParent() *OrderAdjustmentQuery {
	query := &OrderAdjustmentQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderadjustment.Table, orderadjustment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderadjustment.ParentTable, orderadjustment.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (oaq *OrderAdjustmentQuery) QueryChildren() *OrderAdjustmentQuery {
	query := &OrderAdjustmentQuery{config: oaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderadjustment.Table, orderadjustment.FieldID, selector),
			sqlgraph.To(orderadjustment.Table, orderadjustment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, orderadjustment.ChildrenTable, orderadjustment.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderAdjustment entity from the query.
// Returns a *NotFoundError when no OrderAdjustment was found.
func (oaq *OrderAdjustmentQuery) First(ctx context.Context) (*OrderAdjustment, error) {
	nodes, err := oaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderadjustment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) FirstX(ctx context.Context) *OrderAdjustment {
	node, err := oaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderAdjustment ID from the query.
// Returns a *NotFoundError when no OrderAdjustment ID was found.
func (oaq *OrderAdjustmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderadjustment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) FirstIDX(ctx context.Context) int {
	id, err := oaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderAdjustment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrderAdjustment entity is not found.
// Returns a *NotFoundError when no OrderAdjustment entities are found.
func (oaq *OrderAdjustmentQuery) Only(ctx context.Context) (*OrderAdjustment, error) {
	nodes, err := oaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderadjustment.Label}
	default:
		return nil, &NotSingularError{orderadjustment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) OnlyX(ctx context.Context) *OrderAdjustment {
	node, err := oaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderAdjustment ID in the query.
// Returns a *NotSingularError when exactly one OrderAdjustment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oaq *OrderAdjustmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = &NotSingularError{orderadjustment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := oaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderAdjustments.
func (oaq *OrderAdjustmentQuery) All(ctx context.Context) ([]*OrderAdjustment, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) AllX(ctx context.Context) []*OrderAdjustment {
	nodes, err := oaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderAdjustment IDs.
func (oaq *OrderAdjustmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oaq.Select(orderadjustment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) IDsX(ctx context.Context) []int {
	ids, err := oaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oaq *OrderAdjustmentQuery) Count(ctx context.Context) (int, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) CountX(ctx context.Context) int {
	count, err := oaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oaq *OrderAdjustmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := oaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oaq *OrderAdjustmentQuery) ExistX(ctx context.Context) bool {
	exist, err := oaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderAdjustmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oaq *OrderAdjustmentQuery) Clone() *OrderAdjustmentQuery {
	if oaq == nil {
		return nil
	}
	return &OrderAdjustmentQuery{
		config:                      oaq.config,
		limit:                       oaq.limit,
		offset:                      oaq.offset,
		order:                       append([]OrderFunc{}, oaq.order...),
		predicates:                  append([]predicate.OrderAdjustment{}, oaq.predicates...),
		withOrderHeader:             oaq.withOrderHeader.Clone(),
		withOrderItem:               oaq.withOrderItem.Clone(),
		withOrderItemShipGroup:      oaq.withOrderItemShipGroup.Clone(),
		withOrderItemShipGroupAssoc: oaq.withOrderItemShipGroupAssoc.Clone(),
		withParent:                  oaq.withParent.Clone(),
		withChildren:                oaq.withChildren.Clone(),
		// clone intermediate query.
		sql:  oaq.sql.Clone(),
		path: oaq.path,
	}
}

// WithOrderHeader tells the query-builder to eager-load the nodes that are connected to
// the "order_header" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithOrderHeader(opts ...func(*OrderHeaderQuery)) *OrderAdjustmentQuery {
	query := &OrderHeaderQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrderHeader = query
	return oaq
}

// WithOrderItem tells the query-builder to eager-load the nodes that are connected to
// the "order_item" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithOrderItem(opts ...func(*OrderItemQuery)) *OrderAdjustmentQuery {
	query := &OrderItemQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrderItem = query
	return oaq
}

// WithOrderItemShipGroup tells the query-builder to eager-load the nodes that are connected to
// the "order_item_ship_group" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithOrderItemShipGroup(opts ...func(*OrderItemShipGroupQuery)) *OrderAdjustmentQuery {
	query := &OrderItemShipGroupQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrderItemShipGroup = query
	return oaq
}

// WithOrderItemShipGroupAssoc tells the query-builder to eager-load the nodes that are connected to
// the "order_item_ship_group_assoc" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithOrderItemShipGroupAssoc(opts ...func(*OrderItemShipGroupAssocQuery)) *OrderAdjustmentQuery {
	query := &OrderItemShipGroupAssocQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrderItemShipGroupAssoc = query
	return oaq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithParent(opts ...func(*OrderAdjustmentQuery)) *OrderAdjustmentQuery {
	query := &OrderAdjustmentQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withParent = query
	return oaq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAdjustmentQuery) WithChildren(opts ...func(*OrderAdjustmentQuery)) *OrderAdjustmentQuery {
	query := &OrderAdjustmentQuery{config: oaq.config}
	for _, opt := range opts {
		opt(query)
	}
	oaq.withChildren = query
	return oaq
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
//	client.OrderAdjustment.Query().
//		GroupBy(orderadjustment.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oaq *OrderAdjustmentQuery) GroupBy(field string, fields ...string) *OrderAdjustmentGroupBy {
	group := &OrderAdjustmentGroupBy{config: oaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oaq.sqlQuery(ctx), nil
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
//	client.OrderAdjustment.Query().
//		Select(orderadjustment.FieldCreateTime).
//		Scan(ctx, &v)
//
func (oaq *OrderAdjustmentQuery) Select(field string, fields ...string) *OrderAdjustmentSelect {
	oaq.fields = append([]string{field}, fields...)
	return &OrderAdjustmentSelect{OrderAdjustmentQuery: oaq}
}

func (oaq *OrderAdjustmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oaq.fields {
		if !orderadjustment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oaq.path != nil {
		prev, err := oaq.path(ctx)
		if err != nil {
			return err
		}
		oaq.sql = prev
	}
	return nil
}

func (oaq *OrderAdjustmentQuery) sqlAll(ctx context.Context) ([]*OrderAdjustment, error) {
	var (
		nodes       = []*OrderAdjustment{}
		withFKs     = oaq.withFKs
		_spec       = oaq.querySpec()
		loadedTypes = [6]bool{
			oaq.withOrderHeader != nil,
			oaq.withOrderItem != nil,
			oaq.withOrderItemShipGroup != nil,
			oaq.withOrderItemShipGroupAssoc != nil,
			oaq.withParent != nil,
			oaq.withChildren != nil,
		}
	)
	if oaq.withOrderHeader != nil || oaq.withOrderItem != nil || oaq.withOrderItemShipGroup != nil || oaq.withOrderItemShipGroupAssoc != nil || oaq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderadjustment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrderAdjustment{config: oaq.config}
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
	if err := sqlgraph.QueryNodes(ctx, oaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oaq.withOrderHeader; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAdjustment)
		for i := range nodes {
			if nodes[i].order_header_order_adjustments == nil {
				continue
			}
			fk := *nodes[i].order_header_order_adjustments
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
				return nil, fmt.Errorf(`unexpected foreign-key "order_header_order_adjustments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderHeader = n
			}
		}
	}

	if query := oaq.withOrderItem; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAdjustment)
		for i := range nodes {
			if nodes[i].order_item_order_adjustments == nil {
				continue
			}
			fk := *nodes[i].order_item_order_adjustments
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
				return nil, fmt.Errorf(`unexpected foreign-key "order_item_order_adjustments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderItem = n
			}
		}
	}

	if query := oaq.withOrderItemShipGroup; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAdjustment)
		for i := range nodes {
			if nodes[i].order_item_ship_group_order_adjustments == nil {
				continue
			}
			fk := *nodes[i].order_item_ship_group_order_adjustments
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderitemshipgroup.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_item_ship_group_order_adjustments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderItemShipGroup = n
			}
		}
	}

	if query := oaq.withOrderItemShipGroupAssoc; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAdjustment)
		for i := range nodes {
			if nodes[i].order_item_ship_group_assoc_order_adjustments == nil {
				continue
			}
			fk := *nodes[i].order_item_ship_group_assoc_order_adjustments
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderitemshipgroupassoc.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_item_ship_group_assoc_order_adjustments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.OrderItemShipGroupAssoc = n
			}
		}
	}

	if query := oaq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrderAdjustment)
		for i := range nodes {
			if nodes[i].order_adjustment_children == nil {
				continue
			}
			fk := *nodes[i].order_adjustment_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orderadjustment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_adjustment_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := oaq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OrderAdjustment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*OrderAdjustment{}
		}
		query.withFKs = true
		query.Where(predicate.OrderAdjustment(func(s *sql.Selector) {
			s.Where(sql.InValues(orderadjustment.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.order_adjustment_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "order_adjustment_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "order_adjustment_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	return nodes, nil
}

func (oaq *OrderAdjustmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oaq.querySpec()
	return sqlgraph.CountNodes(ctx, oaq.driver, _spec)
}

func (oaq *OrderAdjustmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oaq *OrderAdjustmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderadjustment.Table,
			Columns: orderadjustment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderadjustment.FieldID,
			},
		},
		From:   oaq.sql,
		Unique: true,
	}
	if unique := oaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderadjustment.FieldID)
		for i := range fields {
			if fields[i] != orderadjustment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oaq *OrderAdjustmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oaq.driver.Dialect())
	t1 := builder.Table(orderadjustment.Table)
	columns := oaq.fields
	if len(columns) == 0 {
		columns = orderadjustment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oaq.sql != nil {
		selector = oaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range oaq.predicates {
		p(selector)
	}
	for _, p := range oaq.order {
		p(selector)
	}
	if offset := oaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderAdjustmentGroupBy is the group-by builder for OrderAdjustment entities.
type OrderAdjustmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagb *OrderAdjustmentGroupBy) Aggregate(fns ...AggregateFunc) *OrderAdjustmentGroupBy {
	oagb.fns = append(oagb.fns, fns...)
	return oagb
}

// Scan applies the group-by query and scans the result into the given value.
func (oagb *OrderAdjustmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oagb.path(ctx)
	if err != nil {
		return err
	}
	oagb.sql = query
	return oagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := oagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) StringX(ctx context.Context) string {
	v, err := oagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := oagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) IntX(ctx context.Context) int {
	v, err := oagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oagb.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oagb *OrderAdjustmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oagb *OrderAdjustmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := oagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oagb *OrderAdjustmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oagb.fields {
		if !orderadjustment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oagb *OrderAdjustmentGroupBy) sqlQuery() *sql.Selector {
	selector := oagb.sql.Select()
	aggregation := make([]string, 0, len(oagb.fns))
	for _, fn := range oagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oagb.fields)+len(oagb.fns))
		for _, f := range oagb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oagb.fields...)...)
}

// OrderAdjustmentSelect is the builder for selecting fields of OrderAdjustment entities.
type OrderAdjustmentSelect struct {
	*OrderAdjustmentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oas *OrderAdjustmentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oas.prepareQuery(ctx); err != nil {
		return err
	}
	oas.sql = oas.OrderAdjustmentQuery.sqlQuery(ctx)
	return oas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) StringsX(ctx context.Context) []string {
	v, err := oas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) StringX(ctx context.Context) string {
	v, err := oas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) IntsX(ctx context.Context) []int {
	v, err := oas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) IntX(ctx context.Context) int {
	v, err := oas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) Float64X(ctx context.Context) float64 {
	v, err := oas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oas.fields) > 1 {
		return nil, errors.New("ent: OrderAdjustmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := oas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oas *OrderAdjustmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderadjustment.Label}
	default:
		err = fmt.Errorf("ent: OrderAdjustmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oas *OrderAdjustmentSelect) BoolX(ctx context.Context) bool {
	v, err := oas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oas *OrderAdjustmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oas.sql.Query()
	if err := oas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
