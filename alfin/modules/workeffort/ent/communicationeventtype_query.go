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
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// CommunicationEventTypeQuery is the builder for querying CommunicationEventType entities.
type CommunicationEventTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CommunicationEventType
	// eager-loading edges.
	withParent                        *CommunicationEventTypeQuery
	withChildren                      *CommunicationEventTypeQuery
	withContacMechTypeContactMechType *ContactMechTypeQuery
	withChildCommunicationEventTypes  *CommunicationEventTypeQuery
	withFKs                           bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommunicationEventTypeQuery builder.
func (cetq *CommunicationEventTypeQuery) Where(ps ...predicate.CommunicationEventType) *CommunicationEventTypeQuery {
	cetq.predicates = append(cetq.predicates, ps...)
	return cetq
}

// Limit adds a limit step to the query.
func (cetq *CommunicationEventTypeQuery) Limit(limit int) *CommunicationEventTypeQuery {
	cetq.limit = &limit
	return cetq
}

// Offset adds an offset step to the query.
func (cetq *CommunicationEventTypeQuery) Offset(offset int) *CommunicationEventTypeQuery {
	cetq.offset = &offset
	return cetq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cetq *CommunicationEventTypeQuery) Unique(unique bool) *CommunicationEventTypeQuery {
	cetq.unique = &unique
	return cetq
}

// Order adds an order step to the query.
func (cetq *CommunicationEventTypeQuery) Order(o ...OrderFunc) *CommunicationEventTypeQuery {
	cetq.order = append(cetq.order, o...)
	return cetq
}

