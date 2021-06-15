// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/goodidentificationtype"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// GoodIdentificationTypeUpdate is the builder for updating GoodIdentificationType entities.
type GoodIdentificationTypeUpdate struct {
	config
	hooks    []Hook
	mutation *GoodIdentificationTypeMutation
}

// Where adds a new predicate for the GoodIdentificationTypeUpdate builder.
func (gitu *GoodIdentificationTypeUpdate) Where(ps ...predicate.GoodIdentificationType) *GoodIdentificationTypeUpdate {
	gitu.mutation.predicates = append(gitu.mutation.predicates, ps...)
	return gitu
}

// SetStringRef sets the "string_ref" field.
func (gitu *GoodIdentificationTypeUpdate) SetStringRef(s string) *GoodIdentificationTypeUpdate {
	gitu.mutation.SetStringRef(s)
	return gitu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (gitu *GoodIdentificationTypeUpdate) SetNillableStringRef(s *string) *GoodIdentificationTypeUpdate {
	if s != nil {
		gitu.SetStringRef(*s)
	}
	return gitu
}

// ClearStringRef clears the value of the "string_ref" field.
func (gitu *GoodIdentificationTypeUpdate) ClearStringRef() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearStringRef()
	return gitu
}

// SetHasTable sets the "has_table" field.
func (gitu *GoodIdentificationTypeUpdate) SetHasTable(gt goodidentificationtype.HasTable) *GoodIdentificationTypeUpdate {
	gitu.mutation.SetHasTable(gt)
	return gitu
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (gitu *GoodIdentificationTypeUpdate) SetNillableHasTable(gt *goodidentificationtype.HasTable) *GoodIdentificationTypeUpdate {
	if gt != nil {
		gitu.SetHasTable(*gt)
	}
	return gitu
}

// ClearHasTable clears the value of the "has_table" field.
func (gitu *GoodIdentificationTypeUpdate) ClearHasTable() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearHasTable()
	return gitu
}

// SetDescription sets the "description" field.
func (gitu *GoodIdentificationTypeUpdate) SetDescription(s string) *GoodIdentificationTypeUpdate {
	gitu.mutation.SetDescription(s)
	return gitu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gitu *GoodIdentificationTypeUpdate) SetNillableDescription(s *string) *GoodIdentificationTypeUpdate {
	if s != nil {
		gitu.SetDescription(*s)
	}
	return gitu
}

// ClearDescription clears the value of the "description" field.
func (gitu *GoodIdentificationTypeUpdate) ClearDescription() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearDescription()
	return gitu
}

// SetParentID sets the "parent" edge to the GoodIdentificationType entity by ID.
func (gitu *GoodIdentificationTypeUpdate) SetParentID(id int) *GoodIdentificationTypeUpdate {
	gitu.mutation.SetParentID(id)
	return gitu
}

// SetNillableParentID sets the "parent" edge to the GoodIdentificationType entity by ID if the given value is not nil.
func (gitu *GoodIdentificationTypeUpdate) SetNillableParentID(id *int) *GoodIdentificationTypeUpdate {
	if id != nil {
		gitu = gitu.SetParentID(*id)
	}
	return gitu
}

// SetParent sets the "parent" edge to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) SetParent(g *GoodIdentificationType) *GoodIdentificationTypeUpdate {
	return gitu.SetParentID(g.ID)
}

// AddChildIDs adds the "children" edge to the GoodIdentificationType entity by IDs.
func (gitu *GoodIdentificationTypeUpdate) AddChildIDs(ids ...int) *GoodIdentificationTypeUpdate {
	gitu.mutation.AddChildIDs(ids...)
	return gitu
}

// AddChildren adds the "children" edges to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) AddChildren(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gitu.AddChildIDs(ids...)
}

// AddChildGoodIdentificationTypeIDs adds the "child_good_identification_types" edge to the GoodIdentificationType entity by IDs.
func (gitu *GoodIdentificationTypeUpdate) AddChildGoodIdentificationTypeIDs(ids ...int) *GoodIdentificationTypeUpdate {
	gitu.mutation.AddChildGoodIdentificationTypeIDs(ids...)
	return gitu
}

// AddChildGoodIdentificationTypes adds the "child_good_identification_types" edges to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) AddChildGoodIdentificationTypes(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gitu.AddChildGoodIdentificationTypeIDs(ids...)
}

