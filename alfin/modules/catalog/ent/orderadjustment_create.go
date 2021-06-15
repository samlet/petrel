// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderadjustment"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroup"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitemshipgroupassoc"
)

// OrderAdjustmentCreate is the builder for creating a OrderAdjustment entity.
type OrderAdjustmentCreate struct {
	config
	mutation *OrderAdjustmentMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (oac *OrderAdjustmentCreate) SetCreateTime(t time.Time) *OrderAdjustmentCreate {
	oac.mutation.SetCreateTime(t)
	return oac
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableCreateTime(t *time.Time) *OrderAdjustmentCreate {
	if t != nil {
		oac.SetCreateTime(*t)
	}
	return oac
}

// SetUpdateTime sets the "update_time" field.
func (oac *OrderAdjustmentCreate) SetUpdateTime(t time.Time) *OrderAdjustmentCreate {
	oac.mutation.SetUpdateTime(t)
	return oac
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableUpdateTime(t *time.Time) *OrderAdjustmentCreate {
	if t != nil {
		oac.SetUpdateTime(*t)
	}
	return oac
}

// SetStringRef sets the "string_ref" field.
func (oac *OrderAdjustmentCreate) SetStringRef(s string) *OrderAdjustmentCreate {
	oac.mutation.SetStringRef(s)
	return oac
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableStringRef(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetStringRef(*s)
	}
	return oac
}

// SetOrderAdjustmentTypeID sets the "order_adjustment_type_id" field.
func (oac *OrderAdjustmentCreate) SetOrderAdjustmentTypeID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderAdjustmentTypeID(i)
	return oac
}

// SetNillableOrderAdjustmentTypeID sets the "order_adjustment_type_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderAdjustmentTypeID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetOrderAdjustmentTypeID(*i)
	}
	return oac
}

// SetOrderItemSeqID sets the "order_item_seq_id" field.
func (oac *OrderAdjustmentCreate) SetOrderItemSeqID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderItemSeqID(i)
	return oac
}

// SetNillableOrderItemSeqID sets the "order_item_seq_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderItemSeqID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetOrderItemSeqID(*i)
	}
	return oac
}

// SetShipGroupSeqID sets the "ship_group_seq_id" field.
func (oac *OrderAdjustmentCreate) SetShipGroupSeqID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetShipGroupSeqID(i)
	return oac
}

// SetNillableShipGroupSeqID sets the "ship_group_seq_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableShipGroupSeqID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetShipGroupSeqID(*i)
	}
	return oac
}

// SetComments sets the "comments" field.
func (oac *OrderAdjustmentCreate) SetComments(s string) *OrderAdjustmentCreate {
	oac.mutation.SetComments(s)
	return oac
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableComments(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetComments(*s)
	}
	return oac
}

// SetDescription sets the "description" field.
func (oac *OrderAdjustmentCreate) SetDescription(s string) *OrderAdjustmentCreate {
	oac.mutation.SetDescription(s)
	return oac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableDescription(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetDescription(*s)
	}
	return oac
}

// SetAmount sets the "amount" field.
func (oac *OrderAdjustmentCreate) SetAmount(f float64) *OrderAdjustmentCreate {
	oac.mutation.SetAmount(f)
	return oac
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableAmount(f *float64) *OrderAdjustmentCreate {
	if f != nil {
		oac.SetAmount(*f)
	}
	return oac
}

// SetRecurringAmount sets the "recurring_amount" field.
func (oac *OrderAdjustmentCreate) SetRecurringAmount(f float64) *OrderAdjustmentCreate {
	oac.mutation.SetRecurringAmount(f)
	return oac
}

// SetNillableRecurringAmount sets the "recurring_amount" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableRecurringAmount(f *float64) *OrderAdjustmentCreate {
	if f != nil {
		oac.SetRecurringAmount(*f)
	}
	return oac
}

