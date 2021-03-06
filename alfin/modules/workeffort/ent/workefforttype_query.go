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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
)

// WorkEffortTypeQuery is the builder for querying WorkEffortType entities.
type WorkEffortTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WorkEffortType
	// eager-loading edges.
	withParent               *WorkEffortTypeQuery
	withChildren             *WorkEffortTypeQuery
	withWorkEfforts          *WorkEffortQuery
	withChildWorkEffortTypes *WorkEffortTypeQuery
	withFKs                  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WorkEffortTypeQuery builder.
func (wetq *WorkEffortTypeQuery) Where(ps ...predicate.WorkEffortType) *WorkEffortTypeQuery {
	wetq.predicates = append(wetq.predicates, ps...)
	return wetq
}

// Limit adds a limit step to the query.
func (wetq *WorkEffortTypeQuery) Limit(limit int) *WorkEffortTypeQuery {
	wetq.limit = &limit
	return wetq
}

// Offset adds an offset step to the query.
func (wetq *WorkEffortTypeQuery) Offset(offset int) *WorkEffortTypeQuery {
	wetq.offset = &offset
	return wetq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wetq *WorkEffortTypeQuery) Unique(unique bool) *WorkEffortTypeQuery {
	wetq.unique = &unique
	return wetq
}

// Order adds an order step to the query.
func (wetq *WorkEffortTypeQuery) Order(o ...OrderFunc) *WorkEffortTypeQuery {
	wetq.order = append(wetq.order, o...)
	return wetq
}

