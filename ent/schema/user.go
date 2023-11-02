package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("email").
			Unique(),
		field.String("password").
			Sensitive(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Bool("superuser").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("answers", Answer.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("tags", Tag.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
