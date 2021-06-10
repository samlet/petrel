// Code generated by entc, DO NOT EDIT.

package orderheader

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the orderheader type in the database.
	Label = "order_header"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldOrderTypeID holds the string denoting the order_type_id field in the database.
	FieldOrderTypeID = "order_type_id"
	// FieldOrderName holds the string denoting the order_name field in the database.
	FieldOrderName = "order_name"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// FieldSalesChannelEnumID holds the string denoting the sales_channel_enum_id field in the database.
	FieldSalesChannelEnumID = "sales_channel_enum_id"
	// FieldOrderDate holds the string denoting the order_date field in the database.
	FieldOrderDate = "order_date"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldEntryDate holds the string denoting the entry_date field in the database.
	FieldEntryDate = "entry_date"
	// FieldPickSheetPrintedDate holds the string denoting the pick_sheet_printed_date field in the database.
	FieldPickSheetPrintedDate = "pick_sheet_printed_date"
	// FieldVisitID holds the string denoting the visit_id field in the database.
	FieldVisitID = "visit_id"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldFirstAttemptOrderID holds the string denoting the first_attempt_order_id field in the database.
	FieldFirstAttemptOrderID = "first_attempt_order_id"
	// FieldCurrencyUom holds the string denoting the currency_uom field in the database.
	FieldCurrencyUom = "currency_uom"
	// FieldSyncStatusID holds the string denoting the sync_status_id field in the database.
	FieldSyncStatusID = "sync_status_id"
	// FieldBillingAccountID holds the string denoting the billing_account_id field in the database.
	FieldBillingAccountID = "billing_account_id"
	// FieldOriginFacilityID holds the string denoting the origin_facility_id field in the database.
	FieldOriginFacilityID = "origin_facility_id"
	// FieldWebSiteID holds the string denoting the web_site_id field in the database.
	FieldWebSiteID = "web_site_id"
	// FieldProductStoreID holds the string denoting the product_store_id field in the database.
	FieldProductStoreID = "product_store_id"
	// FieldAgreementID holds the string denoting the agreement_id field in the database.
	FieldAgreementID = "agreement_id"
	// FieldTerminalID holds the string denoting the terminal_id field in the database.
	FieldTerminalID = "terminal_id"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldAutoOrderShoppingListID holds the string denoting the auto_order_shopping_list_id field in the database.
	FieldAutoOrderShoppingListID = "auto_order_shopping_list_id"
	// FieldNeedsInventoryIssuance holds the string denoting the needs_inventory_issuance field in the database.
	FieldNeedsInventoryIssuance = "needs_inventory_issuance"
	// FieldIsRushOrder holds the string denoting the is_rush_order field in the database.
	FieldIsRushOrder = "is_rush_order"
	// FieldInternalCode holds the string denoting the internal_code field in the database.
	FieldInternalCode = "internal_code"
	// FieldRemainingSubTotal holds the string denoting the remaining_sub_total field in the database.
	FieldRemainingSubTotal = "remaining_sub_total"
	// FieldGrandTotal holds the string denoting the grand_total field in the database.
	FieldGrandTotal = "grand_total"
	// FieldIsViewed holds the string denoting the is_viewed field in the database.
	FieldIsViewed = "is_viewed"
	// FieldInvoicePerShipment holds the string denoting the invoice_per_shipment field in the database.
	FieldInvoicePerShipment = "invoice_per_shipment"
	// EdgeItemIssuances holds the string denoting the item_issuances edge name in mutations.
	EdgeItemIssuances = "item_issuances"
	// EdgeOrderContactMeches holds the string denoting the order_contact_meches edge name in mutations.
	EdgeOrderContactMeches = "order_contact_meches"
	// EdgeOrderItems holds the string denoting the order_items edge name in mutations.
	EdgeOrderItems = "order_items"
	// EdgeOrderItemPriceInfos holds the string denoting the order_item_price_infos edge name in mutations.
	EdgeOrderItemPriceInfos = "order_item_price_infos"
	// EdgeOrderItemShipGroups holds the string denoting the order_item_ship_groups edge name in mutations.
	EdgeOrderItemShipGroups = "order_item_ship_groups"
	// EdgeOrderItemShipGroupAssocs holds the string denoting the order_item_ship_group_assocs edge name in mutations.
	EdgeOrderItemShipGroupAssocs = "order_item_ship_group_assocs"
	// EdgeOrderRoles holds the string denoting the order_roles edge name in mutations.
	EdgeOrderRoles = "order_roles"
	// EdgeOrderStatuses holds the string denoting the order_statuses edge name in mutations.
	EdgeOrderStatuses = "order_statuses"
	// EdgePrimaryShipments holds the string denoting the primary_shipments edge name in mutations.
	EdgePrimaryShipments = "primary_shipments"
	// Table holds the table name of the orderheader in the database.
	Table = "order_headers"
	// ItemIssuancesTable is the table the holds the item_issuances relation/edge.
	ItemIssuancesTable = "item_issuances"
	// ItemIssuancesInverseTable is the table name for the ItemIssuance entity.
	// It exists in this package in order to avoid circular dependency with the "itemissuance" package.
	ItemIssuancesInverseTable = "item_issuances"
	// ItemIssuancesColumn is the table column denoting the item_issuances relation/edge.
	ItemIssuancesColumn = "order_header_item_issuances"
	// OrderContactMechesTable is the table the holds the order_contact_meches relation/edge.
	OrderContactMechesTable = "order_contact_meches"
	// OrderContactMechesInverseTable is the table name for the OrderContactMech entity.
	// It exists in this package in order to avoid circular dependency with the "ordercontactmech" package.
	OrderContactMechesInverseTable = "order_contact_meches"
	// OrderContactMechesColumn is the table column denoting the order_contact_meches relation/edge.
	OrderContactMechesColumn = "order_header_order_contact_meches"
	// OrderItemsTable is the table the holds the order_items relation/edge.
	OrderItemsTable = "order_items"
	// OrderItemsInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemsInverseTable = "order_items"
	// OrderItemsColumn is the table column denoting the order_items relation/edge.
	OrderItemsColumn = "order_header_order_items"
	// OrderItemPriceInfosTable is the table the holds the order_item_price_infos relation/edge.
	OrderItemPriceInfosTable = "order_item_price_infos"
	// OrderItemPriceInfosInverseTable is the table name for the OrderItemPriceInfo entity.
	// It exists in this package in order to avoid circular dependency with the "orderitempriceinfo" package.
	OrderItemPriceInfosInverseTable = "order_item_price_infos"
	// OrderItemPriceInfosColumn is the table column denoting the order_item_price_infos relation/edge.
	OrderItemPriceInfosColumn = "order_header_order_item_price_infos"
	// OrderItemShipGroupsTable is the table the holds the order_item_ship_groups relation/edge.
	OrderItemShipGroupsTable = "order_item_ship_groups"
	// OrderItemShipGroupsInverseTable is the table name for the OrderItemShipGroup entity.
	// It exists in this package in order to avoid circular dependency with the "orderitemshipgroup" package.
	OrderItemShipGroupsInverseTable = "order_item_ship_groups"
	// OrderItemShipGroupsColumn is the table column denoting the order_item_ship_groups relation/edge.
	OrderItemShipGroupsColumn = "order_header_order_item_ship_groups"
	// OrderItemShipGroupAssocsTable is the table the holds the order_item_ship_group_assocs relation/edge.
	OrderItemShipGroupAssocsTable = "order_item_ship_group_assocs"
	// OrderItemShipGroupAssocsInverseTable is the table name for the OrderItemShipGroupAssoc entity.
	// It exists in this package in order to avoid circular dependency with the "orderitemshipgroupassoc" package.
	OrderItemShipGroupAssocsInverseTable = "order_item_ship_group_assocs"
	// OrderItemShipGroupAssocsColumn is the table column denoting the order_item_ship_group_assocs relation/edge.
	OrderItemShipGroupAssocsColumn = "order_header_order_item_ship_group_assocs"
	// OrderRolesTable is the table the holds the order_roles relation/edge.
	OrderRolesTable = "order_roles"
	// OrderRolesInverseTable is the table name for the OrderRole entity.
	// It exists in this package in order to avoid circular dependency with the "orderrole" package.
	OrderRolesInverseTable = "order_roles"
	// OrderRolesColumn is the table column denoting the order_roles relation/edge.
	OrderRolesColumn = "order_header_order_roles"
	// OrderStatusesTable is the table the holds the order_statuses relation/edge.
	OrderStatusesTable = "order_status"
	// OrderStatusesInverseTable is the table name for the OrderStatus entity.
	// It exists in this package in order to avoid circular dependency with the "orderstatus" package.
	OrderStatusesInverseTable = "order_status"
	// OrderStatusesColumn is the table column denoting the order_statuses relation/edge.
	OrderStatusesColumn = "order_header_order_statuses"
	// PrimaryShipmentsTable is the table the holds the primary_shipments relation/edge.
	PrimaryShipmentsTable = "shipments"
	// PrimaryShipmentsInverseTable is the table name for the Shipment entity.
	// It exists in this package in order to avoid circular dependency with the "shipment" package.
	PrimaryShipmentsInverseTable = "shipments"
	// PrimaryShipmentsColumn is the table column denoting the primary_shipments relation/edge.
	PrimaryShipmentsColumn = "order_header_primary_shipments"
)

