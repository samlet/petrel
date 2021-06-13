package schemamixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type StringRefMixin struct{
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (StringRefMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("string_ref").
			Optional(),
	}
}
