// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayconfig"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentgatewayups"
)

// ShipmentGatewayUps is the model entity for the ShipmentGatewayUps schema.
type ShipmentGatewayUps struct {
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
	// ConnectTimeout holds the value of the "connect_timeout" field.
	ConnectTimeout int `json:"connect_timeout,omitempty"`
	// ShipperNumber holds the value of the "shipper_number" field.
	ShipperNumber string `json:"shipper_number,omitempty"`
	// BillShipperAccountNumber holds the value of the "bill_shipper_account_number" field.
	BillShipperAccountNumber string `json:"bill_shipper_account_number,omitempty"`
	// AccessLicenseNumber holds the value of the "access_license_number" field.
	AccessLicenseNumber string `json:"access_license_number,omitempty"`
	// AccessUserID holds the value of the "access_user_id" field.
	AccessUserID string `json:"access_user_id,omitempty"`
	// AccessPassword holds the value of the "access_password" field.
	AccessPassword string `json:"access_password,omitempty"`
	// SaveCertInfo holds the value of the "save_cert_info" field.
	SaveCertInfo string `json:"save_cert_info,omitempty"`
	// SaveCertPath holds the value of the "save_cert_path" field.
	SaveCertPath string `json:"save_cert_path,omitempty"`
	// ShipperPickupType holds the value of the "shipper_pickup_type" field.
	ShipperPickupType string `json:"shipper_pickup_type,omitempty"`
	// CustomerClassification holds the value of the "customer_classification" field.
	CustomerClassification string `json:"customer_classification,omitempty"`
	// MaxEstimateWeight holds the value of the "max_estimate_weight" field.
	MaxEstimateWeight float64 `json:"max_estimate_weight,omitempty"`
	// MinEstimateWeight holds the value of the "min_estimate_weight" field.
	MinEstimateWeight float64 `json:"min_estimate_weight,omitempty"`
	// CodAllowCod holds the value of the "cod_allow_cod" field.
	CodAllowCod string `json:"cod_allow_cod,omitempty"`
	// CodSurchargeAmount holds the value of the "cod_surcharge_amount" field.
	CodSurchargeAmount float64 `json:"cod_surcharge_amount,omitempty"`
	// CodSurchargeCurrencyUomID holds the value of the "cod_surcharge_currency_uom_id" field.
	CodSurchargeCurrencyUomID string `json:"cod_surcharge_currency_uom_id,omitempty"`
	// CodSurchargeApplyToPackage holds the value of the "cod_surcharge_apply_to_package" field.
	CodSurchargeApplyToPackage string `json:"cod_surcharge_apply_to_package,omitempty"`
	// CodFundsCode holds the value of the "cod_funds_code" field.
	CodFundsCode string `json:"cod_funds_code,omitempty"`
	// DefaultReturnLabelMemo holds the value of the "default_return_label_memo" field.
	DefaultReturnLabelMemo string `json:"default_return_label_memo,omitempty"`
	// DefaultReturnLabelSubject holds the value of the "default_return_label_subject" field.
	DefaultReturnLabelSubject string `json:"default_return_label_subject,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentGatewayUpsQuery when eager-loading is set.
	Edges                                        ShipmentGatewayUpsEdges `json:"edges"`
	shipment_gateway_config_shipment_gateway_ups *int
}

// ShipmentGatewayUpsEdges holds the relations/edges for other nodes in the graph.
type ShipmentGatewayUpsEdges struct {
	// ShipmentGatewayConfig holds the value of the shipment_gateway_config edge.
	ShipmentGatewayConfig *ShipmentGatewayConfig `json:"shipment_gateway_config,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ShipmentGatewayConfigOrErr returns the ShipmentGatewayConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentGatewayUpsEdges) ShipmentGatewayConfigOrErr() (*ShipmentGatewayConfig, error) {
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
func (*ShipmentGatewayUps) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayups.FieldMaxEstimateWeight, shipmentgatewayups.FieldMinEstimateWeight, shipmentgatewayups.FieldCodSurchargeAmount:
			values[i] = new(sql.NullFloat64)
		case shipmentgatewayups.FieldID, shipmentgatewayups.FieldConnectTimeout:
			values[i] = new(sql.NullInt64)
		case shipmentgatewayups.FieldStringRef, shipmentgatewayups.FieldConnectURL, shipmentgatewayups.FieldShipperNumber, shipmentgatewayups.FieldBillShipperAccountNumber, shipmentgatewayups.FieldAccessLicenseNumber, shipmentgatewayups.FieldAccessUserID, shipmentgatewayups.FieldAccessPassword, shipmentgatewayups.FieldSaveCertInfo, shipmentgatewayups.FieldSaveCertPath, shipmentgatewayups.FieldShipperPickupType, shipmentgatewayups.FieldCustomerClassification, shipmentgatewayups.FieldCodAllowCod, shipmentgatewayups.FieldCodSurchargeCurrencyUomID, shipmentgatewayups.FieldCodSurchargeApplyToPackage, shipmentgatewayups.FieldCodFundsCode, shipmentgatewayups.FieldDefaultReturnLabelMemo, shipmentgatewayups.FieldDefaultReturnLabelSubject:
			values[i] = new(sql.NullString)
		case shipmentgatewayups.FieldCreateTime, shipmentgatewayups.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case shipmentgatewayups.ForeignKeys[0]: // shipment_gateway_config_shipment_gateway_ups
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShipmentGatewayUps", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentGatewayUps fields.
func (sgu *ShipmentGatewayUps) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentgatewayups.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sgu.ID = int(value.Int64)
		case shipmentgatewayups.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sgu.CreateTime = value.Time
			}
		case shipmentgatewayups.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				sgu.UpdateTime = value.Time
			}
		case shipmentgatewayups.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				sgu.StringRef = value.String
			}
		case shipmentgatewayups.FieldConnectURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field connect_url", values[i])
			} else if value.Valid {
				sgu.ConnectURL = value.String
			}
		case shipmentgatewayups.FieldConnectTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field connect_timeout", values[i])
			} else if value.Valid {
				sgu.ConnectTimeout = int(value.Int64)
			}
		case shipmentgatewayups.FieldShipperNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shipper_number", values[i])
			} else if value.Valid {
				sgu.ShipperNumber = value.String
			}
		case shipmentgatewayups.FieldBillShipperAccountNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bill_shipper_account_number", values[i])
			} else if value.Valid {
				sgu.BillShipperAccountNumber = value.String
			}
		case shipmentgatewayups.FieldAccessLicenseNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_license_number", values[i])
			} else if value.Valid {
				sgu.AccessLicenseNumber = value.String
			}
		case shipmentgatewayups.FieldAccessUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_user_id", values[i])
			} else if value.Valid {
				sgu.AccessUserID = value.String
			}
		case shipmentgatewayups.FieldAccessPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_password", values[i])
			} else if value.Valid {
				sgu.AccessPassword = value.String
			}
		case shipmentgatewayups.FieldSaveCertInfo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field save_cert_info", values[i])
			} else if value.Valid {
				sgu.SaveCertInfo = value.String
			}
		case shipmentgatewayups.FieldSaveCertPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field save_cert_path", values[i])
			} else if value.Valid {
				sgu.SaveCertPath = value.String
			}
		case shipmentgatewayups.FieldShipperPickupType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field shipper_pickup_type", values[i])
			} else if value.Valid {
				sgu.ShipperPickupType = value.String
			}
		case shipmentgatewayups.FieldCustomerClassification:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_classification", values[i])
			} else if value.Valid {
				sgu.CustomerClassification = value.String
			}
		case shipmentgatewayups.FieldMaxEstimateWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field max_estimate_weight", values[i])
			} else if value.Valid {
				sgu.MaxEstimateWeight = value.Float64
			}
		case shipmentgatewayups.FieldMinEstimateWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field min_estimate_weight", values[i])
			} else if value.Valid {
				sgu.MinEstimateWeight = value.Float64
			}
		case shipmentgatewayups.FieldCodAllowCod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cod_allow_cod", values[i])
			} else if value.Valid {
				sgu.CodAllowCod = value.String
			}
		case shipmentgatewayups.FieldCodSurchargeAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field cod_surcharge_amount", values[i])
			} else if value.Valid {
				sgu.CodSurchargeAmount = value.Float64
			}
		case shipmentgatewayups.FieldCodSurchargeCurrencyUomID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cod_surcharge_currency_uom_id", values[i])
			} else if value.Valid {
				sgu.CodSurchargeCurrencyUomID = value.String
			}
		case shipmentgatewayups.FieldCodSurchargeApplyToPackage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cod_surcharge_apply_to_package", values[i])
			} else if value.Valid {
				sgu.CodSurchargeApplyToPackage = value.String
			}
		case shipmentgatewayups.FieldCodFundsCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cod_funds_code", values[i])
			} else if value.Valid {
				sgu.CodFundsCode = value.String
			}
		case shipmentgatewayups.FieldDefaultReturnLabelMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_return_label_memo", values[i])
			} else if value.Valid {
				sgu.DefaultReturnLabelMemo = value.String
			}
		case shipmentgatewayups.FieldDefaultReturnLabelSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_return_label_subject", values[i])
			} else if value.Valid {
				sgu.DefaultReturnLabelSubject = value.String
			}
		case shipmentgatewayups.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shipment_gateway_config_shipment_gateway_ups", value)
			} else if value.Valid {
				sgu.shipment_gateway_config_shipment_gateway_ups = new(int)
				*sgu.shipment_gateway_config_shipment_gateway_ups = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryShipmentGatewayConfig queries the "shipment_gateway_config" edge of the ShipmentGatewayUps entity.
func (sgu *ShipmentGatewayUps) QueryShipmentGatewayConfig() *ShipmentGatewayConfigQuery {
	return (&ShipmentGatewayUpsClient{config: sgu.config}).QueryShipmentGatewayConfig(sgu)
}

// Update returns a builder for updating this ShipmentGatewayUps.
// Note that you need to call ShipmentGatewayUps.Unwrap() before calling this method if this ShipmentGatewayUps
// was returned from a transaction, and the transaction was committed or rolled back.
func (sgu *ShipmentGatewayUps) Update() *ShipmentGatewayUpsUpdateOne {
	return (&ShipmentGatewayUpsClient{config: sgu.config}).UpdateOne(sgu)
}

// Unwrap unwraps the ShipmentGatewayUps entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sgu *ShipmentGatewayUps) Unwrap() *ShipmentGatewayUps {
	tx, ok := sgu.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentGatewayUps is not a transactional entity")
	}
	sgu.config.driver = tx.drv
	return sgu
}