// Columns holds all SQL columns for orderheader fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldOrderTypeID,
	FieldOrderName,
	FieldExternalID,
	FieldSalesChannelEnumID,
	FieldOrderDate,
	FieldPriority,
	FieldEntryDate,
	FieldPickSheetPrintedDate,
	FieldVisitID,
	FieldStatusID,
	FieldCreatedBy,
	FieldFirstAttemptOrderID,
	FieldCurrencyUom,
	FieldSyncStatusID,
	FieldBillingAccountID,
	FieldOriginFacilityID,
	FieldWebSiteID,
	FieldProductStoreID,
	FieldAgreementID,
	FieldTerminalID,
	FieldTransactionID,
	FieldAutoOrderShoppingListID,
	FieldNeedsInventoryIssuance,
	FieldIsRushOrder,
	FieldInternalCode,
	FieldRemainingSubTotal,
	FieldGrandTotal,
	FieldIsViewed,
	FieldInvoicePerShipment,
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
	// DefaultOrderDate holds the default value on creation for the "order_date" field.
	DefaultOrderDate func() time.Time
	// DefaultEntryDate holds the default value on creation for the "entry_date" field.
	DefaultEntryDate func() time.Time
	// DefaultPickSheetPrintedDate holds the default value on creation for the "pick_sheet_printed_date" field.
	DefaultPickSheetPrintedDate func() time.Time
	// TerminalIDValidator is a validator for the "terminal_id" field. It is called by the builders before save.
	TerminalIDValidator func(string) error
	// TransactionIDValidator is a validator for the "transaction_id" field. It is called by the builders before save.
	TransactionIDValidator func(string) error
	// InternalCodeValidator is a validator for the "internal_code" field. It is called by the builders before save.
	InternalCodeValidator func(string) error
)

