// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventprptyp"
)

// CommunicationEventPrpTyp is the model entity for the CommunicationEventPrpTyp schema.
type CommunicationEventPrpTyp struct {
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
	HasTable communicationeventprptyp.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommunicationEventPrpTypQuery when eager-loading is set.
	Edges                                CommunicationEventPrpTypEdges `json:"edges"`
	communication_event_prp_typ_children *int
}

// CommunicationEventPrpTypEdges holds the relations/edges for other nodes in the graph.
type CommunicationEventPrpTypEdges struct {
	// Parent holds the value of the parent edge.
	Parent *CommunicationEventPrpTyp `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*CommunicationEventPrpTyp `json:"children,omitempty"`
	// ChildCommunicationEventPrpTyps holds the value of the child_communication_event_prp_typs edge.
	ChildCommunicationEventPrpTyps []*CommunicationEventPrpTyp `json:"child_communication_event_prp_typs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommunicationEventPrpTypEdges) ParentOrErr() (*CommunicationEventPrpTyp, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: communicationeventprptyp.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e CommunicationEventPrpTypEdges) ChildrenOrErr() ([]*CommunicationEventPrpTyp, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ChildCommunicationEventPrpTypsOrErr returns the ChildCommunicationEventPrpTyps value or an error if the edge
// was not loaded in eager-loading.
func (e CommunicationEventPrpTypEdges) ChildCommunicationEventPrpTypsOrErr() ([]*CommunicationEventPrpTyp, error) {
	if e.loadedTypes[2] {
		return e.ChildCommunicationEventPrpTyps, nil
	}
	return nil, &NotLoadedError{edge: "child_communication_event_prp_typs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommunicationEventPrpTyp) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case communicationeventprptyp.FieldID:
			values[i] = new(sql.NullInt64)
		case communicationeventprptyp.FieldStringRef, communicationeventprptyp.FieldHasTable, communicationeventprptyp.FieldDescription:
			values[i] = new(sql.NullString)
		case communicationeventprptyp.FieldCreateTime, communicationeventprptyp.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case communicationeventprptyp.ForeignKeys[0]: // communication_event_prp_typ_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CommunicationEventPrpTyp", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommunicationEventPrpTyp fields.
func (cept *CommunicationEventPrpTyp) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case communicationeventprptyp.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cept.ID = int(value.Int64)
		case communicationeventprptyp.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cept.CreateTime = value.Time
			}
		case communicationeventprptyp.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cept.UpdateTime = value.Time
			}
		case communicationeventprptyp.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				cept.StringRef = value.String
			}
		case communicationeventprptyp.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				cept.HasTable = communicationeventprptyp.HasTable(value.String)
			}
		case communicationeventprptyp.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cept.Description = value.String
			}
		case communicationeventprptyp.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field communication_event_prp_typ_children", value)
			} else if value.Valid {
				cept.communication_event_prp_typ_children = new(int)
				*cept.communication_event_prp_typ_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the CommunicationEventPrpTyp entity.
func (cept *CommunicationEventPrpTyp) QueryParent() *CommunicationEventPrpTypQuery {
	return (&CommunicationEventPrpTypClient{config: cept.config}).QueryParent(cept)
}

// QueryChildren queries the "children" edge of the CommunicationEventPrpTyp entity.
func (cept *CommunicationEventPrpTyp) QueryChildren() *CommunicationEventPrpTypQuery {
	return (&CommunicationEventPrpTypClient{config: cept.config}).QueryChildren(cept)
}

// QueryChildCommunicationEventPrpTyps queries the "child_communication_event_prp_typs" edge of the CommunicationEventPrpTyp entity.
func (cept *CommunicationEventPrpTyp) QueryChildCommunicationEventPrpTyps() *CommunicationEventPrpTypQuery {
	return (&CommunicationEventPrpTypClient{config: cept.config}).QueryChildCommunicationEventPrpTyps(cept)
}

// Update returns a builder for updating this CommunicationEventPrpTyp.
// Note that you need to call CommunicationEventPrpTyp.Unwrap() before calling this method if this CommunicationEventPrpTyp
// was returned from a transaction, and the transaction was committed or rolled back.
func (cept *CommunicationEventPrpTyp) Update() *CommunicationEventPrpTypUpdateOne {
	return (&CommunicationEventPrpTypClient{config: cept.config}).UpdateOne(cept)
}

// Unwrap unwraps the CommunicationEventPrpTyp entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cept *CommunicationEventPrpTyp) Unwrap() *CommunicationEventPrpTyp {
	tx, ok := cept.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommunicationEventPrpTyp is not a transactional entity")
	}
	cept.config.driver = tx.drv
	return cept
}

// String implements the fmt.Stringer.
func (cept *CommunicationEventPrpTyp) String() string {
	var builder strings.Builder
	builder.WriteString("CommunicationEventPrpTyp(")
	builder.WriteString(fmt.Sprintf("id=%v", cept.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(cept.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(cept.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(cept.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", cept.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(cept.Description)
	builder.WriteByte(')')
	return builder.String()
}

// CommunicationEventPrpTyps is a parsable slice of CommunicationEventPrpTyp.
type CommunicationEventPrpTyps []*CommunicationEventPrpTyp

func (cept CommunicationEventPrpTyps) config(cfg config) {
	for _i := range cept {
		cept[_i].config = cfg
	}
}
