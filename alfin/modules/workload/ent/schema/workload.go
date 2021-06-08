package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
    // "entgo.io/ent/schema/index"
    "entgo.io/ent/schema/mixin"
    "time"
)


type WorkloadFeatureApplType struct {
    ent.Schema
}

// Fields of the WorkloadFeatureApplType.
func (WorkloadFeatureApplType) Fields() []ent.Field {
    return []ent.Field{
        field.String("description").Optional(),
    }
}

func (WorkloadFeatureApplType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the WorkloadFeatureApplType.
//  one: WorkloadFeatureApplType
//  many: WorkloadFeatureAppl
func (WorkloadFeatureApplType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", WorkloadFeatureApplType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique(),
                edge.To("workload_feature_appls", WorkloadFeatureAppl.Type),
    }
}


type Workload struct {
    ent.Schema
}

// Fields of the Workload.
func (Workload) Fields() []ent.Field {
    return []ent.Field{
        field.Int("status_id").Optional(),
        field.String("workload_name").Optional(),
        field.String("description").Optional(),
        field.String("long_description").Optional(),
        field.String("comments").Optional(),
        field.Int("workload_size").Optional(),
        field.Time("workload_date").
            Default(time.Now).Optional(),
        field.Time("another_date").
            Default(time.Now).Optional(),
        field.String("another_text").Optional(),
    }
}

func (Workload) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the Workload.
//  one: WorkloadType
//  one: StatusItem
//  many: WorkloadFeatureAppl
//  many: WorkloadItem
//  many: WorkloadStatus
func (Workload) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("workload_type", WorkloadType.Type).Ref("workloads").
                    // Bind the "workloadTypeId" field to this edge.
                    // Field("workload_type_id").
                    Unique(),
                edge.To("workload_feature_appls", WorkloadFeatureAppl.Type),
                edge.To("workload_items", WorkloadItem.Type),
                edge.To("workload_statuses", WorkloadStatus.Type),
    }
}


type WorkloadItem struct {
    ent.Schema
}

// Fields of the WorkloadItem.
func (WorkloadItem) Fields() []ent.Field {
    return []ent.Field{
        field.Int("workload_item_seq_id"),
        field.String("description").Optional(),
        field.Float("amount").Optional(),
        field.Int("amount_uom_id").Optional(),
    }
}

func (WorkloadItem) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (WorkloadItem) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("workload_id", "workload_item_seq_id").
            Unique(),
    }
}

*/



// Edges of the WorkloadItem.
//  one: AmountUom
//  one: Workload
func (WorkloadItem) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("workload", Workload.Type).Ref("workload_items").
                    // Bind the "workloadId" field to this edge.
                    // Field("workload_id").
                    Unique(),
    }
}


type WorkloadStatus struct {
    ent.Schema
}

// Fields of the WorkloadStatus.
func (WorkloadStatus) Fields() []ent.Field {
    return []ent.Field{
        field.Time("status_date").
            Default(time.Now),
        field.Time("status_end_date").
            Default(time.Now).Optional(),
        field.String("change_by_user_login_id").Optional(),
        field.Int("status_id").Optional(),
    }
}

func (WorkloadStatus) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (WorkloadStatus) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("workload_id", "status_date").
            Unique(),
    }
}

*/



// Edges of the WorkloadStatus.
//  one: Workload
//  one: StatusItem
//  one: UserLogin
func (WorkloadStatus) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("workload", Workload.Type).Ref("workload_statuses").
                    // Bind the "workloadId" field to this edge.
                    // Field("workload_id").
                    Unique(),
    }
}


type WorkloadType struct {
    ent.Schema
}

// Fields of the WorkloadType.
func (WorkloadType) Fields() []ent.Field {
    return []ent.Field{
        field.String("description").Optional(),
    }
}

func (WorkloadType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the WorkloadType.
//  one: WorkloadType
//  many: Workload
func (WorkloadType) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("children", WorkloadType.Type).
                    From("parent").
                    // Bind the "parentTypeId" field to this edge.
                    // Field("parent_type_id").
                    Unique(),
                edge.To("workloads", Workload.Type),
    }
}


type WorkloadFeature struct {
    ent.Schema
}

// Fields of the WorkloadFeature.
func (WorkloadFeature) Fields() []ent.Field {
    return []ent.Field{
        field.Int("feature_source_enum_id").Optional(),
        field.String("description").Optional(),
    }
}

func (WorkloadFeature) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
*/



// Edges of the WorkloadFeature.
//  one: Enumeration
//  many: WorkloadFeatureAppl
func (WorkloadFeature) Edges() []ent.Edge {
    return []ent.Edge{
                edge.To("workload_feature_appls", WorkloadFeatureAppl.Type),
    }
}


type WorkloadFeatureAppl struct {
    ent.Schema
}

// Fields of the WorkloadFeatureAppl.
func (WorkloadFeatureAppl) Fields() []ent.Field {
    return []ent.Field{
        field.Time("from_date").
            Default(time.Now),
        field.Time("thru_date").
            Default(time.Now).Optional(),
        field.Int("sequence_num").Optional(),
    }
}

func (WorkloadFeatureAppl) Mixin() []ent.Mixin {
    return []ent.Mixin{
        mixin.Time{},
    }
}

/*
func (WorkloadFeatureAppl) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("workload_id", "workload_feature_id", "from_date").
            Unique(),
    }
}

*/



// Edges of the WorkloadFeatureAppl.
//  one: Workload
//  one: WorkloadFeature
//  one: WorkloadFeatureApplType
func (WorkloadFeatureAppl) Edges() []ent.Edge {
    return []ent.Edge{
                edge.From("workload", Workload.Type).Ref("workload_feature_appls").
                    // Bind the "workloadId" field to this edge.
                    // Field("workload_id").
                    Unique(),
                edge.From("workload_feature", WorkloadFeature.Type).Ref("workload_feature_appls").
                    // Bind the "workloadFeatureId" field to this edge.
                    // Field("workload_feature_id").
                    Unique(),
                edge.From("workload_feature_appl_type", WorkloadFeatureApplType.Type).Ref("workload_feature_appls").
                    // Bind the "workloadFeatureApplTypeId" field to this edge.
                    // Field("workload_feature_appl_type_id").
                    Unique(),
    }
}
