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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyclassificationtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// PartyClassificationTypeQuery is the builder for querying PartyClassificationType entities.
type PartyClassificationTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PartyClassificationType
	// eager-loading edges.
	withParent                        *PartyClassificationTypeQuery
	withChildren                      *PartyClassificationTypeQuery
	withChildPartyClassificationTypes *PartyClassificationTypeQuery
	withFKs                           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PartyClassificationTypeQuery builder.
func (pctq *PartyClassificationTypeQuery) Where(ps ...predicate.PartyClassificationType) *PartyClassificationTypeQuery {
	pctq.predicates = append(pctq.predicates, ps...)
	return pctq
}

// Limit adds a limit step to the query.
func (pctq *PartyClassificationTypeQuery) Limit(limit int) *PartyClassificationTypeQuery {
	pctq.limit = &limit
	return pctq
}

// Offset adds an offset step to the query.
func (pctq *PartyClassificationTypeQuery) Offset(offset int) *PartyClassificationTypeQuery {
	pctq.offset = &offset
	return pctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pctq *PartyClassificationTypeQuery) Unique(unique bool) *PartyClassificationTypeQuery {
	pctq.unique = &unique
	return pctq
}

// Order adds an order step to the query.
func (pctq *PartyClassificationTypeQuery) Order(o ...OrderFunc) *PartyClassificationTypeQuery {
	pctq.order = append(pctq.order, o...)
	return pctq
}

