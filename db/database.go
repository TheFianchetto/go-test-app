package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initializes the database connection and runs migrations
func ConnectDB() *gorm.DB {
	dbURL := "postgres://postgres:@localhost:5432/go_test_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
