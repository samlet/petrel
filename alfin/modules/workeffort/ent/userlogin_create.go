// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partystatus"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userloginsecuritygroup"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userpreference"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortpartyassignment"
)

// UserLoginCreate is the builder for creating a UserLogin entity.
type UserLoginCreate struct {
	config
	mutation *UserLoginMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ulc *UserLoginCreate) SetCreateTime(t time.Time) *UserLoginCreate {
	ulc.mutation.SetCreateTime(t)
	return ulc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableCreateTime(t *time.Time) *UserLoginCreate {
	if t != nil {
		ulc.SetCreateTime(*t)
	}
	return ulc
}

// SetUpdateTime sets the "update_time" field.
func (ulc *UserLoginCreate) SetUpdateTime(t time.Time) *UserLoginCreate {
	ulc.mutation.SetUpdateTime(t)
	return ulc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableUpdateTime(t *time.Time) *UserLoginCreate {
	if t != nil {
		ulc.SetUpdateTime(*t)
	}
	return ulc
}

// SetStringRef sets the "string_ref" field.
func (ulc *UserLoginCreate) SetStringRef(s string) *UserLoginCreate {
	ulc.mutation.SetStringRef(s)
	return ulc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableStringRef(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetStringRef(*s)
	}
	return ulc
}

// SetCurrentPassword sets the "current_password" field.
func (ulc *UserLoginCreate) SetCurrentPassword(s string) *UserLoginCreate {
	ulc.mutation.SetCurrentPassword(s)
	return ulc
}

// SetNillableCurrentPassword sets the "current_password" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableCurrentPassword(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetCurrentPassword(*s)
	}
	return ulc
}

// SetPasswordHint sets the "password_hint" field.
func (ulc *UserLoginCreate) SetPasswordHint(s string) *UserLoginCreate {
	ulc.mutation.SetPasswordHint(s)
	return ulc
}

// SetNillablePasswordHint sets the "password_hint" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillablePasswordHint(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetPasswordHint(*s)
	}
	return ulc
}

// SetIsSystem sets the "is_system" field.
func (ulc *UserLoginCreate) SetIsSystem(us userlogin.IsSystem) *UserLoginCreate {
	ulc.mutation.SetIsSystem(us)
	return ulc
}

// SetNillableIsSystem sets the "is_system" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableIsSystem(us *userlogin.IsSystem) *UserLoginCreate {
	if us != nil {
		ulc.SetIsSystem(*us)
	}
	return ulc
}

// SetEnabled sets the "enabled" field.
func (ulc *UserLoginCreate) SetEnabled(u userlogin.Enabled) *UserLoginCreate {
	ulc.mutation.SetEnabled(u)
	return ulc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableEnabled(u *userlogin.Enabled) *UserLoginCreate {
	if u != nil {
		ulc.SetEnabled(*u)
	}
	return ulc
}

// SetHasLoggedOut sets the "has_logged_out" field.
func (ulc *UserLoginCreate) SetHasLoggedOut(ulo userlogin.HasLoggedOut) *UserLoginCreate {
	ulc.mutation.SetHasLoggedOut(ulo)
	return ulc
}

// SetNillableHasLoggedOut sets the "has_logged_out" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableHasLoggedOut(ulo *userlogin.HasLoggedOut) *UserLoginCreate {
	if ulo != nil {
		ulc.SetHasLoggedOut(*ulo)
	}
	return ulc
}

// SetRequirePasswordChange sets the "require_password_change" field.
func (ulc *UserLoginCreate) SetRequirePasswordChange(upc userlogin.RequirePasswordChange) *UserLoginCreate {
	ulc.mutation.SetRequirePasswordChange(upc)
	return ulc
}

// SetNillableRequirePasswordChange sets the "require_password_change" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableRequirePasswordChange(upc *userlogin.RequirePasswordChange) *UserLoginCreate {
	if upc != nil {
		ulc.SetRequirePasswordChange(*upc)
	}
	return ulc
}

