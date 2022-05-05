package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// LineUser holds the schema definition for the LineUser entity.
type LineUser struct {
	ent.Schema
}

// Fields of the LineUser.
func (LineUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(64).Unique().NotEmpty().Immutable(),
		field.Bool("is_active").Default(true),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the LineUser.
func (LineUser) Edges() []ent.Edge {
	return nil
}
