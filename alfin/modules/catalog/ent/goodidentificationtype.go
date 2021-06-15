// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/goodidentificationtype"
)

// GoodIdentificationType is the model entity for the GoodIdentificationType schema.
type GoodIdentificationType struct {
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
	HasTable goodidentificationtype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GoodIdentificationTypeQuery when eager-loading is set.
	Edges                             GoodIdentificationTypeEdges `json:"edges"`
	good_identification_type_children *int
}

// GoodIdentificationTypeEdges holds the relations/edges for other nodes in the graph.
type GoodIdentificationTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *GoodIdentificationType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*GoodIdentificationType `json:"children,omitempty"`
	// ChildGoodIdentificationTypes holds the value of the child_good_identification_types edge.
	ChildGoodIdentificationTypes []*GoodIdentificationType `json:"child_good_identification_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GoodIdentificationTypeEdges) ParentOrErr() (*GoodIdentificationType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: goodidentificationtype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e GoodIdentificationTypeEdges) ChildrenOrErr() ([]*GoodIdentificationType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ChildGoodIdentificationTypesOrErr returns the ChildGoodIdentificationTypes value or an error if the edge
// was not loaded in eager-loading.
func (e GoodIdentificationTypeEdges) ChildGoodIdentificationTypesOrErr() ([]*GoodIdentificationType, error) {
	if e.loadedTypes[2] {
		return e.ChildGoodIdentificationTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_good_identification_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodIdentificationType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodidentificationtype.FieldID:
			values[i] = new(sql.NullInt64)
		case goodidentificationtype.FieldStringRef, goodidentificationtype.FieldHasTable, goodidentificationtype.FieldDescription:
			values[i] = new(sql.NullString)
		case goodidentificationtype.FieldCreateTime, goodidentificationtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case goodidentificationtype.ForeignKeys[0]: // good_identification_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodIdentificationType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodIdentificationType fields.
func (git *GoodIdentificationType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodidentificationtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			git.ID = int(value.Int64)
		case goodidentificationtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				git.CreateTime = value.Time
			}
		case goodidentificationtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				git.UpdateTime = value.Time
			}
		case goodidentificationtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				git.StringRef = value.String
			}
		case goodidentificationtype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				git.HasTable = goodidentificationtype.HasTable(value.String)
			}
		case goodidentificationtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				git.Description = value.String
			}
		case goodidentificationtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field good_identification_type_children", value)
			} else if value.Valid {
				git.good_identification_type_children = new(int)
				*git.good_identification_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the GoodIdentificationType entity.
func (git *GoodIdentificationType) QueryParent() *GoodIdentificationTypeQuery {
	return (&GoodIdentificationTypeClient{config: git.config}).QueryParent(git)
}

// QueryChildren queries the "children" edge of the GoodIdentificationType entity.
func (git *GoodIdentificationType) QueryChildren() *GoodIdentificationTypeQuery {
	return (&GoodIdentificationTypeClient{config: git.config}).QueryChildren(git)
}

// QueryChildGoodIdentificationTypes queries the "child_good_identification_types" edge of the GoodIdentificationType entity.
func (git *GoodIdentificationType) QueryChildGoodIdentificationTypes() *GoodIdentificationTypeQuery {
	return (&GoodIdentificationTypeClient{config: git.config}).QueryChildGoodIdentificationTypes(git)
}

// Update returns a builder for updating this GoodIdentificationType.
// Note that you need to call GoodIdentificationType.Unwrap() before calling this method if this GoodIdentificationType
// was returned from a transaction, and the transaction was committed or rolled back.
func (git *GoodIdentificationType) Update() *GoodIdentificationTypeUpdateOne {
	return (&GoodIdentificationTypeClient{config: git.config}).UpdateOne(git)
}

// Unwrap unwraps the GoodIdentificationType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (git *GoodIdentificationType) Unwrap() *GoodIdentificationType {
	tx, ok := git.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodIdentificationType is not a transactional entity")
	}
	git.config.driver = tx.drv
	return git
}

// String implements the fmt.Stringer.
func (git *GoodIdentificationType) String() string {
	var builder strings.Builder
	builder.WriteString("GoodIdentificationType(")
	builder.WriteString(fmt.Sprintf("id=%v", git.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(git.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(git.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(git.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", git.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(git.Description)
	builder.WriteByte(')')
	return builder.String()
}

// GoodIdentificationTypes is a parsable slice of GoodIdentificationType.
type GoodIdentificationTypes []*GoodIdentificationType

func (git GoodIdentificationTypes) config(cfg config) {
	for _i := range git {
		git[_i].config = cfg
	}
}
