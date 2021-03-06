// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productprice"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstore"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productstoregroup"
)

// ProductStoreGroupUpdate is the builder for updating ProductStoreGroup entities.
type ProductStoreGroupUpdate struct {
	config
	hooks    []Hook
	mutation *ProductStoreGroupMutation
}

// Where adds a new predicate for the ProductStoreGroupUpdate builder.
func (psgu *ProductStoreGroupUpdate) Where(ps ...predicate.ProductStoreGroup) *ProductStoreGroupUpdate {
	psgu.mutation.predicates = append(psgu.mutation.predicates, ps...)
	return psgu
}

// SetStringRef sets the "string_ref" field.
func (psgu *ProductStoreGroupUpdate) SetStringRef(s string) *ProductStoreGroupUpdate {
	psgu.mutation.SetStringRef(s)
	return psgu
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (psgu *ProductStoreGroupUpdate) SetNillableStringRef(s *string) *ProductStoreGroupUpdate {
	if s != nil {
		psgu.SetStringRef(*s)
	}
	return psgu
}

// ClearStringRef clears the value of the "string_ref" field.
func (psgu *ProductStoreGroupUpdate) ClearStringRef() *ProductStoreGroupUpdate {
	psgu.mutation.ClearStringRef()
	return psgu
}

// SetProductStoreGroupTypeID sets the "product_store_group_type_id" field.
func (psgu *ProductStoreGroupUpdate) SetProductStoreGroupTypeID(i int) *ProductStoreGroupUpdate {
	psgu.mutation.ResetProductStoreGroupTypeID()
	psgu.mutation.SetProductStoreGroupTypeID(i)
	return psgu
}

// SetNillableProductStoreGroupTypeID sets the "product_store_group_type_id" field if the given value is not nil.
func (psgu *ProductStoreGroupUpdate) SetNillableProductStoreGroupTypeID(i *int) *ProductStoreGroupUpdate {
	if i != nil {
		psgu.SetProductStoreGroupTypeID(*i)
	}
	return psgu
}

// AddProductStoreGroupTypeID adds i to the "product_store_group_type_id" field.
func (psgu *ProductStoreGroupUpdate) AddProductStoreGroupTypeID(i int) *ProductStoreGroupUpdate {
	psgu.mutation.AddProductStoreGroupTypeID(i)
	return psgu
}

// ClearProductStoreGroupTypeID clears the value of the "product_store_group_type_id" field.
func (psgu *ProductStoreGroupUpdate) ClearProductStoreGroupTypeID() *ProductStoreGroupUpdate {
	psgu.mutation.ClearProductStoreGroupTypeID()
	return psgu
}

// SetProductStoreGroupName sets the "product_store_group_name" field.
func (psgu *ProductStoreGroupUpdate) SetProductStoreGroupName(s string) *ProductStoreGroupUpdate {
	psgu.mutation.SetProductStoreGroupName(s)
	return psgu
}

// SetNillableProductStoreGroupName sets the "product_store_group_name" field if the given value is not nil.
func (psgu *ProductStoreGroupUpdate) SetNillableProductStoreGroupName(s *string) *ProductStoreGroupUpdate {
	if s != nil {
		psgu.SetProductStoreGroupName(*s)
	}
	return psgu
}

// ClearProductStoreGroupName clears the value of the "product_store_group_name" field.
func (psgu *ProductStoreGroupUpdate) ClearProductStoreGroupName() *ProductStoreGroupUpdate {
	psgu.mutation.ClearProductStoreGroupName()
	return psgu
}

// SetDescription sets the "description" field.
func (psgu *ProductStoreGroupUpdate) SetDescription(s string) *ProductStoreGroupUpdate {
	psgu.mutation.SetDescription(s)
	return psgu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (psgu *ProductStoreGroupUpdate) SetNillableDescription(s *string) *ProductStoreGroupUpdate {
	if s != nil {
		psgu.SetDescription(*s)
	}
	return psgu
}

// ClearDescription clears the value of the "description" field.
func (psgu *ProductStoreGroupUpdate) ClearDescription() *ProductStoreGroupUpdate {
	psgu.mutation.ClearDescription()
	return psgu
}

// SetParentID sets the "parent" edge to the ProductStoreGroup entity by ID.
func (psgu *ProductStoreGroupUpdate) SetParentID(id int) *ProductStoreGroupUpdate {
	psgu.mutation.SetParentID(id)
	return psgu
}

// SetNillableParentID sets the "parent" edge to the ProductStoreGroup entity by ID if the given value is not nil.
func (psgu *ProductStoreGroupUpdate) SetNillableParentID(id *int) *ProductStoreGroupUpdate {
	if id != nil {
		psgu = psgu.SetParentID(*id)
	}
	return psgu
}

// SetParent sets the "parent" edge to the ProductStoreGroup entity.
func (psgu *ProductStoreGroupUpdate) SetParent(p *ProductStoreGroup) *ProductStoreGroupUpdate {
	return psgu.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductStoreGroup entity by IDs.
func (psgu *ProductStoreGroupUpdate) AddChildIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.AddChildIDs(ids...)
	return psgu
}

// AddChildren adds the "children" edges to the ProductStoreGroup entity.
func (psgu *ProductStoreGroupUpdate) AddChildren(p ...*ProductStoreGroup) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.AddChildIDs(ids...)
}

// AddProductPriceIDs adds the "product_prices" edge to the ProductPrice entity by IDs.
func (psgu *ProductStoreGroupUpdate) AddProductPriceIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.AddProductPriceIDs(ids...)
	return psgu
}

