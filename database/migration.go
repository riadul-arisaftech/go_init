package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go_sample/database/migration"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	options := gormigrate.DefaultOptions
	options.UseTransaction = true
	m := gormigrate.New(db, options, []*gormigrate.Migration{
		migration.CreateUsersTable(),
	})

	if err := m.Migrate(); err != nil {
		return err
	}
	return nil
}
