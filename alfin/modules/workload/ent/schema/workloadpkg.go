package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WorkloadPkg holds the schema definition for the WorkloadPkg entity.
type WorkloadPkg struct {
	ent.Schema
}

// Fields of the WorkloadPkg.
func (WorkloadPkg) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("workload"),
	}
}

// Edges of the WorkloadPkg.
func (WorkloadPkg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("assets", Asset.Type),
	}
}

