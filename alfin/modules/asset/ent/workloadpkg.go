// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/asset/ent/workloadpkg"
)

// WorkloadPkg is the model entity for the WorkloadPkg schema.
type WorkloadPkg struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkloadPkgQuery when eager-loading is set.
	Edges WorkloadPkgEdges `json:"edges"`
}

// WorkloadPkgEdges holds the relations/edges for other nodes in the graph.
type WorkloadPkgEdges struct {
	// Assets holds the value of the assets edge.
	Assets []*Asset `json:"assets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssetsOrErr returns the Assets value or an error if the edge
// was not loaded in eager-loading.
func (e WorkloadPkgEdges) AssetsOrErr() ([]*Asset, error) {
	if e.loadedTypes[0] {
		return e.Assets, nil
	}
	return nil, &NotLoadedError{edge: "assets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkloadPkg) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case workloadpkg.FieldID:
			values[i] = new(sql.NullInt64)
		case workloadpkg.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WorkloadPkg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkloadPkg fields.
func (wp *WorkloadPkg) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workloadpkg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wp.ID = int(value.Int64)
		case workloadpkg.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				wp.Name = value.String
			}
		}
	}
	return nil
}

// QueryAssets queries the "assets" edge of the WorkloadPkg entity.
func (wp *WorkloadPkg) QueryAssets() *AssetQuery {
	return (&WorkloadPkgClient{config: wp.config}).QueryAssets(wp)
}

// Update returns a builder for updating this WorkloadPkg.
// Note that you need to call WorkloadPkg.Unwrap() before calling this method if this WorkloadPkg
// was returned from a transaction, and the transaction was committed or rolled back.
func (wp *WorkloadPkg) Update() *WorkloadPkgUpdateOne {
	return (&WorkloadPkgClient{config: wp.config}).UpdateOne(wp)
}

// Unwrap unwraps the WorkloadPkg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wp *WorkloadPkg) Unwrap() *WorkloadPkg {
	tx, ok := wp.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkloadPkg is not a transactional entity")
	}
	wp.config.driver = tx.drv
	return wp
}

// String implements the fmt.Stringer.
func (wp *WorkloadPkg) String() string {
	var builder strings.Builder
	builder.WriteString("WorkloadPkg(")
	builder.WriteString(fmt.Sprintf("id=%v", wp.ID))
	builder.WriteString(", name=")
	builder.WriteString(wp.Name)
	builder.WriteByte(')')
	return builder.String()
}

// WorkloadPkgs is a parsable slice of WorkloadPkg.
type WorkloadPkgs []*WorkloadPkg

func (wp WorkloadPkgs) config(cfg config) {
	for _i := range wp {
		wp[_i].config = cfg
	}
}
