// Code generated by entc, DO NOT EDIT.

package orderitemshipgroupassoc

import (
	"time"
)

const (
	// Label holds the string label denoting the orderitemshipgroupassoc type in the database.
	Label = "order_item_ship_group_assoc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldOrderItemSeqID holds the string denoting the order_item_seq_id field in the database.
	FieldOrderItemSeqID = "order_item_seq_id"
	// FieldShipGroupSeqID holds the string denoting the ship_group_seq_id field in the database.
	FieldShipGroupSeqID = "ship_group_seq_id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldCancelQuantity holds the string denoting the cancel_quantity field in the database.
	FieldCancelQuantity = "cancel_quantity"
	// EdgeOrderHeader holds the string denoting the order_header edge name in mutations.
	EdgeOrderHeader = "order_header"
	// EdgeOrderItem holds the string denoting the order_item edge name in mutations.
	EdgeOrderItem = "order_item"
	// EdgeOrderItemShipGroup holds the string denoting the order_item_ship_group edge name in mutations.
	EdgeOrderItemShipGroup = "order_item_ship_group"
	// EdgeOrderAdjustments holds the string denoting the order_adjustments edge name in mutations.
	EdgeOrderAdjustments = "order_adjustments"
	// EdgeOrderItemShipGrpInvRes holds the string denoting the order_item_ship_grp_inv_res edge name in mutations.
	EdgeOrderItemShipGrpInvRes = "order_item_ship_grp_inv_res"
	// Table holds the table name of the orderitemshipgroupassoc in the database.
	Table = "order_item_ship_group_assocs"
	// OrderHeaderTable is the table the holds the order_header relation/edge.
	OrderHeaderTable = "order_item_ship_group_assocs"
	// OrderHeaderInverseTable is the table name for the OrderHeader entity.
	// It exists in this package in order to avoid circular dependency with the "orderheader" package.
	OrderHeaderInverseTable = "order_headers"
	// OrderHeaderColumn is the table column denoting the order_header relation/edge.
	OrderHeaderColumn = "order_header_order_item_ship_group_assocs"
	// OrderItemTable is the table the holds the order_item relation/edge.
	OrderItemTable = "order_item_ship_group_assocs"
	// OrderItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemInverseTable = "order_items"
	// OrderItemColumn is the table column denoting the order_item relation/edge.
	OrderItemColumn = "order_item_order_item_ship_group_assocs"
	// OrderItemShipGroupTable is the table the holds the order_item_ship_group relation/edge.
	OrderItemShipGroupTable = "order_item_ship_group_assocs"
	// OrderItemShipGroupInverseTable is the table name for the OrderItemShipGroup entity.
	// It exists in this package in order to avoid circular dependency with the "orderitemshipgroup" package.
	OrderItemShipGroupInverseTable = "order_item_ship_groups"
	// OrderItemShipGroupColumn is the table column denoting the order_item_ship_group relation/edge.
	OrderItemShipGroupColumn = "order_item_ship_group_order_item_ship_group_assocs"
	// OrderAdjustmentsTable is the table the holds the order_adjustments relation/edge.
	OrderAdjustmentsTable = "order_adjustments"
	// OrderAdjustmentsInverseTable is the table name for the OrderAdjustment entity.
	// It exists in this package in order to avoid circular dependency with the "orderadjustment" package.
	OrderAdjustmentsInverseTable = "order_adjustments"
	// OrderAdjustmentsColumn is the table column denoting the order_adjustments relation/edge.
	OrderAdjustmentsColumn = "order_item_ship_group_assoc_order_adjustments"
	// OrderItemShipGrpInvResTable is the table the holds the order_item_ship_grp_inv_res relation/edge.
	OrderItemShipGrpInvResTable = "order_item_ship_grp_inv_res"
	// OrderItemShipGrpInvResInverseTable is the table name for the OrderItemShipGrpInvRes entity.
	// It exists in this package in order to avoid circular dependency with the "orderitemshipgrpinvres" package.
	OrderItemShipGrpInvResInverseTable = "order_item_ship_grp_inv_res"
	// OrderItemShipGrpInvResColumn is the table column denoting the order_item_ship_grp_inv_res relation/edge.
	OrderItemShipGrpInvResColumn = "order_item_ship_group_assoc_order_item_ship_grp_inv_res"
)

// Columns holds all SQL columns for orderitemshipgroupassoc fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldOrderItemSeqID,
	FieldShipGroupSeqID,
	FieldQuantity,
	FieldCancelQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_item_ship_group_assocs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_header_order_item_ship_group_assocs",
	"order_item_order_item_ship_group_assocs",
	"order_item_ship_group_order_item_ship_group_assocs",
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