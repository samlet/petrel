// Code generated by entc, DO NOT EDIT.

package orderpaymentpreference

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the orderpaymentpreference type in the database.
	Label = "order_payment_preference"
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
	// FieldPaymentMethodTypeID holds the string denoting the payment_method_type_id field in the database.
	FieldPaymentMethodTypeID = "payment_method_type_id"
	// FieldPaymentMethodID holds the string denoting the payment_method_id field in the database.
	FieldPaymentMethodID = "payment_method_id"
	// FieldFinAccountID holds the string denoting the fin_account_id field in the database.
	FieldFinAccountID = "fin_account_id"
	// FieldSecurityCode holds the string denoting the security_code field in the database.
	FieldSecurityCode = "security_code"
	// FieldTrack2 holds the string denoting the track_2 field in the database.
	FieldTrack2 = "track_2"
	// FieldPresentFlag holds the string denoting the present_flag field in the database.
	FieldPresentFlag = "present_flag"
	// FieldSwipedFlag holds the string denoting the swiped_flag field in the database.
	FieldSwipedFlag = "swiped_flag"
	// FieldOverflowFlag holds the string denoting the overflow_flag field in the database.
	FieldOverflowFlag = "overflow_flag"
	// FieldMaxAmount holds the string denoting the max_amount field in the database.
	FieldMaxAmount = "max_amount"
	// FieldProcessAttempt holds the string denoting the process_attempt field in the database.
	FieldProcessAttempt = "process_attempt"
	// FieldBillingPostalCode holds the string denoting the billing_postal_code field in the database.
	FieldBillingPostalCode = "billing_postal_code"
	// FieldManualAuthCode holds the string denoting the manual_auth_code field in the database.
	FieldManualAuthCode = "manual_auth_code"
	// FieldManualRefNum holds the string denoting the manual_ref_num field in the database.
	FieldManualRefNum = "manual_ref_num"
	// FieldNeedsNsfRetry holds the string denoting the needs_nsf_retry field in the database.
	FieldNeedsNsfRetry = "needs_nsf_retry"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldCreatedByUserLogin holds the string denoting the created_by_user_login field in the database.
	FieldCreatedByUserLogin = "created_by_user_login"
	// FieldLastModifiedDate holds the string denoting the last_modified_date field in the database.
	FieldLastModifiedDate = "last_modified_date"
	// FieldLastModifiedByUserLogin holds the string denoting the last_modified_by_user_login field in the database.
	FieldLastModifiedByUserLogin = "last_modified_by_user_login"
	// EdgeOrderHeader holds the string denoting the order_header edge name in mutations.
	EdgeOrderHeader = "order_header"
	// EdgeOrderItem holds the string denoting the order_item edge name in mutations.
	EdgeOrderItem = "order_item"
	// EdgeOrderItemShipGroup holds the string denoting the order_item_ship_group edge name in mutations.
	EdgeOrderItemShipGroup = "order_item_ship_group"
	// EdgeProductPricePurpose holds the string denoting the product_price_purpose edge name in mutations.
	EdgeProductPricePurpose = "product_price_purpose"
	// EdgeStatusItem holds the string denoting the status_item edge name in mutations.
	EdgeStatusItem = "status_item"
	// EdgeOrderStatuses holds the string denoting the order_statuses edge name in mutations.
	EdgeOrderStatuses = "order_statuses"
	// Table holds the table name of the orderpaymentpreference in the database.
	Table = "order_payment_preferences"
	// OrderHeaderTable is the table the holds the order_header relation/edge.
	OrderHeaderTable = "order_payment_preferences"
	// OrderHeaderInverseTable is the table name for the OrderHeader entity.
	// It exists in this package in order to avoid circular dependency with the "orderheader" package.
	OrderHeaderInverseTable = "order_headers"
	// OrderHeaderColumn is the table column denoting the order_header relation/edge.
	OrderHeaderColumn = "order_header_order_payment_preferences"
	// OrderItemTable is the table the holds the order_item relation/edge.
	OrderItemTable = "order_payment_preferences"
	// OrderItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	OrderItemInverseTable = "order_items"
	// OrderItemColumn is the table column denoting the order_item relation/edge.
	OrderItemColumn = "order_item_order_payment_preferences"
	// OrderItemShipGroupTable is the table the holds the order_item_ship_group relation/edge.
	OrderItemShipGroupTable = "order_payment_preferences"
	// OrderItemShipGroupInverseTable is the table name for the OrderItemShipGroup entity.
	// It exists in this package in order to avoid circular dependency with the "orderitemshipgroup" package.
	OrderItemShipGroupInverseTable = "order_item_ship_groups"
	// OrderItemShipGroupColumn is the table column denoting the order_item_ship_group relation/edge.
	OrderItemShipGroupColumn = "order_item_ship_group_order_payment_preferences"
	// ProductPricePurposeTable is the table the holds the product_price_purpose relation/edge.
	ProductPricePurposeTable = "order_payment_preferences"
	// ProductPricePurposeInverseTable is the table name for the ProductPricePurpose entity.
	// It exists in this package in order to avoid circular dependency with the "productpricepurpose" package.
	ProductPricePurposeInverseTable = "product_price_purposes"
	// ProductPricePurposeColumn is the table column denoting the product_price_purpose relation/edge.
	ProductPricePurposeColumn = "product_price_purpose_order_payment_preferences"
	// StatusItemTable is the table the holds the status_item relation/edge.
	StatusItemTable = "order_payment_preferences"
	// StatusItemInverseTable is the table name for the StatusItem entity.
	// It exists in this package in order to avoid circular dependency with the "statusitem" package.
	StatusItemInverseTable = "status_items"
	// StatusItemColumn is the table column denoting the status_item relation/edge.
	StatusItemColumn = "status_item_order_payment_preferences"
	// OrderStatusesTable is the table the holds the order_statuses relation/edge.
	OrderStatusesTable = "order_status"
	// OrderStatusesInverseTable is the table name for the OrderStatus entity.
	// It exists in this package in order to avoid circular dependency with the "orderstatus" package.
	OrderStatusesInverseTable = "order_status"
	// OrderStatusesColumn is the table column denoting the order_statuses relation/edge.
	OrderStatusesColumn = "order_payment_preference_order_statuses"
)

