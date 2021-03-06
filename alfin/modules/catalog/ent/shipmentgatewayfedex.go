// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfig"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayfedex"
)

// ShipmentGatewayFedex is the model entity for the ShipmentGatewayFedex schema.
type ShipmentGatewayFedex struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// ConnectURL holds the value of the "connect_url" field.
	ConnectURL string `json:"connect_url,omitempty"`
	// ConnectSoapURL holds the value of the "connect_soap_url" field.
	ConnectSoapURL string `json:"connect_soap_url,omitempty"`
	// ConnectTimeout holds the value of the "connect_timeout" field.
	ConnectTimeout int `json:"connect_timeout,omitempty"`
	// AccessAccountNbr holds the value of the "access_account_nbr" field.
	AccessAccountNbr string `json:"access_account_nbr,omitempty"`
	// AccessMeterNumber holds the value of the "access_meter_number" field.
	AccessMeterNumber string `json:"access_meter_number,omitempty"`
	// AccessUserKey holds the value of the "access_user_key" field.
	AccessUserKey string `json:"access_user_key,omitempty"`
	// AccessUserPwd holds the value of the "access_user_pwd" field.
	AccessUserPwd string `json:"access_user_pwd,omitempty"`
	// LabelImageType holds the value of the "label_image_type" field.
	LabelImageType string `json:"label_image_type,omitempty"`
	// DefaultDropoffType holds the value of the "default_dropoff_type" field.
	DefaultDropoffType string `json:"default_dropoff_type,omitempty"`
	// DefaultPackagingType holds the value of the "default_packaging_type" field.
	DefaultPackagingType string `json:"default_packaging_type,omitempty"`
	// TemplateShipment holds the value of the "template_shipment" field.
	TemplateShipment string `json:"template_shipment,omitempty"`
	// TemplateSubscription holds the value of the "template_subscription" field.
	TemplateSubscription string `json:"template_subscription,omitempty"`
	// RateEstimateTemplate holds the value of the "rate_estimate_template" field.
	RateEstimateTemplate string `json:"rate_estimate_template,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentGatewayFedexQuery when eager-loading is set.
	Edges                                          ShipmentGatewayFedexEdges `json:"edges"`
	shipment_gateway_config_shipment_gateway_fedex *int
}

// ShipmentGatewayFedexEdges holds the relations/edges for other nodes in the graph.
type ShipmentGatewayFedexEdges struct {
	// ShipmentGatewayConfig holds the value of the shipment_gateway_config edge.
	ShipmentGatewayConfig *ShipmentGatewayConfig `json:"shipment_gateway_config,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ShipmentGatewayConfigOrErr returns the ShipmentGatewayConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentGatewayFedexEdges) ShipmentGatewayConfigOrErr() (*ShipmentGatewayConfig, error) {
	if e.loadedTypes[0] {
		if e.ShipmentGatewayConfig == nil {
			// The edge shipment_gateway_config was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shipmentgatewayconfig.Label}
		}
		return e.ShipmentGatewayConfig, nil
	}
	return nil, &NotLoadedError{edge: "shipment_gateway_config"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentGatewayFedex) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayfedex.FieldID, shipmentgatewayfedex.FieldConnectTimeout:
			values[i] = new(sql.NullInt64)
		case shipmentgatewayfedex.FieldStringRef, shipmentgatewayfedex.FieldConnectURL, shipmentgatewayfedex.FieldConnectSoapURL, shipmentgatewayfedex.FieldAccessAccountNbr, shipmentgatewayfedex.FieldAccessMeterNumber, shipmentgatewayfedex.FieldAccessUserKey, shipmentgatewayfedex.FieldAccessUserPwd, shipmentgatewayfedex.FieldLabelImageType, shipmentgatewayfedex.FieldDefaultDropoffType, shipmentgatewayfedex.FieldDefaultPackagingType, shipmentgatewayfedex.FieldTemplateShipment, shipmentgatewayfedex.FieldTemplateSubscription, shipmentgatewayfedex.FieldRateEstimateTemplate:
			values[i] = new(sql.NullString)
		case shipmentgatewayfedex.FieldCreateTime, shipmentgatewayfedex.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case shipmentgatewayfedex.ForeignKeys[0]: // shipment_gateway_config_shipment_gateway_fedex
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShipmentGatewayFedex", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentGatewayFedex fields.
func (sgf *ShipmentGatewayFedex) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayfedex.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sgf.ID = int(value.Int64)
		case shipmentgatewayfedex.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sgf.CreateTime = value.Time
			}
		case shipmentgatewayfedex.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				sgf.UpdateTime = value.Time
			}
		case shipmentgatewayfedex.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				sgf.StringRef = value.String
			}
		case shipmentgatewayfedex.FieldConnectURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connect_url", values[i])
			} else if value.Valid {
				sgf.ConnectURL = value.String
			}
		case shipmentgatewayfedex.FieldConnectSoapURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connect_soap_url", values[i])
			} else if value.Valid {
				sgf.ConnectSoapURL = value.String
			}
		case shipmentgatewayfedex.FieldConnectTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connect_timeout", values[i])
			} else if value.Valid {
				sgf.ConnectTimeout = int(value.Int64)
			}
		case shipmentgatewayfedex.FieldAccessAccountNbr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_account_nbr", values[i])
			} else if value.Valid {
				sgf.AccessAccountNbr = value.String
			}
		case shipmentgatewayfedex.FieldAccessMeterNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_meter_number", values[i])
			} else if value.Valid {
				sgf.AccessMeterNumber = value.String
			}
		case shipmentgatewayfedex.FieldAccessUserKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_user_key", values[i])
			} else if value.Valid {
				sgf.AccessUserKey = value.String
			}
		case shipmentgatewayfedex.FieldAccessUserPwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_user_pwd", values[i])
			} else if value.Valid {
				sgf.AccessUserPwd = value.String
			}
		case shipmentgatewayfedex.FieldLabelImageType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label_image_type", values[i])
			} else if value.Valid {
				sgf.LabelImageType = value.String
			}
		case shipmentgatewayfedex.FieldDefaultDropoffType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_dropoff_type", values[i])
			} else if value.Valid {
				sgf.DefaultDropoffType = value.String
			}
		case shipmentgatewayfedex.FieldDefaultPackagingType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_packaging_type", values[i])
			} else if value.Valid {
				sgf.DefaultPackagingType = value.String
			}
		case shipmentgatewayfedex.FieldTemplateShipment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_shipment", values[i])
			} else if value.Valid {
				sgf.TemplateShipment = value.String
			}
		case shipmentgatewayfedex.FieldTemplateSubscription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_subscription", values[i])
			} else if value.Valid {
				sgf.TemplateSubscription = value.String
			}
		case shipmentgatewayfedex.FieldRateEstimateTemplate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rate_estimate_template", values[i])
			} else if value.Valid {
				sgf.RateEstimateTemplate = value.String
			}
		case shipmentgatewayfedex.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shipment_gateway_config_shipment_gateway_fedex", value)
			} else if value.Valid {
				sgf.shipment_gateway_config_shipment_gateway_fedex = new(int)
				*sgf.shipment_gateway_config_shipment_gateway_fedex = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryShipmentGatewayConfig queries the "shipment_gateway_config" edge of the ShipmentGatewayFedex entity.
func (sgf *ShipmentGatewayFedex) QueryShipmentGatewayConfig() *ShipmentGatewayConfigQuery {
	return (&ShipmentGatewayFedexClient{config: sgf.config}).QueryShipmentGatewayConfig(sgf)
}

// Update returns a builder for updating this ShipmentGatewayFedex.
// Note that you need to call ShipmentGatewayFedex.Unwrap() before calling this method if this ShipmentGatewayFedex
// was returned from a transaction, and the transaction was committed or rolled back.
func (sgf *ShipmentGatewayFedex) Update() *ShipmentGatewayFedexUpdateOne {
	return (&ShipmentGatewayFedexClient{config: sgf.config}).UpdateOne(sgf)
}

// Unwrap unwraps the ShipmentGatewayFedex entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sgf *ShipmentGatewayFedex) Unwrap() *ShipmentGatewayFedex {
	tx, ok := sgf.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentGatewayFedex is not a transactional entity")
	}
	sgf.config.driver = tx.drv
	return sgf
}

