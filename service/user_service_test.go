package service

import (
	"strings"
	"testing"
	"time"

	"github.com/sinisaos/fiber-ent-admin/ent/enttest"
	"github.com/sinisaos/fiber-ent-admin/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	userService := NewUserService(client)
	questionService := NewQuestionService(client)
	answerService := NewAnswerService(client)
	// Insert user
	u, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	_, err = userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser2",
		Email:    "testuser2@gmail.com",
		Password: "pass1234",
	})
	assert.NoError(t, err)

	// Check duplicate email error
	_, err = userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser3",
		Email:    "testuser2@gmail.com",
		Password: "pass12345",
	})
	assert.Error(t, err)

	// Check if pagination works
	pageOne, _ := userService.GetAllUsers(0, 1, "", "", "", "", "")
	assert.Len(t, pageOne, 1)
	pageTwo, _ := userService.GetAllUsers(1, 2, "", "", "", "", "")
	assert.Len(t, pageTwo, 1)
	allPages, _ := userService.GetAllUsers(0, 2, "", "", "", "", "")
	assert.Len(t, allPages, 2)

	// Check if sorting works
	sortDesc, _ := userService.GetAllUsers(0, 2, "id", "DESC", "", "", "")
	assert.Equal(t, sortDesc[0].ID, 2)
	sortAsc, _ := userService.GetAllUsers(0, 2, "id", "ASC", "", "", "")
	assert.Equal(t, sortAsc[0].ID, 1)

	// Check if filtering works
	filterUserName, _ := userService.GetAllUsers(0, 2, "", "", "1", "", "")
	assert.Equal(t, filterUserName[0].Username, "TestUser1")
	filterEmail, _ := userService.GetAllUsers(0, 2, "", "", "1", "testuser1", "")
	assert.Equal(t, filterEmail[0].Email, "testuser1@gmail.com")
	delta := u.CreatedAt.Add(time.Hour * 24)
	dateString := strings.Split(delta.String(), " ")[0]
	filterCreatedAt, _ := userService.GetAllUsers(0, 2, "", "", "1", "testuser1", dateString)
	assert.Equal(t, filterCreatedAt[0].ID, 1)

	// Check empty data result
	emptyDataResult, _ := userService.GetAllUsers(10, 20, "", "", "", "", "")
	assert.Len(t, emptyDataResult, 0)

	// Records total for React admin pagination
	total, _ := userService.CountUsers()
	assert.Equal(t, total, 2)

	// Single user
	resultSingleUser, _ := userService.GetUser(1)
	assert.Contains(t, resultSingleUser.Username, "TestUser1")

	// Return error if user does not exist
	_, err = userService.GetUser(10)
	assert.Error(t, err)

	u.Superuser = true

	// Update user if exist
	if u.Superuser {
		_, err = userService.UpdateUser(2, &model.UpdateUserInput{
			UserName:  "TestUser2Updated",
			Superuser: true,
		})
		assert.NoError(t, err)
	} else {
		_, err = userService.UpdateUser(2, &model.UpdateUserInput{
			UserName:  "TestUser2Updated",
			Superuser: true,
		})
		assert.Error(t, err)
	}

	// Check error if user does not exists
	_, err = userService.UpdateUser(3, &model.UpdateUserInput{
		UserName: "TestUser2Updated",
	})
	assert.Error(t, err)

	// Check error if user does not exists
	_, err = userService.UpdateUser(3, &model.UpdateUserInput{
		UserName: "TestUser2Updated",
	})
	assert.Error(t, err)

	// Delete user
	err = userService.DeleteUser(2)
	assert.NoError(t, err)

	// Check error if user does not exists
	err = userService.DeleteUser(10)
	assert.Error(t, err)

	// Check if user is deleted
	deletedUserResult, _ := userService.GetAllUsers(0, 1, "", "", "", "", "")
	assert.Len(t, deletedUserResult, 1)

	// Insert question for checking user questions
	q, err := questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion1",
		Author:  u.ID,
		Content: "Content of question one",
	})
	assert.NoError(t, err)

	// Insert answer for checking user answers
	a, err := answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer one",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.Equal(t, a.ID, 1)
	assert.NoError(t, err)

	// Checking user questions and answers
	userQuestions, _ := userService.GetUserQuestions(u.ID)
	assert.Equal(t, userQuestions.Edges.Questions[0].ID, 1)
	assert.Equal(t, userQuestions.Edges.Questions[0].Title, "TestQuestion1")
	assert.Len(t, userQuestions.Edges.Questions, 1)

	userAnswers, _ := userService.GetUserAnswers(u.ID)
	assert.Equal(t, userAnswers.Edges.Answers[0].ID, 1)
	assert.Equal(t, userAnswers.Edges.Answers[0].Content, "Content of Answer one")
	assert.Len(t, userAnswers.Edges.Answers, 1)

}
