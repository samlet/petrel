// Code generated by entc, DO NOT EDIT.

package partyrole

const (
	// Label holds the string label denoting the partyrole type in the database.
	Label = "party_role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRoleTypeID holds the string denoting the role_type_id field in the database.
	FieldRoleTypeID = "role_type_id"
	// EdgeParty holds the string denoting the party edge name in mutations.
	EdgeParty = "party"
	// EdgeFixedAssets holds the string denoting the fixed_assets edge name in mutations.
	EdgeFixedAssets = "fixed_assets"
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
	// FixedAssetsTable is the table the holds the fixed_assets relation/edge.
	FixedAssetsTable = "fixed_assets"
	// FixedAssetsInverseTable is the table name for the FixedAsset entity.
	// It exists in this package in order to avoid circular dependency with the "fixedasset" package.
	FixedAssetsInverseTable = "fixed_assets"
	// FixedAssetsColumn is the table column denoting the fixed_assets relation/edge.
	FixedAssetsColumn = "party_role_fixed_assets"
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
	FieldRoleTypeID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "party_roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"party_party_roles",
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
