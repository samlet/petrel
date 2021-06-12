// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeature"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappl"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappltype"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloaditem"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadstatus"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadtype"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (w *Workload) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     w.ID,
		Type:   "Workload",
		Fields: make([]*Field, 11),
		Edges:  make([]*Edge, 4),
	}
	var buf []byte
	if buf, err = json.Marshal(w.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.StatusID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "status_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.WorkloadName); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "workload_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.Description); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.LongDescription); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "long_description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.Comments); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "comments",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.WorkloadSize); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "int",
		Name:  "workload_size",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.WorkloadDate); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "workload_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.AnotherDate); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "time.Time",
		Name:  "another_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(w.AnotherText); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "string",
		Name:  "another_text",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WorkloadType",
		Name: "workload_type",
	}
	node.Edges[0].IDs, err = w.QueryWorkloadType().
		Select(workloadtype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "WorkloadFeatureAppl",
		Name: "workload_feature_appls",
	}
	node.Edges[1].IDs, err = w.QueryWorkloadFeatureAppls().
		Select(workloadfeatureappl.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "WorkloadItem",
		Name: "workload_items",
	}
	node.Edges[2].IDs, err = w.QueryWorkloadItems().
		Select(workloaditem.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "WorkloadStatus",
		Name: "workload_statuses",
	}
	node.Edges[3].IDs, err = w.QueryWorkloadStatuses().
		Select(workloadstatus.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (wf *WorkloadFeature) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     wf.ID,
		Type:   "WorkloadFeature",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(wf.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wf.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wf.FeatureSourceEnumID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "feature_source_enum_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wf.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WorkloadFeatureAppl",
		Name: "workload_feature_appls",
	}
	node.Edges[0].IDs, err = wf.QueryWorkloadFeatureAppls().
		Select(workloadfeatureappl.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (wfa *WorkloadFeatureAppl) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     wfa.ID,
		Type:   "WorkloadFeatureAppl",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(wfa.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfa.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfa.FromDate); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "from_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfa.ThruDate); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "thru_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfa.SequenceNum); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "sequence_num",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workload",
		Name: "workload",
	}
	node.Edges[0].IDs, err = wfa.QueryWorkload().
		Select(workload.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "WorkloadFeature",
		Name: "workload_feature",
	}
	node.Edges[1].IDs, err = wfa.QueryWorkloadFeature().
		Select(workloadfeature.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "WorkloadFeatureApplType",
		Name: "workload_feature_appl_type",
	}
	node.Edges[2].IDs, err = wfa.QueryWorkloadFeatureApplType().
		Select(workloadfeatureappltype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (wfat *WorkloadFeatureApplType) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     wfat.ID,
		Type:   "WorkloadFeatureApplType",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(wfat.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfat.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wfat.Description); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WorkloadFeatureApplType",
		Name: "parent",
	}
	node.Edges[0].IDs, err = wfat.QueryParent().
		Select(workloadfeatureappltype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "WorkloadFeatureApplType",
		Name: "children",
	}
	node.Edges[1].IDs, err = wfat.QueryChildren().
		Select(workloadfeatureappltype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "WorkloadFeatureAppl",
		Name: "workload_feature_appls",
	}
	node.Edges[2].IDs, err = wfat.QueryWorkloadFeatureAppls().
		Select(workloadfeatureappl.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (wi *WorkloadItem) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     wi.ID,
		Type:   "WorkloadItem",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(wi.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wi.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wi.WorkloadItemSeqID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "workload_item_seq_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wi.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wi.Amount); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "float64",
		Name:  "amount",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wi.AmountUomID); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "amount_uom_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workload",
		Name: "workload",
	}
	node.Edges[0].IDs, err = wi.QueryWorkload().
		Select(workload.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (ws *WorkloadStatus) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ws.ID,
		Type:   "WorkloadStatus",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ws.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.StatusDate); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "status_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.StatusEndDate); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "status_end_date",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.ChangeByUserLoginID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "change_by_user_login_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ws.StatusID); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "int",
		Name:  "status_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Workload",
		Name: "workload",
	}
	node.Edges[0].IDs, err = ws.QueryWorkload().
		Select(workload.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (wt *WorkloadType) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     wt.ID,
		Type:   "WorkloadType",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(wt.CreateTime); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "create_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wt.UpdateTime); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "update_time",
		Value: string(buf),
	}
	if buf, err = json.Marshal(wt.Description); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "WorkloadType",
		Name: "parent",
	}
	node.Edges[0].IDs, err = wt.QueryParent().
		Select(workloadtype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "WorkloadType",
		Name: "children",
	}
	node.Edges[1].IDs, err = wt.QueryChildren().
		Select(workloadtype.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Workload",
		Name: "workloads",
	}
	node.Edges[2].IDs, err = wt.QueryWorkloads().
		Select(workload.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//		c.Noder(ctx, id)
//		c.Noder(ctx, id, ent.WithNodeType(pet.Table))
//
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case workload.Table:
		n, err := c.Workload.Query().
			Where(workload.ID(id)).
			CollectFields(ctx, "Workload").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloadfeature.Table:
		n, err := c.WorkloadFeature.Query().
			Where(workloadfeature.ID(id)).
			CollectFields(ctx, "WorkloadFeature").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloadfeatureappl.Table:
		n, err := c.WorkloadFeatureAppl.Query().
			Where(workloadfeatureappl.ID(id)).
			CollectFields(ctx, "WorkloadFeatureAppl").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloadfeatureappltype.Table:
		n, err := c.WorkloadFeatureApplType.Query().
			Where(workloadfeatureappltype.ID(id)).
			CollectFields(ctx, "WorkloadFeatureApplType").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloaditem.Table:
		n, err := c.WorkloadItem.Query().
			Where(workloaditem.ID(id)).
			CollectFields(ctx, "WorkloadItem").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloadstatus.Table:
		n, err := c.WorkloadStatus.Query().
			Where(workloadstatus.ID(id)).
			CollectFields(ctx, "WorkloadStatus").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case workloadtype.Table:
		n, err := c.WorkloadType.Query().
			Where(workloadtype.ID(id)).
			CollectFields(ctx, "WorkloadType").
			Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case workload.Table:
		nodes, err := c.Workload.Query().
			Where(workload.IDIn(ids...)).
			CollectFields(ctx, "Workload").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloadfeature.Table:
		nodes, err := c.WorkloadFeature.Query().
			Where(workloadfeature.IDIn(ids...)).
			CollectFields(ctx, "WorkloadFeature").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloadfeatureappl.Table:
		nodes, err := c.WorkloadFeatureAppl.Query().
			Where(workloadfeatureappl.IDIn(ids...)).
			CollectFields(ctx, "WorkloadFeatureAppl").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloadfeatureappltype.Table:
		nodes, err := c.WorkloadFeatureApplType.Query().
			Where(workloadfeatureappltype.IDIn(ids...)).
			CollectFields(ctx, "WorkloadFeatureApplType").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloaditem.Table:
		nodes, err := c.WorkloadItem.Query().
			Where(workloaditem.IDIn(ids...)).
			CollectFields(ctx, "WorkloadItem").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloadstatus.Table:
		nodes, err := c.WorkloadStatus.Query().
			Where(workloadstatus.IDIn(ids...)).
			CollectFields(ctx, "WorkloadStatus").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case workloadtype.Table:
		nodes, err := c.WorkloadType.Query().
			Where(workloadtype.IDIn(ids...)).
			CollectFields(ctx, "WorkloadType").
			All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
