// Code generated by entc, DO NOT EDIT.

package enumeration

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// EnumCode applies equality check predicate on the "enum_code" field. It's identical to EnumCodeEQ.
func EnumCode(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnumCode), v))
	})
}

// SequenceID applies equality check predicate on the "sequence_id" field. It's identical to SequenceIDEQ.
func SequenceID(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceID), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// EnumCodeEQ applies the EQ predicate on the "enum_code" field.
func EnumCodeEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnumCode), v))
	})
}

// EnumCodeNEQ applies the NEQ predicate on the "enum_code" field.
func EnumCodeNEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnumCode), v))
	})
}

// EnumCodeIn applies the In predicate on the "enum_code" field.
func EnumCodeIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnumCode), v...))
	})
}

// EnumCodeNotIn applies the NotIn predicate on the "enum_code" field.
func EnumCodeNotIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnumCode), v...))
	})
}

// EnumCodeGT applies the GT predicate on the "enum_code" field.
func EnumCodeGT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnumCode), v))
	})
}

// EnumCodeGTE applies the GTE predicate on the "enum_code" field.
func EnumCodeGTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnumCode), v))
	})
}

// EnumCodeLT applies the LT predicate on the "enum_code" field.
func EnumCodeLT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnumCode), v))
	})
}

// EnumCodeLTE applies the LTE predicate on the "enum_code" field.
func EnumCodeLTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnumCode), v))
	})
}

// EnumCodeContains applies the Contains predicate on the "enum_code" field.
func EnumCodeContains(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnumCode), v))
	})
}

// EnumCodeHasPrefix applies the HasPrefix predicate on the "enum_code" field.
func EnumCodeHasPrefix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnumCode), v))
	})
}

// EnumCodeHasSuffix applies the HasSuffix predicate on the "enum_code" field.
func EnumCodeHasSuffix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnumCode), v))
	})
}

// EnumCodeIsNil applies the IsNil predicate on the "enum_code" field.
func EnumCodeIsNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEnumCode)))
	})
}

// EnumCodeNotNil applies the NotNil predicate on the "enum_code" field.
func EnumCodeNotNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEnumCode)))
	})
}

// EnumCodeEqualFold applies the EqualFold predicate on the "enum_code" field.
func EnumCodeEqualFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnumCode), v))
	})
}

// EnumCodeContainsFold applies the ContainsFold predicate on the "enum_code" field.
func EnumCodeContainsFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnumCode), v))
	})
}

// SequenceIDEQ applies the EQ predicate on the "sequence_id" field.
func SequenceIDEQ(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSequenceID), v))
	})
}

// SequenceIDNEQ applies the NEQ predicate on the "sequence_id" field.
func SequenceIDNEQ(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSequenceID), v))
	})
}

// SequenceIDIn applies the In predicate on the "sequence_id" field.
func SequenceIDIn(vs ...int) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSequenceID), v...))
	})
}

// SequenceIDNotIn applies the NotIn predicate on the "sequence_id" field.
func SequenceIDNotIn(vs ...int) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSequenceID), v...))
	})
}

// SequenceIDGT applies the GT predicate on the "sequence_id" field.
func SequenceIDGT(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSequenceID), v))
	})
}

// SequenceIDGTE applies the GTE predicate on the "sequence_id" field.
func SequenceIDGTE(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSequenceID), v))
	})
}

// SequenceIDLT applies the LT predicate on the "sequence_id" field.
func SequenceIDLT(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSequenceID), v))
	})
}

// SequenceIDLTE applies the LTE predicate on the "sequence_id" field.
func SequenceIDLTE(v int) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSequenceID), v))
	})
}

// SequenceIDIsNil applies the IsNil predicate on the "sequence_id" field.
func SequenceIDIsNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSequenceID)))
	})
}

