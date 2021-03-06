// Code generated by entc, DO NOT EDIT.

package custommethod

import (
	"time"
)

const (
	// Label holds the string label denoting the custommethod type in the database.
	Label = "custom_method"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldCustomMethodName holds the string denoting the custom_method_name field in the database.
	FieldCustomMethodName = "custom_method_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCustomMethodType holds the string denoting the custom_method_type edge name in mutations.
	EdgeCustomMethodType = "custom_method_type"
	// EdgeProductAssocs holds the string denoting the product_assocs edge name in mutations.
	EdgeProductAssocs = "product_assocs"
	// EdgeProductPrices holds the string denoting the product_prices edge name in mutations.
	EdgeProductPrices = "product_prices"
	// Table holds the table name of the custommethod in the database.
	Table = "custom_methods"
	// CustomMethodTypeTable is the table the holds the custom_method_type relation/edge.
	CustomMethodTypeTable = "custom_methods"
	// CustomMethodTypeInverseTable is the table name for the CustomMethodType entity.
	// It exists in this package in order to avoid circular dependency with the "custommethodtype" package.
	CustomMethodTypeInverseTable = "custom_method_types"
	// CustomMethodTypeColumn is the table column denoting the custom_method_type relation/edge.
	CustomMethodTypeColumn = "custom_method_type_custom_methods"
	// ProductAssocsTable is the table the holds the product_assocs relation/edge.
	ProductAssocsTable = "product_assocs"
	// ProductAssocsInverseTable is the table name for the ProductAssoc entity.
	// It exists in this package in order to avoid circular dependency with the "productassoc" package.
	ProductAssocsInverseTable = "product_assocs"
	// ProductAssocsColumn is the table column denoting the product_assocs relation/edge.
	ProductAssocsColumn = "custom_method_product_assocs"
	// ProductPricesTable is the table the holds the product_prices relation/edge.
	ProductPricesTable = "product_prices"
	// ProductPricesInverseTable is the table name for the ProductPrice entity.
	// It exists in this package in order to avoid circular dependency with the "productprice" package.
	ProductPricesInverseTable = "product_prices"
	// ProductPricesColumn is the table column denoting the product_prices relation/edge.
	ProductPricesColumn = "custom_method_product_prices"
)

// Columns holds all SQL columns for custommethod fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldCustomMethodName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "custom_methods"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"custom_method_type_custom_methods",
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
