// Code generated by entc, DO NOT EDIT.

package partycontenttype

import (
	"time"
)

const (
	// Label holds the string label denoting the partycontenttype type in the database.
	Label = "party_content_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeChildPartyContentTypes holds the string denoting the child_party_content_types edge name in mutations.
	EdgeChildPartyContentTypes = "child_party_content_types"
	// Table holds the table name of the partycontenttype in the database.
	Table = "party_content_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "party_content_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "party_content_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "party_content_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "party_content_type_children"
	// ChildPartyContentTypesTable is the table the holds the child_party_content_types relation/edge. The primary key declared below.
	ChildPartyContentTypesTable = "party_content_type_child_party_content_types"
)

// Columns holds all SQL columns for partycontenttype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "party_content_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"party_content_type_children",
}

var (
	// ChildPartyContentTypesPrimaryKey and ChildPartyContentTypesColumn2 are the table columns denoting the
	// primary key for the child_party_content_types relation (M2M).
	ChildPartyContentTypesPrimaryKey = []string{"party_content_type_id", "child_party_content_type_id"}
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
