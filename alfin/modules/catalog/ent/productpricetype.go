// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productpricetype"
)

// ProductPriceType is the model entity for the ProductPriceType schema.
type ProductPriceType struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductPriceTypeQuery when eager-loading is set.
	Edges ProductPriceTypeEdges `json:"edges"`
}

// ProductPriceTypeEdges holds the relations/edges for other nodes in the graph.
type ProductPriceTypeEdges struct {
	// ProductPrices holds the value of the product_prices edge.
	ProductPrices []*ProductPrice `json:"product_prices,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductPricesOrErr returns the ProductPrices value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPriceTypeEdges) ProductPricesOrErr() ([]*ProductPrice, error) {
	if e.loadedTypes[0] {
		return e.ProductPrices, nil
	}
	return nil, &NotLoadedError{edge: "product_prices"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductPriceType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productpricetype.FieldID:
			values[i] = new(sql.NullInt64)
		case productpricetype.FieldStringRef, productpricetype.FieldDescription:
			values[i] = new(sql.NullString)
		case productpricetype.FieldCreateTime, productpricetype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductPriceType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductPriceType fields.
func (ppt *ProductPriceType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productpricetype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ppt.ID = int(value.Int64)
		case productpricetype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ppt.CreateTime = value.Time
			}
		case productpricetype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ppt.UpdateTime = value.Time
			}
		case productpricetype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				ppt.StringRef = value.String
			}
		case productpricetype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ppt.Description = value.String
			}
		}
	}
	return nil
}

// QueryProductPrices queries the "product_prices" edge of the ProductPriceType entity.
func (ppt *ProductPriceType) QueryProductPrices() *ProductPriceQuery {
	return (&ProductPriceTypeClient{config: ppt.config}).QueryProductPrices(ppt)
}

// Update returns a builder for updating this ProductPriceType.
// Note that you need to call ProductPriceType.Unwrap() before calling this method if this ProductPriceType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ppt *ProductPriceType) Update() *ProductPriceTypeUpdateOne {
	return (&ProductPriceTypeClient{config: ppt.config}).UpdateOne(ppt)
}

// Unwrap unwraps the ProductPriceType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ppt *ProductPriceType) Unwrap() *ProductPriceType {
	tx, ok := ppt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductPriceType is not a transactional entity")
	}
	ppt.config.driver = tx.drv
	return ppt
}

// String implements the fmt.Stringer.
func (ppt *ProductPriceType) String() string {
	var builder strings.Builder
	builder.WriteString("ProductPriceType(")
	builder.WriteString(fmt.Sprintf("id=%v", ppt.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ppt.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ppt.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(ppt.StringRef)
	builder.WriteString(", description=")
	builder.WriteString(ppt.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ProductPriceTypes is a parsable slice of ProductPriceType.
type ProductPriceTypes []*ProductPriceType

func (ppt ProductPriceTypes) config(cfg config) {
	for _i := range ppt {
		ppt[_i].config = cfg
	}
}
