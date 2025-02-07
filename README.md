# Go Test Project

Trying to POC GoLang with various ORM's and migration tooling

Requirements:
- Go 1.23
- Goose
- Sqlc
- GORM

## Goals to evaluate

- Data Abstraction (ORM)
- DB Migrations 
- SQLc vs GORM
- DBO (Database Owners) can be done in SQLc
- go 1.23 webserver 


### Notes:

- Goose is really nice to work with like Rails for migrations
  - Like being able to rollback very easily
  - Raw SQL which is less fun
- GORM works pretty nice but yeah having performance impact sucks


## Commands

```bash
# Create a Migration Files
goose create -s migration_name sql

# Migrate All Changes 
goose up

# Roll Back Changes
goose down

# Generate New ORM Files / Queries
sqlc generate 

```

## Helpful Articles:

- https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj


## Env

```
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://user:pass@localhost:5432/DB_NAME?sslmode=disable
GOOSE_MIGRATION_DIR=./db/migrations
```


## Testing Requests

```bash
# List all Users
curl -X GET http://localhost:8080/users

# Get user by ID
curl -X GET http://localhost:8080/user/1

# Create User
curl -X POST http://localhost:8080/user -d "name=John Doe" -d "email=johndoe@example.com"

```