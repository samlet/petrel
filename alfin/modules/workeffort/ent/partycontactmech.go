// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partycontactmech"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partyrole"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/roletype"
)

// PartyContactMech is the model entity for the PartyContactMech schema.
type PartyContactMech struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// ContactMechID holds the value of the "contact_mech_id" field.
	ContactMechID int `json:"contact_mech_id,omitempty"`
	// FromDate holds the value of the "from_date" field.
	FromDate time.Time `json:"from_date,omitempty"`
	// ThruDate holds the value of the "thru_date" field.
	ThruDate time.Time `json:"thru_date,omitempty"`
	// AllowSolicitation holds the value of the "allow_solicitation" field.
	AllowSolicitation partycontactmech.AllowSolicitation `json:"allow_solicitation,omitempty"`
	// Extension holds the value of the "extension" field.
	Extension string `json:"extension,omitempty"`
	// Verified holds the value of the "verified" field.
	Verified partycontactmech.Verified `json:"verified,omitempty"`
	// Comments holds the value of the "comments" field.
	Comments string `json:"comments,omitempty"`
	// YearsWithContactMech holds the value of the "years_with_contact_mech" field.
	YearsWithContactMech int `json:"years_with_contact_mech,omitempty"`
	// MonthsWithContactMech holds the value of the "months_with_contact_mech" field.
	MonthsWithContactMech int `json:"months_with_contact_mech,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartyContactMechQuery when eager-loading is set.
	Edges                           PartyContactMechEdges `json:"edges"`
	party_party_contact_meches      *int
	party_role_party_contact_meches *int
	person_party_contact_meches     *int
	role_type_party_contact_meches  *int
}

// PartyContactMechEdges holds the relations/edges for other nodes in the graph.
type PartyContactMechEdges struct {
	// Party holds the value of the party edge.
	Party *Party `json:"party,omitempty"`
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// PartyRole holds the value of the party_role edge.
	PartyRole *PartyRole `json:"party_role,omitempty"`
	// RoleType holds the value of the role_type edge.
	RoleType *RoleType `json:"role_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PartyOrErr returns the Party value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyContactMechEdges) PartyOrErr() (*Party, error) {
	if e.loadedTypes[0] {
		if e.Party == nil {
			// The edge party was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: party.Label}
		}
		return e.Party, nil
	}
	return nil, &NotLoadedError{edge: "party"}
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyContactMechEdges) PersonOrErr() (*Person, error) {
	if e.loadedTypes[1] {
		if e.Person == nil {
			// The edge person was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.Person, nil
	}
	return nil, &NotLoadedError{edge: "person"}
}

// PartyRoleOrErr returns the PartyRole value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyContactMechEdges) PartyRoleOrErr() (*PartyRole, error) {
	if e.loadedTypes[2] {
		if e.PartyRole == nil {
			// The edge party_role was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: partyrole.Label}
		}
		return e.PartyRole, nil
	}
	return nil, &NotLoadedError{edge: "party_role"}
}