// SetAmountAlreadyIncluded sets the "amount_already_included" field.
func (oac *OrderAdjustmentCreate) SetAmountAlreadyIncluded(f float64) *OrderAdjustmentCreate {
	oac.mutation.SetAmountAlreadyIncluded(f)
	return oac
}

// SetNillableAmountAlreadyIncluded sets the "amount_already_included" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableAmountAlreadyIncluded(f *float64) *OrderAdjustmentCreate {
	if f != nil {
		oac.SetAmountAlreadyIncluded(*f)
	}
	return oac
}

// SetProductPromoID sets the "product_promo_id" field.
func (oac *OrderAdjustmentCreate) SetProductPromoID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetProductPromoID(i)
	return oac
}

// SetNillableProductPromoID sets the "product_promo_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableProductPromoID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetProductPromoID(*i)
	}
	return oac
}

// SetProductPromoRuleID sets the "product_promo_rule_id" field.
func (oac *OrderAdjustmentCreate) SetProductPromoRuleID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetProductPromoRuleID(i)
	return oac
}

// SetNillableProductPromoRuleID sets the "product_promo_rule_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableProductPromoRuleID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetProductPromoRuleID(*i)
	}
	return oac
}

// SetProductPromoActionSeqID sets the "product_promo_action_seq_id" field.
func (oac *OrderAdjustmentCreate) SetProductPromoActionSeqID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetProductPromoActionSeqID(i)
	return oac
}

// SetNillableProductPromoActionSeqID sets the "product_promo_action_seq_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableProductPromoActionSeqID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetProductPromoActionSeqID(*i)
	}
	return oac
}

// SetProductFeatureID sets the "product_feature_id" field.
func (oac *OrderAdjustmentCreate) SetProductFeatureID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetProductFeatureID(i)
	return oac
}

// SetNillableProductFeatureID sets the "product_feature_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableProductFeatureID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetProductFeatureID(*i)
	}
	return oac
}

// SetCorrespondingProductID sets the "corresponding_product_id" field.
func (oac *OrderAdjustmentCreate) SetCorrespondingProductID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetCorrespondingProductID(i)
	return oac
}

// SetNillableCorrespondingProductID sets the "corresponding_product_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableCorrespondingProductID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetCorrespondingProductID(*i)
	}
	return oac
}

// SetTaxAuthorityRateSeqID sets the "tax_authority_rate_seq_id" field.
func (oac *OrderAdjustmentCreate) SetTaxAuthorityRateSeqID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetTaxAuthorityRateSeqID(i)
	return oac
}

// SetNillableTaxAuthorityRateSeqID sets the "tax_authority_rate_seq_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableTaxAuthorityRateSeqID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetTaxAuthorityRateSeqID(*i)
	}
	return oac
}

// SetSourceReferenceID sets the "source_reference_id" field.
func (oac *OrderAdjustmentCreate) SetSourceReferenceID(s string) *OrderAdjustmentCreate {
	oac.mutation.SetSourceReferenceID(s)
	return oac
}

// SetNillableSourceReferenceID sets the "source_reference_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableSourceReferenceID(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetSourceReferenceID(*s)
	}
	return oac
}

// SetSourcePercentage sets the "source_percentage" field.
func (oac *OrderAdjustmentCreate) SetSourcePercentage(f float64) *OrderAdjustmentCreate {
	oac.mutation.SetSourcePercentage(f)
	return oac
}

// SetNillableSourcePercentage sets the "source_percentage" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableSourcePercentage(f *float64) *OrderAdjustmentCreate {
	if f != nil {
		oac.SetSourcePercentage(*f)
	}
	return oac
}

// SetCustomerReferenceID sets the "customer_reference_id" field.
func (oac *OrderAdjustmentCreate) SetCustomerReferenceID(s string) *OrderAdjustmentCreate {
	oac.mutation.SetCustomerReferenceID(s)
	return oac
}

// SetNillableCustomerReferenceID sets the "customer_reference_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableCustomerReferenceID(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetCustomerReferenceID(*s)
	}
	return oac
}

