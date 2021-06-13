// Code generated by entc, DO NOT EDIT.

package userlogin

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the userlogin type in the database.
	Label = "user_login"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldCurrentPassword holds the string denoting the current_password field in the database.
	FieldCurrentPassword = "current_password"
	// FieldPasswordHint holds the string denoting the password_hint field in the database.
	FieldPasswordHint = "password_hint"
	// FieldIsSystem holds the string denoting the is_system field in the database.
	FieldIsSystem = "is_system"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldHasLoggedOut holds the string denoting the has_logged_out field in the database.
	FieldHasLoggedOut = "has_logged_out"
	// FieldRequirePasswordChange holds the string denoting the require_password_change field in the database.
	FieldRequirePasswordChange = "require_password_change"
	// FieldLastCurrencyUom holds the string denoting the last_currency_uom field in the database.
	FieldLastCurrencyUom = "last_currency_uom"
	// FieldLastLocale holds the string denoting the last_locale field in the database.
	FieldLastLocale = "last_locale"
	// FieldLastTimeZone holds the string denoting the last_time_zone field in the database.
	FieldLastTimeZone = "last_time_zone"
	// FieldDisabledDateTime holds the string denoting the disabled_date_time field in the database.
	FieldDisabledDateTime = "disabled_date_time"
	// FieldSuccessiveFailedLogins holds the string denoting the successive_failed_logins field in the database.
	FieldSuccessiveFailedLogins = "successive_failed_logins"
	// FieldExternalAuthID holds the string denoting the external_auth_id field in the database.
	FieldExternalAuthID = "external_auth_id"
	// FieldUserLdapDn holds the string denoting the user_ldap_dn field in the database.
	FieldUserLdapDn = "user_ldap_dn"
	// FieldDisabledBy holds the string denoting the disabled_by field in the database.
	FieldDisabledBy = "disabled_by"
	// EdgeParty holds the string denoting the party edge name in mutations.
	EdgeParty = "party"
	// EdgePerson holds the string denoting the person edge name in mutations.
	EdgePerson = "person"
	// EdgeCreatedByParties holds the string denoting the created_by_parties edge name in mutations.
	EdgeCreatedByParties = "created_by_parties"
	// EdgeLastModifiedByParties holds the string denoting the last_modified_by_parties edge name in mutations.
	EdgeLastModifiedByParties = "last_modified_by_parties"
	// EdgeChangeByPartyStatuses holds the string denoting the change_by_party_statuses edge name in mutations.
	EdgeChangeByPartyStatuses = "change_by_party_statuses"
	// EdgeUserLoginSecurityGroups holds the string denoting the user_login_security_groups edge name in mutations.
	EdgeUserLoginSecurityGroups = "user_login_security_groups"
	// EdgeUserPreferences holds the string denoting the user_preferences edge name in mutations.
	EdgeUserPreferences = "user_preferences"
	// EdgeAssignedByWorkEffortPartyAssignments holds the string denoting the assigned_by_work_effort_party_assignments edge name in mutations.
	EdgeAssignedByWorkEffortPartyAssignments = "assigned_by_work_effort_party_assignments"
	// Table holds the table name of the userlogin in the database.
	Table = "user_logins"
	// PartyTable is the table the holds the party relation/edge.
	PartyTable = "user_logins"
	// PartyInverseTable is the table name for the Party entity.
	// It exists in this package in order to avoid circular dependency with the "party" package.
	PartyInverseTable = "parties"
	// PartyColumn is the table column denoting the party relation/edge.
	PartyColumn = "party_user_logins"
	// PersonTable is the table the holds the person relation/edge.
	PersonTable = "user_logins"
	// PersonInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	PersonInverseTable = "persons"
	// PersonColumn is the table column denoting the person relation/edge.
	PersonColumn = "person_user_logins"
	// CreatedByPartiesTable is the table the holds the created_by_parties relation/edge.
	CreatedByPartiesTable = "parties"
	// CreatedByPartiesInverseTable is the table name for the Party entity.
	// It exists in this package in order to avoid circular dependency with the "party" package.
	CreatedByPartiesInverseTable = "parties"
	// CreatedByPartiesColumn is the table column denoting the created_by_parties relation/edge.
	CreatedByPartiesColumn = "user_login_created_by_parties"
	// LastModifiedByPartiesTable is the table the holds the last_modified_by_parties relation/edge.
	LastModifiedByPartiesTable = "parties"
	// LastModifiedByPartiesInverseTable is the table name for the Party entity.
	// It exists in this package in order to avoid circular dependency with the "party" package.
	LastModifiedByPartiesInverseTable = "parties"
	// LastModifiedByPartiesColumn is the table column denoting the last_modified_by_parties relation/edge.
	LastModifiedByPartiesColumn = "user_login_last_modified_by_parties"
	// ChangeByPartyStatusesTable is the table the holds the change_by_party_statuses relation/edge.
	ChangeByPartyStatusesTable = "party_status"
	// ChangeByPartyStatusesInverseTable is the table name for the PartyStatus entity.
	// It exists in this package in order to avoid circular dependency with the "partystatus" package.
	ChangeByPartyStatusesInverseTable = "party_status"
	// ChangeByPartyStatusesColumn is the table column denoting the change_by_party_statuses relation/edge.
	ChangeByPartyStatusesColumn = "user_login_change_by_party_statuses"
	// UserLoginSecurityGroupsTable is the table the holds the user_login_security_groups relation/edge.
	UserLoginSecurityGroupsTable = "user_login_security_groups"
	// UserLoginSecurityGroupsInverseTable is the table name for the UserLoginSecurityGroup entity.
	// It exists in this package in order to avoid circular dependency with the "userloginsecuritygroup" package.
	UserLoginSecurityGroupsInverseTable = "user_login_security_groups"
	// UserLoginSecurityGroupsColumn is the table column denoting the user_login_security_groups relation/edge.
	UserLoginSecurityGroupsColumn = "user_login_user_login_security_groups"
	// UserPreferencesTable is the table the holds the user_preferences relation/edge.
	UserPreferencesTable = "user_preferences"
	// UserPreferencesInverseTable is the table name for the UserPreference entity.
	// It exists in this package in order to avoid circular dependency with the "userpreference" package.
	UserPreferencesInverseTable = "user_preferences"
	// UserPreferencesColumn is the table column denoting the user_preferences relation/edge.
	UserPreferencesColumn = "user_login_user_preferences"
	// AssignedByWorkEffortPartyAssignmentsTable is the table the holds the assigned_by_work_effort_party_assignments relation/edge.
	AssignedByWorkEffortPartyAssignmentsTable = "work_effort_party_assignments"
	// AssignedByWorkEffortPartyAssignmentsInverseTable is the table name for the WorkEffortPartyAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "workeffortpartyassignment" package.
	AssignedByWorkEffortPartyAssignmentsInverseTable = "work_effort_party_assignments"
	// AssignedByWorkEffortPartyAssignmentsColumn is the table column denoting the assigned_by_work_effort_party_assignments relation/edge.
	AssignedByWorkEffortPartyAssignmentsColumn = "user_login_assigned_by_work_effort_party_assignments"
)

