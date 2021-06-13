// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontactmech"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PersonCreate) SetCreateTime(t time.Time) *PersonCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PersonCreate) SetNillableCreateTime(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PersonCreate) SetUpdateTime(t time.Time) *PersonCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUpdateTime(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetStringRef sets the "string_ref" field.
func (pc *PersonCreate) SetStringRef(s string) *PersonCreate {
	pc.mutation.SetStringRef(s)
	return pc
}

// SetNillableStringRef sets the "string_ref" field if the given value is not nil.
func (pc *PersonCreate) SetNillableStringRef(s *string) *PersonCreate {
	if s != nil {
		pc.SetStringRef(*s)
	}
	return pc
}

// SetSalutation sets the "salutation" field.
func (pc *PersonCreate) SetSalutation(s string) *PersonCreate {
	pc.mutation.SetSalutation(s)
	return pc
}

// SetNillableSalutation sets the "salutation" field if the given value is not nil.
func (pc *PersonCreate) SetNillableSalutation(s *string) *PersonCreate {
	if s != nil {
		pc.SetSalutation(*s)
	}
	return pc
}

// SetFirstName sets the "first_name" field.
func (pc *PersonCreate) SetFirstName(s string) *PersonCreate {
	pc.mutation.SetFirstName(s)
	return pc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableFirstName(s *string) *PersonCreate {
	if s != nil {
		pc.SetFirstName(*s)
	}
	return pc
}

// SetMiddleName sets the "middle_name" field.
func (pc *PersonCreate) SetMiddleName(s string) *PersonCreate {
	pc.mutation.SetMiddleName(s)
	return pc
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMiddleName(s *string) *PersonCreate {
	if s != nil {
		pc.SetMiddleName(*s)
	}
	return pc
}

// SetLastName sets the "last_name" field.
func (pc *PersonCreate) SetLastName(s string) *PersonCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableLastName(s *string) *PersonCreate {
	if s != nil {
		pc.SetLastName(*s)
	}
	return pc
}

// SetPersonalTitle sets the "personal_title" field.
func (pc *PersonCreate) SetPersonalTitle(s string) *PersonCreate {
	pc.mutation.SetPersonalTitle(s)
	return pc
}

// SetNillablePersonalTitle sets the "personal_title" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePersonalTitle(s *string) *PersonCreate {
	if s != nil {
		pc.SetPersonalTitle(*s)
	}
	return pc
}

// SetSuffix sets the "suffix" field.
func (pc *PersonCreate) SetSuffix(s string) *PersonCreate {
	pc.mutation.SetSuffix(s)
	return pc
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (pc *PersonCreate) SetNillableSuffix(s *string) *PersonCreate {
	if s != nil {
		pc.SetSuffix(*s)
	}
	return pc
}

// SetNickname sets the "nickname" field.
func (pc *PersonCreate) SetNickname(s string) *PersonCreate {
	pc.mutation.SetNickname(s)
	return pc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (pc *PersonCreate) SetNillableNickname(s *string) *PersonCreate {
	if s != nil {
		pc.SetNickname(*s)
	}
	return pc
}

// SetFirstNameLocal sets the "first_name_local" field.
func (pc *PersonCreate) SetFirstNameLocal(s string) *PersonCreate {
	pc.mutation.SetFirstNameLocal(s)
	return pc
}

// SetNillableFirstNameLocal sets the "first_name_local" field if the given value is not nil.
func (pc *PersonCreate) SetNillableFirstNameLocal(s *string) *PersonCreate {
	if s != nil {
		pc.SetFirstNameLocal(*s)
	}
	return pc
}

// SetMiddleNameLocal sets the "middle_name_local" field.
func (pc *PersonCreate) SetMiddleNameLocal(s string) *PersonCreate {
	pc.mutation.SetMiddleNameLocal(s)
	return pc
}

// SetNillableMiddleNameLocal sets the "middle_name_local" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMiddleNameLocal(s *string) *PersonCreate {
	if s != nil {
		pc.SetMiddleNameLocal(*s)
	}
	return pc
}

// SetLastNameLocal sets the "last_name_local" field.
func (pc *PersonCreate) SetLastNameLocal(s string) *PersonCreate {
	pc.mutation.SetLastNameLocal(s)
	return pc
}

// SetNillableLastNameLocal sets the "last_name_local" field if the given value is not nil.
func (pc *PersonCreate) SetNillableLastNameLocal(s *string) *PersonCreate {
	if s != nil {
		pc.SetLastNameLocal(*s)
	}
	return pc
}

// SetOtherLocal sets the "other_local" field.
func (pc *PersonCreate) SetOtherLocal(s string) *PersonCreate {
	pc.mutation.SetOtherLocal(s)
	return pc
}

// SetNillableOtherLocal sets the "other_local" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOtherLocal(s *string) *PersonCreate {
	if s != nil {
		pc.SetOtherLocal(*s)
	}
	return pc
}

// SetMemberID sets the "member_id" field.
func (pc *PersonCreate) SetMemberID(i int) *PersonCreate {
	pc.mutation.SetMemberID(i)
	return pc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMemberID(i *int) *PersonCreate {
	if i != nil {
		pc.SetMemberID(*i)
	}
	return pc
}

// SetGender sets the "gender" field.
func (pc *PersonCreate) SetGender(pe person.Gender) *PersonCreate {
	pc.mutation.SetGender(pe)
	return pc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (pc *PersonCreate) SetNillableGender(pe *person.Gender) *PersonCreate {
	if pe != nil {
		pc.SetGender(*pe)
	}
	return pc
}

// SetBirthDate sets the "birth_date" field.
func (pc *PersonCreate) SetBirthDate(t time.Time) *PersonCreate {
	pc.mutation.SetBirthDate(t)
	return pc
}

// SetNillableBirthDate sets the "birth_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableBirthDate(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetBirthDate(*t)
	}
	return pc
}

// SetDeceasedDate sets the "deceased_date" field.
func (pc *PersonCreate) SetDeceasedDate(t time.Time) *PersonCreate {
	pc.mutation.SetDeceasedDate(t)
	return pc
}

// SetNillableDeceasedDate sets the "deceased_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDeceasedDate(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetDeceasedDate(*t)
	}
	return pc
}

