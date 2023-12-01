package service

import (
	"context"
	"time"

	"github.com/sinisaos/fiber-ent-admin/ent"
	"github.com/sinisaos/fiber-ent-admin/ent/predicate"
	"github.com/sinisaos/fiber-ent-admin/ent/user"
	"github.com/sinisaos/fiber-ent-admin/pkg/model"
	"github.com/sinisaos/fiber-ent-admin/pkg/utils"
)

type UserService struct {
	Client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		Client: client,
	}
}

func (s UserService) GetAllUsers(start int, end int, sort string, order string, title string, email string, createdAt string) ([]*ent.User, error) {
	var orderFunc []user.OrderOption
	var predicateUser []predicate.User

	if title != "" {
		predicateUser = append(predicateUser, user.UsernameContainsFold(title))
	}
	if email != "" {
		predicateUser = append(predicateUser, user.EmailContainsFold(email))
	}
	if createdAt != "" {
		createDate, _ := time.Parse("2006-01-02", createdAt)
		predicateUser = append(predicateUser, user.CreatedAtLTE(createDate))
	}

	if order != "" {
		if order == "ASC" {
			orderFunc = append(orderFunc, ent.Asc(sort))
		} else {
			orderFunc = append(orderFunc, ent.Desc(sort))
		}
	}

	users, _ := s.Client.User.Query().
		Where(predicateUser...).
		WithQuestions().
		WithAnswers().
		WithTags().
		Limit(end - start).
		Offset(start).
		Order(orderFunc...).
		All(context.Background())

	return users, nil
}

func (s UserService) CountUsers() (int, error) {
	count, _ := s.Client.User.Query().
		Aggregate(
			ent.Count(),
		).
		Int(context.Background())

	return count, nil
}

func (s UserService) GetUser(id int) (*ent.User, error) {
	user, err := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithAnswers().
		WithQuestions().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) CreateUser(payload *model.NewUserInput) (*ent.User, error) {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.Client.User.Create().
		SetEmail(payload.Email).
		SetUsername(payload.UserName).
		SetPassword(hashedPassword).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) UpdateUser(id int, payload *model.UpdateUserInput) (*ent.User, error) {
	hashedPassword, _ := utils.HashPassword(payload.Password)
	user, err := s.Client.User.UpdateOneID(id).
		SetUsername(payload.UserName).
		SetEmail(payload.Email).
		SetPassword(hashedPassword).
		SetSuperuser(payload.Superuser).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s UserService) DeleteUser(id int) error {
	err := s.Client.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) GetUserQuestions(id int) (*ent.User, error) {
	user, _ := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithQuestions().
		Only(context.Background())

	return user, nil
}

func (s UserService) GetUserAnswers(id int) (*ent.User, error) {
	user, _ := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithAnswers().
		Only(context.Background())

	return user, nil
}
