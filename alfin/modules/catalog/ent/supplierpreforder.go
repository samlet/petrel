// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/supplierpreforder"
)

// SupplierPrefOrder is the model entity for the SupplierPrefOrder schema.
type SupplierPrefOrder struct {
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
func (*SupplierPrefOrder) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case supplierpreforder.FieldID:
			values[i] = new(sql.NullInt64)
		case supplierpreforder.FieldStringRef, supplierpreforder.FieldDescription:
			values[i] = new(sql.NullString)
		case supplierpreforder.FieldCreateTime, supplierpreforder.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SupplierPrefOrder", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SupplierPrefOrder fields.
func (spo *SupplierPrefOrder) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case supplierpreforder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			spo.ID = int(value.Int64)
		case supplierpreforder.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				spo.CreateTime = value.Time
			}
		case supplierpreforder.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				spo.UpdateTime = value.Time
			}
		case supplierpreforder.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				spo.StringRef = value.String
			}
		case supplierpreforder.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				spo.Description = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SupplierPrefOrder.
// Note that you need to call SupplierPrefOrder.Unwrap() before calling this method if this SupplierPrefOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (spo *SupplierPrefOrder) Update() *SupplierPrefOrderUpdateOne {
	return (&SupplierPrefOrderClient{config: spo.config}).UpdateOne(spo)
}

// Unwrap unwraps the SupplierPrefOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (spo *SupplierPrefOrder) Unwrap() *SupplierPrefOrder {
	tx, ok := spo.config.driver.(*txDriver)
	if !ok {
		panic("ent: SupplierPrefOrder is not a transactional entity")
	}
	spo.config.driver = tx.drv
	return spo
}

// String implements the fmt.Stringer.
func (spo *SupplierPrefOrder) String() string {
	var builder strings.Builder
	builder.WriteString("SupplierPrefOrder(")
	builder.WriteString(fmt.Sprintf("id=%v", spo.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(spo.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(spo.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(spo.StringRef)
	builder.WriteString(", description=")
	builder.WriteString(spo.Description)
	builder.WriteByte(')')
	return builder.String()
}

// SupplierPrefOrders is a parsable slice of SupplierPrefOrder.
type SupplierPrefOrders []*SupplierPrefOrder

func (spo SupplierPrefOrders) config(cfg config) {
	for _i := range spo {
		spo[_i].config = cfg
	}
}