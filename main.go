package main

import (
	"fmt"
	"log"

	"github.com/TheFianchetto/go-test-app/db"
	"github.com/TheFianchetto/go-test-app/db/migrations"
	"github.com/TheFianchetto/go-test-app/models"

	"github.com/go-gormigrate/gormigrate/v2"
)

func main() {
	// Connect to DB
	database := db.ConnectDB()

	// Run Migrations
	m := gormigrate.New(database, gormigrate.DefaultOptions, migrations.GetMigrations())
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migrations applied successfully!")

	// Check if the users table exists
	if !database.Migrator().HasTable(&models.User{}) {
		log.Fatalf("Users table does not exist")
	} else {
		fmt.Println("Users table exists")
	}

	// Reset the database
	// database.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	// print all users in db
	fmt.Print("All users in db: " + "\n" + "ID | Name | Email | Age" + "\n")
	var users []models.User
	database.Find(&users)
	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Email, user.Age)
	}

	// Create a new user (ActiveRecord-like)
	user := &models.User{Name: "Alice", Email: "alice@example.com", Age: 25}

	if err := user.Create(database); err != nil {
		log.Fatalf("User creation failed: %v", err)
	}

	fmt.Println("User created:", user)
}
