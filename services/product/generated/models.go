// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package productStore

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       *string    `json:"image"`
	Price       float64    `json:"price"`
	Quantity    int32      `json:"quantity"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
