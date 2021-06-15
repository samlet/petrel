// Code generated by entc, DO NOT EDIT.

package prodcatalogcategorytype

import (
	"time"
)

const (
	// Label holds the string label denoting the prodcatalogcategorytype type in the database.
	Label = "prod_catalog_category_type"
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
	// EdgeChildProdCatalogCategoryTypes holds the string denoting the child_prod_catalog_category_types edge name in mutations.
	EdgeChildProdCatalogCategoryTypes = "child_prod_catalog_category_types"
	// Table holds the table name of the prodcatalogcategorytype in the database.
	Table = "prod_catalog_category_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "prod_catalog_category_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "prod_catalog_category_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "prod_catalog_category_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "prod_catalog_category_type_children"
	// ChildProdCatalogCategoryTypesTable is the table the holds the child_prod_catalog_category_types relation/edge. The primary key declared below.
	ChildProdCatalogCategoryTypesTable = "prod_catalog_category_type_child_prod_catalog_category_types"
)

// Columns holds all SQL columns for prodcatalogcategorytype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prod_catalog_category_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"prod_catalog_category_type_children",
}

var (
	// ChildProdCatalogCategoryTypesPrimaryKey and ChildProdCatalogCategoryTypesColumn2 are the table columns denoting the
	// primary key for the child_prod_catalog_category_types relation (M2M).
	ChildProdCatalogCategoryTypesPrimaryKey = []string{"prod_catalog_category_type_id", "child_prod_catalog_category_type_id"}
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