// SetPrimaryGeoID sets the "primary_geo_id" field.
func (oac *OrderAdjustmentCreate) SetPrimaryGeoID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetPrimaryGeoID(i)
	return oac
}

// SetNillablePrimaryGeoID sets the "primary_geo_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillablePrimaryGeoID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetPrimaryGeoID(*i)
	}
	return oac
}

// SetSecondaryGeoID sets the "secondary_geo_id" field.
func (oac *OrderAdjustmentCreate) SetSecondaryGeoID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetSecondaryGeoID(i)
	return oac
}

// SetNillableSecondaryGeoID sets the "secondary_geo_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableSecondaryGeoID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetSecondaryGeoID(*i)
	}
	return oac
}

// SetExemptAmount sets the "exempt_amount" field.
func (oac *OrderAdjustmentCreate) SetExemptAmount(f float64) *OrderAdjustmentCreate {
	oac.mutation.SetExemptAmount(f)
	return oac
}

// SetNillableExemptAmount sets the "exempt_amount" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableExemptAmount(f *float64) *OrderAdjustmentCreate {
	if f != nil {
		oac.SetExemptAmount(*f)
	}
	return oac
}

// SetTaxAuthGeoID sets the "tax_auth_geo_id" field.
func (oac *OrderAdjustmentCreate) SetTaxAuthGeoID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetTaxAuthGeoID(i)
	return oac
}

// SetNillableTaxAuthGeoID sets the "tax_auth_geo_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableTaxAuthGeoID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetTaxAuthGeoID(*i)
	}
	return oac
}

// SetTaxAuthPartyID sets the "tax_auth_party_id" field.
func (oac *OrderAdjustmentCreate) SetTaxAuthPartyID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetTaxAuthPartyID(i)
	return oac
}

// SetNillableTaxAuthPartyID sets the "tax_auth_party_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableTaxAuthPartyID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetTaxAuthPartyID(*i)
	}
	return oac
}

// SetOverrideGlAccountID sets the "override_gl_account_id" field.
func (oac *OrderAdjustmentCreate) SetOverrideGlAccountID(i int) *OrderAdjustmentCreate {
	oac.mutation.SetOverrideGlAccountID(i)
	return oac
}

// SetNillableOverrideGlAccountID sets the "override_gl_account_id" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOverrideGlAccountID(i *int) *OrderAdjustmentCreate {
	if i != nil {
		oac.SetOverrideGlAccountID(*i)
	}
	return oac
}

// SetIncludeInTax sets the "include_in_tax" field.
func (oac *OrderAdjustmentCreate) SetIncludeInTax(oit orderadjustment.IncludeInTax) *OrderAdjustmentCreate {
	oac.mutation.SetIncludeInTax(oit)
	return oac
}

// SetNillableIncludeInTax sets the "include_in_tax" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableIncludeInTax(oit *orderadjustment.IncludeInTax) *OrderAdjustmentCreate {
	if oit != nil {
		oac.SetIncludeInTax(*oit)
	}
	return oac
}

// SetIncludeInShipping sets the "include_in_shipping" field.
func (oac *OrderAdjustmentCreate) SetIncludeInShipping(ois orderadjustment.IncludeInShipping) *OrderAdjustmentCreate {
	oac.mutation.SetIncludeInShipping(ois)
	return oac
}

// SetNillableIncludeInShipping sets the "include_in_shipping" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableIncludeInShipping(ois *orderadjustment.IncludeInShipping) *OrderAdjustmentCreate {
	if ois != nil {
		oac.SetIncludeInShipping(*ois)
	}
	return oac
}

// SetIsManual sets the "is_manual" field.
func (oac *OrderAdjustmentCreate) SetIsManual(om orderadjustment.IsManual) *OrderAdjustmentCreate {
	oac.mutation.SetIsManual(om)
	return oac
}

// SetNillableIsManual sets the "is_manual" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableIsManual(om *orderadjustment.IsManual) *OrderAdjustmentCreate {
	if om != nil {
		oac.SetIsManual(*om)
	}
	return oac
}

