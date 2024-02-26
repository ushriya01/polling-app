package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Int("poll_id"),
		field.String("user_id"),
		field.Int("option_id"),
	}
}
