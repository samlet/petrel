// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/shipmenttype"
)

// ShipmentType is the model entity for the ShipmentType schema.
type ShipmentType struct {
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
	HasTable shipmenttype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShipmentTypeQuery when eager-loading is set.
	Edges                  ShipmentTypeEdges `json:"edges"`
	shipment_type_children *int
}

// ShipmentTypeEdges holds the relations/edges for other nodes in the graph.
type ShipmentTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ShipmentType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*ShipmentType `json:"children,omitempty"`
	// ChildShipmentTypes holds the value of the child_shipment_types edge.
	ChildShipmentTypes []*ShipmentType `json:"child_shipment_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShipmentTypeEdges) ParentOrErr() (*ShipmentType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shipmenttype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentTypeEdges) ChildrenOrErr() ([]*ShipmentType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ChildShipmentTypesOrErr returns the ChildShipmentTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ShipmentTypeEdges) ChildShipmentTypesOrErr() ([]*ShipmentType, error) {
	if e.loadedTypes[2] {
		return e.ChildShipmentTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_shipment_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShipmentType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case shipmenttype.FieldID:
			values[i] = new(sql.NullInt64)
		case shipmenttype.FieldStringRef, shipmenttype.FieldHasTable, shipmenttype.FieldDescription:
			values[i] = new(sql.NullString)
		case shipmenttype.FieldCreateTime, shipmenttype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case shipmenttype.ForeignKeys[0]: // shipment_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ShipmentType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShipmentType fields.
func (st *ShipmentType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shipmenttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			st.ID = int(value.Int64)
		case shipmenttype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				st.CreateTime = value.Time
			}
		case shipmenttype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				st.UpdateTime = value.Time
			}
		case shipmenttype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				st.StringRef = value.String
			}
		case shipmenttype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				st.HasTable = shipmenttype.HasTable(value.String)
			}
		case shipmenttype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				st.Description = value.String
			}
		case shipmenttype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field shipment_type_children", value)
			} else if value.Valid {
				st.shipment_type_children = new(int)
				*st.shipment_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ShipmentType entity.
func (st *ShipmentType) QueryParent() *ShipmentTypeQuery {
	return (&ShipmentTypeClient{config: st.config}).QueryParent(st)
}

// QueryChildren queries the "children" edge of the ShipmentType entity.
func (st *ShipmentType) QueryChildren() *ShipmentTypeQuery {
	return (&ShipmentTypeClient{config: st.config}).QueryChildren(st)
}

// QueryChildShipmentTypes queries the "child_shipment_types" edge of the ShipmentType entity.
func (st *ShipmentType) QueryChildShipmentTypes() *ShipmentTypeQuery {
	return (&ShipmentTypeClient{config: st.config}).QueryChildShipmentTypes(st)
}

// Update returns a builder for updating this ShipmentType.
// Note that you need to call ShipmentType.Unwrap() before calling this method if this ShipmentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *ShipmentType) Update() *ShipmentTypeUpdateOne {
	return (&ShipmentTypeClient{config: st.config}).UpdateOne(st)
}

// Unwrap unwraps the ShipmentType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *ShipmentType) Unwrap() *ShipmentType {
	tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShipmentType is not a transactional entity")
	}
	st.config.driver = tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *ShipmentType) String() string {
	var builder strings.Builder
	builder.WriteString("ShipmentType(")
	builder.WriteString(fmt.Sprintf("id=%v", st.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(st.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(st.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(st.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", st.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(st.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ShipmentTypes is a parsable slice of ShipmentType.
type ShipmentTypes []*ShipmentType

func (st ShipmentTypes) config(cfg config) {
	for _i := range st {
		st[_i].config = cfg
	}
}