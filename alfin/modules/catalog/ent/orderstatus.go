// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderheader"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderitem"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderpaymentpreference"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/orderstatus"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/statusitem"
)

// OrderStatus is the model entity for the OrderStatus schema.
type OrderStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// OrderItemSeqID holds the value of the "order_item_seq_id" field.
	OrderItemSeqID int `json:"order_item_seq_id,omitempty"`
	// StatusDatetime holds the value of the "status_datetime" field.
	StatusDatetime time.Time `json:"status_datetime,omitempty"`
	// StatusUserLogin holds the value of the "status_user_login" field.
	StatusUserLogin string `json:"status_user_login,omitempty"`
	// ChangeReason holds the value of the "change_reason" field.
	ChangeReason string `json:"change_reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderStatusQuery when eager-loading is set.
	Edges                                   OrderStatusEdges `json:"edges"`
	order_header_order_statuses             *int
	order_item_order_statuses               *int
	order_payment_preference_order_statuses *int
	status_item_order_statuses              *int
}

// OrderStatusEdges holds the relations/edges for other nodes in the graph.
type OrderStatusEdges struct {
	// StatusItem holds the value of the status_item edge.
	StatusItem *StatusItem `json:"status_item,omitempty"`
	// OrderHeader holds the value of the order_header edge.
	OrderHeader *OrderHeader `json:"order_header,omitempty"`
	// OrderItem holds the value of the order_item edge.
	OrderItem *OrderItem `json:"order_item,omitempty"`
	// OrderPaymentPreference holds the value of the order_payment_preference edge.
	OrderPaymentPreference *OrderPaymentPreference `json:"order_payment_preference,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// StatusItemOrErr returns the StatusItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderStatusEdges) StatusItemOrErr() (*StatusItem, error) {
	if e.loadedTypes[0] {
		if e.StatusItem == nil {
			// The edge status_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statusitem.Label}
		}
		return e.StatusItem, nil
	}
	return nil, &NotLoadedError{edge: "status_item"}
}

// OrderHeaderOrErr returns the OrderHeader value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderStatusEdges) OrderHeaderOrErr() (*OrderHeader, error) {
	if e.loadedTypes[1] {
		if e.OrderHeader == nil {
			// The edge order_header was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orderheader.Label}
		}
		return e.OrderHeader, nil
	}
	return nil, &NotLoadedError{edge: "order_header"}
}

// OrderItemOrErr returns the OrderItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderStatusEdges) OrderItemOrErr() (*OrderItem, error) {
	if e.loadedTypes[2] {
		if e.OrderItem == nil {
			// The edge order_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orderitem.Label}
		}
		return e.OrderItem, nil
	}
	return nil, &NotLoadedError{edge: "order_item"}
}

