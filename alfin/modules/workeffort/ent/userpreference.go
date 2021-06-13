// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userlogin"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/userpreference"
)

// UserPreference is the model entity for the UserPreference schema.
type UserPreference struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// UserPrefTypeID holds the value of the "user_pref_type_id" field.
	UserPrefTypeID string `json:"user_pref_type_id,omitempty"`
	// UserPrefGroupTypeID holds the value of the "user_pref_group_type_id" field.
	UserPrefGroupTypeID string `json:"user_pref_group_type_id,omitempty"`
	// UserPrefValue holds the value of the "user_pref_value" field.
	UserPrefValue string `json:"user_pref_value,omitempty"`
	// UserPrefDataType holds the value of the "user_pref_data_type" field.
	UserPrefDataType string `json:"user_pref_data_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserPreferenceQuery when eager-loading is set.
	Edges                       UserPreferenceEdges `json:"edges"`
	user_login_user_preferences *int
}

// UserPreferenceEdges holds the relations/edges for other nodes in the graph.
type UserPreferenceEdges struct {
	// UserLogin holds the value of the user_login edge.
	UserLogin *UserLogin `json:"user_login,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserLoginOrErr returns the UserLogin value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserPreferenceEdges) UserLoginOrErr() (*UserLogin, error) {
	if e.loadedTypes[0] {
		if e.UserLogin == nil {
			// The edge user_login was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: userlogin.Label}
		}
		return e.UserLogin, nil
	}
	return nil, &NotLoadedError{edge: "user_login"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserPreference) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userpreference.FieldID:
			values[i] = new(sql.NullInt64)
		case userpreference.FieldStringRef, userpreference.FieldUserPrefTypeID, userpreference.FieldUserPrefGroupTypeID, userpreference.FieldUserPrefValue, userpreference.FieldUserPrefDataType:
			values[i] = new(sql.NullString)
		case userpreference.FieldCreateTime, userpreference.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case userpreference.ForeignKeys[0]: // user_login_user_preferences
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserPreference", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserPreference fields.
func (up *UserPreference) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userpreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case userpreference.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				up.CreateTime = value.Time
			}
		case userpreference.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				up.UpdateTime = value.Time
			}
		case userpreference.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				up.StringRef = value.String
			}
		case userpreference.FieldUserPrefTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_pref_type_id", values[i])
			} else if value.Valid {
				up.UserPrefTypeID = value.String
			}
		case userpreference.FieldUserPrefGroupTypeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_pref_group_type_id", values[i])
			} else if value.Valid {
				up.UserPrefGroupTypeID = value.String
			}
		case userpreference.FieldUserPrefValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_pref_value", values[i])
			} else if value.Valid {
				up.UserPrefValue = value.String
			}
		case userpreference.FieldUserPrefDataType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_pref_data_type", values[i])
			} else if value.Valid {
				up.UserPrefDataType = value.String
			}
		case userpreference.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_login_user_preferences", value)
			} else if value.Valid {
				up.user_login_user_preferences = new(int)
				*up.user_login_user_preferences = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUserLogin queries the "user_login" edge of the UserPreference entity.
func (up *UserPreference) QueryUserLogin() *UserLoginQuery {
	return (&UserPreferenceClient{config: up.config}).QueryUserLogin(up)
}

// Update returns a builder for updating this UserPreference.
// Note that you need to call UserPreference.Unwrap() before calling this method if this UserPreference
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UserPreference) Update() *UserPreferenceUpdateOne {
	return (&UserPreferenceClient{config: up.config}).UpdateOne(up)
}

// Unwrap unwraps the UserPreference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UserPreference) Unwrap() *UserPreference {
	tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserPreference is not a transactional entity")
	}
	up.config.driver = tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UserPreference) String() string {
	var builder strings.Builder
	builder.WriteString("UserPreference(")
	builder.WriteString(fmt.Sprintf("id=%v", up.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(up.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(up.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(up.StringRef)
	builder.WriteString(", user_pref_type_id=")
	builder.WriteString(up.UserPrefTypeID)
	builder.WriteString(", user_pref_group_type_id=")
	builder.WriteString(up.UserPrefGroupTypeID)
	builder.WriteString(", user_pref_value=")
	builder.WriteString(up.UserPrefValue)
	builder.WriteString(", user_pref_data_type=")
	builder.WriteString(up.UserPrefDataType)
	builder.WriteByte(')')
	return builder.String()
}

// UserPreferences is a parsable slice of UserPreference.
type UserPreferences []*UserPreference

func (up UserPreferences) config(cfg config) {
	for _i := range up {
		up[_i].config = cfg
	}
}
