-- name: GetProducts :many
SELECT id, name, description, image, price, quantity, created_at, updated_at
FROM products
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetProductById :one
SELECT id, name, description, image, price, quantity, created_at, updated_at
FROM products
WHERE id = $1
AND deleted_at IS NULL;

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

-- name: DeleteProduct :exec
UPDATE products
set deleted_at = $2,
    updated_at = $3
WHERE id = $1;

-- name: GetProductsByIds :many
SELECT id, name, description, image, price, quantity, created_at, updated_at
FROM products
WHERE id = ANY($1::uuid[])
AND deleted_at IS NULL;

-- name: DecreaseProductStock :one
UPDATE products
SET quantity = quantity - $2,
    updated_at = $3
WHERE id = $1
  AND deleted_at IS NULL
RETURNING id, name, description, image, price, quantity, created_at, updated_at;
