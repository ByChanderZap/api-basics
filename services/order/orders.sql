-- name: CreateOrder :one
INSERT INTO orders (id, user_id, total, status, address, created_at, updated_at, deleted_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: CreateOrderItem :one
INSERT INTO order_items (id, order_id, product_id, quantity, price) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
