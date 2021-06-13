// Code generated by entc, DO NOT EDIT.

package partyrole

import (
	"time"
)

const (
	// Label holds the string label denoting the partyrole type in the database.
	Label = "party_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// EdgeParty holds the string denoting the party edge name in mutations.
	EdgeParty = "party"
	// EdgeRoleType holds the string denoting the role_type edge name in mutations.
	EdgeRoleType = "role_type"
	// EdgeFixedAssets holds the string denoting the fixed_assets edge name in mutations.
	EdgeFixedAssets = "fixed_assets"
	// EdgePartyContactMeches holds the string denoting the party_contact_meches edge name in mutations.
	EdgePartyContactMeches = "party_contact_meches"
	// EdgeWorkEffortPartyAssignments holds the string denoting the work_effort_party_assignments edge name in mutations.
	EdgeWorkEffortPartyAssignments = "work_effort_party_assignments"
	// Table holds the table name of the partyrole in the database.
	Table = "party_roles"
	// PartyTable is the table the holds the party relation/edge.
	PartyTable = "party_roles"
	// PartyInverseTable is the table name for the Party entity.
	// It exists in this package in order to avoid circular dependency with the "party" package.
	PartyInverseTable = "parties"
	// PartyColumn is the table column denoting the party relation/edge.
	PartyColumn = "party_party_roles"
	// RoleTypeTable is the table the holds the role_type relation/edge.
	RoleTypeTable = "party_roles"
	// RoleTypeInverseTable is the table name for the RoleType entity.
	// It exists in this package in order to avoid circular dependency with the "roletype" package.
	RoleTypeInverseTable = "role_types"
	// RoleTypeColumn is the table column denoting the role_type relation/edge.
	RoleTypeColumn = "role_type_party_roles"
	// FixedAssetsTable is the table the holds the fixed_assets relation/edge.
	FixedAssetsTable = "fixed_assets"
	// FixedAssetsInverseTable is the table name for the FixedAsset entity.
	// It exists in this package in order to avoid circular dependency with the "fixedasset" package.
	FixedAssetsInverseTable = "fixed_assets"
	// FixedAssetsColumn is the table column denoting the fixed_assets relation/edge.
	FixedAssetsColumn = "party_role_fixed_assets"
	// PartyContactMechesTable is the table the holds the party_contact_meches relation/edge.
	PartyContactMechesTable = "party_contact_meches"
	// PartyContactMechesInverseTable is the table name for the PartyContactMech entity.
	// It exists in this package in order to avoid circular dependency with the "partycontactmech" package.
	PartyContactMechesInverseTable = "party_contact_meches"
	// PartyContactMechesColumn is the table column denoting the party_contact_meches relation/edge.
	PartyContactMechesColumn = "party_role_party_contact_meches"
	// WorkEffortPartyAssignmentsTable is the table the holds the work_effort_party_assignments relation/edge.
	WorkEffortPartyAssignmentsTable = "work_effort_party_assignments"
	// WorkEffortPartyAssignmentsInverseTable is the table name for the WorkEffortPartyAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "workeffortpartyassignment" package.
	WorkEffortPartyAssignmentsInverseTable = "work_effort_party_assignments"
	// WorkEffortPartyAssignmentsColumn is the table column denoting the work_effort_party_assignments relation/edge.
	WorkEffortPartyAssignmentsColumn = "party_role_work_effort_party_assignments"
)

// Columns holds all SQL columns for partyrole fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "party_roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"party_party_roles",
	"role_type_party_roles",
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
)
