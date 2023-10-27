package services

import (
	"context"
	"errors"
	"github.com/go_sample/api/interfaces"
	"github.com/go_sample/core/token"
	"github.com/go_sample/database"
)

type AuthService struct {
	Store database.Store
	Token *token.Maker
}

func NewAuthService(store database.Store, maker *token.Maker) interfaces.IAuthService {
	return &AuthService{
		Store: store,
		Token: maker,
	}
}

func (s AuthService) CreateToken() error {
	return nil
}

func (s AuthService) CreateUser() error {
	tx := s.Store.DB().WithContext(context.Background())
	if tx.Create("").Error != nil {
		tx.Rollback()
		return errors.New("error")
	}
	tx.Commit()
	return nil
}
