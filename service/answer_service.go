package service

import (
	"context"
	"time"

	"github.com/sinisaos/fiber-ent-admin/ent"
	"github.com/sinisaos/fiber-ent-admin/ent/answer"
	"github.com/sinisaos/fiber-ent-admin/ent/predicate"
	"github.com/sinisaos/fiber-ent-admin/model"
)

type AnswerService struct {
	Client *ent.Client
}

func NewAnswerService(client *ent.Client) *AnswerService {
	return &AnswerService{
		Client: client,
	}
}

func (s AnswerService) GetAllAnswers(start int, end int, sort string, order string, content string) ([]*ent.Answer, error) {
	var orderFunc []answer.OrderOption
	var predicateAnswer []predicate.Answer

	if content != "" {
		predicateAnswer = append(predicateAnswer, answer.ContentContainsFold(content))
	}

	if order != "" {
		if order == "ASC" {
			orderFunc = append(orderFunc, ent.Asc(sort))
		} else {
			orderFunc = append(orderFunc, ent.Desc(sort))
		}
	}

	answers, _ := s.Client.Answer.Query().
		Where(predicateAnswer...).
		WithAuthor().
		WithQuestion().
		Limit(end - start).
		Offset(start).
		Order(orderFunc...).
		All(context.Background())

	return answers, nil
}

func (s AnswerService) CountAnswers() (int, error) {
	count, _ := s.Client.Answer.Query().
		Aggregate(
			ent.Count(),
		).
		Int(context.Background())
	return count, nil
}

func (s AnswerService) GetAnswer(id int) (*ent.Answer, error) {
	answer, err := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithAuthor().
		WithQuestion().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) CreateAnswer(payload *model.NewAnswerInput) (*ent.Answer, error) {
	answer, err := s.Client.Answer.Create().
		SetContent(payload.Content).
		SetCreatedAt(time.Now()).
		SetAuthorID(payload.Author).
		SetQuestionID(payload.Question).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) UpdateAnswer(id int, payload *model.UpdateAnswerInput) (*ent.Answer, error) {
	answer, err := s.Client.Answer.UpdateOneID(id).
		SetContent(payload.Content).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) DeleteAnswer(id int) error {
	err := s.Client.Answer.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s AnswerService) GetAnswerAuthor(id int) (*ent.Answer, error) {
	answer, _ := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithAuthor().
		Only(context.Background())

	return answer, nil
}

func (s AnswerService) GetAnswerQuestion(id int) (*ent.Answer, error) {
	answer, _ := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithQuestion().
		Only(context.Background())

	return answer, nil
}
