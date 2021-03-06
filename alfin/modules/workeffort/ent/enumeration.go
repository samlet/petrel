// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumeration"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/enumerationtype"
)

// Enumeration is the model entity for the Enumeration schema.
type Enumeration struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// EnumCode holds the value of the "enum_code" field.
	EnumCode string `json:"enum_code,omitempty"`
	// SequenceID holds the value of the "sequence_id" field.
	SequenceID int `json:"sequence_id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EnumerationQuery when eager-loading is set.
	Edges                         EnumerationEdges `json:"edges"`
	enumeration_type_enumerations *int
}

// EnumerationEdges holds the relations/edges for other nodes in the graph.
type EnumerationEdges struct {
	// EnumerationType holds the value of the enumeration_type edge.
	EnumerationType *EnumerationType `json:"enumeration_type,omitempty"`
	// ClassFixedAssets holds the value of the class_fixed_assets edge.
	ClassFixedAssets []*FixedAsset `json:"class_fixed_assets,omitempty"`
	// EmploymentStatusPeople holds the value of the employment_status_people edge.
	EmploymentStatusPeople []*Person `json:"employment_status_people,omitempty"`
	// ResidenceStatusPeople holds the value of the residence_status_people edge.
	ResidenceStatusPeople []*Person `json:"residence_status_people,omitempty"`
	// MaritalStatusPeople holds the value of the marital_status_people edge.
	MaritalStatusPeople []*Person `json:"marital_status_people,omitempty"`
	// ScopeWorkEfforts holds the value of the scope_work_efforts edge.
	ScopeWorkEfforts []*WorkEffort `json:"scope_work_efforts,omitempty"`
	// ExpectationWorkEffortPartyAssignments holds the value of the expectation_work_effort_party_assignments edge.
	ExpectationWorkEffortPartyAssignments []*WorkEffortPartyAssignment `json:"expectation_work_effort_party_assignments,omitempty"`
	// DelegateReasonWorkEffortPartyAssignments holds the value of the delegate_reason_work_effort_party_assignments edge.
	DelegateReasonWorkEffortPartyAssignments []*WorkEffortPartyAssignment `json:"delegate_reason_work_effort_party_assignments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
}

// EnumerationTypeOrErr returns the EnumerationType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EnumerationEdges) EnumerationTypeOrErr() (*EnumerationType, error) {
	if e.loadedTypes[0] {
		if e.EnumerationType == nil {
			// The edge enumeration_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: enumerationtype.Label}
		}
		return e.EnumerationType, nil
	}
	return nil, &NotLoadedError{edge: "enumeration_type"}
}

// ClassFixedAssetsOrErr returns the ClassFixedAssets value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) ClassFixedAssetsOrErr() ([]*FixedAsset, error) {
	if e.loadedTypes[1] {
		return e.ClassFixedAssets, nil
	}
	return nil, &NotLoadedError{edge: "class_fixed_assets"}
}

// EmploymentStatusPeopleOrErr returns the EmploymentStatusPeople value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) EmploymentStatusPeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[2] {
		return e.EmploymentStatusPeople, nil
	}
	return nil, &NotLoadedError{edge: "employment_status_people"}
}

// ResidenceStatusPeopleOrErr returns the ResidenceStatusPeople value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) ResidenceStatusPeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[3] {
		return e.ResidenceStatusPeople, nil
	}
	return nil, &NotLoadedError{edge: "residence_status_people"}
}

// MaritalStatusPeopleOrErr returns the MaritalStatusPeople value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) MaritalStatusPeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[4] {
		return e.MaritalStatusPeople, nil
	}
	return nil, &NotLoadedError{edge: "marital_status_people"}
}

// ScopeWorkEffortsOrErr returns the ScopeWorkEfforts value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) ScopeWorkEffortsOrErr() ([]*WorkEffort, error) {
	if e.loadedTypes[5] {
		return e.ScopeWorkEfforts, nil
	}
	return nil, &NotLoadedError{edge: "scope_work_efforts"}
}

// ExpectationWorkEffortPartyAssignmentsOrErr returns the ExpectationWorkEffortPartyAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) ExpectationWorkEffortPartyAssignmentsOrErr() ([]*WorkEffortPartyAssignment, error) {
	if e.loadedTypes[6] {
		return e.ExpectationWorkEffortPartyAssignments, nil
	}
	return nil, &NotLoadedError{edge: "expectation_work_effort_party_assignments"}
}

