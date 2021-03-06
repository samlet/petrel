// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtypepurpose"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ContactMechTypePurposeDelete is the builder for deleting a ContactMechTypePurpose entity.
type ContactMechTypePurposeDelete struct {
	config
	hooks    []Hook
	mutation *ContactMechTypePurposeMutation
}

// Where adds a new predicate to the ContactMechTypePurposeDelete builder.
func (cmtpd *ContactMechTypePurposeDelete) Where(ps ...predicate.ContactMechTypePurpose) *ContactMechTypePurposeDelete {
	cmtpd.mutation.predicates = append(cmtpd.mutation.predicates, ps...)
	return cmtpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmtpd *ContactMechTypePurposeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmtpd.hooks) == 0 {
		affected, err = cmtpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContactMechTypePurposeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmtpd.mutation = mutation
			affected, err = cmtpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmtpd.hooks) - 1; i >= 0; i-- {
			mut = cmtpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmtpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmtpd *ContactMechTypePurposeDelete) ExecX(ctx context.Context) int {
	n, err := cmtpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmtpd *ContactMechTypePurposeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: contactmechtypepurpose.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: contactmechtypepurpose.FieldID,
			},
		},
	}
	if ps := cmtpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cmtpd.driver, _spec)
}

// ContactMechTypePurposeDeleteOne is the builder for deleting a single ContactMechTypePurpose entity.
type ContactMechTypePurposeDeleteOne struct {
	cmtpd *ContactMechTypePurposeDelete
}

// Exec executes the deletion query.
func (cmtpdo *ContactMechTypePurposeDeleteOne) Exec(ctx context.Context) error {
	n, err := cmtpdo.cmtpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contactmechtypepurpose.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmtpdo *ContactMechTypePurposeDeleteOne) ExecX(ctx context.Context) {
	cmtpdo.cmtpd.ExecX(ctx)
}