// Mutation returns the GoodIdentificationTypeMutation object of the builder.
func (gitu *GoodIdentificationTypeUpdate) Mutation() *GoodIdentificationTypeMutation {
	return gitu.mutation
}

// ClearParent clears the "parent" edge to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) ClearParent() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearParent()
	return gitu
}

// ClearChildren clears all "children" edges to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) ClearChildren() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearChildren()
	return gitu
}

// RemoveChildIDs removes the "children" edge to GoodIdentificationType entities by IDs.
func (gitu *GoodIdentificationTypeUpdate) RemoveChildIDs(ids ...int) *GoodIdentificationTypeUpdate {
	gitu.mutation.RemoveChildIDs(ids...)
	return gitu
}

// RemoveChildren removes "children" edges to GoodIdentificationType entities.
func (gitu *GoodIdentificationTypeUpdate) RemoveChildren(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gitu.RemoveChildIDs(ids...)
}

// ClearChildGoodIdentificationTypes clears all "child_good_identification_types" edges to the GoodIdentificationType entity.
func (gitu *GoodIdentificationTypeUpdate) ClearChildGoodIdentificationTypes() *GoodIdentificationTypeUpdate {
	gitu.mutation.ClearChildGoodIdentificationTypes()
	return gitu
}

// RemoveChildGoodIdentificationTypeIDs removes the "child_good_identification_types" edge to GoodIdentificationType entities by IDs.
func (gitu *GoodIdentificationTypeUpdate) RemoveChildGoodIdentificationTypeIDs(ids ...int) *GoodIdentificationTypeUpdate {
	gitu.mutation.RemoveChildGoodIdentificationTypeIDs(ids...)
	return gitu
}

// RemoveChildGoodIdentificationTypes removes "child_good_identification_types" edges to GoodIdentificationType entities.
func (gitu *GoodIdentificationTypeUpdate) RemoveChildGoodIdentificationTypes(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gitu.RemoveChildGoodIdentificationTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gitu *GoodIdentificationTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gitu.defaults()
	if len(gitu.hooks) == 0 {
		if err = gitu.check(); err != nil {
			return 0, err
		}
		affected, err = gitu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodIdentificationTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gitu.check(); err != nil {
				return 0, err
			}
			gitu.mutation = mutation
			affected, err = gitu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gitu.hooks) - 1; i >= 0; i-- {
			mut = gitu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gitu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gitu *GoodIdentificationTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := gitu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gitu *GoodIdentificationTypeUpdate) Exec(ctx context.Context) error {
	_, err := gitu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gitu *GoodIdentificationTypeUpdate) ExecX(ctx context.Context) {
	if err := gitu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gitu *GoodIdentificationTypeUpdate) defaults() {
	if _, ok := gitu.mutation.UpdateTime(); !ok {
		v := goodidentificationtype.UpdateDefaultUpdateTime()
		gitu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gitu *GoodIdentificationTypeUpdate) check() error {
	if v, ok := gitu.mutation.HasTable(); ok {
		if err := goodidentificationtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (gitu *GoodIdentificationTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodidentificationtype.Table,
			Columns: goodidentificationtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodidentificationtype.FieldID,
			},
		},
	}
	if ps := gitu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gitu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodidentificationtype.FieldUpdateTime,
		})
	}
	if value, ok := gitu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodidentificationtype.FieldStringRef,
		})
	}
	if gitu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodidentificationtype.FieldStringRef,
		})
	}
	if value, ok := gitu.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: goodidentificationtype.FieldHasTable,
		})
	}
	if gitu.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: goodidentificationtype.FieldHasTable,
		})
	}
	if value, ok := gitu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodidentificationtype.FieldDescription,
		})
	}
	if gitu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodidentificationtype.FieldDescription,
		})
	}
	if gitu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodidentificationtype.ParentTable,
			Columns: []string{goodidentificationtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gitu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodidentificationtype.ParentTable,
			Columns: []string{goodidentificationtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gitu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gitu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !gitu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gitu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gitu.mutation.ChildGoodIdentificationTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gitu.mutation.RemovedChildGoodIdentificationTypesIDs(); len(nodes) > 0 && !gitu.mutation.ChildGoodIdentificationTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gitu.mutation.ChildGoodIdentificationTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gitu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodidentificationtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GoodIdentificationTypeUpdateOne is the builder for updating a single GoodIdentificationType entity.
type GoodIdentificationTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GoodIdentificationTypeMutation
}

// SetStringRef sets the "string_ref" field.
func (gituo *GoodIdentificationTypeUpdateOne) SetStringRef(s string) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.SetStringRef(s)
	return gituo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (gituo *GoodIdentificationTypeUpdateOne) SetNillableStringRef(s *string) *GoodIdentificationTypeUpdateOne {
	if s != nil {
		gituo.SetStringRef(*s)
	}
	return gituo
}

// ClearStringRef clears the value of the "string_ref" field.
func (gituo *GoodIdentificationTypeUpdateOne) ClearStringRef() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearStringRef()
	return gituo
}

