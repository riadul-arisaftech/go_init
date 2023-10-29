package services

import (
	"github.com/go_sample/api/interfaces"
	"github.com/go_sample/core/token"
	"github.com/go_sample/database/repository"
)

type AuthService struct {
	Store repository.Store
	Token *token.Maker
}

func NewAuthService(store repository.Store, maker *token.Maker) interfaces.IAuthService {
	return &AuthService{
		Store: store,
		Token: maker,
	}
}

func (s AuthService) CreateToken() error {
	return nil
}

func (s AuthService) CreateUser() error {
	err := s.Store.CreateUser()
	return err
}
