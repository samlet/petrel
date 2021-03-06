// Code generated by entc, DO NOT EDIT.

package productreview

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the productreview type in the database.
	Label = "product_review"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldUserLoginID holds the string denoting the user_login_id field in the database.
	FieldUserLoginID = "user_login_id"
	// FieldPostedAnonymous holds the string denoting the posted_anonymous field in the database.
	FieldPostedAnonymous = "posted_anonymous"
	// FieldPostedDateTime holds the string denoting the posted_date_time field in the database.
	FieldPostedDateTime = "posted_date_time"
	// FieldProductRating holds the string denoting the product_rating field in the database.
	FieldProductRating = "product_rating"
	// FieldProductReview holds the string denoting the product_review field in the database.
	FieldProductReview = "product_review"
	// EdgeProductStore holds the string denoting the product_store edge name in mutations.
	EdgeProductStore = "product_store"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeStatusItem holds the string denoting the status_item edge name in mutations.
	EdgeStatusItem = "status_item"
	// Table holds the table name of the productreview in the database.
	Table = "product_reviews"
	// ProductStoreTable is the table the holds the product_store relation/edge.
	ProductStoreTable = "product_reviews"
	// ProductStoreInverseTable is the table name for the ProductStore entity.
	// It exists in this package in order to avoid circular dependency with the "productstore" package.
	ProductStoreInverseTable = "product_stores"
	// ProductStoreColumn is the table column denoting the product_store relation/edge.
	ProductStoreColumn = "product_store_product_reviews"
	// ProductTable is the table the holds the product relation/edge.
	ProductTable = "product_reviews"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_product_reviews"
	// StatusItemTable is the table the holds the status_item relation/edge.
	StatusItemTable = "product_reviews"
	// StatusItemInverseTable is the table name for the StatusItem entity.
	// It exists in this package in order to avoid circular dependency with the "statusitem" package.
	StatusItemInverseTable = "status_items"
	// StatusItemColumn is the table column denoting the status_item relation/edge.
	StatusItemColumn = "status_item_product_reviews"
)

// Columns holds all SQL columns for productreview fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldUserLoginID,
	FieldPostedAnonymous,
	FieldPostedDateTime,
	FieldProductRating,
	FieldProductReview,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "product_reviews"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_product_reviews",
	"product_store_product_reviews",
	"status_item_product_reviews",
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
	// DefaultPostedDateTime holds the default value on creation for the "posted_date_time" field.
	DefaultPostedDateTime func() time.Time
)

// PostedAnonymous defines the type for the "posted_anonymous" enum field.
type PostedAnonymous string

// PostedAnonymous values.
const (
	PostedAnonymousYes     PostedAnonymous = "Yes"
	PostedAnonymousNo      PostedAnonymous = "No"
	PostedAnonymousUnknown PostedAnonymous = "Unknown"
)

func (pa PostedAnonymous) String() string {
	return string(pa)
}

// PostedAnonymousValidator is a validator for the "posted_anonymous" field enum values. It is called by the builders before save.
func PostedAnonymousValidator(pa PostedAnonymous) error {
	switch pa {
	case PostedAnonymousYes, PostedAnonymousNo, PostedAnonymousUnknown:
		return nil
	default:
		return fmt.Errorf("productreview: invalid enum value for posted_anonymous field: %q", pa)
	}
}