// SetHasTable sets the "has_table" field.
func (gituo *GoodIdentificationTypeUpdateOne) SetHasTable(gt goodidentificationtype.HasTable) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.SetHasTable(gt)
	return gituo
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (gituo *GoodIdentificationTypeUpdateOne) SetNillableHasTable(gt *goodidentificationtype.HasTable) *GoodIdentificationTypeUpdateOne {
	if gt != nil {
		gituo.SetHasTable(*gt)
	}
	return gituo
}

// ClearHasTable clears the value of the "has_table" field.
func (gituo *GoodIdentificationTypeUpdateOne) ClearHasTable() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearHasTable()
	return gituo
}

// SetDescription sets the "description" field.
func (gituo *GoodIdentificationTypeUpdateOne) SetDescription(s string) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.SetDescription(s)
	return gituo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gituo *GoodIdentificationTypeUpdateOne) SetNillableDescription(s *string) *GoodIdentificationTypeUpdateOne {
	if s != nil {
		gituo.SetDescription(*s)
	}
	return gituo
}

// ClearDescription clears the value of the "description" field.
func (gituo *GoodIdentificationTypeUpdateOne) ClearDescription() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearDescription()
	return gituo
}

// SetParentID sets the "parent" edge to the GoodIdentificationType entity by ID.
func (gituo *GoodIdentificationTypeUpdateOne) SetParentID(id int) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.SetParentID(id)
	return gituo
}

// SetNillableParentID sets the "parent" edge to the GoodIdentificationType entity by ID if the given value is not nil.
func (gituo *GoodIdentificationTypeUpdateOne) SetNillableParentID(id *int) *GoodIdentificationTypeUpdateOne {
	if id != nil {
		gituo = gituo.SetParentID(*id)
	}
	return gituo
}

// SetParent sets the "parent" edge to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) SetParent(g *GoodIdentificationType) *GoodIdentificationTypeUpdateOne {
	return gituo.SetParentID(g.ID)
}

// AddChildIDs adds the "children" edge to the GoodIdentificationType entity by IDs.
func (gituo *GoodIdentificationTypeUpdateOne) AddChildIDs(ids ...int) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.AddChildIDs(ids...)
	return gituo
}

// AddChildren adds the "children" edges to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) AddChildren(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gituo.AddChildIDs(ids...)
}

// AddChildGoodIdentificationTypeIDs adds the "child_good_identification_types" edge to the GoodIdentificationType entity by IDs.
func (gituo *GoodIdentificationTypeUpdateOne) AddChildGoodIdentificationTypeIDs(ids ...int) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.AddChildGoodIdentificationTypeIDs(ids...)
	return gituo
}

// AddChildGoodIdentificationTypes adds the "child_good_identification_types" edges to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) AddChildGoodIdentificationTypes(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gituo.AddChildGoodIdentificationTypeIDs(ids...)
}

// Mutation returns the GoodIdentificationTypeMutation object of the builder.
func (gituo *GoodIdentificationTypeUpdateOne) Mutation() *GoodIdentificationTypeMutation {
	return gituo.mutation
}

// ClearParent clears the "parent" edge to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) ClearParent() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearParent()
	return gituo
}

// ClearChildren clears all "children" edges to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) ClearChildren() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearChildren()
	return gituo
}

// RemoveChildIDs removes the "children" edge to GoodIdentificationType entities by IDs.
func (gituo *GoodIdentificationTypeUpdateOne) RemoveChildIDs(ids ...int) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.RemoveChildIDs(ids...)
	return gituo
}

// RemoveChildren removes "children" edges to GoodIdentificationType entities.
func (gituo *GoodIdentificationTypeUpdateOne) RemoveChildren(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gituo.RemoveChildIDs(ids...)
}

