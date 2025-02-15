-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, email, password, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users 
WHERE email = $1;

-- name: GetUserById :one
SELECT *
FROM users 
WHERE id = $1;
