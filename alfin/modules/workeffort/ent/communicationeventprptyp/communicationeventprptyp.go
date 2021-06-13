// Code generated by entc, DO NOT EDIT.

package communicationeventprptyp

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the communicationeventprptyp type in the database.
	Label = "communication_event_prp_typ"
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
	// EdgeChildCommunicationEventPrpTyps holds the string denoting the child_communication_event_prp_typs edge name in mutations.
	EdgeChildCommunicationEventPrpTyps = "child_communication_event_prp_typs"
	// Table holds the table name of the communicationeventprptyp in the database.
	Table = "communication_event_prp_typs"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "communication_event_prp_typs"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "communication_event_prp_typ_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "communication_event_prp_typs"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "communication_event_prp_typ_children"
	// ChildCommunicationEventPrpTypsTable is the table the holds the child_communication_event_prp_typs relation/edge. The primary key declared below.
	ChildCommunicationEventPrpTypsTable = "communication_event_prp_typ_child_communication_event_prp_typs"
)

// Columns holds all SQL columns for communicationeventprptyp fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldHasTable,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "communication_event_prp_typs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"communication_event_prp_typ_children",
}

var (
	// ChildCommunicationEventPrpTypsPrimaryKey and ChildCommunicationEventPrpTypsColumn2 are the table columns denoting the
	// primary key for the child_communication_event_prp_typs relation (M2M).
	ChildCommunicationEventPrpTypsPrimaryKey = []string{"communication_event_prp_typ_id", "child_communication_event_prp_typ_id"}
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
		return fmt.Errorf("communicationeventprptyp: invalid enum value for has_table field: %q", ht)
	}
}
