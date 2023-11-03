package service

import (
	"context"

	"github.com/sinisaos/fiber-ent-admin/ent"
	"github.com/sinisaos/fiber-ent-admin/ent/user"
	"github.com/sinisaos/fiber-ent-admin/model"
)

type AuthService struct {
	Client *ent.Client
}

func NewAuthService(client *ent.Client) *AuthService {
	return &AuthService{
		Client: client,
	}
}

func (s AuthService) Login(payload *model.LoginUserInput) (*ent.User, error) {
	user, err := s.Client.User.Query().
		Where(user.Username(payload.UserName)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}
