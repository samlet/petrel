// Code generated by entc, DO NOT EDIT.

package payment

import (
	"time"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldPaymentTypeID holds the string denoting the payment_type_id field in the database.
	FieldPaymentTypeID = "payment_type_id"
	// FieldPaymentMethodTypeID holds the string denoting the payment_method_type_id field in the database.
	FieldPaymentMethodTypeID = "payment_method_type_id"
	// FieldPaymentMethodID holds the string denoting the payment_method_id field in the database.
	FieldPaymentMethodID = "payment_method_id"
	// FieldPaymentGatewayResponseID holds the string denoting the payment_gateway_response_id field in the database.
	FieldPaymentGatewayResponseID = "payment_gateway_response_id"
	// FieldPaymentPreferenceID holds the string denoting the payment_preference_id field in the database.
	FieldPaymentPreferenceID = "payment_preference_id"
	// FieldPartyIDFrom holds the string denoting the party_id_from field in the database.
	FieldPartyIDFrom = "party_id_from"
	// FieldPartyIDTo holds the string denoting the party_id_to field in the database.
	FieldPartyIDTo = "party_id_to"
	// FieldRoleTypeIDTo holds the string denoting the role_type_id_to field in the database.
	FieldRoleTypeIDTo = "role_type_id_to"
	// FieldStatusID holds the string denoting the status_id field in the database.
	FieldStatusID = "status_id"
	// FieldEffectiveDate holds the string denoting the effective_date field in the database.
	FieldEffectiveDate = "effective_date"
	// FieldPaymentRefNum holds the string denoting the payment_ref_num field in the database.
	FieldPaymentRefNum = "payment_ref_num"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCurrencyUomID holds the string denoting the currency_uom_id field in the database.
	FieldCurrencyUomID = "currency_uom_id"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldFinAccountTransID holds the string denoting the fin_account_trans_id field in the database.
	FieldFinAccountTransID = "fin_account_trans_id"
	// FieldOverrideGlAccountID holds the string denoting the override_gl_account_id field in the database.
	FieldOverrideGlAccountID = "override_gl_account_id"
	// FieldActualCurrencyAmount holds the string denoting the actual_currency_amount field in the database.
	FieldActualCurrencyAmount = "actual_currency_amount"
	// FieldActualCurrencyUomID holds the string denoting the actual_currency_uom_id field in the database.
	FieldActualCurrencyUomID = "actual_currency_uom_id"
	// Table holds the table name of the payment in the database.
	Table = "payments"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldPaymentTypeID,
	FieldPaymentMethodTypeID,
	FieldPaymentMethodID,
	FieldPaymentGatewayResponseID,
	FieldPaymentPreferenceID,
	FieldPartyIDFrom,
	FieldPartyIDTo,
	FieldRoleTypeIDTo,
	FieldStatusID,
	FieldEffectiveDate,
	FieldPaymentRefNum,
	FieldAmount,
	FieldCurrencyUomID,
	FieldComments,
	FieldFinAccountTransID,
	FieldOverrideGlAccountID,
	FieldActualCurrencyAmount,
	FieldActualCurrencyUomID,
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
	// DefaultEffectiveDate holds the default value on creation for the "effective_date" field.
	DefaultEffectiveDate func() time.Time
)
