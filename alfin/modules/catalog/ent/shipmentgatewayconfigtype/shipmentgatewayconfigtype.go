// Code generated by entc, DO NOT EDIT.

package shipmentgatewayconfigtype

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the shipmentgatewayconfigtype type in the database.
	Label = "shipment_gateway_config_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStringRef holds the string denoting the string_ref field in the database.
	FieldStringRef = "string_ref"
	// FieldHasTable holds the string denoting the has_table field in the database.
	FieldHasTable = "has_table"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeSiblingShipmentGatewayConfigTypes holds the string denoting the sibling_shipment_gateway_config_types edge name in mutations.
	EdgeSiblingShipmentGatewayConfigTypes = "sibling_shipment_gateway_config_types"
	// EdgeShipmentGatewayConfigs holds the string denoting the shipment_gateway_configs edge name in mutations.
	EdgeShipmentGatewayConfigs = "shipment_gateway_configs"
	// EdgeChildShipmentGatewayConfigTypes holds the string denoting the child_shipment_gateway_config_types edge name in mutations.
	EdgeChildShipmentGatewayConfigTypes = "child_shipment_gateway_config_types"
	// Table holds the table name of the shipmentgatewayconfigtype in the database.
	Table = "shipment_gateway_config_types"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "shipment_gateway_config_types"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "shipment_gateway_config_type_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "shipment_gateway_config_types"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "shipment_gateway_config_type_children"
	// SiblingShipmentGatewayConfigTypesTable is the table the holds the sibling_shipment_gateway_config_types relation/edge. The primary key declared below.
	SiblingShipmentGatewayConfigTypesTable = "shipment_gateway_config_type_sibling_shipment_gateway_config_types"
	// ShipmentGatewayConfigsTable is the table the holds the shipment_gateway_configs relation/edge.
	ShipmentGatewayConfigsTable = "shipment_gateway_configs"
	// ShipmentGatewayConfigsInverseTable is the table name for the ShipmentGatewayConfig entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentgatewayconfig" package.
	ShipmentGatewayConfigsInverseTable = "shipment_gateway_configs"
	// ShipmentGatewayConfigsColumn is the table column denoting the shipment_gateway_configs relation/edge.
	ShipmentGatewayConfigsColumn = "shipment_gateway_config_type_shipment_gateway_configs"
	// ChildShipmentGatewayConfigTypesTable is the table the holds the child_shipment_gateway_config_types relation/edge. The primary key declared below.
	ChildShipmentGatewayConfigTypesTable = "shipment_gateway_config_type_child_shipment_gateway_config_types"
)

// Columns holds all SQL columns for shipmentgatewayconfigtype fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStringRef,
	FieldHasTable,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "shipment_gateway_config_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"shipment_gateway_config_type_children",
}

var (
	// SiblingShipmentGatewayConfigTypesPrimaryKey and SiblingShipmentGatewayConfigTypesColumn2 are the table columns denoting the
	// primary key for the sibling_shipment_gateway_config_types relation (M2M).
	SiblingShipmentGatewayConfigTypesPrimaryKey = []string{"shipment_gateway_config_type_id", "sibling_shipment_gateway_config_type_id"}
	// ChildShipmentGatewayConfigTypesPrimaryKey and ChildShipmentGatewayConfigTypesColumn2 are the table columns denoting the
	// primary key for the child_shipment_gateway_config_types relation (M2M).
	ChildShipmentGatewayConfigTypesPrimaryKey = []string{"shipment_gateway_config_type_id", "child_shipment_gateway_config_type_id"}
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

// HasTable defines the type for the "has_table" enum field.
type HasTable string

// HasTable values.
const (
	HasTableYes     HasTable = "Yes"
	HasTableNo      HasTable = "No"
	HasTableUnknown HasTable = "Unknown"
)

func (ht HasTable) String() string {
	return string(ht)
}

// HasTableValidator is a validator for the "has_table" field enum values. It is called by the builders before save.
func HasTableValidator(ht HasTable) error {
	switch ht {
	case HasTableYes, HasTableNo, HasTableUnknown:
		return nil
	default:
		return fmt.Errorf("shipmentgatewayconfigtype: invalid enum value for has_table field: %q", ht)
	}
}
