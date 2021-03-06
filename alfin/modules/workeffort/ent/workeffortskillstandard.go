// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/skilltype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffort"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workeffortskillstandard"
)

// WorkEffortSkillStandard is the model entity for the WorkEffortSkillStandard schema.
type WorkEffortSkillStandard struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// EstimatedNumPeople holds the value of the "estimated_num_people" field.
	EstimatedNumPeople float64 `json:"estimated_num_people,omitempty"`
	// EstimatedDuration holds the value of the "estimated_duration" field.
	EstimatedDuration float64 `json:"estimated_duration,omitempty"`
	// EstimatedCost holds the value of the "estimated_cost" field.
	EstimatedCost float64 `json:"estimated_cost,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkEffortSkillStandardQuery when eager-loading is set.
	Edges                                   WorkEffortSkillStandardEdges `json:"edges"`
	skill_type_work_effort_skill_standards  *int
	work_effort_work_effort_skill_standards *int
}

// WorkEffortSkillStandardEdges holds the relations/edges for other nodes in the graph.
type WorkEffortSkillStandardEdges struct {
	// WorkEffort holds the value of the work_effort edge.
	WorkEffort *WorkEffort `json:"work_effort,omitempty"`
	// SkillType holds the value of the skill_type edge.
	SkillType *SkillType `json:"skill_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// WorkEffortOrErr returns the WorkEffort value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEffortSkillStandardEdges) WorkEffortOrErr() (*WorkEffort, error) {
	if e.loadedTypes[0] {
		if e.WorkEffort == nil {
			// The edge work_effort was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workeffort.Label}
		}
		return e.WorkEffort, nil
	}
	return nil, &NotLoadedError{edge: "work_effort"}
}

// SkillTypeOrErr returns the SkillType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEffortSkillStandardEdges) SkillTypeOrErr() (*SkillType, error) {
	if e.loadedTypes[1] {
		if e.SkillType == nil {
			// The edge skill_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: skilltype.Label}
		}
		return e.SkillType, nil
	}
	return nil, &NotLoadedError{edge: "skill_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkEffortSkillStandard) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workeffortskillstandard.FieldEstimatedNumPeople, workeffortskillstandard.FieldEstimatedDuration, workeffortskillstandard.FieldEstimatedCost:
			values[i] = new(sql.NullFloat64)
		case workeffortskillstandard.FieldID:
			values[i] = new(sql.NullInt64)
		case workeffortskillstandard.FieldStringRef:
			values[i] = new(sql.NullString)
		case workeffortskillstandard.FieldCreateTime, workeffortskillstandard.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case workeffortskillstandard.ForeignKeys[0]: // skill_type_work_effort_skill_standards
			values[i] = new(sql.NullInt64)
		case workeffortskillstandard.ForeignKeys[1]: // work_effort_work_effort_skill_standards
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkEffortSkillStandard", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkEffortSkillStandard fields.
func (wess *WorkEffortSkillStandard) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workeffortskillstandard.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wess.ID = int(value.Int64)
		case workeffortskillstandard.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wess.CreateTime = value.Time
			}
		case workeffortskillstandard.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wess.UpdateTime = value.Time
			}
		case workeffortskillstandard.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				wess.StringRef = value.String
			}
		case workeffortskillstandard.FieldEstimatedNumPeople:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_num_people", values[i])
			} else if value.Valid {
				wess.EstimatedNumPeople = value.Float64
			}
		case workeffortskillstandard.FieldEstimatedDuration:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_duration", values[i])
			} else if value.Valid {
				wess.EstimatedDuration = value.Float64
			}
		case workeffortskillstandard.FieldEstimatedCost:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field estimated_cost", values[i])
			} else if value.Valid {
				wess.EstimatedCost = value.Float64
			}
		case workeffortskillstandard.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field skill_type_work_effort_skill_standards", value)
			} else if value.Valid {
				wess.skill_type_work_effort_skill_standards = new(int)
				*wess.skill_type_work_effort_skill_standards = int(value.Int64)
			}
		case workeffortskillstandard.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field work_effort_work_effort_skill_standards", value)
			} else if value.Valid {
				wess.work_effort_work_effort_skill_standards = new(int)
				*wess.work_effort_work_effort_skill_standards = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryWorkEffort queries the "work_effort" edge of the WorkEffortSkillStandard entity.
func (wess *WorkEffortSkillStandard) QueryWorkEffort() *WorkEffortQuery {
	return (&WorkEffortSkillStandardClient{config: wess.config}).QueryWorkEffort(wess)
}

// QuerySkillType queries the "skill_type" edge of the WorkEffortSkillStandard entity.
func (wess *WorkEffortSkillStandard) QuerySkillType() *SkillTypeQuery {
	return (&WorkEffortSkillStandardClient{config: wess.config}).QuerySkillType(wess)
}

// Update returns a builder for updating this WorkEffortSkillStandard.
// Note that you need to call WorkEffortSkillStandard.Unwrap() before calling this method if this WorkEffortSkillStandard
// was returned from a transaction, and the transaction was committed or rolled back.
func (wess *WorkEffortSkillStandard) Update() *WorkEffortSkillStandardUpdateOne {
	return (&WorkEffortSkillStandardClient{config: wess.config}).UpdateOne(wess)
}

// Unwrap unwraps the WorkEffortSkillStandard entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wess *WorkEffortSkillStandard) Unwrap() *WorkEffortSkillStandard {
	tx, ok := wess.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkEffortSkillStandard is not a transactional entity")
	}
	wess.config.driver = tx.drv
	return wess
}

// String implements the fmt.Stringer.
func (wess *WorkEffortSkillStandard) String() string {
	var builder strings.Builder
	builder.WriteString("WorkEffortSkillStandard(")
	builder.WriteString(fmt.Sprintf("id=%v", wess.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(wess.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(wess.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(wess.StringRef)
	builder.WriteString(", estimated_num_people=")
	builder.WriteString(fmt.Sprintf("%v", wess.EstimatedNumPeople))
	builder.WriteString(", estimated_duration=")
	builder.WriteString(fmt.Sprintf("%v", wess.EstimatedDuration))
	builder.WriteString(", estimated_cost=")
	builder.WriteString(fmt.Sprintf("%v", wess.EstimatedCost))
	builder.WriteByte(')')
	return builder.String()
}

// WorkEffortSkillStandards is a parsable slice of WorkEffortSkillStandard.
type WorkEffortSkillStandards []*WorkEffortSkillStandard

func (wess WorkEffortSkillStandards) config(cfg config) {
	for _i := range wess {
		wess[_i].config = cfg
	}
}