// SequenceIDNotNil applies the NotNil predicate on the "sequence_id" field.
func SequenceIDNotNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSequenceID)))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Enumeration {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Enumeration(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasEnumerationType applies the HasEdge predicate on the "enumeration_type" edge.
func HasEnumerationType() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EnumerationTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EnumerationTypeTable, EnumerationTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEnumerationTypeWith applies the HasEdge predicate on the "enumeration_type" edge with a given conditions (other predicates).
func HasEnumerationTypeWith(preds ...predicate.EnumerationType) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EnumerationTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EnumerationTypeTable, EnumerationTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClassFixedAssets applies the HasEdge predicate on the "class_fixed_assets" edge.
func HasClassFixedAssets() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassFixedAssetsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassFixedAssetsTable, ClassFixedAssetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassFixedAssetsWith applies the HasEdge predicate on the "class_fixed_assets" edge with a given conditions (other predicates).
func HasClassFixedAssetsWith(preds ...predicate.FixedAsset) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassFixedAssetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassFixedAssetsTable, ClassFixedAssetsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEmploymentStatusPeople applies the HasEdge predicate on the "employment_status_people" edge.
func HasEmploymentStatusPeople() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmploymentStatusPeopleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmploymentStatusPeopleTable, EmploymentStatusPeopleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmploymentStatusPeopleWith applies the HasEdge predicate on the "employment_status_people" edge with a given conditions (other predicates).
func HasEmploymentStatusPeopleWith(preds ...predicate.Person) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmploymentStatusPeopleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmploymentStatusPeopleTable, EmploymentStatusPeopleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResidenceStatusPeople applies the HasEdge predicate on the "residence_status_people" edge.
func HasResidenceStatusPeople() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResidenceStatusPeopleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResidenceStatusPeopleTable, ResidenceStatusPeopleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResidenceStatusPeopleWith applies the HasEdge predicate on the "residence_status_people" edge with a given conditions (other predicates).
func HasResidenceStatusPeopleWith(preds ...predicate.Person) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ResidenceStatusPeopleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ResidenceStatusPeopleTable, ResidenceStatusPeopleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMaritalStatusPeople applies the HasEdge predicate on the "marital_status_people" edge.
func HasMaritalStatusPeople() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaritalStatusPeopleTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MaritalStatusPeopleTable, MaritalStatusPeopleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMaritalStatusPeopleWith applies the HasEdge predicate on the "marital_status_people" edge with a given conditions (other predicates).
func HasMaritalStatusPeopleWith(preds ...predicate.Person) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaritalStatusPeopleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MaritalStatusPeopleTable, MaritalStatusPeopleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasScopeWorkEfforts applies the HasEdge predicate on the "scope_work_efforts" edge.
func HasScopeWorkEfforts() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScopeWorkEffortsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ScopeWorkEffortsTable, ScopeWorkEffortsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScopeWorkEffortsWith applies the HasEdge predicate on the "scope_work_efforts" edge with a given conditions (other predicates).
func HasScopeWorkEffortsWith(preds ...predicate.WorkEffort) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScopeWorkEffortsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ScopeWorkEffortsTable, ScopeWorkEffortsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExpectationWorkEffortPartyAssignments applies the HasEdge predicate on the "expectation_work_effort_party_assignments" edge.
func HasExpectationWorkEffortPartyAssignments() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExpectationWorkEffortPartyAssignmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExpectationWorkEffortPartyAssignmentsTable, ExpectationWorkEffortPartyAssignmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExpectationWorkEffortPartyAssignmentsWith applies the HasEdge predicate on the "expectation_work_effort_party_assignments" edge with a given conditions (other predicates).
func HasExpectationWorkEffortPartyAssignmentsWith(preds ...predicate.WorkEffortPartyAssignment) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ExpectationWorkEffortPartyAssignmentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExpectationWorkEffortPartyAssignmentsTable, ExpectationWorkEffortPartyAssignmentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDelegateReasonWorkEffortPartyAssignments applies the HasEdge predicate on the "delegate_reason_work_effort_party_assignments" edge.
func HasDelegateReasonWorkEffortPartyAssignments() predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DelegateReasonWorkEffortPartyAssignmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DelegateReasonWorkEffortPartyAssignmentsTable, DelegateReasonWorkEffortPartyAssignmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDelegateReasonWorkEffortPartyAssignmentsWith applies the HasEdge predicate on the "delegate_reason_work_effort_party_assignments" edge with a given conditions (other predicates).
func HasDelegateReasonWorkEffortPartyAssignmentsWith(preds ...predicate.WorkEffortPartyAssignment) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DelegateReasonWorkEffortPartyAssignmentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DelegateReasonWorkEffortPartyAssignmentsTable, DelegateReasonWorkEffortPartyAssignmentsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Enumeration) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Enumeration) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
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
func Not(p predicate.Enumeration) predicate.Enumeration {
	return predicate.Enumeration(func(s *sql.Selector) {
		p(s.Not())
	})
}