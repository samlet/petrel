// Code generated by entc, DO NOT EDIT.

package itemissuance

import (
	"time"
)

const (
	// Label holds the string label denoting the itemissuance type in the database.
	Label = "item_issuance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldOrderItemSeqID holds the string denoting the order_item_seq_id field in the database.
	FieldOrderItemSeqID = "order_item_seq_id"
	// FieldShipGroupSeqID holds the string denoting the ship_group_seq_id field in the database.
	FieldShipGroupSeqID = "ship_group_seq_id"
	// FieldInventoryItemID holds the string denoting the inventory_item_id field in the database.
	FieldInventoryItemID = "inventory_item_id"
	// FieldShipmentItemSeqID holds the string denoting the shipment_item_seq_id field in the database.
	FieldShipmentItemSeqID = "shipment_item_seq_id"
	// FieldFixedAssetID holds the string denoting the fixed_asset_id field in the database.
	FieldFixedAssetID = "fixed_asset_id"
	// FieldMaintHistSeqID holds the string denoting the maint_hist_seq_id field in the database.
	FieldMaintHistSeqID = "maint_hist_seq_id"
	// FieldIssuedDateTime holds the string denoting the issued_date_time field in the database.
	FieldIssuedDateTime = "issued_date_time"
	// FieldIssuedByUserLoginID holds the string denoting the issued_by_user_login_id field in the database.
	FieldIssuedByUserLoginID = "issued_by_user_login_id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldCancelQuantity holds the string denoting the cancel_quantity field in the database.
	FieldCancelQuantity = "cancel_quantity"
	// EdgeShipment holds the string denoting the shipment edge name in mutations.
	EdgeShipment = "shipment"
	// EdgeShipmentItem holds the string denoting the shipment_item edge name in mutations.
	EdgeShipmentItem = "shipment_item"
	// EdgeOrderHeader holds the string denoting the order_header edge name in mutations.
	EdgeOrderHeader = "order_header"
	// EdgeOrderItem holds the string denoting the order_item edge name in mutations.
	EdgeOrderItem = "order_item"
	// Table holds the table name of the itemissuance in the database.
	Table = "item_issuances"
	// ShipmentTable is the table the holds the shipment relation/edge.
	ShipmentTable = "item_issuances"
	// ShipmentInverseTable is the table name for the Shipment entity.
	// It exists in this package in order to avoid circular dependency with the "shipment" package.
	ShipmentInverseTable = "shipments"
	// ShipmentColumn is the table column denoting the shipment relation/edge.
	ShipmentColumn = "shipment_item_issuances"
	// ShipmentItemTable is the table the holds the shipment_item relation/edge.
	ShipmentItemTable = "item_issuances"
	// ShipmentItemInverseTable is the table name for the ShipmentItem entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentitem" package.
	ShipmentItemInverseTable = "shipment_items"
	// ShipmentItemColumn is the table column denoting the shipment_item relation/edge.
	ShipmentItemColumn = "shipment_item_item_issuances"
	// OrderHeaderTable is the table the holds the order_header relation/edge.
	OrderHeaderTable = "item_issuances"
	// OrderHeaderInverseTable is the table name for the OrderHeader entity.
	// It exists in this package in order to avoid circular dependency with the "orderheader" package.
	OrderHeaderInverseTable = "order_headers"
	// OrderHeaderColumn is the table column denoting the order_header relation/edge.
	OrderHeaderColumn = "order_header_item_issuances"
	// OrderItemTable is the table the holds the order_item relation/edge.
	OrderItemTable = "item_issuances"
	// OrderItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemInverseTable = "order_items"
	// OrderItemColumn is the table column denoting the order_item relation/edge.
	OrderItemColumn = "order_item_item_issuances"
)

// Columns holds all SQL columns for itemissuance fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldOrderItemSeqID,
	FieldShipGroupSeqID,
	FieldInventoryItemID,
	FieldShipmentItemSeqID,
	FieldFixedAssetID,
	FieldMaintHistSeqID,
	FieldIssuedDateTime,
	FieldIssuedByUserLoginID,
	FieldQuantity,
	FieldCancelQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "item_issuances"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_header_item_issuances",
	"order_item_item_issuances",
	"shipment_item_issuances",
	"shipment_item_item_issuances",
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
	// DefaultIssuedDateTime holds the default value on creation for the "issued_date_time" field.
	DefaultIssuedDateTime func() time.Time
)