// SetCreatedDate sets the "created_date" field.
func (oac *OrderAdjustmentCreate) SetCreatedDate(t time.Time) *OrderAdjustmentCreate {
	oac.mutation.SetCreatedDate(t)
	return oac
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableCreatedDate(t *time.Time) *OrderAdjustmentCreate {
	if t != nil {
		oac.SetCreatedDate(*t)
	}
	return oac
}

// SetCreatedByUserLogin sets the "created_by_user_login" field.
func (oac *OrderAdjustmentCreate) SetCreatedByUserLogin(s string) *OrderAdjustmentCreate {
	oac.mutation.SetCreatedByUserLogin(s)
	return oac
}

// SetNillableCreatedByUserLogin sets the "created_by_user_login" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableCreatedByUserLogin(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetCreatedByUserLogin(*s)
	}
	return oac
}

// SetLastModifiedDate sets the "last_modified_date" field.
func (oac *OrderAdjustmentCreate) SetLastModifiedDate(t time.Time) *OrderAdjustmentCreate {
	oac.mutation.SetLastModifiedDate(t)
	return oac
}

// SetNillableLastModifiedDate sets the "last_modified_date" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableLastModifiedDate(t *time.Time) *OrderAdjustmentCreate {
	if t != nil {
		oac.SetLastModifiedDate(*t)
	}
	return oac
}

// SetLastModifiedByUserLogin sets the "last_modified_by_user_login" field.
func (oac *OrderAdjustmentCreate) SetLastModifiedByUserLogin(s string) *OrderAdjustmentCreate {
	oac.mutation.SetLastModifiedByUserLogin(s)
	return oac
}

// SetNillableLastModifiedByUserLogin sets the "last_modified_by_user_login" field if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableLastModifiedByUserLogin(s *string) *OrderAdjustmentCreate {
	if s != nil {
		oac.SetLastModifiedByUserLogin(*s)
	}
	return oac
}

// SetOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID.
func (oac *OrderAdjustmentCreate) SetOrderHeaderID(id int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderHeaderID(id)
	return oac
}

// SetNillableOrderHeaderID sets the "order_header" edge to the OrderHeader entity by ID if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderHeaderID(id *int) *OrderAdjustmentCreate {
	if id != nil {
		oac = oac.SetOrderHeaderID(*id)
	}
	return oac
}

// SetOrderHeader sets the "order_header" edge to the OrderHeader entity.
func (oac *OrderAdjustmentCreate) SetOrderHeader(o *OrderHeader) *OrderAdjustmentCreate {
	return oac.SetOrderHeaderID(o.ID)
}

// SetOrderItemID sets the "order_item" edge to the OrderItem entity by ID.
func (oac *OrderAdjustmentCreate) SetOrderItemID(id int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderItemID(id)
	return oac
}

// SetNillableOrderItemID sets the "order_item" edge to the OrderItem entity by ID if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderItemID(id *int) *OrderAdjustmentCreate {
	if id != nil {
		oac = oac.SetOrderItemID(*id)
	}
	return oac
}

// SetOrderItem sets the "order_item" edge to the OrderItem entity.
func (oac *OrderAdjustmentCreate) SetOrderItem(o *OrderItem) *OrderAdjustmentCreate {
	return oac.SetOrderItemID(o.ID)
}

// SetOrderItemShipGroupID sets the "order_item_ship_group" edge to the OrderItemShipGroup entity by ID.
func (oac *OrderAdjustmentCreate) SetOrderItemShipGroupID(id int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderItemShipGroupID(id)
	return oac
}

// SetNillableOrderItemShipGroupID sets the "order_item_ship_group" edge to the OrderItemShipGroup entity by ID if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderItemShipGroupID(id *int) *OrderAdjustmentCreate {
	if id != nil {
		oac = oac.SetOrderItemShipGroupID(*id)
	}
	return oac
}

