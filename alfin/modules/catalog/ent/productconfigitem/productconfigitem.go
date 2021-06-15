// Code generated by entc, DO NOT EDIT.

package productconfigitem

import (
	"time"
)

const (
	// Label holds the string label denoting the productconfigitem type in the database.
	Label = "product_config_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldConfigItemTypeID holds the string denoting the config_item_type_id field in the database.
	FieldConfigItemTypeID = "config_item_type_id"
	// FieldConfigItemName holds the string denoting the config_item_name field in the database.
	FieldConfigItemName = "config_item_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLongDescription holds the string denoting the long_description field in the database.
	FieldLongDescription = "long_description"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// Table holds the table name of the productconfigitem in the database.
	Table = "product_config_items"
)

// Columns holds all SQL columns for productconfigitem fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldConfigItemTypeID,
	FieldConfigItemName,
	FieldDescription,
	FieldLongDescription,
	FieldImageURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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