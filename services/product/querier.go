// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package product

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	GetProducts(ctx context.Context) ([]Product, error)
}

var _ Querier = (*Queries)(nil)