// SetHeight sets the "height" field.
func (pc *PersonCreate) SetHeight(f float64) *PersonCreate {
	pc.mutation.SetHeight(f)
	return pc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHeight(f *float64) *PersonCreate {
	if f != nil {
		pc.SetHeight(*f)
	}
	return pc
}

// SetWeight sets the "weight" field.
func (pc *PersonCreate) SetWeight(f float64) *PersonCreate {
	pc.mutation.SetWeight(f)
	return pc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (pc *PersonCreate) SetNillableWeight(f *float64) *PersonCreate {
	if f != nil {
		pc.SetWeight(*f)
	}
	return pc
}

// SetMothersMaidenName sets the "mothers_maiden_name" field.
func (pc *PersonCreate) SetMothersMaidenName(s string) *PersonCreate {
	pc.mutation.SetMothersMaidenName(s)
	return pc
}

// SetNillableMothersMaidenName sets the "mothers_maiden_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMothersMaidenName(s *string) *PersonCreate {
	if s != nil {
		pc.SetMothersMaidenName(*s)
	}
	return pc
}

// SetOldMaritalStatus sets the "old_marital_status" field.
func (pc *PersonCreate) SetOldMaritalStatus(pms person.OldMaritalStatus) *PersonCreate {
	pc.mutation.SetOldMaritalStatus(pms)
	return pc
}

// SetNillableOldMaritalStatus sets the "old_marital_status" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOldMaritalStatus(pms *person.OldMaritalStatus) *PersonCreate {
	if pms != nil {
		pc.SetOldMaritalStatus(*pms)
	}
	return pc
}

// SetSocialSecurityNumber sets the "social_security_number" field.
func (pc *PersonCreate) SetSocialSecurityNumber(s string) *PersonCreate {
	pc.mutation.SetSocialSecurityNumber(s)
	return pc
}

// SetNillableSocialSecurityNumber sets the "social_security_number" field if the given value is not nil.
func (pc *PersonCreate) SetNillableSocialSecurityNumber(s *string) *PersonCreate {
	if s != nil {
		pc.SetSocialSecurityNumber(*s)
	}
	return pc
}

// SetPassportNumber sets the "passport_number" field.
func (pc *PersonCreate) SetPassportNumber(s string) *PersonCreate {
	pc.mutation.SetPassportNumber(s)
	return pc
}

// SetNillablePassportNumber sets the "passport_number" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePassportNumber(s *string) *PersonCreate {
	if s != nil {
		pc.SetPassportNumber(*s)
	}
	return pc
}