// AddProductPrices adds the "product_prices" edges to the ProductPrice entity.
func (psgu *ProductStoreGroupUpdate) AddProductPrices(p ...*ProductPrice) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.AddProductPriceIDs(ids...)
}

// AddPrimaryProductStoreIDs adds the "primary_product_stores" edge to the ProductStore entity by IDs.
func (psgu *ProductStoreGroupUpdate) AddPrimaryProductStoreIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.AddPrimaryProductStoreIDs(ids...)
	return psgu
}

// AddPrimaryProductStores adds the "primary_product_stores" edges to the ProductStore entity.
func (psgu *ProductStoreGroupUpdate) AddPrimaryProductStores(p ...*ProductStore) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.AddPrimaryProductStoreIDs(ids...)
}

// Mutation returns the ProductStoreGroupMutation object of the builder.
func (psgu *ProductStoreGroupUpdate) Mutation() *ProductStoreGroupMutation {
	return psgu.mutation
}

// ClearParent clears the "parent" edge to the ProductStoreGroup entity.
func (psgu *ProductStoreGroupUpdate) ClearParent() *ProductStoreGroupUpdate {
	psgu.mutation.ClearParent()
	return psgu
}

// ClearChildren clears all "children" edges to the ProductStoreGroup entity.
func (psgu *ProductStoreGroupUpdate) ClearChildren() *ProductStoreGroupUpdate {
	psgu.mutation.ClearChildren()
	return psgu
}

// RemoveChildIDs removes the "children" edge to ProductStoreGroup entities by IDs.
func (psgu *ProductStoreGroupUpdate) RemoveChildIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.RemoveChildIDs(ids...)
	return psgu
}

// RemoveChildren removes "children" edges to ProductStoreGroup entities.
func (psgu *ProductStoreGroupUpdate) RemoveChildren(p ...*ProductStoreGroup) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.RemoveChildIDs(ids...)
}

// ClearProductPrices clears all "product_prices" edges to the ProductPrice entity.
func (psgu *ProductStoreGroupUpdate) ClearProductPrices() *ProductStoreGroupUpdate {
	psgu.mutation.ClearProductPrices()
	return psgu
}

// RemoveProductPriceIDs removes the "product_prices" edge to ProductPrice entities by IDs.
func (psgu *ProductStoreGroupUpdate) RemoveProductPriceIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.RemoveProductPriceIDs(ids...)
	return psgu
}

// RemoveProductPrices removes "product_prices" edges to ProductPrice entities.
func (psgu *ProductStoreGroupUpdate) RemoveProductPrices(p ...*ProductPrice) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.RemoveProductPriceIDs(ids...)
}

// ClearPrimaryProductStores clears all "primary_product_stores" edges to the ProductStore entity.
func (psgu *ProductStoreGroupUpdate) ClearPrimaryProductStores() *ProductStoreGroupUpdate {
	psgu.mutation.ClearPrimaryProductStores()
	return psgu
}

// RemovePrimaryProductStoreIDs removes the "primary_product_stores" edge to ProductStore entities by IDs.
func (psgu *ProductStoreGroupUpdate) RemovePrimaryProductStoreIDs(ids ...int) *ProductStoreGroupUpdate {
	psgu.mutation.RemovePrimaryProductStoreIDs(ids...)
	return psgu
}

