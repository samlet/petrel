// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadstatus"
)

// WorkloadStatus is the model entity for the WorkloadStatus schema.
type WorkloadStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StatusDate holds the value of the "status_date" field.
	StatusDate time.Time `json:"status_date,omitempty"`
	// StatusEndDate holds the value of the "status_end_date" field.
	StatusEndDate time.Time `json:"status_end_date,omitempty"`
	// ChangeByUserLoginID holds the value of the "change_by_user_login_id" field.
	ChangeByUserLoginID string `json:"change_by_user_login_id,omitempty"`
	// StatusID holds the value of the "status_id" field.
	StatusID int `json:"status_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkloadStatusQuery when eager-loading is set.
	Edges                      WorkloadStatusEdges `json:"edges"`
	workload_workload_statuses *int
}

// WorkloadStatusEdges holds the relations/edges for other nodes in the graph.
type WorkloadStatusEdges struct {
	// Workload holds the value of the workload edge.
	Workload *Workload `json:"workload,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// WorkloadOrErr returns the Workload value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkloadStatusEdges) WorkloadOrErr() (*Workload, error) {
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
func (*WorkloadStatus) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workloadstatus.FieldID, workloadstatus.FieldStatusID:
			values[i] = new(sql.NullInt64)
		case workloadstatus.FieldChangeByUserLoginID:
			values[i] = new(sql.NullString)
		case workloadstatus.FieldCreateTime, workloadstatus.FieldUpdateTime, workloadstatus.FieldStatusDate, workloadstatus.FieldStatusEndDate:
			values[i] = new(sql.NullTime)
		case workloadstatus.ForeignKeys[0]: // workload_workload_statuses
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkloadStatus", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkloadStatus fields.
func (ws *WorkloadStatus) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workloadstatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ws.ID = int(value.Int64)
		case workloadstatus.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ws.CreateTime = value.Time
			}
		case workloadstatus.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ws.UpdateTime = value.Time
			}
		case workloadstatus.FieldStatusDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field status_date", values[i])
			} else if value.Valid {
				ws.StatusDate = value.Time
			}
		case workloadstatus.FieldStatusEndDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field status_end_date", values[i])
			} else if value.Valid {
				ws.StatusEndDate = value.Time
			}
		case workloadstatus.FieldChangeByUserLoginID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field change_by_user_login_id", values[i])
			} else if value.Valid {
				ws.ChangeByUserLoginID = value.String
			}
		case workloadstatus.FieldStatusID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status_id", values[i])
			} else if value.Valid {
				ws.StatusID = int(value.Int64)
			}
		case workloadstatus.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workload_workload_statuses", value)
			} else if value.Valid {
				ws.workload_workload_statuses = new(int)
				*ws.workload_workload_statuses = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWorkload queries the "workload" edge of the WorkloadStatus entity.
func (ws *WorkloadStatus) QueryWorkload() *WorkloadQuery {
	return (&WorkloadStatusClient{config: ws.config}).QueryWorkload(ws)
}

// Update returns a builder for updating this WorkloadStatus.
// Note that you need to call WorkloadStatus.Unwrap() before calling this method if this WorkloadStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (ws *WorkloadStatus) Update() *WorkloadStatusUpdateOne {
	return (&WorkloadStatusClient{config: ws.config}).UpdateOne(ws)
}

// Unwrap unwraps the WorkloadStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ws *WorkloadStatus) Unwrap() *WorkloadStatus {
	tx, ok := ws.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkloadStatus is not a transactional entity")
	}
	ws.config.driver = tx.drv
	return ws
}

// String implements the fmt.Stringer.
func (ws *WorkloadStatus) String() string {
	var builder strings.Builder
	builder.WriteString("WorkloadStatus(")
	builder.WriteString(fmt.Sprintf("id=%v", ws.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ws.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ws.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", status_date=")
	builder.WriteString(ws.StatusDate.Format(time.ANSIC))
	builder.WriteString(", status_end_date=")
	builder.WriteString(ws.StatusEndDate.Format(time.ANSIC))
	builder.WriteString(", change_by_user_login_id=")
	builder.WriteString(ws.ChangeByUserLoginID)
	builder.WriteString(", status_id=")
	builder.WriteString(fmt.Sprintf("%v", ws.StatusID))
	builder.WriteByte(')')
	return builder.String()
}

// WorkloadStatusSlice is a parsable slice of WorkloadStatus.
type WorkloadStatusSlice []*WorkloadStatus

func (ws WorkloadStatusSlice) config(cfg config) {
	for _i := range ws {
		ws[_i].config = cfg
	}
}