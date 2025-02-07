package migrations

import (
	"github.com/TheFianchetto/go-test-app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// CreateUsersTable creates the users table
func CreateUsersTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240206_create_users_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
