-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateProduct :one
INSERT INTO products (id, name, description, image, price, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name = $2,
    description = $3,
    image = $4,
    price = $5,
    quantity = $6,
    updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :one
UPDATE products
set deleted_at = $2,
    updated_at = $3
WHERE id = $1
RETURNING *;