// RemovePrimaryProductStores removes "primary_product_stores" edges to ProductStore entities.
func (psgu *ProductStoreGroupUpdate) RemovePrimaryProductStores(p ...*ProductStore) *ProductStoreGroupUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psgu.RemovePrimaryProductStoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psgu *ProductStoreGroupUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	psgu.defaults()
	if len(psgu.hooks) == 0 {
		affected, err = psgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductStoreGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psgu.mutation = mutation
			affected, err = psgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psgu.hooks) - 1; i >= 0; i-- {
			mut = psgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psgu *ProductStoreGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := psgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psgu *ProductStoreGroupUpdate) Exec(ctx context.Context) error {
	_, err := psgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psgu *ProductStoreGroupUpdate) ExecX(ctx context.Context) {
	if err := psgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psgu *ProductStoreGroupUpdate) defaults() {
	if _, ok := psgu.mutation.UpdateTime(); !ok {
		v := productstoregroup.UpdateDefaultUpdateTime()
		psgu.mutation.SetUpdateTime(v)
	}
}

func (psgu *ProductStoreGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productstoregroup.Table,
			Columns: productstoregroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productstoregroup.FieldID,
			},
		},
	}
	if ps := psgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psgu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productstoregroup.FieldUpdateTime,
		})
	}
	if value, ok := psgu.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldStringRef,
		})
	}
	if psgu.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldStringRef,
		})
	}
	if value, ok := psgu.mutation.ProductStoreGroupTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if value, ok := psgu.mutation.AddedProductStoreGroupTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if psgu.mutation.ProductStoreGroupTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if value, ok := psgu.mutation.ProductStoreGroupName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupName,
		})
	}
	if psgu.mutation.ProductStoreGroupNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldProductStoreGroupName,
		})
	}
	if value, ok := psgu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldDescription,
		})
	}
	if psgu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldDescription,
		})
	}
	if psgu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productstoregroup.ParentTable,
			Columns: []string{productstoregroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productstoregroup.ParentTable,
			Columns: []string{productstoregroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psgu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !psgu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psgu.mutation.ProductPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.RemovedProductPricesIDs(); len(nodes) > 0 && !psgu.mutation.ProductPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.ProductPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psgu.mutation.PrimaryProductStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.RemovedPrimaryProductStoresIDs(); len(nodes) > 0 && !psgu.mutation.PrimaryProductStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psgu.mutation.PrimaryProductStoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, psgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productstoregroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductStoreGroupUpdateOne is the builder for updating a single ProductStoreGroup entity.
type ProductStoreGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductStoreGroupMutation
}

// SetStringRef sets the "string_ref" field.
func (psguo *ProductStoreGroupUpdateOne) SetStringRef(s string) *ProductStoreGroupUpdateOne {
	psguo.mutation.SetStringRef(s)
	return psguo
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (psguo *ProductStoreGroupUpdateOne) SetNillableStringRef(s *string) *ProductStoreGroupUpdateOne {
	if s != nil {
		psguo.SetStringRef(*s)
	}
	return psguo
}

// ClearStringRef clears the value of the "string_ref" field.
func (psguo *ProductStoreGroupUpdateOne) ClearStringRef() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearStringRef()
	return psguo
}

// SetProductStoreGroupTypeID sets the "product_store_group_type_id" field.
func (psguo *ProductStoreGroupUpdateOne) SetProductStoreGroupTypeID(i int) *ProductStoreGroupUpdateOne {
	psguo.mutation.ResetProductStoreGroupTypeID()
	psguo.mutation.SetProductStoreGroupTypeID(i)
	return psguo
}

// SetNillableProductStoreGroupTypeID sets the "product_store_group_type_id" field if the given value is not nil.
func (psguo *ProductStoreGroupUpdateOne) SetNillableProductStoreGroupTypeID(i *int) *ProductStoreGroupUpdateOne {
	if i != nil {
		psguo.SetProductStoreGroupTypeID(*i)
	}
	return psguo
}

// AddProductStoreGroupTypeID adds i to the "product_store_group_type_id" field.
func (psguo *ProductStoreGroupUpdateOne) AddProductStoreGroupTypeID(i int) *ProductStoreGroupUpdateOne {
	psguo.mutation.AddProductStoreGroupTypeID(i)
	return psguo
}

// ClearProductStoreGroupTypeID clears the value of the "product_store_group_type_id" field.
func (psguo *ProductStoreGroupUpdateOne) ClearProductStoreGroupTypeID() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearProductStoreGroupTypeID()
	return psguo
}

// SetProductStoreGroupName sets the "product_store_group_name" field.
func (psguo *ProductStoreGroupUpdateOne) SetProductStoreGroupName(s string) *ProductStoreGroupUpdateOne {
	psguo.mutation.SetProductStoreGroupName(s)
	return psguo
}

// SetNillableProductStoreGroupName sets the "product_store_group_name" field if the given value is not nil.
func (psguo *ProductStoreGroupUpdateOne) SetNillableProductStoreGroupName(s *string) *ProductStoreGroupUpdateOne {
	if s != nil {
		psguo.SetProductStoreGroupName(*s)
	}
	return psguo
}

