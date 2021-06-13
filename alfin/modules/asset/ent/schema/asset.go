package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/samlet/petrel/alfin/schemamixins"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pkg", WorkloadPkg.Type).
			Ref("assets").
			Unique()}
}

func (Asset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		schemamixins.StringRefMixin{},
	}
}

