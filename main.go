package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	user, err := queries.GetUserByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Role, "User:", user.Name, user.Email)
}
