package service

import (
	"testing"

	"github.com/sinisaos/fiber-ent-admin/ent/enttest"
	"github.com/sinisaos/fiber-ent-admin/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestQuestionService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	userService := NewUserService(client)
	questionService := NewQuestionService(client)
	tagService := NewTagService(client)
	answerService := NewAnswerService(client)
	// Insert user
	u, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)
	// Insert tags
	firstTag, err := tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag1",
	})
	assert.NoError(t, err)

	secondTag, err := tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag2",
	})
	assert.NoError(t, err)
	// Tags slice
	var tags []int
	tags = append(tags, firstTag.ID, secondTag.ID)

	// Insert question
	q, err := questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion1",
		Content: "Content od question one",
		Author:  u.ID,
		Tags:    tags,
	})
	assert.Equal(t, q.ID, 1)
	assert.NoError(t, err)

	_, err = questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion2",
		Content: "Content od question two",
		Author:  u.ID,
		Tags:    tags,
	})
	assert.NoError(t, err)

	// Check if pagination works
	pageOne, _ := questionService.GetAllQuestions(0, 1, "", "", "", "")
	assert.Len(t, pageOne, 1)
	pageTwo, _ := questionService.GetAllQuestions(1, 2, "", "", "", "")
	assert.Len(t, pageTwo, 1)
	allPages, _ := questionService.GetAllQuestions(0, 2, "", "", "", "")
	assert.Len(t, allPages, 2)

	// Check if sorting works
	sortDesc, _ := questionService.GetAllQuestions(0, 2, "id", "DESC", "", "")
	assert.Equal(t, sortDesc[0].ID, 2)
	sortAsc, _ := questionService.GetAllQuestions(0, 2, "id", "ASC", "", "")
	assert.Equal(t, sortAsc[0].ID, 1)

	// Check if filtering works
	filterTitle, _ := questionService.GetAllQuestions(0, 2, "", "", "1", "")
	assert.Equal(t, filterTitle[0].Title, "TestQuestion1")
	filterTitleContent, _ := questionService.GetAllQuestions(0, 2, "", "", "1", "one")
	assert.Equal(t, filterTitleContent[0].Content, "Content od question one")

	// Check empty data result
	emptyDataResult, _ := questionService.GetAllQuestions(10, 20, "", "", "", "")
	assert.Len(t, emptyDataResult, 0)

	// Records total for React admin pagination
	total, _ := questionService.CountQuestions()
	assert.Equal(t, total, 2)

	// Single question
	resultSingleQuestion, _ := questionService.GetQuestion(1)
	assert.Contains(t, resultSingleQuestion.Title, "TestQuestion1")

	// Return error if Question does not exist
	_, err = questionService.GetQuestion(10)
	assert.Error(t, err)

	// Insert updated tags
	thirdTag, err := tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag3",
	})
	assert.NoError(t, err)

	fourthTag, err := tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag4",
	})
	assert.NoError(t, err)
	// Update tags slice
	var updatedTags []int
	updatedTags = append(updatedTags, thirdTag.ID, fourthTag.ID)

	// Update question if exist
	_, err = questionService.UpdateQuestion(2, &model.UpdateQuestionInput{
		Title:   "TestQuestion2Updated",
		Content: "Updated content od question two",
		Author:  u.ID,
		Tags:    updatedTags,
	})
	assert.NoError(t, err)

	// Check error if question does not exists
	_, err = questionService.UpdateQuestion(3, &model.UpdateQuestionInput{
		Title:   "TestQuestion3Updated",
		Content: "Updated content od question three",
		Author:  u.ID,
		Tags:    updatedTags,
	})
	assert.Error(t, err)

	// Delete question
	err = questionService.DeleteQuestion(2)
	assert.NoError(t, err)

	// Check error if question does not exists
	err = questionService.DeleteQuestion(10)
	assert.Error(t, err)

	// Check if question is deleted
	deletedQuestionResult, _ := questionService.GetAllQuestions(0, 1, "", "", "", "")
	assert.Len(t, deletedQuestionResult, 1)

	// Insert answer for checking question answers
	_, err = answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer one",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.NoError(t, err)

	// Checking question author, tags and answers
	questionAuthor, _ := questionService.GetQuestionAuthor(q.ID)
	assert.Equal(t, questionAuthor.Edges.Author.ID, 1)
	assert.Equal(t, questionAuthor.Edges.Author.Username, "TestUser1")

	questionAnswers, _ := questionService.GetQuestionAnswers(q.ID)
	assert.Equal(t, questionAnswers.Edges.Answers[0].ID, 1)
	assert.Equal(t, questionAnswers.Edges.Answers[0].Content, "Content of Answer one")
	assert.Len(t, questionAnswers.Edges.Answers, 1)

	questionTags, _ := questionService.GetQuestionTags(q.ID)
	assert.Equal(t, questionTags.Edges.Tags[0].ID, 1)
	assert.Equal(t, questionTags.Edges.Tags[0].Name, "TestTag1")
	assert.Len(t, questionTags.Edges.Tags, 2)
}
