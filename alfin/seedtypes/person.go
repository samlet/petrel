package seedtypes
import "github.com/samlet/petrel/services"


type Person struct {
    PartyId *string `json:"partyId,omitempty" xml:"partyId,attr"`
    Salutation *string `json:"salutation,omitempty" xml:"salutation,attr"`
    FirstName *string `json:"firstName,omitempty" xml:"firstName,attr"`
    MiddleName *string `json:"middleName,omitempty" xml:"middleName,attr"`
    LastName *string `json:"lastName,omitempty" xml:"lastName,attr"`
    PersonalTitle *string `json:"personalTitle,omitempty" xml:"personalTitle,attr"`
    Suffix *string `json:"suffix,omitempty" xml:"suffix,attr"`
    Nickname *string `json:"nickname,omitempty" xml:"nickname,attr"`
    FirstNameLocal *string `json:"firstNameLocal,omitempty" xml:"firstNameLocal,attr"`
    MiddleNameLocal *string `json:"middleNameLocal,omitempty" xml:"middleNameLocal,attr"`
    LastNameLocal *string `json:"lastNameLocal,omitempty" xml:"lastNameLocal,attr"`
    OtherLocal *string `json:"otherLocal,omitempty" xml:"otherLocal,attr"`
    MemberId *string `json:"memberId,omitempty" xml:"memberId,attr"`
    Gender *byte `json:"gender,omitempty" xml:"gender,attr"`
    BirthDate *services.DateTime `json:"birthDate,omitempty" xml:"birthDate,attr"`
    DeceasedDate *services.DateTime `json:"deceasedDate,omitempty" xml:"deceasedDate,attr"`
    Height *float64 `json:"height,omitempty" xml:"height,attr"`
    Weight *float64 `json:"weight,omitempty" xml:"weight,attr"`
    MothersMaidenName *string `json:"mothersMaidenName,omitempty" xml:"mothersMaidenName,attr"`
    OldMaritalStatus *byte `json:"oldMaritalStatus,omitempty" xml:"oldMaritalStatus,attr"`
    MaritalStatusEnumId *string `json:"maritalStatusEnumId,omitempty" xml:"maritalStatusEnumId,attr"`
    SocialSecurityNumber *string `json:"socialSecurityNumber,omitempty" xml:"socialSecurityNumber,attr"`
    PassportNumber *string `json:"passportNumber,omitempty" xml:"passportNumber,attr"`
    PassportExpireDate *services.DateTime `json:"passportExpireDate,omitempty" xml:"passportExpireDate,attr"`
    TotalYearsWorkExperience *float64 `json:"totalYearsWorkExperience,omitempty" xml:"totalYearsWorkExperience,attr"`
    Comments *string `json:"comments,omitempty" xml:"comments,attr"`
    EmploymentStatusEnumId *string `json:"employmentStatusEnumId,omitempty" xml:"employmentStatusEnumId,attr"`
    ResidenceStatusEnumId *string `json:"residenceStatusEnumId,omitempty" xml:"residenceStatusEnumId,attr"`
    Occupation *string `json:"occupation,omitempty" xml:"occupation,attr"`
    YearsWithEmployer *int64 `json:"yearsWithEmployer,omitempty" xml:"yearsWithEmployer,attr"`
    MonthsWithEmployer *int64 `json:"monthsWithEmployer,omitempty" xml:"monthsWithEmployer,attr"`
    ExistingCustomer *byte `json:"existingCustomer,omitempty" xml:"existingCustomer,attr"`
    CardId *string `json:"cardId,omitempty" xml:"cardId,attr"`
    LastUpdatedStamp *services.DateTime `json:"lastUpdatedStamp,omitempty" xml:"lastUpdatedStamp,attr"`
    LastUpdatedTxStamp *services.DateTime `json:"lastUpdatedTxStamp,omitempty" xml:"lastUpdatedTxStamp,attr"`
    CreatedStamp *services.DateTime `json:"createdStamp,omitempty" xml:"createdStamp,attr"`
    CreatedTxStamp *services.DateTime `json:"createdTxStamp,omitempty" xml:"createdTxStamp,attr"`
}

