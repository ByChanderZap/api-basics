// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: orders.sql

package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (id, user_id, total, status, address, created_at, updated_at, deleted_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, user_id, total, status, address, created_at, updated_at, deleted_at
`

type CreateOrderParams struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"user_id"`
	Total     string       `json:"total"`
	Status    OrderStatus  `json:"status"`
	Address   string       `json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.UserID,
		arg.Total,
		arg.Status,
		arg.Address,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Total,
		&i.Status,
		&i.Address,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (id, order_id, product_id, quantity, price) 
VALUES ($1, $2, $3, $4, $5)
RETURNING id, order_id, product_id, quantity, price
`

type CreateOrderItemParams struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int32     `json:"quantity"`
	Price     string    `json:"price"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.ID,
		arg.OrderID,
		arg.ProductID,
		arg.Quantity,
		arg.Price,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.OrderID,
		&i.ProductID,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}