// DelegateReasonWorkEffortPartyAssignmentsOrErr returns the DelegateReasonWorkEffortPartyAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e EnumerationEdges) DelegateReasonWorkEffortPartyAssignmentsOrErr() ([]*WorkEffortPartyAssignment, error) {
	if e.loadedTypes[7] {
		return e.DelegateReasonWorkEffortPartyAssignments, nil
	}
	return nil, &NotLoadedError{edge: "delegate_reason_work_effort_party_assignments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Enumeration) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case enumeration.FieldID, enumeration.FieldSequenceID:
			values[i] = new(sql.NullInt64)
		case enumeration.FieldStringRef, enumeration.FieldEnumCode, enumeration.FieldDescription:
			values[i] = new(sql.NullString)
		case enumeration.FieldCreateTime, enumeration.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case enumeration.ForeignKeys[0]: // enumeration_type_enumerations
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Enumeration", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Enumeration fields.
func (e *Enumeration) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case enumeration.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case enumeration.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		case enumeration.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = value.Time
			}
		case enumeration.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				e.StringRef = value.String
			}
		case enumeration.FieldEnumCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field enum_code", values[i])
			} else if value.Valid {
				e.EnumCode = value.String
			}
		case enumeration.FieldSequenceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequence_id", values[i])
			} else if value.Valid {
				e.SequenceID = int(value.Int64)
			}
		case enumeration.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case enumeration.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field enumeration_type_enumerations", value)
			} else if value.Valid {
				e.enumeration_type_enumerations = new(int)
				*e.enumeration_type_enumerations = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEnumerationType queries the "enumeration_type" edge of the Enumeration entity.
func (e *Enumeration) QueryEnumerationType() *EnumerationTypeQuery {
	return (&EnumerationClient{config: e.config}).QueryEnumerationType(e)
}

// QueryClassFixedAssets queries the "class_fixed_assets" edge of the Enumeration entity.
func (e *Enumeration) QueryClassFixedAssets() *FixedAssetQuery {
	return (&EnumerationClient{config: e.config}).QueryClassFixedAssets(e)
}

// QueryEmploymentStatusPeople queries the "employment_status_people" edge of the Enumeration entity.
func (e *Enumeration) QueryEmploymentStatusPeople() *PersonQuery {
	return (&EnumerationClient{config: e.config}).QueryEmploymentStatusPeople(e)
}

// QueryResidenceStatusPeople queries the "residence_status_people" edge of the Enumeration entity.
func (e *Enumeration) QueryResidenceStatusPeople() *PersonQuery {
	return (&EnumerationClient{config: e.config}).QueryResidenceStatusPeople(e)
}

// QueryMaritalStatusPeople queries the "marital_status_people" edge of the Enumeration entity.
func (e *Enumeration) QueryMaritalStatusPeople() *PersonQuery {
	return (&EnumerationClient{config: e.config}).QueryMaritalStatusPeople(e)
}

// QueryScopeWorkEfforts queries the "scope_work_efforts" edge of the Enumeration entity.
func (e *Enumeration) QueryScopeWorkEfforts() *WorkEffortQuery {
	return (&EnumerationClient{config: e.config}).QueryScopeWorkEfforts(e)
}

// QueryExpectationWorkEffortPartyAssignments queries the "expectation_work_effort_party_assignments" edge of the Enumeration entity.
func (e *Enumeration) QueryExpectationWorkEffortPartyAssignments() *WorkEffortPartyAssignmentQuery {
	return (&EnumerationClient{config: e.config}).QueryExpectationWorkEffortPartyAssignments(e)
}

// QueryDelegateReasonWorkEffortPartyAssignments queries the "delegate_reason_work_effort_party_assignments" edge of the Enumeration entity.
func (e *Enumeration) QueryDelegateReasonWorkEffortPartyAssignments() *WorkEffortPartyAssignmentQuery {
	return (&EnumerationClient{config: e.config}).QueryDelegateReasonWorkEffortPartyAssignments(e)
}

// Update returns a builder for updating this Enumeration.
// Note that you need to call Enumeration.Unwrap() before calling this method if this Enumeration
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Enumeration) Update() *EnumerationUpdateOne {
	return (&EnumerationClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Enumeration entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Enumeration) Unwrap() *Enumeration {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Enumeration is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Enumeration) String() string {
	var builder strings.Builder
	builder.WriteString("Enumeration(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(e.StringRef)
	builder.WriteString(", enum_code=")
	builder.WriteString(e.EnumCode)
	builder.WriteString(", sequence_id=")
	builder.WriteString(fmt.Sprintf("%v", e.SequenceID))
	builder.WriteString(", description=")
	builder.WriteString(e.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Enumerations is a parsable slice of Enumeration.
type Enumerations []*Enumeration

func (e Enumerations) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
