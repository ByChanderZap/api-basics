package cart

import (
	"fmt"

	"github.com/ByChanderZap/api-basics/types"
	"github.com/google/uuid"
)

func getCartItemsIds(items []types.CartItem) ([]uuid.UUID, error) {
	productIds := make([]uuid.UUID, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product %v", item.ProductId)
		}

		productIds[i] = item.ProductId
	}
	return productIds, nil
}
