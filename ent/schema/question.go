package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Question holds the schema definition for the Question entity.
type Question struct {
	ent.Schema
}

// Fields of the Question.
func (Question) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Text("content"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
	}
}

// Edges of the Question.
func (Question) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("answers", Answer.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("author", User.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("questions").
			Unique(),
		edge.From("tags", Tag.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Ref("questions"),
	}
}