// SetPassportExpireDate sets the "passport_expire_date" field.
func (pc *PersonCreate) SetPassportExpireDate(t time.Time) *PersonCreate {
	pc.mutation.SetPassportExpireDate(t)
	return pc
}

// SetNillablePassportExpireDate sets the "passport_expire_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePassportExpireDate(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetPassportExpireDate(*t)
	}
	return pc
}

// SetTotalYearsWorkExperience sets the "total_years_work_experience" field.
func (pc *PersonCreate) SetTotalYearsWorkExperience(f float64) *PersonCreate {
	pc.mutation.SetTotalYearsWorkExperience(f)
	return pc
}

// SetNillableTotalYearsWorkExperience sets the "total_years_work_experience" field if the given value is not nil.
func (pc *PersonCreate) SetNillableTotalYearsWorkExperience(f *float64) *PersonCreate {
	if f != nil {
		pc.SetTotalYearsWorkExperience(*f)
	}
	return pc
}

// SetComments sets the "comments" field.
func (pc *PersonCreate) SetComments(s string) *PersonCreate {
	pc.mutation.SetComments(s)
	return pc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (pc *PersonCreate) SetNillableComments(s *string) *PersonCreate {
	if s != nil {
		pc.SetComments(*s)
	}
	return pc
}

// SetOccupation sets the "occupation" field.
func (pc *PersonCreate) SetOccupation(s string) *PersonCreate {
	pc.mutation.SetOccupation(s)
	return pc
}

// SetNillableOccupation sets the "occupation" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOccupation(s *string) *PersonCreate {
	if s != nil {
		pc.SetOccupation(*s)
	}
	return pc
}

// SetYearsWithEmployer sets the "years_with_employer" field.
func (pc *PersonCreate) SetYearsWithEmployer(i int) *PersonCreate {
	pc.mutation.SetYearsWithEmployer(i)
	return pc
}

// SetNillableYearsWithEmployer sets the "years_with_employer" field if the given value is not nil.
func (pc *PersonCreate) SetNillableYearsWithEmployer(i *int) *PersonCreate {
	if i != nil {
		pc.SetYearsWithEmployer(*i)
	}
	return pc
}

// SetMonthsWithEmployer sets the "months_with_employer" field.
func (pc *PersonCreate) SetMonthsWithEmployer(i int) *PersonCreate {
	pc.mutation.SetMonthsWithEmployer(i)
	return pc
}

// SetNillableMonthsWithEmployer sets the "months_with_employer" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMonthsWithEmployer(i *int) *PersonCreate {
	if i != nil {
		pc.SetMonthsWithEmployer(*i)
	}
	return pc
}

// SetExistingCustomer sets the "existing_customer" field.
func (pc *PersonCreate) SetExistingCustomer(value person.ExistingCustomer) *PersonCreate {
	pc.mutation.SetExistingCustomer(value)
	return pc
}

// SetNillableExistingCustomer sets the "existing_customer" field if the given value is not nil.
func (pc *PersonCreate) SetNillableExistingCustomer(value *person.ExistingCustomer) *PersonCreate {
	if value != nil {
		pc.SetExistingCustomer(*value)
	}
	return pc
}

// SetCardID sets the "card_id" field.
func (pc *PersonCreate) SetCardID(s string) *PersonCreate {
	pc.mutation.SetCardID(s)
	return pc
}

// SetNillableCardID sets the "card_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableCardID(s *string) *PersonCreate {
	if s != nil {
		pc.SetCardID(*s)
	}
	return pc
}

// SetPartyID sets the "party" edge to the Party entity by ID.
func (pc *PersonCreate) SetPartyID(id int) *PersonCreate {
	pc.mutation.SetPartyID(id)
	return pc
}

// SetNillablePartyID sets the "party" edge to the Party entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillablePartyID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetPartyID(*id)
	}
	return pc
}

// SetParty sets the "party" edge to the Party entity.
func (pc *PersonCreate) SetParty(p *Party) *PersonCreate {
	return pc.SetPartyID(p.ID)
}

// SetEmploymentStatusEnumerationID sets the "employment_status_enumeration" edge to the Enumeration entity by ID.
func (pc *PersonCreate) SetEmploymentStatusEnumerationID(id int) *PersonCreate {
	pc.mutation.SetEmploymentStatusEnumerationID(id)
	return pc
}

// SetNillableEmploymentStatusEnumerationID sets the "employment_status_enumeration" edge to the Enumeration entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableEmploymentStatusEnumerationID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetEmploymentStatusEnumerationID(*id)
	}
	return pc
}

