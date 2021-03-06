// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpriceactiontype"
)

// ProductPriceActionType is the model entity for the ProductPriceActionType schema.
type ProductPriceActionType struct {
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
func (*ProductPriceActionType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productpriceactiontype.FieldID:
			values[i] = new(sql.NullInt64)
		case productpriceactiontype.FieldStringRef, productpriceactiontype.FieldDescription:
			values[i] = new(sql.NullString)
		case productpriceactiontype.FieldCreateTime, productpriceactiontype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductPriceActionType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductPriceActionType fields.
func (ppat *ProductPriceActionType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productpriceactiontype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ppat.ID = int(value.Int64)
		case productpriceactiontype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ppat.CreateTime = value.Time
			}
		case productpriceactiontype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ppat.UpdateTime = value.Time
			}
		case productpriceactiontype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				ppat.StringRef = value.String
			}
		case productpriceactiontype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ppat.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ProductPriceActionType.
// Note that you need to call ProductPriceActionType.Unwrap() before calling this method if this ProductPriceActionType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ppat *ProductPriceActionType) Update() *ProductPriceActionTypeUpdateOne {
	return (&ProductPriceActionTypeClient{config: ppat.config}).UpdateOne(ppat)
}

// Unwrap unwraps the ProductPriceActionType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ppat *ProductPriceActionType) Unwrap() *ProductPriceActionType {
	tx, ok := ppat.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductPriceActionType is not a transactional entity")
	}
	ppat.config.driver = tx.drv
	return ppat
}

// String implements the fmt.Stringer.
func (ppat *ProductPriceActionType) String() string {
	var builder strings.Builder
	builder.WriteString("ProductPriceActionType(")
	builder.WriteString(fmt.Sprintf("id=%v", ppat.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ppat.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ppat.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(ppat.StringRef)
	builder.WriteString(", description=")
	builder.WriteString(ppat.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ProductPriceActionTypes is a parsable slice of ProductPriceActionType.
type ProductPriceActionTypes []*ProductPriceActionType

func (ppat ProductPriceActionTypes) config(cfg config) {
	for _i := range ppat {
		ppat[_i].config = cfg
	}
}
