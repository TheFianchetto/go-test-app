package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/TheFianchetto/go-test-app/db"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:@localhost:5432/go_test_db?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}

	// Example goroutine 1: Print a message every second
	go func() {
		for {
			fmt.Println("Goroutine 1: Running...")
			time.Sleep(1 * time.Second)
		}
	}()

	// Example goroutine 2: Print a different message every two seconds
	go func() {
		for {
			fmt.Println("Goroutine 2: Running...")
			time.Sleep(2 * time.Second)
		}
	}()

	// Basic web server for testing user actions
	// GET /users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		users, err := queries.ListUsers(ctx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Users: %v", users)
		defer func() {
			duration := time.Since(start)
			fmt.Printf("Served request in %v\n", duration)
		}()
	})

	// GET /user/:id
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/user/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		userID := int32(id)
		user, err := queries.GetUserByID(ctx, userID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "User: %v", user)
	})

	// POST /user
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")
		user, err := queries.CreateUser(ctx, db.CreateUserParams{
			Name:  name,
			Email: email,
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "User created: %v", user)
	})

	http.ListenAndServe(":8080", nil)

}
