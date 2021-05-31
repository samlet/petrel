// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alfin

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
	BirthDate                DateTime
	DeceasedDate             DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       DateTime
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
	APIResponse
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
	BirthDate                DateTime
	DeceasedDate             DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       DateTime
	TotalYearsWorkExperience float64
	Comments                 string
	EmploymentStatusEnumId   string
	ResidenceStatusEnumId    string
	Occupation               string
	YearsWithEmployer        int64
	MonthsWithEmployer       int64
	ExistingCustomer         string
	CardId                   string
	UserLogin                MetaValue
}

type CreateUpdatePersonResult struct {
	APIResponse
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
	BirthDate                DateTime
	DeceasedDate             DateTime
	Height                   float64
	Weight                   float64
	MothersMaidenName        string
	OldMaritalStatus         string
	MaritalStatusEnumId      string
	SocialSecurityNumber     string
	PassportNumber           string
	PassportExpireDate       DateTime
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
