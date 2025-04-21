package types

import "github.com/google/uuid"

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}

type CartItem struct {
	ProductId uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,number"`
}
