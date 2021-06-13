// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/skilltype"
)

// SkillTypeDelete is the builder for deleting a SkillType entity.
type SkillTypeDelete struct {
	config
	hooks    []Hook
	mutation *SkillTypeMutation
}

// Where adds a new predicate to the SkillTypeDelete builder.
func (std *SkillTypeDelete) Where(ps ...predicate.SkillType) *SkillTypeDelete {
	std.mutation.predicates = append(std.mutation.predicates, ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *SkillTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(std.hooks) == 0 {
		affected, err = std.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkillTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			std.mutation = mutation
			affected, err = std.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(std.hooks) - 1; i >= 0; i-- {
			mut = std.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, std.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (std *SkillTypeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *SkillTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: skilltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skilltype.FieldID,
			},
		},
	}
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, std.driver, _spec)
}

// SkillTypeDeleteOne is the builder for deleting a single SkillType entity.
type SkillTypeDeleteOne struct {
	std *SkillTypeDelete
}

// Exec executes the deletion query.
func (stdo *SkillTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{skilltype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *SkillTypeDeleteOne) ExecX(ctx context.Context) {
	stdo.std.ExecX(ctx)
}
