package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PollOption holds the schema definition for the PollOption entity.
type PollOption struct {
	ent.Schema
}

// Fields of the PollOption.
func (PollOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Int("votes").Default(0),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the PollOption.
func (PollOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).
			Ref("options").
			Unique(),
	}
}
