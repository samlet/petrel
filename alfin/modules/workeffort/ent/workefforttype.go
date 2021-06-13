// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/workefforttype"
)

// WorkEffortType is the model entity for the WorkEffortType schema.
type WorkEffortType struct {
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
	HasTable workefforttype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkEffortTypeQuery when eager-loading is set.
	Edges                     WorkEffortTypeEdges `json:"edges"`
	work_effort_type_children *int
}

// WorkEffortTypeEdges holds the relations/edges for other nodes in the graph.
type WorkEffortTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *WorkEffortType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*WorkEffortType `json:"children,omitempty"`
	// WorkEfforts holds the value of the work_efforts edge.
	WorkEfforts []*WorkEffort `json:"work_efforts,omitempty"`
	// ChildWorkEffortTypes holds the value of the child_work_effort_types edge.
	ChildWorkEffortTypes []*WorkEffortType `json:"child_work_effort_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEffortTypeEdges) ParentOrErr() (*WorkEffortType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workefforttype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEffortTypeEdges) ChildrenOrErr() ([]*WorkEffortType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// WorkEffortsOrErr returns the WorkEfforts value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEffortTypeEdges) WorkEffortsOrErr() ([]*WorkEffort, error) {
	if e.loadedTypes[2] {
		return e.WorkEfforts, nil
	}
	return nil, &NotLoadedError{edge: "work_efforts"}
}

// ChildWorkEffortTypesOrErr returns the ChildWorkEffortTypes value or an error if the edge
// was not loaded in eager-loading.
func (e WorkEffortTypeEdges) ChildWorkEffortTypesOrErr() ([]*WorkEffortType, error) {
	if e.loadedTypes[3] {
		return e.ChildWorkEffortTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_work_effort_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkEffortType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workefforttype.FieldID:
			values[i] = new(sql.NullInt64)
		case workefforttype.FieldStringRef, workefforttype.FieldHasTable, workefforttype.FieldDescription:
			values[i] = new(sql.NullString)
		case workefforttype.FieldCreateTime, workefforttype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case workefforttype.ForeignKeys[0]: // work_effort_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkEffortType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkEffortType fields.
func (wet *WorkEffortType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workefforttype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wet.ID = int(value.Int64)
		case workefforttype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				wet.CreateTime = value.Time
			}
		case workefforttype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				wet.UpdateTime = value.Time
			}
		case workefforttype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				wet.StringRef = value.String
			}
		case workefforttype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				wet.HasTable = workefforttype.HasTable(value.String)
			}
		case workefforttype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				wet.Description = value.String
			}
		case workefforttype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field work_effort_type_children", value)
			} else if value.Valid {
				wet.work_effort_type_children = new(int)
				*wet.work_effort_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the WorkEffortType entity.
func (wet *WorkEffortType) QueryParent() *WorkEffortTypeQuery {
	return (&WorkEffortTypeClient{config: wet.config}).QueryParent(wet)
}

// QueryChildren queries the "children" edge of the WorkEffortType entity.
func (wet *WorkEffortType) QueryChildren() *WorkEffortTypeQuery {
	return (&WorkEffortTypeClient{config: wet.config}).QueryChildren(wet)
}

// QueryWorkEfforts queries the "work_efforts" edge of the WorkEffortType entity.
func (wet *WorkEffortType) QueryWorkEfforts() *WorkEffortQuery {
	return (&WorkEffortTypeClient{config: wet.config}).QueryWorkEfforts(wet)
}

// QueryChildWorkEffortTypes queries the "child_work_effort_types" edge of the WorkEffortType entity.
func (wet *WorkEffortType) QueryChildWorkEffortTypes() *WorkEffortTypeQuery {
	return (&WorkEffortTypeClient{config: wet.config}).QueryChildWorkEffortTypes(wet)
}

// Update returns a builder for updating this WorkEffortType.
// Note that you need to call WorkEffortType.Unwrap() before calling this method if this WorkEffortType
// was returned from a transaction, and the transaction was committed or rolled back.
func (wet *WorkEffortType) Update() *WorkEffortTypeUpdateOne {
	return (&WorkEffortTypeClient{config: wet.config}).UpdateOne(wet)
}

// Unwrap unwraps the WorkEffortType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wet *WorkEffortType) Unwrap() *WorkEffortType {
	tx, ok := wet.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkEffortType is not a transactional entity")
	}
	wet.config.driver = tx.drv
	return wet
}

// String implements the fmt.Stringer.
func (wet *WorkEffortType) String() string {
	var builder strings.Builder
	builder.WriteString("WorkEffortType(")
	builder.WriteString(fmt.Sprintf("id=%v", wet.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(wet.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(wet.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(wet.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", wet.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(wet.Description)
	builder.WriteByte(')')
	return builder.String()
}

// WorkEffortTypes is a parsable slice of WorkEffortType.
type WorkEffortTypes []*WorkEffortType

func (wet WorkEffortTypes) config(cfg config) {
	for _i := range wet {
		wet[_i].config = cfg
	}
}
