// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloaditem"
)

// WorkloadItem is the model entity for the WorkloadItem schema.
type WorkloadItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// WorkloadItemSeqID holds the value of the "workload_item_seq_id" field.
	WorkloadItemSeqID int `json:"workload_item_seq_id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount,omitempty"`
	// AmountUomID holds the value of the "amount_uom_id" field.
	AmountUomID int `json:"amount_uom_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkloadItemQuery when eager-loading is set.
	Edges                   WorkloadItemEdges `json:"edges"`
	workload_workload_items *int
}

// WorkloadItemEdges holds the relations/edges for other nodes in the graph.
type WorkloadItemEdges struct {
	// Workload holds the value of the workload edge.
	Workload *Workload `json:"workload,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkloadOrErr returns the Workload value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkloadItemEdges) WorkloadOrErr() (*Workload, error) {
	if e.loadedTypes[0] {
		if e.Workload == nil {
			// The edge workload was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workload.Label}
		}
		return e.Workload, nil
	}
	return nil, &NotLoadedError{edge: "workload"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkloadItem) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workloaditem.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case workloaditem.FieldID, workloaditem.FieldWorkloadItemSeqID, workloaditem.FieldAmountUomID:
			values[i] = new(sql.NullInt64)
		case workloaditem.FieldDescription:
			values[i] = new(sql.NullString)
		case workloaditem.FieldCreateTime, workloaditem.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case workloaditem.ForeignKeys[0]: // workload_workload_items
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkloadItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkloadItem fields.
func (wi *WorkloadItem) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workloaditem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wi.ID = int(value.Int64)
		case workloaditem.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wi.CreateTime = value.Time
			}
		case workloaditem.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wi.UpdateTime = value.Time
			}
		case workloaditem.FieldWorkloadItemSeqID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field workload_item_seq_id", values[i])
			} else if value.Valid {
				wi.WorkloadItemSeqID = int(value.Int64)
			}
		case workloaditem.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				wi.Description = value.String
			}
		case workloaditem.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				wi.Amount = value.Float64
			}
		case workloaditem.FieldAmountUomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount_uom_id", values[i])
			} else if value.Valid {
				wi.AmountUomID = int(value.Int64)
			}
		case workloaditem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workload_workload_items", value)
			} else if value.Valid {
				wi.workload_workload_items = new(int)
				*wi.workload_workload_items = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWorkload queries the "workload" edge of the WorkloadItem entity.
func (wi *WorkloadItem) QueryWorkload() *WorkloadQuery {
	return (&WorkloadItemClient{config: wi.config}).QueryWorkload(wi)
}

// Update returns a builder for updating this WorkloadItem.
// Note that you need to call WorkloadItem.Unwrap() before calling this method if this WorkloadItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (wi *WorkloadItem) Update() *WorkloadItemUpdateOne {
	return (&WorkloadItemClient{config: wi.config}).UpdateOne(wi)
}

// Unwrap unwraps the WorkloadItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wi *WorkloadItem) Unwrap() *WorkloadItem {
	tx, ok := wi.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkloadItem is not a transactional entity")
	}
	wi.config.driver = tx.drv
	return wi
}

// String implements the fmt.Stringer.
func (wi *WorkloadItem) String() string {
	var builder strings.Builder
	builder.WriteString("WorkloadItem(")
	builder.WriteString(fmt.Sprintf("id=%v", wi.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(wi.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(wi.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", workload_item_seq_id=")
	builder.WriteString(fmt.Sprintf("%v", wi.WorkloadItemSeqID))
	builder.WriteString(", description=")
	builder.WriteString(wi.Description)
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", wi.Amount))
	builder.WriteString(", amount_uom_id=")
	builder.WriteString(fmt.Sprintf("%v", wi.AmountUomID))
	builder.WriteByte(')')
	return builder.String()
}

// WorkloadItems is a parsable slice of WorkloadItem.
type WorkloadItems []*WorkloadItem

func (wi WorkloadItems) config(cfg config) {
	for _i := range wi {
		wi[_i].config = cfg
	}
}