// SetOrderItemShipGroup sets the "order_item_ship_group" edge to the OrderItemShipGroup entity.
func (oac *OrderAdjustmentCreate) SetOrderItemShipGroup(o *OrderItemShipGroup) *OrderAdjustmentCreate {
	return oac.SetOrderItemShipGroupID(o.ID)
}

// SetOrderItemShipGroupAssocID sets the "order_item_ship_group_assoc" edge to the OrderItemShipGroupAssoc entity by ID.
func (oac *OrderAdjustmentCreate) SetOrderItemShipGroupAssocID(id int) *OrderAdjustmentCreate {
	oac.mutation.SetOrderItemShipGroupAssocID(id)
	return oac
}

// SetNillableOrderItemShipGroupAssocID sets the "order_item_ship_group_assoc" edge to the OrderItemShipGroupAssoc entity by ID if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableOrderItemShipGroupAssocID(id *int) *OrderAdjustmentCreate {
	if id != nil {
		oac = oac.SetOrderItemShipGroupAssocID(*id)
	}
	return oac
}

// SetOrderItemShipGroupAssoc sets the "order_item_ship_group_assoc" edge to the OrderItemShipGroupAssoc entity.
func (oac *OrderAdjustmentCreate) SetOrderItemShipGroupAssoc(o *OrderItemShipGroupAssoc) *OrderAdjustmentCreate {
	return oac.SetOrderItemShipGroupAssocID(o.ID)
}

// SetParentID sets the "parent" edge to the OrderAdjustment entity by ID.
func (oac *OrderAdjustmentCreate) SetParentID(id int) *OrderAdjustmentCreate {
	oac.mutation.SetParentID(id)
	return oac
}

// SetNillableParentID sets the "parent" edge to the OrderAdjustment entity by ID if the given value is not nil.
func (oac *OrderAdjustmentCreate) SetNillableParentID(id *int) *OrderAdjustmentCreate {
	if id != nil {
		oac = oac.SetParentID(*id)
	}
	return oac
}

// SetParent sets the "parent" edge to the OrderAdjustment entity.
func (oac *OrderAdjustmentCreate) SetParent(o *OrderAdjustment) *OrderAdjustmentCreate {
	return oac.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the OrderAdjustment entity by IDs.
func (oac *OrderAdjustmentCreate) AddChildIDs(ids ...int) *OrderAdjustmentCreate {
	oac.mutation.AddChildIDs(ids...)
	return oac
}

// AddChildren adds the "children" edges to the OrderAdjustment entity.
func (oac *OrderAdjustmentCreate) AddChildren(o ...*OrderAdjustment) *OrderAdjustmentCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oac.AddChildIDs(ids...)
}

// Mutation returns the OrderAdjustmentMutation object of the builder.
func (oac *OrderAdjustmentCreate) Mutation() *OrderAdjustmentMutation {
	return oac.mutation
}