// RoleTypeOrErr returns the RoleType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyContactMechEdges) RoleTypeOrErr() (*RoleType, error) {
	if e.loadedTypes[3] {
		if e.RoleType == nil {
			// The edge role_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roletype.Label}
		}
		return e.RoleType, nil
	}
	return nil, &NotLoadedError{edge: "role_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PartyContactMech) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case partycontactmech.FieldID, partycontactmech.FieldContactMechID, partycontactmech.FieldYearsWithContactMech, partycontactmech.FieldMonthsWithContactMech:
			values[i] = new(sql.NullInt64)
		case partycontactmech.FieldStringRef, partycontactmech.FieldAllowSolicitation, partycontactmech.FieldExtension, partycontactmech.FieldVerified, partycontactmech.FieldComments:
			values[i] = new(sql.NullString)
		case partycontactmech.FieldCreateTime, partycontactmech.FieldUpdateTime, partycontactmech.FieldFromDate, partycontactmech.FieldThruDate:
			values[i] = new(sql.NullTime)
		case partycontactmech.ForeignKeys[0]: // party_party_contact_meches
			values[i] = new(sql.NullInt64)
		case partycontactmech.ForeignKeys[1]: // party_role_party_contact_meches
			values[i] = new(sql.NullInt64)
		case partycontactmech.ForeignKeys[2]: // person_party_contact_meches
			values[i] = new(sql.NullInt64)
		case partycontactmech.ForeignKeys[3]: // role_type_party_contact_meches
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PartyContactMech", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PartyContactMech fields.
func (pcm *PartyContactMech) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case partycontactmech.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pcm.ID = int(value.Int64)
		case partycontactmech.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pcm.CreateTime = value.Time
			}
		case partycontactmech.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pcm.UpdateTime = value.Time
			}
		case partycontactmech.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				pcm.StringRef = value.String
			}
		case partycontactmech.FieldContactMechID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contact_mech_id", values[i])
			} else if value.Valid {
				pcm.ContactMechID = int(value.Int64)
			}
		case partycontactmech.FieldFromDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from_date", values[i])
			} else if value.Valid {
				pcm.FromDate = value.Time
			}
		case partycontactmech.FieldThruDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field thru_date", values[i])
			} else if value.Valid {
				pcm.ThruDate = value.Time
			}
		case partycontactmech.FieldAllowSolicitation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field allow_solicitation", values[i])
			} else if value.Valid {
				pcm.AllowSolicitation = partycontactmech.AllowSolicitation(value.String)
			}
		case partycontactmech.FieldExtension:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension", values[i])
			} else if value.Valid {
				pcm.Extension = value.String
			}
		case partycontactmech.FieldVerified:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field verified", values[i])
			} else if value.Valid {
				pcm.Verified = partycontactmech.Verified(value.String)
			}
		case partycontactmech.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				pcm.Comments = value.String
			}
		case partycontactmech.FieldYearsWithContactMech:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field years_with_contact_mech", values[i])
			} else if value.Valid {
				pcm.YearsWithContactMech = int(value.Int64)
			}
		case partycontactmech.FieldMonthsWithContactMech:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field months_with_contact_mech", values[i])
			} else if value.Valid {
				pcm.MonthsWithContactMech = int(value.Int64)
			}
		case partycontactmech.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field party_party_contact_meches", value)
			} else if value.Valid {
				pcm.party_party_contact_meches = new(int)
				*pcm.party_party_contact_meches = int(value.Int64)
			}
		case partycontactmech.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field party_role_party_contact_meches", value)
			} else if value.Valid {
				pcm.party_role_party_contact_meches = new(int)
				*pcm.party_role_party_contact_meches = int(value.Int64)
			}
		case partycontactmech.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field person_party_contact_meches", value)
			} else if value.Valid {
				pcm.person_party_contact_meches = new(int)
				*pcm.person_party_contact_meches = int(value.Int64)
			}
		case partycontactmech.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_type_party_contact_meches", value)
			} else if value.Valid {
				pcm.role_type_party_contact_meches = new(int)
				*pcm.role_type_party_contact_meches = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParty queries the "party" edge of the PartyContactMech entity.
func (pcm *PartyContactMech) QueryParty() *PartyQuery {
	return (&PartyContactMechClient{config: pcm.config}).QueryParty(pcm)
}

// QueryPerson queries the "person" edge of the PartyContactMech entity.
func (pcm *PartyContactMech) QueryPerson() *PersonQuery {
	return (&PartyContactMechClient{config: pcm.config}).QueryPerson(pcm)
}

// QueryPartyRole queries the "party_role" edge of the PartyContactMech entity.
func (pcm *PartyContactMech) QueryPartyRole() *PartyRoleQuery {
	return (&PartyContactMechClient{config: pcm.config}).QueryPartyRole(pcm)
}

// QueryRoleType queries the "role_type" edge of the PartyContactMech entity.
func (pcm *PartyContactMech) QueryRoleType() *RoleTypeQuery {
	return (&PartyContactMechClient{config: pcm.config}).QueryRoleType(pcm)
}

// Update returns a builder for updating this PartyContactMech.
// Note that you need to call PartyContactMech.Unwrap() before calling this method if this PartyContactMech
// was returned from a transaction, and the transaction was committed or rolled back.
func (pcm *PartyContactMech) Update() *PartyContactMechUpdateOne {
	return (&PartyContactMechClient{config: pcm.config}).UpdateOne(pcm)
}

// Unwrap unwraps the PartyContactMech entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pcm *PartyContactMech) Unwrap() *PartyContactMech {
	tx, ok := pcm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PartyContactMech is not a transactional entity")
	}
	pcm.config.driver = tx.drv
	return pcm
}

// String implements the fmt.Stringer.
func (pcm *PartyContactMech) String() string {
	var builder strings.Builder
	builder.WriteString("PartyContactMech(")
	builder.WriteString(fmt.Sprintf("id=%v", pcm.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pcm.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pcm.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(pcm.StringRef)
	builder.WriteString(", contact_mech_id=")
	builder.WriteString(fmt.Sprintf("%v", pcm.ContactMechID))
	builder.WriteString(", from_date=")
	builder.WriteString(pcm.FromDate.Format(time.ANSIC))
	builder.WriteString(", thru_date=")
	builder.WriteString(pcm.ThruDate.Format(time.ANSIC))
	builder.WriteString(", allow_solicitation=")
	builder.WriteString(fmt.Sprintf("%v", pcm.AllowSolicitation))
	builder.WriteString(", extension=")
	builder.WriteString(pcm.Extension)
	builder.WriteString(", verified=")
	builder.WriteString(fmt.Sprintf("%v", pcm.Verified))
	builder.WriteString(", comments=")
	builder.WriteString(pcm.Comments)
	builder.WriteString(", years_with_contact_mech=")
	builder.WriteString(fmt.Sprintf("%v", pcm.YearsWithContactMech))
	builder.WriteString(", months_with_contact_mech=")
	builder.WriteString(fmt.Sprintf("%v", pcm.MonthsWithContactMech))
	builder.WriteByte(')')
	return builder.String()
}

// PartyContactMeches is a parsable slice of PartyContactMech.
type PartyContactMeches []*PartyContactMech

func (pcm PartyContactMeches) config(cfg config) {
	for _i := range pcm {
		pcm[_i].config = cfg
	}
}
