package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

// GetMigrations collects all migrations
func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		CreateUsersTable(),
	}
}
