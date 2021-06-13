// Code generated by entc, DO NOT EDIT.

package securitygroup

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// StringRef applies equality check predicate on the "string_ref" field. It's identical to StringRefEQ.
func StringRef(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// GroupName applies equality check predicate on the "group_name" field. It's identical to GroupNameEQ.
func GroupName(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// StringRefEQ applies the EQ predicate on the "string_ref" field.
func StringRefEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStringRef), v))
	})
}

// StringRefNEQ applies the NEQ predicate on the "string_ref" field.
func StringRefNEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStringRef), v))
	})
}

// StringRefIn applies the In predicate on the "string_ref" field.
func StringRefIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func StringRefNotIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func StringRefGT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStringRef), v))
	})
}

// StringRefGTE applies the GTE predicate on the "string_ref" field.
func StringRefGTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStringRef), v))
	})
}

// StringRefLT applies the LT predicate on the "string_ref" field.
func StringRefLT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStringRef), v))
	})
}

// StringRefLTE applies the LTE predicate on the "string_ref" field.
func StringRefLTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStringRef), v))
	})
}

// StringRefContains applies the Contains predicate on the "string_ref" field.
func StringRefContains(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStringRef), v))
	})
}

// StringRefHasPrefix applies the HasPrefix predicate on the "string_ref" field.
func StringRefHasPrefix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStringRef), v))
	})
}

// StringRefHasSuffix applies the HasSuffix predicate on the "string_ref" field.
func StringRefHasSuffix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStringRef), v))
	})
}

// StringRefIsNil applies the IsNil predicate on the "string_ref" field.
func StringRefIsNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStringRef)))
	})
}

// StringRefNotNil applies the NotNil predicate on the "string_ref" field.
func StringRefNotNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStringRef)))
	})
}

// StringRefEqualFold applies the EqualFold predicate on the "string_ref" field.
func StringRefEqualFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStringRef), v))
	})
}

// StringRefContainsFold applies the ContainsFold predicate on the "string_ref" field.
func StringRefContainsFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStringRef), v))
	})
}

// GroupNameEQ applies the EQ predicate on the "group_name" field.
func GroupNameEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupName), v))
	})
}

// GroupNameNEQ applies the NEQ predicate on the "group_name" field.
func GroupNameNEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupName), v))
	})
}

// GroupNameIn applies the In predicate on the "group_name" field.
func GroupNameIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupName), v...))
	})
}

// GroupNameNotIn applies the NotIn predicate on the "group_name" field.
func GroupNameNotIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupName), v...))
	})
}

// GroupNameGT applies the GT predicate on the "group_name" field.
func GroupNameGT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupName), v))
	})
}

// GroupNameGTE applies the GTE predicate on the "group_name" field.
func GroupNameGTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupName), v))
	})
}

// GroupNameLT applies the LT predicate on the "group_name" field.
func GroupNameLT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupName), v))
	})
}

// GroupNameLTE applies the LTE predicate on the "group_name" field.
func GroupNameLTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupName), v))
	})
}

// GroupNameContains applies the Contains predicate on the "group_name" field.
func GroupNameContains(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupName), v))
	})
}

// GroupNameHasPrefix applies the HasPrefix predicate on the "group_name" field.
func GroupNameHasPrefix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupName), v))
	})
}

// GroupNameHasSuffix applies the HasSuffix predicate on the "group_name" field.
func GroupNameHasSuffix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupName), v))
	})
}

// GroupNameIsNil applies the IsNil predicate on the "group_name" field.
func GroupNameIsNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGroupName)))
	})
}

// GroupNameNotNil applies the NotNil predicate on the "group_name" field.
func GroupNameNotNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGroupName)))
	})
}

// GroupNameEqualFold applies the EqualFold predicate on the "group_name" field.
func GroupNameEqualFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupName), v))
	})
}

// GroupNameContainsFold applies the ContainsFold predicate on the "group_name" field.
func GroupNameContainsFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func DescriptionNotIn(vs ...string) predicate.SecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func DescriptionGT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// HasSecurityGroupPermissions applies the HasEdge predicate on the "security_group_permissions" edge.
func HasSecurityGroupPermissions() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SecurityGroupPermissionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SecurityGroupPermissionsTable, SecurityGroupPermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecurityGroupPermissionsWith applies the HasEdge predicate on the "security_group_permissions" edge with a given conditions (other predicates).
func HasSecurityGroupPermissionsWith(preds ...predicate.SecurityGroupPermission) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SecurityGroupPermissionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SecurityGroupPermissionsTable, SecurityGroupPermissionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserLoginSecurityGroups applies the HasEdge predicate on the "user_login_security_groups" edge.
func HasUserLoginSecurityGroups() predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserLoginSecurityGroupsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserLoginSecurityGroupsTable, UserLoginSecurityGroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserLoginSecurityGroupsWith applies the HasEdge predicate on the "user_login_security_groups" edge with a given conditions (other predicates).
func HasUserLoginSecurityGroupsWith(preds ...predicate.UserLoginSecurityGroup) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserLoginSecurityGroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserLoginSecurityGroupsTable, UserLoginSecurityGroupsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SecurityGroup) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SecurityGroup) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
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
func Not(p predicate.SecurityGroup) predicate.SecurityGroup {
	return predicate.SecurityGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
