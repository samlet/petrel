// Code generated by entc, DO NOT EDIT.

package productcategorytype

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the productcategorytype type in the database.
	Label = "product_category_type"
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
	// EdgeProductCategories holds the string denoting the product_categories edge name in mutations.
	EdgeProductCategories = "product_categories"
	// EdgeChildProductCategoryTypes holds the string denoting the child_product_category_types edge name in mutations.
	EdgeChildProductCategoryTypes = "child_product_category_types"
	// Table holds the table name of the productcategorytype in the database.
	Table = "product_category_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "product_category_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "product_category_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "product_category_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "product_category_type_children"
	// ProductCategoriesTable is the table the holds the product_categories relation/edge.
	ProductCategoriesTable = "product_categories"
	// ProductCategoriesInverseTable is the table name for the ProductCategory entity.
	// It exists in this package in order to avoid circular dependency with the "productcategory" package.
	ProductCategoriesInverseTable = "product_categories"
	// ProductCategoriesColumn is the table column denoting the product_categories relation/edge.
	ProductCategoriesColumn = "product_category_type_product_categories"
	// ChildProductCategoryTypesTable is the table the holds the child_product_category_types relation/edge. The primary key declared below.
	ChildProductCategoryTypesTable = "product_category_type_child_product_category_types"
)

// Columns holds all SQL columns for productcategorytype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldHasTable,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "product_category_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_category_type_children",
}

var (
	// ChildProductCategoryTypesPrimaryKey and ChildProductCategoryTypesColumn2 are the table columns denoting the
	// primary key for the child_product_category_types relation (M2M).
	ChildProductCategoryTypesPrimaryKey = []string{"product_category_type_id", "child_product_category_type_id"}
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
		return fmt.Errorf("productcategorytype: invalid enum value for has_table field: %q", ht)
	}
}
