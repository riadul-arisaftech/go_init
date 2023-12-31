package services

import (
	"github.com/go_sample/api/interfaces"
	"github.com/go_sample/core/token"
	"github.com/go_sample/core/workers"
	"github.com/go_sample/database/repository/irepo"
)

type AuthService struct {
	Store       irepo.IStore
	Token       *token.Maker
	Distributor workers.TaskDistributor
}

func NewAuthService(store irepo.IStore, maker *token.Maker, dist workers.TaskDistributor) interfaces.IAuthService {
	return &AuthService{
		Store:       store,
		Token:       maker,
		Distributor: dist,
	}
}

func (s AuthService) CreateToken() error {
	return nil
}

func (s AuthService) CreateUser() error {
	err := s.Store.CreateUser()
	return err
}