// SetLastCurrencyUom sets the "last_currency_uom" field.
func (ulc *UserLoginCreate) SetLastCurrencyUom(i int) *UserLoginCreate {
	ulc.mutation.SetLastCurrencyUom(i)
	return ulc
}

// SetNillableLastCurrencyUom sets the "last_currency_uom" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableLastCurrencyUom(i *int) *UserLoginCreate {
	if i != nil {
		ulc.SetLastCurrencyUom(*i)
	}
	return ulc
}

// SetLastLocale sets the "last_locale" field.
func (ulc *UserLoginCreate) SetLastLocale(s string) *UserLoginCreate {
	ulc.mutation.SetLastLocale(s)
	return ulc
}

// SetNillableLastLocale sets the "last_locale" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableLastLocale(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetLastLocale(*s)
	}
	return ulc
}

// SetLastTimeZone sets the "last_time_zone" field.
func (ulc *UserLoginCreate) SetLastTimeZone(s string) *UserLoginCreate {
	ulc.mutation.SetLastTimeZone(s)
	return ulc
}

// SetNillableLastTimeZone sets the "last_time_zone" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableLastTimeZone(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetLastTimeZone(*s)
	}
	return ulc
}

// SetDisabledDateTime sets the "disabled_date_time" field.
func (ulc *UserLoginCreate) SetDisabledDateTime(t time.Time) *UserLoginCreate {
	ulc.mutation.SetDisabledDateTime(t)
	return ulc
}

// SetNillableDisabledDateTime sets the "disabled_date_time" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableDisabledDateTime(t *time.Time) *UserLoginCreate {
	if t != nil {
		ulc.SetDisabledDateTime(*t)
	}
	return ulc
}

// SetSuccessiveFailedLogins sets the "successive_failed_logins" field.
func (ulc *UserLoginCreate) SetSuccessiveFailedLogins(i int) *UserLoginCreate {
	ulc.mutation.SetSuccessiveFailedLogins(i)
	return ulc
}

// SetNillableSuccessiveFailedLogins sets the "successive_failed_logins" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableSuccessiveFailedLogins(i *int) *UserLoginCreate {
	if i != nil {
		ulc.SetSuccessiveFailedLogins(*i)
	}
	return ulc
}

// SetExternalAuthID sets the "external_auth_id" field.
func (ulc *UserLoginCreate) SetExternalAuthID(s string) *UserLoginCreate {
	ulc.mutation.SetExternalAuthID(s)
	return ulc
}

// SetNillableExternalAuthID sets the "external_auth_id" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableExternalAuthID(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetExternalAuthID(*s)
	}
	return ulc
}

// SetUserLdapDn sets the "user_ldap_dn" field.
func (ulc *UserLoginCreate) SetUserLdapDn(s string) *UserLoginCreate {
	ulc.mutation.SetUserLdapDn(s)
	return ulc
}

// SetNillableUserLdapDn sets the "user_ldap_dn" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableUserLdapDn(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetUserLdapDn(*s)
	}
	return ulc
}

// SetDisabledBy sets the "disabled_by" field.
func (ulc *UserLoginCreate) SetDisabledBy(s string) *UserLoginCreate {
	ulc.mutation.SetDisabledBy(s)
	return ulc
}

