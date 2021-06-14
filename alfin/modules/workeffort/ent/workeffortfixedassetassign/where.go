// Code generated by entc, DO NOT EDIT.

package workeffortfixedassetassign

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func IDGT(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// ThruDate applies equality check predicate on the "thru_date" field. It's identical to ThruDateEQ.
func ThruDate(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// AllocatedCost applies equality check predicate on the "allocated_cost" field. It's identical to AllocatedCostEQ.
func AllocatedCost(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocatedCost), v))
	})
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComments), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromDate), v))
	})
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func FromDateNotIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func FromDateGT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromDate), v))
	})
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromDate), v))
	})
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromDate), v))
	})
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromDate), v))
	})
}

// ThruDateEQ applies the EQ predicate on the "thru_date" field.
func ThruDateEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// ThruDateNEQ applies the NEQ predicate on the "thru_date" field.
func ThruDateNEQ(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThruDate), v))
	})
}

// ThruDateIn applies the In predicate on the "thru_date" field.
func ThruDateIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func ThruDateNotIn(vs ...time.Time) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func ThruDateGT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThruDate), v))
	})
}

// ThruDateGTE applies the GTE predicate on the "thru_date" field.
func ThruDateGTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThruDate), v))
	})
}

// ThruDateLT applies the LT predicate on the "thru_date" field.
func ThruDateLT(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThruDate), v))
	})
}

// ThruDateLTE applies the LTE predicate on the "thru_date" field.
func ThruDateLTE(v time.Time) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThruDate), v))
	})
}

// ThruDateIsNil applies the IsNil predicate on the "thru_date" field.
func ThruDateIsNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThruDate)))
	})
}

// ThruDateNotNil applies the NotNil predicate on the "thru_date" field.
func ThruDateNotNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThruDate)))
	})
}

// AllocatedCostEQ applies the EQ predicate on the "allocated_cost" field.
func AllocatedCostEQ(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostNEQ applies the NEQ predicate on the "allocated_cost" field.
func AllocatedCostNEQ(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostIn applies the In predicate on the "allocated_cost" field.
func AllocatedCostIn(vs ...float64) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAllocatedCost), v...))
	})
}

// AllocatedCostNotIn applies the NotIn predicate on the "allocated_cost" field.
func AllocatedCostNotIn(vs ...float64) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAllocatedCost), v...))
	})
}

// AllocatedCostGT applies the GT predicate on the "allocated_cost" field.
func AllocatedCostGT(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostGTE applies the GTE predicate on the "allocated_cost" field.
func AllocatedCostGTE(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostLT applies the LT predicate on the "allocated_cost" field.
func AllocatedCostLT(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostLTE applies the LTE predicate on the "allocated_cost" field.
func AllocatedCostLTE(v float64) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAllocatedCost), v))
	})
}

// AllocatedCostIsNil applies the IsNil predicate on the "allocated_cost" field.
func AllocatedCostIsNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAllocatedCost)))
	})
}

// AllocatedCostNotNil applies the NotNil predicate on the "allocated_cost" field.
func AllocatedCostNotNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAllocatedCost)))
	})
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComments), v))
	})
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComments), v))
	})
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldComments), v...))
	})
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.WorkEffortFixedAssetAssign {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldComments), v...))
	})
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComments), v))
	})
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComments), v))
	})
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComments), v))
	})
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComments), v))
	})
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComments), v))
	})
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComments), v))
	})
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComments), v))
	})
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldComments)))
	})
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldComments)))
	})
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComments), v))
	})
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComments), v))
	})
}

// HasWorkEffort applies the HasEdge predicate on the "work_effort" edge.
func HasWorkEffort() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkEffortTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkEffortTable, WorkEffortColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkEffortWith applies the HasEdge predicate on the "work_effort" edge with a given conditions (other predicates).
func HasWorkEffortWith(preds ...predicate.WorkEffort) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WorkEffortInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkEffortTable, WorkEffortColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFixedAsset applies the HasEdge predicate on the "fixed_asset" edge.
func HasFixedAsset() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixedAssetTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixedAssetTable, FixedAssetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFixedAssetWith applies the HasEdge predicate on the "fixed_asset" edge with a given conditions (other predicates).
func HasFixedAssetWith(preds ...predicate.FixedAsset) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FixedAssetInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FixedAssetTable, FixedAssetColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusItem applies the HasEdge predicate on the "status_item" edge.
func HasStatusItem() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusItemTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusItemTable, StatusItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusItemWith applies the HasEdge predicate on the "status_item" edge with a given conditions (other predicates).
func HasStatusItemWith(preds ...predicate.StatusItem) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, StatusItemTable, StatusItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAvailabilityStatusItem applies the HasEdge predicate on the "availability_status_item" edge.
func HasAvailabilityStatusItem() predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AvailabilityStatusItemTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AvailabilityStatusItemTable, AvailabilityStatusItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAvailabilityStatusItemWith applies the HasEdge predicate on the "availability_status_item" edge with a given conditions (other predicates).
func HasAvailabilityStatusItemWith(preds ...predicate.StatusItem) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AvailabilityStatusItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AvailabilityStatusItemTable, AvailabilityStatusItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkEffortFixedAssetAssign) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkEffortFixedAssetAssign) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
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
func Not(p predicate.WorkEffortFixedAssetAssign) predicate.WorkEffortFixedAssetAssign {
	return predicate.WorkEffortFixedAssetAssign(func(s *sql.Selector) {
		p(s.Not())
	})
}