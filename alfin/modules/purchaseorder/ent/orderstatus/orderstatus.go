// Code generated by entc, DO NOT EDIT.

package orderstatus

import (
	"time"
)

const (
	// Label holds the string label denoting the orderstatus type in the database.
	Label = "order_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldOrderItemSeqID holds the string denoting the order_item_seq_id field in the database.
	FieldOrderItemSeqID = "order_item_seq_id"
	// FieldOrderPaymentPreferenceID holds the string denoting the order_payment_preference_id field in the database.
	FieldOrderPaymentPreferenceID = "order_payment_preference_id"
	// FieldStatusDatetime holds the string denoting the status_datetime field in the database.
	FieldStatusDatetime = "status_datetime"
	// FieldStatusUserLogin holds the string denoting the status_user_login field in the database.
	FieldStatusUserLogin = "status_user_login"
	// FieldChangeReason holds the string denoting the change_reason field in the database.
	FieldChangeReason = "change_reason"
	// EdgeOrderHeader holds the string denoting the order_header edge name in mutations.
	EdgeOrderHeader = "order_header"
	// EdgeOrderItem holds the string denoting the order_item edge name in mutations.
	EdgeOrderItem = "order_item"
	// Table holds the table name of the orderstatus in the database.
	Table = "order_status"
	// OrderHeaderTable is the table the holds the order_header relation/edge.
	OrderHeaderTable = "order_status"
	// OrderHeaderInverseTable is the table name for the OrderHeader entity.
	// It exists in this package in order to avoid circular dependency with the "orderheader" package.
	OrderHeaderInverseTable = "order_headers"
	// OrderHeaderColumn is the table column denoting the order_header relation/edge.
	OrderHeaderColumn = "order_header_order_statuses"
	// OrderItemTable is the table the holds the order_item relation/edge.
	OrderItemTable = "order_status"
	// OrderItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemInverseTable = "order_items"
	// OrderItemColumn is the table column denoting the order_item relation/edge.
	OrderItemColumn = "order_item_order_statuses"
)

// Columns holds all SQL columns for orderstatus fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStatusID,
	FieldOrderItemSeqID,
	FieldOrderPaymentPreferenceID,
	FieldStatusDatetime,
	FieldStatusUserLogin,
	FieldChangeReason,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_status"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_header_order_statuses",
	"order_item_order_statuses",
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
	// DefaultStatusDatetime holds the default value on creation for the "status_datetime" field.
	DefaultStatusDatetime func() time.Time
)
