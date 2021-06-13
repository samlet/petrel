// Code generated by entc, DO NOT EDIT.

package securitygrouppermission

import (
	"time"
)

const (
	// Label holds the string label denoting the securitygrouppermission type in the database.
	Label = "security_group_permission"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldPermissionID holds the string denoting the permission_id field in the database.
	FieldPermissionID = "permission_id"
	// FieldFromDate holds the string denoting the from_date field in the database.
	FieldFromDate = "from_date"
	// FieldThruDate holds the string denoting the thru_date field in the database.
	FieldThruDate = "thru_date"
	// EdgeSecurityGroup holds the string denoting the security_group edge name in mutations.
	EdgeSecurityGroup = "security_group"
	// Table holds the table name of the securitygrouppermission in the database.
	Table = "security_group_permissions"
	// SecurityGroupTable is the table the holds the security_group relation/edge.
	SecurityGroupTable = "security_group_permissions"
	// SecurityGroupInverseTable is the table name for the SecurityGroup entity.
	// It exists in this package in order to avoid circular dependency with the "securitygroup" package.
	SecurityGroupInverseTable = "security_groups"
	// SecurityGroupColumn is the table column denoting the security_group relation/edge.
	SecurityGroupColumn = "security_group_security_group_permissions"
)

// Columns holds all SQL columns for securitygrouppermission fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldPermissionID,
	FieldFromDate,
	FieldThruDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "security_group_permissions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"security_group_security_group_permissions",
	"user_login_security_group_security_group_permissions",
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
	// PermissionIDValidator is a validator for the "permission_id" field. It is called by the builders before save.
	PermissionIDValidator func(string) error
	// DefaultFromDate holds the default value on creation for the "from_date" field.
	DefaultFromDate func() time.Time
	// DefaultThruDate holds the default value on creation for the "thru_date" field.
	DefaultThruDate func() time.Time
)