// Save creates the OrderAdjustment in the database.
func (oac *OrderAdjustmentCreate) Save(ctx context.Context) (*OrderAdjustment, error) {
	var (
		err  error
		node *OrderAdjustment
	)
	oac.defaults()
	if len(oac.hooks) == 0 {
		if err = oac.check(); err != nil {
			return nil, err
		}
		node, err = oac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderAdjustmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oac.check(); err != nil {
				return nil, err
			}
			oac.mutation = mutation
			if node, err = oac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oac.hooks) - 1; i >= 0; i-- {
			mut = oac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OrderAdjustmentCreate) SaveX(ctx context.Context) *OrderAdjustment {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (oac *OrderAdjustmentCreate) defaults() {
	if _, ok := oac.mutation.CreateTime(); !ok {
		v := orderadjustment.DefaultCreateTime()
		oac.mutation.SetCreateTime(v)
	}
	if _, ok := oac.mutation.UpdateTime(); !ok {
		v := orderadjustment.DefaultUpdateTime()
		oac.mutation.SetUpdateTime(v)
	}
	if _, ok := oac.mutation.CreatedDate(); !ok {
		v := orderadjustment.DefaultCreatedDate()
		oac.mutation.SetCreatedDate(v)
	}
	if _, ok := oac.mutation.LastModifiedDate(); !ok {
		v := orderadjustment.DefaultLastModifiedDate()
		oac.mutation.SetLastModifiedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oac *OrderAdjustmentCreate) check() error {
	if _, ok := oac.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := oac.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := oac.mutation.SourceReferenceID(); ok {
		if err := orderadjustment.SourceReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "source_reference_id", err: fmt.Errorf("ent: validator failed for field \"source_reference_id\": %w", err)}
		}
	}
	if v, ok := oac.mutation.CustomerReferenceID(); ok {
		if err := orderadjustment.CustomerReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "customer_reference_id", err: fmt.Errorf("ent: validator failed for field \"customer_reference_id\": %w", err)}
		}
	}
	if v, ok := oac.mutation.IncludeInTax(); ok {
		if err := orderadjustment.IncludeInTaxValidator(v); err != nil {
			return &ValidationError{Name: "include_in_tax", err: fmt.Errorf("ent: validator failed for field \"include_in_tax\": %w", err)}
		}
	}
	if v, ok := oac.mutation.IncludeInShipping(); ok {
		if err := orderadjustment.IncludeInShippingValidator(v); err != nil {
			return &ValidationError{Name: "include_in_shipping", err: fmt.Errorf("ent: validator failed for field \"include_in_shipping\": %w", err)}
		}
	}
	if v, ok := oac.mutation.IsManual(); ok {
		if err := orderadjustment.IsManualValidator(v); err != nil {
			return &ValidationError{Name: "is_manual", err: fmt.Errorf("ent: validator failed for field \"is_manual\": %w", err)}
		}
	}
	return nil
}