// String implements the fmt.Stringer.
func (sgf *ShipmentGatewayFedex) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentGatewayFedex(")
	builder.WriteString(fmt.Sprintf("id=%v", sgf.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(sgf.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(sgf.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(sgf.StringRef)
	builder.WriteString(", connect_url=")
	builder.WriteString(sgf.ConnectURL)
	builder.WriteString(", connect_soap_url=")
	builder.WriteString(sgf.ConnectSoapURL)
	builder.WriteString(", connect_timeout=")
	builder.WriteString(fmt.Sprintf("%v", sgf.ConnectTimeout))
	builder.WriteString(", access_account_nbr=")
	builder.WriteString(sgf.AccessAccountNbr)
	builder.WriteString(", access_meter_number=")
	builder.WriteString(sgf.AccessMeterNumber)
	builder.WriteString(", access_user_key=")
	builder.WriteString(sgf.AccessUserKey)
	builder.WriteString(", access_user_pwd=")
	builder.WriteString(sgf.AccessUserPwd)
	builder.WriteString(", label_image_type=")
	builder.WriteString(sgf.LabelImageType)
	builder.WriteString(", default_dropoff_type=")
	builder.WriteString(sgf.DefaultDropoffType)
	builder.WriteString(", default_packaging_type=")
	builder.WriteString(sgf.DefaultPackagingType)
	builder.WriteString(", template_shipment=")
	builder.WriteString(sgf.TemplateShipment)
	builder.WriteString(", template_subscription=")
	builder.WriteString(sgf.TemplateSubscription)
	builder.WriteString(", rate_estimate_template=")
	builder.WriteString(sgf.RateEstimateTemplate)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentGatewayFedexes is a parsable slice of ShipmentGatewayFedex.
type ShipmentGatewayFedexes []*ShipmentGatewayFedex

func (sgf ShipmentGatewayFedexes) config(cfg config) {
	for _i := range sgf {
		sgf[_i].config = cfg
	}
}
