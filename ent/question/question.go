// Code generated by ent, DO NOT EDIT.

package question

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the question type in the database.
	Label = "question"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeAnswers holds the string denoting the answers edge name in mutations.
	EdgeAnswers = "answers"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// Table holds the table name of the question in the database.
	Table = "questions"
	// AnswersTable is the table that holds the answers relation/edge.
	AnswersTable = "answers"
	// AnswersInverseTable is the table name for the Answer entity.
	// It exists in this package in order to avoid circular dependency with the "answer" package.
	AnswersInverseTable = "answers"
	// AnswersColumn is the table column denoting the answers relation/edge.
	AnswersColumn = "question_answers"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "questions"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_questions"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "tag_questions"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
)

// Columns holds all SQL columns for question fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "questions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_questions",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"tag_id", "question_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Question queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByAnswersCount orders the results by answers count.
func ByAnswersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnswersStep(), opts...)
	}
}

// ByAnswers orders the results by answers terms.
func ByAnswers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnswersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAnswersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnswersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AnswersTable, AnswersColumn),
	)
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TagsTable, TagsPrimaryKey...),
	)
}