func (oac *OrderAdjustmentCreate) sqlSave(ctx context.Context) (*OrderAdjustment, error) {
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oac *OrderAdjustmentCreate) createSpec() (*OrderAdjustment, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderAdjustment{config: oac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderadjustment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderadjustment.FieldID,
			},
		}
	)
	if value, ok := oac.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderadjustment.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oac.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderadjustment.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := oac.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := oac.mutation.OrderAdjustmentTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldOrderAdjustmentTypeID,
		})
		_node.OrderAdjustmentTypeID = value
	}
	if value, ok := oac.mutation.OrderItemSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldOrderItemSeqID,
		})
		_node.OrderItemSeqID = value
	}
	if value, ok := oac.mutation.ShipGroupSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldShipGroupSeqID,
		})
		_node.ShipGroupSeqID = value
	}
	if value, ok := oac.mutation.Comments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldComments,
		})
		_node.Comments = value
	}
	if value, ok := oac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := oac.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderadjustment.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := oac.mutation.RecurringAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderadjustment.FieldRecurringAmount,
		})
		_node.RecurringAmount = value
	}
	if value, ok := oac.mutation.AmountAlreadyIncluded(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderadjustment.FieldAmountAlreadyIncluded,
		})
		_node.AmountAlreadyIncluded = value
	}
	if value, ok := oac.mutation.ProductPromoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldProductPromoID,
		})
		_node.ProductPromoID = value
	}
	if value, ok := oac.mutation.ProductPromoRuleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldProductPromoRuleID,
		})
		_node.ProductPromoRuleID = value
	}
	if value, ok := oac.mutation.ProductPromoActionSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldProductPromoActionSeqID,
		})
		_node.ProductPromoActionSeqID = value
	}
	if value, ok := oac.mutation.ProductFeatureID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldProductFeatureID,
		})
		_node.ProductFeatureID = value
	}
	if value, ok := oac.mutation.CorrespondingProductID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldCorrespondingProductID,
		})
		_node.CorrespondingProductID = value
	}
	if value, ok := oac.mutation.TaxAuthorityRateSeqID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldTaxAuthorityRateSeqID,
		})
		_node.TaxAuthorityRateSeqID = value
	}
	if value, ok := oac.mutation.SourceReferenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldSourceReferenceID,
		})
		_node.SourceReferenceID = value
	}
	if value, ok := oac.mutation.SourcePercentage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderadjustment.FieldSourcePercentage,
		})
		_node.SourcePercentage = value
	}
	if value, ok := oac.mutation.CustomerReferenceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldCustomerReferenceID,
		})
		_node.CustomerReferenceID = value
	}
	if value, ok := oac.mutation.PrimaryGeoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldPrimaryGeoID,
		})
		_node.PrimaryGeoID = value
	}
	if value, ok := oac.mutation.SecondaryGeoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldSecondaryGeoID,
		})
		_node.SecondaryGeoID = value
	}
	if value, ok := oac.mutation.ExemptAmount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: orderadjustment.FieldExemptAmount,
		})
		_node.ExemptAmount = value
	}
	if value, ok := oac.mutation.TaxAuthGeoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldTaxAuthGeoID,
		})
		_node.TaxAuthGeoID = value
	}
	if value, ok := oac.mutation.TaxAuthPartyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldTaxAuthPartyID,
		})
		_node.TaxAuthPartyID = value
	}
	if value, ok := oac.mutation.OverrideGlAccountID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderadjustment.FieldOverrideGlAccountID,
		})
		_node.OverrideGlAccountID = value
	}
	if value, ok := oac.mutation.IncludeInTax(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: orderadjustment.FieldIncludeInTax,
		})
		_node.IncludeInTax = value
	}
	if value, ok := oac.mutation.IncludeInShipping(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: orderadjustment.FieldIncludeInShipping,
		})
		_node.IncludeInShipping = value
	}
	if value, ok := oac.mutation.IsManual(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: orderadjustment.FieldIsManual,
		})
		_node.IsManual = value
	}
	if value, ok := oac.mutation.CreatedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderadjustment.FieldCreatedDate,
		})
		_node.CreatedDate = value
	}
	if value, ok := oac.mutation.CreatedByUserLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldCreatedByUserLogin,
		})
		_node.CreatedByUserLogin = value
	}
	if value, ok := oac.mutation.LastModifiedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderadjustment.FieldLastModifiedDate,
		})
		_node.LastModifiedDate = value
	}
	if value, ok := oac.mutation.LastModifiedByUserLogin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderadjustment.FieldLastModifiedByUserLogin,
		})
		_node.LastModifiedByUserLogin = value
	}
	if nodes := oac.mutation.OrderHeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderadjustment.OrderHeaderTable,
			Columns: []string{orderadjustment.OrderHeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderheader.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_header_order_adjustments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.OrderItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderadjustment.OrderItemTable,
			Columns: []string{orderadjustment.OrderItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_order_adjustments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.OrderItemShipGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderadjustment.OrderItemShipGroupTable,
			Columns: []string{orderadjustment.OrderItemShipGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitemshipgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_ship_group_order_adjustments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.OrderItemShipGroupAssocIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderadjustment.OrderItemShipGroupAssocTable,
			Columns: []string{orderadjustment.OrderItemShipGroupAssocColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderitemshipgroupassoc.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_ship_group_assoc_order_adjustments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderadjustment.ParentTable,
			Columns: []string{orderadjustment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderadjustment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_adjustment_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderadjustment.ChildrenTable,
			Columns: []string{orderadjustment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderadjustment.FieldID,
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

// OrderAdjustmentCreateBulk is the builder for creating many OrderAdjustment entities in bulk.
type OrderAdjustmentCreateBulk struct {
	config
	builders []*OrderAdjustmentCreate
}

// Save creates the OrderAdjustment entities in the database.
func (oacb *OrderAdjustmentCreateBulk) Save(ctx context.Context) ([]*OrderAdjustment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OrderAdjustment, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderAdjustmentMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OrderAdjustmentCreateBulk) SaveX(ctx context.Context) []*OrderAdjustment {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
