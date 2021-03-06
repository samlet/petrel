// Code generated by entc, DO NOT EDIT.

package workeffortassoc

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func IDGT(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// WorkEffortAssocTypeID applies equality check predicate on the "work_effort_assoc_type_id" field. It's identical to WorkEffortAssocTypeIDEQ.
func WorkEffortAssocTypeID(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// SequenceNum applies equality check predicate on the "sequence_num" field. It's identical to SequenceNumEQ.
func SequenceNum(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceNum), v))
	})
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// ThruDate applies equality check predicate on the "thru_date" field. It's identical to ThruDateEQ.
func ThruDate(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// WorkEffortAssocTypeIDEQ applies the EQ predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDEQ(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// WorkEffortAssocTypeIDNEQ applies the NEQ predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDNEQ(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// WorkEffortAssocTypeIDIn applies the In predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDIn(vs ...int) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWorkEffortAssocTypeID), v...))
	})
}

// WorkEffortAssocTypeIDNotIn applies the NotIn predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDNotIn(vs ...int) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWorkEffortAssocTypeID), v...))
	})
}

// WorkEffortAssocTypeIDGT applies the GT predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDGT(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// WorkEffortAssocTypeIDGTE applies the GTE predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDGTE(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// WorkEffortAssocTypeIDLT applies the LT predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDLT(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// WorkEffortAssocTypeIDLTE applies the LTE predicate on the "work_effort_assoc_type_id" field.
func WorkEffortAssocTypeIDLTE(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWorkEffortAssocTypeID), v))
	})
}

// SequenceNumEQ applies the EQ predicate on the "sequence_num" field.
func SequenceNumEQ(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumNEQ applies the NEQ predicate on the "sequence_num" field.
func SequenceNumNEQ(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumIn applies the In predicate on the "sequence_num" field.
func SequenceNumIn(vs ...int) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSequenceNum), v...))
	})
}

// SequenceNumNotIn applies the NotIn predicate on the "sequence_num" field.
func SequenceNumNotIn(vs ...int) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSequenceNum), v...))
	})
}

// SequenceNumGT applies the GT predicate on the "sequence_num" field.
func SequenceNumGT(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumGTE applies the GTE predicate on the "sequence_num" field.
func SequenceNumGTE(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumLT applies the LT predicate on the "sequence_num" field.
func SequenceNumLT(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumLTE applies the LTE predicate on the "sequence_num" field.
func SequenceNumLTE(v int) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSequenceNum), v))
	})
}

// SequenceNumIsNil applies the IsNil predicate on the "sequence_num" field.
func SequenceNumIsNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSequenceNum)))
	})
}

// SequenceNumNotNil applies the NotNil predicate on the "sequence_num" field.
func SequenceNumNotNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSequenceNum)))
	})
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromDate), v))
	})
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFromDate), v...))
	})
}

// FromDateNotIn applies the NotIn predicate on the "from_date" field.
func FromDateNotIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFromDate), v...))
	})
}

// FromDateGT applies the GT predicate on the "from_date" field.
func FromDateGT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromDate), v))
	})
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromDate), v))
	})
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromDate), v))
	})
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromDate), v))
	})
}

// ThruDateEQ applies the EQ predicate on the "thru_date" field.
func ThruDateEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// ThruDateNEQ applies the NEQ predicate on the "thru_date" field.
func ThruDateNEQ(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThruDate), v))
	})
}

// ThruDateIn applies the In predicate on the "thru_date" field.
func ThruDateIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThruDate), v...))
	})
}

// ThruDateNotIn applies the NotIn predicate on the "thru_date" field.
func ThruDateNotIn(vs ...time.Time) predicate.WorkEffortAssoc {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThruDate), v...))
	})
}

// ThruDateGT applies the GT predicate on the "thru_date" field.
func ThruDateGT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThruDate), v))
	})
}

// ThruDateGTE applies the GTE predicate on the "thru_date" field.
func ThruDateGTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThruDate), v))
	})
}

// ThruDateLT applies the LT predicate on the "thru_date" field.
func ThruDateLT(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThruDate), v))
	})
}

// ThruDateLTE applies the LTE predicate on the "thru_date" field.
func ThruDateLTE(v time.Time) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThruDate), v))
	})
}

// ThruDateIsNil applies the IsNil predicate on the "thru_date" field.
func ThruDateIsNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThruDate)))
	})
}

// ThruDateNotNil applies the NotNil predicate on the "thru_date" field.
func ThruDateNotNil() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThruDate)))
	})
}

// HasFromWorkEffort applies the HasEdge predicate on the "from_work_effort" edge.
func HasFromWorkEffort() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromWorkEffortTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromWorkEffortTable, FromWorkEffortColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFromWorkEffortWith applies the HasEdge predicate on the "from_work_effort" edge with a given conditions (other predicates).
func HasFromWorkEffortWith(preds ...predicate.WorkEffort) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FromWorkEffortInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FromWorkEffortTable, FromWorkEffortColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasToWorkEffort applies the HasEdge predicate on the "to_work_effort" edge.
func HasToWorkEffort() predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ToWorkEffortTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToWorkEffortTable, ToWorkEffortColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToWorkEffortWith applies the HasEdge predicate on the "to_work_effort" edge with a given conditions (other predicates).
func HasToWorkEffortWith(preds ...predicate.WorkEffort) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ToWorkEffortInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ToWorkEffortTable, ToWorkEffortColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkEffortAssoc) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkEffortAssoc) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
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
func Not(p predicate.WorkEffortAssoc) predicate.WorkEffortAssoc {
	return predicate.WorkEffortAssoc(func(s *sql.Selector) {
		p(s.Not())
	})
}
