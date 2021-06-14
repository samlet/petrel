// Code generated by entc, DO NOT EDIT.

package roletype

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the roletype type in the database.
	Label = "role_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldHasTable holds the string denoting the has_table field in the database.
	FieldHasTable = "has_table"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeFixedAssets holds the string denoting the fixed_assets edge name in mutations.
	EdgeFixedAssets = "fixed_assets"
	// EdgePartyContactMeches holds the string denoting the party_contact_meches edge name in mutations.
	EdgePartyContactMeches = "party_contact_meches"
	// EdgeValidFromPartyRelationshipTypes holds the string denoting the valid_from_party_relationship_types edge name in mutations.
	EdgeValidFromPartyRelationshipTypes = "valid_from_party_relationship_types"
	// EdgeValidToPartyRelationshipTypes holds the string denoting the valid_to_party_relationship_types edge name in mutations.
	EdgeValidToPartyRelationshipTypes = "valid_to_party_relationship_types"
	// EdgePartyRoles holds the string denoting the party_roles edge name in mutations.
	EdgePartyRoles = "party_roles"
	// EdgeChildRoleTypes holds the string denoting the child_role_types edge name in mutations.
	EdgeChildRoleTypes = "child_role_types"
	// EdgeWorkEffortPartyAssignments holds the string denoting the work_effort_party_assignments edge name in mutations.
	EdgeWorkEffortPartyAssignments = "work_effort_party_assignments"
	// Table holds the table name of the roletype in the database.
	Table = "role_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "role_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "role_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "role_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "role_type_children"
	// FixedAssetsTable is the table the holds the fixed_assets relation/edge.
	FixedAssetsTable = "fixed_assets"
	// FixedAssetsInverseTable is the table name for the FixedAsset entity.
	// It exists in this package in order to avoid circular dependency with the "fixedasset" package.
	FixedAssetsInverseTable = "fixed_assets"
	// FixedAssetsColumn is the table column denoting the fixed_assets relation/edge.
	FixedAssetsColumn = "role_type_fixed_assets"
	// PartyContactMechesTable is the table the holds the party_contact_meches relation/edge.
	PartyContactMechesTable = "party_contact_meches"
	// PartyContactMechesInverseTable is the table name for the PartyContactMech entity.
	// It exists in this package in order to avoid circular dependency with the "partycontactmech" package.
	PartyContactMechesInverseTable = "party_contact_meches"
	// PartyContactMechesColumn is the table column denoting the party_contact_meches relation/edge.
	PartyContactMechesColumn = "role_type_party_contact_meches"
	// ValidFromPartyRelationshipTypesTable is the table the holds the valid_from_party_relationship_types relation/edge.
	ValidFromPartyRelationshipTypesTable = "party_relationship_types"
	// ValidFromPartyRelationshipTypesInverseTable is the table name for the PartyRelationshipType entity.
	// It exists in this package in order to avoid circular dependency with the "partyrelationshiptype" package.
	ValidFromPartyRelationshipTypesInverseTable = "party_relationship_types"
	// ValidFromPartyRelationshipTypesColumn is the table column denoting the valid_from_party_relationship_types relation/edge.
	ValidFromPartyRelationshipTypesColumn = "role_type_valid_from_party_relationship_types"
	// ValidToPartyRelationshipTypesTable is the table the holds the valid_to_party_relationship_types relation/edge.
	ValidToPartyRelationshipTypesTable = "party_relationship_types"
	// ValidToPartyRelationshipTypesInverseTable is the table name for the PartyRelationshipType entity.
	// It exists in this package in order to avoid circular dependency with the "partyrelationshiptype" package.
	ValidToPartyRelationshipTypesInverseTable = "party_relationship_types"
	// ValidToPartyRelationshipTypesColumn is the table column denoting the valid_to_party_relationship_types relation/edge.
	ValidToPartyRelationshipTypesColumn = "role_type_valid_to_party_relationship_types"
	// PartyRolesTable is the table the holds the party_roles relation/edge.
	PartyRolesTable = "party_roles"
	// PartyRolesInverseTable is the table name for the PartyRole entity.
	// It exists in this package in order to avoid circular dependency with the "partyrole" package.
	PartyRolesInverseTable = "party_roles"
	// PartyRolesColumn is the table column denoting the party_roles relation/edge.
	PartyRolesColumn = "role_type_party_roles"
	// ChildRoleTypesTable is the table the holds the child_role_types relation/edge. The primary key declared below.
	ChildRoleTypesTable = "role_type_child_role_types"
	// WorkEffortPartyAssignmentsTable is the table the holds the work_effort_party_assignments relation/edge.
	WorkEffortPartyAssignmentsTable = "work_effort_party_assignments"
	// WorkEffortPartyAssignmentsInverseTable is the table name for the WorkEffortPartyAssignment entity.
	// It exists in this package in order to avoid circular dependency with the "workeffortpartyassignment" package.
	WorkEffortPartyAssignmentsInverseTable = "work_effort_party_assignments"
	// WorkEffortPartyAssignmentsColumn is the table column denoting the work_effort_party_assignments relation/edge.
	WorkEffortPartyAssignmentsColumn = "role_type_work_effort_party_assignments"
)

// Columns holds all SQL columns for roletype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldHasTable,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "role_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"role_type_children",
}

var (
	// ChildRoleTypesPrimaryKey and ChildRoleTypesColumn2 are the table columns denoting the
	// primary key for the child_role_types relation (M2M).
	ChildRoleTypesPrimaryKey = []string{"role_type_id", "child_role_type_id"}
)

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

// HasTable defines the type for the "has_table" enum field.
type HasTable string

// HasTable values.
const (
	HasTableYes     HasTable = "Yes"
	HasTableNo      HasTable = "No"
	HasTableUnknown HasTable = "Unknown"
)

func (ht HasTable) String() string {
	return string(ht)
}

// HasTableValidator is a validator for the "has_table" field enum values. It is called by the builders before save.
func HasTableValidator(ht HasTable) error {
	switch ht {
	case HasTableYes, HasTableNo, HasTableUnknown:
		return nil
	default:
		return fmt.Errorf("roletype: invalid enum value for has_table field: %q", ht)
	}
}