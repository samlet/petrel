// Code generated by entc, DO NOT EDIT.

package asset

import (
	"time"
)

const (
	// Label holds the string label denoting the asset type in the database.
	Label = "asset"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldRegisteredAt holds the string denoting the registered_at field in the database.
	FieldRegisteredAt = "registered_at"
	// EdgePkg holds the string denoting the pkg edge name in mutations.
	EdgePkg = "pkg"
	// Table holds the table name of the asset in the database.
	Table = "assets"
	// PkgTable is the table the holds the pkg relation/edge.
	PkgTable = "assets"
	// PkgInverseTable is the table name for the WorkloadPkg entity.
	// It exists in this package in order to avoid circular dependency with the "workloadpkg" package.
	PkgInverseTable = "workload_pkgs"
	// PkgColumn is the table column denoting the pkg relation/edge.
	PkgColumn = "workload_pkg_assets"
)

// Columns holds all SQL columns for asset fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldModel,
	FieldRegisteredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "assets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workload_pkg_assets",
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