// String implements the fmt.Stringer.
func (sgu *ShipmentGatewayUps) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentGatewayUps(")
	builder.WriteString(fmt.Sprintf("id=%v", sgu.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(sgu.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(sgu.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(sgu.StringRef)
	builder.WriteString(", connect_url=")
	builder.WriteString(sgu.ConnectURL)
	builder.WriteString(", connect_timeout=")
	builder.WriteString(fmt.Sprintf("%v", sgu.ConnectTimeout))
	builder.WriteString(", shipper_number=")
	builder.WriteString(sgu.ShipperNumber)
	builder.WriteString(", bill_shipper_account_number=")
	builder.WriteString(sgu.BillShipperAccountNumber)
	builder.WriteString(", access_license_number=")
	builder.WriteString(sgu.AccessLicenseNumber)
	builder.WriteString(", access_user_id=")
	builder.WriteString(sgu.AccessUserID)
	builder.WriteString(", access_password=")
	builder.WriteString(sgu.AccessPassword)
	builder.WriteString(", save_cert_info=")
	builder.WriteString(sgu.SaveCertInfo)
	builder.WriteString(", save_cert_path=")
	builder.WriteString(sgu.SaveCertPath)
	builder.WriteString(", shipper_pickup_type=")
	builder.WriteString(sgu.ShipperPickupType)
	builder.WriteString(", customer_classification=")
	builder.WriteString(sgu.CustomerClassification)
	builder.WriteString(", max_estimate_weight=")
	builder.WriteString(fmt.Sprintf("%v", sgu.MaxEstimateWeight))
	builder.WriteString(", min_estimate_weight=")
	builder.WriteString(fmt.Sprintf("%v", sgu.MinEstimateWeight))
	builder.WriteString(", cod_allow_cod=")
	builder.WriteString(sgu.CodAllowCod)
	builder.WriteString(", cod_surcharge_amount=")
	builder.WriteString(fmt.Sprintf("%v", sgu.CodSurchargeAmount))
	builder.WriteString(", cod_surcharge_currency_uom_id=")
	builder.WriteString(sgu.CodSurchargeCurrencyUomID)
	builder.WriteString(", cod_surcharge_apply_to_package=")
	builder.WriteString(sgu.CodSurchargeApplyToPackage)
	builder.WriteString(", cod_funds_code=")
	builder.WriteString(sgu.CodFundsCode)
	builder.WriteString(", default_return_label_memo=")
	builder.WriteString(sgu.DefaultReturnLabelMemo)
	builder.WriteString(", default_return_label_subject=")
	builder.WriteString(sgu.DefaultReturnLabelSubject)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentGatewayUpsSlice is a parsable slice of ShipmentGatewayUps.
type ShipmentGatewayUpsSlice []*ShipmentGatewayUps

func (sgu ShipmentGatewayUpsSlice) config(cfg config) {
	for _i := range sgu {
		sgu[_i].config = cfg
	}
}