// SetNillableDisabledBy sets the "disabled_by" field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableDisabledBy(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetDisabledBy(*s)
	}
	return ulc
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (ulc *UserLoginCreate) SetPartyID(id int) *UserLoginCreate {
	ulc.mutation.SetPartyID(id)
	return ulc
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (ulc *UserLoginCreate) SetNillablePartyID(id *int) *UserLoginCreate {
	if id != nil {
		ulc = ulc.SetPartyID(*id)
	}
	return ulc
}

// SetParty sets the "party" edge to the Party entity.
func (ulc *UserLoginCreate) SetParty(p *Party) *UserLoginCreate {
	return ulc.SetPartyID(p.ID)
}

// SetPersonID sets the "person" edge to the Person entity by ID.
func (ulc *UserLoginCreate) SetPersonID(id int) *UserLoginCreate {
	ulc.mutation.SetPersonID(id)
	return ulc
}

// SetNillablePersonID sets the "person" edge to the Person entity by ID if the given value is not nil.
func (ulc *UserLoginCreate) SetNillablePersonID(id *int) *UserLoginCreate {
	if id != nil {
		ulc = ulc.SetPersonID(*id)
	}
	return ulc
}

// SetPerson sets the "person" edge to the Person entity.
func (ulc *UserLoginCreate) SetPerson(p *Person) *UserLoginCreate {
	return ulc.SetPersonID(p.ID)
}

// AddCreatedByPartyIDs adds the "created_by_parties" edge to the Party entity by IDs.
func (ulc *UserLoginCreate) AddCreatedByPartyIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddCreatedByPartyIDs(ids...)
	return ulc
}

// AddCreatedByParties adds the "created_by_parties" edges to the Party entity.
func (ulc *UserLoginCreate) AddCreatedByParties(p ...*Party) *UserLoginCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ulc.AddCreatedByPartyIDs(ids...)
}

// AddLastModifiedByPartyIDs adds the "last_modified_by_parties" edge to the Party entity by IDs.
func (ulc *UserLoginCreate) AddLastModifiedByPartyIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddLastModifiedByPartyIDs(ids...)
	return ulc
}

// AddLastModifiedByParties adds the "last_modified_by_parties" edges to the Party entity.
func (ulc *UserLoginCreate) AddLastModifiedByParties(p ...*Party) *UserLoginCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ulc.AddLastModifiedByPartyIDs(ids...)
}

// AddChangeByPartyStatusIDs adds the "change_by_party_statuses" edge to the PartyStatus entity by IDs.
func (ulc *UserLoginCreate) AddChangeByPartyStatusIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddChangeByPartyStatusIDs(ids...)
	return ulc
}

// AddChangeByPartyStatuses adds the "change_by_party_statuses" edges to the PartyStatus entity.
func (ulc *UserLoginCreate) AddChangeByPartyStatuses(p ...*PartyStatus) *UserLoginCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ulc.AddChangeByPartyStatusIDs(ids...)
}

// AddUserLoginSecurityGroupIDs adds the "user_login_security_groups" edge to the UserLoginSecurityGroup entity by IDs.
func (ulc *UserLoginCreate) AddUserLoginSecurityGroupIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddUserLoginSecurityGroupIDs(ids...)
	return ulc
}

// AddUserLoginSecurityGroups adds the "user_login_security_groups" edges to the UserLoginSecurityGroup entity.
func (ulc *UserLoginCreate) AddUserLoginSecurityGroups(u ...*UserLoginSecurityGroup) *UserLoginCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ulc.AddUserLoginSecurityGroupIDs(ids...)
}

// AddUserPreferenceIDs adds the "user_preferences" edge to the UserPreference entity by IDs.
func (ulc *UserLoginCreate) AddUserPreferenceIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddUserPreferenceIDs(ids...)
	return ulc
}

// AddUserPreferences adds the "user_preferences" edges to the UserPreference entity.
func (ulc *UserLoginCreate) AddUserPreferences(u ...*UserPreference) *UserLoginCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ulc.AddUserPreferenceIDs(ids...)
}

// AddAssignedByWorkEffortPartyAssignmentIDs adds the "assigned_by_work_effort_party_assignments" edge to the WorkEffortPartyAssignment entity by IDs.
func (ulc *UserLoginCreate) AddAssignedByWorkEffortPartyAssignmentIDs(ids ...int) *UserLoginCreate {
	ulc.mutation.AddAssignedByWorkEffortPartyAssignmentIDs(ids...)
	return ulc
}

