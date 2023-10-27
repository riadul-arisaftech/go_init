package services

import (
	"context"
	"errors"
	"go_sample/api/interfaces"
	"go_sample/core/token"
	"go_sample/database"
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
