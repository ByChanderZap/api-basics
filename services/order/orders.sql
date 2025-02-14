-- name: CreateOrder :one
INSERT INTO orders (id, userId, total, status, address, createdAt, updatedAt, deletedAt) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: CreateOrderItem :one
INSERT INTO order_items (id, orderId, productId, quantity, price) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
