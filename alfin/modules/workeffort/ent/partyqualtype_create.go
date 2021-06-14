// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyqualtype"
)

// PartyQualTypeCreate is the builder for creating a PartyQualType entity.
type PartyQualTypeCreate struct {
	config
	mutation *PartyQualTypeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pqtc *PartyQualTypeCreate) SetCreateTime(t time.Time) *PartyQualTypeCreate {
	pqtc.mutation.SetCreateTime(t)
	return pqtc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableCreateTime(t *time.Time) *PartyQualTypeCreate {
	if t != nil {
		pqtc.SetCreateTime(*t)
	}
	return pqtc
}

// SetUpdateTime sets the "update_time" field.
func (pqtc *PartyQualTypeCreate) SetUpdateTime(t time.Time) *PartyQualTypeCreate {
	pqtc.mutation.SetUpdateTime(t)
	return pqtc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableUpdateTime(t *time.Time) *PartyQualTypeCreate {
	if t != nil {
		pqtc.SetUpdateTime(*t)
	}
	return pqtc
}

// SetStringRef sets the "string_ref" field.
func (pqtc *PartyQualTypeCreate) SetStringRef(s string) *PartyQualTypeCreate {
	pqtc.mutation.SetStringRef(s)
	return pqtc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableStringRef(s *string) *PartyQualTypeCreate {
	if s != nil {
		pqtc.SetStringRef(*s)
	}
	return pqtc
}

// SetHasTable sets the "has_table" field.
func (pqtc *PartyQualTypeCreate) SetHasTable(pt partyqualtype.HasTable) *PartyQualTypeCreate {
	pqtc.mutation.SetHasTable(pt)
	return pqtc
}

// SetNillableHasTable sets the "has_table" field if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableHasTable(pt *partyqualtype.HasTable) *PartyQualTypeCreate {
	if pt != nil {
		pqtc.SetHasTable(*pt)
	}
	return pqtc
}

// SetDescription sets the "description" field.
func (pqtc *PartyQualTypeCreate) SetDescription(s string) *PartyQualTypeCreate {
	pqtc.mutation.SetDescription(s)
	return pqtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableDescription(s *string) *PartyQualTypeCreate {
	if s != nil {
		pqtc.SetDescription(*s)
	}
	return pqtc
}

// SetParentID sets the "parent" edge to the PartyQualType entity by ID.
func (pqtc *PartyQualTypeCreate) SetParentID(id int) *PartyQualTypeCreate {
	pqtc.mutation.SetParentID(id)
	return pqtc
}

// SetNillableParentID sets the "parent" edge to the PartyQualType entity by ID if the given value is not nil.
func (pqtc *PartyQualTypeCreate) SetNillableParentID(id *int) *PartyQualTypeCreate {
	if id != nil {
		pqtc = pqtc.SetParentID(*id)
	}
	return pqtc
}

// SetParent sets the "parent" edge to the PartyQualType entity.
func (pqtc *PartyQualTypeCreate) SetParent(p *PartyQualType) *PartyQualTypeCreate {
	return pqtc.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the PartyQualType entity by IDs.
func (pqtc *PartyQualTypeCreate) AddChildIDs(ids ...int) *PartyQualTypeCreate {
	pqtc.mutation.AddChildIDs(ids...)
	return pqtc
}

// AddChildren adds the "children" edges to the PartyQualType entity.
func (pqtc *PartyQualTypeCreate) AddChildren(p ...*PartyQualType) *PartyQualTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtc.AddChildIDs(ids...)
}

// AddChildPartyQualTypeIDs adds the "child_party_qual_types" edge to the PartyQualType entity by IDs.
func (pqtc *PartyQualTypeCreate) AddChildPartyQualTypeIDs(ids ...int) *PartyQualTypeCreate {
	pqtc.mutation.AddChildPartyQualTypeIDs(ids...)
	return pqtc
}

// AddChildPartyQualTypes adds the "child_party_qual_types" edges to the PartyQualType entity.
func (pqtc *PartyQualTypeCreate) AddChildPartyQualTypes(p ...*PartyQualType) *PartyQualTypeCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pqtc.AddChildPartyQualTypeIDs(ids...)
}

// Mutation returns the PartyQualTypeMutation object of the builder.
func (pqtc *PartyQualTypeCreate) Mutation() *PartyQualTypeMutation {
	return pqtc.mutation
}

// Save creates the PartyQualType in the database.
func (pqtc *PartyQualTypeCreate) Save(ctx context.Context) (*PartyQualType, error) {
	var (
		err  error
		node *PartyQualType
	)
	pqtc.defaults()
	if len(pqtc.hooks) == 0 {
		if err = pqtc.check(); err != nil {
			return nil, err
		}
		node, err = pqtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PartyQualTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pqtc.check(); err != nil {
				return nil, err
			}
			pqtc.mutation = mutation
			if node, err = pqtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pqtc.hooks) - 1; i >= 0; i-- {
			mut = pqtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pqtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pqtc *PartyQualTypeCreate) SaveX(ctx context.Context) *PartyQualType {
	v, err := pqtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pqtc *PartyQualTypeCreate) defaults() {
	if _, ok := pqtc.mutation.CreateTime(); !ok {
		v := partyqualtype.DefaultCreateTime()
		pqtc.mutation.SetCreateTime(v)
	}
	if _, ok := pqtc.mutation.UpdateTime(); !ok {
		v := partyqualtype.DefaultUpdateTime()
		pqtc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pqtc *PartyQualTypeCreate) check() error {
	if _, ok := pqtc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pqtc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pqtc.mutation.HasTable(); ok {
		if err := partyqualtype.HasTableValidator(v); err != nil {
			return &ValidationError{Name: "has_table", err: fmt.Errorf("ent: validator failed for field \"has_table\": %w", err)}
		}
	}
	return nil
}

func (pqtc *PartyQualTypeCreate) sqlSave(ctx context.Context) (*PartyQualType, error) {
	_node, _spec := pqtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pqtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pqtc *PartyQualTypeCreate) createSpec() (*PartyQualType, *sqlgraph.CreateSpec) {
	var (
		_node = &PartyQualType{config: pqtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: partyqualtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: partyqualtype.FieldID,
			},
		}
	)
	if value, ok := pqtc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyqualtype.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pqtc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: partyqualtype.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pqtc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pqtc.mutation.HasTable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: partyqualtype.FieldHasTable,
		})
		_node.HasTable = value
	}
	if value, ok := pqtc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: partyqualtype.FieldDescription,
		})
		_node.Description = value
	}
	if nodes := pqtc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   partyqualtype.ParentTable,
			Columns: []string{partyqualtype.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.party_qual_type_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pqtc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   partyqualtype.ChildrenTable,
			Columns: []string{partyqualtype.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pqtc.mutation.ChildPartyQualTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   partyqualtype.ChildPartyQualTypesTable,
			Columns: partyqualtype.ChildPartyQualTypesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partyqualtype.FieldID,
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

// PartyQualTypeCreateBulk is the builder for creating many PartyQualType entities in bulk.
type PartyQualTypeCreateBulk struct {
	config
	builders []*PartyQualTypeCreate
}

// Save creates the PartyQualType entities in the database.
func (pqtcb *PartyQualTypeCreateBulk) Save(ctx context.Context) ([]*PartyQualType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pqtcb.builders))
	nodes := make([]*PartyQualType, len(pqtcb.builders))
	mutators := make([]Mutator, len(pqtcb.builders))
	for i := range pqtcb.builders {
		func(i int, root context.Context) {
			builder := pqtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartyQualTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, pqtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pqtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pqtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pqtcb *PartyQualTypeCreateBulk) SaveX(ctx context.Context) []*PartyQualType {
	v, err := pqtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}