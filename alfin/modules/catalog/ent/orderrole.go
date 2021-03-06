// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderrole"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/roletype"
)

// OrderRole is the model entity for the OrderRole schema.
type OrderRole struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// PartyID holds the value of the "party_id" field.
	PartyID int `json:"party_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderRoleQuery when eager-loading is set.
	Edges                    OrderRoleEdges `json:"edges"`
	order_header_order_roles *int
	role_type_order_roles    *int
}

// OrderRoleEdges holds the relations/edges for other nodes in the graph.
type OrderRoleEdges struct {
	// OrderHeader holds the value of the order_header edge.
	OrderHeader *OrderHeader `json:"order_header,omitempty"`
	// RoleType holds the value of the role_type edge.
	RoleType *RoleType `json:"role_type,omitempty"`
	// OrderItems holds the value of the order_items edge.
	OrderItems []*OrderItem `json:"order_items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OrderHeaderOrErr returns the OrderHeader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderRoleEdges) OrderHeaderOrErr() (*OrderHeader, error) {
	if e.loadedTypes[0] {
		if e.OrderHeader == nil {
			// The edge order_header was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orderheader.Label}
		}
		return e.OrderHeader, nil
	}
	return nil, &NotLoadedError{edge: "order_header"}
}

// RoleTypeOrErr returns the RoleType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderRoleEdges) RoleTypeOrErr() (*RoleType, error) {
	if e.loadedTypes[1] {
		if e.RoleType == nil {
			// The edge role_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roletype.Label}
		}
		return e.RoleType, nil
	}
	return nil, &NotLoadedError{edge: "role_type"}
}

// OrderItemsOrErr returns the OrderItems value or an error if the edge
// was not loaded in eager-loading.
func (e OrderRoleEdges) OrderItemsOrErr() ([]*OrderItem, error) {
	if e.loadedTypes[2] {
		return e.OrderItems, nil
	}
	return nil, &NotLoadedError{edge: "order_items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderRole) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderrole.FieldID, orderrole.FieldPartyID:
			values[i] = new(sql.NullInt64)
		case orderrole.FieldStringRef:
			values[i] = new(sql.NullString)
		case orderrole.FieldCreateTime, orderrole.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case orderrole.ForeignKeys[0]: // order_header_order_roles
			values[i] = new(sql.NullInt64)
		case orderrole.ForeignKeys[1]: // role_type_order_roles
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderRole", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderRole fields.
func (or *OrderRole) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case orderrole.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				or.CreateTime = value.Time
			}
		case orderrole.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				or.UpdateTime = value.Time
			}
		case orderrole.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				or.StringRef = value.String
			}
		case orderrole.FieldPartyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field party_id", values[i])
			} else if value.Valid {
				or.PartyID = int(value.Int64)
			}
		case orderrole.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_header_order_roles", value)
			} else if value.Valid {
				or.order_header_order_roles = new(int)
				*or.order_header_order_roles = int(value.Int64)
			}
		case orderrole.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_type_order_roles", value)
			} else if value.Valid {
				or.role_type_order_roles = new(int)
				*or.role_type_order_roles = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrderHeader queries the "order_header" edge of the OrderRole entity.
func (or *OrderRole) QueryOrderHeader() *OrderHeaderQuery {
	return (&OrderRoleClient{config: or.config}).QueryOrderHeader(or)
}

// QueryRoleType queries the "role_type" edge of the OrderRole entity.
func (or *OrderRole) QueryRoleType() *RoleTypeQuery {
	return (&OrderRoleClient{config: or.config}).QueryRoleType(or)
}

// QueryOrderItems queries the "order_items" edge of the OrderRole entity.
func (or *OrderRole) QueryOrderItems() *OrderItemQuery {
	return (&OrderRoleClient{config: or.config}).QueryOrderItems(or)
}

// Update returns a builder for updating this OrderRole.
// Note that you need to call OrderRole.Unwrap() before calling this method if this OrderRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OrderRole) Update() *OrderRoleUpdateOne {
	return (&OrderRoleClient{config: or.config}).UpdateOne(or)
}

// Unwrap unwraps the OrderRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OrderRole) Unwrap() *OrderRole {
	tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderRole is not a transactional entity")
	}
	or.config.driver = tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OrderRole) String() string {
	var builder strings.Builder
	builder.WriteString("OrderRole(")
	builder.WriteString(fmt.Sprintf("id=%v", or.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(or.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(or.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(or.StringRef)
	builder.WriteString(", party_id=")
	builder.WriteString(fmt.Sprintf("%v", or.PartyID))
	builder.WriteByte(')')
	return builder.String()
}

// OrderRoles is a parsable slice of OrderRole.
type OrderRoles []*OrderRole

func (or OrderRoles) config(cfg config) {
	for _i := range or {
		or[_i].config = cfg
	}
}
