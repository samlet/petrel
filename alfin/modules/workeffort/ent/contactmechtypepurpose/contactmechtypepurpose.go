// Code generated by entc, DO NOT EDIT.

package contactmechtypepurpose

import (
	"time"
)

const (
	// Label holds the string label denoting the contactmechtypepurpose type in the database.
	Label = "contact_mech_type_purpose"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// EdgeContactMechType holds the string denoting the contact_mech_type edge name in mutations.
	EdgeContactMechType = "contact_mech_type"
	// EdgeContactMechPurposeType holds the string denoting the contact_mech_purpose_type edge name in mutations.
	EdgeContactMechPurposeType = "contact_mech_purpose_type"
	// Table holds the table name of the contactmechtypepurpose in the database.
	Table = "contact_mech_type_purposes"
	// ContactMechTypeTable is the table the holds the contact_mech_type relation/edge.
	ContactMechTypeTable = "contact_mech_type_purposes"
	// ContactMechTypeInverseTable is the table name for the ContactMechType entity.
	// It exists in this package in order to avoid circular dependency with the "contactmechtype" package.
	ContactMechTypeInverseTable = "contact_mech_types"
	// ContactMechTypeColumn is the table column denoting the contact_mech_type relation/edge.
	ContactMechTypeColumn = "contact_mech_type_contact_mech_type_purposes"
	// ContactMechPurposeTypeTable is the table the holds the contact_mech_purpose_type relation/edge.
	ContactMechPurposeTypeTable = "contact_mech_type_purposes"
	// ContactMechPurposeTypeInverseTable is the table name for the ContactMechPurposeType entity.
	// It exists in this package in order to avoid circular dependency with the "contactmechpurposetype" package.
	ContactMechPurposeTypeInverseTable = "contact_mech_purpose_types"
	// ContactMechPurposeTypeColumn is the table column denoting the contact_mech_purpose_type relation/edge.
	ContactMechPurposeTypeColumn = "contact_mech_purpose_type_contact_mech_type_purposes"
)

// Columns holds all SQL columns for contactmechtypepurpose fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "contact_mech_type_purposes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"contact_mech_purpose_type_contact_mech_type_purposes",
	"contact_mech_type_contact_mech_type_purposes",
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
