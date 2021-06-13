// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/workeffort/ent/roletype"
)

// RoleType is the model entity for the RoleType schema.
type RoleType struct {
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
	HasTable roletype.HasTable `json:"has_table,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleTypeQuery when eager-loading is set.
	Edges              RoleTypeEdges `json:"edges"`
	role_type_children *int
}

// RoleTypeEdges holds the relations/edges for other nodes in the graph.
type RoleTypeEdges struct {
	// Parent holds the value of the parent edge.
	Parent *RoleType `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*RoleType `json:"children,omitempty"`
	// FixedAssets holds the value of the fixed_assets edge.
	FixedAssets []*FixedAsset `json:"fixed_assets,omitempty"`
	// PartyContactMeches holds the value of the party_contact_meches edge.
	PartyContactMeches []*PartyContactMech `json:"party_contact_meches,omitempty"`
	// ValidFromPartyRelationshipTypes holds the value of the valid_from_party_relationship_types edge.
	ValidFromPartyRelationshipTypes []*PartyRelationshipType `json:"valid_from_party_relationship_types,omitempty"`
	// ValidToPartyRelationshipTypes holds the value of the valid_to_party_relationship_types edge.
	ValidToPartyRelationshipTypes []*PartyRelationshipType `json:"valid_to_party_relationship_types,omitempty"`
	// PartyRoles holds the value of the party_roles edge.
	PartyRoles []*PartyRole `json:"party_roles,omitempty"`
	// ChildRoleTypes holds the value of the child_role_types edge.
	ChildRoleTypes []*RoleType `json:"child_role_types,omitempty"`
	// WorkEffortPartyAssignments holds the value of the work_effort_party_assignments edge.
	WorkEffortPartyAssignments []*WorkEffortPartyAssignment `json:"work_effort_party_assignments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleTypeEdges) ParentOrErr() (*RoleType, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roletype.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) ChildrenOrErr() ([]*RoleType, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// FixedAssetsOrErr returns the FixedAssets value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) FixedAssetsOrErr() ([]*FixedAsset, error) {
	if e.loadedTypes[2] {
		return e.FixedAssets, nil
	}
	return nil, &NotLoadedError{edge: "fixed_assets"}
}

// PartyContactMechesOrErr returns the PartyContactMeches value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) PartyContactMechesOrErr() ([]*PartyContactMech, error) {
	if e.loadedTypes[3] {
		return e.PartyContactMeches, nil
	}
	return nil, &NotLoadedError{edge: "party_contact_meches"}
}

// ValidFromPartyRelationshipTypesOrErr returns the ValidFromPartyRelationshipTypes value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) ValidFromPartyRelationshipTypesOrErr() ([]*PartyRelationshipType, error) {
	if e.loadedTypes[4] {
		return e.ValidFromPartyRelationshipTypes, nil
	}
	return nil, &NotLoadedError{edge: "valid_from_party_relationship_types"}
}

// ValidToPartyRelationshipTypesOrErr returns the ValidToPartyRelationshipTypes value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) ValidToPartyRelationshipTypesOrErr() ([]*PartyRelationshipType, error) {
	if e.loadedTypes[5] {
		return e.ValidToPartyRelationshipTypes, nil
	}
	return nil, &NotLoadedError{edge: "valid_to_party_relationship_types"}
}

// PartyRolesOrErr returns the PartyRoles value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) PartyRolesOrErr() ([]*PartyRole, error) {
	if e.loadedTypes[6] {
		return e.PartyRoles, nil
	}
	return nil, &NotLoadedError{edge: "party_roles"}
}

// ChildRoleTypesOrErr returns the ChildRoleTypes value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) ChildRoleTypesOrErr() ([]*RoleType, error) {
	if e.loadedTypes[7] {
		return e.ChildRoleTypes, nil
	}
	return nil, &NotLoadedError{edge: "child_role_types"}
}

// WorkEffortPartyAssignmentsOrErr returns the WorkEffortPartyAssignments value or an error if the edge
// was not loaded in eager-loading.
func (e RoleTypeEdges) WorkEffortPartyAssignmentsOrErr() ([]*WorkEffortPartyAssignment, error) {
	if e.loadedTypes[8] {
		return e.WorkEffortPartyAssignments, nil
	}
	return nil, &NotLoadedError{edge: "work_effort_party_assignments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleType) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case roletype.FieldID:
			values[i] = new(sql.NullInt64)
		case roletype.FieldStringRef, roletype.FieldHasTable, roletype.FieldDescription:
			values[i] = new(sql.NullString)
		case roletype.FieldCreateTime, roletype.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case roletype.ForeignKeys[0]: // role_type_children
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RoleType", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleType fields.
func (rt *RoleType) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roletype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rt.ID = int(value.Int64)
		case roletype.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rt.CreateTime = value.Time
			}
		case roletype.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				rt.UpdateTime = value.Time
			}
		case roletype.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				rt.StringRef = value.String
			}
		case roletype.FieldHasTable:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field has_table", values[i])
			} else if value.Valid {
				rt.HasTable = roletype.HasTable(value.String)
			}
		case roletype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				rt.Description = value.String
			}
		case roletype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field role_type_children", value)
			} else if value.Valid {
				rt.role_type_children = new(int)
				*rt.role_type_children = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the RoleType entity.
func (rt *RoleType) QueryParent() *RoleTypeQuery {
	return (&RoleTypeClient{config: rt.config}).QueryParent(rt)
}

// QueryChildren queries the "children" edge of the RoleType entity.
func (rt *RoleType) QueryChildren() *RoleTypeQuery {
	return (&RoleTypeClient{config: rt.config}).QueryChildren(rt)
}

// QueryFixedAssets queries the "fixed_assets" edge of the RoleType entity.
func (rt *RoleType) QueryFixedAssets() *FixedAssetQuery {
	return (&RoleTypeClient{config: rt.config}).QueryFixedAssets(rt)
}

// QueryPartyContactMeches queries the "party_contact_meches" edge of the RoleType entity.
func (rt *RoleType) QueryPartyContactMeches() *PartyContactMechQuery {
	return (&RoleTypeClient{config: rt.config}).QueryPartyContactMeches(rt)
}

// QueryValidFromPartyRelationshipTypes queries the "valid_from_party_relationship_types" edge of the RoleType entity.
func (rt *RoleType) QueryValidFromPartyRelationshipTypes() *PartyRelationshipTypeQuery {
	return (&RoleTypeClient{config: rt.config}).QueryValidFromPartyRelationshipTypes(rt)
}

// QueryValidToPartyRelationshipTypes queries the "valid_to_party_relationship_types" edge of the RoleType entity.
func (rt *RoleType) QueryValidToPartyRelationshipTypes() *PartyRelationshipTypeQuery {
	return (&RoleTypeClient{config: rt.config}).QueryValidToPartyRelationshipTypes(rt)
}

// QueryPartyRoles queries the "party_roles" edge of the RoleType entity.
func (rt *RoleType) QueryPartyRoles() *PartyRoleQuery {
	return (&RoleTypeClient{config: rt.config}).QueryPartyRoles(rt)
}

// QueryChildRoleTypes queries the "child_role_types" edge of the RoleType entity.
func (rt *RoleType) QueryChildRoleTypes() *RoleTypeQuery {
	return (&RoleTypeClient{config: rt.config}).QueryChildRoleTypes(rt)
}

// QueryWorkEffortPartyAssignments queries the "work_effort_party_assignments" edge of the RoleType entity.
func (rt *RoleType) QueryWorkEffortPartyAssignments() *WorkEffortPartyAssignmentQuery {
	return (&RoleTypeClient{config: rt.config}).QueryWorkEffortPartyAssignments(rt)
}

// Update returns a builder for updating this RoleType.
// Note that you need to call RoleType.Unwrap() before calling this method if this RoleType
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RoleType) Update() *RoleTypeUpdateOne {
	return (&RoleTypeClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the RoleType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rt *RoleType) Unwrap() *RoleType {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleType is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RoleType) String() string {
	var builder strings.Builder
	builder.WriteString("RoleType(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(rt.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(rt.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(rt.StringRef)
	builder.WriteString(", has_table=")
	builder.WriteString(fmt.Sprintf("%v", rt.HasTable))
	builder.WriteString(", description=")
	builder.WriteString(rt.Description)
	builder.WriteByte(')')
	return builder.String()
}

// RoleTypes is a parsable slice of RoleType.
type RoleTypes []*RoleType

func (rt RoleTypes) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
