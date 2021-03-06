// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtype"
)

// ContactMechType is the model entity for the ContactMechType schema.
type ContactMechType struct {
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
	HasTable contactmechtype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContactMechTypeQuery when eager-loading is set.
	Edges                      ContactMechTypeEdges `json:"edges"`
	contact_mech_type_children *int
}

// ContactMechTypeEdges holds the relations/edges for other nodes in the graph.
type ContactMechTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ContactMechType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*ContactMechType `json:"children,omitempty"`
	// ContacMechTypeCommunicationEventTypes holds the value of the contac_mech_type_communication_event_types edge.
	ContacMechTypeCommunicationEventTypes []*CommunicationEventType `json:"contac_mech_type_communication_event_types,omitempty"`
	// ChildContactMechTypes holds the value of the child_contact_mech_types edge.
	ChildContactMechTypes []*ContactMechType `json:"child_contact_mech_types,omitempty"`
	// ContactMechTypePurposes holds the value of the contact_mech_type_purposes edge.
	ContactMechTypePurposes []*ContactMechTypePurpose `json:"contact_mech_type_purposes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContactMechTypeEdges) ParentOrErr() (*ContactMechType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: contactmechtype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ContactMechTypeEdges) ChildrenOrErr() ([]*ContactMechType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ContacMechTypeCommunicationEventTypesOrErr returns the ContacMechTypeCommunicationEventTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ContactMechTypeEdges) ContacMechTypeCommunicationEventTypesOrErr() ([]*CommunicationEventType, error) {
	if e.loadedTypes[2] {
		return e.ContacMechTypeCommunicationEventTypes, nil
	}
	return nil, &NotLoadedError{edge: "contac_mech_type_communication_event_types"}
}

// ChildContactMechTypesOrErr returns the ChildContactMechTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ContactMechTypeEdges) ChildContactMechTypesOrErr() ([]*ContactMechType, error) {
	if e.loadedTypes[3] {
		return e.ChildContactMechTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_contact_mech_types"}
}

// ContactMechTypePurposesOrErr returns the ContactMechTypePurposes value or an error if the edge
// was not loaded in eager-loading.
func (e ContactMechTypeEdges) ContactMechTypePurposesOrErr() ([]*ContactMechTypePurpose, error) {
	if e.loadedTypes[4] {
		return e.ContactMechTypePurposes, nil
	}
	return nil, &NotLoadedError{edge: "contact_mech_type_purposes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContactMechType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case contactmechtype.FieldID:
			values[i] = new(sql.NullInt64)
		case contactmechtype.FieldStringRef, contactmechtype.FieldHasTable, contactmechtype.FieldDescription:
			values[i] = new(sql.NullString)
		case contactmechtype.FieldCreateTime, contactmechtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case contactmechtype.ForeignKeys[0]: // contact_mech_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ContactMechType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContactMechType fields.
func (cmt *ContactMechType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contactmechtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cmt.ID = int(value.Int64)
		case contactmechtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cmt.CreateTime = value.Time
			}
		case contactmechtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cmt.UpdateTime = value.Time
			}
		case contactmechtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				cmt.StringRef = value.String
			}
		case contactmechtype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				cmt.HasTable = contactmechtype.HasTable(value.String)
			}
		case contactmechtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cmt.Description = value.String
			}
		case contactmechtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field contact_mech_type_children", value)
			} else if value.Valid {
				cmt.contact_mech_type_children = new(int)
				*cmt.contact_mech_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ContactMechType entity.
func (cmt *ContactMechType) QueryParent() *ContactMechTypeQuery {
	return (&ContactMechTypeClient{config: cmt.config}).QueryParent(cmt)
}

// QueryChildren queries the "children" edge of the ContactMechType entity.
func (cmt *ContactMechType) QueryChildren() *ContactMechTypeQuery {
	return (&ContactMechTypeClient{config: cmt.config}).QueryChildren(cmt)
}

// QueryContacMechTypeCommunicationEventTypes queries the "contac_mech_type_communication_event_types" edge of the ContactMechType entity.
func (cmt *ContactMechType) QueryContacMechTypeCommunicationEventTypes() *CommunicationEventTypeQuery {
	return (&ContactMechTypeClient{config: cmt.config}).QueryContacMechTypeCommunicationEventTypes(cmt)
}

// QueryChildContactMechTypes queries the "child_contact_mech_types" edge of the ContactMechType entity.
func (cmt *ContactMechType) QueryChildContactMechTypes() *ContactMechTypeQuery {
	return (&ContactMechTypeClient{config: cmt.config}).QueryChildContactMechTypes(cmt)
}

// QueryContactMechTypePurposes queries the "contact_mech_type_purposes" edge of the ContactMechType entity.
func (cmt *ContactMechType) QueryContactMechTypePurposes() *ContactMechTypePurposeQuery {
	return (&ContactMechTypeClient{config: cmt.config}).QueryContactMechTypePurposes(cmt)
}

// Update returns a builder for updating this ContactMechType.
// Note that you need to call ContactMechType.Unwrap() before calling this method if this ContactMechType
// was returned from a transaction, and the transaction was committed or rolled back.
func (cmt *ContactMechType) Update() *ContactMechTypeUpdateOne {
	return (&ContactMechTypeClient{config: cmt.config}).UpdateOne(cmt)
}

// Unwrap unwraps the ContactMechType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cmt *ContactMechType) Unwrap() *ContactMechType {
	tx, ok := cmt.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContactMechType is not a transactional entity")
	}
	cmt.config.driver = tx.drv
	return cmt
}

// String implements the fmt.Stringer.
func (cmt *ContactMechType) String() string {
	var builder strings.Builder
	builder.WriteString("ContactMechType(")
	builder.WriteString(fmt.Sprintf("id=%v", cmt.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(cmt.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(cmt.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(cmt.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", cmt.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(cmt.Description)
	builder.WriteByte(')')
	return builder.String()
}

// ContactMechTypes is a parsable slice of ContactMechType.
type ContactMechTypes []*ContactMechType

func (cmt ContactMechTypes) config(cfg config) {
	for _i := range cmt {
		cmt[_i].config = cfg
	}
}
