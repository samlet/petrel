// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workload"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeature"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappl"
	"github.com/samlet/petrel/alfin/modules/workload/ent/workloadfeatureappltype"
)

// WorkloadFeatureAppl is the model entity for the WorkloadFeatureAppl schema.
type WorkloadFeatureAppl struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// FromDate holds the value of the "from_date" field.
	FromDate time.Time `json:"from_date,omitempty"`
	// ThruDate holds the value of the "thru_date" field.
	ThruDate time.Time `json:"thru_date,omitempty"`
	// SequenceNum holds the value of the "sequence_num" field.
	SequenceNum int `json:"sequence_num,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkloadFeatureApplQuery when eager-loading is set.
	Edges                                             WorkloadFeatureApplEdges `json:"edges"`
	workload_workload_feature_appls                   *int
	workload_feature_workload_feature_appls           *int
	workload_feature_appl_type_workload_feature_appls *int
}

// WorkloadFeatureApplEdges holds the relations/edges for other nodes in the graph.
type WorkloadFeatureApplEdges struct {
	// Workload holds the value of the workload edge.
	Workload *Workload `json:"workload,omitempty"`
	// WorkloadFeature holds the value of the workload_feature edge.
	WorkloadFeature *WorkloadFeature `json:"workload_feature,omitempty"`
	// WorkloadFeatureApplType holds the value of the workload_feature_appl_type edge.
	WorkloadFeatureApplType *WorkloadFeatureApplType `json:"workload_feature_appl_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// WorkloadOrErr returns the Workload value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkloadFeatureApplEdges) WorkloadOrErr() (*Workload, error) {
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

// WorkloadFeatureOrErr returns the WorkloadFeature value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkloadFeatureApplEdges) WorkloadFeatureOrErr() (*WorkloadFeature, error) {
	if e.loadedTypes[1] {
		if e.WorkloadFeature == nil {
			// The edge workload_feature was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workloadfeature.Label}
		}
		return e.WorkloadFeature, nil
	}
	return nil, &NotLoadedError{edge: "workload_feature"}
}

// WorkloadFeatureApplTypeOrErr returns the WorkloadFeatureApplType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkloadFeatureApplEdges) WorkloadFeatureApplTypeOrErr() (*WorkloadFeatureApplType, error) {
	if e.loadedTypes[2] {
		if e.WorkloadFeatureApplType == nil {
			// The edge workload_feature_appl_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workloadfeatureappltype.Label}
		}
		return e.WorkloadFeatureApplType, nil
	}
	return nil, &NotLoadedError{edge: "workload_feature_appl_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkloadFeatureAppl) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workloadfeatureappl.FieldID, workloadfeatureappl.FieldSequenceNum:
			values[i] = new(sql.NullInt64)
		case workloadfeatureappl.FieldCreateTime, workloadfeatureappl.FieldUpdateTime, workloadfeatureappl.FieldFromDate, workloadfeatureappl.FieldThruDate:
			values[i] = new(sql.NullTime)
		case workloadfeatureappl.ForeignKeys[0]: // workload_workload_feature_appls
			values[i] = new(sql.NullInt64)
		case workloadfeatureappl.ForeignKeys[1]: // workload_feature_workload_feature_appls
			values[i] = new(sql.NullInt64)
		case workloadfeatureappl.ForeignKeys[2]: // workload_feature_appl_type_workload_feature_appls
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkloadFeatureAppl", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkloadFeatureAppl fields.
func (wfa *WorkloadFeatureAppl) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workloadfeatureappl.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wfa.ID = int(value.Int64)
		case workloadfeatureappl.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wfa.CreateTime = value.Time
			}
		case workloadfeatureappl.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wfa.UpdateTime = value.Time
			}
		case workloadfeatureappl.FieldFromDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from_date", values[i])
			} else if value.Valid {
				wfa.FromDate = value.Time
			}
		case workloadfeatureappl.FieldThruDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field thru_date", values[i])
			} else if value.Valid {
				wfa.ThruDate = value.Time
			}
		case workloadfeatureappl.FieldSequenceNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence_num", values[i])
			} else if value.Valid {
				wfa.SequenceNum = int(value.Int64)
			}
		case workloadfeatureappl.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workload_workload_feature_appls", value)
			} else if value.Valid {
				wfa.workload_workload_feature_appls = new(int)
				*wfa.workload_workload_feature_appls = int(value.Int64)
			}
		case workloadfeatureappl.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workload_feature_workload_feature_appls", value)
			} else if value.Valid {
				wfa.workload_feature_workload_feature_appls = new(int)
				*wfa.workload_feature_workload_feature_appls = int(value.Int64)
			}
		case workloadfeatureappl.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field workload_feature_appl_type_workload_feature_appls", value)
			} else if value.Valid {
				wfa.workload_feature_appl_type_workload_feature_appls = new(int)
				*wfa.workload_feature_appl_type_workload_feature_appls = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWorkload queries the "workload" edge of the WorkloadFeatureAppl entity.
func (wfa *WorkloadFeatureAppl) QueryWorkload() *WorkloadQuery {
	return (&WorkloadFeatureApplClient{config: wfa.config}).QueryWorkload(wfa)
}

// QueryWorkloadFeature queries the "workload_feature" edge of the WorkloadFeatureAppl entity.
func (wfa *WorkloadFeatureAppl) QueryWorkloadFeature() *WorkloadFeatureQuery {
	return (&WorkloadFeatureApplClient{config: wfa.config}).QueryWorkloadFeature(wfa)
}

// QueryWorkloadFeatureApplType queries the "workload_feature_appl_type" edge of the WorkloadFeatureAppl entity.
func (wfa *WorkloadFeatureAppl) QueryWorkloadFeatureApplType() *WorkloadFeatureApplTypeQuery {
	return (&WorkloadFeatureApplClient{config: wfa.config}).QueryWorkloadFeatureApplType(wfa)
}

// Update returns a builder for updating this WorkloadFeatureAppl.
// Note that you need to call WorkloadFeatureAppl.Unwrap() before calling this method if this WorkloadFeatureAppl
// was returned from a transaction, and the transaction was committed or rolled back.
func (wfa *WorkloadFeatureAppl) Update() *WorkloadFeatureApplUpdateOne {
	return (&WorkloadFeatureApplClient{config: wfa.config}).UpdateOne(wfa)
}

// Unwrap unwraps the WorkloadFeatureAppl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wfa *WorkloadFeatureAppl) Unwrap() *WorkloadFeatureAppl {
	tx, ok := wfa.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkloadFeatureAppl is not a transactional entity")
	}
	wfa.config.driver = tx.drv
	return wfa
}

// String implements the fmt.Stringer.
func (wfa *WorkloadFeatureAppl) String() string {
	var builder strings.Builder
	builder.WriteString("WorkloadFeatureAppl(")
	builder.WriteString(fmt.Sprintf("id=%v", wfa.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(wfa.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(wfa.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", from_date=")
	builder.WriteString(wfa.FromDate.Format(time.ANSIC))
	builder.WriteString(", thru_date=")
	builder.WriteString(wfa.ThruDate.Format(time.ANSIC))
	builder.WriteString(", sequence_num=")
	builder.WriteString(fmt.Sprintf("%v", wfa.SequenceNum))
	builder.WriteByte(')')
	return builder.String()
}

// WorkloadFeatureAppls is a parsable slice of WorkloadFeatureAppl.
type WorkloadFeatureAppls []*WorkloadFeatureAppl

func (wfa WorkloadFeatureAppls) config(cfg config) {
	for _i := range wfa {
		wfa[_i].config = cfg
	}
}