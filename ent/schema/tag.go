package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("questions", Question.Type),
		// edge.From("questions", Question.Type).
		// 	Annotations(entsql.OnDelete(entsql.Cascade)).
		// 	Ref("tags").
		// 	Through("tag_question", QuestionTag.Type),
	}
}
