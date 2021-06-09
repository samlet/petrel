// Code generated by entc, DO NOT EDIT.

package orderitempriceinfo

import (
	"time"
)

const (
	// Label holds the string label denoting the orderitempriceinfo type in the database.
	Label = "order_item_price_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldOrderItemSeqID holds the string denoting the order_item_seq_id field in the database.
	FieldOrderItemSeqID = "order_item_seq_id"
	// FieldProductPriceRuleID holds the string denoting the product_price_rule_id field in the database.
	FieldProductPriceRuleID = "product_price_rule_id"
	// FieldProductPriceActionSeqID holds the string denoting the product_price_action_seq_id field in the database.
	FieldProductPriceActionSeqID = "product_price_action_seq_id"
	// FieldModifyAmount holds the string denoting the modify_amount field in the database.
	FieldModifyAmount = "modify_amount"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldRateCode holds the string denoting the rate_code field in the database.
	FieldRateCode = "rate_code"
	// EdgeOrderHeader holds the string denoting the order_header edge name in mutations.
	EdgeOrderHeader = "order_header"
	// EdgeOrderItem holds the string denoting the order_item edge name in mutations.
	EdgeOrderItem = "order_item"
	// Table holds the table name of the orderitempriceinfo in the database.
	Table = "order_item_price_infos"
	// OrderHeaderTable is the table the holds the order_header relation/edge.
	OrderHeaderTable = "order_item_price_infos"
	// OrderHeaderInverseTable is the table name for the OrderHeader entity.
	// It exists in this package in order to avoid circular dependency with the "orderheader" package.
	OrderHeaderInverseTable = "order_headers"
	// OrderHeaderColumn is the table column denoting the order_header relation/edge.
	OrderHeaderColumn = "order_header_order_item_price_infos"
	// OrderItemTable is the table the holds the order_item relation/edge.
	OrderItemTable = "order_item_price_infos"
	// OrderItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemInverseTable = "order_items"
	// OrderItemColumn is the table column denoting the order_item relation/edge.
	OrderItemColumn = "order_item_order_item_price_infos"
)

// Columns holds all SQL columns for orderitempriceinfo fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldOrderItemSeqID,
	FieldProductPriceRuleID,
	FieldProductPriceActionSeqID,
	FieldModifyAmount,
	FieldDescription,
	FieldRateCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_item_price_infos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_header_order_item_price_infos",
	"order_item_order_item_price_infos",
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
