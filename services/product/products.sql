-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateProduct :one
INSERT INTO products (id, name, description, image, price, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;