// QueryParent chains the current query on the "parent" edge.
func (cetq *CommunicationEventTypeQuery) QueryParent() *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(communicationeventtype.Table, communicationeventtype.FieldID, selector),
			sqlgraph.To(communicationeventtype.Table, communicationeventtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, communicationeventtype.ParentTable, communicationeventtype.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (cetq *CommunicationEventTypeQuery) QueryChildren() *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(communicationeventtype.Table, communicationeventtype.FieldID, selector),
			sqlgraph.To(communicationeventtype.Table, communicationeventtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, communicationeventtype.ChildrenTable, communicationeventtype.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(cetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContacMechTypeContactMechType chains the current query on the "contac_mech_type_contact_mech_type" edge.
func (cetq *CommunicationEventTypeQuery) QueryContacMechTypeContactMechType() *ContactMechTypeQuery {
	query := &ContactMechTypeQuery{config: cetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(communicationeventtype.Table, communicationeventtype.FieldID, selector),
			sqlgraph.To(contactmechtype.Table, contactmechtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, communicationeventtype.ContacMechTypeContactMechTypeTable, communicationeventtype.ContacMechTypeContactMechTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(cetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildCommunicationEventTypes chains the current query on the "child_communication_event_types" edge.
func (cetq *CommunicationEventTypeQuery) QueryChildCommunicationEventTypes() *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cetq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(communicationeventtype.Table, communicationeventtype.FieldID, selector),
			sqlgraph.To(communicationeventtype.Table, communicationeventtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, communicationeventtype.ChildCommunicationEventTypesTable, communicationeventtype.ChildCommunicationEventTypesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cetq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommunicationEventType entity from the query.
// Returns a *NotFoundError when no CommunicationEventType was found.
func (cetq *CommunicationEventTypeQuery) First(ctx context.Context) (*CommunicationEventType, error) {
	nodes, err := cetq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{communicationeventtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) FirstX(ctx context.Context) *CommunicationEventType {
	node, err := cetq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommunicationEventType ID from the query.
// Returns a *NotFoundError when no CommunicationEventType ID was found.
func (cetq *CommunicationEventTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cetq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{communicationeventtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := cetq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommunicationEventType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one CommunicationEventType entity is not found.
// Returns a *NotFoundError when no CommunicationEventType entities are found.
func (cetq *CommunicationEventTypeQuery) Only(ctx context.Context) (*CommunicationEventType, error) {
	nodes, err := cetq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{communicationeventtype.Label}
	default:
		return nil, &NotSingularError{communicationeventtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) OnlyX(ctx context.Context) *CommunicationEventType {
	node, err := cetq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommunicationEventType ID in the query.
// Returns a *NotSingularError when exactly one CommunicationEventType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cetq *CommunicationEventTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cetq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = &NotSingularError{communicationeventtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := cetq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommunicationEventTypes.
func (cetq *CommunicationEventTypeQuery) All(ctx context.Context) ([]*CommunicationEventType, error) {
	if err := cetq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cetq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) AllX(ctx context.Context) []*CommunicationEventType {
	nodes, err := cetq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommunicationEventType IDs.
func (cetq *CommunicationEventTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cetq.Select(communicationeventtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := cetq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cetq *CommunicationEventTypeQuery) Count(ctx context.Context) (int, error) {
	if err := cetq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cetq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) CountX(ctx context.Context) int {
	count, err := cetq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cetq *CommunicationEventTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := cetq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cetq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cetq *CommunicationEventTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := cetq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommunicationEventTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cetq *CommunicationEventTypeQuery) Clone() *CommunicationEventTypeQuery {
	if cetq == nil {
		return nil
	}
	return &CommunicationEventTypeQuery{
		config:                            cetq.config,
		limit:                             cetq.limit,
		offset:                            cetq.offset,
		order:                             append([]OrderFunc{}, cetq.order...),
		predicates:                        append([]predicate.CommunicationEventType{}, cetq.predicates...),
		withParent:                        cetq.withParent.Clone(),
		withChildren:                      cetq.withChildren.Clone(),
		withContacMechTypeContactMechType: cetq.withContacMechTypeContactMechType.Clone(),
		withChildCommunicationEventTypes:  cetq.withChildCommunicationEventTypes.Clone(),
		// clone intermediate query.
		sql:  cetq.sql.Clone(),
		path: cetq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (cetq *CommunicationEventTypeQuery) WithParent(opts ...func(*CommunicationEventTypeQuery)) *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	for _, opt := range opts {
		opt(query)
	}
	cetq.withParent = query
	return cetq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (cetq *CommunicationEventTypeQuery) WithChildren(opts ...func(*CommunicationEventTypeQuery)) *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	for _, opt := range opts {
		opt(query)
	}
	cetq.withChildren = query
	return cetq
}

// WithContacMechTypeContactMechType tells the query-builder to eager-load the nodes that are connected to
// the "contac_mech_type_contact_mech_type" edge. The optional arguments are used to configure the query builder of the edge.
func (cetq *CommunicationEventTypeQuery) WithContacMechTypeContactMechType(opts ...func(*ContactMechTypeQuery)) *CommunicationEventTypeQuery {
	query := &ContactMechTypeQuery{config: cetq.config}
	for _, opt := range opts {
		opt(query)
	}
	cetq.withContacMechTypeContactMechType = query
	return cetq
}

// WithChildCommunicationEventTypes tells the query-builder to eager-load the nodes that are connected to
// the "child_communication_event_types" edge. The optional arguments are used to configure the query builder of the edge.
func (cetq *CommunicationEventTypeQuery) WithChildCommunicationEventTypes(opts ...func(*CommunicationEventTypeQuery)) *CommunicationEventTypeQuery {
	query := &CommunicationEventTypeQuery{config: cetq.config}
	for _, opt := range opts {
		opt(query)
	}
	cetq.withChildCommunicationEventTypes = query
	return cetq
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
//	client.CommunicationEventType.Query().
//		GroupBy(communicationeventtype.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cetq *CommunicationEventTypeQuery) GroupBy(field string, fields ...string) *CommunicationEventTypeGroupBy {
	group := &CommunicationEventTypeGroupBy{config: cetq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cetq.sqlQuery(ctx), nil
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
//	client.CommunicationEventType.Query().
//		Select(communicationeventtype.FieldCreateTime).
//		Scan(ctx, &v)
//
func (cetq *CommunicationEventTypeQuery) Select(field string, fields ...string) *CommunicationEventTypeSelect {
	cetq.fields = append([]string{field}, fields...)
	return &CommunicationEventTypeSelect{CommunicationEventTypeQuery: cetq}
}

func (cetq *CommunicationEventTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cetq.fields {
		if !communicationeventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cetq.path != nil {
		prev, err := cetq.path(ctx)
		if err != nil {
			return err
		}
		cetq.sql = prev
	}
	return nil
}

func (cetq *CommunicationEventTypeQuery) sqlAll(ctx context.Context) ([]*CommunicationEventType, error) {
	var (
		nodes       = []*CommunicationEventType{}
		withFKs     = cetq.withFKs
		_spec       = cetq.querySpec()
		loadedTypes = [4]bool{
			cetq.withParent != nil,
			cetq.withChildren != nil,
			cetq.withContacMechTypeContactMechType != nil,
			cetq.withChildCommunicationEventTypes != nil,
		}
	)
	if cetq.withParent != nil || cetq.withContacMechTypeContactMechType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, communicationeventtype.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CommunicationEventType{config: cetq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cetq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cetq.withParent; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CommunicationEventType)
		for i := range nodes {
			if nodes[i].communication_event_type_children == nil {
				continue
			}
			fk := *nodes[i].communication_event_type_children
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(communicationeventtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "communication_event_type_children" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := cetq.withChildren; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CommunicationEventType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children = []*CommunicationEventType{}
		}
		query.withFKs = true
		query.Where(predicate.CommunicationEventType(func(s *sql.Selector) {
			s.Where(sql.InValues(communicationeventtype.ChildrenColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.communication_event_type_children
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "communication_event_type_children" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "communication_event_type_children" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children = append(node.Edges.Children, n)
		}
	}

	if query := cetq.withContacMechTypeContactMechType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CommunicationEventType)
		for i := range nodes {
			if nodes[i].contact_mech_type_contac_mech_type_communication_event_types == nil {
				continue
			}
			fk := *nodes[i].contact_mech_type_contac_mech_type_communication_event_types
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(contactmechtype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "contact_mech_type_contac_mech_type_communication_event_types" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ContacMechTypeContactMechType = n
			}
		}
	}

	if query := cetq.withChildCommunicationEventTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*CommunicationEventType, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.ChildCommunicationEventTypes = []*CommunicationEventType{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*CommunicationEventType)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   communicationeventtype.ChildCommunicationEventTypesTable,
				Columns: communicationeventtype.ChildCommunicationEventTypesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(communicationeventtype.ChildCommunicationEventTypesPrimaryKey[0], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, cetq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "child_communication_event_types": %w`, err)
		}
		query.Where(communicationeventtype.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "child_communication_event_types" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ChildCommunicationEventTypes = append(nodes[i].Edges.ChildCommunicationEventTypes, n)
			}
		}
	}

	return nodes, nil
}

func (cetq *CommunicationEventTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cetq.querySpec()
	return sqlgraph.CountNodes(ctx, cetq.driver, _spec)
}

func (cetq *CommunicationEventTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cetq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cetq *CommunicationEventTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   communicationeventtype.Table,
			Columns: communicationeventtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: communicationeventtype.FieldID,
			},
		},
		From:   cetq.sql,
		Unique: true,
	}
	if unique := cetq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cetq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, communicationeventtype.FieldID)
		for i := range fields {
			if fields[i] != communicationeventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cetq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cetq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cetq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cetq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cetq *CommunicationEventTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cetq.driver.Dialect())
	t1 := builder.Table(communicationeventtype.Table)
	columns := cetq.fields
	if len(columns) == 0 {
		columns = communicationeventtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cetq.sql != nil {
		selector = cetq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range cetq.predicates {
		p(selector)
	}
	for _, p := range cetq.order {
		p(selector)
	}
	if offset := cetq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cetq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommunicationEventTypeGroupBy is the group-by builder for CommunicationEventType entities.
type CommunicationEventTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cetgb *CommunicationEventTypeGroupBy) Aggregate(fns ...AggregateFunc) *CommunicationEventTypeGroupBy {
	cetgb.fns = append(cetgb.fns, fns...)
	return cetgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cetgb *CommunicationEventTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cetgb.path(ctx)
	if err != nil {
		return err
	}
	cetgb.sql = query
	return cetgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cetgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cetgb.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := cetgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cetgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) StringX(ctx context.Context) string {
	v, err := cetgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cetgb.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := cetgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cetgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) IntX(ctx context.Context) int {
	v, err := cetgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cetgb.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cetgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cetgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cetgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cetgb.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cetgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cetgb *CommunicationEventTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cetgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cetgb *CommunicationEventTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := cetgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cetgb *CommunicationEventTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cetgb.fields {
		if !communicationeventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cetgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cetgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cetgb *CommunicationEventTypeGroupBy) sqlQuery() *sql.Selector {
	selector := cetgb.sql.Select()
	aggregation := make([]string, 0, len(cetgb.fns))
	for _, fn := range cetgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cetgb.fields)+len(cetgb.fns))
		for _, f := range cetgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cetgb.fields...)...)
}

// CommunicationEventTypeSelect is the builder for selecting fields of CommunicationEventType entities.
type CommunicationEventTypeSelect struct {
	*CommunicationEventTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cets *CommunicationEventTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cets.prepareQuery(ctx); err != nil {
		return err
	}
	cets.sql = cets.CommunicationEventTypeQuery.sqlQuery(ctx)
	return cets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cets.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) StringsX(ctx context.Context) []string {
	v, err := cets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cets.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) StringX(ctx context.Context) string {
	v, err := cets.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cets.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) IntsX(ctx context.Context) []int {
	v, err := cets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cets.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) IntX(ctx context.Context) int {
	v, err := cets.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cets.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cets.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := cets.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cets.fields) > 1 {
		return nil, errors.New("ent: CommunicationEventTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := cets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cets *CommunicationEventTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cets.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{communicationeventtype.Label}
	default:
		err = fmt.Errorf("ent: CommunicationEventTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cets *CommunicationEventTypeSelect) BoolX(ctx context.Context) bool {
	v, err := cets.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cets *CommunicationEventTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cets.sql.Query()
	if err := cets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}