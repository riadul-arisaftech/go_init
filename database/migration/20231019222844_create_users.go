package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"go_sample/database/models"
	"gorm.io/gorm"
)

func CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20231019222844_create_users",
		Migrate: func(tx *gorm.DB) error {
			// Add migration logic here
			if err := tx.Table("users").AutoMigrate(&models.User{}); err != nil {
				return err
			}
			if err := tx.Exec(`CREATE UNIQUE INDEX users_email_uidx ON users (email);`).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Add rollback logic here
			return tx.Migrator().DropTable("users")
		},
	}
}
