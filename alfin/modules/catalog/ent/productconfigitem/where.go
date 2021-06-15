// Code generated by entc, DO NOT EDIT.

package productconfigitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// ConfigItemTypeID applies equality check predicate on the "config_item_type_id" field. It's identical to ConfigItemTypeIDEQ.
func ConfigItemTypeID(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemName applies equality check predicate on the "config_item_name" field. It's identical to ConfigItemNameEQ.
func ConfigItemName(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigItemName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// LongDescription applies equality check predicate on the "long_description" field. It's identical to LongDescriptionEQ.
func LongDescription(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongDescription), v))
	})
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// ConfigItemTypeIDEQ applies the EQ predicate on the "config_item_type_id" field.
func ConfigItemTypeIDEQ(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDNEQ applies the NEQ predicate on the "config_item_type_id" field.
func ConfigItemTypeIDNEQ(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDIn applies the In predicate on the "config_item_type_id" field.
func ConfigItemTypeIDIn(vs ...int) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfigItemTypeID), v...))
	})
}

// ConfigItemTypeIDNotIn applies the NotIn predicate on the "config_item_type_id" field.
func ConfigItemTypeIDNotIn(vs ...int) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfigItemTypeID), v...))
	})
}

// ConfigItemTypeIDGT applies the GT predicate on the "config_item_type_id" field.
func ConfigItemTypeIDGT(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDGTE applies the GTE predicate on the "config_item_type_id" field.
func ConfigItemTypeIDGTE(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDLT applies the LT predicate on the "config_item_type_id" field.
func ConfigItemTypeIDLT(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDLTE applies the LTE predicate on the "config_item_type_id" field.
func ConfigItemTypeIDLTE(v int) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfigItemTypeID), v))
	})
}

// ConfigItemTypeIDIsNil applies the IsNil predicate on the "config_item_type_id" field.
func ConfigItemTypeIDIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldConfigItemTypeID)))
	})
}

// ConfigItemTypeIDNotNil applies the NotNil predicate on the "config_item_type_id" field.
func ConfigItemTypeIDNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldConfigItemTypeID)))
	})
}

// ConfigItemNameEQ applies the EQ predicate on the "config_item_name" field.
func ConfigItemNameEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameNEQ applies the NEQ predicate on the "config_item_name" field.
func ConfigItemNameNEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameIn applies the In predicate on the "config_item_name" field.
func ConfigItemNameIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfigItemName), v...))
	})
}

// ConfigItemNameNotIn applies the NotIn predicate on the "config_item_name" field.
func ConfigItemNameNotIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfigItemName), v...))
	})
}

// ConfigItemNameGT applies the GT predicate on the "config_item_name" field.
func ConfigItemNameGT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameGTE applies the GTE predicate on the "config_item_name" field.
func ConfigItemNameGTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameLT applies the LT predicate on the "config_item_name" field.
func ConfigItemNameLT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameLTE applies the LTE predicate on the "config_item_name" field.
func ConfigItemNameLTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameContains applies the Contains predicate on the "config_item_name" field.
func ConfigItemNameContains(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameHasPrefix applies the HasPrefix predicate on the "config_item_name" field.
func ConfigItemNameHasPrefix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameHasSuffix applies the HasSuffix predicate on the "config_item_name" field.
func ConfigItemNameHasSuffix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameIsNil applies the IsNil predicate on the "config_item_name" field.
func ConfigItemNameIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldConfigItemName)))
	})
}

// ConfigItemNameNotNil applies the NotNil predicate on the "config_item_name" field.
func ConfigItemNameNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldConfigItemName)))
	})
}

// ConfigItemNameEqualFold applies the EqualFold predicate on the "config_item_name" field.
func ConfigItemNameEqualFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfigItemName), v))
	})
}

// ConfigItemNameContainsFold applies the ContainsFold predicate on the "config_item_name" field.
func ConfigItemNameContainsFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfigItemName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func DescriptionNotIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func DescriptionGT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// LongDescriptionEQ applies the EQ predicate on the "long_description" field.
func LongDescriptionEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionNEQ applies the NEQ predicate on the "long_description" field.
func LongDescriptionNEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionIn applies the In predicate on the "long_description" field.
func LongDescriptionIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLongDescription), v...))
	})
}

// LongDescriptionNotIn applies the NotIn predicate on the "long_description" field.
func LongDescriptionNotIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLongDescription), v...))
	})
}

// LongDescriptionGT applies the GT predicate on the "long_description" field.
func LongDescriptionGT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionGTE applies the GTE predicate on the "long_description" field.
func LongDescriptionGTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionLT applies the LT predicate on the "long_description" field.
func LongDescriptionLT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionLTE applies the LTE predicate on the "long_description" field.
func LongDescriptionLTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionContains applies the Contains predicate on the "long_description" field.
func LongDescriptionContains(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionHasPrefix applies the HasPrefix predicate on the "long_description" field.
func LongDescriptionHasPrefix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionHasSuffix applies the HasSuffix predicate on the "long_description" field.
func LongDescriptionHasSuffix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionIsNil applies the IsNil predicate on the "long_description" field.
func LongDescriptionIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLongDescription)))
	})
}

// LongDescriptionNotNil applies the NotNil predicate on the "long_description" field.
func LongDescriptionNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLongDescription)))
	})
}

// LongDescriptionEqualFold applies the EqualFold predicate on the "long_description" field.
func LongDescriptionEqualFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLongDescription), v))
	})
}

// LongDescriptionContainsFold applies the ContainsFold predicate on the "long_description" field.
func LongDescriptionContainsFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLongDescription), v))
	})
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImageURL), v))
	})
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImageURL), v))
	})
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImageURL), v...))
	})
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.ProductConfigItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImageURL), v...))
	})
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImageURL), v))
	})
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImageURL), v))
	})
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImageURL), v))
	})
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImageURL), v))
	})
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImageURL), v))
	})
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImageURL), v))
	})
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImageURL), v))
	})
}

// ImageURLIsNil applies the IsNil predicate on the "image_url" field.
func ImageURLIsNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldImageURL)))
	})
}

// ImageURLNotNil applies the NotNil predicate on the "image_url" field.
func ImageURLNotNil() predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldImageURL)))
	})
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImageURL), v))
	})
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImageURL), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductConfigItem) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductConfigItem) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
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
func Not(p predicate.ProductConfigItem) predicate.ProductConfigItem {
	return predicate.ProductConfigItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
