package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.String("question"),
		field.String("created_by"),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("options", PollOption.Type),
	}
}
