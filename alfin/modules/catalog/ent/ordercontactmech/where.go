// Code generated by entc, DO NOT EDIT.

package ordercontactmech

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func IDGT(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// ContactMechPurposeTypeID applies equality check predicate on the "contact_mech_purpose_type_id" field. It's identical to ContactMechPurposeTypeIDEQ.
func ContactMechPurposeTypeID(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechID applies equality check predicate on the "contact_mech_id" field. It's identical to ContactMechIDEQ.
func ContactMechID(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactMechID), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// ContactMechPurposeTypeIDEQ applies the EQ predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDEQ(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechPurposeTypeIDNEQ applies the NEQ predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDNEQ(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechPurposeTypeIDIn applies the In predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDIn(vs ...int) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContactMechPurposeTypeID), v...))
	})
}

// ContactMechPurposeTypeIDNotIn applies the NotIn predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDNotIn(vs ...int) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContactMechPurposeTypeID), v...))
	})
}

// ContactMechPurposeTypeIDGT applies the GT predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDGT(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechPurposeTypeIDGTE applies the GTE predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDGTE(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechPurposeTypeIDLT applies the LT predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDLT(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechPurposeTypeIDLTE applies the LTE predicate on the "contact_mech_purpose_type_id" field.
func ContactMechPurposeTypeIDLTE(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactMechPurposeTypeID), v))
	})
}

// ContactMechIDEQ applies the EQ predicate on the "contact_mech_id" field.
func ContactMechIDEQ(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContactMechID), v))
	})
}

// ContactMechIDNEQ applies the NEQ predicate on the "contact_mech_id" field.
func ContactMechIDNEQ(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContactMechID), v))
	})
}

// ContactMechIDIn applies the In predicate on the "contact_mech_id" field.
func ContactMechIDIn(vs ...int) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContactMechID), v...))
	})
}

// ContactMechIDNotIn applies the NotIn predicate on the "contact_mech_id" field.
func ContactMechIDNotIn(vs ...int) predicate.OrderContactMech {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderContactMech(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContactMechID), v...))
	})
}

// ContactMechIDGT applies the GT predicate on the "contact_mech_id" field.
func ContactMechIDGT(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContactMechID), v))
	})
}

// ContactMechIDGTE applies the GTE predicate on the "contact_mech_id" field.
func ContactMechIDGTE(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContactMechID), v))
	})
}

// ContactMechIDLT applies the LT predicate on the "contact_mech_id" field.
func ContactMechIDLT(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContactMechID), v))
	})
}

// ContactMechIDLTE applies the LTE predicate on the "contact_mech_id" field.
func ContactMechIDLTE(v int) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContactMechID), v))
	})
}

// HasOrderHeader applies the HasEdge predicate on the "order_header" edge.
func HasOrderHeader() predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrderHeaderTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderHeaderTable, OrderHeaderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderHeaderWith applies the HasEdge predicate on the "order_header" edge with a given conditions (other predicates).
func HasOrderHeaderWith(preds ...predicate.OrderHeader) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderContactMech) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderContactMech) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
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
func Not(p predicate.OrderContactMech) predicate.OrderContactMech {
	return predicate.OrderContactMech(func(s *sql.Selector) {
		p(s.Not())
	})
}