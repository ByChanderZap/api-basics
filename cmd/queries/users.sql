-- name: CreateUser :one
INSERT INTO users (id, firstName, lastName, email, password, createdAt, updatedAt) 
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
