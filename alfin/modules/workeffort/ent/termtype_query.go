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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/termtype"
)

// TermTypeQuery is the builder for querying TermType entities.
type TermTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TermType
	// eager-loading edges.
	withParent         *TermTypeQuery
	withChildren       *TermTypeQuery
	withChildTermTypes *TermTypeQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TermTypeQuery builder.
func (ttq *TermTypeQuery) Where(ps ...predicate.TermType) *TermTypeQuery {
	ttq.predicates = append(ttq.predicates, ps...)
	return ttq
}

// Limit adds a limit step to the query.
func (ttq *TermTypeQuery) Limit(limit int) *TermTypeQuery {
	ttq.limit = &limit
	return ttq
}

// Offset adds an offset step to the query.
func (ttq *TermTypeQuery) Offset(offset int) *TermTypeQuery {
	ttq.offset = &offset
	return ttq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ttq *TermTypeQuery) Unique(unique bool) *TermTypeQuery {
	ttq.unique = &unique
	return ttq
}

// Order adds an order step to the query.
func (ttq *TermTypeQuery) Order(o ...OrderFunc) *TermTypeQuery {
	ttq.order = append(ttq.order, o...)
	return ttq
}

// QueryParent chains the current query on the "parent" edge.
func (ttq *TermTypeQuery) QueryParent() *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(termtype.Table, termtype.FieldID, selector),
			sqlgraph.To(termtype.Table, termtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, termtype.ParentTable, termtype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (ttq *TermTypeQuery) QueryChildren() *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(termtype.Table, termtype.FieldID, selector),
			sqlgraph.To(termtype.Table, termtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, termtype.ChildrenTable, termtype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildTermTypes chains the current query on the "child_term_types" edge.
func (ttq *TermTypeQuery) QueryChildTermTypes() *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ttq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(termtype.Table, termtype.FieldID, selector),
			sqlgraph.To(termtype.Table, termtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, termtype.ChildTermTypesTable, termtype.ChildTermTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ttq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TermType entity from the query.
// Returns a *NotFoundError when no TermType was found.
func (ttq *TermTypeQuery) First(ctx context.Context) (*TermType, error) {
	nodes, err := ttq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{termtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ttq *TermTypeQuery) FirstX(ctx context.Context) *TermType {
	node, err := ttq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TermType ID from the query.
// Returns a *NotFoundError when no TermType ID was found.
func (ttq *TermTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{termtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ttq *TermTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ttq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TermType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TermType entity is not found.
// Returns a *NotFoundError when no TermType entities are found.
func (ttq *TermTypeQuery) Only(ctx context.Context) (*TermType, error) {
	nodes, err := ttq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{termtype.Label}
	default:
		return nil, &NotSingularError{termtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ttq *TermTypeQuery) OnlyX(ctx context.Context) *TermType {
	node, err := ttq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TermType ID in the query.
// Returns a *NotSingularError when exactly one TermType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ttq *TermTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ttq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = &NotSingularError{termtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ttq *TermTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ttq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TermTypes.
func (ttq *TermTypeQuery) All(ctx context.Context) ([]*TermType, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ttq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ttq *TermTypeQuery) AllX(ctx context.Context) []*TermType {
	nodes, err := ttq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TermType IDs.
func (ttq *TermTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ttq.Select(termtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ttq *TermTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ttq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ttq *TermTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ttq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ttq *TermTypeQuery) CountX(ctx context.Context) int {
	count, err := ttq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ttq *TermTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ttq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ttq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ttq *TermTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ttq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TermTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ttq *TermTypeQuery) Clone() *TermTypeQuery {
	if ttq == nil {
		return nil
	}
	return &TermTypeQuery{
		config:             ttq.config,
		limit:              ttq.limit,
		offset:             ttq.offset,
		order:              append([]OrderFunc{}, ttq.order...),
		predicates:         append([]predicate.TermType{}, ttq.predicates...),
		withParent:         ttq.withParent.Clone(),
		withChildren:       ttq.withChildren.Clone(),
		withChildTermTypes: ttq.withChildTermTypes.Clone(),
		// clone intermediate query.
		sql:  ttq.sql.Clone(),
		path: ttq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TermTypeQuery) WithParent(opts ...func(*TermTypeQuery)) *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withParent = query
	return ttq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TermTypeQuery) WithChildren(opts ...func(*TermTypeQuery)) *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withChildren = query
	return ttq
}

// WithChildTermTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_term_types" edge. The optional arguments are used to configure the query builder of the edge.
func (ttq *TermTypeQuery) WithChildTermTypes(opts ...func(*TermTypeQuery)) *TermTypeQuery {
	query := &TermTypeQuery{config: ttq.config}
	for _, opt := range opts {
		opt(query)
	}
	ttq.withChildTermTypes = query
	return ttq
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
//	client.TermType.Query().
//		GroupBy(termtype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ttq *TermTypeQuery) GroupBy(field string, fields ...string) *TermTypeGroupBy {
	group := &TermTypeGroupBy{config: ttq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ttq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ttq.sqlQuery(ctx), nil
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
//	client.TermType.Query().
//		Select(termtype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (ttq *TermTypeQuery) Select(field string, fields ...string) *TermTypeSelect {
	ttq.fields = append([]string{field}, fields...)
	return &TermTypeSelect{TermTypeQuery: ttq}
}

func (ttq *TermTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ttq.fields {
		if !termtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ttq.path != nil {
		prev, err := ttq.path(ctx)
		if err != nil {
			return err
		}
		ttq.sql = prev
	}
	return nil
}

func (ttq *TermTypeQuery) sqlAll(ctx context.Context) ([]*TermType, error) {
	var (
		nodes       = []*TermType{}
		withFKs     = ttq.withFKs
		_spec       = ttq.querySpec()
		loadedTypes = [3]bool{
			ttq.withParent != nil,
			ttq.withChildren != nil,
			ttq.withChildTermTypes != nil,
		}
	)
	if ttq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, termtype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TermType{config: ttq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ttq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ttq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*TermType)
		for i := range nodes {
			if nodes[i].term_type_children == nil {
				continue
			}
			fk := *nodes[i].term_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(termtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "term_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ttq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*TermType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*TermType{}
		}
		query.withFKs = true
		query.Where(predicate.TermType(func(s *sql.Selector) {
			s.Where(sql.InValues(termtype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.term_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "term_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "term_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := ttq.withChildTermTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*TermType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildTermTypes = []*TermType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*TermType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   termtype.ChildTermTypesTable,
				Columns: termtype.ChildTermTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(termtype.ChildTermTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, ttq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_term_types": %w`, err)
		}
		query.Where(termtype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_term_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildTermTypes = append(nodes[i].Edges.ChildTermTypes, n)
			}
		}
	}

	return nodes, nil
}

func (ttq *TermTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ttq.querySpec()
	return sqlgraph.CountNodes(ctx, ttq.driver, _spec)
}

func (ttq *TermTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ttq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ttq *TermTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   termtype.Table,
			Columns: termtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: termtype.FieldID,
			},
		},
		From:   ttq.sql,
		Unique: true,
	}
	if unique := ttq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ttq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, termtype.FieldID)
		for i := range fields {
			if fields[i] != termtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ttq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ttq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ttq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ttq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ttq *TermTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ttq.driver.Dialect())
	t1 := builder.Table(termtype.Table)
	columns := ttq.fields
	if len(columns) == 0 {
		columns = termtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ttq.sql != nil {
		selector = ttq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ttq.predicates {
		p(selector)
	}
	for _, p := range ttq.order {
		p(selector)
	}
	if offset := ttq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ttq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TermTypeGroupBy is the group-by builder for TermType entities.
type TermTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ttgb *TermTypeGroupBy) Aggregate(fns ...AggregateFunc) *TermTypeGroupBy {
	ttgb.fns = append(ttgb.fns, fns...)
	return ttgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ttgb *TermTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ttgb.path(ctx)
	if err != nil {
		return err
	}
	ttgb.sql = query
	return ttgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ttgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TermTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ttgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ttgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ttgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TermTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ttgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ttgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ttgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TermTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ttgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ttgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ttgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ttgb.fields) > 1 {
		return nil, errors.New("ent: TermTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ttgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ttgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ttgb *TermTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ttgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ttgb *TermTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ttgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ttgb *TermTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ttgb.fields {
		if !termtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ttgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ttgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ttgb *TermTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ttgb.sql.Select()
	aggregation := make([]string, 0, len(ttgb.fns))
	for _, fn := range ttgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ttgb.fields)+len(ttgb.fns))
		for _, f := range ttgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ttgb.fields...)...)
}

// TermTypeSelect is the builder for selecting fields of TermType entities.
type TermTypeSelect struct {
	*TermTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tts *TermTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tts.prepareQuery(ctx); err != nil {
		return err
	}
	tts.sql = tts.TermTypeQuery.sqlQuery(ctx)
	return tts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tts *TermTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TermTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tts *TermTypeSelect) StringsX(ctx context.Context) []string {
	v, err := tts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tts *TermTypeSelect) StringX(ctx context.Context) string {
	v, err := tts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TermTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tts *TermTypeSelect) IntsX(ctx context.Context) []int {
	v, err := tts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tts *TermTypeSelect) IntX(ctx context.Context) int {
	v, err := tts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TermTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tts *TermTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tts *TermTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := tts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tts.fields) > 1 {
		return nil, errors.New("ent: TermTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tts *TermTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := tts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tts *TermTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{termtype.Label}
	default:
		err = fmt.Errorf("ent: TermTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tts *TermTypeSelect) BoolX(ctx context.Context) bool {
	v, err := tts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tts *TermTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tts.sql.Query()
	if err := tts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