// AddAssignedByWorkEffortPartyAssignments adds the "assigned_by_work_effort_party_assignments" edges to the WorkEffortPartyAssignment entity.
func (ulc *UserLoginCreate) AddAssignedByWorkEffortPartyAssignments(w ...*WorkEffortPartyAssignment) *UserLoginCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ulc.AddAssignedByWorkEffortPartyAssignmentIDs(ids...)
}

// Mutation returns the UserLoginMutation object of the builder.
func (ulc *UserLoginCreate) Mutation() *UserLoginMutation {
	return ulc.mutation
}

// Save creates the UserLogin in the database.
func (ulc *UserLoginCreate) Save(ctx context.Context) (*UserLogin, error) {
	var (
		err  error
		node *UserLogin
	)
	ulc.defaults()
	if len(ulc.hooks) == 0 {
		if err = ulc.check(); err != nil {
			return nil, err
		}
		node, err = ulc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ulc.check(); err != nil {
				return nil, err
			}
			ulc.mutation = mutation
			if node, err = ulc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ulc.hooks) - 1; i >= 0; i-- {
			mut = ulc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ulc *UserLoginCreate) SaveX(ctx context.Context) *UserLogin {
	v, err := ulc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ulc *UserLoginCreate) defaults() {
	if _, ok := ulc.mutation.CreateTime(); !ok {
		v := userlogin.DefaultCreateTime()
		ulc.mutation.SetCreateTime(v)
	}
	if _, ok := ulc.mutation.UpdateTime(); !ok {
		v := userlogin.DefaultUpdateTime()
		ulc.mutation.SetUpdateTime(v)
	}
	if _, ok := ulc.mutation.DisabledDateTime(); !ok {
		v := userlogin.DefaultDisabledDateTime()
		ulc.mutation.SetDisabledDateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulc *UserLoginCreate) check() error {
	if _, ok := ulc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := ulc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := ulc.mutation.IsSystem(); ok {
		if err := userlogin.IsSystemValidator(v); err != nil {
			return &ValidationError{Name: "is_system", err: fmt.Errorf("ent: validator failed for field \"is_system\": %w", err)}
		}
	}
	if v, ok := ulc.mutation.Enabled(); ok {
		if err := userlogin.EnabledValidator(v); err != nil {
			return &ValidationError{Name: "enabled", err: fmt.Errorf("ent: validator failed for field \"enabled\": %w", err)}
		}
	}
	if v, ok := ulc.mutation.HasLoggedOut(); ok {
		if err := userlogin.HasLoggedOutValidator(v); err != nil {
			return &ValidationError{Name: "has_logged_out", err: fmt.Errorf("ent: validator failed for field \"has_logged_out\": %w", err)}
		}
	}
	if v, ok := ulc.mutation.RequirePasswordChange(); ok {
		if err := userlogin.RequirePasswordChangeValidator(v); err != nil {
			return &ValidationError{Name: "require_password_change", err: fmt.Errorf("ent: validator failed for field \"require_password_change\": %w", err)}
		}
	}
	if v, ok := ulc.mutation.LastLocale(); ok {
		if err := userlogin.LastLocaleValidator(v); err != nil {
			return &ValidationError{Name: "last_locale", err: fmt.Errorf("ent: validator failed for field \"last_locale\": %w", err)}
		}
	}
	if v, ok := ulc.mutation.LastTimeZone(); ok {
		if err := userlogin.LastTimeZoneValidator(v); err != nil {
			return &ValidationError{Name: "last_time_zone", err: fmt.Errorf("ent: validator failed for field \"last_time_zone\": %w", err)}
		}
	}
	return nil
}

func (ulc *UserLoginCreate) sqlSave(ctx context.Context) (*UserLogin, error) {
	_node, _spec := ulc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ulc *UserLoginCreate) createSpec() (*UserLogin, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLogin{config: ulc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userlogin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlogin.FieldID,
			},
		}
	)
	if value, ok := ulc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlogin.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ulc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlogin.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ulc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := ulc.mutation.CurrentPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldCurrentPassword,
		})
		_node.CurrentPassword = value
	}
	if value, ok := ulc.mutation.PasswordHint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldPasswordHint,
		})
		_node.PasswordHint = value
	}
	if value, ok := ulc.mutation.IsSystem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: userlogin.FieldIsSystem,
		})
		_node.IsSystem = value
	}
	if value, ok := ulc.mutation.Enabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: userlogin.FieldEnabled,
		})
		_node.Enabled = value
	}
	if value, ok := ulc.mutation.HasLoggedOut(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: userlogin.FieldHasLoggedOut,
		})
		_node.HasLoggedOut = value
	}
	if value, ok := ulc.mutation.RequirePasswordChange(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: userlogin.FieldRequirePasswordChange,
		})
		_node.RequirePasswordChange = value
	}
	if value, ok := ulc.mutation.LastCurrencyUom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userlogin.FieldLastCurrencyUom,
		})
		_node.LastCurrencyUom = value
	}
	if value, ok := ulc.mutation.LastLocale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldLastLocale,
		})
		_node.LastLocale = value
	}
	if value, ok := ulc.mutation.LastTimeZone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldLastTimeZone,
		})
		_node.LastTimeZone = value
	}
	if value, ok := ulc.mutation.DisabledDateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlogin.FieldDisabledDateTime,
		})
		_node.DisabledDateTime = value
	}
	if value, ok := ulc.mutation.SuccessiveFailedLogins(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userlogin.FieldSuccessiveFailedLogins,
		})
		_node.SuccessiveFailedLogins = value
	}
	if value, ok := ulc.mutation.ExternalAuthID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldExternalAuthID,
		})
		_node.ExternalAuthID = value
	}
	if value, ok := ulc.mutation.UserLdapDn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldUserLdapDn,
		})
		_node.UserLdapDn = value
	}
	if value, ok := ulc.mutation.DisabledBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldDisabledBy,
		})
		_node.DisabledBy = value
	}
	if nodes := ulc.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userlogin.PartyTable,
			Columns: []string{userlogin.PartyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.party_user_logins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.PersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userlogin.PersonTable,
			Columns: []string{userlogin.PersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.person_user_logins = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.CreatedByPartiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.CreatedByPartiesTable,
			Columns: []string{userlogin.CreatedByPartiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.LastModifiedByPartiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.LastModifiedByPartiesTable,
			Columns: []string{userlogin.LastModifiedByPartiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: party.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.ChangeByPartyStatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.ChangeByPartyStatusesTable,
			Columns: []string{userlogin.ChangeByPartyStatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partystatus.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.UserLoginSecurityGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.UserLoginSecurityGroupsTable,
			Columns: []string{userlogin.UserLoginSecurityGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userloginsecuritygroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.UserPreferencesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.UserPreferencesTable,
			Columns: []string{userlogin.UserPreferencesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userpreference.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulc.mutation.AssignedByWorkEffortPartyAssignmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   userlogin.AssignedByWorkEffortPartyAssignmentsTable,
			Columns: []string{userlogin.AssignedByWorkEffortPartyAssignmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workeffortpartyassignment.FieldID,
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

// UserLoginCreateBulk is the builder for creating many UserLogin entities in bulk.
type UserLoginCreateBulk struct {
	config
	builders []*UserLoginCreate
}

// Save creates the UserLogin entities in the database.
func (ulcb *UserLoginCreateBulk) Save(ctx context.Context) ([]*UserLogin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ulcb.builders))
	nodes := make([]*UserLogin, len(ulcb.builders))
	mutators := make([]Mutator, len(ulcb.builders))
	for i := range ulcb.builders {
		func(i int, root context.Context) {
			builder := ulcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLoginMutation)
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
					_, err = mutators[i+1].Mutate(root, ulcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ulcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulcb *UserLoginCreateBulk) SaveX(ctx context.Context) []*UserLogin {
	v, err := ulcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