// OrderPaymentPreferenceOrErr returns the OrderPaymentPreference value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderStatusEdges) OrderPaymentPreferenceOrErr() (*OrderPaymentPreference, error) {
	if e.loadedTypes[3] {
		if e.OrderPaymentPreference == nil {
			// The edge order_payment_preference was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orderpaymentpreference.Label}
		}
		return e.OrderPaymentPreference, nil
	}
	return nil, &NotLoadedError{edge: "order_payment_preference"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderStatus) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderstatus.FieldID, orderstatus.FieldOrderItemSeqID:
			values[i] = new(sql.NullInt64)
		case orderstatus.FieldStringRef, orderstatus.FieldStatusUserLogin, orderstatus.FieldChangeReason:
			values[i] = new(sql.NullString)
		case orderstatus.FieldCreateTime, orderstatus.FieldUpdateTime, orderstatus.FieldStatusDatetime:
			values[i] = new(sql.NullTime)
		case orderstatus.ForeignKeys[0]: // order_header_order_statuses
			values[i] = new(sql.NullInt64)
		case orderstatus.ForeignKeys[1]: // order_item_order_statuses
			values[i] = new(sql.NullInt64)
		case orderstatus.ForeignKeys[2]: // order_payment_preference_order_statuses
			values[i] = new(sql.NullInt64)
		case orderstatus.ForeignKeys[3]: // status_item_order_statuses
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderStatus fields.
func (os *OrderStatus) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderstatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int(value.Int64)
		case orderstatus.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				os.CreateTime = value.Time
			}
		case orderstatus.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				os.UpdateTime = value.Time
			}
		case orderstatus.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				os.StringRef = value.String
			}
		case orderstatus.FieldOrderItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_item_seq_id", values[i])
			} else if value.Valid {
				os.OrderItemSeqID = int(value.Int64)
			}
		case orderstatus.FieldStatusDatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field status_datetime", values[i])
			} else if value.Valid {
				os.StatusDatetime = value.Time
			}
		case orderstatus.FieldStatusUserLogin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status_user_login", values[i])
			} else if value.Valid {
				os.StatusUserLogin = value.String
			}
		case orderstatus.FieldChangeReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_reason", values[i])
			} else if value.Valid {
				os.ChangeReason = value.String
			}
		case orderstatus.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_header_order_statuses", value)
			} else if value.Valid {
				os.order_header_order_statuses = new(int)
				*os.order_header_order_statuses = int(value.Int64)
			}
		case orderstatus.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_item_order_statuses", value)
			} else if value.Valid {
				os.order_item_order_statuses = new(int)
				*os.order_item_order_statuses = int(value.Int64)
			}
		case orderstatus.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field order_payment_preference_order_statuses", value)
			} else if value.Valid {
				os.order_payment_preference_order_statuses = new(int)
				*os.order_payment_preference_order_statuses = int(value.Int64)
			}
		case orderstatus.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field status_item_order_statuses", value)
			} else if value.Valid {
				os.status_item_order_statuses = new(int)
				*os.status_item_order_statuses = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryStatusItem queries the "status_item" edge of the OrderStatus entity.
func (os *OrderStatus) QueryStatusItem() *StatusItemQuery {
	return (&OrderStatusClient{config: os.config}).QueryStatusItem(os)
}

// QueryOrderHeader queries the "order_header" edge of the OrderStatus entity.
func (os *OrderStatus) QueryOrderHeader() *OrderHeaderQuery {
	return (&OrderStatusClient{config: os.config}).QueryOrderHeader(os)
}

// QueryOrderItem queries the "order_item" edge of the OrderStatus entity.
func (os *OrderStatus) QueryOrderItem() *OrderItemQuery {
	return (&OrderStatusClient{config: os.config}).QueryOrderItem(os)
}

// QueryOrderPaymentPreference queries the "order_payment_preference" edge of the OrderStatus entity.
func (os *OrderStatus) QueryOrderPaymentPreference() *OrderPaymentPreferenceQuery {
	return (&OrderStatusClient{config: os.config}).QueryOrderPaymentPreference(os)
}

// Update returns a builder for updating this OrderStatus.
// Note that you need to call OrderStatus.Unwrap() before calling this method if this OrderStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OrderStatus) Update() *OrderStatusUpdateOne {
	return (&OrderStatusClient{config: os.config}).UpdateOne(os)
}

// Unwrap unwraps the OrderStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OrderStatus) Unwrap() *OrderStatus {
	tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderStatus is not a transactional entity")
	}
	os.config.driver = tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OrderStatus) String() string {
	var builder strings.Builder
	builder.WriteString("OrderStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", os.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(os.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(os.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(os.StringRef)
	builder.WriteString(", order_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", os.OrderItemSeqID))
	builder.WriteString(", status_datetime=")
	builder.WriteString(os.StatusDatetime.Format(time.ANSIC))
	builder.WriteString(", status_user_login=")
	builder.WriteString(os.StatusUserLogin)
	builder.WriteString(", change_reason=")
	builder.WriteString(os.ChangeReason)
	builder.WriteByte(')')
	return builder.String()
}

// OrderStatusSlice is a parsable slice of OrderStatus.
type OrderStatusSlice []*OrderStatus

func (os OrderStatusSlice) config(cfg config) {
	for _i := range os {
		os[_i].config = cfg
	}
}