// Columns holds all SQL columns for orderpaymentpreference fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldOrderItemSeqID,
	FieldShipGroupSeqID,
	FieldPaymentMethodTypeID,
	FieldPaymentMethodID,
	FieldFinAccountID,
	FieldSecurityCode,
	FieldTrack2,
	FieldPresentFlag,
	FieldSwipedFlag,
	FieldOverflowFlag,
	FieldMaxAmount,
	FieldProcessAttempt,
	FieldBillingPostalCode,
	FieldManualAuthCode,
	FieldManualRefNum,
	FieldNeedsNsfRetry,
	FieldCreatedDate,
	FieldCreatedByUserLogin,
	FieldLastModifiedDate,
	FieldLastModifiedByUserLogin,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_payment_preferences"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_header_order_payment_preferences",
	"order_item_order_payment_preferences",
	"order_item_ship_group_order_payment_preferences",
	"product_price_purpose_order_payment_preferences",
	"status_item_order_payment_preferences",
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
	// DefaultCreatedDate holds the default value on creation for the "created_date" field.
	DefaultCreatedDate func() time.Time
	// DefaultLastModifiedDate holds the default value on creation for the "last_modified_date" field.
	DefaultLastModifiedDate func() time.Time
)

// PresentFlag defines the type for the "present_flag" enum field.
type PresentFlag string

// PresentFlag values.
const (
	PresentFlagYes     PresentFlag = "Yes"
	PresentFlagNo      PresentFlag = "No"
	PresentFlagUnknown PresentFlag = "Unknown"
)

func (pf PresentFlag) String() string {
	return string(pf)
}

// PresentFlagValidator is a validator for the "present_flag" field enum values. It is called by the builders before save.
func PresentFlagValidator(pf PresentFlag) error {
	switch pf {
	case PresentFlagYes, PresentFlagNo, PresentFlagUnknown:
		return nil
	default:
		return fmt.Errorf("orderpaymentpreference: invalid enum value for present_flag field: %q", pf)
	}
}

// SwipedFlag defines the type for the "swiped_flag" enum field.
type SwipedFlag string

// SwipedFlag values.
const (
	SwipedFlagYes     SwipedFlag = "Yes"
	SwipedFlagNo      SwipedFlag = "No"
	SwipedFlagUnknown SwipedFlag = "Unknown"
)

func (sf SwipedFlag) String() string {
	return string(sf)
}

// SwipedFlagValidator is a validator for the "swiped_flag" field enum values. It is called by the builders before save.
func SwipedFlagValidator(sf SwipedFlag) error {
	switch sf {
	case SwipedFlagYes, SwipedFlagNo, SwipedFlagUnknown:
		return nil
	default:
		return fmt.Errorf("orderpaymentpreference: invalid enum value for swiped_flag field: %q", sf)
	}
}

// OverflowFlag defines the type for the "overflow_flag" enum field.
type OverflowFlag string

// OverflowFlag values.
const (
	OverflowFlagYes     OverflowFlag = "Yes"
	OverflowFlagNo      OverflowFlag = "No"
	OverflowFlagUnknown OverflowFlag = "Unknown"
)

func (of OverflowFlag) String() string {
	return string(of)
}

// OverflowFlagValidator is a validator for the "overflow_flag" field enum values. It is called by the builders before save.
func OverflowFlagValidator(of OverflowFlag) error {
	switch of {
	case OverflowFlagYes, OverflowFlagNo, OverflowFlagUnknown:
		return nil
	default:
		return fmt.Errorf("orderpaymentpreference: invalid enum value for overflow_flag field: %q", of)
	}
}

// NeedsNsfRetry defines the type for the "needs_nsf_retry" enum field.
type NeedsNsfRetry string

// NeedsNsfRetry values.
const (
	NeedsNsfRetryYes     NeedsNsfRetry = "Yes"
	NeedsNsfRetryNo      NeedsNsfRetry = "No"
	NeedsNsfRetryUnknown NeedsNsfRetry = "Unknown"
)

func (nnr NeedsNsfRetry) String() string {
	return string(nnr)
}

// NeedsNsfRetryValidator is a validator for the "needs_nsf_retry" field enum values. It is called by the builders before save.
func NeedsNsfRetryValidator(nnr NeedsNsfRetry) error {
	switch nnr {
	case NeedsNsfRetryYes, NeedsNsfRetryNo, NeedsNsfRetryUnknown:
		return nil
	default:
		return fmt.Errorf("orderpaymentpreference: invalid enum value for needs_nsf_retry field: %q", nnr)
	}
}