// Priority defines the type for the "priority" enum field.
type Priority string

// Priority values.
const (
	PriorityY Priority = "Y"
	PriorityN Priority = "N"
	Priority  Priority = "-"
)

func (pr Priority) String() string {
	return string(pr)
}

// PriorityValidator is a validator for the "priority" field enum values. It is called by the builders before save.
func PriorityValidator(pr Priority) error {
	switch pr {
	case PriorityY, PriorityN, Priority:
		return nil
	default:
		return fmt.Errorf("orderheader: invalid enum value for priority field: %q", pr)
	}
}

// NeedsInventoryIssuance defines the type for the "needs_inventory_issuance" enum field.
type NeedsInventoryIssuance string

// NeedsInventoryIssuance values.
const (
	NeedsInventoryIssuanceY NeedsInventoryIssuance = "Y"
	NeedsInventoryIssuanceN NeedsInventoryIssuance = "N"
	NeedsInventoryIssuance  NeedsInventoryIssuance = "-"
)

func (nii NeedsInventoryIssuance) String() string {
	return string(nii)
}

// NeedsInventoryIssuanceValidator is a validator for the "needs_inventory_issuance" field enum values. It is called by the builders before save.
func NeedsInventoryIssuanceValidator(nii NeedsInventoryIssuance) error {
	switch nii {
	case NeedsInventoryIssuanceY, NeedsInventoryIssuanceN, NeedsInventoryIssuance:
		return nil
	default:
		return fmt.Errorf("orderheader: invalid enum value for needs_inventory_issuance field: %q", nii)
	}
}

// IsRushOrder defines the type for the "is_rush_order" enum field.
type IsRushOrder string

// IsRushOrder values.
const (
	IsRushOrderY IsRushOrder = "Y"
	IsRushOrderN IsRushOrder = "N"
	IsRushOrder  IsRushOrder = "-"
)

func (iro IsRushOrder) String() string {
	return string(iro)
}

// IsRushOrderValidator is a validator for the "is_rush_order" field enum values. It is called by the builders before save.
func IsRushOrderValidator(iro IsRushOrder) error {
	switch iro {
	case IsRushOrderY, IsRushOrderN, IsRushOrder:
		return nil
	default:
		return fmt.Errorf("orderheader: invalid enum value for is_rush_order field: %q", iro)
	}
}

// IsViewed defines the type for the "is_viewed" enum field.
type IsViewed string

// IsViewed values.
const (
	IsViewedY IsViewed = "Y"
	IsViewedN IsViewed = "N"
	IsViewed  IsViewed = "-"
)

func (iv IsViewed) String() string {
	return string(iv)
}

// IsViewedValidator is a validator for the "is_viewed" field enum values. It is called by the builders before save.
func IsViewedValidator(iv IsViewed) error {
	switch iv {
	case IsViewedY, IsViewedN, IsViewed:
		return nil
	default:
		return fmt.Errorf("orderheader: invalid enum value for is_viewed field: %q", iv)
	}
}

// InvoicePerShipment defines the type for the "invoice_per_shipment" enum field.
type InvoicePerShipment string

// InvoicePerShipment values.
const (
	InvoicePerShipmentY InvoicePerShipment = "Y"
	InvoicePerShipmentN InvoicePerShipment = "N"
	InvoicePerShipment  InvoicePerShipment = "-"
)

func (ips InvoicePerShipment) String() string {
	return string(ips)
}

// InvoicePerShipmentValidator is a validator for the "invoice_per_shipment" field enum values. It is called by the builders before save.
func InvoicePerShipmentValidator(ips InvoicePerShipment) error {
	switch ips {
	case InvoicePerShipmentY, InvoicePerShipmentN, InvoicePerShipment:
		return nil
	default:
		return fmt.Errorf("orderheader: invalid enum value for invoice_per_shipment field: %q", ips)
	}
}