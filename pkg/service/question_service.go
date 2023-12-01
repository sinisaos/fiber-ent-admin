package service

import (
	"context"
	"time"

	"github.com/sinisaos/fiber-ent-admin/ent"
	"github.com/sinisaos/fiber-ent-admin/ent/predicate"
	"github.com/sinisaos/fiber-ent-admin/ent/question"
	"github.com/sinisaos/fiber-ent-admin/pkg/model"
)

type QuestionService struct {
	Client *ent.Client
}

func NewQuestionService(client *ent.Client) *QuestionService {
	return &QuestionService{
		Client: client,
	}
}

func (s QuestionService) GetAllQuestions(start int, end int, sort string, order string, title string, content string) ([]*ent.Question, error) {
	var orderFunc []question.OrderOption
	var predicateQuestion []predicate.Question

	if title != "" {
		predicateQuestion = append(predicateQuestion, question.TitleContainsFold(title))
	}
	if content != "" {
		predicateQuestion = append(predicateQuestion, question.ContentContainsFold(content))
	}

	if order != "" {
		if order == "ASC" {
			orderFunc = append(orderFunc, ent.Asc(sort))
		} else {
			orderFunc = append(orderFunc, ent.Desc(sort))
		}
	}

	questions, _ := s.Client.Question.Query().
		Where(predicateQuestion...).
		WithAuthor().
		WithTags().
		Limit(end - start).
		Offset(start).
		Order(orderFunc...).
		All(context.Background())

	return questions, nil
}

func (s QuestionService) CountQuestions() (int, error) {
	count, _ := s.Client.Question.Query().
		Aggregate(
			ent.Count(),
		).
		Int(context.Background())
	return count, nil
}

func (s QuestionService) GetQuestion(id int) (*ent.Question, error) {
	question, err := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAuthor().
		WithAnswers().
		WithTags().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) CreateQuestion(payload *model.NewQuestionInput) (*ent.Question, error) {
	question, err := s.Client.Question.Create().
		SetTitle(payload.Title).
		SetContent(payload.Content).
		SetCreatedAt(time.Now()).
		SetAuthorID(payload.Author).
		AddTagIDs(payload.Tags...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) UpdateQuestion(id int, payload *model.UpdateQuestionInput) (*ent.Question, error) {
	existingQuestion, err := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithTags().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	var tagsSlice []int

	for i := range existingQuestion.Edges.Tags {
		tagsSlice = append(tagsSlice, existingQuestion.Edges.Tags[i].ID)
	}

	question, err := existingQuestion.Update().
		SetTitle(payload.Title).
		SetContent(payload.Content).
		SetAuthorID(payload.Author).
		RemoveTagIDs(tagsSlice...).
		AddTagIDs(payload.Tags...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) DeleteQuestion(id int) error {
	err := s.Client.Question.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s QuestionService) GetQuestionAnswers(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAnswers().
		Only(context.Background())

	return question, nil
}

func (s QuestionService) GetQuestionAuthor(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAuthor().
		Only(context.Background())

	return question, nil
}

func (s QuestionService) GetQuestionTags(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithTags().
		Only(context.Background())

	return question, nil
}
