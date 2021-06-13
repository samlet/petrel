// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyclassificationtype"
)

// PartyClassificationType is the model entity for the PartyClassificationType schema.
type PartyClassificationType struct {
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
	HasTable partyclassificationtype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartyClassificationTypeQuery when eager-loading is set.
	Edges                              PartyClassificationTypeEdges `json:"edges"`
	party_classification_type_children *int
}

// PartyClassificationTypeEdges holds the relations/edges for other nodes in the graph.
type PartyClassificationTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *PartyClassificationType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*PartyClassificationType `json:"children,omitempty"`
	// ChildPartyClassificationTypes holds the value of the child_party_classification_types edge.
	ChildPartyClassificationTypes []*PartyClassificationType `json:"child_party_classification_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyClassificationTypeEdges) ParentOrErr() (*PartyClassificationType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: partyclassificationtype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e PartyClassificationTypeEdges) ChildrenOrErr() ([]*PartyClassificationType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ChildPartyClassificationTypesOrErr returns the ChildPartyClassificationTypes value or an error if the edge
// was not loaded in eager-loading.
func (e PartyClassificationTypeEdges) ChildPartyClassificationTypesOrErr() ([]*PartyClassificationType, error) {
	if e.loadedTypes[2] {
		return e.ChildPartyClassificationTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_party_classification_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartyClassificationType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case partyclassificationtype.FieldID:
			values[i] = new(sql.NullInt64)
		case partyclassificationtype.FieldStringRef, partyclassificationtype.FieldHasTable, partyclassificationtype.FieldDescription:
			values[i] = new(sql.NullString)
		case partyclassificationtype.FieldCreateTime, partyclassificationtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case partyclassificationtype.ForeignKeys[0]: // party_classification_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PartyClassificationType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartyClassificationType fields.
func (pct *PartyClassificationType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partyclassificationtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pct.ID = int(value.Int64)
		case partyclassificationtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pct.CreateTime = value.Time
			}
		case partyclassificationtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pct.UpdateTime = value.Time
			}
		case partyclassificationtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				pct.StringRef = value.String
			}
		case partyclassificationtype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				pct.HasTable = partyclassificationtype.HasTable(value.String)
			}
		case partyclassificationtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pct.Description = value.String
			}
		case partyclassificationtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field party_classification_type_children", value)
			} else if value.Valid {
				pct.party_classification_type_children = new(int)
				*pct.party_classification_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the PartyClassificationType entity.
func (pct *PartyClassificationType) QueryParent() *PartyClassificationTypeQuery {
	return (&PartyClassificationTypeClient{config: pct.config}).QueryParent(pct)
}

// QueryChildren queries the "children" edge of the PartyClassificationType entity.
func (pct *PartyClassificationType) QueryChildren() *PartyClassificationTypeQuery {
	return (&PartyClassificationTypeClient{config: pct.config}).QueryChildren(pct)
}

// QueryChildPartyClassificationTypes queries the "child_party_classification_types" edge of the PartyClassificationType entity.
func (pct *PartyClassificationType) QueryChildPartyClassificationTypes() *PartyClassificationTypeQuery {
	return (&PartyClassificationTypeClient{config: pct.config}).QueryChildPartyClassificationTypes(pct)
}

// Update returns a builder for updating this PartyClassificationType.
// Note that you need to call PartyClassificationType.Unwrap() before calling this method if this PartyClassificationType
// was returned from a transaction, and the transaction was committed or rolled back.
func (pct *PartyClassificationType) Update() *PartyClassificationTypeUpdateOne {
	return (&PartyClassificationTypeClient{config: pct.config}).UpdateOne(pct)
}

// Unwrap unwraps the PartyClassificationType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pct *PartyClassificationType) Unwrap() *PartyClassificationType {
	tx, ok := pct.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartyClassificationType is not a transactional entity")
	}
	pct.config.driver = tx.drv
	return pct
}

// String implements the fmt.Stringer.
func (pct *PartyClassificationType) String() string {
	var builder strings.Builder
	builder.WriteString("PartyClassificationType(")
	builder.WriteString(fmt.Sprintf("id=%v", pct.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pct.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pct.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(pct.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", pct.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(pct.Description)
	builder.WriteByte(')')
	return builder.String()
}

// PartyClassificationTypes is a parsable slice of PartyClassificationType.
type PartyClassificationTypes []*PartyClassificationType

func (pct PartyClassificationTypes) config(cfg config) {
	for _i := range pct {
		pct[_i].config = cfg
	}
}
