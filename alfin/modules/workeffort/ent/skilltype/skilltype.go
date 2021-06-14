// Code generated by entc, DO NOT EDIT.

package skilltype

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the skilltype type in the database.
	Label = "skill_type"
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
	// EdgeChildSkillTypes holds the string denoting the child_skill_types edge name in mutations.
	EdgeChildSkillTypes = "child_skill_types"
	// EdgeWorkEffortSkillStandards holds the string denoting the work_effort_skill_standards edge name in mutations.
	EdgeWorkEffortSkillStandards = "work_effort_skill_standards"
	// Table holds the table name of the skilltype in the database.
	Table = "skill_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "skill_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "skill_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "skill_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "skill_type_children"
	// ChildSkillTypesTable is the table the holds the child_skill_types relation/edge. The primary key declared below.
	ChildSkillTypesTable = "skill_type_child_skill_types"
	// WorkEffortSkillStandardsTable is the table the holds the work_effort_skill_standards relation/edge.
	WorkEffortSkillStandardsTable = "work_effort_skill_standards"
	// WorkEffortSkillStandardsInverseTable is the table name for the WorkEffortSkillStandard entity.
	// It exists in this package in order to avoid circular dependency with the "workeffortskillstandard" package.
	WorkEffortSkillStandardsInverseTable = "work_effort_skill_standards"
	// WorkEffortSkillStandardsColumn is the table column denoting the work_effort_skill_standards relation/edge.
	WorkEffortSkillStandardsColumn = "skill_type_work_effort_skill_standards"
)

// Columns holds all SQL columns for skilltype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldHasTable,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "skill_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"skill_type_children",
}

var (
	// ChildSkillTypesPrimaryKey and ChildSkillTypesColumn2 are the table columns denoting the
	// primary key for the child_skill_types relation (M2M).
	ChildSkillTypesPrimaryKey = []string{"skill_type_id", "child_skill_type_id"}
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
		return fmt.Errorf("skilltype: invalid enum value for has_table field: %q", ht)
	}
}