// ClearProductStoreGroupName clears the value of the "product_store_group_name" field.
func (psguo *ProductStoreGroupUpdateOne) ClearProductStoreGroupName() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearProductStoreGroupName()
	return psguo
}

// SetDescription sets the "description" field.
func (psguo *ProductStoreGroupUpdateOne) SetDescription(s string) *ProductStoreGroupUpdateOne {
	psguo.mutation.SetDescription(s)
	return psguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (psguo *ProductStoreGroupUpdateOne) SetNillableDescription(s *string) *ProductStoreGroupUpdateOne {
	if s != nil {
		psguo.SetDescription(*s)
	}
	return psguo
}

// ClearDescription clears the value of the "description" field.
func (psguo *ProductStoreGroupUpdateOne) ClearDescription() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearDescription()
	return psguo
}

// SetParentID sets the "parent" edge to the ProductStoreGroup entity by ID.
func (psguo *ProductStoreGroupUpdateOne) SetParentID(id int) *ProductStoreGroupUpdateOne {
	psguo.mutation.SetParentID(id)
	return psguo
}

// SetNillableParentID sets the "parent" edge to the ProductStoreGroup entity by ID if the given value is not nil.
func (psguo *ProductStoreGroupUpdateOne) SetNillableParentID(id *int) *ProductStoreGroupUpdateOne {
	if id != nil {
		psguo = psguo.SetParentID(*id)
	}
	return psguo
}

// SetParent sets the "parent" edge to the ProductStoreGroup entity.
func (psguo *ProductStoreGroupUpdateOne) SetParent(p *ProductStoreGroup) *ProductStoreGroupUpdateOne {
	return psguo.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductStoreGroup entity by IDs.
func (psguo *ProductStoreGroupUpdateOne) AddChildIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.AddChildIDs(ids...)
	return psguo
}

// AddChildren adds the "children" edges to the ProductStoreGroup entity.
func (psguo *ProductStoreGroupUpdateOne) AddChildren(p ...*ProductStoreGroup) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.AddChildIDs(ids...)
}

// AddProductPriceIDs adds the "product_prices" edge to the ProductPrice entity by IDs.
func (psguo *ProductStoreGroupUpdateOne) AddProductPriceIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.AddProductPriceIDs(ids...)
	return psguo
}

// AddProductPrices adds the "product_prices" edges to the ProductPrice entity.
func (psguo *ProductStoreGroupUpdateOne) AddProductPrices(p ...*ProductPrice) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.AddProductPriceIDs(ids...)
}

// AddPrimaryProductStoreIDs adds the "primary_product_stores" edge to the ProductStore entity by IDs.
func (psguo *ProductStoreGroupUpdateOne) AddPrimaryProductStoreIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.AddPrimaryProductStoreIDs(ids...)
	return psguo
}

// AddPrimaryProductStores adds the "primary_product_stores" edges to the ProductStore entity.
func (psguo *ProductStoreGroupUpdateOne) AddPrimaryProductStores(p ...*ProductStore) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.AddPrimaryProductStoreIDs(ids...)
}

// Mutation returns the ProductStoreGroupMutation object of the builder.
func (psguo *ProductStoreGroupUpdateOne) Mutation() *ProductStoreGroupMutation {
	return psguo.mutation
}

// ClearParent clears the "parent" edge to the ProductStoreGroup entity.
func (psguo *ProductStoreGroupUpdateOne) ClearParent() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearParent()
	return psguo
}

// ClearChildren clears all "children" edges to the ProductStoreGroup entity.
func (psguo *ProductStoreGroupUpdateOne) ClearChildren() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearChildren()
	return psguo
}

// RemoveChildIDs removes the "children" edge to ProductStoreGroup entities by IDs.
func (psguo *ProductStoreGroupUpdateOne) RemoveChildIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.RemoveChildIDs(ids...)
	return psguo
}

// RemoveChildren removes "children" edges to ProductStoreGroup entities.
func (psguo *ProductStoreGroupUpdateOne) RemoveChildren(p ...*ProductStoreGroup) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.RemoveChildIDs(ids...)
}

// ClearProductPrices clears all "product_prices" edges to the ProductPrice entity.
func (psguo *ProductStoreGroupUpdateOne) ClearProductPrices() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearProductPrices()
	return psguo
}

// RemoveProductPriceIDs removes the "product_prices" edge to ProductPrice entities by IDs.
func (psguo *ProductStoreGroupUpdateOne) RemoveProductPriceIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.RemoveProductPriceIDs(ids...)
	return psguo
}

