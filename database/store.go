package database

import "gorm.io/gorm"

type GormStore struct {
	conn *gorm.DB
}

type Store interface {
	DB() *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &GormStore{
		conn: db,
	}
}

func (r GormStore) DB() *gorm.DB {
	return r.conn
}
