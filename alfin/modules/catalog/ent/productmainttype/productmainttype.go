// Code generated by entc, DO NOT EDIT.

package productmainttype

import (
	"time"
)

const (
	// Label holds the string label denoting the productmainttype type in the database.
	Label = "product_maint_type"
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
	// EdgeChildProductMaintTypes holds the string denoting the child_product_maint_types edge name in mutations.
	EdgeChildProductMaintTypes = "child_product_maint_types"
	// Table holds the table name of the productmainttype in the database.
	Table = "product_maint_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "product_maint_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "product_maint_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "product_maint_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "product_maint_type_children"
	// ChildProductMaintTypesTable is the table the holds the child_product_maint_types relation/edge. The primary key declared below.
	ChildProductMaintTypesTable = "product_maint_type_child_product_maint_types"
)

// Columns holds all SQL columns for productmainttype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "product_maint_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_maint_type_children",
}

var (
	// ChildProductMaintTypesPrimaryKey and ChildProductMaintTypesColumn2 are the table columns denoting the
	// primary key for the child_product_maint_types relation (M2M).
	ChildProductMaintTypesPrimaryKey = []string{"product_maint_type_id", "child_product_maint_type_id"}
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
