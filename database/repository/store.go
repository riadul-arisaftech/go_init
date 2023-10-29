package repository

import (
	"github.com/go_sample/database/migration"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormStore struct {
	conn *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(70)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// DB migrations
	if err := migration.Migrate(db); err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return &GormStore{
		conn: db,
	}
}

func (store GormStore) DB() *gorm.DB {
	return store.conn
}
