// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/samlet/petrel/alfin/modules/asset/ent/asset"
	"github.com/samlet/petrel/alfin/modules/asset/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAsset       = "Asset"
	TypeWorkloadPkg = "WorkloadPkg"
)

// AssetMutation represents an operation that mutates the Asset nodes in the graph.
type AssetMutation struct {
	config
	op            Op
	typ           string
	id            *int
	create_time   *time.Time
	update_time   *time.Time
	string_ref    *string
	model         *string
	registered_at *time.Time
	clearedFields map[string]struct{}
	pkg           *int
	clearedpkg    bool
	done          bool
	oldValue      func(context.Context) (*Asset, error)
	predicates    []predicate.Asset
}

var _ ent.Mutation = (*AssetMutation)(nil)

// assetOption allows management of the mutation configuration using functional options.
type assetOption func(*AssetMutation)

// newAssetMutation creates new mutation for the Asset entity.
func newAssetMutation(c config, op Op, opts ...assetOption) *AssetMutation {
	m := &AssetMutation{
		config:        c,
		op:            op,
		typ:           TypeAsset,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAssetID sets the ID field of the mutation.
func withAssetID(id int) assetOption {
	return func(m *AssetMutation) {
		var (
			err   error
			once  sync.Once
			value *Asset
		)
		m.oldValue = func(ctx context.Context) (*Asset, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Asset.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAsset sets the old Asset of the mutation.
func withAsset(node *Asset) assetOption {
	return func(m *AssetMutation) {
		m.oldValue = func(context.Context) (*Asset, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AssetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AssetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AssetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreateTime sets the "create_time" field.
func (m *AssetMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *AssetMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Asset entity.
// If the Asset object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssetMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *AssetMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *AssetMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *AssetMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Asset entity.
// If the Asset object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssetMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *AssetMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetStringRef sets the "string_ref" field.
func (m *AssetMutation) SetStringRef(s string) {
	m.string_ref = &s
}

// StringRef returns the value of the "string_ref" field in the mutation.
func (m *AssetMutation) StringRef() (r string, exists bool) {
	v := m.string_ref
	if v == nil {
		return
	}
	return *v, true
}

// OldStringRef returns the old "string_ref" field's value of the Asset entity.
// If the Asset object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssetMutation) OldStringRef(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldStringRef is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldStringRef requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStringRef: %w", err)
	}
	return oldValue.StringRef, nil
}

// ClearStringRef clears the value of the "string_ref" field.
func (m *AssetMutation) ClearStringRef() {
	m.string_ref = nil
	m.clearedFields[asset.FieldStringRef] = struct{}{}
}

// StringRefCleared returns if the "string_ref" field was cleared in this mutation.
func (m *AssetMutation) StringRefCleared() bool {
	_, ok := m.clearedFields[asset.FieldStringRef]
	return ok
}

// ResetStringRef resets all changes to the "string_ref" field.
func (m *AssetMutation) ResetStringRef() {
	m.string_ref = nil
	delete(m.clearedFields, asset.FieldStringRef)
}

// SetModel sets the "model" field.
func (m *AssetMutation) SetModel(s string) {
	m.model = &s
}

// Model returns the value of the "model" field in the mutation.
func (m *AssetMutation) Model() (r string, exists bool) {
	v := m.model
	if v == nil {
		return
	}
	return *v, true
}

// OldModel returns the old "model" field's value of the Asset entity.
// If the Asset object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssetMutation) OldModel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldModel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldModel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldModel: %w", err)
	}
	return oldValue.Model, nil
}

// ResetModel resets all changes to the "model" field.
func (m *AssetMutation) ResetModel() {
	m.model = nil
}

// SetRegisteredAt sets the "registered_at" field.
func (m *AssetMutation) SetRegisteredAt(t time.Time) {
	m.registered_at = &t
}

// RegisteredAt returns the value of the "registered_at" field in the mutation.
func (m *AssetMutation) RegisteredAt() (r time.Time, exists bool) {
	v := m.registered_at
	if v == nil {
		return
	}
	return *v, true
}

// OldRegisteredAt returns the old "registered_at" field's value of the Asset entity.
// If the Asset object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AssetMutation) OldRegisteredAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRegisteredAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRegisteredAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRegisteredAt: %w", err)
	}
	return oldValue.RegisteredAt, nil
}

// ResetRegisteredAt resets all changes to the "registered_at" field.
func (m *AssetMutation) ResetRegisteredAt() {
	m.registered_at = nil
}

// SetPkgID sets the "pkg" edge to the WorkloadPkg entity by id.
func (m *AssetMutation) SetPkgID(id int) {
	m.pkg = &id
}

// ClearPkg clears the "pkg" edge to the WorkloadPkg entity.
func (m *AssetMutation) ClearPkg() {
	m.clearedpkg = true
}

// PkgCleared reports if the "pkg" edge to the WorkloadPkg entity was cleared.
func (m *AssetMutation) PkgCleared() bool {
	return m.clearedpkg
}

// PkgID returns the "pkg" edge ID in the mutation.
func (m *AssetMutation) PkgID() (id int, exists bool) {
	if m.pkg != nil {
		return *m.pkg, true
	}
	return
}

// PkgIDs returns the "pkg" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// PkgID instead. It exists only for internal usage by the builders.
func (m *AssetMutation) PkgIDs() (ids []int) {
	if id := m.pkg; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPkg resets all changes to the "pkg" edge.
func (m *AssetMutation) ResetPkg() {
	m.pkg = nil
	m.clearedpkg = false
}

// Op returns the operation name.
func (m *AssetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Asset).
func (m *AssetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AssetMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.create_time != nil {
		fields = append(fields, asset.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, asset.FieldUpdateTime)
	}
	if m.string_ref != nil {
		fields = append(fields, asset.FieldStringRef)
	}
	if m.model != nil {
		fields = append(fields, asset.FieldModel)
	}
	if m.registered_at != nil {
		fields = append(fields, asset.FieldRegisteredAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AssetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case asset.FieldCreateTime:
		return m.CreateTime()
	case asset.FieldUpdateTime:
		return m.UpdateTime()
	case asset.FieldStringRef:
		return m.StringRef()
	case asset.FieldModel:
		return m.Model()
	case asset.FieldRegisteredAt:
		return m.RegisteredAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AssetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case asset.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case asset.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case asset.FieldStringRef:
		return m.OldStringRef(ctx)
	case asset.FieldModel:
		return m.OldModel(ctx)
	case asset.FieldRegisteredAt:
		return m.OldRegisteredAt(ctx)
	}
	return nil, fmt.Errorf("unknown Asset field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AssetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case asset.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case asset.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case asset.FieldStringRef:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStringRef(v)
		return nil
	case asset.FieldModel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetModel(v)
		return nil
	case asset.FieldRegisteredAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRegisteredAt(v)
		return nil
	}
	return fmt.Errorf("unknown Asset field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AssetMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AssetMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AssetMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Asset numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AssetMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(asset.FieldStringRef) {
		fields = append(fields, asset.FieldStringRef)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AssetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AssetMutation) ClearField(name string) error {
	switch name {
	case asset.FieldStringRef:
		m.ClearStringRef()
		return nil
	}
	return fmt.Errorf("unknown Asset nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AssetMutation) ResetField(name string) error {
	switch name {
	case asset.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case asset.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case asset.FieldStringRef:
		m.ResetStringRef()
		return nil
	case asset.FieldModel:
		m.ResetModel()
		return nil
	case asset.FieldRegisteredAt:
		m.ResetRegisteredAt()
		return nil
	}
	return fmt.Errorf("unknown Asset field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AssetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.pkg != nil {
		edges = append(edges, asset.EdgePkg)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AssetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case asset.EdgePkg:
		if id := m.pkg; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AssetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AssetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AssetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpkg {
		edges = append(edges, asset.EdgePkg)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AssetMutation) EdgeCleared(name string) bool {
	switch name {
	case asset.EdgePkg:
		return m.clearedpkg
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AssetMutation) ClearEdge(name string) error {
	switch name {
	case asset.EdgePkg:
		m.ClearPkg()
		return nil
	}
	return fmt.Errorf("unknown Asset unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AssetMutation) ResetEdge(name string) error {
	switch name {
	case asset.EdgePkg:
		m.ResetPkg()
		return nil
	}
	return fmt.Errorf("unknown Asset edge %s", name)
}

// WorkloadPkgMutation represents an operation that mutates the WorkloadPkg nodes in the graph.
type WorkloadPkgMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	assets        map[int]struct{}
	removedassets map[int]struct{}
	clearedassets bool
	done          bool
	oldValue      func(context.Context) (*WorkloadPkg, error)
	predicates    []predicate.WorkloadPkg
}

var _ ent.Mutation = (*WorkloadPkgMutation)(nil)

// workloadpkgOption allows management of the mutation configuration using functional options.
type workloadpkgOption func(*WorkloadPkgMutation)

// newWorkloadPkgMutation creates new mutation for the WorkloadPkg entity.
func newWorkloadPkgMutation(c config, op Op, opts ...workloadpkgOption) *WorkloadPkgMutation {
	m := &WorkloadPkgMutation{
		config:        c,
		op:            op,
		typ:           TypeWorkloadPkg,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWorkloadPkgID sets the ID field of the mutation.
func withWorkloadPkgID(id int) workloadpkgOption {
	return func(m *WorkloadPkgMutation) {
		var (
			err   error
			once  sync.Once
			value *WorkloadPkg
		)
		m.oldValue = func(ctx context.Context) (*WorkloadPkg, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().WorkloadPkg.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWorkloadPkg sets the old WorkloadPkg of the mutation.
func withWorkloadPkg(node *WorkloadPkg) workloadpkgOption {
	return func(m *WorkloadPkgMutation) {
		m.oldValue = func(context.Context) (*WorkloadPkg, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m WorkloadPkgMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m WorkloadPkgMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *WorkloadPkgMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *WorkloadPkgMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *WorkloadPkgMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the WorkloadPkg entity.
// If the WorkloadPkg object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *WorkloadPkgMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *WorkloadPkgMutation) ResetName() {
	m.name = nil
}

// AddAssetIDs adds the "assets" edge to the Asset entity by ids.
func (m *WorkloadPkgMutation) AddAssetIDs(ids ...int) {
	if m.assets == nil {
		m.assets = make(map[int]struct{})
	}
	for i := range ids {
		m.assets[ids[i]] = struct{}{}
	}
}

// ClearAssets clears the "assets" edge to the Asset entity.
func (m *WorkloadPkgMutation) ClearAssets() {
	m.clearedassets = true
}

// AssetsCleared reports if the "assets" edge to the Asset entity was cleared.
func (m *WorkloadPkgMutation) AssetsCleared() bool {
	return m.clearedassets
}

// RemoveAssetIDs removes the "assets" edge to the Asset entity by IDs.
func (m *WorkloadPkgMutation) RemoveAssetIDs(ids ...int) {
	if m.removedassets == nil {
		m.removedassets = make(map[int]struct{})
	}
	for i := range ids {
		m.removedassets[ids[i]] = struct{}{}
	}
}

// RemovedAssets returns the removed IDs of the "assets" edge to the Asset entity.
func (m *WorkloadPkgMutation) RemovedAssetsIDs() (ids []int) {
	for id := range m.removedassets {
		ids = append(ids, id)
	}
	return
}

// AssetsIDs returns the "assets" edge IDs in the mutation.
func (m *WorkloadPkgMutation) AssetsIDs() (ids []int) {
	for id := range m.assets {
		ids = append(ids, id)
	}
	return
}

// ResetAssets resets all changes to the "assets" edge.
func (m *WorkloadPkgMutation) ResetAssets() {
	m.assets = nil
	m.clearedassets = false
	m.removedassets = nil
}

// Op returns the operation name.
func (m *WorkloadPkgMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (WorkloadPkg).
func (m *WorkloadPkgMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *WorkloadPkgMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, workloadpkg.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *WorkloadPkgMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case workloadpkg.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *WorkloadPkgMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case workloadpkg.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown WorkloadPkg field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkloadPkgMutation) SetField(name string, value ent.Value) error {
	switch name {
	case workloadpkg.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown WorkloadPkg field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *WorkloadPkgMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *WorkloadPkgMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *WorkloadPkgMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown WorkloadPkg numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *WorkloadPkgMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *WorkloadPkgMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *WorkloadPkgMutation) ClearField(name string) error {
	return fmt.Errorf("unknown WorkloadPkg nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *WorkloadPkgMutation) ResetField(name string) error {
	switch name {
	case workloadpkg.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown WorkloadPkg field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *WorkloadPkgMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.assets != nil {
		edges = append(edges, workloadpkg.EdgeAssets)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *WorkloadPkgMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case workloadpkg.EdgeAssets:
		ids := make([]ent.Value, 0, len(m.assets))
		for id := range m.assets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *WorkloadPkgMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedassets != nil {
		edges = append(edges, workloadpkg.EdgeAssets)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *WorkloadPkgMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case workloadpkg.EdgeAssets:
		ids := make([]ent.Value, 0, len(m.removedassets))
		for id := range m.removedassets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *WorkloadPkgMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedassets {
		edges = append(edges, workloadpkg.EdgeAssets)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *WorkloadPkgMutation) EdgeCleared(name string) bool {
	switch name {
	case workloadpkg.EdgeAssets:
		return m.clearedassets
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *WorkloadPkgMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown WorkloadPkg unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *WorkloadPkgMutation) ResetEdge(name string) error {
	switch name {
	case workloadpkg.EdgeAssets:
		m.ResetAssets()
		return nil
	}
	return fmt.Errorf("unknown WorkloadPkg edge %s", name)
}
