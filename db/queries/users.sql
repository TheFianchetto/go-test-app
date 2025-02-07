-- name: GetUserByID :one
SELECT id, name, email, age, created_at, updated_at, role FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, age, role) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUserByID :exec
UPDATE users SET name = $2, email = $3, age = $4, updated_at = now() WHERE id = $1 RETURNING updated_at;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, age, created_at, updated_at, role FROM users ORDER BY id;

-- name: GetUserByEmail :one
SELECT id, name, email, age, created_at, updated_at, role FROM users WHERE email = $1;
