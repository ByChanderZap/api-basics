// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: products.sql

package product

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (id, name, description, image, price, quantity, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, description, image, price, quantity, created_at, updated_at, deleted_at
`

type CreateProductParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Price,
		arg.Quantity,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :one
UPDATE products
set deleted_at = $2,
    updated_at = $3
WHERE id = $1
RETURNING id, name, description, image, price, quantity, created_at, updated_at, deleted_at
`

type DeleteProductParams struct {
	ID        uuid.UUID        `json:"id"`
	DeletedAt pgtype.Timestamp `json:"deleted_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func (q *Queries) DeleteProduct(ctx context.Context, arg DeleteProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, deleteProduct, arg.ID, arg.DeletedAt, arg.UpdatedAt)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, description, image, price, quantity, created_at, updated_at, deleted_at
FROM products
WHERE id = $1
AND deleted_at IS NULL
`

func (q *Queries) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, name, description, image, price, quantity, created_at, updated_at, deleted_at 
FROM products
WHERE deleted_at IS NULL
ORDER BY created_at DESC
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Image,
			&i.Price,
			&i.Quantity,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $2,
    description = $3,
    image = $4,
    price = $5,
    quantity = $6,
    updated_at = $7
WHERE id = $1
RETURNING id, name, description, image, price, quantity, created_at, updated_at, deleted_at
`

type UpdateProductParams struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int32     `json:"quantity"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Image,
		arg.Price,
		arg.Quantity,
		arg.UpdatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Image,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
