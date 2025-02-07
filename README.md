# Go Test Project

Trying to POC GoLang with various ORM's and migration tooling

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



## Helpful Articles:

- https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj


## Env

```
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://user:pass@localhost:5432/DB_NAME?sslmode=disable
GOOSE_MIGRATION_DIR=./db/migrations
```