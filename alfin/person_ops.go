// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

type CreatePersonParams struct {
	PartyId                  interface{}
	Salutation               interface{}
	FirstName                interface{}
	MiddleName               interface{}
	LastName                 interface{}
	PersonalTitle            interface{}
	Suffix                   interface{}
	Nickname                 interface{}
	FirstNameLocal           interface{}
	MiddleNameLocal          interface{}
	LastNameLocal            interface{}
	OtherLocal               interface{}
	MemberId                 interface{}
	Gender                   interface{}
	BirthDate                interface{}
	DeceasedDate             interface{}
	Height                   interface{}
	Weight                   interface{}
	MothersMaidenName        interface{}
	OldMaritalStatus         interface{}
	MaritalStatusEnumId      interface{}
	SocialSecurityNumber     interface{}
	PassportNumber           interface{}
	PassportExpireDate       interface{}
	TotalYearsWorkExperience interface{}
	Comments                 interface{}
	EmploymentStatusEnumId   interface{}
	ResidenceStatusEnumId    interface{}
	Occupation               interface{}
	YearsWithEmployer        interface{}
	MonthsWithEmployer       interface{}
	ExistingCustomer         interface{}
	CardId                   interface{}
	PreferredCurrencyUomId   interface{}
	Description              interface{}
	ExternalId               interface{}
	StatusId                 interface{}
}

type CreatePersonResult struct {
	APIResponse
	PartyId interface{}
}

type CreateUpdatePersonParams struct {
	PartyId                  interface{}
	Salutation               interface{}
	FirstName                interface{}
	MiddleName               interface{}
	LastName                 interface{}
	PersonalTitle            interface{}
	Suffix                   interface{}
	Nickname                 interface{}
	FirstNameLocal           interface{}
	MiddleNameLocal          interface{}
	LastNameLocal            interface{}
	OtherLocal               interface{}
	MemberId                 interface{}
	Gender                   interface{}
	BirthDate                interface{}
	DeceasedDate             interface{}
	Height                   interface{}
	Weight                   interface{}
	MothersMaidenName        interface{}
	OldMaritalStatus         interface{}
	MaritalStatusEnumId      interface{}
	SocialSecurityNumber     interface{}
	PassportNumber           interface{}
	PassportExpireDate       interface{}
	TotalYearsWorkExperience interface{}
	Comments                 interface{}
	EmploymentStatusEnumId   interface{}
	ResidenceStatusEnumId    interface{}
	Occupation               interface{}
	YearsWithEmployer        interface{}
	MonthsWithEmployer       interface{}
	ExistingCustomer         interface{}
	CardId                   interface{}
	UserLogin                interface{}
}

type CreateUpdatePersonResult struct {
	APIResponse
	PartyId interface{}
}

type UpdatePersonParams struct {
	PartyId                  interface{}
	Salutation               interface{}
	FirstName                interface{}
	MiddleName               interface{}
	LastName                 interface{}
	PersonalTitle            interface{}
	Suffix                   interface{}
	Nickname                 interface{}
	FirstNameLocal           interface{}
	MiddleNameLocal          interface{}
	LastNameLocal            interface{}
	OtherLocal               interface{}
	MemberId                 interface{}
	Gender                   interface{}
	BirthDate                interface{}
	DeceasedDate             interface{}
	Height                   interface{}
	Weight                   interface{}
	MothersMaidenName        interface{}
	OldMaritalStatus         interface{}
	MaritalStatusEnumId      interface{}
	SocialSecurityNumber     interface{}
	PassportNumber           interface{}
	PassportExpireDate       interface{}
	TotalYearsWorkExperience interface{}
	Comments                 interface{}
	EmploymentStatusEnumId   interface{}
	ResidenceStatusEnumId    interface{}
	Occupation               interface{}
	YearsWithEmployer        interface{}
	MonthsWithEmployer       interface{}
	ExistingCustomer         interface{}
	CardId                   interface{}
	PreferredCurrencyUomId   interface{}
	Description              interface{}
	ExternalId               interface{}
	StatusId                 interface{}
}

type UpdatePersonResult struct {
	APIResponse
}

// Interface
type PersonOps interface {
	// CreatePerson Create a Person
	CreatePerson(params *CreatePersonParams) (*CreatePersonResult, error)
	// CreateUpdatePerson Create and Update a person
	CreateUpdatePerson(params *CreateUpdatePersonParams) (*CreateUpdatePersonResult, error)
	// UpdatePerson Update a Person
	UpdatePerson(params *UpdatePersonParams) (*UpdatePersonResult, error)
}
