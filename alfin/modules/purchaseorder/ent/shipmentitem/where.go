// Code generated by entc, DO NOT EDIT.

package shipmentitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/purchaseorder/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ShipmentItemSeqID applies equality check predicate on the "shipment_item_seq_id" field. It's identical to ShipmentItemSeqIDEQ.
func ShipmentItemSeqID(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentItemSeqID), v))
	})
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductID), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// ShipmentContentDescription applies equality check predicate on the "shipment_content_description" field. It's identical to ShipmentContentDescriptionEQ.
func ShipmentContentDescription(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentContentDescription), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ShipmentItemSeqIDEQ applies the EQ predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDEQ(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentItemSeqID), v))
	})
}

// ShipmentItemSeqIDNEQ applies the NEQ predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDNEQ(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipmentItemSeqID), v))
	})
}

// ShipmentItemSeqIDIn applies the In predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDIn(vs ...int) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShipmentItemSeqID), v...))
	})
}

// ShipmentItemSeqIDNotIn applies the NotIn predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDNotIn(vs ...int) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShipmentItemSeqID), v...))
	})
}

// ShipmentItemSeqIDGT applies the GT predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDGT(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipmentItemSeqID), v))
	})
}

// ShipmentItemSeqIDGTE applies the GTE predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDGTE(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipmentItemSeqID), v))
	})
}

// ShipmentItemSeqIDLT applies the LT predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDLT(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipmentItemSeqID), v))
	})
}

// ShipmentItemSeqIDLTE applies the LTE predicate on the "shipment_item_seq_id" field.
func ShipmentItemSeqIDLTE(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipmentItemSeqID), v))
	})
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductID), v))
	})
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductID), v))
	})
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...int) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductID), v...))
	})
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...int) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductID), v...))
	})
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProductID), v))
	})
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProductID), v))
	})
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProductID), v))
	})
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v int) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProductID), v))
	})
}

// ProductIDIsNil applies the IsNil predicate on the "product_id" field.
func ProductIDIsNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProductID)))
	})
}

// ProductIDNotNil applies the NotNil predicate on the "product_id" field.
func ProductIDNotNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProductID)))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...float64) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func QuantityNotIn(vs ...float64) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func QuantityGT(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v float64) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// QuantityIsNil applies the IsNil predicate on the "quantity" field.
func QuantityIsNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldQuantity)))
	})
}

// QuantityNotNil applies the NotNil predicate on the "quantity" field.
func QuantityNotNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldQuantity)))
	})
}

// ShipmentContentDescriptionEQ applies the EQ predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionEQ(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionNEQ applies the NEQ predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionNEQ(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionIn applies the In predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionIn(vs ...string) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShipmentContentDescription), v...))
	})
}

// ShipmentContentDescriptionNotIn applies the NotIn predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionNotIn(vs ...string) predicate.ShipmentItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShipmentItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShipmentContentDescription), v...))
	})
}

// ShipmentContentDescriptionGT applies the GT predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionGT(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionGTE applies the GTE predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionGTE(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionLT applies the LT predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionLT(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionLTE applies the LTE predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionLTE(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionContains applies the Contains predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionContains(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionHasPrefix applies the HasPrefix predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionHasPrefix(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionHasSuffix applies the HasSuffix predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionHasSuffix(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionIsNil applies the IsNil predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionIsNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShipmentContentDescription)))
	})
}

// ShipmentContentDescriptionNotNil applies the NotNil predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionNotNil() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShipmentContentDescription)))
	})
}

// ShipmentContentDescriptionEqualFold applies the EqualFold predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionEqualFold(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShipmentContentDescription), v))
	})
}

// ShipmentContentDescriptionContainsFold applies the ContainsFold predicate on the "shipment_content_description" field.
func ShipmentContentDescriptionContainsFold(v string) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShipmentContentDescription), v))
	})
}

// HasShipment applies the HasEdge predicate on the "shipment" edge.
func HasShipment() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShipmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShipmentTable, ShipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentWith applies the HasEdge predicate on the "shipment" edge with a given conditions (other predicates).
func HasShipmentWith(preds ...predicate.Shipment) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShipmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShipmentTable, ShipmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItemIssuances applies the HasEdge predicate on the "item_issuances" edge.
func HasItemIssuances() predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemIssuancesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemIssuancesTable, ItemIssuancesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemIssuancesWith applies the HasEdge predicate on the "item_issuances" edge with a given conditions (other predicates).
func HasItemIssuancesWith(preds ...predicate.ItemIssuance) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemIssuancesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemIssuancesTable, ItemIssuancesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ShipmentItem) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ShipmentItem) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
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
func Not(p predicate.ShipmentItem) predicate.ShipmentItem {
	return predicate.ShipmentItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
