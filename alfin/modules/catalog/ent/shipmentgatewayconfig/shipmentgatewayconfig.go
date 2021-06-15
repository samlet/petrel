// Code generated by entc, DO NOT EDIT.

package shipmentgatewayconfig

import (
	"time"
)

const (
	// Label holds the string label denoting the shipmentgatewayconfig type in the database.
	Label = "shipment_gateway_config"
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
	// EdgeShipmentGatewayConfigType holds the string denoting the shipment_gateway_config_type edge name in mutations.
	EdgeShipmentGatewayConfigType = "shipment_gateway_config_type"
	// EdgeShipmentGatewayDhl holds the string denoting the shipment_gateway_dhl edge name in mutations.
	EdgeShipmentGatewayDhl = "shipment_gateway_dhl"
	// EdgeShipmentGatewayFedex holds the string denoting the shipment_gateway_fedex edge name in mutations.
	EdgeShipmentGatewayFedex = "shipment_gateway_fedex"
	// EdgeShipmentGatewayUps holds the string denoting the shipment_gateway_ups edge name in mutations.
	EdgeShipmentGatewayUps = "shipment_gateway_ups"
	// EdgeShipmentGatewayUsps holds the string denoting the shipment_gateway_usps edge name in mutations.
	EdgeShipmentGatewayUsps = "shipment_gateway_usps"
	// Table holds the table name of the shipmentgatewayconfig in the database.
	Table = "shipment_gateway_configs"
	// ShipmentGatewayConfigTypeTable is the table the holds the shipment_gateway_config_type relation/edge.
	ShipmentGatewayConfigTypeTable = "shipment_gateway_configs"
	// ShipmentGatewayConfigTypeInverseTable is the table name for the ShipmentGatewayConfigType entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewayconfigtype" package.
	ShipmentGatewayConfigTypeInverseTable = "shipment_gateway_config_types"
	// ShipmentGatewayConfigTypeColumn is the table column denoting the shipment_gateway_config_type relation/edge.
	ShipmentGatewayConfigTypeColumn = "shipment_gateway_config_type_shipment_gateway_configs"
	// ShipmentGatewayDhlTable is the table the holds the shipment_gateway_dhl relation/edge.
	ShipmentGatewayDhlTable = "shipment_gateway_dhls"
	// ShipmentGatewayDhlInverseTable is the table name for the ShipmentGatewayDhl entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewaydhl" package.
	ShipmentGatewayDhlInverseTable = "shipment_gateway_dhls"
	// ShipmentGatewayDhlColumn is the table column denoting the shipment_gateway_dhl relation/edge.
	ShipmentGatewayDhlColumn = "shipment_gateway_config_shipment_gateway_dhl"
	// ShipmentGatewayFedexTable is the table the holds the shipment_gateway_fedex relation/edge.
	ShipmentGatewayFedexTable = "shipment_gateway_fedexes"
	// ShipmentGatewayFedexInverseTable is the table name for the ShipmentGatewayFedex entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewayfedex" package.
	ShipmentGatewayFedexInverseTable = "shipment_gateway_fedexes"
	// ShipmentGatewayFedexColumn is the table column denoting the shipment_gateway_fedex relation/edge.
	ShipmentGatewayFedexColumn = "shipment_gateway_config_shipment_gateway_fedex"
	// ShipmentGatewayUpsTable is the table the holds the shipment_gateway_ups relation/edge.
	ShipmentGatewayUpsTable = "shipment_gateway_ups"
	// ShipmentGatewayUpsInverseTable is the table name for the ShipmentGatewayUps entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewayups" package.
	ShipmentGatewayUpsInverseTable = "shipment_gateway_ups"
	// ShipmentGatewayUpsColumn is the table column denoting the shipment_gateway_ups relation/edge.
	ShipmentGatewayUpsColumn = "shipment_gateway_config_shipment_gateway_ups"
	// ShipmentGatewayUspsTable is the table the holds the shipment_gateway_usps relation/edge.
	ShipmentGatewayUspsTable = "shipment_gateway_usps"
	// ShipmentGatewayUspsInverseTable is the table name for the ShipmentGatewayUsps entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewayusps" package.
	ShipmentGatewayUspsInverseTable = "shipment_gateway_usps"
	// ShipmentGatewayUspsColumn is the table column denoting the shipment_gateway_usps relation/edge.
	ShipmentGatewayUspsColumn = "shipment_gateway_config_shipment_gateway_usps"
)

// Columns holds all SQL columns for shipmentgatewayconfig fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "shipment_gateway_configs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"shipment_gateway_config_type_shipment_gateway_configs",
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
