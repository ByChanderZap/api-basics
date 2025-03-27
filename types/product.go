package types

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity" validate:"required,number"`
	Price       float64 `json:"price" validate:"required,number"`
}
