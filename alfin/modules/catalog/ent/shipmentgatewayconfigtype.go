// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfigtype"
)

// ShipmentGatewayConfigType is the model entity for the ShipmentGatewayConfigType schema.
type ShipmentGatewayConfigType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// HasTable holds the value of the "has_table" field.
	HasTable shipmentgatewayconfigtype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentGatewayConfigTypeQuery when eager-loading is set.
	Edges                                 ShipmentGatewayConfigTypeEdges `json:"edges"`
	shipment_gateway_config_type_children *int
}

// ShipmentGatewayConfigTypeEdges holds the relations/edges for other nodes in the graph.
type ShipmentGatewayConfigTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ShipmentGatewayConfigType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*ShipmentGatewayConfigType `json:"children,omitempty"`
	// SiblingShipmentGatewayConfigTypes holds the value of the sibling_shipment_gateway_config_types edge.
	SiblingShipmentGatewayConfigTypes []*ShipmentGatewayConfigType `json:"sibling_shipment_gateway_config_types,omitempty"`
	// ShipmentGatewayConfigs holds the value of the shipment_gateway_configs edge.
	ShipmentGatewayConfigs []*ShipmentGatewayConfig `json:"shipment_gateway_configs,omitempty"`
	// ChildShipmentGatewayConfigTypes holds the value of the child_shipment_gateway_config_types edge.
	ChildShipmentGatewayConfigTypes []*ShipmentGatewayConfigType `json:"child_shipment_gateway_config_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentGatewayConfigTypeEdges) ParentOrErr() (*ShipmentGatewayConfigType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shipmentgatewayconfigtype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentGatewayConfigTypeEdges) ChildrenOrErr() ([]*ShipmentGatewayConfigType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// SiblingShipmentGatewayConfigTypesOrErr returns the SiblingShipmentGatewayConfigTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentGatewayConfigTypeEdges) SiblingShipmentGatewayConfigTypesOrErr() ([]*ShipmentGatewayConfigType, error) {
	if e.loadedTypes[2] {
		return e.SiblingShipmentGatewayConfigTypes, nil
	}
	return nil, &NotLoadedError{edge: "sibling_shipment_gateway_config_types"}
}

// ShipmentGatewayConfigsOrErr returns the ShipmentGatewayConfigs value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentGatewayConfigTypeEdges) ShipmentGatewayConfigsOrErr() ([]*ShipmentGatewayConfig, error) {
	if e.loadedTypes[3] {
		return e.ShipmentGatewayConfigs, nil
	}
	return nil, &NotLoadedError{edge: "shipment_gateway_configs"}
}

// ChildShipmentGatewayConfigTypesOrErr returns the ChildShipmentGatewayConfigTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentGatewayConfigTypeEdges) ChildShipmentGatewayConfigTypesOrErr() ([]*ShipmentGatewayConfigType, error) {
	if e.loadedTypes[4] {
		return e.ChildShipmentGatewayConfigTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_shipment_gateway_config_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentGatewayConfigType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayconfigtype.FieldID:
			values[i] = new(sql.NullInt64)
		case shipmentgatewayconfigtype.FieldStringRef, shipmentgatewayconfigtype.FieldHasTable, shipmentgatewayconfigtype.FieldDescription:
			values[i] = new(sql.NullString)
		case shipmentgatewayconfigtype.FieldCreateTime, shipmentgatewayconfigtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case shipmentgatewayconfigtype.ForeignKeys[0]: // shipment_gateway_config_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShipmentGatewayConfigType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentGatewayConfigType fields.
func (sgct *ShipmentGatewayConfigType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayconfigtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sgct.ID = int(value.Int64)
		case shipmentgatewayconfigtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sgct.CreateTime = value.Time
			}
		case shipmentgatewayconfigtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				sgct.UpdateTime = value.Time
			}
		case shipmentgatewayconfigtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				sgct.StringRef = value.String
			}
		case shipmentgatewayconfigtype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				sgct.HasTable = shipmentgatewayconfigtype.HasTable(value.String)
			}
		case shipmentgatewayconfigtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sgct.Description = value.String
			}
		case shipmentgatewayconfigtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shipment_gateway_config_type_children", value)
			} else if value.Valid {
				sgct.shipment_gateway_config_type_children = new(int)
				*sgct.shipment_gateway_config_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ShipmentGatewayConfigType entity.
func (sgct *ShipmentGatewayConfigType) QueryParent() *ShipmentGatewayConfigTypeQuery {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).QueryParent(sgct)
}

// QueryChildren queries the "children" edge of the ShipmentGatewayConfigType entity.
func (sgct *ShipmentGatewayConfigType) QueryChildren() *ShipmentGatewayConfigTypeQuery {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).QueryChildren(sgct)
}

// QuerySiblingShipmentGatewayConfigTypes queries the "sibling_shipment_gateway_config_types" edge of the ShipmentGatewayConfigType entity.
func (sgct *ShipmentGatewayConfigType) QuerySiblingShipmentGatewayConfigTypes() *ShipmentGatewayConfigTypeQuery {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).QuerySiblingShipmentGatewayConfigTypes(sgct)
}

// QueryShipmentGatewayConfigs queries the "shipment_gateway_configs" edge of the ShipmentGatewayConfigType entity.
func (sgct *ShipmentGatewayConfigType) QueryShipmentGatewayConfigs() *ShipmentGatewayConfigQuery {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).QueryShipmentGatewayConfigs(sgct)
}

// QueryChildShipmentGatewayConfigTypes queries the "child_shipment_gateway_config_types" edge of the ShipmentGatewayConfigType entity.
func (sgct *ShipmentGatewayConfigType) QueryChildShipmentGatewayConfigTypes() *ShipmentGatewayConfigTypeQuery {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).QueryChildShipmentGatewayConfigTypes(sgct)
}

// Update returns a builder for updating this ShipmentGatewayConfigType.
// Note that you need to call ShipmentGatewayConfigType.Unwrap() before calling this method if this ShipmentGatewayConfigType
// was returned from a transaction, and the transaction was committed or rolled back.
func (sgct *ShipmentGatewayConfigType) Update() *ShipmentGatewayConfigTypeUpdateOne {
	return (&ShipmentGatewayConfigTypeClient{config: sgct.config}).UpdateOne(sgct)
}

// Unwrap unwraps the ShipmentGatewayConfigType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sgct *ShipmentGatewayConfigType) Unwrap() *ShipmentGatewayConfigType {
	tx, ok := sgct.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentGatewayConfigType is not a transactional entity")
	}
	sgct.config.driver = tx.drv
	return sgct
}

// String implements the fmt.Stringer.
func (sgct *ShipmentGatewayConfigType) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentGatewayConfigType(")
	builder.WriteString(fmt.Sprintf("id=%v", sgct.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(sgct.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(sgct.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(sgct.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", sgct.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(sgct.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentGatewayConfigTypes is a parsable slice of ShipmentGatewayConfigType.
type ShipmentGatewayConfigTypes []*ShipmentGatewayConfigType

func (sgct ShipmentGatewayConfigTypes) config(cfg config) {
	for _i := range sgct {
		sgct[_i].config = cfg
	}
}
