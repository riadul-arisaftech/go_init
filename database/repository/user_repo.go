package repository

import (
	"context"
	"github.com/go_sample/database/models"
)

func (store *GormStore) CreateUser() error {
	tx := store.DB().WithContext(context.Background())
	tx.Create(models.User{})
	tx.Commit()
	return nil
}
func (store *GormStore) ReadUser() error {
	return nil
}
func (store *GormStore) UpdateUser() error {
	return nil
}
func (store *GormStore) DeleteUser() error {
	return nil
}