// ClearChildGoodIdentificationTypes clears all "child_good_identification_types" edges to the GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) ClearChildGoodIdentificationTypes() *GoodIdentificationTypeUpdateOne {
	gituo.mutation.ClearChildGoodIdentificationTypes()
	return gituo
}

// RemoveChildGoodIdentificationTypeIDs removes the "child_good_identification_types" edge to GoodIdentificationType entities by IDs.
func (gituo *GoodIdentificationTypeUpdateOne) RemoveChildGoodIdentificationTypeIDs(ids ...int) *GoodIdentificationTypeUpdateOne {
	gituo.mutation.RemoveChildGoodIdentificationTypeIDs(ids...)
	return gituo
}

// RemoveChildGoodIdentificationTypes removes "child_good_identification_types" edges to GoodIdentificationType entities.
func (gituo *GoodIdentificationTypeUpdateOne) RemoveChildGoodIdentificationTypes(g ...*GoodIdentificationType) *GoodIdentificationTypeUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gituo.RemoveChildGoodIdentificationTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gituo *GoodIdentificationTypeUpdateOne) Select(field string, fields ...string) *GoodIdentificationTypeUpdateOne {
	gituo.fields = append([]string{field}, fields...)
	return gituo
}

// Save executes the query and returns the updated GoodIdentificationType entity.
func (gituo *GoodIdentificationTypeUpdateOne) Save(ctx context.Context) (*GoodIdentificationType, error) {
	var (
		err  error
		node *GoodIdentificationType
	)
	gituo.defaults()
	if len(gituo.hooks) == 0 {
		if err = gituo.check(); err != nil {
			return nil, err
		}
		node, err = gituo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodIdentificationTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gituo.check(); err != nil {
				return nil, err
			}
			gituo.mutation = mutation
			node, err = gituo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gituo.hooks) - 1; i >= 0; i-- {
			mut = gituo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gituo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gituo *GoodIdentificationTypeUpdateOne) SaveX(ctx context.Context) *GoodIdentificationType {
	node, err := gituo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gituo *GoodIdentificationTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := gituo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gituo *GoodIdentificationTypeUpdateOne) ExecX(ctx context.Context) {
	if err := gituo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gituo *GoodIdentificationTypeUpdateOne) defaults() {
	if _, ok := gituo.mutation.UpdateTime(); !ok {
		v := goodidentificationtype.UpdateDefaultUpdateTime()
		gituo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gituo *GoodIdentificationTypeUpdateOne) check() error {
	if v, ok := gituo.mutation.HasTable(); ok {
		if err := goodidentificationtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (gituo *GoodIdentificationTypeUpdateOne) sqlSave(ctx context.Context) (_node *GoodIdentificationType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodidentificationtype.Table,
			Columns: goodidentificationtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goodidentificationtype.FieldID,
			},
		},
	}
	id, ok := gituo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GoodIdentificationType.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := gituo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodidentificationtype.FieldID)
		for _, f := range fields {
			if !goodidentificationtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodidentificationtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gituo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gituo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: goodidentificationtype.FieldUpdateTime,
		})
	}
	if value, ok := gituo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodidentificationtype.FieldStringRef,
		})
	}
	if gituo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodidentificationtype.FieldStringRef,
		})
	}
	if value, ok := gituo.mutation.HasTable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: goodidentificationtype.FieldHasTable,
		})
	}
	if gituo.mutation.HasTableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: goodidentificationtype.FieldHasTable,
		})
	}
	if value, ok := gituo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodidentificationtype.FieldDescription,
		})
	}
	if gituo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodidentificationtype.FieldDescription,
		})
	}
	if gituo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodidentificationtype.ParentTable,
			Columns: []string{goodidentificationtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gituo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   goodidentificationtype.ParentTable,
			Columns: []string{goodidentificationtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gituo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gituo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !gituo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gituo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildrenTable,
			Columns: []string{goodidentificationtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gituo.mutation.ChildGoodIdentificationTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gituo.mutation.RemovedChildGoodIdentificationTypesIDs(); len(nodes) > 0 && !gituo.mutation.ChildGoodIdentificationTypesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gituo.mutation.ChildGoodIdentificationTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   goodidentificationtype.ChildGoodIdentificationTypesTable,
			Columns: goodidentificationtype.ChildGoodIdentificationTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: goodidentificationtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GoodIdentificationType{config: gituo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gituo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodidentificationtype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
