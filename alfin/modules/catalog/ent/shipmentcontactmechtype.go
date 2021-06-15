// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmentcontactmechtype"
)

// ShipmentContactMechType is the model entity for the ShipmentContactMechType schema.
type ShipmentContactMechType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentContactMechType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmentcontactmechtype.FieldID:
			values[i] = new(sql.NullInt64)
		case shipmentcontactmechtype.FieldStringRef, shipmentcontactmechtype.FieldDescription:
			values[i] = new(sql.NullString)
		case shipmentcontactmechtype.FieldCreateTime, shipmentcontactmechtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShipmentContactMechType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentContactMechType fields.
func (scmt *ShipmentContactMechType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmentcontactmechtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			scmt.ID = int(value.Int64)
		case shipmentcontactmechtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				scmt.CreateTime = value.Time
			}
		case shipmentcontactmechtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				scmt.UpdateTime = value.Time
			}
		case shipmentcontactmechtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				scmt.StringRef = value.String
			}
		case shipmentcontactmechtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				scmt.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ShipmentContactMechType.
// Note that you need to call ShipmentContactMechType.Unwrap() before calling this method if this ShipmentContactMechType
// was returned from a transaction, and the transaction was committed or rolled back.
func (scmt *ShipmentContactMechType) Update() *ShipmentContactMechTypeUpdateOne {
	return (&ShipmentContactMechTypeClient{config: scmt.config}).UpdateOne(scmt)
}

// Unwrap unwraps the ShipmentContactMechType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (scmt *ShipmentContactMechType) Unwrap() *ShipmentContactMechType {
	tx, ok := scmt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentContactMechType is not a transactional entity")
	}
	scmt.config.driver = tx.drv
	return scmt
}

// String implements the fmt.Stringer.
func (scmt *ShipmentContactMechType) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentContactMechType(")
	builder.WriteString(fmt.Sprintf("id=%v", scmt.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(scmt.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(scmt.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(scmt.StringRef)
	builder.WriteString(", description=")
	builder.WriteString(scmt.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentContactMechTypes is a parsable slice of ShipmentContactMechType.
type ShipmentContactMechTypes []*ShipmentContactMechType

func (scmt ShipmentContactMechTypes) config(cfg config) {
	for _i := range scmt {
		scmt[_i].config = cfg
	}
}
