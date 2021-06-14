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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumerationtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// EnumerationTypeQuery is the builder for querying EnumerationType entities.
type EnumerationTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EnumerationType
	// eager-loading edges.
	withParent                *EnumerationTypeQuery
	withChildren              *EnumerationTypeQuery
	withEnumerations          *EnumerationQuery
	withChildEnumerationTypes *EnumerationTypeQuery
	withFKs                   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EnumerationTypeQuery builder.
func (etq *EnumerationTypeQuery) Where(ps ...predicate.EnumerationType) *EnumerationTypeQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit adds a limit step to the query.
func (etq *EnumerationTypeQuery) Limit(limit int) *EnumerationTypeQuery {
	etq.limit = &limit
	return etq
}

// Offset adds an offset step to the query.
func (etq *EnumerationTypeQuery) Offset(offset int) *EnumerationTypeQuery {
	etq.offset = &offset
	return etq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (etq *EnumerationTypeQuery) Unique(unique bool) *EnumerationTypeQuery {
	etq.unique = &unique
	return etq
}

// Order adds an order step to the query.
func (etq *EnumerationTypeQuery) Order(o ...OrderFunc) *EnumerationTypeQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// QueryParent chains the current query on the "parent" edge.
func (etq *EnumerationTypeQuery) QueryParent() *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(enumerationtype.Table, enumerationtype.FieldID, selector),
			sqlgraph.To(enumerationtype.Table, enumerationtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, enumerationtype.ParentTable, enumerationtype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (etq *EnumerationTypeQuery) QueryChildren() *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(enumerationtype.Table, enumerationtype.FieldID, selector),
			sqlgraph.To(enumerationtype.Table, enumerationtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, enumerationtype.ChildrenTable, enumerationtype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEnumerations chains the current query on the "enumerations" edge.
func (etq *EnumerationTypeQuery) QueryEnumerations() *EnumerationQuery {
	query := &EnumerationQuery{config: etq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(enumerationtype.Table, enumerationtype.FieldID, selector),
			sqlgraph.To(enumeration.Table, enumeration.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, enumerationtype.EnumerationsTable, enumerationtype.EnumerationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildEnumerationTypes chains the current query on the "child_enumeration_types" edge.
func (etq *EnumerationTypeQuery) QueryChildEnumerationTypes() *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(enumerationtype.Table, enumerationtype.FieldID, selector),
			sqlgraph.To(enumerationtype.Table, enumerationtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, enumerationtype.ChildEnumerationTypesTable, enumerationtype.ChildEnumerationTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EnumerationType entity from the query.
// Returns a *NotFoundError when no EnumerationType was found.
func (etq *EnumerationTypeQuery) First(ctx context.Context) (*EnumerationType, error) {
	nodes, err := etq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enumerationtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EnumerationTypeQuery) FirstX(ctx context.Context) *EnumerationType {
	node, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EnumerationType ID from the query.
// Returns a *NotFoundError when no EnumerationType ID was found.
func (etq *EnumerationTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = etq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enumerationtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (etq *EnumerationTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EnumerationType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one EnumerationType entity is not found.
// Returns a *NotFoundError when no EnumerationType entities are found.
func (etq *EnumerationTypeQuery) Only(ctx context.Context) (*EnumerationType, error) {
	nodes, err := etq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enumerationtype.Label}
	default:
		return nil, &NotSingularError{enumerationtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EnumerationTypeQuery) OnlyX(ctx context.Context) *EnumerationType {
	node, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EnumerationType ID in the query.
// Returns a *NotSingularError when exactly one EnumerationType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (etq *EnumerationTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = etq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = &NotSingularError{enumerationtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (etq *EnumerationTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EnumerationTypes.
func (etq *EnumerationTypeQuery) All(ctx context.Context) ([]*EnumerationType, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return etq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (etq *EnumerationTypeQuery) AllX(ctx context.Context) []*EnumerationType {
	nodes, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EnumerationType IDs.
func (etq *EnumerationTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := etq.Select(enumerationtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EnumerationTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EnumerationTypeQuery) Count(ctx context.Context) (int, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return etq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EnumerationTypeQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EnumerationTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return etq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EnumerationTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EnumerationTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EnumerationTypeQuery) Clone() *EnumerationTypeQuery {
	if etq == nil {
		return nil
	}
	return &EnumerationTypeQuery{
		config:                    etq.config,
		limit:                     etq.limit,
		offset:                    etq.offset,
		order:                     append([]OrderFunc{}, etq.order...),
		predicates:                append([]predicate.EnumerationType{}, etq.predicates...),
		withParent:                etq.withParent.Clone(),
		withChildren:              etq.withChildren.Clone(),
		withEnumerations:          etq.withEnumerations.Clone(),
		withChildEnumerationTypes: etq.withChildEnumerationTypes.Clone(),
		// clone intermediate query.
		sql:  etq.sql.Clone(),
		path: etq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EnumerationTypeQuery) WithParent(opts ...func(*EnumerationTypeQuery)) *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	for _, opt := range opts {
		opt(query)
	}
	etq.withParent = query
	return etq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EnumerationTypeQuery) WithChildren(opts ...func(*EnumerationTypeQuery)) *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	for _, opt := range opts {
		opt(query)
	}
	etq.withChildren = query
	return etq
}

// WithEnumerations tells the query-builder to eager-load the nodes that are connected to
// the "enumerations" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EnumerationTypeQuery) WithEnumerations(opts ...func(*EnumerationQuery)) *EnumerationTypeQuery {
	query := &EnumerationQuery{config: etq.config}
	for _, opt := range opts {
		opt(query)
	}
	etq.withEnumerations = query
	return etq
}

// WithChildEnumerationTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_enumeration_types" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EnumerationTypeQuery) WithChildEnumerationTypes(opts ...func(*EnumerationTypeQuery)) *EnumerationTypeQuery {
	query := &EnumerationTypeQuery{config: etq.config}
	for _, opt := range opts {
		opt(query)
	}
	etq.withChildEnumerationTypes = query
	return etq
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
//	client.EnumerationType.Query().
//		GroupBy(enumerationtype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (etq *EnumerationTypeQuery) GroupBy(field string, fields ...string) *EnumerationTypeGroupBy {
	group := &EnumerationTypeGroupBy{config: etq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return etq.sqlQuery(ctx), nil
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
//	client.EnumerationType.Query().
//		Select(enumerationtype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (etq *EnumerationTypeQuery) Select(field string, fields ...string) *EnumerationTypeSelect {
	etq.fields = append([]string{field}, fields...)
	return &EnumerationTypeSelect{EnumerationTypeQuery: etq}
}

func (etq *EnumerationTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range etq.fields {
		if !enumerationtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if etq.path != nil {
		prev, err := etq.path(ctx)
		if err != nil {
			return err
		}
		etq.sql = prev
	}
	return nil
}

func (etq *EnumerationTypeQuery) sqlAll(ctx context.Context) ([]*EnumerationType, error) {
	var (
		nodes       = []*EnumerationType{}
		withFKs     = etq.withFKs
		_spec       = etq.querySpec()
		loadedTypes = [4]bool{
			etq.withParent != nil,
			etq.withChildren != nil,
			etq.withEnumerations != nil,
			etq.withChildEnumerationTypes != nil,
		}
	)
	if etq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, enumerationtype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EnumerationType{config: etq.config}
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
	if err := sqlgraph.QueryNodes(ctx, etq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := etq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EnumerationType)
		for i := range nodes {
			if nodes[i].enumeration_type_children == nil {
				continue
			}
			fk := *nodes[i].enumeration_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(enumerationtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "enumeration_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := etq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*EnumerationType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*EnumerationType{}
		}
		query.withFKs = true
		query.Where(predicate.EnumerationType(func(s *sql.Selector) {
			s.Where(sql.InValues(enumerationtype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.enumeration_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "enumeration_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "enumeration_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := etq.withEnumerations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*EnumerationType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Enumerations = []*Enumeration{}
		}
		query.withFKs = true
		query.Where(predicate.Enumeration(func(s *sql.Selector) {
			s.Where(sql.InValues(enumerationtype.EnumerationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.enumeration_type_enumerations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "enumeration_type_enumerations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "enumeration_type_enumerations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Enumerations = append(node.Edges.Enumerations, n)
		}
	}

	if query := etq.withChildEnumerationTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*EnumerationType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildEnumerationTypes = []*EnumerationType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*EnumerationType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   enumerationtype.ChildEnumerationTypesTable,
				Columns: enumerationtype.ChildEnumerationTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(enumerationtype.ChildEnumerationTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, etq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_enumeration_types": %w`, err)
		}
		query.Where(enumerationtype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_enumeration_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildEnumerationTypes = append(nodes[i].Edges.ChildEnumerationTypes, n)
			}
		}
	}

	return nodes, nil
}

func (etq *EnumerationTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := etq.querySpec()
	return sqlgraph.CountNodes(ctx, etq.driver, _spec)
}

func (etq *EnumerationTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := etq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (etq *EnumerationTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   enumerationtype.Table,
			Columns: enumerationtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enumerationtype.FieldID,
			},
		},
		From:   etq.sql,
		Unique: true,
	}
	if unique := etq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := etq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enumerationtype.FieldID)
		for i := range fields {
			if fields[i] != enumerationtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := etq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := etq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := etq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := etq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (etq *EnumerationTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(enumerationtype.Table)
	columns := etq.fields
	if len(columns) == 0 {
		columns = enumerationtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EnumerationTypeGroupBy is the group-by builder for EnumerationType entities.
type EnumerationTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EnumerationTypeGroupBy) Aggregate(fns ...AggregateFunc) *EnumerationTypeGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the group-by query and scans the result into the given value.
func (etgb *EnumerationTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := etgb.path(ctx)
	if err != nil {
		return err
	}
	etgb.sql = query
	return etgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := etgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := etgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = etgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) StringX(ctx context.Context) string {
	v, err := etgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := etgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = etgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) IntX(ctx context.Context) int {
	v, err := etgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := etgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = etgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := etgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := etgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EnumerationTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = etgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (etgb *EnumerationTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := etgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (etgb *EnumerationTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range etgb.fields {
		if !enumerationtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := etgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := etgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (etgb *EnumerationTypeGroupBy) sqlQuery() *sql.Selector {
	selector := etgb.sql.Select()
	aggregation := make([]string, 0, len(etgb.fns))
	for _, fn := range etgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(etgb.fields)+len(etgb.fns))
		for _, f := range etgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(etgb.fields...)...)
}

// EnumerationTypeSelect is the builder for selecting fields of EnumerationType entities.
type EnumerationTypeSelect struct {
	*EnumerationTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ets *EnumerationTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ets.prepareQuery(ctx); err != nil {
		return err
	}
	ets.sql = ets.EnumerationTypeQuery.sqlQuery(ctx)
	return ets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ets *EnumerationTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ets *EnumerationTypeSelect) StringsX(ctx context.Context) []string {
	v, err := ets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ets.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ets *EnumerationTypeSelect) StringX(ctx context.Context) string {
	v, err := ets.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ets *EnumerationTypeSelect) IntsX(ctx context.Context) []int {
	v, err := ets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ets.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ets *EnumerationTypeSelect) IntX(ctx context.Context) int {
	v, err := ets.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ets *EnumerationTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ets.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ets *EnumerationTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := ets.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EnumerationTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ets *EnumerationTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ets *EnumerationTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ets.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{enumerationtype.Label}
	default:
		err = fmt.Errorf("ent: EnumerationTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ets *EnumerationTypeSelect) BoolX(ctx context.Context) bool {
	v, err := ets.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ets *EnumerationTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ets.sql.Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}