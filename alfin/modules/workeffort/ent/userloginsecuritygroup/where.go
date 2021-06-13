// Code generated by entc, DO NOT EDIT.

package userloginsecuritygroup

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func IDGT(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FromDate applies equality check predicate on the "from_date" field. It's identical to FromDateEQ.
func FromDate(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// ThruDate applies equality check predicate on the "thru_date" field. It's identical to ThruDateEQ.
func ThruDate(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// FromDateEQ applies the EQ predicate on the "from_date" field.
func FromDateEQ(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFromDate), v))
	})
}

// FromDateNEQ applies the NEQ predicate on the "from_date" field.
func FromDateNEQ(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFromDate), v))
	})
}

// FromDateIn applies the In predicate on the "from_date" field.
func FromDateIn(vs ...time.Time) predicate.UserLoginSecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func FromDateNotIn(vs ...time.Time) predicate.UserLoginSecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func FromDateGT(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFromDate), v))
	})
}

// FromDateGTE applies the GTE predicate on the "from_date" field.
func FromDateGTE(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFromDate), v))
	})
}

// FromDateLT applies the LT predicate on the "from_date" field.
func FromDateLT(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFromDate), v))
	})
}

// FromDateLTE applies the LTE predicate on the "from_date" field.
func FromDateLTE(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFromDate), v))
	})
}

// ThruDateEQ applies the EQ predicate on the "thru_date" field.
func ThruDateEQ(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThruDate), v))
	})
}

// ThruDateNEQ applies the NEQ predicate on the "thru_date" field.
func ThruDateNEQ(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThruDate), v))
	})
}

// ThruDateIn applies the In predicate on the "thru_date" field.
func ThruDateIn(vs ...time.Time) predicate.UserLoginSecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func ThruDateNotIn(vs ...time.Time) predicate.UserLoginSecurityGroup {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func ThruDateGT(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThruDate), v))
	})
}

// ThruDateGTE applies the GTE predicate on the "thru_date" field.
func ThruDateGTE(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThruDate), v))
	})
}

// ThruDateLT applies the LT predicate on the "thru_date" field.
func ThruDateLT(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThruDate), v))
	})
}

// ThruDateLTE applies the LTE predicate on the "thru_date" field.
func ThruDateLTE(v time.Time) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThruDate), v))
	})
}

// ThruDateIsNil applies the IsNil predicate on the "thru_date" field.
func ThruDateIsNil() predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThruDate)))
	})
}

// ThruDateNotNil applies the NotNil predicate on the "thru_date" field.
func ThruDateNotNil() predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThruDate)))
	})
}

// HasUserLogin applies the HasEdge predicate on the "user_login" edge.
func HasUserLogin() predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserLoginTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserLoginTable, UserLoginColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserLoginWith applies the HasEdge predicate on the "user_login" edge with a given conditions (other predicates).
func HasUserLoginWith(preds ...predicate.UserLogin) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserLoginInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserLoginTable, UserLoginColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSecurityGroup applies the HasEdge predicate on the "security_group" edge.
func HasSecurityGroup() predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SecurityGroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SecurityGroupTable, SecurityGroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecurityGroupWith applies the HasEdge predicate on the "security_group" edge with a given conditions (other predicates).
func HasSecurityGroupWith(preds ...predicate.SecurityGroup) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SecurityGroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SecurityGroupTable, SecurityGroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSecurityGroupPermissions applies the HasEdge predicate on the "security_group_permissions" edge.
func HasSecurityGroupPermissions() predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SecurityGroupPermissionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SecurityGroupPermissionsTable, SecurityGroupPermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSecurityGroupPermissionsWith applies the HasEdge predicate on the "security_group_permissions" edge with a given conditions (other predicates).
func HasSecurityGroupPermissionsWith(preds ...predicate.SecurityGroupPermission) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserLoginSecurityGroup) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserLoginSecurityGroup) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
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
func Not(p predicate.UserLoginSecurityGroup) predicate.UserLoginSecurityGroup {
	return predicate.UserLoginSecurityGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