// RemoveProductPrices removes "product_prices" edges to ProductPrice entities.
func (psguo *ProductStoreGroupUpdateOne) RemoveProductPrices(p ...*ProductPrice) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.RemoveProductPriceIDs(ids...)
}

// ClearPrimaryProductStores clears all "primary_product_stores" edges to the ProductStore entity.
func (psguo *ProductStoreGroupUpdateOne) ClearPrimaryProductStores() *ProductStoreGroupUpdateOne {
	psguo.mutation.ClearPrimaryProductStores()
	return psguo
}

// RemovePrimaryProductStoreIDs removes the "primary_product_stores" edge to ProductStore entities by IDs.
func (psguo *ProductStoreGroupUpdateOne) RemovePrimaryProductStoreIDs(ids ...int) *ProductStoreGroupUpdateOne {
	psguo.mutation.RemovePrimaryProductStoreIDs(ids...)
	return psguo
}

// RemovePrimaryProductStores removes "primary_product_stores" edges to ProductStore entities.
func (psguo *ProductStoreGroupUpdateOne) RemovePrimaryProductStores(p ...*ProductStore) *ProductStoreGroupUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psguo.RemovePrimaryProductStoreIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psguo *ProductStoreGroupUpdateOne) Select(field string, fields ...string) *ProductStoreGroupUpdateOne {
	psguo.fields = append([]string{field}, fields...)
	return psguo
}

// Save executes the query and returns the updated ProductStoreGroup entity.
func (psguo *ProductStoreGroupUpdateOne) Save(ctx context.Context) (*ProductStoreGroup, error) {
	var (
		err  error
		node *ProductStoreGroup
	)
	psguo.defaults()
	if len(psguo.hooks) == 0 {
		node, err = psguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductStoreGroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psguo.mutation = mutation
			node, err = psguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psguo.hooks) - 1; i >= 0; i-- {
			mut = psguo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psguo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psguo *ProductStoreGroupUpdateOne) SaveX(ctx context.Context) *ProductStoreGroup {
	node, err := psguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psguo *ProductStoreGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := psguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psguo *ProductStoreGroupUpdateOne) ExecX(ctx context.Context) {
	if err := psguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psguo *ProductStoreGroupUpdateOne) defaults() {
	if _, ok := psguo.mutation.UpdateTime(); !ok {
		v := productstoregroup.UpdateDefaultUpdateTime()
		psguo.mutation.SetUpdateTime(v)
	}
}

func (psguo *ProductStoreGroupUpdateOne) sqlSave(ctx context.Context) (_node *ProductStoreGroup, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productstoregroup.Table,
			Columns: productstoregroup.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productstoregroup.FieldID,
			},
		},
	}
	id, ok := psguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProductStoreGroup.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := psguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productstoregroup.FieldID)
		for _, f := range fields {
			if !productstoregroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productstoregroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psguo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productstoregroup.FieldUpdateTime,
		})
	}
	if value, ok := psguo.mutation.StringRef(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldStringRef,
		})
	}
	if psguo.mutation.StringRefCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldStringRef,
		})
	}
	if value, ok := psguo.mutation.ProductStoreGroupTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if value, ok := psguo.mutation.AddedProductStoreGroupTypeID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if psguo.mutation.ProductStoreGroupTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: productstoregroup.FieldProductStoreGroupTypeID,
		})
	}
	if value, ok := psguo.mutation.ProductStoreGroupName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldProductStoreGroupName,
		})
	}
	if psguo.mutation.ProductStoreGroupNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldProductStoreGroupName,
		})
	}
	if value, ok := psguo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productstoregroup.FieldDescription,
		})
	}
	if psguo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productstoregroup.FieldDescription,
		})
	}
	if psguo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productstoregroup.ParentTable,
			Columns: []string{productstoregroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productstoregroup.ParentTable,
			Columns: []string{productstoregroup.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psguo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !psguo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ChildrenTable,
			Columns: []string{productstoregroup.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstoregroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psguo.mutation.ProductPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.RemovedProductPricesIDs(); len(nodes) > 0 && !psguo.mutation.ProductPricesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.ProductPricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.ProductPricesTable,
			Columns: []string{productstoregroup.ProductPricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productprice.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psguo.mutation.PrimaryProductStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.RemovedPrimaryProductStoresIDs(); len(nodes) > 0 && !psguo.mutation.PrimaryProductStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psguo.mutation.PrimaryProductStoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productstoregroup.PrimaryProductStoresTable,
			Columns: []string{productstoregroup.PrimaryProductStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productstore.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProductStoreGroup{config: psguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productstoregroup.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
