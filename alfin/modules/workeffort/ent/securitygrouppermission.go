// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygroup"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/securitygrouppermission"
)

// SecurityGroupPermission is the model entity for the SecurityGroupPermission schema.
type SecurityGroupPermission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// PermissionID holds the value of the "permission_id" field.
	PermissionID string `json:"permission_id,omitempty"`
	// FromDate holds the value of the "from_date" field.
	FromDate time.Time `json:"from_date,omitempty"`
	// ThruDate holds the value of the "thru_date" field.
	ThruDate time.Time `json:"thru_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SecurityGroupPermissionQuery when eager-loading is set.
	Edges                                                SecurityGroupPermissionEdges `json:"edges"`
	security_group_security_group_permissions            *int
	user_login_security_group_security_group_permissions *int
}

// SecurityGroupPermissionEdges holds the relations/edges for other nodes in the graph.
type SecurityGroupPermissionEdges struct {
	// SecurityGroup holds the value of the security_group edge.
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SecurityGroupOrErr returns the SecurityGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SecurityGroupPermissionEdges) SecurityGroupOrErr() (*SecurityGroup, error) {
	if e.loadedTypes[0] {
		if e.SecurityGroup == nil {
			// The edge security_group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: securitygroup.Label}
		}
		return e.SecurityGroup, nil
	}
	return nil, &NotLoadedError{edge: "security_group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SecurityGroupPermission) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case securitygrouppermission.FieldID:
			values[i] = new(sql.NullInt64)
		case securitygrouppermission.FieldStringRef, securitygrouppermission.FieldPermissionID:
			values[i] = new(sql.NullString)
		case securitygrouppermission.FieldCreateTime, securitygrouppermission.FieldUpdateTime, securitygrouppermission.FieldFromDate, securitygrouppermission.FieldThruDate:
			values[i] = new(sql.NullTime)
		case securitygrouppermission.ForeignKeys[0]: // security_group_security_group_permissions
			values[i] = new(sql.NullInt64)
		case securitygrouppermission.ForeignKeys[1]: // user_login_security_group_security_group_permissions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SecurityGroupPermission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SecurityGroupPermission fields.
func (sgp *SecurityGroupPermission) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case securitygrouppermission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sgp.ID = int(value.Int64)
		case securitygrouppermission.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sgp.CreateTime = value.Time
			}
		case securitygrouppermission.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				sgp.UpdateTime = value.Time
			}
		case securitygrouppermission.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				sgp.StringRef = value.String
			}
		case securitygrouppermission.FieldPermissionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permission_id", values[i])
			} else if value.Valid {
				sgp.PermissionID = value.String
			}
		case securitygrouppermission.FieldFromDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from_date", values[i])
			} else if value.Valid {
				sgp.FromDate = value.Time
			}
		case securitygrouppermission.FieldThruDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field thru_date", values[i])
			} else if value.Valid {
				sgp.ThruDate = value.Time
			}
		case securitygrouppermission.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field security_group_security_group_permissions", value)
			} else if value.Valid {
				sgp.security_group_security_group_permissions = new(int)
				*sgp.security_group_security_group_permissions = int(value.Int64)
			}
		case securitygrouppermission.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_login_security_group_security_group_permissions", value)
			} else if value.Valid {
				sgp.user_login_security_group_security_group_permissions = new(int)
				*sgp.user_login_security_group_security_group_permissions = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySecurityGroup queries the "security_group" edge of the SecurityGroupPermission entity.
func (sgp *SecurityGroupPermission) QuerySecurityGroup() *SecurityGroupQuery {
	return (&SecurityGroupPermissionClient{config: sgp.config}).QuerySecurityGroup(sgp)
}

// Update returns a builder for updating this SecurityGroupPermission.
// Note that you need to call SecurityGroupPermission.Unwrap() before calling this method if this SecurityGroupPermission
// was returned from a transaction, and the transaction was committed or rolled back.
func (sgp *SecurityGroupPermission) Update() *SecurityGroupPermissionUpdateOne {
	return (&SecurityGroupPermissionClient{config: sgp.config}).UpdateOne(sgp)
}

// Unwrap unwraps the SecurityGroupPermission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sgp *SecurityGroupPermission) Unwrap() *SecurityGroupPermission {
	tx, ok := sgp.config.driver.(*txDriver)
	if !ok {
		panic("ent: SecurityGroupPermission is not a transactional entity")
	}
	sgp.config.driver = tx.drv
	return sgp
}

// String implements the fmt.Stringer.
func (sgp *SecurityGroupPermission) String() string {
	var builder strings.Builder
	builder.WriteString("SecurityGroupPermission(")
	builder.WriteString(fmt.Sprintf("id=%v", sgp.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(sgp.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(sgp.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(sgp.StringRef)
	builder.WriteString(", permission_id=")
	builder.WriteString(sgp.PermissionID)
	builder.WriteString(", from_date=")
	builder.WriteString(sgp.FromDate.Format(time.ANSIC))
	builder.WriteString(", thru_date=")
	builder.WriteString(sgp.ThruDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SecurityGroupPermissions is a parsable slice of SecurityGroupPermission.
type SecurityGroupPermissions []*SecurityGroupPermission

func (sgp SecurityGroupPermissions) config(cfg config) {
	for _i := range sgp {
		sgp[_i].config = cfg
	}
}
