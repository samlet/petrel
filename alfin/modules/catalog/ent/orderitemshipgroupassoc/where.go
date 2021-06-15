// Code generated by entc, DO NOT EDIT.

package orderitemshipgroupassoc

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// OrderItemSeqID applies equality check predicate on the "order_item_seq_id" field. It's identical to OrderItemSeqIDEQ.
func OrderItemSeqID(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderItemSeqID), v))
	})
}

// ShipGroupSeqID applies equality check predicate on the "ship_group_seq_id" field. It's identical to ShipGroupSeqIDEQ.
func ShipGroupSeqID(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipGroupSeqID), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// CancelQuantity applies equality check predicate on the "cancel_quantity" field. It's identical to CancelQuantityEQ.
func CancelQuantity(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelQuantity), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStringRef), v...))
	})
}

// StringRefNotIn applies the NotIn predicate on the "string_ref" field.
func StringRefNotIn(vs ...string) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStringRef), v...))
	})
}

// StringRefGT applies the GT predicate on the "string_ref" field.
func StringRefGT(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// OrderItemSeqIDEQ applies the EQ predicate on the "order_item_seq_id" field.
func OrderItemSeqIDEQ(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderItemSeqID), v))
	})
}

// OrderItemSeqIDNEQ applies the NEQ predicate on the "order_item_seq_id" field.
func OrderItemSeqIDNEQ(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderItemSeqID), v))
	})
}

// OrderItemSeqIDIn applies the In predicate on the "order_item_seq_id" field.
func OrderItemSeqIDIn(vs ...int) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderItemSeqID), v...))
	})
}

// OrderItemSeqIDNotIn applies the NotIn predicate on the "order_item_seq_id" field.
func OrderItemSeqIDNotIn(vs ...int) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderItemSeqID), v...))
	})
}

// OrderItemSeqIDGT applies the GT predicate on the "order_item_seq_id" field.
func OrderItemSeqIDGT(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderItemSeqID), v))
	})
}

// OrderItemSeqIDGTE applies the GTE predicate on the "order_item_seq_id" field.
func OrderItemSeqIDGTE(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderItemSeqID), v))
	})
}

// OrderItemSeqIDLT applies the LT predicate on the "order_item_seq_id" field.
func OrderItemSeqIDLT(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderItemSeqID), v))
	})
}

// OrderItemSeqIDLTE applies the LTE predicate on the "order_item_seq_id" field.
func OrderItemSeqIDLTE(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderItemSeqID), v))
	})
}

// ShipGroupSeqIDEQ applies the EQ predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDEQ(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipGroupSeqID), v))
	})
}

// ShipGroupSeqIDNEQ applies the NEQ predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDNEQ(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipGroupSeqID), v))
	})
}

// ShipGroupSeqIDIn applies the In predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDIn(vs ...int) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShipGroupSeqID), v...))
	})
}

// ShipGroupSeqIDNotIn applies the NotIn predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDNotIn(vs ...int) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShipGroupSeqID), v...))
	})
}

// ShipGroupSeqIDGT applies the GT predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDGT(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipGroupSeqID), v))
	})
}

// ShipGroupSeqIDGTE applies the GTE predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDGTE(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipGroupSeqID), v))
	})
}

// ShipGroupSeqIDLT applies the LT predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDLT(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipGroupSeqID), v))
	})
}

// ShipGroupSeqIDLTE applies the LTE predicate on the "ship_group_seq_id" field.
func ShipGroupSeqIDLTE(v int) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipGroupSeqID), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...float64) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...float64) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// QuantityIsNil applies the IsNil predicate on the "quantity" field.
func QuantityIsNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldQuantity)))
	})
}

// QuantityNotNil applies the NotNil predicate on the "quantity" field.
func QuantityNotNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldQuantity)))
	})
}

// CancelQuantityEQ applies the EQ predicate on the "cancel_quantity" field.
func CancelQuantityEQ(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityNEQ applies the NEQ predicate on the "cancel_quantity" field.
func CancelQuantityNEQ(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityIn applies the In predicate on the "cancel_quantity" field.
func CancelQuantityIn(vs ...float64) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCancelQuantity), v...))
	})
}

// CancelQuantityNotIn applies the NotIn predicate on the "cancel_quantity" field.
func CancelQuantityNotIn(vs ...float64) predicate.OrderItemShipGroupAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCancelQuantity), v...))
	})
}

// CancelQuantityGT applies the GT predicate on the "cancel_quantity" field.
func CancelQuantityGT(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityGTE applies the GTE predicate on the "cancel_quantity" field.
func CancelQuantityGTE(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityLT applies the LT predicate on the "cancel_quantity" field.
func CancelQuantityLT(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityLTE applies the LTE predicate on the "cancel_quantity" field.
func CancelQuantityLTE(v float64) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCancelQuantity), v))
	})
}

// CancelQuantityIsNil applies the IsNil predicate on the "cancel_quantity" field.
func CancelQuantityIsNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCancelQuantity)))
	})
}

// CancelQuantityNotNil applies the NotNil predicate on the "cancel_quantity" field.
func CancelQuantityNotNil() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCancelQuantity)))
	})
}

// HasOrderHeader applies the HasEdge predicate on the "order_header" edge.
func HasOrderHeader() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderHeaderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderHeaderTable, OrderHeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderHeaderWith applies the HasEdge predicate on the "order_header" edge with a given conditions (other predicates).
func HasOrderHeaderWith(preds ...predicate.OrderHeader) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderHeaderInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderHeaderTable, OrderHeaderColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItem applies the HasEdge predicate on the "order_item" edge.
func HasOrderItem() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderItemTable, OrderItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemWith applies the HasEdge predicate on the "order_item" edge with a given conditions (other predicates).
func HasOrderItemWith(preds ...predicate.OrderItem) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderItemTable, OrderItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItemShipGroup applies the HasEdge predicate on the "order_item_ship_group" edge.
func HasOrderItemShipGroup() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemShipGroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderItemShipGroupTable, OrderItemShipGroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemShipGroupWith applies the HasEdge predicate on the "order_item_ship_group" edge with a given conditions (other predicates).
func HasOrderItemShipGroupWith(preds ...predicate.OrderItemShipGroup) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemShipGroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderItemShipGroupTable, OrderItemShipGroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderAdjustments applies the HasEdge predicate on the "order_adjustments" edge.
func HasOrderAdjustments() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderAdjustmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderAdjustmentsTable, OrderAdjustmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderAdjustmentsWith applies the HasEdge predicate on the "order_adjustments" edge with a given conditions (other predicates).
func HasOrderAdjustmentsWith(preds ...predicate.OrderAdjustment) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderAdjustmentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderAdjustmentsTable, OrderAdjustmentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrderItemShipGrpInvRes applies the HasEdge predicate on the "order_item_ship_grp_inv_res" edge.
func HasOrderItemShipGrpInvRes() predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemShipGrpInvResTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderItemShipGrpInvResTable, OrderItemShipGrpInvResColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderItemShipGrpInvResWith applies the HasEdge predicate on the "order_item_ship_grp_inv_res" edge with a given conditions (other predicates).
func HasOrderItemShipGrpInvResWith(preds ...predicate.OrderItemShipGrpInvRes) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderItemShipGrpInvResInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OrderItemShipGrpInvResTable, OrderItemShipGrpInvResColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItemShipGroupAssoc) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItemShipGroupAssoc) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderItemShipGroupAssoc) predicate.OrderItemShipGroupAssoc {
	return predicate.OrderItemShipGroupAssoc(func(s *sql.Selector) {
		p(s.Not())
	})
}
