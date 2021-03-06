// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/communicationeventtype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/contactmechtype"
)

// CommunicationEventType is the model entity for the CommunicationEventType schema.
type CommunicationEventType struct {
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
	HasTable communicationeventtype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommunicationEventTypeQuery when eager-loading is set.
	Edges                                                        CommunicationEventTypeEdges `json:"edges"`
	communication_event_type_children                            *int
	contact_mech_type_contac_mech_type_communication_event_types *int
}

// CommunicationEventTypeEdges holds the relations/edges for other nodes in the graph.
type CommunicationEventTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *CommunicationEventType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*CommunicationEventType `json:"children,omitempty"`
	// ContacMechTypeContactMechType holds the value of the contac_mech_type_contact_mech_type edge.
	ContacMechTypeContactMechType *ContactMechType `json:"contac_mech_type_contact_mech_type,omitempty"`
	// ChildCommunicationEventTypes holds the value of the child_communication_event_types edge.
	ChildCommunicationEventTypes []*CommunicationEventType `json:"child_communication_event_types,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommunicationEventTypeEdges) ParentOrErr() (*CommunicationEventType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: communicationeventtype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e CommunicationEventTypeEdges) ChildrenOrErr() ([]*CommunicationEventType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ContacMechTypeContactMechTypeOrErr returns the ContacMechTypeContactMechType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommunicationEventTypeEdges) ContacMechTypeContactMechTypeOrErr() (*ContactMechType, error) {
	if e.loadedTypes[2] {
		if e.ContacMechTypeContactMechType == nil {
			// The edge contac_mech_type_contact_mech_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: contactmechtype.Label}
		}
		return e.ContacMechTypeContactMechType, nil
	}
	return nil, &NotLoadedError{edge: "contac_mech_type_contact_mech_type"}
}

// ChildCommunicationEventTypesOrErr returns the ChildCommunicationEventTypes value or an error if the edge
// was not loaded in eager-loading.
func (e CommunicationEventTypeEdges) ChildCommunicationEventTypesOrErr() ([]*CommunicationEventType, error) {
	if e.loadedTypes[3] {
		return e.ChildCommunicationEventTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_communication_event_types"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CommunicationEventType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case communicationeventtype.FieldID:
			values[i] = new(sql.NullInt64)
		case communicationeventtype.FieldStringRef, communicationeventtype.FieldHasTable, communicationeventtype.FieldDescription:
			values[i] = new(sql.NullString)
		case communicationeventtype.FieldCreateTime, communicationeventtype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case communicationeventtype.ForeignKeys[0]: // communication_event_type_children
			values[i] = new(sql.NullInt64)
		case communicationeventtype.ForeignKeys[1]: // contact_mech_type_contac_mech_type_communication_event_types
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CommunicationEventType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CommunicationEventType fields.
func (cet *CommunicationEventType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case communicationeventtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cet.ID = int(value.Int64)
		case communicationeventtype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				cet.CreateTime = value.Time
			}
		case communicationeventtype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				cet.UpdateTime = value.Time
			}
		case communicationeventtype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				cet.StringRef = value.String
			}
		case communicationeventtype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				cet.HasTable = communicationeventtype.HasTable(value.String)
			}
		case communicationeventtype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				cet.Description = value.String
			}
		case communicationeventtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field communication_event_type_children", value)
			} else if value.Valid {
				cet.communication_event_type_children = new(int)
				*cet.communication_event_type_children = int(value.Int64)
			}
		case communicationeventtype.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field contact_mech_type_contac_mech_type_communication_event_types", value)
			} else if value.Valid {
				cet.contact_mech_type_contac_mech_type_communication_event_types = new(int)
				*cet.contact_mech_type_contac_mech_type_communication_event_types = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the CommunicationEventType entity.
func (cet *CommunicationEventType) QueryParent() *CommunicationEventTypeQuery {
	return (&CommunicationEventTypeClient{config: cet.config}).QueryParent(cet)
}

// QueryChildren queries the "children" edge of the CommunicationEventType entity.
func (cet *CommunicationEventType) QueryChildren() *CommunicationEventTypeQuery {
	return (&CommunicationEventTypeClient{config: cet.config}).QueryChildren(cet)
}

// QueryContacMechTypeContactMechType queries the "contac_mech_type_contact_mech_type" edge of the CommunicationEventType entity.
func (cet *CommunicationEventType) QueryContacMechTypeContactMechType() *ContactMechTypeQuery {
	return (&CommunicationEventTypeClient{config: cet.config}).QueryContacMechTypeContactMechType(cet)
}

// QueryChildCommunicationEventTypes queries the "child_communication_event_types" edge of the CommunicationEventType entity.
func (cet *CommunicationEventType) QueryChildCommunicationEventTypes() *CommunicationEventTypeQuery {
	return (&CommunicationEventTypeClient{config: cet.config}).QueryChildCommunicationEventTypes(cet)
}

// Update returns a builder for updating this CommunicationEventType.
// Note that you need to call CommunicationEventType.Unwrap() before calling this method if this CommunicationEventType
// was returned from a transaction, and the transaction was committed or rolled back.
func (cet *CommunicationEventType) Update() *CommunicationEventTypeUpdateOne {
	return (&CommunicationEventTypeClient{config: cet.config}).UpdateOne(cet)
}

// Unwrap unwraps the CommunicationEventType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cet *CommunicationEventType) Unwrap() *CommunicationEventType {
	tx, ok := cet.config.driver.(*txDriver)
	if !ok {
		panic("ent: CommunicationEventType is not a transactional entity")
	}
	cet.config.driver = tx.drv
	return cet
}

// String implements the fmt.Stringer.
func (cet *CommunicationEventType) String() string {
	var builder strings.Builder
	builder.WriteString("CommunicationEventType(")
	builder.WriteString(fmt.Sprintf("id=%v", cet.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(cet.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(cet.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(cet.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", cet.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(cet.Description)
	builder.WriteByte(')')
	return builder.String()
}

// CommunicationEventTypes is a parsable slice of CommunicationEventType.
type CommunicationEventTypes []*CommunicationEventType

func (cet CommunicationEventTypes) config(cfg config) {
	for _i := range cet {
		cet[_i].config = cfg
	}
}