// Columns holds all SQL columns for userlogin fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldCurrentPassword,
	FieldPasswordHint,
	FieldIsSystem,
	FieldEnabled,
	FieldHasLoggedOut,
	FieldRequirePasswordChange,
	FieldLastCurrencyUom,
	FieldLastLocale,
	FieldLastTimeZone,
	FieldDisabledDateTime,
	FieldSuccessiveFailedLogins,
	FieldExternalAuthID,
	FieldUserLdapDn,
	FieldDisabledBy,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "user_logins"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"party_user_logins",
	"person_user_logins",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// LastLocaleValidator is a validator for the "last_locale" field. It is called by the builders before save.
	LastLocaleValidator func(string) error
	// LastTimeZoneValidator is a validator for the "last_time_zone" field. It is called by the builders before save.
	LastTimeZoneValidator func(string) error
	// DefaultDisabledDateTime holds the default value on creation for the "disabled_date_time" field.
	DefaultDisabledDateTime func() time.Time
)

// IsSystem defines the type for the "is_system" enum field.
type IsSystem string

// IsSystem values.
const (
	IsSystemYes     IsSystem = "Yes"
	IsSystemNo      IsSystem = "No"
	IsSystemUnknown IsSystem = "Unknown"
)

func (is IsSystem) String() string {
	return string(is)
}

// IsSystemValidator is a validator for the "is_system" field enum values. It is called by the builders before save.
func IsSystemValidator(is IsSystem) error {
	switch is {
	case IsSystemYes, IsSystemNo, IsSystemUnknown:
		return nil
	default:
		return fmt.Errorf("userlogin: invalid enum value for is_system field: %q", is)
	}
}

// Enabled defines the type for the "enabled" enum field.
type Enabled string

// Enabled values.
const (
	EnabledYes     Enabled = "Yes"
	EnabledNo      Enabled = "No"
	EnabledUnknown Enabled = "Unknown"
)

func (e Enabled) String() string {
	return string(e)
}

// EnabledValidator is a validator for the "enabled" field enum values. It is called by the builders before save.
func EnabledValidator(e Enabled) error {
	switch e {
	case EnabledYes, EnabledNo, EnabledUnknown:
		return nil
	default:
		return fmt.Errorf("userlogin: invalid enum value for enabled field: %q", e)
	}
}

// HasLoggedOut defines the type for the "has_logged_out" enum field.
type HasLoggedOut string

// HasLoggedOut values.
const (
	HasLoggedOutYes     HasLoggedOut = "Yes"
	HasLoggedOutNo      HasLoggedOut = "No"
	HasLoggedOutUnknown HasLoggedOut = "Unknown"
)

func (hlo HasLoggedOut) String() string {
	return string(hlo)
}

// HasLoggedOutValidator is a validator for the "has_logged_out" field enum values. It is called by the builders before save.
func HasLoggedOutValidator(hlo HasLoggedOut) error {
	switch hlo {
	case HasLoggedOutYes, HasLoggedOutNo, HasLoggedOutUnknown:
		return nil
	default:
		return fmt.Errorf("userlogin: invalid enum value for has_logged_out field: %q", hlo)
	}
}

// RequirePasswordChange defines the type for the "require_password_change" enum field.
type RequirePasswordChange string

// RequirePasswordChange values.
const (
	RequirePasswordChangeYes     RequirePasswordChange = "Yes"
	RequirePasswordChangeNo      RequirePasswordChange = "No"
	RequirePasswordChangeUnknown RequirePasswordChange = "Unknown"
)

func (rpc RequirePasswordChange) String() string {
	return string(rpc)
}

// RequirePasswordChangeValidator is a validator for the "require_password_change" field enum values. It is called by the builders before save.
func RequirePasswordChangeValidator(rpc RequirePasswordChange) error {
	switch rpc {
	case RequirePasswordChangeYes, RequirePasswordChangeNo, RequirePasswordChangeUnknown:
		return nil
	default:
		return fmt.Errorf("userlogin: invalid enum value for require_password_change field: %q", rpc)
	}
}
