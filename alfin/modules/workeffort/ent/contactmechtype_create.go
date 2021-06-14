// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtypepurpose"
)

// ContactMechTypeCreate is the builder for creating a ContactMechType entity.
type ContactMechTypeCreate struct {
	config
	mutation *ContactMechTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cmtc *ContactMechTypeCreate) SetCreateTime(t time.Time) *ContactMechTypeCreate {
	cmtc.mutation.SetCreateTime(t)
	return cmtc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableCreateTime(t *time.Time) *ContactMechTypeCreate {
	if t != nil {
		cmtc.SetCreateTime(*t)
	}
	return cmtc
}

// SetUpdateTime sets the "update_time" field.
func (cmtc *ContactMechTypeCreate) SetUpdateTime(t time.Time) *ContactMechTypeCreate {
	cmtc.mutation.SetUpdateTime(t)
	return cmtc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableUpdateTime(t *time.Time) *ContactMechTypeCreate {
	if t != nil {
		cmtc.SetUpdateTime(*t)
	}
	return cmtc
}

// SetStringRef sets the "string_ref" field.
func (cmtc *ContactMechTypeCreate) SetStringRef(s string) *ContactMechTypeCreate {
	cmtc.mutation.SetStringRef(s)
	return cmtc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableStringRef(s *string) *ContactMechTypeCreate {
	if s != nil {
		cmtc.SetStringRef(*s)
	}
	return cmtc
}

// SetHasTable sets the "has_table" field.
func (cmtc *ContactMechTypeCreate) SetHasTable(ct contactmechtype.HasTable) *ContactMechTypeCreate {
	cmtc.mutation.SetHasTable(ct)
	return cmtc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableHasTable(ct *contactmechtype.HasTable) *ContactMechTypeCreate {
	if ct != nil {
		cmtc.SetHasTable(*ct)
	}
	return cmtc
}

// SetDescription sets the "description" field.
func (cmtc *ContactMechTypeCreate) SetDescription(s string) *ContactMechTypeCreate {
	cmtc.mutation.SetDescription(s)
	return cmtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableDescription(s *string) *ContactMechTypeCreate {
	if s != nil {
		cmtc.SetDescription(*s)
	}
	return cmtc
}

// SetParentID sets the "parent" edge to the ContactMechType entity by ID.
func (cmtc *ContactMechTypeCreate) SetParentID(id int) *ContactMechTypeCreate {
	cmtc.mutation.SetParentID(id)
	return cmtc
}

// SetNillableParentID sets the "parent" edge to the ContactMechType entity by ID if the given value is not nil.
func (cmtc *ContactMechTypeCreate) SetNillableParentID(id *int) *ContactMechTypeCreate {
	if id != nil {
		cmtc = cmtc.SetParentID(*id)
	}
	return cmtc
}

// SetParent sets the "parent" edge to the ContactMechType entity.
func (cmtc *ContactMechTypeCreate) SetParent(c *ContactMechType) *ContactMechTypeCreate {
	return cmtc.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the ContactMechType entity by IDs.
func (cmtc *ContactMechTypeCreate) AddChildIDs(ids ...int) *ContactMechTypeCreate {
	cmtc.mutation.AddChildIDs(ids...)
	return cmtc
}

// AddChildren adds the "children" edges to the ContactMechType entity.
func (cmtc *ContactMechTypeCreate) AddChildren(c ...*ContactMechType) *ContactMechTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtc.AddChildIDs(ids...)
}

// AddContacMechTypeCommunicationEventTypeIDs adds the "contac_mech_type_communication_event_types" edge to the CommunicationEventType entity by IDs.
func (cmtc *ContactMechTypeCreate) AddContacMechTypeCommunicationEventTypeIDs(ids ...int) *ContactMechTypeCreate {
	cmtc.mutation.AddContacMechTypeCommunicationEventTypeIDs(ids...)
	return cmtc
}

// AddContacMechTypeCommunicationEventTypes adds the "contac_mech_type_communication_event_types" edges to the CommunicationEventType entity.
func (cmtc *ContactMechTypeCreate) AddContacMechTypeCommunicationEventTypes(c ...*CommunicationEventType) *ContactMechTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtc.AddContacMechTypeCommunicationEventTypeIDs(ids...)
}

// AddChildContactMechTypeIDs adds the "child_contact_mech_types" edge to the ContactMechType entity by IDs.
func (cmtc *ContactMechTypeCreate) AddChildContactMechTypeIDs(ids ...int) *ContactMechTypeCreate {
	cmtc.mutation.AddChildContactMechTypeIDs(ids...)
	return cmtc
}

// AddChildContactMechTypes adds the "child_contact_mech_types" edges to the ContactMechType entity.
func (cmtc *ContactMechTypeCreate) AddChildContactMechTypes(c ...*ContactMechType) *ContactMechTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtc.AddChildContactMechTypeIDs(ids...)
}

// AddContactMechTypePurposeIDs adds the "contact_mech_type_purposes" edge to the ContactMechTypePurpose entity by IDs.
func (cmtc *ContactMechTypeCreate) AddContactMechTypePurposeIDs(ids ...int) *ContactMechTypeCreate {
	cmtc.mutation.AddContactMechTypePurposeIDs(ids...)
	return cmtc
}

// AddContactMechTypePurposes adds the "contact_mech_type_purposes" edges to the ContactMechTypePurpose entity.
func (cmtc *ContactMechTypeCreate) AddContactMechTypePurposes(c ...*ContactMechTypePurpose) *ContactMechTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cmtc.AddContactMechTypePurposeIDs(ids...)
}

// Mutation returns the ContactMechTypeMutation object of the builder.
func (cmtc *ContactMechTypeCreate) Mutation() *ContactMechTypeMutation {
	return cmtc.mutation
}

// Save creates the ContactMechType in the database.
func (cmtc *ContactMechTypeCreate) Save(ctx context.Context) (*ContactMechType, error) {
	var (
		err  error
		node *ContactMechType
	)
	cmtc.defaults()
	if len(cmtc.hooks) == 0 {
		if err = cmtc.check(); err != nil {
			return nil, err
		}
		node, err = cmtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMechTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cmtc.check(); err != nil {
				return nil, err
			}
			cmtc.mutation = mutation
			if node, err = cmtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cmtc.hooks) - 1; i >= 0; i-- {
			mut = cmtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cmtc *ContactMechTypeCreate) SaveX(ctx context.Context) *ContactMechType {
	v, err := cmtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cmtc *ContactMechTypeCreate) defaults() {
	if _, ok := cmtc.mutation.CreateTime(); !ok {
		v := contactmechtype.DefaultCreateTime()
		cmtc.mutation.SetCreateTime(v)
	}
	if _, ok := cmtc.mutation.UpdateTime(); !ok {
		v := contactmechtype.DefaultUpdateTime()
		cmtc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmtc *ContactMechTypeCreate) check() error {
	if _, ok := cmtc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := cmtc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := cmtc.mutation.HasTable(); ok {
		if err := contactmechtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (cmtc *ContactMechTypeCreate) sqlSave(ctx context.Context) (*ContactMechType, error) {
	_node, _spec := cmtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cmtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cmtc *ContactMechTypeCreate) createSpec() (*ContactMechType, *sqlgraph.CreateSpec) {
	var (
		_node = &ContactMechType{config: cmtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: contactmechtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contactmechtype.FieldID,
			},
		}
	)
	if value, ok := cmtc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contactmechtype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cmtc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: contactmechtype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cmtc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechtype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := cmtc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: contactmechtype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := cmtc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contactmechtype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := cmtc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contactmechtype.ParentTable,
			Columns: []string{contactmechtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.contact_mech_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cmtc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechtype.ChildrenTable,
			Columns: []string{contactmechtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cmtc.mutation.ContacMechTypeCommunicationEventTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechtype.ContacMechTypeCommunicationEventTypesTable,
			Columns: []string{contactmechtype.ContacMechTypeCommunicationEventTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: communicationeventtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cmtc.mutation.ChildContactMechTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   contactmechtype.ChildContactMechTypesTable,
			Columns: contactmechtype.ChildContactMechTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cmtc.mutation.ContactMechTypePurposesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contactmechtype.ContactMechTypePurposesTable,
			Columns: []string{contactmechtype.ContactMechTypePurposesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: contactmechtypepurpose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContactMechTypeCreateBulk is the builder for creating many ContactMechType entities in bulk.
type ContactMechTypeCreateBulk struct {
	config
	builders []*ContactMechTypeCreate
}

// Save creates the ContactMechType entities in the database.
func (cmtcb *ContactMechTypeCreateBulk) Save(ctx context.Context) ([]*ContactMechType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cmtcb.builders))
	nodes := make([]*ContactMechType, len(cmtcb.builders))
	mutators := make([]Mutator, len(cmtcb.builders))
	for i := range cmtcb.builders {
		func(i int, root context.Context) {
			builder := cmtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactMechTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cmtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cmtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cmtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cmtcb *ContactMechTypeCreateBulk) SaveX(ctx context.Context) []*ContactMechType {
	v, err := cmtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}