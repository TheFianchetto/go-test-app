package main

import (
	"fmt"
	"log"

	"github.com/TheFianchetto/go-test-app/db"
	"github.com/TheFianchetto/go-test-app/models"
)

func main() {
	// Connect to DB
	database := db.ConnectDB()

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