// SetEmploymentStatusEnumeration sets the "employment_status_enumeration" edge to the Enumeration entity.
func (pc *PersonCreate) SetEmploymentStatusEnumeration(e *Enumeration) *PersonCreate {
	return pc.SetEmploymentStatusEnumerationID(e.ID)
}

// SetResidenceStatusEnumerationID sets the "residence_status_enumeration" edge to the Enumeration entity by ID.
func (pc *PersonCreate) SetResidenceStatusEnumerationID(id int) *PersonCreate {
	pc.mutation.SetResidenceStatusEnumerationID(id)
	return pc
}

// SetNillableResidenceStatusEnumerationID sets the "residence_status_enumeration" edge to the Enumeration entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableResidenceStatusEnumerationID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetResidenceStatusEnumerationID(*id)
	}
	return pc
}

// SetResidenceStatusEnumeration sets the "residence_status_enumeration" edge to the Enumeration entity.
func (pc *PersonCreate) SetResidenceStatusEnumeration(e *Enumeration) *PersonCreate {
	return pc.SetResidenceStatusEnumerationID(e.ID)
}

// SetMaritalStatusEnumerationID sets the "marital_status_enumeration" edge to the Enumeration entity by ID.
func (pc *PersonCreate) SetMaritalStatusEnumerationID(id int) *PersonCreate {
	pc.mutation.SetMaritalStatusEnumerationID(id)
	return pc
}

// SetNillableMaritalStatusEnumerationID sets the "marital_status_enumeration" edge to the Enumeration entity by ID if the given value is not nil.
func (pc *PersonCreate) SetNillableMaritalStatusEnumerationID(id *int) *PersonCreate {
	if id != nil {
		pc = pc.SetMaritalStatusEnumerationID(*id)
	}
	return pc
}

// SetMaritalStatusEnumeration sets the "marital_status_enumeration" edge to the Enumeration entity.
func (pc *PersonCreate) SetMaritalStatusEnumeration(e *Enumeration) *PersonCreate {
	return pc.SetMaritalStatusEnumerationID(e.ID)
}

// AddPartyContactMechIDs adds the "party_contact_meches" edge to the PartyContactMech entity by IDs.
func (pc *PersonCreate) AddPartyContactMechIDs(ids ...int) *PersonCreate {
	pc.mutation.AddPartyContactMechIDs(ids...)
	return pc
}

// AddPartyContactMeches adds the "party_contact_meches" edges to the PartyContactMech entity.
func (pc *PersonCreate) AddPartyContactMeches(p ...*PartyContactMech) *PersonCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPartyContactMechIDs(ids...)
}

// AddUserLoginIDs adds the "user_logins" edge to the UserLogin entity by IDs.
func (pc *PersonCreate) AddUserLoginIDs(ids ...int) *PersonCreate {
	pc.mutation.AddUserLoginIDs(ids...)
	return pc
}

