// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeature"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeaturecategory"
	"github.com/samlet/petrel/alfin/modules/catalog/ent/productfeaturetype"
)

// ProductFeature is the model entity for the ProductFeature schema.
type ProductFeature struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// StringRef holds the value of the "string_ref" field.
	StringRef string `json:"string_ref,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// UomID holds the value of the "uom_id" field.
	UomID int `json:"uom_id,omitempty"`
	// NumberSpecified holds the value of the "number_specified" field.
	NumberSpecified float64 `json:"number_specified,omitempty"`
	// DefaultAmount holds the value of the "default_amount" field.
	DefaultAmount float64 `json:"default_amount,omitempty"`
	// DefaultSequenceNum holds the value of the "default_sequence_num" field.
	DefaultSequenceNum int `json:"default_sequence_num,omitempty"`
	// Abbrev holds the value of the "abbrev" field.
	Abbrev int `json:"abbrev,omitempty"`
	// IDCode holds the value of the "id_code" field.
	IDCode string `json:"id_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductFeatureQuery when eager-loading is set.
	Edges                                     ProductFeatureEdges `json:"edges"`
	product_feature_category_product_features *int
	product_feature_type_product_features     *int
}

// ProductFeatureEdges holds the relations/edges for other nodes in the graph.
type ProductFeatureEdges struct {
	// ProductFeatureCategory holds the value of the product_feature_category edge.
	ProductFeatureCategory *ProductFeatureCategory `json:"product_feature_category,omitempty"`
	// ProductFeatureType holds the value of the product_feature_type edge.
	ProductFeatureType *ProductFeatureType `json:"product_feature_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductFeatureCategoryOrErr returns the ProductFeatureCategory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductFeatureEdges) ProductFeatureCategoryOrErr() (*ProductFeatureCategory, error) {
	if e.loadedTypes[0] {
		if e.ProductFeatureCategory == nil {
			// The edge product_feature_category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: productfeaturecategory.Label}
		}
		return e.ProductFeatureCategory, nil
	}
	return nil, &NotLoadedError{edge: "product_feature_category"}
}

// ProductFeatureTypeOrErr returns the ProductFeatureType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductFeatureEdges) ProductFeatureTypeOrErr() (*ProductFeatureType, error) {
	if e.loadedTypes[1] {
		if e.ProductFeatureType == nil {
			// The edge product_feature_type was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: productfeaturetype.Label}
		}
		return e.ProductFeatureType, nil
	}
	return nil, &NotLoadedError{edge: "product_feature_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductFeature) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case productfeature.FieldNumberSpecified, productfeature.FieldDefaultAmount:
			values[i] = new(sql.NullFloat64)
		case productfeature.FieldID, productfeature.FieldUomID, productfeature.FieldDefaultSequenceNum, productfeature.FieldAbbrev:
			values[i] = new(sql.NullInt64)
		case productfeature.FieldStringRef, productfeature.FieldDescription, productfeature.FieldIDCode:
			values[i] = new(sql.NullString)
		case productfeature.FieldCreateTime, productfeature.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case productfeature.ForeignKeys[0]: // product_feature_category_product_features
			values[i] = new(sql.NullInt64)
		case productfeature.ForeignKeys[1]: // product_feature_type_product_features
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductFeature", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductFeature fields.
func (pf *ProductFeature) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productfeature.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pf.ID = int(value.Int64)
		case productfeature.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pf.CreateTime = value.Time
			}
		case productfeature.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pf.UpdateTime = value.Time
			}
		case productfeature.FieldStringRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field string_ref", values[i])
			} else if value.Valid {
				pf.StringRef = value.String
			}
		case productfeature.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pf.Description = value.String
			}
		case productfeature.FieldUomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uom_id", values[i])
			} else if value.Valid {
				pf.UomID = int(value.Int64)
			}
		case productfeature.FieldNumberSpecified:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field number_specified", values[i])
			} else if value.Valid {
				pf.NumberSpecified = value.Float64
			}
		case productfeature.FieldDefaultAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field default_amount", values[i])
			} else if value.Valid {
				pf.DefaultAmount = value.Float64
			}
		case productfeature.FieldDefaultSequenceNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field default_sequence_num", values[i])
			} else if value.Valid {
				pf.DefaultSequenceNum = int(value.Int64)
			}
		case productfeature.FieldAbbrev:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field abbrev", values[i])
			} else if value.Valid {
				pf.Abbrev = int(value.Int64)
			}
		case productfeature.FieldIDCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id_code", values[i])
			} else if value.Valid {
				pf.IDCode = value.String
			}
		case productfeature.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_feature_category_product_features", value)
			} else if value.Valid {
				pf.product_feature_category_product_features = new(int)
				*pf.product_feature_category_product_features = int(value.Int64)
			}
		case productfeature.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field product_feature_type_product_features", value)
			} else if value.Valid {
				pf.product_feature_type_product_features = new(int)
				*pf.product_feature_type_product_features = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProductFeatureCategory queries the "product_feature_category" edge of the ProductFeature entity.
func (pf *ProductFeature) QueryProductFeatureCategory() *ProductFeatureCategoryQuery {
	return (&ProductFeatureClient{config: pf.config}).QueryProductFeatureCategory(pf)
}

// QueryProductFeatureType queries the "product_feature_type" edge of the ProductFeature entity.
func (pf *ProductFeature) QueryProductFeatureType() *ProductFeatureTypeQuery {
	return (&ProductFeatureClient{config: pf.config}).QueryProductFeatureType(pf)
}

// Update returns a builder for updating this ProductFeature.
// Note that you need to call ProductFeature.Unwrap() before calling this method if this ProductFeature
// was returned from a transaction, and the transaction was committed or rolled back.
func (pf *ProductFeature) Update() *ProductFeatureUpdateOne {
	return (&ProductFeatureClient{config: pf.config}).UpdateOne(pf)
}

// Unwrap unwraps the ProductFeature entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pf *ProductFeature) Unwrap() *ProductFeature {
	tx, ok := pf.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductFeature is not a transactional entity")
	}
	pf.config.driver = tx.drv
	return pf
}

// String implements the fmt.Stringer.
func (pf *ProductFeature) String() string {
	var builder strings.Builder
	builder.WriteString("ProductFeature(")
	builder.WriteString(fmt.Sprintf("id=%v", pf.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(pf.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pf.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", string_ref=")
	builder.WriteString(pf.StringRef)
	builder.WriteString(", description=")
	builder.WriteString(pf.Description)
	builder.WriteString(", uom_id=")
	builder.WriteString(fmt.Sprintf("%v", pf.UomID))
	builder.WriteString(", number_specified=")
	builder.WriteString(fmt.Sprintf("%v", pf.NumberSpecified))
	builder.WriteString(", default_amount=")
	builder.WriteString(fmt.Sprintf("%v", pf.DefaultAmount))
	builder.WriteString(", default_sequence_num=")
	builder.WriteString(fmt.Sprintf("%v", pf.DefaultSequenceNum))
	builder.WriteString(", abbrev=")
	builder.WriteString(fmt.Sprintf("%v", pf.Abbrev))
	builder.WriteString(", id_code=")
	builder.WriteString(pf.IDCode)
	builder.WriteByte(')')
	return builder.String()
}

// ProductFeatures is a parsable slice of ProductFeature.
type ProductFeatures []*ProductFeature

func (pf ProductFeatures) config(cfg config) {
	for _i := range pf {
		pf[_i].config = cfg
	}
}