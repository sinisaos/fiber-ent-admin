package service

import (
	"testing"

	"github.com/sinisaos/fiber-ent-admin/ent/enttest"
	"github.com/sinisaos/fiber-ent-admin/pkg/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestTagService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	tagService := NewTagService(client)
	userService := NewUserService(client)
	questionService := NewQuestionService(client)
	// Insert tag
	_, err := tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag1",
	})
	assert.NoError(t, err)

	_, err = tagService.CreateTag(&model.NewTagInput{
		Name: "TestTag2",
	})
	assert.NoError(t, err)

	// Check if pagination works
	pageOne, _ := tagService.GetAllTags(0, 1, "", "", "")
	assert.Len(t, pageOne, 1)
	pageTwo, _ := tagService.GetAllTags(1, 2, "", "", "")
	assert.Len(t, pageTwo, 1)
	allPages, _ := tagService.GetAllTags(0, 2, "", "", "")
	assert.Len(t, allPages, 2)

	// Check if sorting works
	sortDesc, _ := tagService.GetAllTags(0, 2, "id", "DESC", "")
	assert.Equal(t, sortDesc[0].ID, 2)
	sortAsc, _ := tagService.GetAllTags(0, 2, "id", "ASC", "")
	assert.Equal(t, sortAsc[0].ID, 1)

	// Check if filtering works
	filterName, _ := tagService.GetAllTags(0, 2, "", "", "1")
	assert.Equal(t, filterName[0].Name, "TestTag1")

	// Check empty data result
	emptyDataResult, _ := tagService.GetAllTags(10, 20, "", "", "")
	assert.Len(t, emptyDataResult, 0)

	// Records total for React admin pagination
	total, _ := tagService.CountTags()
	assert.Equal(t, total, 2)

	// Single Tag
	resultSingleTag, _ := tagService.GetTag(1)
	assert.Contains(t, resultSingleTag.Name, "TestTag1")

	// Return error if Tag does not exist
	_, err = tagService.GetTag(10)
	assert.Error(t, err)

	// Update Tag if exist
	_, err = tagService.UpdateTag(2, &model.UpdateTagInput{
		Name: "TestTag2Updated",
	})
	assert.NoError(t, err)

	// Check error if Tag does not exists
	_, err = tagService.UpdateTag(3, &model.UpdateTagInput{
		Name: "TestTag2Updated",
	})
	assert.Error(t, err)

	// Delete Tag
	err = tagService.DeleteTag(2)
	assert.NoError(t, err)

	// Check error if Tag does not exists
	err = tagService.DeleteTag(10)
	assert.Error(t, err)

	// Check if Tag is deleted
	deletedTagResult, _ := tagService.GetAllTags(0, 1, "", "", "")
	assert.Len(t, deletedTagResult, 1)

	// Insert user for checking tags tags
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

	// Insert question for checking tags questions
	_, err = questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion1",
		Author:  u.ID,
		Content: "Content of question one",
		Tags:    tags,
	})
	assert.NoError(t, err)

	// Checking tag questions
	tagQuestions, _ := tagService.GetTagQuestions(firstTag.ID)
	assert.Equal(t, tagQuestions.Edges.Questions[0].ID, 1)
	assert.Equal(t, tagQuestions.Edges.Questions[0].Title, "TestQuestion1")
	assert.Len(t, tagQuestions.Edges.Questions, 1)
}