// AddUserLogins adds the "user_logins" edges to the UserLogin entity.
func (pc *PersonCreate) AddUserLogins(u ...*UserLogin) *PersonCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserLoginIDs(ids...)
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	var (
		err  error
		node *Person
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PersonCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := person.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := person.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.BirthDate(); !ok {
		v := person.DefaultBirthDate()
		pc.mutation.SetBirthDate(v)
	}
	if _, ok := pc.mutation.DeceasedDate(); !ok {
		v := person.DefaultDeceasedDate()
		pc.mutation.SetDeceasedDate(v)
	}
	if _, ok := pc.mutation.PassportExpireDate(); !ok {
		v := person.DefaultPassportExpireDate()
		pc.mutation.SetPassportExpireDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := person.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	if v, ok := pc.mutation.OldMaritalStatus(); ok {
		if err := person.OldMaritalStatusValidator(v); err != nil {
			return &ValidationError{Name: "old_marital_status", err: fmt.Errorf("ent: validator failed for field \"old_marital_status\": %w", err)}
		}
	}
	if v, ok := pc.mutation.ExistingCustomer(); ok {
		if err := person.ExistingCustomerValidator(v); err != nil {
			return &ValidationError{Name: "existing_customer", err: fmt.Errorf("ent: validator failed for field \"existing_customer\": %w", err)}
		}
	}
	if v, ok := pc.mutation.CardID(); ok {
		if err := person.CardIDValidator(v); err != nil {
			return &ValidationError{Name: "card_id", err: fmt.Errorf("ent: validator failed for field \"card_id\": %w", err)}
		}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: person.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: person.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: person.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: person.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.StringRef(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldStringRef,
		})
		_node.StringRef = value
	}
	if value, ok := pc.mutation.Salutation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldSalutation,
		})
		_node.Salutation = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := pc.mutation.MiddleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldMiddleName,
		})
		_node.MiddleName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := pc.mutation.PersonalTitle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldPersonalTitle,
		})
		_node.PersonalTitle = value
	}
	if value, ok := pc.mutation.Suffix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldSuffix,
		})
		_node.Suffix = value
	}
	if value, ok := pc.mutation.Nickname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldNickname,
		})
		_node.Nickname = value
	}
	if value, ok := pc.mutation.FirstNameLocal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldFirstNameLocal,
		})
		_node.FirstNameLocal = value
	}
	if value, ok := pc.mutation.MiddleNameLocal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldMiddleNameLocal,
		})
		_node.MiddleNameLocal = value
	}
	if value, ok := pc.mutation.LastNameLocal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldLastNameLocal,
		})
		_node.LastNameLocal = value
	}
	if value, ok := pc.mutation.OtherLocal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldOtherLocal,
		})
		_node.OtherLocal = value
	}
	if value, ok := pc.mutation.MemberID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: person.FieldMemberID,
		})
		_node.MemberID = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: person.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := pc.mutation.BirthDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: person.FieldBirthDate,
		})
		_node.BirthDate = value
	}
	if value, ok := pc.mutation.DeceasedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: person.FieldDeceasedDate,
		})
		_node.DeceasedDate = value
	}
	if value, ok := pc.mutation.Height(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: person.FieldHeight,
		})
		_node.Height = value
	}
	if value, ok := pc.mutation.Weight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: person.FieldWeight,
		})
		_node.Weight = value
	}
	if value, ok := pc.mutation.MothersMaidenName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldMothersMaidenName,
		})
		_node.MothersMaidenName = value
	}
	if value, ok := pc.mutation.OldMaritalStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: person.FieldOldMaritalStatus,
		})
		_node.OldMaritalStatus = value
	}
	if value, ok := pc.mutation.SocialSecurityNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldSocialSecurityNumber,
		})
		_node.SocialSecurityNumber = value
	}
	if value, ok := pc.mutation.PassportNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldPassportNumber,
		})
		_node.PassportNumber = value
	}
	if value, ok := pc.mutation.PassportExpireDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: person.FieldPassportExpireDate,
		})
		_node.PassportExpireDate = value
	}
	if value, ok := pc.mutation.TotalYearsWorkExperience(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: person.FieldTotalYearsWorkExperience,
		})
		_node.TotalYearsWorkExperience = value
	}
	if value, ok := pc.mutation.Comments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldComments,
		})
		_node.Comments = value
	}
	if value, ok := pc.mutation.Occupation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldOccupation,
		})
		_node.Occupation = value
	}
	if value, ok := pc.mutation.YearsWithEmployer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: person.FieldYearsWithEmployer,
		})
		_node.YearsWithEmployer = value
	}
	if value, ok := pc.mutation.MonthsWithEmployer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: person.FieldMonthsWithEmployer,
		})
		_node.MonthsWithEmployer = value
	}
	if value, ok := pc.mutation.ExistingCustomer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: person.FieldExistingCustomer,
		})
		_node.ExistingCustomer = value
	}
	if value, ok := pc.mutation.CardID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: person.FieldCardID,
		})
		_node.CardID = value
	}
	if nodes := pc.mutation.PartyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   person.PartyTable,
			Columns: []string{person.PartyColumn},
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
		_node.party_person = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EmploymentStatusEnumerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.EmploymentStatusEnumerationTable,
			Columns: []string{person.EmploymentStatusEnumerationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumeration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_employment_status_people = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ResidenceStatusEnumerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.ResidenceStatusEnumerationTable,
			Columns: []string{person.ResidenceStatusEnumerationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumeration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_residence_status_people = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MaritalStatusEnumerationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   person.MaritalStatusEnumerationTable,
			Columns: []string{person.MaritalStatusEnumerationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: enumeration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.enumeration_marital_status_people = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PartyContactMechesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.PartyContactMechesTable,
			Columns: []string{person.PartyContactMechesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: partycontactmech.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UserLoginsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   person.UserLoginsTable,
			Columns: []string{person.UserLoginsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlogin.FieldID,
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

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	builders []*PersonCreate
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
