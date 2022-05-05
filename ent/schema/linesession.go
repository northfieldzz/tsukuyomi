package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// LineSession holds the schema definition for the LineSession entity.
type LineSession struct {
	ent.Schema
}

// Fields of the LineSession.
func (LineSession) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("type").Immutable(),
		field.String("user_id").MaxLen(64).NotEmpty().Immutable(),
		field.String("group_id").MaxLen(64).Immutable(),
		field.String("room_id").MaxLen(64).Immutable(),
		field.Time("create_at").Default(time.Now).Immutable(),
		field.Time("update_at").Default(time.Now),
	}
}

// Edges of the LineSession.
func (LineSession) Edges() []ent.Edge {
	return nil
}
