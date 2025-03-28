package types

import "time"

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity" validate:"required,number"`
	Price       float64 `json:"price" validate:"required,number"`
}

type UpdateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image" validate:"omitempty,url"`
	Quantity    int     `json:"quantity" validate:"required,number"`
	Price       float64 `json:"price" validate:"required,number"`
}

type ProductResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       *string    `json:"image,omitempty"`
	Price       float64    `json:"price"`
	Quantity    int32      `json:"quantity"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
