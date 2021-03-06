// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/party"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/partytype"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/person"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/statusitem"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
)

// Party is the model entity for the Party schema.
type Party struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// ExternalID holds the value of the "external_id" field.
	ExternalID int `json:"external_id,omitempty"`
	// PreferredCurrencyUomID holds the value of the "preferred_currency_uom_id" field.
	PreferredCurrencyUomID int `json:"preferred_currency_uom_id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedDate holds the value of the "created_date" field.
	CreatedDate time.Time `json:"created_date,omitempty"`
	// LastModifiedDate holds the value of the "last_modified_date" field.
	LastModifiedDate time.Time `json:"last_modified_date,omitempty"`
	// DataSourceID holds the value of the "data_source_id" field.
	DataSourceID int `json:"data_source_id,omitempty"`
	// IsUnread holds the value of the "is_unread" field.
	IsUnread party.IsUnread `json:"is_unread,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PartyQuery when eager-loading is set.
	Edges                               PartyEdges `json:"edges"`
	party_type_parties                  *int
	status_item_parties                 *int
	user_login_created_by_parties       *int
	user_login_last_modified_by_parties *int
}

// PartyEdges holds the relations/edges for other nodes in the graph.
type PartyEdges struct {
	// PartyType holds the value of the party_type edge.
	PartyType *PartyType `json:"party_type,omitempty"`
	// CreatedByUserLogin holds the value of the created_by_user_login edge.
	CreatedByUserLogin *UserLogin `json:"created_by_user_login,omitempty"`
	// LastModifiedByUserLogin holds the value of the last_modified_by_user_login edge.
	LastModifiedByUserLogin *UserLogin `json:"last_modified_by_user_login,omitempty"`
	// StatusItem holds the value of the status_item edge.
	StatusItem *StatusItem `json:"status_item,omitempty"`
	// FixedAssets holds the value of the fixed_assets edge.
	FixedAssets []*FixedAsset `json:"fixed_assets,omitempty"`
	// PartyContactMeches holds the value of the party_contact_meches edge.
	PartyContactMeches []*PartyContactMech `json:"party_contact_meches,omitempty"`
	// PartyRoles holds the value of the party_roles edge.
	PartyRoles []*PartyRole `json:"party_roles,omitempty"`
	// PartyStatuses holds the value of the party_statuses edge.
	PartyStatuses []*PartyStatus `json:"party_statuses,omitempty"`
	// Person holds the value of the person edge.
	Person *Person `json:"person,omitempty"`
	// UserLogins holds the value of the user_logins edge.
	UserLogins []*UserLogin `json:"user_logins,omitempty"`
	// WorkEffortPartyAssignments holds the value of the work_effort_party_assignments edge.
	WorkEffortPartyAssignments []*WorkEffortPartyAssignment `json:"work_effort_party_assignments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [11]bool
}

// PartyTypeOrErr returns the PartyType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyEdges) PartyTypeOrErr() (*PartyType, error) {
	if e.loadedTypes[0] {
		if e.PartyType == nil {
			// The edge party_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: partytype.Label}
		}
		return e.PartyType, nil
	}
	return nil, &NotLoadedError{edge: "party_type"}
}

// CreatedByUserLoginOrErr returns the CreatedByUserLogin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyEdges) CreatedByUserLoginOrErr() (*UserLogin, error) {
	if e.loadedTypes[1] {
		if e.CreatedByUserLogin == nil {
			// The edge created_by_user_login was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userlogin.Label}
		}
		return e.CreatedByUserLogin, nil
	}
	return nil, &NotLoadedError{edge: "created_by_user_login"}
}

// LastModifiedByUserLoginOrErr returns the LastModifiedByUserLogin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyEdges) LastModifiedByUserLoginOrErr() (*UserLogin, error) {
	if e.loadedTypes[2] {
		if e.LastModifiedByUserLogin == nil {
			// The edge last_modified_by_user_login was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userlogin.Label}
		}
		return e.LastModifiedByUserLogin, nil
	}
	return nil, &NotLoadedError{edge: "last_modified_by_user_login"}
}

// StatusItemOrErr returns the StatusItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyEdges) StatusItemOrErr() (*StatusItem, error) {
	if e.loadedTypes[3] {
		if e.StatusItem == nil {
			// The edge status_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: statusitem.Label}
		}
		return e.StatusItem, nil
	}
	return nil, &NotLoadedError{edge: "status_item"}
}

// FixedAssetsOrErr returns the FixedAssets value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) FixedAssetsOrErr() ([]*FixedAsset, error) {
	if e.loadedTypes[4] {
		return e.FixedAssets, nil
	}
	return nil, &NotLoadedError{edge: "fixed_assets"}
}

// PartyContactMechesOrErr returns the PartyContactMeches value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) PartyContactMechesOrErr() ([]*PartyContactMech, error) {
	if e.loadedTypes[5] {
		return e.PartyContactMeches, nil
	}
	return nil, &NotLoadedError{edge: "party_contact_meches"}
}

// PartyRolesOrErr returns the PartyRoles value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) PartyRolesOrErr() ([]*PartyRole, error) {
	if e.loadedTypes[6] {
		return e.PartyRoles, nil
	}
	return nil, &NotLoadedError{edge: "party_roles"}
}

// PartyStatusesOrErr returns the PartyStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) PartyStatusesOrErr() ([]*PartyStatus, error) {
	if e.loadedTypes[7] {
		return e.PartyStatuses, nil
	}
	return nil, &NotLoadedError{edge: "party_statuses"}
}

// PersonOrErr returns the Person value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PartyEdges) PersonOrErr() (*Person, error) {
	if e.loadedTypes[8] {
		if e.Person == nil {
			// The edge person was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: person.Label}
		}
		return e.Person, nil
	}
	return nil, &NotLoadedError{edge: "person"}
}

// UserLoginsOrErr returns the UserLogins value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) UserLoginsOrErr() ([]*UserLogin, error) {
	if e.loadedTypes[9] {
		return e.UserLogins, nil
	}
	return nil, &NotLoadedError{edge: "user_logins"}
}

// WorkEffortPartyAssignmentsOrErr returns the WorkEffortPartyAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e PartyEdges) WorkEffortPartyAssignmentsOrErr() ([]*WorkEffortPartyAssignment, error) {
	if e.loadedTypes[10] {
		return e.WorkEffortPartyAssignments, nil
	}
	return nil, &NotLoadedError{edge: "work_effort_party_assignments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Party) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case party.FieldID, party.FieldExternalID, party.FieldPreferredCurrencyUomID, party.FieldDataSourceID:
			values[i] = new(sql.NullInt64)
		case party.FieldStringRef, party.FieldDescription, party.FieldIsUnread:
			values[i] = new(sql.NullString)
		case party.FieldCreateTime, party.FieldUpdateTime, party.FieldCreatedDate, party.FieldLastModifiedDate:
			values[i] = new(sql.NullTime)
		case party.ForeignKeys[0]: // party_type_parties
			values[i] = new(sql.NullInt64)
		case party.ForeignKeys[1]: // status_item_parties
			values[i] = new(sql.NullInt64)
		case party.ForeignKeys[2]: // user_login_created_by_parties
			values[i] = new(sql.NullInt64)
		case party.ForeignKeys[3]: // user_login_last_modified_by_parties
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Party", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Party fields.
func (pa *Party) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case party.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case party.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pa.CreateTime = value.Time
			}
		case party.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pa.UpdateTime = value.Time
			}
		case party.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				pa.StringRef = value.String
			}
		case party.FieldExternalID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field external_id", values[i])
			} else if value.Valid {
				pa.ExternalID = int(value.Int64)
			}
		case party.FieldPreferredCurrencyUomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_currency_uom_id", values[i])
			} else if value.Valid {
				pa.PreferredCurrencyUomID = int(value.Int64)
			}
		case party.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pa.Description = value.String
			}
		case party.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_date", values[i])
			} else if value.Valid {
				pa.CreatedDate = value.Time
			}
		case party.FieldLastModifiedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified_date", values[i])
			} else if value.Valid {
				pa.LastModifiedDate = value.Time
			}
		case party.FieldDataSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field data_source_id", values[i])
			} else if value.Valid {
				pa.DataSourceID = int(value.Int64)
			}
		case party.FieldIsUnread:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_unread", values[i])
			} else if value.Valid {
				pa.IsUnread = party.IsUnread(value.String)
			}
		case party.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field party_type_parties", value)
			} else if value.Valid {
				pa.party_type_parties = new(int)
				*pa.party_type_parties = int(value.Int64)
			}
		case party.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field status_item_parties", value)
			} else if value.Valid {
				pa.status_item_parties = new(int)
				*pa.status_item_parties = int(value.Int64)
			}
		case party.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_login_created_by_parties", value)
			} else if value.Valid {
				pa.user_login_created_by_parties = new(int)
				*pa.user_login_created_by_parties = int(value.Int64)
			}
		case party.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_login_last_modified_by_parties", value)
			} else if value.Valid {
				pa.user_login_last_modified_by_parties = new(int)
				*pa.user_login_last_modified_by_parties = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPartyType queries the "party_type" edge of the Party entity.
func (pa *Party) QueryPartyType() *PartyTypeQuery {
	return (&PartyClient{config: pa.config}).QueryPartyType(pa)
}

// QueryCreatedByUserLogin queries the "created_by_user_login" edge of the Party entity.
func (pa *Party) QueryCreatedByUserLogin() *UserLoginQuery {
	return (&PartyClient{config: pa.config}).QueryCreatedByUserLogin(pa)
}

// QueryLastModifiedByUserLogin queries the "last_modified_by_user_login" edge of the Party entity.
func (pa *Party) QueryLastModifiedByUserLogin() *UserLoginQuery {
	return (&PartyClient{config: pa.config}).QueryLastModifiedByUserLogin(pa)
}

// QueryStatusItem queries the "status_item" edge of the Party entity.
func (pa *Party) QueryStatusItem() *StatusItemQuery {
	return (&PartyClient{config: pa.config}).QueryStatusItem(pa)
}

// QueryFixedAssets queries the "fixed_assets" edge of the Party entity.
func (pa *Party) QueryFixedAssets() *FixedAssetQuery {
	return (&PartyClient{config: pa.config}).QueryFixedAssets(pa)
}

// QueryPartyContactMeches queries the "party_contact_meches" edge of the Party entity.
func (pa *Party) QueryPartyContactMeches() *PartyContactMechQuery {
	return (&PartyClient{config: pa.config}).QueryPartyContactMeches(pa)
}

// QueryPartyRoles queries the "party_roles" edge of the Party entity.
func (pa *Party) QueryPartyRoles() *PartyRoleQuery {
	return (&PartyClient{config: pa.config}).QueryPartyRoles(pa)
}

// QueryPartyStatuses queries the "party_statuses" edge of the Party entity.
func (pa *Party) QueryPartyStatuses() *PartyStatusQuery {
	return (&PartyClient{config: pa.config}).QueryPartyStatuses(pa)
}

// QueryPerson queries the "person" edge of the Party entity.
func (pa *Party) QueryPerson() *PersonQuery {
	return (&PartyClient{config: pa.config}).QueryPerson(pa)
}

// QueryUserLogins queries the "user_logins" edge of the Party entity.
func (pa *Party) QueryUserLogins() *UserLoginQuery {
	return (&PartyClient{config: pa.config}).QueryUserLogins(pa)
}

// QueryWorkEffortPartyAssignments queries the "work_effort_party_assignments" edge of the Party entity.
func (pa *Party) QueryWorkEffortPartyAssignments() *WorkEffortPartyAssignmentQuery {
	return (&PartyClient{config: pa.config}).QueryWorkEffortPartyAssignments(pa)
}

// Update returns a builder for updating this Party.
// Note that you need to call Party.Unwrap() before calling this method if this Party
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Party) Update() *PartyUpdateOne {
	return (&PartyClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Party entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Party) Unwrap() *Party {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Party is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Party) String() string {
	var builder strings.Builder
	builder.WriteString("Party(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pa.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pa.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(pa.StringRef)
	builder.WriteString(", external_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.ExternalID))
	builder.WriteString(", preferred_currency_uom_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.PreferredCurrencyUomID))
	builder.WriteString(", description=")
	builder.WriteString(pa.Description)
	builder.WriteString(", created_date=")
	builder.WriteString(pa.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", last_modified_date=")
	builder.WriteString(pa.LastModifiedDate.Format(time.ANSIC))
	builder.WriteString(", data_source_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.DataSourceID))
	builder.WriteString(", is_unread=")
	builder.WriteString(fmt.Sprintf("%v", pa.IsUnread))
	builder.WriteByte(')')
	return builder.String()
}

// Parties is a parsable slice of Party.
type Parties []*Party

func (pa Parties) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
