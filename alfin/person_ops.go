// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

import "github.com/samlet/petrel/services"

type CreatePersonParams struct {
	PartyId                  string
	Salutation               string
	FirstName                string
	MiddleName               string
	LastName                 string
	PersonalTitle            string
	Suffix                   string
	Nickname                 string
	FirstNameLocal           string
	MiddleNameLocal          string
	LastNameLocal            string
	OtherLocal               string
	MemberId                 string
	Gender                   string
	BirthDate                services.DateTime
	DeceasedDate             services.DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       services.DateTime
	TotalYearsWorkExperience float64
	Comments                 string
	EmploymentStatusEnumId   string
	ResidenceStatusEnumId    string
	Occupation               string
	YearsWithEmployer        int64
	MonthsWithEmployer       int64
	ExistingCustomer         string
	CardId                   string
	PreferredCurrencyUomId   string
	Description              string
	ExternalId               string
	StatusId                 string
}

type CreatePersonResult struct {
	services.APIResponse
	PartyId string
}

type CreateUpdatePersonParams struct {
	PartyId                  string
	Salutation               string
	FirstName                string
	MiddleName               string
	LastName                 string
	PersonalTitle            string
	Suffix                   string
	Nickname                 string
	FirstNameLocal           string
	MiddleNameLocal          string
	LastNameLocal            string
	OtherLocal               string
	MemberId                 string
	Gender                   string
	BirthDate                services.DateTime
	DeceasedDate             services.DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       services.DateTime
	TotalYearsWorkExperience float64
	Comments                 string
	EmploymentStatusEnumId   string
	ResidenceStatusEnumId    string
	Occupation               string
	YearsWithEmployer        int64
	MonthsWithEmployer       int64
	ExistingCustomer         string
	CardId                   string
	UserLogin                services.MetaValue
}

type CreateUpdatePersonResult struct {
	services.APIResponse
	PartyId string
}

type UpdatePersonParams struct {
	PartyId                  string
	Salutation               string
	FirstName                string
	MiddleName               string
	LastName                 string
	PersonalTitle            string
	Suffix                   string
	Nickname                 string
	FirstNameLocal           string
	MiddleNameLocal          string
	LastNameLocal            string
	OtherLocal               string
	MemberId                 string
	Gender                   string
	BirthDate                services.DateTime
	DeceasedDate             services.DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       services.DateTime
	TotalYearsWorkExperience float64
	Comments                 string
	EmploymentStatusEnumId   string
	ResidenceStatusEnumId    string
	Occupation               string
	YearsWithEmployer        int64
	MonthsWithEmployer       int64
	ExistingCustomer         string
	CardId                   string
	PreferredCurrencyUomId   string
	Description              string
	ExternalId               string
	StatusId                 string
}

type UpdatePersonResult struct {
	services.APIResponse
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