// QueryParent chains the current query on the "parent" edge.
func (pctq *PartyClassificationTypeQuery) QueryParent() *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partyclassificationtype.Table, partyclassificationtype.FieldID, selector),
			sqlgraph.To(partyclassificationtype.Table, partyclassificationtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, partyclassificationtype.ParentTable, partyclassificationtype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(pctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (pctq *PartyClassificationTypeQuery) QueryChildren() *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partyclassificationtype.Table, partyclassificationtype.FieldID, selector),
			sqlgraph.To(partyclassificationtype.Table, partyclassificationtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, partyclassificationtype.ChildrenTable, partyclassificationtype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(pctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildPartyClassificationTypes chains the current query on the "child_party_classification_types" edge.
func (pctq *PartyClassificationTypeQuery) QueryChildPartyClassificationTypes() *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(partyclassificationtype.Table, partyclassificationtype.FieldID, selector),
			sqlgraph.To(partyclassificationtype.Table, partyclassificationtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, partyclassificationtype.ChildPartyClassificationTypesTable, partyclassificationtype.ChildPartyClassificationTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PartyClassificationType entity from the query.
// Returns a *NotFoundError when no PartyClassificationType was found.
func (pctq *PartyClassificationTypeQuery) First(ctx context.Context) (*PartyClassificationType, error) {
	nodes, err := pctq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{partyclassificationtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) FirstX(ctx context.Context) *PartyClassificationType {
	node, err := pctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PartyClassificationType ID from the query.
// Returns a *NotFoundError when no PartyClassificationType ID was found.
func (pctq *PartyClassificationTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pctq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{partyclassificationtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := pctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PartyClassificationType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one PartyClassificationType entity is not found.
// Returns a *NotFoundError when no PartyClassificationType entities are found.
func (pctq *PartyClassificationTypeQuery) Only(ctx context.Context) (*PartyClassificationType, error) {
	nodes, err := pctq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{partyclassificationtype.Label}
	default:
		return nil, &NotSingularError{partyclassificationtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) OnlyX(ctx context.Context) *PartyClassificationType {
	node, err := pctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PartyClassificationType ID in the query.
// Returns a *NotSingularError when exactly one PartyClassificationType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pctq *PartyClassificationTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pctq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = &NotSingularError{partyclassificationtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := pctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PartyClassificationTypes.
func (pctq *PartyClassificationTypeQuery) All(ctx context.Context) ([]*PartyClassificationType, error) {
	if err := pctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pctq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) AllX(ctx context.Context) []*PartyClassificationType {
	nodes, err := pctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PartyClassificationType IDs.
func (pctq *PartyClassificationTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pctq.Select(partyclassificationtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := pctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pctq *PartyClassificationTypeQuery) Count(ctx context.Context) (int, error) {
	if err := pctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pctq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) CountX(ctx context.Context) int {
	count, err := pctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pctq *PartyClassificationTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pctq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pctq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pctq *PartyClassificationTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := pctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PartyClassificationTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pctq *PartyClassificationTypeQuery) Clone() *PartyClassificationTypeQuery {
	if pctq == nil {
		return nil
	}
	return &PartyClassificationTypeQuery{
		config:                            pctq.config,
		limit:                             pctq.limit,
		offset:                            pctq.offset,
		order:                             append([]OrderFunc{}, pctq.order...),
		predicates:                        append([]predicate.PartyClassificationType{}, pctq.predicates...),
		withParent:                        pctq.withParent.Clone(),
		withChildren:                      pctq.withChildren.Clone(),
		withChildPartyClassificationTypes: pctq.withChildPartyClassificationTypes.Clone(),
		// clone intermediate query.
		sql:  pctq.sql.Clone(),
		path: pctq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (pctq *PartyClassificationTypeQuery) WithParent(opts ...func(*PartyClassificationTypeQuery)) *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pctq.withParent = query
	return pctq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (pctq *PartyClassificationTypeQuery) WithChildren(opts ...func(*PartyClassificationTypeQuery)) *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pctq.withChildren = query
	return pctq
}

// WithChildPartyClassificationTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_party_classification_types" edge. The optional arguments are used to configure the query builder of the edge.
func (pctq *PartyClassificationTypeQuery) WithChildPartyClassificationTypes(opts ...func(*PartyClassificationTypeQuery)) *PartyClassificationTypeQuery {
	query := &PartyClassificationTypeQuery{config: pctq.config}
	for _, opt := range opts {
		opt(query)
	}
	pctq.withChildPartyClassificationTypes = query
	return pctq
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
//	client.PartyClassificationType.Query().
//		GroupBy(partyclassificationtype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pctq *PartyClassificationTypeQuery) GroupBy(field string, fields ...string) *PartyClassificationTypeGroupBy {
	group := &PartyClassificationTypeGroupBy{config: pctq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pctq.sqlQuery(ctx), nil
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
//	client.PartyClassificationType.Query().
//		Select(partyclassificationtype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pctq *PartyClassificationTypeQuery) Select(field string, fields ...string) *PartyClassificationTypeSelect {
	pctq.fields = append([]string{field}, fields...)
	return &PartyClassificationTypeSelect{PartyClassificationTypeQuery: pctq}
}

func (pctq *PartyClassificationTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pctq.fields {
		if !partyclassificationtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pctq.path != nil {
		prev, err := pctq.path(ctx)
		if err != nil {
			return err
		}
		pctq.sql = prev
	}
	return nil
}

func (pctq *PartyClassificationTypeQuery) sqlAll(ctx context.Context) ([]*PartyClassificationType, error) {
	var (
		nodes       = []*PartyClassificationType{}
		withFKs     = pctq.withFKs
		_spec       = pctq.querySpec()
		loadedTypes = [3]bool{
			pctq.withParent != nil,
			pctq.withChildren != nil,
			pctq.withChildPartyClassificationTypes != nil,
		}
	)
	if pctq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, partyclassificationtype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &PartyClassificationType{config: pctq.config}
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
	if err := sqlgraph.QueryNodes(ctx, pctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pctq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PartyClassificationType)
		for i := range nodes {
			if nodes[i].party_classification_type_children == nil {
				continue
			}
			fk := *nodes[i].party_classification_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(partyclassificationtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "party_classification_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := pctq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*PartyClassificationType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*PartyClassificationType{}
		}
		query.withFKs = true
		query.Where(predicate.PartyClassificationType(func(s *sql.Selector) {
			s.Where(sql.InValues(partyclassificationtype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.party_classification_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "party_classification_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "party_classification_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := pctq.withChildPartyClassificationTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*PartyClassificationType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildPartyClassificationTypes = []*PartyClassificationType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*PartyClassificationType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   partyclassificationtype.ChildPartyClassificationTypesTable,
				Columns: partyclassificationtype.ChildPartyClassificationTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(partyclassificationtype.ChildPartyClassificationTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, pctq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_party_classification_types": %w`, err)
		}
		query.Where(partyclassificationtype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_party_classification_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildPartyClassificationTypes = append(nodes[i].Edges.ChildPartyClassificationTypes, n)
			}
		}
	}

	return nodes, nil
}

func (pctq *PartyClassificationTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pctq.querySpec()
	return sqlgraph.CountNodes(ctx, pctq.driver, _spec)
}

func (pctq *PartyClassificationTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pctq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pctq *PartyClassificationTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   partyclassificationtype.Table,
			Columns: partyclassificationtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyclassificationtype.FieldID,
			},
		},
		From:   pctq.sql,
		Unique: true,
	}
	if unique := pctq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pctq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, partyclassificationtype.FieldID)
		for i := range fields {
			if fields[i] != partyclassificationtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pctq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pctq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pctq *PartyClassificationTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pctq.driver.Dialect())
	t1 := builder.Table(partyclassificationtype.Table)
	columns := pctq.fields
	if len(columns) == 0 {
		columns = partyclassificationtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pctq.sql != nil {
		selector = pctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pctq.predicates {
		p(selector)
	}
	for _, p := range pctq.order {
		p(selector)
	}
	if offset := pctq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pctq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PartyClassificationTypeGroupBy is the group-by builder for PartyClassificationType entities.
type PartyClassificationTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pctgb *PartyClassificationTypeGroupBy) Aggregate(fns ...AggregateFunc) *PartyClassificationTypeGroupBy {
	pctgb.fns = append(pctgb.fns, fns...)
	return pctgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pctgb *PartyClassificationTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pctgb.path(ctx)
	if err != nil {
		return err
	}
	pctgb.sql = query
	return pctgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pctgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pctgb.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := pctgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pctgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) StringX(ctx context.Context) string {
	v, err := pctgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pctgb.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := pctgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pctgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) IntX(ctx context.Context) int {
	v, err := pctgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pctgb.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pctgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pctgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pctgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pctgb.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pctgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pctgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pctgb *PartyClassificationTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pctgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pctgb *PartyClassificationTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := pctgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pctgb *PartyClassificationTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pctgb.fields {
		if !partyclassificationtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pctgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pctgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pctgb *PartyClassificationTypeGroupBy) sqlQuery() *sql.Selector {
	selector := pctgb.sql.Select()
	aggregation := make([]string, 0, len(pctgb.fns))
	for _, fn := range pctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pctgb.fields)+len(pctgb.fns))
		for _, f := range pctgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pctgb.fields...)...)
}

// PartyClassificationTypeSelect is the builder for selecting fields of PartyClassificationType entities.
type PartyClassificationTypeSelect struct {
	*PartyClassificationTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcts *PartyClassificationTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcts.prepareQuery(ctx); err != nil {
		return err
	}
	pcts.sql = pcts.PartyClassificationTypeQuery.sqlQuery(ctx)
	return pcts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pcts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pcts.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) StringsX(ctx context.Context) []string {
	v, err := pcts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pcts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) StringX(ctx context.Context) string {
	v, err := pcts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pcts.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) IntsX(ctx context.Context) []int {
	v, err := pcts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pcts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) IntX(ctx context.Context) int {
	v, err := pcts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pcts.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pcts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pcts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := pcts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pcts.fields) > 1 {
		return nil, errors.New("ent: PartyClassificationTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pcts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := pcts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pcts *PartyClassificationTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pcts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{partyclassificationtype.Label}
	default:
		err = fmt.Errorf("ent: PartyClassificationTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pcts *PartyClassificationTypeSelect) BoolX(ctx context.Context) bool {
	v, err := pcts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pcts *PartyClassificationTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcts.sql.Query()
	if err := pcts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}