// QueryParent chains the current query on the "parent" edge.
func (wetq *WorkEffortTypeQuery) QueryParent() *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workefforttype.Table, workefforttype.FieldID, selector),
			sqlgraph.To(workefforttype.Table, workefforttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, workefforttype.ParentTable, workefforttype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(wetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (wetq *WorkEffortTypeQuery) QueryChildren() *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workefforttype.Table, workefforttype.FieldID, selector),
			sqlgraph.To(workefforttype.Table, workefforttype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workefforttype.ChildrenTable, workefforttype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(wetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkEfforts chains the current query on the "work_efforts" edge.
func (wetq *WorkEffortTypeQuery) QueryWorkEfforts() *WorkEffortQuery {
	query := &WorkEffortQuery{config: wetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workefforttype.Table, workefforttype.FieldID, selector),
			sqlgraph.To(workeffort.Table, workeffort.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workefforttype.WorkEffortsTable, workefforttype.WorkEffortsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildWorkEffortTypes chains the current query on the "child_work_effort_types" edge.
func (wetq *WorkEffortTypeQuery) QueryChildWorkEffortTypes() *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workefforttype.Table, workefforttype.FieldID, selector),
			sqlgraph.To(workefforttype.Table, workefforttype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, workefforttype.ChildWorkEffortTypesTable, workefforttype.ChildWorkEffortTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkEffortType entity from the query.
// Returns a *NotFoundError when no WorkEffortType was found.
func (wetq *WorkEffortTypeQuery) First(ctx context.Context) (*WorkEffortType, error) {
	nodes, err := wetq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{workefforttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) FirstX(ctx context.Context) *WorkEffortType {
	node, err := wetq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WorkEffortType ID from the query.
// Returns a *NotFoundError when no WorkEffortType ID was found.
func (wetq *WorkEffortTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wetq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workefforttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := wetq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WorkEffortType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WorkEffortType entity is not found.
// Returns a *NotFoundError when no WorkEffortType entities are found.
func (wetq *WorkEffortTypeQuery) Only(ctx context.Context) (*WorkEffortType, error) {
	nodes, err := wetq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{workefforttype.Label}
	default:
		return nil, &NotSingularError{workefforttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) OnlyX(ctx context.Context) *WorkEffortType {
	node, err := wetq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WorkEffortType ID in the query.
// Returns a *NotSingularError when exactly one WorkEffortType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wetq *WorkEffortTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wetq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = &NotSingularError{workefforttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := wetq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkEffortTypes.
func (wetq *WorkEffortTypeQuery) All(ctx context.Context) ([]*WorkEffortType, error) {
	if err := wetq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wetq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) AllX(ctx context.Context) []*WorkEffortType {
	nodes, err := wetq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WorkEffortType IDs.
func (wetq *WorkEffortTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wetq.Select(workefforttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := wetq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wetq *WorkEffortTypeQuery) Count(ctx context.Context) (int, error) {
	if err := wetq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wetq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) CountX(ctx context.Context) int {
	count, err := wetq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wetq *WorkEffortTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wetq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wetq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wetq *WorkEffortTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := wetq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WorkEffortTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wetq *WorkEffortTypeQuery) Clone() *WorkEffortTypeQuery {
	if wetq == nil {
		return nil
	}
	return &WorkEffortTypeQuery{
		config:                   wetq.config,
		limit:                    wetq.limit,
		offset:                   wetq.offset,
		order:                    append([]OrderFunc{}, wetq.order...),
		predicates:               append([]predicate.WorkEffortType{}, wetq.predicates...),
		withParent:               wetq.withParent.Clone(),
		withChildren:             wetq.withChildren.Clone(),
		withWorkEfforts:          wetq.withWorkEfforts.Clone(),
		withChildWorkEffortTypes: wetq.withChildWorkEffortTypes.Clone(),
		// clone intermediate query.
		sql:  wetq.sql.Clone(),
		path: wetq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (wetq *WorkEffortTypeQuery) WithParent(opts ...func(*WorkEffortTypeQuery)) *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	for _, opt := range opts {
		opt(query)
	}
	wetq.withParent = query
	return wetq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (wetq *WorkEffortTypeQuery) WithChildren(opts ...func(*WorkEffortTypeQuery)) *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	for _, opt := range opts {
		opt(query)
	}
	wetq.withChildren = query
	return wetq
}

// WithWorkEfforts tells the query-builder to eager-load the nodes that are connected to
// the "work_efforts" edge. The optional arguments are used to configure the query builder of the edge.
func (wetq *WorkEffortTypeQuery) WithWorkEfforts(opts ...func(*WorkEffortQuery)) *WorkEffortTypeQuery {
	query := &WorkEffortQuery{config: wetq.config}
	for _, opt := range opts {
		opt(query)
	}
	wetq.withWorkEfforts = query
	return wetq
}

// WithChildWorkEffortTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_work_effort_types" edge. The optional arguments are used to configure the query builder of the edge.
func (wetq *WorkEffortTypeQuery) WithChildWorkEffortTypes(opts ...func(*WorkEffortTypeQuery)) *WorkEffortTypeQuery {
	query := &WorkEffortTypeQuery{config: wetq.config}
	for _, opt := range opts {
		opt(query)
	}
	wetq.withChildWorkEffortTypes = query
	return wetq
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
//	client.WorkEffortType.Query().
//		GroupBy(workefforttype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wetq *WorkEffortTypeQuery) GroupBy(field string, fields ...string) *WorkEffortTypeGroupBy {
	group := &WorkEffortTypeGroupBy{config: wetq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wetq.sqlQuery(ctx), nil
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
//	client.WorkEffortType.Query().
//		Select(workefforttype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (wetq *WorkEffortTypeQuery) Select(field string, fields ...string) *WorkEffortTypeSelect {
	wetq.fields = append([]string{field}, fields...)
	return &WorkEffortTypeSelect{WorkEffortTypeQuery: wetq}
}

func (wetq *WorkEffortTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wetq.fields {
		if !workefforttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wetq.path != nil {
		prev, err := wetq.path(ctx)
		if err != nil {
			return err
		}
		wetq.sql = prev
	}
	return nil
}

func (wetq *WorkEffortTypeQuery) sqlAll(ctx context.Context) ([]*WorkEffortType, error) {
	var (
		nodes       = []*WorkEffortType{}
		withFKs     = wetq.withFKs
		_spec       = wetq.querySpec()
		loadedTypes = [4]bool{
			wetq.withParent != nil,
			wetq.withChildren != nil,
			wetq.withWorkEfforts != nil,
			wetq.withChildWorkEffortTypes != nil,
		}
	)
	if wetq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workefforttype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WorkEffortType{config: wetq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wetq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wetq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkEffortType)
		for i := range nodes {
			if nodes[i].work_effort_type_children == nil {
				continue
			}
			fk := *nodes[i].work_effort_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(workefforttype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_effort_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := wetq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkEffortType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*WorkEffortType{}
		}
		query.withFKs = true
		query.Where(predicate.WorkEffortType(func(s *sql.Selector) {
			s.Where(sql.InValues(workefforttype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.work_effort_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "work_effort_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_effort_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := wetq.withWorkEfforts; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkEffortType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.WorkEfforts = []*WorkEffort{}
		}
		query.withFKs = true
		query.Where(predicate.WorkEffort(func(s *sql.Selector) {
			s.Where(sql.InValues(workefforttype.WorkEffortsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.work_effort_type_work_efforts
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "work_effort_type_work_efforts" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_effort_type_work_efforts" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.WorkEfforts = append(node.Edges.WorkEfforts, n)
		}
	}

	if query := wetq.withChildWorkEffortTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*WorkEffortType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildWorkEffortTypes = []*WorkEffortType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*WorkEffortType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   workefforttype.ChildWorkEffortTypesTable,
				Columns: workefforttype.ChildWorkEffortTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(workefforttype.ChildWorkEffortTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, wetq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_work_effort_types": %w`, err)
		}
		query.Where(workefforttype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_work_effort_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildWorkEffortTypes = append(nodes[i].Edges.ChildWorkEffortTypes, n)
			}
		}
	}

	return nodes, nil
}

func (wetq *WorkEffortTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wetq.querySpec()
	return sqlgraph.CountNodes(ctx, wetq.driver, _spec)
}

func (wetq *WorkEffortTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wetq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wetq *WorkEffortTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workefforttype.Table,
			Columns: workefforttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workefforttype.FieldID,
			},
		},
		From:   wetq.sql,
		Unique: true,
	}
	if unique := wetq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wetq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workefforttype.FieldID)
		for i := range fields {
			if fields[i] != workefforttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wetq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wetq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wetq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wetq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wetq *WorkEffortTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wetq.driver.Dialect())
	t1 := builder.Table(workefforttype.Table)
	columns := wetq.fields
	if len(columns) == 0 {
		columns = workefforttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wetq.sql != nil {
		selector = wetq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range wetq.predicates {
		p(selector)
	}
	for _, p := range wetq.order {
		p(selector)
	}
	if offset := wetq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wetq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkEffortTypeGroupBy is the group-by builder for WorkEffortType entities.
type WorkEffortTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wetgb *WorkEffortTypeGroupBy) Aggregate(fns ...AggregateFunc) *WorkEffortTypeGroupBy {
	wetgb.fns = append(wetgb.fns, fns...)
	return wetgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wetgb *WorkEffortTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wetgb.path(ctx)
	if err != nil {
		return err
	}
	wetgb.sql = query
	return wetgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wetgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wetgb.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wetgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wetgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) StringX(ctx context.Context) string {
	v, err := wetgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wetgb.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wetgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wetgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) IntX(ctx context.Context) int {
	v, err := wetgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wetgb.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wetgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wetgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wetgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wetgb.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wetgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wetgb *WorkEffortTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wetgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wetgb *WorkEffortTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wetgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wetgb *WorkEffortTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wetgb.fields {
		if !workefforttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wetgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wetgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wetgb *WorkEffortTypeGroupBy) sqlQuery() *sql.Selector {
	selector := wetgb.sql.Select()
	aggregation := make([]string, 0, len(wetgb.fns))
	for _, fn := range wetgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wetgb.fields)+len(wetgb.fns))
		for _, f := range wetgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wetgb.fields...)...)
}

// WorkEffortTypeSelect is the builder for selecting fields of WorkEffortType entities.
type WorkEffortTypeSelect struct {
	*WorkEffortTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wets *WorkEffortTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wets.prepareQuery(ctx); err != nil {
		return err
	}
	wets.sql = wets.WorkEffortTypeQuery.sqlQuery(ctx)
	return wets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wets.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) StringsX(ctx context.Context) []string {
	v, err := wets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wets.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) StringX(ctx context.Context) string {
	v, err := wets.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wets.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) IntsX(ctx context.Context) []int {
	v, err := wets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wets.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) IntX(ctx context.Context) int {
	v, err := wets.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wets.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wets.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := wets.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wets.fields) > 1 {
		return nil, errors.New("ent: WorkEffortTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := wets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wets *WorkEffortTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wets.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workefforttype.Label}
	default:
		err = fmt.Errorf("ent: WorkEffortTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wets *WorkEffortTypeSelect) BoolX(ctx context.Context) bool {
	v, err := wets.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wets *WorkEffortTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wets.sql.Query()
	if err := wets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
