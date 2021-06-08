// Code generated by entc, DO NOT EDIT.

package workloadfeature

import (
	"time"
)

const (
	// Label holds the string label denoting the workloadfeature type in the database.
	Label = "workload_feature"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldFeatureSourceEnumID holds the string denoting the feature_source_enum_id field in the database.
	FieldFeatureSourceEnumID = "feature_source_enum_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeWorkloadFeatureAppls holds the string denoting the workload_feature_appls edge name in mutations.
	EdgeWorkloadFeatureAppls = "workload_feature_appls"
	// Table holds the table name of the workloadfeature in the database.
	Table = "workload_features"
	// WorkloadFeatureApplsTable is the table the holds the workload_feature_appls relation/edge.
	WorkloadFeatureApplsTable = "workload_feature_appls"
	// WorkloadFeatureApplsInverseTable is the table name for the WorkloadFeatureAppl entity.
	// It exists in this package in order to avoid circular dependency with the "workloadfeatureappl" package.
	WorkloadFeatureApplsInverseTable = "workload_feature_appls"
	// WorkloadFeatureApplsColumn is the table column denoting the workload_feature_appls relation/edge.
	WorkloadFeatureApplsColumn = "workload_feature_workload_feature_appls"
)

// Columns holds all SQL columns for workloadfeature fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldFeatureSourceEnumID,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
